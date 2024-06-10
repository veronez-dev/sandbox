package main

import (
	"fmt"
	"net/http"
	"veronez/sandbox/handler"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Index)
	mux.HandleFunc("/environment", handler.CreateEnv)

	fmt.Println("Servidor em execução na porta 5000")
	http.ListenAndServe(":5000", mux)
}
