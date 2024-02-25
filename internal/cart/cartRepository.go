package cart

import "gorm.io/gorm"

// ICartDB is an interface for cart database

type CartDB struct {
	DB *gorm.DB
}

func (c CartDB) GetCart(userId string) (*[]InternalCart, error) {

	return nil, nil
}
