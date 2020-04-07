package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("public" + r.URL.Path)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}

		defer f.Close()
		io.Copy(w, f)
	})
	http.ListenAndServe(":8000", nil)
}
