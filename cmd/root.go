package cmd

import (
	"fmt"
	"github.com/machine-brest/k8s-helloworld/internal/infra"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Short:         "API Backend",
	SilenceUsage:  true,
	SilenceErrors: false,
	RunE: func(cmd *cobra.Command, args []string) error {

		db := infra.NewDatabase()
		dbConn, err := db.GetConnection()
		if err != nil {
			panic(err)
		}

		srv := infra.NewServer(dbConn)

		httpListen := ":8080"
		if val, ok := os.LookupEnv("PORT"); ok && len(val) > 0 {
			httpListen = strings.TrimSpace(val)
		}

		return srv.Start(httpListen)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
