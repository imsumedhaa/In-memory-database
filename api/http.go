package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/imsumedhaa/In-memory-database/database"
	"github.com/imsumedhaa/In-memory-database/pkg/client/postgres"
)

type Response struct {
	Message string `json:"message"`
}

type Request struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type Http struct {
	client postgres.Client
}

// Create implements database.Database.
func (h *Http) Create(key string, value string) error {
	panic("unimplemented")
}

// Delete implements database.Database.
func (h *Http) Delete(key string) error {
	panic("unimplemented")
}

// Exit implements database.Database.
func (h *Http) Exit() error {
	panic("unimplemented")
}

// Get implements database.Database.
func (h *Http) Get(key string) error {
	panic("unimplemented")
}

// Show implements database.Database.
func (h *Http) Show() error {
	panic("unimplemented")
}

// Update implements database.Database.
func (h *Http) Update(key string, value string) error {
	panic("unimplemented")
}

func NewHttp(port, username, password, dbname string) (database.Database, error) {
	dbClient, err := postgres.NewClient(port, username, password, dbname)

	if err != nil {
		return nil, fmt.Errorf("failed to connect %w", err)
	}
	return &Http{client: dbClient}, nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	routes()
	dbClient, _ := postgres.NewClient("5431", "admin", "Sumedha1234", "mydb")
	dbClient.CreatePostgresRow("key1", "value1")

	response := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Run() {

	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() {
	http.HandleFunc("/Create", helloHandler)
	http.HandleFunc("/Update", helloHandler)
	http.HandleFunc("/Delete", helloHandler)
	http.HandleFunc("/Get", helloHandler)
	http.HandleFunc("/Show", helloHandler)
}
