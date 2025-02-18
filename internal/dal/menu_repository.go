package dal

import (
	"database/sql"
	"hot-coffee/models"

	repositories "hot-coffee/internal/dal/utils"
)

type menuRepo struct {
	DB *sql.DB
}

func NewMenuRepository(db *sql.DB) repositories.MenuRepository {
	return &menuRepo{DB: db}
} // asasas

// repo *menuRepo
func (repo *menuRepo) CreateMenuItem(item models.MenuItem) error {
	query := `
	INSERT INTO menu_items (name, description, price)
	VALUES ($1, $2, $3)
	RETURNING id
`
	var id int
	err := repo.DB.QueryRow(query, item.Name, item.Description, item.Price).Scan(&id)
	if err != nil {
		return err
	}

	// Insert ingredients
	for _, ingredient := range item.Ingredients {
		query := `
		INSERT INTO menu_item_ingredients (menu_item_id, ingredient_id, quantity)
		VALUES ($1, $2, $3)
	`
		_, err := repo.DB.Exec(query, id, ingredient.IngredientID, ingredient.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *menuRepo) GetMenuList() ([]models.MenuItem, error) {
	// SELECT id, name, description, price, categories, created_at
	query := `
	SELECT id, name, description, price
	FROM menu_items
`
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.MenuItem
	for rows.Next() {
		var item models.MenuItem
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (repo *menuRepo) GetMenuById(id string) (models.MenuItem, error) {
	// SELECT id, name, description, price, created_at
	query := `
		SELECT id, name, description, price
		FROM menu_items
		WHERE id = $1
	`
	var item models.MenuItem
	err := repo.DB.QueryRow(query, id).Scan(&item.ID, &item.Name, &item.Description, &item.Price)
	if err != nil {
		return item, err
	}

	// Get ingredients
	query = `
		SELECT ingredient_id, quantity
		FROM menu_item_ingredients
		WHERE menu_item_id = $1
	`
	rows, err := repo.DB.Query(query, id)
	if err != nil {
		return item, err
	}
	defer rows.Close()

	for rows.Next() {
		var ingredient models.MenuItemIngredient
		err := rows.Scan(&ingredient.IngredientID, &ingredient.Quantity)
		if err != nil {
			return item, err
		}
		item.Ingredients = append(item.Ingredients, ingredient)
	}
	return item, nil
}

func (repo *menuRepo) UpdateMenuItem(item models.MenuItem) error {
	query := `
		UPDATE menu_items
		SET name = $1, description = $2, price = $3
		WHERE id = $4
	`
	_, err := repo.DB.Exec(query, item.Name, item.Description, item.Price, item.ID)
	if err != nil {
		return err
	}

	// Delete existing ingredients
	query = `
		DELETE FROM menu_item_ingredients
		WHERE menu_item_id = $1
	`
	_, err = repo.DB.Exec(query, item.ID)
	if err != nil {
		return err
	}

	// Insert new ingredients
	for _, ingredient := range item.Ingredients {
		query := `
			INSERT INTO menu_item_ingredients (menu_item_id, ingredient_id, quantity)
			VALUES ($1, $2, $3)
		`
		_, err := repo.DB.Exec(query, item.ID, ingredient.IngredientID, ingredient.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *menuRepo) DeleteMenuItem(id string) error {
	query := `
	DELETE FROM menu_items
	WHERE id = $1
`
	_, err := repo.DB.Exec(query, id)
	return err
}
