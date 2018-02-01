package main

import (
	"net/http"
	"path"
	"os"
	"fmt"
)

func main() {
	port := "8080"
	fmt.Println("listening on ", port)
	panic(http.ListenAndServe(":"+port, http.FileServer(http.Dir(pwd()))))
}

func pwd() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir := path.Dir(ex)
	return dir
}
