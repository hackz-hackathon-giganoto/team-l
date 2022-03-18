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

func BooksArticleHandlerTest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, books!")
}

type PostArticle struct {
	Isbn    string `json:"isbn"`
	Article string `json:"article"`
	Lend    bool   `json:"lend"`
}

// 投稿詳細取得(GET)
func GetArticleHandler(c echo.Context) error {
	articleId := c.Param("id")

	// sql初期化
	var connectionString = mysqlgo.InitDB()
	db, err := sql.Open("mysql", connectionString)
	mysqlgo.CheckError(err)
	defer db.Close()

	// DB構造体定義
	var (
		article_id string
		user_id    string
		isbn       string
		article    string
		lend       string
	)
	rows, err := db.Query(fmt.Sprintf("SELECT * from books_article WHERE article_id = '%s';", articleId))
	mysqlgo.CheckError(err)
	defer rows.Close()
	fmt.Println("Reading data:")
	for rows.Next() {
		err := rows.Scan(&article_id, &user_id, &isbn, &article, &lend)
		mysqlgo.CheckError(err)
		fmt.Printf("Data row = (%s, %s, %s, %s, %s)\n", article_id, user_id, isbn, article, lend)
	}
	err = rows.Err()
	mysqlgo.CheckError(err)

	fetchData, err := tools.FetchBooksData(isbn)
	if err != nil {
		return err
	}

	type ResData struct {
		Article_id string        `json:"articleId"`
		User_id    string        `json:"userId"`
		Article    string        `json:"article"`
		Lend       bool          `json:"lend"`
		Book_data  tools.Summary `json:"bookData"`
	}

	var response ResData
	response.Article_id = article_id
	response.User_id = user_id
	response.Article = article
	if lend == "true" {
		response.Lend = true
	} else {
		response.Lend = false
	}
	response.Book_data = fetchData

	return c.JSON(http.StatusOK, response)
}
