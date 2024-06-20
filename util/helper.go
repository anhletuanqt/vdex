package util

import (
	"github.com/cxptek/vdex/models"
	"github.com/shopspring/decimal"
)

var (
	SegmentVip0 = 0
	SegmentVip1 = 1
	SegmentVip2 = 2
	SegmentVip3 = 3
	SegmentVip4 = 4

	N_10M  = decimal.NewFromInt(10000000)
	N_200M = decimal.NewFromInt(200000000)
	N_5B   = decimal.NewFromInt(5000000000)
	N_100B = decimal.NewFromInt(100000000000)
)

func GetSegmentTypeByVolume(volume decimal.Decimal) int64 {
	segmentType := SegmentVip0

	if volume.GreaterThanOrEqual(N_10M) && volume.LessThan(N_200M) {
		segmentType = SegmentVip1
	} else if volume.GreaterThanOrEqual(N_200M) && volume.LessThan(N_5B) {
		segmentType = SegmentVip2
	} else if volume.GreaterThanOrEqual(N_5B) && volume.LessThan(N_100B) {
		segmentType = SegmentVip3
	} else if volume.GreaterThanOrEqual(N_100B) {
		segmentType = SegmentVip4
	}

	return int64(segmentType)
}

func GetSegmentFeeByType(segmentType int64) (makerFee decimal.Decimal, takerFee decimal.Decimal) {
	switch int(segmentType) {
	case SegmentVip0:
		makerFee = decimal.NewFromFloat(0.1)
		takerFee = decimal.NewFromFloat(0.2)
	case SegmentVip1:
		makerFee = decimal.NewFromFloat(0.09)
		takerFee = decimal.NewFromFloat(0.18)
	case SegmentVip2:
		makerFee = decimal.NewFromFloat(0.07)
		takerFee = decimal.NewFromFloat(0.14)
	case SegmentVip3:
		makerFee = decimal.NewFromFloat(0.05)
		takerFee = decimal.NewFromFloat(0.1)
	case SegmentVip4:
		makerFee = decimal.NewFromFloat(0.03)
		takerFee = decimal.NewFromFloat(0.06)
	}

	return makerFee, takerFee
}

func GetUserByID(id int64, users []models.User) models.User {
	for _, u := range users {
		if u.ID == id {
			return u
		}
	}
	return models.User{}
}

func OrderTypeToNumber(orderType models.OrderType) uint8 {
	// 0: market, 1: limit
	if orderType == models.OrderTypeLimit {
		return 1
	}

	return 0
}

func OrderSideToNumber(orderSide models.Side) uint8 {
	// 0: buy, 1: sell
	if orderSide == models.SideSell {
		return 1
	}

	return 0
}
