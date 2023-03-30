
### Creando tu primer servidor web en Go

#

Creando tu primer servidor web en Go
¡Hola a todos y bienvenidos al episodio de hoy en el que crearemos su primer servidor web en Go! Esta sesión está diseñada para nuevos desarrolladores en la industria tecnológica y aquellos que quieren aprender sobre el lenguaje.

En esta sesión, utilizaremos la biblioteca estándar en Go y mostraremos el poder de GitHub Copilot, una poderosa herramienta que puede ayudarlo a escribir código de manera más rápida y eficiente.

Al final de esta sesión, tendrá una sólida comprensión de cómo construir un servidor web básico en Go y estará equipado con las herramientas para continuar explorando este emocionante lenguaje. ¡Entonces, tome su café y únase a nosotros para una sesión divertida e informativa!

Para ejecutar esto en Docker o Podman localmente

[link](https://www.youtube.com/live/uhhxPZNKRWM?feature=share)

``` sh

// navigate to the root of the project file
cd web-server

// build the image
docker build --tag docker.io/goshow:v1 .   

// run the image
docker run -dt -p 8080:8080/tcp docker.io/library/goshow:v1

```