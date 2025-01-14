package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/bufbuild/connect-go"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
)

func TestNewUserServer(t *testing.T) {
	tests := []struct {
		name string
		want *userServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userServer_GetUser(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[dynamicv1.GetUserRequest]
	}
	tests := []struct {
		name    string
		s       *userServer
		args    args
		want    *connect.Response[dynamicv1.GetUserResponse]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("userServer.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userServer.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
