package server_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/bufbuild/connect-go"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/golang/mock/gomock"
	"github.com/mamoruuji/dynamic_novel_api/mocks"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	argsMock := m.Called(ctx, query, args)
	return argsMock.Get(0).(sql.Result), argsMock.Error(1)
}

func (m *MockDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	argsMock := m.Called(ctx, query, args)
	return argsMock.Get(0).(*sql.Rows), argsMock.Error(1)
}

func (m *MockDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	argsMock := m.Called(ctx, query, args)
	return argsMock.Get(0).(*sql.Row)
}

func (m *MockDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	argsMock := m.Called(ctx, opts)
	return argsMock.Get(0).(*sql.Tx), argsMock.Error(1)
}

func TestListTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHandler := mocks.NewMockTagServiceHandler(ctrl)

	var mockTags []*dynamicv1.TagData
	mockTags = append(mockTags, &dynamicv1.TagData{
		TagId: 1,
		Name:  "tag1",
	})
	mockTags = append(mockTags, &dynamicv1.TagData{
		TagId: 2,
		Name:  "tag2",
	})

	res := connect.NewResponse(&dynamicv1.ListTagsResponse{
		Tags: mockTags,
	})

	mockHandler.EXPECT().
		ListTags(gomock.Any(), gomock.Any()).
		Return(res, nil).Times(1)

	ctx := context.Background()
	req := connect.NewRequest(&dynamicv1.ListTagsRequest{})

	resp, err := mockHandler.ListTags(ctx, req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedTags := []*dynamicv1.TagData{
		{TagId: 1, Name: "tag1"},
		{TagId: 2, Name: "tag2"},
	}

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, len(expectedTags), len(resp.Msg.Tags))
	assert.Equal(t, "tag1", resp.Msg.Tags[0].Name)
}

func TestSetDynamicOnTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHandler := mocks.NewMockTagServiceHandler(ctrl)

	ctx := context.Background()
	dynamicID := int32(1)
	tagNames := []string{"Tag1", "Tag2", "Tag3"}
	req := connect.NewRequest(&dynamicv1.SetDynamicOnTagRequest{
		DynamicId: dynamicID,
		TagNames:  tagNames,
	})

	mockDB := new(MockDB)
	mockTx := new(MockDB)
	mockDB.On("BeginTx", ctx, mock.Anything).Return(mockTx, nil)

	mockTx.On("ExecContext", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	mockTx.On("Commit").Return(nil)

	for _, tagName := range tagNames {
		mockTx.On("ExecContext", ctx, mock.Anything, []interface{}{tagName}).Return(nil, nil)
		mockTx.On("ExecContext", ctx, mock.Anything, mock.Anything).Return(nil, nil)
	}

	mockHandler.EXPECT().SetDynamicOnTag(gomock.Any(), gomock.Any()).Return(connect.NewResponse(&dynamicv1.SetDynamicOnTagResponse{}), nil).Times(1)
	resp, err := mockHandler.SetDynamicOnTag(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
