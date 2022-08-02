package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	port := os.Getenv("PORT")

	e.GET("/", func(c echo.Context) error {
		res := fmt.Sprintf("ok. port=%s", port)
		return c.JSON(http.StatusOK, res)
	})

	// クエリでの出し分け
	e.GET("/query", func(c echo.Context) error {
		name := c.QueryParam("name")
		res := fmt.Sprintf("query. name=%s, port=%s", name, port)
		return c.JSON(http.StatusOK, res)
	})

	// パスパラメタでの出し分け
	e.GET("/path/:name", func(c echo.Context) error {
		name := c.Param("name")
		res := fmt.Sprintf("query. name=%s, port=%s", name, port)
		return c.JSON(http.StatusOK, res)
	})

	log.Fatalln(e.Start(fmt.Sprintf(":%s", port)))
}
