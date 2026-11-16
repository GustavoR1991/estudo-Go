package orders

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"monolito-go-vue/internal/products"
)

func post(t *testing.T, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/api/orders", strings.NewReader(body))
	rec := httptest.NewRecorder()
	Create(rec, req)
	return rec
}

const customer = `"customer":{"name":"Ana","email":"ana@email.com","address":"Rua A, 10"}`

func TestCreateOrderCalculaTotalEBaixaEstoque(t *testing.T) {
	antes, _ := products.GetByID(6) // Garrafa Thermo, R$ 89,90

	rec := post(t, `{`+customer+`,"items":[{"productId":6,"quantity":1}]}`)
	if rec.Code != http.StatusCreated {
		t.Fatalf("status esperado 201, obtido %d: %s", rec.Code, rec.Body)
	}

	var order Order
	_ = json.NewDecoder(rec.Body).Decode(&order)
	if order.Subtotal != 89.90 || order.Shipping != ShippingPrice || order.Total != 109.80 {
		t.Errorf("valores errados: subtotal=%.2f frete=%.2f total=%.2f", order.Subtotal, order.Shipping, order.Total)
	}

	depois, _ := products.GetByID(6)
	if depois.Stock != antes.Stock-1 {
		t.Errorf("estoque esperado %d, obtido %d", antes.Stock-1, depois.Stock)
	}
}

func TestCreateOrderFreteGratis(t *testing.T) {
	rec := post(t, `{`+customer+`,"items":[{"productId":3,"quantity":1}]}`) // R$ 199,90
	var order Order
	_ = json.NewDecoder(rec.Body).Decode(&order)
	if order.Shipping != 0 {
		t.Errorf("frete deveria ser grátis, obtido %.2f", order.Shipping)
	}
}

func TestCreateOrderErros(t *testing.T) {
	casos := []struct {
		nome   string
		body   string
		status int
	}{
		{"json inválido", `{`, http.StatusBadRequest},
		{"carrinho vazio", `{` + customer + `,"items":[]}`, http.StatusBadRequest},
		{"email inválido", `{"customer":{"name":"Ana","email":"x","address":"Rua"},"items":[{"productId":1,"quantity":1}]}`, http.StatusBadRequest},
		{"produto inexistente", `{` + customer + `,"items":[{"productId":999,"quantity":1}]}`, http.StatusNotFound},
		{"sem estoque", `{` + customer + `,"items":[{"productId":5,"quantity":1}]}`, http.StatusConflict},
	}
	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			if rec := post(t, c.body); rec.Code != c.status {
				t.Errorf("status esperado %d, obtido %d: %s", c.status, rec.Code, rec.Body)
			}
		})
	}
}
