package router

import (
	"main/controller"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/user", controller.UserHandlerTest)
	e.GET("/user/:id", controller.UserGetHandler)   // ユーザー詳細情報取得
	e.POST("/user/:id", controller.UserPostHandler) // ユーザ登録 (LineId)
	// e.PUT("/user/:id",) // ユーザ情報更新

	e.GET("/books", controller.BooksArticleHandlerTest) // 投稿一覧取得
	e.GET("/books/:id", controller.GetArticleHandler)   // 投稿詳細取得
	e.POST("/books/:id", controller.PostArticleHandler) // 投稿
	// e.PUT("/books/:id",) // 投稿修正
	// e.DELETE("/books/:id",) // 投稿削除
	e.Logger.Fatal(e.Start(":80"))
}
