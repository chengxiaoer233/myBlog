vue + golang + mysql 实现一个博客

* 目标：主要是用来自己总结和学习相关技术

* [原博客链接🔗](https://github.com/wejectchen/ginblog)

<div align="center">

# myBlog

</div>

<div align="center">

![](https://img.shields.io/github/languages/code-size/chengxiaoer233/myBlog?label=CodeSize)
![](https://img.shields.io/github/stars/chengxiaoer233/myBlog?label=GitHub)
![](https://img.shields.io/github/watchers/chengxiaoer233/myBlog?label=Watch)
[![Go Report Card](https://goreportcard.com/badge/github.com/chengxiaoer233/myBlog)](https://goreportcard.com/report/github.com/chengxiaoer233/myBlog)
[![LICENSE](https://img.shields.io/badge/license-MIT-green)](https://mit-license.org/)

</div>

<div align="center">

<img  src="https://my-source666.obs.cn-south-1.myhuaweicloud.com/myBlog/golang-jixiangwu-image.png" width="600" height="350"/>

</div>

#### p1
- 修改1：重新梳理文件目录结构，更适合工程化，增加编译、启动脚本
- 修改2：重新定义一些数据结构、增加异常处理

#### p2
- 完善mysql数据库相关配置,数据库定义个操作相分离，model为数据结构定义层，dao为操作层

#### p3、p4、p5
- 完善用户模块相关接口，采用面向对象的方式来阻止代码结构
- 完善用户新增用户时密文存储用户密码

#### p6
- 删除用户改为 delete + json body，需要校验密码，同时需要校验要用户是否已经存在 
- 完善各种异常或错误逻辑，满足一个线上代码的要求
- 用户编辑接口，增加校验逻辑，校验用户名是否存在及用户id是否匹配一致

## 接口测试方法
- 用户管理
- 1: 新增用户 AddUser
```shell script
Path: /api/v1/user/add
Method: post 
Header: "Content-Type:application/json"

payload示例:
  {"userName":"666","password":"666","role":0}

curl请求示例: 
  curl -X POST "localhost:3000/api/v1/user/add" -d '{"userName":"test666","password":"test666","role":0}' -H "Content-Type:application/json"

返回示例:
{"status":200,"msg":"OK","data":null}

{"status":1001,"msg":"用户名已存在！","data":null}
```

- 2：删除单个用户  DeleteUser
```shell script
Path: /api/v1/admin/user/
Method: delete 
Header: "Content-Type:application/json"

payload示例:
  {"id":1,"password":"666"}

curl请求示例: 
curl -X DELETE localhost:3000/api/v1/admin/user/ -d '{"id":1,"password":"666"}' -H "Content-Type:application/json"

返回示例:
{"status":1006,"msg":"TOKEN不正确,请重新登陆","data":null}%

{"status":200,"msg":"ok","data":null}
```

- 3：修改用户信息
```shell script
Path: api/v1/admin/user
Method: PUT 
payload示例:
  {"id":1,"userName":"666"}

curl请求示例
curl -X PUT localhost:3000/api/v1/admin/user/ -d '{"id":1,"userName":"666","role":2}' -H "Content-Type:application/json"

返回结果示例
{"status":1001,"msg":"用户名已存在！","data":null}

{"status":200,"msg":"OK","data":null}
```

- 4：查询单个用户信息 GetOneUserInfo
```shell script
Path: /api/v1/user/:id
Method: get 

curl请求示例
curl localhost:3000/api/v1/user/1

返回结果示例
{"status":200,"msg":"OK","data":{"ID":1,"CreatedAt":"***","UpdatedAt":"***",
"DeletedAt":null,"userName":"test","password":"********","role":0}}

{"status":1003,"msg":"用户不存在","data":null}
```

- 5：查询所有用户 GetUsers
```shell script
Path: /api/v1/user/users
Path: /api/v1/admin/user/users
Method: get 

curl请求带参数示例
curl "localhost:3000/api/v1/user/users?pageSize=1&pageNum=1&userName=test"
curl "localhost:3000/api/v1/admin/user/users?pageSize=1&pageNum=1&userName=test"

curl请求不带参数示例
curl "localhost:3000/api/v1/user/users"

返回示例: 
{"data":[],"msg":"OK","status":200,"total":0}

{"data":[{"ID":1,"CreatedAt":"***","UpdatedAt":"***","DeletedAt":null,"userName":"666",
"password":"***","role":0}],"msg":"OK","status":200,"total":1}
```

- 文章分类
- 1：新增文章分类
```shell script
Path: api/v1/admin/category/add
Method: POST 
payload示例:
  {"name":"666"}

curl示例
curl -X POST localhost:3000/api/v1/admin/category/add -d '{"name":"666"}' -H "Content-Type:application/json"

返回示例
{"status":200,"msg":"OK","data":null}

{"status":3001,"msg":"该分类已存在","data":null}
```
- 2：删除文章分类
```shell script
Path: api/v1/admin/category/:ID
Method: DELETE 

curl示例
curl -X DELETE localhost:3000/api/v1/admin/category/1

返回示例
{"status":200,"msg":"OK","data":null}

{"status":3002,"msg":"该分类不存在","data":null}
```
- 3：编辑文章分类
```shell script
Path: api/v1/admin/category/:ID
Method: PUT 

curl示例
curl -X PUT localhost:3000/api/v1/admin/category/1 -d '{"name":"666"}' -H "Content-Type:application/json"

返回示例
{"status":200,"msg":"OK","data":null}

{"status":3002,"msg":"该分类不存在","data":null}
```

- 4：返回某个分类
```shell script
Path: api/v1/category/:ID
Method: GET 

curl示例
curl localhost:3000/api/v1/category/1

返回示例
{"status":200,"msg":"OK","data":{"ID":2,"CreatedAt":"***","UpdatedAt":"***","DeletedAt":null,"name":"666"}}

{"status":3002,"msg":"该分类不存在","data":null}
```

- 5：返回所有分类
```shell script
Path: api/v1/category/
Path: api/v1/admin/category/
Method: GET 

curl示例
curl localhost:3000/api/v1/category/
curl localhost:3000/api/v1/admin/category/

返回示例
{"data":[{"ID":2,"CreatedAt":"***","UpdatedAt":"***","DeletedAt":null,"name":"666"}],"msg":"OK","status":200,"total":1}

{"data":[],"msg":"OK","status":200,"total":0}
```