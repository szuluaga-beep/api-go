# Pruebas de Carga y Testing - API REST Go con Concurrencia

Este documento describe las pruebas implementadas para validar las mejoras de concurrencia y realizar pruebas de carga en la API REST.

## 📋 Tipos de Pruebas Implementadas

### 1. Pruebas Unitarias (`*_test.go`)
- Tests de endpoints individuales
- Validación de respuestas JSON
- Tests de concurrencia y thread-safety

### 2. Pruebas de Carga
- Scripts para simular múltiples usuarios concurrentes
- Medición de rendimiento con goroutines
- Comparación antes/después de la implementación de concurrencia

### 3. Benchmarks de Performance
- Medición de throughput (requests/segundo)
- Latencia promedio y percentiles
- Uso de memoria y goroutines

## 🚀 Configuración de Pruebas

### Instalar Dependencias de Testing

```bash
# Dependencias para testing
go get github.com/stretchr/testify/assert
go get github.com/stretchr/testify/suite

# Herramientas de load testing
go install github.com/rakyll/hey@latest
go install github.com/tsenart/vegeta@latest

# Para stress testing
go install golang.org/x/tools/cmd/stress@latest