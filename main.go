package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		println(request.Header.Get("key"))
		writer.Write([]byte("hello"))
	})
	err := http.ListenAndServe("127.0.0.1:8001",nil)
	if err!=nil {
		fmt.Println("网络错误",err.Error())
		return
	}
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<p style='color:red;'>hello</p>"))
}
