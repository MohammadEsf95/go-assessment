package main

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	contract "service3/contract/proto"
	"service3/entity"
	"testing"
)

type MockRepo struct {
	mock.Mock
}

func (r *MockRepo) Insert(data []entity.Model) error {
	args := r.Called(data)
	return args.Error(0)
}

type mockService1Client struct {
	mock.Mock
}

func (m *mockService1Client) GetData(ctx context.Context, req *contract.Service1Request, opts ...grpc.CallOption) (*contract.Service1Response, error) {
	args := m.Called(ctx, req, opts)
	if resp, ok := args.Get(0).(*contract.Service1Response); ok {
		return resp, args.Error(1)
	}
	return nil, args.Error(1)
}

type mockService2Client struct {
	mock.Mock
}

func (m *mockService2Client) GetData(ctx context.Context, req *contract.Service2Request, opts ...grpc.CallOption) (*contract.Service2Response, error) {
	args := m.Called(ctx, req, opts)
	if resp, ok := args.Get(0).(*contract.Service2Response); ok {
		return resp, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestGetResult_Success(t *testing.T) {
	// Setup mocks
	mockRepo := new(MockRepo)
	mockSrv1Client := new(mockService1Client)
	mockSrv2Client := new(mockService2Client)

	// Setup expectations
	mockRepo.On("Insert", mock.Anything).Return(nil)
	mockSrv1Client.On("GetData", mock.Anything, &contract.Service1Request{Id: 1}, mock.Anything).
		Return(&contract.Service1Response{Message: "one"}, nil)
	mockSrv2Client.On("GetData", mock.Anything, &contract.Service2Request{Id: 1}, mock.Anything).
		Return(&contract.Service2Response{Message: "one"}, nil)

	// Execute
	GetResult(mockRepo, mockSrv1Client, mockSrv2Client)

	// Assert
	mockRepo.AssertExpectations(t)
	mockSrv1Client.AssertExpectations(t)
	mockSrv2Client.AssertExpectations(t)
}

func TestGetDataFromServiceOneClient_Success(t *testing.T) {
	mockSrv1Client := new(mockService1Client)
	expectedResp := &contract.Service1Response{Message: "one"}

	mockSrv1Client.On("GetData", mock.Anything, &contract.Service1Request{Id: 1}, mock.Anything).
		Return(expectedResp, nil)

	resp, err := getDataFromServiceOneClient(context.Background(), mockSrv1Client)

	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)
	mockSrv1Client.AssertExpectations(t)
}

func TestGetDataFromServiceOneClient_Error(t *testing.T) {
	mockSrv1Client := new(mockService1Client)
	expectedErr := errors.New("service1 error")

	mockSrv1Client.On("GetData", mock.Anything, &contract.Service1Request{Id: 4}, mock.Anything).
		Return(nil, expectedErr)

	resp, err := getDataFromServiceOneClient(context.Background(), mockSrv1Client)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, expectedErr, err)
	mockSrv1Client.AssertExpectations(t)
}

func TestGetDataFromServiceTwoClient_Success(t *testing.T) {
	mockSrv2Client := new(mockService2Client)
	expectedResp := &contract.Service2Response{Message: "success"}

	mockSrv2Client.On("GetData", mock.Anything, &contract.Service2Request{Id: 1}, mock.Anything).
		Return(expectedResp, nil)

	resp, err := getDataFromServiceTwoClient(mockSrv2Client)

	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)
	mockSrv2Client.AssertExpectations(t)
}

func TestDatabaseOperation_Success(t *testing.T) {
	mockRepo := new(MockRepo)
	mockRepo.On("Insert", mock.Anything).Return(nil)

	err := databaseOperation(mockRepo)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
