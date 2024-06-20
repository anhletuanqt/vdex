package orderbook

type OrderBookSnapshot struct {
	Bids []*Order
	Asks []*Order
}

func (o *OrderBook) Snapshot() *OrderBookSnapshot {
	snapshot := &OrderBookSnapshot{
		Asks: make([]*Order, o.asks.numOrders),
		Bids: make([]*Order, o.bids.numOrders),
	}

	for i, order := range o.asks.Orders() {
		o, _ := order.Value.(*Order)
		snapshot.Asks[i] = o
	}
	for i, order := range o.bids.Orders() {
		o, _ := order.Value.(*Order)
		snapshot.Bids[i] = o
	}

	return snapshot
}

func (o *OrderBook) Restore(snapshot *OrderBookSnapshot) {
	for _, order := range snapshot.Asks {
		o.orders[order.id] = o.asks.Append(order)
	}
	for i, order := range snapshot.Bids {
		o.orders[order.id] = o.bids.Append(snapshot.Bids[i])
	}
}
