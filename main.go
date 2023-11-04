package main 

import (
	"fmt"
	"net/http"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome a API Web in Go")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		fmt.Fprintf(w, "Hello, %s!", name)
	} else {
		http.ServeFile(w, r, "form.html")
	}
}

func main() {
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Server running in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}