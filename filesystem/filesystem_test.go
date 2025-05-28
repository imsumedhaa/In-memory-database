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
			// Create a virtual file system
			fs := afero.NewMemMapFs()
			filename := "test.json"

			// Initialize file with initialStore if provided
			if tt.initialStore != nil {
				data, _ := json.MarshalIndent(tt.initialStore, "", "  ")
				_ = afero.WriteFile(fs, filename, data, 0644)
			} else {
				_, _ = fs.Create(filename)
			}

			// Initialize FileSystem with virtual FS
			fsdb, err := NewFileSystemWithFS(filename, fs)
			if err != nil {
				t.Fatalf("setup failed: %v", err)
			}

			// Call the method under test
			err = fsdb.CreateWithInput(tt.key, tt.value)

			// Check for expected error message
			if tt.expectedError != "" {
				if err == nil || err.Error() != tt.expectedError {
					t.Errorf("expected error %q, got %v", tt.expectedError, err)
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			// Check final contents of the store
			data, _ := afero.ReadFile(fs, filename)
			var got map[string]string
			_ = json.Unmarshal(data, &got)

			if len(got) != len(tt.expectedStore) {
				t.Errorf("expected store length %d, got %d", len(tt.expectedStore), len(got))
			}

			for k, v := range tt.expectedStore {
				if got[k] != v {
					t.Errorf("expected key %q to have value %q, got %q", k, v, got[k])
				}
			}
		})
	}
}

