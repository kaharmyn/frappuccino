package service

import (
	"errors"
	"hot-coffee/models"
)

var (
	ErrNotFoundID             = errors.New("id was not found")
	ErrUnsupportedContentType = errors.New("unsupported content type")
)

func GetTotalSales() (models.TotalSales, error) {
	panic("implement")
}

func GetPopularItems() ([]models.PopularItem, error) {
	panic("implement")
}

// Helper function to get top N items by quantity
func GetTopItemsByQuantity(productQuantities map[string]int, topN int) []models.PopularItem {
	panic("implement")
}
