package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/bufbuild/connect-go"
	. "github.com/mamoruuji/dynamic_novel_api/db/models"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
)

func TestNewImageServer(t *testing.T) {
	tests := []struct {
		name string
		want *imageServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewImageServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImageServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_imageServer_ListUserImages(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[dynamicv1.ListUserImagesRequest]
	}
	tests := []struct {
		name    string
		s       *imageServer
		args    args
		want    *connect.Response[dynamicv1.ListUserImagesResponse]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ListUserImages(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("imageServer.ListUserImages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageServer.ListUserImages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertFolder(t *testing.T) {
	type args struct {
		f *Folder
	}
	tests := []struct {
		name string
		args args
		want *dynamicv1.FolderData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertFolder(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
