package postgres

import (
	"database/sql"
	"fmt"
)

type Client interface {
	CreatePostgressRow(key, val string) error
	DeletePostgressRow(key string) error
	UpdatePostgressRow(key, value string) error
	GetPostgressRow(key string) (string, error)
	ShowPostgressRow() (map[string]string,error)
}

// NewClient creates new HCloud clients.
func NewClient(port, username, password, dbname string) (Client, error) {
	// Build connection string
	connStr := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		port, username, password, dbname)

	database, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	// Create a simple table if not exists
	_, err = database.Exec(`CREATE TABLE IF NOT EXISTS kvstore (
		key TEXT PRIMARY KEY,
		value TEXT
	)`)
	if err != nil {
		return nil, fmt.Errorf("failed to create key value: %w", err)
	}

	fmt.Println("Connected to Postgres successfully.")
	return &realClient{db: database}, nil
}

type realClient struct {
	db *sql.DB
}

func (r *realClient) CreatePostgressRow(key, val string) error {

	// Check if key already exists
	var existing string
	err := r.db.QueryRow("SELECT key FROM kvstore WHERE key = $1", key).Scan(&existing)
	if err == nil {
		return fmt.Errorf("key already exists. Use 'update' to change the value")
	}

	// Insert new key-value pair
	_, err = r.db.Exec("INSERT INTO kvstore (key, value) VALUES ($1, $2)", key, val)
	if err != nil {
		return fmt.Errorf("error inserting data: %w", err)
	}

	return nil
}

func (r *realClient) DeletePostgressRow(key string) error {

	var existing string

	err := r.db.QueryRow("SELECT key FROM kvstore WHERE key = $1", key).Scan(&existing)

	if err == sql.ErrNoRows {
		return fmt.Errorf("key does not exist")

	} else if err != nil {
		fmt.Println("Error is ", err)
	} else {
		// Delete new key-value pair
		_, err = r.db.Exec("DELETE FROM kvstore WHERE key = $1", key)
		if err != nil {
			return fmt.Errorf("error deleting data: %w", err)
		}
		fmt.Println("Key deleted successfully.")
	}
	return nil
}

func (r *realClient) UpdatePostgressRow(key, value string) error {

	var existing string

	err := r.db.QueryRow("SELECT key FROM kvstore WHERE key = $1", key).Scan(&existing)

	if err == sql.ErrNoRows {

		return fmt.Errorf("key does not found")
	} else if err != nil {
		fmt.Println("Error while checking: ", err)
	} else {
		//Key found
		fmt.Println("Enter the value you want to update:")

		if value == "" {
			return fmt.Errorf("value can not be empty")
		} else {
			_, err = r.db.Exec("UPDATE kvstore SET value = $1 WHERE key = $2", value, key)
			if err != nil {
				return fmt.Errorf("error updating data: %w", err)
			}
			fmt.Println("Key value updated successfully.")

		}

	}
	return nil
}

func (r *realClient) GetPostgressRow(key string) (string, error) {

	var value string
	err := r.db.QueryRow("SELECT value FROM kvstore WHERE key = $1", key).Scan(&value)

	if err == sql.ErrNoRows {
		return "", fmt.Errorf("key not found")
	} else if err != nil {
		return "", fmt.Errorf("error while checking the key: %w", err)
	}


	return value, nil
}

func (r *realClient) ShowPostgressRow() (map[string]string, error) {

	store := make(map[string]string)

	rows, err := r.db.Query("SELECT key, value from kvstore")
	if err != nil {
		return nil, fmt.Errorf("error retrieving data %w", err)
	}
	defer rows.Close()

	var key, value string
	for rows.Next() {
		err := rows.Scan(&key, &value)

		if err != nil {
			fmt.Printf("error while scaning the data: %v\n", err) ////????
			continue                                              //if any point cannot scan, skip that particular row and will execute the rest.
		}
		store[key]=value

	}

	if err = rows.Err(); err != nil {

		return nil, fmt.Errorf("error iterating over rows: %w", err)

	}

	return store,nil
}

func (r *realClient) ExitPostgressRow() error {

	return nil
}
