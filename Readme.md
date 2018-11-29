# clu Blog

登陆，登出，发表信息，查看信息

API接口

- 登陆
```
请求: curl -X POST -d "username=xiaowang&pwd=123456" http://localhost:8001/sessions/new

> {"result":"success","message":"ok","data":{"jwt":"4ac7b248-f392-11e8-8ebb-9801a7a96639"}}
```


- posts列表
```
curl --header "jwt:11f34a2a-f394-11e8-afa9-9801a7a96639" -X GET http://localhost:8001/posts
>{
    "result": "success",
    "message": "",
    "data": {
        "current": 1,
        "list": [
            {
                "author_id": 1,
                "author_name": "xiaowang",
                "content": "111111111111111111111111111111111111111111111111111111111111111111111111",
                "id": 1,
                "title": "helloworld",
                "updated_at": "2018-10-19T00:00:00+08:00"
            }
        ],
        "total": 1
    }
}
```

- posts详情


```
请求：curl --header "jwt:11f34a2a-f394-11e8-afa9-9801a7a96639" -X GET http://localhost:8001/posts/1
> {
      "result": "success",
      "message": "",
      "data": {
          "author_id": 1,
          "author_name": "",
          "content": "111111111111111111111111111111111111111111111111111111111111111111111111",
          "id": 1,
          "title": "helloworld",
          "updated_at": "2018-10-19T00:00:00+08:00"
      }
  }
```
