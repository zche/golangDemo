package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	n, err := io.WriteString(w, "Hello,world!")
	if err != nil {
		log.Fatal("错误码:", n)
	}
	//log.Fatal("test:")
}
func main() {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
