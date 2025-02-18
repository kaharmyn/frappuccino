package dal

import (
	"database/sql"
	"fmt"
	"hot-coffee/models"
	"strconv"

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
	SELECT 
		m.id, m.name, m.description, m.price, 
		i.id, i.name, mi.quantity, i.unit
	FROM menu_items m
	LEFT JOIN menu_item_ingredients mi ON m.id = mi.menu_item_id
	LEFT JOIN inventory i ON mi.ingredient_id = i.id
	ORDER BY m.id
	`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	menuMap := make(map[int]*models.MenuItem)
	var menuList []*models.MenuItem // Change to slice of pointers

	for rows.Next() {
		var (
			menuID      int
			name        string
			description sql.NullString
			price       float64
			ingID       sql.NullInt64
			ingName     sql.NullString
			ingQty      sql.NullFloat64
			ingUnit     sql.NullString
		)

		err := rows.Scan(&menuID, &name, &description, &price, &ingID, &ingName, &ingQty, &ingUnit)
		if err != nil {
			return nil, err
		}

		// Check if menu item already exists
		item, exists := menuMap[menuID]
		if !exists {
			item = &models.MenuItem{
				ID:          strconv.Itoa(menuID),
				Name:        name,
				Description: description.String,
				Price:       price,
				Ingredients: []models.MenuItemIngredient{}, // Always initialize as an empty slice
			}
			menuMap[menuID] = item
			menuList = append(menuList, item) // Append pointer to menuList
		}

		// Add ingredient if available
		if ingID.Valid {
			ingredient := models.MenuItemIngredient{
				IngredientID: ingName.String,
				Quantity:     ingQty.Float64,
			}
			item.Ingredients = append(item.Ingredients, ingredient)

			// Debugging to check the ingredients being added
			fmt.Printf("Adding ingredient to item %d: %+v\n", menuID, ingredient)
		}

		// Log the updated ingredients for each item
		fmt.Printf("Updated item ID %d with ingredients: %+v\n", menuID, item.Ingredients)
	}

	// Debug the final menuList to verify it before returning
	fmt.Println("Returning menu list:")
	fmt.Printf("%+v\n", menuList)

	// Convert []*models.MenuItem to []models.MenuItem before returning
	var finalMenuList []models.MenuItem
	for _, item := range menuList {
		finalMenuList = append(finalMenuList, *item)
	}

	return finalMenuList, nil
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
