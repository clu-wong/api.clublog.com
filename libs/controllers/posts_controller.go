package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"api.clublog.com/libs/configs"
	"api.clublog.com/libs/models"
	_ "encoding/json"
	"api.clublog.com/libs/utils"
	"fmt"
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
	fmt.Println(r.Header)
	if !CheckJwt(r){
		Render400(w, r)
		return
	}
	var posts []models.Post
	var counts int
	r.ParseForm()
	page := utils.ParseInt(r.Form.Get("page"))

	configs.Db.Model(models.Post{}).Count(&counts)
	configs.Db.Model(models.Post{}).Limit(20).Offset(20*(page-1)).Find(&posts)
	trans_rows := models.MappingPostsRows(posts, 20)
	result := ResultData2Json{200,"success", "", models.ListData(trans_rows, counts, page)}
	Render2Json(w, r, &result)
}

func ShowPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	if CheckJwt(r){
		var post models.Post
		var user models.User
		configs.Db.Model(models.Post{}).Where("id = ?", ps.ByName("id")).Find(&post)//.Related(&user)
		data := post.JsonMap(user)
		result := ResultData2Json{200,"success", "", data}
		Render2Json(w, r, &result)
	}else{
		Render400(w, r)
	}
}
func CreatePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	if CheckJwt(r){
		var post models.Post
		var data interface{}
		r.ParseForm()
		post.New(r.Form)
		err := configs.Db.Create(&post)
		if err.Error != nil{
			data = err.Error
		}else{
			data = post.JsonMap(models.User{})
		}
		result := ResultData2Json{200,"success", "", data}
		Render2Json(w, r, &result)
	}else{
		Render400(w, r)
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

}

func DeletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params){

}
