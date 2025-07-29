package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

// 🧠 Ejemplos de Memory Leaks para detectar y resolver

// DataProcessor - Servicio con múltiples memory leaks intencionados
type DataProcessor struct {
	mu        sync.RWMutex
	cache     map[string][]byte // LEAK: nunca se limpia
	workers   []chan []byte     // LEAK: workers crecen infinitamente
	callbacks []func([]byte)    // LEAK: callbacks nunca se remueven
	stats     []ProcessingStat  // LEAK: estadísticas crecen sin límite
}

type ProcessingStat struct {
	Timestamp time.Time
	DataSize  int
	Duration  time.Duration
}

func NewDataProcessor() *DataProcessor {
	return &DataProcessor{
		cache:     make(map[string][]byte),
		workers:   make([]chan []byte, 0),
		callbacks: make([]func([]byte), 0),
		stats:     make([]ProcessingStat, 0),
	}
}

// ProcessData - Método con múltiples memory leaks
func (dp *DataProcessor) ProcessData(key string, data []byte) {
	start := time.Now()

	// LEAK 1: Cache crece infinitamente
	dp.mu.Lock()
	dp.cache[key] = make([]byte, len(data))
	copy(dp.cache[key], data)
	dp.mu.Unlock()

	// LEAK 2: Workers crecen cuando están ocupados
	var workerChan chan []byte
	foundWorker := false

	for _, worker := range dp.workers {
		select {
		case worker <- data:
			workerChan = worker
			foundWorker = true
		default:
			continue
		}
		if foundWorker {
			break
		}
	}

	if !foundWorker {
		// Crear nuevo worker si todos están ocupados
		newWorker := make(chan []byte, 100)
		dp.workers = append(dp.workers, newWorker)
		workerChan = newWorker

		go func(ch <-chan []byte) {
			for data := range ch {
				// Simular procesamiento
				time.Sleep(time.Duration(len(data)) * time.Microsecond)

				// Notificar callbacks
				dp.mu.RLock()
				for _, callback := range dp.callbacks {
					callback(data)
				}
				dp.mu.RUnlock()
			}
		}(newWorker)
	}

	if workerChan != nil {
		workerChan <- data
	}

	// LEAK 3: Estadísticas crecen sin límite
	stat := ProcessingStat{
		Timestamp: start,
		DataSize:  len(data),
		Duration:  time.Since(start),
	}

	dp.mu.Lock()
	dp.stats = append(dp.stats, stat)
	dp.mu.Unlock()
}

func (dp *DataProcessor) RegisterCallback(cb func([]byte)) {
	dp.mu.Lock()
	defer dp.mu.Unlock()

	// LEAK 4: Callbacks nunca se remueven
	dp.callbacks = append(dp.callbacks, cb)
}

func (dp *DataProcessor) GetStats() []ProcessingStat {
	dp.mu.RLock()
	defer dp.mu.RUnlock()

	// LEAK 5: Devuelve todo el slice, no hay límite
	result := make([]ProcessingStat, len(dp.stats))
	copy(result, dp.stats)
	return result
}

// 🚨 WebCrawler con goroutine leaks
type WebCrawler struct {
	client     *http.Client
	activeJobs sync.WaitGroup
	results    chan CrawlResult
	shutdown   chan struct{} // Para versión fixed
}

type CrawlResult struct {
	URL    string
	Status int
	Error  error
}

func NewWebCrawler() *WebCrawler {
	return &WebCrawler{
		client:   &http.Client{Timeout: 10 * time.Second},
		results:  make(chan CrawlResult, 1000),
		shutdown: make(chan struct{}),
	}
}

// CrawlURL - Versión con goroutine leak
func (wc *WebCrawler) CrawlURL(url string) {
	// LEAK: Goroutine sin context para cancelación
	go func() {
		defer wc.activeJobs.Done()

		// LEAK: No hay timeout, puede colgarse indefinidamente
		resp, err := wc.client.Get(url)
		if err != nil {
			wc.results <- CrawlResult{URL: url, Error: err}
			return
		}
		defer resp.Body.Close()

		wc.results <- CrawlResult{
			URL:    url,
			Status: resp.StatusCode,
		}
	}()

	wc.activeJobs.Add(1)
}

// StartWorkers - Versión con goroutine leak
func (wc *WebCrawler) StartWorkers(numWorkers int) {
	// LEAK: Workers sin mecanismo de shutdown
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for {
				select {
				case result := <-wc.results:
					fmt.Printf("Worker %d processed: %+v\n", workerID, result)
					// LEAK: No hay case para shutdown
				}
			}
		}(i)
	}
}

// 🛡️ Versiones corregidas (sin leaks)

// FixedDataProcessor - Versión sin memory leaks
type FixedDataProcessor struct {
	mu           sync.RWMutex
	cache        map[string]*CacheEntry  // Con TTL
	workers      []chan []byte           // Limitado
	callbacks    map[string]func([]byte) // Removables
	stats        []ProcessingStat        // Con rotación
	maxWorkers   int
	maxCacheSize int
	maxStats     int
}

type CacheEntry struct {
	Data    []byte
	Expires time.Time
}

func NewFixedDataProcessor() *FixedDataProcessor {
	fdp := &FixedDataProcessor{
		cache:        make(map[string]*CacheEntry),
		workers:      make([]chan []byte, 0),
		callbacks:    make(map[string]func([]byte)),
		stats:        make([]ProcessingStat, 0),
		maxWorkers:   10,   // Límite de workers
		maxCacheSize: 1000, // Límite de cache
		maxStats:     5000, // Límite de stats
	}

	// Iniciar limpieza automática
	go fdp.cleanupLoop()

	return fdp
}

func (fdp *FixedDataProcessor) cleanupLoop() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		fdp.cleanup()
	}
}

func (fdp *FixedDataProcessor) cleanup() {
	fdp.mu.Lock()
	defer fdp.mu.Unlock()

	now := time.Now()

	// Limpiar cache expirado
	for key, entry := range fdp.cache {
		if now.After(entry.Expires) {
			delete(fdp.cache, key)
		}
	}

	// Limpiar estadísticas viejas
	if len(fdp.stats) > fdp.maxStats {
		copy(fdp.stats, fdp.stats[len(fdp.stats)-fdp.maxStats:])
		fdp.stats = fdp.stats[:fdp.maxStats]
	}
}

func (fdp *FixedDataProcessor) ProcessData(key string, data []byte) {
	start := time.Now()

	// Cache con TTL y límite de tamaño
	fdp.mu.Lock()
	if len(fdp.cache) < fdp.maxCacheSize {
		fdp.cache[key] = &CacheEntry{
			Data:    append([]byte(nil), data...),
			Expires: time.Now().Add(5 * time.Minute),
		}
	}
	fdp.mu.Unlock()

	// Workers con límite
	var workerChan chan []byte
	if len(fdp.workers) < fdp.maxWorkers {
		// Buscar worker disponible
		for _, worker := range fdp.workers {
			select {
			case worker <- data:
				workerChan = worker
				goto process
			default:
				continue
			}
		}

		// Crear nuevo worker solo si no alcanzamos el límite
		if len(fdp.workers) < fdp.maxWorkers {
			newWorker := make(chan []byte, 100)
			fdp.workers = append(fdp.workers, newWorker)
			workerChan = newWorker

			go func(ch <-chan []byte) {
				for data := range ch {
					time.Sleep(time.Duration(len(data)) * time.Microsecond)
					// Procesar callbacks de forma segura
				}
			}(newWorker)
		}
	}

process:
	if workerChan != nil {
		workerChan <- data
	}

	// Stats con rotación
	stat := ProcessingStat{
		Timestamp: start,
		DataSize:  len(data),
		Duration:  time.Since(start),
	}

	fdp.mu.Lock()
	fdp.stats = append(fdp.stats, stat)
	fdp.mu.Unlock()
}

func (fdp *FixedDataProcessor) RegisterCallback(id string, cb func([]byte)) {
	fdp.mu.Lock()
	defer fdp.mu.Unlock()

	fdp.callbacks[id] = cb
}

func (fdp *FixedDataProcessor) UnregisterCallback(id string) {
	fdp.mu.Lock()
	defer fdp.mu.Unlock()

	delete(fdp.callbacks, id)
}

// 📊 Memory monitoring utilities

func PrintMemoryStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("📊 Memory Statistics:\n")
	fmt.Printf("  Allocated memory: %s\n", formatBytes(m.Alloc))
	fmt.Printf("  Total allocated: %s\n", formatBytes(m.TotalAlloc))
	fmt.Printf("  System memory: %s\n", formatBytes(m.Sys))
	fmt.Printf("  Number of GCs: %d\n", m.NumGC)
	fmt.Printf("  Goroutines: %d\n", runtime.NumGoroutine())
	fmt.Println()
}

func formatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

// Demo functions

func MemoryLeakDemo() {
	fmt.Println("🧠 MEMORY LEAK DEMO")
	fmt.Println("This will demonstrate growing memory usage")
	fmt.Println()

	processor := NewDataProcessor()

	fmt.Println("Initial memory state:")
	PrintMemoryStats()

	// Registrar muchos callbacks para simular leak
	for i := 0; i < 100; i++ {
		processor.RegisterCallback(func(data []byte) {
			// Callback que mantiene referencia a data
			_ = append([]byte(nil), data...)
			time.Sleep(1 * time.Millisecond)
		})
	}

	// Procesar muchos datos
	fmt.Println("Processing data (creating leaks)...")
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("data_%d", i)
		data := make([]byte, 1024) // 1KB por request
		processor.ProcessData(key, data)

		if i%100 == 0 {
			fmt.Printf("Processed %d items\n", i)
			PrintMemoryStats()
		}
	}

	fmt.Println("Final memory state (with leaks):")
	PrintMemoryStats()

	stats := processor.GetStats()
	fmt.Printf("Stats collected: %d entries\n", len(stats))
}

func FixedMemoryDemo() {
	fmt.Println("🛡️ FIXED MEMORY DEMO")
	fmt.Println("This demonstrates controlled memory usage")
	fmt.Println()

	processor := NewFixedDataProcessor()

	fmt.Println("Initial memory state:")
	PrintMemoryStats()

	// Registrar callbacks con IDs para poder removerlos
	for i := 0; i < 100; i++ {
		id := fmt.Sprintf("callback_%d", i)
		processor.RegisterCallback(id, func(data []byte) {
			// Callback más eficiente
			time.Sleep(1 * time.Millisecond)
		})
	}

	// Procesar datos con limpieza automática
	fmt.Println("Processing data (with leak prevention)...")
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("data_%d", i%100) // Reutilizar keys
		data := make([]byte, 1024)
		processor.ProcessData(key, data)

		if i%100 == 0 {
			fmt.Printf("Processed %d items\n", i)
			PrintMemoryStats()
		}

		// Remover callbacks viejos
		if i%50 == 0 && i > 0 {
			oldId := fmt.Sprintf("callback_%d", (i/50-1)%100)
			processor.UnregisterCallback(oldId)
		}
	}

	fmt.Println("Final memory state (fixed version):")
	PrintMemoryStats()
}

func GoroutineLeakDemo() {
	fmt.Println("🔄 GOROUTINE LEAK DEMO")
	fmt.Println()

	fmt.Printf("Initial goroutines: %d\n", runtime.NumGoroutine())

	crawler := NewWebCrawler()
	crawler.StartWorkers(10)

	// URLs que pueden causar timeout
	urls := []string{
		"http://httpbin.org/delay/1",
		"http://httpbin.org/delay/2",
		"http://httpbin.org/status/404",
		"http://nonexistent-domain-12345.com",
	}

	fmt.Println("Starting crawls (may create goroutine leaks)...")
	for i := 0; i < 20; i++ {
		url := urls[i%len(urls)]
		crawler.CrawlURL(url)

		if i%5 == 0 {
			fmt.Printf("Started %d crawls, goroutines: %d\n",
				i+1, runtime.NumGoroutine())
		}
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("Final goroutines: %d\n", runtime.NumGoroutine())
	fmt.Println("💡 Some goroutines may be leaked (workers without shutdown)")
}

// MemoryMain - función principal para demos de memory leaks
func MemoryMain() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "leak":
			MemoryLeakDemo()
		case "fixed":
			FixedMemoryDemo()
		case "goroutine":
			GoroutineLeakDemo()
		case "stats":
			PrintMemoryStats()
		default:
			fmt.Println("Usage: go run memory_leak_examples.go [leak|fixed|goroutine|stats]")
		}
	} else {
		fmt.Println("🧠 Memory Leak Detection Examples")
		fmt.Println()
		fmt.Println("Available demos:")
		fmt.Println("  leak      - Demonstrate memory leaks")
		fmt.Println("  fixed     - Demonstrate fixed version")
		fmt.Println("  goroutine - Demonstrate goroutine leaks")
		fmt.Println("  stats     - Show current memory stats")
		fmt.Println()
		fmt.Println("Memory profiling commands:")
		fmt.Println("  go build -o app memory_leak_examples.go")
		fmt.Println("  ./app leak &")
		fmt.Println("  go tool pprof http://localhost:6060/debug/pprof/heap")
		fmt.Println()
		fmt.Println("Monitoring commands:")
		fmt.Println("  watch -n 1 'ps aux | grep app'")
		fmt.Println("  htop")
	}
}
