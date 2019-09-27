package main // import "github.com/yoseplee/go-echo-vue"

import (
	"net/http"
	conn "github.com/yoseplee/go-echo-vue/dbConnection"
	"github.com/labstack/echo"
)

type user struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

//PostUser ...
func PostUser(c echo.Context) error {
	return c.JSON(http.StatusCreated, "")
}

func GetUser(c echo.Context) error {
	u := &user{
		ID: 1,
		Name: "test",
	}

	return c.JSON(http.StatusCreated, u)
}

func main() {

	conn.InitDB()
	conn.PrintVersionOfDB()
	
	e := echo.New()
	e.GET("/user", GetUser)
	e.POST("/test", PostUser)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

	defer conn.GetDB().Close()
}