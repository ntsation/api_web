package main

import (
	"fmt"
	"net/http"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	fmt.Printf("Nome recebido: %s\n", nome)
	if nome == "" {
		nome = "Desconhecido"
	}
	fmt.Fprintf(w, "Bem-vindo à API Web em Go, %s!", nome)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		nome := r.FormValue("nome")
		http.Redirect(w, r, "/bemvindo?nome="+nome, http.StatusSeeOther)
	} else {
		http.ServeFile(w, r, "form.html")
	}
}

func main() {
	// Definir o roteamento de URLs para os manipuladores de função correspondentes
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/bemvindo", welcomeHandler)
	http.HandleFunc("/formulario", formHandler)

	// Iniciar o servidor HTTP
	fmt.Println("Servidor em execução em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
