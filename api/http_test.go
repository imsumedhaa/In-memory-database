package api

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/imsumedhaa/In-memory-database/pkg/client/postgres/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHttp_Create(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		requestBody  string
		mockFunc     func(m *mocks.Client)
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Invalid JSON",
			method:       http.MethodPost,
			requestBody:  `invalid-json`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid create body request",
		},
		{
			name:         "Wrong Http Method",
			method:       http.MethodGet,
			requestBody:  `{"Key":"Hello","Value":"World"}`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "Method not allowed",
		},
		{
			name:         "Missing key",
			method:       http.MethodPost,
			requestBody:  `{"Key":"","Value":"World"}`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Key cannot be empty",
		},
		{
			name:        "Create Failure - db error",
			method:      http.MethodPost,
			requestBody: `{"Key":"Hello","Value":"World"}`,
			mockFunc: func(m *mocks.Client) {
				m.On("CreatePostgresRow", "Hello", "World").Return(errors.New("db error")).Times(1)
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: "Failed to create row: db error",
		},
		{
			name:        "Create Success",
			method:      http.MethodPost,
			requestBody: `{"Key":"Hello","Value":"World"}`,
			mockFunc: func(m *mocks.Client) {
				m.On("CreatePostgresRow", "Hello", "World").Return(nil).Times(1)
			},
			expectedCode: http.StatusOK,
			expectedBody: "Row created succesfully",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockClient := mocks.NewClient(t)
			tt.mockFunc(mockClient)

			handler := &Http{client: mockClient}

			req := httptest.NewRequest(tt.method, "/create", bytes.NewBuffer([]byte(tt.requestBody)))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			handler.create(rec, req)

			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.expectedBody)

			mockClient.AssertExpectations(t)
		})
	}
}

func TestHttp_Delete(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		requestBody  string
		mockFunc     func(m *mocks.Client)
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Invalid JSON",
			method:       http.MethodDelete,
			requestBody:  `invalid-json`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid Delete body request",
		},
		{
			name:         "Missig Key",
			method:       http.MethodDelete,
			requestBody:  `{"Key":""}`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Key cannot be empty",
		},
		{
			name:         "Wrong Http method",
			method:       http.MethodGet,
			requestBody:  `{"Key":"Hello"}`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "Method not allowed",
		},
		{
			name:        "Delete Failure - db error",
			method:      http.MethodDelete,
			requestBody: `{"Key":"Hello"}`,
			mockFunc: func(m *mocks.Client) {
				m.On("DeletePostgresRow", "Hello").Return(errors.New("db error")).Times(1)
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: "Failed to delete row: db error",
		},
		{
			name:        "Delete Success",
			method:      http.MethodDelete,
			requestBody: `{"Key":"Hello"}`,
			mockFunc: func(m *mocks.Client) {
				m.On("DeletePostgresRow", "Hello").Return(nil).Times(1)
			},
			expectedCode: http.StatusOK,
			expectedBody: "Row deleted succesfully",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := mocks.NewClient(t)
			tt.mockFunc(mockClient)

			handler := &Http{client: mockClient}

			req := httptest.NewRequest(tt.method, "/delete", bytes.NewBufferString(tt.requestBody))
			rec := httptest.NewRecorder()
			handler.delete(rec, req)

			assert.Equal(t, rec.Code, tt.expectedCode)
			assert.Contains(t, rec.Body.String(), tt.expectedBody)

			mockClient.AssertExpectations(t)
		})
	}
}

func TestHttp_Update(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		requestBody  string
		mockFunc     func(m *mocks.Client)
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Invalid json",
			method:       http.MethodPut,
			requestBody:  `invalid-json`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid update body request",
		},
		{
			name:         "Missing Key",
			method:       http.MethodPut,
			requestBody:  `{"Key":"","Value": "World"}`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Key cannot be empty",
		},
		{
			name:         "Wrong Http Method",
			method:       http.MethodGet,
			requestBody:  `{"Key":"Hello","Value": "World"}`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "Method not allowed",
		},
		{
			name:        "Update Failure",
			method:      http.MethodPut,
			requestBody: `{"Key":"Hello","Value": "World"}`,
			mockFunc: func(m *mocks.Client) {
				m.On("UpdatePostgresRow", "Hello", "World").Return(errors.New("db error")).Times(1)
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: "Failed to update row: db error",
		},
		{
			name:        "Update Success",
			method:      http.MethodPut,
			requestBody: `{"Key":"Hello","Value": "World"}`,
			mockFunc: func(m *mocks.Client) {
				m.On("UpdatePostgresRow", "Hello", "World").Return(nil).Times(1)
			},
			expectedCode: http.StatusOK,
			expectedBody: "Row updated succesfully",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := mocks.NewClient(t)
			tt.mockFunc(mockClient)

			handler := &Http{client: mockClient}

			req := httptest.NewRequest(tt.method, "/update", bytes.NewBufferString(tt.requestBody))
			rec := httptest.NewRecorder()

			handler.update(rec, req)
			assert.Equal(t, rec.Code, tt.expectedCode)
			assert.Contains(t, rec.Body.String(), tt.expectedBody)

			mockClient.AssertExpectations(t)
		})
	}
}

func TestHttp_Get(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		requestBody  string
		mockFunc     func(m *mocks.Client)
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Invalid json",
			method:       http.MethodGet,
			requestBody:  `invalid-json`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid get body request",
		},
		{
			name:         "Missing Key",
			method:       http.MethodGet,
			requestBody:  `{"Key":""}`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Key cannot be empty",
		},
		{
			name:         "Wrong Http Method",
			method:       http.MethodPost,
			requestBody:  `{"Key":"Hello"}`,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "Method not allowed",
		},
		{
			name:        "Get Failure",
			method:      http.MethodGet,
			requestBody: `{"Key":"Hello"}`,
			mockFunc: func(m *mocks.Client) {
				m.On("GetPostgresRow", "Hello").Return("", errors.New("db error")).Times(1)
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: "Failed to get the row: db error",
		},
		{
			name:        "Get Success",
			method:      http.MethodGet,
			requestBody: `{"Key":"Hello"}`,
			mockFunc: func(m *mocks.Client) {
				m.On("GetPostgresRow", "Hello").Return("World", nil).Times(1)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"Key":"Hello","Value":"World"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := mocks.NewClient(t)
			tt.mockFunc(mockClient)

			handler := &Http{client: mockClient}

			req := httptest.NewRequest(tt.method, "/get", bytes.NewBufferString(tt.requestBody))
			rec := httptest.NewRecorder()

			handler.show(rec, req)
			assert.Equal(t, rec.Code, tt.expectedCode)
			assert.Contains(t, rec.Body.String(), tt.expectedBody)

			mockClient.AssertExpectations(t)
		})
	}
}

func TestHttp_Show(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		mockFunc     func(m *mocks.Client)
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Wrong Http Method",
			method:       http.MethodPost,
			mockFunc:     func(m *mocks.Client) {},
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: "Method not allowed",
		},
		{
			name:   "Show Failure",
			method: http.MethodGet,
			mockFunc: func(m *mocks.Client) {
				m.On("ShowPostgresRow").Return(nil, errors.New("db error")).Times(1)
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: "Failed to show row: db error",
		},
		{
			name:        "Show Success",
			method:      http.MethodGet,
			mockFunc: func(m *mocks.Client) {
				m.On("ShowPostgresRow").Return(map[string]string{"Hello": "World"}, nil).Times(1)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"Hello":"World"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := mocks.NewClient(t)
			tt.mockFunc(mockClient)

			handler := &Http{client: mockClient}

			req := httptest.NewRequest(tt.method, "/show", nil)
			rec := httptest.NewRecorder()

			handler.show(rec, req)
			assert.Equal(t, rec.Code, tt.expectedCode)
			assert.Contains(t, rec.Body.String(), tt.expectedBody)

			mockClient.AssertExpectations(t)
		})
	}
}
