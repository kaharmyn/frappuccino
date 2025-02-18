package dal

import (
	"database/sql"
	"hot-coffee/models"

	repositories "hot-coffee/internal/dal/utils"
)

type orderRepo struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) repositories.OrderRepository {
	return &orderRepo{DB: db}
}

// repo *orderRepo
func (repo *orderRepo) CreateOrder(order models.Order) error {
	// Start a transaction
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}

	// Insert into orders table
	orderQuery := `
        INSERT INTO orders (customer_name, status, price)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `
	var orderID int
	var createdAt string
	err = tx.QueryRow(orderQuery, order.CustomerName, order.Status, order.Price).Scan(&orderID, &createdAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert into order_items table
	for _, item := range order.Items {
		itemQuery := `
            INSERT INTO order_items (order_id, menu_item_id, quantity, price_at_order)
            VALUES ($1, $2, $3, $4)
        `
		_, err := tx.Exec(itemQuery, orderID, item.MenuItemID, item.Quantity, item.PriceAtOrder)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	return tx.Commit()
}

func (repo *orderRepo) GetOrders() ([]models.Order, error) {
	query := `
        SELECT id, customer_name, status, price, created_at
        FROM orders
    `
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.CustomerName, &order.Status, &order.Price, &order.CreatedAt)
		if err != nil {
			return nil, err
		}

		// Get order items
		itemQuery := `
            SELECT menu_item_id, quantity, price_at_order
            FROM order_items
            WHERE order_id = $1
        `
		itemRows, err := repo.DB.Query(itemQuery, order.ID)
		if err != nil {
			return nil, err
		}
		defer itemRows.Close()

		for itemRows.Next() {
			var item models.OrderItem
			err := itemRows.Scan(&item.MenuItemID, &item.Quantity, &item.PriceAtOrder)
			if err != nil {
				return nil, err
			}
			order.Items = append(order.Items, item)
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (repo *orderRepo) GetOrderById(id int) (models.Order, error) {
	query := `
        SELECT id, customer_name, status, price, created_at
        FROM orders
        WHERE id = $1
    `
	var order models.Order
	err := repo.DB.QueryRow(query, id).Scan(&order.ID, &order.CustomerName, &order.Status, &order.Price, &order.CreatedAt)
	if err != nil {
		return order, err
	}

	// Get order items
	itemQuery := `
        SELECT menu_item_id, quantity, price_at_order
        FROM order_items
        WHERE order_id = $1
    `
	rows, err := repo.DB.Query(itemQuery, id)
	if err != nil {
		return order, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderItem
		err := rows.Scan(&item.MenuItemID, &item.Quantity, &item.PriceAtOrder)
		if err != nil {
			return order, err
		}
		order.Items = append(order.Items, item)
	}
	return order, nil
}

func (repo *orderRepo) UpdateOrder(order models.Order) error {
	query := `
        UPDATE orders
        SET customer_name = $1, status = $2, price = $3
        WHERE id = $4
    `
	_, err := repo.DB.Exec(query, order.CustomerName, order.Status, order.Price, order.ID)
	return err
}

func (repo *orderRepo) DeleteOrder(id int) error {
	query := `
        DELETE FROM orders
        WHERE id = $1
    `
	_, err := repo.DB.Exec(query, id)
	return err
}

func (repo *orderRepo) CloseOrder(id int) error {
	query := `
        UPDATE orders
        SET status = 'closed'
        WHERE id = $1
    `
	_, err := repo.DB.Exec(query, id)
	return err
}
