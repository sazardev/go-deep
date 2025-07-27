# 🧪 Makefile para Testing en Go Deep - Lección 15
# ==================================================

.PHONY: help test test-short test-race test-cover test-stress bench clean install-deps setup

# Variables
PACKAGE := ./...
COVERAGE_OUT := reports/coverage.out
COVERAGE_HTML := reports/coverage.html

# Colores para output
GREEN := \033[0;32m
YELLOW := \033[1;33m
BLUE := \033[0;34m
NC := \033[0m

# Target por defecto
help: ## Mostrar ayuda
	@echo -e "$(BLUE)🧪 Go Deep - Testing Makefile$(NC)"
	@echo "=============================="
	@echo ""
	@echo "Targets disponibles:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(NC) %s\n", $$1, $$2}'
	@echo ""
	@echo "Ejemplos:"
	@echo "  make test          # Ejecutar todos los tests"
	@echo "  make test-race     # Tests con race detection"
	@echo "  make bench         # Solo benchmarks"
	@echo "  make coverage      # Tests + coverage HTML"

# =============================================================================
# SETUP Y DEPENDENCIAS
# =============================================================================

setup: ## Configurar entorno de testing
	@echo -e "$(BLUE)📦 Configurando entorno...$(NC)"
	@mkdir -p reports
	@go mod tidy
	@echo -e "$(GREEN)✅ Entorno configurado$(NC)"

install-deps: ## Instalar dependencias de testing
	@echo -e "$(BLUE)📦 Instalando dependencias...$(NC)"
	@go install -a std
	@echo -e "$(GREEN)✅ Dependencias instaladas$(NC)"

# =============================================================================
# TESTING BÁSICO
# =============================================================================

test: setup ## Ejecutar todos los tests
	@echo -e "$(BLUE)🧪 Ejecutando tests...$(NC)"
	@go test -v $(PACKAGE)

test-short: setup ## Ejecutar solo tests rápidos
	@echo -e "$(BLUE)⚡ Ejecutando tests rápidos...$(NC)"
	@go test -short -v $(PACKAGE)

test-verbose: setup ## Tests con output verbose
	@echo -e "$(BLUE)📝 Tests con output detallado...$(NC)"
	@go test -v -count=1 $(PACKAGE)

# =============================================================================
# CONCURRENCIA Y RACE CONDITIONS
# =============================================================================

test-race: setup ## Tests con race detection
	@echo -e "$(BLUE)🏃 Detectando race conditions...$(NC)"
	@go test -race -v $(PACKAGE)

test-stress: setup ## Ejecutar tests de stress
	@echo -e "$(BLUE)💪 Tests de stress...$(NC)"
	@go test -v -timeout=30s -run="Stress|Concurrent" $(PACKAGE)

# =============================================================================
# COVERAGE
# =============================================================================

coverage: setup ## Generar reporte de coverage
	@echo -e "$(BLUE)📊 Generando coverage...$(NC)"
	@go test -coverprofile=$(COVERAGE_OUT) -coverpkg=$(PACKAGE) $(PACKAGE)
	@go tool cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)
	@echo -e "$(GREEN)✅ Coverage HTML: $(COVERAGE_HTML)$(NC)"

coverage-func: setup ## Mostrar coverage por función
	@echo -e "$(BLUE)📈 Coverage por función...$(NC)"
	@go test -coverprofile=$(COVERAGE_OUT) -coverpkg=$(PACKAGE) $(PACKAGE) > /dev/null
	@go tool cover -func=$(COVERAGE_OUT)

coverage-total: setup ## Mostrar solo coverage total
	@go test -coverprofile=$(COVERAGE_OUT) -coverpkg=$(PACKAGE) $(PACKAGE) > /dev/null 2>&1
	@go tool cover -func=$(COVERAGE_OUT) | grep total | awk '{print "Coverage: " $$3}'

test-cover: coverage ## Alias para coverage

# =============================================================================
# BENCHMARKS
# =============================================================================

bench: setup ## Ejecutar benchmarks
	@echo -e "$(BLUE)⚡ Ejecutando benchmarks...$(NC)"
	@go test -bench=. -benchmem $(PACKAGE)

bench-cpu: setup ## Benchmarks con CPU profiling
	@echo -e "$(BLUE)🔥 Benchmarks con CPU profiling...$(NC)"
	@mkdir -p reports
	@go test -bench=. -benchmem -cpuprofile=reports/cpu.prof $(PACKAGE)

bench-mem: setup ## Benchmarks con memory profiling
	@echo -e "$(BLUE)🧠 Benchmarks con memory profiling...$(NC)"
	@mkdir -p reports
	@go test -bench=. -benchmem -memprofile=reports/mem.prof $(PACKAGE)

bench-compare: setup ## Comparar benchmarks (ejecutar dos veces)
	@echo -e "$(BLUE)📊 Benchmarks de comparación...$(NC)"
	@mkdir -p reports
	@go test -bench=. -count=5 $(PACKAGE) > reports/bench_new.txt
	@echo -e "$(GREEN)✅ Resultados en reports/bench_new.txt$(NC)"

# =============================================================================
# LINTING Y FORMATO
# =============================================================================

fmt: ## Formatear código
	@echo -e "$(BLUE)🎨 Formateando código...$(NC)"
	@gofmt -w .
	@echo -e "$(GREEN)✅ Código formateado$(NC)"

vet: ## Ejecutar go vet
	@echo -e "$(BLUE)🔍 Ejecutando go vet...$(NC)"
	@go vet $(PACKAGE)
	@echo -e "$(GREEN)✅ go vet completado$(NC)"

lint: fmt vet ## Formatear + vet

# =============================================================================
# FUZZING (Go 1.18+)
# =============================================================================

fuzz: setup ## Ejecutar fuzzing tests
	@echo -e "$(BLUE)🔀 Ejecutando fuzzing...$(NC)"
	@timeout 30s go test -fuzz=. -fuzztime=30s $(PACKAGE) || true

# =============================================================================
# TARGETS COMBINADOS
# =============================================================================

full: lint test test-race coverage bench ## Suite completa de testing
	@echo -e "$(GREEN)🎉 Suite completa ejecutada$(NC)"

ci: setup test-short test-race coverage-total ## Tests para CI/CD
	@echo -e "$(GREEN)✅ CI tests completados$(NC)"

dev: test-short coverage ## Tests rápidos para desarrollo
	@echo -e "$(GREEN)⚡ Dev tests completados$(NC)"

# =============================================================================
# UTILIDADES
# =============================================================================

clean: ## Limpiar archivos generados
	@echo -e "$(BLUE)🧹 Limpiando archivos...$(NC)"
	@rm -rf reports/
	@go clean -testcache
	@echo -e "$(GREEN)✅ Archivos limpiados$(NC)"

watch: ## Ejecutar tests cuando cambie el código (requiere entr)
	@echo -e "$(BLUE)👀 Watching for changes...$(NC)"
	@find . -name "*.go" | entr -c make test-short

deps-check: ## Verificar dependencias
	@echo -e "$(BLUE)🔍 Verificando dependencias...$(NC)"
	@go mod verify
	@go mod tidy
	@echo -e "$(GREEN)✅ Dependencias verificadas$(NC)"

# =============================================================================
# TARGETS ESPECÍFICOS POR EJERCICIO
# =============================================================================

test-book: setup ## Tests para BookValidator
	@go test -v -run TestBookValidator

test-inventory: setup ## Tests para InventoryService
	@go test -v -run TestInventoryService

test-cart: setup ## Tests para ShoppingCart
	@go test -v -run TestShoppingCart

test-payment: setup ## Tests para PaymentProcessor
	@go test -v -run TestPaymentProcessor

test-search: setup ## Tests para SearchEngine
	@go test -v -run TestSearchEngine

test-api: setup ## Tests para APIServer
	@go test -v -run TestAPIServer

test-calculator: setup ## Tests para Calculator
	@go test -v -run TestCalculator

test-worker: setup ## Tests para WorkerPool
	@go test -v -run TestWorkerPool

# =============================================================================
# BENCHMARKS ESPECÍFICOS
# =============================================================================

bench-book: setup ## Benchmark BookValidator
	@go test -bench=BenchmarkBookValidator -benchmem

bench-inventory: setup ## Benchmark InventoryService
	@go test -bench=BenchmarkInventoryService -benchmem

bench-cart: setup ## Benchmark ShoppingCart
	@go test -bench=BenchmarkShoppingCart -benchmem

bench-search: setup ## Benchmark SearchEngine
	@go test -bench=BenchmarkSearchEngine -benchmem

# =============================================================================
# TARGETS DE EJEMPLO
# =============================================================================

example: ## Ejecutar ejemplos
	@echo -e "$(BLUE)📖 Ejecutando ejemplos...$(NC)"
	@go test -v -run Example

demo: setup ## Demostración de testing
	@echo -e "$(BLUE)🎬 Demostración de testing...$(NC)"
	@echo ""
	@echo "1. Tests básicos:"
	@make test-short
	@echo ""
	@echo "2. Coverage:"
	@make coverage-total
	@echo ""
	@echo "3. Benchmarks rápidos:"
	@go test -bench=BenchmarkAdd -benchtime=1s
	@echo ""
	@echo -e "$(GREEN)✅ Demostración completada$(NC)"

# =============================================================================
# INFORMACIÓN DEL SISTEMA
# =============================================================================

info: ## Información del entorno
	@echo -e "$(BLUE)ℹ️  Información del entorno$(NC)"
	@echo "=========================="
	@echo "Go version: $(shell go version)"
	@echo "GOOS: $(shell go env GOOS)"
	@echo "GOARCH: $(shell go env GOARCH)"
	@echo "CGO_ENABLED: $(shell go env CGO_ENABLED)"
	@echo "Test cache: $(shell go env GOCACHE)"
	@echo ""
	@echo "Módulos:"
	@go list -m all | head -5

# =============================================================================
# SCRIPTS PERSONALIZADOS
# =============================================================================

suite: ## Ejecutar suite completa con script
	@echo -e "$(BLUE)🚀 Ejecutando suite completa...$(NC)"
	@./coverage.sh

# Target para mostrar estadísticas
stats: setup ## Mostrar estadísticas de tests
	@echo -e "$(BLUE)📊 Estadísticas de tests$(NC)"
	@echo "========================"
	@echo "Archivos de test: $(shell find . -name "*_test.go" | wc -l)"
	@echo "Funciones de test: $(shell grep -r "^func Test" . | wc -l)"
	@echo "Benchmarks: $(shell grep -r "^func Benchmark" . | wc -l)"
	@echo "Examples: $(shell grep -r "^func Example" . | wc -l)"
