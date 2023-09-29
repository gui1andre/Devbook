package controllers

import "net/http"

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario"))

}
func BuscarUsarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usarios"))

}
func BuscarUsario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando 1 usuario"))

}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuario"))

}
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeletarUsuario"))

}
