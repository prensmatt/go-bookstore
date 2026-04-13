package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/prensmatt/go-bookstore/pkg/routers"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	log.Printf("Starting server at port 9010\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}