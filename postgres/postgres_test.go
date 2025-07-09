package postgres

import (
	"errors"
	"testing"

	"github.com/imsumedhaa/In-memory-database/pkg/client/postgres/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPostgres_Create(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		value         string
		mockFunc      func(m *mocks.Client)
		expectedError string
	}{
		{
			name:          "Empty Key",
			key:           "",
			value:         "World",
			mockFunc:      func(m *mocks.Client) {},
			expectedError: "key cannot be empty",
		},
		{
			name:  "Create Failure",
			key:   "Hello",
			value: "World",
			mockFunc: func(m *mocks.Client) {
				m.On("CreatePostgresRow", "Hello", "World").Return(errors.New("db error")).Times(1)
			},
			expectedError: "failed to create postgres row: db error",
		},
		{
			name:  "Create Success",
			key:   "Hello",
			value: "World",
			mockFunc: func(m *mocks.Client) {
				m.On("CreatePostgresRow", "Hello", "World").Return(nil).Times(1)
			},
			expectedError: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockClient := mocks.NewClient(t)
			tt.mockFunc(mockClient)

			db := &Postgres{client: mockClient}

			err := db.Create(tt.key, tt.value)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockClient.AssertExpectations(t)

		})
	}
}

func TestPostgres_Delete(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		mockFunc      func(m *mocks.Client)
		expectedError string
	}{
		{
			name:          "Empty Key",
			key:           "",
			mockFunc:      func(m *mocks.Client) {},
			expectedError: "key cannot be empty",
		},
		{
			name: "Delete Failure",
			key:  "Hello",
			mockFunc: func(m *mocks.Client) {
				m.On("DeletePostgresRow", "Hello").Return(errors.New("db error")).Times(1)
			},
			expectedError: "failed to delete postgres row: db error",
		},
		{
			name: "Delete Success",
			key:  "Hello",
			mockFunc: func(m *mocks.Client) {
				m.On("DeletePostgresRow", "Hello").Return(nil).Times(1)
			},
			expectedError: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockClient := mocks.NewClient(t)
			tt.mockFunc(mockClient)

			db := &Postgres{client: mockClient}

			err := db.Delete(tt.key)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockClient.AssertExpectations(t)

		})
	}
}

func TestPostgres_Update(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		value         string
		mockFunc      func(m *mocks.Client)
		expectedError string
	}{
		{
			name:          "Empty Key",
			key:           "",
			value:         "World",
			mockFunc:      func(m *mocks.Client) {},
			expectedError: "key cannot be empty",
		},
		{
			name:  "Update Failure",
			key:   "Hello",
			value: "World",
			mockFunc: func(m *mocks.Client) {
				m.On("UpdatePostgresRow", "Hello", "World").Return(errors.New("db error")).Times(1)
			},
			expectedError: "failed to update postgres row: db error",
		},
		{
			name:  "Update Success",
			key:   "Hello",
			value: "World",
			mockFunc: func(m *mocks.Client) {
				m.On("UpdatePostgresRow", "Hello", "World").Return(nil).Times(1)
			},
			expectedError: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockClient := mocks.NewClient(t)
			tt.mockFunc(mockClient)

			db := &Postgres{client: mockClient}

			err := db.Update(tt.key, tt.value)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockClient.AssertExpectations(t)

		})
	}
}

func TestPostgres_Get(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		mockFunc      func(m *mocks.Client)
		expectedError string
	}{
		{
			name:          "Empty Key",
			key:           "",
			mockFunc:      func(m *mocks.Client) {},
			expectedError: "key cannot be empty",
		},
		{
			name: "Get Failure",
			key:  "Hello",
			mockFunc: func(m *mocks.Client) {
				m.On("GetPostgresRow", "Hello").Return("", errors.New("db error")).Times(1)
			},
			expectedError: "failed to get postgres row: db error",
		},
		{
			name: "Get Success",
			key:  "Hello",
			mockFunc: func(m *mocks.Client) {
				m.On("GetPostgresRow", "Hello").Return("World", nil).Times(1)
			},
			expectedError: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockClient := mocks.NewClient(t)
			tt.mockFunc(mockClient)

			db := &Postgres{client: mockClient}

			err := db.Get(tt.key)

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockClient.AssertExpectations(t)

		})
	}
}

func TestPostgres_Show(t *testing.T) {
	tests := []struct {
		name          string
		mockFunc      func(m *mocks.Client)
		expectedError string
	}{
		{
			name: "Show Failure",
			mockFunc: func(m *mocks.Client) {
				m.On("ShowPostgresRow").Return(nil, errors.New("db error")).Times(1)
			},
			expectedError: "failed to show postgres row: db error",
		},
		{
			name: "Get Success",
			mockFunc: func(m *mocks.Client) {
				m.On("ShowPostgresRow").Return(map[string]string{"Hello": "World"}, nil).Times(1)
			},
			expectedError: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockClient := mocks.NewClient(t)
			tt.mockFunc(mockClient)

			db := &Postgres{client: mockClient}

			err := db.Show()

			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockClient.AssertExpectations(t)

		})
	}
}
