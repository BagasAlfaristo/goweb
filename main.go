package main

import (
	//"goweb/config"
	"goweb/config"
	"goweb/controllers/customercontroller"
	"goweb/models/customermodel"
	"net/http"

	//"goweb/controllers"
	"log"
)

func main() {
	config.Dbconnection()
	//customercontroller.AddCustomer()
	//http.HandleFunc("/", customercontroller.Index)
	http.HandleFunc("/", customercontroller.Index)
	http.HandleFunc("/addcustomer", customercontroller.FormHandler)
	http.HandleFunc("/add", customermodel.AddHandler)

	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
