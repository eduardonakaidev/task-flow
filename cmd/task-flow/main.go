package main

import "net/http"

func main() {
	http.HandleFunc("POST /api/task", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":3000", nil)
}
