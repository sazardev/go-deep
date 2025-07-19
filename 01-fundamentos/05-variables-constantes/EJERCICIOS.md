# 🧪 Ejercicios: Variables y Constantes

> **Instrucciones**: Copia el código de cada ejercicio en un archivo `.go` separado y completa los TODOs.

## 📋 Lista de Ejercicios

- [x] **Ejercicio 1**: Declaraciones de Variables
- [x] **Ejercicio 2**: Scope y Shadowing  
- [x] **Ejercicio 3**: Zero Values
- [x] **Ejercicio 4**: Constantes con iota
- [x] **Ejercicio 5**: Bit Flags para Permisos
- [x] **Ejercicio 6**: Zero Values en Structs
- [x] **Ejercicio 7**: Constantes Complejas
- [x] **Ejercicio 8**: Intercambio de Variables
- [x] **Ejercicio 9**: Tipos Custom
- [x] **Ejercicio 10**: Sistema de Configuración

---

## 🎯 Ejercicio 1: Declaraciones de Variables

```go
package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicio 1: Declaraciones de Variables ===")
	
	// TODO: Completa estas declaraciones:
	
	// 1. Variable 'nombre' tipo string usando var (sin inicializar)
	// var nombre string
	
	// 2. Variable 'edad' tipo int inicializada a 25 usando var  
	// var edad int = 25
	
	// 3. Variable 'activo' tipo bool usando := con valor true
	// activo := true
	
	// 4. Variables múltiples: x, y, z con valores 1, 2, 3
	// x, y, z := 1, 2, 3
	
	// TODO: Descomena cuando tengas las variables:
	// fmt.Printf("Nombre: '%s', Edad: %d, Activo: %t\n", nombre, edad, activo)
	// fmt.Printf("Coordenadas: x=%d, y=%d, z=%d\n", x, y, z)
	
	fmt.Println("✏️ Completa las declaraciones arriba")
}
```

**Objetivo**: Practicar diferentes formas de declarar variables en Go.

---

## 🔍 Ejercicio 2: Scope y Shadowing

```go
package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicio 2: Scope y Shadowing ===")
	
	mensaje := "exterior"
	fmt.Printf("Mensaje inicial: %s\n", mensaje)
	
	// TODO: Crea un bloque {} y dentro:
	// 1. Declara una nueva variable 'mensaje' con valor "interior"
	// 2. Imprime el mensaje dentro del bloque
	// 3. Observa cómo cambia el scope
	
	// Tu código aquí:
	
	fmt.Printf("Mensaje final: %s\n", mensaje)
}
```

**Objetivo**: Entender scope y variable shadowing.

---

## 🔢 Ejercicio 3: Zero Values

```go
package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicio 3: Zero Values ===")
	
	// TODO: Declara variables sin inicializar y muestra sus zero values:
	// var b bool
	// var i int  
	// var f float64
	// var s string
	// var slice []int
	// var m map[string]int
	// var ptr *int
	
	// TODO: Imprime cada variable mostrando su zero value:
	// fmt.Printf("bool: %t\n", b)
	// fmt.Printf("int: %d\n", i)
	// ... etc
	
	fmt.Println("✏️ Declara variables y muestra sus zero values")
}
```

**Objetivo**: Comprender los zero values de diferentes tipos.

---

## 📊 Ejercicio 4: Constantes con iota

```go
package main

import "fmt"

// TODO: Define aquí las constantes para días de la semana:
// const (
//     Domingo = iota
//     Lunes
//     Martes
//     Miercoles
//     Jueves  
//     Viernes
//     Sabado
// )

// TODO: Define constantes para tamaños de archivo:
// const (
//     Byte = 1
//     KB   = 1024 * Byte
//     MB   = 1024 * KB
//     GB   = 1024 * MB
// )

func main() {
	fmt.Println("=== Ejercicio 4: Constantes con iota ===")
	
	// TODO: Descomena cuando tengas las constantes:
	// fmt.Printf("Miércoles: %d\n", Miercoles)
	// fmt.Printf("1 MB: %d bytes\n", MB)
	// fmt.Printf("1 GB: %d bytes\n", GB)
}
```

**Objetivo**: Dominar el uso de `iota` para generar constantes.

---

## 🔐 Ejercicio 5: Bit Flags para Permisos

```go
package main

import "fmt"

// TODO: Define el tipo y constantes:
// type Permission int
// const (
//     Lectura Permission = 1 << iota
//     Escritura
//     Ejecucion  
//     Eliminacion
// )

// TODO: Función para verificar permisos:
// func tienePermiso(permisos, permiso Permission) bool {
//     return permisos&permiso != 0
// }

func main() {
	fmt.Println("=== Ejercicio 5: Sistema de Permisos ===")
	
	// TODO: Crea roles combinando permisos:
	// usuario := Lectura | Escritura
	// moderador := Lectura | Escritura | Eliminacion
	// admin := Lectura | Escritura | Ejecucion | Eliminacion
	
	// TODO: Prueba los permisos:
	// fmt.Printf("Usuario puede leer: %t\n", tienePermiso(usuario, Lectura))
	// fmt.Printf("Usuario puede eliminar: %t\n", tienePermiso(usuario, Eliminacion))
}
```

**Objetivo**: Implementar sistema de permisos con bit flags.

---

## 📦 Ejercicio 6: Zero Values en Structs

```go
package main

import "fmt"

// TODO: Define el struct Contador:
// type Contador struct {
//     valor int
//     items []string  
// }

// TODO: Implementa métodos:
// func (c *Contador) Incrementar() { c.valor++ }
// func (c *Contador) Valor() int { return c.valor }
// func (c *Contador) AgregarItem(item string) { 
//     c.items = append(c.items, item) 
// }
// func (c *Contador) Items() []string { return c.items }

func main() {
	fmt.Println("=== Ejercicio 6: Zero Values en Structs ===")
	
	// TODO: Usa el struct sin inicialización:
	// var contador Contador
	// fmt.Printf("Valor inicial: %d\n", contador.Valor())
	// contador.Incrementar()
	// contador.Incrementar()
	// contador.AgregarItem("test")
	// fmt.Printf("Valor final: %d\n", contador.Valor())
	// fmt.Printf("Items: %v\n", contador.Items())
}
```

**Objetivo**: Crear structs que aprovechen zero values efectivamente.

---

## ⚡ Ejercicio 7: Constantes Complejas

```go
package main

import "fmt"

// TODO: Define constantes con expresiones:
// const (
//     SegundosEnDia = 60 * 60 * 24
//     MilisegundosEnHora = 60 * 60 * 1000
//     BytesEnKB = 1024
//     BytesEnMB = BytesEnKB * 1024
//     BytesEnGB = BytesEnMB * 1024
// )

func main() {
	fmt.Println("=== Ejercicio 7: Constantes Complejas ===")
	
	// TODO: Muestra los cálculos:
	// fmt.Printf("Segundos en un día: %d\n", SegundosEnDia)
	// fmt.Printf("Milisegundos en una hora: %d\n", MilisegundosEnHora)
	// fmt.Printf("1 GB en bytes: %d\n", BytesEnGB)
	
	// TODO: Calcula cuántas horas de video caben en 1GB (1MB por minuto):
	// minutosEnGB := BytesEnGB / BytesEnMB
	// horasEnGB := minutosEnGB / 60
	// fmt.Printf("Horas de video en 1GB: %d\n", horasEnGB)
}
```

**Objetivo**: Usar expresiones matemáticas en constantes.

---

## 🔄 Ejercicio 8: Intercambio de Variables

```go
package main

import "fmt"

func main() {
	fmt.Println("=== Ejercicio 8: Intercambio de Variables ===")
	
	a, b := 10, 20
	fmt.Printf("Antes: a=%d, b=%d\n", a, b)
	
	// TODO: Método 1 - con variable temporal:
	// temp := a
	// a = b  
	// b = temp
	// fmt.Printf("Después método 1: a=%d, b=%d\n", a, b)
	
	// TODO: Método 2 - asignación múltiple:
	// a, b = b, a
	// fmt.Printf("Después método 2: a=%d, b=%d\n", a, b)
	
	// TODO: Prueba con 3 variables (rotación):
	// x, y, z := 1, 2, 3
	// fmt.Printf("Antes rotación: x=%d, y=%d, z=%d\n", x, y, z)
	// x, y, z = z, x, y  // x->y, y->z, z->x
	// fmt.Printf("Después rotación: x=%d, y=%d, z=%d\n", x, y, z)
}
```

**Objetivo**: Practicar asignación múltiple y variables temporales.

---

## 🌡️ Ejercicio 9: Tipos Custom

```go
package main

import "fmt"

// TODO: Define tipos custom:
// type TemperaturaType float64
// type EstadoType string

// TODO: Define constantes:
// const (
//     Frio   EstadoType = "frío"
//     Tibio  EstadoType = "tibio" 
//     Calido EstadoType = "cálido"
// )

// TODO: Implementa método:
// func (t TemperaturaType) ToFahrenheit() TemperaturaType {
//     return t*9/5 + 32
// }

// func (t TemperaturaType) Estado() EstadoType {
//     switch {
//     case t < 10:
//         return Frio
//     case t < 25:
//         return Tibio
//     default:
//         return Calido
//     }
// }

func main() {
	fmt.Println("=== Ejercicio 9: Tipos Custom ===")
	
	// TODO: Usa los tipos:
	// var temp TemperaturaType = 25.0
	// fmt.Printf("%.1f°C = %.1f°F (%s)\n", 
	//     temp, temp.ToFahrenheit(), temp.Estado())
	
	// TODO: Prueba diferentes temperaturas:
	// temperaturas := []TemperaturaType{-5, 5, 15, 25, 35}
	// for _, t := range temperaturas {
	//     fmt.Printf("%.0f°C = %.1f°F (%s)\n", 
	//         t, t.ToFahrenheit(), t.Estado())
	// }
}
```

**Objetivo**: Crear tipos custom con métodos y constantes.

---

## ⚙️ Ejercicio 10: Sistema de Configuración

```go
package main

import (
	"fmt"
	"time"
)

// TODO: Define struct de configuración:
// type Configuracion struct {
//     Puerto   int               // 0 = puerto automático
//     Debug    bool              // false = modo producción
//     Timeout  time.Duration     // 0 = sin timeout
//     Hosts    []string          // nil = todos los hosts
//     Features map[string]bool   // nil = sin features
//     LogLevel string            // "" = nivel por defecto
// }

// TODO: Implementa métodos:
// func (c *Configuracion) CargarDefaults() {
//     if c.Puerto == 0 {
//         c.Puerto = 8080
//     }
//     if c.Timeout == 0 {
//         c.Timeout = 30 * time.Second
//     }
//     if c.LogLevel == "" {
//         c.LogLevel = "INFO"
//     }
//     if c.Hosts == nil {
//         c.Hosts = []string{"localhost"}
//     }
// }

// func (c *Configuracion) HabilitarFeature(feature string) {
//     if c.Features == nil {
//         c.Features = make(map[string]bool)
//     }
//     c.Features[feature] = true
// }

// func (c *Configuracion) String() string {
//     return fmt.Sprintf("Config{Puerto: %d, Debug: %t, Timeout: %v, LogLevel: %s}",
//         c.Puerto, c.Debug, c.Timeout, c.LogLevel)
// }

func main() {
	fmt.Println("=== Ejercicio 10: Sistema de Configuración ===")
	
	// TODO: Demuestra zero values:
	// var config Configuracion
	// fmt.Printf("Inicial: %s\n", config.String())
	
	// config.CargarDefaults()
	// fmt.Printf("Con defaults: %s\n", config.String())
	
	// config.HabilitarFeature("cache")
	// config.HabilitarFeature("metrics")
	// config.Debug = true
	// fmt.Printf("Final: %s\n", config.String())
}
```

**Objetivo**: Crear un sistema robusto que aproveche zero values.

---

## 🎯 Consejos para Completar

1. **Ejecuta frecuentemente**: Compila y ejecuta después de cada cambio
2. **Lee los errores**: Go tiene mensajes de error muy descriptivos
3. **Experimenta**: Prueba variaciones de cada ejercicio
4. **Compara**: Mira `soluciones.go` si te atascas

## 🏆 Criterios de Éxito

- ✅ **Todos los ejercicios compilan** sin errores
- ✅ **La salida es correcta** comparada con las soluciones
- ✅ **Entiendes cada concepto** que practicaste
- ✅ **Puedes explicar** las diferencias entre var y :=
- ✅ **Sabes cuándo usar** cada forma de declaración

---

*¡Practica hasta dominar cada concepto! 💪*