# Configuración del servidor

## Start the API server

```bash
go run main.go
```


## RESTful Routes

#### GET /persons - Obtener todas las personas


Llama a la API con curl:
```bash
curl -X GET http://localhost:8080/persons
```


Devuelve una rebanada de objetos JSON:
```json
[
    {
        "id": "1234",
        "name": "Liam Hampton",
        "fact": "Likes F1",
        "location": "London"
    },
    {
        "id": "1234",
        "name": "Adelina Simion",
        "fact": "Likes coding",
        "location": "London"
    }
]
```


### POST /persons - Crear una persona

Llama a la API con curl:

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "name": "Adelina Simion",
    "fact": "Likes coding",
    "location": "London"
}' localhost:8080/persons
```

### TODO: POST /person/{id} - Eliminar una persona

Llama a la API con curl:

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "id": "1234"
}' localhost:8080/person/{id}
```
