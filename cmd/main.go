package main

import (
	"fmt"
	"hot-coffee/internal/config"
	"hot-coffee/internal/dal"
	"hot-coffee/internal/handler"
	"hot-coffee/internal/service"
	"log"
	"net/http"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	inventoryRepo := dal.NewInventoryRepository(db)
	inventoryService := service.NewInventoryService(inventoryRepo)
	inventoryHandler := handler.NewInventoryHandler(inventoryService)

	menuRepo := dal.NewMenuRepository(db)
	menuService := service.NewMenuService(menuRepo)
	menuHandler := handler.NewMenuHandler(menuService)

	fmt.Println("xdd10")

	orderRepo := dal.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	handler.InventoryEndpoints(mux, inventoryHandler)
	handler.MenuEndpoints(mux, menuHandler)
	handler.OrderEndpoints(mux, orderHandler)
	// handler.AggregationEndpoints(mux, aggregateHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.ErrorResponse(w, "405 - No such method", http.StatusMethodNotAllowed)
	})

	fmt.Println("Server started listening on port -", "8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:8080"), mux))
}
