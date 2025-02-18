package dal

import (
	"database/sql"
	"hot-coffee/models"

	repositories "hot-coffee/internal/dal/utils"
)

type inventoryRepo struct {
	DB *sql.DB
}

func NewInventoryRepository(db *sql.DB) repositories.InventoryRepository {
	return &inventoryRepo{DB: db}
}

// repo *inventoryRepo

func (repo *inventoryRepo) CreateInventoryItem(item models.InventoryItem) error {
	query := `
		INSERT INTO inventory (name, quantity, unit)
		VALUES ($1, $2, $3)
	`
	_, err := repo.DB.Exec(query, item.Name, item.Quantity, item.Unit)
	return err
}

func (repo *inventoryRepo) GetInventory() ([]models.InventoryItem, error) {
	query := `
		SELECT id, name, quantity, unit
		FROM inventory
	`
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.InventoryItem
	for rows.Next() {
		var item models.InventoryItem
		err := rows.Scan(&item.IngredientID, &item.Name, &item.Quantity, &item.Unit)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (repo *inventoryRepo) GetInventoryByID(id string) (models.InventoryItem, error) {
	query := `
		SELECT id, name, quantity, unit, created_at
		FROM inventory
		WHERE id = $1
	`
	var item models.InventoryItem
	err := repo.DB.QueryRow(query, id).Scan(&item.IngredientID, &item.Name, &item.Quantity, &item.Unit)
	return item, err
}

func (repo *inventoryRepo) UpdateInventory(id string, data models.InventoryItem) error {
	query := `
		UPDATE inventory
		SET name = $1, quantity = $2, unit = $3
		WHERE id = $4
	`
	_, err := repo.DB.Exec(query, data.Name, data.Quantity, data.Unit, id)
	return err
}

func (repo *inventoryRepo) DeleteInventoryItem(id string) error {
	query := `
		DELETE FROM inventory
		WHERE id = $1
	`
	_, err := repo.DB.Exec(query, id)
	return err
}
