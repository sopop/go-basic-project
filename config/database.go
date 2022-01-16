package config

// 数据库配置
const (
	driver   = "mysql"
	host     = "127.0.0.1"
	port     = "3306"
	user     = "root"
	password = "root"
	dbname   = "demo_db"
	prefix   = "cms_"
)

func InitDBConf() map[string]string {
	args := map[string]string{
		"driver":   driver,
		"host":     host,
		"port":     port,
		"user":     user,
		"password": password,
		"dbname":   dbname,
		"prefix":   prefix,
	}
	return args
}
