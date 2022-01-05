# Operación Fuego de Quasar


> Como jefe de comunicaciones rebelde, tu misión es crear un programa en Golang que retorne la fuente y contenido del mensaje de  auxilio . Para esto, cuentas con tres satélites que te permitirán triangular la posición, ¡pero cuidado! el mensaje puede no llegar completo a cada satélite debido al campo de asteroides frente a la nave.

## Modo de uso

### Ejecución local

Descar
Ingresar al directorio `quasar-fire`
```
$ git clone https://github.com/betoscopio/quasar-fire.git
$ cd quasar-fire/
$ go run main.go
```

### Una vez ejecutandose el servicio

Visitar URLs:
- http://localhost:8080/, para mensaje de bienvenida
- http://localhost:8080/satellites, para ver los satelites preingresados

Mediante línea de comando, con cliente `curl`.

Mensaje de Bienvenida
```
$ curl localhost:8080/
```

Retorna listado de satelites predefinidos en formato JSON.
``` 
$ curl localhost:8080/satellites
```