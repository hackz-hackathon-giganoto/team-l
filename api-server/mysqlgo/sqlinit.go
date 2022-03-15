package mysqlgo

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var host string
var database string
var user string
var password string

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func InitDB() string {
	LoadEnv()
	// 接続文字列の初期化
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)
	return connectionString
}

func LoadEnv() {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("fail read env: %v", err)
	}

	host = os.Getenv("GIGANOTO_HOST")
	database = os.Getenv("GIGANOTO_DATABASE")
	user = os.Getenv("GIGANOTO_USER")
	password = os.Getenv("GIGANOTO_PASSWORD")
}
