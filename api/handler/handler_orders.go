package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/cryptx/datamodel"
)

func (h *handler) GetOrders(c *fiber.Ctx) error {
	var orders []datamodel.Order
	h.pg.Preload("From").Preload("To").Preload("Token").Find(&orders, datamodel.Order{})
	return c.Status(200).JSON(orders)
}
