package main

import (
	"log"
	"net/http"

	"github.com/jevitapearl/go_webProj/goBookInventory/pkg/config"
	"github.com/gorilla/mux"
	"github.com/jevitapearl/go_webProj/goBookInventory/pkg/routes"
	"github.com/jinzhu/gorm/dialect/mysql"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080",r))
}