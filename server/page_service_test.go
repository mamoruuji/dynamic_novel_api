package server

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bufbuild/connect-go"
	. "github.com/mamoruuji/dynamic_novel_api/config"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestNewPageServer(t *testing.T) {
	tests := []struct {
		name string
		want *pageServer
	}{
		{
			name: "正常系: DBへの接続",
			want: &pageServer{
				db: GetDB(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPageServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPageServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pageServer_AddPage(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "failed to create mock db")
	defer db.Close()

	boil.SetDB(db)

	ctx := context.Background()
	tests := []struct {
		name           string
		req            *connect.Request[dynamicv1.AddPageRequest]
		setupMock      func()
		wantErr        bool
		expectedPageID int32
		expectedName   string
		expectedOrder  int32
	}{
		{
			name: "正常系: ページが存在しない場合",
			req: &connect.Request[dynamicv1.AddPageRequest]{
				Msg: &dynamicv1.AddPageRequest{
					ChapterId: 1,
				},
			},
			setupMock: func() {
				mock.ExpectQuery(regexp.
					QuoteMeta(`SELECT "pages".* FROM "pages" WHERE ("pages"."chapter_id" = $1) ORDER BY "order" DESC;`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"page_id", "order"}))

				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "pages" ("order","chapter_id","created_at","updated_at","name") VALUES ($1,$2,$3,$4,$5) RETURNING "page_id"`)).
					WithArgs(1, 1, sqlmock.AnyArg(), sqlmock.AnyArg(), "新しいページ").
					WillReturnRows(sqlmock.NewRows([]string{"page_id"}).
						AddRow(1))

			},
			wantErr:        false,
			expectedPageID: 1,
			expectedName:   "新しいページ",
			expectedOrder:  1,
		},
		{
			name: "正常系: ページが存在する場合 連番確認①",
			req: &connect.Request[dynamicv1.AddPageRequest]{
				Msg: &dynamicv1.AddPageRequest{
					ChapterId: 1,
				},
			},
			setupMock: func() {
				mock.ExpectQuery(regexp.
					QuoteMeta(`SELECT "pages".* FROM "pages" WHERE ("pages"."chapter_id" = $1) ORDER BY "order" DESC;`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"page_id", "order", "name"}).
						AddRow(1, 1, "新しいページ"))

				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "pages" ("order","chapter_id","created_at","updated_at","name") VALUES ($1,$2,$3,$4,$5) RETURNING "page_id"`)).
					WithArgs(2, 1, sqlmock.AnyArg(), sqlmock.AnyArg(), "新しいページ 2").
					WillReturnRows(sqlmock.NewRows([]string{"page_id"}).
						AddRow(2))

			},
			wantErr:        false,
			expectedPageID: 2,
			expectedName:   "新しいページ 2",
			expectedOrder:  2,
		},
		{
			name: "正常系: ページが存在する場合 連番確認②",
			req: &connect.Request[dynamicv1.AddPageRequest]{
				Msg: &dynamicv1.AddPageRequest{
					ChapterId: 1,
				},
			},
			setupMock: func() {
				mock.ExpectQuery(regexp.
					QuoteMeta(`SELECT "pages".* FROM "pages" WHERE ("pages"."chapter_id" = $1) ORDER BY "order" DESC;`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"page_id", "order", "name"}).
						AddRow(2, 2, "新しいページ 3").
						AddRow(1, 1, "新しいページ"))

				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "pages" ("order","chapter_id","created_at","updated_at","name") VALUES ($1,$2,$3,$4,$5) RETURNING "page_id"`)).
					WithArgs(3, 1, sqlmock.AnyArg(), sqlmock.AnyArg(), "新しいページ 4").
					WillReturnRows(sqlmock.NewRows([]string{"page_id"}).
						AddRow(3))

			},
			wantErr:        false,
			expectedPageID: 3,
			expectedName:   "新しいページ 4",
			expectedOrder:  3,
		},
		{
			name: "異常系: SELECT クエリでエラー発生",
			req: &connect.Request[dynamicv1.AddPageRequest]{
				Msg: &dynamicv1.AddPageRequest{
					ChapterId: 1,
				},
			},
			setupMock: func() {
				mock.ExpectQuery(regexp.
					QuoteMeta(`SELECT "pages".* FROM "pages" WHERE ("pages"."chapter_id" = $1) ORDER BY "order" DESC;`)).
					WithArgs(1).
					WillReturnError(fmt.Errorf("db error"))
			},
			wantErr: true,
		},
		{
			name: "異常系: INSERT クエリでエラー発生",
			req: &connect.Request[dynamicv1.AddPageRequest]{
				Msg: &dynamicv1.AddPageRequest{
					ChapterId: 1,
				},
			},
			setupMock: func() {
				mock.ExpectQuery(regexp.
					QuoteMeta(`SELECT "pages".* FROM "pages" WHERE ("pages"."chapter_id" = $1) ORDER BY "order" DESC;`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"page_id", "order", "name"}).
						AddRow(1, 2, "新しいページ"))

				mock.ExpectExec(regexp.QuoteMeta(
					`INSERT INTO "pages" ("order","chapter_id","name") VALUES ($1,$2,$3)`)).
					WithArgs(3, 1, "新しいページ 2").
					WillReturnError(fmt.Errorf("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			s := &pageServer{db: db}

			res, err := s.AddPage(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddPage() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				assert.Equal(t, tt.expectedPageID, res.Msg.PageId)
				assert.Equal(t, tt.expectedName, res.Msg.Name)
				assert.Equal(t, tt.expectedOrder, res.Msg.Order)
			}
		})
	}
}

func Test_pageServer_DeletePage(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "failed to create mock db")
	defer db.Close()

	boil.SetDB(db)

	ctx := context.Background()

	tests := []struct {
		name      string
		req       *connect.Request[dynamicv1.DeletePageRequest]
		setupMock func()
		wantRes   *dynamicv1.DeletePageResponse
		wantErr   bool
	}{
		{
			name: "正常系: ページのみ削除",
			req: connect.NewRequest(&dynamicv1.DeletePageRequest{
				PageId: 1,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "pages" where "page_id"=$1`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"page_id", "name", "order", "chapter_id"}).
							AddRow(1, "テストページ", 1, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "sections".* FROM "sections" WHERE ("sections"."page_id" = $1)`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{}),
					)

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "pages" WHERE "page_id"=$1`)).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
			},
			wantRes: &dynamicv1.DeletePageResponse{
				Status:  "success",
				Message: "ページが削除されました",
			},
			wantErr: false,
		},
		{
			name: "正常系: ページとセクションを削除",
			req: connect.NewRequest(&dynamicv1.DeletePageRequest{
				PageId: 1,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "pages" where "page_id"=$1`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"page_id", "name", "order", "chapter_id"}).
							AddRow(1, "テストページ", 1, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "sections".* FROM "sections" WHERE ("sections"."page_id" = $1)`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"section_id", "page_id"}).
							AddRow(1, 1).
							AddRow(2, 1),
					)

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "sections" WHERE ("section_id"=$1) OR ("section_id"=$2)`)).
					WithArgs(1, 2).
					WillReturnResult(sqlmock.NewResult(2, 2))

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "pages" WHERE "page_id"=$1`)).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
			},
			wantRes: &dynamicv1.DeletePageResponse{
				Status:  "success",
				Message: "ページが削除されました",
			},
			wantErr: false,
		},
		{
			name: "異常系: トランザクション開始エラー",
			req: connect.NewRequest(&dynamicv1.DeletePageRequest{
				PageId: 1,
			}),
			setupMock: func() {
				mock.ExpectBegin().
					WillReturnError(fmt.Errorf("トランザクション開始エラー"))
			},
			wantRes: &dynamicv1.DeletePageResponse{
				Status:  "failure",
				Message: "トランザクション開始時にエラーが発生しました",
			},
			wantErr: true,
		},
		{
			name: "異常系: 存在しないページを検索してエラー発生",
			req: connect.NewRequest(&dynamicv1.DeletePageRequest{
				PageId: 999,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "pages" where "page_id"=$1`)).
					WithArgs(999).
					WillReturnError(fmt.Errorf("ページが見つかりません"))

				mock.ExpectRollback()
			},
			wantRes: &dynamicv1.DeletePageResponse{
				Status:  "failure",
				Message: "削除対象のページ検索に失敗しました",
			},
			wantErr: true,
		},
		{
			name: "異常系: セクションを検索してエラー発生",
			req: connect.NewRequest(&dynamicv1.DeletePageRequest{
				PageId: 999,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "pages" where "page_id"=$1`)).
					WithArgs(999).
					WillReturnRows(
						sqlmock.NewRows([]string{"page_id", "name", "order", "chapter_id"}).
							AddRow(1, "テストページ", 1, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "sections".* FROM "sections" WHERE ("sections"."page_id" = $1)`)).
					WithArgs(999).
					WillReturnError(fmt.Errorf("セクション検索エラー"))

				mock.ExpectRollback()

			},
			wantRes: &dynamicv1.DeletePageResponse{
				Status:  "failure",
				Message: "削除対象のセクション検索に失敗しました",
			},
			wantErr: true,
		},
		{
			name: "異常系: セクション削除でエラー発生",
			req: connect.NewRequest(&dynamicv1.DeletePageRequest{
				PageId: 999,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "pages" where "page_id"=$1`)).
					WithArgs(999).
					WillReturnRows(
						sqlmock.NewRows([]string{"page_id", "name", "order", "chapter_id"}).
							AddRow(999, "テストページ", 1, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "sections".* FROM "sections" WHERE ("sections"."page_id" = $1)`)).
					WithArgs(999).
					WillReturnRows(
						sqlmock.NewRows([]string{"section_id", "page_id"}).
							AddRow(1, 999).
							AddRow(2, 999),
					)

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "sections" WHERE ("section_id"=$1) OR ("section_id"=$2)`)).
					WithArgs(1, 2).
					WillReturnError(fmt.Errorf("セクション削除エラー"))

				mock.ExpectRollback()

			},
			wantRes: &dynamicv1.DeletePageResponse{
				Status:  "failure",
				Message: "セクション削除に失敗しました",
			},
			wantErr: true,
		},
		{
			name: "異常系: ページ削除でエラー発生",
			req: connect.NewRequest(&dynamicv1.DeletePageRequest{
				PageId: 999,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "pages" where "page_id"=$1`)).
					WithArgs(999).
					WillReturnRows(
						sqlmock.NewRows([]string{"page_id", "name", "order", "chapter_id"}).
							AddRow(999, "テストページ", 1, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "sections".* FROM "sections" WHERE ("sections"."page_id" = $1)`)).
					WithArgs(999).
					WillReturnRows(
						sqlmock.NewRows([]string{"section_id", "page_id"}).
							AddRow(1, 999).
							AddRow(2, 999),
					)

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "sections" WHERE ("section_id"=$1) OR ("section_id"=$2)`)).
					WithArgs(1, 2).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "pages" WHERE ("page_id"=$1)`)).
					WithArgs(999).
					WillReturnError(fmt.Errorf("ページ削除エラー"))

				mock.ExpectRollback()

			},
			wantRes: &dynamicv1.DeletePageResponse{
				Status:  "failure",
				Message: "ページ削除に失敗しました",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			s := &pageServer{db: db}

			res, err := s.DeletePage(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeletePage() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				assert.Equal(t, tt.wantRes.Status, res.Msg.Status)
				assert.Equal(t, tt.wantRes.Message, res.Msg.Message)
			} else if res != nil {
				assert.Equal(t, tt.wantRes.Status, res.Msg.Status)
				assert.Equal(t, tt.wantRes.Message, res.Msg.Message)
			}
		})
	}
}

func Test_pageServer_UpdatePageName(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "failed to create mock db")
	defer db.Close()

	boil.SetDB(db)

	ctx := context.Background()

	tests := []struct {
		name         string
		req          *connect.Request[dynamicv1.UpdatePageNameRequest]
		setupMock    func()
		wantErr      bool
		expectedName string
	}{
		{
			name: "正常系: ページ名を更新",
			req: connect.NewRequest(&dynamicv1.UpdatePageNameRequest{
				PageId: 1,
				Name:   "更新されたページ名",
			}),
			setupMock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "pages" where "page_id"=$1`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"page_id", "name", "order", "chapter_id"}).
							AddRow(1, "古いページ名", 1, 1),
					)

				mock.ExpectExec(regexp.QuoteMeta(
					`UPDATE "pages" SET "name"=$1 WHERE "page_id"=$2`)).
					WithArgs("更新されたページ名", 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr:      false,
			expectedName: "更新されたページ名",
		},
		{
			name: "異常系: 存在しないページを検索してエラー発生",
			req: connect.NewRequest(&dynamicv1.UpdatePageNameRequest{
				PageId: 999,
				Name:   "新しいページ名",
			}),
			setupMock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "pages" WHERE "page_id"=$1`)).
					WithArgs(999).
					WillReturnError(fmt.Errorf("ページが見つかりません"))
			},
			wantErr: true,
		},
		{
			name: "異常系: ページの名前更新時にエラー発生",
			req: connect.NewRequest(&dynamicv1.UpdatePageNameRequest{
				PageId: 999,
				Name:   "更新されたページ名",
			}),
			setupMock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "pages" WHERE "page_id"=$1`)).
					WithArgs(999).
					WillReturnRows(
						sqlmock.NewRows([]string{"page_id", "name", "order", "chapter_id"}).
							AddRow(999, "古いページ名", 1, 1),
					)
				mock.ExpectExec(regexp.QuoteMeta(
					`UPDATE "pages" SET "name"=$1 WHERE "page_id"=$2`)).
					WithArgs("更新されたページ名", 999).
					WillReturnError(fmt.Errorf("ページ更新エラー"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			s := &pageServer{db: db}

			_, err := s.UpdatePageName(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePageName() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				assert.Equal(t, tt.expectedName, tt.req.Msg.Name)
			}
		})
	}
}
