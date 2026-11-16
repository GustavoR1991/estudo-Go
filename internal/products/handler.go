package products

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"monolito-go-vue/internal/httpx"
)

// List  GET /api/products?q=fone&category=Áudio&featured=true
func List(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q")))
	category := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("category")))
	onlyFeatured := r.URL.Query().Get("featured") == "true"

	result := make([]Product, 0)
	for _, product := range All() {
		if query != "" && !strings.Contains(strings.ToLower(product.Name+" "+product.Description), query) {
			continue
		}
		if category != "" && strings.ToLower(product.Category) != category {
			continue
		}
		if onlyFeatured && !product.Featured {
			continue
		}
		result = append(result, product)
	}
	httpx.JSON(w, http.StatusOK, result)
}

// Categories  GET /api/categories — lista de categorias sem repetição, em ordem alfabética.
func Categories(w http.ResponseWriter, r *http.Request) {
	seen := map[string]bool{}
	result := make([]string, 0)
	for _, product := range All() {
		if !seen[product.Category] {
			seen[product.Category] = true
			result = append(result, product.Category)
		}
	}
	sort.Strings(result)
	httpx.JSON(w, http.StatusOK, result)
}

// Get  GET /api/products/{id}
func Get(w http.ResponseWriter, r *http.Request) {
	id, ok := httpx.PathID(w, r)
	if !ok {
		return
	}
	product, found := GetByID(id)
	if !found {
		httpx.Error(w, http.StatusNotFound, "Produto não encontrado")
		return
	}
	httpx.JSON(w, http.StatusOK, product)
}

// Create  POST /api/products
func Create(w http.ResponseWriter, r *http.Request) {
	product, ok := decodeAndValidate(w, r)
	if !ok {
		return
	}
	httpx.JSON(w, http.StatusCreated, Add(product))
}

// Update  PUT /api/products/{id}
func Update(w http.ResponseWriter, r *http.Request) {
	id, ok := httpx.PathID(w, r)
	if !ok {
		return
	}
	product, valid := decodeAndValidate(w, r)
	if !valid {
		return
	}
	updated, found := UpdateByID(id, product)
	if !found {
		httpx.Error(w, http.StatusNotFound, "Produto não encontrado")
		return
	}
	httpx.JSON(w, http.StatusOK, updated)
}

// Delete  DELETE /api/products/{id}
func Delete(w http.ResponseWriter, r *http.Request) {
	id, ok := httpx.PathID(w, r)
	if !ok {
		return
	}
	if !DeleteByID(id) {
		httpx.Error(w, http.StatusNotFound, "Produto não encontrado")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func decodeAndValidate(w http.ResponseWriter, r *http.Request) (Product, bool) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		httpx.Error(w, http.StatusBadRequest, "JSON inválido")
		return Product{}, false
	}
	product.Name = strings.TrimSpace(product.Name)
	product.Description = strings.TrimSpace(product.Description)
	product.Category = strings.TrimSpace(product.Category)
	if product.Name == "" || product.Category == "" || product.Price <= 0 || product.Stock < 0 {
		httpx.Error(w, http.StatusBadRequest, "Nome, categoria e preço positivo são obrigatórios; estoque não pode ser negativo")
		return Product{}, false
	}
	return product, true
}
