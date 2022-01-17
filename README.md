### go基础项目 从零开始上手MVC框架

#### 这是用于学习GO语言的项目
#### 实现了MVC增删改查功能用于刚接触Go练手

#### 注意：
#### 这里前端用了vue，由于更直观的展示Go的应用，没用使用npm没有前后端分离，建议开发项目使用npm引入vue

#### 设置数据库参数在config中的database文件中
#### 访问：http://localhost:8089/

#### 国内仓库设置：https://goproxy.io/zh/

#### 用到的框架：
#### Gorm：https://gorm.io/zh_CN/docs/
#### Gin：https://gin-gonic.com/zh-cn/docs/

#### 用到的库：
#### github.com/joho/godotenv
#### github.com/swaggo/swag/cmd/swag
#### github.com/swaggo/files
#### github.com/swaggo/gin-swagger

#### 推荐的库
#### github.com/syyongx/php2go
#### github.com/jiazhoulvke/table2struct

#### swagger使用方法
##### 编写完注释后，使用以下命令安装swag工具：
go get -u github.com/swaggo/swag/cmd/swag
##### 在项目根目录执行以下命令，使用swag工具生成接口文档数据。
swag init
#### 访问：http://localhost:8089/swagger/index.html

#### 其中Dockerfile文件是docker build时用到

