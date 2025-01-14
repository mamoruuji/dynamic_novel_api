package handler

import (
	"net/http"

	"github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1/dynamicv1connect"
	. "github.com/mamoruuji/dynamic_novel_api/server"
)

func Server() http.Handler {
	mux := http.NewServeMux()

	path, handler := dynamicv1connect.NewDynamicServiceHandler(NewDynamicServer())
	mux.Handle(path, handler)
	path, handler = dynamicv1connect.NewImageServiceHandler(NewImageServer())
	mux.Handle(path, handler)
	path, handler = dynamicv1connect.NewPageServiceHandler(NewPageServer())
	mux.Handle(path, handler)
	path, handler = dynamicv1connect.NewSortServiceHandler(NewSortServer())
	mux.Handle(path, handler)
	path, handler = dynamicv1connect.NewTagServiceHandler(NewTagServer())
	mux.Handle(path, handler)
	path, handler = dynamicv1connect.NewUserServiceHandler(NewUserServer())
	mux.Handle(path, handler)

	return mux
}
