package booksArticleHandler

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"main/mysqlgo"
	"main/tools"

	"github.com/labstack/echo/v4"
)

// 投稿(POST)
func PostArticleHandler(c echo.Context) error {
	post := new(PostArticle)
	userId := c.Param("id")
	articleId := tools.GenId()

	if err := c.Bind(post); err != nil {
		return err
	}
	var lend string

	isbn := post.Isbn
	article := post.Article
	if post.Lend {
		lend = "true"
	} else {
		lend = "false"
	}

	//接続文字列を初期化します。
	var connectionString = mysqlgo.InitDB()

	// 接続オブジェクトを初期化する。
	db, err := sql.Open("mysql", connectionString)
	mysqlgo.CheckError(err)

	defer db.Close()

	tx, _ := db.Begin()

	// テーブルにデータを挿入する。
	sqlStatement, _ := tx.Prepare("INSERT INTO books_article (article_id, user_id, isbn, article, lend) VALUES (?, ?, ?, ?, ?);")
	res, err := sqlStatement.Exec(articleId, userId, isbn, article, lend)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowCount, _ := res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	tx.Commit()
	return c.JSON(http.StatusCreated, "ok")
}
