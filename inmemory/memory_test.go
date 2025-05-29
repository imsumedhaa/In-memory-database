package inmemory

import (
	"bufio"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		value         string
		initialStore  map[string]string // initial store state
		expectedError string            // "" means expect no error
		expectedStore map[string]string // expected final store state
	}{
		{
			name:          "Create new key-value pair",
			key:           "name",
			value:         "abc",
			initialStore:  map[string]string{},
			expectedError: "",
			expectedStore: map[string]string{"name": "abc"},
		},
		{
			name:          "Empty key",
			key:           "",
			value:         "foo",
			initialStore:  map[string]string{},
			expectedError: "require the key and value",
			expectedStore: map[string]string{},
		},
		{
			name:          "Empty value",
			key:           "name",
			value:         "",
			initialStore:  map[string]string{},
			expectedError: "require the key and value",
			expectedStore: map[string]string{},
		},
		{
			name:          "Empty key and value",
			key:           "",
			value:         "",
			initialStore:  map[string]string{},
			expectedError: "require the key and value",
			expectedStore: map[string]string{},
		},
		{
			name:          "Duplicate key does not overwrite",
			key:           "name",
			value:         "abc",
			initialStore:  map[string]string{"name": "Alice"},
			expectedError: "", // no error returned, but duplicate is not overwritten
			expectedStore: map[string]string{"name": "Alice"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inmem := &Inmemory{ //inmemory struct
				reader: bufio.NewReader(strings.NewReader(tt.key + "\n" + tt.value)),
				store:  copyMap(tt.initialStore),
			}

			err := inmem.Create(tt.key, tt.value)

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

func TestUpdate(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		value         string
		initialStore  map[string]string // initial store state
		expectedError string            // "" means expect no error
		expectedStore map[string]string // expected final store state
	}{
		{
			name:          "Update value",
			key:           "name",
			value:         "abc",
			initialStore:  map[string]string{"name": "Alice"},
			expectedError: "", //No error
			expectedStore: map[string]string{"name": "abc"},
		},
		{
			name:          "Empty key",
			key:           "",
			value:         "abc",
			initialStore:  map[string]string{"name": "Alice"},
			expectedError: "require the key",
			expectedStore: map[string]string{"name": "Alice"},
		},
		{
			name:          "Empty value",
			key:           "name",
			value:         "",
			initialStore:  map[string]string{"name": "Alice"},
			expectedError: "require the value",
			expectedStore: map[string]string{"name": "Alice"},
		},
		{
			name:          "Key is not there in the map",
			key:           "age",
			value:         "19",
			initialStore:  map[string]string{"name": "Alice"},
			expectedError: "key not found",
			expectedStore: map[string]string{"name": "Alice"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inmem := &Inmemory{
				reader: bufio.NewReader(strings.NewReader(tt.key + "\n" + tt.value)),
				store:  copyMap(tt.initialStore),
			}

			err := inmem.Update(tt.key, tt.value)

			if tt.expectedError == "" {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			} else {
				if err == nil || err.Error() != tt.expectedError {
					t.Errorf("expected error '%s', got '%v'", tt.expectedError, err)
				}
			}

			if !mapsEqual(inmem.store, tt.expectedStore) {
				t.Errorf("expected store: %v, got: %v", tt.expectedStore, inmem.store)
			}
		})
	}
}
func TestGet(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		initialStore  map[string]string
		expectedError string
	}{
		{
			name:          "Get value",
			key:           "name",
			initialStore:  map[string]string{"name": "Alice"},
			expectedError: "", //No error
		},
		{
			name:          "Empty key",
			key:           "",
			initialStore:  map[string]string{"name": "Alice"},
			expectedError: "require the key",
		},
		{
			name:          "Key not found",
			key:           "age",
			initialStore:  map[string]string{"name": "Alice"},
			expectedError: "key not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inmem := &Inmemory{
				reader: bufio.NewReader(strings.NewReader(tt.key)),
				store:  copyMap(tt.initialStore),
			}

			err := inmem.Get(tt.key)

			if tt.expectedError == "" {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			} else {
				if err == nil || err.Error() != tt.expectedError {
					t.Errorf("expected error '%s', got '%v'", tt.expectedError, err)
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		initialStore  map[string]string
		expectedError string
	}{
		{
			name:          "Delete key value pair",
			key:           "name",
			initialStore:  map[string]string{"name": "Alice"},
			expectedError: "", //No error
		},
		{
			name:          "Empty key",
			key:           "",
			initialStore:  map[string]string{"name": "Alice"},
			expectedError: "require the key",
		},
		{
			name:          "Key not found",
			key:           "age\n\n",
			initialStore:  map[string]string{"name": "Alice"},
			expectedError: "key not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inmem := &Inmemory{
				reader: bufio.NewReader(strings.NewReader(tt.key)),
				store:  copyMap(tt.initialStore),
			}

			err := inmem.Get(tt.key)

			if tt.expectedError == "" {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			} else {
				if err == nil || err.Error() != tt.expectedError {
					t.Errorf("expected error '%s', got '%v'", tt.expectedError, err)
				}
			}
		})
	}
}

func TestShow(t *testing.T) {
	tests := []struct {
		name          string
		initialStore  map[string]string // initial store state
		expectedError string            // "" means expect no error
		expectedStore map[string]string // expected final store state
	}{
		{
			name:          "Show the map",
			initialStore:  map[string]string{"name": "Alice"},
			expectedError: "", //No error
			expectedStore: map[string]string{"name": "Alice"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inmem := &Inmemory{
				store: copyMap(tt.initialStore),
			}

			err := inmem.Show()

			if tt.expectedError == "" {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			} else {
				if err == nil || err.Error() != tt.expectedError {
					t.Errorf("expected error '%s', got '%v'", tt.expectedError, err)
				}
			}

			if !mapsEqual(inmem.store, tt.expectedStore) {
				t.Errorf("expected store: %v, got: %v", tt.expectedStore, inmem.store)
			}
		})
	}
}
