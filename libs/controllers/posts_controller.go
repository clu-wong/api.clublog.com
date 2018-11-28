package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"api.clublog.com/libs/configs"
	"api.clublog.com/libs/models"
	_ "encoding/json"
)

func Posts(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	switch r.Method {
	case "GET":
		if ps.ByName("id") != ""{
			ShowPost(w, r, ps)
		}else{
			ListAllPosts(w, r, ps)
		}
	case "POST":
		CreatePost(w, r, ps)
	case "DELETE":
		DeletePost(w, r, ps)
	case "PUT":
		UpdatePost(w, r, ps)
	default:
		ListAllPosts(w, r, ps)
	}
}

func ListAllPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	var posts []models.Post
	configs.Db.Model(models.Post{}).Find(&posts)
	result := ResultData2Json{200,"success", "", posts}
	Render2Json(w, r, &result)
}

func ShowPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

}
func CreatePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

}

func UpdatePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

}

func DeletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

}
