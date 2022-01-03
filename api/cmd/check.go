package main

import (
	"encoding/json"
	"fmt"

	"github.com/yimsoijoi/cryptx/datamodel"
	"gorm.io/gorm"
)

func check(db *gorm.DB) {
	metaheroToken := datamodel.NewToken("0xD40bEDb44C081D2935eebA6eF5a3c8A31A1bBE13", 18)
	a := datamodel.NewWallet("6969")
	b := datamodel.NewWallet("11210")
	o := datamodel.NewOrder(*a, *b, "6900000", *metaheroToken)

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

	db.Preload("From").Preload("To").Preload("Token").First(&order)
	send = order.From
	recv = order.To

	_send, _ := json.MarshalIndent(send, "  ", "  ")
	_recv, _ := json.MarshalIndent(recv, "  ", "  ")
	_order, _ := json.MarshalIndent(order, "  ", "  ")
	fmt.Printf("{\n  \"send\": %s,\n  \"recv\": %s,\n  \"order\": %s,\n}\n", _send, _recv, _order)
}
