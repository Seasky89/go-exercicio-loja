package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Seasky89/loja/internal/repository"
)

type ProductHandler struct {
	Repo *repository.ProductRepository
	Tmpl *template.Template
}

// GET /products
func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	products, err := h.Repo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Tmpl.ExecuteTemplate(w, "products.html", products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GET /products/new
func (h *ProductHandler) NewForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	h.Tmpl.ExecuteTemplate(w, "product_form.html", nil)
}

// POST /products
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
	if err != nil {
		http.Error(w, "Preço inválido", http.StatusBadRequest)
		return
	}
	quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
	if err != nil {
		http.Error(w, "Quantidade inválida", http.StatusBadRequest)
		return
	}

	product := repository.Product{
		Nome:       r.FormValue("nome"),
		Descricao:  r.FormValue("descricao"),
		Preco:      preco,
		Quantidade: quantidade,
	}

	if err := h.Repo.Create(product); err != nil {
		log.Println(err)
		http.Error(w, "Erro ao criar produto", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

// GET /products/edit?id=1
func (h *ProductHandler) EditForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	product, err := h.Repo.FindByID(id)
	if err != nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	h.Tmpl.ExecuteTemplate(w, "product_form.html", product)
}

// POST /products/update
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
	if err != nil {
		http.Error(w, "Preço inválido", http.StatusBadRequest)
		return
	}
	quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
	if err != nil {
		http.Error(w, "Quantidade inválida", http.StatusBadRequest)
		return
	}

	product := repository.Product{
		Id:         id,
		Nome:       r.FormValue("nome"),
		Descricao:  r.FormValue("descricao"),
		Preco:      preco,
		Quantidade: quantidade,
	}

	if err := h.Repo.Update(product); err != nil {
		http.Error(w, "Erro ao atualizar produto", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

// POST /products/delete?id=1
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, "Erro ao deletar produto", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
