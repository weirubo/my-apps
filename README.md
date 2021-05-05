##### 介绍
仅供读者参考的练手项目，不建议在生产环境中使用。
##### 开源库
* gin
* gorm
* validator
* viper
##### 路由
```
v1 := r.Group("/v1")
{
    v1.POST("/user/add", user.Add)
    v1.GET("/user/list", user.List)
    v1.PUT("/user/edit", user.Edit)
    v1.DELETE("/user/remove", user.Remove)
    v1.POST("/category/add", category.Add)
    v1.GET("/category/list", category.List)
    v1.PUT("/category/edit", category.Edit)
    v1.DELETE("/category/remove", category.Remove)
    v1.POST("/tag/add", tag.Add)
    v1.GET("/tag/list", tag.List)
    v1.PUT("/tag/edit", tag.Edit)
    v1.DELETE("/tag/remove", tag.Remove)
    v1.POST("/article/add", article.Add)
    v1.GET("/article/list", article.List)
    v1.GET("/article/view", article.View)
    v1.PUT("/article/edit", article.Edit)
    v1.DELETE("/article/remove", article.Remove)
    v1.POST("/comment/add", comment.Add)
    v1.GET("/comment/list", comment.List)
    v1.PUT("/comment/edit", comment.Edit)
    v1.DELETE("/comment/remove", comment.Remove)
}
```
##### 目录
```
.
├── LICENSE
├── README.md
├── api
│   └── cms
│       ├── configs
│       │   └── config.yaml
│       ├── internal
│       │   ├── dao
│       │   │   ├── article.go
│       │   │   ├── category.go
│       │   │   ├── comment.go
│       │   │   ├── dao.go
│       │   │   ├── tag.go
│       │   │   └── user.go
│       │   ├── model
│       │   │   ├── article-view.go
│       │   │   ├── article.go
│       │   │   ├── category.go
│       │   │   ├── comment.go
│       │   │   ├── tag.go
│       │   │   └── user.go
│       │   └── service
│       │       ├── article.go
│       │       ├── category.go
│       │       ├── comment.go
│       │       ├── service.go
│       │       ├── tag.go
│       │       └── user.go
│       ├── pkg
│       │   ├── common
│       │   │   └── common.go
│       │   ├── config.go
│       │   ├── mysql.go
│       │   └── trans.go
│       └── routers
│           ├── router.go
│           └── version1.0
│               ├── article.go
│               ├── category.go
│               ├── comment.go
│               ├── tag.go
│               └── user.go
├── cmd
│   └── cms
│       ├── cms
│       └── main.go
├── go.mod
└── go.sum
```
