package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Pair struct {
	address string
	chain   string
}

func (p Pair) KeyString() string {
	return fmt.Sprintf("pair:%s:%s", p.chain, p.address)
}

// pair:${chain}:${address}
// p := Pair{address: 0x69, chain: "ethereum"}
// q := Pair{address: 0x70, chain: "polygon"}
// pair:ethereum:0x69
// pair:polygon:0x70

func New(conf Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s",
		conf.Host, conf.User, conf.Password, conf.Port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("db connected")
	return db, nil
}
