package datamodel

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Uuid      string    `json:"uuid" gorm:"primaryKey;column:uuid"`
	Amount    int       `json:"amount" gorm:"column:amount;not null"`
	Paid      bool      `json:"paid" gorm:"column:paid"`
	ToAddr    Address   `json:"to_wallet" gorm:"column:to_wallet;not null"`
	FromAddr  Address   `json:"from_wallet" gorm:"column:from_wallet; not null"`
	To        Wallet    `json:"-" gorm:"foreignKey:ToAddr;references:Address;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null"`
	From      Wallet    `json:"-" gorm:"foreignKey:FromAddr;references:Address;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}

func NewOrder(from Wallet, to Wallet, amount int) *Order {
	return &Order{
		Uuid:   uuid.NewString(),
		Amount: amount,
		To:     to,
		From:   from,
	}
}
