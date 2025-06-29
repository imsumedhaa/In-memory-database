package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/imsumedhaa/In-memory-database/pkg/client/postgres/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHttp_Create_Success(t *testing.T) {

	mockClient := mocks.NewClient(t)
	mockClient.On("CreatePostgresRow", "Hello", "World").Return(nil).Times(1)

	handler := &Http{client: mockClient}

	reqBody := Request{
		Key:   "Hello",
		Value: "World",
	}
	byteBody, _ := json.Marshal(reqBody)

	//fake http req using in build httptest
	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(byteBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder() //Pretends to be a Response Writer. because in testing there is no such server

	handler.create(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "expected 200 status code")

	expectedResponse := Response{Message: "Row created succesfully"}

	var actualResponse Response
	err := json.NewDecoder(rec.Body).Decode(&actualResponse)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	mockClient.AssertExpectations(t)

}

func TestHttp_Create_Failure(t *testing.T) {
	mockClient := mocks.NewClient(t) //Creates a mock version of your postgres.Client interface.

	// Simulate DB error
	mockClient.On("CreatePostgresRow", "Hello", "World").Return(errors.New("server error")).Once()

	handler := &Http{client: mockClient}

	reqBody := Request{
		Key:   "Hello",
		Value: "World",
	}
	byteBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(byteBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.create(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code, "Expected status 500")
	assert.Contains(t, rec.Body.String(), "Failed to create row: server error")

	mockClient.AssertExpectations(t)
}

func TestHttp_Update_Success(t *testing.T) {

	mockClient := mocks.NewClient(t)
	mockClient.On("UpdatePostgresRow", "Hello", "World").Return(nil).Times(1)

	handler := &Http{client: mockClient}

	reqBody := Request{
		Key:   "Hello",
		Value: "World",
	}
	byteReqBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, "/update", bytes.NewBuffer(byteReqBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler.update(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Epected status code 200")

	expectedResponse := Response{Message: "Row updated succesfully"}

	var actualResponse Response
	err := json.NewDecoder(rec.Body).Decode(&actualResponse)

	assert.NoError(t, err)
	assert.Equal(t,expectedResponse,actualResponse)

	mockClient.AssertExpectations(t)
}
