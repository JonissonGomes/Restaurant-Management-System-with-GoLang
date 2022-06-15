package routes

import (
	"github.com/JonissonGomes/Restaurant-Management-System-with-GoLang/package/controllers"
	"github.com/gorilla/mux"
)

var RegisterRestaurantRoutes = func(router *mux.Router) {
	router.HandleFunc("/restaurants/", controllers.CreateRestaurant).Methods("POST")
	router.HandleFunc("/restaurants/", controllers.GetRestaurants).Methods("GET")
	router.HandleFunc("/restaurants/{id}", controllers.GetRestaurantById).Methods("GET")
	router.HandleFunc("/restaurants/{id}", controllers.UpdateRestaurant).Methods("PUT")
	router.HandleFunc("/restaurants/{id}", controllers.DeleteRestaurant).Methods("DELETE")
}
