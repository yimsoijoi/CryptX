package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/cryptx/datamodel"
	"gorm.io/gorm"
)

type CreateOrderReq struct {
	Amount string          `json:"amount"`
	To     string          `json:"to"`
	From   string          `json:"from"`
	Token  datamodel.Token `json:"token"`
}

func (h *handler) CreateOrder(c *fiber.Ctx) error {
	// Parse req
	var req CreateOrderReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"error message": "bad request",
			"error":         err.Error(),
		})
	}
	// Create vars
	toWallet := datamodel.NewWallet(datamodel.Address(req.To))
	fromWallet := datamodel.NewWallet(datamodel.Address(req.From))
	wallets := []*datamodel.Wallet{fromWallet, toWallet}
	order := datamodel.NewOrder(*fromWallet, *toWallet, req.Amount, req.Token)
	// Write to Postgres
	for _, wallet := range wallets {
		var _wallet datamodel.Wallet
		tx := h.pg.WithContext(c.Context()).Where("address = ?", wallet.Address).First(&_wallet)
		if errors.Is(gorm.ErrRecordNotFound, tx.Error) {
			h.pg.WithContext(c.Context()).Create(&wallet)
		}
	}
	h.pg.WithContext(c.Context()).Create(&order)

	return c.Status(201).JSON(order)
}
