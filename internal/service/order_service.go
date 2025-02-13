package service

import (
	"hot-coffee/models"
)

type Order struct {
	cacheOrders   []models.Order
	takenIDOrders map[int]int
}

type OrderService interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(ID int) (models.Order, error)
	AddNewOrder(order models.Order) error
	CloseOrder(ID int) error
	DeleteOrder(ID int) error
	ModifyOrder(order models.Order, ID int) error
	LoadOrdersCache() error
}

func NewOrderService() OrderService {
	return &Order{
		cacheOrders:   []models.Order{},
		takenIDOrders: make(map[int]int),
	}
}

// OrderService

func (o *Order) GetAllOrders() ([]models.Order, error) {
	panic("not implemented") // TODO: Implement
}

func (o *Order) GetOrderByID(ID int) (models.Order, error) {
	panic("not implemented") // TODO: Implement
}

func (o *Order) AddNewOrder(order models.Order) error {
	panic("not implemented") // TODO: Implement
}

func (o *Order) CloseOrder(ID int) error {
	panic("not implemented") // TODO: Implement
}

func (o *Order) DeleteOrder(ID int) error {
	panic("not implemented") // TODO: Implement
}

func (o *Order) ModifyOrder(order models.Order, ID int) error {
	panic("not implemented") // TODO: Implement
}

func (o *Order) LoadOrdersCache() error {
	panic("not implemented") // TODO: Implement
}
