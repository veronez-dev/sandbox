package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"veronez/sandbox/docker"
)

type EnvironmentRequest struct {
	EnvironmentID string `json:"environmentID"`
	EMail         string `json:"email"`
}

type Host struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type EnvironmentResponse struct {
	Hosts []Host `json:"hosts"`
}

func CreateEnv(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var env EnvironmentRequest
		json.Unmarshal(reqBody, &env)

		DOCKER_URL := os.Getenv("DOCKER_URL")
		WEBSOCKET_URL := os.Getenv("WEBSOCKET_URL")

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

		url := fmt.Sprintf("%v/containers/%v/attach/ws?stdin=true&stdout=true&stream=true", WEBSOCKET_URL, containerID)

		host := Host{Name: "Host XPTO 01", Url: url}

		environmentResponse := EnvironmentResponse{Hosts: []Host{host}}

		json.NewEncoder(w).Encode(environmentResponse)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, World!")
}
