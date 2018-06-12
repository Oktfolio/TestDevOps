package main

import (
	"net/http"
	"io"
	"os/exec"
	"log"
)

func main() {

	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":5000", nil)
}

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello, this is my first page!</h>")
	reLaunch()
}

func reLaunch() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}
