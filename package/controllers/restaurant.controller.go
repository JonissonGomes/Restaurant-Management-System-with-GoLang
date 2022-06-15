package controllers

import (
	"enconding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JonissonGomes/Restaurant-Management-System-with-GoLang/package/entities"
	"github.com/gorilla/mux"
)

var NewRestaurant entities.Restaurant

func GetRestaurants(w http.ResponseWriter, r *http.Request) {
	restaurants := entities.GetAllRestaurant()
	res, _ := json.Marshal(restaurants)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRestaurantById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	restaurantId, _ := params["restaurantId"]
	ID, err := strconv.ParseInt(restaurantId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		return
	}
	restaurantDetails, _ := entities.GetRestaurantById(int64(ID))

	res, _ := json.Marshal(restaurantDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
