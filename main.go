package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	key string
	value string
	store= make(map[string]string)  //global map declaration
)

func main(){
	

	if len(os.Args) < 2 {
		fmt.Println("Expected subcommand: create, update, delete, get")
		os.Exit(1)
	}

	cmd := os.Args[1]

	// Create a new FlagSet to parse only the flags after the subcommand
	flags := flag.NewFlagSet(cmd, flag.ExitOnError)

	flags.StringVar(&key,"key","","this is the key")
	flags.StringVar(&value,"value","","this is the value")

	flags.Parse(os.Args[2:]) // Parse args after the subcommand


	switch cmd{
	case "create":
		create(key,value)
	case "update":
		update()
	case "delete":
		delete(key)
	case "get":
		get(key)
	default:
		fmt.Println("Wrong command")
	}	
}

func create(key,value string){
	fmt.Printf("We have a key %s and value %s\n", key,value)
	if key=="" || value==""{
		fmt.Println("Key and Value required to store in the map")
	}

	store[key]=value
	fmt.Println(store)

}
func update(){
	fmt.Printf("Update value %s for %s key\n",value,key)

	if key=="" || value==""{
		fmt.Println("Key and Value required to update in the map")
	}
	store[key]= value
	fmt.Println(store)


}
func delete(key string){
	fmt.Printf("Delete key %s\n",key)

	if key==""{
		fmt.Println("Key required to delete in the map")
	}
	//delete(store,fooo)
}

func get(key string){
	fmt.Printf("Get the value for %s key\n",key)
	
}