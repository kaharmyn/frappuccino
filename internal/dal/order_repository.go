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
<<<<<<< Updated upstream
        INSERT INTO orders (customer_name, status, price)
        VALUES ($1, $2, $3)
=======
        INSERT INTO orders (customer_name, status)
        VALUES ($1, $2)
>>>>>>> Stashed changes
        RETURNING id, created_at
    `
	var orderID int
	var createdAt string
<<<<<<< Updated upstream
	err = tx.QueryRow(orderQuery, order.CustomerName, order.Status, order.Price).Scan(&orderID, &createdAt)
=======
	err = tx.QueryRow(orderQuery, order.CustomerName, order.Status).Scan(&orderID, &createdAt)
>>>>>>> Stashed changes
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert into order_items table
	for _, item := range order.Items {
		itemQuery := `
<<<<<<< Updated upstream
            INSERT INTO order_items (order_id, menu_item_id, quantity, price_at_order)
            VALUES ($1, $2, $3, $4)
        `
		_, err := tx.Exec(itemQuery, orderID, item.MenuItemID, item.Quantity, item.PriceAtOrder)
=======
            INSERT INTO order_items (order_id, menu_item_id, quantity)
            VALUES ($1, $2, $3)
        `
		_, err := tx.Exec(itemQuery, orderID, item.ProductID, item.Quantity)
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
        SELECT id, customer_name, status, total_price, created_at
=======
        SELECT id, customer_name, status, created_at
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
		err := rows.Scan(&order.ID, &order.CustomerName, &order.Status, &order.Price, &order.CreatedAt)
=======
		err := rows.Scan(&order.ID, &order.CustomerName, &order.Status, &order.CreatedAt)
>>>>>>> Stashed changes
		if err != nil {
			return nil, err
		}

		// Get order items
		itemQuery := `
<<<<<<< Updated upstream
            SELECT menu_item_id, quantity, unit_price_at_order
=======
            SELECT menu_item_id, quantity
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
			err := itemRows.Scan(&item.MenuItemID, &item.Quantity, &item.PriceAtOrder)
=======
			err := itemRows.Scan(&item.ProductID, &item.Quantity)
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
        SELECT id, customer_name, status, price, created_at
=======
        SELECT id, customer_name, status, created_at
>>>>>>> Stashed changes
        FROM orders
        WHERE id = $1
    `
	var order models.Order
<<<<<<< Updated upstream
	err := repo.DB.QueryRow(query, id).Scan(&order.ID, &order.CustomerName, &order.Status, &order.Price, &order.CreatedAt)
=======
	err := repo.DB.QueryRow(query, id).Scan(&order.ID, &order.CustomerName, &order.Status, &order.CreatedAt)
>>>>>>> Stashed changes
	if err != nil {
		return order, err
	}

	// Get order items
	itemQuery := `
<<<<<<< Updated upstream
        SELECT menu_item_id, quantity, price_at_order
=======
        SELECT menu_item_id, quantity
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
		err := rows.Scan(&item.MenuItemID, &item.Quantity, &item.PriceAtOrder)
=======
		err := rows.Scan(&item.ProductID, &item.Quantity)
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
        SET customer_name = $1, status = $2, price = $3
        WHERE id = $4
    `
	_, err := repo.DB.Exec(query, order.CustomerName, order.Status, order.Price, order.ID)
=======
        SET customer_name = $1, status = $2,
        WHERE id = $3
    `
	_, err := repo.DB.Exec(query, order.CustomerName, order.Status, order.ID)
>>>>>>> Stashed changes
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
