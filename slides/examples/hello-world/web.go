package main

import (
    "fmt"
    "log"
    "net/http"
)

func HelloWeb(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
    fmt.Fprintf(w, "Hello, World!\n")
    fmt.Fprintf(w, "Hello, ನಮ್ಮ ಬೆಂಗಳೂರು!!\n")
}

func main() {
    fmt.Println("Visit http://localhost:8080/hello")
    http.HandleFunc("/hello", HelloWeb)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
