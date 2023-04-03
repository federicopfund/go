# Los conceptos fundamentales dentro de Go
El propósito de esta sección es equipar a nuestros oyentes con todo el conocimiento que requieren para nuestra sesión, pero no les enseñaremos el Go desde cero.

## Escritura fuerte
- El compilador es tu mejor amigo. 
- En todos los puntos, conoceremos el tipo de nuestras variables, y qué comportamiento exponen.
- La escritura dinámica es más lenta y el compilador evita errores de tiempo de ejecución y errores fatales debido a comportamientos indefinidos.
- El paquete fmt es parte de la librería estándar y nos permite formatear e imprimir cadenas.
- La cadena de herramientas Go construye y ejecuta nuestros programas.

```bash
$ go run fundamentals/strong-typing/main.go
```
## Funciones
- Las funciones go son soportadas de forma nativa y se pueden pasar como variables, tipos de retorno y parámetros para invocación posterior. 
- También se permiten funciones anónimas.
- Composición de la función es fácil de hacer en Go. 
- Las funciones diferidas son útiles para garantizar las tareas de limpieza. 

```bash
$ go run fundamentals/functions/main.go
```

## Error handling
- Go functions can return multiple values.
- By convention, the error is the last returned value using the built-in error type. 
- The zero value of the error type is nil.
- Errors should be handled first, keeping code minimally indented.


```bash
$ go run fundamentals/error-handling/main.go
```


## Manejo de errores
- Las funciones go pueden devolver varios valores.
- Por convención, el error es el último valor devuelto usando el tipo de error incorporado. 
- El valor cero del tipo de error es nulo.
- Los errores deben ser manejados primero, manteniendo el código mínimamente sangrado.

```bash
$ go run fundamentals/structs/main.go
```


## Visibilidad
- Go code está organizado en paquetes, que controlan la visibilidad de las variables, tipos y funciones que contienen.
- Una carpeta puede contener solo un paquete, pero el paquete no necesita ser nombrado después del directorio.
- Los nombres solo pueden usarse una vez dentro del mismo paquete.
- Los programas ejecutables tienen una función principal definida en un paquete principal.
- Podemos exportar campos fuera de su paquete con la primera letra de su nombre. 

```bash
$ go run fundamentals/visibility/main.go
```
## Interfaces
- Las interfaces son colecciones de firmas de métodos. 
- Se implementan automáticamente por el compilador en tipos que satisfacen toda la colección de métodos. 
- Son la principal forma de implementar el polimorfismo en Go.
- Las interfaces a menudo se exportan, mientras que las estructuras permanecen visibles solo dentro del paquete. 

```bash
$ go run fundamentals/visibility/main.go
```


## Goroutines
- Los goroutines se conocen como hilos ligeros. Se utilizan para ejecutar funciones simultáneamente dentro de nuestros programas Go.
- Le indicamos al tiempo de ejecución de Go que ejecute una función en una nueva goroutine usando la palabra clave `go`.
- Iniciar un goroutine es no bloquear por diseño, de lo contrario estaríamos ejecutando las cosas secuencialmente.
- El programa funciona en su propio goroutine, conocido como goroutine principal. 
- El goroutine principal tiene una relación de padre e hijo con las goroutines que inicia.

```bash
$ go run fundamentals/goroutines/main.go
```
## Canales
- Se desaconseja pasar información entre goroutines usando variables de memoria compartida.
- Los canales son tubos que permiten pasar información de una manera segura para los hilos.
- El tipo de variable que soporta el canal forma parte de su inicialización.
- La operación de envío escribe información a través de un canal, mientras que la operación de recepción lee información del canal.
- Los envíos y las recepciones en un canal son operaciones de bloqueo. Se pueden utilizar para la sincronización de goroutines.
- Los mensajes solo se leen una vez.
- Una vez completadas las operaciones, los canales pueden cerrarse para indicar a otros que no se enviarán más valores a través de ellos.

```bash
$ go run fundamentals/channels/main.go
```

## Canales en búfer
- De forma predeterminada, los canales no están conectados. 
- Requieren que tanto el emisor como el receptor estén disponibles para la operación. Estas operaciones son sincrónicas.
- Si un lado está disponible sin el otro, se bloqueará hasta que sea posible la operación opuesta correspondiente.
- Los canales pueden ser almacenados con una capacidad predeterminada para mantener los valores de los remitentes hasta que lleguen los receptores.
- Si hay espacio en la cola del canal, la operación se completa inmediatamente.

```bash
$ go run fundamentals/buffered-channels/main.go
```


## Pruebas de unidad
- El paquete de pruebas de Go nos permite escribir pruebas, verificaciones y referencias.
- Viniendo de otros idiomas, puede parecer que el paquete de pruebas estándar de Go es barebones.
- Podemos complementarlo con otras bibliotecas de terceros, pero es bueno empezar por entender cómo escribir pruebas primero.
- Probar código concurrente no puede probar la ausencia de errores, pero puede darnos una confianza estadística del comportamiento de nuestro código bajo ciertas condiciones.

```bash
$ go test -run=TestSayHi ./fundamentals/unit-test
```


## El paquete net/http
- El paquete HTTP de Go es fácil de usar y una de las razones por las que Go es tan ampliamente utilizado en el desarrollo de aplicaciones web.
- Los controladores responden a las peticiones HTTP. Las funciones que sirven como controladores toman dos parámetros: `http.ResponseWriter` y `http.Request`.
- Las funciones del controlador están registradas en una ruta HTTP particular usando `http.HandleFunc`.
- El paquete `net/http` es responsable de pasar los encabezados y las solicitudes a nuestras funciones de control registradas personalizadas.

```bash
$ go run fundamentals/server/main.go
```