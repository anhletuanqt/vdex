package matching

import (
	"time"

	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/orderbook"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Engine struct {
	productID     string
	OrderBook     *orderbook.OrderBook
	orderReader   OrderReader
	orderOffset   int64
	orderCh       chan *offsetOrder
	logCh         chan LogOrder
	snapshotReqCh chan *Snapshot
	snapshotCh    chan *Snapshot
	logStore      LogStore
	snapshotStore SnapshotStore
}

type offsetOrder struct {
	Offset int64
	Order  *models.Order
}

type Snapshot struct {
	OrderBookSnapshot *orderbook.OrderBookSnapshot
	OrderOffset       int64
}

func NewEngine(product *models.Product, orderReader OrderReader, logStore LogStore, snapshotStore SnapshotStore) *Engine {
	e := &Engine{
		productID:     product.ID,
		OrderBook:     orderbook.NewOrderBook(product.ID),
		orderCh:       make(chan *offsetOrder, 10000),
		logCh:         make(chan LogOrder, 10000),
		snapshotCh:    make(chan *Snapshot, 32),
		snapshotReqCh: make(chan *Snapshot, 32),
		orderReader:   orderReader,
		logStore:      logStore,
		snapshotStore: snapshotStore,
	}
	// restore
	snapshot, err := snapshotStore.GetLatest()
	if err != nil {
		logrus.Fatalf("get latest snapshot error: %v", err)
	}
	if snapshot != nil {
		e.restore(snapshot)
	}

	return e
}

func (e *Engine) Start() {
	go e.runFetcher()
	go e.runApplier()
	go e.runCommitter()
	go e.runSnapshots()
}

func (e *Engine) runFetcher() {
	var offset = e.orderOffset
	if offset > 0 {
		offset = offset + 1
	}
	err := e.orderReader.SetOffset(offset)
	if err != nil {
		logrus.Fatalf("[runFetcher] set order reader offset error: %v", err)
	}

	for {
		offset, order, err := e.orderReader.FetchOrder()
		if err != nil {
			logrus.Errorln(err)
			continue
		}
		e.orderCh <- &offsetOrder{offset, order}
	}
}

func (e *Engine) runApplier() {
	var orderOffset int64
	for {
		select {
		case offsetOrder := <-e.orderCh:
			var matchedOrder *MatchedOrder
			var err error
			if offsetOrder.Order.Status == models.OrderStatusCancelling {
				matchedOrder = e.cancelOrder(offsetOrder.Order)
			} else {
				matchedOrder, err = e.applyOrder(offsetOrder.Order)
			}
			if err != nil {
				logrus.Errorln(err)
				continue
			}
			// fmt.Println(e.OrderBook)
			logs := e.createLogOrderByMatchedOrder(offsetOrder.Order.ID, offsetOrder.Order.Type, matchedOrder)
			// request snapshot
			for _, log := range logs {
				if log != nil {
					e.logCh <- *log
				}
			}
			orderOffset = offsetOrder.Offset
		case snapshot := <-e.snapshotReqCh:
			delta := orderOffset - snapshot.OrderOffset
			// if delta <= 1000 {
			// 	continue
			// }
			if delta <= 0 {
				continue
			}
			snapshot.OrderBookSnapshot = e.OrderBook.Snapshot()
			snapshot.OrderOffset = orderOffset
			e.snapshotCh <- snapshot
		}
	}
}

func (e *Engine) runCommitter() {
	var logs []interface{}

	for {
		select {
		case log := <-e.logCh:
			logs = append(logs, log)
			// chan is not empty and buffer is not full, continue read.
			if len(e.logCh) > 0 && len(logs) < 100 {
				continue
			}
			// store log, clean buffer
			err := e.logStore.Store(logs)
			if err != nil {
				logrus.Errorln(err)
				panic(err)
			}
			logs = nil
		}
	}
}

func (e *Engine) runSnapshots() {
	orderOffset := e.orderOffset

	for {
		select {
		case <-time.After(10 * time.Second):
			// make a new snapshot request
			e.snapshotReqCh <- &Snapshot{
				OrderOffset: orderOffset,
			}

		case snapshot := <-e.snapshotCh:
			// store snapshot
			err := e.snapshotStore.Store(snapshot)
			if err != nil {
				logrus.Warnf("store snapshot failed: %v", err)
				continue
			}

			// update offset for next snapshot request
			orderOffset = snapshot.OrderOffset
		}
	}
}

func (e *Engine) applyOrder(order *models.Order) (*MatchedOrder, error) {
	if order.Type == models.OrderTypeLimit {
		done, partial, partialQuantityProcessed, err := e.OrderBook.ProcessLimitOrderV1(orderbook.SideFromStr(string(order.Side)), order.ID, order.Size, order.Price)
		matchedOrder := &MatchedOrder{
			done:                     done,
			partial:                  partial,
			partialQuantityProcessed: partialQuantityProcessed,
			quantityLeft:             decimal.NewFromInt(0),
		}
		if len(matchedOrder.done) == 0 && err == nil {
			matchedOrder.open = orderbook.NewOrder(order.ID, orderbook.Side(order.Side), order.Size, order.Price, time.Now())
		}

		return matchedOrder, err
	}

	done, partial, partialQuantityProcessed, quantityLeft, err := e.OrderBook.ProcessMarketOrderV1(order.ID, orderbook.Side(order.Side), order.Size)
	return &MatchedOrder{
		done:                     done,
		partial:                  partial,
		partialQuantityProcessed: partialQuantityProcessed,
		quantityLeft:             quantityLeft,
	}, err
}

func (e *Engine) cancelOrder(order *models.Order) *MatchedOrder {
	cancelOrder := e.OrderBook.CancelOrder(order.ID)

	return &MatchedOrder{
		cancel: cancelOrder,
	}
}

func (e *Engine) createLogOrderByMatchedOrder(takerOrderID int64, orderType models.OrderType, matchedOrder *MatchedOrder) []*LogOrder {
	logs := []*LogOrder{}
	if matchedOrder.cancel != nil {
		order := matchedOrder.cancel
		logs = append(logs, &LogOrder{
			LogType:   LogTypeCancel,
			OrderID:   order.ID(),
			Quantity:  order.Quantity(),
			Price:     order.Price(),
			Side:      order.Side(),
			Time:      order.Time(),
			ProductID: e.productID,
		})
	}

	if matchedOrder.open != nil {
		order := matchedOrder.open
		logs = append(logs, &LogOrder{
			LogType:   LogTypeOpen,
			OrderID:   order.ID(),
			Quantity:  order.Quantity(),
			Price:     order.Price(),
			Side:      order.Side(),
			Time:      order.Time(),
			ProductID: e.productID,
		})
	}

	if len(matchedOrder.done) != 0 {
		for _, v := range matchedOrder.done {
			logs = append(logs, &LogOrder{
				LogType:      LogTypeDone,
				TakerOrderID: v.TakerOrderID,
				MakerOrderID: v.MakerOrderID,
				Side:         v.Side,
				Time:         v.Time,
				ProductID:    e.productID,
				Price:        v.Price,
				Quantity:     v.Quantity,
			})
		}
	}

	return logs
}

func (e *Engine) restore(snapshot *Snapshot) {
	logrus.Infof("restoring: %+v", *snapshot)
	e.orderOffset = snapshot.OrderOffset
	e.OrderBook.Restore(snapshot.OrderBookSnapshot)
}
