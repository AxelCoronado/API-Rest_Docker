package main

import (
	"net/http"

	"github.com/AxelCoronad/go-restapi/db"
	"github.com/AxelCoronad/go-restapi/models"
	"github.com/AxelCoronad/go-restapi/routes"
	"github.com/gorilla/mux"
)

func main(){
	db.DBconnection()
	db.DB.AutoMigrate(models.IceCream{})

	r := mux.NewRouter()

	r.HandleFunc("/icecreams", routes.PostIceCream).Methods("POST")
	r.HandleFunc("/icecreams", routes.GetAll).Methods("GET")
	r.HandleFunc("/icecreams/{id}", routes.GetIceCream).Methods("GET")
	r.HandleFunc("/icecreams/{id}", routes.UpdateIceCream).Methods("PUT")
	r.HandleFunc("/icecreams/{id}", routes.DeleteIceCream).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}