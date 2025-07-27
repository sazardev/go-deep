# 🧪 Lección 15: Testing en Go - Archivos de Ejercicios

## 📁 Estructura del Proyecto

```
15-testing/
├── README.md                    # Tutorial completo de testing
├── ejercicios.go               # 8 ejercicios prácticos
├── book_validator_test.go      # Ejemplo: Tests unitarios con table-driven tests
├── inventory_service_test.go   # Ejemplo: Mocking y dependency injection
├── shopping_cart_test.go       # Ejemplo: Tests de concurrencia
├── coverage.sh                 # Script completo de testing y coverage
├── Makefile                    # Automatización de tareas de testing
├── .vscode/
│   ├── settings.json          # Configuración de Go para VS Code
│   └── tasks.json             # Tareas de testing para VS Code
└── reports/                   # Directorio para reportes (se crea automáticamente)
```

## 🚀 Inicio Rápido

### 1. Ejecutar Tests Básicos
```bash
# Todos los tests
make test

# Solo tests rápidos
make test-short

# Tests específicos
make test-book
make test-inventory
make test-cart
```

### 2. Coverage y Quality
```bash
# Coverage HTML
make coverage

# Suite completa
make full

# Script completo con reportes
./coverage.sh
```

### 3. Tests de Concurrencia
```bash
# Race detection
make test-race

# Tests de stress
make test-stress
```

### 4. Benchmarks
```bash
# Todos los benchmarks
make bench

# Benchmarks específicos
make bench-book
make bench-cart
```

## 📚 Ejercicios Incluidos

### 🎯 Ejercicio 1: BookValidator - Tests Unitarios
- **Archivo**: `book_validator_test.go` (ejemplo implementado)
- **Conceptos**: Table-driven tests, validaciones, edge cases
- **Tests**: Validación de libros e ISBN-10/ISBN-13

### 🎯 Ejercicio 2: InventoryService - Mocking
- **Archivo**: `inventory_service_test.go` (ejemplo implementado)
- **Conceptos**: Dependency injection, mocks manuales, verificación de llamadas
- **Tests**: Servicios con múltiples dependencias

### 🎯 Ejercicio 3: ShoppingCart - Concurrencia
- **Archivo**: `shopping_cart_test.go` (ejemplo implementado)
- **Conceptos**: Race conditions, sync.RWMutex, tests paralelos
- **Tests**: Operaciones concurrentes seguras

### 🎯 Ejercicio 4: PaymentProcessor - Integration Testing
- **Conceptos**: HTTP testing, context, timeouts
- **Tu tarea**: Implementar tests de integración

### 🎯 Ejercicio 5: SearchEngine - Benchmarks
- **Conceptos**: Performance testing, algoritmos de búsqueda
- **Tu tarea**: Optimizar y benchmarkear

### 🎯 Ejercicio 6: APIServer - HTTP Testing
- **Conceptos**: httptest, REST APIs, middlewares
- **Tu tarea**: Tests end-to-end de API

### 🎯 Ejercicio 7: Calculator - TDD
- **Conceptos**: Test-Driven Development, ciclo Red-Green-Refactor
- **Tu tarea**: Implementar usando TDD estricto

### 🎯 Ejercicio 8: WorkerPool - Concurrencia Avanzada
- **Conceptos**: Goroutines, channels, worker patterns
- **Tu tarea**: Tests de pools concurrentes

## 🔧 Comandos Útiles

### Testing Básico
```bash
# Ejecutar tests con verbose
go test -v ./...

# Solo tests que coincidan con patrón
go test -run TestBookValidator

# Tests con timeout
go test -timeout 30s ./...

# Limpiar cache de tests
go clean -testcache
```

### Race Detection
```bash
# Detectar race conditions
go test -race ./...

# Con output detallado
go test -race -v ./...
```

### Coverage
```bash
# Coverage básico
go test -cover ./...

# Coverage detallado
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Coverage por función
go tool cover -func=coverage.out
```

### Benchmarks
```bash
# Todos los benchmarks
go test -bench=. ./...

# Con información de memoria
go test -bench=. -benchmem ./...

# Benchmarks específicos
go test -bench=BenchmarkAdd

# Múltiples ejecuciones para precisión
go test -bench=. -count=5
```

### Fuzzing (Go 1.18+)
```bash
# Fuzzing por tiempo limitado
go test -fuzz=FuzzValidateEmail -fuzztime=30s

# Fuzzing hasta encontrar problema
go test -fuzz=FuzzValidateEmail
```

## 🎯 Criterios de Evaluación

### ✅ Tests Unitarios (25%)
- [ ] Table-driven tests implementados
- [ ] Edge cases cubiertos
- [ ] Validaciones correctas
- [ ] Subtests organizados

### ✅ Mocking y DI (25%)
- [ ] Mocks manuales implementados
- [ ] Dependency injection correcta
- [ ] Verificación de llamadas a mocks
- [ ] Casos de error testeados

### ✅ Concurrencia (20%)
- [ ] Tests sin race conditions
- [ ] Operaciones concurrentes seguras
- [ ] Stress tests implementados
- [ ] Benchmarks paralelos

### ✅ Coverage y Quality (15%)
- [ ] Coverage > 80%
- [ ] No dead code
- [ ] go vet sin warnings
- [ ] Código formateado

### ✅ Integration y E2E (10%)
- [ ] HTTP tests implementados
- [ ] Context y timeouts manejados
- [ ] Mocks de servicios externos
- [ ] Tests de flujos completos

### ✅ TDD y Metodología (5%)
- [ ] TDD aplicado correctamente
- [ ] Refactoring con tests verdes
- [ ] Tests como documentación
- [ ] Commits atómicos

## 📊 Métricas de Éxito

### 🎯 Targets Mínimos
- **Coverage**: ≥ 80%
- **Tests**: Todos pasando
- **Race conditions**: 0 detectadas
- **Performance**: Benchmarks estables

### 🏆 Targets Excelencia
- **Coverage**: ≥ 90%
- **Property-based tests**: Implementados
- **Mutation testing**: Sin mutantes sobrevivientes
- **CI/CD**: Pipeline completo configurado

## 🛠️ Troubleshooting

### ❌ "Test file not found"
```bash
# Asegúrate de que los archivos terminan en _test.go
ls -la *_test.go
```

### ❌ "Race condition detected"
```bash
# Ejecuta con race detection para ver detalles
go test -race -v ./...
```

### ❌ "Coverage too low"
```bash
# Ver qué líneas no están cubiertas
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### ❌ "Tests timeout"
```bash
# Aumentar timeout
go test -timeout 60s ./...

# Ejecutar solo tests rápidos
go test -short ./...
```

### ❌ "Import cycle"
```bash
# Verificar dependencies
go mod graph | grep cyclic
```

## 🚀 Siguientes Pasos

1. **Completar todos los ejercicios** siguiendo los ejemplos
2. **Mejorar coverage** añadiendo más casos de prueba
3. **Optimizar performance** usando los benchmarks
4. **Configurar CI/CD** con GitHub Actions
5. **Explorar fuzzing** para encontrar edge cases
6. **Implementar mutation testing** para validar quality

## 📖 Recursos Adicionales

- [Testing Package Documentation](https://golang.org/pkg/testing/)
- [Go Testing By Example](https://github.com/golang/go/wiki/TableDrivenTests)
- [Testify Framework](https://github.com/stretchr/testify)
- [Go Fuzzing](https://go.dev/security/fuzz/)
- [Benchmark Comparison Tools](https://pkg.go.dev/golang.org/x/perf)

---

**¡Feliz Testing! 🧪✨**

> *"Los tests no son solo código que verifica otro código, son la documentación ejecutable de lo que tu software debería hacer."*
