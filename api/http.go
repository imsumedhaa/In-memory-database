package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func NewHttp(host,port, username, password, dbname string) (*Http, error) {
	dbClient, err := postgres.NewClient(host,port, username, password, dbname)

	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}
	return &Http{client: dbClient}, nil
}

func (h *Http) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid create body request: %s", err), http.StatusBadRequest)
		return
	}

	if req.Key == "" {
		http.Error(w, "Key cannot be empty", http.StatusBadRequest)
		return
	}

	if err := h.client.CreatePostgresRow(req.Key, req.Value); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create row: %s", err), http.StatusInternalServerError)
		return
	}

	response := Response{Message: "Row created succesfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Http) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid update body request: %s", err), http.StatusBadRequest)
		return
	}

	if req.Key == "" {
		http.Error(w, "Key cannot be empty", http.StatusBadRequest)
		return
	}

	if err := h.client.UpdatePostgresRow(req.Key, req.Value); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update row: %s ", err), http.StatusInternalServerError)
		return
	}

	response := Response{Message: "Row updated succesfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Http) delete(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid Delete body request: %s", err), http.StatusBadRequest)
		return
	}
	if req.Key == "" {
		http.Error(w, "Key cannot be empty", http.StatusBadRequest)
		return
	}


	if err := h.client.DeletePostgresRow(req.Key); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete row: %s", err), http.StatusInternalServerError)
		return
	}

	response := Response{Message: "Row deleted succesfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Http) get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid get body request: %s", err), http.StatusBadRequest)
		return
	}

	if req.Key == "" {
		http.Error(w, "Key cannot be empty", http.StatusBadRequest)
		return
	}

	value, err := h.client.GetPostgresRow(req.Key)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get the row: %s", err), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"Key":   req.Key,
		"Value": value,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func (h *Http) show(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	store, err := h.client.ShowPostgresRow()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to show row: %s", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store)
}

func (h *Http) Run() error {
	h.routes()
	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}
	return nil
}

func (h *Http) routes() {
	http.HandleFunc("/create", h.create)
	http.HandleFunc("/update", h.update)
	http.HandleFunc("/delete", h.delete)
	http.HandleFunc("/get", h.get)
	http.HandleFunc("/show", h.show)
}

