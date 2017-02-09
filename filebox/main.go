package main

import (
	"github.com/qor/filebox"
	"net/http"

	"fmt"
)

func main() {
	mux := http.NewServeMux()

	Filebox := filebox.New("./downloads")
	Filebox.MountTo("/downloads", mux)

	port := ":9090"
	fmt.Println("will listen port ", port)
	fmt.Println("download by cmd: curl http://localhost:9090/downloads/test.txt")
	http.ListenAndServe(port, mux)
}
