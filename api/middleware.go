package api

import (
	"net/http"
	"strings"

	"github.com/cxptek/vdex/service"
	"github.com/cxptek/vdex/util"
	"github.com/gofiber/fiber/v2"
)

type Header struct {
	Authorization string `reqHeader:"Authorization"`
}

func requireJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		h := new(Header)
		if err := c.ReqHeaderParser(h); err != nil {
			return c.Status(http.StatusForbidden).JSON(util.ErrorResponse(http.StatusForbidden, err.Error()))
		}
		bearerToken := strings.Split(h.Authorization, " ")
		if len(bearerToken) != 2 {
			return c.Status(http.StatusForbidden).JSON(util.ErrorResponse(http.StatusForbidden, "invalid token"))
		}
		token := bearerToken[1]
		user, err := service.CheckJWT(token)
		if err != nil {
			return c.Status(http.StatusForbidden).JSON(util.ErrorResponse(http.StatusForbidden, err.Error()))
		}
		c.Locals("user", user)
		c.Next()

		return nil
	}
}
