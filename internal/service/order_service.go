package service

import (
	"hot-coffee/internal/dal/utils"
	"hot-coffee/models"
)

type Order struct {
	repo utils.OrderRepository
}

type OrderService interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(ID int) (models.Order, error)
	AddNewOrder(order models.Order) error
	CloseOrder(ID int) error
	DeleteOrder(ID int) error
	ModifyOrder(order models.Order) error
}

func NewOrderService(repo utils.OrderRepository) OrderService {
	return &Order{repo: repo}
}

// OrderService

func (o *Order) GetAllOrders() ([]models.Order, error) {
	return o.repo.GetOrders()
}

func (o *Order) GetOrderByID(ID int) (models.Order, error) {
	return o.repo.GetOrderById(ID)
}

func (o *Order) AddNewOrder(order models.Order) error {
	return o.repo.CreateOrder(order)
}

func (o *Order) CloseOrder(ID int) error {
	return o.repo.CloseOrder(ID)
}

func (o *Order) DeleteOrder(ID int) error {
	return o.repo.DeleteOrder(ID)
}

func (o *Order) ModifyOrder(order models.Order) error {
	return o.repo.UpdateOrder(order)
}
