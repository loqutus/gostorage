package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

func save_file(name string, content []byte) error {
	err := ioutil.WriteFile("data/"+name, content, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func read_file(name string) ([]byte, error) {
	data, err := ioutil.ReadFile("data/" + name)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return data, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi. This is storage")
	})

	http.ListenAndServe(":8888", nil)
}
