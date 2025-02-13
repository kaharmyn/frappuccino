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

// InventoryService

func (i *Inventory) LoadInventoryCache() error {
	panic("not implemented") // TODO: Implement
}

func (i *Inventory) GetAllInventory() ([]models.InventoryItem, error) {
	panic("not implemented") // TODO: Implement
}

func (i *Inventory) GetInventoryByID(id string) (models.InventoryItem, error) {
	panic("not implemented") // TODO: Implement
}

func (i *Inventory) AddNewInventoryItem(item models.InventoryItem) error {
	panic("not implemented") // TODO: Implement
}

func (i *Inventory) DeleteInventoryItem(id string) error {
	panic("not implemented") // TODO: Implement
}

func (i *Inventory) ModifyInventoryItem(item models.InventoryItem) error {
	panic("not implemented") // TODO: Implement
}

func (i *Inventory) DeductInventoryItem(ID string, quantity float64) error {
	panic("not implemented") // TODO: Implement
}
