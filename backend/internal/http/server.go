package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"re/internal/service"

	"github.com/gorilla/handlers"
)

type Server struct {
	service *service.Service
}

// NewServer initializes a new Server instance
func NewServer() *Server {
	return &Server{
		service: service.NewService(),
	}
}

// Algorithm defines the expected structure for the algorithm input
type Algorithm struct {
	Number int `json:"number"`
}

// EditInput defines the expected structure for the edit input
type EditInput struct {
	PackSizes []int `json:"pack_sizes"`
}

// SolveAlgorithmHandler handler, calls the algorithm service to solve the algorithm
func (s *Server) SolveAlgorithmHandler(w http.ResponseWriter, r *http.Request) {
	var input Algorithm
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding request body: %v", err)
		return
	}

	result := s.service.SolveAlgorithm(input.Number)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"result": result}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error encoding response: %v", err)
		return
	}
}

// EditInputParametersHandler handler, edit pack sizes with newly provided ones
func (s *Server) EditInputParametersHandler(w http.ResponseWriter, r *http.Request) {
	var input EditInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding request body: %v", err)
		return
	}

	s.service.EditPackSizes(input.PackSizes)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"message": "Pack sizes updated successfully"}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error encoding response: %v", err)
		return
	}
}

// GetParams handler, gets current pack sizes
func (s *Server) GetParams(w http.ResponseWriter, _ *http.Request) {
	result := s.service.GetPackSizes()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"result": result}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error encoding response: %v", err)
		return
	}
}

// StartServer starts the server, listens and serves
func (s *Server) StartServer() {
	router := http.NewServeMux()
	router.HandleFunc("/solve", s.SolveAlgorithmHandler)
	router.HandleFunc("/get-parameters", s.GetParams)
	router.HandleFunc("/parameters", s.EditInputParametersHandler)

	// CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
	)

	// Wrap the router with the CORS middleware
	http.Handle("/", corsHandler(router))

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
