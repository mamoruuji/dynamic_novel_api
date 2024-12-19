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

	for _, name := range req.Msg.TagNames {
		tag := &Tag{Name: name}
		if err = s.upsertTag(ctx, tx, tag); err != nil {
			log.Printf("failed to upsert tag: %v", err)
			return nil, err
		}
		if err = s.upsertDynamicsOnTag(ctx, tx, req.Msg.DynamicId, tag.TagID); err != nil {
			log.Printf("failed to upsert dynamic-tag association: %v", err)
			return nil, err
		}
	}

	log.Println("Batch operation completed successfully")
	return connect.NewResponse(&dynamicv1.SetDynamicOnTagResponse{}), nil
}

func (s *tagServer) upsertTag(ctx context.Context, tx boil.ContextExecutor, tag *Tag) error {
	conflictColumns := []string{TagColumns.Name}
	insertColumns := boil.Infer()
	updateColumns := boil.Whitelist()

	return tag.Upsert(ctx, tx, false, conflictColumns, insertColumns, updateColumns)
}

func (s *tagServer) upsertDynamicsOnTag(ctx context.Context, tx boil.ContextExecutor, dynamicID, tagID int32) error {
	conflictColumns := []string{
		DynamicsOnTagColumns.DynamicID,
		DynamicsOnTagColumns.TagID,
	}
	insertColumns := boil.Infer()
	updateColumns := boil.Whitelist()

	dynamicsOnTag := &DynamicsOnTag{
		DynamicID: dynamicID,
		TagID:     tagID,
	}

	return dynamicsOnTag.Upsert(ctx, tx, false, conflictColumns, insertColumns, updateColumns)
}
