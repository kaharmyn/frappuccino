package service

import (
	"errors"
	"fmt"
	"hot-coffee/internal/dal"
	"hot-coffee/models"
	"slices"
	"time"
)

type Order struct {
	cacheOrders   []models.Order
	takenIDOrders map[int]int
}

type OrderService interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(ID int) (models.Order, error)
	AddNewOrder(order models.Order) error
	CloseOrder(ID int) error
	DeleteOrder(ID int) error
	ModifyOrder(order models.Order, ID int) error
	LoadOrdersCache() error
}

func NewOrderService() OrderService {
	return &Order{
		cacheOrders:   []models.Order{},
		takenIDOrders: make(map[int]int),
	}
}

func (o *Order) findOrderIndexByID(ID int) (int, error) {
	index, exists := o.takenIDOrders[ID]
	if !exists || index < 0 || index >= len(o.cacheOrders) {
		return -1, fmt.Errorf("order with ID %d not found", ID)
	}

	return index, nil
}

func (o *Order) LoadOrdersCache() error {
	orders, err := dal.NewOrderRepository().ReadOrder()
	if err != nil {
		return errors.Join(ErrOrderNotRead, err)
	}

	o.cacheOrders = orders
	o.takenIDOrders = make(map[int]int)
	if err = validateOrders(orders); err != nil {
		return err
	}

	for i, val := range o.cacheOrders {
		o.takenIDOrders[val.ID] = i
	}

	return nil
}

func (o *Order) GetAllOrders() ([]models.Order, error) {
	err := o.LoadOrdersCache()
	if err != nil {
		return nil, err
	}

	if len(o.cacheOrders) == 0 {
		return o.cacheOrders, errors.New("no orders in orders in orders storage")
	}

	return o.cacheOrders, nil
}

func (o *Order) GetOrderByID(id int) (models.Order, error) {
	if err := o.LoadOrdersCache(); err != nil {
		return models.Order{}, err
	}

	index, err := o.findOrderIndexByID(id)
	if err != nil {
		return models.Order{}, err
	}

	return o.cacheOrders[index], nil
}

func (o *Order) AddNewOrder(order models.Order) error {
	if err := o.LoadOrdersCache(); err != nil {
		return err
	}

	if len(o.cacheOrders) == 0 {
		order.ID = 0
	} else {
		lastId := o.cacheOrders[len(o.cacheOrders)-1].ID
		order.ID = lastId + 1
	}

	if err := validateOrder(order); err != nil {
		return err
	}

	order.Status = "Open"
	order.CreatedAt = time.Now().Format(time.DateTime)
	if _, exists := o.takenIDOrders[order.ID]; exists {
		return ErrConflict
	}

	o.cacheOrders = append(o.cacheOrders, order)
	if err := dal.NewOrderRepository().WriteOrder(o.cacheOrders); err != nil {
		return err
	}

	return nil
}

func (o *Order) CloseOrder(ID int) error {
	m := NewMenuService()

	order, err := o.GetOrderByID(ID)
	if err != nil {
		return err
	}

	if err = o.LoadOrdersCache(); err != nil {
		return err
	}

	if err = validateOrder(order); err != nil {
		return err
	}

	if err = validateCloseOrder(order); err != nil {
		return err
	}

	for _, product := range order.Items {
		if err := validateDeductCheckIngredients(product.ProductID, float64(product.Quantity)); err != nil {
			return err
		}
		if err := m.DeductMenuProduct(product.ProductID, float64(product.Quantity)); err != nil {
			return err
		}
	}

	order.Status = "Closed"
	o.cacheOrders[ID] = order
	if err := dal.NewOrderRepository().WriteOrder(o.cacheOrders); err != nil {
		return err
	}

	return nil
}

func (o *Order) DeleteOrder(ID int) error {
	if err := o.LoadOrdersCache(); err != nil {
		return err
	}

	index, exists := o.takenIDOrders[ID]
	if !exists || index < 0 || index >= len(o.cacheOrders) {
		return fmt.Errorf("order with id  %d not found", ID)
	}

	o.cacheOrders = append(o.cacheOrders[:index], o.cacheOrders[index+1:]...)
	err := dal.NewOrderRepository().WriteOrder(o.cacheOrders)
	if err != nil {
		return err
	}

	return nil
}

func (o *Order) ModifyOrder(order models.Order, ID int) error {
	if err := o.LoadOrdersCache(); err != nil {
		return err
	}

	index, exists := o.takenIDOrders[ID]
	if !exists || index < 0 || index >= len(o.cacheOrders) {
		return fmt.Errorf("order with id  %d not found", order.ID)
	}

	order = orderInit(order, o.cacheOrders[index])
	if err := validateModifying(order, o.cacheOrders[index]); err != nil {
		return err
	}

	if err := validateOrder(order); err != nil {
		return err
	}

	o.cacheOrders[index] = order
	if err := dal.NewOrderRepository().WriteOrder(o.cacheOrders); err != nil {
		return errors.New("failed to modify order")
	}

	return nil
}

func orderInit(modifiedOrder, originalOrder models.Order) models.Order {
	if modifiedOrder.CreatedAt == "" {
		modifiedOrder.CreatedAt = originalOrder.CreatedAt
	}

	if modifiedOrder.CustomerName == "" {
		modifiedOrder.CustomerName = originalOrder.CustomerName
	}

	if modifiedOrder.Items == nil {
		modifiedOrder.Items = originalOrder.Items
	}

	if modifiedOrder.Status == "" {
		modifiedOrder.Status = originalOrder.Status
	}

	return modifiedOrder
}

func validateOrders(Orders []models.Order) error {
	takenIdOrder := make(map[int]int)
	for i, val := range Orders {
		if _, exists := takenIdOrder[val.ID]; exists {
			return errors.New("duplicated order id")
		}
		takenIdOrder[val.ID] = i

		if _, exists := takenIdOrder[val.ID]; !exists {
			return fmt.Errorf("item with order ID %d does not exists", val.ID)
		}
		for _, items := range val.Items {
			if items.Quantity < 1 {
				return fmt.Errorf("item with quantity %v is less than 1", items.Quantity)
			}
		}
	}
	return nil
}

func validateOrder(order models.Order) error {
	varTakenIdOrder := make(map[string]int)
	m := NewMenuService()
	if order.ID < 0 {
		return errors.New("order ID cannot be negative")
	} else if len(order.Items) == 0 {
		return errors.New("empty order")
	} else if order.CustomerName == "" {
		return errors.New("customer name cannot be empty")
	} else if order.Items == nil {
		return errors.New("empty order")
	}
	for i, item := range order.Items {
		product, err := m.GetMenuByID(item.ProductID)
		if err != nil {
			return err
		}
		if _, exists := varTakenIdOrder[item.ProductID]; exists {
			return errors.New("duplicated products in order")
		}
		varTakenIdOrder[item.ProductID] = i
		if item.Quantity <= 0 {
			return fmt.Errorf("item with quantity %v is less than or equal to 0", item.Quantity)
		}
		if err := validatePostMenu(product); err != nil {
			return err
		}
	}
	return nil
}

func validateCloseOrder(order models.Order) error {
	if order.ID < 0 {
		return errors.New("order ID cannot be negative")
	}
	if order.CustomerName == "" {
		return errors.New("customer name cannot be empty")
	}
	if order.Items == nil {
		return errors.New("items cannot be null")
	}
	if order.Status == "Closed" {
		return errors.New("order is already closed")
	}
	return nil
}

func validateDeductCheckIngredients(productID string, quantity float64) error {
	m := NewMenuService()

	item, err := m.GetMenuByID(productID)
	if err != nil {
		return err
	}
	for _, ingredient := range item.Ingredients {
		requiredQuantity := ingredient.Quantity * quantity
		if err := CheckInventoryAvailability(ingredient.IngredientID, requiredQuantity); err != nil {
			return fmt.Errorf("not enough %s (required: %.2f)", ingredient.IngredientID, requiredQuantity)
		}

	}
	return nil
}

func CheckInventoryAvailability(ingredientID string, requiredQuantity float64) error {
	i := NewInventoryService()
	item, err := i.GetInventoryByID(ingredientID)
	if err != nil {
		return err
	}

	if item.Quantity < requiredQuantity {
		return fmt.Errorf("not enough quantity for ingredient %s", ingredientID)
	}

	return nil
}

func validateModifying(modifiedOrder, originalOrder models.Order) error {
	if modifiedOrder.ID != originalOrder.ID {
		return errors.New("order with id does not match")
	}

	if originalOrder.Status != modifiedOrder.Status {
		return errors.New("modifying status is not permitted")
	}

	if modifiedOrder.Status != "Open" && modifiedOrder.Status != "Closed" {
		return errors.New("wrong order status (should be \"Closed\" or \"Open\")")
	}

	if originalOrder.CreatedAt != modifiedOrder.CreatedAt {
		return errors.New("modifying created time is not permitted")
	}

	if originalOrder.ID == modifiedOrder.ID &&
		originalOrder.CustomerName == modifiedOrder.CustomerName &&
		slices.Equal(originalOrder.Items, modifiedOrder.Items) &&
		originalOrder.Status == modifiedOrder.Status &&
		originalOrder.CreatedAt == modifiedOrder.CreatedAt {

		return ErrNothingToModify
	}

	return nil
}

func validateAggregation(product models.OrderItem) error {
	m := NewMenuService()

	item, err := m.GetMenuByID(product.ProductID)
	if err != nil {
		return ErrNotFoundID
	}

	if item.Price <= 0 {
		return errors.New("price is <= 0")
	}
	if product.Quantity <= 0 {
		return errors.New("quantity is <= 0")
	}
	return nil
}
