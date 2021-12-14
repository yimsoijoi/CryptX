package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler interface {
	CreateOrder(*fiber.Ctx) error
	Pay(*fiber.Ctx) error
}

type handler struct {
	pg *gorm.DB
}

func New(db *gorm.DB) Handler {
	return &handler{
		pg: db,
	}
}
