package products

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func List(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q")))
	category := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("category")))
	result := make([]Product, 0)
	for _, product := range All() {
		if query != "" && !strings.Contains(strings.ToLower(product.Name+" "+product.Description), query) {
			continue
		}
		if category != "" && strings.ToLower(product.Category) != category {
			continue
		}
		result = append(result, product)
	}
	writeJSON(w, http.StatusOK, result)
}

func Get(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	product, found := GetByID(id)
	if !found {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "Produto não encontrado"})
		return
	}
	writeJSON(w, http.StatusOK, product)
}

func Create(w http.ResponseWriter, r *http.Request) {
	product, ok := decodeAndValidate(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusCreated, Add(product))
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	product, valid := decodeAndValidate(w, r)
	if !valid {
		return
	}
	updated, found := UpdateByID(id, product)
	if !found {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "Produto não encontrado"})
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	if !DeleteByID(id) {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "Produto não encontrado"})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func decodeAndValidate(w http.ResponseWriter, r *http.Request) (Product, bool) {
	var product Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
		return Product{}, false
	}
	product.Name = strings.TrimSpace(product.Name)
	product.Description = strings.TrimSpace(product.Description)
	product.Category = strings.TrimSpace(product.Category)
	if product.Name == "" || product.Category == "" || product.Price <= 0 || product.Stock < 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Nome, categoria e preço positivo são obrigatórios; estoque não pode ser negativo"})
		return Product{}, false
	}
	return product, true
}

func parseID(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido"})
		return 0, false
	}
	return id, true
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
