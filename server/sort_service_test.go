package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/bufbuild/connect-go"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestNewSortServer(t *testing.T) {
	tests := []struct {
		name string
		want *sortServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSortServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSortServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortServer_ListSorts(t *testing.T) {
	type args struct {
		ctx context.Context
		in1 *connect.Request[dynamicv1.ListSortsRequest]
	}
	tests := []struct {
		name    string
		s       *sortServer
		args    args
		want    *connect.Response[dynamicv1.ListSortsResponse]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ListSorts(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("sortServer.ListSorts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortServer.ListSorts() = %v, want %v", got, tt.want)
			}
		})
	}
}
