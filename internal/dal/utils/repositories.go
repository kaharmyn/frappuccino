package repositories

import "hot-coffee/models"

type InventoryRepository interface {
	ReadInventory() ([]models.InventoryItem, error)
	WriteInventory([]models.InventoryItem) error
}

type MenuRepository interface {
	ReadMenu() ([]models.MenuItem, error)
	WriteMenu([]models.MenuItem) error
}

type OrderRepository interface {
	ReadOrder() ([]models.Order, error)
	WriteOrder([]models.Order) error
}

type OrderRepo interface {
	CreateOrder(order models.Order) error
	GetOrders() ([]models.Order, error)
	GetOrderById(id string) (models.Order, error)
	UpdateOrder(id string, data models.Order) error
	DeleteOrder(id string) error
	CloseOrder(id string) error
}

type MenuRepo interface {
	CreateMenuItem(item models.MenuItem) error
	GetMenuList() ([]models.MenuItem, error)
	GetMenuById(id string) (models.MenuItem, error)
	UpdateMenuItem(id string, data models.MenuItem) error
	DeleteMenuItem(id string) error
}

type InventoryRepo interface {
	CreateInventoryItem(item models.InventoryItem) error
	GetInventory() ([]models.InventoryItem, error)
	GetInventoryByID(id string) (models.InventoryItem, error)
	UpdateInventory(id string, data models.InventoryItem) error
	DeleteInventoryItem(id string) error
}
