package inmemory

import (
	"bufio"
	"fmt"
	"os"
	

	"github.com/imsumedhaa/In-memory-database/database"
)

//struct name Inmemory

type Inmemory struct {
	store  map[string]string
	reader *bufio.Reader
}

//Constructor -> A function which returns a pointer to the struct Inmemory

func NewInmemory() (database.Database, error) {
	return &Inmemory{
		store:  make(map[string]string),
		reader: bufio.NewReader(os.Stdin),
	}, nil
}

//struct Receiver

func (i *Inmemory) Create(key, value string) error {

	if key == "" || value == "" {
		return fmt.Errorf("require the key and value")
	}

	if existValue, exists := i.store[key]; exists {
		fmt.Println("Key already exists. Use 'update' to change the value.")
		fmt.Println(existValue)
	} else {
		i.store[key] = value
		fmt.Println("Created successfully.")
	}
	return nil
}

func (i *Inmemory) Get(key string) error {

	if key == "" {
		return fmt.Errorf("require the key")
	}

	if val, ok := i.store[key]; ok {
		fmt.Printf("Value: %s\n", val)
	} else {
		return fmt.Errorf("key not found")
	}
	return nil
}


func (i *Inmemory) Update(key, value string) error {
	
	if key == "" {
		return fmt.Errorf("require the key")
	}
	if value == "" {
		return fmt.Errorf("require the value")
	}
	
	if _, ok := i.store[key]; ok {
		i.store[key] = value
	} else {
		return fmt.Errorf("key not found")
	}
	return nil
}


func (i *Inmemory) Delete(key string) error {
	

	if key == "" {
		return fmt.Errorf("require the key")
	}
	if _, ok := i.store[key]; ok {
		delete(i.store, key)
		fmt.Println("Succesfully deleted")
	} else {
		return fmt.Errorf("key not found")
	}

	return nil
}

func (i *Inmemory) Show() error {
	fmt.Println("The full map is:")
	fmt.Println(i.store)

	return nil
}

func (i *Inmemory) Exit() error {
	fmt.Println("Exiting program.")
	os.Exit(0)
	return nil
}
