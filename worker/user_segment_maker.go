package worker

import (
	"context"
	"database/sql"
	"time"

	"github.com/cxptek/vdex/matching"
	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/models/pg"
	"github.com/cxptek/vdex/service"
	"github.com/cxptek/vdex/util"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type UserSegmentLog struct {
	OrderID      int64
	TakerOrderID int64
	MakerOrderID int64
	LogType      matching.LogType
	Size         decimal.Decimal
	Price        decimal.Decimal
	Time         time.Time
	LogOffset    int64
}

type UserSegmentMaker struct {
	segmentCh chan *UserSegmentLog
	logReader matching.LogReader
	logOffset int64
}

func NewUserSegmentMaker(logReader matching.LogReader) *UserSegmentMaker {
	ctx := context.Background()

	t := &UserSegmentMaker{
		segmentCh: make(chan *UserSegmentLog, 1000),
		logReader: logReader,
	}

	lastSegment, err := pg.SharedStore().GetLastUserSegment(ctx)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if lastSegment != nil {
		t.logOffset = lastSegment.LogOffset
	}
	t.logReader.RegisterObserver(t)

	return t
}

func (t *UserSegmentMaker) Start() {
	if t.logOffset > 0 {
		t.logOffset++
	}

	go t.logReader.Run(t.logOffset)
	go t.flusher()
}

func (t *UserSegmentMaker) OnOpenLog(log *matching.LogOrder, offset int64) {
	// do nothing
}

func (t *UserSegmentMaker) OnCancelLog(log *matching.LogOrder, offset int64) {
	t.segmentCh <- &UserSegmentLog{
		LogType:   log.LogType,
		OrderID:   log.OrderID,
		Size:      log.Quantity,
		Price:     log.Price,
		Time:      log.Time,
		LogOffset: offset,
	}
}

func (t *UserSegmentMaker) OnDoneLog(log *matching.LogOrder, offset int64) {
	t.segmentCh <- &UserSegmentLog{
		LogType:      log.LogType,
		OrderID:      log.OrderID,
		TakerOrderID: log.TakerOrderID,
		MakerOrderID: log.MakerOrderID,
		Size:         log.Quantity,
		Price:        log.Price,
		Time:         log.Time,
		LogOffset:    offset,
	}
}

func (t *UserSegmentMaker) flusher() {
	for {
		select {
		case segmentLog := <-t.segmentCh:
			executeSegment(segmentLog)
		}
	}
}

func executeSegment(segmentLog *UserSegmentLog) error {
	ctx := context.Background()

	tickTime := segmentLog.Time.UTC().Truncate(util.UserSegmentTime).Unix()
	// tx
	db, err := pg.SharedStore().BeginTx(ctx)
	if err != nil {
		logrus.Error(err, "OrderID: ", segmentLog.OrderID)
		return err
	}
	defer func() { _ = db.Rollback() }()

	if segmentLog.LogType == matching.LogTypeCancel {
		order, err := service.GetOrderByID(ctx, segmentLog.OrderID)
		if err != nil {
			logrus.Error(err, "OrderID: ", segmentLog.OrderID)
			return err
		}
		newSegment := &models.UserSegment{
			UserID: order.UserID,
			Volume: decimal.Zero,
			Time:   tickTime,
		}
		segment, err := db.GetLastUserSegmentByUserIDForUpdate(ctx, order.UserID)
		volume := segmentLog.Price.Mul(segmentLog.Size)
		if err != nil {
			if err != sql.ErrNoRows {
				logrus.Error(err)
				return err
			}
			segment = &models.UserSegment{
				Volume: decimal.Zero,
			}
		}
		// segment
		newSegment.Volume = segment.Volume.Add(volume)
		newSegment.Type = util.GetSegmentTypeByVolume(newSegment.Volume)
		newSegment.LogOffset = segmentLog.LogOffset
		newSegment.Time = tickTime
		if _, err := db.AddUserSegment(ctx, newSegment); err != nil {
			logrus.Error(err)
			return err
		}
	} else if segmentLog.LogType == matching.LogTypeDone {
		orderIDs := []int64{segmentLog.TakerOrderID, segmentLog.MakerOrderID}
		orders, err := service.GetOrdersByIDs(ctx, orderIDs)
		if err != nil {
			logrus.Error(err, "OrderIDs: ", orderIDs)
			return err
		}

		// segment0
		newSegment0 := &models.UserSegment{
			UserID: orders[0].UserID,
			Volume: decimal.Zero,
			Time:   tickTime,
		}
		volume := segmentLog.Price.Mul(segmentLog.Size)
		segment0, err := db.GetLastUserSegmentByUserIDForUpdate(ctx, orders[0].UserID)
		if err != nil {
			if err != sql.ErrNoRows {
				logrus.Error(err)
				return err
			}
			segment0 = &models.UserSegment{
				Volume: decimal.Zero,
			}
		}
		newSegment0.Volume = segment0.Volume.Add(volume)
		newSegment0.Type = util.GetSegmentTypeByVolume(newSegment0.Volume)
		newSegment0.LogOffset = segmentLog.LogOffset
		newSegment0.Time = tickTime
		if _, err := db.AddUserSegment(ctx, newSegment0); err != nil {
			logrus.Error(err)
			return err
		}

		// segment1
		segment1, err := db.GetLastUserSegmentByUserIDForUpdate(ctx, orders[1].UserID)
		if err != nil {
			if err != sql.ErrNoRows {
				logrus.Error(err)
				return err
			}
			segment1 = &models.UserSegment{
				Volume: decimal.Zero,
			}
		}
		newSegment1 := &models.UserSegment{
			UserID: orders[1].UserID,
			Volume: decimal.Zero,
			Time:   tickTime,
		}
		newSegment1.Volume = segment1.Volume.Add(volume)
		newSegment1.Type = util.GetSegmentTypeByVolume(newSegment1.Volume)
		newSegment1.LogOffset = segmentLog.LogOffset
		newSegment1.Time = tickTime
		if _, err := db.AddUserSegment(ctx, newSegment1); err != nil {
			logrus.Error(err)
			return err
		}
	}

	return db.CommitTx()
}
