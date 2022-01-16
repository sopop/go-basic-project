package service

import (
	"demo/dao"
)

// 成功后返回数据结构
type Response struct {
	Code int         `json:"code"` //为0时表示成功
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 失败后返回数据结构
type ResponseError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 请求信息
type RequestInfo struct {
	Page     int    `json:"page,omitempty"`     //页码
	PageSize int    `json:"pageSize,omitempty"` //每页多少条
	Content  string `json:"content,omitempty"`  //指定工作的服务
	Sreach   string `json:"sreach,omitempty"`   //搜索条件
	MsgType  int    `json:"msgType,omitempty"`  //是否返回数据
}

func InitGormDB() {
	dao.GormInit()
}
