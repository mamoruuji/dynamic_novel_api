package server

import (
	"context"
	"database/sql"
	"fmt"
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

func TestNewChapterServer(t *testing.T) {
	tests := []struct {
		name string
		want *chapterServer
	}{
		{
			name: "正常系: DBへの接続",
			want: &chapterServer{
				db: GetDB(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChapterServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChapterServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chapterServer_AddChapter(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "failed to create mock db")
	defer db.Close()

	boil.SetDB(db)

	ctx := context.Background()
	tests := []struct {
		name              string
		req               *connect.Request[dynamicv1.AddChapterRequest]
		setupMock         func()
		wantErr           bool
		expectedChapterID int32
		expectedName      string
		expectedOrder     int32
		wantRes           *dynamicv1.AddChapterResponse
	}{
		{
			name: "正常系: 章が存在しない場合",
			req: &connect.Request[dynamicv1.AddChapterRequest]{
				Msg: &dynamicv1.AddChapterRequest{
					DynamicId: 1,
				},
			},
			setupMock: func() {
				mock.ExpectQuery(regexp.
					QuoteMeta(`SELECT "chapters".* FROM "chapters" WHERE ("chapters"."dynamic_id" = $1) ORDER BY "order" DESC;`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"chapter_id", "order"}))

				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "chapters" ("order","dynamic_id","created_at","updated_at","name") VALUES ($1,$2,$3,$4,$5) RETURNING "chapter_id"`)).
					WithArgs(1, 1, sqlmock.AnyArg(), sqlmock.AnyArg(), "新しい章").
					WillReturnRows(sqlmock.NewRows([]string{"chapter_id"}).
						AddRow(1))

			},
			wantRes: &dynamicv1.AddChapterResponse{
				ChapterId: 1,
				Name:      "新しい章",
				Order:     1,
			},
			wantErr: false,
		},
		{
			name: "正常系: 章が存在する場合 連番確認①",
			req: &connect.Request[dynamicv1.AddChapterRequest]{
				Msg: &dynamicv1.AddChapterRequest{
					DynamicId: 1,
				},
			},
			setupMock: func() {
				mock.ExpectQuery(regexp.
					QuoteMeta(`SELECT "chapters".* FROM "chapters" WHERE ("chapters"."dynamic_id" = $1) ORDER BY "order" DESC;`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"chapter_id", "order", "name"}).
						AddRow(1, 1, "新しい章"))

				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "chapters" ("order","dynamic_id","created_at","updated_at","name") VALUES ($1,$2,$3,$4,$5) RETURNING "chapter_id"`)).
					WithArgs(2, 1, sqlmock.AnyArg(), sqlmock.AnyArg(), "新しい章 2").
					WillReturnRows(sqlmock.NewRows([]string{"chapter_id"}).
						AddRow(2))

			},
			wantRes: &dynamicv1.AddChapterResponse{
				ChapterId: 2,
				Name:      "新しい章 2",
				Order:     2,
			},
			wantErr: false,
		},
		{
			name: "正常系: ページが存在する場合 連番確認②",
			req: &connect.Request[dynamicv1.AddChapterRequest]{
				Msg: &dynamicv1.AddChapterRequest{
					DynamicId: 1,
				},
			},
			setupMock: func() {
				mock.ExpectQuery(regexp.
					QuoteMeta(`SELECT "chapters".* FROM "chapters" WHERE ("chapters"."dynamic_id" = $1) ORDER BY "order" DESC;`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"chapter_id", "order", "name"}).
						AddRow(2, 2, "新しい章 3").
						AddRow(1, 1, "新しい章"))

				mock.ExpectQuery(regexp.QuoteMeta(
					`INSERT INTO "chapters" ("order","dynamic_id","created_at","updated_at","name") VALUES ($1,$2,$3,$4,$5) RETURNING "chapter_id"`)).
					WithArgs(3, 1, sqlmock.AnyArg(), sqlmock.AnyArg(), "新しい章 4").
					WillReturnRows(sqlmock.NewRows([]string{"chapter_id"}).
						AddRow(3))

			},
			wantRes: &dynamicv1.AddChapterResponse{
				ChapterId: 3,
				Name:      "新しい章 4",
				Order:     3,
			},
			wantErr: false,
		},
		{
			name: "異常系: ページの最大order値を取得でエラー発生",
			req: &connect.Request[dynamicv1.AddChapterRequest]{
				Msg: &dynamicv1.AddChapterRequest{
					DynamicId: 1,
				},
			},
			setupMock: func() {
				mock.ExpectQuery(regexp.
					QuoteMeta(`SELECT "chapters".* FROM "chapters" WHERE ("chapters"."dynamic_id" = $1) ORDER BY "order" DESC;`)).
					WithArgs(1).
					WillReturnError(fmt.Errorf("db error"))
			},
			wantErr: true,
		},
		{
			name: "異常系: INSERT クエリでエラー発生",
			req: &connect.Request[dynamicv1.AddChapterRequest]{
				Msg: &dynamicv1.AddChapterRequest{
					DynamicId: 1,
				},
			},
			setupMock: func() {
				mock.ExpectQuery(regexp.
					QuoteMeta(`SELECT "chapters".* FROM "chapters" WHERE ("chapters"."dynamic_id" = $1) ORDER BY "order" DESC;`)).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"chapter_id", "order", "name"}).
						AddRow(1, 2, "新しい章"))

				mock.ExpectExec(regexp.QuoteMeta(
					`INSERT INTO "chapters" ("order","dynamic_id","name") VALUES ($1,$2,$3)`)).
					WithArgs(3, 1, "新しい章 2").
					WillReturnError(fmt.Errorf("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			s := &chapterServer{db: db}

			res, err := s.AddChapter(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddChapter() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				assert.Equal(t, tt.wantRes.ChapterId, res.Msg.ChapterId)
				assert.Equal(t, tt.wantRes.Name, res.Msg.Name)
				assert.Equal(t, tt.wantRes.Order, res.Msg.Order)
			}
		})
	}
}

func Test_chapterServer_DeleteChapter(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "failed to create mock db")
	defer db.Close()

	boil.SetDB(db)

	ctx := context.Background()
	tests := []struct {
		name      string
		req       *connect.Request[dynamicv1.DeleteChapterRequest]
		setupMock func()
		wantRes   *dynamicv1.DeleteChapterResponse
		wantErr   bool
	}{
		{
			name: "正常系: 章のみを削除",
			req: connect.NewRequest(&dynamicv1.DeleteChapterRequest{
				ChapterId: 1,
			}),
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
			wantRes: &dynamicv1.DeleteChapterResponse{
				Status:  "success",
				Message: "章が削除されました",
			},
			wantErr: false,
		},
		{
			name: "正常系: 章とページを削除",
			req: connect.NewRequest(&dynamicv1.DeleteChapterRequest{
				ChapterId: 1,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "chapters" where "chapter_id"=$1`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"chapter_id", "name", "order", "dynamic_id"}).
							AddRow(1, "テスト章", 1, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "pages".* FROM "pages" WHERE ("pages"."chapter_id" = $1)`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"page_id", "chapter_id"}).
							AddRow(1, 1).
							AddRow(2, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "sections".* FROM "sections" WHERE ("sections"."page_id" = $1) AND ("sections"."page_id" = $2)`)).
					WithArgs(1, 2).
					WillReturnRows(
						sqlmock.NewRows([]string{"section_id", "page_id"}),
					)

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "pages" WHERE ("page_id"=$1) OR ("page_id"=$2)`)).
					WithArgs(1, 2).
					WillReturnResult(sqlmock.NewResult(2, 2))

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "chapters" WHERE "chapter_id"=$1`)).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
			},
			wantRes: &dynamicv1.DeleteChapterResponse{
				Status:  "success",
				Message: "章が削除されました",
			},
			wantErr: false,
		},
		{
			name: "正常系: 章とページとセクションを削除",
			req: connect.NewRequest(&dynamicv1.DeleteChapterRequest{
				ChapterId: 1,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "chapters" where "chapter_id"=$1`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"chapter_id", "name", "order", "dynamic_id"}).
							AddRow(1, "テスト章", 1, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "pages".* FROM "pages" WHERE ("pages"."chapter_id" = $1)`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"page_id", "chapter_id"}).
							AddRow(1, 1).
							AddRow(2, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "sections".* FROM "sections" WHERE ("sections"."page_id" = $1) AND ("sections"."page_id" = $2)`)).
					WithArgs(1, 2).
					WillReturnRows(
						sqlmock.NewRows([]string{"section_id", "page_id"}).
							AddRow(3, 1).
							AddRow(4, 1),
					)

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "sections" WHERE ("section_id"=$1) OR ("section_id"=$2)`)).
					WithArgs(3, 4).
					WillReturnResult(sqlmock.NewResult(2, 2))

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "pages" WHERE ("page_id"=$1) OR ("page_id"=$2)`)).
					WithArgs(1, 2).
					WillReturnResult(sqlmock.NewResult(2, 2))

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "chapters" WHERE "chapter_id"=$1`)).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))

				mock.ExpectCommit()
			},
			wantRes: &dynamicv1.DeleteChapterResponse{
				Status:  "success",
				Message: "章が削除されました",
			},
			wantErr: false,
		},
		{
			name: "異常系: トランザクション開始エラー",
			req: connect.NewRequest(&dynamicv1.DeleteChapterRequest{
				ChapterId: 1,
			}),
			setupMock: func() {
				mock.ExpectBegin().
					WillReturnError(fmt.Errorf("トランザクション開始エラー"))
			},
			wantRes: &dynamicv1.DeleteChapterResponse{
				Status:  "failure",
				Message: "トランザクション開始時にエラーが発生しました",
			},
			wantErr: true,
		},
		{
			name: "異常系: 章を検索してエラー発生",
			req: connect.NewRequest(&dynamicv1.DeleteChapterRequest{
				ChapterId: 999,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "chapters" where "chapter_id"=$1`)).
					WithArgs(999).
					WillReturnError(fmt.Errorf("章検索エラー"))

				mock.ExpectRollback()
			},
			wantRes: &dynamicv1.DeleteChapterResponse{
				Status:  "failure",
				Message: "削除対象の章検索に失敗しました",
			},
			wantErr: true,
		},
		{
			name: "異常系: ページ検索してエラー発生",
			req: connect.NewRequest(&dynamicv1.DeleteChapterRequest{
				ChapterId: 999,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "chapters" where "chapter_id"=$1`)).
					WithArgs(999).
					WillReturnRows(
						sqlmock.NewRows([]string{"chapter_id", "name", "order", "dynamic_id"}).
							AddRow(999, "テスト章", 1, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "pages".* FROM "pages" WHERE ("pages"."chapter_id" = $1)`)).
					WithArgs(999).
					WillReturnError(fmt.Errorf("ページ検索エラー"))

				mock.ExpectRollback()
			},
			wantRes: &dynamicv1.DeleteChapterResponse{
				Status:  "failure",
				Message: "削除対象のページ検索に失敗しました",
			},
			wantErr: true,
		},
		{
			name: "異常系: セクション検索してエラー発生",
			req: connect.NewRequest(&dynamicv1.DeleteChapterRequest{
				ChapterId: 999,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "chapters" where "chapter_id"=$1`)).
					WithArgs(999).
					WillReturnRows(
						sqlmock.NewRows([]string{"chapter_id", "name", "order", "dynamic_id"}).
							AddRow(999, "テスト章", 1, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "pages".* FROM "pages" WHERE ("pages"."chapter_id" = $1)`)).
					WithArgs(999).
					WillReturnRows(
						sqlmock.NewRows([]string{"page_id", "chapter_id"}).
							AddRow(1, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "sections".* FROM "sections" WHERE ("sections"."page_id" = $1)`)).
					WithArgs(1).
					WillReturnError(fmt.Errorf("セクション検索エラー"))

				mock.ExpectRollback()
			},
			wantRes: &dynamicv1.DeleteChapterResponse{
				Status:  "failure",
				Message: "削除対象のセクション検索に失敗しました",
			},
			wantErr: true,
		},
		{
			name: "異常系: セクション削除でエラー",
			req: connect.NewRequest(&dynamicv1.DeleteChapterRequest{
				ChapterId: 1,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "chapters" where "chapter_id"=$1`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"chapter_id", "name", "order", "dynamic_id"}).
							AddRow(1, "テスト章", 1, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "pages".* FROM "pages" WHERE ("pages"."chapter_id" = $1)`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"page_id", "chapter_id"}).
							AddRow(1, 1).
							AddRow(2, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "sections".* FROM "sections" WHERE ("sections"."page_id" = $1) AND ("sections"."page_id" = $2)`)).
					WithArgs(1, 2).
					WillReturnRows(
						sqlmock.NewRows([]string{"section_id", "page_id"}).
							AddRow(3, 1).
							AddRow(4, 1),
					)

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "sections" WHERE ("section_id"=$1) OR ("section_id"=$2)`)).
					WithArgs(3, 4).
					WillReturnError(fmt.Errorf("セクション削除エラー"))

				mock.ExpectRollback()
			},
			wantRes: &dynamicv1.DeleteChapterResponse{
				Status:  "failure",
				Message: "セクション削除に失敗しました",
			},
			wantErr: true,
		},
		{
			name: "異常系: ページ削除でエラー",
			req: connect.NewRequest(&dynamicv1.DeleteChapterRequest{
				ChapterId: 1,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "chapters" where "chapter_id"=$1`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"chapter_id", "name", "order", "dynamic_id"}).
							AddRow(1, "テスト章", 1, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "pages".* FROM "pages" WHERE ("pages"."chapter_id" = $1)`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"page_id", "chapter_id"}).
							AddRow(1, 1).
							AddRow(2, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "sections".* FROM "sections" WHERE ("sections"."page_id" = $1) AND ("sections"."page_id" = $2)`)).
					WithArgs(1, 2).
					WillReturnRows(
						sqlmock.NewRows([]string{"section_id", "page_id"}).
							AddRow(3, 1).
							AddRow(4, 1),
					)

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "sections" WHERE ("section_id"=$1) OR ("section_id"=$2)`)).
					WithArgs(3, 4).
					WillReturnResult(sqlmock.NewResult(2, 2))

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "pages" WHERE ("page_id"=$1) OR ("page_id"=$2)`)).
					WithArgs(1, 2).
					WillReturnError(fmt.Errorf("ページ削除エラー"))

				mock.ExpectRollback()
			},
			wantRes: &dynamicv1.DeleteChapterResponse{
				Status:  "failure",
				Message: "ページ削除に失敗しました",
			},
			wantErr: true,
		},
		{
			name: "異常系: 章削除でエラー",
			req: connect.NewRequest(&dynamicv1.DeleteChapterRequest{
				ChapterId: 1,
			}),
			setupMock: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "chapters" where "chapter_id"=$1`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"chapter_id", "name", "order", "dynamic_id"}).
							AddRow(1, "テスト章", 1, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "pages".* FROM "pages" WHERE ("pages"."chapter_id" = $1)`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"page_id", "chapter_id"}).
							AddRow(1, 1).
							AddRow(2, 1),
					)

				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT "sections".* FROM "sections" WHERE ("sections"."page_id" = $1) AND ("sections"."page_id" = $2)`)).
					WithArgs(1, 2).
					WillReturnRows(
						sqlmock.NewRows([]string{"section_id", "page_id"}).
							AddRow(3, 1).
							AddRow(4, 1),
					)

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "sections" WHERE ("section_id"=$1) OR ("section_id"=$2)`)).
					WithArgs(3, 4).
					WillReturnResult(sqlmock.NewResult(2, 2))

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "pages" WHERE ("page_id"=$1) OR ("page_id"=$2)`)).
					WithArgs(1, 2).
					WillReturnResult(sqlmock.NewResult(2, 2))

				mock.ExpectExec(regexp.QuoteMeta(
					`DELETE FROM "chapters" WHERE "chapter_id"=$1`)).
					WithArgs(1).
					WillReturnError(fmt.Errorf("章削除エラー"))

				mock.ExpectRollback()
			},
			wantRes: &dynamicv1.DeleteChapterResponse{
				Status:  "failure",
				Message: "章削除に失敗しました",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			s := &chapterServer{db: db}

			res, err := s.DeleteChapter(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteChapter() error = %v, wantErr %v", err, tt.wantErr)
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

func Test_chapterServer_UpdateChapterName(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, "failed to create mock db")
	defer db.Close()

	boil.SetDB(db)

	ctx := context.Background()

	tests := []struct {
		name         string
		req          *connect.Request[dynamicv1.UpdateChapterNameRequest]
		setupMock    func()
		wantErr      bool
		expectedName string
	}{
		{
			name: "正常系: 章名を更新",
			req: connect.NewRequest(&dynamicv1.UpdateChapterNameRequest{
				ChapterId: 1,
				Name:      "更新された章名",
			}),
			setupMock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(
					`select * from "chapters" where "chapter_id"=$1`)).
					WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"chapter_id", "name", "order", "dynamic_id"}).
							AddRow(1, "古い章名", 1, 1),
					)

				mock.ExpectExec(regexp.QuoteMeta(
					`UPDATE "chapters" SET "name"=$1 WHERE "chapter_id"=$2`)).
					WithArgs("更新された章名", 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr:      false,
			expectedName: "更新された章名",
		},
		{
			name: "異常系: 存在しない章を検索してエラー発生",
			req: connect.NewRequest(&dynamicv1.UpdateChapterNameRequest{
				ChapterId: 999,
				Name:      "更新された章名",
			}),
			setupMock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "chapters" WHERE "chapter_id"=$1`)).
					WithArgs(999).
					WillReturnError(fmt.Errorf("章が見つかりません"))
			},
			wantErr: true,
		},
		{
			name: "異常系: 章の名前更新時にエラー発生",
			req: connect.NewRequest(&dynamicv1.UpdateChapterNameRequest{
				ChapterId: 999,
				Name:      "更新された章名",
			}),
			setupMock: func() {
				mock.ExpectQuery(regexp.QuoteMeta(
					`SELECT * FROM "chapters" WHERE "chapter_id"=$1`)).
					WithArgs(999).
					WillReturnRows(
						sqlmock.NewRows([]string{"chapter_id", "name", "order", "dynamic_id"}).
							AddRow(999, "古い章名", 1, 1),
					)
				mock.ExpectExec(regexp.QuoteMeta(
					`UPDATE "chapters" SET "name"=$1 WHERE "chapter_id"=$2`)).
					WithArgs("更新された章名", 999).
					WillReturnError(fmt.Errorf("章更新エラー"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			s := &chapterServer{db: db}

			_, err := s.UpdateChapterName(ctx, tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateChapterName() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				assert.Equal(t, tt.expectedName, tt.req.Msg.Name)
			}
		})
	}

}

func Test_handleChapterDeleteError(t *testing.T) {
	type args struct {
		tx  *sql.Tx
		msg string
		err error
	}
	tests := []struct {
		name    string
		args    args
		want    *connect.Response[dynamicv1.DeleteChapterResponse]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handleChapterDeleteError(tt.args.tx, tt.args.msg, tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleChapterDeleteError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleChapterDeleteError() = %v, want %v", got, tt.want)
			}
		})
	}
}
