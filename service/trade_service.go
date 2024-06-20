package service

import (
	"context"

	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/models/pg"
	"github.com/sirupsen/logrus"
)

func AddTrades(ctx context.Context, trades []*models.Trade) error {
	if len(trades) == 0 {
		return nil
	}
	return pg.SharedStore().AddTrades(ctx, trades)
}

func GetTradesByProductID(ctx context.Context, productID string, limit int64) ([]models.Trade, error) {
	return pg.SharedStore().GetTradesByProductID(ctx, productID, limit)
}

func GetTradeByID(ctx context.Context, tradeID int64) (*models.Trade, error) {
	return pg.SharedStore().GetTradeByID(ctx, tradeID)
}

func ExecuteTrades(ctx context.Context, trades []*models.Trade) error {
	db, err := pg.SharedStore().BeginTx(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = db.Rollback() }()

	if err := db.AddTrades(ctx, trades); err != nil {
		logrus.Errorln(err)
		return err
	}
	settlements := make([]models.Settlement, len(trades))
	for i := range trades {
		settlements[i] = models.Settlement{
			RefID:   trades[i].ID,
			Settled: false,
			Type:    models.SettlementTrade,
			Time:    *trades[i].Time,
		}
	}
	if err := db.AddSettlements(ctx, settlements); err != nil {
		logrus.Errorln(err)
		return err
	}

	return db.CommitTx()
}
