package customercontroller

import (
	"goweb/models/customermodel"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	customers := customermodel.GetAll()
	data := map[string]interface{}{
		"customers": customers,
	}

	temp, err := template.ParseFiles("views/customer/test.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
