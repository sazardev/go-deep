// 🎯 Archivo de coordinación para todos los ejemplos de profiling y debugging

// Para ejecutar los ejemplos, usa los siguientes comandos:

/*

📊 BENCHMARK EXAMPLES:
go run benchmark_examples.go profile
go run benchmark_examples.go memory
go test -bench=. -benchmem

🐛 DEADLOCK EXAMPLES:
go run deadlock_examples.go deadlock
go run deadlock_examples.go safe
go run -race deadlock_examples.go race
go run deadlock_examples.go goroutines

🧠 MEMORY LEAK EXAMPLES:
go run memory_leak_examples.go leak
go run memory_leak_examples.go fixed
go run memory_leak_examples.go goroutine
go run memory_leak_examples.go stats

� PROFILING COMMANDS:
go test -bench=. -cpuprofile=cpu.prof
go test -bench=. -memprofile=mem.prof
go tool pprof cpu.prof
go tool pprof mem.prof

� DEBUGGING COMMANDS:
dlv debug
(dlv) b main.main
(dlv) c
(dlv) goroutines
(dlv) goroutine 1 bt

*/

package main
