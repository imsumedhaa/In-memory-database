package filesystem

import (
	"fmt"
	"os"
)

// struct name Inmemory
type FileSystem struct { //find out what is necessary for file system
	FileName string
	//Key      string
	//Value    string
}

func NewFileSystem(name string) (*FileSystem, error){  //create some file name database.json only if it is not exist
	if _,err:=os.Stat(name);err == nil{
		fmt.Println("File successfully")
	}else{
		if os.IsNotExist(err){
			fmt.Println("File note creates yet, creating new one...")
			_,err:= os.Create(name)
			if err!=nil{
				return nil, fmt.Errorf("failed to create file: %w", err)
			}
		}
		return nil, fmt.Errorf("failed to get file: %w", err)
	}

	return &FileSystem{
		FileName : name,
	},nil
	
} 

func (i *FileSystem) Create() {

}

func (i *FileSystem) Update() {
}

func (i *FileSystem) Delete() {
}

func (i *FileSystem) Get() {
}

func (i *FileSystem) Show() {
}

func (i *FileSystem) Exit() {
}
