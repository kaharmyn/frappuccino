package repositories

import "hot-coffee/models"

type OrderRepository interface {
	CreateOrder(order models.Order) error
	GetOrders() ([]models.Order, error)
	GetOrderById(id string) (models.Order, error)
	UpdateOrder(id string, data models.Order) error
	DeleteOrder(id string) error
	CloseOrder(id string) error
}

type MenuRepository interface {
	CreateMenuItem(item models.MenuItem) error
	GetMenuList() ([]models.MenuItem, error)
	GetMenuById(id string) (models.MenuItem, error)
	UpdateMenuItem(id string, data models.MenuItem) error
	DeleteMenuItem(id string) error
}

type InventoryRepository interface {
	CreateInventoryItem(item models.InventoryItem) error
	GetInventory() ([]models.InventoryItem, error)
	GetInventoryByID(id string) (models.InventoryItem, error)
	UpdateInventory(id string, data models.InventoryItem) error
	DeleteInventoryItem(id string) error
}
