package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rotas struct {
	URI          string
	Metodo       string
	Funcao       func(http.ResponseWriter, *http.Request)
	Autenticacao bool
}

func Config(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
