package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/imsumedhaa/In-memory-database/api"
	"github.com/imsumedhaa/In-memory-database/database"
	"github.com/imsumedhaa/In-memory-database/filesystem"
	"github.com/imsumedhaa/In-memory-database/inmemory"
	"github.com/imsumedhaa/In-memory-database/postgres"
)

var (
	name string
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected subcommand 'filesystem' or 'inmemory' or 'postgres'")
		os.Exit(1)
	}

	cmd := os.Args[1]

	flags := flag.NewFlagSet(cmd, flag.ExitOnError)
	flags.StringVar(&name, "name", "database.json", "this is the name")
	flags.Parse(os.Args[2:]) // Parse args after the subcommand

	var operation database.Database = nil
	var err error

	switch cmd {
	case "filesystem":
		operation, err = filesystem.NewFileSystem(name)
		if err != nil {
			fmt.Printf("Error creating filesystem: %v\n", err)
			os.Exit(1)
		}

	case "inmemory":
		operation, err = inmemory.NewInmemory()
		if err != nil {
			fmt.Printf("Error creating the map: %v", err)
			os.Exit(1)
		}

	case "postgres":

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

		operation, err = postgres.NewPostgres(port, username, password, dbname)
		if err != nil {
			fmt.Printf("Error creating the connection: %v\n", err)
			os.Exit(1)
		}

	case "server":

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

		httpConfig, err := api.NewHttp(port, username, password, dbname)

		if err != nil {
			fmt.Printf("Error creating the http connection: %v\n", err)
			os.Exit(1)
		}

		if err := httpConfig.Run(); err != nil {
			fmt.Printf("Error run http server: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Println("Wrong Command. Should be either 'filesystem' or 'inmemory' or 'postgres'")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter subcommand: create,update,get,delete,show & exit to quit the program")
		input, _ := reader.ReadString('\n') //to read the input from user and store into input var
		command := strings.TrimSpace(input) //delete the "\n" from input var and store it in command var

		switch command {
		case "create":
			fmt.Println("Enter the key:")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)

			fmt.Println("Enter the value:")
			value, _ := reader.ReadString('\n')
			value = strings.TrimSpace(value)

			err := operation.Create(key, value)
			if err != nil {
				fmt.Printf("Error while creating the map: %v\n", err)
				fmt.Println("try using another key...")
			}

		case "get":
			fmt.Println("Enter the key:")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)

			err := operation.Get(key)
			if err != nil {
				fmt.Printf("Error while getting the value: %v\n", err)
				os.Exit(1)
			}

		case "update":
			fmt.Println("Enter the key:")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)

			fmt.Println("Enter the value:")
			value, _ := reader.ReadString('\n')
			value = strings.TrimSpace(value)

			err := operation.Update(key, value)
			if err != nil {
				fmt.Printf("Error while updating the value: %v\n", err)
				os.Exit(1)
			}

		case "delete":
			fmt.Println("Enter the key you want to delete:")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)

			err := operation.Delete(key)
			if err != nil {
				fmt.Printf("Error while Deleteing the pair: %v\n", err)
				os.Exit(1)
			}

		case "show":
			err := operation.Show()
			if err != nil {
				fmt.Printf("Error while showing the map: %v\n", err)
				os.Exit(1)
			}

		case "exit":
			err := operation.Exit()

			if err != nil {
				fmt.Printf("Error while exiting the program: %v\n", err)
				os.Exit(1)
			}

		default:
			fmt.Println("Wrong Command.")
		}
	}
}
