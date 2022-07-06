package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
		<html>
			<head>
				<title>Chat</title>
			</head>
			<body>
				Let's chat!
			</body>
		</html>
		`)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
