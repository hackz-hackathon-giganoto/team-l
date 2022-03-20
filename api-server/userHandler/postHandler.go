package userHandler

import (
	"database/sql"
	"main/mysqlgo"
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
	userId := c.Param("id")
	post := new(PostUser)

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

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	sqlStatement, _ := tx.Prepare("DELETE FROM user WHERE user_id = ?;")
	if _, err := sqlStatement.Exec(userId); err != nil {
		tx.Rollback()
		return err
	}

	// テーブルにデータを挿入する。
	sqlStatement2, _ := tx.Prepare(
		"INSERT INTO user (user_id, name, introduction, twitter, github, image) VALUES (?, ?, ?, ?, ?, ?);")
	if _, err := sqlStatement2.Exec(userId, userName, introduction, twitter, github, image); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return c.JSON(http.StatusCreated, "ok")
}
