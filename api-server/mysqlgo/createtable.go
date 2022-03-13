package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var host string
var database string
var user string
var password string

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	loadEnv()

	// 接続文字列を初期化します。
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	// 接続オブジェクトを初期化する。
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")

	// 同名の以前のテーブルがあれば削除する。
	_, err = db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("Finished dropping table (if existed).")

	// テーブルを作成します。
	_, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
	checkError(err)
	fmt.Println("Finished creating table.")

	// テーブルにデータを挿入する。
	sqlStatement, err := db.Prepare("INSERT INTO inventory (name, quantity) VALUES (?, ?);")
	res, err := sqlStatement.Exec("banana", 150)
	checkError(err)
	rowCount, err := res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	res, err = sqlStatement.Exec("orange", 154)
	checkError(err)
	rowCount, err = res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	res, err = sqlStatement.Exec("apple", 100)
	checkError(err)
	rowCount, err = res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)
	fmt.Println("Done.")
}

func loadEnv() {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	host = os.Getenv("GIGANOTO_HOST")
	database = os.Getenv("GIGANOTO_DATABASE")
	user = os.Getenv("GIGANOTO_USER")
	password = os.Getenv("GIGANOTO_PASSWORD")

	fmt.Println(host)
	fmt.Println(database)
	fmt.Println(user)
	fmt.Println(password)
}
