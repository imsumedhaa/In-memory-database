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
		return fmt.Errorf("error while deleting pairs %w", err)
	}
	return nil
}

func (p *Postgres) Update(key, value string) error {

	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	err := p.client.UpdatePostgressRow(key, value)
	if err != nil {
		return fmt.Errorf("error while updating the value %w", err)
	}
	return nil

}

func (p *Postgres) Get(key string) error {

	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	err := p.client.GetPostgressRow(key)
	if err != nil {
		return fmt.Errorf("error while getting the value %w", err)
	}

	return nil
}

func (p *Postgres) Show() error {

	err := p.client.ShowPostgressRow()
	if err != nil {
		return fmt.Errorf("error while showing %w", err)
	}

	return nil
}

func (p *Postgres) Exit() error {
	fmt.Println("Exiting program...")
	os.Exit(0)
	return nil
}
