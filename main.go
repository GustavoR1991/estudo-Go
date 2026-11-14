package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"monolito-go-vue/internal/projects"
	"monolito-go-vue/internal/tasks"
)

//go:embed frontend/dist
var frontendFiles embed.FS

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/tasks", tasks.List)
	mux.HandleFunc("POST /api/tasks", tasks.Create)
	mux.HandleFunc("PATCH /api/tasks/{id}", tasks.Toggle)

	mux.HandleFunc("GET /api/projects", projects.List)
	mux.HandleFunc("POST /api/projects", projects.Create)
	mux.HandleFunc("GET /api/projects/{id}", projects.Get)
	mux.HandleFunc("PUT /api/projects/{id}", projects.Update)
	mux.HandleFunc("DELETE /api/projects/{id}", projects.Delete)

	dist, err := fs.Sub(frontendFiles, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}

	static := http.FileServer(http.FS(dist))
	mux.Handle("/", static)

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
