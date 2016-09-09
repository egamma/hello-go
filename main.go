package main

import (
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    var msg = "Hello C&E from VS Code and Go!"

}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
