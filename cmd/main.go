package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/prensmatt/go-bookstore/pkg/routes"
)

func main(){
	route := mux.NewRouter()
	routes.RegisterBookStoreRoutes(route)
	http.Handle("/",route)
	log.Fatal(http.ListenAndServe("localhost:9010", route))
}