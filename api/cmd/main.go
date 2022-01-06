package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/cryptx/api/handler"
	"github.com/yimsoijoi/cryptx/config"
	"github.com/yimsoijoi/cryptx/datamodel"
	"github.com/yimsoijoi/cryptx/lib/postgres"
)

func main() {
	conf, err := config.Load()
	db, err := postgres.New(conf.Postgres)
	if err != nil {
		log.Fatalln("can't connect db")
	}
	fmt.Println("db connected")
	db.AutoMigrate(&datamodel.Wallet{}, &datamodel.Order{}, &datamodel.Token{})
	// check(db)

	handler := handler.New(db, conf.WalletPrivateKey)
	app := fiber.New()
	orderAPI := app.Group("/orders")
	orderAPI.Post("/", handler.CreateOrder)
	orderAPI.Post("/pay/:uuid", handler.Pay)
	orderAPI.Get("/", handler.GetOrders)
	orderAPI.Get("/:uuid", handler.GetOrder)

	log.Fatal(app.Listen(":8000"))
}
