package booksArticleHandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func BooksArticleHandlerTest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, books!")
}
