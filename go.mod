module github.com/sazardev/go-deep

go 1.24.5

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/gorilla/mux v1.8.1
	github.com/stretchr/testify v1.9.0
	github.com/spf13/cobra v1.8.1
	github.com/spf13/viper v1.19.0
	go.uber.org/zap v1.27.0
	gorm.io/gorm v1.25.12
	gorm.io/driver/postgres v1.5.9
	github.com/redis/go-redis/v9 v9.6.1
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
	github.com/prometheus/client_golang v1.20.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/jaegertracing/jaeger-client-go v2.30.0+incompatible
	k8s.io/client-go v0.30.3
	github.com/docker/docker v27.1.1+incompatible
	golang.org/x/sync v0.8.0
	golang.org/x/time v0.6.0
	golang.org/x/crypto v0.25.0
	golang.org/x/net v0.27.0
)

// Módulos para desarrollo y testing
require (
	github.com/golangci/golangci-lint v1.59.1
	github.com/go-delve/delve v1.23.0
	github.com/rakyll/hey v0.1.4
	github.com/dave/dst v0.27.3
	honnef.co/go/tools v0.4.7
)

// Herramientas para documentación y generación de código
require (
	github.com/swaggo/swag v1.16.3
	github.com/deepmap/oapi-codegen v1.16.3
	github.com/99designs/gqlgen v0.17.49
	github.com/vektah/gqlparser/v2 v2.5.16
)

// Para ejemplos avanzados de sistemas distribuidos
require (
	go.etcd.io/etcd/client/v3 v3.5.15
	github.com/nats-io/nats.go v1.36.0
	github.com/segmentio/kafka-go v0.4.47
	github.com/hashicorp/consul/api v1.29.2
	github.com/hashicorp/vault/api v1.14.0
)
