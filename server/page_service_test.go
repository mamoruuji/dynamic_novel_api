package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/bufbuild/connect-go"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestNewPageServer(t *testing.T) {
	tests := []struct {
		name string
		want *pageServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPageServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPageServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pageServer_ListPages(t *testing.T) {
	type args struct {
		ctx context.Context
		in1 *connect.Request[dynamicv1.ListPagesRequest]
	}
	tests := []struct {
		name    string
		s       *pageServer
		args    args
		want    *connect.Response[dynamicv1.ListPagesResponse]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ListPages(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("pageServer.ListPages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pageServer.ListPages() = %v, want %v", got, tt.want)
			}
		})
	}
}
