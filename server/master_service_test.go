package server

import (
	"context"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bufbuild/connect-go"
	. "github.com/mamoruuji/dynamic_novel_api/config"
	_ "github.com/mamoruuji/dynamic_novel_api/db/models"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestNewMasterServer(t *testing.T) {
	tests := []struct {
		name string
		want *masterServer
	}{
		{
			name: "正常系: DBへの接続",
			want: &masterServer{
				db: GetDB(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMasterServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMasterServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_masterServer_ListSectionMaster(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "failed to create mock db")
	defer db.Close()

	boil.SetDB(db)

	ctx := context.Background()

	tests := []struct {
		name      string
		req       *connect.Request[dynamicv1.ListSectionMastersRequest]
		setupMock func()
		wantRes   dynamicv1.ListSectionMastersResponse
		wantErr   bool
	}{
		{
			name: "正常系: 章のみを削除",
			req:  connect.NewRequest(&dynamicv1.ListSectionMastersRequest{}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "chapters" where "chapter_id"=$1`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"chapter_id"}).
							AddRow(1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "pages".* FROM "pages" WHERE ("pages"."chapter_id" = $1)`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{}),
					)

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "chapters" WHERE "chapter_id"=$1`)).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			// wantRes: &dynamicv1.ListSectionMastersResponse{
			// 	SectionTypeMaster:
			// 		[]*SectionTypeMaster{},
			// },
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			s := &masterServer{db: db}

			res, err := s.ListSectionMasters(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSectionMaster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, tt.wantRes.Msg.SectionTypeMaster, res.Msg.ChapterId)
			}
		})
	}
}

func Test_masterServer_ListSortMaster(t *testing.T) {
	type args struct {
		ctx context.Context
		in1 *connect.Request[dynamicv1.ListSortMasterRequest]
	}
	tests := []struct {
		name    string
		s       *masterServer
		args    args
		want    *connect.Response[dynamicv1.ListSortMasterResponse]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ListSortMaster(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("masterServer.ListSortMaster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("masterServer.ListSortMaster() = %v, want %v", got, tt.want)
			}
		})
	}
}
