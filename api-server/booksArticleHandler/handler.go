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

// 投稿(POST)
func PostArticleHandler(c echo.Context) error {
	post := new(PostArticle)
	userId := c.Param("id")

	if err := c.Bind(post); err != nil {
		return err
	}
	var lend string
	articleId := tools.GenId()
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

	err = db.Ping()
	mysqlgo.CheckError(err)
	fmt.Println("Successfully created connection to database.")

	// テーブルにデータを挿入する。
	sqlStatement, err := db.Prepare(
		"INSERT INTO books_article (article_id, user_id, isbn, article, lend) VALUES (?, ?, ?, ?, ?);")
	res, err := sqlStatement.Exec(articleId, userId, isbn, article, lend)
	mysqlgo.CheckError(err)
	rowCount, err := res.RowsAffected()
	fmt.Printf("Inserted %d row(s) of data.\n", rowCount)

	return c.JSON(http.StatusCreated, "ok")
}
