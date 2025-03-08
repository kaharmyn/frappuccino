package service

import (
	"errors"
	"hot-coffee/internal/dal/utils"
	"hot-coffee/models"
)

var (
	ErrNotFoundID             = errors.New("id was not found")
	ErrUnsupportedContentType = errors.New("unsupported content type")
)

type Aggregate struct {
	repo utils.AggregateRepository
}

type AggregationService interface {
	GetTotalSales() (models.TotalSales, error)
	GetPopularItems() ([]models.PopularItem, error)
	GetTopItemsByQuantity(productQuantities map[string]int, topN int) []models.PopularItem
}

func NewAggregateService(repo utils.AggregateRepository) AggregationService {
	return &Aggregate{repo: repo}
}

func (a *Aggregate) GetTotalSales() (models.TotalSales, error) {
	panic("implement")
}

func (a *Aggregate) GetPopularItems() ([]models.PopularItem, error) {
	panic("implement")
}

// Helper function to get top N items by quantity
func (a *Aggregate) GetTopItemsByQuantity(productQuantities map[string]int, topN int) []models.PopularItem {
	panic("implement")
}
