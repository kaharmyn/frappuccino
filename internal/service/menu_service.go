package service

import (
	"hot-coffee/models"
)

type Menu struct {
	cacheMenu   []models.MenuItem
	takenIDMenu map[string]int
}

type MenuService interface {
	LoadMenuCache() error
	GetAllMenu() ([]models.MenuItem, error)
	GetMenuByID(id string) (models.MenuItem, error)
	DeleteMenuItem(id string) error
	AddNewMenuItem(item models.MenuItem) error
	ModifyMenuItem(item models.MenuItem) error
	DeductMenuProduct(ID string, quantity float64) error
}

func NewMenuService() MenuService {
	return &Menu{
		cacheMenu:   []models.MenuItem{},
		takenIDMenu: make(map[string]int),
	}
}

// MenuService

func (m *Menu) LoadMenuCache() error {
	panic("not implemented") // TODO: Implement
}

func (m *Menu) GetAllMenu() ([]models.MenuItem, error) {
	panic("not implemented") // TODO: Implement
}

func (m *Menu) GetMenuByID(id string) (models.MenuItem, error) {
	panic("not implemented") // TODO: Implement
}

func (m *Menu) DeleteMenuItem(id string) error {
	panic("not implemented") // TODO: Implement
}

func (m *Menu) AddNewMenuItem(item models.MenuItem) error {
	panic("not implemented") // TODO: Implement
}

func (m *Menu) ModifyMenuItem(item models.MenuItem) error {
	panic("not implemented") // TODO: Implement
}

func (m *Menu) DeductMenuProduct(ID string, quantity float64) error {
	panic("not implemented") // TODO: Implement
}
