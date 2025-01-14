# Indexador

Indexa una base de datos de correos en formato [RFC 5322](https://datatracker.ietf.org/doc/html/rfc5322) en un índice de [ZincSearch](https://zincsearch-docs.zinc.dev/getting-started/).

- Base de datos: [enron_mail_20110402](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz).

## Generalidades

- El programa fue desarrollado en un equipo Windows, por lo tanto los flags de todos los comandos usados deben de tener doble guión (--).
- Para los paths usados en los comandos validar si funciona con (/) o con (\\), ya que varía según el sistema operativo.

## ZincSearch

TODO... Configuración, Valores por defecto

- Agregar las variables de entorno:

  - `EMAIL_INDEX_ZS_HOST`: Host donde se aloja ZincSearch
  - `EMAIL_INDEX_ZS_USER`: Usuario configurado en ZincSearch
  - `EMAIL_INDEX_ZS_PASS`: Contraseña del usuario

## Flags

El indexador cuenta con algunos flags que permiten definir la configuración a utilizar durante la ejecución de la indexación.

- `--i`: (String) **Index**: `Default=emails`: Nombre con el que se creará el índice en ZincSearch.
- `--p`: (String) **Path**: `Default=./mock/maildir`: Ruta del directorio de emails a indexar.
- `--r`: (Int) **Routines**: `Default=10`: Cantidad de procesos concurrentes que se utilizan al momento de procesar e indexar los emails. Dependiendo del equipo en donde se ejecute el programa, se debe de buscar un valor equilibrado para evitar que se bloquee el sistema.
- `--b`: (Bool) **IndexByBatch**: `Default=true`: Flag para determinar si el consumo del api de ZincSearch se hace por lotes.
- `--s`: (String) **BatchSize**: `Default=10000`: Cantidad de correos a procesar por lote. Es importante buscar un valor adecuado que no afecte el envío a ZincSearch, ya que un lote muy grande, implica un `body` muy grande, que podría llegar a fallar.

## Ejecución

- Ejecutar el programa sin compilar

```bash
go run main.go [flags...]
go run main.go --i=emails --p=./mock/maildir --r=10 --b=true --s=10000
```

- Compilar el indexador

```bash
go build --o indexer.exe
```

- Ejecutar el indexador compilado

```bash
./indexer.exe --i=emails --p=./mock/maildir --r=10 --b=true --s=10000
```

## Profiling

Para visualizar gráficamente el reporte se debe de instalar la herramienta [Graphviz](https://graphviz.org/download/).

- Ejecutar test con profiling

```bash
cd test

go test --cpuprofile=prof/cpu.out --memprofile=prof/mem.out --mutexprofile=prof/mutex.out --blockprofile=prof/block.out
```

- Visualizar el reporte interactivo web

```bash
go tool pprof --http=:[port] [profile-file-name].out

go tool pprof --http=:8081 prof/cpu.out
go tool pprof --http=:8082 prof/mem.out
go tool pprof --http=:8083 prof/mutex.out
go tool pprof --http=:8084 prof/block.out
```

- Generar PDF

```bash
go tool pprof prof/cpu.out
go tool pprof prof/mem.out
go tool pprof prof/mutex.out
go tool pprof prof/block.out
```
