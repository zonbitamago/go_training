package main

import (
	"fmt"
	"go_training/03_webapi/handlers"
	"go_training/03_webapi/logger"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	router.GET("/", logger.Logging(handlers.HelloIndex, "index"))

	port := ":8080"

	fmt.Println("start!", "http://localhost"+port)
	log.Fatal(http.ListenAndServe(port, router))
}
