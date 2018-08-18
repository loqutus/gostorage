package main

import (
	"net/http"
	"fmt"
	"bytes"
	"strings"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	file, header, err := r.FormFile(filename)
	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()

}

func downloadHandler(w http.ResponseWriter, r *http.Request) {

}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi. This is storage")
	})
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/download", downloadHandler)
	http.ListenAndServe(":8888", nil)
}
