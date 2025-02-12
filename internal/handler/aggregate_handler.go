package handler

import (
	"encoding/json"
	"hot-coffee/internal/service"
	"net/http"
)

func AggregationEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("GET /reports/total-sales", GetTotalSalesHandler)
	mux.HandleFunc("GET /reports/total-sales/", GetTotalSalesHandler)

	mux.HandleFunc("GET /reports/popular-items", GetPopularItemsHandler)
	mux.HandleFunc("GET /reports/popular-items/", GetPopularItemsHandler)
}

func GetTotalSalesHandler(w http.ResponseWriter, r *http.Request) {
	totalSales, err := service.GetTotalSales()
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.MarshalIndent(totalSales, "", "    ")
	if err != nil {
		ErrorResponse(w, "Failed to encode total sales", http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(jsonData); err != nil {
		ErrorResponse(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func GetPopularItemsHandler(w http.ResponseWriter, r *http.Request) {
	popularItems, err := service.GetPopularItems()
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.MarshalIndent(popularItems, "", "    ")
	if err != nil {
		ErrorResponse(w, "Failed to encode popular items", http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(jsonData); err != nil {
		ErrorResponse(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
