// Archivo: proyecto_ecommerce.go
// Proyecto: Sistema de E-commerce Completo usando Structs
// Demuestra: embedding, tags, validación, patrones de diseño, manejo de estados

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// ==============================================
// ENTIDADES BASE Y COMPONENTES COMPARTIDOS
// ==============================================

// Componente base para entidades con ID
type Identificable struct {
	ID string `json:"id"`
}

// Componente base para timestamping
type Timestampable struct {
	CreadoEn      time.Time  `json:"created_at"`
	ActualizadoEn time.Time  `json:"updated_at"`
	EliminadoEn   *time.Time `json:"deleted_at,omitempty"`
}

func (t *Timestampable) MarcarCreado() {
	now := time.Now()
	t.CreadoEn = now
	t.ActualizadoEn = now
}

func (t *Timestampable) MarcarActualizado() {
	t.ActualizadoEn = time.Now()
}

func (t *Timestampable) MarcarEliminado() {
	now := time.Now()
	t.EliminadoEn = &now
}

func (t Timestampable) EstaEliminado() bool {
	return t.EliminadoEn != nil
}

// Componente para direcciones
type Direccion struct {
	Nombre    string `json:"nombre" validate:"required"`
	Calle     string `json:"calle" validate:"required"`
	Ciudad    string `json:"ciudad" validate:"required"`
	Estado    string `json:"estado" validate:"required"`
	CodigoP   string `json:"codigo_postal" validate:"required"`
	Pais      string `json:"pais" validate:"required"`
	Telefono  string `json:"telefono"`
	Principal bool   `json:"principal"`
}

func (d Direccion) String() string {
	return fmt.Sprintf("%s\n%s\n%s, %s %s\n%s",
		d.Nombre, d.Calle, d.Ciudad, d.Estado, d.CodigoP, d.Pais)
}

func (d Direccion) Validar() error {
	if strings.TrimSpace(d.Nombre) == "" {
		return errors.New("nombre es requerido")
	}
	if strings.TrimSpace(d.Calle) == "" {
		return errors.New("calle es requerida")
	}
	if strings.TrimSpace(d.Ciudad) == "" {
		return errors.New("ciudad es requerida")
	}
	return nil
}

// ==============================================
// SISTEMA DE USUARIOS
// ==============================================

type Usuario struct {
	Identificable   `json:",inline"`
	Timestampable   `json:",inline"`
	Email           string                 `json:"email" validate:"required,email"`
	Username        string                 `json:"username" validate:"required,min=3"`
	PasswordHash    string                 `json:"-"`
	Nombre          string                 `json:"nombre" validate:"required"`
	Apellido        string                 `json:"apellido" validate:"required"`
	FechaNacimiento time.Time              `json:"fecha_nacimiento"`
	Telefono        string                 `json:"telefono"`
	Direcciones     []Direccion            `json:"direcciones"`
	Activo          bool                   `json:"activo"`
	Verificado      bool                   `json:"verificado"`
	Preferencias    map[string]interface{} `json:"preferencias"`
}

func NewUsuario(email, username, nombre, apellido string) *Usuario {
	usuario := &Usuario{
		Identificable: Identificable{ID: generarID("USR")},
		Email:         email,
		Username:      username,
		Nombre:        nombre,
		Apellido:      apellido,
		Direcciones:   []Direccion{},
		Activo:        true,
		Verificado:    false,
		Preferencias:  make(map[string]interface{}),
	}
	usuario.MarcarCreado()
	return usuario
}

func (u Usuario) NombreCompleto() string {
	return fmt.Sprintf("%s %s", u.Nombre, u.Apellido)
}

func (u Usuario) Validar() error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return errors.New("email no válido")
	}

	if len(u.Username) < 3 {
		return errors.New("username debe tener al menos 3 caracteres")
	}

	if strings.TrimSpace(u.Nombre) == "" {
		return errors.New("nombre es requerido")
	}

	return nil
}

func (u *Usuario) AgregarDireccion(direccion Direccion) {
	// Si es la primera dirección, marcarla como principal
	if len(u.Direcciones) == 0 {
		direccion.Principal = true
	}
	u.Direcciones = append(u.Direcciones, direccion)
	u.MarcarActualizado()
}

func (u Usuario) DireccionPrincipal() (Direccion, bool) {
	for _, dir := range u.Direcciones {
		if dir.Principal {
			return dir, true
		}
	}
	return Direccion{}, false
}

// ==============================================
// SISTEMA DE PRODUCTOS
// ==============================================

type Categoria struct {
	Identificable `json:",inline"`
	Timestampable `json:",inline"`
	Nombre        string `json:"nombre" validate:"required"`
	Descripcion   string `json:"descripcion"`
	Slug          string `json:"slug" validate:"required"`
	PadreID       string `json:"padre_id,omitempty"`
	Activa        bool   `json:"activa"`
}

type Variante struct {
	ID          string                 `json:"id"`
	Nombre      string                 `json:"nombre"` // ej: "Talla", "Color"
	Valor       string                 `json:"valor"`  // ej: "M", "Rojo"
	PrecioExtra float64                `json:"precio_extra"`
	Stock       int                    `json:"stock"`
	Atributos   map[string]interface{} `json:"atributos"`
}

type Review struct {
	Identificable `json:",inline"`
	Timestampable `json:",inline"`
	UsuarioID     string `json:"usuario_id"`
	ProductoID    string `json:"producto_id"`
	Calificacion  int    `json:"calificacion" validate:"min=1,max=5"`
	Titulo        string `json:"titulo"`
	Comentario    string `json:"comentario"`
	Verificado    bool   `json:"verificado"`
	Util          int    `json:"votos_util"`
}

type Producto struct {
	Identificable `json:",inline"`
	Timestampable `json:",inline"`
	Nombre        string                 `json:"nombre" validate:"required"`
	Descripcion   string                 `json:"descripcion"`
	SKU           string                 `json:"sku" validate:"required"`
	CategoriaID   string                 `json:"categoria_id" validate:"required"`
	Precio        float64                `json:"precio" validate:"min=0"`
	PrecioOferta  float64                `json:"precio_oferta,omitempty"`
	Stock         int                    `json:"stock" validate:"min=0"`
	StockMinimo   int                    `json:"stock_minimo"`
	Peso          float64                `json:"peso"`
	Dimensiones   map[string]float64     `json:"dimensiones"`
	Imagenes      []string               `json:"imagenes"`
	Variantes     []Variante             `json:"variantes"`
	Tags          []string               `json:"tags"`
	Atributos     map[string]interface{} `json:"atributos"`
	Activo        bool                   `json:"activo"`
	Destacado     bool                   `json:"destacado"`
	Reviews       []Review               `json:"reviews"`
}

func NewProducto(nombre, descripcion, sku, categoriaID string, precio float64, stock int) *Producto {
	producto := &Producto{
		Identificable: Identificable{ID: generarID("PRD")},
		Nombre:        nombre,
		Descripcion:   descripcion,
		SKU:           sku,
		CategoriaID:   categoriaID,
		Precio:        precio,
		Stock:         stock,
		StockMinimo:   5,
		Dimensiones:   make(map[string]float64),
		Imagenes:      []string{},
		Variantes:     []Variante{},
		Tags:          []string{},
		Atributos:     make(map[string]interface{}),
		Activo:        true,
		Reviews:       []Review{},
	}
	producto.MarcarCreado()
	return producto
}

func (p Producto) TieneStock(cantidad int) bool {
	return p.Stock >= cantidad
}

func (p Producto) PrecioFinal() float64 {
	if p.PrecioOferta > 0 && p.PrecioOferta < p.Precio {
		return p.PrecioOferta
	}
	return p.Precio
}

func (p Producto) DescuentoPorcentaje() float64 {
	if p.PrecioOferta > 0 && p.PrecioOferta < p.Precio {
		return ((p.Precio - p.PrecioOferta) / p.Precio) * 100
	}
	return 0
}

func (p Producto) CalificacionPromedio() float64 {
	if len(p.Reviews) == 0 {
		return 0
	}

	total := 0
	for _, review := range p.Reviews {
		total += review.Calificacion
	}

	return float64(total) / float64(len(p.Reviews))
}

func (p *Producto) AgregarReview(review Review) {
	review.ProductoID = p.ID
	review.MarcarCreado()
	p.Reviews = append(p.Reviews, review)
	p.MarcarActualizado()
}

func (p *Producto) ReducirStock(cantidad int) error {
	if !p.TieneStock(cantidad) {
		return fmt.Errorf("stock insuficiente: disponible %d, solicitado %d", p.Stock, cantidad)
	}
	p.Stock -= cantidad
	p.MarcarActualizado()
	return nil
}

func (p *Producto) AumentarStock(cantidad int) {
	p.Stock += cantidad
	p.MarcarActualizado()
}

// ==============================================
// SISTEMA DE CARRITO DE COMPRAS
// ==============================================

type ItemCarrito struct {
	ProductoID     string                 `json:"producto_id"`
	VarianteID     string                 `json:"variante_id,omitempty"`
	Cantidad       int                    `json:"cantidad" validate:"min=1"`
	PrecioUnitario float64                `json:"precio_unitario"`
	Descuento      float64                `json:"descuento"`
	Atributos      map[string]interface{} `json:"atributos"`
	FechaAgregado  time.Time              `json:"fecha_agregado"`
}

func (i ItemCarrito) Subtotal() float64 {
	return (i.PrecioUnitario * float64(i.Cantidad)) - i.Descuento
}

type Carrito struct {
	Identificable  `json:",inline"`
	Timestampable  `json:",inline"`
	UsuarioID      string        `json:"usuario_id"`
	Items          []ItemCarrito `json:"items"`
	CuponAplicado  string        `json:"cupon_aplicado,omitempty"`
	DescuentoCupon float64       `json:"descuento_cupon"`
	SessionID      string        `json:"session_id"`
	Activo         bool          `json:"activo"`
}

func NewCarrito(usuarioID string) *Carrito {
	carrito := &Carrito{
		Identificable: Identificable{ID: generarID("CART")},
		UsuarioID:     usuarioID,
		Items:         []ItemCarrito{},
		Activo:        true,
		SessionID:     generarSessionID(),
	}
	carrito.MarcarCreado()
	return carrito
}

func (c *Carrito) AgregarItem(productoID string, cantidad int, precioUnitario float64) {
	// Buscar si el producto ya existe en el carrito
	for i, item := range c.Items {
		if item.ProductoID == productoID && item.VarianteID == "" {
			c.Items[i].Cantidad += cantidad
			c.Items[i].FechaAgregado = time.Now()
			c.MarcarActualizado()
			return
		}
	}

	// Agregar nuevo item
	nuevoItem := ItemCarrito{
		ProductoID:     productoID,
		Cantidad:       cantidad,
		PrecioUnitario: precioUnitario,
		Atributos:      make(map[string]interface{}),
		FechaAgregado:  time.Now(),
	}

	c.Items = append(c.Items, nuevoItem)
	c.MarcarActualizado()
}

func (c *Carrito) EliminarItem(productoID string) {
	for i, item := range c.Items {
		if item.ProductoID == productoID {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			c.MarcarActualizado()
			return
		}
	}
}

func (c *Carrito) ModificarCantidad(productoID string, nuevaCantidad int) {
	if nuevaCantidad <= 0 {
		c.EliminarItem(productoID)
		return
	}

	for i, item := range c.Items {
		if item.ProductoID == productoID {
			c.Items[i].Cantidad = nuevaCantidad
			c.Items[i].FechaAgregado = time.Now()
			c.MarcarActualizado()
			return
		}
	}
}

func (c Carrito) Subtotal() float64 {
	total := 0.0
	for _, item := range c.Items {
		total += item.Subtotal()
	}
	return total
}

func (c Carrito) DescuentoTotal() float64 {
	descuento := 0.0
	for _, item := range c.Items {
		descuento += item.Descuento
	}
	return descuento + c.DescuentoCupon
}

func (c Carrito) Total() float64 {
	return c.Subtotal() - c.DescuentoTotal()
}

func (c Carrito) TotalItems() int {
	total := 0
	for _, item := range c.Items {
		total += item.Cantidad
	}
	return total
}

func (c *Carrito) AplicarCupon(codigo string, descuento float64) {
	c.CuponAplicado = codigo
	c.DescuentoCupon = descuento
	c.MarcarActualizado()
}

func (c *Carrito) Vaciar() {
	c.Items = []ItemCarrito{}
	c.CuponAplicado = ""
	c.DescuentoCupon = 0
	c.MarcarActualizado()
}

// ==============================================
// SISTEMA DE ÓRDENES
// ==============================================

type EstadoOrden string

const (
	OrdenPendiente   EstadoOrden = "pendiente"
	OrdenProcesando  EstadoOrden = "procesando"
	OrdenConfirmada  EstadoOrden = "confirmada"
	OrdenEnEnvio     EstadoOrden = "en_envio"
	OrdenEntregada   EstadoOrden = "entregada"
	OrdenCancelada   EstadoOrden = "cancelada"
	OrdenReembolsada EstadoOrden = "reembolsada"
)

type ItemOrden struct {
	ProductoID     string  `json:"producto_id"`
	VarianteID     string  `json:"variante_id,omitempty"`
	Nombre         string  `json:"nombre"`
	SKU            string  `json:"sku"`
	Cantidad       int     `json:"cantidad"`
	PrecioUnitario float64 `json:"precio_unitario"`
	Descuento      float64 `json:"descuento"`
	Subtotal       float64 `json:"subtotal"`
}

type MetodoPago struct {
	Tipo           string                 `json:"tipo"` // "tarjeta", "paypal", "transferencia"
	Referencia     string                 `json:"referencia"`
	UltimosDigitos string                 `json:"ultimos_digitos,omitempty"`
	Estado         string                 `json:"estado"`
	Atributos      map[string]interface{} `json:"atributos"`
}

type Envio struct {
	Metodo         string     `json:"metodo"`
	Costo          float64    `json:"costo"`
	TiempoEstimado string     `json:"tiempo_estimado"`
	Direccion      Direccion  `json:"direccion"`
	NumeroTracking string     `json:"numero_tracking,omitempty"`
	FechaEnvio     *time.Time `json:"fecha_envio,omitempty"`
	FechaEntrega   *time.Time `json:"fecha_entrega,omitempty"`
}

type Orden struct {
	Identificable    `json:",inline"`
	Timestampable    `json:",inline"`
	NumeroOrden      string            `json:"numero_orden"`
	UsuarioID        string            `json:"usuario_id"`
	Items            []ItemOrden       `json:"items"`
	Estado           EstadoOrden       `json:"estado"`
	Subtotal         float64           `json:"subtotal"`
	DescuentoTotal   float64           `json:"descuento_total"`
	CostoEnvio       float64           `json:"costo_envio"`
	Impuestos        float64           `json:"impuestos"`
	Total            float64           `json:"total"`
	MetodoPago       MetodoPago        `json:"metodo_pago"`
	Envio            Envio             `json:"envio"`
	CuponAplicado    string            `json:"cupon_aplicado,omitempty"`
	Notas            string            `json:"notas"`
	HistorialEstados []EstadoHistorial `json:"historial_estados"`
}

type EstadoHistorial struct {
	Estado     EstadoOrden `json:"estado"`
	Fecha      time.Time   `json:"fecha"`
	Comentario string      `json:"comentario"`
	UsuarioID  string      `json:"usuario_id"`
}

func NewOrden(usuarioID string, carrito *Carrito) *Orden {
	numeroOrden := generarNumeroOrden()

	orden := &Orden{
		Identificable:    Identificable{ID: generarID("ORD")},
		NumeroOrden:      numeroOrden,
		UsuarioID:        usuarioID,
		Items:            []ItemOrden{},
		Estado:           OrdenPendiente,
		Subtotal:         carrito.Subtotal(),
		DescuentoTotal:   carrito.DescuentoTotal(),
		Total:            carrito.Total(),
		CuponAplicado:    carrito.CuponAplicado,
		HistorialEstados: []EstadoHistorial{},
	}

	// Convertir items del carrito a items de orden
	for _, item := range carrito.Items {
		itemOrden := ItemOrden{
			ProductoID:     item.ProductoID,
			VarianteID:     item.VarianteID,
			Cantidad:       item.Cantidad,
			PrecioUnitario: item.PrecioUnitario,
			Descuento:      item.Descuento,
			Subtotal:       item.Subtotal(),
		}
		orden.Items = append(orden.Items, itemOrden)
	}

	orden.MarcarCreado()
	orden.AgregarEstadoHistorial(OrdenPendiente, "Orden creada", usuarioID)

	return orden
}

func (o *Orden) CambiarEstado(nuevoEstado EstadoOrden, comentario, usuarioID string) {
	o.Estado = nuevoEstado
	o.AgregarEstadoHistorial(nuevoEstado, comentario, usuarioID)
	o.MarcarActualizado()
}

func (o *Orden) AgregarEstadoHistorial(estado EstadoOrden, comentario, usuarioID string) {
	historial := EstadoHistorial{
		Estado:     estado,
		Fecha:      time.Now(),
		Comentario: comentario,
		UsuarioID:  usuarioID,
	}
	o.HistorialEstados = append(o.HistorialEstados, historial)
}

func (o *Orden) PuedeSerCancelada() bool {
	return o.Estado == OrdenPendiente || o.Estado == OrdenProcesando
}

func (o *Orden) Cancelar(motivo, usuarioID string) error {
	if !o.PuedeSerCancelada() {
		return fmt.Errorf("la orden en estado %s no puede ser cancelada", o.Estado)
	}

	o.CambiarEstado(OrdenCancelada, fmt.Sprintf("Cancelada: %s", motivo), usuarioID)
	return nil
}

func (o *Orden) MarcarComoEnviada(numeroTracking, usuarioID string) {
	o.Envio.NumeroTracking = numeroTracking
	fechaEnvio := time.Now()
	o.Envio.FechaEnvio = &fechaEnvio
	o.CambiarEstado(OrdenEnEnvio, fmt.Sprintf("Enviada con tracking: %s", numeroTracking), usuarioID)
}

func (o *Orden) MarcarComoEntregada(usuarioID string) {
	fechaEntrega := time.Now()
	o.Envio.FechaEntrega = &fechaEntrega
	o.CambiarEstado(OrdenEntregada, "Orden entregada exitosamente", usuarioID)
}

// ==============================================
// SISTEMA DE INVENTARIO
// ==============================================

type MovimientoInventario struct {
	Identificable  `json:",inline"`
	Timestampable  `json:",inline"`
	ProductoID     string `json:"producto_id"`
	TipoMovimiento string `json:"tipo_movimiento"` // "entrada", "salida", "reserva", "liberacion"
	Cantidad       int    `json:"cantidad"`
	StockAnterior  int    `json:"stock_anterior"`
	StockNuevo     int    `json:"stock_nuevo"`
	Motivo         string `json:"motivo"`
	ReferenciaID   string `json:"referencia_id"` // ID de orden, carrito, etc.
	UsuarioID      string `json:"usuario_id"`
}

type Inventario struct {
	productos   map[string]*Producto
	movimientos []MovimientoInventario
	reservas    map[string]ReservaStock
}

type ReservaStock struct {
	ProductoID  string
	Cantidad    int
	UsuarioID   string
	FechaExpira time.Time
	Activa      bool
}

func NewInventario() *Inventario {
	return &Inventario{
		productos:   make(map[string]*Producto),
		movimientos: []MovimientoInventario{},
		reservas:    make(map[string]ReservaStock),
	}
}

func (inv *Inventario) AgregarProducto(producto *Producto) {
	inv.productos[producto.ID] = producto
}

func (inv *Inventario) ReservarStock(productoID, usuarioID string, cantidad int, duracion time.Duration) error {
	producto, existe := inv.productos[productoID]
	if !existe {
		return errors.New("producto no encontrado")
	}

	if !producto.TieneStock(cantidad) {
		return fmt.Errorf("stock insuficiente: disponible %d, solicitado %d", producto.Stock, cantidad)
	}

	// Crear reserva
	reservaID := generarID("RES")
	reserva := ReservaStock{
		ProductoID:  productoID,
		Cantidad:    cantidad,
		UsuarioID:   usuarioID,
		FechaExpira: time.Now().Add(duracion),
		Activa:      true,
	}

	inv.reservas[reservaID] = reserva

	// Reducir stock temporalmente
	if err := producto.ReducirStock(cantidad); err != nil {
		return err
	}

	// Registrar movimiento
	inv.registrarMovimiento(productoID, "reserva", cantidad,
		producto.Stock+cantidad, producto.Stock, "Reserva de stock", reservaID, usuarioID)

	return nil
}

func (inv *Inventario) LiberarReserva(reservaID string) error {
	reserva, existe := inv.reservas[reservaID]
	if !existe {
		return errors.New("reserva no encontrada")
	}

	if !reserva.Activa {
		return errors.New("reserva ya liberada")
	}

	producto := inv.productos[reserva.ProductoID]
	producto.AumentarStock(reserva.Cantidad)

	// Marcar reserva como inactiva
	reserva.Activa = false
	inv.reservas[reservaID] = reserva

	// Registrar movimiento
	inv.registrarMovimiento(reserva.ProductoID, "liberacion", reserva.Cantidad,
		producto.Stock-reserva.Cantidad, producto.Stock, "Liberación de reserva", reservaID, reserva.UsuarioID)

	return nil
}

func (inv *Inventario) registrarMovimiento(productoID, tipo string, cantidad, stockAnterior, stockNuevo int, motivo, referenciaID, usuarioID string) {
	movimiento := MovimientoInventario{
		Identificable:  Identificable{ID: generarID("MOV")},
		ProductoID:     productoID,
		TipoMovimiento: tipo,
		Cantidad:       cantidad,
		StockAnterior:  stockAnterior,
		StockNuevo:     stockNuevo,
		Motivo:         motivo,
		ReferenciaID:   referenciaID,
		UsuarioID:      usuarioID,
	}
	movimiento.MarcarCreado()

	inv.movimientos = append(inv.movimientos, movimiento)
}

func (inv *Inventario) LimpiarReservasExpiradas() int {
	liberadas := 0
	ahora := time.Now()

	for reservaID, reserva := range inv.reservas {
		if reserva.Activa && ahora.After(reserva.FechaExpira) {
			if err := inv.LiberarReserva(reservaID); err == nil {
				liberadas++
			}
		}
	}

	return liberadas
}

// ==============================================
// SISTEMA PRINCIPAL DE E-COMMERCE
// ==============================================

type Ecommerce struct {
	usuarios   map[string]*Usuario
	productos  map[string]*Producto
	categorias map[string]*Categoria
	carritos   map[string]*Carrito
	ordenes    map[string]*Orden
	inventario *Inventario
}

func NewEcommerce() *Ecommerce {
	return &Ecommerce{
		usuarios:   make(map[string]*Usuario),
		productos:  make(map[string]*Producto),
		categorias: make(map[string]*Categoria),
		carritos:   make(map[string]*Carrito),
		ordenes:    make(map[string]*Orden),
		inventario: NewInventario(),
	}
}

func (e *Ecommerce) RegistrarUsuario(email, username, nombre, apellido string) (*Usuario, error) {
	// Verificar que el email no exista
	for _, usuario := range e.usuarios {
		if usuario.Email == email {
			return nil, errors.New("email ya registrado")
		}
		if usuario.Username == username {
			return nil, errors.New("username ya existe")
		}
	}

	usuario := NewUsuario(email, username, nombre, apellido)
	if err := usuario.Validar(); err != nil {
		return nil, err
	}

	e.usuarios[usuario.ID] = usuario
	return usuario, nil
}

func (e *Ecommerce) CrearProducto(nombre, descripcion, sku, categoriaID string, precio float64, stock int) (*Producto, error) {
	// Verificar que la categoría exista
	if _, existe := e.categorias[categoriaID]; !existe {
		return nil, errors.New("categoría no encontrada")
	}

	// Verificar que el SKU no exista
	for _, producto := range e.productos {
		if producto.SKU == sku {
			return nil, errors.New("SKU ya existe")
		}
	}

	producto := NewProducto(nombre, descripcion, sku, categoriaID, precio, stock)
	e.productos[producto.ID] = producto
	e.inventario.AgregarProducto(producto)

	return producto, nil
}

func (e *Ecommerce) ObtenerCarritoUsuario(usuarioID string) *Carrito {
	// Buscar carrito activo del usuario
	for _, carrito := range e.carritos {
		if carrito.UsuarioID == usuarioID && carrito.Activo {
			return carrito
		}
	}

	// Crear nuevo carrito si no existe
	carrito := NewCarrito(usuarioID)
	e.carritos[carrito.ID] = carrito
	return carrito
}

func (e *Ecommerce) AgregarAlCarrito(usuarioID, productoID string, cantidad int) error {
	producto, existe := e.productos[productoID]
	if !existe {
		return errors.New("producto no encontrado")
	}

	if !producto.TieneStock(cantidad) {
		return fmt.Errorf("stock insuficiente: disponible %d, solicitado %d", producto.Stock, cantidad)
	}

	carrito := e.ObtenerCarritoUsuario(usuarioID)
	carrito.AgregarItem(productoID, cantidad, producto.PrecioFinal())

	return nil
}

func (e *Ecommerce) ProcesarOrden(usuarioID string, direccionEnvio Direccion, metodoPago MetodoPago) (*Orden, error) {
	carrito := e.ObtenerCarritoUsuario(usuarioID)
	if len(carrito.Items) == 0 {
		return nil, errors.New("carrito vacío")
	}

	// Validar stock disponible
	for _, item := range carrito.Items {
		producto := e.productos[item.ProductoID]
		if !producto.TieneStock(item.Cantidad) {
			return nil, fmt.Errorf("stock insuficiente para producto %s", producto.Nombre)
		}
	}

	// Crear orden
	orden := NewOrden(usuarioID, carrito)
	orden.MetodoPago = metodoPago
	orden.Envio = Envio{
		Direccion: direccionEnvio,
		Costo:     calcularCostoEnvio(direccionEnvio, carrito.Total()),
	}
	orden.Total += orden.Envio.Costo

	// Reservar stock
	for _, item := range carrito.Items {
		err := e.inventario.ReservarStock(item.ProductoID, usuarioID, item.Cantidad, 24*time.Hour)
		if err != nil {
			return nil, fmt.Errorf("error reservando stock: %v", err)
		}
	}

	// Guardar orden y limpiar carrito
	e.ordenes[orden.ID] = orden
	carrito.Vaciar()

	// Cambiar estado a procesando
	orden.CambiarEstado(OrdenProcesando, "Orden en procesamiento", usuarioID)

	return orden, nil
}

func (e *Ecommerce) ConfirmarPago(ordenID string) error {
	orden, existe := e.ordenes[ordenID]
	if !existe {
		return errors.New("orden no encontrada")
	}

	if orden.Estado != OrdenProcesando {
		return fmt.Errorf("orden en estado %s no puede ser confirmada", orden.Estado)
	}

	orden.CambiarEstado(OrdenConfirmada, "Pago confirmado exitosamente", "sistema")
	return nil
}

// ==============================================
// FUNCIONES AUXILIARES
// ==============================================

var contadorGlobal = 0

func generarID(prefijo string) string {
	contadorGlobal++
	return fmt.Sprintf("%s_%06d", prefijo, contadorGlobal)
}

func generarSessionID() string {
	return fmt.Sprintf("SESS_%d", time.Now().Unix())
}

func generarNumeroOrden() string {
	return fmt.Sprintf("ORD-%d-%04d", time.Now().Year(), contadorGlobal+1000)
}

func calcularCostoEnvio(direccion Direccion, subtotal float64) float64 {
	// Lógica simple de cálculo de envío
	costoBase := 10.0
	if subtotal > 100 {
		return 0 // Envío gratis para compras mayores a $100
	}
	return costoBase
}

// ==============================================
// DEMOSTRACIÓN DEL SISTEMA
// ==============================================

func main() {
	fmt.Println("🛒 SISTEMA DE E-COMMERCE COMPLETO")
	fmt.Println("==================================")

	// Crear sistema de e-commerce
	ecommerce := NewEcommerce()

	// Crear categorías
	categoriaElectronicos := &Categoria{
		Identificable: Identificable{ID: generarID("CAT")},
		Nombre:        "Electrónicos",
		Descripcion:   "Dispositivos electrónicos",
		Slug:          "electronicos",
		Activa:        true,
	}
	categoriaElectronicos.MarcarCreado()
	ecommerce.categorias[categoriaElectronicos.ID] = categoriaElectronicos

	categoriaRopa := &Categoria{
		Identificable: Identificable{ID: generarID("CAT")},
		Nombre:        "Ropa",
		Descripcion:   "Ropa y accesorios",
		Slug:          "ropa",
		Activa:        true,
	}
	categoriaRopa.MarcarCreado()
	ecommerce.categorias[categoriaRopa.ID] = categoriaRopa

	fmt.Printf("✅ Categorías creadas: %s, %s\n", categoriaElectronicos.Nombre, categoriaRopa.Nombre)

	// Registrar usuarios
	usuario1, err := ecommerce.RegistrarUsuario("ana@ejemplo.com", "ana_garcia", "Ana", "García")
	if err != nil {
		fmt.Printf("Error registrando usuario: %v\n", err)
		return
	}

	usuario2, err := ecommerce.RegistrarUsuario("carlos@ejemplo.com", "carlos_lopez", "Carlos", "López")
	if err != nil {
		fmt.Printf("Error registrando usuario: %v\n", err)
		return
	}

	fmt.Printf("✅ Usuarios registrados: %s, %s\n", usuario1.NombreCompleto(), usuario2.NombreCompleto())

	// Agregar direcciones
	direccion1 := Direccion{
		Nombre:    "Ana García",
		Calle:     "Av. Principal 123",
		Ciudad:    "Madrid",
		Estado:    "Madrid",
		CodigoP:   "28001",
		Pais:      "España",
		Telefono:  "+34612345678",
		Principal: true,
	}
	usuario1.AgregarDireccion(direccion1)

	// Crear productos
	laptop, err := ecommerce.CrearProducto(
		"Laptop Gaming Pro",
		"Laptop para gaming de alta gama con procesador Intel i7",
		"LAP-001",
		categoriaElectronicos.ID,
		1299.99,
		10,
	)
	if err != nil {
		fmt.Printf("Error creando producto: %v\n", err)
		return
	}

	mouse, err := ecommerce.CrearProducto(
		"Mouse Gaming RGB",
		"Mouse óptico con iluminación RGB programable",
		"MOU-001",
		categoriaElectronicos.ID,
		79.99,
		50,
	)
	if err != nil {
		fmt.Printf("Error creando producto: %v\n", err)
		return
	}

	camiseta, err := ecommerce.CrearProducto(
		"Camiseta Básica",
		"Camiseta de algodón 100% en varios colores",
		"CAM-001",
		categoriaRopa.ID,
		19.99,
		100,
	)
	if err != nil {
		fmt.Printf("Error creando producto: %v\n", err)
		return
	}

	fmt.Printf("✅ Productos creados: %s ($%.2f), %s ($%.2f), %s ($%.2f)\n",
		laptop.Nombre, laptop.Precio,
		mouse.Nombre, mouse.Precio,
		camiseta.Nombre, camiseta.Precio)

	// Agregar productos al carrito
	fmt.Println("\n🛒 Simulando compra...")

	err = ecommerce.AgregarAlCarrito(usuario1.ID, laptop.ID, 1)
	if err != nil {
		fmt.Printf("Error agregando al carrito: %v\n", err)
		return
	}

	err = ecommerce.AgregarAlCarrito(usuario1.ID, mouse.ID, 2)
	if err != nil {
		fmt.Printf("Error agregando al carrito: %v\n", err)
		return
	}

	err = ecommerce.AgregarAlCarrito(usuario1.ID, camiseta.ID, 3)
	if err != nil {
		fmt.Printf("Error agregando al carrito: %v\n", err)
		return
	}

	// Obtener carrito y mostrar resumen
	carrito := ecommerce.ObtenerCarritoUsuario(usuario1.ID)
	fmt.Printf("📦 Items en carrito: %d\n", carrito.TotalItems())
	fmt.Printf("💰 Subtotal: $%.2f\n", carrito.Subtotal())
	fmt.Printf("💰 Total: $%.2f\n", carrito.Total())

	// Aplicar cupón de descuento
	carrito.AplicarCupon("DESCUENTO10", 50.00)
	fmt.Printf("🎫 Cupón aplicado: %s (descuento: $%.2f)\n", carrito.CuponAplicado, carrito.DescuentoCupon)
	fmt.Printf("💰 Nuevo total: $%.2f\n", carrito.Total())

	// Procesar orden
	metodoPago := MetodoPago{
		Tipo:           "tarjeta",
		Referencia:     "VISA-4532",
		UltimosDigitos: "1234",
		Estado:         "aprobado",
		Atributos:      make(map[string]interface{}),
	}

	orden, err := ecommerce.ProcesarOrden(usuario1.ID, direccion1, metodoPago)
	if err != nil {
		fmt.Printf("Error procesando orden: %v\n", err)
		return
	}

	fmt.Printf("\n📋 Orden creada: %s\n", orden.NumeroOrden)
	fmt.Printf("🏷️ Estado: %s\n", orden.Estado)
	fmt.Printf("📦 Items: %d\n", len(orden.Items))
	fmt.Printf("💰 Total final: $%.2f\n", orden.Total)

	// Confirmar pago
	err = ecommerce.ConfirmarPago(orden.ID)
	if err != nil {
		fmt.Printf("Error confirmando pago: %v\n", err)
		return
	}

	fmt.Printf("✅ Pago confirmado, estado actual: %s\n", orden.Estado)

	// Simular envío
	orden.MarcarComoEnviada("TRACK123456789", "sistema")
	fmt.Printf("📦 Orden marcada como enviada, tracking: %s\n", orden.Envio.NumeroTracking)

	// Simular entrega
	time.Sleep(1 * time.Second) // Simular tiempo
	orden.MarcarComoEntregada("sistema")
	fmt.Printf("✅ Orden entregada exitosamente\n")

	// Mostrar historial de estados
	fmt.Println("\n📊 Historial de estados:")
	for _, estado := range orden.HistorialEstados {
		fmt.Printf("  %s: %s (%s)\n",
			estado.Fecha.Format("2006-01-02 15:04:05"),
			estado.Estado,
			estado.Comentario)
	}

	// Agregar review al producto
	review := Review{
		Identificable: Identificable{ID: generarID("REV")},
		UsuarioID:     usuario1.ID,
		Calificacion:  5,
		Titulo:        "Excelente producto",
		Comentario:    "La laptop funciona perfectamente, muy recomendada",
		Verificado:    true,
	}
	laptop.AgregarReview(review)

	fmt.Printf("\n⭐ Review agregada a %s\n", laptop.Nombre)
	fmt.Printf("📊 Calificación promedio: %.1f/5\n", laptop.CalificacionPromedio())

	// Mostrar inventario
	fmt.Println("\n📦 Estado del inventario:")
	for id, producto := range ecommerce.productos {
		fmt.Printf("  %s: %d unidades (mín: %d)\n",
			producto.Nombre, producto.Stock, producto.StockMinimo)
		_ = id
	}

	// Limpiar reservas expiradas
	liberadas := ecommerce.inventario.LimpiarReservasExpiradas()
	fmt.Printf("🔄 Reservas expiradas liberadas: %d\n", liberadas)

	// Serializar datos a JSON (ejemplo)
	usuarioJSON, _ := json.MarshalIndent(usuario1, "", "  ")
	fmt.Println("\n📄 Datos del usuario (JSON):")
	fmt.Println(string(usuarioJSON))

	fmt.Println("\n🎉 ¡Sistema de e-commerce funcionando completamente!")
	fmt.Println("✅ Todas las funcionalidades con structs demostradas exitosamente")
}
