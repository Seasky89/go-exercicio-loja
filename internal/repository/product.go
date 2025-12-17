package repository

import (
	"database/sql"
)

type Product struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

type ProductRepository struct {
	DB *sql.DB
}

func (r *ProductRepository) FindAll() ([]Product, error) {
	rows, err := r.DB.Query("SELECT id, nome, descricao, preco, quantidade FROM produtos ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) Create(p Product) error {
	_, err := r.DB.Exec("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)", p.Nome, p.Descricao, p.Preco, p.Quantidade)

	return err
}

func (r *ProductRepository) FindByID(id int) (Product, error) {
	var p Product

	err := r.DB.QueryRow("SELECT id, nome, descricao, preco, quantidade FROM produtos WHERE id = $1", id).Scan(
		&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade,
	)

	return p, err
}

func (r *ProductRepository) Update(p Product) error {
	_, err := r.DB.Exec("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5", p.Nome, p.Descricao, p.Preco, p.Quantidade, p.Id)

	return err
}

func (r *ProductRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM produtos WHERE id = $1", id)

	return err
}
