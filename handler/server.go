package handler

import (
	"net/http"

	"github.com/mamoruuji/dynamic_novel_api/gen/proto/dynamic/v1/dynamicv1connect"
	. "github.com/mamoruuji/dynamic_novel_api/server"
)

func Server() http.Handler {
	mux := http.NewServeMux()

	services := []func() (string, http.Handler){
		func() (string, http.Handler) { return dynamicv1connect.NewDynamicServiceHandler(NewDynamicServer()) },
		func() (string, http.Handler) { return dynamicv1connect.NewImageServiceHandler(NewImageServer()) },
		func() (string, http.Handler) { return dynamicv1connect.NewChapterServiceHandler(NewChapterServer()) },
		func() (string, http.Handler) { return dynamicv1connect.NewPageServiceHandler(NewPageServer()) },
		func() (string, http.Handler) { return dynamicv1connect.NewTagServiceHandler(NewTagServer()) },
		func() (string, http.Handler) { return dynamicv1connect.NewUserServiceHandler(NewUserServer()) },
		func() (string, http.Handler) { return dynamicv1connect.NewMasterServiceHandler(NewMasterServer()) },
	}

	for _, service := range services {
		path, handler := service()
		mux.Handle(path, handler)
	}

	return mux
}
