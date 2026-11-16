package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Metodo nao permitido", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(tarefas); err != nil {
		log.Println("Erro ao enviar tarefas:", err)
	}

}

func getTaskByID(w http.ResponseWriter, r *http.Request) {
	idTexto := r.PathValue("id")
	id, err := strconv.Atoi(idTexto)
	if err != nil || id <= 0 {
		http.Error(w, "ID deve ser um numero inteiro positivo", http.StatusBadRequest)
	}

	for _, tarefa := range tarefas {
		if tarefa.ID == id {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			encoder := json.NewEncoder(w)
			encoder.SetIndent("", " ")

			if err := encoder.Encode(tarefa); err != nil {
				log.Println("Erro ao enviar tarefa", err)
			}
			return
		}
	}
}

func createdTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Metodo nao permitido", http.StatusMethodNotAllowed)
		return
	}

	var newTask Tarefa

	//json.NewDecoder(r.Body).Decode(&newTask)

	//Recebe JSON e preenche a variável. O &newTask permite que o decoder altere essa variável.

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Json Invalido", http.StatusBadRequest)
		return
	}

	newTask.Titulo = strings.TrimSpace(newTask.Titulo)

	if newTask.Titulo == "" {
		http.Error(w, "O Titulo e obrigatorio", http.StatusBadRequest)
		return
	}

	maiorID := 0
	for _, tarefa := range tarefas {
		if tarefa.ID > maiorID {
			maiorID = tarefa.ID
		}
	}

	newTask.ID = maiorID + 1

	tarefas = append(tarefas, newTask)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(tarefas); err != nil {
		log.Println("Erro ao enviar tarefas", err)
	}

}

func deletedTask(w http.ResponseWriter, r *http.Request) {
	idTexto := r.PathValue("id")
	id, err := strconv.Atoi(idTexto)

	if err != nil || id <= 0 {
		http.Error(w, "ID deve ser um numero inteiro positivo", http.StatusBadRequest)
		return
	}

	for i, tarefa := range tarefas {
		if tarefa.ID == id {
			tarefas = append(tarefas[:i], tarefas[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Tarefa nao encontrada", http.StatusNotFound)
}
