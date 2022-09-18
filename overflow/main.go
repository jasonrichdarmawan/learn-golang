package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/overflow", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(1)
		if err != nil {
			http.Error(w, "stack overflow", http.StatusInternalServerError)
			return
		}
		filesHeader := r.MultipartForm.File["files"]
		_ = filesHeader
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8080", nil)
}
