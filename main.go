package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"apiTest_wicho/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// datos estáticos
var users = []models.User{
	{ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
	{ID: 3, Name: "Alice Johnson", Email: "alice.johnson@example.com"},
	{ID: 4, Name: "Bob Brown", Email: "bob.brown@example.com"},
	{ID: 5, Name: "Charlie White", Email: "charlie.white@example.com"},
	{ID: 6, Name: "David Green", Email: "david.green@example.com"},
	{ID: 7, Name: "Eve Black", Email: "eve.black@example.com"},
	{ID: 8, Name: "Frank Blue", Email: "frank.blue@example.com"},
	{ID: 9, Name: "Grace Red", Email: "grace.red@example.com"},
	{ID: 10, Name: "Hank Yellow", Email: "hank.yellow@example.com"},
	{ID: 11, Name: "John Doe", Email: "john.doe@example.com"},
	{ID: 12, Name: "Jane Smith", Email: "jane.smith@example.com"},
	{ID: 13, Name: "Alice Johnson", Email: "alice.johnson@example.com"},
	{ID: 14, Name: "Bob Brown", Email: "bob.brown@example.com"},
	{ID: 15, Name: "Charlie White", Email: "charlie.white@example.com"},
	{ID: 16, Name: "David Green", Email: "david.green@example.com"},
	{ID: 17, Name: "Eve Black", Email: "eve.black@example.com"},
	{ID: 18, Name: "Frank Blue", Email: "frank.blue@example.com"},
	{ID: 19, Name: "Grace Red", Email: "grace.red@example.com"},
	{ID: 20, Name: "Hank Yellow", Email: "hank.yellow@example.com"},
	{ID: 21, Name: "John Doe", Email: "john.doe@example.com"},
	{ID: 22, Name: "Jane Smith", Email: "jane.smith@example.com"},
	{ID: 23, Name: "Alice Johnson", Email: "alice.johnson@example.com"},
	{ID: 24, Name: "Bob Brown", Email: "bob.brown@example.com"},
	{ID: 25, Name: "Charlie White", Email: "charlie.white@example.com"},
	{ID: 26, Name: "David Green", Email: "david.green@example.com"},
	{ID: 27, Name: "Eve Black", Email: "eve.black@example.com"},
	{ID: 28, Name: "Frank Blue", Email: "frank.blue@example.com"},
	{ID: 29, Name: "Grace Red", Email: "grace.red@example.com"},
	{ID: 30, Name: "Hank Yellow", Email: "hank.yellow@example.com"},
	{ID: 31, Name: "Charlie White", Email: "charlie.white@example.com"},
	{ID: 32, Name: "David Green", Email: "david.green@example.com"},
	{ID: 33, Name: "Eve Black", Email: "eve.black@example.com"},
	{ID: 34, Name: "Frank Blue", Email: "frank.blue@example.com"},
}

// Handler para obtener usuarios con paginación
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Leer parámetros
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("size")

	// Valores predeterminados
	if pageStr == "" {
		pageStr = "1"
	}
	if limitStr == "" {
		limitStr = "5"
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
		return
	}

	// Calcular los índices
	start := (page - 1) * limit
	end := start + limit
	if start > len(users) {
		start = len(users)
	}
	if end > len(users) {
		end = len(users)
	}

	// Obtener la parte de la lista correspondiente a la página solicitada
	paginatedUsers := users[start:end]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paginatedUsers)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/users", GetUsers).Methods("GET")

	// Configurar el middleware CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Cambia esto según tu frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	fmt.Println("Servidor corriendo en http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
