package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	key string
	value string	
)

func main(){
	store:= make(map[string]string)  //global map declaration
	reader := bufio.NewReader(os.Stdin)     //creates a new buffered reader that reads input from the terminal
	
	for{
		fmt.Println("Enter subcommand: create,update,get,delete,show & exit to quit the program")
		input, _ := reader.ReadString('\n')     //to read the input from user and store into input var
		command := strings.TrimSpace(input)     //delete the "\n" from input var and store it in command var

		switch command{
		case "create":
			fmt.Println("Enter the key:")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)

			fmt.Println("Enter the value:")
			val,_:=reader.ReadString('\n')
			val = strings.TrimSpace(val)

			if key=="" || val==""{                               
				fmt.Println("Require the key and value.")
				os.Exit(0)
			}

			if existValue, exists := store[key]; exists {
				fmt.Println("Key already exists. Use 'update' to change the value.")               ///
				fmt.Println(existValue)
			} else {
				store[key] = val
				fmt.Println("Created successfully.")
			}

		case "get":
			fmt.Println("Enter the key:")
			key,_:= reader.ReadString('\n')
			key= strings.TrimSpace(key)

			if key==""{
				fmt.Println("Enter the key properly.")
				os.Exit(0)
			}

			if val,ok:=store[key];ok{
				fmt.Printf("Value is %s\n",val)
			}else{
				fmt.Println("key not found")
			}

		case "update":
			fmt.Println("Enter the key:")
			key,_:= reader.ReadString('\n')
			key = strings.TrimSpace(key)

			if key==""{
				fmt.Println("require the key.")
				os.Exit(0)
			}

			if _,ok:=store[key];ok{
				fmt.Println("Enter new value:")
				value,_:=reader.ReadString('\n')
				value= strings.TrimSpace(value)
				store[key]=value
			}else{
				fmt.Println("Key not found.")
			}
		
		case "delete":
			fmt.Println("Enter the key you want to delete:")
			key,_:=reader.ReadString('\n')
			key=strings.TrimSpace(key)

			if key==""{
				fmt.Println("require the key.")
			}else{
				delete(store,key)
			 	fmt.Println("Succesfully deleted")
			}
		
		case "show":
			fmt.Println("The full map is:")
			fmt.Println(store)
		
		case "exit":
			fmt.Println("Exiting program.")
            os.Exit(0)

		default:
			fmt.Println("Wrong Command.")
		}
	}
}


