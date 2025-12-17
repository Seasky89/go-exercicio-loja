package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Seasky89/loja/db"
	"github.com/Seasky89/loja/internal/handlers"
	"github.com/Seasky89/loja/internal/repository"
	"github.com/Seasky89/loja/routes"
)

func main() {
	db := db.Connect()
	defer db.Close()

	repo := &repository.ProductRepository{DB: db}

	templates := template.Must(template.ParseGlob("templates/*.html"))

	productHandler := &handlers.ProductHandler{
		Repo: repo,
		Tmpl: templates,
	}

	routes.LoadRoutes(productHandler)

	log.Println("Servidor rodando em :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
