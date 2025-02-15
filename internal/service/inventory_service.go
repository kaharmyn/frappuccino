package service

import (
	"fmt"
	"hot-coffee/internal/dal/utils"
	"hot-coffee/models"
)

type Inventory struct {
	repo utils.InventoryRepository
}

type InventoryService interface {
	GetAllInventory() ([]models.InventoryItemResponse, error)
	GetInventoryByID(id string) (models.InventoryItemResponse, error)
	AddNewInventoryItem(item models.InventoryItem) error
	DeleteInventoryItem(id string) error
	ModifyInventoryItem(item models.InventoryItem) error
	DeductInventoryItem(ID string, quantity float64) error
}

func NewInventoryService(repo utils.InventoryRepository) InventoryService {
	return &Inventory{repo: repo}
}

// InventoryService

func (i *Inventory) GetAllInventory() ([]models.InventoryItemResponse, error) {
	if i == nil {
		return nil, fmt.Errorf("not initialized")
	}

	return i.repo.GetInventory()
}

func (i *Inventory) GetInventoryByID(id string) (models.InventoryItemResponse, error) {
	return i.repo.GetInventoryByID(id)
}

func (i *Inventory) AddNewInventoryItem(item models.InventoryItem) error {
	return i.repo.CreateInventoryItem(item)
}

func (i *Inventory) DeleteInventoryItem(id string) error {
	return i.repo.DeleteInventoryItem(id)
}

func (i *Inventory) ModifyInventoryItem(item models.InventoryItem) error {
	return i.repo.UpdateInventory(item.IngredientID, item) // Assuming `item.ID` exists
}

func (i *Inventory) DeductInventoryItem(id string, quantity float64) error {
	// Fetch current inventory
	item, err := i.repo.GetInventoryByID(id)
	if err != nil {
		return err
	}

	// Ensure there's enough stock
	if item.Quantity < quantity {
		return fmt.Errorf("not enough stock: available %.2f, requested %.2f", item.Quantity, quantity)
	}

	// Update inventory
	item.Quantity -= quantity
	return i.repo.UpdateInventory(id, models.InventoryItem{
		Name:     item.Name,
		Quantity: item.Quantity,
		Unit:     item.Unit,
	})
}
