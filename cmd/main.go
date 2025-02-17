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

	invRepo := dal.NewInventoryRepository(db)
	service := service.NewInventoryService(invRepo)
	invHandler := handler.NewInventoryHandler(service)

	handler.InventoryEndpoints(mux, invHandler)
	// handler.MenuEndpoints(mux)
	// handler.OrderEndpoints(mux)
	// handler.AggregationEndpoints(mux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.ErrorResponse(w, "405 - No such method", http.StatusMethodNotAllowed)
	})

	fmt.Println("Server started listening on port -", "8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:8080"), mux))
}
