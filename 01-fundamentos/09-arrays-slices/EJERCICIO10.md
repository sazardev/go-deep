# 🔟 Ejercicio 10: Algoritmos de Cadenas con Slices

## 📋 Resumen

El **Ejercicio 10** es el ejercicio más avanzado de la lección de Arrays y Slices, enfocándose en algoritmos sofisticados para procesamiento de cadenas usando slices como estructura de datos fundamental.

## 🎯 Objetivos

- **Implementar algoritmos clásicos** de ciencias de la computación
- **Optimizar operaciones** con cadenas usando slices
- **Aplicar programación dinámica** para resolver problemas complejos
- **Entender análisis de complejidad** temporal y espacial
- **Resolver problemas reales** de procesamiento de texto

## 🧩 Algoritmos Implementados

### 1. 🔍 Algoritmo KMP (Knuth-Morris-Pratt)

**Propósito**: Búsqueda eficiente de patrones en texto.

**Complejidad**: O(n + m) donde n = longitud del texto, m = longitud del patrón

#### Características Técnicas:
- **Tabla LPS**: Longest Proper Prefix que es también Suffix
- **Sin backtracking**: Evita volver atrás en el texto
- **Optimización**: Reutiliza información de coincidencias parciales

#### Implementación:
```go
func buscarPatronKMP(texto, patron string) []int
func construirTablaLPS(patron string) []int
```

#### Casos de Uso:
- Búsqueda en editores de texto
- Análisis de secuencias biológicas
- Detección de plagios
- Compresión de datos

### 2. 🧬 Algoritmo LCS (Longest Common Subsequence)

**Propósito**: Encontrar la subsecuencia común más larga entre dos cadenas.

**Complejidad**: O(n * m) donde n, m = longitudes de las cadenas

#### Características Técnicas:
- **Programación dinámica**: Tabla bidimensional para almacenar resultados
- **Reconstrucción**: Backtracking para obtener la secuencia real
- **Optimización de espacio**: Posible reducir a O(min(n,m))

#### Implementación:
```go
func subsecuenciaComunMasLarga(s1, s2 string) string
```

#### Casos de Uso:
- Comparación de secuencias de ADN
- Control de versiones (diff)
- Detección de similitudes en texto
- Análisis filogenético

## 🔧 Algoritmos Complementarios

### 3. 🔄 Palíndromo Más Largo
- **Técnica**: Expansión desde el centro
- **Complejidad**: O(n²)
- **Aplicación**: Análisis de simetría en datos

### 4. 📏 Distancia de Levenshtein
- **Técnica**: Programación dinámica
- **Complejidad**: O(n * m)
- **Aplicación**: Corrección ortográfica, bioinformática

### 5. 🔀 Detección de Anagramas
- **Técnica**: Mapas de frecuencia
- **Complejidad**: O(n)
- **Aplicación**: Juegos de palabras, análisis lingüístico

### 6. 📝 Generación de Subcadenas
- **Técnica**: Doble bucle anidado
- **Complejidad**: O(n²)
- **Aplicación**: Análisis exhaustivo de patrones

## 📊 Resultados de la Demostración

### Algoritmo KMP
```
Texto: ABABDABACDABABCABCABCABCABC
Patrón: ABCABC
Tabla LPS: [0 0 0 1 2 3]
Ocurrencias: [12 15 18 21]
```

### Algoritmo LCS
```
'ABCDGH' ∩ 'AEDFHR' = 'ADH' (longitud: 3)
'PROGRAMMING' ∩ 'ALGORITHM' = 'GRI' (27.3% similitud)
```

### Algoritmos Adicionales
```
Palíndromo más largo de 'babad': 'bab'
Distancia Levenshtein 'kitten' ↔ 'sitting': 3
Anagramas: 'listen' ↔ 'silent' ✅
```

## ⏱️ Análisis de Rendimiento

### Benchmarks Ejecutados

| Algoritmo | Tamaño Input | Tiempo | Complejidad |
|-----------|--------------|--------|-------------|
| KMP | 110,013 chars | 278.89µs | O(n + m) |
| strings.Index | 110,013 chars | 1.112µs | O(n * m) worst case |
| LCS | 800 x 800 chars | 4.294ms | O(n * m) |

### Observaciones:
- **KMP vs strings.Index**: Go's strings.Index está altamente optimizado para casos promedio
- **LCS**: Escalabilidad cuadrática, crítica para inputs grandes
- **Memoria**: KMP usa O(m) extra, LCS usa O(n * m)

## 🎓 Conceptos Pedagógicos Demostrados

### 1. **Programación Dinámica**
- Descomposición en subproblemas
- Memoización de resultados
- Reconstrucción de soluciones

### 2. **Análisis de Algoritmos**
- Complejidad temporal y espacial
- Trade-offs entre tiempo y memoria
- Benchmarking práctico

### 3. **Manipulación Avanzada de Slices**
- Slices como buffer de caracteres
- Operaciones in-place vs creación de nuevos slices
- Manejo eficiente de memoria

### 4. **Patrones de Diseño**
- Separación de construcción y búsqueda (KMP)
- Builder pattern para reconstrucción (LCS)
- Factory methods para algoritmos parametrizables

## 💡 Aplicaciones Prácticas

### 🧬 Bioinformática
- **Alineamiento de secuencias**: LCS para comparar ADN/proteínas
- **Búsqueda de motivos**: KMP para encontrar secuencias específicas
- **Análisis evolutivo**: Distancia de edición entre especies

### 📝 Procesamiento de Texto
- **Búsqueda en documentos**: KMP para encontrar términos
- **Detección de plagios**: LCS para medir similitud
- **Corrección automática**: Levenshtein para sugerencias

### 🔍 Análisis de Datos
- **Pattern matching**: KMP en logs y streams
- **Clustering textual**: LCS para agrupar documentos similares
- **Compresión**: Identificación de patrones repetitivos

### 🎮 Aplicaciones Lúdicas
- **Juegos de palabras**: Anagramas y palíndromos
- **Puzzles**: Búsqueda de patrones en matrices
- **Generadores**: Creación de contenido procedural

## 📚 Extensiones Posibles

### 1. **Optimizaciones Avanzadas**
- LCS con optimización de espacio O(min(n,m))
- KMP con múltiples patrones (Aho-Corasick)
- Algoritmos aproximados para cadenas grandes

### 2. **Paralelización**
- División del texto para búsqueda paralela
- LCS paralelo usando programación dinámica diagonal
- Map-Reduce para procesamiento masivo

### 3. **Algoritmos Relacionados**
- Boyer-Moore para búsqueda de patrones
- Suffix arrays y suffix trees
- Rolling hash para comparaciones rápidas

### 4. **Aplicaciones Específicas**
- Análisis de sentimientos usando LCS
- Compresión de texto con patterns KMP
- Indexación de documentos para búsqueda

## 🔬 Análisis Técnico

### Ventajas del Enfoque con Slices

1. **Eficiencia de Memoria**
   - Reutilización de buffers
   - Evitar copias innecesarias
   - Gestión automática de capacidad

2. **Flexibilidad**
   - Fácil manipulación de subcadenas
   - Operaciones in-place cuando es posible
   - Interfaz consistente con otros tipos Go

3. **Rendimiento**
   - Acceso directo por índice O(1)
   - Operaciones de slice muy optimizadas
   - Integración natural con strings

### Consideraciones de Diseño

1. **Trade-offs Tiempo vs Espacio**
   - KMP: O(m) espacio extra para O(n+m) tiempo
   - LCS: O(n*m) espacio para evitar recalcular

2. **Robustez**
   - Validación de inputs
   - Manejo de casos edge
   - Gestión de memoria en inputs grandes

3. **Mantenibilidad**
   - Código auto-documentado
   - Separación de responsabilidades
   - Tests exhaustivos

## 🎯 Métricas de Éxito

Al completar este ejercicio, habrás demostrado:

- ✅ **Implementación correcta** de algoritmos clásicos
- ✅ **Optimización práctica** usando slices eficientemente
- ✅ **Análisis crítico** de complejidad y rendimiento
- ✅ **Aplicación creativa** a problemas reales
- ✅ **Código profesional** con buenas prácticas

## 🔗 Referencias y Lectura Adicional

- **Algoritmo KMP**: "Introduction to Algorithms" (CLRS), Capítulo 32
- **LCS**: "Dynamic Programming" patterns y aplicaciones
- **String Algorithms**: "Algorithms on Strings, Trees and Sequences" - Dan Gusfield
- **Go Performance**: "The Go Programming Language" - Donovan & Kernighan

---

Este ejercicio representa la culminación del aprendizaje de arrays y slices en Go, demostrando cómo estructuras de datos simples pueden resolver problemas algorítmicos complejos con elegancia y eficiencia.
