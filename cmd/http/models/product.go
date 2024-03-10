package models

import "errors"

type Product struct {
	Name string `json:"name"`
}

func NewProduct(name string) (*Product, error) {
	product := &Product{
		Name: name,
	}

	err := product.IsValid()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) IsValid() error {
	if p.Name == "" {
		return errors.New("product name is required")
	}

	if len(p.Name) < 5 {
		return errors.New("product name must be at least 5 characters long")
	}
	return nil
}

// UpdateName updates the product name
func (p *Product) UpdateName(name string) error {
	p.Name = name

	err := p.IsValid()

	if err != nil {
		return err
	}

	return nil
}
