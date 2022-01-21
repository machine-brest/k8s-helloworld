module github.com/machine-brest/k8s-helloworld

go 1.16

require (
	github.com/labstack/echo/v4 v4.5.0
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	github.com/uptrace/bun v1.0.3
	github.com/uptrace/bun/dialect/pgdialect v1.0.3
	github.com/uptrace/bun/driver/pgdriver v1.0.3
	go.opentelemetry.io/otel v1.3.0
	go.opentelemetry.io/otel/exporters/jaeger v1.3.0
	go.opentelemetry.io/otel/sdk v1.3.0
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)
