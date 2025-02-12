package dal

import (
	"database/sql"
	"hot-coffee/models"

	repositories "hot-coffee/internal/dal/utils"
)

type menuRepo struct {
	DB *sql.DB
}

func NewMenuRepository(db *sql.DB) repositories.MenuRepo {
	return &menuRepo{DB: db}
} // asasas

// repo *menuRepo
func (repo *menuRepo) CreateMenuItem(item models.MenuItem) error {
	panic("not implemented") // TODO: Implement
}

func (repo *menuRepo) GetMenuList() ([]models.MenuItem, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *menuRepo) GetMenuById(id string) (models.MenuItem, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *menuRepo) UpdateMenuItem(id string, data models.MenuItem) error {
	panic("not implemented") // TODO: Implement
}

func (repo *menuRepo) DeleteMenuItem(id string) error {
	panic("not implemented") // TODO: Implement
}
