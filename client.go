package main

import (
	"flag"
	"fmt"
	"os"
	"net/http"
	"io"
	"errors"
	"io/ioutil"
)

func download(server *string, port *string, filename *string) error {
	if _, err := os.Stat(*filename); err == nil {
		return errors.New("file already exists")
	}
	resp, err := http.Get("http://" + *server + ":" + *port + "/download/" + *filename)
	if err != nil {
		fmt.Println("get error")
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(*filename)
	if err != nil {
		fmt.Println("file create error")
		return err
	}
	defer out.Close()
	io.Copy(out, resp.Body)
	return nil
}

func upload(server *string, port *string, filename *string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return errors.New("file does no exist")
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("file open error")
		return err
	}
	defer file.Close()
	res, err := http.Post("http://"+*server+":"+*port+"/upload/"+*filename, "binary/octet-stream", file)
	if err != nil {
		fmt.Println("post error")
		return err
	}
	defer res.Body.Close()
	message, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("")
		return err
	}
	fmt.Println(message)
	return nil
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
		err := download(server, port, file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	if *action == "upload" {
		err := upload(server, port, file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
