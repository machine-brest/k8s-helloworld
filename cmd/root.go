package cmd

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Short: "Very sample web server",
	RunE: func(cmd *cobra.Command, args []string) error {

		e := echo.New()
		e.Use(middleware.Recover())

		httpListen := ":8080"
		if val, ok := os.LookupEnv("PORT"); ok && len(val) > 0 {
			httpListen = strings.TrimSpace(val)
		}

		e.GET("/health", func(c echo.Context) error {
			return c.String(http.StatusOK, "OK")
		})

		group := e.Group("/v1")

		group.Use(middleware.RequestID())
		group.Use(middleware.RemoveTrailingSlash())
		group.Use(middleware.Logger())

		group.GET("/hello", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "ok")
		})

		return e.Start(httpListen)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
