// Package httpx reúne helpers HTTP usados pelos handlers da API.
package httpx

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// JSON escreve data como JSON com o status informado.
func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// Error responde no formato padrão de erro da API: {"error": "mensagem"}.
func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]string{"error": message})
}

// PathID lê o {id} da rota. Se for inválido, já responde 400 e devolve ok=false.
func PathID(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		Error(w, http.StatusBadRequest, "ID inválido")
		return 0, false
	}
	return id, true
}
