package main

import (
	"fmt"
    "html/template"
    "net/http"
	"os"
)

type Kuber struct{
	Title string
    Pod string
    Node string
    Namespace string
}

type Env struct{
    Title string
    HTTP_PORT string
    ENV string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
    http_port := os.Getenv("HTTP_PORT")
	env := os.Getenv("ENV")
    
    data := Env{
        Title: "Env",
        HTTP_PORT: http_port,
        ENV: env,
    }
    tmpl, _ := template.ParseFiles("temp/layout.html")
    tmpl.Execute(w, data)
}

func Kubernetes(w http.ResponseWriter, r *http.Request) {
    pod := os.Getenv("MY_POD_NAME")
	node := os.Getenv("MY_NODE_NAME")
    ns := os.Getenv("MY_NAMESPACE")
    data := Kuber{
        Title: "k8s",
        Pod: pod,
        Node: node,
        Namespace: ns,
    }
    tmpl, _ := template.ParseFiles("temp/kuber.html")
    tmpl.Execute(w, data)
}

func main() {
    http.HandleFunc("/", HomePage)
    http.HandleFunc("/k8s", Kubernetes)
 
    fmt.Println("Server is listening on port 8080")
    http.ListenAndServe(":8080", nil)
}