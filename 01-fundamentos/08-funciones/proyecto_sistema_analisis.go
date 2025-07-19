// 🚀 PROYECTO INTEGRADOR: SISTEMA DE ANÁLISIS DE DATOS
// ===================================================
//
// Este proyecto demuestra el uso práctico de funciones en Go
// implementando un sistema completo de análisis de datos que incluye:
// - Carga y procesamiento de datos
// - Análisis estadístico avanzado
// - Visualización básica en texto
// - Sistema de reportes
// - Pipeline de transformaciones

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ===== TIPOS Y ESTRUCTURAS =====

// Registro representa una fila de datos
type Registro struct {
	ID       int
	Nombre   string
	Edad     int
	Salario  float64
	Ciudad   string
	Fecha    time.Time
}

// EstadisticasDescriptivas contiene métricas estadísticas
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

// AnalizadorDatos es el componente principal del sistema
type AnalizadorDatos struct {
	datos       []Registro
	filtros     []func(Registro) bool
	transformadores []func(Registro) Registro
	logger      func(string)
}

// ===== CONSTRUCTORES Y CONFIGURACIÓN =====

// NuevoAnalizador crea una nueva instancia del analizador
func NuevoAnalizador() *AnalizadorDatos {
	return &AnalizadorDatos{
		datos:       make([]Registro, 0),
		filtros:     make([]func(Registro) bool, 0),
		transformadores: make([]func(Registro) Registro, 0),
		logger:      func(msg string) { fmt.Printf("[LOG] %s\n", msg) },
	}
}

// ConfigurarLogger permite personalizar el sistema de logging
func (a *AnalizadorDatos) ConfigurarLogger(logger func(string)) {
	a.logger = logger
}

// ===== CARGA DE DATOS =====

// CargarDatosCSV carga datos desde un archivo CSV
func (a *AnalizadorDatos) CargarDatosCSV(nombreArchivo string) error {
	a.logger(fmt.Sprintf("Cargando datos desde %s", nombreArchivo))
	
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return fmt.Errorf("error abriendo archivo: %w", err)
	}
	defer archivo.Close()

	reader := csv.NewReader(archivo)
	reader.Comma = ','
	reader.Comment = '#'

	// Leer encabezados
	_, err = reader.Read()
	if err != nil {
		return fmt.Errorf("error leyendo encabezados: %w", err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			a.logger(fmt.Sprintf("Error leyendo fila: %v", err))
			continue
		}

		registro, err := a.parsearRegistro(record)
		if err != nil {
			a.logger(fmt.Sprintf("Error parseando registro: %v", err))
			continue
		}

		a.datos = append(a.datos, registro)
	}

	a.logger(fmt.Sprintf("Cargados %d registros exitosamente", len(a.datos)))
	return nil
}

// parsearRegistro convierte una fila CSV en un Registro
func (a *AnalizadorDatos) parsearRegistro(record []string) (Registro, error) {
	if len(record) < 6 {
		return Registro{}, fmt.Errorf("registro incompleto: %v", record)
	}

	id, err := strconv.Atoi(record[0])
	if err != nil {
		return Registro{}, fmt.Errorf("ID inválido: %s", record[0])
	}

	edad, err := strconv.Atoi(record[2])
	if err != nil {
		return Registro{}, fmt.Errorf("edad inválida: %s", record[2])
	}

	salario, err := strconv.ParseFloat(record[3], 64)
	if err != nil {
		return Registro{}, fmt.Errorf("salario inválido: %s", record[3])
	}

	fecha, err := time.Parse("2006-01-02", record[5])
	if err != nil {
		return Registro{}, fmt.Errorf("fecha inválida: %s", record[5])
	}

	return Registro{
		ID:      id,
		Nombre:  strings.TrimSpace(record[1]),
		Edad:    edad,
		Salario: salario,
		Ciudad:  strings.TrimSpace(record[4]),
		Fecha:   fecha,
	}, nil
}

// GenerarDatosPrueba crea datos de ejemplo para demostración
func (a *AnalizadorDatos) GenerarDatosPrueba() {
	a.logger("Generando datos de prueba")
	
	nombres := []string{"Ana García", "Carlos López", "María Rodríguez", "Juan Pérez", "Laura Martín"}
	ciudades := []string{"Madrid", "Barcelona", "Valencia", "Sevilla", "Bilbao"}
	
	for i := 1; i <= 100; i++ {
		registro := Registro{
			ID:      i,
			Nombre:  nombres[i%len(nombres)] + fmt.Sprintf(" %d", i),
			Edad:    20 + (i%40),
			Salario: 30000 + float64(i%50)*1000,
			Ciudad:  ciudades[i%len(ciudades)],
			Fecha:   time.Date(2020+(i%4), time.Month(1+(i%12)), 1+(i%28), 0, 0, 0, 0, time.UTC),
		}
		a.datos = append(a.datos, registro)
	}
	
	a.logger(fmt.Sprintf("Generados %d registros de prueba", len(a.datos)))
}

// ===== SISTEMA DE FILTROS =====

// AgregarFiltro añade un filtro al pipeline de procesamiento
func (a *AnalizadorDatos) AgregarFiltro(filtro func(Registro) bool) {
	a.filtros = append(a.filtros, filtro)
}

// FiltrarPorEdad crea un filtro por rango de edad
func FiltrarPorEdad(min, max int) func(Registro) bool {
	return func(r Registro) bool {
		return r.Edad >= min && r.Edad <= max
	}
}

// FiltrarPorSalario crea un filtro por rango de salario
func FiltrarPorSalario(min, max float64) func(Registro) bool {
	return func(r Registro) bool {
		return r.Salario >= min && r.Salario <= max
	}
}

// FiltrarPorCiudad crea un filtro por ciudad
func FiltrarPorCiudad(ciudades ...string) func(Registro) bool {
	ciudadSet := make(map[string]bool)
	for _, ciudad := range ciudades {
		ciudadSet[ciudad] = true
	}
	
	return func(r Registro) bool {
		return ciudadSet[r.Ciudad]
	}
}

// FiltrarPorFecha crea un filtro por rango de fechas
func FiltrarPorFecha(inicio, fin time.Time) func(Registro) bool {
	return func(r Registro) bool {
		return r.Fecha.After(inicio) && r.Fecha.Before(fin)
	}
}

// aplicarFiltros aplica todos los filtros configurados
func (a *AnalizadorDatos) aplicarFiltros(datos []Registro) []Registro {
	resultado := make([]Registro, 0, len(datos))
	
	for _, registro := range datos {
		cumpleTodos := true
		for _, filtro := range a.filtros {
			if !filtro(registro) {
				cumpleTodos = false
				break
			}
		}
		if cumpleTodos {
			resultado = append(resultado, registro)
		}
	}
	
	return resultado
}

// ===== SISTEMA DE TRANSFORMACIONES =====

// AgregarTransformador añade un transformador al pipeline
func (a *AnalizadorDatos) AgregarTransformador(transformador func(Registro) Registro) {
	a.transformadores = append(a.transformadores, transformador)
}

// TransformadorBonificacion añade bonificación basada en edad
func TransformadorBonificacion(factor float64) func(Registro) Registro {
	return func(r Registro) Registro {
		if r.Edad > 30 {
			r.Salario = r.Salario * (1 + factor)
		}
		return r
	}
}

// TransformadorNormalizarCiudad normaliza nombres de ciudades
func TransformadorNormalizarCiudad() func(Registro) Registro {
	return func(r Registro) Registro {
		r.Ciudad = strings.Title(strings.ToLower(r.Ciudad))
		return r
	}
}

// aplicarTransformaciones aplica todas las transformaciones configuradas
func (a *AnalizadorDatos) aplicarTransformaciones(datos []Registro) []Registro {
	resultado := make([]Registro, len(datos))
	copy(resultado, datos)
	
	for i, registro := range resultado {
		for _, transformador := range a.transformadores {
			registro = transformador(registro)
		}
		resultado[i] = registro
	}
	
	return resultado
}

// ===== ANÁLISIS ESTADÍSTICO =====

// ObtenerDatosProcesados devuelve los datos después de aplicar filtros y transformaciones
func (a *AnalizadorDatos) ObtenerDatosProcesados() []Registro {
	datos := a.aplicarFiltros(a.datos)
	datos = a.aplicarTransformaciones(datos)
	return datos
}

// CalcularEstadisticasSalario calcula estadísticas descriptivas de salarios
func (a *AnalizadorDatos) CalcularEstadisticasSalario() EstadisticasDescriptivas {
	datos := a.ObtenerDatosProcesados()
	salarios := make([]float64, len(datos))
	
	for i, registro := range datos {
		salarios[i] = registro.Salario
	}
	
	return calcularEstadisticasDescriptivas(salarios)
}

// CalcularEstadisticasEdad calcula estadísticas descriptivas de edades
func (a *AnalizadorDatos) CalcularEstadisticasEdad() EstadisticasDescriptivas {
	datos := a.ObtenerDatosProcesados()
	edades := make([]float64, len(datos))
	
	for i, registro := range datos {
		edades[i] = float64(registro.Edad)
	}
	
	return calcularEstadisticasDescriptivas(edades)
}

// calcularEstadisticasDescriptivas es la función base para el análisis estadístico
func calcularEstadisticasDescriptivas(valores []float64) EstadisticasDescriptivas {
	if len(valores) == 0 {
		return EstadisticasDescriptivas{}
	}

	// Copiar y ordenar para no modificar el original
	sorted := make([]float64, len(valores))
	copy(sorted, valores)
	sort.Float64s(sorted)

	// Cálculos básicos
	suma := 0.0
	for _, v := range valores {
		suma += v
	}
	media := suma / float64(len(valores))

	// Mediana
	var mediana float64
	if len(sorted)%2 == 0 {
		mediana = (sorted[len(sorted)/2-1] + sorted[len(sorted)/2]) / 2
	} else {
		mediana = sorted[len(sorted)/2]
	}

	// Moda (valor más frecuente)
	frecuencias := make(map[float64]int)
	maxFrecuencia := 0
	var moda float64
	for _, v := range valores {
		frecuencias[v]++
		if frecuencias[v] > maxFrecuencia {
			maxFrecuencia = frecuencias[v]
			moda = v
		}
	}

	// Varianza y desviación estándar
	sumaCuadrados := 0.0
	for _, v := range valores {
		diff := v - media
		sumaCuadrados += diff * diff
	}
	varianza := sumaCuadrados / float64(len(valores))
	desviacion := math.Sqrt(varianza)

	// Percentiles
	p25 := calcularPercentil(sorted, 25)
	p75 := calcularPercentil(sorted, 75)

	return EstadisticasDescriptivas{
		Media:       media,
		Mediana:     mediana,
		Moda:        moda,
		Minimo:      sorted[0],
		Maximo:      sorted[len(sorted)-1],
		Desviacion:  desviacion,
		Varianza:    varianza,
		Rango:       sorted[len(sorted)-1] - sorted[0],
		Percentil25: p25,
		Percentil75: p75,
		Conteo:      len(valores),
	}
}

// calcularPercentil calcula el percentil especificado
func calcularPercentil(sortedValues []float64, percentil float64) float64 {
	if len(sortedValues) == 0 {
		return 0
	}
	
	index := (percentil / 100) * float64(len(sortedValues)-1)
	lower := int(math.Floor(index))
	upper := int(math.Ceil(index))

	if lower == upper {
		return sortedValues[lower]
	}

	weight := index - float64(lower)
	return sortedValues[lower]*(1-weight) + sortedValues[upper]*weight
}

// ===== ANÁLISIS AGRUPADO =====

// AgruparPorCiudad agrupa registros por ciudad
func (a *AnalizadorDatos) AgruparPorCiudad() map[string][]Registro {
	datos := a.ObtenerDatosProcesados()
	grupos := make(map[string][]Registro)
	
	for _, registro := range datos {
		grupos[registro.Ciudad] = append(grupos[registro.Ciudad], registro)
	}
	
	return grupos
}

// AgruparPorRangoEdad agrupa registros por rangos de edad
func (a *AnalizadorDatos) AgruparPorRangoEdad(tamañoRango int) map[string][]Registro {
	datos := a.ObtenerDatosProcesados()
	grupos := make(map[string][]Registro)
	
	for _, registro := range datos {
		rangoInicio := (registro.Edad / tamañoRango) * tamañoRango
		rangoFin := rangoInicio + tamañoRango - 1
		clave := fmt.Sprintf("%d-%d", rangoInicio, rangoFin)
		grupos[clave] = append(grupos[clave], registro)
	}
	
	return grupos
}

// ===== REPORTING Y VISUALIZACIÓN =====

// GenerarReporteCompleto genera un reporte detallado del análisis
func (a *AnalizadorDatos) GenerarReporteCompleto() string {
	var reporte strings.Builder
	
	reporte.WriteString("📊 REPORTE COMPLETO DE ANÁLISIS DE DATOS\n")
	reporte.WriteString("==========================================\n\n")
	
	// Información general
	totalRegistros := len(a.datos)
	registrosProcesados := len(a.ObtenerDatosProcesados())
	
	reporte.WriteString("📈 RESUMEN GENERAL\n")
	reporte.WriteString(fmt.Sprintf("Total de registros: %d\n", totalRegistros))
	reporte.WriteString(fmt.Sprintf("Registros procesados: %d\n", registrosProcesados))
	reporte.WriteString(fmt.Sprintf("Filtros aplicados: %d\n", len(a.filtros)))
	reporte.WriteString(fmt.Sprintf("Transformaciones aplicadas: %d\n\n", len(a.transformadores)))
	
	// Estadísticas de salarios
	statsSalario := a.CalcularEstadisticasSalario()
	reporte.WriteString("💰 ANÁLISIS DE SALARIOS\n")
	reporte.WriteString(fmt.Sprintf("Media: €%.2f\n", statsSalario.Media))
	reporte.WriteString(fmt.Sprintf("Mediana: €%.2f\n", statsSalario.Mediana))
	reporte.WriteString(fmt.Sprintf("Mínimo: €%.2f\n", statsSalario.Minimo))
	reporte.WriteString(fmt.Sprintf("Máximo: €%.2f\n", statsSalario.Maximo))
	reporte.WriteString(fmt.Sprintf("Desviación estándar: €%.2f\n", statsSalario.Desviacion))
	reporte.WriteString(fmt.Sprintf("Rango: €%.2f\n\n", statsSalario.Rango))
	
	// Estadísticas de edades
	statsEdad := a.CalcularEstadisticasEdad()
	reporte.WriteString("👥 ANÁLISIS DE EDADES\n")
	reporte.WriteString(fmt.Sprintf("Media: %.1f años\n", statsEdad.Media))
	reporte.WriteString(fmt.Sprintf("Mediana: %.1f años\n", statsEdad.Mediana))
	reporte.WriteString(fmt.Sprintf("Mínimo: %.0f años\n", statsEdad.Minimo))
	reporte.WriteString(fmt.Sprintf("Máximo: %.0f años\n", statsEdad.Maximo))
	reporte.WriteString(fmt.Sprintf("Desviación estándar: %.1f años\n\n", statsEdad.Desviacion))
	
	// Análisis por ciudad
	gruposCiudad := a.AgruparPorCiudad()
	reporte.WriteString("🏙️ DISTRIBUCIÓN POR CIUDAD\n")
	for ciudad, registros := range gruposCiudad {
		salarioPromedio := 0.0
		for _, r := range registros {
			salarioPromedio += r.Salario
		}
		salarioPromedio /= float64(len(registros))
		
		reporte.WriteString(fmt.Sprintf("%s: %d empleados (salario promedio: €%.2f)\n", 
			ciudad, len(registros), salarioPromedio))
	}
	
	return reporte.String()
}

// VisualizarHistograma crea una visualización simple en texto de un histograma
func (a *AnalizadorDatos) VisualizarHistograma(campo string, bins int) string {
	datos := a.ObtenerDatosProcesados()
	
	var valores []float64
	switch campo {
	case "salario":
		for _, r := range datos {
			valores = append(valores, r.Salario)
		}
	case "edad":
		for _, r := range datos {
			valores = append(valores, float64(r.Edad))
		}
	default:
		return "Campo no soportado para histograma"
	}
	
	return generarHistogramaTexto(valores, bins, campo)
}

// generarHistogramaTexto crea una representación visual en texto
func generarHistogramaTexto(valores []float64, bins int, titulo string) string {
	if len(valores) == 0 {
		return "No hay datos para mostrar"
	}
	
	// Encontrar min y max
	min, max := valores[0], valores[0]
	for _, v := range valores {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	
	// Crear bins
	binSize := (max - min) / float64(bins)
	counts := make([]int, bins)
	
	for _, v := range valores {
		binIndex := int((v - min) / binSize)
		if binIndex >= bins {
			binIndex = bins - 1
		}
		counts[binIndex]++
	}
	
	// Encontrar el máximo count para escalar
	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}
	
	// Generar histograma
	var resultado strings.Builder
	resultado.WriteString(fmt.Sprintf("\n📊 HISTOGRAMA: %s\n", strings.ToUpper(titulo)))
	resultado.WriteString(strings.Repeat("=", 50) + "\n")
	
	for i, count := range counts {
		rangeStart := min + float64(i)*binSize
		rangeEnd := rangeStart + binSize
		
		// Calcular barras (máximo 40 caracteres)
		barLength := int(float64(count) / float64(maxCount) * 40)
		bar := strings.Repeat("█", barLength)
		
		resultado.WriteString(fmt.Sprintf("[%8.1f-%8.1f]: %s (%d)\n", 
			rangeStart, rangeEnd, bar, count))
	}
	
	return resultado.String()
}

// ===== PIPELINE DE DATOS =====

// Pipeline representa una secuencia de operaciones de procesamiento
type Pipeline struct {
	operaciones []func([]Registro) []Registro
}

// NuevoPipeline crea un nuevo pipeline de procesamiento
func NuevoPipeline() *Pipeline {
	return &Pipeline{
		operaciones: make([]func([]Registro) []Registro, 0),
	}
}

// Filtrar añade una operación de filtrado al pipeline
func (p *Pipeline) Filtrar(filtro func(Registro) bool) *Pipeline {
	operacion := func(datos []Registro) []Registro {
		resultado := make([]Registro, 0)
		for _, registro := range datos {
			if filtro(registro) {
				resultado = append(resultado, registro)
			}
		}
		return resultado
	}
	p.operaciones = append(p.operaciones, operacion)
	return p
}

// Transformar añade una operación de transformación al pipeline
func (p *Pipeline) Transformar(transformador func(Registro) Registro) *Pipeline {
	operacion := func(datos []Registro) []Registro {
		resultado := make([]Registro, len(datos))
		for i, registro := range datos {
			resultado[i] = transformador(registro)
		}
		return resultado
	}
	p.operaciones = append(p.operaciones, operacion)
	return p
}

// Ordenar añade una operación de ordenamiento al pipeline
func (p *Pipeline) Ordenar(comparador func(Registro, Registro) bool) *Pipeline {
	operacion := func(datos []Registro) []Registro {
		resultado := make([]Registro, len(datos))
		copy(resultado, datos)
		
		sort.Slice(resultado, func(i, j int) bool {
			return comparador(resultado[i], resultado[j])
		})
		
		return resultado
	}
	p.operaciones = append(p.operaciones, operacion)
	return p
}

// Ejecutar ejecuta todo el pipeline sobre los datos
func (p *Pipeline) Ejecutar(datos []Registro) []Registro {
	resultado := datos
	for _, operacion := range p.operaciones {
		resultado = operacion(resultado)
	}
	return resultado
}

// EjecutarPipeline aplica un pipeline a los datos del analizador
func (a *AnalizadorDatos) EjecutarPipeline(pipeline *Pipeline) []Registro {
	return pipeline.Ejecutar(a.datos)
}

// ===== FUNCIONES DE DEMOSTRACIÓN =====

func main() {
	fmt.Println("🚀 SISTEMA DE ANÁLISIS DE DATOS - PROYECTO INTEGRADOR")
	fmt.Println("=====================================================")
	
	// Crear analizador
	analizador := NuevoAnalizador()
	
	// Configurar logger personalizado
	analizador.ConfigurarLogger(func(msg string) {
		fmt.Printf("🔍 [%s] %s\n", time.Now().Format("15:04:05"), msg)
	})
	
	// Generar datos de prueba
	analizador.GenerarDatosPrueba()
	
	fmt.Println("\n" + strings.Repeat("=", 60))
	demoFiltrosBasicos(analizador)
	
	fmt.Println("\n" + strings.Repeat("=", 60))
	demoTransformaciones(analizador)
	
	fmt.Println("\n" + strings.Repeat("=", 60))
	demoAnalisisEstadistico(analizador)
	
	fmt.Println("\n" + strings.Repeat("=", 60))
	demoVisualizacion(analizador)
	
	fmt.Println("\n" + strings.Repeat("=", 60))
	demoPipeline(analizador)
	
	fmt.Println("\n" + strings.Repeat("=", 60))
	demoReporteCompleto(analizador)
}

func demoFiltrosBasicos(analizador *AnalizadorDatos) {
	fmt.Println("📊 DEMO: Sistema de Filtros")
	fmt.Println("===========================")
	
	// Limpiar filtros anteriores
	analizador.filtros = nil
	
	fmt.Printf("Datos originales: %d registros\n", len(analizador.datos))
	
	// Filtro por edad
	analizador.AgregarFiltro(FiltrarPorEdad(25, 40))
	datos := analizador.ObtenerDatosProcesados()
	fmt.Printf("Después de filtrar por edad (25-40): %d registros\n", len(datos))
	
	// Filtro adicional por salario
	analizador.AgregarFiltro(FiltrarPorSalario(35000, 50000))
	datos = analizador.ObtenerDatosProcesados()
	fmt.Printf("Después de filtrar por salario (35k-50k): %d registros\n", len(datos))
	
	// Mostrar algunos ejemplos
	fmt.Println("\nEjemplos de registros filtrados:")
	for i, registro := range datos[:min(5, len(datos))] {
		fmt.Printf("%d. %s - Edad: %d, Salario: €%.0f\n", 
			i+1, registro.Nombre, registro.Edad, registro.Salario)
	}
}

func demoTransformaciones(analizador *AnalizadorDatos) {
	fmt.Println("🔄 DEMO: Sistema de Transformaciones")
	fmt.Println("====================================")
	
	// Limpiar transformadores anteriores
	analizador.transformadores = nil
	
	// Añadir bonificación para mayores de 30
	analizador.AgregarTransformador(TransformadorBonificacion(0.1)) // 10% de bonificación
	
	// Normalizar nombres de ciudades
	analizador.AgregarTransformador(TransformadorNormalizarCiudad())
	
	datos := analizador.ObtenerDatosProcesados()
	
	fmt.Println("Ejemplos con transformaciones aplicadas:")
	for i, registro := range datos[:min(3, len(datos))] {
		fmt.Printf("%d. %s - Edad: %d, Salario: €%.0f, Ciudad: %s\n", 
			i+1, registro.Nombre, registro.Edad, registro.Salario, registro.Ciudad)
	}
}

func demoAnalisisEstadistico(analizador *AnalizadorDatos) {
	fmt.Println("📈 DEMO: Análisis Estadístico")
	fmt.Println("=============================")
	
	// Estadísticas de salarios
	statsSalario := analizador.CalcularEstadisticasSalario()
	fmt.Println("ESTADÍSTICAS DE SALARIOS:")
	fmt.Printf("  Media: €%.2f\n", statsSalario.Media)
	fmt.Printf("  Mediana: €%.2f\n", statsSalario.Mediana)
	fmt.Printf("  Desviación estándar: €%.2f\n", statsSalario.Desviacion)
	fmt.Printf("  Rango: €%.2f - €%.2f\n", statsSalario.Minimo, statsSalario.Maximo)
	
	// Estadísticas de edades
	statsEdad := analizador.CalcularEstadisticasEdad()
	fmt.Println("\nESTADÍSTICAS DE EDADES:")
	fmt.Printf("  Media: %.1f años\n", statsEdad.Media)
	fmt.Printf("  Mediana: %.1f años\n", statsEdad.Mediana)
	fmt.Printf("  Rango: %.0f - %.0f años\n", statsEdad.Minimo, statsEdad.Maximo)
	
	// Análisis por grupos
	fmt.Println("\nDISTRIBUCIÓN POR CIUDAD:")
	grupos := analizador.AgruparPorCiudad()
	for ciudad, registros := range grupos {
		fmt.Printf("  %s: %d empleados\n", ciudad, len(registros))
	}
}

func demoVisualizacion(analizador *AnalizadorDatos) {
	fmt.Println("📊 DEMO: Visualización")
	fmt.Println("======================")
	
	// Histograma de salarios
	histogramaSalarios := analizador.VisualizarHistograma("salario", 8)
	fmt.Println(histogramaSalarios)
	
	// Histograma de edades
	histogramaEdades := analizador.VisualizarHistograma("edad", 6)
	fmt.Println(histogramaEdades)
}

func demoPipeline(analizador *AnalizadorDatos) {
	fmt.Println("🔄 DEMO: Pipeline de Datos")
	fmt.Println("==========================")
	
	// Crear pipeline complejo
	pipeline := NuevoPipeline().
		Filtrar(func(r Registro) bool { return r.Edad >= 30 }).
		Filtrar(func(r Registro) bool { return r.Salario > 40000 }).
		Transformar(func(r Registro) Registro {
			r.Salario = r.Salario * 1.05 // Aumento del 5%
			return r
		}).
		Ordenar(func(r1, r2 Registro) bool { return r1.Salario > r2.Salario })
	
	resultados := analizador.EjecutarPipeline(pipeline)
	
	fmt.Printf("Pipeline ejecutado: %d registros resultantes\n", len(resultados))
	fmt.Println("Top 5 salarios después del pipeline:")
	for i, registro := range resultados[:min(5, len(resultados))] {
		fmt.Printf("%d. %s - €%.0f\n", i+1, registro.Nombre, registro.Salario)
	}
}

func demoReporteCompleto(analizador *AnalizadorDatos) {
	fmt.Println("📋 DEMO: Reporte Completo")
	fmt.Println("=========================")
	
	reporte := analizador.GenerarReporteCompleto()
	fmt.Println(reporte)
}

// Función auxiliar
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
