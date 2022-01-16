package router

import (
	"demo/config"
	"demo/controller/goods"
	"demo/controller/home"
	"demo/controller/mysql"
	"demo/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Strat() *gin.Engine {
	route := gin.Default()
	// 开发模式
	conf := config.InitConf()
	if conf["debug"] == "false" {
		fmt.Println(conf)
		// 生产模式
		gin.SetMode(gin.ReleaseMode)
	}

	// 中间件验证apiToken
	route.Use(middleware.Token())
	// 注册模板
	route.LoadHTMLGlob("view/**/*")

	// 注册静态文件
	route.Static("./static", "static")
	// 注册路由
	suffix := ".html"
	route.GET("/", home.Index)
	route.GET("/home"+suffix, home.Index)
	route.GET("/goods"+suffix, goods.Goods)
	// 路由分组
	api := route.Group("/api")
	{
		// 创建数据库
		api.POST("/create_mysql", mysql.CreateMysqlDB)
		// 商品
		api.POST("/goods_lists", goods.GoodsLists)
		api.POST("/goods_detail", goods.GoodsDetail)
		api.POST("/goods_operate", goods.Operate)
		api.POST("/goods_del", goods.GoodsDel)
	}
	return route
}
