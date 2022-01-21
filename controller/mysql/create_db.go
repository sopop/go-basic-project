package mysql

import (
	"database/sql"
	"demo/config"
	"demo/service"
	"demo/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type DBTask struct {
	DBCon *sql.DB
}

var (
	//导入配置文件
	configMap = config.InitDBConf()
	//查看配置文件里所有键值对
	dbDriver = configMap["driver"]
	dbHost   = configMap["host"]
	dbPort   = configMap["port"]
	dbName   = configMap["dbname"]
	dbUser   = configMap["user"]
	dbPass   = configMap["password"]
)

func CreateMysqlDB(c *gin.Context) {
	dbTask := &DBTask{}
	var (
		err  error
		code int
		msg  string
	)
	// 连接数据库
	dataSourceName := strings.Join([]string{dbUser, ":", dbPass,
		"@tcp(", dbHost, ":", dbPort, ")/mysql?charset=utf8"}, "")
	dbTask.DBCon, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println("failed to open database:", err.Error())
		return
	}
	defer dbTask.DBCon.Close()
	// 创建数据库
	err = dbTask.createDB()

	if err != nil {
		code = 1
		msg = "操作失败,检查数据库是否连接"
	} else {
		copy()
		appendToFile()
		service.InitGormDB()
		code = 0
		msg = "提交成功"
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func (d *DBTask) createDB() error {
	createDbSql := strings.Join([]string{"CREATE DATABASE IF NOT EXISTS ", dbName,
		" DEFAULT CHARSET UTF8MB4 COLLATE UTF8MB4_UNICODE_CI;"}, "")
	_, err := d.DBCon.Exec(createDbSql)
	if err != nil {
		log.Println("failed to create databases", err.Error())
		return err
	}

	useDbSql := strings.Join([]string{"USE ", dbName}, "")
	_, err = d.DBCon.Exec(useDbSql)
	if err != nil {
		log.Println("select database failed")
		return err
	}

	createTbSql := strings.Join([]string{"CREATE TABLE IF NOT EXISTS cms_goods(",
		"id int(10) primary key auto_increment COMMENT '商品ID',",
		"cat_id int(10) COMMENT '分类ID',",
		"name varchar(20) COMMENT '商品名称',",
		"detail varchar(255) COMMENT '商品详情',",
		"create_time int(10) COMMENT '创建时间',",
		"update_time int(10) COMMENT '更新时间',",
		"status tinyint(1) COMMENT '商品状态'",
		") ENGINE=InnoDB;"}, "")
	_, err = d.DBCon.Exec(createTbSql)
	if err != nil {
		log.Println("create table  failed:", err.Error())
		return err
	}
	return nil
}

func copy() {
	var dir string
	dir, _ = exec.LookPath(os.Args[0])
	file1 := "/temp/index.tmp"
	file2 := "/view/home/index.html"
	file1Path := path.Join(path.Dir(dir), file1)
	file2Path := path.Join(path.Dir(dir), file2)

	// 检查文件是否存在
	_, err := os.Stat(file1Path)
	if err != nil && os.IsNotExist(err) {
		mainPath := utils.GetCurrentAbsPath()
		file1Path = path.Join(mainPath, file1)
		file2Path = path.Join(mainPath, file2)
	}
	log.Println(dir)
	log.Println(file1Path)

	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		log.Printf("文件打开失败=%v\n", err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		log.Printf("文件打开失败=%v\n", err)
	}
}

func appendToFile() {
	file := "/.env"
	str := "\n#初始化完成\n" + "APP_LOCK=true"
	var dir string
	dir, _ = exec.LookPath(os.Args[0])
	filePath := path.Join(path.Dir(dir), file)

	// 检查文件是否存在
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		mainPath := utils.GetCurrentAbsPath()
		filePath = path.Join(mainPath, file)
	}
	log.Println(filePath)

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		log.Printf("Cannot open file %s!\n", file)
		return
	}
	defer f.Close()
	f.WriteString(str)
}
