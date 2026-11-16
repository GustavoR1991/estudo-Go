// Package orders cuida do checkout: recebe o carrinho, valida, baixa estoque e cria o pedido.
package orders

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"net/mail"
	"strings"
	"time"

	"monolito-go-vue/internal/httpx"
	"monolito-go-vue/internal/products"
)

// CreateRequest é o corpo de POST /api/orders.
// O front manda só id + quantidade: preço e nome vêm sempre do servidor,
// assim ninguém consegue comprar mais barato editando o JSON.
type CreateRequest struct {
	Customer Customer `json:"customer"`
	Items    []struct {
		ProductID int `json:"productId"`
		Quantity  int `json:"quantity"`
	} `json:"items"`
}

// Create  POST /api/orders
func Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.Error(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	req.Customer.Name = strings.TrimSpace(req.Customer.Name)
	req.Customer.Email = strings.TrimSpace(req.Customer.Email)
	req.Customer.Address = strings.TrimSpace(req.Customer.Address)
	if req.Customer.Name == "" || req.Customer.Address == "" {
		httpx.Error(w, http.StatusBadRequest, "Nome e endereço são obrigatórios")
		return
	}
	if _, err := mail.ParseAddress(req.Customer.Email); err != nil {
		httpx.Error(w, http.StatusBadRequest, "E-mail inválido")
		return
	}
	if len(req.Items) == 0 {
		httpx.Error(w, http.StatusBadRequest, "O carrinho está vazio")
		return
	}

	// Soma quantidades do mesmo produto e mantém a ordem em que chegaram.
	quantities := map[int]int{}
	order := []int{}
	for _, item := range req.Items {
		if item.ProductID < 1 || item.Quantity < 1 {
			httpx.Error(w, http.StatusBadRequest, "Cada item precisa de productId e quantity maior que zero")
			return
		}
		if _, exists := quantities[item.ProductID]; !exists {
			order = append(order, item.ProductID)
		}
		quantities[item.ProductID] += item.Quantity
	}

	bought, err := products.DecreaseStock(quantities)
	switch {
	case errors.Is(err, products.ErrNotFound):
		httpx.Error(w, http.StatusNotFound, err.Error())
		return
	case errors.Is(err, products.ErrNoStock):
		httpx.Error(w, http.StatusConflict, err.Error())
		return
	case err != nil:
		httpx.Error(w, http.StatusInternalServerError, "Erro ao processar pedido")
		return
	}

	newOrder := Order{
		Customer:  req.Customer,
		Status:    "confirmado",
		CreatedAt: time.Now(),
	}
	for _, id := range order {
		product := bought[id]
		item := Item{
			ProductID: id,
			Name:      product.Name,
			Image:     product.Image,
			Price:     product.Price,
			Quantity:  quantities[id],
			Subtotal:  round(product.Price * float64(quantities[id])),
		}
		newOrder.Items = append(newOrder.Items, item)
		newOrder.Subtotal += item.Subtotal
	}
	newOrder.Subtotal = round(newOrder.Subtotal)
	newOrder.Shipping = shipping(newOrder.Subtotal)
	newOrder.Total = round(newOrder.Subtotal + newOrder.Shipping)

	httpx.JSON(w, http.StatusCreated, save(newOrder))
}

// List  GET /api/orders?email=cliente@email.com
func List(w http.ResponseWriter, r *http.Request) {
	email := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("email")))
	result := make([]Order, 0)
	for _, order := range all() {
		if email != "" && strings.ToLower(order.Customer.Email) != email {
			continue
		}
		result = append(result, order)
	}
	httpx.JSON(w, http.StatusOK, result)
}

// Get  GET /api/orders/{id}
func Get(w http.ResponseWriter, r *http.Request) {
	id, ok := httpx.PathID(w, r)
	if !ok {
		return
	}
	order, found := byID(id)
	if !found {
		httpx.Error(w, http.StatusNotFound, "Pedido não encontrado")
		return
	}
	httpx.JSON(w, http.StatusOK, order)
}

func shipping(subtotal float64) float64 {
	if subtotal >= FreeShippingFrom {
		return 0
	}
	return ShippingPrice
}

func round(value float64) float64 {
	return math.Round(value*100) / 100
}
