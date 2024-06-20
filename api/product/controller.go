package product

import (
	"net/http"

	"github.com/cxptek/vdex/service"
	"github.com/cxptek/vdex/util"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	products, err := service.GetProducts(c.Context())
	if err != nil {
		customErr, ok := err.(util.CustomErr)
		if ok {
			return c.Status(customErr.Code).JSON(util.ErrorResponse(customErr.Code, customErr.Error()))
		}

		return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(util.SuccessResponse(products))
}

func GetProductByID(c *fiber.Ctx) error {
	params := &IDParams{}
	if err := c.ParamsParser(params); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	products, err := service.GetProductByID(c.Context(), params.ID)
	if err != nil {
		customErr, ok := err.(util.CustomErr)
		if ok {
			return c.Status(customErr.Code).JSON(util.ErrorResponse(customErr.Code, customErr.Error()))
		}

		return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(util.SuccessResponse(products))
}

func GetTrades(c *fiber.Ctx) error {
	query := &GetTradesQuery{}
	if err := c.QueryParser(query); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	params := IDParams{}
	if err := c.ParamsParser(&params); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	products, err := service.GetTradesByProductID(c.Context(), params.ID, query.Limit)
	if err != nil {
		customErr, ok := err.(util.CustomErr)
		if ok {
			return c.Status(customErr.Code).JSON(util.ErrorResponse(customErr.Code, customErr.Error()))
		}

		return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(util.SuccessResponse(products))
}

func GetDepths(c *fiber.Ctx) error {
	query := &GetDepthQuery{}
	if err := c.QueryParser(query); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	params := &IDParams{}
	if err := c.ParamsParser(params); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	limit := query.Limit
	if limit > 1000 {
		limit = 1000
	}
	asks, bids := service.GetProductDepth(params.ID, query.Limit)
	return c.Status(http.StatusOK).JSON(util.SuccessResponse(fiber.Map{
		"asks": asks,
		"bids": bids,
	}))
}

func GetCandles(c *fiber.Ctx) error {
	query := &GetCandlesQuery{}
	if err := c.QueryParser(query); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	params := &IDParams{}
	if err := c.ParamsParser(params); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}
	if query.Granularity == 0 {
		query.Granularity = 60
	}

	tickRes := [][6]float64{}

	ticks, err := service.GetTicksByProductID(c.Context(), params.ID, query.Granularity, query.Limit)
	if err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}

	for _, tick := range ticks {
		tickRes = append(tickRes, [6]float64{float64(tick.Time), util.DToF64(tick.Low), util.DToF64(tick.High),
			util.DToF64(tick.Open), util.DToF64(tick.Close), util.DToF64(tick.Volume)})
	}

	return c.Status(http.StatusOK).JSON(util.SuccessResponse(tickRes))
}
