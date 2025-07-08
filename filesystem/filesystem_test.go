package filesystem

import (
	"encoding/json"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)
func TestCreate(t *testing.T){
	tests := []struct{
		name          string
		key           string  
		value         string             
		initialStore  map[string]string 
		expectedError string
		expectedStore map[string]string 
	}{
		{
			name: "Create new key-value pair",
			key:  "name",
			value: "foo",
			initialStore: map[string]string{},
			expectedError: "",
			expectedStore: map[string]string{"name":"foo"},
		},
		{
			name: "Empty key",
			key:  "",
			value: "foo",
			initialStore: map[string]string{},
			expectedError: "key and value cannot be empty",
			expectedStore: map[string]string{},
		},
		{
			name: "Empty value",
			key:  "name",
			value: "",
			initialStore: map[string]string{},
			expectedError: "key and value cannot be empty",
			expectedStore: map[string]string{},
		},
		{
			name: "Duplicate key",
			key:  "name",
			value: "Max",
			initialStore: map[string]string{"name":"abc"},
			expectedError: "key already exists",
			expectedStore: map[string]string{"name":"abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			fs := afero.NewMemMapFs()     //creates a virtual(in-memory) file system using the Afero library.Does not touch the real disk.
			filename := "test.json"

			if len(tt.initialStore) > 0 {
				data, _ := json.MarshalIndent(tt.initialStore, "", "  ")    //map is converted to JSON using json.MarshalIndent
				_ = afero.WriteFile(fs, filename, data, 0644)    //data is written to test.json filr and 0644 is the file permission
			} else {
				// Ensure file is created even if empty
				fs.Create(filename)
			}
			
			// Create FileSystem instance with mocked fs
			store, err := NewFileSystemWithFS(filename, fs)
			assert.NoError(t, err)


			// Call Create
			err = store.Create(tt.key, tt.value)

			// Check error expectation
			if tt.expectedError != "" {
				assert.Error(t, err)      //make sure there is an error..returns bool
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			// Read file content
			data, _ := afero.ReadFile(fs, filename)
			var actual map[string]string
			_ = json.Unmarshal(data, &actual)

			// Check store content
			assert.Equal(t, normalizeMap(tt.expectedStore), normalizeMap(actual))
		})
	}
}
func normalizeMap(m map[string]string) map[string]string {
	if m == nil {
		return map[string]string{}
	}
	return m
}


func TestUpdate(t *testing.T){
	tests := []struct{
		name          string
		key           string  
		value         string             
		initialStore  map[string]string 
		expectedError string
		expectedStore map[string]string 
	}{
		{
			name: "Update value for the key",
			key:  "name",
			value: "abc",
			initialStore: map[string]string{"name":"abc"},
			expectedError: "",
			expectedStore: map[string]string{"name":"abc"},
		},
		{
			name: "Empty key",
			key:  "",
			value: "foo",
			initialStore: map[string]string{"name":"abc"},
			expectedError: "key cannot be empty",
			expectedStore: map[string]string{"name":"abc"},
		},
		{
			name: "Empty value",
			key:  "name",
			value: "",
			initialStore: map[string]string{"name":"abc"},
			expectedError: "value cannot be empty",
			expectedStore: map[string]string{"name":"abc"},
		},
		{
			name: "Key not there in the map",
			key:  "age",
			value: "19",
			initialStore: map[string]string{"name":"abc"},
			expectedError: "key not found",
			expectedStore: map[string]string{"name":"abc"},
		},
	}
	for _,tt := range tests{
		t.Run(tt.name , func(t *testing.T){
			fs:= afero.NewMemMapFs()
			filename := "test.json"
			
			if len(tt.initialStore) > 0{
				data,_ := json.MarshalIndent(tt.initialStore, "", "  ")
				_ =afero.WriteFile(fs, filename, data, 0644)
			}

			//create a instance of file system using NewFileSystemWithFS

	 		store, err := NewFileSystemWithFS(filename, fs)
			assert.NoError(t,err)    //Check there is no error

			err = store.Update(tt.key,tt.value)

			if tt.expectedError != ""{
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else{
				assert.NoError(t, err)
			}

			// Read file content
			data, _ := afero.ReadFile(fs, filename)
			var actual map[string]string
			_ = json.Unmarshal(data, &actual)

			// Check store content
			assert.Equal(t, normalizeMap(tt.expectedStore), normalizeMap(actual))
		})
	}
}

func TestDelete(t *testing.T){
	tests := []struct{
		name   string
		key    string
		initialStore    map[string]string
		expectedError string
		expectedStore map[string]string
	}{
		{
			name: "Delete the key value pair",
			key: "name",
			initialStore: map[string]string{"name":"abc",},
			expectedError: "",
			expectedStore: map[string]string{},
		},
		{
			name: "Empty key",
			key: "",
			initialStore: map[string]string{"name":"abc"},
			expectedError: "key cannot be empty",
			expectedStore: map[string]string{"name":"abc"},
		},
		{
			name: "Key not found",
			key: "age",
			initialStore: map[string]string{"name":"abc"},
			expectedError: "key not found",
			expectedStore: map[string]string{"name":"abc"},
		},
	}
	for _,tt := range tests{
		t.Run(tt.name , func(t *testing.T){
			fs := afero.NewMemMapFs()  //creates virtual filesystem
			filename := "test.json"

			if len(tt.initialStore)>0 {
				data,_ := json.MarshalIndent(tt.initialStore, "", "  ")
				_ = afero.WriteFile(fs, filename, data, 0644)
			}

			store, err := NewFileSystemWithFS(filename, fs)

			assert.NoError(t, err)

			err = store.Delete(tt.key)

			if tt.expectedError != ""{
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}else {
				assert.NoError(t, err)
			}
			data,_ := afero.ReadFile(fs,filename)

			var actual map[string]string
			_ = json.Unmarshal(data, &actual)

			assert.Equal(t, normalizeMap(actual), normalizeMap(tt.expectedStore))


		})
	}
}

func TestGet(t *testing.T){
	tests := []struct{
		name           string
		key            string
		initialStore   map[string]string
		expectedError  string
	}{
		{
			name:         "Get the value",
			key:          "name",
			initialStore:  map[string]string{"name":"abc"},
			expectedError: "",
		},
		{
			name:         "Empty Key",
		    key:          "",
		    initialStore:  map[string]string{"name":"abc"},
		    expectedError: "key cannot be empty",
		},
		{
			name:         "Key is not in the map",
		    key:          "age",
		    initialStore:  map[string]string{"name":"abc"},
		    expectedError: "key not found",
		},
	}
	for _, tt := range tests{
		t.Run(tt.name , func(t *testing.T){
			fs := afero.NewMemMapFs()
			filename := "test.json"

			if len(tt.initialStore)> 0{
				data,_ := json.MarshalIndent(tt.initialStore, "", "  ")
				_= afero.WriteFile(fs, filename, data, 0644)
			}
			
			store, err := NewFileSystemWithFS(filename, fs)

			assert.NoError(t, err)

			err = store.Get(tt.key)

			if tt.expectedError != ""{
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestShow(t *testing.T){
	tests := []struct{
		name           string
		initialStore   map[string]string
		expectedError  string
		expectedStore  map[string]string
	}{
		{
			name:         "Show the full map",
			initialStore:  map[string]string{"name":"abc"},
			expectedError: "",
			expectedStore:  map[string]string{"name":"abc"},
		},
	}
	for _, tt := range tests{
		t.Run(tt.name , func(t *testing.T){
			fs := afero.NewMemMapFs()
			filename := "test.json"

			if len(tt.initialStore)> 0{
				data,_ := json.MarshalIndent(tt.initialStore, "", "  ")
				_= afero.WriteFile(fs, filename, data, 0644)
			}
			
			store, err := NewFileSystemWithFS(filename, fs)

			assert.NoError(t, err)
			err = store.Show()

			if tt.expectedError != ""{
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}else {
				assert.NoError(t, err)
			}
			data,_ := afero.ReadFile(fs,filename)

			var actual map[string]string
			_ = json.Unmarshal(data, &actual)

			assert.Equal(t, normalizeMap(actual), normalizeMap(tt.expectedStore))

		})
	}
}
