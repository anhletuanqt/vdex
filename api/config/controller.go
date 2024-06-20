package config

import (
	"net/http"

	"github.com/cxptek/vdex/config"
	"github.com/cxptek/vdex/util"
	"github.com/gofiber/fiber/v2"
)

func Addresses(c *fiber.Ctx) error {
	params := &IDParams{}
	if err := c.ParamsParser(params); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}

	resp := fiber.Map{
		"vdex": config.GetConfig().Contracts.Vdex,
		"tokens": fiber.Map{
			"vic": config.GetAddressBySymbol("VIC"),
			"bnb": config.GetAddressBySymbol("BNB"),
		},
	}

	return c.Status(http.StatusOK).JSON(util.SuccessResponse(resp))
}
