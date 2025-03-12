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

type chapterServer struct {
	dynamicv1connect.ChapterServiceHandler
	db boil.ContextExecutor
}

func NewChapterServer() *chapterServer {
	return &chapterServer{
		db: GetDB(),
	}
}

func (s *chapterServer) AddChapter(
	ctx context.Context,
	req *connect.Request[dynamicv1.AddChapterRequest],
) (*connect.Response[dynamicv1.AddChapterResponse], error) {

	// ページの最大order値を取得
	chapters, err := Chapters(
		ChapterWhere.DynamicID.EQ(req.Msg.GetDynamicId()),
		OrderBy("\""+ChapterColumns.Order+"\" DESC"),
	).All(ctx, s.db)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("Error fetching max order for chapters: %v", err))
	}

	var maxOrder int32
	if len(chapters) == 0 {
		maxOrder = 0
	} else {
		maxOrder = chapters[0].Order
	}

	// ユニーク名を生成
	baseName := "新しい章"
	maxNumber := 1

	namePattern := regexp.MustCompile(fmt.Sprintf(`^%s(?: (\d+))?$`, regexp.QuoteMeta(baseName)))

	for _, chapter := range chapters {
		matches := namePattern.FindStringSubmatch(chapter.Name)
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
	chapter := &models.Chapter{
		DynamicID: req.Msg.GetDynamicId(),
		Name:      uniqueName,
		Order:     maxOrder + 1,
	}

	err = chapter.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("Error Add to chapter: %v", err))
	}

	res := connect.NewResponse(&dynamicv1.AddChapterResponse{
		ChapterId:   chapter.ID,
		Name:        chapter.Name,
		Order:       chapter.Order,
		CreatedTime: chapter.CreatedAt.Format(time.RFC3339),
	})

	return res, nil
}

func (s *chapterServer) DeleteChapter(
	ctx context.Context,
	req *connect.Request[dynamicv1.DeleteChapterRequest],
) (*connect.Response[dynamicv1.DeleteChapterResponse], error) {

	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		res := connect.NewResponse(&dynamicv1.DeleteChapterResponse{
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

	// 削除対象検索 章
	chapter, err := FindChapter(ctx, tx, req.Msg.GetChapterId())
	if err != nil {
		return handleChapterDeleteError(tx, "削除対象の章検索に失敗しました", fmt.Errorf("Error delete to find chapter: %w", err))
	}

	// 削除対象検索 ページ
	pages, err := Pages(PageWhere.ChapterID.EQ(req.Msg.GetChapterId())).All(ctx, tx)
	if err != nil {
		return handleChapterDeleteError(tx, "削除対象のページ検索に失敗しました", fmt.Errorf("Error delete chapter to find pages: %v\n", err))
	}

	// 削除対象検索 セクション
	var sections SectionSlice
	if pages != nil {
		var modifiers []QueryMod
		for _, page := range pages {
			modifiers = append(modifiers, SectionWhere.PageID.EQ(page.ID))
		}
		sections, err = Sections(modifiers...).All(ctx, tx)
		if err != nil {
			return handleChapterDeleteError(tx, "削除対象のセクション検索に失敗しました", fmt.Errorf("Error delete chapter to find sections: %v\n", err))
		}
	}

	// 削除実施 セクション
	if sections != nil {
		_, err = sections.DeleteAll(ctx, tx)
		if err != nil {
			return handleChapterDeleteError(tx, "セクション削除に失敗しました", fmt.Errorf("Error delete chapter to delete sections: %v\n", err))
		}
	}

	// 削除実施 ページ
	if pages != nil {
		_, err = pages.DeleteAll(ctx, tx)
		if err != nil {
			return handleChapterDeleteError(tx, "ページ削除に失敗しました", fmt.Errorf("Error delete chapter to delete pages: %v\n", err))
		}
	}

	// 削除実施 章
	_, err = chapter.Delete(ctx, tx)
	if err != nil {
		return handleChapterDeleteError(tx, "章削除に失敗しました", fmt.Errorf("Error delete chapter: %v\n", err))
	}

	res := connect.NewResponse(&dynamicv1.DeleteChapterResponse{
		Status:  "success",
		Message: "章が削除されました",
	})

	return res, nil
}

func handleChapterDeleteError(tx *sql.Tx, msg string, err error) (*connect.Response[dynamicv1.DeleteChapterResponse], error) {

	_ = tx.Rollback()
	res := connect.NewResponse(&dynamicv1.DeleteChapterResponse{
		Status:  "failure",
		Message: msg,
	})
	return res, connect.NewError(connect.CodeInternal, err)
}

func (s *chapterServer) UpdateChapterName(
	ctx context.Context,
	req *connect.Request[dynamicv1.UpdateChapterNameRequest],
) (*connect.Response[dynamicv1.UpdateChapterNameResponse], error) {

	chapter, err := FindChapter(ctx, s.db, req.Msg.GetChapterId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("Error delete to find chapter: %v\n", err))
	}

	chapter.Name = req.Msg.GetName()
	chapterUpdateColumns := boil.Whitelist(ChapterColumns.Name)
	_, err = chapter.Update(ctx, s.db, chapterUpdateColumns)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("Error update chapter: %v\n", err))
	}

	res := connect.NewResponse(&dynamicv1.UpdateChapterNameResponse{
		ChapterId:   chapter.ID,
		UpdatedTime: chapter.UpdatedAt.Format(time.RFC3339),
	})

	return res, nil
}
