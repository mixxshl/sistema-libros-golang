# 📚 Sistema de Gestión de Libros Electrónicos

## 👨‍💻 Integrantes

- Mishel Cumbal 

---

## 📅 Fecha

Julio 2026

---

## 🎯 Objetivo del Proyecto

Desarrollar un Sistema de Gestión de Libros Electrónicos utilizando el lenguaje de programación Go (Golang), aplicando los conceptos de programación funcional, modularización, estructuras de datos, servicios web REST, manejo de errores, serialización mediante JSON y concurrencia.

---

## 📖 Descripción del Sistema

El Sistema de Gestión de Libros Electrónicos permite administrar una biblioteca digital mediante la gestión de libros, usuarios y préstamos.

Durante la primera etapa del proyecto el sistema fue desarrollado como una aplicación de consola. Posteriormente, en la Unidad 4, se implementó una API REST utilizando la biblioteca estándar de Go (`net/http`), permitiendo acceder a la información mediante servicios web.

Actualmente el sistema almacena la información en memoria utilizando slices y estructuras propias del lenguaje Go.

---

## ⚙️ Tecnologías Utilizadas

- Go (Golang)
- net/http
- encoding/json
- sync (Mutex)
- Visual Studio Code
- Git
- GitHub

---

## 📂 Estructura del Proyecto

```
sistema-libros/

│── main.go
│── go.mod

├── handlers/
│   ├── libro_handler.go
│   ├── usuario_handler.go
│   └── prestamo_handler.go

├── models/
│   ├── libro.go
│   ├── usuario.go
│   └── prestamo.go

├── routes/
│   └── routes.go

├── services/
│   ├── libro_service.go
│   ├── usuario_service.go
│   └── prestamo_service.go

├── utils/
│   └── menu.go

└── README.md
```

---

## 🌐 Servicios Web Implementados

### Gestión de Libros

- GET /libros
- POST /libros/agregar
- PUT /libros/actualizar
- DELETE /libros/eliminar

### Gestión de Usuarios

- GET /usuarios
- POST /usuarios/agregar

### Gestión de Préstamos

- GET /prestamos
- POST /prestamos/agregar

En total se implementaron ocho servicios web para la administración de la biblioteca digital.

---

## 📦 Formato de intercambio de datos

La comunicación entre el servidor y los clientes se realiza mediante JSON (JavaScript Object Notation), utilizando el paquete estándar `encoding/json` del lenguaje Go.

---

## 🔄 Funcionalidades Implementadas

- Registro de libros.
- Consulta de libros.
- Actualización de libros.
- Eliminación de libros.
- Registro de usuarios.
- Consulta de usuarios.
- Registro de préstamos.
- Consulta de préstamos.
- API REST.
- Respuestas en formato JSON.
- Manejo de errores.
- Programación funcional.
- Arquitectura modular.
- Concurrencia mediante Mutex.

---

## 🚀 Ejecución del Proyecto

Desde la carpeta principal ejecutar:

```bash
go run .
```

El servidor iniciará en:

```
http://localhost:8080
```

---

## 🧪 Pruebas Realizadas

Se realizaron pruebas utilizando peticiones HTTP para verificar:

- Consulta de libros.
- Registro de libros.
- Registro de usuarios.
- Registro de préstamos.
- Actualización de libros.
- Eliminación de libros.
- Respuestas JSON.
- Validación de errores.

---

## 🔐 Concurrencia

El sistema implementa mecanismos básicos de concurrencia mediante `sync.Mutex`, permitiendo proteger las estructuras de datos compartidas cuando múltiples solicitudes son atendidas simultáneamente por el servidor.

---

## 📈 Posibles Mejoras Futuras

- Integración con una base de datos PostgreSQL o MySQL.
- Implementación de autenticación mediante JWT.
- Desarrollo de una interfaz web.
- Desarrollo de una aplicación móvil.
- Implementación de almacenamiento en la nube.
- Gestión de roles de usuarios.
- Reportes estadísticos.
- Integración con servicios de inteligencia artificial para recomendaciones de lectura.

---

## ✅ Conclusión

Este proyecto permitió aplicar los conocimientos adquiridos durante las cuatro unidades de la asignatura, integrando programación funcional, modularización, estructuras de datos, servicios web REST, manejo de JSON y concurrencia. El sistema constituye una base sólida para futuras mejoras y demuestra el potencial de Go para el desarrollo de aplicaciones modernas.

---

## 🔗 Repositorio

https://github.com/mixxshl/sistema-libros-golang
