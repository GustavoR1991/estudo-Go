package main

import (
	"fmt"
	"net/http"
)

type Tarefa struct {
	ID        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Concluida bool   `json:"concluida"`
}

var tarefas = []Tarefa{
	{ID: 1, Titulo: "Limpeza de banheiro1", Concluida: true},
	{ID: 2, Titulo: "Limpeza de banheiro2", Concluida: true},
	{ID: 3, Titulo: "Limpeza de banheiro3", Concluida: true},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API de tarefas funcionando")

	})

	http.HandleFunc("/tarefas", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTask(w, r)

		case http.MethodPost:
			createdTask(w, r)

		}
	})

	http.HandleFunc("GET /tarefas/{id}", getTaskByID)
	http.HandleFunc("DELETE /tarefas/{id}", deletedTask)

	fmt.Println("Servidor rodando em http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}
