// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: review.proto

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

// ReviewServiceClient is the client API for ReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewServiceClient interface {
	// general
	// create
	Create(ctx context.Context, in *Review, opts ...grpc.CallOption) (*Review, error)
	// update
	Update(ctx context.Context, in *Review, opts ...grpc.CallOption) (*ReviewBoolResponse, error)
	// delete
	Delete(ctx context.Context, in *ReviewIdRequest, opts ...grpc.CallOption) (*ReviewBoolResponse, error)
	// get by id
	GetById(ctx context.Context, in *ReviewIdRequest, opts ...grpc.CallOption) (*Review, error)
	// get by user
	GetListByUser(ctx context.Context, in *ReviewIdRequest, opts ...grpc.CallOption) (*ReviewList, error)
	// get by item
	GetListByItem(ctx context.Context, in *ReviewIdRequest, opts ...grpc.CallOption) (*ReviewList, error)
	// get list by id list
	GetListByIdList(ctx context.Context, in *ReviewIdListRequest, opts ...grpc.CallOption) (*ReviewList, error)
	// admin
	// get all
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ReviewList, error)
}

type reviewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewServiceClient(cc grpc.ClientConnInterface) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) Create(ctx context.Context, in *Review, opts ...grpc.CallOption) (*Review, error) {
	out := new(Review)
	err := c.cc.Invoke(ctx, "/review.ReviewService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) Update(ctx context.Context, in *Review, opts ...grpc.CallOption) (*ReviewBoolResponse, error) {
	out := new(ReviewBoolResponse)
	err := c.cc.Invoke(ctx, "/review.ReviewService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) Delete(ctx context.Context, in *ReviewIdRequest, opts ...grpc.CallOption) (*ReviewBoolResponse, error) {
	out := new(ReviewBoolResponse)
	err := c.cc.Invoke(ctx, "/review.ReviewService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetById(ctx context.Context, in *ReviewIdRequest, opts ...grpc.CallOption) (*Review, error) {
	out := new(Review)
	err := c.cc.Invoke(ctx, "/review.ReviewService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetListByUser(ctx context.Context, in *ReviewIdRequest, opts ...grpc.CallOption) (*ReviewList, error) {
	out := new(ReviewList)
	err := c.cc.Invoke(ctx, "/review.ReviewService/GetListByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetListByItem(ctx context.Context, in *ReviewIdRequest, opts ...grpc.CallOption) (*ReviewList, error) {
	out := new(ReviewList)
	err := c.cc.Invoke(ctx, "/review.ReviewService/GetListByItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetListByIdList(ctx context.Context, in *ReviewIdListRequest, opts ...grpc.CallOption) (*ReviewList, error) {
	out := new(ReviewList)
	err := c.cc.Invoke(ctx, "/review.ReviewService/GetListByIdList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ReviewList, error) {
	out := new(ReviewList)
	err := c.cc.Invoke(ctx, "/review.ReviewService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServiceServer is the server API for ReviewService service.
// All implementations should embed UnimplementedReviewServiceServer
// for forward compatibility
type ReviewServiceServer interface {
	// general
	// create
	Create(context.Context, *Review) (*Review, error)
	// update
	Update(context.Context, *Review) (*ReviewBoolResponse, error)
	// delete
	Delete(context.Context, *ReviewIdRequest) (*ReviewBoolResponse, error)
	// get by id
	GetById(context.Context, *ReviewIdRequest) (*Review, error)
	// get by user
	GetListByUser(context.Context, *ReviewIdRequest) (*ReviewList, error)
	// get by item
	GetListByItem(context.Context, *ReviewIdRequest) (*ReviewList, error)
	// get list by id list
	GetListByIdList(context.Context, *ReviewIdListRequest) (*ReviewList, error)
	// admin
	// get all
	GetAll(context.Context, *emptypb.Empty) (*ReviewList, error)
}

// UnimplementedReviewServiceServer should be embedded to have forward compatible implementations.
type UnimplementedReviewServiceServer struct {
}

func (UnimplementedReviewServiceServer) Create(context.Context, *Review) (*Review, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedReviewServiceServer) Update(context.Context, *Review) (*ReviewBoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedReviewServiceServer) Delete(context.Context, *ReviewIdRequest) (*ReviewBoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedReviewServiceServer) GetById(context.Context, *ReviewIdRequest) (*Review, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedReviewServiceServer) GetListByUser(context.Context, *ReviewIdRequest) (*ReviewList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByUser not implemented")
}
func (UnimplementedReviewServiceServer) GetListByItem(context.Context, *ReviewIdRequest) (*ReviewList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByItem not implemented")
}
func (UnimplementedReviewServiceServer) GetListByIdList(context.Context, *ReviewIdListRequest) (*ReviewList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByIdList not implemented")
}
func (UnimplementedReviewServiceServer) GetAll(context.Context, *emptypb.Empty) (*ReviewList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

// UnsafeReviewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewServiceServer will
// result in compilation errors.
type UnsafeReviewServiceServer interface {
	mustEmbedUnimplementedReviewServiceServer()
}

func RegisterReviewServiceServer(s grpc.ServiceRegistrar, srv ReviewServiceServer) {
	s.RegisterService(&ReviewService_ServiceDesc, srv)
}

func _ReviewService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Review)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).Create(ctx, req.(*Review))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Review)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).Update(ctx, req.(*Review))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).Delete(ctx, req.(*ReviewIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetById(ctx, req.(*ReviewIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetListByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetListByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/GetListByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetListByUser(ctx, req.(*ReviewIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetListByItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetListByItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/GetListByItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetListByItem(ctx, req.(*ReviewIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetListByIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewIdListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetListByIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/GetListByIdList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetListByIdList(ctx, req.(*ReviewIdListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.ReviewService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ReviewService_ServiceDesc is the grpc.ServiceDesc for ReviewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReviewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "review.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ReviewService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ReviewService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReviewService_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ReviewService_GetById_Handler,
		},
		{
			MethodName: "GetListByUser",
			Handler:    _ReviewService_GetListByUser_Handler,
		},
		{
			MethodName: "GetListByItem",
			Handler:    _ReviewService_GetListByItem_Handler,
		},
		{
			MethodName: "GetListByIdList",
			Handler:    _ReviewService_GetListByIdList_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ReviewService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "review.proto",
}