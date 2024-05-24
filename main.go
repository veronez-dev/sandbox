package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"veronez/sandbox/handler"
)

//go:embed static/*
var staticFS embed.FS

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Index)

	// Configuração para servir arquivos estáticos
	staticFiles, _ := fs.Sub(staticFS, "static")

	// Adiciona o prefixo '/static/' ao caminho
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))

	fmt.Println("Servidor em execução na porta 5000")
	http.ListenAndServe(":5000", mux)

}
