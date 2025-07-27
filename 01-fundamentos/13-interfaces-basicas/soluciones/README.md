# 🎯 Soluciones - Lección 13: Interfaces Básicas

Este directorio contiene las **soluciones completas** de todos los ejercicios de la Lección 13.

## 📁 Estructura de la Lección

```
13-interfaces-basicas/
├── README.md              # Tutorial completo de interfaces
├── ejercicios.go          # Ejercicios con TODO para completar
└── soluciones/
    ├── README.md          # Este archivo
    └── main.go            # Soluciones completas
```

## 🚀 Cómo Ejecutar las Soluciones

### Ejecutar Todas las Soluciones
```bash
cd soluciones
go run main.go
```

### Ejecutar Ejercicios Específicos
Edita `main.go` y comenta/descomenta los ejercicios que quieras ejecutar en la función `main()`:

```go
func main() {
    fmt.Println("🎪 Soluciones: Interfaces Básicas en Go")
    fmt.Println("======================================")
    fmt.Println()

    ejercicio1()    // Descomenta para ejecutar
    // ejercicio2() // Comenta para saltar
    ejercicio3()
    // ... resto de ejercicios
}
```

## 📚 Ejercicios Incluidos

### 1. **Interface Básica - Sistema de Formas Geométricas**
- ✅ Interface `Forma` con métodos `Area()`, `Perimetro()` y `Descripcion()`
- ✅ Structs `Rectangulo` y `Circulo` implementando la interface
- ✅ Función polimórfica `MostrarInformacionForma()`

### 2. **Polimorfismo - Sistema de Animales**
- ✅ Interface `Animal` con métodos de comportamiento
- ✅ Implementaciones para `Perro`, `Gato` y `Pajaro`
- ✅ Función `CuidarAnimal()` que funciona con cualquier animal

### 3. **Interfaces Estándar - fmt.Stringer y sort.Interface**
- ✅ Implementación de `fmt.Stringer` para printing customizado
- ✅ Implementación de `sort.Interface` para ordenamiento
- ✅ Sistema de productos con ordenamiento por precio

### 4. **Type Assertions - Procesador de Datos**
- ✅ Switch con type assertions para múltiples tipos
- ✅ Procesamiento inteligente según el tipo de dato
- ✅ Manejo de tipos desconocidos

### 5. **Empty Interface - Sistema de Logging**
- ✅ Interface `Logger` que acepta `interface{}`
- ✅ Implementaciones `ConsoleLogger` y `FileLogger`
- ✅ Función genérica para loggear diferentes tipos de datos

### 6. **Strategy Pattern - Sistema de Descuentos**
- ✅ Interface `EstrategiaDescuento` para diferentes estrategias
- ✅ Múltiples implementaciones: sin descuento, porcentaje, fijo, por cantidad
- ✅ `CalculadoraPrecio` que cambia estrategias dinámicamente

### 7. **Observer Pattern - Sistema de Notificaciones**
- ✅ Interface `Observer` y `Sujeto` para el patrón Observer
- ✅ `GestorEventos` que maneja suscripciones y notificaciones
- ✅ Múltiples observadores: Email, SMS, Analytics

### 8. **Factory Pattern - Conexiones de Base de Datos**
- ✅ Interface `BaseDatos` común para diferentes DB
- ✅ Implementaciones para MySQL, PostgreSQL y MongoDB
- ✅ Factory que crea conexiones según configuración

## 🔍 Comparar con Ejercicios

Para entender mejor las soluciones:

1. **Abre dos ventanas en tu editor:**
   - `../ejercicios.go` (con los TODO)
   - `main.go` (con las soluciones)

2. **Compara lado a lado:**
   - Ve cómo se completan los TODO
   - Observa las implementaciones completas
   - Entiende los patrones aplicados

## 💡 Conceptos Demostrados

- **🔀 Polimorfismo:** Un mismo código funciona con diferentes tipos
- **🧩 Interfaces:** Contratos que definen comportamiento
- **🎯 Type Assertions:** Verificación y conversión de tipos
- **📦 Empty Interface:** Flexibilidad máxima con `interface{}`
- **🏗️ Design Patterns:** Strategy, Observer, Factory usando interfaces

## 🎓 Aprendizajes Clave

1. **Las interfaces en Go son implícitas** - no necesitas declarar que las implementas
2. **Diseña interfaces pequeñas** - muchas interfaces de 1-3 métodos
3. **Acepta interfaces, devuelve structs** - principio de diseño Go
4. **Las interfaces permiten testing fácil** - puedes crear mocks fácilmente

## 🔄 Siguientes Pasos

Después de revisar estas soluciones:

1. **Intenta los ejercicios** por tu cuenta antes de ver las soluciones
2. **Modifica las implementaciones** para añadir nueva funcionalidad
3. **Crea tus propias interfaces** para problemas que tengas
4. **Practica los design patterns** en proyectos reales

## 🆘 Ayuda Adicional

Si tienes dudas sobre alguna implementación:

1. **Lee el tutorial completo** en `../README.md`
2. **Experimenta con el código** - modificalo y ve qué pasa
3. **Consulta la documentación** oficial de Go sobre interfaces
4. **Practica, practica, practica** - las interfaces son la clave del Go idiomático

---

¡Felicidades por completar la lección de interfaces! 🎉

Las interfaces son una de las características más poderosas de Go. Con estas bases sólidas, estás listo para crear código más flexible, testeable y mantenible.
