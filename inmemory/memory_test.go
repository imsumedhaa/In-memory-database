package inmemory

import (
	"bufio"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name          string
		input         string            // simulated user input: key\nvalue\n
		initialStore  map[string]string // initial store state
		expectedError string            // "" means expect no error
		expectedStore map[string]string // expected final store state
	}{
		{
			name:          "Create new key-value pair",
			input:         "name\nAlice\n",
			initialStore:  map[string]string{},
			expectedError: "",
			expectedStore: map[string]string{"name": "Alice"},
		},
		{
			name:          "Empty key",
			input:         "\nAlice\n",
			initialStore:  map[string]string{},
			expectedError: "require the key and value",
			expectedStore: map[string]string{},
		},
		{
			name:          "Empty value",
			input:         "name\n\n",
			initialStore:  map[string]string{},
			expectedError: "require the key and value",
			expectedStore: map[string]string{},
		},
		{
			name:          "Empty key and value",
			input:         "\n\n\n",
			initialStore:  map[string]string{},
			expectedError: "require the key and value",
			expectedStore: map[string]string{},
		},
		{
			name:          "Duplicate key does not overwrite",
			input:         "name\nBob\n",
			initialStore:  map[string]string{"name": "Alice"},
			expectedError: "", // no error returned, but duplicate is not overwritten
			expectedStore: map[string]string{"name": "Alice"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inmem := &Inmemory{
				reader: bufio.NewReader(strings.NewReader(tt.input)),
				store:  copyMap(tt.initialStore),
			}

			err := inmem.Create()

			// Check error
			if tt.expectedError == "" {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			} else {
				if err == nil || err.Error() != tt.expectedError {
					t.Errorf("expected error '%s', got '%v'", tt.expectedError, err)
				}
			}

			// Check store state
			if !mapsEqual(inmem.store, tt.expectedStore) {
				t.Errorf("expected store: %v, got: %v", tt.expectedStore, inmem.store)
			}
		})
	}
}

// Helper: deep copy a map
func copyMap(m map[string]string) map[string]string {
	c := make(map[string]string)
	for k, v := range m {
		c[k] = v
	}
	return c
}

// Helper: compare two maps
func mapsEqual(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

