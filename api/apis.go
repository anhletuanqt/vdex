package api

import (
	"github.com/cxptek/vdex/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"

	configAPI "github.com/cxptek/vdex/api/config"
	orderAPI "github.com/cxptek/vdex/api/order"
	productAPI "github.com/cxptek/vdex/api/product"
	userAPI "github.com/cxptek/vdex/api/user"
)

func StartServer() {
	engine := fiber.New()
	engine.Use(cors.New())
	engine.Use(logger.New())
	engine.Use(recover.New())

	registerAPIs(engine)

	logrus.Fatal(engine.Listen(config.GetConfig().RestServer.Port))
}

func registerAPIs(app *fiber.App) {
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	productAPIs(app)
	userAPIs(app)
	orderAPIs(app)
	configAPIs(app)
}

func productAPIs(app *fiber.App) {
	group := app.Group("/v1/products")

	group.Get("/", productAPI.GetProducts)
	group.Get("/:id", productAPI.GetProductByID)
	group.Get("/:id/depths", productAPI.GetDepths)
	group.Get("/:id/trades", productAPI.GetTrades)
	group.Get("/:id/candles", productAPI.GetCandles)
}

func userAPIs(app *fiber.App) {
	group := app.Group("/v1/users")

	group.Post("/login", userAPI.Login)
	group.Get("/:id", requireJWT(), userAPI.GetByID)
}

func orderAPIs(app *fiber.App) {
	group := app.Group("/v1/orders")

	if config.GetConfig().Env == "local" {
		group.Post("/", requireJWT(), orderAPI.CreateOrder)
		group.Post("/all", requireJWT(), orderAPI.CreateOrder1)
	}
	group.Get("/", requireJWT(), orderAPI.GetOrders)
	group.Put("/:id", requireJWT(), orderAPI.CancelOrder)
	group.Post("/gasless", requireJWT(), orderAPI.CreateGaslessOrder)
}

func configAPIs(app *fiber.App) {
	group := app.Group("/v1/configs")

	group.Get("/addrs", configAPI.Addresses)
}
