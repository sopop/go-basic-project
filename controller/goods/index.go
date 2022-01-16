package goods

import (
	"demo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Goods(c *gin.Context) {
	data := map[string]interface{}{
		"name":  "《go语言开发》",
		"level": 2,
	}
	title := "商品标题"
	c.HTML(http.StatusOK, "goods/goods.html", gin.H{
		"title":  title,
		"result": data,
	})
}

func GoodsLists(c *gin.Context) {
	page, _ := strconv.Atoi(c.PostForm("page"))
	size, _ := strconv.Atoi(c.PostForm("pageSize"))
	data := service.RequestInfo{
		Page:     page,
		PageSize: size,
	}
	lists, _ := service.NewGoods().GetLists(data)
	c.JSON(http.StatusOK, gin.H{
		"lists": lists,
	})
}

func GoodsDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	goods, _ := service.NewGoods().GetDetail(id)
	c.JSON(http.StatusOK, gin.H{
		"lists": goods,
	})
}

func Operate(c *gin.Context) {
	var (
		code int
		msg  string
		err  error
	)
	data := map[string]interface{}{
		"name":   c.PostForm("name"),
		"detail": c.PostForm("detail"),
	}

	id, _ := strconv.Atoi(c.PostForm("id"))
	if id > 0 {
		err = service.NewGoods().EditGoods(id, data)
	} else {
		err = service.NewGoods().AddGoods(data)
	}

	if err != nil {
		code = 1
		msg = "提交失败"
	} else {
		code = 0
		msg = "提交成功"
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func GoodsDel(c *gin.Context) {
	var (
		code int
		msg  string
		err  error
	)
	id, _ := strconv.Atoi(c.PostForm("id"))
	err = service.NewGoods().DelGoods(id)
	if err != nil {
		code = 1
		msg = "删除失败"
	} else {
		code = 0
		msg = "删除成功"
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
