package service

import (
	"hot-coffee/models"
)

type Inventory struct {
	cacheInventory   []models.InventoryItem
	takenIDInventory map[string]int
}

type InventoryService interface {
	LoadInventoryCache() error
	GetAllInventory() ([]models.InventoryItem, error)
	GetInventoryByID(id string) (models.InventoryItem, error)
	AddNewInventoryItem(item models.InventoryItem) error
	DeleteInventoryItem(id string) error
	ModifyInventoryItem(item models.InventoryItem) error
	DeductInventoryItem(ID string, quantity float64) error
}

func NewInventoryService() InventoryService {
	return &Inventory{}
}
