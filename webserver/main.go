package main

import "io"
import "net/http"

func main()  {
	http.HandleFunc("/",firstPage)
	http.ListenAndServe(":8000", nil)

}
func firstPage(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer,"<h1> Hello,this is my first page!</h1>")
}
