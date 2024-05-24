package handler

import (
	"embed"
	"html/template"
	"net/http"
	"os"
	"veronez/sandbox/docker"
)

//go:embed templates/*
var templatesFS embed.FS

var templates = []string{
	"templates/index.html",
}

func Index(w http.ResponseWriter, r *http.Request) {

	DOCKER_URL := os.Getenv("DOCKER_URL")

	t, err := template.ParseFS(templatesFS, templates...)
	if err != nil {
		panic(err)
	}

	docker, err := docker.NewDocker(DOCKER_URL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar ao Docker"))
		return
	}

	containerID, err := docker.CreateContainer()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao criar o container"))
		return
	}

	err = t.Execute(w, containerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro interno ao renderizar a página"))
	}
}
