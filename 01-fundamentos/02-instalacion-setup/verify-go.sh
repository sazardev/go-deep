#!/bin/bash

# 🔧 Script de Verificación de Instalación de Go
# ==============================================
# Descripción: Verifica que Go esté instalado y configurado correctamente
# Compatibilidad: Linux, macOS, Windows (WSL/Git Bash)
# Uso: ./verify-go.sh

set -e

# Colors para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Función para log con colores
log() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

success() {
    echo -e "${GREEN}[✓]${NC} $1"
}

warning() {
    echo -e "${YELLOW}[!]${NC} $1"
}

error() {
    echo -e "${RED}[✗]${NC} $1"
}

header() {
    echo -e "\n${PURPLE}$1${NC}"
    echo -e "${PURPLE}$(echo "$1" | sed 's/./=/g')${NC}\n"
}

# Contadores
TESTS_PASSED=0
TESTS_FAILED=0
TOTAL_TESTS=0

# Función para ejecutar test
run_test() {
    local test_name="$1"
    local test_command="$2"
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    
    printf "%-50s " "$test_name..."
    
    if eval "$test_command" >/dev/null 2>&1; then
        echo -e "${GREEN}✓ PASS${NC}"
        TESTS_PASSED=$((TESTS_PASSED + 1))
        return 0
    else
        echo -e "${RED}✗ FAIL${NC}"
        TESTS_FAILED=$((TESTS_FAILED + 1))
        return 1
    fi
}

# Función para mostrar información detallada
show_detailed_info() {
    local test_name="$1"
    local test_command="$2"
    
    echo -e "${CYAN}$test_name:${NC}"
    if eval "$test_command" 2>/dev/null; then
        echo ""
    else
        error "No disponible o error"
        echo ""
    fi
}

# Header del script
header "🔧 VERIFICADOR DE INSTALACIÓN DE GO"

# 1. Verificar si Go está instalado
header "1. Verificación de Instalación Básica"

run_test "Go está instalado" "command -v go"
run_test "Go es ejecutable" "go version"

if command -v go >/dev/null 2>&1; then
    show_detailed_info "Versión de Go" "go version"
    show_detailed_info "Ubicación del binario" "which go"
else
    error "Go no está instalado o no está en PATH"
    echo ""
    log "Para instalar Go, visita: https://golang.org/dl/"
    log "O usa nuestros scripts de instalación:"
    log "  - Linux/macOS: ./install-go.sh"
    log "  - Windows: install-go.bat o install-go.ps1"
    exit 1
fi

# 2. Verificar variables de entorno
header "2. Verificación de Variables de Entorno"

run_test "GOROOT está configurado" "[ -n \"\$GOROOT\" ] || go env GOROOT"
run_test "GOPATH está configurado" "[ -n \"\$GOPATH\" ] || go env GOPATH"
run_test "GOBIN está configurado" "[ -n \"\$GOBIN\" ] || go env GOBIN"

show_detailed_info "GOROOT" "go env GOROOT"
show_detailed_info "GOPATH" "go env GOPATH"
show_detailed_info "GOBIN" "go env GOBIN"
show_detailed_info "GOOS" "go env GOOS"
show_detailed_info "GOARCH" "go env GOARCH"

# 3. Verificar estructura de directorios
header "3. Verificación de Estructura de Directorios"

GOPATH=$(go env GOPATH)
run_test "Directorio GOPATH existe" "[ -d \"$GOPATH\" ]"
run_test "Directorio GOPATH/bin existe" "[ -d \"$GOPATH/bin\" ]"
run_test "Directorio GOPATH/src existe" "[ -d \"$GOPATH/src\" ]"
run_test "Directorio GOPATH/pkg existe" "[ -d \"$GOPATH/pkg\" ]"

if [ -d "$GOPATH" ]; then
    show_detailed_info "Contenido de GOPATH" "ls -la \"$GOPATH\""
else
    warning "GOPATH no existe. Creando estructura..."
    mkdir -p "$GOPATH/bin" "$GOPATH/src" "$GOPATH/pkg"
    success "Estructura de GOPATH creada"
fi

# 4. Verificar PATH
header "4. Verificación de PATH"

GOROOT=$(go env GOROOT)
GOBIN=$(go env GOBIN)

run_test "GOROOT/bin está en PATH" "echo \$PATH | grep -q \"$GOROOT/bin\""
run_test "GOBIN está en PATH" "echo \$PATH | grep -q \"$GOBIN\""

echo -e "${CYAN}PATH actual:${NC}"
echo "$PATH" | tr ':' '\n' | grep -E "(go|Go)" || warning "No se encontraron rutas de Go en PATH"
echo ""

# 5. Verificar herramientas básicas
header "5. Verificación de Herramientas Básicas"

run_test "go build funciona" "echo 'package main; func main() {}' | go build -o /tmp/test-go -"
run_test "go run funciona" "echo 'package main; import \"fmt\"; func main() { fmt.Println(\"test\") }' | go run -"
run_test "go fmt funciona" "echo 'package main; func main(){fmt.Println(\"test\")}' | go fmt"

# Limpiar archivo de test
rm -f /tmp/test-go

# 6. Verificar herramientas adicionales
header "6. Verificación de Herramientas Adicionales"

tools=(
    "goimports:golang.org/x/tools/cmd/goimports"
    "staticcheck:honnef.co/go/tools/cmd/staticcheck"
    "golangci-lint:github.com/golangci/golangci-lint/cmd/golangci-lint"
    "air:github.com/air-verse/air"
    "gopls:golang.org/x/tools/gopls"
)

for tool_info in "${tools[@]}"; do
    tool_name=$(echo "$tool_info" | cut -d':' -f1)
    run_test "$tool_name está instalado" "command -v \"$tool_name\""
done

# 7. Test de compilación y ejecución
header "7. Test de Compilación y Ejecución"

# Crear programa de test temporal
TEST_DIR="/tmp/go-verification-test"
rm -rf "$TEST_DIR"
mkdir -p "$TEST_DIR"

cat > "$TEST_DIR/main.go" << 'EOF'
package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    fmt.Printf("🚀 Go Verification Test\n")
    fmt.Printf("======================\n")
    fmt.Printf("Go Version: %s\n", runtime.Version())
    fmt.Printf("OS: %s\n", runtime.GOOS)
    fmt.Printf("Architecture: %s\n", runtime.GOARCH)
    fmt.Printf("CPUs: %d\n", runtime.NumCPU())
    fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
    
    // Test goroutine
    done := make(chan bool)
    go func() {
        time.Sleep(100 * time.Millisecond)
        fmt.Println("✓ Goroutine funcionando")
        done <- true
    }()
    
    <-done
    fmt.Println("✓ Test completado exitosamente")
}
EOF

run_test "Compilación de programa de test" "cd \"$TEST_DIR\" && go build -o test-program main.go"
run_test "Ejecución de programa de test" "cd \"$TEST_DIR\" && ./test-program"

# Limpiar
rm -rf "$TEST_DIR"

# 8. Verificar módulos de Go
header "8. Verificación de Módulos de Go"

run_test "go mod está disponible" "go help mod"
run_test "go mod init funciona" "cd /tmp && mkdir -p go-mod-test && cd go-mod-test && go mod init test && rm -rf /tmp/go-mod-test"

# 9. Test de descarga de dependencias
header "9. Test de Descarga de Dependencias"

TEST_MOD_DIR="/tmp/go-mod-verification"
rm -rf "$TEST_MOD_DIR"
mkdir -p "$TEST_MOD_DIR"

cat > "$TEST_MOD_DIR/main.go" << 'EOF'
package main

import (
    "fmt"
    "github.com/fatih/color"
)

func main() {
    c := color.New(color.FgGreen).Add(color.Underline)
    c.Println("Dependencias funcionando correctamente!")
}
EOF

cd "$TEST_MOD_DIR"
run_test "go mod init funciona" "go mod init test-module"
run_test "go mod tidy funciona" "go mod tidy"
run_test "Descarga de dependencias" "go get github.com/fatih/color"
run_test "Compilación con dependencias" "go build -o test-deps main.go"

# Limpiar
rm -rf "$TEST_MOD_DIR"

# 10. Verificar configuración del editor
header "10. Verificación de Configuración del Editor"

# Verificar VS Code Go extension si VS Code está instalado
if command -v code >/dev/null 2>&1; then
    run_test "VS Code está instalado" "command -v code"
    
    # Verificar si la extensión de Go está instalada
    if code --list-extensions 2>/dev/null | grep -q golang.go; then
        success "Extensión de Go para VS Code está instalada"
    else
        warning "Extensión de Go para VS Code no encontrada"
        log "Instalar con: code --install-extension golang.Go"
    fi
else
    warning "VS Code no está instalado"
fi

# Verificar Vim/Neovim Go plugins
if command -v vim >/dev/null 2>&1 || command -v nvim >/dev/null 2>&1; then
    run_test "Vim/Neovim está instalado" "command -v vim || command -v nvim"
else
    log "Vim/Neovim no está instalado"
fi

# Resumen final
header "📊 RESUMEN DE VERIFICACIÓN"

echo -e "${CYAN}Estadísticas:${NC}"
echo "  Total de tests: $TOTAL_TESTS"
echo -e "  Tests pasados: ${GREEN}$TESTS_PASSED${NC}"
echo -e "  Tests fallidos: ${RED}$TESTS_FAILED${NC}"

if [ $TESTS_FAILED -eq 0 ]; then
    echo ""
    success "🎉 ¡Todas las verificaciones pasaron! Go está correctamente instalado y configurado."
    echo ""
    log "🚀 Estás listo para comenzar a programar en Go!"
    log "📚 Próximos pasos sugeridos:"
    echo "   1. Visita https://tour.golang.org/ para el tour interactivo"
    echo "   2. Lee la documentación en https://golang.org/doc/"
    echo "   3. Explora ejemplos en https://gobyexample.com/"
    echo "   4. Únete a la comunidad en https://gophers.slack.com/"
else
    echo ""
    warning "⚠️  Algunas verificaciones fallaron. Revisa los errores arriba."
    echo ""
    log "🔧 Posibles soluciones:"
    echo "   1. Reinstalar Go usando nuestros scripts de instalación"
    echo "   2. Verificar variables de entorno en tu shell profile"
    echo "   3. Reiniciar tu terminal después de la instalación"
    echo "   4. Verificar que el PATH incluya los directorios de Go"
    
    # Mostrar comandos de diagnóstico
    echo ""
    log "🔍 Comandos de diagnóstico útiles:"
    echo "   go env                    # Mostrar todas las variables de Go"
    echo "   echo \$PATH               # Verificar PATH"
    echo "   which go                  # Ubicación del binario de Go"
    echo "   go version                # Versión instalada"
fi

echo ""
header "¡Verificación completada! 🏁"

exit $TESTS_FAILED
