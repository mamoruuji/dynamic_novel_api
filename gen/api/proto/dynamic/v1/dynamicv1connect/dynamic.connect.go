// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/proto/dynamic/v1/dynamic.proto

package dynamicv1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/amomon/dynamic_novel_api/gen/api/proto/dynamic/v1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// DynamicServiceName is the fully-qualified name of the DynamicService service.
	DynamicServiceName = "api.proto.dynamic.v1.DynamicService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DynamicServiceGetDynamicsProcedure is the fully-qualified name of the DynamicService's
	// GetDynamics RPC.
	DynamicServiceGetDynamicsProcedure = "/api.proto.dynamic.v1.DynamicService/GetDynamics"
	// DynamicServiceAddDynamicProcedure is the fully-qualified name of the DynamicService's AddDynamic
	// RPC.
	DynamicServiceAddDynamicProcedure = "/api.proto.dynamic.v1.DynamicService/AddDynamic"
	// DynamicServiceDeleteDynamicProcedure is the fully-qualified name of the DynamicService's
	// DeleteDynamic RPC.
	DynamicServiceDeleteDynamicProcedure = "/api.proto.dynamic.v1.DynamicService/DeleteDynamic"
	// DynamicServiceUpdateDynamicStatusProcedure is the fully-qualified name of the DynamicService's
	// UpdateDynamicStatus RPC.
	DynamicServiceUpdateDynamicStatusProcedure = "/api.proto.dynamic.v1.DynamicService/UpdateDynamicStatus"
)

// DynamicServiceClient is a client for the api.proto.dynamic.v1.DynamicService service.
type DynamicServiceClient interface {
	GetDynamics(context.Context, *connect_go.Request[v1.GetDynamicsRequest]) (*connect_go.Response[v1.GetDynamicsResponse], error)
	AddDynamic(context.Context, *connect_go.Request[v1.AddDynamicRequest]) (*connect_go.Response[v1.AddDynamicResponse], error)
	DeleteDynamic(context.Context, *connect_go.Request[v1.DeleteDynamicRequest]) (*connect_go.Response[v1.DeleteDynamicResponse], error)
	UpdateDynamicStatus(context.Context, *connect_go.Request[v1.UpdateDynamicStatusRequest]) (*connect_go.Response[v1.UpdateDynamicStatusResponse], error)
}

// NewDynamicServiceClient constructs a client for the api.proto.dynamic.v1.DynamicService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDynamicServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) DynamicServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &dynamicServiceClient{
		getDynamics: connect_go.NewClient[v1.GetDynamicsRequest, v1.GetDynamicsResponse](
			httpClient,
			baseURL+DynamicServiceGetDynamicsProcedure,
			opts...,
		),
		addDynamic: connect_go.NewClient[v1.AddDynamicRequest, v1.AddDynamicResponse](
			httpClient,
			baseURL+DynamicServiceAddDynamicProcedure,
			opts...,
		),
		deleteDynamic: connect_go.NewClient[v1.DeleteDynamicRequest, v1.DeleteDynamicResponse](
			httpClient,
			baseURL+DynamicServiceDeleteDynamicProcedure,
			opts...,
		),
		updateDynamicStatus: connect_go.NewClient[v1.UpdateDynamicStatusRequest, v1.UpdateDynamicStatusResponse](
			httpClient,
			baseURL+DynamicServiceUpdateDynamicStatusProcedure,
			opts...,
		),
	}
}

// dynamicServiceClient implements DynamicServiceClient.
type dynamicServiceClient struct {
	getDynamics         *connect_go.Client[v1.GetDynamicsRequest, v1.GetDynamicsResponse]
	addDynamic          *connect_go.Client[v1.AddDynamicRequest, v1.AddDynamicResponse]
	deleteDynamic       *connect_go.Client[v1.DeleteDynamicRequest, v1.DeleteDynamicResponse]
	updateDynamicStatus *connect_go.Client[v1.UpdateDynamicStatusRequest, v1.UpdateDynamicStatusResponse]
}

// GetDynamics calls api.proto.dynamic.v1.DynamicService.GetDynamics.
func (c *dynamicServiceClient) GetDynamics(ctx context.Context, req *connect_go.Request[v1.GetDynamicsRequest]) (*connect_go.Response[v1.GetDynamicsResponse], error) {
	return c.getDynamics.CallUnary(ctx, req)
}

// AddDynamic calls api.proto.dynamic.v1.DynamicService.AddDynamic.
func (c *dynamicServiceClient) AddDynamic(ctx context.Context, req *connect_go.Request[v1.AddDynamicRequest]) (*connect_go.Response[v1.AddDynamicResponse], error) {
	return c.addDynamic.CallUnary(ctx, req)
}

// DeleteDynamic calls api.proto.dynamic.v1.DynamicService.DeleteDynamic.
func (c *dynamicServiceClient) DeleteDynamic(ctx context.Context, req *connect_go.Request[v1.DeleteDynamicRequest]) (*connect_go.Response[v1.DeleteDynamicResponse], error) {
	return c.deleteDynamic.CallUnary(ctx, req)
}

// UpdateDynamicStatus calls api.proto.dynamic.v1.DynamicService.UpdateDynamicStatus.
func (c *dynamicServiceClient) UpdateDynamicStatus(ctx context.Context, req *connect_go.Request[v1.UpdateDynamicStatusRequest]) (*connect_go.Response[v1.UpdateDynamicStatusResponse], error) {
	return c.updateDynamicStatus.CallUnary(ctx, req)
}

// DynamicServiceHandler is an implementation of the api.proto.dynamic.v1.DynamicService service.
type DynamicServiceHandler interface {
	GetDynamics(context.Context, *connect_go.Request[v1.GetDynamicsRequest]) (*connect_go.Response[v1.GetDynamicsResponse], error)
	AddDynamic(context.Context, *connect_go.Request[v1.AddDynamicRequest]) (*connect_go.Response[v1.AddDynamicResponse], error)
	DeleteDynamic(context.Context, *connect_go.Request[v1.DeleteDynamicRequest]) (*connect_go.Response[v1.DeleteDynamicResponse], error)
	UpdateDynamicStatus(context.Context, *connect_go.Request[v1.UpdateDynamicStatusRequest]) (*connect_go.Response[v1.UpdateDynamicStatusResponse], error)
}

// NewDynamicServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDynamicServiceHandler(svc DynamicServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	dynamicServiceGetDynamicsHandler := connect_go.NewUnaryHandler(
		DynamicServiceGetDynamicsProcedure,
		svc.GetDynamics,
		opts...,
	)
	dynamicServiceAddDynamicHandler := connect_go.NewUnaryHandler(
		DynamicServiceAddDynamicProcedure,
		svc.AddDynamic,
		opts...,
	)
	dynamicServiceDeleteDynamicHandler := connect_go.NewUnaryHandler(
		DynamicServiceDeleteDynamicProcedure,
		svc.DeleteDynamic,
		opts...,
	)
	dynamicServiceUpdateDynamicStatusHandler := connect_go.NewUnaryHandler(
		DynamicServiceUpdateDynamicStatusProcedure,
		svc.UpdateDynamicStatus,
		opts...,
	)
	return "/api.proto.dynamic.v1.DynamicService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DynamicServiceGetDynamicsProcedure:
			dynamicServiceGetDynamicsHandler.ServeHTTP(w, r)
		case DynamicServiceAddDynamicProcedure:
			dynamicServiceAddDynamicHandler.ServeHTTP(w, r)
		case DynamicServiceDeleteDynamicProcedure:
			dynamicServiceDeleteDynamicHandler.ServeHTTP(w, r)
		case DynamicServiceUpdateDynamicStatusProcedure:
			dynamicServiceUpdateDynamicStatusHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDynamicServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDynamicServiceHandler struct{}

func (UnimplementedDynamicServiceHandler) GetDynamics(context.Context, *connect_go.Request[v1.GetDynamicsRequest]) (*connect_go.Response[v1.GetDynamicsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.proto.dynamic.v1.DynamicService.GetDynamics is not implemented"))
}

func (UnimplementedDynamicServiceHandler) AddDynamic(context.Context, *connect_go.Request[v1.AddDynamicRequest]) (*connect_go.Response[v1.AddDynamicResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.proto.dynamic.v1.DynamicService.AddDynamic is not implemented"))
}

func (UnimplementedDynamicServiceHandler) DeleteDynamic(context.Context, *connect_go.Request[v1.DeleteDynamicRequest]) (*connect_go.Response[v1.DeleteDynamicResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.proto.dynamic.v1.DynamicService.DeleteDynamic is not implemented"))
}

func (UnimplementedDynamicServiceHandler) UpdateDynamicStatus(context.Context, *connect_go.Request[v1.UpdateDynamicStatusRequest]) (*connect_go.Response[v1.UpdateDynamicStatusResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.proto.dynamic.v1.DynamicService.UpdateDynamicStatus is not implemented"))
}
