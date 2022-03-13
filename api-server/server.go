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
	e.GET("/test", userHandler.UserHandlerTest)
	e.GET("/books", booksArticleHandler.BooksArticleHandlerTest)

	e.Logger.Fatal(e.Start(":8080"))
}

// ハンドラー
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
