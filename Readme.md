# clu Blog

登陆，登出，发表信息，查看信息

API接口

- 登陆
```
请求: curl -X POST -d "username=xiaowang&pwd=123456" http://localhost:8001/sessions/new

返回: {"result":"success","message":"ok","data":{"jwt":"0255e7f8-f376-11e8-83c2-9801a7a96639"}}
```


- 发布post