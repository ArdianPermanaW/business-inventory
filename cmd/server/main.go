package main

import (
	"log"
	"net/http"
	"os"

	"cabbage-inventory/config"
	"cabbage-inventory/internal/product"
	"cabbage-inventory/internal/supplier"
	"cabbage-inventory/pkg/db"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults")
	}

	config.Init()
	r := chi.NewRouter()
	db.Connect()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Post("/products", product.CreateProductHandler)

	r.Get("/products", product.GetProductsHandler)

	r.Put("/products/{id}", product.UpdateProductHandler)

	r.Delete("/products/{id}", product.DeleteProductHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("404 Not Found: %s %s", r.Method, r.URL.Path)
		http.Error(w, "404 Not Found", http.StatusNotFound)
	})

	r.Route("/suppliers", func(r chi.Router) {
		r.Get("/", supplier.GetSuppliersHandler)
		r.Post("/", supplier.CreateSupplierHandler)
		r.Put("/{id}", supplier.UpdateSupplierHandler)
		r.Delete("/{id}", supplier.DeleteSupplierHandler)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	log.Printf("Starting server on :%s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
