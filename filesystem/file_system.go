package filesystem

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/imsumedhaa/In-memory-database/database"
	"github.com/spf13/afero"
)


type FileSystem struct { //find out what is necessary for file system
	FileName string
	store    map[string]string
	fs       afero.Fs      //afero.Fs is an interface defined by the Afero library  and here f.fs comes
}

func NewFileSystemWithFS(name string, fs afero.Fs) (*FileSystem, error){
	if exists, _ := afero.Exists(fs, name); !exists{
		_,err := fs.Create(name)
			if err != nil{
				return nil, fmt.Errorf("failed to create file: %w", err)
			
		}		
	}
	return &FileSystem{
		FileName: name,
		store:    make(map[string]string),
		fs:       fs,
	},nil
}


func NewFileSystem(name string) (database.Database, error) { //create some file name database.json only if it is not exist
	fs := afero.NewOsFs()

	if _,err := fs.Stat(name); err!= nil{
		if os.IsNotExist(err){
			fmt.Println("File not created yet, Creating new one....")
			_, err = fs.Create(name)
				if err != nil{
					return nil, fmt.Errorf("failed to create file: %w", err)
				}
		} else{
			return nil, fmt.Errorf("failed to get the file: %w", err)
		}
	}else{
		fmt.Println("File already exists")
	}
	return &FileSystem{
		FileName: name,
		store: make(map[string]string),
		fs : fs,
	},nil
}

func (f *FileSystem) Create(key, value string) error {

	file, err := afero.ReadFile(f.fs , f.FileName)       //f.fs means “use the file system instance (real or virtual) stored in this struct.”
	if err == nil && len(file) > 0 {
		json.Unmarshal(file, &f.store)     //
	}

	if key=="" || value==""{
		return fmt.Errorf("key and value cannot be empty")
	}
	if _, exists := f.store[key];exists{
		return fmt.Errorf("key already exists")
	}

	// Add and save
	f.store[key] = value

	updateData,err := json.MarshalIndent(f.store,"","  ")
	if err != nil{
		return fmt.Errorf("error while encoding  %w",err)
	}
	err= afero.WriteFile(f.fs, f.FileName,updateData,0644)

	if err != nil{
		return fmt.Errorf("error while writing the file: %w",err)
	}
	fmt.Println("Key value pair successfully created")

	return nil
}

func (f *FileSystem) Update(key,value string) error {
	// Step 1: Load existing data
	file, err := afero.ReadFile(f.fs, f.FileName)    //f.fs means “use the file system instance (real or virtual) stored in this struct.”
	if err == nil && len(file) > 0 {
		json.Unmarshal(file, &f.store)       //json.Unmarshal: Convert JSON ➡️ Go data
	}

	if key == "" {
		return fmt.Errorf("key cannot be empty")

	}

	// Step 3: Check if key exists
	if _, exists := f.store[key]; !exists {
		return fmt.Errorf("key not found")
	}

	if value == "" {
		return fmt.Errorf("value cannot be empty")
	}

	// Step 5: Update value in store
	f.store[key] = value

	// Step 6: Write updated data to file
	updatedData, err := json.MarshalIndent(f.store, "", "  ")       // Convert Go data ➡️  JSON with indent means space
	if err != nil {
		return fmt.Errorf("error encoding data: %w", err)
	}

	err = afero.WriteFile(f.fs, f.FileName, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	fmt.Println("Value updated successfully.")

	return nil
}

func (f *FileSystem) Delete(key string) error {

	file, err := afero.ReadFile(f.fs, f.FileName)
	if err == nil && len(file) > 0 { //check if the error is nil and the file has some data to decode
		json.Unmarshal(file, &f.store) //decode the data
	}

	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	if _, exists := f.store[key]; !exists {   ///////
		return fmt.Errorf("key not found")
	}

	delete(f.store, key)

	updatedData, err := json.MarshalIndent(f.store, "", "  ")

	if err != nil {
		return fmt.Errorf("error encoding data: %w", err)
	}

	err = afero.WriteFile(f.fs, f.FileName, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	fmt.Println("Value deleted successfully.")

	return nil

}

func (f *FileSystem) Get(key string) error {

	file, err := afero.ReadFile(f.fs, f.FileName)
	if err == nil && len(file) > 0 { //check if the error is nil and the file has some data to decode
		json.Unmarshal(file, &f.store) //decode the data
	}

	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	if _, exists := f.store[key]; !exists {   
		return fmt.Errorf("key not found")
	}

	if val, ok := f.store[key]; ok {
		fmt.Printf("Value: %s\n", val)
	} 
	return nil
}

func (f *FileSystem) Show() error {

	file, err := afero.ReadFile(f.fs, f.FileName)
	if err!= nil{
		return fmt.Errorf("error while reading from the file: %w",err)
	}
	if len(file) > 0{
		err = json.Unmarshal(file, &f.store)     //data will store from json file to map
		if err != nil{
			return fmt.Errorf("failed to decode JSON: %w", err)
		}
	}

	fmt.Println("The full map is:")
	fmt.Println(f.store)

	return nil
}

func (f *FileSystem) Exit() error {

	fmt.Println("Exiting program.")
	os.Exit(0)

	return nil
}
