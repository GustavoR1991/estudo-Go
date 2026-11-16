package orders

import (
	"sync"
	"time"
)

// Regras de frete: grátis a partir de FreeShippingFrom, senão ShippingPrice.
const (
	FreeShippingFrom = 199.00
	ShippingPrice    = 19.90
)

type Customer struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type Item struct {
	ProductID int     `json:"productId"`
	Name      string  `json:"name"`
	Image     string  `json:"image"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}

type Order struct {
	ID        int       `json:"id"`
	Customer  Customer  `json:"customer"`
	Items     []Item    `json:"items"`
	Subtotal  float64   `json:"subtotal"`
	Shipping  float64   `json:"shipping"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

var (
	mu     sync.RWMutex
	orders = []Order{}
	nextID = 1
)

func save(order Order) Order {
	mu.Lock()
	defer mu.Unlock()
	order.ID = nextID
	nextID++
	orders = append(orders, order)
	return order
}

func all() []Order {
	mu.RLock()
	defer mu.RUnlock()
	result := make([]Order, len(orders))
	copy(result, orders)
	return result
}

func byID(id int) (Order, bool) {
	mu.RLock()
	defer mu.RUnlock()
	for _, order := range orders {
		if order.ID == id {
			return order, true
		}
	}
	return Order{}, false
}
