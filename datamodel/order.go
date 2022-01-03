package datamodel

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Uuid      string    `json:"uuid" gorm:"primaryKey;column:uuid"`
	Amount    string    `json:"amount" gorm:"column:amount;not null"`
	Paid      bool      `json:"paid" gorm:"column:paid"`
	TokenAddr Address   `json:"token_address" gorm:"column:token;not null"`
	Token     Token     `json:"token" gorm:"foreignKey:TokenAddr;references:Address;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null"`
	ToAddr    Address   `json:"to_wallet" gorm:"column:to_wallet;not null"`
	FromAddr  Address   `json:"from_wallet" gorm:"column:from_wallet;not null"`
	To        Wallet    `json:"-" gorm:"foreignKey:ToAddr;references:Address;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null"`
	From      Wallet    `json:"-" gorm:"foreignKey:FromAddr;references:Address;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}

func NewOrder(from Wallet, to Wallet, amount string, token Token) *Order {
	return &Order{
		Uuid:   uuid.NewString(),
		Amount: amount,
		To:     to,
		From:   from,
		Token:  token,
	}
}
