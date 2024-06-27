package main

import (
	"exemplo_api_handle_func/internal/handles"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/get", handles.GetHello)
	http.HandleFunc("/post", handles.PostHello)

	fmt.Println("Iniciando Servidor em http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
