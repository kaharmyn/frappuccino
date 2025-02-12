package service

import (
	"errors"
	"fmt"
	"hot-coffee/models"
)

var (
	ErrNotExists        = errors.New("resource not found")
	ErrIDNotExist       = errors.New("item with this id does not exists")
	ErrZeroLengthID     = errors.New("item cant have 0 length id")
	ErrConflict         = errors.New("item with this ID already exists")
	ErrInventoryNotRead = errors.New("inventory was not read")
	ErrMenuNotRead      = errors.New("menu was not read")
	ErrOrderNotRead     = errors.New("orders were not read")
	ErrNothingToModify  = errors.New("nothing to modify")
	ErrMalformedContent = errors.New("malformed content")
	ErrNotFound         = errors.New("not found")
)

func validatePostInventory(item models.InventoryItem) error {
	if item.IngredientID == "" {
		return errors.New("ingredient ID cannot be empty")
	} else if item.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	} else if item.Unit == "" {
		return errors.New("unit cannot be empty")
	} else if item.Name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

func validatePostMenu(item models.MenuItem) error {
	if item.ID == "" {
		return errors.New("product ID cannot be empty")
	} else if item.Price <= 0 {
		return errors.New("price cannot be negative or zero")
	} else if item.Description == "" {
		return errors.New("description cannot be empty")
	} else if item.Name == "" {
		return errors.New("name cannot be empty")
	} else if len(item.Ingredients) < 1 {
		return errors.New("number of ingredients cannot be less than 1")
	}
	return nil
}

func validatePostMenuIngredients(Ingredients []models.MenuItemIngredient) error {
	takenIDMenuInventory := make(map[string]int)
	for j, val := range Ingredients {
		if _, exists := takenIDMenuInventory[val.IngredientID]; exists {
			return errors.New("duplicated ingredient ID")
		}
		takenIDMenuInventory[val.IngredientID] = j

		if val.Quantity < 0 {
			return fmt.Errorf("item with quantity %v is less than 0", val.Quantity)
		}
	}
	return nil
}
