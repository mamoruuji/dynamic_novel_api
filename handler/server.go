package handler

import (
	"net/http"

	"github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1/dynamicv1connect"
	"github.com/mamoruuji/dynamic_novel_api/server"
)

func Server() http.Handler {
	mux := http.NewServeMux()

	path, handler := dynamicv1connect.NewDynamicServiceHandler(server.NewDynamicServer())
	mux.Handle(path, handler)
	path, handler = dynamicv1connect.NewPageServiceHandler(server.NewPageServer())
	mux.Handle(path, handler)
	path, handler = dynamicv1connect.NewSortServiceHandler(server.NewSortServer())
	mux.Handle(path, handler)
	path, handler = dynamicv1connect.NewImageServiceHandler(server.NewImageServer())
	mux.Handle(path, handler)
	path, handler = dynamicv1connect.NewUserServiceHandler(server.NewUserServer())
	mux.Handle(path, handler)

	return mux
}
