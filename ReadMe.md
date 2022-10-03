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

### p7
- 完善文章分类模块的编写，并实现工程化

### p8
- 完善文章相关接口,重新组织readme文档，增加接口调用相关规则
&emsp; 

## 接口测试方法
- **用户管理相关接口**

* 1: **新增用户** ```AddUser```
    
    - Path:   ```/api/v1/user/add```
    - Method: ```POST```
    - payload示例
    ```shell script
     {
         "userName": "666",
         "password": "666",
         "role": 0
     }
    ```
    &emsp; 
    
    -  curl请求示例
    ```shell script
      curl -X POST "localhost:3000/api/v1/user/add" -d '{"userName":"test666","password":"test666","role":0}' -H "Content-Type:application/json"
    ```
    &emsp; 
    
    - 返回示例
    ```json
      {
        "status": 200,
        "msg": "OK",
        "data": null
      }
    ```
    &emsp; 
    
* 2: **删除单个用户** ```DeleteUser```
    
    - Path:   ```/api/v1/admin/user/```
    - Method: ```DELETE```
    - Header: ```"Content-Type:application/json"```
    - payload示例
    ```shell script
     {
     	"id": 1,
     	"password": "666"
     }
    ```
    &emsp; 
  
    - curl请求示例
    ```shell script
      curl -X DELETE localhost:3000/api/v1/admin/user/ -d '{"id":1,"password":"666"}' -H "Content-Type:application/json"
    ```
    &emsp; 
    
    - 返回示例
    ```json
    {
        "status": 200,
        "msg": "OK",
        "data": null
    }
    ```
    &emsp; 
    
* 3: **修改用户信息** ```EditUser```

    - Path:   ```api/v1/admin/user```
    - Method: ```PUT```
    - payload示例
    ```shell script
     {
     	"id": 1,
     	"password": "666"
     }
    ```
    &emsp; 
  
    - curl请求示例
    ```shell script
      curl -X PUT localhost:3000/api/v1/admin/user/ -d '{"id":1,"userName":"666","role":2}' -H "Content-Type:application/json"
    ```
    &emsp;
     
    - 返回示例
    ```json
    {
        "status": 200,
        "msg": "OK",
        "data": null
    }
    ```  
    &emsp; 
    
 * 4: **查询单个用户信息** ```GetOneUserInfo```
     
     - Path:   ```/api/v1/user/:id```
     - Method: ```GET```
     &emsp; 
     
     - curl请求示例
     ```shell script
       curl localhost:3000/api/v1/user/1
     ```
     &emsp;
     
     - 返回示例
     ```json
     {
     	"status": 200,
     	"msg": "OK",
     	"data": {
     		"ID": 1,
     		"CreatedAt": "***",
     		"UpdatedAt": "***",
     		"DeletedAt": null,
     		"userName": "test",
     		"password": "********",
     		"role": 0
     	}
     }
    ```
    &emsp; 
   
 * 5: **查询所有用户** ```GetUsers```
     
     - Path:   ```/api/v1/user/users```
     - Path:   ```/api/v1/admin/user/users```
     - Method: ```GET```
     - Query参数：```pagesize,pagenum```
     &emsp; 
     
     - curl请求示例
     ```shell script
       curl "localhost:3000/api/v1/user/users?pageSize=1&pageNum=1&userName=test"
       curl "localhost:3000/api/v1/admin/user/users?pageSize=1&pageNum=1&userName=test"
       curl "localhost:3000/api/v1/user/users"
     ```
     &emsp; 
     
     - 返回示例
     ```json
     {
     	"data": [{
     		"ID": 1,
     		"CreatedAt": "***",
     		"UpdatedAt": "***",
     		"DeletedAt": null,
     		"userName": "666",
     		"password": "***",
     		"role": 0
     	}],
     	"msg": "OK",
     	"status": 200,
     	"total": 1
     }
    ```   
   &emsp; 
   
- **文章分类相关接口**

* 1: **新增文章分类** ```AddCategory```
    
    - Path:   ```api/v1/admin/category/add```
    - Method: ```POST```
    - payload示例
    ```shell script
     {
         "name": "666"
     }
    ```
    &emsp; 
    
    - curl请求示例
    ```shell script
      curl -X POST localhost:3000/api/v1/admin/category/add -d '{"name":"666"}' -H "Content-Type:application/json"
    ```
    &emsp; 
    
    - 返回示例
    ```json
    {
        "status": 200,
        "msg": "OK",
        "data": null
    }
    ```
    &emsp; 
    
* 2: **删除文章分类** ```DeleteCate```
    
    - Path:   ```api/v1/admin/category/:id```
    - Method: ```DELETE```
    &emsp;
    
    - curl请求示例
    ```shell script
      curl -X DELETE localhost:3000/api/v1/admin/category/1
    ```
    &emsp;
    
    - 返回示例
    ```json
    {
        "status": 200,
        "msg": "OK",
        "data": null
    }
    ``` 
    &emsp;
    
* 3: **编辑文章分类** ```EditCate```
    
    - Path:   ```api/v1/admin/category/:ID```
    - Method: ```PUT```
    - payload示例
    ```shell script
     {
         "name": "666"
     }
    ```
    &emsp;
    
    - curl请求示例
    ```shell script
      curl -X PUT localhost:3000/api/v1/admin/category/1 -d '{"name":"666"}' -H "Content-Type:application/json"
    ```
    &emsp;
    
    - 返回示例
    ```json
    {
        "status": 200,
        "msg": "OK",
        "data": null
    }
    ``` 
    &emsp;
    
* 4: **返回某个分类** ```GetCateInfo```
    
    - Path:   ```api/v1/category/:id```
    - Method: ```GET```
    &emsp;
    
    - curl请求示例
    ```shell script
      curl localhost:3000/api/v1/category/1
    ```
    &emsp;
    
    - 返回示例
    ```json
    {
    	"status": 200,
    	"msg": "OK",
    	"data": {
    		"ID": 2,
    		"CreatedAt": "***",
    		"UpdatedAt": "***",
    		"DeletedAt": null,
    		"name": "666"
    	}
    }
    ``` 
    &emsp;
    
* 5: **返回所有分类** ```GetCateS```
    
    - Path:   ```api/v1/category/```
    - Path:   ```api/v1/admin/category/```
    - Method: ```GET```
    - Query参数: ```pagesize,pagenum```  
    &emsp;
    
    - curl请求示例
    ```shell script
      curl localhost:3000/api/v1/category/
      curl localhost:3000/api/v1/admin/category/
    ```
    &emsp;
  
    - 返回示例
    ```json
    {
    	"data": [{
    		"ID": 2,
    		"CreatedAt": "***",
    		"UpdatedAt": "***",
    		"DeletedAt": null,
    		"name": "666"
    	}],
    	"msg": "OK",
    	"status": 200,
    	"total": 1
    }
    ``` 
    &emsp; 
  

- **文章相关接口**

* 1: **新建文章** ```AddArticle```
    
    - Path:   ```api/v1/admin/article/add```
    - Method: ```POST```
    - payload示例
    ```shell script
     {
     	"title": "test",
     	"cid": 1,
     	"desc": "test",
     	"content": "666",
     	"img": "test",
     	"commentCount": 666,
     	"readCount": 666
     }
    ```
    &emsp;
    
    - curl请求示例
    ```shell script
      curl -X POST localhost:3000/api/v1/admin/article/add -d '{"title":"test","cid":1,"desc":"test","content":"test666","img":"test","commentCount":6,"readCount":6}' -H "Content-Type:application/json"
    ```
    &emsp;
    
    - 返回示例
    ```json
   {
   	"status": 200,
   	"msg": "OK",
   	"data": {
   		"Category": {
   			"ID": 0,
   			"CreatedAt": "***",
   			"UpdatedAt": "***",
   			"DeletedAt": null,
   			"name": ""
   		},
   		"ID": 4,
   		"CreatedAt": "***",
   		"UpdatedAt": "***",
   		"DeletedAt": null,
   		"title": "test",
   		"cid": 1,
   		"desc": "test",
   		"content": "666",
   		"img": "test",
   		"commentCount": 666,
   		"readCount": 666
   	}
   }
  ```
  &emsp;
    
* 2: **删除文章** ```DeleteArticle```
    
    - Path:   ```api/v1/admin/article/:id```
    - Method: ```DELETE```
    &emsp;
    
    - curl请求示例
     ```shell script
      curl -X DELETE localhost:3000/api/v1/admin/article/1
    ```
    &emsp;
    
    - 返回示例
    ```json
    {
        "status": 200,
        "msg": "OK",
        "data": null
    }
    ```   
    &emsp;
    
* 3: **编辑、修改文章** ```EditArticle```
    
    - Path:   ```api/v1/admin/article/:id```
    - Method: ```PUT```
    - payload示例
    ```json
      {
      	"title": "test",
      	"cid": 1,
      	"desc": "test",
      	"content": "666",
      	"img": "test",
      	"commentCount": 999,
      	"readCount": 999
      }
    ```
    &emsp;
    
    - curl请求示例
    ```shell script
      curl -X PUT localhost:3000/api/v1/admin/article/2 -d '{"title":"test","cid":1,"desc":"test","content":"666","img":"test","commentCount":999,"readCount":999}' -H "Content-Type:application/json"
    ```
    &emsp;
    
    - 返回示例
    ```json
    {
        "status": 200,
        "msg": "OK",
        "data": null
    }
    ```
    &emsp;
             
* 4: **查询单个文章** ```GetOneArtInfo```
    
    - Path:   ```/api/v1/article/info/:id```
    - Method: ```GET```
    &emsp;
    
    - curl请求示例
    ```shell script
      curl localhost:3000/api/v1/article/info/2
    ```
    &emsp;
    
    - 返回示例
    ```json
    {
    	"status": 200,
    	"msg": "OK",
    	"data": {
    		"Category": {
    			"ID": 0,
    			"CreatedAt": "***",
    			"UpdatedAt": "***",
    			"DeletedAt": null,
    			"name": ""
    		},
    		"ID": 6,
    		"CreatedAt": "***",
    		"UpdatedAt": "***",
    		"DeletedAt": null,
    		"title": "test",
    		"cid": 1,
    		"desc": "test",
    		"content": "test666",
    		"img": "test",
    		"comment_count": 0,
    		"read_count": 1
    	}
    }
    ```
    &emsp;
     
* 5: **查询所有文章** ```GetArts```
    
    - Path:     ```/api/v1/article/```
    - Method:   ```GET```
    - Query参数: ```pagesize,pagenum,title```
    &emsp;
    
    - curl请求示例
    ```shell script
      curl localhost:3000/api/v1/article/
    ```
    &emsp;
    
    - 返回示例
    ```json
    {
    	"status": 200,
    	"msg": "OK",
    	"data": {
    		"Category": {
    			"ID": 0,
    			"CreatedAt": "***",
    			"UpdatedAt": "***",
    			"DeletedAt": null,
    			"name": ""
    		},
    		"ID": 6,
    		"CreatedAt": "***",
    		"UpdatedAt": "***",
    		"DeletedAt": null,
    		"title": "test",
    		"cid": 1,
    		"desc": "test",
    		"content": "test666",
    		"img": "test",
    		"comment_count": 0,
    		"read_count": 1
    	}
    }
    ```
    &emsp;
      
* 6: **查询分类的所有文章** ```GetCateArt```
    
    - Path:     ```/api/v1/article/list/:id```
    - Method:   ```GET```
    - Query参数: ```pagesize,pagenum```
    &emsp;
    
    - curl请求示例
    ```shell script
      curl localhost:3000/api/v1/article/list/1
    ```
    &emsp;
    
    - 返回示例
    ```json
    {
    	"data": [{
    		"Category": {
    			"id": 1,
    			"name": "666"
    		},
    		"ID": 7,
    		"CreatedAt": "2022-10-03T22:15:54.197+08:00",
    		"UpdatedAt": "2022-10-03T22:15:54.197+08:00",
    		"DeletedAt": null,
    		"title": "test",
    		"cid": 1,
    		"desc": "test",
    		"content": "test666",
    		"img": "test",
    		"comment_count": 0,
    		"read_count": 0
    	}],
    	"msg": "OK",
    	"status": 200,
    	"total": 1
    }
    ```
    &emsp;       