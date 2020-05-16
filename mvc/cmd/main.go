package main

// import "net/http"

import (
	"github.com/stevedesilva/golang-microservices/app"
)

func main() {
	app.StartApp()
}

// func main() {
// 	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
// 		writer.Write([]byte("hello world!\n"))
// 	})

// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		panic(err)
// 	}

// }
