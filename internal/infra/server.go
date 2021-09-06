package infra

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
	"net/http"
)

func NewServer(db *bun.DB) *echo.Echo {

	// todo: use db
	_ = db

	s := echo.New()
	s.Use(middleware.Recover())

	s.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	group := s.Group("/v1")

	group.Use(middleware.RequestID())
	group.Use(middleware.RemoveTrailingSlash())
	group.Use(middleware.Logger())

	group.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	return s
}
