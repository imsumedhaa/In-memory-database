package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/imsumedhaa/In-memory-database/database"
	"github.com/imsumedhaa/In-memory-database/pkg/client/postgres"
	_ "github.com/lib/pq"
)

type Postgres struct {
	client postgres.Client
}

func NewPostgres(port, username, password, dbname string) (database.Database, error) {

	dbClient, err := postgres.NewClient(port, username, password, dbname)

	if err != nil {
		return nil, fmt.Errorf("failed to connect %w", err)
	}
	return &Postgres{client: dbClient}, nil
}

func (p *Postgres) Create(key, value string) error {

	if key == "" || value == "" {
		return fmt.Errorf("key and value cannot be empty")
	}

	err := CreatePostgressRow(key, value)
	if err != nil {

	}

	fmt.Println("Data inserted successfully.")

	return nil
}

func (p *Postgres) Delete(key string) error {

	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	var existing string

	err := p.db.QueryRow("SELECT key FROM kvstore WHERE key = $1", key).Scan(&existing)

	if err == sql.ErrNoRows {
		return fmt.Errorf("key does not exist")

	} else if err != nil {
		fmt.Println("Error is ", err)
	} else {
		// Delete new key-value pair
		_, err = p.db.Exec("DELETE FROM kvstore WHERE key = $1", key)
		if err != nil {
			return fmt.Errorf("error deleting data: %w", err)
		}
		fmt.Println("Key deleted successfully.")
	}
	return nil
}

func (p *Postgres) Update(key, value string) error {

	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	var existing string

	err := p.db.QueryRow("SELECT key FROM kvstore WHERE key = $1", key).Scan(&existing)

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
			_, err = p.db.Exec("UPDATE kvstore SET value = $1 WHERE key = $2", value, key)
			if err != nil {
				return fmt.Errorf("error updating data: %w", err)
			}
			fmt.Println("Key value updated successfully.")

		}

	}
	return nil

}

func (p *Postgres) Get(key string) error {

	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	var value string
	err := p.db.QueryRow("SELECT value FROM kvstore WHERE key = $1", key).Scan(&value)

	if err == sql.ErrNoRows {
		return fmt.Errorf("key not found")
	} else if err != nil {
		return fmt.Errorf("error while checking the key: %w", err)
	}

	// If we reach here, the key exists and have the value
	fmt.Printf("Value for key '%s': %s\n", key, value)

	return nil
}

func (p *Postgres) Show() error {
	rows, err := p.db.Query("SELECT key, value from kvstore")
	if err != nil {
		return fmt.Errorf("error retrieving data %w", err)
	}
	defer rows.Close()

	var key, value string
	for rows.Next() {
		err := rows.Scan(&key, &value)

		if err != nil {
			fmt.Printf("error while scaning the data: %v\n", err) ////????
			continue                                              //if any point cannot scan, skip that particular row and will execute the rest.
		}
		//if reach here means no error, can print the key value
		fmt.Printf("Key : %s, Value : %s\n", key, value)
	}

	if err = rows.Err(); err != nil {

		return fmt.Errorf("error iterating over rows: %w", err)

	}
	return nil
}

func (p *Postgres) Exit() error {
	fmt.Println("Exiting program...")
	os.Exit(0)
	return nil
}
