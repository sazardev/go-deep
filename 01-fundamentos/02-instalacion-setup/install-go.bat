@echo off
:: 🚀 Script de Instalación Automática de Go para Windows
:: ====================================================
:: Descripción: Instala la última versión de Go en Windows
:: Compatibilidad: Windows 10/11 (PowerShell 5.1+)
:: Uso: Ejecutar como Administrador

setlocal enabledelayedexpansion

:: Colores para output (usando echo con códigos ANSI)
set "RED=31m"
set "GREEN=32m"
set "YELLOW=33m"
set "BLUE=34m"
set "PURPLE=35m"
set "NC=0m"

:: Variables globales
set "INSTALL_DIR=C:\Go"
set "TEMP_DIR=%TEMP%\go-installer"
set "GO_VERSION="

:: Función para mostrar mensajes con colores
:log
echo [INFO] %~1
goto :eof

:success
echo [SUCCESS] %~1
goto :eof

:warning
echo [WARNING] %~1
goto :eof

:error
echo [ERROR] %~1
goto :eof

:header
echo.
echo %~1
echo ===============================
echo.
goto :eof

:: Header del script
call :header "🚀 INSTALADOR AUTOMATICO DE GO PARA WINDOWS"

:: Verificar permisos de administrador
:check_admin
net session >nul 2>&1
if %errorLevel% == 0 (
    call :success "Ejecutando con permisos de administrador"
) else (
    call :error "Este script debe ejecutarse como Administrador"
    echo Haz clic derecho en el archivo y selecciona "Ejecutar como administrador"
    pause
    exit /b 1
)

:: Detectar arquitectura
:detect_arch
if "%PROCESSOR_ARCHITECTURE%"=="AMD64" (
    set "ARCH=amd64"
) else if "%PROCESSOR_ARCHITECTURE%"=="ARM64" (
    set "ARCH=arm64"
) else if "%PROCESSOR_ARCHITECTURE%"=="x86" (
    set "ARCH=386"
) else (
    call :error "Arquitectura no soportada: %PROCESSOR_ARCHITECTURE%"
    pause
    exit /b 1
)

call :log "Arquitectura detectada: %ARCH%"

:: Verificar si Go ya está instalado
:check_existing_go
where go >nul 2>&1
if %errorLevel% == 0 (
    for /f "tokens=3" %%i in ('go version 2^>nul') do set "CURRENT_VERSION=%%i"
    call :warning "Go ya está instalado (versión !CURRENT_VERSION!)"
    echo.
    set /p "response=¿Deseas continuar con la instalación? (y/N): "
    if /i not "!response!"=="y" (
        call :log "Instalación cancelada por el usuario"
        pause
        exit /b 0
    )
)

:: Crear directorio temporal
:create_temp_dir
if exist "%TEMP_DIR%" rd /s /q "%TEMP_DIR%"
mkdir "%TEMP_DIR%"
call :log "Directorio temporal creado: %TEMP_DIR%"

:: Obtener la última versión de Go
:get_latest_version
call :log "Obteniendo información de la última versión..."

:: Usar PowerShell para obtener la última versión
powershell -Command "try { $version = (Invoke-WebRequest -Uri 'https://go.dev/VERSION?m=text' -UseBasicParsing).Content.Trim(); Write-Output $version } catch { Write-Output 'go1.24.5' }" > "%TEMP_DIR%\version.txt"

set /p GO_VERSION=<"%TEMP_DIR%\version.txt"
if "!GO_VERSION!"=="" set "GO_VERSION=go1.24.5"

call :success "Versión a instalar: !GO_VERSION!"

:: Descargar Go
:download_go
set "DOWNLOAD_URL=https://go.dev/dl/!GO_VERSION!.windows-%ARCH%.zip"
set "FILENAME=!GO_VERSION!.windows-%ARCH%.zip"
set "FILEPATH=%TEMP_DIR%\!FILENAME!"

call :log "Descargando Go desde: !DOWNLOAD_URL!"

:: Usar PowerShell para descargar
powershell -Command "try { Invoke-WebRequest -Uri '!DOWNLOAD_URL!' -OutFile '!FILEPATH!' -UseBasicParsing; Write-Output 'SUCCESS' } catch { Write-Output 'ERROR' }" > "%TEMP_DIR%\download_result.txt"

set /p DOWNLOAD_RESULT=<"%TEMP_DIR%\download_result.txt"
if "!DOWNLOAD_RESULT!"=="ERROR" (
    call :error "Fallo en la descarga"
    pause
    exit /b 1
)

call :success "Descarga completada: !FILEPATH!"

:: Instalar Go
:install_go
call :log "Instalando Go..."

:: Remover instalación anterior si existe
if exist "%INSTALL_DIR%" (
    call :log "Removiendo instalación anterior de Go..."
    rd /s /q "%INSTALL_DIR%"
)

:: Extraer archivo usando PowerShell
call :log "Extrayendo archivo..."
powershell -Command "try { Expand-Archive -Path '!FILEPATH!' -DestinationPath 'C:\' -Force; Write-Output 'SUCCESS' } catch { Write-Output 'ERROR' }" > "%TEMP_DIR%\extract_result.txt"

set /p EXTRACT_RESULT=<"%TEMP_DIR%\extract_result.txt"
if "!EXTRACT_RESULT!"=="ERROR" (
    call :error "Fallo en la extracción"
    pause
    exit /b 1
)

:: Verificar instalación
if exist "%INSTALL_DIR%" (
    call :success "Go instalado correctamente en %INSTALL_DIR%"
) else (
    call :error "Fallo en la instalación"
    pause
    exit /b 1
)

:: Configurar variables de entorno
:setup_environment
call :log "Configurando variables de entorno del sistema..."

:: Agregar Go al PATH del sistema
setx /M PATH "%PATH%;%INSTALL_DIR%\bin" >nul 2>&1

:: Configurar GOPATH
set "GOPATH=%USERPROFILE%\go"
setx /M GOPATH "%GOPATH%" >nul 2>&1

:: Configurar GOBIN
set "GOBIN=%GOPATH%\bin"
setx /M GOBIN "%GOBIN%" >nul 2>&1

:: Agregar GOBIN al PATH
setx /M PATH "%PATH%;%GOBIN%" >nul 2>&1

call :success "Variables de entorno configuradas"

:: Crear directorios de workspace
if not exist "%GOPATH%" mkdir "%GOPATH%"
if not exist "%GOPATH%\bin" mkdir "%GOPATH%\bin"
if not exist "%GOPATH%\src" mkdir "%GOPATH%\src"
if not exist "%GOPATH%\pkg" mkdir "%GOPATH%\pkg"

call :success "Directorios de workspace creados en %GOPATH%"

:: Actualizar PATH para la sesión actual
set "PATH=%PATH%;%INSTALL_DIR%\bin;%GOBIN%"

:: Verificar instalación
:verify_installation
call :log "Verificando instalación..."

"%INSTALL_DIR%\bin\go.exe" version >nul 2>&1
if %errorLevel% == 0 (
    for /f "tokens=*" %%i in ('"%INSTALL_DIR%\bin\go.exe" version') do set "GO_VERSION_INSTALLED=%%i"
    call :success "¡Instalación exitosa!"
    call :success "!GO_VERSION_INSTALLED!"
) else (
    call :error "La instalación falló. Go no se encuentra accesible"
    pause
    exit /b 1
)

:: Limpiar archivos temporales
:cleanup
call :log "Limpiando archivos temporales..."
if exist "%TEMP_DIR%" rd /s /q "%TEMP_DIR%"

:: Mostrar información post-instalación
:show_post_install_info
echo.
call :header "🎉 ¡INSTALACIÓN COMPLETADA!"
echo.

call :success "Go ha sido instalado exitosamente"
for /f "tokens=*" %%i in ('"%INSTALL_DIR%\bin\go.exe" version') do call :success "Versión: %%i"
echo.

call :log "📍 Ubicación: %INSTALL_DIR%"
call :log "🏠 GOPATH: %GOPATH%"
call :log "🔧 Variables de entorno: Configuradas en el sistema"
echo.

call :warning "⚠️  IMPORTANTE: Reinicia tu terminal o Command Prompt"
echo.

call :log "🚀 Para verificar la instalación en una nueva terminal:"
echo    go version
echo    go env
echo.

call :log "📚 Próximos pasos sugeridos:"
echo    1. Abrir nueva terminal/Command Prompt
echo    2. Ejecutar: go version
echo    3. Crear tu primer programa Go
echo    4. Explorar: https://tour.golang.org/
echo.

call :success "¡Felicitaciones! Estás listo para programar en Go 🚀"
echo.

call :header "¡Gracias por usar el instalador automático de Go! 🎊"

pause
exit /b 0
