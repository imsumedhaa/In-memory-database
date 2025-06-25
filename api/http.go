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

func NewHttp(port, username, password, dbname string)(*Http, error){
	dbClient, err :=postgres.NewClient(port, username, password, dbname)

	if err != nil{
		return nil,fmt.Errorf("failed to connect: %w",err)
	}
	return &Http{client: dbClient},nil
}


func (h *Http)create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	
	if err := h.client.CreatePostgresRow("key1", "value1");err!=nil{
		//http.Error()
	}

	response := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Http)delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := h.client.DeletePostgresRow("key1"); err != nil{
		//
	}


}

func (h *Http)update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := h.client.UpdatePostgresRow("key1","val1"); err != nil{
		//
	}


}

func (h *Http)get(w http.ResponseWriter, r *http.Request)() {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if _, err := h.client.GetPostgresRow("key1"); err != nil{
		//
	}


}

func (h *Http)show(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if _, err := h.client.ShowPostgresRow(); err != nil{
		//
	}


}




func (h *Http)Run() error {

	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}
	return nil
}

func (h *Http)routes() {
	http.HandleFunc("/Create", h.create)
	http.HandleFunc("/Update", h.update)
	http.HandleFunc("/Delete", h.delete)
	http.HandleFunc("/Get", h,get)
	//http.HandleFunc("/Show", helloHandler)
}
