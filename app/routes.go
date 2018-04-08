package app

import (
"github.com/gorilla/mux"
_"gopkg.in/mgo.v2"
_"gopkg.in/mgo.v2/bson"
_"time"
_"log"
)
// NewRouter ...
func AllRoutes() *mux.Router {
	//Create main router
	r := mux.NewRouter().StrictSlash(true)

	/**
	 * Routes
	 */

	r.Methods("GET").Path("/").HandlerFunc(HelloWorld)
	r.Methods("GET").Path("/getItems").HandlerFunc(GetAllItems)
	r.Methods("GET").Path("/getItems/{id}").HandlerFunc(GetItemById)
	r.Methods("POST").Path("/addItems").HandlerFunc(AddItems)
	r.Methods("PUT").Path("/putItems").HandlerFunc(UpdateItems)
	r.Methods("DELETE").Path("/deleteItems/{id}").HandlerFunc(DeleteItems)


	return r
}