package filesystem

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/imsumedhaa/In-memory-database/database"
)

// struct name Inmemory
type FileSystem struct { //find out what is necessary for file system
	FileName string
	store map[string]string
	//Key      string
	//Value    string
}

func NewFileSystem(name string) (database.Database, error){  //create some file name database.json only if it is not exist
	if _,err:=os.Stat(name);err == nil{
		fmt.Println("File created successfully")
	}else{
		if os.IsNotExist(err){
			fmt.Println("File not created yet, creating new one...")
			_,err:= os.Create(name)
			if err!=nil{
				return nil, fmt.Errorf("failed to create file: %w", err)
			}
		}else{
			return nil, fmt.Errorf("failed to get file: %w", err)	
		}
	}

	return &FileSystem{
		FileName : name,
		store : make(map[string]string),
		
	},nil
	
} 

func (f *FileSystem) Create() error{
	file, err := os.ReadFile(f.FileName)
	if err == nil && len(file) > 0 {
		json.Unmarshal(file, &f.store)     //
	}
	reader := bufio.NewReader(os.Stdin)    //for input the data from user

	
	fmt.Println("Enter the key:")
	key, _ := reader.ReadString('\n')      //read input from user and stored in key variable
	key = strings.TrimSpace(key)           //trim the space which is added next to the key

	fmt.Println("Enter the value:")       
	value,_:= reader.ReadString('\n')
	value = strings.TrimSpace(value)

	if key=="" || value==""{
		fmt.Println("Key and value cannot be empty")
		os.Exit(1)
	}
	if _, exists := f.store[key];exists{
		fmt.Println("key already exists")
		os.Exit(1)
	}

	// Add and save
	f.store[key] = value

	updateData,err:=json.MarshalIndent(f.store,"","  ")
	if err!= nil{
		return fmt.Errorf("error while encoding  %w",err)
	}
	err= os.WriteFile(f.FileName,updateData,0644)
	if err!=nil{
		return fmt.Errorf("Error while writing the file: %w\n",err)
	}
	fmt.Println("Key value pair successfully created")

	return nil
}

func (f *FileSystem) Update() error {
	// Step 1: Load existing data
	file, err := os.ReadFile(f.FileName)
	if err == nil && len(file) > 0 {
		json.Unmarshal(file, &f.store)
	}

	// Step 2: Ask for key
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the key to update:")
	key, _ := reader.ReadString('\n')
	key = strings.TrimSpace(key)

	if key == "" {
		return fmt.Errorf("Key cannot be empty.")

	}

	// Step 3: Check if key exists
	if _, exists := f.store[key]; !exists {
		return fmt.Errorf("Key not found.")
	}

	// Step 4: Get new value
	fmt.Println("Enter the new value:")
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)

	if value == "" {
		return fmt.Errorf("Value cannot be empty.")
	}

	// Step 5: Update value in store
	f.store[key] = value

	// Step 6: Write updated data to file
	updatedData, err := json.MarshalIndent(f.store, "", "  ")
	if err != nil {
		return fmt.Errorf("Error encoding data:", err)
	}

	err = os.WriteFile(f.FileName, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %w", err)
	}

	fmt.Println("Value updated successfully.")

	return nil
}


func (f *FileSystem) Delete() error{


	file, err := os.ReadFile(f.FileName)
	if err == nil && len(file) > 0 {     //check if the error is nil and the file has some data to decode
		json.Unmarshal(file, &f.store)    //decode the data
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the key you want to delete:")
	key,_:= reader.ReadString('\n')
	key= strings.TrimSpace(key)

	if key==""{
		return fmt.Errorf("Key value cannot be empty...")
	}

	if _, exists := f.store[key]; exists {
		return fmt.Errorf("Key not found.")
	}

	delete(f.store,key)

	updatedData, err := json.MarshalIndent(f.store, "", "  ")

	if err!=nil{
		return fmt.Errorf("Error encoding data: %w",err)
	}

	
	err = os.WriteFile(f.FileName, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %w", err)
	}

	fmt.Println("Value updated successfully.")

	return nil

}

func (f *FileSystem) Get() error {

	file, err := os.ReadFile(f.FileName)
	if err == nil && len(file) > 0 {     //check if the error is nil and the file has some data to decode
		json.Unmarshal(file, &f.store)    //decode the data
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the key you want to get:")
	key,_:= reader.ReadString('\n')
	key = strings.TrimSpace(key)

	if key==""{
		return fmt.Errorf("Key cannot be empty...")
	}

	if val,ok:=f.store[key];ok{
		fmt.Printf("Value: %s\n", val)
	}else{
		fmt.Println("key not found")
		}
		return nil
}

func (f *FileSystem) Show() error{

	fmt.Println("The full map is:")
	fmt.Println(f.store)

	return nil
}

func (f *FileSystem) Exit() error{

	fmt.Println("Exiting program.")
    os.Exit(0)

    return nil
}
