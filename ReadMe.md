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

## 接口测试方法

- 1:新增用户 AddUser
```shell script
curl -X POST "localhost:3000/api/v1/user/add" -d '{"userName":"test666","password":"test666","role":0}' -H "Content-Type:application/json"
```

- 2：查询所有用户 GetUsers
```shell script
带参数
curl "localhost:3000/api/v1/user/users?pageSize=1&pageNum=1&userName=test"

不带参数
curl "localhost:3000/api/v1/user/users"
```

