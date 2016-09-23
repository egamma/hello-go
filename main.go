package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	var msg = "<h1>Happy Anniversary VIS and IAETH!</h1>"
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
