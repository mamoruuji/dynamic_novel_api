package server

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"

	. "github.com/mamoruuji/dynamic_novel_api/config"

	"github.com/bufbuild/connect-go"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
	"github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1/dynamicv1connect"

	. "github.com/mamoruuji/dynamic_novel_api/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type dynamicServer struct {
	dynamicv1connect.DynamicServiceHandler
	db boil.ContextExecutor
}

func NewDynamicServer() *dynamicServer {
	return &dynamicServer{
		db: GetDB(),
	}
}

func (s *dynamicServer) GetDynamic(
	ctx context.Context,
	req *connect.Request[dynamicv1.GetDynamicRequest],
) (*connect.Response[dynamicv1.GetDynamicResponse], error) {
	modifiers := []QueryMod{
		Load(DynamicRels.User),
		Load(DynamicRels.Terms, OrderBy(TermTableColumns.Order)),
		Load(DynamicRels.Chapters, OrderBy(ChapterTableColumns.Order)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages), OrderBy(PageTableColumns.Order)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Terms), OrderBy(TermTableColumns.Order)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections), OrderBy(SectionTableColumns.Order)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Terms), OrderBy(TermTableColumns.Order)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.TypeOfSection)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.TypeOfAnimation)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.TypeOfPosition)),
		// Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.Image)),
		// Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.Image, ImageRels.TypeOfImage)),
		// Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.Image, ImageRels.User)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.TypeOfFont)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.FrameColor)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.TextColor)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.Terms), OrderBy(TermTableColumns.Order)),
		DynamicWhere.DynamicID.EQ(req.Msg.DynamicId),
	}

	dynamic, err := Dynamics(modifiers...).One(ctx, s.db)

	if err != nil {
		log.Printf("failed to get dynamics: %v", err)
		return nil, err
	}

	var pbChapters []*dynamicv1.ChapterData
	for _, chapter := range dynamic.R.Chapters {
		var pbPages []*dynamicv1.PageData
		for _, page := range chapter.R.Pages {
			var pbSections []*dynamicv1.SectionData
			for _, section := range page.R.Sections {

				pbTerms := SetTermData(section.R.GetTerms())

				pbSection := &dynamicv1.SectionData{
					SectionId:     section.SectionID,
					Name:          section.Name,
					Order:         section.Order,
					TypeSection:   section.R.GetTypeOfSection().Name,
					TypePosition:  section.R.GetTypeOfPosition().Name,
					TypeAnimation: section.R.GetTypeOfAnimation().Name,
					FrameColor:    section.R.GetFrameColor().Name,
					Text:          section.Text,
					TextColor:     section.R.GetTextColor().Name,
					TextSize:      section.TextSize,
					Font:          section.R.GetTypeOfFont().Name,
					ImageUrl:      section.ImageURL,
					Terms:         pbTerms,
				}
				pbSections = append(pbSections, pbSection)
			}
			pbTerms := SetTermData(page.R.GetTerms())

			pbPage := &dynamicv1.PageData{
				PageId:   page.PageID,
				Name:     page.Name,
				Order:    page.Order,
				Sections: pbSections,
				Terms:    pbTerms,
			}
			pbPages = append(pbPages, pbPage)
		}
		pbTerms := SetTermData(chapter.R.GetTerms())

		pbChapter := &dynamicv1.ChapterData{
			ChapterId: chapter.ChapterID,
			Name:      chapter.Name,
			Order:     chapter.Order,
			Pages:     pbPages,
			Terms:     pbTerms,
		}
		pbChapters = append(pbChapters, pbChapter)
	}
	pbTerms := SetTermData(dynamic.R.GetTerms())
	createdAT := timestamppb.New(dynamic.CreatedAt)
	updatedAT := timestamppb.New(dynamic.UpdatedAt)

	pbTags, _ := SetTagData(req.Msg.DynamicId, ctx, s.db)
	res := connect.NewResponse(&dynamicv1.GetDynamicResponse{
		DynamicId:   dynamic.DynamicID,
		Name:        dynamic.Name,
		Overview:    dynamic.Overview,
		UserId:      dynamic.UserID,
		Published:   dynamic.Published,
		PenName:     dynamic.R.GetUser().PenName,
		ImageUrl:    dynamic.ImageURL,
		Chapters:    pbChapters,
		Terms:       pbTerms,
		Tags:        pbTags,
		CreatedTime: createdAT,
		UpdatedTime: updatedAT,
	})

	return res, nil
}

func (s *dynamicServer) ListDynamics(
	ctx context.Context,
	req *connect.Request[dynamicv1.ListDynamicsRequest],
) (*connect.Response[dynamicv1.ListDynamicsResponse], error) {
	sortCategory, err := FindTypeOfSort(ctx, s.db, req.Msg.SortCategory)

	if err != nil {
		log.Printf("failed to get sorts: %v", err)
		return nil, err
	}

	modifiers := []QueryMod{
		Load(DynamicRels.User),
		Load(DynamicRels.DynamicsOnTags),
		Load(Rels(DynamicRels.DynamicsOnTags, DynamicsOnTagRels.Tag)),
		// LeftOuterJoin(TableNames.Users + " on " + TableNames.Dynamics + "." + DynamicColumns.UserID + " = " + TableNames.Users + "." + UserColumns.UserID),
		// LeftOuterJoin(TableNames.DynamicsOnTags + " on " + TableNames.Dynamics + "." + DynamicColumns.DynamicID + " = " + TableNames.DynamicsOnTags + "." + DynamicsOnTagColumns.DynamicID),
		// LeftOuterJoin(TableNames.Tags + " on " + TableNames.DynamicsOnTags + "." + DynamicsOnTagColumns.DynamicID + " = " + TableNames.Tags + "." + TagColumns.TagID),
		OrderBy(fmt.Sprintf("%s %s", sortCategory.SQL, req.Msg.SortOrder)),
	}

	if req.Msg.UserId != "" {
		modifiers = append(modifiers, DynamicWhere.UserID.EQ(req.Msg.UserId))
	} else {
		modifiers = append(modifiers, DynamicWhere.Published.EQ(true))
	}

	if len(req.Msg.SearchKeywords) != 0 {
		// var orConditions, Dwhere, Uwhere, Twhere []QueryMod
		// var orConditions, Dwhere []QueryMod
		var orConditions, Dwhere, Uwhere []QueryMod
		for key, searchKeyword := range req.Msg.SearchKeywords {
			keyword := "%" + searchKeyword + "%"
			if key == 0 {
				Dwhere = append(Dwhere, DynamicWhere.Name.ILIKE(keyword))
				// Uwhere = append(Uwhere, UserWhere.Name.ILIKE(keyword))
				// Twhere = append(Twhere, Expr(TagWhere.Name.ILIKE(keyword), Or2(Where(TagColumns.Name+" IS NULL"))))
			} else {
				Dwhere = append(Dwhere, Or2(DynamicWhere.Name.ILIKE(keyword)))
				// Uwhere = append(Uwhere, Or2(UserWhere.Name.ILIKE(keyword)))
				// Twhere = append(Twhere, Or2(TagWhere.Name.ILIKE(keyword)))
			}
			// LeftOuterJoin用の条件追加
			// Dwhere = append(Dwhere, Or2(UserWhere.Name.ILIKE(keyword)))
			// Dwhere = append(Dwhere, Or2(TagWhere.Name.ILIKE(keyword)))
		}
		orConditions = append(orConditions, Expr(Dwhere...))
		orConditions = append(orConditions, Load(DynamicRels.User, Expr(Uwhere...)))
		// orConditions = append(orConditions, Load(DynamicRels.User))
		// orConditions = append(orConditions, Load(DynamicRels.DynamicsOnTags)) 不要
		// orConditions = append(orConditions, Load(Rels(DynamicRels.DynamicsOnTags, DynamicsOnTagRels.Tag), Expr(Twhere...)))
		// orConditions = append(orConditions, Load(Rels(DynamicRels.DynamicsOnTags, DynamicsOnTagRels.Tag)))
		modifiers = append(modifiers, orConditions...)
	}

	if len(req.Msg.FilterKeywords) != 0 {
		// var orConditions, Dwhere, Uwhere, Twhere []QueryMod
		var orConditions, Dwhere []QueryMod
		for _, filterKeyword := range req.Msg.FilterKeywords {
			keyword := "%" + filterKeyword + "%"
			Dwhere = append(Dwhere, DynamicWhere.Name.NILIKE(keyword))
			// Uwhere = append(Uwhere, UserWhere.Name.NILIKE(keyword))
			// Twhere = append(Twhere, TagWhere.Name.NILIKE(keyword))
			// LeftOuterJoin用の条件追加
			// Dwhere = append(Dwhere, UserWhere.Name.NILIKE(keyword))
			// Dwhere = append(Dwhere, Expr(TagWhere.Name.NILIKE(keyword), Or2(Where(TableNames.Tags+"."+TagColumns.Name+" IS NULL"))))
		}
		orConditions = append(orConditions, Expr(Dwhere...))
		// orConditions = append(orConditions, Load(DynamicRels.User, Expr(Uwhere...)))
		// orConditions = append(orConditions, Load(Rels(DynamicRels.DynamicsOnTags, DynamicsOnTagRels.Tag), Expr(Twhere...)))
		modifiers = append(modifiers, orConditions...)
	}

	if req.Msg.FilterStartDate != "" {
		condition := DynamicWhere.UpdatedAt.GTE(ConvertToPostgresTimestamp(req.Msg.FilterStartDate))
		modifiers = append(modifiers, condition)
	}

	if req.Msg.FilterEndDate != "" {
		condition := DynamicWhere.UpdatedAt.LT(ConvertToPostgresTimestamp(req.Msg.FilterEndDate))
		modifiers = append(modifiers, condition)
	}

	// boil.DebugMode = true
	dynamics, err := Dynamics(modifiers...).All(ctx, s.db)
	// spew.Dump(dynamics)

	if err != nil {
		log.Printf("failed to get dynamics: %v", err)
		return nil, err
	}

	var pbDynamics []*dynamicv1.ListDynamicData
	for _, dynamic := range dynamics {
		pbTags, _ := SetTagData(dynamic.DynamicID, ctx, s.db)
		createdAT := timestamppb.New(dynamic.CreatedAt)
		updatedAT := timestamppb.New(dynamic.UpdatedAt)
		pbDynamic := &dynamicv1.ListDynamicData{
			DynamicId:   dynamic.DynamicID,
			Name:        dynamic.Name,
			Overview:    dynamic.Overview,
			PenName:     dynamic.R.GetUser().PenName,
			ImageUrl:    dynamic.ImageURL,
			Tags:        pbTags,
			CreatedTime: createdAT,
			UpdatedTime: updatedAT,
		}
		pbDynamics = append(pbDynamics, pbDynamic)
	}

	res := connect.NewResponse(&dynamicv1.ListDynamicsResponse{
		Dynamics: pbDynamics,
	})

	return res, nil
}

func (s *dynamicServer) UpdateContentsOrder(
	ctx context.Context,
	req *connect.Request[dynamicv1.UpdateContentsOrderRequest],
) (*connect.Response[dynamicv1.UpdateContentsOrderResponse], error) {
	chapters := req.Msg.Chapters

	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to start transaction: %w", err))
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

	pageUpdateColumns := boil.Whitelist(PageColumns.ChapterID, PageColumns.Order)
	chapterUpdateColumns := boil.Whitelist(ChapterColumns.Order)
	for i, chapter := range chapters {
		for j, page := range chapter.Pages {
			pageOrder := int32(j + 1)

			_, err = (&Page{
				PageID:    page.PageId,
				ChapterID: chapter.ChapterId,
				Order:     pageOrder,
			}).Update(ctx, tx, pageUpdateColumns)
			if err != nil {
				return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to update page: %w", err))
			}
		}
		chapterOrder := int32(i + 1)
		_, err = (&Chapter{
			ChapterID: chapter.ChapterId,
			Order:     chapterOrder,
		}).Update(ctx, tx, chapterUpdateColumns)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to update chapter: %w", err))
		}
	}
	return connect.NewResponse(&dynamicv1.UpdateContentsOrderResponse{}), nil
}
