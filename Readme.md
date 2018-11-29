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

- posts 创建

```
请求 curl --header "jwt:11f34a2a-f394-11e8-afa9-9801a7a96639" -X POST -d"title=测试M&content=此次AWS发布的AI芯片主要通过其云业务向外提供服务，并不直接出售这些芯片，因此，Andy口中的这一“行业颠覆者”并不会对芯片巨头英特尔和英伟达构成直接威胁。Andy表示，客户可以通过AWS的EC2 instances、Amazon SageMaker和Elastic Inference等产品获得这一服务。与其他AWS服务一样，用户可以根据使用量付费。&author_id=1" http://localhost:8001/posts

>{
    "result":"success",
    "message":"",
    "data":{
        "author_id":1,
        "author_name":"",
        "content":"此次AWS发布的AI芯片主要通过其云业务向外提供服务，并不直接出售这些芯片，因此，Andy口中的这一“行业颠覆者”并不会对芯片巨头英特尔和英伟达构成直接威胁。Andy表示，客户可以通过AWS的EC2 instances、Amazon SageMaker和Elastic Inference等产品获得这一服务。与其他AWS服务一样，用户可以根据使用量付费。",
        "id":25,
        "title":"测试M",
        "updated_at":"2018-11-29T17:45:09.474123+08:00"
    }
}

```