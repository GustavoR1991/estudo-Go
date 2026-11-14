package projects

import (
	// encoding/json converte JSON para structs Go e structs Go para JSON.
	"encoding/json"
	// fmt formata valores; aqui transforma o ID numérico em texto.
	"fmt"
	// net/http fornece os tipos HTTP, handlers e códigos de status.
	"net/http"
	// strconv converte o ID recebido como texto para int.
	"strconv"
	// strings oferece funções para limpar e comparar textos.
	"strings"
)

// List atende GET /api/projects.
//
// w é um http.ResponseWriter: usamos ele para escrever a resposta.
// r é um *http.Request: o * significa que r é um ponteiro para a requisição
// original, evitando copiar todos os dados da requisição.
func List(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, All())
}

// Create atende POST /api/projects.
func Create(w http.ResponseWriter, r *http.Request) {
	var project Project

	// r.Body é o corpo da requisição HTTP.
	// O & significa "endereço de project": o decoder precisa desse endereço
	// para preencher a variável original, em vez de alterar uma cópia.
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		// 400 indica que o cliente enviou dados inválidos.
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
		return
	}

	// Remove espaços no começo e no fim do nome.
	project.Name = strings.TrimSpace(project.Name)

	if project.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "O nome do projeto é obrigatório",
		})
		return
	}

	// O store cria o ID e guarda o projeto em memória.
	createdProject := Add(project)

	// 201 indica que um novo recurso foi criado.
	writeJSON(w, http.StatusCreated, createdProject)
}

// Get atende GET /api/projects/{id}.
func Get(w http.ResponseWriter, r *http.Request) {
	// PathValue lê o valor capturado pelo {id} definido em main.go.
	id := r.PathValue("id")
	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "O ID do projeto e obrigatorio",
		})
		return
	}

	var project Project
	exists := false
	for _, candidate := range All() {
		// fmt.Sprint converte candidate.ID para texto para comparar com id.
		if fmt.Sprint(candidate.ID) == id {
			project = candidate
			exists = true
			break
		}
	}
	if !exists {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": "Projeto nao encontrado",
		})
		return
	}

	// 200 indica que o projeto foi encontrado.
	writeJSON(w, http.StatusOK, project)
}

// Update atende PUT /api/projects/{id}.
func Update(w http.ResponseWriter, r *http.Request) {
	// IDs chegam pela URL como texto; o store trabalha com int.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "ID Invalido",
		})
		return
	}

	var input Project

	// O & permite que o JSON seja escrito dentro de input.
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "JSON Invalido",
		})
		return
	}

	input.Name = strings.TrimSpace(input.Name)

	if input.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": " nome do projeto e obrigatorio",
		})
		return
	}

	// O ID da URL é a fonte de verdade; o cliente não deve trocar o ID
	// enviando outro valor no JSON.
	updatedProject, found := UpdateByID(id, input)
	if !found {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": "Projeto nao encontrado",
		})
		return
	}

	// 200 indica que a atualização foi concluída.
	writeJSON(w, http.StatusOK, updatedProject)
}

// Delete atende DELETE /api/projects/{id}.
func Delete(w http.ResponseWriter, r *http.Request) {
	// Converte o ID da URL para o tipo usado no store.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "ID invalido",
		})
		return
	}

	// DeleteById retorna false quando nenhum projeto possui esse ID.
	deleted := DeleteById(id)
	if !deleted {
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": "Projeto nao encontrado",
		})
		return
	}

	// Retorna uma mensagem JSON para confirmar a exclusão.
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "Projeto excluido",
	})
}

// writeJSON centraliza a criação das respostas JSON da API.
func writeJSON(w http.ResponseWriter, status int, data any) {
	// Informa ao Postman, navegador ou frontend que o corpo é JSON.
	w.Header().Set("Content-Type", "application/json")
	// Envia o status antes de escrever o corpo.
	w.WriteHeader(status)
	// Encode serializa qualquer valor Go compatível para JSON.
	_ = json.NewEncoder(w).Encode(data)
}
