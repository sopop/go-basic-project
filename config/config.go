package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitConf() map[string]string {
	arr := make(map[string]string)
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		return arr
	}

	arr["debug"] = os.Getenv("APP_DEBUG")
	return arr
}
