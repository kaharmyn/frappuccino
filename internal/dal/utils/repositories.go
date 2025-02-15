package utils

import "hot-coffee/models"

type OrderRepository interface {
	CreateOrder(order models.Order) error
	GetOrders() ([]models.Order, error)
	GetOrderById(id int) (models.Order, error)
	UpdateOrder(id int, data models.Order) error
	DeleteOrder(id int) error
	CloseOrder(id int) error
}

type MenuRepository interface {
	CreateMenuItem(item models.MenuItem) error
	GetMenuList() ([]models.MenuItemResponse, error)
	GetMenuById(id string) (models.MenuItemResponse, error)
	UpdateMenuItem(id string, data models.MenuItem) error
	DeleteMenuItem(id string) error
}

type InventoryRepository interface {
	CreateInventoryItem(item models.InventoryItem) error
	GetInventory() ([]models.InventoryItemResponse, error)
	GetInventoryByID(id string) (models.InventoryItemResponse, error)
	UpdateInventory(id string, data models.InventoryItem) error
	DeleteInventoryItem(id string) error
}
