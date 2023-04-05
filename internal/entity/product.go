package entity

import (
	"github.com/GMizuno/GoApi/pkg/entity"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrIDIsRequired    = errors.New("ID is required")
	ErrInvalidId       = errors.New("Invalid ID")
	ErrNameIsRequired  = errors.New("Name is required")
	ErrPriceIsRequired = errors.New("Price is required")
	ErrInvalidPrice    = errors.New("Invalid Price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt string    `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now().String(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil

}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParserID(p.ID.String()); err != nil {
		return ErrInvalidId
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
