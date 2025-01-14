package server

import (
	"context"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bufbuild/connect-go"
	. "github.com/mamoruuji/dynamic_novel_api/config"
	. "github.com/mamoruuji/dynamic_novel_api/db/models"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/boil"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"
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
		WillReturnRows(sqlmock.NewRows([]string{"dynamic_id", "title"}).
			AddRow(1, "Dynamic Title 1"))

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
		WillReturnRows(sqlmock.NewRows([]string{"chapter_id", "dynamic_id", "title", "order"}).
			AddRow(1, 1, "Chapter Title 1", 1).
			AddRow(2, 1, "Chapter Title 2", 2))

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT * FROM "terms" WHERE ("terms"."chapter_id" IN ($1,$2)) ORDER BY terms.order;`)).
		WithArgs(1, 2).
		WillReturnRows(sqlmock.NewRows([]string{"term_id", "name"}).
			AddRow(1, "Name").
			AddRow(2, "Name"))

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT * FROM "pages" WHERE ("pages"."chapter_id" IN ($1,$2)) ORDER BY pages.order;`)).
		WithArgs(1, 2).
		WillReturnRows(sqlmock.NewRows([]string{"page_id", "chapter_id", "title", "order"}).
			AddRow(1, 1, "Page Title 1", 1).
			AddRow(2, 2, "Page Title 2", 2))

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT * FROM "terms" WHERE ("terms"."page_id" IN ($1,$2)) ORDER BY terms.order;`)).
		WithArgs(1, 2).
		WillReturnRows(sqlmock.NewRows([]string{"term_id", "name"}).
			AddRow(1, "Name").
			AddRow(2, "Name"))

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT * FROM "sections" WHERE ("sections"."page_id" IN ($1,$2)) ORDER BY sections.order;`)).
		WithArgs(1, 2).
		WillReturnRows(sqlmock.NewRows([]string{"section_id", "title"}).
			AddRow(1, "Section Title 1"))

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
				Title:     "Dynamic Title 1",
				Overview:  "Text",
				UserId:    "use id",
				Published: true,
				PenName:   "pen name",
				ImageUrl:  "image url",
				Chapters: []*dynamicv1.ChapterData{
					{
						ChapterId: 1,
						Title:     "Chapter Title 1",
						Order:     1,
						Pages: []*dynamicv1.PageData{
							{
								PageId: 1,
								Title:  "Page Title 1",
								Order:  1,
								Sections: []*dynamicv1.SectionData{
									{
										SectionId:     1,
										Name:          "Section Title 1",
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
	req := &connect.Request[dynamicv1.GetDynamicRequest]{
		Msg: &dynamicv1.GetDynamicRequest{
			DynamicId: 1,
		},
	}

	mock.MatchExpectationsInOrder(false)

	mock.ExpectQuery(regexp.
		QuoteMeta(`SELECT "dynamics".* FROM "dynamics" WHERE ("dynamics"."dynamic_id" = $1) LIMIT 1;`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"dynamic_id", "title"}).
			AddRow(1, "Dynamic Title 1"))

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
			want: connect.NewResponse(&dynamicv1.GetDynamicResponse{
				DynamicId: 1,
				Title:     "Dynamic Title 1",
				Overview:  "Text",
				UserId:    "use id",
				Published: true,
				PenName:   "pen name",
				ImageUrl:  "image url",
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
