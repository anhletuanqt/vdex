package user

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/cxptek/vdex/models"
	"github.com/cxptek/vdex/service"
	"github.com/cxptek/vdex/util"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	req := &LoginReq{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}

	// verify signature
	if err := verifyLoginEIP712(req.Address, req.Signature); err != nil {
		return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
	}

	user, err := service.GetUserByAddress(c.Context(), strings.ToLower(req.Address))
	if err != nil {
		if err != sql.ErrNoRows {
			return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
		}

		// insert new user
		user = &models.User{
			Address: strings.ToLower(req.Address),
		}
		if err := service.AddUser(c.Context(), user); err != nil {
			return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
		}
	}

	// sign JWT
	accessToken, err := signJWT(user)
	if err != nil {
		return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(util.SuccessResponse(fiber.Map{
		"jwt": accessToken,
	}))
}

func GetByID(c *fiber.Ctx) error {
	params := &IDParams{}
	if err := c.ParamsParser(params); err != nil {
		return c.Status(200).JSON(util.ErrorResponse(400, err.Error()))
	}

	user, err := service.GetUserByID(c.Context(), params.ID)
	if err != nil {
		return c.Status(http.StatusOK).JSON(util.ErrorResponse(400, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(util.SuccessResponse(user))
}
