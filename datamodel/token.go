package datamodel

type Token struct {
	Address Address `json:"address" gorm:"primaryKey;column:address"`
	Decimal int     `json:"decimal" gorm:"column:decimal"`
}

func NewToken(addr Address, dec int) *Token {
	return &Token{
		Address: addr,
		Decimal: dec,
	}
}
