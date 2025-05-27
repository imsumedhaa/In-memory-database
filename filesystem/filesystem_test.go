package filesystem

import(
	"testing"
	"encoding/json"
	"github.com/spf13/afero"

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
		t.Run(tt.name, func(t *testing.T) {
		
        }
}
