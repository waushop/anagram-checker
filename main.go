package main

import (
	"anagramchecker/services"
	"crypto/subtle"
	"net/http"

	echotemplate "github.com/foolin/echo-template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	e.Static("/assets", "assets")

	e.Renderer = echotemplate.Default()

	// Enable request logging middleware
	e.Use(middleware.Logger())

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("admin")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("password")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	e.GET("/", func(c echo.Context) error {
		ac := &services.AnagramChecker{}
		ac.IsAnagram = false
		data := echo.Map{
			"isAnagram": ac.IsAnagram,
		}
		return c.Render(http.StatusOK, "checkAnagram.html", data)
	})
	e.POST("/check-anagram", services.CheckAnagram)

	e.Logger.Fatal(e.Start(":9000"))
}
