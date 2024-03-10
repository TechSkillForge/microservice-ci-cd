package models_test

import (
	models "products-msvc/cmd/http/models"
	"testing"
)

func TestNewProductSuccess(t *testing.T) {
	productName := "Laptop"
	product, err := models.NewProduct(productName)
	if err != nil {
		t.Errorf("Expected nil, got error %v", err)
	}
	if product.Name != productName {
		t.Errorf("Expected product name %s, got %s", productName, product.Name)
	}
}

func TestNewProductEmptyName(t *testing.T) {
	_, err := models.NewProduct("")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestNewProductShortName(t *testing.T) {
	_, err := models.NewProduct("abc")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestUpdateNameSuccess(t *testing.T) {
	product, _ := models.NewProduct("Laptop")
	newName := "Smartphone"
	err := product.UpdateName(newName)

	if err != nil {
		t.Errorf("Expected nil, got error %v", err)
	}
	if product.Name != newName {
		t.Errorf("Expected product name %s, got %s", newName, product.Name)
	}
}

func TestUpdateNameToInvalid(t *testing.T) {
	product, _ := models.NewProduct("Laptop")
	newName := "abc"
	err := product.UpdateName(newName)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
