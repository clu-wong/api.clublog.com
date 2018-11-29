package configs

import (
	"github.com/pelletier/go-toml"
	"github.com/jinzhu/gorm"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var Config *toml.Tree
var Db *gorm.DB
var RedisClient *redis.Client
var SessionExpired time.Duration

func Init(){
	var err error
	Config, err = toml.LoadFile("config.toml")
	if err != nil{
		panic("Config Init error: "+ err.Error())
	}

	envConfig := Config.Get("development").(*toml.Tree)
	Db, err = gorm.Open("mysql", envConfig.Get("mysql_config").(string))
	if err != nil{
		panic("Db Init error: "+ err.Error())
	}
	expire_val := envConfig.Get("session_expire").(int64)
	SessionExpired = time.Duration(expire_val)

	RedisClient = redis.NewClient(&redis.Options{
		Addr: envConfig.Get("redis_config").(string),
		Password: "",
		DB: 0})

}