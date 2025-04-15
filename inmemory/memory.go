package inmemory

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/imsumedhaa/In-memory-database/database"
)

//struct name Inmemory

type Inmemory struct{
	store map[string]string
	reader *bufio.Reader
}
//Constructor -> A function which returns a pointer to the struct Inmemory

func NewInmemory() database.Database{
	return &Inmemory{
		store:  make(map[string]string),
		reader: bufio.NewReader(os.Stdin),
	}
}
//struct Receiver

func (i *Inmemory)Create(){
	fmt.Println("Enter the key:")
	key, _ := i.reader.ReadString('\n')
	key = strings.TrimSpace(key)

	fmt.Println("Enter the value:")
	val,_:=i.reader.ReadString('\n')
	val = strings.TrimSpace(val)

	if key=="" || val==""{                               
		fmt.Println("Require the key and value.")
		os.Exit(0)
	}

	if existValue, exists := i.store[key]; exists {
		fmt.Println("Key already exists. Use 'update' to change the value.")               
		fmt.Println(existValue)
	}else {
		i.store[key] = val
		fmt.Println("Created successfully.")
	}
}

func (i *Inmemory)Get(){
	fmt.Println("Enter the key:")
	key,_:= i.reader.ReadString('\n')
	key= strings.TrimSpace(key)

	if key==""{
		fmt.Println("Enter the key properly.")
		os.Exit(0)
	}

	if val,ok:=i.store[key];ok{
		fmt.Printf("Value is %s\n",val)
	}else{
		fmt.Println("key not found")
		}
}

func (i *Inmemory)Update(){
	fmt.Println("Enter the key:")
	key,_:= i.reader.ReadString('\n')
	key = strings.TrimSpace(key)

	if key==""{
		fmt.Println("require the key.")
		os.Exit(0)
	}

	if _,ok:= i.store[key];ok{
		fmt.Println("Enter new value:")
		value,_:= i.reader.ReadString('\n')
		value = strings.TrimSpace(value)
		i.store[key] = value
		}else{
		fmt.Println("Key not found.")
	}
}

func (i *Inmemory)Delete(){
	fmt.Println("Enter the key you want to delete:")
	key,_:= i.reader.ReadString('\n')
	key=strings.TrimSpace(key)

	if key==""{
		fmt.Println("require the key.")
		}else{
			delete(i.store,key)
			fmt.Println("Succesfully deleted")
		}
}

func (i *Inmemory)Show(){
	fmt.Println("The full map is:")
	fmt.Println(i.store)
}

func (i *Inmemory)Exit(){
	fmt.Println("Exiting program.")
    os.Exit(0)
}