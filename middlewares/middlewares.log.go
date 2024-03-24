package middlewares

import (
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

func Log(f echo.HandlerFunc) echo.HandlerFunc {
	starTime := time.Now()
	return func(c echo.Context) error {
		duration := time.Since(starTime)
		log.Printf("Peticion a: %v, metodo: %v duration: %v", c.Request().URL.RequestURI(), c.Request().Method, duration)

		return f(c)
	}
}
