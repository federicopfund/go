package main

/// Uncomment to use the standar library///
import (
	"github.com/gin-gonic/gin"
)

//"fmt"
//"net/http"

//"github.com/gofiber/fiber/v2"

//func main(){
//	http.HandleFunc("/",handler)
//	http.ListenAndServe(":8080", nil)
//}

//func handler(w http.ResponseWriter, r *http.Request){
//	if r.URL.Path != "/"{
//		http.NotFound(w,r)
//		return
//	}
//	fmt.Fprint(w,"Hello, World")
//}

///// USE Gin FRAMEWORK

func main() {
 	r := gin.Default()
 	r.GET("/", handler)
 	r.Run(":8080")
 }

 func handler(c *gin.Context) {
 	c.String(200, "Hello, World")
 }

// Uncomment to use fiber

// func main() {
//  	app := fiber.New()

// app.Get("/", func(c *fiber.Ctx) error {
// 	return c.SendString("Hello, World!")
// })

// 	app.Listen(":8080")
// }


// func main(){
// 	http.HandleFunc("/",handler)
// 	http.ListenAndServe(":8080",nil)

// }

// func handler(w http.ResponseWriter, r *http.Request){
// 	//Leer el contenido del archivo HTML
// 	html, err := os.ReadFile("index.html")
// 	if err != nil {
// 		http.Error(w,"No se puede cargar el archivo HTML", http.StatusInternalServerError)
// 		return
// 	}

// 	// Estableca el encabeado del tipo de contenido en HTML
// 	w.Header().Set("Content-Type","text/html")

// 	// Escriba el archivo HTML en la respuesta

// 	fmt.Fprintf(w,"%s",html)
// }
