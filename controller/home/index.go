package home

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
