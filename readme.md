# KALYPSO APP | VICENT COLL Y ALEX ZARAZUA

**INDEX**

  * Tecnologías Implicadas
    * GO
    * LARAVEL
    * REDIS 
    * ANGULAR 
    * PROMETHEUS
    * GRAFANA
  * Proceso de Desarrollo

# Tecnologías Implicadas 

## Lenguaje de Backend Go

 **Go o Golang** es un lenguaje de programación de código abierto creado por Google  en 2009 y que ayuda en la creación de software de manera fácil, eficiente y de alto rendimiento.

_A pesar de ser un lenguaje muy joven tiene un rendimiento similar a C pero con la sintaxis amigable parecida a Python, actualmente es utilizado en aquellos programas que requieran alto rendimiento y un ejemplo de ello es Docker_

* VENTAJAS DE GO
    * Velocidad similar a C pero con una sintaxis amigable como Python.
    * Facilita el uso de buenas prácticas en el código.
    * Mantiene su rendimiento con grandes volúmenes de información.
    * Su curva de aprendizaje es suave en comparación con Java o C.
    * Es un lenguaje multiparadigma.

 Una de las principales funciones de _Golang_ es que maneja las múltiples tareas de forma concurrente, es decir, que el procesador divide el trabajo en múltiples hilos y a cada uno le dedica un tiempo en milisegundos acelerando el tiempo de ejecución en cada tarea sin necesidad de realizar configuraciones adicionales ya que _Golang_ lo traen en sus librerías preinstaladas.

 La manera como _Golang_ trabaja esto es a través de Goroutines y Canales. Puedes tener múltiples Goroutines por canal y transportar el Goroutine de un canal a otro. Un excelente uso de Goroutines es al momento de trabajar transferencia de datos en tiempo real desde un cliente a otro, como por ejemplo un chat o un sistema de ubicación en tiempo real.

###  Frameworks populares

     Echo ->  Minimalista, escalable y de alto rendimiento y si lo combinas con Fresh podrás hacer que cada cambio guardado se recompile y ejecute.

     Revel -> Es bastante completo para hacer proyectos complejos.

     Gin -> Es rápido en su ejecución e ideal para hacer prototipos de de microservicios RESTful API.

## Lenguaje de Backend Laravel

**Laravel** es un framework open-source de PHP desarrollado por el MIT en 2011. Su principal objetivo es _simplificar y estilizar el código y que sea más sencillo de leer, interpretar y programar, virtud de la que carecía la programación del lenguaje nativo PHP._

Consigue su objetivo gracias a componentes y dependencias,con el transcurso de los años ha conseguido crear dependencias propias mejorando las de otros frameworks. Muchas de estas dependencias son del pionero francés Symfony.

Así como lo era y es para PHP, el principal target de Laravel son las páginas y aplicaciones web.

### Puntos fuertes de hacer un backend con Laravel

    Documentación --> Si es verdad que para un programador junior o principiante la palabra framework asusta, en Laravel no. Su documentación oficial es de las más completas, de las más simples y de las mejor explicadas --> https://laravel.com .

    Eloquent ORM --> Cómo ha sido nombrado anteriormente, este súper-paquete de Laravel nos permite sustituir las consultas SQL y de MongoDB por una simple Programación Orientada a Objetos. De este modo, Eloquent se encarga de convertir nuestros objectos a queries para cualquiera de las bases de datos que acepta.

    Routing o rutas -->  Si bien en Node.js, por convención, las rutas se definen por entidades o modelos en ficheros separados, en Laravel los endpoints se agrupan por funcionalidades de la aplicación (web.php si es en un entorno web o api.php si es para una API), para poder organizar mejor las rutas de nuestra aplicación.
    Además, nos da la posibilidad de agrupar las rutas, permitiéndonos así asignarles prefijos, sufijos y middlewares.

    Middlewares más sencillos --> Son controladores que se ejecutan antes o después de una petición para evitar repetirlos en nuestras funciones sistemáticamente. Laravel nos da ejemplos y una estructura para aprender como usarlos. Los más comunes son las autenticaciones de tokens de sesión, los permisos de un usuario, etc.

    La consola artisan -->  Laravel incorpora el cli artisan, que nos permite en comandos muy lógicos y sencillos realizar acciones que otros frameworks no nos permiten, como crear ficheros (modelos, controladores, providers …), limpiar la cache, etc. En definitiva, una maravilla, no tenemos que tenerle miedo a la terminal, es nuestra amiga --> https://laravel.com/docs/7.x/artisan 







