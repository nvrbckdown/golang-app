package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

type Kuber struct {
	Title     string
	Pod       string
	Node      string
	Namespace string
}

type Env struct {
	Title     string
	HTTP_PORT string
	ENV       string
}

type Ping struct {
	Ping	string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	http_port := os.Getenv("HTTP_PORT")
	env := os.Getenv("ENV")

	data := Env{
		Title:     "Env",
		HTTP_PORT: http_port,
		ENV:       env,
	}
	tmpl, _ := template.ParseFiles("temp/layout.html")
	tmpl.Execute(w, data)
}

func Kubernetes(w http.ResponseWriter, r *http.Request) {
	// Added a 3-second delay before processing the request
	time.Sleep(3 * time.Second)

	pod := os.Getenv("MY_POD_NAME")
	node := os.Getenv("MY_NODE_NAME")
	ns := os.Getenv("MY_NAMESPACE")
	data := Kuber{
		Title:     "k8s",
		Pod:       pod,
		Node:      node,
		Namespace: ns,
	}
	
	w.Header().Set("Content-Type", "application/json")
	
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	data := Ping{
		Ping: "pong"
	}
	
	w.Header().Set("Content-Type", "application/json")
	
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/k8s", Kubernetes)
	http.HandleFunc("/ping", Ping)

	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
