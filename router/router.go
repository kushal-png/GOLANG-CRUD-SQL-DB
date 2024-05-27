package router

import (
	"goserver/controller"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {

	r := mux.NewRouter()
	getHandlers(r)
	return r
}

func getHandlers(r *mux.Router) {
	r.HandleFunc("/user/{id}", controller.GetUser).Methods("GET")
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/user", controller.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user/{id}", controller.UpdateUser).Methods("PUT")
}
