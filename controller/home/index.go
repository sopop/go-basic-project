package home

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 首页
// @Description 首页
// @Accept  plain
// @Produce  html
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} service.ResponseError
// @Router / [get]
func Index(c *gin.Context) {
	data := map[string]interface{}{
		"name":    "测试项目",
		"address": "Beijing china",
		"age":     23,
	}
	// 消除html字符转义
	content := "<div><span>测试</span><br><span>这是HTML内容，消除html字符转义</span></div>"
	html := template.HTML(content)
	title := "首页"
	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"result":  data,
		"count":   len(data),
		"content": html,
		"title":   title,
	})
}
