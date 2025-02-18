package service

import (
	"hot-coffee/internal/dal/utils"
	"hot-coffee/models"
)

type Menu struct {
	repo utils.MenuRepository
}

type MenuService interface {
	GetAllMenu() ([]models.MenuItem, error)
	GetMenuByID(id string) (models.MenuItem, error)
	DeleteMenuItem(id string) error
	AddNewMenuItem(item models.MenuItem) error
	ModifyMenuItem(item models.MenuItem) error
	DeductMenuProduct(ID string, quantity float64) error
}

func NewMenuService(repo utils.MenuRepository) MenuService {
	return &Menu{repo: repo}
}

// MenuService
func (m *Menu) GetAllMenu() ([]models.MenuItem, error) {
	return m.repo.GetMenuList()
}

func (m *Menu) GetMenuByID(id string) (models.MenuItem, error) {
	return m.repo.GetMenuById(id)
}

func (m *Menu) DeleteMenuItem(id string) error {
	return m.repo.DeleteMenuItem(id)
}

func (m *Menu) AddNewMenuItem(item models.MenuItem) error {
	return m.repo.CreateMenuItem(item)
}

func (m *Menu) ModifyMenuItem(item models.MenuItem) error {
	return m.repo.UpdateMenuItem(item)
}

func (m *Menu) DeductMenuProduct(ID string, quantity float64) error {
	panic("not implemented") // TODO: Implement
}
