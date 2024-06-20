package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"github.com/cxptek/vdex/config"
	"github.com/cxptek/vdex/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var bunDB *bun.DB
var store models.Store
var storeOnce sync.Once

type Store struct {
	db bun.IDB
}

func SharedStore() models.Store {
	storeOnce.Do(func() {
		err := initDb()
		if err != nil {
			panic(fmt.Sprintf("connect db err: %v", err))
		}
		store = NewStore(bunDB)
	})
	return store
}

func NewStore(db bun.IDB) *Store {
	return &Store{
		db: db,
	}
}

func initDb() error {
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.GetConfig().DbURL)))
	bunDB = bun.NewDB(sqlDB, pgdialect.New())
	if err := bunDB.Ping(); err != nil {
		return err
	}
	bunDB.SetMaxIdleConns(2)
	bunDB.SetMaxOpenConns(50)
	// bunDB.AddQueryHook(bundebug.NewQueryHook())

	return nil
}

func (s *Store) BeginTx(ctx context.Context) (models.Store, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return NewStore(tx), nil
}

func (s *Store) Rollback() error {
	tx, ok := s.db.(bun.Tx)
	if !ok {
		return errors.New("Store not tx type")
	}

	return tx.Rollback()
}

func (s *Store) CommitTx() error {
	tx, ok := s.db.(bun.Tx)
	if !ok {
		return errors.New("Store not tx type")
	}

	return tx.Commit()
}
