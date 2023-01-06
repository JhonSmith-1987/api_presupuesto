package routes

import "net/http"

func HomeRoutes(response http.ResponseWriter, reques *http.Request) {
	response.Write([]byte("sin palabras"))
}
