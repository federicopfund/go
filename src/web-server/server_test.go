package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	
// Crear una nueva solicitud con método GET y sin cuerpo de solicitud
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear una nueva grabadora para registrar la respuesta
	rr := httptest.NewRecorder()

	// Llamar al controlador con la solicitud y la grabadora
	handler(rr, req)

	
	// Comprobar el código de estado de la respuesta
	gotStatus := rr.Code
	wantStatus := http.StatusOK
	if gotStatus != wantStatus {
		t.Errorf("El controlador devolvió un código de estado incorrecto: got %v, want %v", gotStatus, wantStatus)
	}

	// Revisar el cuerpo de la respuesta
	gotBody := rr.Body.String()
	wantBody := "Hola Mundo!"
	if gotBody != wantBody {
		t.Errorf("El controlador devolvió un cuerpo inesperado: got %v, want %v", gotBody, wantBody)
	}
}