package goods

import (
	"demo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 商品页
// @Description 展示商品
// @Accept  plain
// @Produce  html
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} service.ResponseError
// @Router /api/goods_lists [get]
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

// @Summary 商品列表
// @Description 获取商品列表
// @Accept  json
// @Produce  json
// @param ApiToken header string true "apitoken验证"
// @Param  page query string true "页码"
// @Param  pageSize query string true "每页条数"
// @Success 200 {object} service.Response{data=model.Goods}
// @Failure 500 {object} service.ResponseError
// @Router /api/goods_lists [post]
func GoodsLists(c *gin.Context) {
	var (
		code int
		msg  string
		err  error
	)
	page, _ := strconv.Atoi(c.PostForm("page"))
	size, _ := strconv.Atoi(c.PostForm("pageSize"))
	data := service.RequestInfo{
		Page:     page,
		PageSize: size,
	}
	lists, err := service.NewGoods().GetLists(data)
	if err != nil {
		code = 1
		msg = "提交失败"
	} else {
		code = 0
		msg = "提交成功"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   msg,
		"lists": lists,
	})
}

// @Summary 商品详情
// @Description 获取商品信息
// @Accept  json
// @Produce  json
// @param ApiToken header string true "apitoken验证"
// @Param  id query string true "商品ID"
// @Success 200 {object} service.Response{data=model.Goods}
// @Failure 500 {object} service.ResponseError
// @Router /api/goods_detail [post]
func GoodsDetail(c *gin.Context) {
	var (
		code int
		msg  string
		err  error
	)
	id, _ := strconv.Atoi(c.PostForm("id"))
	goods, err := service.NewGoods().GetDetail(id)
	if err != nil {
		code = 1
		msg = "提交失败"
	} else {
		code = 0
		msg = "提交成功"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   msg,
		"lists": goods,
	})
}

// @Summary 添加、编辑商品
// @Description 添加、编辑商品信息
// @Accept  json
// @Produce  json
// @param ApiToken header string true "apitoken验证"
// @Param  id query string true "商品ID"
// @Param  name query string true "商品名称"
// @Param  detail query string true "商品描述"
// @Success 200 {object} service.Response
// @Failure 500 {object} service.ResponseError
// @Router /api/goods_operate [post]
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

// @Summary 删除商品
// @Description 删除商品信息
// @Accept  json
// @Produce  json
// @param ApiToken header string true "apitoken验证"
// @Param  id query string true "商品ID"
// @Success 200 {object} service.Response
// @Failure 500 {object} service.ResponseError
// @Router /api/goods_del [post]
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
