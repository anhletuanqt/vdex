package matching

import "github.com/cxptek/vdex/models"

type OrderReader interface {
	SetOffset(offset int64) error
	FetchOrder() (offset int64, order *models.Order, err error)
}

type LogStore interface {
	Store(logs []interface{}) error
}

type LogReader interface {
	GetProductId() string
	RegisterObserver(observer LogObserver)
	Run(offset int64)
}

type LogObserver interface {
	OnOpenLog(log *LogOrder, offset int64)
	OnDoneLog(log *LogOrder, offset int64)
	OnCancelLog(log *LogOrder, offset int64)
}

type SnapshotStore interface {
	Store(snapshot *Snapshot) error
	GetLatest() (*Snapshot, error)
}
