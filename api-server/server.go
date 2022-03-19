package main

import (
	"net/http"

	"main/booksArticleHandler"
	"main/userHandler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/user", userHandler.UserHandlerTest)
	e.GET("/user/:id", userHandler.UserGetHandler)   // ユーザー詳細情報取得
	e.POST("/user/:id", userHandler.UserPostHandler) // ユーザ登録 (LineId)
	// e.PUT("/user/:id",) // ユーザ情報更新

	e.GET("/books", booksArticleHandler.BooksArticleHandlerTest) // 投稿一覧取得
	e.GET("/books/:id", booksArticleHandler.GetArticleHandler)   // 投稿詳細取得
	e.POST("/books/:id", booksArticleHandler.PostArticleHandler) // 投稿
	// e.PUT("/books/:id",) // 投稿修正
	// e.DELETE("/books/:id",) // 投稿削除
	e.Logger.Fatal(e.Start(":80"))
}

// ハンドラー
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
