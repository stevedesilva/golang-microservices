package app

import "github.com/stevedesilva/golang-microservices/mvc/controllers"

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
