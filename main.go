package main

import (
	"demo/router"
)

// @title go基础项目
// @version 1.0
// @description go 基础项目 api docs.
// @license.name Apache 2.0
// @contact.name go-swagger帮助文档
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host localhost:8089
// @BasePath /api
func main() {
	router := router.Strat()
	router.Run(":8089")
}
