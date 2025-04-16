package main

import (
	"bufio"
	"fmt"
	"os"
	"flag"
	

	"github.com/imsumedhaa/In-memory-database/inmemory"
)

var (
	name string
	
)

func main(){
	mem := inmemory.NewInmemory()

	//store:= make(map[string]string)  //global map declaration
	//reader := bufio.NewReader(os.Stdin)     //creates a new buffered reader that reads input from the terminal


	if len(os.Args)<2{
		fmt.Println("Expected subcommand 'filesystem'")
		os.Exit(1)
	}

	cmd:= os.Args[1]
	
	flags := flag.NewFlagSet(cmd, flag.ExitOnError)

	flags.StringVar(&name,"name","database.json","this is the name")
 
	flags.Parse(os.Args[2:]) // Parse args after the subcommand


		switch cmd{
		case "filesystem":


		case "inmemory":
			mem := inmemory.NewInmemory()

	store:= make(map[string]string)  //global map declaration
	reader := bufio.NewReader(os.Stdin)     //creates a new buffered reader that reads input from the terminal


					

		default:
			fmt.Println("Wrong Command.")
		}
	}



