package userHandler

import (
	"database/sql"
	"fmt"
	"main/mysqlgo"
	"main/tools"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostUser struct {
	UserName     string
	Introduction string
	Image        string
	Twitter      string
	Github       string
}

func UserPostHandler(c echo.Context) error {
	post := new(PostUser)
	userId := tools.GenId()

	if err := c.Bind(post); err != nil {
		return err
	}

	userName := post.UserName
	introduction := post.Introduction
	image := post.Image
	twitter := post.Twitter
	github := post.Github

	//接続文字列を初期化します。
	var connectionString = mysqlgo.InitDB()

	// 接続オブジェクトを初期化する。
	db, err := sql.Open("mysql", connectionString)
	mysqlgo.CheckError(err)

	defer db.Close()

	err = db.Ping()
	mysqlgo.CheckError(err)

	// テーブルにデータを挿入する。
	sqlStatement, _ := db.Prepare(
		"INSERT INTO user (user_id, name, introduction, twitter, github, image) VALUES (?, ?, ?, ?, ?, ?);")
	res, err := sqlStatement.Exec(userId, userName, introduction, twitter, github, image)
	mysqlgo.CheckError(err)
	rowCount, _ := res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	return c.JSON(http.StatusCreated, "ok")
}
