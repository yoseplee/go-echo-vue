package main // import "github.com/yoseplee/go-echo-vue"

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	conn "github.com/yoseplee/go-echo-vue/dbConnection"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//PostUser ...
func PostUser(c echo.Context) error {
	return c.JSON(http.StatusCreated, "")
}

func GetUser(c echo.Context) error {
	u := &user{
		ID:   1,
		Name: "test",
	}
	return c.JSON(http.StatusCreated, u)
}

func main() {

	/*
		conn.InitDB()
		conn.PrintVersionOfDB()
	*/

	e := echo.New()
	e.Use(middleware.CORS()) //for vue-cli dev server

	e.Static("/", "dist")
	e.File("/", "dist/index.html")

	e.GET("/user", GetUser)
	e.POST("/test", PostUser)
	e.GET("/api", func(c echo.Context) error {
		u := user{ID: 3, Name: "Lee"}
		return c.JSON(http.StatusCreated, u)
	})
	e.Logger.Fatal(e.Start(":1323"))

	defer conn.GetDB().Close()
}
