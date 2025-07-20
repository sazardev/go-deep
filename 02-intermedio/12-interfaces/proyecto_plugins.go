// Archivo: proyecto_plugins.go
// Proyecto: Sistema de Plugins Modulares con Interfaces
// Demuestra el poder de las interfaces en Go mediante un sistema extensible

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

// ==============================================
// INTERFACES BASE DEL SISTEMA
// ==============================================

// Plugin base que todos los plugins deben implementar
type Plugin interface {
	Name() string
	Version() string
	Initialize(config map[string]interface{}) error
	Shutdown() error
	IsEnabled() bool
}

// Plugin con metadatos extendidos
type PluginInfo interface {
	Plugin
	Description() string
	Author() string
	Dependencies() []string
}

// ==============================================
// INTERFACES ESPECÍFICAS POR FUNCIONALIDAD
// ==============================================

// Niveles de log
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Interface para procesadores de datos
type DataProcessor interface {
	PluginInfo
	SupportedFormats() []string
	Process(data []byte, format string) ([]byte, error)
	Validate(data []byte, format string) error
}

// Interface para loggers
type Logger interface {
	PluginInfo
	Log(level LogLevel, message string, fields map[string]interface{})
	SetLevel(level LogLevel)
	GetLevel() LogLevel
}

// Tipos para autenticación
type User struct {
	ID       string
	Username string
	Email    string
	Roles    []string
}

// Interface para autenticadores
type Authenticator interface {
	PluginInfo
	Authenticate(credentials map[string]string) (User, error)
	ValidateToken(token string) (User, error)
	RefreshToken(token string) (string, error)
}

// Tipos para notificaciones
type MessageType string

const (
	EMAIL MessageType = "email"
	SMS   MessageType = "sms"
	PUSH  MessageType = "push"
	SLACK MessageType = "slack"
)

type Message struct {
	ID       string
	Subject  string
	Body     string
	Type     MessageType
	Priority int
}

type DeliveryStatus struct {
	Delivered bool
	Timestamp time.Time
	Error     string
}

// Interface para notificadores
type Notifier interface {
	PluginInfo
	Send(recipient string, message Message) error
	SupportedTypes() []MessageType
	GetDeliveryStatus(messageID string) DeliveryStatus
}

// ==============================================
// SISTEMA CORE - GESTIÓN DE PLUGINS
// ==============================================

type PluginManager struct {
	plugins        map[string]PluginInfo
	processors     map[string]DataProcessor
	loggers        map[string]Logger
	authenticators map[string]Authenticator
	notifiers      map[string]Notifier
	mu             sync.RWMutex
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins:        make(map[string]PluginInfo),
		processors:     make(map[string]DataProcessor),
		loggers:        make(map[string]Logger),
		authenticators: make(map[string]Authenticator),
		notifiers:      make(map[string]Notifier),
	}
}

func (pm *PluginManager) RegisterPlugin(plugin PluginInfo) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	name := plugin.Name()
	pm.plugins[name] = plugin

	// Registrar en categoría específica según tipo
	if processor, ok := plugin.(DataProcessor); ok {
		pm.processors[name] = processor
	}
	if logger, ok := plugin.(Logger); ok {
		pm.loggers[name] = logger
	}
	if auth, ok := plugin.(Authenticator); ok {
		pm.authenticators[name] = auth
	}
	if notifier, ok := plugin.(Notifier); ok {
		pm.notifiers[name] = notifier
	}

	return plugin.Initialize(map[string]interface{}{
		"plugin_name": name,
		"timestamp":   time.Now(),
	})
}

func (pm *PluginManager) GetProcessor(name string) (DataProcessor, bool) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	processor, exists := pm.processors[name]
	return processor, exists
}

func (pm *PluginManager) GetLogger(name string) (Logger, bool) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	logger, exists := pm.loggers[name]
	return logger, exists
}

func (pm *PluginManager) GetAuthenticator(name string) (Authenticator, bool) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	auth, exists := pm.authenticators[name]
	return auth, exists
}

func (pm *PluginManager) GetNotifier(name string) (Notifier, bool) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	notifier, exists := pm.notifiers[name]
	return notifier, exists
}

func (pm *PluginManager) ListPlugins() []string {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	var names []string
	for name := range pm.plugins {
		names = append(names, name)
	}
	return names
}

// ==============================================
// IMPLEMENTACIÓN: DATA PROCESSORS
// ==============================================

// Base común para processors
type BaseProcessor struct {
	name         string
	version      string
	description  string
	author       string
	dependencies []string
	enabled      bool
}

func (bp *BaseProcessor) Name() string           { return bp.name }
func (bp *BaseProcessor) Version() string        { return bp.version }
func (bp *BaseProcessor) Description() string    { return bp.description }
func (bp *BaseProcessor) Author() string         { return bp.author }
func (bp *BaseProcessor) Dependencies() []string { return bp.dependencies }
func (bp *BaseProcessor) IsEnabled() bool        { return bp.enabled }

func (bp *BaseProcessor) Initialize(config map[string]interface{}) error {
	bp.enabled = true
	return nil
}

func (bp *BaseProcessor) Shutdown() error {
	bp.enabled = false
	return nil
}

// JSON Processor
type JSONProcessor struct {
	BaseProcessor
}

func NewJSONProcessor() *JSONProcessor {
	return &JSONProcessor{
		BaseProcessor: BaseProcessor{
			name:        "JSONProcessor",
			version:     "1.0.0",
			description: "Procesador de datos JSON",
			author:      "Go Deep Team",
		},
	}
}

func (jp *JSONProcessor) SupportedFormats() []string {
	return []string{"json"}
}

func (jp *JSONProcessor) Process(data []byte, format string) ([]byte, error) {
	if format != "json" {
		return nil, fmt.Errorf("formato no soportado: %s", format)
	}

	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, fmt.Errorf("error parseando JSON: %v", err)
	}

	// Procesar (en este caso, pretty print)
	processed, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error formateando JSON: %v", err)
	}

	return processed, nil
}

func (jp *JSONProcessor) Validate(data []byte, format string) error {
	if format != "json" {
		return fmt.Errorf("formato no soportado: %s", format)
	}

	var jsonData interface{}
	return json.Unmarshal(data, &jsonData)
}

// XML Processor (simulado)
type XMLProcessor struct {
	BaseProcessor
}

func NewXMLProcessor() *XMLProcessor {
	return &XMLProcessor{
		BaseProcessor: BaseProcessor{
			name:        "XMLProcessor",
			version:     "1.0.0",
			description: "Procesador de datos XML",
			author:      "Go Deep Team",
		},
	}
}

func (xp *XMLProcessor) SupportedFormats() []string {
	return []string{"xml"}
}

func (xp *XMLProcessor) Process(data []byte, format string) ([]byte, error) {
	if format != "xml" {
		return nil, fmt.Errorf("formato no soportado: %s", format)
	}

	// Simulación de procesamiento XML
	processed := []byte(fmt.Sprintf("<!-- Procesado por XMLProcessor -->\n%s", string(data)))
	return processed, nil
}

func (xp *XMLProcessor) Validate(data []byte, format string) error {
	if format != "xml" {
		return fmt.Errorf("formato no soportado: %s", format)
	}

	// Validación básica de XML
	content := string(data)
	if !strings.Contains(content, "<") || !strings.Contains(content, ">") {
		return fmt.Errorf("formato XML inválido")
	}

	return nil
}

// ==============================================
// IMPLEMENTACIÓN: LOGGERS
// ==============================================

// Base común para loggers
type BaseLogger struct {
	BaseProcessor
	level LogLevel
}

func (bl *BaseLogger) SetLevel(level LogLevel) {
	bl.level = level
}

func (bl *BaseLogger) GetLevel() LogLevel {
	return bl.level
}

// Console Logger
type ConsoleLogger struct {
	BaseLogger
}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{
		BaseLogger: BaseLogger{
			BaseProcessor: BaseProcessor{
				name:        "ConsoleLogger",
				version:     "1.0.0",
				description: "Logger que envía output a la consola",
				author:      "Go Deep Team",
			},
			level: INFO,
		},
	}
}

func (cl *ConsoleLogger) Log(level LogLevel, message string, fields map[string]interface{}) {
	if level < cl.level {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	levelStr := level.String()

	// Colores para diferentes niveles
	var color string
	switch level {
	case DEBUG:
		color = "\033[36m" // Cyan
	case INFO:
		color = "\033[32m" // Green
	case WARN:
		color = "\033[33m" // Yellow
	case ERROR:
		color = "\033[31m" // Red
	case FATAL:
		color = "\033[35m" // Magenta
	}
	reset := "\033[0m"

	fmt.Printf("%s[%s] %s%s%s: %s", timestamp, levelStr, color, levelStr, reset, message)

	if len(fields) > 0 {
		fieldsJSON, _ := json.Marshal(fields)
		fmt.Printf(" | Fields: %s", string(fieldsJSON))
	}
	fmt.Println()
}

// File Logger (simulado)
type FileLogger struct {
	BaseLogger
	filename string
}

func NewFileLogger() *FileLogger {
	return &FileLogger{
		BaseLogger: BaseLogger{
			BaseProcessor: BaseProcessor{
				name:        "FileLogger",
				version:     "1.0.0",
				description: "Logger que escribe a archivos",
				author:      "Go Deep Team",
			},
			level: INFO,
		},
		filename: "app.log",
	}
}

func (fl *FileLogger) Log(level LogLevel, message string, fields map[string]interface{}) {
	if level < fl.level {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s: %s", timestamp, level.String(), message)

	if len(fields) > 0 {
		fieldsJSON, _ := json.Marshal(fields)
		logEntry += fmt.Sprintf(" | %s", string(fieldsJSON))
	}

	// Simulamos escritura a archivo
	fmt.Printf("📄 [FileLogger -> %s] %s\n", fl.filename, logEntry)
}

// ==============================================
// IMPLEMENTACIÓN: AUTHENTICATORS
// ==============================================

// JWT Authenticator (simulado)
type JWTAuthenticator struct {
	BaseProcessor
	secretKey string
}

func NewJWTAuthenticator() *JWTAuthenticator {
	return &JWTAuthenticator{
		BaseProcessor: BaseProcessor{
			name:        "JWTAuthenticator",
			version:     "1.0.0",
			description: "Autenticador basado en JWT",
			author:      "Go Deep Team",
		},
		secretKey: "super-secret-key",
	}
}

func (ja *JWTAuthenticator) Authenticate(credentials map[string]string) (User, error) {
	username, ok := credentials["username"]
	if !ok {
		return User{}, fmt.Errorf("username requerido")
	}

	password, ok := credentials["password"]
	if !ok {
		return User{}, fmt.Errorf("password requerido")
	}

	// Simulación de autenticación
	if username == "admin" && password == "secret" {
		return User{
			ID:       "1",
			Username: username,
			Email:    "admin@example.com",
			Roles:    []string{"admin", "user"},
		}, nil
	}

	if username == "user" && password == "pass" {
		return User{
			ID:       "2",
			Username: username,
			Email:    "user@example.com",
			Roles:    []string{"user"},
		}, nil
	}

	return User{}, fmt.Errorf("credenciales inválidas")
}

func (ja *JWTAuthenticator) ValidateToken(token string) (User, error) {
	// Simulación de validación de token
	if token == "valid-jwt-token" {
		return User{
			ID:       "1",
			Username: "admin",
			Email:    "admin@example.com",
			Roles:    []string{"admin", "user"},
		}, nil
	}

	return User{}, fmt.Errorf("token inválido")
}

func (ja *JWTAuthenticator) RefreshToken(token string) (string, error) {
	user, err := ja.ValidateToken(token)
	if err != nil {
		return "", err
	}

	// Simulación de refresh
	newToken := fmt.Sprintf("new-jwt-token-for-%s", user.Username)
	return newToken, nil
}

// ==============================================
// IMPLEMENTACIÓN: NOTIFIERS
// ==============================================

// Email Notifier
type EmailNotifier struct {
	BaseProcessor
	smtpServer string
	port       int
}

func NewEmailNotifier() *EmailNotifier {
	return &EmailNotifier{
		BaseProcessor: BaseProcessor{
			name:        "EmailNotifier",
			version:     "1.0.0",
			description: "Notificador vía email",
			author:      "Go Deep Team",
		},
		smtpServer: "smtp.example.com",
		port:       587,
	}
}

func (en *EmailNotifier) Send(recipient string, message Message) error {
	// Simulación de envío de email
	fmt.Printf("📧 [EmailNotifier] Enviando email a: %s\n", recipient)
	fmt.Printf("   Asunto: %s\n", message.Subject)
	fmt.Printf("   Mensaje: %s\n", message.Body)
	fmt.Printf("   Servidor: %s:%d\n", en.smtpServer, en.port)

	return nil
}

func (en *EmailNotifier) SupportedTypes() []MessageType {
	return []MessageType{EMAIL}
}

func (en *EmailNotifier) GetDeliveryStatus(messageID string) DeliveryStatus {
	return DeliveryStatus{
		Delivered: true,
		Timestamp: time.Now(),
		Error:     "",
	}
}

// Slack Notifier
type SlackNotifier struct {
	BaseProcessor
	webhookURL string
	channel    string
}

func NewSlackNotifier() *SlackNotifier {
	return &SlackNotifier{
		BaseProcessor: BaseProcessor{
			name:        "SlackNotifier",
			version:     "1.0.0",
			description: "Notificador vía Slack",
			author:      "Go Deep Team",
		},
		webhookURL: "https://hooks.slack.com/services/...",
		channel:    "#general",
	}
}

func (sn *SlackNotifier) Send(recipient string, message Message) error {
	// Simulación de envío a Slack
	fmt.Printf("💬 [SlackNotifier] Enviando a Slack\n")
	fmt.Printf("   Canal: %s\n", sn.channel)
	fmt.Printf("   Mensaje: %s\n", message.Body)
	fmt.Printf("   Webhook: %s\n", sn.webhookURL)

	return nil
}

func (sn *SlackNotifier) SupportedTypes() []MessageType {
	return []MessageType{SLACK}
}

func (sn *SlackNotifier) GetDeliveryStatus(messageID string) DeliveryStatus {
	return DeliveryStatus{
		Delivered: true,
		Timestamp: time.Now(),
		Error:     "",
	}
}

// ==============================================
// SISTEMA DE PIPELINE DE PROCESAMIENTO
// ==============================================

type ProcessingPipeline struct {
	manager *PluginManager
	steps   []PipelineStep
}

type PipelineStep struct {
	Name       string
	PluginName string
	PluginType string
	Config     map[string]interface{}
}

func NewProcessingPipeline(manager *PluginManager) *ProcessingPipeline {
	return &ProcessingPipeline{
		manager: manager,
		steps:   make([]PipelineStep, 0),
	}
}

func (pp *ProcessingPipeline) AddStep(name, pluginName, pluginType string, config map[string]interface{}) {
	pp.steps = append(pp.steps, PipelineStep{
		Name:       name,
		PluginName: pluginName,
		PluginType: pluginType,
		Config:     config,
	})
}

func (pp *ProcessingPipeline) Execute(data []byte, format string) error {
	fmt.Printf("🚀 Iniciando pipeline de procesamiento...\n")
	fmt.Printf("📊 Datos de entrada: %d bytes, formato: %s\n\n", len(data), format)

	currentData := data

	for i, step := range pp.steps {
		fmt.Printf("📍 Paso %d: %s (Plugin: %s)\n", i+1, step.Name, step.PluginName)

		switch step.PluginType {
		case "processor":
			if processor, exists := pp.manager.GetProcessor(step.PluginName); exists {
				processed, err := processor.Process(currentData, format)
				if err != nil {
					return fmt.Errorf("error en paso %d: %v", i+1, err)
				}
				currentData = processed
				fmt.Printf("   ✅ Procesado exitosamente (%d bytes)\n", len(processed))
			}

		case "logger":
			if logger, exists := pp.manager.GetLogger(step.PluginName); exists {
				logger.Log(INFO, fmt.Sprintf("Pipeline paso %d completado", i+1), map[string]interface{}{
					"step":      step.Name,
					"data_size": len(currentData),
				})
			}

		case "notifier":
			if notifier, exists := pp.manager.GetNotifier(step.PluginName); exists {
				message := Message{
					ID:      fmt.Sprintf("pipeline-%d", time.Now().Unix()),
					Subject: "Pipeline Step Completed",
					Body:    fmt.Sprintf("Completado paso: %s", step.Name),
					Type:    EMAIL,
				}
				notifier.Send("admin@example.com", message)
			}
		}

		fmt.Println()
	}

	fmt.Printf("🎉 Pipeline completado exitosamente!\n")
	fmt.Printf("📤 Datos finales: %d bytes\n\n", len(currentData))

	return nil
}

// ==============================================
// SISTEMA DE DEMOSTRACIÓN
// ==============================================

func main() {
	fmt.Println("🔌 SISTEMA DE PLUGINS MODULARES")
	fmt.Println("===============================\n")

	// Crear manager y registrar plugins
	manager := NewPluginManager()

	// Registrar procesadores
	fmt.Println("📦 Registrando plugins...")
	plugins := []PluginInfo{
		NewJSONProcessor(),
		NewXMLProcessor(),
		NewConsoleLogger(),
		NewFileLogger(),
		NewJWTAuthenticator(),
		NewEmailNotifier(),
		NewSlackNotifier(),
	}

	for _, plugin := range plugins {
		if err := manager.RegisterPlugin(plugin); err != nil {
			log.Printf("Error registrando plugin %s: %v", plugin.Name(), err)
		} else {
			fmt.Printf("✅ %s v%s registrado\n", plugin.Name(), plugin.Version())
		}
	}

	fmt.Printf("\n📋 Total de plugins registrados: %d\n", len(manager.ListPlugins()))
	fmt.Println()

	// Demostración 1: Procesamiento de datos
	fmt.Println("🔄 DEMO 1: Procesamiento de Datos")
	fmt.Println("=================================")

	jsonData := `{"name":"John","age":30,"active":true}`

	if processor, exists := manager.GetProcessor("JSONProcessor"); exists {
		fmt.Printf("📊 Procesando JSON con %s...\n", processor.Name())

		if err := processor.Validate([]byte(jsonData), "json"); err != nil {
			fmt.Printf("❌ Validación falló: %v\n", err)
		} else {
			fmt.Println("✅ Validación exitosa")

			processed, err := processor.Process([]byte(jsonData), "json")
			if err != nil {
				fmt.Printf("❌ Procesamiento falló: %v\n", err)
			} else {
				fmt.Printf("✅ Datos procesados:\n%s\n", string(processed))
			}
		}
	}
	fmt.Println()

	// Demostración 2: Sistema de logging
	fmt.Println("📝 DEMO 2: Sistema de Logging")
	fmt.Println("=============================")

	if logger, exists := manager.GetLogger("ConsoleLogger"); exists {
		logger.Log(INFO, "Sistema iniciado correctamente", map[string]interface{}{
			"plugins": len(manager.ListPlugins()),
			"tiempo":  time.Now(),
		})

		logger.Log(WARN, "Plugin no encontrado", map[string]interface{}{
			"plugin": "NonExistentPlugin",
		})

		logger.Log(ERROR, "Error simulado de conexión", map[string]interface{}{
			"error": "connection timeout",
			"retry": 3,
		})
	}
	fmt.Println()

	// Demostración 3: Autenticación
	fmt.Println("🔐 DEMO 3: Sistema de Autenticación")
	fmt.Println("===================================")

	if auth, exists := manager.GetAuthenticator("JWTAuthenticator"); exists {
		// Autenticación exitosa
		credentials := map[string]string{
			"username": "admin",
			"password": "secret",
		}

		user, err := auth.Authenticate(credentials)
		if err != nil {
			fmt.Printf("❌ Autenticación falló: %v\n", err)
		} else {
			fmt.Printf("✅ Usuario autenticado: %s (%s)\n", user.Username, user.Email)
			fmt.Printf("   Roles: %v\n", user.Roles)

			// Validar token
			token := "valid-jwt-token"
			if validUser, err := auth.ValidateToken(token); err == nil {
				fmt.Printf("✅ Token válido para: %s\n", validUser.Username)
			}
		}
	}
	fmt.Println()

	// Demostración 4: Notificaciones
	fmt.Println("📢 DEMO 4: Sistema de Notificaciones")
	fmt.Println("====================================")

	message := Message{
		ID:       "msg-123",
		Subject:  "Sistema Iniciado",
		Body:     "El sistema de plugins ha sido iniciado correctamente",
		Type:     EMAIL,
		Priority: 1,
	}

	// Enviar por email
	if emailNotifier, exists := manager.GetNotifier("EmailNotifier"); exists {
		emailNotifier.Send("admin@example.com", message)
	}

	// Enviar por Slack
	if slackNotifier, exists := manager.GetNotifier("SlackNotifier"); exists {
		slackMessage := message
		slackMessage.Type = SLACK
		slackNotifier.Send("#general", slackMessage)
	}
	fmt.Println()

	// Demostración 5: Pipeline de procesamiento
	fmt.Println("⚡ DEMO 5: Pipeline de Procesamiento")
	fmt.Println("===================================")

	pipeline := NewProcessingPipeline(manager)
	pipeline.AddStep("Procesar JSON", "JSONProcessor", "processor", nil)
	pipeline.AddStep("Log resultado", "ConsoleLogger", "logger", nil)
	pipeline.AddStep("Notificar admin", "EmailNotifier", "notifier", nil)

	complexJSON := `{"users":[{"id":1,"name":"Ana"},{"id":2,"name":"Carlos"}],"total":2}`
	pipeline.Execute([]byte(complexJSON), "json")

	// Resumen final
	fmt.Println("📊 RESUMEN DEL SISTEMA")
	fmt.Println("=====================")
	fmt.Printf("🔌 Total plugins: %d\n", len(manager.ListPlugins()))
	fmt.Printf("📊 Procesadores: %d\n", len(manager.processors))
	fmt.Printf("📝 Loggers: %d\n", len(manager.loggers))
	fmt.Printf("🔐 Autenticadores: %d\n", len(manager.authenticators))
	fmt.Printf("📢 Notificadores: %d\n", len(manager.notifiers))

	fmt.Println("\n✅ Sistema de plugins funcionando correctamente!")
	fmt.Println("🎉 Demostración completada exitosamente!")
}
