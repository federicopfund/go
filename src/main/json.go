package main

import (
	"encoding/json"
	"fmt"
)

// Definir el tipo de estructura a codificar
type Person struct {
	Name string `json:"name"` // La etiqueta json se usa para especificar el nombre del campo en la salida JSON, pero no es obligatoria.
	Age  int    `json:"age"`
}

func main() {
	encodeJson() // convertir una estructura de datos a JSON
	decodeJson() // convertir JSON a una estructura de datos
	decodeIntoMap()
	encodingSliceOfStructs()
	encodingJsonWithIndent()
	decodeIntoSlice()
	encodeStructWithOmitEmpty()
}

func encodeJson() {
	// Cree una nueva instancia de la estructura y rellénela
	person := Person{
		Name: "John Doe",
		Age:  32,
	}

	// Codifique la estructura a JSON usando la función json.Marshal
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Encode JSON:")
	fmt.Println(string(jsonData))
}

func decodeJson() {
	// Crear una cadena JSON
	jsonData := []byte(`{"name":"John Doe","age":32}`)

	// Decode the JSON string to a struct using the json.Unmarshal function
	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nDecode JSON:")
	fmt.Printf("Name: %s\nAge: %d\n", person.Name, person.Age)
}

func decodeIntoMap() {
	// Create a JSON string
	jsonData := []byte(`{"name":"John Doe","age":32}`)

	// Decodifique la cadena JSON en un mapa usando la función json.Unmarshal
	// El tipo de interfaz{} se usa para representar valores de cualquier tipo. Esto es útil cuando no conoce el tipo de datos JSON.
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data) // & se utiliza para pasar un puntero a la variable de datos para cambiar directamente
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nDecode JSON into Map:")
	fmt.Println("Name:", data["name"]) // Las claves del mapa son cadenas y los valores son de tipo interfaz{}
	fmt.Println("Age:", data["age"])
}

func encodingSliceOfStructs() {
	// Crear una porción de estructuras de persona
	people := []Person{
		{Name: "John Doe", Age: 32},
		{Name: "Jane Doe", Age: 28},
	}

	// Codifique el segmento en JSON usando la función json.Marshal
	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nEncoding slice of structs:")
	fmt.Println(string(jsonData))
}

func encodingJsonWithIndent() {
	//Cree una nueva instancia de la estructura y rellénela

	person := Person{
		Name: "John Doe",
		Age:  32,
	}

	
// Codificar la estructura a JSON usando la función json.MarshalIndent
// El primer argumento es el valor a codificar
// El segundo argumento es un prefijo para anteponer a cada línea de salida
// El tercer argumento es la cadena de sangría a usar
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nEncoding JSON with indent:")
	fmt.Println(string(jsonData))
}

func decodeIntoSlice() {
	// Crear una cadena JSON
	jsonData := []byte(`[{"name":"John Doe","age":32},{"name":"Jane Doe","age":28}]`)

 
	//Decodifique la cadena JSON en un segmento usando la función json.Unmarshal
	var people []Person                      // 
											// El segmento se inicializa en un segmento vacío de estructuras de persona.
	err := json.Unmarshal(jsonData, &people) // & se utiliza para pasar un puntero a la variable de personas para cambiar directamente
	if err != nil {
		fmt.Println(err)
		return
	}

	// Recorra el corte e imprima los valores
	fmt.Println("\nDecoding JSON into slice:")
	for _, person := range people {
		fmt.Printf("Name: %s\nAge: %d\n", person.Name, person.Age)
	}
}

func encodeStructWithOmitEmpty() {
	// Cree una nueva instancia de la estructura y complétela con solo 1 campo
	person := Person{
		Name: "John Doe",
	}

	//Codifique la estructura a JSON usando la función json.Marshal
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	// El campo Edad no se incluye en la salida JSON porque está vacío
	fmt.Println("\nEncode struct with OmitEmpty:")
	fmt.Println(string(jsonData))
}