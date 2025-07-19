package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

// ==============================================
// PROYECTO: Sistema de Gestión de Biblioteca Digital
// ==============================================

// Estructuras de datos principales
type Libro struct {
	ID          string                 `json:"id"`
	ISBN        string                 `json:"isbn"`
	Titulo      string                 `json:"titulo"`
	Autores     []string               `json:"autores"`
	Generos     []string               `json:"generos"`
	Editorial   string                 `json:"editorial"`
	AnoPublic   int                    `json:"ano_publicacion"`
	Paginas     int                    `json:"paginas"`
	Idioma      string                 `json:"idioma"`
	Descripcion string                 `json:"descripcion"`
	Tags        []string               `json:"tags"`
	Disponible  bool                   `json:"disponible"`
	Ubicacion   string                 `json:"ubicacion"`
	Estado      string                 `json:"estado"` // "nuevo", "bueno", "regular", "malo"
	FechaAdq    time.Time              `json:"fecha_adquisicion"`
	Precio      float64                `json:"precio"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type Usuario struct {
	ID           string                 `json:"id"`
	Email        string                 `json:"email"`
	Nombre       string                 `json:"nombre"`
	Apellido     string                 `json:"apellido"`
	Telefono     string                 `json:"telefono"`
	Direccion    string                 `json:"direccion"`
	FechaNac     time.Time              `json:"fecha_nacimiento"`
	FechaReg     time.Time              `json:"fecha_registro"`
	TipoUsuario  string                 `json:"tipo_usuario"` // "estudiante", "profesor", "externo"
	Activo       bool                   `json:"activo"`
	LimiteLibros int                    `json:"limite_libros"`
	Multas       float64                `json:"multas"`
	Historial    []string               `json:"historial"` // IDs de préstamos
	Preferencias map[string]interface{} `json:"preferencias"`
}

type Prestamo struct {
	ID            string     `json:"id"`
	UsuarioID     string     `json:"usuario_id"`
	LibroID       string     `json:"libro_id"`
	FechaPrestamo time.Time  `json:"fecha_prestamo"`
	FechaVenc     time.Time  `json:"fecha_vencimiento"`
	FechaDevol    *time.Time `json:"fecha_devolucion,omitempty"`
	Estado        string     `json:"estado"` // "activo", "devuelto", "vencido", "renovado"
	Renovaciones  int        `json:"renovaciones"`
	Multa         float64    `json:"multa"`
	Notas         string     `json:"notas"`
}

type Reserva struct {
	ID        string    `json:"id"`
	UsuarioID string    `json:"usuario_id"`
	LibroID   string    `json:"libro_id"`
	FechaRes  time.Time `json:"fecha_reserva"`
	Prioridad int       `json:"prioridad"`
	Estado    string    `json:"estado"` // "pendiente", "disponible", "cancelada", "cumplida"
}

// Estructuras para reportes y estadísticas
type EstadisticasPeriodo struct {
	PrestamosRealizados int                    `json:"prestamos_realizados"`
	LibrosDevueltos     int                    `json:"libros_devueltos"`
	MultasGeneradas     float64                `json:"multas_generadas"`
	UsuariosActivos     int                    `json:"usuarios_activos"`
	LibrosMasPopulares  []LibroPopularidad     `json:"libros_mas_populares"`
	GenerosMasLeidos    map[string]int         `json:"generos_mas_leidos"`
	Detalles            map[string]interface{} `json:"detalles"`
}

type LibroPopularidad struct {
	LibroID    string  `json:"libro_id"`
	Titulo     string  `json:"titulo"`
	Prestamos  int     `json:"prestamos"`
	Puntuacion float64 `json:"puntuacion"`
}

type UsuarioActividad struct {
	UsuarioID string  `json:"usuario_id"`
	Nombre    string  `json:"nombre"`
	Prestamos int     `json:"prestamos"`
	Multas    float64 `json:"multas"`
}

// Sistema de caché
type CacheBiblioteca struct {
	busquedas    map[string][]Libro     `json:"busquedas"`
	estadisticas map[string]interface{} `json:"estadisticas"`
	ttl          time.Duration          `json:"ttl"`
	lastUpdate   map[string]time.Time   `json:"last_update"`
	mu           sync.RWMutex           `json:"-"`
}

// Configuración del sistema
type ConfigBiblioteca struct {
	MaxPrestamosEstudiante int     `json:"max_prestamos_estudiante"`
	MaxPrestamosProfesor   int     `json:"max_prestamos_profesor"`
	MaxPrestamosExterno    int     `json:"max_prestamos_externo"`
	DiasPrestamoEstudiante int     `json:"dias_prestamo_estudiante"`
	DiasPrestamoProfesor   int     `json:"dias_prestamo_profesor"`
	DiasPrestamoExterno    int     `json:"dias_prestamo_externo"`
	MultaDiaria            float64 `json:"multa_diaria"`
	MaxRenovaciones        int     `json:"max_renovaciones"`
	DiasReserva            int     `json:"dias_reserva"`
}

// Estructura principal del sistema
type Biblioteca struct {
	// Almacenamiento principal
	libros    map[string]Libro    `json:"libros"`
	usuarios  map[string]Usuario  `json:"usuarios"`
	prestamos map[string]Prestamo `json:"prestamos"`
	reservas  map[string]Reserva  `json:"reservas"`

	// Índices para búsquedas rápidas
	indicesTitulo map[string][]string `json:"indices_titulo"` // título -> []libro_id
	indicesAutor  map[string][]string `json:"indices_autor"`  // autor -> []libro_id
	indicesGenero map[string][]string `json:"indices_genero"` // género -> []libro_id
	indicesISBN   map[string]string   `json:"indices_isbn"`   // isbn -> libro_id

	indicesEmail    map[string]string   `json:"indices_email"`     // email -> usuario_id
	indicesTipoUser map[string][]string `json:"indices_tipo_user"` // tipo -> []usuario_id

	// Índice invertido para búsqueda de texto
	indiceInvertido map[string]map[string]int `json:"indice_invertido"` // palabra -> libro_id -> frecuencia

	// Sistema de caché
	cache *CacheBiblioteca `json:"cache"`

	// Configuración
	config *ConfigBiblioteca `json:"config"`

	// Thread safety
	mu sync.RWMutex `json:"-"`

	// Contadores para IDs únicos
	contadorLibros    int64 `json:"contador_libros"`
	contadorUsuarios  int64 `json:"contador_usuarios"`
	contadorPrestamos int64 `json:"contador_prestamos"`
	contadorReservas  int64 `json:"contador_reservas"`
}

// Constructor
func NewBiblioteca() *Biblioteca {
	return &Biblioteca{
		libros:    make(map[string]Libro),
		usuarios:  make(map[string]Usuario),
		prestamos: make(map[string]Prestamo),
		reservas:  make(map[string]Reserva),

		indicesTitulo: make(map[string][]string),
		indicesAutor:  make(map[string][]string),
		indicesGenero: make(map[string][]string),
		indicesISBN:   make(map[string]string),

		indicesEmail:    make(map[string]string),
		indicesTipoUser: make(map[string][]string),

		indiceInvertido: make(map[string]map[string]int),

		cache: &CacheBiblioteca{
			busquedas:    make(map[string][]Libro),
			estadisticas: make(map[string]interface{}),
			ttl:          30 * time.Minute,
			lastUpdate:   make(map[string]time.Time),
		},

		config: &ConfigBiblioteca{
			MaxPrestamosEstudiante: 3,
			MaxPrestamosProfesor:   5,
			MaxPrestamosExterno:    2,
			DiasPrestamoEstudiante: 14,
			DiasPrestamoProfesor:   21,
			DiasPrestamoExterno:    7,
			MultaDiaria:            0.50,
			MaxRenovaciones:        2,
			DiasReserva:            3,
		},
	}
}

// ==============================================
// GESTIÓN DE LIBROS
// ==============================================

func (b *Biblioteca) AgregarLibro(libro Libro) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Generar ID único si no tiene
	if libro.ID == "" {
		b.contadorLibros++
		libro.ID = fmt.Sprintf("LIB_%06d", b.contadorLibros)
	}

	// Validaciones básicas
	if libro.Titulo == "" {
		return fmt.Errorf("el título es requerido")
	}
	if libro.ISBN == "" {
		return fmt.Errorf("el ISBN es requerido")
	}

	// Verificar ISBN único
	if existeID, existe := b.indicesISBN[libro.ISBN]; existe && existeID != libro.ID {
		return fmt.Errorf("ya existe un libro con ISBN %s", libro.ISBN)
	}

	// Establecer valores por defecto
	if libro.FechaAdq.IsZero() {
		libro.FechaAdq = time.Now()
	}
	if libro.Estado == "" {
		libro.Estado = "bueno"
	}
	if libro.Idioma == "" {
		libro.Idioma = "español"
	}
	libro.Disponible = true

	// Guardar en almacenamiento principal
	b.libros[libro.ID] = libro

	// Actualizar índices
	b.actualizarIndicesLibro(libro)

	// Invalidar caché de búsquedas
	b.invalidarCacheBusquedas()

	return nil
}

func (b *Biblioteca) actualizarIndicesLibro(libro Libro) {
	// Índice de título (normalizado)
	tituloNorm := strings.ToLower(libro.Titulo)
	if !contiene(b.indicesTitulo[tituloNorm], libro.ID) {
		b.indicesTitulo[tituloNorm] = append(b.indicesTitulo[tituloNorm], libro.ID)
	}

	// Índices de autores
	for _, autor := range libro.Autores {
		autorNorm := strings.ToLower(autor)
		if !contiene(b.indicesAutor[autorNorm], libro.ID) {
			b.indicesAutor[autorNorm] = append(b.indicesAutor[autorNorm], libro.ID)
		}
	}

	// Índices de géneros
	for _, genero := range libro.Generos {
		generoNorm := strings.ToLower(genero)
		if !contiene(b.indicesGenero[generoNorm], libro.ID) {
			b.indicesGenero[generoNorm] = append(b.indicesGenero[generoNorm], libro.ID)
		}
	}

	// Índice de ISBN
	b.indicesISBN[libro.ISBN] = libro.ID

	// Actualizar índice invertido
	b.actualizarIndiceInvertido(libro)
}

func (b *Biblioteca) actualizarIndiceInvertido(libro Libro) {
	// Extraer todas las palabras del libro
	texto := strings.ToLower(libro.Titulo + " " + libro.Descripcion + " " + strings.Join(libro.Autores, " "))
	palabras := strings.Fields(texto)

	for _, palabra := range palabras {
		// Limpiar palabra de puntuación
		palabra = limpiarPalabra(palabra)
		if len(palabra) < 3 { // Ignorar palabras muy cortas
			continue
		}

		if b.indiceInvertido[palabra] == nil {
			b.indiceInvertido[palabra] = make(map[string]int)
		}
		b.indiceInvertido[palabra][libro.ID]++
	}
}

func (b *Biblioteca) ActualizarLibro(id string, libro Libro) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Verificar que existe
	if _, existe := b.libros[id]; !existe {
		return fmt.Errorf("libro con ID %s no encontrado", id)
	}

	// Mantener el ID original
	libro.ID = id

	// Limpiar índices antiguos
	b.limpiarIndicesLibro(id)

	// Actualizar libro
	b.libros[id] = libro

	// Recrear índices
	b.actualizarIndicesLibro(libro)

	// Invalidar caché
	b.invalidarCacheBusquedas()

	return nil
}

func (b *Biblioteca) EliminarLibro(id string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Verificar que existe
	libro, existe := b.libros[id]
	if !existe {
		return fmt.Errorf("libro con ID %s no encontrado", id)
	}

	// Verificar que no esté prestado
	for _, prestamo := range b.prestamos {
		if prestamo.LibroID == id && prestamo.Estado == "activo" {
			return fmt.Errorf("no se puede eliminar: libro está prestado (préstamo %s)", prestamo.ID)
		}
	}

	// Eliminar de almacenamiento
	delete(b.libros, id)

	// Limpiar índices
	b.limpiarIndicesLibro(id)

	// Limpiar índice invertido
	for palabra, docs := range b.indiceInvertido {
		delete(docs, id)
		if len(docs) == 0 {
			delete(b.indiceInvertido, palabra)
		}
	}

	// Cancelar reservas pendientes
	for reservaID, reserva := range b.reservas {
		if reserva.LibroID == id && reserva.Estado == "pendiente" {
			reserva.Estado = "cancelada"
			b.reservas[reservaID] = reserva
		}
	}

	// Invalidar caché
	b.invalidarCacheBusquedas()

	fmt.Printf("Libro '%s' eliminado exitosamente\n", libro.Titulo)
	return nil
}

func (b *Biblioteca) limpiarIndicesLibro(libroID string) {
	libro := b.libros[libroID]

	// Limpiar índice de título
	tituloNorm := strings.ToLower(libro.Titulo)
	b.indicesTitulo[tituloNorm] = removerString(b.indicesTitulo[tituloNorm], libroID)
	if len(b.indicesTitulo[tituloNorm]) == 0 {
		delete(b.indicesTitulo, tituloNorm)
	}

	// Limpiar índices de autores
	for _, autor := range libro.Autores {
		autorNorm := strings.ToLower(autor)
		b.indicesAutor[autorNorm] = removerString(b.indicesAutor[autorNorm], libroID)
		if len(b.indicesAutor[autorNorm]) == 0 {
			delete(b.indicesAutor, autorNorm)
		}
	}

	// Limpiar índices de géneros
	for _, genero := range libro.Generos {
		generoNorm := strings.ToLower(genero)
		b.indicesGenero[generoNorm] = removerString(b.indicesGenero[generoNorm], libroID)
		if len(b.indicesGenero[generoNorm]) == 0 {
			delete(b.indicesGenero, generoNorm)
		}
	}

	// Limpiar índice de ISBN
	delete(b.indicesISBN, libro.ISBN)
}

func (b *Biblioteca) ObtenerLibro(id string) (Libro, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	libro, existe := b.libros[id]
	return libro, existe
}

// ==============================================
// BÚSQUEDAS DE LIBROS
// ==============================================

func (b *Biblioteca) BuscarPorTitulo(titulo string) []Libro {
	b.mu.RLock()
	defer b.mu.RUnlock()

	tituloNorm := strings.ToLower(titulo)

	// Verificar caché primero
	cacheKey := "titulo:" + tituloNorm
	if resultados := b.obtenerDesdCache(cacheKey); resultados != nil {
		return resultados
	}

	var resultados []Libro

	// Búsqueda exacta
	if ids, existe := b.indicesTitulo[tituloNorm]; existe {
		for _, id := range ids {
			if libro, existe := b.libros[id]; existe {
				resultados = append(resultados, libro)
			}
		}
	}

	// Búsqueda parcial si no hay resultados exactos
	if len(resultados) == 0 {
		for tituloIndice, ids := range b.indicesTitulo {
			if strings.Contains(tituloIndice, tituloNorm) {
				for _, id := range ids {
					if libro, existe := b.libros[id]; existe {
						resultados = append(resultados, libro)
					}
				}
			}
		}
	}

	// Guardar en caché
	b.guardarEnCache(cacheKey, resultados)

	return resultados
}

func (b *Biblioteca) BuscarPorAutor(autor string) []Libro {
	b.mu.RLock()
	defer b.mu.RUnlock()

	autorNorm := strings.ToLower(autor)

	// Verificar caché
	cacheKey := "autor:" + autorNorm
	if resultados := b.obtenerDesdCache(cacheKey); resultados != nil {
		return resultados
	}

	var resultados []Libro

	// Búsqueda exacta
	if ids, existe := b.indicesAutor[autorNorm]; existe {
		for _, id := range ids {
			if libro, existe := b.libros[id]; existe {
				resultados = append(resultados, libro)
			}
		}
	}

	// Búsqueda parcial
	if len(resultados) == 0 {
		for autorIndice, ids := range b.indicesAutor {
			if strings.Contains(autorIndice, autorNorm) {
				for _, id := range ids {
					if libro, existe := b.libros[id]; existe && !contieneLibro(resultados, libro) {
						resultados = append(resultados, libro)
					}
				}
			}
		}
	}

	// Guardar en caché
	b.guardarEnCache(cacheKey, resultados)

	return resultados
}

func (b *Biblioteca) BuscarPorGenero(genero string) []Libro {
	b.mu.RLock()
	defer b.mu.RUnlock()

	generoNorm := strings.ToLower(genero)

	var resultados []Libro

	if ids, existe := b.indicesGenero[generoNorm]; existe {
		for _, id := range ids {
			if libro, existe := b.libros[id]; existe {
				resultados = append(resultados, libro)
			}
		}
	}

	return resultados
}

func (b *Biblioteca) BuscarPorISBN(isbn string) (Libro, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if id, existe := b.indicesISBN[isbn]; existe {
		if libro, existe := b.libros[id]; existe {
			return libro, true
		}
	}

	return Libro{}, false
}

func (b *Biblioteca) BuscarTexto(query string) []Libro {
	b.mu.RLock()
	defer b.mu.RUnlock()

	queryNorm := strings.ToLower(query)
	palabras := strings.Fields(queryNorm)

	if len(palabras) == 0 {
		return []Libro{}
	}

	// Verificar caché
	cacheKey := "texto:" + queryNorm
	if resultados := b.obtenerDesdCache(cacheKey); resultados != nil {
		return resultados
	}

	// Contar relevancia por documento
	relevancia := make(map[string]int)

	for _, palabra := range palabras {
		palabra = limpiarPalabra(palabra)
		if docs, existe := b.indiceInvertido[palabra]; existe {
			for docID, frecuencia := range docs {
				relevancia[docID] += frecuencia
			}
		}
	}

	// Convertir a slice y ordenar por relevancia
	type resultado struct {
		libro      Libro
		relevancia int
	}

	var resultados []resultado
	for libroID, rel := range relevancia {
		if libro, existe := b.libros[libroID]; existe {
			resultados = append(resultados, resultado{libro, rel})
		}
	}

	// Ordenar por relevancia descendente
	sort.Slice(resultados, func(i, j int) bool {
		return resultados[i].relevancia > resultados[j].relevancia
	})

	// Extraer solo los libros
	var librosFinales []Libro
	for _, res := range resultados {
		librosFinales = append(librosFinales, res.libro)
	}

	// Guardar en caché
	b.guardarEnCache(cacheKey, librosFinales)

	return librosFinales
}

// ==============================================
// GESTIÓN DE USUARIOS
// ==============================================

func (b *Biblioteca) RegistrarUsuario(usuario Usuario) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Generar ID único si no tiene
	if usuario.ID == "" {
		b.contadorUsuarios++
		usuario.ID = fmt.Sprintf("USR_%06d", b.contadorUsuarios)
	}

	// Validaciones
	if usuario.Email == "" {
		return fmt.Errorf("el email es requerido")
	}
	if usuario.Nombre == "" {
		return fmt.Errorf("el nombre es requerido")
	}

	// Verificar email único
	if existeID, existe := b.indicesEmail[usuario.Email]; existe && existeID != usuario.ID {
		return fmt.Errorf("ya existe un usuario con email %s", usuario.Email)
	}

	// Establecer valores por defecto
	if usuario.FechaReg.IsZero() {
		usuario.FechaReg = time.Now()
	}
	if usuario.TipoUsuario == "" {
		usuario.TipoUsuario = "estudiante"
	}
	usuario.Activo = true
	usuario.Multas = 0

	// Establecer límite según tipo
	switch usuario.TipoUsuario {
	case "estudiante":
		usuario.LimiteLibros = b.config.MaxPrestamosEstudiante
	case "profesor":
		usuario.LimiteLibros = b.config.MaxPrestamosProfesor
	case "externo":
		usuario.LimiteLibros = b.config.MaxPrestamosExterno
	default:
		usuario.LimiteLibros = b.config.MaxPrestamosEstudiante
	}

	if usuario.Historial == nil {
		usuario.Historial = []string{}
	}
	if usuario.Preferencias == nil {
		usuario.Preferencias = make(map[string]interface{})
	}

	// Guardar usuario
	b.usuarios[usuario.ID] = usuario

	// Actualizar índices
	b.indicesEmail[usuario.Email] = usuario.ID
	b.indicesTipoUser[usuario.TipoUsuario] = append(b.indicesTipoUser[usuario.TipoUsuario], usuario.ID)

	return nil
}

func (b *Biblioteca) ObtenerUsuario(id string) (Usuario, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	usuario, existe := b.usuarios[id]
	return usuario, existe
}

func (b *Biblioteca) BuscarUsuarioPorEmail(email string) (Usuario, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if id, existe := b.indicesEmail[email]; existe {
		if usuario, existe := b.usuarios[id]; existe {
			return usuario, true
		}
	}

	return Usuario{}, false
}

// ==============================================
// SISTEMA DE PRÉSTAMOS
// ==============================================

func (b *Biblioteca) CrearPrestamo(usuarioID, libroID string) (string, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Validar usuario
	usuario, existe := b.usuarios[usuarioID]
	if !existe {
		return "", fmt.Errorf("usuario %s no encontrado", usuarioID)
	}
	if !usuario.Activo {
		return "", fmt.Errorf("usuario %s está inactivo", usuarioID)
	}

	// Validar libro
	libro, existe := b.libros[libroID]
	if !existe {
		return "", fmt.Errorf("libro %s no encontrado", libroID)
	}
	if !libro.Disponible {
		return "", fmt.Errorf("libro %s no está disponible", libroID)
	}

	// Verificar límites del usuario
	prestamosActivos := b.contarPrestamosActivos(usuarioID)
	if prestamosActivos >= usuario.LimiteLibros {
		return "", fmt.Errorf("usuario ha alcanzado el límite de préstamos (%d)", usuario.LimiteLibros)
	}

	// Verificar multas pendientes
	if usuario.Multas > 0 {
		return "", fmt.Errorf("usuario tiene multas pendientes ($%.2f)", usuario.Multas)
	}

	// Crear préstamo
	b.contadorPrestamos++
	prestamoID := fmt.Sprintf("PREST_%06d", b.contadorPrestamos)

	// Calcular fecha de vencimiento según tipo de usuario
	var diasPrestamo int
	switch usuario.TipoUsuario {
	case "estudiante":
		diasPrestamo = b.config.DiasPrestamoEstudiante
	case "profesor":
		diasPrestamo = b.config.DiasPrestamoProfesor
	case "externo":
		diasPrestamo = b.config.DiasPrestamoExterno
	default:
		diasPrestamo = b.config.DiasPrestamoEstudiante
	}

	fechaVenc := time.Now().AddDate(0, 0, diasPrestamo)

	prestamo := Prestamo{
		ID:            prestamoID,
		UsuarioID:     usuarioID,
		LibroID:       libroID,
		FechaPrestamo: time.Now(),
		FechaVenc:     fechaVenc,
		Estado:        "activo",
		Renovaciones:  0,
		Multa:         0,
	}

	// Guardar préstamo
	b.prestamos[prestamoID] = prestamo

	// Marcar libro como no disponible
	libro.Disponible = false
	b.libros[libroID] = libro

	// Actualizar historial del usuario
	usuario.Historial = append(usuario.Historial, prestamoID)
	b.usuarios[usuarioID] = usuario

	fmt.Printf("Préstamo creado: %s prestó '%s' hasta %s\n",
		usuario.Nombre, libro.Titulo, fechaVenc.Format("2006-01-02"))

	return prestamoID, nil
}

func (b *Biblioteca) DevolverLibro(prestamoID string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Verificar préstamo
	prestamo, existe := b.prestamos[prestamoID]
	if !existe {
		return fmt.Errorf("préstamo %s no encontrado", prestamoID)
	}
	if prestamo.Estado != "activo" {
		return fmt.Errorf("préstamo %s ya fue devuelto", prestamoID)
	}

	// Marcar como devuelto
	ahora := time.Now()
	prestamo.FechaDevol = &ahora
	prestamo.Estado = "devuelto"

	// Calcular multa si está vencido
	if ahora.After(prestamo.FechaVenc) {
		diasVencido := int(ahora.Sub(prestamo.FechaVenc).Hours() / 24)
		multa := float64(diasVencido) * b.config.MultaDiaria
		prestamo.Multa = multa

		// Agregar multa al usuario
		usuario := b.usuarios[prestamo.UsuarioID]
		usuario.Multas += multa
		b.usuarios[prestamo.UsuarioID] = usuario

		fmt.Printf("⚠️ Libro devuelto con %d días de retraso. Multa: $%.2f\n", diasVencido, multa)
	}

	// Actualizar préstamo
	b.prestamos[prestamoID] = prestamo

	// Marcar libro como disponible
	libro := b.libros[prestamo.LibroID]
	libro.Disponible = true
	b.libros[prestamo.LibroID] = libro

	// Procesar reservas pendientes
	b.procesarReservasPendientes(prestamo.LibroID)

	fmt.Printf("✅ Libro '%s' devuelto exitosamente\n", libro.Titulo)

	return nil
}

func (b *Biblioteca) contarPrestamosActivos(usuarioID string) int {
	count := 0
	for _, prestamo := range b.prestamos {
		if prestamo.UsuarioID == usuarioID && prestamo.Estado == "activo" {
			count++
		}
	}
	return count
}

func (b *Biblioteca) procesarReservasPendientes(libroID string) {
	// Buscar reservas pendientes para este libro
	var reservasPendientes []Reserva

	for _, reserva := range b.reservas {
		if reserva.LibroID == libroID && reserva.Estado == "pendiente" {
			reservasPendientes = append(reservasPendientes, reserva)
		}
	}

	if len(reservasPendientes) == 0 {
		return
	}

	// Ordenar por prioridad
	sort.Slice(reservasPendientes, func(i, j int) bool {
		return reservasPendientes[i].Prioridad < reservasPendientes[j].Prioridad
	})

	// Notificar al primer usuario en la cola
	primeraReserva := reservasPendientes[0]
	primeraReserva.Estado = "disponible"
	b.reservas[primeraReserva.ID] = primeraReserva

	usuario := b.usuarios[primeraReserva.UsuarioID]
	libro := b.libros[libroID]

	fmt.Printf("📬 Reserva disponible: %s puede retirar '%s' (Reserva ID: %s)\n",
		usuario.Nombre, libro.Titulo, primeraReserva.ID)
}

// ==============================================
// UTILIDADES Y HELPERS
// ==============================================

func contiene(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func removerString(slice []string, item string) []string {
	for i, s := range slice {
		if s == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func contieneLibro(libros []Libro, libro Libro) bool {
	for _, l := range libros {
		if l.ID == libro.ID {
			return true
		}
	}
	return false
}

func limpiarPalabra(palabra string) string {
	return strings.Trim(palabra, ".,!?;:()[]{}\"'")
}

// Sistema de caché
func (b *Biblioteca) obtenerDesdCache(clave string) []Libro {
	b.cache.mu.RLock()
	defer b.cache.mu.RUnlock()

	if lastUpdate, existe := b.cache.lastUpdate[clave]; existe {
		if time.Since(lastUpdate) < b.cache.ttl {
			if resultados, existe := b.cache.busquedas[clave]; existe {
				return resultados
			}
		}
	}

	return nil
}

func (b *Biblioteca) guardarEnCache(clave string, resultados []Libro) {
	b.cache.mu.Lock()
	defer b.cache.mu.Unlock()

	b.cache.busquedas[clave] = resultados
	b.cache.lastUpdate[clave] = time.Now()
}

func (b *Biblioteca) invalidarCacheBusquedas() {
	b.cache.mu.Lock()
	defer b.cache.mu.Unlock()

	// Limpiar todo el caché de búsquedas
	b.cache.busquedas = make(map[string][]Libro)
	b.cache.lastUpdate = make(map[string]time.Time)
}

// ==============================================
// EJEMPLO DE USO
// ==============================================

func ejemploUso() {
	fmt.Println("🏛️ SISTEMA DE GESTIÓN DE BIBLIOTECA DIGITAL")
	fmt.Println("============================================")

	// Crear biblioteca
	biblioteca := NewBiblioteca()

	// Agregar algunos libros
	libros := []Libro{
		{
			Titulo:      "El Quijote de la Mancha",
			Autores:     []string{"Miguel de Cervantes"},
			Generos:     []string{"Novela", "Clásico"},
			ISBN:        "978-84-376-0494-7",
			Editorial:   "Real Academia Española",
			AnoPublic:   1605,
			Paginas:     1200,
			Descripcion: "La obra cumbre de la literatura española",
			Estado:      "bueno",
			Precio:      25.99,
		},
		{
			Titulo:      "Programación en Go",
			Autores:     []string{"Alan Donovan", "Brian Kernighan"},
			Generos:     []string{"Tecnología", "Programación"},
			ISBN:        "978-0-13-419044-0",
			Editorial:   "Addison-Wesley",
			AnoPublic:   2015,
			Paginas:     380,
			Descripcion: "Guía completa del lenguaje Go",
			Estado:      "nuevo",
			Precio:      45.50,
		},
		{
			Titulo:      "Cien años de soledad",
			Autores:     []string{"Gabriel García Márquez"},
			Generos:     []string{"Novela", "Realismo mágico"},
			ISBN:        "978-84-376-0495-4",
			Editorial:   "Sudamericana",
			AnoPublic:   1967,
			Paginas:     471,
			Descripcion: "Obra maestra del realismo mágico",
			Estado:      "bueno",
			Precio:      18.75,
		},
	}

	for _, libro := range libros {
		if err := biblioteca.AgregarLibro(libro); err != nil {
			fmt.Printf("Error agregando libro: %v\n", err)
		} else {
			fmt.Printf("✅ Libro agregado: %s\n", libro.Titulo)
		}
	}

	// Registrar usuarios
	usuarios := []Usuario{
		{
			Email:       "ana.garcia@universidad.edu",
			Nombre:      "Ana",
			Apellido:    "García",
			TipoUsuario: "estudiante",
			FechaNac:    time.Date(2000, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			Email:       "carlos.lopez@universidad.edu",
			Nombre:      "Carlos",
			Apellido:    "López",
			TipoUsuario: "profesor",
			FechaNac:    time.Date(1975, 8, 22, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, usuario := range usuarios {
		if err := biblioteca.RegistrarUsuario(usuario); err != nil {
			fmt.Printf("Error registrando usuario: %v\n", err)
		} else {
			fmt.Printf("✅ Usuario registrado: %s %s (%s)\n",
				usuario.Nombre, usuario.Apellido, usuario.TipoUsuario)
		}
	}

	// Realizar búsquedas
	fmt.Println("\n🔍 BÚSQUEDAS")
	fmt.Println("-------------")

	// Búsqueda por título
	resultados := biblioteca.BuscarPorTitulo("Quijote")
	fmt.Printf("Búsqueda por título 'Quijote': %d resultados\n", len(resultados))
	for _, libro := range resultados {
		fmt.Printf("  - %s por %s\n", libro.Titulo, strings.Join(libro.Autores, ", "))
	}

	// Búsqueda por autor
	resultados = biblioteca.BuscarPorAutor("García Márquez")
	fmt.Printf("Búsqueda por autor 'García Márquez': %d resultados\n", len(resultados))
	for _, libro := range resultados {
		fmt.Printf("  - %s\n", libro.Titulo)
	}

	// Búsqueda de texto completo
	resultados = biblioteca.BuscarTexto("programación go lenguaje")
	fmt.Printf("Búsqueda de texto 'programación go lenguaje': %d resultados\n", len(resultados))
	for _, libro := range resultados {
		fmt.Printf("  - %s\n", libro.Titulo)
	}

	// Préstamos
	fmt.Println("\n📚 PRÉSTAMOS")
	fmt.Println("-------------")

	// Buscar usuarios e IDs de libros para préstamos
	var usuarioID, libroID string

	for id, usuario := range biblioteca.usuarios {
		if usuario.Email == "ana.garcia@universidad.edu" {
			usuarioID = id
			break
		}
	}

	for id, libro := range biblioteca.libros {
		if libro.Titulo == "Programación en Go" {
			libroID = id
			break
		}
	}

	if usuarioID != "" && libroID != "" {
		if prestamoID, err := biblioteca.CrearPrestamo(usuarioID, libroID); err != nil {
			fmt.Printf("Error creando préstamo: %v\n", err)
		} else {
			fmt.Printf("✅ Préstamo creado con ID: %s\n", prestamoID)

			// Simular devolución después de un tiempo
			fmt.Println("\n📤 Simulando devolución...")
			if err := biblioteca.DevolverLibro(prestamoID); err != nil {
				fmt.Printf("Error devolviendo libro: %v\n", err)
			}
		}
	}

	// Mostrar estadísticas
	fmt.Println("\n📊 ESTADÍSTICAS")
	fmt.Println("----------------")
	fmt.Printf("Total libros: %d\n", len(biblioteca.libros))
	fmt.Printf("Total usuarios: %d\n", len(biblioteca.usuarios))
	fmt.Printf("Total préstamos: %d\n", len(biblioteca.prestamos))

	librosDisponibles := 0
	for _, libro := range biblioteca.libros {
		if libro.Disponible {
			librosDisponibles++
		}
	}
	fmt.Printf("Libros disponibles: %d\n", librosDisponibles)
}

func main() {
	ejemploUso()
}
