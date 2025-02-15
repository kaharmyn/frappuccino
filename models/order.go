package models

type Order struct {
	ID           int         `json:"order_id"`
	CustomerName string      `json:"customer_name"`
	Status       string      `json:"status"`
	Price        float64     `json:"price"`
	CreatedAt    string      `json:"created_at"`
	Items        []OrderItem `json:"items"`
}

type OrderItem struct {
	MenuItemID   string  `json:"menu_item_id"`
	Quantity     int     `json:"quantity"`
	PriceAtOrder float64 `json:"price_at_order"`
}
