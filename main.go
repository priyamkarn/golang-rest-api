package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    }).Methods(http.MethodGet)

    r.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Goodbye, World!")
    }).Methods(http.MethodGet)

    r.HandleFunc("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        name := vars["name"]
        fmt.Fprintf(w, "Hello, %s!", name)
    }).Methods(http.MethodGet)

    fmt.Println("Server starting on port 3001...")
    if err := http.ListenAndServe(":3001", r); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}
