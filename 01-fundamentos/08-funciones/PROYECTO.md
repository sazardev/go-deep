# 🚀 PROYECTO INTEGRADOR: Sistema de Análisis de Datos

Este proyecto demuestra la aplicación práctica de todos los conceptos de funciones en Go mediante un sistema completo de análisis de datos.

## 📋 Descripción

El Sistema de Análisis de Datos es una aplicación completa que incluye:

- **Carga y procesamiento de datos** desde archivos CSV
- **Sistema de filtros** configurables y componibles
- **Transformaciones** de datos con pipeline personalizable
- **Análisis estadístico** avanzado con métricas descriptivas
- **Visualización básica** en formato texto
- **Sistema de reportes** automáticos
- **Pipeline de procesamiento** funcional

## 🏗️ Arquitectura del Sistema

### Componentes Principales

1. **AnalizadorDatos**: Componente central que gestiona el procesamiento
2. **Sistema de Filtros**: Funciones de orden superior para filtrado
3. **Sistema de Transformaciones**: Pipeline de transformaciones
4. **Motor Estadístico**: Cálculos estadísticos avanzados
5. **Generador de Reportes**: Sistema de visualización y reportes

### Conceptos de Funciones Demostrados

#### 1. Funciones Básicas
```go
func (a *AnalizadorDatos) GenerarDatosPrueba()
func calcularPercentil(sortedValues []float64, percentil float64) float64
```

#### 2. Múltiples Valores de Retorno
```go
func (a *AnalizadorDatos) parsearRegistro(record []string) (Registro, error)
```

#### 3. Funciones Variádicas
```go
func FiltrarPorCiudad(ciudades ...string) func(Registro) bool
```

#### 4. Funciones de Orden Superior
```go
func (a *AnalizadorDatos) AgregarFiltro(filtro func(Registro) bool)
func (p *Pipeline) Filtrar(filtro func(Registro) bool) *Pipeline
```

#### 5. Closures
```go
func FiltrarPorEdad(min, max int) func(Registro) bool {
    return func(r Registro) bool {
        return r.Edad >= min && r.Edad <= max
    }
}
```

#### 6. Métodos en Estructuras
```go
func (a *AnalizadorDatos) CalcularEstadisticasSalario() EstadisticasDescriptivas
func (p *Pipeline) Ejecutar(datos []Registro) []Registro
```

#### 7. Funciones Anónimas
```go
analizador.ConfigurarLogger(func(msg string) {
    fmt.Printf("🔍 [%s] %s\n", time.Now().Format("15:04:05"), msg)
})
```

#### 8. Defer para Recursos
```go
defer archivo.Close()
```

## 🚀 Funcionalidades

### Sistema de Filtros
- Filtro por rango de edad
- Filtro por rango de salario
- Filtro por ciudades específicas
- Filtro por rango de fechas
- Composición de múltiples filtros

### Sistema de Transformaciones
- Bonificaciones automáticas
- Normalización de datos
- Transformaciones personalizadas
- Pipeline configurable

### Análisis Estadístico
- Estadísticas descriptivas completas
- Percentiles y cuartiles
- Análisis de varianza
- Agrupación por categorías

### Visualización
- Histogramas en formato texto
- Reportes detallados
- Métricas de resumen

## 💻 Uso del Sistema

### Ejemplo Básico
```go
// Crear analizador
analizador := NuevoAnalizador()

// Cargar datos
analizador.GenerarDatosPrueba()

// Aplicar filtros
analizador.AgregarFiltro(FiltrarPorEdad(25, 45))
analizador.AgregarFiltro(FiltrarPorSalario(30000, 60000))

// Calcular estadísticas
stats := analizador.CalcularEstadisticasSalario()

// Generar reporte
reporte := analizador.GenerarReporteCompleto()
fmt.Println(reporte)
```

### Pipeline Avanzado
```go
pipeline := NuevoPipeline().
    Filtrar(func(r Registro) bool { return r.Edad >= 30 }).
    Transformar(TransformadorBonificacion(0.1)).
    Ordenar(func(r1, r2 Registro) bool { return r1.Salario > r2.Salario })

resultados := analizador.EjecutarPipeline(pipeline)
```

## 🎯 Objetivos Pedagógicos

Este proyecto demuestra:

1. **Composición de funciones** para crear sistemas complejos
2. **Encapsulación** mediante métodos y closures
3. **Reutilización** con funciones de orden superior
4. **Flexibilidad** con interfaces funcionales
5. **Mantenibilidad** con arquitectura modular
6. **Escalabilidad** con patrones funcionales

## 📊 Tipos de Datos

### Registro
```go
type Registro struct {
    ID      int
    Nombre  string
    Edad    int
    Salario float64
    Ciudad  string
    Fecha   time.Time
}
```

### Estadísticas Descriptivas
```go
type EstadisticasDescriptivas struct {
    Media       float64
    Mediana     float64
    Moda        float64
    Minimo      float64
    Maximo      float64
    Desviacion  float64
    Varianza    float64
    Rango       float64
    Percentil25 float64
    Percentil75 float64
    Conteo      int
}
```

## 🔧 Extensiones Posibles

- Soporte para más formatos de datos (JSON, XML)
- Gráficos reales con bibliotecas externas
- Persistencia en base de datos
- API REST para acceso remoto
- Exportación de reportes a PDF
- Análisis predictivo básico
- Paralelización del procesamiento

## 📚 Conceptos Avanzados Aplicados

1. **Factory Pattern** con constructores
2. **Builder Pattern** con pipeline fluido
3. **Strategy Pattern** con funciones intercambiables
4. **Observer Pattern** con logging configurable
5. **Chain of Responsibility** con filtros encadenados

Este proyecto integra todos los conceptos de funciones vistos en la lección, demostrando su aplicación práctica en un sistema real y complejo.
