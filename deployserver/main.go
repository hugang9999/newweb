package main

import (
	"io"
	"net/http"
	"os/exec"
	"log"
)

func reLaunch() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	print(err)
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":5000", nil)

}
func firstPage(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "<h1> Hello,this is deployserver 5000!</h1>")
	reLaunch()
}
