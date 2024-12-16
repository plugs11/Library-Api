package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"manan.tola/routes"
)

func main() {
	router := gin.Default()
	file, _ := os.Create("logfile.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	router.Use(gin.Logger())
	routes.RegisterBookRoutes(router) // passes the routing data ie. the url request to the routes.go file
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":4040", router))
} //creating the route and provinding the localhost
