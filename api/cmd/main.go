package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/cryptx/api/handler"
	"github.com/yimsoijoi/cryptx/datamodel"
	"github.com/yimsoijoi/cryptx/lib/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := postgres.New()
	if err != nil {
		log.Fatalln("can't connect db")
	}
	fmt.Println("db connected")
	db.AutoMigrate(&datamodel.Wallet{}, &datamodel.Order{})

	handler := handler.New(db)
	app := fiber.New()
	orderAPI := app.Group("/orders")
	orderAPI.Post("/", handler.CreateOrder)
	orderAPI.Post("/pay/:uuid", handler.Pay)
	log.Fatal(app.Listen(":8000"))
}

func check(db *gorm.DB) {
	a := datamodel.NewWallet("6969")
	b := datamodel.NewWallet("11210")
	o := datamodel.NewOrder(*a, *b, 27)

	_a, _ := json.MarshalIndent(a, "  ", "  ")
	_b, _ := json.MarshalIndent(b, "  ", "  ")
	_o, _ := json.MarshalIndent(o, "  ", "  ")
	fmt.Printf("{\n  \"a\": %s,\n  \"b\": %s,\n  \"o\": %s,\n}\n", _a, _b, _o)

	db.Create(&a)
	db.Create(&b)
	db.Create(&o)

	var send datamodel.Wallet
	var recv datamodel.Wallet
	var order datamodel.Order

	db.Preload("From").Preload("To").First(&order)
	send = order.From
	recv = order.To

	_send, _ := json.MarshalIndent(send, "  ", "  ")
	_recv, _ := json.MarshalIndent(recv, "  ", "  ")
	_order, _ := json.MarshalIndent(order, "  ", "  ")
	fmt.Printf("{\n  \"send\": %s,\n  \"recv\": %s,\n  \"order\": %s,\n}\n", _send, _recv, _order)
}
