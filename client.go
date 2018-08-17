package main

import (
	"flag"
	"fmt"
	"os"
	"net/http"
)

func download(server string, port string, filename string) Error {
	resp, err := http.Get("http://" + *server + ":" + *port + "/download/" + filename)
	if err != nil {
		fmt.Println("get error")
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filename)
	if err != nil{
		fmt.Println("file write error")
		return err
	}


}

func main() {
	server := flag.String("server", "localhost", "server to connect to")
	port := flag.String("port", "8888", "port to connect to")
	action := flag.String("action", "", "action to do: upload or download")
	file := flag.String("file", "", "file name to use")
	flag.Parse()
	if *action == "" {
		fmt.Println("missing action")
		os.Exit(1)
	}
	if *action != "upload" && *action != "download" {
		fmt.Println("wrong action")
	}
	if *file == "" {
		fmt.Println("missing file")
		os.Exit(1)
	}
	if *action == "download" {
		err := download(*server, *port, *file)
		if err := nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
