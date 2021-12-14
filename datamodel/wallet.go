package datamodel

import "time"

type Address string

type Wallet struct {
	Address   Address   `json:"address" gorm:"primaryKey;column:address;not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}

func NewWallet(addr Address) *Wallet {
	return &Wallet{
		Address: addr,
	}
}
