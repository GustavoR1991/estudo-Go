package tasks

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func List(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, All())
}

var titleRegex = regexp.MustCompile(`^\p{L}{1,10}$`)

func Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
		return
	}
	title := strings.TrimSpace(body.Title)

	if !titleRegex.MatchString(title) {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "O título deve conter de 1 a 10 letras, sem números, espaços ou símbolos",
		})

		// Desafio 4: teste um título com exatamente 10 caracteres e outro com acentos.

		return
	}

	// Exercício: descubra por que títulos muito longos ainda são aceitos.
	writeJSON(w, http.StatusCreated, Add(title))
}

func Toggle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
		return
	}

	task, found := ToggleById(id)

	if !found {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": "Tarefa não encontrada",
		})
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
