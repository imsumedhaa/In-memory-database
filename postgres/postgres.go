package postgres

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/imsumedhaa/In-memory-database/database"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() (database.Database, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Postgres Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter Postgres Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Print("Enter Postgres DB Name: ")
	dbname, _ := reader.ReadString('\n')
	dbname = strings.TrimSpace(dbname)

	fmt.Print("Enter Postgres Port (default 5432): ")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSpace(port)
	if port == "" {
		port = "5432"
	}

	// Build connection string
	connStr := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		port, username, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	// Create a simple table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS kvstore (
		key TEXT PRIMARY KEY,
		value TEXT
	)`)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	fmt.Println("Connected to Postgres successfully.")

	return &Postgres{db: db}, nil
}

func (p *Postgres) Create() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the key:")
	key, _ := reader.ReadString('\n')
	key = strings.TrimSpace(key)

	fmt.Println("Enter the value:")
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)

	if key == "" || value == "" {
		fmt.Println("Key and value cannot be empty.")
		return
	}

	// Check if key already exists
	var existing string
	err := p.db.QueryRow("SELECT key FROM kvstore WHERE key = $1", key).Scan(&existing)
	if err == nil {
		fmt.Println("Key already exists. Use 'update' to change the value.")
		return
	}

	// Insert new key-value pair
	_, err = p.db.Exec("INSERT INTO kvstore (key, value) VALUES ($1, $2)", key, value)
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return
	}

	fmt.Println("Data inserted successfully.")
}

func (p *Postgres) Delete() {

	reader:= bufio.NewReader((os.Stdin))

	fmt.Println("Enter the key you want to delete:")
	key,_:= reader.ReadString('\n')
	key = strings.TrimSpace(key)

	if key == "" {
		fmt.Println("Key cannot be empty.")
		return
	}

	var existing string

	err := p.db.QueryRow("SELECT key FROM kvstore WHERE key = $1", key).Scan(&existing)

	if err==sql.ErrNoRows{
		fmt.Println("Key does not exists.")
		return
	}else if err !=  nil{
		fmt.Println("Error is ", err)
	}else{
		// Delete new key-value pair
	_, err = p.db.Exec("DELETE FROM kvstore WHERE key = $1", key, )
	if err != nil {
		fmt.Println("Error deleting data:", err)
		return
	}
	fmt.Println("Key deleted successfully.")
	}

}

func (p *Postgres) Update() {
	

}

func (p *Postgres) Get() {
}

func (p *Postgres) Show() {
}

func (p *Postgres) Exit() {
}

