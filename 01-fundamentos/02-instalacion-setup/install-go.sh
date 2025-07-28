#!/bin/bash

# 🚀 Script de Instalación Automática de Go para Linux/macOS
# ==========================================================
# Descripción: Instala la última versión de Go de forma automatizada
# Compatibilidad: Linux (Ubuntu/Debian/CentOS/Arch) y macOS
# Uso: curl -fsSL <URL> | bash

set -e  # Exit on any error

# Colors para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
NC='\033[0m' # No Color

# Función para log con colores
log() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

header() {
    echo -e "${PURPLE}$1${NC}"
}

# Header del script
echo ""
header "🚀 INSTALADOR AUTOMÁTICO DE GO"
header "==============================="
echo ""

# Detectar sistema operativo
detect_os() {
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        if command -v apt-get >/dev/null 2>&1; then
            OS="ubuntu"
        elif command -v yum >/dev/null 2>&1; then
            OS="centos"
        elif command -v pacman >/dev/null 2>&1; then
            OS="arch"
        else
            OS="linux"
        fi
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        OS="macos"
    else
        error "Sistema operativo no soportado: $OSTYPE"
        exit 1
    fi
    
    log "Sistema detectado: $OS"
}

# Detectar arquitectura
detect_arch() {
    ARCH=$(uname -m)
    case $ARCH in
        x86_64)
            ARCH="amd64"
            ;;
        arm64|aarch64)
            ARCH="arm64"
            ;;
        i386|i686)
            ARCH="386"
            ;;
        *)
            error "Arquitectura no soportada: $ARCH"
            exit 1
            ;;
    esac
    
    log "Arquitectura detectada: $ARCH"
}

# Verificar si Go ya está instalado
check_existing_go() {
    if command -v go >/dev/null 2>&1; then
        CURRENT_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
        warning "Go ya está instalado (versión $CURRENT_VERSION)"
        echo ""
        echo "¿Deseas continuar con la instalación? (y/N)"
        read -r response
        if [[ ! "$response" =~ ^[Yy]$ ]]; then
            log "Instalación cancelada por el usuario"
            exit 0
        fi
    fi
}

# Obtener la última versión de Go
get_latest_version() {
    log "Obteniendo información de la última versión..."
    
    # Try to get latest version from Go website
    if command -v curl >/dev/null 2>&1; then
        GO_VERSION=$(curl -s https://go.dev/VERSION?m=text | head -n1)
    elif command -v wget >/dev/null 2>&1; then
        GO_VERSION=$(wget -qO- https://go.dev/VERSION?m=text | head -n1)
    else
        warning "curl/wget no encontrado, usando versión por defecto"
        GO_VERSION="go1.24.5"
    fi
    
    if [[ -z "$GO_VERSION" ]]; then
        warning "No se pudo obtener la última versión, usando 1.24.5"
        GO_VERSION="go1.24.5"
    fi
    
    success "Versión a instalar: $GO_VERSION"
}

# Descargar Go
download_go() {
    local os_name
    case $OS in
        ubuntu|centos|arch|linux)
            os_name="linux"
            ;;
        macos)
            os_name="darwin"
            ;;
    esac
    
    DOWNLOAD_URL="https://go.dev/dl/${GO_VERSION}.${os_name}-${ARCH}.tar.gz"
    FILENAME="${GO_VERSION}.${os_name}-${ARCH}.tar.gz"
    
    log "Descargando Go desde: $DOWNLOAD_URL"
    
    if command -v curl >/dev/null 2>&1; then
        curl -L "$DOWNLOAD_URL" -o "/tmp/$FILENAME"
    elif command -v wget >/dev/null 2>&1; then
        wget "$DOWNLOAD_URL" -O "/tmp/$FILENAME"
    else
        error "Necesitas curl o wget para descargar Go"
        exit 1
    fi
    
    success "Descarga completada: /tmp/$FILENAME"
}

# Instalar Go
install_go() {
    log "Instalando Go..."
    
    # Remover instalación anterior si existe
    if [[ -d "/usr/local/go" ]]; then
        log "Removiendo instalación anterior de Go..."
        sudo rm -rf /usr/local/go
    fi
    
    # Extraer archivo
    log "Extrayendo archivo..."
    sudo tar -C /usr/local -xzf "/tmp/$FILENAME"
    
    # Verificar instalación
    if [[ -d "/usr/local/go" ]]; then
        success "Go instalado correctamente en /usr/local/go"
    else
        error "Fallo en la instalación"
        exit 1
    fi
    
    # Limpiar archivo descargado
    rm "/tmp/$FILENAME"
}

# Configurar variables de entorno
setup_environment() {
    log "Configurando variables de entorno..."
    
    # Detectar shell
    if [[ -n "$ZSH_VERSION" ]]; then
        SHELL_RC="$HOME/.zshrc"
    elif [[ -n "$BASH_VERSION" ]]; then
        SHELL_RC="$HOME/.bashrc"
    else
        SHELL_RC="$HOME/.profile"
    fi
    
    # Backup del archivo de configuración
    if [[ -f "$SHELL_RC" ]]; then
        cp "$SHELL_RC" "${SHELL_RC}.backup.$(date +%Y%m%d_%H%M%S)"
        log "Backup creado: ${SHELL_RC}.backup.$(date +%Y%m%d_%H%M%S)"
    fi
    
    # Remover configuraciones anteriores de Go
    if [[ -f "$SHELL_RC" ]]; then
        grep -v "# Go configuration" "$SHELL_RC" > "${SHELL_RC}.tmp" || true
        grep -v "export PATH=.*go/bin" "${SHELL_RC}.tmp" > "$SHELL_RC" || true
        rm -f "${SHELL_RC}.tmp"
    fi
    
    # Agregar nueva configuración
    cat >> "$SHELL_RC" << 'EOF'

# Go configuration
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
EOF

    success "Variables de entorno configuradas en $SHELL_RC"
    
    # Crear directorio GOPATH
    mkdir -p "$HOME/go/bin"
    mkdir -p "$HOME/go/src"
    mkdir -p "$HOME/go/pkg"
    
    success "Directorios de workspace creados en $HOME/go"
}

# Verificar instalación
verify_installation() {
    log "Verificando instalación..."
    
    # Source del archivo de configuración
    export PATH=$PATH:/usr/local/go/bin
    
    if command -v /usr/local/go/bin/go >/dev/null 2>&1; then
        GO_VERSION_INSTALLED=$(/usr/local/go/bin/go version)
        success "¡Instalación exitosa!"
        success "$GO_VERSION_INSTALLED"
    else
        error "La instalación falló. Go no se encuentra en PATH"
        exit 1
    fi
}

# Instalar herramientas adicionales
install_tools() {
    log "Instalando herramientas adicionales..."
    
    export PATH=$PATH:/usr/local/go/bin
    export GOPATH=$HOME/go
    export GOBIN=$GOPATH/bin
    
    # Lista de herramientas útiles
    tools=(
        "golang.org/x/tools/cmd/goimports@latest"
        "honnef.co/go/tools/cmd/staticcheck@latest"
        "github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
        "github.com/air-verse/air@latest"
    )
    
    for tool in "${tools[@]}"; do
        log "Instalando $(echo $tool | cut -d'/' -f3 | cut -d'@' -f1)..."
        /usr/local/go/bin/go install "$tool" 2>/dev/null || warning "Fallo al instalar $tool"
    done
    
    success "Herramientas adicionales instaladas"
}

# Mostrar información post-instalación
show_post_install_info() {
    echo ""
    header "🎉 ¡INSTALACIÓN COMPLETADA!"
    header "=========================="
    echo ""
    
    success "Go ha sido instalado exitosamente"
    success "Versión: $(/usr/local/go/bin/go version)"
    echo ""
    
    log "📍 Ubicación: /usr/local/go"
    log "🏠 GOPATH: $HOME/go"
    log "🔧 Configuración: $SHELL_RC"
    echo ""
    
    warning "⚠️  IMPORTANTE: Reinicia tu terminal o ejecuta:"
    echo "   source $SHELL_RC"
    echo ""
    
    log "🚀 Para verificar la instalación:"
    echo "   go version"
    echo "   go env"
    echo ""
    
    log "📚 Próximos pasos sugeridos:"
    echo "   1. Reiniciar terminal"
    echo "   2. Ejecutar: go version"
    echo "   3. Crear tu primer programa Go"
    echo "   4. Explorar: https://tour.golang.org/"
    echo ""
    
    success "¡Felicitaciones! Estás listo para programar en Go 🚀"
}

# Función principal
main() {
    detect_os
    detect_arch
    check_existing_go
    get_latest_version
    download_go
    install_go
    setup_environment
    verify_installation
    install_tools
    show_post_install_info
}

# Verificar permisos de sudo
if ! sudo -n true 2>/dev/null; then
    warning "Este script requiere permisos de sudo para instalar Go en /usr/local/"
    echo "Se te pedirá la contraseña de sudo..."
    echo ""
fi

# Ejecutar instalación
main

echo ""
header "¡Gracias por usar el instalador automático de Go! 🎊"
