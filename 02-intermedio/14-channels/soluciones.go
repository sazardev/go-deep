// ==============================================
// LECCIÓN 14: Channels - Soluciones
// ==============================================

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// ==============================================
// EJERCICIO 1: Primer Channel - SOLUCIÓN
// ==============================================

func emisor(ch chan<- string) {
	mensajes := []string{"Hola", "Mundo", "Go"}
	for i, mensaje := range mensajes {
		fmt.Printf("📤 Enviando mensaje %d: %s\n", i+1, mensaje)
		ch <- mensaje
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func solucion1() {
	fmt.Println("📡 Ejercicio 1: Primer Channel - SOLUCIÓN")
	fmt.Println("========================================")

	// Crear channel unbuffered
	ch := make(chan string)

	// Lanzar emisor en goroutine
	go emisor(ch)

	// Recibir mensajes en main
	fmt.Println("📥 Esperando mensajes...")
	for mensaje := range ch {
		fmt.Printf("📥 Recibido: %s\n", mensaje)
		time.Sleep(300 * time.Millisecond) // Simular procesamiento
	}

	fmt.Println("✅ Channel cerrado, comunicación completada\n")
}

// ==============================================
// EJERCICIO 2: Buffered vs Unbuffered - SOLUCIÓN
// ==============================================

func probarUnbuffered() {
	fmt.Println("🔄 Probando channel UNBUFFERED:")
	ch := make(chan string)

	// Intentar enviar en main (esto bloquearía sin goroutine)
	go func() {
		fmt.Println("📤 Enviando a unbuffered...")
		ch <- "Mensaje 1"
		fmt.Println("✅ Mensaje 1 enviado")
		ch <- "Mensaje 2"
		fmt.Println("✅ Mensaje 2 enviado")
		close(ch)
	}()

	time.Sleep(100 * time.Millisecond) // Pequeño delay
	for msg := range ch {
		fmt.Printf("📥 Recibido: %s\n", msg)
		time.Sleep(200 * time.Millisecond) // Simular procesamiento lento
	}
}

func probarBuffered() {
	fmt.Println("\n📦 Probando channel BUFFERED:")
	ch := make(chan string, 3) // Buffer de 3

	// Enviar múltiples mensajes sin bloquear
	fmt.Println("📤 Enviando mensajes a buffer...")
	ch <- "Mensaje A"
	fmt.Println("✅ Mensaje A enviado (no bloquea)")
	ch <- "Mensaje B"
	fmt.Println("✅ Mensaje B enviado (no bloquea)")
	ch <- "Mensaje C"
	fmt.Println("✅ Mensaje C enviado (no bloquea)")
	close(ch)

	// Leer después
	fmt.Println("📥 Leyendo del buffer:")
	for msg := range ch {
		fmt.Printf("📥 Recibido: %s\n", msg)
	}
}

func solucion2() {
	fmt.Println("📦 Ejercicio 2: Buffered vs Unbuffered - SOLUCIÓN")
	fmt.Println("================================================")

	probarUnbuffered()
	probarBuffered()
	fmt.Println()
}

// ==============================================
// EJERCICIO 3: Productor-Consumidor - SOLUCIÓN
// ==============================================

func productor(ch chan<- int) {
	fmt.Println("🏭 Productor iniciando...")
	for i := 1; i <= 10; i++ {
		fmt.Printf("📤 Produciendo: %d\n", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
	fmt.Println("🏭 Productor terminado")
}

func consumidor(ch <-chan int) {
	fmt.Println("🍽️ Consumidor iniciando...")
	for numero := range ch {
		fmt.Printf("📥 Consumiendo: %d\n", numero)
		time.Sleep(150 * time.Millisecond) // Consumidor más lento
	}
	fmt.Println("🍽️ Consumidor terminado")
}

func solucion3() {
	fmt.Println("🏭 Ejercicio 3: Productor-Consumidor - SOLUCIÓN")
	fmt.Println("==============================================")

	ch := make(chan int)

	// Lanzar productor y consumidor
	go productor(ch)
	go consumidor(ch)

	// Esperar a que terminen
	time.Sleep(3 * time.Second)
	fmt.Println()
}

// ==============================================
// EJERCICIO 4: Select Statement - SOLUCIÓN
// ==============================================

func emisorConDelay(ch chan<- string, mensaje string, delay time.Duration) {
	time.Sleep(delay)
	ch <- mensaje
	close(ch)
}

func solucion4() {
	fmt.Println("🎛️ Ejercicio 4: Select Statement - SOLUCIÓN")
	fmt.Println("=========================================")

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	// Lanzar emisores con diferentes delays
	go emisorConDelay(ch1, "Respuesta rápida", 100*time.Millisecond)
	go emisorConDelay(ch2, "Respuesta media", 200*time.Millisecond)
	go emisorConDelay(ch3, "Respuesta lenta", 300*time.Millisecond)

	// Recibir usando select
	mensajesRecibidos := 0
	for mensajesRecibidos < 3 {
		select {
		case msg, ok := <-ch1:
			if ok {
				fmt.Printf("🔵 Canal 1: %s\n", msg)
				mensajesRecibidos++
			}
		case msg, ok := <-ch2:
			if ok {
				fmt.Printf("🟢 Canal 2: %s\n", msg)
				mensajesRecibidos++
			}
		case msg, ok := <-ch3:
			if ok {
				fmt.Printf("🟡 Canal 3: %s\n", msg)
				mensajesRecibidos++
			}
		}
	}

	fmt.Println("✅ Todos los mensajes recibidos\n")
}

// ==============================================
// EJERCICIO 5: Select con Timeout - SOLUCIÓN
// ==============================================

func operacionLenta(ch chan<- string, delay time.Duration) {
	fmt.Printf("⏳ Iniciando operación (duración: %v)\n", delay)
	time.Sleep(delay)
	ch <- "Operación completada"
}

func probarConTimeout(delay time.Duration) {
	ch := make(chan string)
	go operacionLenta(ch, delay)

	select {
	case resultado := <-ch:
		fmt.Printf("✅ %s\n", resultado)
	case <-time.After(2 * time.Second):
		fmt.Println("⏰ Timeout: Operación cancelada después de 2 segundos")
	}
}

func solucion5() {
	fmt.Println("⏰ Ejercicio 5: Select con Timeout - SOLUCIÓN")
	fmt.Println("===========================================")

	fmt.Println("🧪 Prueba 1 - Operación rápida (1 segundo):")
	probarConTimeout(1 * time.Second)

	fmt.Println("\n🧪 Prueba 2 - Operación lenta (3 segundos):")
	probarConTimeout(3 * time.Second)

	fmt.Println()
}

// ==============================================
// EJERCICIO 6: Pipeline de Datos - SOLUCIÓN
// ==============================================

// Etapa 1: Generar números
func generarNumeros() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		fmt.Println("🔢 Etapa 1: Generando números del 1 al 20")
		for i := 1; i <= 20; i++ {
			ch <- i
		}
	}()
	return ch
}

// Etapa 2: Multiplicar por 2
func multiplicarPorDos(input <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		fmt.Println("✖️ Etapa 2: Multiplicando por 2")
		for num := range input {
			ch <- num * 2
		}
	}()
	return ch
}

// Etapa 3: Filtrar divisibles por 4
func filtrarDivisiblesPor4(input <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		fmt.Println("🔍 Etapa 3: Filtrando divisibles por 4")
		for num := range input {
			if num%4 == 0 {
				ch <- num
			}
		}
	}()
	return ch
}

// Etapa 4: Formatear como string
func formatearString(input <-chan int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		fmt.Println("📝 Etapa 4: Formateando como string")
		for num := range input {
			ch <- fmt.Sprintf("Número procesado: %d", num)
		}
	}()
	return ch
}

func solucion6() {
	fmt.Println("🔄 Ejercicio 6: Pipeline de Datos - SOLUCIÓN")
	fmt.Println("==========================================")

	// Construir pipeline
	numeros := generarNumeros()
	multiplicados := multiplicarPorDos(numeros)
	filtrados := filtrarDivisiblesPor4(multiplicados)
	formateados := formatearString(filtrados)

	// Procesar resultados
	fmt.Println("\n📊 Resultados del pipeline:")
	for resultado := range formateados {
		fmt.Printf("  %s\n", resultado)
	}

	fmt.Println("✅ Pipeline completado\n")
}

// ==============================================
// EJERCICIO 7: Worker Pool - SOLUCIÓN
// ==============================================

type TrabajoWorker struct {
	ID   int
	Dato int
}

type ResultadoWorker struct {
	TrabajoID int
	Resultado int
}

func worker(id int, trabajos <-chan TrabajoWorker, resultados chan<- ResultadoWorker) {
	for trabajo := range trabajos {
		fmt.Printf("👷 Worker %d procesando trabajo %d\n", id, trabajo.ID)
		
		// Simular procesamiento
		time.Sleep(100 * time.Millisecond)
		
		// Calcular cuadrado
		resultado := ResultadoWorker{
			TrabajoID: trabajo.ID,
			Resultado: trabajo.Dato * trabajo.Dato,
		}
		
		resultados <- resultado
	}
	fmt.Printf("👷 Worker %d terminando\n", id)
}

func solucion7() {
	fmt.Println("👷 Ejercicio 7: Worker Pool - SOLUCIÓN")
	fmt.Println("====================================")

	trabajos := make(chan TrabajoWorker, 5)
	resultados := make(chan ResultadoWorker, 15)

	// Lanzar 3 workers
	fmt.Println("🏭 Lanzando 3 workers...")
	for i := 1; i <= 3; i++ {
		go worker(i, trabajos, resultados)
	}

	// Enviar trabajos
	fmt.Println("📋 Enviando 15 trabajos...")
	for j := 1; j <= 15; j++ {
		trabajo := TrabajoWorker{ID: j, Dato: j}
		trabajos <- trabajo
	}
	close(trabajos)

	// Recoger resultados
	fmt.Println("\n📊 Recogiendo resultados:")
	for r := 1; r <= 15; r++ {
		resultado := <-resultados
		fmt.Printf("📈 Trabajo %d: %d² = %d\n", 
			resultado.TrabajoID, int(resultado.TrabajoID), resultado.Resultado)
	}

	fmt.Println("✅ Worker pool completado\n")
}

// ==============================================
// EJERCICIO 8: Fan-Out/Fan-In - SOLUCIÓN
// ==============================================

func generarNumerosParaPrimos() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 50; i++ {
			ch <- i
		}
	}()
	return ch
}

func workerPrimos(id int, numeros <-chan int, primos chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for num := range numeros {
		if esPrimoSolucion(num) {
			fmt.Printf("🔍 Worker %d encontró primo: %d\n", id, num)
			primos <- num
		}
	}
}

func merge(canales ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	merged := make(chan int)

	copiar := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			merged <- n
		}
	}

	wg.Add(len(canales))
	for _, c := range canales {
		go copiar(c)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func solucion8() {
	fmt.Println("📊 Ejercicio 8: Fan-Out/Fan-In - SOLUCIÓN")
	fmt.Println("=======================================")

	// Generar números
	numeros := generarNumerosParaPrimos()

	// Fan-Out: Distribuir entre 5 workers
	var wg sync.WaitGroup
	canalesPrimos := make([]chan int, 5)

	fmt.Println("🏭 Distribuyendo trabajo entre 5 workers...")
	for i := 0; i < 5; i++ {
		primos := make(chan int)
		canalesPrimos[i] = primos
		wg.Add(1)
		go workerPrimos(i+1, numeros, primos, &wg)
	}

	// Cerrar channels cuando terminen los workers
	go func() {
		wg.Wait()
		for _, ch := range canalesPrimos {
			close(ch)
		}
	}()

	// Fan-In: Combinar resultados
	canalesReadOnly := make([]<-chan int, 5)
	for i, ch := range canalesPrimos {
		canalesReadOnly[i] = ch
	}
	merged := merge(canalesReadOnly...)

	// Recoger primos
	fmt.Println("\n🎯 Números primos encontrados:")
	var primosEncontrados []int
	for primo := range merged {
		primosEncontrados = append(primosEncontrados, primo)
	}

	// Ordenar y mostrar
	fmt.Printf("📋 Lista de primos del 1 al 50: %v\n", primosEncontrados)
	fmt.Println("✅ Fan-Out/Fan-In completado\n")
}

// ==============================================
// EJERCICIO 9: Quit Channel - SOLUCIÓN
// ==============================================

func servidor(quit <-chan bool) {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	contador := 0
	fmt.Println("🚀 Servidor iniciado")

	for {
		select {
		case <-ticker.C:
			contador++
			fmt.Printf("⚡ Servidor tick #%d\n", contador)

		case <-quit:
			fmt.Println("🛑 Señal de quit recibida")
			fmt.Println("🧹 Realizando cleanup...")
			time.Sleep(200 * time.Millisecond) // Simular cleanup
			fmt.Println("✅ Servidor terminado correctamente")
			return
		}
	}
}

func solucion9() {
	fmt.Println("🛑 Ejercicio 9: Quit Channel - SOLUCIÓN")
	fmt.Println("====================================")

	quit := make(chan bool)

	// Lanzar servidor
	go servidor(quit)

	// Dejar correr por 3 segundos
	fmt.Println("⏱️ Servidor corriendo por 3 segundos...")
	time.Sleep(3 * time.Second)

	// Enviar quit
	fmt.Println("📤 Enviando señal de quit...")
	quit <- true

	// Dar tiempo para cleanup
	time.Sleep(500 * time.Millisecond)
	fmt.Println()
}

// ==============================================
// EJERCICIO 10: Channel de Channels - SOLUCIÓN
// ==============================================

func generadorCanales() <-chan (<-chan int) {
	canales := make(chan (<-chan int))
	
	go func() {
		defer close(canales)
		
		// Crear 3 channels con diferentes secuencias
		for i := 1; i <= 3; i++ {
			ch := make(chan int)
			canales <- ch
			
			// Cada channel genera su secuencia
			go func(id int, canal chan int) {
				defer close(canal)
				base := id * 10
				for j := 1; j <= 5; j++ {
					valor := base + j
					fmt.Printf("📡 Canal %d enviando: %d\n", id, valor)
					canal <- valor
					time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				}
			}(i, ch)
		}
	}()
	
	return canales
}

func multiplexor(canales <-chan (<-chan int)) <-chan int {
	salida := make(chan int)
	var wg sync.WaitGroup
	
	// Función para leer de un channel específico
	leer := func(ch <-chan int) {
		defer wg.Done()
		for valor := range ch {
			salida <- valor
		}
	}
	
	// Lanzar goroutine para cada channel que llegue
	go func() {
		for ch := range canales {
			wg.Add(1)
			go leer(ch)
		}
		
		// Cerrar salida cuando todos terminen
		go func() {
			wg.Wait()
			close(salida)
		}()
	}()
	
	return salida
}

func solucion10() {
	fmt.Println("📡 Ejercicio 10: Channel de Channels - SOLUCIÓN")
	fmt.Println("=============================================")

	rand.Seed(time.Now().UnixNano())

	// Generar channels dinámicamente
	canales := generadorCanales()
	
	// Multiplexar todos los channels
	valores := multiplexor(canales)
	
	// Recibir valores intercalados
	fmt.Println("\n📊 Valores intercalados:")
	var todosLosValores []int
	for valor := range valores {
		fmt.Printf("📥 Recibido: %d\n", valor)
		todosLosValores = append(todosLosValores, valor)
	}
	
	fmt.Printf("\n📋 Total de valores recibidos: %v\n", len(todosLosValores))
	fmt.Println("✅ Multiplexación completada\n")
}

// ==============================================
// FUNCIÓN PRINCIPAL Y UTILIDADES
// ==============================================

func ejecutarSoluciones() {
	fmt.Println("📡 SOLUCIONES: Channels")
	fmt.Println("======================")
	fmt.Printf("Usando Go %s con %d CPUs\n\n", runtime.Version(), runtime.NumCPU())

	soluciones := []func(){
		solucion1,
		solucion2,
		solucion3,
		solucion4,
		solucion5,
		solucion6,
		solucion7,
		solucion8,
		solucion9,
		solucion10,
	}

	for i, solucion := range soluciones {
		fmt.Printf("🔧 Ejecutando solución %d...\n", i+1)
		solucion()
		time.Sleep(1 * time.Second) // Pausa entre soluciones
	}

	fmt.Println("🎓 Conceptos dominados:")
	fmt.Println("   ✅ Channels básicos (buffered/unbuffered)")
	fmt.Println("   ✅ Comunicación productor-consumidor")
	fmt.Println("   ✅ Select statement y multiplexación")
	fmt.Println("   ✅ Timeouts y cancelación")
	fmt.Println("   ✅ Pipelines de transformación")
	fmt.Println("   ✅ Worker pools con channels")
	fmt.Println("   ✅ Fan-Out/Fan-In patterns")
	fmt.Println("   ✅ Quit channels para shutdown")
	fmt.Println("   ✅ Channels de channels dinámicos")
	fmt.Println("   ✅ Multiplexación avanzada")

	fmt.Println("\n💡 Próximo paso: Lección 15 - Context Package")
	fmt.Println("   🎯 Cancelación de operaciones")
	fmt.Println("   ⏰ Timeouts y deadlines")
	fmt.Println("   📦 Propagación de valores")
	fmt.Println("   🛡️ Best practices con context")
}

// ==============================================
// FUNCIONES AUXILIARES
// ==============================================

func esPrimoSolucion(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
