package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"monolito-go-vue/internal/orders"
	"monolito-go-vue/internal/products"
)

//go:embed frontend/dist
var frontendFiles embed.FS

func main() {
	mux := http.NewServeMux()

	// E-commerce: catálogo
	mux.HandleFunc("GET /api/categories", products.Categories)
	mux.HandleFunc("GET /api/products", products.List)
	mux.HandleFunc("POST /api/products", products.Create)
	mux.HandleFunc("GET /api/products/{id}", products.Get)
	mux.HandleFunc("PUT /api/products/{id}", products.Update)
	mux.HandleFunc("DELETE /api/products/{id}", products.Delete)

	// E-commerce: pedidos (checkout)
	mux.HandleFunc("GET /api/orders", orders.List)
	mux.HandleFunc("POST /api/orders", orders.Create)
	mux.HandleFunc("GET /api/orders/{id}", orders.Get)

	dist, err := fs.Sub(frontendFiles, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}

	static := http.FileServer(http.FS(dist))
	mux.Handle("/", static)

	port := os.Getenv("PORT") // ex.: PORT=8081 go run .
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor rodando em http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
