package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ServerResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Data    string `json:"data,omitempty"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		
		params := c.QueryParams()

		// параметров вообще нет 
		if len(params) == 0 {
			return c.JSON(http.StatusOK, ServerResponse{
				Status:  "success",
				Message: "Hello, World!",
			})
		}

		
		// есть ли параметр name
		_, hasName := params["name"] // [value, doesExist]

		
		if !hasName {
			return c.JSON(http.StatusBadRequest, ServerResponse{
				Status:  "error",
				Message: "Bad Request: no 'name' parameter provided",
			})
		}

		
		name := c.QueryParam("name")
		return c.JSON(http.StatusOK, ServerResponse{
			Status:  "success",
			Message: "Hello, " + name + "!",
			Data: c.RealIP() + name,
		})
	})
	e.Logger.Fatal(e.Start(":1323"))
}