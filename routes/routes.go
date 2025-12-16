package routes

import (
	"net/http"

	"github.com/Seasky89/loja/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.ListProducts)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
}
