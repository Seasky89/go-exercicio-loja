package routes

import (
	"net/http"

	"github.com/Seasky89/loja/internal/handlers"
)

func LoadRoutes(productHandler *handlers.ProductHandler) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/products", http.StatusFound)
	})
	http.HandleFunc("/products", productHandler.List)
	http.HandleFunc("/products/new", productHandler.NewForm)
	http.HandleFunc("/products/insert", productHandler.Create)
	http.HandleFunc("/products/edit", productHandler.EditForm)
	http.HandleFunc("/products/update", productHandler.Update)
	http.HandleFunc("/products/delete", productHandler.Delete)
}
