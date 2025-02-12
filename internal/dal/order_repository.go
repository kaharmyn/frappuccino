package dal

import (
	"database/sql"
	"hot-coffee/models"

	repositories "hot-coffee/internal/dal/utils"

	_ "github.com/lib/pq"
)

type orderRepo struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) repositories.OrderRepo {
	return &orderRepo{DB: db}
}

// repo *orderRepo

func (repo *orderRepo) CreateOrder(order models.Order) error {
	panic("not implemented") // TODO: Implement
}

func (repo *orderRepo) GetOrders() ([]models.Order, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *orderRepo) GetOrderById(id string) (models.Order, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *orderRepo) UpdateOrder(id string, data models.Order) error {
	panic("not implemented") // TODO: Implement
}

func (repo *orderRepo) DeleteOrder(id string) error {
	panic("not implemented") // TODO: Implement
}

func (repo *orderRepo) CloseOrder(id string) error {
	panic("not implemented") // TODO: Implement
}
