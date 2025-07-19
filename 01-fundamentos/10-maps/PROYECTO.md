# 🗺️ Proyecto Integrador: Sistema de Gestión de Biblioteca Digital

## 📋 Descripción del Proyecto

Desarrollarás un **Sistema de Gestión de Biblioteca Digital** que utiliza maps de manera intensiva para gestionar libros, usuarios, préstamos y estadísticas. Este proyecto integra todos los conceptos aprendidos sobre maps en Go.

## 🎯 Objetivos de Aprendizaje

- Aplicar maps en un sistema completo de gestión de datos
- Implementar índices múltiples para búsquedas eficientes  
- Crear un sistema de caché para optimizar consultas frecuentes
- Desarrollar un sistema de configuración jerárquico
- Manejar concurrencia segura con maps thread-safe
- Generar estadísticas y reportes usando agrupaciones

## 🏗️ Arquitectura del Sistema

```
Sistema de Biblioteca Digital
├── Gestión de Libros
│   ├── Catálogo principal (ID -> Libro)
│   ├── Índices por título, autor, género, ISBN
│   └── Sistema de tags y categorías
├── Gestión de Usuarios
│   ├── Base de usuarios (ID -> Usuario)
│   ├── Índices por email, nombre, tipo
│   └── Historial de actividades
├── Sistema de Préstamos
│   ├── Préstamos activos (ID -> Préstamo)
│   ├── Historial de préstamos
│   └── Cola de reservas
├── Búsquedas y Filtros
│   ├── Índice invertido para búsqueda de texto
│   ├── Filtros por múltiples criterios
│   └── Búsquedas combinadas (AND/OR)
├── Sistema de Caché
│   ├── Caché de búsquedas frecuentes
│   ├── Caché de estadísticas
│   └── Invalidación inteligente
├── Estadísticas y Reportes
│   ├── Agrupaciones por período, género, autor
│   ├── Rankings de popularidad
│   └── Métricas de uso
└── Configuración
    ├── Configuración por ambiente
    ├── Parámetros del sistema
    └── Límites y quotas
```

## 📚 Estructura de Datos

### Entidades Principales

```go
type Libro struct {
    ID          string
    ISBN        string
    Titulo      string
    Autores     []string
    Generos     []string
    Editorial   string
    AnoPublic   int
    Paginas     int
    Idioma      string
    Descripcion string
    Tags        []string
    Disponible  bool
    Ubicacion   string
    Estado      string // "nuevo", "bueno", "regular", "malo"
    FechaAdq    time.Time
    Precio      float64
    Metadata    map[string]interface{}
}

type Usuario struct {
    ID           string
    Email        string
    Nombre       string
    Apellido     string
    Telefono     string
    Direccion    string
    FechaNac     time.Time
    FechaReg     time.Time
    TipoUsuario  string // "estudiante", "profesor", "externo"
    Activo       bool
    LimiteLibros int
    Multas       float64
    Historial    []string // IDs de préstamos
    Preferencias map[string]interface{}
}

type Prestamo struct {
    ID            string
    UsuarioID     string
    LibroID       string
    FechaPrestamo time.Time
    FechaVenc     time.Time
    FechaDevol    *time.Time
    Estado        string // "activo", "devuelto", "vencido", "renovado"
    Renovaciones  int
    Multa         float64
    Notas         string
}

type Reserva struct {
    ID        string
    UsuarioID string
    LibroID   string
    FechaRes  time.Time
    Prioridad int
    Estado    string // "pendiente", "disponible", "cancelada", "cumplida"
}
```

## 🔧 Funcionalidades Requeridas

### 1. Gestión de Libros (25 puntos)

```go
// Operaciones CRUD básicas
func (b *Biblioteca) AgregarLibro(libro Libro) error
func (b *Biblioteca) ActualizarLibro(id string, libro Libro) error
func (b *Biblioteca) EliminarLibro(id string) error
func (b *Biblioteca) ObtenerLibro(id string) (Libro, bool)

// Búsquedas por índices
func (b *Biblioteca) BuscarPorTitulo(titulo string) []Libro
func (b *Biblioteca) BuscarPorAutor(autor string) []Libro
func (b *Biblioteca) BuscarPorGenero(genero string) []Libro
func (b *Biblioteca) BuscarPorISBN(isbn string) (Libro, bool)

// Búsqueda de texto completo
func (b *Biblioteca) BuscarTexto(query string) []Libro
func (b *Biblioteca) BuscarAvanzada(criterios map[string]interface{}) []Libro

// Gestión de disponibilidad
func (b *Biblioteca) MarcarDisponible(id string) error
func (b *Biblioteca) MarcarNoDisponible(id string) error
func (b *Biblioteca) LibrosDisponibles() []Libro
```

### 2. Gestión de Usuarios (20 puntos)

```go
// CRUD de usuarios
func (b *Biblioteca) RegistrarUsuario(usuario Usuario) error
func (b *Biblioteca) ActualizarUsuario(id string, usuario Usuario) error
func (b *Biblioteca) DesactivarUsuario(id string) error
func (b *Biblioteca) ObtenerUsuario(id string) (Usuario, bool)

// Búsquedas de usuarios
func (b *Biblioteca) BuscarUsuarioPorEmail(email string) (Usuario, bool)
func (b *Biblioteca) UsuariosPorTipo(tipo string) []Usuario
func (b *Biblioteca) UsuariosActivos() []Usuario

// Gestión de multas y límites
func (b *Biblioteca) AgregarMulta(usuarioID string, monto float64) error
func (b *Biblioteca) PagarMulta(usuarioID string, monto float64) error
func (b *Biblioteca) VerificarLimites(usuarioID string) (bool, string)
```

### 3. Sistema de Préstamos (25 puntos)

```go
// Operaciones de préstamo
func (b *Biblioteca) CrearPrestamo(usuarioID, libroID string) (string, error)
func (b *Biblioteca) DevolverLibro(prestamoID string) error
func (b *Biblioteca) RenovarPrestamo(prestamoID string) error

// Gestión de reservas
func (b *Biblioteca) CrearReserva(usuarioID, libroID string) (string, error)
func (b *Biblioteca) CancelarReserva(reservaID string) error
func (b *Biblioteca) ProcesarReservas() []Reserva

// Consultas de préstamos
func (b *Biblioteca) PrestamosActivos(usuarioID string) []Prestamo
func (b *Biblioteca) PrestamosVencidos() []Prestamo
func (b *Biblioteca) HistorialPrestamos(usuarioID string) []Prestamo

// Cálculo de multas
func (b *Biblioteca) CalcularMultas() map[string]float64
func (b *Biblioteca) GenerarNotificaciones() []Notificacion
```

### 4. Sistema de Búsqueda y Filtros (15 puntos)

```go
// Índice invertido para búsqueda de texto
func (b *Biblioteca) IndexarLibros() error
func (b *Biblioteca) BusquedaBooleana(query string) []Libro // AND, OR, NOT
func (b *Biblioteca) BusquedaFuzzy(term string) []Libro

// Filtros combinados
func (b *Biblioteca) FiltrarLibros(filtros map[string]interface{}) []Libro
func (b *Biblioteca) OrdenarResultados(libros []Libro, criterio string) []Libro

// Autocompletado y sugerencias
func (b *Biblioteca) Autocompletar(prefijo string) []string
func (b *Biblioteca) LibrosSimilares(libroID string) []Libro
func (b *Biblioteca) SugerenciasPorUsuario(usuarioID string) []Libro
```

### 5. Estadísticas y Reportes (15 puntos)

```go
// Estadísticas de uso
func (b *Biblioteca) EstadisticasPorPeriodo(inicio, fin time.Time) EstadisticasPeriodo
func (b *Biblioteca) LibrosMasPopulares(limite int) []LibroPopularidad
func (b *Biblioteca) UsuariosMasActivos(limite int) []UsuarioActividad

// Reportes por categorías
func (b *Biblioteca) ReportePorGenero() map[string]ReporteGenero
func (b *Biblioteca) ReportePorAutor() map[string]ReporteAutor
func (b *Biblioteca) ReporteInventario() ReporteInventario

// Métricas del sistema
func (b *Biblioteca) TasaDevolucion() float64
func (b *Biblioteca) TiempoPromedioRetencion() time.Duration
func (b *Biblioteca) AnalisisMultas() AnalisisMultas
```

## 🎯 Implementación Detallada

### Fase 1: Estructura Base (20%)

```go
type Biblioteca struct {
    // Almacenamiento principal
    libros   map[string]Libro
    usuarios map[string]Usuario
    prestamos map[string]Prestamo
    reservas  map[string]Reserva
    
    // Índices para búsquedas rápidas
    indicesTitulo map[string][]string  // título -> []libro_id
    indicesAutor  map[string][]string  // autor -> []libro_id
    indicesGenero map[string][]string  // género -> []libro_id
    indicesISBN   map[string]string    // isbn -> libro_id
    
    indicesEmail     map[string]string   // email -> usuario_id
    indicesTipoUser  map[string][]string // tipo -> []usuario_id
    
    // Índice invertido para búsqueda de texto
    indiceInvertido map[string]map[string]int // palabra -> libro_id -> frecuencia
    
    // Sistema de caché
    cache *CacheBiblioteca
    
    // Configuración
    config *ConfigBiblioteca
    
    // Thread safety
    mu sync.RWMutex
    
    // Contadores para IDs únicos
    contadorLibros   int64
    contadorUsuarios int64
    contadorPrestamos int64
    contadorReservas int64
}

type CacheBiblioteca struct {
    busquedas    map[string][]Libro
    estadisticas map[string]interface{}
    ttl          time.Duration
    lastUpdate   map[string]time.Time
    mu           sync.RWMutex
}

type ConfigBiblioteca struct {
    MaxPrestamosEstudiante int
    MaxPrestamosProfesor   int
    MaxPrestamosExterno    int
    DiasPrestamoEstudiante int
    DiasPrestamoProfesor   int
    DiasPrestamoExterno    int
    MultaDiaria            float64
    MaxRenovaciones        int
    DiasReserva            int
}
```

### Fase 2: Operaciones CRUD (30%)

Implementa todas las operaciones básicas de creación, lectura, actualización y eliminación para libros, usuarios, préstamos y reservas.

### Fase 3: Sistema de Índices y Búsquedas (25%)

Desarrolla el sistema de índices múltiples y el motor de búsqueda con soporte para:
- Búsquedas exactas por campo
- Búsqueda de texto completo
- Búsquedas booleanas (AND, OR, NOT)
- Filtros combinados
- Autocompletado

### Fase 4: Lógica de Negocio (15%)

Implementa las reglas de negocio:
- Validación de límites de préstamo
- Cálculo automático de multas
- Gestión de reservas
- Notificaciones automáticas

### Fase 5: Estadísticas y Reportes (10%)

Desarrolla el sistema de reporting con agrupaciones complejas y métricas del sistema.

## 🧪 Casos de Prueba

### Test Case 1: Gestión Básica de Libros
```go
func TestGestionLibros(t *testing.T) {
    biblioteca := NewBiblioteca()
    
    // Agregar libros
    libro1 := Libro{
        Titulo:  "El Quijote",
        Autores: []string{"Miguel de Cervantes"},
        Generos: []string{"Novela", "Clásico"},
        ISBN:    "978-84-376-0494-7",
    }
    
    err := biblioteca.AgregarLibro(libro1)
    assert.NoError(t, err)
    
    // Buscar por título
    resultados := biblioteca.BuscarPorTitulo("Quijote")
    assert.Len(t, resultados, 1)
    assert.Equal(t, "El Quijote", resultados[0].Titulo)
}
```

### Test Case 2: Sistema de Préstamos
```go
func TestSistemaPrestamos(t *testing.T) {
    biblioteca := setupBibliotecaConDatos()
    
    // Crear préstamo
    prestamoID, err := biblioteca.CrearPrestamo("user1", "book1")
    assert.NoError(t, err)
    assert.NotEmpty(t, prestamoID)
    
    // Verificar disponibilidad del libro
    libro, _ := biblioteca.ObtenerLibro("book1")
    assert.False(t, libro.Disponible)
    
    // Devolver libro
    err = biblioteca.DevolverLibro(prestamoID)
    assert.NoError(t, err)
    
    // Verificar disponibilidad restaurada
    libro, _ = biblioteca.ObtenerLibro("book1")
    assert.True(t, libro.Disponible)
}
```

### Test Case 3: Búsquedas Avanzadas
```go
func TestBusquedasAvanzadas(t *testing.T) {
    biblioteca := setupBibliotecaConDatos()
    
    // Búsqueda booleana
    resultados := biblioteca.BusquedaBooleana("programación AND go")
    assert.NotEmpty(t, resultados)
    
    // Filtros combinados
    filtros := map[string]interface{}{
        "genero":     "Tecnología",
        "ano_desde":  2020,
        "disponible": true,
    }
    resultados = biblioteca.FiltrarLibros(filtros)
    assert.NotEmpty(t, resultados)
}
```

## 📊 Métricas de Evaluación

### Funcionalidad (70%)
- ✅ CRUD completo para todas las entidades (20%)
- ✅ Sistema de índices funcionando (15%)
- ✅ Búsquedas y filtros correctos (15%)
- ✅ Lógica de préstamos completa (10%)
- ✅ Estadísticas y reportes (10%)

### Calidad del Código (20%)
- ✅ Uso apropiado de maps en estructuras de datos
- ✅ Manejo de concurrencia con mutex
- ✅ Gestión de errores consistente
- ✅ Código limpio y bien documentado

### Rendimiento (10%)
- ✅ Búsquedas eficientes O(1) o O(log n)
- ✅ Uso de caché para consultas frecuentes
- ✅ Optimización de memoria con maps

## 🚀 Extensiones Opcionales

### Nivel Avanzado (+20% extra)
1. **Persistencia**: Guardar/cargar datos en JSON/CSV
2. **API REST**: Crear endpoints HTTP para el sistema
3. **Concurrencia Avanzada**: Soporte para múltiples usuarios simultáneos
4. **Sistema de Recomendaciones**: ML básico para sugerencias
5. **Dashboard**: Interfaz web simple para visualizar estadísticas

### Nivel Experto (+30% extra)
1. **Base de Datos**: Integración con PostgreSQL/MongoDB
2. **Microservicios**: Dividir en servicios independientes
3. **Mensajería**: Sistema de notificaciones asíncronas
4. **Métricas**: Integración con Prometheus/Grafana
5. **Testing Avanzado**: Benchmarks y tests de carga

## 📝 Entregables

1. **Código fuente completo** en `proyecto_biblioteca.go`
2. **Tests unitarios** en `proyecto_biblioteca_test.go`
3. **Documentación** detallada en `BIBLIOTECA.md`
4. **Ejemplos de uso** en `ejemplos_biblioteca.go`
5. **Análisis de rendimiento** en `PERFORMANCE.md`

## ⏰ Timeline Sugerido

- **Semana 1**: Estructura base y CRUD básico
- **Semana 2**: Sistema de índices y búsquedas
- **Semana 3**: Lógica de préstamos y reservas
- **Semana 4**: Estadísticas, testing y documentación

## 🎯 Objetivos Específicos por Semana

### Semana 1: Fundación
- ✅ Definir todas las estructuras de datos
- ✅ Implementar operaciones CRUD básicas
- ✅ Crear sistema de IDs únicos
- ✅ Tests unitarios para operaciones básicas

### Semana 2: Búsquedas
- ✅ Implementar todos los índices
- ✅ Sistema de búsqueda de texto completo
- ✅ Filtros y búsquedas combinadas
- ✅ Caché de búsquedas frecuentes

### Semana 3: Lógica de Negocio
- ✅ Sistema completo de préstamos
- ✅ Gestión de reservas y colas
- ✅ Cálculo automático de multas
- ✅ Validaciones y reglas de negocio

### Semana 4: Finalización
- ✅ Estadísticas y reportes completos
- ✅ Suite completa de tests
- ✅ Documentación técnica
- ✅ Análisis de rendimiento

---

**¡Este proyecto te permitirá dominar completamente el uso de maps en Go mientras construyes un sistema real y útil!**

Los maps son la columna vertebral de este sistema, siendo utilizados para almacenamiento principal, índices de búsqueda, caché, configuración y estadísticas. Al completar este proyecto, habrás aplicado todos los patrones y técnicas avanzadas de maps en un contexto práctico y realista.
