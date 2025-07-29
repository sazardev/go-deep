package main

import (
	"sync"
	"testing"
)

// 🧪 Ejercicio 3: Object Pooling para optimización de memoria

// 🔄 Objeto pesado que es caro de crear
type ExpensiveObject struct {
	Buffer []byte
	Data   map[string]interface{}
}

// 🛠️ Constructor para objetos nuevos
func NewExpensiveObject() *ExpensiveObject {
	return &ExpensiveObject{
		Buffer: make([]byte, 1024*1024), // 1MB buffer
		Data:   make(map[string]interface{}),
	}
}

// 🧹 Método para limpiar/resetear el objeto
func (e *ExpensiveObject) Reset() {
	// Limpiar buffer
	for i := range e.Buffer {
		e.Buffer[i] = 0
	}
	// Limpiar mapa
	for k := range e.Data {
		delete(e.Data, k)
	}
}

// 🚀 Pool de objetos optimizado
var expensiveObjectPool = sync.Pool{
	New: func() interface{} {
		return NewExpensiveObject()
	},
}

// 🔄 Funciones para obtener y devolver objetos del pool
func GetExpensiveObject() *ExpensiveObject {
	return expensiveObjectPool.Get().(*ExpensiveObject)
}

func PutExpensiveObject(obj *ExpensiveObject) {
	obj.Reset()
	expensiveObjectPool.Put(obj)
}

// 🐌 Función que crea objeto cada vez (sin pool)
func ProcessDataWithoutPool(iterations int) {
	for i := 0; i < iterations; i++ {
		obj := NewExpensiveObject()

		// Simular trabajo con el objeto
		obj.Data["iteration"] = i
		obj.Buffer[0] = byte(i % 256)

		// Objeto se libera automáticamente (GC)
	}
}

// 🚀 Función que usa object pool
func ProcessDataWithPool(iterations int) {
	for i := 0; i < iterations; i++ {
		obj := GetExpensiveObject()

		// Simular trabajo con el objeto
		obj.Data["iteration"] = i
		obj.Buffer[0] = byte(i % 256)

		// Devolver al pool
		PutExpensiveObject(obj)
	}
}

// 📊 Benchmarks para comparar performance
func BenchmarkWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessDataWithoutPool(100)
	}
}

func BenchmarkWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessDataWithPool(100)
	}
}

// 🎯 Ejemplo de pool genérico (Go 1.18+)
type Pool[T any] struct {
	pool sync.Pool
}

func NewPool[T any](newFunc func() T) *Pool[T] {
	return &Pool[T]{
		pool: sync.Pool{
			New: func() interface{} {
				return newFunc()
			},
		},
	}
}

func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}

func (p *Pool[T]) Put(item T) {
	p.pool.Put(item)
}

// 🧪 Ejemplo de uso del pool genérico
func DemoGenericPool() {
	// Pool de slices de bytes
	byteSlicePool := NewPool(func() []byte {
		return make([]byte, 0, 1024)
	})

	// Usar el pool
	slice := byteSlicePool.Get()
	slice = append(slice, []byte("Hello, World!")...)

	// Limpiar y devolver
	slice = slice[:0] // Reset length pero mantiene capacity
	byteSlicePool.Put(slice)
}
