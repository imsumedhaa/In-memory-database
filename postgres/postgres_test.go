package postgres

import (
	"errors"
	"testing"

	"github.com/imsumedhaa/In-memory-database/pkg/client/postgres/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPostgres_Create_Success(t *testing.T) {
	mockClient := mocks.NewClient(t)

	mockClient.On("CreatePostgresRow", "key1", "value1").Return(nil).Times(1)

	db := Postgres{client: mockClient}

	err := db.Create("key1", "value1")

	assert.NoError(t, err)

	mockClient.AssertExpectations(t)

}

func TestPostgres_Create_Failure(t *testing.T) {
	mockClient := mocks.NewClient(t)

	mockClient.On("CreatePostgresRow", "key1", "val1").Return(errors.New("db error")).Times(1)

	db := Postgres{client: mockClient}

	err := db.Create("key1", "val1")

	assert.EqualError(t, err, "failed to create postgres row: db error")

	mockClient.AssertExpectations(t)
}

func TestPostgres_Delete_Succes(t *testing.T) {
	mockClient := mocks.NewClient(t)
	mockClient.On("DeletePostgresRow", "key1").Return(nil).Times(1)

	db := Postgres{client: mockClient}
	err := db.Delete("key1")

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestPostgres_Delete_Failure(t *testing.T){
	mockClient := mocks.NewClient(t)
	mockClient.On("DeletePostgresRow", "key1").Return(errors.New("db error")).Times(1)

	db:= Postgres{client : mockClient}
	err := db.Delete("key1")

	assert.EqualError(t, err, "failed to delete postgres row: db error")
	mockClient.AssertExpectations(t)

}

func TestPostgres_Update_Success(t *testing.T){
	mockClient := mocks.NewClient(t)

	mockClient.On("UpdatePostgresRow","key1","val1").Return(nil).Times(1)

	db := Postgres{client: mockClient}
	err :=db.Update("key1","val1")

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)

}

func TestPostgres_Update_Failure(t *testing.T){
	mockClient := mocks.NewClient(t)

	mockClient.On("UpdatePostgresRow","key1","val1").Return(errors.New("db error")).Times(1)

	db := Postgres{client: mockClient}
	err :=db.Update("key1","val1")

	assert.EqualError(t, err, "failed to update postgres row: db error")
	mockClient.AssertExpectations(t)

}

func TestPostgres_Get_Success(t *testing.T){
	mockClient := mocks.NewClient(t)

	mockClient.On("GetPostgresRow","key1").Return("value1",nil).Times(1)

	db := Postgres{client: mockClient}
	err :=db.Get("key1")

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)

}

func TestPostgres_Get_Failure(t *testing.T){
	mockClient := mocks.NewClient(t)

	mockClient.On("GetPostgresRow","key1").Return("",errors.New("db error")).Times(1)

	db := Postgres{client: mockClient}
	err :=db.Get("key1")

	assert.EqualError(t, err, "failed to get postgres row: db error")
	mockClient.AssertExpectations(t)
}



func TestPostgres_Show_Success(t *testing.T) {
	mockClient := mocks.NewClient(t)
	expectedStore := map[string]string{
		"key1": "val1",
		"key2": "val2",
	}

	mockClient.On("ShowPostgresRow").Return(expectedStore, nil).Once()

	db := Postgres{client: mockClient}

	err := db.Show()

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestPostgres_Show_Failure(t *testing.T) {
	mockClient := mocks.NewClient(t)
	mockClient.On("ShowPostgresRow").Return(nil, errors.New("some DB error")).Once()

	db := Postgres{client: mockClient}

	err := db.Show()

	assert.EqualError(t, err, "failed to show postgres row: some DB error")
	mockClient.AssertExpectations(t)
}

