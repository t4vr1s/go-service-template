# go-service-template

Proyecto base en Go para servicios modulares con una estructura inspirada en arquitectura limpia.

## Estructura inicial

```text
cmd/
  service/
internal/
  domain/
  usecase/
  infrastructure/
pkg/
api/
```

## Ejecutar localmente

```bash
go run ./cmd/service
```

La aplicación expone por defecto `http://localhost:8080` y permite configurar el puerto con `PORT`.

## Endpoints de ejemplo

- `GET /` retorna información básica del servicio.
- `GET /health` retorna el estado del servicio.
- `GET /greet?name=travis` retorna un saludo de ejemplo.

## Validación

```bash
go test ./...
go build ./...
```
