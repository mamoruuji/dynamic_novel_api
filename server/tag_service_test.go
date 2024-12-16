package server_test

import (
	"context"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/golang/mock/gomock"

	"github.com/mamoruuji/dynamic_novel_api/db/models"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
	"github.com/mamoruuji/dynamic_novel_api/mocks"
	"github.com/mamoruuji/dynamic_novel_api/server"
	"github.com/stretchr/testify/assert"
)

func TestListTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockContextExecutor(ctrl)

	mockTags := []*models.Tag{
		{TagID: 1, Name: "tag1"},
		{TagID: 2, Name: "tag2"},
	}

	mockDB.EXPECT().
		All(gomock.Any(), gomock.Any()).
		Return(mockTags, nil).Times(1)

	tagServer := server.NewTagServer(mockDB)

	// tags, err := mockDB.All(context.Background(), nil)
	// if err != nil {
	// 	t.Errorf("unexpected error: %v", err)
	// }

	req := &connect.Request[dynamicv1.ListTagsRequest]{}
	resp, err := tagServer.ListTags(context.Background(), req)

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
