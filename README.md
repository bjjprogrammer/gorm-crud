# Proyecto GORM CRUD

Este proyecto es una aplicación web en Go que implementa operaciones CRUD (Crear, Leer, Actualizar, Eliminar) para usuarios en una base de datos MySQL utilizando GORM, un ORM (Object-Relational Mapping) para Go.

## Estructura del Proyecto

El proyecto está organizado en varias carpetas:

- `database/`: Contiene `connection.go` que gestiona la conexión a la base de datos.
- `handlers/`: Contiene `user.go` que define los manejadores de las rutas para las operaciones CRUD de usuarios.
- `models/`: Contiene `user.go` que define el modelo de datos para un usuario.

## Tecnologías Utilizadas

- **Go**: Lenguaje de programación utilizado para desarrollar la aplicación.
- **GORM**: ORM utilizado para interactuar con la base de datos MySQL.
- **MySQL**: Sistema de gestión de bases de datos.
- **Gorilla Mux**: Paquete para el enrutamiento de solicitudes HTTP.

## Configuración

Para configurar el proyecto, es necesario ajustar las variables de entorno definidas en el archivo `.env` basándose en el archivo de ejemplo `.env.example`.

```plaintext
#DATABASE
DB_HOST=localhost
DB_USER=root
DB_PASS=
DB_NAME=database
```

## Ejecución

Para ejecutar el proyecto, asegúrate de tener Go instalado y configurado correctamente. Luego, ejecuta el siguiente comando en la terminal:

```sh
go run main.go
```

Esto iniciará el servidor en el puerto `8080`, y podrás acceder a las rutas definidas para realizar operaciones CRUD sobre los usuarios.

## Desarrollo

El proyecto utiliza [Air](https://github.com/cosmtrek/air) para el desarrollo en vivo, lo que permite recargar automáticamente la aplicación al detectar cambios en el código. La configuración de Air se encuentra en `.air.toml`.

## Contribuir

Para contribuir al proyecto, por favor crea un fork del repositorio, realiza tus cambios y envía un Pull Request con una descripción detallada de tus cambios.

## Licencia

Este proyecto está licenciado bajo [MIT License](LICENSE).
