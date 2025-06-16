package postgres

import (
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

	err := p.client.CreatePostgressRow(key, value)
	if err != nil {
		return fmt.Errorf("error creating new row %w", err)
	}

	fmt.Println("Data inserted successfully.")

	return nil
}

func (p *Postgres) Delete(key string) error {

	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	err := p.client.DeletePostgressRow(key)
	if err != nil {
		return fmt.Errorf("failed to create postgres row: %w", err)
	}
	return nil
}

func (p *Postgres) Update(key, value string) error {

	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	err := p.client.UpdatePostgressRow(key, value)
	if err != nil {
		return fmt.Errorf("failed to update postgres row %w", err)
	}
	return nil

}

func (p *Postgres) Get(key string) error {

	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	value,err := p.client.GetPostgressRow(key)
	if err != nil {
		return fmt.Errorf("failed to get postgres row: %w", err)
	}
	// If we reach here, the key exists and have the value
	fmt.Printf("Value for key '%s': %s\n", key, value)

	return nil
}

func (p *Postgres) Show() error {

	store,err := p.client.ShowPostgressRow()
	if err != nil {
		return fmt.Errorf("failed to show postgres row %w", err)
	}
	//if reach here means no error, can print the key value
	fmt.Println("Map is", store)

	return nil
}

func (p *Postgres) Exit() error {
	fmt.Println("Exiting program...")
	os.Exit(0)
	return nil
}
