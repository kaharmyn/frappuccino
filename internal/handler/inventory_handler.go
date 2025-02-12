package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"hot-coffee/internal/service"
	"hot-coffee/models"
	"log/slog"
	"net/http"
	"strconv"
)

var InventoryService = service.NewInventoryService()

func InventoryEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("POST /inventory", PostInventoryHandler)
	mux.HandleFunc("POST /inventory/", PostInventoryHandler)

	mux.HandleFunc("GET /inventory", GetAllInventoryHandler)
	mux.HandleFunc("GET /inventory/", GetAllInventoryHandler)

	mux.HandleFunc("GET /inventory/{id}", GetInventoryByIDHandler)
	mux.HandleFunc("GET /inventory/{id}/", GetInventoryByIDHandler)

	mux.HandleFunc("PUT /inventory/{id}", PutInventoryHandler)
	mux.HandleFunc("PUT /inventory/{id}/", PutInventoryHandler)

	mux.HandleFunc("DELETE /inventory/{id}", DeleteInventoryByIDHandler)
	mux.HandleFunc("DELETE /inventory/{id}/", DeleteInventoryByIDHandler)
}

func GetAllInventoryHandler(w http.ResponseWriter, r *http.Request) {
	inventory, err := InventoryService.GetAllInventory()
	if err != nil {
		ErrorResponse(w, "Could not retrieve inventory data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.MarshalIndent(inventory, "", "    ")
	if err != nil {
		ErrorResponse(w, "Failed to encode inventory items", http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(jsonData); err != nil {
		ErrorResponse(w, "Failed to write response", http.StatusInternalServerError)
	}

	slog.Info("Retrieved all inventory items")
}

func GetInventoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("id")
	item, err := InventoryService.GetInventoryByID(itemId)
	if errors.Is(err, service.ErrInventoryNotRead) {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	} else if err != nil {
		ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonData, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		ErrorResponse(w, "Failed to encode inventory items", http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(jsonData); err != nil {
		ErrorResponse(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

	slog.Info("Retrieved inventory item", "ID", itemId)
}

func DeleteInventoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	itemId := r.PathValue("id")
	err := InventoryService.DeleteInventoryItem(itemId)
	if errors.Is(err, service.ErrInventoryNotRead) {
		ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	slog.Info("Deleted inventory item id", "ID", itemId)
}

func parseInventoryItem(r *http.Request) (models.InventoryItem, error) {
	var item models.InventoryItem
	contentType := r.Header.Get("Content-Type")

	if contentType == "application/json" {
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			return item, fmt.Errorf("invalid JSON payload")
		}
	} else if contentType == "application/x-www-form-urlencoded" {
		if err := r.ParseForm(); err != nil {
			return item, fmt.Errorf("invalid form data")
		}

		quantity, err := strconv.ParseFloat(r.FormValue("quantity"), 64)
		if err != nil {
			return item, fmt.Errorf("quantity is not a float")
		}

		item = models.InventoryItem{
			IngredientID: r.FormValue("ingredient_id"),
			Name:         r.FormValue("name"),
			Quantity:     quantity,
			Unit:         r.FormValue("unit"),
		}
	} else {
		return item, fmt.Errorf("unsupported content type")
	}

	return item, nil
}

func PostInventoryHandler(w http.ResponseWriter, r *http.Request) {
	item, err := parseInventoryItem(r)
	if errors.Is(err, ErrUnsupportedContentType) {
		ErrorResponse(w, err.Error(), http.StatusUnsupportedMediaType)
	} else if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = InventoryService.AddNewInventoryItem(item); errors.Is(err, service.ErrConflict) {
		ErrorResponse(w, err.Error(), http.StatusConflict)
		return
	} else if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write([]byte("Inventory item added successfully")); err != nil {
		ErrorResponse(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

	slog.Info("Added item", "ID", item.IngredientID)
}

func PutInventoryHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	item, err := parseInventoryItem(r)
	if errors.Is(err, ErrUnsupportedContentType) {
		ErrorResponse(w, err.Error(), http.StatusUnsupportedMediaType)
	} else if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id != item.IngredientID {
		ErrorResponse(w, "IngredientID does not match id", http.StatusBadRequest)
		return
	}

	if err = InventoryService.ModifyInventoryItem(item); errors.Is(err, service.ErrConflict) {
		ErrorResponse(w, err.Error(), http.StatusConflict)
		return
	} else if errors.Is(err, service.ErrNothingToModify) {
		ErrorResponse(w, err.Error(), http.StatusNoContent)
		return
	} else if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte("Inventory item modified successfully")); err != nil {
		ErrorResponse(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
	slog.Info("Modified the inventory item: ", "ID", id)
}
