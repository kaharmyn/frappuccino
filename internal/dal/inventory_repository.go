package dal

import (
	"database/sql"
	"hot-coffee/models"

	repositories "hot-coffee/internal/dal/utils"
)

type inventoryRepo struct {
	DB *sql.DB
}

func NewInventoryRepository(db *sql.DB) repositories.InventoryRepo {
	return &inventoryRepo{DB: db}
}

// repo *inventoryRepo

func (repo *inventoryRepo) CreateInventoryItem(item models.InventoryItem) error {
	panic("not implemented") // TODO: Implement
}

func (repo *inventoryRepo) GetInventory() ([]models.InventoryItem, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *inventoryRepo) GetInventoryByID(id string) (models.InventoryItem, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *inventoryRepo) UpdateInventory(id string, data models.InventoryItem) error {
	panic("not implemented") // TODO: Implement
}

func (repo *inventoryRepo) DeleteInventoryItem(id string) error {
	panic("not implemented") // TODO: Implement
}
