package userHandler

import (
	"database/sql"
	"fmt"
	"main/mysqlgo"
	"main/tools"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserHandlerTest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, user!")
}

// ユーザー情報取得(GET)
func UserGetHandler(c echo.Context) error {
	userId := c.Param("id")

	// sql初期化
	var connectionString = mysqlgo.InitDB()
	db, err := sql.Open("mysql", connectionString)
	mysqlgo.CheckError(err)
	defer db.Close()

	var (
		user_id      string
		name         string
		introduction string
		twitter      string
		github       string
		image        string
	)

	// ユーザーデータの取得
	rows, err := db.Query(fmt.Sprintf("SELECT * from user WHERE user_id = '%s';", userId))
	mysqlgo.CheckError(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user_id, &name, &introduction, &twitter, &github, &image)
		mysqlgo.CheckError(err)
	}
	err = rows.Err()
	mysqlgo.CheckError(err)

	type ArticleData struct {
		Article_id string        `json:"articleId"`
		Isbn       string        `json:"isbn"`
		Article    string        `json:"article"`
		Lend       bool          `json:"lend"`
		Book_data  tools.Summary `json:"bookData"`
	}

	type ResData struct {
		UserId       string        `json:"userId"`
		UserName     string        `json:"name"`
		Introduction string        `json:"introduction"`
		Image        string        `json:"image"`
		Twitter      string        `json:"twitter"`
		Github       string        `json:"Github"`
		ArticleList  []ArticleData `json:"ArticleList"`
	}

	var articleData ArticleData
	var response ResData
	articleList := []ArticleData{}

	response.UserId = userId
	response.UserName = name
	response.Introduction = introduction
	response.Image = image
	response.Twitter = twitter
	response.Github = github

	var (
		article_id string
		isbn       string
		article    string
		lend       string
	)

	articleRows, err := db.Query(fmt.Sprintf("SELECT article_id, isbn, article, lend from books_article WHERE user_id = '%s';", userId))
	mysqlgo.CheckError(err)
	defer articleRows.Close()
	for articleRows.Next() {
		err := articleRows.Scan(&article_id, &isbn, &article, &lend)
		mysqlgo.CheckError(err)
		articleData.Article_id = article_id
		articleData.Isbn = isbn
		articleData.Article = article
		articleData.Book_data, err = tools.FetchBooksData(isbn)
		if err != nil {
			return err
		}
		articleList = append(articleList, articleData)
	}

	response.ArticleList = articleList

	err = articleRows.Err()
	mysqlgo.CheckError(err)

	return c.JSON(http.StatusOK, response)
}
