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

	// パスパラメタでの出し分け
	e.GET("/path/:name", func(c echo.Context) error {
		name := c.Param("name")
		res := fmt.Sprintf("path. name=%s, port=%s", name, port)
		return c.JSON(http.StatusOK, res)
	})

	// クエリでの出し分け
	e.GET("/query", func(c echo.Context) error {
		name := c.QueryParam("name")
		res := fmt.Sprintf("query. name=%s, port=%s", name, port)
		return c.JSON(http.StatusOK, res)
	})

	// URLに特定の文字列が含まれているかどうかで出し分ける
	e.GET("/match/:name", func(c echo.Context) error {
		name := c.Param("name")
		res := fmt.Sprintf("path2. name=%s, port=%s", name, port)
		return c.JSON(http.StatusOK, res)
	})
	e.GET("/match/:name/hoge", func(c echo.Context) error {
		name := c.Param("name")
		res := fmt.Sprintf("path2 hoge. name=%s, port=%s", name, port)
		return c.JSON(http.StatusOK, res)
	})

	log.Fatalln(e.Start(fmt.Sprintf(":%s", port)))
}
