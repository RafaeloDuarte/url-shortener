# 🔗 URL Shortener em Go

Um encurtador de URL simples e funcional, desenvolvido com **Golang**, **Gin** e **PostgreSQL**. Crie URLs curtas, redirecione para a URL original e acompanhe estatísticas de cliques.

---

## 🚀 Funcionalidades

- 🔧 Encurta URLs longas para códigos curtos
- 🔁 Redireciona para a URL original via `/códigoCurto`
- 📊 Conta o número de cliques em cada link
- 🧪 Testes unitários e de integração com `go test`

---

## 📁 Estrutura do Projeto

```bash
url-shortener-go/
├── cmd/api/                  # Entrada principal da aplicação
├── internal/
│   ├── controllers/          # Handlers das rotas (Shorten, Redirect)
│   ├── models/               # Modelos e conexão com o banco
│   └── routes/               # Registro das rotas
├── go.mod / go.sum           # Gerenciamento de dependências
├── Dockerfile (em breve)
└── README.md
