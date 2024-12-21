package server

import (
	"context"
	"log"

	. "github.com/mamoruuji/dynamic_novel_api/config"

	"github.com/bufbuild/connect-go"
	dynamicv1 "github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1"
	"github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1/dynamicv1connect"

	. "github.com/mamoruuji/dynamic_novel_api/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type tagServer struct {
	dynamicv1connect.TagServiceHandler
	db boil.ContextExecutor
}

func NewTagServer(db boil.ContextExecutor) *tagServer {
	return &tagServer{
		db: GetDB(),
	}
}

func (s *tagServer) ListTags(
	ctx context.Context,
	req *connect.Request[dynamicv1.ListTagsRequest],
) (*connect.Response[dynamicv1.ListTagsResponse], error) {
	tags, err := Tags().All(ctx, s.db)
	if err != nil {
		log.Printf("failed to get tags: %v", err)
		return nil, err
	}

	var pbTags []*dynamicv1.TagData
	for _, tag := range tags {
		pbTag := &dynamicv1.TagData{
			TagId: tag.TagID,
			Name:  tag.Name,
		}
		pbTags = append(pbTags, pbTag)
	}

	res := connect.NewResponse(&dynamicv1.ListTagsResponse{
		Tags: pbTags,
	})

	return res, nil
}

func (s *tagServer) SetDynamicOnTag(
	ctx context.Context,
	req *connect.Request[dynamicv1.SetDynamicOnTagRequest],
) (*connect.Response[dynamicv1.SetDynamicOnTagResponse], error) {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("failed to begin transaction: %v", err)
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			log.Printf("transaction rollback due to panic: %v", p)
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
			log.Printf("transaction rollback due to error: %v", err)
		} else {
			err = tx.Commit()
		}
	}()

	if _, err = DynamicsOnTags(DynamicsOnTagWhere.DynamicID.EQ(req.Msg.DynamicId)).DeleteAll(ctx, tx); err != nil {
		log.Printf("failed to reset dynamic-tag association: %v", err)
		return nil, err
	}

	insertColumns := boil.Whitelist(DynamicsOnTagColumns.DynamicID, DynamicsOnTagColumns.TagID)

	for _, name := range req.Msg.TagNames {
		tag := &Tag{Name: name}
		if err = s.upsertTag(ctx, tx, tag); err != nil {
			log.Printf("failed to upsert tag: %v", err)
			return nil, err
		}

		if upsertedTag, err := Tags(TagWhere.Name.EQ(name)).One(ctx, tx); err != nil {
			log.Fatalf("failed to fetch inserted tag: %v", err)
			return nil, err
		} else {
			tag.TagID = upsertedTag.TagID
		}

		dynamicsOnTag := &DynamicsOnTag{
			DynamicID: req.Msg.DynamicId,
			TagID:     tag.TagID,
		}
		if err = dynamicsOnTag.Insert(ctx, tx, insertColumns); err != nil {
			log.Printf("failed to set dynamic-tag association: %v", err)
			return nil, err
		}
	}

	log.Println("Batch operation completed successfully")
	return connect.NewResponse(&dynamicv1.SetDynamicOnTagResponse{}), nil
}

func (s *tagServer) upsertTag(ctx context.Context, tx boil.ContextExecutor, tag *Tag) error {
	conflictColumns := []string{TagColumns.Name}
	insertColumns := boil.Whitelist(TagColumns.Name)
	updateColumns := boil.Whitelist(TagColumns.Name)

	return tag.Upsert(ctx, tx, false, conflictColumns, insertColumns, updateColumns)
}
