// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: view.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ViewServiceClient is the client API for ViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ViewServiceClient interface {
	// create
	Create(ctx context.Context, in *View, opts ...grpc.CallOption) (*View, error)
	// update
	Update(ctx context.Context, in *View, opts ...grpc.CallOption) (*ViewBoolResponse, error)
	// delete
	Delete(ctx context.Context, in *ViewIdRequest, opts ...grpc.CallOption) (*ViewBoolResponse, error)
	// get by id
	GetById(ctx context.Context, in *ViewIdRequest, opts ...grpc.CallOption) (*View, error)
	// get by user
	GetListByUser(ctx context.Context, in *ViewUserIdRequest, opts ...grpc.CallOption) (*ViewList, error)
	// get by item
	GetListByItem(ctx context.Context, in *ViewItemIdRequest, opts ...grpc.CallOption) (*ViewList, error)
	// get list by id list
	GetListByIdList(ctx context.Context, in *ViewIdListRequest, opts ...grpc.CallOption) (*ViewList, error)
	// get all
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ViewList, error)
}

type viewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewViewServiceClient(cc grpc.ClientConnInterface) ViewServiceClient {
	return &viewServiceClient{cc}
}

func (c *viewServiceClient) Create(ctx context.Context, in *View, opts ...grpc.CallOption) (*View, error) {
	out := new(View)
	err := c.cc.Invoke(ctx, "/view.ViewService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) Update(ctx context.Context, in *View, opts ...grpc.CallOption) (*ViewBoolResponse, error) {
	out := new(ViewBoolResponse)
	err := c.cc.Invoke(ctx, "/view.ViewService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) Delete(ctx context.Context, in *ViewIdRequest, opts ...grpc.CallOption) (*ViewBoolResponse, error) {
	out := new(ViewBoolResponse)
	err := c.cc.Invoke(ctx, "/view.ViewService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) GetById(ctx context.Context, in *ViewIdRequest, opts ...grpc.CallOption) (*View, error) {
	out := new(View)
	err := c.cc.Invoke(ctx, "/view.ViewService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) GetListByUser(ctx context.Context, in *ViewUserIdRequest, opts ...grpc.CallOption) (*ViewList, error) {
	out := new(ViewList)
	err := c.cc.Invoke(ctx, "/view.ViewService/GetListByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) GetListByItem(ctx context.Context, in *ViewItemIdRequest, opts ...grpc.CallOption) (*ViewList, error) {
	out := new(ViewList)
	err := c.cc.Invoke(ctx, "/view.ViewService/GetListByItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) GetListByIdList(ctx context.Context, in *ViewIdListRequest, opts ...grpc.CallOption) (*ViewList, error) {
	out := new(ViewList)
	err := c.cc.Invoke(ctx, "/view.ViewService/GetListByIdList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viewServiceClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ViewList, error) {
	out := new(ViewList)
	err := c.cc.Invoke(ctx, "/view.ViewService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViewServiceServer is the server API for ViewService service.
// All implementations should embed UnimplementedViewServiceServer
// for forward compatibility
type ViewServiceServer interface {
	// create
	Create(context.Context, *View) (*View, error)
	// update
	Update(context.Context, *View) (*ViewBoolResponse, error)
	// delete
	Delete(context.Context, *ViewIdRequest) (*ViewBoolResponse, error)
	// get by id
	GetById(context.Context, *ViewIdRequest) (*View, error)
	// get by user
	GetListByUser(context.Context, *ViewUserIdRequest) (*ViewList, error)
	// get by item
	GetListByItem(context.Context, *ViewItemIdRequest) (*ViewList, error)
	// get list by id list
	GetListByIdList(context.Context, *ViewIdListRequest) (*ViewList, error)
	// get all
	GetAll(context.Context, *emptypb.Empty) (*ViewList, error)
}

// UnimplementedViewServiceServer should be embedded to have forward compatible implementations.
type UnimplementedViewServiceServer struct {
}

func (UnimplementedViewServiceServer) Create(context.Context, *View) (*View, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedViewServiceServer) Update(context.Context, *View) (*ViewBoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedViewServiceServer) Delete(context.Context, *ViewIdRequest) (*ViewBoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedViewServiceServer) GetById(context.Context, *ViewIdRequest) (*View, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedViewServiceServer) GetListByUser(context.Context, *ViewUserIdRequest) (*ViewList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByUser not implemented")
}
func (UnimplementedViewServiceServer) GetListByItem(context.Context, *ViewItemIdRequest) (*ViewList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByItem not implemented")
}
func (UnimplementedViewServiceServer) GetListByIdList(context.Context, *ViewIdListRequest) (*ViewList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByIdList not implemented")
}
func (UnimplementedViewServiceServer) GetAll(context.Context, *emptypb.Empty) (*ViewList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

// UnsafeViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ViewServiceServer will
// result in compilation errors.
type UnsafeViewServiceServer interface {
	mustEmbedUnimplementedViewServiceServer()
}

func RegisterViewServiceServer(s grpc.ServiceRegistrar, srv ViewServiceServer) {
	s.RegisterService(&ViewService_ServiceDesc, srv)
}

func _ViewService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(View)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/view.ViewService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).Create(ctx, req.(*View))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(View)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/view.ViewService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).Update(ctx, req.(*View))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/view.ViewService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).Delete(ctx, req.(*ViewIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/view.ViewService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).GetById(ctx, req.(*ViewIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_GetListByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).GetListByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/view.ViewService/GetListByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).GetListByUser(ctx, req.(*ViewUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_GetListByItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewItemIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).GetListByItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/view.ViewService/GetListByItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).GetListByItem(ctx, req.(*ViewItemIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_GetListByIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewIdListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).GetListByIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/view.ViewService/GetListByIdList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).GetListByIdList(ctx, req.(*ViewIdListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViewService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViewServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/view.ViewService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViewServiceServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ViewService_ServiceDesc is the grpc.ServiceDesc for ViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "view.ViewService",
	HandlerType: (*ViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ViewService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ViewService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ViewService_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ViewService_GetById_Handler,
		},
		{
			MethodName: "GetListByUser",
			Handler:    _ViewService_GetListByUser_Handler,
		},
		{
			MethodName: "GetListByItem",
			Handler:    _ViewService_GetListByItem_Handler,
		},
		{
			MethodName: "GetListByIdList",
			Handler:    _ViewService_GetListByIdList_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ViewService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "view.proto",
}
