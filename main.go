package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func modsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("Mod %s", id)
	fmt.Fprint(w, response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/mods/{id:[0-9]+}", modsHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
