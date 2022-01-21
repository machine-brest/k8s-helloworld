package cmd

import (
	"fmt"
	"github.com/machine-brest/k8s-helloworld/internal/infra"
	"github.com/spf13/cobra"
	"os"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

const (
	service     = "k8s-helloworld"
	environment = "dev"
	id          = 1
)

var rootCmd = &cobra.Command{
	Short:         "API Backend",
	SilenceUsage:  true,
	SilenceErrors: false,
	PreRunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		tp, err := tracerProvider("http://http://jaeger-agent.default.svc:14268/api/traces")
		if err != nil {
			panic(err)
		}

		db := infra.NewDatabase()
		dbConn, err := db.GetConnection()
		if err != nil {
			panic(err)
		}

		// Register our TracerProvider as the global so any imported
		// instrumentation in the future will default to using it.
		otel.SetTracerProvider(tp)

		srv := infra.NewServer(dbConn)

		httpListen := ":8080"
		if val, ok := os.LookupEnv("PORT"); ok && len(val) > 0 {
			httpListen = strings.TrimSpace(val)
		}

		return srv.Start(httpListen)
	},
}

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
