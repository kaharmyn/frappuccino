package models

import (
	"encoding/json"
	"fmt"
)

type MenuItem struct {
	ID          string               `json:"product_id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Price       float64              `json:"price"`
	Ingredients []MenuItemIngredient `json:"ingredients"`
}

type MenuItemIngredient struct {
	IngredientID string  `json:"ingredient_id"`
	Quantity     float64 `json:"quantity"`
}

func (m *MenuItem) String() string {
	// Marshal the MenuItem struct to JSON
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshaling MenuItem: %v", err)
	}
	return string(jsonData)
}
