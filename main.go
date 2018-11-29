package main

import (

	"api.clublog.com/libs/configs"
	_ "api.clublog.com/libs/models"
	"net/http"
	"api.clublog.com/libs/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	configs.Init()
	router := httprouter.New()
	router.POST("/sessions/new", controllers.NewSession)
	router.GET("/posts", controllers.ListAllPosts)
	router.GET("/posts/:id", controllers.ShowPost)
	router.POST("/posts", controllers.CreatePost)

	http.ListenAndServe(":8001", router)
}
