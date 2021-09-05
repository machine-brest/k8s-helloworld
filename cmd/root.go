package cmd

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/machine-brest/k8s-helloworld/srv"
	"github.com/machine-brest/k8s-helloworld/srv/data"
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

		db := pg.Connect(&pg.Options{
			Addr:     "postgres:5432",
			User:     "dbuser",
			Password: "dbpass",
			Database: "sample",
		})

		err := createSchema(db)
		if err != nil {
			return err
		}

		defer func() {
			_ = db.Close()
		}()

		group.POST("/login", srv.Login(db))

		group.GET("/hello", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "ok")
		})

		return e.Start(httpListen)
	},
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*data.User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
