package server

import (
	"context"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bufbuild/connect-go"
	. "github.com/mamoruuji/dynamic_novel_api/config"
	_ "github.com/mamoruuji/dynamic_novel_api/db/models"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewDynamicServer(t *testing.T) {
	tests := []struct {
		name string
		want *dynamicServer
	}{
		{
			name: "正常系: DBへの接続",
			want: &dynamicServer{
				db: GetDB(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDynamicServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDynamicServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dynamicServer_GetDynamic(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "failed to create mock db")
	defer db.Close()

	boil.SetDB(db)

	ctx := context.Background()
	req := &connect.Request[dynamicv1.GetDynamicRequest]{
		Msg: &dynamicv1.GetDynamicRequest{
			DynamicId: 1,
		},
	}

	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT "dynamics".* FROM "dynamics" WHERE ("dynamics"."dynamic_id" = $1) LIMIT 1;`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"dynamic_id", "name"}).
			AddRow(1, "Dynamic Name 1"))

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT * FROM "users" WHERE ("users"."user_id" IN ($1));`)).
		WithArgs("").
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "name"}).
			AddRow("", "User Name"))

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT * FROM "terms" WHERE ("terms"."dynamic_id" IN ($1)) ORDER BY terms.order;`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"term_id", "name"}).
			AddRow(1, "Name"))

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT * FROM "chapters" WHERE ("chapters"."dynamic_id" IN ($1)) ORDER BY chapters.order;`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"chapter_id", "dynamic_id", "name", "order"}).
			AddRow(1, 1, "Chapter Name 1", 1).
			AddRow(2, 1, "Chapter Name 2", 2))

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT * FROM "terms" WHERE ("terms"."chapter_id" IN ($1,$2)) ORDER BY terms.order;`)).
		WithArgs(1, 2).
		WillReturnRows(sqlmock.NewRows([]string{"term_id", "name"}).
			AddRow(1, "Name").
			AddRow(2, "Name"))

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT * FROM "pages" WHERE ("pages"."chapter_id" IN ($1,$2)) ORDER BY pages.order;`)).
		WithArgs(1, 2).
		WillReturnRows(sqlmock.NewRows([]string{"page_id", "chapter_id", "name", "order"}).
			AddRow(1, 1, "Page Name 1", 1).
			AddRow(2, 2, "Page Name 2", 2))

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT * FROM "terms" WHERE ("terms"."page_id" IN ($1,$2)) ORDER BY terms.order;`)).
		WithArgs(1, 2).
		WillReturnRows(sqlmock.NewRows([]string{"term_id", "name"}).
			AddRow(1, "Name").
			AddRow(2, "Name"))

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT * FROM "sections" WHERE ("sections"."page_id" IN ($1,$2)) ORDER BY sections.order;`)).
		WithArgs(1, 2).
		WillReturnRows(sqlmock.NewRows([]string{"section_id", "name"}).
			AddRow(1, "Section Name 1"))

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT "dynamics_on_tags".* FROM "dynamics_on_tags" WHERE ("dynamics_on_tags"."dynamic_id" = $1)`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"dynamic_id", "tag_id"}).
			AddRow(1, 1))

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT * FROM "tags" WHERE ("tags"."tag_id" IN ($1))`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"tag_id", "name"}).
			AddRow(1, "name"))

	type args struct {
		ctx context.Context
		req *connect.Request[dynamicv1.GetDynamicRequest]
	}
	tests := []struct {
		name    string
		s       *dynamicServer
		args    args
		want    *connect.Response[dynamicv1.GetDynamicResponse]
		wantErr bool
	}{
		{
			name: "正常系: ユーザー ID に紐づくデータを取得",
			s: &dynamicServer{
				db: db,
			},
			args: args{
				ctx: ctx,
				req: req,
			},
			want: connect.NewResponse(&dynamicv1.GetDynamicResponse{
				DynamicId: 1,
				Name:      "Dynamic Name 1",
				Overview:  "Text",
				UserId:    "use id",
				Published: true,
				PenName:   "pen name",
				ImageUrl:  "image url",
				Chapters: []*dynamicv1.ChapterData{
					{
						ChapterId: 1,
						Name:      "Chapter Name 1",
						Order:     1,
						Pages: []*dynamicv1.PageData{
							{
								PageId: 1,
								Name:   "Page Name 1",
								Order:  1,
								Sections: []*dynamicv1.SectionData{
									{
										SectionId:     1,
										Name:          "Section Name 1",
										TypeSection:   "monologue",
										TypePosition:  "center",
										TypeAnimation: "none",
										FrameColor:    "red",
										Text:          "text",
										TextColor:     "blue",
										TextSize:      12,
										Font:          "monologueFont",
										ImageUrl:      "image url",
										Terms: []*dynamicv1.TermData{
											{
												TermId:   4,
												Name:     "term section 4",
												Text:     "text",
												Order:    1,
												ImageUrl: "image url",
											},
										},
									},
								},
								Terms: []*dynamicv1.TermData{
									{
										TermId:   3,
										Name:     "term page 3",
										Text:     "text",
										Order:    1,
										ImageUrl: "image url",
									},
								},
							},
						},
						Terms: []*dynamicv1.TermData{
							{
								TermId:   2,
								Name:     "term chapter 2",
								Text:     "text",
								Order:    1,
								ImageUrl: "image url",
							},
						},
					},
				},
				Terms: []*dynamicv1.TermData{
					{
						TermId:   1,
						Name:     "term dynamic 1",
						Text:     "text",
						Order:    1,
						ImageUrl: "image url",
					},
				},
				Tags: []*dynamicv1.TagData{
					{
						TagId: 1,
						Name:  "tag dynamic 1",
					},
				},
				// CreatedTime: time.Now(),
				// UpdatedTime: time.Now(),
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetDynamic(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("dynamicServer.GetDynamic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dynamicServer.GetDynamic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dynamicServer_ListDynamics(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "failed to create mock db")
	defer db.Close()

	boil.SetDB(db)

	ctx := context.Background()
	req := &connect.Request[dynamicv1.ListDynamicsRequest]{
		Msg: &dynamicv1.ListDynamicsRequest{
			UserId:       "user id",
			SortCategory: 4,
		},
	}

	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.
		QuoteMeta(`select * from "type_of_sorts" where "type_of_sort_id"=$1`)).
		WithArgs(4).
		WillReturnRows(sqlmock.NewRows([]string{"type_of_sort_id", "name", "sql"}).
			AddRow(4, "お気に入り数", "COUNT(marks.mark_id)"))

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT "dynamics".* FROM "dynamics" WHERE ("dynamics"."user_id" = $1) ORDER BY COUNT(marks.mark_id)`)).
		WithArgs("user id").
		WillReturnRows(sqlmock.NewRows([]string{"dynamic_id", "name", "overview", "pen_name", "image_url", "user_id"}).
			AddRow(1, "Dynamic Name 1", "Text", "pen name", "image url", "user_id").
			AddRow(2, "Dynamic Name 2", "Text", "pen name", "image url", "user_id"))

		// キーワード検索用
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE ("users"."user_id" IN ($1))`)).
		WithArgs("user_id").
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "username", "pen_name", "email"}).
			AddRow("user_id", "test_user", "pen name", "test@example.com"))

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "dynamics_on_tags" WHERE ("dynamics_on_tags"."dynamic_id" IN ($1,$2));`)).
		WithArgs(1, 2).
		WillReturnRows(sqlmock.NewRows([]string{"dynamic_id", "tag_id"}).
			AddRow(1, 1).
			AddRow(2, 2))

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "tags" WHERE ("tags"."tag_id" IN ($1,$2));`)).
		WithArgs(1, 2).
		WillReturnRows(sqlmock.NewRows([]string{"tag_id", "name"}).
			AddRow(1, "name1").
			AddRow(2, "name2"))

		// SetTagData()
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT "dynamics_on_tags".* FROM "dynamics_on_tags" WHERE ("dynamics_on_tags"."dynamic_id" = $1)`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"dynamic_id", "tag_id"}).
			AddRow(1, 101))

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "tags" WHERE ("tags"."tag_id" IN ($1));`)).
		WithArgs(101).
		WillReturnRows(sqlmock.NewRows([]string{"tag_id", "name"}).
			AddRow(101, "tag dynamic 1"))

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT "dynamics_on_tags".* FROM "dynamics_on_tags" WHERE ("dynamics_on_tags"."dynamic_id" = $1)`)).
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{"dynamic_id", "tag_id"}).
			AddRow(2, 102))

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "tags" WHERE ("tags"."tag_id" IN ($1));`)).
		WithArgs(102).
		WillReturnRows(sqlmock.NewRows([]string{"tag_id", "name"}).
			AddRow(102, "tag dynamic 2"))

	type args struct {
		ctx context.Context
		req *connect.Request[dynamicv1.ListDynamicsRequest]
	}
	tests := []struct {
		name    string
		s       *dynamicServer
		args    args
		want    *connect.Response[dynamicv1.ListDynamicsResponse]
		wantErr bool
	}{
		{
			name: "正常系: ユーザー ID に紐づくデータを取得",
			s: &dynamicServer{
				db: db,
			},
			args: args{
				ctx: ctx,
				req: req,
			},
			want: connect.NewResponse(&dynamicv1.ListDynamicsResponse{
				Dynamics: []*dynamicv1.ListDynamicData{
					{
						DynamicId: 1,
						Name:      "Dynamic Name 1",
						Overview:  "Text",
						PenName:   "pen name",
						ImageUrl:  "image url",
						Tags: []*dynamicv1.TagData{
							{
								TagId: 101,
								Name:  "tag dynamic 1",
							},
						},
						CreatedTime: timestamppb.New(time.Unix(-62135596800, 0)),
						UpdatedTime: timestamppb.New(time.Unix(-62135596800, 0)),
					},
					{
						DynamicId: 2,
						Name:      "Dynamic Name 2",
						Overview:  "Text",
						PenName:   "pen name",
						ImageUrl:  "image url",
						Tags: []*dynamicv1.TagData{
							{
								TagId: 102,
								Name:  "tag dynamic 2",
							},
						},
						CreatedTime: timestamppb.New(time.Unix(-62135596800, 0)),
						UpdatedTime: timestamppb.New(time.Unix(-62135596800, 0)),
					},
				},
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ListDynamics(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("dynamicServer.ListDynamics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dynamicServer.ListDynamics() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dynamicServer_UpdateContentsOrder(t *testing.T) {
	type args struct {
		ctx context.Context
		req *connect.Request[dynamicv1.UpdateContentsOrderRequest]
	}
	tests := []struct {
		name    string
		s       *dynamicServer
		args    args
		want    *connect.Response[dynamicv1.UpdateContentsOrderResponse]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateContentsOrder(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("dynamicServer.UpdateContentsOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dynamicServer.UpdateContentsOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
