package bootstrap

import (
	"database/sql"
	"wallet-api/app/routes"

	"github.com/gofiber/fiber/v2"
)

type bootstrap struct {
	c   *fiber.App
	mysqlDB *sql.DB
}

func NewBootstrap(c *fiber.App, mysqlDB *sql.DB) *bootstrap {
	return &bootstrap{
		c:       c,
		mysqlDB: mysqlDB,
	}
}

func (b *bootstrap) Run() error {
	//init route
	routes.NewRouter(b.c, b.mysqlDB).Init()

	b.c.Listen(":3000")
	return nil
}