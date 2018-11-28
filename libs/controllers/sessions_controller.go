package controllers

import (
	"net/http"
	_ "api.clublog.com/libs/configs"
	"api.clublog.com/libs/configs"
	"clublog/lib/models"
	"github.com/satori/go.uuid"
	"fmt"
	"encoding/json"
)

type ResultData2Json struct{
	status_code int
	Result string `json:"result"`// success | fail
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
func Render2Json(w http.ResponseWriter, r *http.Request, resp_data *ResultData2Json){
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if resp_data != nil{
		w.WriteHeader(resp_data.status_code)
	}

	json.NewEncoder(w).Encode(resp_data)
}

func CheckJwt(r *http.Request) bool{
	jwt := r.Header.Get("jwt")
	if jwt == ""{
		return false
	}
	result := configs.RedisClient.Get(jwt)
	if jwt != "" && result.Val() != ""{
		configs.RedisClient.Set(jwt, result.Val(), configs.SessionExpired)
		return true
	}else{
		return false
	}
}

func FilterProtect(w http.ResponseWriter, r *http.Request){
	if !CheckJwt(r){
		var result ResultData2Json
		result= ResultData2Json{400, "fail", "token error",nil}
		Render2Json(w, r, &result)
		return
	}
}


func NewSession(w http.ResponseWriter, r *http.Request){
	var result ResultData2Json
	if r.Method != "POST"{
		result= ResultData2Json{400, "fail", "access denied",nil}
		Render2Json(w, r, &result)
		return
	}
	r.ParseForm()
	var user models.User
	var jwt string
	configs.Db.Where("name =?", r.Form.Get("username")).First(&user)
	if user.Password == r.Form.Get("pwd"){
		jwt = uuid.Must(uuid.NewV1()).String()
		configs.RedisClient.Set(jwt, user.ID, configs.SessionExpired)
		result = ResultData2Json{200,"success", "ok", map[string]interface{}{"jwt": jwt}}
	}else{
		result = ResultData2Json{200,"fail", "password error", map[string]interface{}{}}

	}
	fmt.Println(result)
	Render2Json(w, r, &result)
}
