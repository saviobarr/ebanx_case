package app

import (
	"ebanx_case/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//StartApp Start listener and routers
func StartApp() {

	r := mux.NewRouter()

	r.HandleFunc("/balance", controllers.GetBalance).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
