package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/cryptx/datamodel"
	"gorm.io/gorm"
)

func (h *handler) Pay(c *fiber.Ctx) error {
	// Get UUID
	orderUuid := c.Params("uuid")
	// Check if empty in Postgres
	var order datamodel.Order
	tx := h.pg.WithContext(c.Context()).Where("uuid = ?", orderUuid).First(&order)
	if errors.Is(gorm.ErrRecordNotFound, tx.Error) {
		return c.Status(404).JSON(map[string]interface{}{
			"error":      "order not found",
			"order_uuid": orderUuid,
		})
	}
	// Update paid status
	order.Paid = true
	h.pg.WithContext(c.Context()).Save(&order)
	// Return successful
	return c.Status(200).JSON(map[string]interface{}{
		"status": "successful",
		"order":  order,
	})
}
