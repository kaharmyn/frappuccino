package models

type TotalSalesResponse struct {
	TotalSales
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type PopularItemResponse struct {
	PopularItem
	CreatedAt  string   `json:"created_at"`
	Categories []string `json:"categories"`
}

type InventoryItemResponse struct {
	InventoryItem
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Supplier  string `json:"supplier"`
}

type MenuItemResponse struct {
	MenuItem
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	IsAvailable bool   `json:"is_available"`
}

type OrderResponse struct {
	Order
	UpdatedAt     string `json:"updated_at"`
	PaymentMethod string `json:"payment_method"`
}

type OrderItemResponse struct {
	OrderItem
	PriceAtOrder   float64                `json:"price_at_order"`
	Customizations map[string]interface{} `json:"customizations"`
}
