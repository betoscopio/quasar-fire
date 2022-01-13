# Operación Fuego de Quasar


> Como jefe de comunicaciones rebelde, tu misión es crear un programa en Golang que retorne la fuente y contenido del mensaje de  auxilio . Para esto, cuentas con tres satélites que te permitirán triangular la posición, ¡pero cuidado! el mensaje puede no llegar completo a cada satélite debido al campo de asteroides frente a la nave.

## Modo de uso

### Ejecución local

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

## Generación de Imagen Docker (local)


### Requisitos

Tener instalado Docker de manera local.

```
$ docker build -f Dockerfile -t usuario/quasar-fire:tag .
```

### Ejecución

Si se ejecutó el *build* local.
```
$ docker run -p 8080:8080 usuario/quasar-fire:tag
```

Si se desa usar la imagen de DockerHub
```
$ docker run -p 8080:8080 betoscopio/quasar-fire:1.1.1
```

## Despliegue y versión en línea

Se ha definido la infraestructura siguiente:

- Github + Github Actions, para almacenamiento de código, generación de imágenes Docker y despliegue a Google Cloud Platform
- Google Container Registry, que tiene buena integración con otros servicios, para almacenamiento de imagenes Docker generadas. Adicionalmente se sube a Docker Registry, para posible uso en otro tipo de despligue, pero no se usa en este caso.
- *Google Cloud Run*, el cual hace uso de la imagen subida al *Container Registry* y hace el servicio público.

### Generación de tags

Se ha definido un proceso automatizado de generación de tags tanto para relases en *main*, como para las imágenes docker que son subidas al *Container Registry*. Esto se puede encontrar en el directorio `scripts/` el cual está se basa en https://github.com/antonputra/lesson-087/blob/main/scripts/git_update.sh. 

## Versión en línea

Se puede visitar en la URL https://quasar-fire-d4z3rckaeq-uc.a.run.app/.

## Modo de uso

Urls existentes

- `/` : página de bienvenida.
- `/topsecret`: permite envío de señales en grupo para obtener una respuesta de ubicación si existen los datos necesarios (usado método POST). //TODO
- `/topsecret_split/{satellite_name}`: Permite interactuar con un satelite en particular. //TODO
  - Método GET, permite obtener la ubicación de la nave si existen los datos suficientes.
  - Método POST, permite enviar una señal única nueva al satélite en particular
- `/satellites`: muestra los satelites ingresados al sistema.
- `/signals`: muesta las señales enviadas. 
  - Método GET, permite ver cuales son las ingresadas.
  - Método POST, permite enviar una señal única nueva.

Obtener satélites ingresados:
```
$ curl -i localhost:8080/satellites
```

Revisar señales ingresadas:
```
$ curl -i localhost:8080/signals
```

Ingresar nueva señal:
```
$ curl localhost:8080/signals -X POST -H "Content-Type: application/json"  -d '{"name": "fisto","distance": 80, "message": ["este", "es", "un", "","secreto"]}'
```

Ingresar multiples señales:

```
$ curl localhost:8080/signals -X POST -H "Content-Type: application/json"  -d '{"name": "fisto","distance": 80, "message": ["este", "es", "un", "","secreto"]}, '
```
{ "satellites": [
        {"name": "fisto","distance": 80, "message": ["este", "es", "un", "","secreto"]},
        {"name": "koon","distance": 125, "message": ["", "", "un", "mensaje","secreto"]},
        {"name": "windu","distance": 145, "message": ["", "", "un", "mensaje","secreto"]}
    ]
}

```
$ curl localhost:8080/topsecret -X POST -H "Content-Type: application/json"  -d           
'{ "satellites": [                                                                      
        {"name": "fisto","distance": 80, "message": ["este", "es", "un", "","secreto"]},
        {"name": "koon","distance": 125, "message": ["", "", "un", "mensaje","secreto"]},
        {"name": "windu","distance": 145, "message": ["", "", "un", "mensaje","secreto"]}
    ]
}'
```

```
$ curl localhost:8080/topsecret -X POST -H "Content-Type: application/json"  -d '{ "satellites": [{"name": "fisto","distance": 80, "message": ["este", "es", "un", "","secreto"]},{"name": "koon","distance": 125, "message": ["", "", "un", "mensaje","secreto"]},{"name": "windu","distance": 145, "message": ["", "", "un", "mensaje","secreto"]}]}'
```

```
$ curl localhost:8080/topsecret -X POST -H "Content-Type: application/json"  -d '{ "satellites": ["a","b","c","d"]}'
```