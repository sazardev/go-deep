# 🎯 PROYECTO: Sistema de Testing Integral

## Descripción del Proyecto

Este proyecto integra todos los conceptos de testing avanzado aprendidos en la lección:

- **TDD (Test-Driven Development)**: Desarrollo guiado por tests
- **Mocking**: Simulación de dependencias externas  
- **Property Testing**: Tests basados en propiedades matemáticas
- **Integration Testing**: Tests de integración con servidores HTTP
- **Benchmark Testing**: Medición de rendimiento
- **Test Suites**: Organización de tests complejos

## 📋 Ejercicios Completados

### ✅ Ejercicio 1: TDD - Calculadora Científica
- Implementación siguiendo el ciclo Red-Green-Refactor
- Tests para operaciones básicas y científicas
- Manejo de errores (división por cero, raíz negativa)
- Historial de operaciones

### ✅ Ejercicio 2: Mocking - Sistema de Notificaciones  
- Interfaces para Email, SMS y Push notifications
- Mocks que capturan llamadas y parámetros
- Simulación de errores en servicios externos
- Verificación de comportamiento de dependencias

### ✅ Ejercicio 3: Property Testing - Lista Ordenada
- Tests basados en invariantes matemáticas
- Verificación que la lista siempre está ordenada
- Tests de idempotencia y ciclos insert/remove
- Generación automática de casos de prueba

### ✅ Ejercicio 4: Integration Testing - API Client
- Tests con servidores HTTP reales usando httptest
- Verificación de serialización/deserialización JSON
- Manejo de errores HTTP (404, 500, timeout)
- Tests end-to-end de flujos completos

## 🚀 Cómo Ejecutar

### Ejecutar el programa principal:
```bash
go run ejercicios.go
```

### Ejecutar tests específicos:
```bash
# Tests de calculadora
go test -v -run TestScientificCalculator

# Tests de notificaciones (mocking)
go test -v -run TestNotificationService

# Property tests de lista ordenada
go test -v -run TestSortedList

# Integration tests de API
go test -v -run TestAPIClient
```

### Ejecutar benchmarks:
```bash
# Benchmarks de calculadora
go test -bench BenchmarkCalculator

# Benchmarks de lista ordenada
go test -bench BenchmarkSortedList

# Benchmarks de API client
go test -bench BenchmarkAPIClient
```

### Ejecutar todos los tests:
```bash
go test -v ./...
```

## 📊 Métricas de Testing

### Cobertura de Código
```bash
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Análisis de Performance
```bash
go test -bench=. -benchmem
go test -cpuprofile=cpu.prof -memprofile=mem.prof -bench=.
```

## 🎓 Conceptos Aplicados

### 1. **Test-Driven Development (TDD)**
```go
// 1. RED: Escribir test que falla
func TestCalculator_Add(t *testing.T) {
    calc := NewScientificCalculator()
    result := calc.Add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %f", result)
    }
}

// 2. GREEN: Implementar código mínimo
func (sc *ScientificCalculator) Add(a, b float64) float64 {
    return a + b
}

// 3. REFACTOR: Mejorar sin romper tests
```

### 2. **Mocking para Aislar Dependencias**
```go
type EmailSenderMock struct {
    calls []EmailCall
    shouldFail bool
}

func (m *EmailSenderMock) SendEmail(to, subject, body string) error {
    m.calls = append(m.calls, EmailCall{To: to, Subject: subject, Body: body})
    if m.shouldFail {
        return errors.New("mock error")
    }
    return nil
}
```

### 3. **Property Testing**
```go
func TestSortedList_AlwaysSorted(t *testing.T) {
    for i := 0; i < 100; i++ {
        list := NewSortedList()
        // Insertar números aleatorios
        for j := 0; j < 20; j++ {
            list.Insert(rand.Intn(100))
            // Property: Lista siempre ordenada
            if !isSorted(list.ToSlice()) {
                t.Error("Lista no está ordenada")
            }
        }
    }
}
```

### 4. **Integration Testing**
```go
func TestAPIClient_Integration(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Simular respuesta de API real
        user := User{ID: 1, Name: "Test User"}
        json.NewEncoder(w).Encode(user)
    }))
    defer server.Close()
    
    client := NewAPIClient(server.URL)
    user, err := client.GetUser(1)
    // Verificar comportamiento completo
}
```

## 📈 Beneficios del Testing Avanzado

### 🔒 **Confiabilidad**
- Detecta regresiones automáticamente
- Garantiza que el código funciona como se espera
- Reduce bugs en producción

### 🚀 **Refactoring Seguro**
- Los tests actúan como red de seguridad
- Permite mejoras sin miedo a romper funcionalidad
- Facilita evolución del código

### 📚 **Documentación Viva**
- Los tests describen el comportamiento esperado
- Ejemplos de uso de las APIs
- Especificaciones ejecutables

### ⚡ **Desarrollo Más Rápido**
- Feedback inmediato sobre cambios
- Menos tiempo debuggeando
- Mayor confianza al hacer cambios

## 🎯 Próximos Pasos

1. **Testify Framework**: Usar bibliotecas de testing más avanzadas
2. **Fuzzing**: Testing con datos aleatorios más sofisticado  
3. **Contract Testing**: Verificar contratos entre servicios
4. **Load Testing**: Tests de carga y estrés
5. **Mutation Testing**: Verificar calidad de los tests

## 💡 Mejores Prácticas Aplicadas

- ✅ Tests independientes y deterministas
- ✅ Nombres descriptivos de tests y funciones
- ✅ Arrange-Act-Assert pattern
- ✅ Tests rápidos y confiables
- ✅ Un assert por test (cuando es posible)
- ✅ Mocks para dependencias externas
- ✅ Property testing para algoritmos complejos
- ✅ Integration tests para flujos críticos
- ✅ Benchmarks para código de alto rendimiento

---

**¡Felicitaciones!** Has completado la lección de Testing Avanzado y tienes una base sólida para escribir tests robustos y maintener código de alta calidad en Go. 🎉
