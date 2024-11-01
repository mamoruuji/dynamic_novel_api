package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	_ "github.com/lib/pq"
	"github.com/mamoruuji/dynamic_novel_api/db/models"
	. "github.com/mamoruuji/dynamic_novel_api/db/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/bufbuild/connect-go"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/davecgh/go-spew/spew"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"        // generated by protoc-gen-go
	"github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1/dynamicv1connect" // generated by protoc-gen-connect-go
)

type dynamicServer struct {
	dynamicv1connect.DynamicServiceHandler
}

type pageServer struct {
	dynamicv1connect.PageServiceHandler
}

type sortServer struct {
	dynamicv1connect.SortServiceHandler
}

type imageServer struct {
	dynamicv1connect.ImageServiceHandler
}

type userServer struct {
	dynamicv1connect.UserServiceHandler
}

func (s *dynamicServer) GetDynamic(
	ctx context.Context,
	req *connect.Request[dynamicv1.GetDynamicRequest],
) (*connect.Response[dynamicv1.GetDynamicResponse], error) {

	modifiers := []QueryMod{
		Load(DynamicRels.User),
		Load(DynamicRels.Image),
		Load(Rels(DynamicRels.Image, ImageRels.TypeOfImage)),
		// Load(Rels(DynamicRels.Image, ImageRels.User)),
		Load(DynamicRels.Terms, OrderBy(TermTableColumns.Order)),
		Load(DynamicRels.Chapters, OrderBy(ChapterTableColumns.Order)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages), OrderBy(PageTableColumns.Order)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Terms), OrderBy(TermTableColumns.Order)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections), OrderBy(SectionTableColumns.Order)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Terms), OrderBy(TermTableColumns.Order)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.TypeOfSection)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.TypeOfAnimation)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.TypeOfPosition)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.Image)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.Image, ImageRels.TypeOfImage)),
		// Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.Image, ImageRels.User)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.TypeOfFont)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.FrameColor)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.TextColor)),
		Load(Rels(DynamicRels.Chapters, ChapterRels.Pages, PageRels.Sections, SectionRels.Terms), OrderBy(TermTableColumns.Order)),
		DynamicWhere.DynamicID.EQ(req.Msg.DynamicId),
	}

	dynamic, err := Dynamics(modifiers...).One(ctx, db)

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

				pbImage := SetImageData(section.R.Image)
				pbTerms := SetTermData(section.R.Terms)

				pbSection := &dynamicv1.SectionData{
					SectionId:     section.SectionID,
					Name:          section.Name,
					Order:         section.Order,
					TypeSection:   section.R.TypeOfSection.Name,
					TypePosition:  section.R.TypeOfPosition.Name,
					TypeAnimation: section.R.TypeOfAnimation.Name,
					FrameColor:    section.R.FrameColor.Name,
					Text:          section.Text,
					TextColor:     section.R.TextColor.Name,
					TextSize:      section.TextSize,
					Font:          section.R.TypeOfFont.Name,
					Image:         pbImage,
					Terms:         pbTerms,
				}
				pbSections = append(pbSections, pbSection)
			}
			pbTerms := SetTermData(page.R.Terms)

			pbPage := &dynamicv1.PageData{
				PageId:   page.PageID,
				Title:    page.Title,
				Order:    page.Order,
				Sections: pbSections,
				Terms:    pbTerms,
			}
			pbPages = append(pbPages, pbPage)
		}
		pbTerms := SetTermData(chapter.R.Terms)

		pbChapter := &dynamicv1.ChapterData{
			ChapterId: chapter.ChapterID,
			Title:     chapter.Title,
			Order:     chapter.Order,
			Pages:     pbPages,
			Terms:     pbTerms,
		}
		pbChapters = append(pbChapters, pbChapter)
	}
	pbTerms := SetTermData(dynamic.R.Terms)
	createdAT := timestamppb.New(dynamic.CreatedAt)
	updatedAT := timestamppb.New(dynamic.UpdatedAt)

	pbTags := SetTagData(req.Msg.DynamicId, ctx)
	pbImage := SetImageData(dynamic.R.Image)
	res := connect.NewResponse(&dynamicv1.GetDynamicResponse{
		DynamicId:   dynamic.DynamicID,
		Title:       dynamic.Title,
		Overview:    dynamic.Overview,
		UserId:      dynamic.UserID,
		Published:   dynamic.Published,
		UserName:    dynamic.R.User.Name,
		Image:       pbImage,
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
	sortCategory, err := FindTypeOfSort(ctx, db, req.Msg.SortCategory)

	if err != nil {
		log.Printf("failed to get sorts: %v", err)
		return nil, err
	}

	modifiers := []QueryMod{
		Load(DynamicRels.User),
		Load(DynamicRels.Image),
		Load(Rels(DynamicRels.Image, ImageRels.TypeOfImage)),
		// Load(DynamicRels.DynamicsOnTags),
		// Load(Rels(DynamicRels.DynamicsOnTags, DynamicsOnTagRels.Tag)),
		// LeftOuterJoin(TableNames.Users + " on " + TableNames.Dynamics + "." + DynamicColumns.UserID + " = " + TableNames.Users + "." + UserColumns.UserID),
		// LeftOuterJoin(TableNames.DynamicsOnTags + " on " + TableNames.Dynamics + "." + DynamicColumns.DynamicID + " = " + TableNames.DynamicsOnTags + "." + DynamicsOnTagColumns.DynamicID),
		// LeftOuterJoin(TableNames.Tags + " on " + TableNames.DynamicsOnTags + "." + DynamicsOnTagColumns.DynamicID + " = " + TableNames.Tags + "." + TagColumns.TagID),
		DynamicWhere.Published.EQ(true),
		OrderBy(fmt.Sprintf("%s %s", sortCategory.SQL, req.Msg.SortOrder)),
	}

	if req.Msg.UserId != "" {
		condition := DynamicWhere.UserID.EQ(req.Msg.UserId)
		modifiers = append(modifiers, condition)
	}

	if len(req.Msg.SearchKeywords) != 0 {
		// var orConditions, Dwhere, Uwhere, Twhere []QueryMod
		var orConditions, Dwhere []QueryMod
		for key, searchKeyword := range req.Msg.SearchKeywords {
			keyword := "%" + searchKeyword + "%"
			if key == 0 {
				Dwhere = append(Dwhere, DynamicWhere.Title.ILIKE(keyword))
				// Uwhere = append(Uwhere, UserWhere.Name.ILIKE(keyword))
				// Twhere = append(Twhere, Expr(TagWhere.Name.ILIKE(keyword), Or2(Where(TagColumns.Name+" IS NULL"))))
			} else {
				Dwhere = append(Dwhere, Or2(DynamicWhere.Title.ILIKE(keyword)))
				// Uwhere = append(Uwhere, Or2(UserWhere.Name.ILIKE(keyword)))
				// Twhere = append(Twhere, Or2(TagWhere.Name.ILIKE(keyword)))
			}
			// LeftOuterJoin用の条件追加
			// Dwhere = append(Dwhere, Or2(UserWhere.Name.ILIKE(keyword)))
			// Dwhere = append(Dwhere, Or2(TagWhere.Name.ILIKE(keyword)))
		}
		orConditions = append(orConditions, Expr(Dwhere...))
		// orConditions = append(orConditions, Load(DynamicRels.User, Expr(Uwhere...)))
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
			Dwhere = append(Dwhere, DynamicWhere.Title.NILIKE(keyword))
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
	dynamics, err := Dynamics(modifiers...).All(ctx, db)

	if err != nil {
		log.Printf("failed to get dynamics: %v", err)
		return nil, err
	}

	var pbDynamics []*dynamicv1.ListDynamicData
	for _, dynamic := range dynamics {
		pbTags := SetTagData(dynamic.DynamicID, ctx)
		pbImage := SetImageData(dynamic.R.Image)
		// spew.Dump(dynamic.R.User)
		createdAT := timestamppb.New(dynamic.CreatedAt)
		updatedAT := timestamppb.New(dynamic.UpdatedAt)
		pbDynamic := &dynamicv1.ListDynamicData{
			DynamicId: dynamic.DynamicID,
			Title:     dynamic.Title,
			Overview:  dynamic.Overview,
			// UserName:    dynamic.R.User.Name,
			Image:       pbImage,
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

func (s *pageServer) ListPages(
	ctx context.Context,
	_ *connect.Request[dynamicv1.ListPagesRequest],
) (*connect.Response[dynamicv1.ListPagesResponse], error) {
	pages, err := models.Pages().All(ctx, db)
	if err != nil {
		log.Printf("failed to get pages: %v", err)
		return nil, err
	}

	var pbPages []*dynamicv1.ListPageData
	for _, page := range pages {
		pbPage := &dynamicv1.ListPageData{
			PageId:    page.PageID,
			Title:     page.Title,
			Order:     page.Order,
			ChapterId: page.ChapterID,
		}
		pbPages = append(pbPages, pbPage)
	}

	res := connect.NewResponse(&dynamicv1.ListPagesResponse{
		Pages: pbPages,
	})

	return res, nil
}

func (s *sortServer) ListSorts(
	ctx context.Context,
	_ *connect.Request[dynamicv1.ListSortsRequest],
) (*connect.Response[dynamicv1.ListSortsResponse], error) {
	sorts, err := TypeOfSorts().All(ctx, db)
	if err != nil {
		log.Printf("failed to get sorts: %v", err)
		return nil, err
	}

	var pbSorts []*dynamicv1.SortData
	for _, sort := range sorts {
		pbSort := &dynamicv1.SortData{
			SortId: sort.TypeOfSortID,
			Name:   sort.Name,
		}
		pbSorts = append(pbSorts, pbSort)
	}

	res := connect.NewResponse(&dynamicv1.ListSortsResponse{
		Sorts: pbSorts,
	})

	return res, nil
}

func (s *dynamicServer) AddDynamic(
	ctx context.Context,
	req *connect.Request[dynamicv1.AddDynamicRequest],
) (*connect.Response[dynamicv1.AddDynamicResponse], error) {
	dynamic := models.Dynamic{
		Title:     req.Msg.Title,
		UserID:    req.Msg.UserId,
		Published: false,
	}
	err := dynamic.Insert(ctx, db, boil.Infer())
	if err != nil {
		log.Printf("failed to add dynamic: %v", err)
		return nil, err
	}

	res := connect.NewResponse(&dynamicv1.AddDynamicResponse{
		DynamicId: dynamic.DynamicID,
	})

	return res, nil
}

func (s *dynamicServer) DeleteDynamic(
	ctx context.Context,
	req *connect.Request[dynamicv1.DeleteDynamicRequest],
) (*connect.Response[dynamicv1.DeleteDynamicResponse], error) {
	dynamic := models.Dynamic{
		DynamicID: req.Msg.DynamicId,
	}

	_, err := dynamic.Delete(ctx, db)
	if err != nil {
		log.Printf("failed to delete dynamic: %v", err)
		return nil, err
	}

	return connect.NewResponse(&dynamicv1.DeleteDynamicResponse{}), nil
}

func (s *dynamicServer) UpdateDynamicStatus(
	ctx context.Context,
	req *connect.Request[dynamicv1.UpdateDynamicStatusRequest],
) (*connect.Response[dynamicv1.UpdateDynamicStatusResponse], error) {
	// 作品を作成
	dynamic := models.Dynamic{
		DynamicID: req.Msg.DynamicId,
		Title:     req.Msg.Title,
		Overview:  req.Msg.Overview,
		Published: req.Msg.Published,
	}
	_, err := dynamic.Update(ctx, db, boil.Infer())
	if err != nil {
		log.Printf("failed to update dynamic status: %v", err)
		return nil, err
	}

	// // 表紙画像を追加・更新
	// // 表紙画像がない場合 新規追加
	// // 既に表紙画像を持っている場合 更新
	// image := models.Image{
	// 	ImageID: req.Msg.ImageId,
	// 	Name:    req.Msg.imageName,
	// }
	// _, err := image.Upsert(ctx, db, boil.Infer())
	// if err != nil {
	// 	log.Printf("failed to update image status: %v", err)
	// 	return nil, err
	// }

	// // タグを追加・更新
	// // タグがない場合 新規追加
	// // 既に存在しているタグを付与する場合 更新
	// tag := models.Tag{
	// 	Title:     req.Msg.Title,
	// 	Overview:  req.Msg.Overview,
	// 	Published: req.Msg.Published,
	// }
	// _, err := tag.Update(ctx, db, boil.Infer())
	// if err != nil {
	// 	log.Printf("failed to update tag status: %v", err)
	// 	return nil, err
	// }

	return connect.NewResponse(&dynamicv1.UpdateDynamicStatusResponse{}), nil
}

func (s *imageServer) ListUserImages(
	ctx context.Context,
	req *connect.Request[dynamicv1.ListUserImagesRequest],
) (*connect.Response[dynamicv1.ListUserImagesResponse], error) {
	boil.DebugMode = true
	rootFolders, err := Folders(
		Load(FolderRels.Images),
		Load(Rels(FolderRels.Images, ImageRels.TypeOfImage)),
		FolderWhere.UserID.EQ(req.Msg.UserId),
		OrderBy(FolderColumns.FolderID),
	).All(ctx, db)

	if err != nil {
		log.Printf("failed to get images: %v", err)
		return nil, err
	}
	// fmt.Println(rootFolders)
	// フォルダIDをキーにしたマップを作成
	folderMap := GetFoldersMap(rootFolders)

	// 親IDが0（ルート）のフォルダから階層構造を構築
	folderTree := BuildTree(folderMap, 0)
	// spew.Dump(folderTree)
	fmt.Println(folderTree)
	for _, folder := range folderTree {
		spew.Dump(folder.Images)
	}

	rootImages, err := Images(
		Load(ImageRels.TypeOfImage),
		ImageWhere.UserID.EQ(req.Msg.UserId),
		ImageWhere.FolderID.IsNull(),
		OrderBy(ImageColumns.ImageID),
	).All(ctx, db)

	if err != nil {
		log.Printf("failed to get images: %v", err)
		return nil, err
	}
	fmt.Println(rootImages)

	var folders []*dynamicv1.FolderData
	for _, root := range rootFolders {
		folders = append(folders, convertFolder(root))
	}
	pbImages := SetImagesData(rootImages)

	res := connect.NewResponse(&dynamicv1.ListUserImagesResponse{
		Folders: folders,
		Images:  pbImages,
	})

	return res, nil
}

// GetFoldersMap creates a map of folder_id to Folder
func GetFoldersMap(folders []*models.Folder) map[int]*models.Folder {
	folderMap := make(map[int]*models.Folder)
	for _, folder := range folders {
		folderMap[int(folder.FolderID)] = folder
	}
	return folderMap
}

// BuildTree recursively builds the folder tree
func BuildTree(folderMap map[int]*models.Folder, parentID int) []*dynamicv1.FolderData {
	var result []*dynamicv1.FolderData
	for _, folder := range folderMap {
		if int(folder.ParentID.Int32) == parentID {
			pbImages := SetImagesData(folder.R.Images)
			node := &dynamicv1.FolderData{
				FolderId: folder.FolderID,
				Name:     folder.Name,
				UserId:   folder.UserID,
				ParentId: folder.FolderID,
				Children: BuildTree(folderMap, int(folder.FolderID)),
				Images:   pbImages,
			}
			result = append(result, node)
		}
	}
	return result
}

func convertFolder(f *models.Folder) *dynamicv1.FolderData {
	var children []*dynamicv1.FolderData
	for _, child := range f.R.ParentFolders {
		children = append(children, convertFolder(child))
	}

	var pbImages []*dynamicv1.ImageData
	for _, image := range f.R.Images {
		pbImage := SetImageData(image)
		pbImages = append(pbImages, pbImage)
	}

	return &dynamicv1.FolderData{
		FolderId: f.FolderID,
		Name:     f.Name,
		UserId:   f.UserID,
		ParentId: f.ParentID.Int32,
		Children: children,
		Images:   pbImages,
	}
}

func (s *userServer) GetUser(
	ctx context.Context,
	req *connect.Request[dynamicv1.GetUserRequest],
) (*connect.Response[dynamicv1.GetUserResponse], error) {

	modifiers := []QueryMod{
		UserWhere.UserID.EQ(req.Msg.UserId),
	}

	user, err := Users(modifiers...).One(ctx, db)

	if err != nil {
		log.Printf("failed to get user: %v", err)
		return nil, err
	}

	res := connect.NewResponse(&dynamicv1.GetUserResponse{
		UserId: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		Image:  NullStringToEmptyString(user.Image),
	})

	return res, nil
}

var db boil.ContextExecutor

func InitDB() {
	db = boil.GetContextDB()
}

func server() http.Handler {
	InitDB()

	mux := http.NewServeMux()
	path, handler := dynamicv1connect.NewDynamicServiceHandler(&dynamicServer{})
	mux.Handle(path, handler)
	path, handler = dynamicv1connect.NewPageServiceHandler(&pageServer{})
	mux.Handle(path, handler)
	path, handler = dynamicv1connect.NewSortServiceHandler(&sortServer{})
	mux.Handle(path, handler)
	path, handler = dynamicv1connect.NewImageServiceHandler(&imageServer{})
	mux.Handle(path, handler)
	path, handler = dynamicv1connect.NewUserServiceHandler(&userServer{})
	mux.Handle(path, handler)
	return mux
}

func ConvertToPostgresTimestamp(date string) time.Time {
	layout := "2006/01/02"
	datetime, _ := time.Parse(layout, date)
	return datetime
}

func NullStringToEmptyString(str null.String) string {
	if !str.Valid {
		return ""
	}
	return str.String
}

func SetImageData(image *Image) *dynamicv1.ImageData {
	var pbImage *dynamicv1.ImageData
	if image != nil {
		imagePath := image.Path + image.Name
		pbImage = &dynamicv1.ImageData{
			ImageId:   image.ImageID,
			Name:      image.Name,
			ImagePath: imagePath,
			Type:      image.R.TypeOfImage.Name,
		}
	}
	return pbImage
}

func SetImagesData(images []*Image) []*dynamicv1.ImageData {
	var pbImages []*dynamicv1.ImageData
	for _, image := range images {
		pbImage := SetImageData(image)
		pbImages = append(pbImages, pbImage)
	}
	return pbImages
}

func SetTermData(terms []*Term) []*dynamicv1.TermData {
	var pbTerms []*dynamicv1.TermData
	for _, term := range terms {
		pbTermImage := SetImageData(term.R.Image)
		pbTerm := &dynamicv1.TermData{
			TermId: term.TermID,
			Name:   term.Name,
			Text:   term.Text,
			Order:  term.Order,
			Image:  pbTermImage,
		}
		pbTerms = append(pbTerms, pbTerm)
	}

	return pbTerms
}

func SetTagData(dynamicId int32, ctx context.Context) []*dynamicv1.TagData {
	modifiers := []QueryMod{
		Load(DynamicsOnTagRels.Tag),
		DynamicsOnTagWhere.DynamicID.EQ(dynamicId),
	}
	DynamicsOnTags, _ := DynamicsOnTags(modifiers...).All(ctx, db)

	// if err != nil {
	// 	log.Printf("failed to get tags: %v", err)
	// 	return nil, err
	// }
	var pbTags []*dynamicv1.TagData
	for _, dynamicsOnTag := range DynamicsOnTags {

		pbTag := &dynamicv1.TagData{
			TagId: dynamicsOnTag.TagID,
			Name:  dynamicsOnTag.R.Tag.Name,
		}
		pbTags = append(pbTags, pbTag)
	}
	return pbTags
}

func main() {
	db, err := sql.Open("postgres", "host=db port=5432 user=postgres password=pass dbname=dynamic_novel sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)

	mux := server()

	err = http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)

	if err != nil {
		log.Fatalf("failed to listen(tcp, :8080)")
	}

}
