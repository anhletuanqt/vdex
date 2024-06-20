package service

import (
	"context"

	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/models/pg"
	"github.com/cxptek/vdex/util"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func UpdateSettlementFee(ctx context.Context, settlement *models.Settlement) error {
	db, err := pg.SharedStore().BeginTx(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = db.Rollback() }()
	if settlement.Type == models.SettlementOrder {
		// add settlement fee
		order, err := pg.SharedStore().GetOrderByID(ctx, settlement.RefID)
		if err != nil {
			logrus.Errorln(err)
			return err
		}
		// find segment
		tickTime := settlement.Time.UTC().Truncate(util.UserSegmentTime).Unix()
		segment, err := pg.SharedStore().GetUserSegmentByUserIDAndTime(ctx, order.UserID, tickTime)
		if err != nil {
			logrus.Errorln(err)
			return err
		}
		makerFee, _ := util.GetSegmentFeeByType(segment.Type)
		// buy fee
		fee := order.Size.Sub(order.FilledSize).Mul(order.Price).Mul(makerFee).Div(decimal.NewFromInt(100))
		if order.Side == models.SideSell {
			fee = order.Size.Sub(order.FilledSize).Mul(makerFee).Div(decimal.NewFromInt(100))
		}
		settlement.MakerFee = fee
		if err := db.UpdateOrderFeeByID(ctx, order.ID, fee); err != nil {
			return err
		}
	} else if settlement.Type == models.SettlementTrade {
		trade, err := pg.SharedStore().GetTradeByID(ctx, settlement.RefID)
		if err != nil {
			logrus.Errorln(err)
			return err
		}
		orderIDs := []int64{trade.TakerOrderID, trade.MakerOrderID}
		orders, err := pg.SharedStore().GetOrdersByIDs(ctx, orderIDs)
		if err != nil {
			logrus.Errorln(err)
			return err
		}
		// find segment
		takerOrder, _ := lo.Find[models.Order](orders, func(item models.Order) bool {
			return item.ID == trade.TakerOrderID
		})
		makerOrder, _ := lo.Find[models.Order](orders, func(item models.Order) bool {
			return item.ID == trade.MakerOrderID
		})
		tickTime := settlement.Time.UTC().Truncate(util.UserSegmentTime).Unix()
		takerSegment, err := pg.SharedStore().GetUserSegmentByUserIDAndTime(ctx, takerOrder.UserID, tickTime)
		if err != nil {
			logrus.Errorln(err)
			return err
		}
		makerSegment, err := pg.SharedStore().GetUserSegmentByUserIDAndTime(ctx, makerOrder.UserID, tickTime)
		if err != nil {
			logrus.Errorln(err)
			return err
		}
		// Taker fee
		_, takerFee := util.GetSegmentFeeByType(takerSegment.Type)
		// buy fee
		tFee := trade.Size.Mul(takerFee).Div(decimal.NewFromInt(100))
		if takerOrder.Side == models.SideSell {
			tFee = trade.Size.Mul(trade.Price).Mul(takerFee).Div(decimal.NewFromInt(100))
		}
		// Maker fee
		makerFee, _ := util.GetSegmentFeeByType(makerSegment.Type)
		// buy fee
		mFee := trade.Size.Mul(makerFee).Div(decimal.NewFromInt(100))
		if makerOrder.Side == models.SideSell {
			mFee = trade.Size.Mul(trade.Price).Mul(makerFee).Div(decimal.NewFromInt(100))
		}
		settlement.TakerFee = tFee
		settlement.MakerFee = mFee
		// update order fee
		if err := db.UpdateOrderFeeByID(ctx, takerOrder.ID, tFee); err != nil {
			return err
		}
		if err := db.UpdateOrderFeeByID(ctx, makerOrder.ID, mFee); err != nil {
			return err
		}
	}

	if err := db.UpdateSettlement(ctx, settlement); err != nil {
		logrus.Errorln(err)
		return err
	}

	return db.CommitTx()
}
