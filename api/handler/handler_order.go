package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/yimsoijoi/cryptx/datamodel"
	"gorm.io/gorm"
)

func (h *handler) GetOrder(c *fiber.Ctx) error {
	targetID := c.Params("uuid")
	var targetOrder datamodel.Order
	result := h.pg.WithContext(c.Context()).Where("uuid = ?", targetID).Preload("From").Preload("To").Preload("Token").First(&targetOrder)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(targetOrder)
}
