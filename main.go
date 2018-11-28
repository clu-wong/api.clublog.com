package main

import (

	_ "clublog/lib/models"

	"api.clublog.com/libs/configs"
	"net/http"
	"api.clublog.com/libs/controllers"
)

func main() {
	configs.Init()
	http.HandleFunc("/sessions/new", controllers.NewSession)
	//router := httprouter.New()
	//router.GET("/posts/:id", controllers.Posts)

	http.ListenAndServe(":8001", nil)
}
