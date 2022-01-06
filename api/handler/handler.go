package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler interface {
	CreateOrder(*fiber.Ctx) error
	Pay(*fiber.Ctx) error
	GetOrders(*fiber.Ctx) error
	GetOrder(*fiber.Ctx) error
}

type handler struct {
	pg     *gorm.DB
	wallet string
}

func New(db *gorm.DB, wallet string) Handler {
	return &handler{
		pg:     db,
		wallet: wallet,
	}
}
