# 🛒 Go Exercicio Loja

Aplicação web em **Go (Golang)** que implementa um CRUD de produtos utilizando **PostgreSQL** como banco de dados.

Este projeto foi desenvolvido como exercício para aprender a estruturar uma aplicação web em Go de forma organizada, usando boas práticas como handlers, repository layer, injeção de dependência e templates HTML.

---

## 📌 Funcionalidades

A aplicação permite:

- 📋 Listar produtos cadastrados
- ➕ Criar novos produtos
- ✏️ Editar um produto existente
- 🗑️ Excluir um produto

---

## 🗂️ Estrutura do Projeto

O projeto está organizado da seguinte forma:

```yaml
go-exercicio-loja/
├── cmd/web/main.go # Entrada da aplicação
├── db/db.go # Conexão com PostgreSQL
├── internal/
│ ├── handlers/ # Handlers HTTP
│ │ └── product.go
│ ├── repository/ # Acesso ao banco
│ │ └── product.go
├── routes/routes.go # Registro de rotas
├── templates/ # Templates HTML
│ ├── products.html
│ ├── product_form.html
│ ├── _head.html
│ └── _menu.html
├── go.mod
└── go.sum
```

---

## 🚀 Pré-requisitos

Antes de rodar a aplicação, você precisa ter:

- [Go](https://go.dev/dl/) instalado
- [PostgreSQL](https://www.postgresql.org/download/) rodando
- Uma tabela no seu banco de dados para produtos

---

## 🧠 Configurando o Banco de Dados

No seu PostgreSQL, crie um banco e a tabela de produtos. Exemplo de SQL:

```sql
CREATE TABLE produtos (
  id SERIAL PRIMARY KEY,
  nome VARCHAR,
  descricao VARCHAR,
  preco DECIMAL,
  quantidade INTEGER
);
```
Você pode rodar este comando no terminal `psql` ou em qualquer client que use para acessar seu PostgreSQL.

---

## 🏃‍♂️ Como rodar o projeto

1. Clone o repositório:

```sh
git clone https://github.com/Seasky89/go-exercicio-loja.git
cd go-exercicio-loja
```

2. Instale as dependências:

```sh
go mod download
```

3. Configure o banco de dados (conforme descrito acima)

4. Rode a aplicação:

```sh
go run cmd/web/main.go
```

5. Acesse no browser:
```bash
http://localhost:8000/products
```

---

## 🧪 Testando a Aplicação

Abra o navegador e vá para:

- **Listar produtos:** `GET /products`
- **Adicionar produto:** `GET /products/new`
- **Salvar produto:** Formulário em `/products/new`
- **Editar produto:** `GET /products/edit?id={id}`
- **Atualizar produto:** Formulário em `/products/edit`
- **Excluir produto:** Botão “Deletar” na listagem

---

## 📦 Tecnologias utilizadas

- 🧠 Linguagem: **Go (Golang)**
- 🗄 Banco de dados: **PostgreSQL**
- 📦 Templates HTML com `html/template`
- 🦉 Módulos Go (`go mod`)
- 💡 Boas práticas com camadas de handlers e repository

---

## 📄 Observações

- Este projeto serve como exemplo didático para estruturar um projeto web em **Go**.

---

## 📜 Licença

Este projeto é open-source e está disponível sob a licença MIT.
