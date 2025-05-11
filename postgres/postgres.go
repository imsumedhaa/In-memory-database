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

func (p *Postgres) Create() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the key:")
	key, _ := reader.ReadString('\n')
	key = strings.TrimSpace(key)

	fmt.Println("Enter the value:")
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)

	if key == "" || value == "" {
		return fmt.Errorf("Key and value cannot be empty.")
	}

	// Check if key already exists
	var existing string
	err := p.db.QueryRow("SELECT key FROM kvstore WHERE key = $1", key).Scan(&existing)
	if err == nil {
		return fmt.Errorf("Key already exists. Use 'update' to change the value.")
	}

	// Insert new key-value pair
	_, err = p.db.Exec("INSERT INTO kvstore (key, value) VALUES ($1, $2)", key, value)
	if err != nil {
		return fmt.Errorf("Error inserting data:", err)
	}

	fmt.Println("Data inserted successfully.")

	return nil
}

func (p *Postgres) Delete() error {

	reader := bufio.NewReader((os.Stdin))

	fmt.Println("Enter the key you want to delete:")
	key,_ := reader.ReadString('\n')
	key = strings.TrimSpace(key)

	if key == "" {
		return fmt.Errorf("Key cannot be empty.")
	}

	var existing string

	err := p.db.QueryRow("SELECT key FROM kvstore WHERE key = $1", key).Scan(&existing)

	if err == sql.ErrNoRows{
		return fmt.Errorf("Key does not exists.")

	}else if err !=  nil{
		fmt.Println("Error is ", err)
	}else{
		// Delete new key-value pair
	_, err = p.db.Exec("DELETE FROM kvstore WHERE key = $1", key, )
	if err != nil {
		return fmt.Errorf("Error deleting data:", err)
	}
	fmt.Println("Key deleted successfully.")
	}
	return nil
}

func (p *Postgres) Update() error {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the key to update the value :")
	key,_:=reader.ReadString('\n')
	key= strings.TrimSpace(key)

	if key == ""{
		return fmt.Errorf("Key cannot be empty.")
	}
	var existing string

	err := p.db.QueryRow("SELECT key FROM kvstore WHERE key = $1", key).Scan(&existing)

	if err == sql.ErrNoRows{
		
		return fmt.Errorf("Key does not found..")
	}else if err != nil{
		fmt.Println("Error while checking: ",err)
	}else{
		//Key found
		fmt.Println("Enter the value you want to update:")

		newVal,_ := reader.ReadString('\n')
		newVal = strings.TrimSpace(newVal)

		if newVal==""{
			return fmt.Errorf("Value can not be empty...")
		}else{
			_, err = p.db.Exec("UPDATE kvstore SET value = $1 WHERE key = $2", newVal, key)
			if err != nil {
				return fmt.Errorf("Error updating data:", err)
			}
			fmt.Println("Key value updated successfully.")

		}		
	
	}
	return nil

}

func (p *Postgres) Get() error {

	reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter the key:")
    key, _ := reader.ReadString('\n')
    key = strings.TrimSpace(key)

    if key == "" {
        return fmt.Errorf("Key cannot be empty")
    }

    var value string
    err := p.db.QueryRow("SELECT value FROM kvstore WHERE key = $1", key).Scan(&value)

    if err == sql.ErrNoRows {
        return fmt.Errorf("Key not found.")
    } else if err != nil {
        return fmt.Errorf("Error while checking the key:", err)
    }

    // If we reach here, the key exists and have the value
    fmt.Printf("Value for key '%s': %s\n", key, value)

	return nil
}


func (p *Postgres) Show() error {
	rows,err := p.db.Query("SELECT key, value from kvstore")
	if err != nil{
		return fmt.Errorf("Erroe retrieving data %w",err)
	}
	defer rows.Close()

	var key, value string 
	for rows.Next(){
		err := rows.Scan(&key, &value)

		if err != nil{
			fmt.Errorf("Error while scaning the data: %w", err)
			continue  //if any point cannot scan, skip that particular row and will execute the rest.
		}
		//if reach here means no error, can print the key value
		fmt.Printf("Key : %s, Value : %s\n", key, value)
	}

	if err = rows.Err(); err != nil {

        return fmt.Errorf("Error iterating over rows:", err)

    }
	return nil
}

func (p *Postgres) Exit() error{
	fmt.Println("Exiting program...")
	os.Exit(0)
	return nil
}

