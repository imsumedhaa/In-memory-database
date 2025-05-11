package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

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
		fmt.Println("Expected subcommand 'filesystem'")
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
		if err!= nil{
			fmt.Printf("Error creating the map: %v", err)
			os.Exit(1)
		}


	case "postgres":
		operation,err = postgres.NewPostgres()
		if err != nil{
			fmt.Printf("Error creating the connection: %v\n", err)
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
			operation.Create()

		case "get":
			operation.Get()

		case "update":
			operation.Update()

		case "delete":
			operation.Delete()

		case "show":
			operation.Show()

		case "exit":
			operation.Exit()

		default:
			fmt.Println("Wrong Command.")
		}
	}
}