package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//struct name Inmemory

type FileSystem struct{       //find out what is necessary for file system
	Data map[string]string `json:"store"`
	Key string
	Value string
	
}

func NewFileSystem() *FileSystem{       //create some file name database.json only if it is not exist

	if _,err:= os.Stat("jsonfile.txt");err==nil{
		fmt.Println("File exists")
	}else{
		os.Create("database.json")
	}
	
	return &FileSystem{
		Data:  make(map[string]string),
			
	}
}
func (f *FileSystem)Create(){
	fmt.Println("Enter the key:")
	fmt.Println("Enter the key:")
	store.Data[key]=value

	jsonData, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

}
