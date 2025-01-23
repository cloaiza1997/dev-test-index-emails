# Indexación de Correos

Sistema que permite indexar, visualizar y buscar en una base de datos de correos.

| Column 1            | Column 2                                                        |
| ------------------- | --------------------------------------------------------------- |
| Índice              | [ZincSearch](https://zincsearch-docs.zinc.dev/getting-started/) |
| Backend             | [Go](https://go.dev/)                                           |
| Librería API Router | [Chi](https://go-chi.io/)                                       |
| Frontend            | [Vue.js](https://vuejs.org/)                                    |

## Generalidades

- El programa fue desarrollado en un equipo Windows, por lo tanto los flags de todos los comandos usados deben de tener doble guión (--).
- Para los paths usados en los comandos validar si funciona con (/) o con (\\), ya que varía según el sistema operativo.

## Visualizador

### Web

![web-01](/assets/web-01.png)
![web-02](/assets/web-02.png)
![web-03](/assets/web-03.png)

### Responsive

![mobile-01](/assets/mobile-01.png)
![mobile-02](/assets/mobile-02.png)
![mobile-03](/assets/mobile-03.png)
![mobile-04](/assets/mobile-04.png)

## Variables de entonor

Para ejecutar el proyecto de forma manual y local, se deben de crear las siguientes variables de entorno. Sus valores están en el archivo [docker-compose.yml](/docker-compose.yml).

- `EMAIL_INDEX_ZS_HOST = 8080`: Host donde se aloja ZincSearch.
- `EMAIL_INDEX_API_PORT = http://<_IP_ADDRESS_>:4080/api`: Puerto donde se expone el api.
- `ZINC_FIRST_ADMIN_USER = root`: Usuario configurado en ZincSearch.
- `ZINC_FIRST_ADMIN_PASSWORD = root`: Contraseña del usuario

## Puesta en marcha

Para facilitar el proceso se recomienda usar Linux, si es un equipo Windows, se puede usar la línea de comandos de GitBash o usar directamente el WSL (Puede ser con Ubuntu).

```shell
# Verificar si se tiene instalado
wsl --list --verbose

# En caso de no tener instalado Ubuntu
wsl --install

# Abrir la terminal
wsl -d Ubuntu
```

1. Se debe de descargar la base de datos [enron_mail_20110402](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz) y luego descomprimir (Este proceso puede tomar unos 10 minutos, debido a la cantidad de correos que hay en la base de datos).

```shell
# Abrir la carpeta en donde se guardarán los correos
cd 00-indexer/mock/email-data

# Descargar la base de datos
curl -L -o enron_mail_20110402.tgz http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz

# Descomprimir
tar -xvzf enron_mail_20110402.tgz

# Regresar a la raíz del proyecto
cd ../../../
```

2. En los siguientes archivos buscar y reemplazar `<_IP_ADDRESS_>` por la url con la ip del equipo en donde se va a ejecutar el proyecto. Esto para permitir la comunicación entre los diferentes contenedores que se crean.

   - [docker-compose.yml](/docker-compose.yml)
   - [02-ui/src/config/environment.ts](/02-ui/src/config/environment.ts): En caso de ser requerido. Por defecto usa [http://localhost:8080/](http://localhost:8080/).

```JavaScript
  // const HOST: '<_IP_ADDRESS_>:8080',
  const HOST: 'http://192.168.0.1:8080'
```

3. Desde la raíz del proyecto subir todos los servicios ejecutar el archivo [docker-compose.yml](/docker-compose.yml).

```Shell
docker-compose up -d --build
```

4. Se crearán 4 contenedores.

| Contenedor     | Acceso                                                             | Descripción                                                                                                                                                                                                                                                                                                                                                                                                 |
| -------------- | ------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **zincsearch** | [http://localhost:4080/ui/search](http://localhost:4080/ui/search) | Índice de ZincSearch. Las credenciales se pueden obtener del arhchivo [docker-compose.yml](/docker-compose.yml) de las variables `ZINC_FIRST_ADMIN_USER` y `ZINC_FIRST_ADMIN_PASSWORD`.                                                                                                                                                                                                                     |
| **indexer**    | Se accede desde la terminal del contenedor a través de docker      | Contiene el programa que indexa una base de datos de correos en formato [RFC 5322](https://datatracker.ietf.org/doc/html/rfc5322) en ZincSerach. El indexador es dinámico, se puede indexar cualquier base de datos de correos en formato RFC 5322, siempre y cuando esté dentro de la carpeta [00-indexer/mock/email-data](/00-indexer/mock/email-data), ya que esta es montada como volumen del contedor. |
| **api**        | [http://localhost:8080/](http://localhost:8080/)                   | Servicio que consulta y filtra los correos indexados.                                                                                                                                                                                                                                                                                                                                                       |
| **ui**         | [http://localhost:3000/](http://localhost:3000/)                   | Interfaz gráfica para visualizar y filtrar los correos.                                                                                                                                                                                                                                                                                                                                                     |

5. Para cargar la base de datos se debe ejecutar el programa almacenado en el contenedor del indexador. Tener en cuenta que cargar la base de datos completa puede tomar varios minutos, por lo tanto se deja una base de datos pequeña para pruebas: [00-indexer/mock/email-data/maildir](/00-indexer/mock/email-data/maildir).

```Shell
# Ejecuta la indexación con los valores por defecto y la base de datos de prueba
docker exec -it dev-test-tr-indexer ./00-indexer
```

```Shell
# Ejecuta la indexación de la base de datos completa
# Esto puede tardar entre 15 y 20 minutos según los recursos del equipo
docker exec -it dev-test-tr-indexer ./00-indexer --p=./mock/email-data/enron_mail_20110402
```

---

# 00 - Indexador

## Flags

El indexador cuenta con algunos flags que permiten definir la configuración a utilizar durante la ejecución de la indexación.

- `--i`: (String) **Index**: `Default=emails`: Nombre con el que se creará el índice en ZincSearch.
- `--p`: (String) **Path**: `Default=./mock/email-data/maildir`: Ruta del directorio de emails a indexar.
- `--r`: (Int) **Routines**: `Default=10`: Cantidad de procesos concurrentes que se utilizan al momento de procesar e indexar los emails. Dependiendo del equipo en donde se ejecute el programa, se debe de buscar un valor equilibrado para evitar que se bloquee el sistema.
- `--b`: (Bool) **IndexByBatch**: `Default=true`: Flag para determinar si el consumo del api de ZincSearch se hace por lotes.
- `--s`: (String) **BatchSize**: `Default=10000`: Cantidad de correos a procesar por lote. Es importante buscar un valor adecuado que no afecte el envío a ZincSearch, ya que un lote muy grande, implica un `body` muy grande, que podría llegar a fallar.

## Ejecución local

```Shell
cd 00-indexer

# go run main.go [flags...]
go run main.go

# Modificando los valores por defecto
go run main.go --i=emails --p=./mock/email-data/maildir --r=10 --b=true --s=10000
```

- Compilar el indexador

```Shell
go build --o indexer.exe
```

- Ejecutar el indexador compilado

```Shell
./indexer.exe --i=emails --p=./mock/email-data/maildir --r=10 --b=true --s=10000
```

## Profiling

Para visualizar gráficamente el reporte se debe de instalar la herramienta [Graphviz](https://graphviz.org/download/).

- Ejecutar test con profiling

```Shell
cd test

go test --cpuprofile=prof/cpu.out --memprofile=prof/mem.out --mutexprofile=prof/mutex.out --blockprofile=prof/block.out
```

- Visualizar el reporte interactivo web

```Shell
go tool pprof --http=:[port] [profile-file-name].out

go tool pprof --http=:8081 prof/cpu.out
go tool pprof --http=:8082 prof/mem.out
go tool pprof --http=:8083 prof/mutex.out
go tool pprof --http=:8084 prof/block.out
```

- Generar PDF

```Shell
go tool pprof prof/cpu.out # >> pdf
go tool pprof prof/mem.out # >> pdf
go tool pprof prof/mutex.out # >> pdf
go tool pprof prof/block.out # >> pdf
```

- Archivos generados de profiling

  - [block.pdf](/00-indexer/test/prof/block.pdf)
  - [cpu.pdf](/00-indexer/test/prof/cpu.pdf)
  - [mem.pdf](/00-indexer/test/prof/mem.pdf)
  - [mutex.pdf](/00-indexer/test/prof/mutex.pdf)

---

# 01 - Api

- Documentación del API para importar con Postman: [api.json](/api.json).

## Ejecución local

```Shell
cd 01-api/cmd

go run main.go
```

# 02 - Ui

Funcionalidades:

- Buscador
- Paginación
- Navegación de forma individual
- Visualización del contenido del archivo

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) to make the TypeScript language service aware of `.vue` types.

## Customize configuration

See [Vite Configuration Reference](https://vite.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```
