package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Person representa un documento de persona en la base de datos
type Person struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Fact     string             `bson:"fact" json:"fact"`
	Location string             `bson:"location" json:"location"`
}

func main() {
	
// Establecer el URI de conexión de MongoDB, incluido el nombre de usuario, la contraseña y el nombre de la base de datos
	uri := os.Getenv("MONGODB_CONNECTION_STRING")

	// Establecer las opciones del cliente
	clientOptions := options.Client().ApplyURI(uri)

	
	// Establecer el contexto con un tiempo de espera de 10 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	/// Conéctese a la instancia de MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	
	// Comprobar la conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Cosmos DB MongoDB instance!")

	http.HandleFunc("/persons", func(w http.ResponseWriter, r *http.Request) {
		
	// Establecer encabezados necesarios para manejar solicitudes CORS// Set necessary headers for handling CORS requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		switch r.Method {
		case "GET":
			getAllPeople(w, r, client)
		case "POST":
			createPerson(w, r, client)
		case "OPTIONS":
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Dejo esto por ahora, pero lo eliminaré más tarde
	http.HandleFunc("/person/delete", func(w http.ResponseWriter, r *http.Request) {
		log.Print("deletePerson called")
		deletePerson(w, r, client)
	})

	
	// Inicie el servidor HTTP
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


// createPerson crea un nuevo documento de persona en la base de datos
func createPerson(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	
// Analizar el cuerpo de la solicitud en una estructura de persona
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	
	// Establezca el campo ID en una nueva ID única
	person.ID = primitive.NewObjectID()

	
	// Obtener la colección MongoDB del cliente
	collection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION"))

	// Insertar el nuevo documento de persona en la colección
	_, err = collection.InsertOne(context.Background(), &person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
// Establecer el código de estado de respuesta en 201 Creado y devolver la identificación de la persona recién creada
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct{ ID string }{person.ID.Hex()})
}

// deletePerson elimina un documento de persona de la base de datos
func deletePerson(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	
// Leer el cuerpo de la solicitud
	decoder := json.NewDecoder(r.Body)
	var data struct {
		ID string `json:"id"`
	}
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	
	// Validar el parámetro ID
	if data.ID == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}
	id, err := primitive.ObjectIDFromHex(data.ID)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	
	// Obtener la colección MongoDB del cliente
	collection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION"))

	
	// Eliminar el documento de la persona con el ID especificado
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if result.DeletedCount == 0 {
		http.Error(w, "Person document not found", http.StatusNotFound)
		return
	}

	
// Establecer el código de estado de respuesta en 204 Sin contenido para indicar una eliminación exitosa
	w.WriteHeader(http.StatusNoContent)
}

// getAllPeople devuelve todos los documentos de personas de la base de datos
func getAllPeople(w http.ResponseWriter, r *http.Request, client *mongo.Client) {
	
	// Obtener la colección MongoDB del cliente
	collection := client.Database(os.Getenv("MONGODB_DATABASE")).Collection(os.Getenv("MONGODB_COLLECTION"))

	
	// Encuentra todos los documentos de personas en la colección
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	
// Iterar a través del cursor y decodificar cada documento de persona en una estructura de persona
	var people []Person
	for cursor.Next(context.Background()) {
		var person Person
		err = cursor.Decode(&person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		people = append(people, person)
	}

	// Devuelve el segmento de personas como JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(people)
}