package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/imsumedhaa/In-memory-database/inmemory"
)

var (
	key string
	value string	
)

func main(){
	//store:= make(map[string]string)  //global map declaration
	reader := bufio.NewReader(os.Stdin)     //creates a new buffered reader that reads input from the terminal
	
	for{
		fmt.Println("Enter subcommand: create,update,get,delete,show & exit to quit the program")
		input, _ := reader.ReadString('\n')     //to read the input from user and store into input var
		command := strings.TrimSpace(input)     //delete the "\n" from input var and store it in command var

		switch command{
		case "create":
			inmemory.Create()

		case "get":
			inmemory.Get()

		case "update":
			inmemory.Update()
		
		case "delete":
			inmemory.Delete()
		
		case "show":
			inmemory.Show()
		
		case "exit":
			inmemory.Exit()

		default:
			fmt.Println("Wrong Command.")
		}
	}
}


