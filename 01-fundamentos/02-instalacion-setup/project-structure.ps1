# 📁 Estructura de Carpetas para Proyecto Go
# ==========================================
# Script para crear la estructura de carpetas recomendada para proyectos Go

# Crear estructura básica de proyecto
function Create-GoProjectStructure {
    param(
        [Parameter(Mandatory=$true)]
        [string]$ProjectName,
        
        [string]$ProjectPath = ".",
        [switch]$WithExamples = $false
    )
    
    $fullPath = Join-Path $ProjectPath $ProjectName
    
    Write-Host "🏗️  Creando estructura de proyecto: $ProjectName" -ForegroundColor Green
    Write-Host "📍 Ubicación: $fullPath" -ForegroundColor Cyan
    
    # Crear directorio principal
    if (Test-Path $fullPath) {
        Write-Host "⚠️  El directorio ya existe: $fullPath" -ForegroundColor Yellow
        $response = Read-Host "¿Continuar? (y/N)"
        if ($response -ne "y" -and $response -ne "Y") {
            return
        }
    }
    
    New-Item -ItemType Directory -Path $fullPath -Force | Out-Null
    
    # Estructura de directorios estándar
    $directories = @(
        "cmd/$ProjectName",      # Aplicaciones principales
        "internal",              # Código privado de la aplicación
        "pkg",                   # Código que pueden usar aplicaciones externas
        "api",                   # Definiciones de API
        "web",                   # Recursos web (templates, assets)
        "configs",               # Archivos de configuración
        "init",                  # Scripts de inicialización
        "scripts",               # Scripts de build, deploy, etc.
        "build",                 # Packaging y CI
        "deployments",           # Configuraciones de deployment
        "test",                  # Datos de test adicionales
        "docs",                  # Documentación
        "tools",                 # Herramientas de soporte
        "examples",              # Ejemplos de uso
        "third_party",           # Herramientas externas
        "githooks",              # Git hooks
        "assets",                # Otros assets (imágenes, logos, etc.)
        "website"                # Datos del sitio web del proyecto
    )
    
    foreach ($dir in $directories) {
        $dirPath = Join-Path $fullPath $dir
        New-Item -ItemType Directory -Path $dirPath -Force | Out-Null
        Write-Host "📁 $dir" -ForegroundColor Blue
    }
    
    # Crear archivos básicos
    Push-Location $fullPath
    
    # go.mod
    $moduleName = $ProjectName
    if ($ProjectName -notmatch "github.com|gitlab.com|bitbucket.org") {
        $moduleName = "github.com/tu-usuario/$ProjectName"
    }
    
    & go mod init $moduleName 2>$null
    Write-Host "📄 go.mod creado" -ForegroundColor Green
    
    # main.go básico
    $mainGoContent = @"
package main

import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("🚀 Bienvenido a $ProjectName!")
    log.Println("Aplicación iniciada correctamente")
}
"@
    
    $mainGoContent | Out-File -FilePath "cmd/$ProjectName/main.go" -Encoding UTF8
    Write-Host "📄 cmd/$ProjectName/main.go creado" -ForegroundColor Green
    
    # README.md
    $readmeContent = @"
# $ProjectName

Descripción breve del proyecto.

## 🚀 Instalación

``````bash
go mod download
``````

## 🏃‍♂️ Uso

``````bash
go run cmd/$ProjectName/main.go
``````

## 🔧 Desarrollo

``````bash
# Compilar
go build -o bin/$ProjectName cmd/$ProjectName/main.go

# Tests
go test ./...

# Linting
golangci-lint run
``````

## 📁 Estructura del Proyecto

``````
$ProjectName/
├── cmd/                    # Aplicaciones principales
│   └── $ProjectName/       # Aplicación principal
├── internal/               # Código privado de la aplicación
├── pkg/                    # Bibliotecas públicas
├── api/                    # Definiciones de API
├── web/                    # Recursos web
├── configs/                # Configuraciones
├── scripts/                # Scripts de automatización
├── test/                   # Datos de test adicionales
├── docs/                   # Documentación
└── README.md
``````

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama feature (`git checkout -b feature/nueva-caracteristica`)
3. Commit tus cambios (`git commit -am 'Agregar nueva característica'`)
4. Push a la rama (`git push origin feature/nueva-caracteristica`)
5. Abre un Pull Request

## 📝 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para detalles.
"@
    
    $readmeContent | Out-File -FilePath "README.md" -Encoding UTF8
    Write-Host "📄 README.md creado" -ForegroundColor Green
    
    # .gitignore
    $gitignoreContent = @"
# Binarios compilados
*.exe
*.exe~
*.dll
*.so
*.dylib
bin/
dist/

# Test binarios, construidos con `go test -c`
*.test

# Output de coverage
*.out
coverage.html

# Directorios de dependencias
vendor/

# Archivos de configuración del IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS generated files
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db

# Logs
*.log

# Variables de entorno
.env
.env.local
.env.*.local

# Archivos temporales
tmp/
temp/
"@
    
    $gitignoreContent | Out-File -FilePath ".gitignore" -Encoding UTF8
    Write-Host "📄 .gitignore creado" -ForegroundColor Green
    
    # Makefile
    $makefileContent = @"
# Makefile para $ProjectName

# Variables
BINARY_NAME=$ProjectName
BINARY_PATH=bin/`$(BINARY_NAME)
MAIN_PATH=cmd/`$(BINARY_NAME)/main.go

# Comandos por defecto
.PHONY: all build clean test coverage help

all: clean build

## build: Compilar la aplicación
build:
	@echo "🔨 Compilando..."
	@go build -o `$(BINARY_PATH) `$(MAIN_PATH)
	@echo "✅ Compilación completada: `$(BINARY_PATH)"

## run: Ejecutar la aplicación
run:
	@echo "🚀 Ejecutando..."
	@go run `$(MAIN_PATH)

## clean: Limpiar archivos compilados
clean:
	@echo "🧹 Limpiando..."
	@go clean
	@rm -f `$(BINARY_PATH)

## test: Ejecutar tests
test:
	@echo "🧪 Ejecutando tests..."
	@go test -v ./...

## coverage: Ejecutar tests con coverage
coverage:
	@echo "📊 Generando coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "📈 Coverage generado: coverage.html"

## lint: Ejecutar linter
lint:
	@echo "🔍 Ejecutando linter..."
	@golangci-lint run

## fmt: Formatear código
fmt:
	@echo "💅 Formateando código..."
	@go fmt ./...

## mod-tidy: Limpiar módulos
mod-tidy:
	@echo "📦 Limpiando módulos..."
	@go mod tidy

## deps: Descargar dependencias
deps:
	@echo "📥 Descargando dependencias..."
	@go mod download

## install: Instalar la aplicación
install: build
	@echo "📦 Instalando..."
	@go install `$(MAIN_PATH)

## help: Mostrar ayuda
help:
	@echo "Comandos disponibles:"
	@sed -n 's/^##//p' `$(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'
"@
    
    $makefileContent | Out-File -FilePath "Makefile" -Encoding UTF8
    Write-Host "📄 Makefile creado" -ForegroundColor Green
    
    # docker-compose.yml (opcional)
    $dockerComposeContent = @"
version: '3.8'

services:
  $ProjectName:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - GO_ENV=development
    volumes:
      - .:/app
    restart: unless-stopped

  # Ejemplo de base de datos
  # postgres:
  #   image: postgres:13
  #   environment:
  #     POSTGRES_DB: ${ProjectName}_db
  #     POSTGRES_USER: user
  #     POSTGRES_PASSWORD: password
  #   ports:
  #     - "5432:5432"
  #   volumes:
  #     - postgres_data:/var/lib/postgresql/data

# volumes:
#   postgres_data:
"@
    
    $dockerComposeContent | Out-File -FilePath "docker-compose.yml" -Encoding UTF8
    Write-Host "📄 docker-compose.yml creado" -ForegroundColor Green
    
    # Dockerfile
    $dockerfileContent = @"
# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copiar go mod y sum files
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/$ProjectName/main.go

# Runtime stage
FROM alpine:latest

# Instalar ca-certificates para HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar el binario desde el builder stage
COPY --from=builder /app/main .

# Exponer puerto
EXPOSE 8080

# Comando para ejecutar
CMD ["./main"]
"@
    
    $dockerfileContent | Out-File -FilePath "Dockerfile" -Encoding UTF8
    Write-Host "📄 Dockerfile creado" -ForegroundColor Green
    
    # .env.example
    $envExampleContent = @"
# Configuración de la aplicación
APP_NAME=$ProjectName
APP_VERSION=1.0.0
APP_ENV=development
APP_PORT=8080

# Base de datos
DB_HOST=localhost
DB_PORT=5432
DB_NAME=${ProjectName}_db
DB_USER=user
DB_PASSWORD=password

# Logs
LOG_LEVEL=info
LOG_FORMAT=json

# API Keys (agregar las necesarias)
API_KEY_EXAMPLE=your-api-key-here
"@
    
    $envExampleContent | Out-File -FilePath ".env.example" -Encoding UTF8
    Write-Host "📄 .env.example creado" -ForegroundColor Green
    
    if ($WithExamples) {
        # Crear ejemplos básicos
        $exampleContent = @"
package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "¡Hola desde $ProjectName!")
    })
    
    fmt.Println("🌐 Servidor iniciado en http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
"@
        
        $exampleContent | Out-File -FilePath "examples/http-server.go" -Encoding UTF8
        Write-Host "📄 examples/http-server.go creado" -ForegroundColor Green
    }
    
    Pop-Location
    
    Write-Host ""
    Write-Host "🎉 ¡Estructura del proyecto creada exitosamente!" -ForegroundColor Green
    Write-Host ""
    Write-Host "📋 Próximos pasos:" -ForegroundColor Cyan
    Write-Host "   1. cd $ProjectName" -ForegroundColor White
    Write-Host "   2. go mod tidy" -ForegroundColor White
    Write-Host "   3. go run cmd/$ProjectName/main.go" -ForegroundColor White
    Write-Host ""
    Write-Host "🔧 Comandos útiles:" -ForegroundColor Cyan
    Write-Host "   make build    # Compilar" -ForegroundColor White
    Write-Host "   make run      # Ejecutar" -ForegroundColor White
    Write-Host "   make test     # Tests" -ForegroundColor White
    Write-Host "   make help     # Ver todos los comandos" -ForegroundColor White
}

# Función para crear estructura simple
function Create-SimpleGoProject {
    param(
        [Parameter(Mandatory=$true)]
        [string]$ProjectName
    )
    
    Write-Host "🏗️  Creando proyecto simple: $ProjectName" -ForegroundColor Green
    
    if (Test-Path $ProjectName) {
        Write-Host "⚠️  El directorio ya existe: $ProjectName" -ForegroundColor Yellow
        return
    }
    
    New-Item -ItemType Directory -Path $ProjectName | Out-Null
    Push-Location $ProjectName
    
    # Inicializar módulo
    & go mod init $ProjectName
    
    # main.go simple
    $mainContent = @"
package main

import "fmt"

func main() {
    fmt.Println("¡Hola, $ProjectName!")
}
"@
    
    $mainContent | Out-File -FilePath "main.go" -Encoding UTF8
    
    Pop-Location
    
    Write-Host "✅ Proyecto simple creado" -ForegroundColor Green
    Write-Host "▶️  cd $ProjectName && go run main.go" -ForegroundColor Cyan
}

# Ejemplo de uso
if ($args.Count -eq 0) {
    Write-Host "📁 Generador de Estructura de Proyectos Go" -ForegroundColor Magenta
    Write-Host "=========================================" -ForegroundColor Magenta
    Write-Host ""
    Write-Host "Uso:" -ForegroundColor Cyan
    Write-Host "  . .\project-structure.ps1" -ForegroundColor White
    Write-Host "  Create-GoProjectStructure -ProjectName 'mi-proyecto'" -ForegroundColor White
    Write-Host "  Create-GoProjectStructure -ProjectName 'mi-proyecto' -WithExamples" -ForegroundColor White
    Write-Host "  Create-SimpleGoProject -ProjectName 'proyecto-simple'" -ForegroundColor White
    Write-Host ""
    Write-Host "Parámetros:" -ForegroundColor Cyan
    Write-Host "  -ProjectName    : Nombre del proyecto (requerido)" -ForegroundColor White
    Write-Host "  -ProjectPath    : Ruta donde crear el proyecto (opcional)" -ForegroundColor White
    Write-Host "  -WithExamples   : Incluir archivos de ejemplo (opcional)" -ForegroundColor White
}

# Exportar funciones
Export-ModuleMember -Function Create-GoProjectStructure, Create-SimpleGoProject
