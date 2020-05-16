package main

import (
	"net/http"

	"github.com/stevedesilva/golang-microservices/controllers"
)

// StartApp entry point
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

	// http.HandleFunc("/users", func(writer http.ResponseWriter, request *http.Request) {
	// 	writer.Write([]byte("hello world!\n"))
	// })

	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	panic(err)
	// }

}
