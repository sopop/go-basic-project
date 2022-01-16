package dao

import (
	"database/sql"
	"demo/config"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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
	GormDB   *gorm.DB
)

func init() {
	if os.Getenv("APP_LOCK") == "true" {
		GormInit()
	}
}

func GormInit() {
	sqlDb := ormInit()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{
		// GORM日志级别：Silent、Error、Warn、Info
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("failed to connect database")
	}
	log.Println("init gorm success")
	GormDB = gormDB
}

func ormInit() *sql.DB {
	dataSourceName := strings.Join([]string{dbUser, ":", dbPass,
		"@tcp(", dbHost, ":", dbPort, ")/", dbName, "?charset=utf8&parseTime=True&loc=Local"}, "")
	db, err := sql.Open(dbDriver, dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(120 * time.Second)

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	return db
}
