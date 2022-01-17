// +build doc

package router

// 引入gin-swagger渲染文档数据, 利用条件编译来决定是否编译swagger，因为线上运行时，不需要swagger
// 每次修改注释都要执行swag init才能生效
// 加了+build doc编译时不生成
