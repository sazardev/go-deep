# 🚀 Instalador de Go para Windows (PowerShell)
# ==============================================
# Descripción: Script de PowerShell para instalar Go en Windows
# Compatibilidad: Windows 10/11 con PowerShell 5.1+
# Uso: Ejecutar en PowerShell como Administrador
# Comando: Set-ExecutionPolicy Bypass -Scope Process -Force; .\install-go.ps1

[CmdletBinding()]
param(
    [string]$Version = "latest",
    [string]$InstallPath = "C:\Go",
    [switch]$Force = $false
)

# Verificar si se ejecuta como administrador
function Test-Administrator {
    $currentUser = [Security.Principal.WindowsIdentity]::GetCurrent()
    $principal = New-Object Security.Principal.WindowsPrincipal($currentUser)
    return $principal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
}

# Función para escribir mensajes con colores
function Write-ColorMessage {
    param(
        [string]$Message,
        [string]$Type = "Info"
    )
    
    switch ($Type) {
        "Success" { Write-Host "✅ [SUCCESS] $Message" -ForegroundColor Green }
        "Warning" { Write-Host "⚠️  [WARNING] $Message" -ForegroundColor Yellow }
        "Error"   { Write-Host "❌ [ERROR] $Message" -ForegroundColor Red }
        "Info"    { Write-Host "ℹ️  [INFO] $Message" -ForegroundColor Cyan }
        "Header"  { 
            Write-Host ""
            Write-Host $Message -ForegroundColor Magenta
            Write-Host ("=" * $Message.Length) -ForegroundColor Magenta
            Write-Host ""
        }
    }
}

# Header del script
Write-ColorMessage "🚀 INSTALADOR AUTOMÁTICO DE GO PARA WINDOWS (PowerShell)" "Header"

# Verificar permisos de administrador
if (-not (Test-Administrator)) {
    Write-ColorMessage "Este script debe ejecutarse como Administrador" "Error"
    Write-ColorMessage "Haz clic derecho en PowerShell y selecciona 'Ejecutar como administrador'" "Warning"
    Write-ColorMessage "Luego ejecuta: Set-ExecutionPolicy Bypass -Scope Process -Force" "Info"
    exit 1
}

Write-ColorMessage "Ejecutando con permisos de administrador" "Success"

# Detectar arquitectura
$arch = switch ($env:PROCESSOR_ARCHITECTURE) {
    "AMD64" { "amd64" }
    "ARM64" { "arm64" }
    "x86"   { "386" }
    default { 
        Write-ColorMessage "Arquitectura no soportada: $env:PROCESSOR_ARCHITECTURE" "Error"
        exit 1
    }
}

Write-ColorMessage "Arquitectura detectada: $arch" "Info"

# Verificar si Go ya está instalado
$goExisting = Get-Command go -ErrorAction SilentlyContinue
if ($goExisting -and -not $Force) {
    $currentVersion = (go version).Split()[2]
    Write-ColorMessage "Go ya está instalado (versión $currentVersion)" "Warning"
    
    $response = Read-Host "¿Deseas continuar con la instalación? (y/N)"
    if ($response -ne "y" -and $response -ne "Y") {
        Write-ColorMessage "Instalación cancelada por el usuario" "Info"
        exit 0
    }
}

# Obtener la última versión de Go
Write-ColorMessage "Obteniendo información de la última versión..." "Info"

try {
    if ($Version -eq "latest") {
        $latestVersion = (Invoke-WebRequest -Uri "https://go.dev/VERSION?m=text" -UseBasicParsing).Content.Trim()
        $goVersion = $latestVersion
    } else {
        $goVersion = $Version
        if (-not $goVersion.StartsWith("go")) {
            $goVersion = "go$goVersion"
        }
    }
} catch {
    Write-ColorMessage "No se pudo obtener la última versión, usando 1.24.5" "Warning"
    $goVersion = "go1.24.5"
}

Write-ColorMessage "Versión a instalar: $goVersion" "Success"

# Configurar URLs y paths
$downloadUrl = "https://go.dev/dl/$goVersion.windows-$arch.zip"
$tempDir = "$env:TEMP\go-installer"
$zipFile = "$tempDir\$goVersion.windows-$arch.zip"

# Crear directorio temporal
if (Test-Path $tempDir) {
    Remove-Item $tempDir -Recurse -Force
}
New-Item -ItemType Directory -Path $tempDir -Force | Out-Null
Write-ColorMessage "Directorio temporal creado: $tempDir" "Info"

# Descargar Go
Write-ColorMessage "Descargando Go desde: $downloadUrl" "Info"

try {
    # Usar WebClient para mostrar progreso
    $webClient = New-Object System.Net.WebClient
    
    # Registrar evento de progreso
    Register-ObjectEvent -InputObject $webClient -EventName DownloadProgressChanged -Action {
        $progress = [math]::Round($Event.SourceEventArgs.ProgressPercentage)
        Write-Progress -Activity "Descargando Go" -Status "$progress% completado" -PercentComplete $progress
    } | Out-Null
    
    $webClient.DownloadFile($downloadUrl, $zipFile)
    $webClient.Dispose()
    
    Write-Progress -Activity "Descargando Go" -Completed
    Write-ColorMessage "Descarga completada: $zipFile" "Success"
} catch {
    Write-ColorMessage "Error al descargar: $($_.Exception.Message)" "Error"
    exit 1
}

# Instalar Go
Write-ColorMessage "Instalando Go..." "Info"

# Remover instalación anterior si existe
if (Test-Path $InstallPath) {
    Write-ColorMessage "Removiendo instalación anterior de Go..." "Info"
    Remove-Item $InstallPath -Recurse -Force
}

# Extraer archivo
Write-ColorMessage "Extrayendo archivo..." "Info"
try {
    Expand-Archive -Path $zipFile -DestinationPath (Split-Path $InstallPath -Parent) -Force
    Write-ColorMessage "Go instalado correctamente en $InstallPath" "Success"
} catch {
    Write-ColorMessage "Error al extraer: $($_.Exception.Message)" "Error"
    exit 1
}

# Configurar variables de entorno
Write-ColorMessage "Configurando variables de entorno del sistema..." "Info"

$goPath = "$env:USERPROFILE\go"
$goBin = "$goPath\bin"

# Configurar variables de entorno del sistema
[Environment]::SetEnvironmentVariable("GOPATH", $goPath, "Machine")
[Environment]::SetEnvironmentVariable("GOBIN", $goBin, "Machine")

# Obtener PATH actual y agregar Go
$currentPath = [Environment]::GetEnvironmentVariable("PATH", "Machine")
$newPaths = @("$InstallPath\bin", $goBin)

foreach ($newPath in $newPaths) {
    if ($currentPath -notlike "*$newPath*") {
        $currentPath = "$currentPath;$newPath"
    }
}

[Environment]::SetEnvironmentVariable("PATH", $currentPath, "Machine")

Write-ColorMessage "Variables de entorno configuradas" "Success"

# Crear directorios de workspace
$directories = @("$goPath", "$goPath\bin", "$goPath\src", "$goPath\pkg")
foreach ($dir in $directories) {
    if (-not (Test-Path $dir)) {
        New-Item -ItemType Directory -Path $dir -Force | Out-Null
    }
}

Write-ColorMessage "Directorios de workspace creados en $goPath" "Success"

# Actualizar PATH para la sesión actual
$env:PATH = "$env:PATH;$InstallPath\bin;$goBin"

# Verificar instalación
Write-ColorMessage "Verificando instalación..." "Info"

try {
    $goVersionOutput = & "$InstallPath\bin\go.exe" version
    Write-ColorMessage "¡Instalación exitosa!" "Success"
    Write-ColorMessage "$goVersionOutput" "Success"
} catch {
    Write-ColorMessage "La instalación falló. Go no se encuentra accesible" "Error"
    exit 1
}

# Instalar herramientas adicionales
Write-ColorMessage "Instalando herramientas adicionales..." "Info"

$tools = @(
    "golang.org/x/tools/cmd/goimports@latest",
    "honnef.co/go/tools/cmd/staticcheck@latest",
    "github.com/golangci/golangci-lint/cmd/golangci-lint@latest",
    "github.com/air-verse/air@latest"
)

foreach ($tool in $tools) {
    $toolName = ($tool -split '/')[-1] -replace '@.*', ''
    Write-ColorMessage "Instalando $toolName..." "Info"
    
    try {
        & "$InstallPath\bin\go.exe" install $tool 2>$null
        Write-ColorMessage "$toolName instalado" "Success"
    } catch {
        Write-ColorMessage "Fallo al instalar $toolName" "Warning"
    }
}

# Limpiar archivos temporales
Write-ColorMessage "Limpiando archivos temporales..." "Info"
if (Test-Path $tempDir) {
    Remove-Item $tempDir -Recurse -Force
}

# Información post-instalación
Write-ColorMessage "🎉 ¡INSTALACIÓN COMPLETADA!" "Header"

Write-ColorMessage "Go ha sido instalado exitosamente" "Success"
$finalVersion = & "$InstallPath\bin\go.exe" version
Write-ColorMessage "Versión: $finalVersion" "Success"
Write-Host ""

Write-ColorMessage "📍 Ubicación: $InstallPath" "Info"
Write-ColorMessage "🏠 GOPATH: $goPath" "Info"
Write-ColorMessage "🔧 Variables de entorno: Configuradas en el sistema" "Info"
Write-Host ""

Write-ColorMessage "⚠️  IMPORTANTE: Reinicia tu PowerShell o Command Prompt" "Warning"
Write-Host ""

Write-ColorMessage "🚀 Para verificar la instalación en una nueva terminal:" "Info"
Write-Host "   go version"
Write-Host "   go env"
Write-Host ""

Write-ColorMessage "📚 Próximos pasos sugeridos:" "Info"
Write-Host "   1. Abrir nueva terminal/PowerShell"
Write-Host "   2. Ejecutar: go version"
Write-Host "   3. Crear tu primer programa Go"
Write-Host "   4. Explorar: https://tour.golang.org/"
Write-Host ""

Write-ColorMessage "🛠️  Herramientas instaladas:" "Info"
Write-Host "   • goimports (formateo de imports)"
Write-Host "   • staticcheck (análisis estático)"
Write-Host "   • golangci-lint (linter)"
Write-Host "   • air (hot reload)"
Write-Host ""

Write-ColorMessage "¡Felicitaciones! Estás listo para programar en Go 🚀" "Success"
Write-Host ""
Write-ColorMessage "¡Gracias por usar el instalador automático de Go! 🎊" "Header"

# Preguntar si desea abrir documentación
$openDocs = Read-Host "¿Deseas abrir la documentación oficial de Go? (y/N)"
if ($openDocs -eq "y" -or $openDocs -eq "Y") {
    Start-Process "https://golang.org/doc/"
}
