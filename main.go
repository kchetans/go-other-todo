package main

import(
	"github.com/gorilla/mux"
	"github.com/go-to-do/app"
	"net/http"
	"time"
	"log"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func main(){

	err := godotenv.Load()
	if err != nil {
		logrus.Info("Error loading .env file")
	}
	r := mux.NewRouter()
	r.HandleFunc("/", app.HelloWorld)
	r.Methods("GET").Path("/getItems").HandlerFunc(app.GetAllItems)
	r.Methods("GET").Path("/getItem/{id}").HandlerFunc(app.GetItemById)
	r.Methods("POST").Path("/addItems").HandlerFunc(app.AddItems)
	r.Methods("PUT").Path("/putItems").HandlerFunc(app.UpdateItems)
	r.Methods("DELETE").Path("/deleteItems/{id}").HandlerFunc(app.DeleteItems)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logrus.Info("Starting Server on http://localhost:", os.Getenv("PORT"))

	log.Fatal(srv.ListenAndServe())
}