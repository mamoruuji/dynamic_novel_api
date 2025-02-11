package server

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"time"

	. "github.com/mamoruuji/dynamic_novel_api/config"

	"github.com/bufbuild/connect-go"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
	"github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1/dynamicv1connect" // generated by protoc-gen-connect-go

	"github.com/mamoruuji/dynamic_novel_api/db/models"
	. "github.com/mamoruuji/dynamic_novel_api/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type pageServer struct {
	dynamicv1connect.PageServiceHandler
	db boil.ContextExecutor
}

func NewPageServer() *pageServer {
	return &pageServer{
		db: GetDB(),
	}
}

func (s *pageServer) AddPage(
	ctx context.Context,
	req *connect.Request[dynamicv1.AddPageRequest],
) (*connect.Response[dynamicv1.AddPageResponse], error) {

	// ページの最大order値を取得
	pages, err := Pages(
		PageWhere.ChapterID.EQ(req.Msg.GetChapterId()),
		OrderBy("\""+PageColumns.Order+"\" DESC"),
	).All(ctx, s.db)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("Error fetching max order for pages: %v", err))
	}

	var maxOrder int32
	if len(pages) == 0 {
		maxOrder = 0
	} else {
		maxOrder = pages[0].Order
	}

	// ユニーク名を生成
	baseName := "新しいページ"
	maxNumber := 1

	namePattern := regexp.MustCompile(fmt.Sprintf(`^%s(?: (\d+))?$`, regexp.QuoteMeta(baseName)))

	for _, page := range pages {
		matches := namePattern.FindStringSubmatch(page.Name)
		if matches == nil {
			continue // マッチしない場合や数字部分が無い場合はスキップ
		}

		if matches[1] == "" {
			if maxNumber == 1 {
				maxNumber = 2 // 「新しい章」が存在する場合、次は「新しい章2」
			}
			continue
		}

		// 数字部分を解析して最大値を更新
		if number, err := strconv.Atoi(matches[1]); err == nil && number >= maxNumber {
			maxNumber = number + 1
		}
	}

	// ユニークな名前を生成
	uniqueName := baseName
	if maxNumber > 1 {
		uniqueName = fmt.Sprintf("%s %d", baseName, maxNumber)
	}

	// 新しいページデータを作成
	page := &models.Page{
		ChapterID: req.Msg.GetChapterId(),
		Name:      uniqueName,
		Order:     maxOrder + 1,
	}

	err = page.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("Error Add to page: %v", err))
	}

	res := connect.NewResponse(&dynamicv1.AddPageResponse{
		PageId:      page.PageID,
		Name:        page.Name,
		Order:       page.Order,
		CreatedTime: page.CreatedAt.Format(time.RFC3339),
	})

	return res, nil
}

func (s *pageServer) DeletePage(
	ctx context.Context,
	req *connect.Request[dynamicv1.DeletePageRequest],
) (*connect.Response[dynamicv1.DeletePageResponse], error) {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		res := connect.NewResponse(&dynamicv1.DeletePageResponse{
			Status:  "failure",
			Message: "トランザクション開始時にエラーが発生しました",
		})

		return res, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to start transaction: %w", err))
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// 削除対象 ページ
	page, err := FindPage(ctx, s.db, req.Msg.GetPageId())

	if err != nil {
		return handlePageDeleteError(tx, "削除対象のページ検索に失敗しました", fmt.Errorf("Error delete to find page: %v\n", err))
	}

	// 削除対象検索 ページ
	sections, err := Sections(SectionWhere.PageID.EQ(req.Msg.GetPageId())).All(ctx, tx)
	if err != nil {
		return handlePageDeleteError(tx, "削除対象のセクション検索に失敗しました", fmt.Errorf("Error delete page to find sections: %v\n", err))
	}

	// 削除実施 セクション
	if sections != nil {
		_, err = sections.DeleteAll(ctx, tx)
	}
	if err != nil {
		return handlePageDeleteError(tx, "セクション削除に失敗しました", fmt.Errorf("Error delete page to delete sections: %v\n", err))
	}

	// 削除実施 ページ
	_, err = page.Delete(ctx, s.db)
	if err != nil {
		return handlePageDeleteError(tx, "ページ削除に失敗しました", fmt.Errorf("Error delete page: %v\n", err))
	}

	res := connect.NewResponse(&dynamicv1.DeletePageResponse{
		Status:  "success",
		Message: "ページが削除されました",
	})

	return res, nil
}
func handlePageDeleteError(tx *sql.Tx, msg string, err error) (*connect.Response[dynamicv1.DeletePageResponse], error) {

	_ = tx.Rollback()
	res := connect.NewResponse(&dynamicv1.DeletePageResponse{
		Status:  "failure",
		Message: msg,
	})
	return res, connect.NewError(connect.CodeInternal, err)
}

func (s *pageServer) UpdatePageName(
	ctx context.Context,
	req *connect.Request[dynamicv1.UpdatePageNameRequest],
) (*connect.Response[dynamicv1.UpdatePageNameResponse], error) {

	page, err := FindPage(ctx, s.db, req.Msg.GetPageId())

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("Error update to find page: %v\n", err))
	}
	page.Name = req.Msg.GetName()

	pageUpdateColumns := boil.Whitelist(PageColumns.Name)
	_, err = page.Update(ctx, s.db, pageUpdateColumns)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("Error update page: %v\n", err))
	}

	res := connect.NewResponse(&dynamicv1.UpdatePageNameResponse{
		PageId:      page.PageID,
		UpdatedTime: page.UpdatedAt.Format(time.RFC3339),
	})

	return res, nil
}
