package main

import(
	"github.com/gorilla/mux"
	"github.com/go-to-do/app"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/products/{key}", HelloWorld)
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)


}