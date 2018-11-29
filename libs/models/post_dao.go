package models

import (
	"github.com/jinzhu/gorm"
	"api.clublog.com/libs/configs"
)

type Post struct{
	gorm.Model
	Title string `gorm:"type:varchar(100);unique_index"`
	Content string `gorm:"type:text"`
	UserId int
	User User
}


func (post Post) JsonMap(author User) map[string]interface{}{
	var data_mapping = make(map[string]interface{}, 6)
	data_mapping["id"] = post.ID
	data_mapping["title"] = post.Title
	data_mapping["content"] = post.Content
	data_mapping["updated_at"] = post.UpdatedAt
	data_mapping["author_id"] = post.UserId
	data_mapping["author_name"] = author.Name
	return data_mapping
}

func (post Post) user() User{
	var user User
	configs.Db.Where("id = ? ", post.UserId).First(&user)
	return user
}

func ListData(data_mapping interface{}, total int, current int) map[string]interface{}{
	var data = make(map[string]interface{})
	data["list"] = data_mapping
	data["total"] = total
	data["current"] = current
	return data
}

func MappingPostsRows(posts []Post, count int) interface{}{
	var data = make([]interface{},len(posts),count)
	for i :=0; i < len(posts); i++ {
		data[i] = posts[i].JsonMap(posts[i].user())
	}
	return data
}