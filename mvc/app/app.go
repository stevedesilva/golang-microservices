package app

import (
	"net/http"

	"github.com/stevedesilva/golang-microservices/mvc/controllers"
)

// StartApp entry point
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
