package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	//ServerPort 端口号
	ServerPort = 7777
	//ServerDomain 域名
	ServerDomain = "localhost"
	//ResponseTemplate 响应内容
	ResponseTemplate = "hello"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(ResponseTemplate)))
	w.Write([]byte(ResponseTemplate))
}
func main() {
	http.HandleFunc(fmt.Sprintf("%s:%d/", ServerDomain, ServerPort), rootHandler)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", ServerPort), "rui.crt", "rui.key", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS:", err.Error())
	}
}
