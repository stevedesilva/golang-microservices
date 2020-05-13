package main

import "net/http"
/*
 create binary from main.go:
 go build -o webserver


 run server:
  ./webserver

 call endpoint:
 curl -v localhost:8080/hello

*/
func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world!\n"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
