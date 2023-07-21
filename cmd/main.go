package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/amomon/dynamic_novel_api/api/dynamic"
	"github.com/amomon/dynamic_novel_api/db/models"
	"google.golang.org/grpc"

	// . "github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type dynamicServer struct {
	dynamic.UnimplementedDynamicServer
}

func (s *dynamicServer) GetDynamics(ctx context.Context, _ *dynamic.GetDynamicsRequest) (*dynamic.GetDynamicsResponse, error) {
	dynamics, err := models.Dynamics().All(ctx, boil.GetContextDB())
	if err != nil {
		log.Printf("failed to get dynamics: %v", err)
		return nil, err
	}

	var pbDynamics []*dynamic.DynamicData
	for _, d := range dynamics {
		pbDynamic := &dynamic.DynamicData{
			Id:        int32(d.DynamicID),
			Title:     d.Title,
			Overview:  d.Overview,
			UserId:    d.UserID,
			Published: d.Published,
		}
		pbDynamics = append(pbDynamics, pbDynamic)
	}

	return &dynamic.GetDynamicsResponse{
		Dynamics: pbDynamics,
	}, nil
}

func (s *dynamicServer) AddDynamic(ctx context.Context, req *dynamic.AddDynamicRequest) (*dynamic.AddDynamicResponse, error) {
	d := models.Dynamic{
		Title:     req.Title,
		UserID:    req.UserId,
		Published: false,
	}
	err := d.Insert(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		log.Printf("failed to add dynamic: %v", err)
		return nil, err
	}

	return &dynamic.AddDynamicResponse{
		Id: int32(d.DynamicID),
	}, nil
}

func (s *dynamicServer) DeleteDynamic(ctx context.Context, req *dynamic.DeleteDynamicRequest) (*dynamic.DeleteDynamicResponse, error) {
	dynamicID := int(req.Id)
	d := models.Dynamic{
		DynamicID: dynamicID,
	}

	_, err := d.Delete(ctx, boil.GetContextDB())
	if err != nil {
		log.Printf("failed to delete dynamic: %v", err)
		return nil, err
	}

	return &dynamic.DeleteDynamicResponse{}, nil
}

func (s *dynamicServer) UpdateDynamicStatus(ctx context.Context, req *dynamic.UpdateDynamicStatusRequest) (*dynamic.UpdateDynamicStatusResponse, error) {
	dynamicID := int(req.Id)
	d := models.Dynamic{
		DynamicID: dynamicID,
	}

	d.Title = req.Title
	d.Overview = req.Overview
	d.Published = req.Published
	_, err := d.Update(ctx, boil.GetContextDB(), boil.Infer())
	if err != nil {
		log.Printf("failed to update dynamic status: %v", err)
		return nil, err
	}

	return &dynamic.UpdateDynamicStatusResponse{}, nil
}

func main() {
	db, err := sql.Open("postgres", "host=db port=5432 user=postgres password=pass dbname=dynamic_novel sslmodedisable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)

	server := grpc.NewServer()

	dynamicServer := &dynamicServer{}
	dynamic.RegisterDynamicServer(server, dynamicServer)

	// r := router.Group("/dynamic")
	// {
	// 	r.GET("/select", func(c *gin.Context) {
	// 		GetDynamics(ctx, c)
	// 	})
	// 	r.POST("/insert", func(c *gin.Context) {
	// 		AddDynamic(ctx, c)
	// 	})
	// 	r.DELETE("/delete/:id", func(c *gin.Context) {
	// 		DeleteDynamic(ctx, c)
	// 	})
	// 	r.PUT("/update/:id", func(c *gin.Context) {
	// 		UpdateDynamicStatus(ctx, c)
	// 	})
	// }

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err = server.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

// func dieIf(err error, message string){
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// }
