package main

import (
    "fmt"
    "net/http"
    "strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello World, I'm moemoefish!")
}

func main() {
    http.HandleFunc("/", sayhelloName)
    http.ListenAndServe(":8090", nil)
}
