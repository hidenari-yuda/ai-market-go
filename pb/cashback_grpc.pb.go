// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: cashback.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CashbackServiceClient is the client API for CashbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CashbackServiceClient interface {
	// general
	// create
	Create(ctx context.Context, in *Cashback, opts ...grpc.CallOption) (*Cashback, error)
	// update
	Update(ctx context.Context, in *Cashback, opts ...grpc.CallOption) (*CashbackBoolResponse, error)
	// delete
	Delete(ctx context.Context, in *CashbackIdRequest, opts ...grpc.CallOption) (*CashbackBoolResponse, error)
	// get by id
	GetById(ctx context.Context, in *CashbackIdRequest, opts ...grpc.CallOption) (*Cashback, error)
	// get onsale list by user
	GetSoldListByUser(ctx context.Context, in *CashbackUserIdRequest, opts ...grpc.CallOption) (*CashbackList, error)
	// get purchaed list by user
	GetPurchasedListByUser(ctx context.Context, in *CashbackUserIdRequest, opts ...grpc.CallOption) (*CashbackList, error)
	// get list by id list
	GetListByIdList(ctx context.Context, in *CashbackIdListRequest, opts ...grpc.CallOption) (*CashbackList, error)
}

type cashbackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCashbackServiceClient(cc grpc.ClientConnInterface) CashbackServiceClient {
	return &cashbackServiceClient{cc}
}

func (c *cashbackServiceClient) Create(ctx context.Context, in *Cashback, opts ...grpc.CallOption) (*Cashback, error) {
	out := new(Cashback)
	err := c.cc.Invoke(ctx, "/cashback.CashbackService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashbackServiceClient) Update(ctx context.Context, in *Cashback, opts ...grpc.CallOption) (*CashbackBoolResponse, error) {
	out := new(CashbackBoolResponse)
	err := c.cc.Invoke(ctx, "/cashback.CashbackService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashbackServiceClient) Delete(ctx context.Context, in *CashbackIdRequest, opts ...grpc.CallOption) (*CashbackBoolResponse, error) {
	out := new(CashbackBoolResponse)
	err := c.cc.Invoke(ctx, "/cashback.CashbackService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashbackServiceClient) GetById(ctx context.Context, in *CashbackIdRequest, opts ...grpc.CallOption) (*Cashback, error) {
	out := new(Cashback)
	err := c.cc.Invoke(ctx, "/cashback.CashbackService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashbackServiceClient) GetSoldListByUser(ctx context.Context, in *CashbackUserIdRequest, opts ...grpc.CallOption) (*CashbackList, error) {
	out := new(CashbackList)
	err := c.cc.Invoke(ctx, "/cashback.CashbackService/GetSoldListByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashbackServiceClient) GetPurchasedListByUser(ctx context.Context, in *CashbackUserIdRequest, opts ...grpc.CallOption) (*CashbackList, error) {
	out := new(CashbackList)
	err := c.cc.Invoke(ctx, "/cashback.CashbackService/GetPurchasedListByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashbackServiceClient) GetListByIdList(ctx context.Context, in *CashbackIdListRequest, opts ...grpc.CallOption) (*CashbackList, error) {
	out := new(CashbackList)
	err := c.cc.Invoke(ctx, "/cashback.CashbackService/GetListByIdList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CashbackServiceServer is the server API for CashbackService service.
// All implementations should embed UnimplementedCashbackServiceServer
// for forward compatibility
type CashbackServiceServer interface {
	// general
	// create
	Create(context.Context, *Cashback) (*Cashback, error)
	// update
	Update(context.Context, *Cashback) (*CashbackBoolResponse, error)
	// delete
	Delete(context.Context, *CashbackIdRequest) (*CashbackBoolResponse, error)
	// get by id
	GetById(context.Context, *CashbackIdRequest) (*Cashback, error)
	// get onsale list by user
	GetSoldListByUser(context.Context, *CashbackUserIdRequest) (*CashbackList, error)
	// get purchaed list by user
	GetPurchasedListByUser(context.Context, *CashbackUserIdRequest) (*CashbackList, error)
	// get list by id list
	GetListByIdList(context.Context, *CashbackIdListRequest) (*CashbackList, error)
}

// UnimplementedCashbackServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCashbackServiceServer struct {
}

func (UnimplementedCashbackServiceServer) Create(context.Context, *Cashback) (*Cashback, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCashbackServiceServer) Update(context.Context, *Cashback) (*CashbackBoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCashbackServiceServer) Delete(context.Context, *CashbackIdRequest) (*CashbackBoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCashbackServiceServer) GetById(context.Context, *CashbackIdRequest) (*Cashback, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedCashbackServiceServer) GetSoldListByUser(context.Context, *CashbackUserIdRequest) (*CashbackList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSoldListByUser not implemented")
}
func (UnimplementedCashbackServiceServer) GetPurchasedListByUser(context.Context, *CashbackUserIdRequest) (*CashbackList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPurchasedListByUser not implemented")
}
func (UnimplementedCashbackServiceServer) GetListByIdList(context.Context, *CashbackIdListRequest) (*CashbackList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByIdList not implemented")
}

// UnsafeCashbackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CashbackServiceServer will
// result in compilation errors.
type UnsafeCashbackServiceServer interface {
	mustEmbedUnimplementedCashbackServiceServer()
}

func RegisterCashbackServiceServer(s grpc.ServiceRegistrar, srv CashbackServiceServer) {
	s.RegisterService(&CashbackService_ServiceDesc, srv)
}

func _CashbackService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cashback)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CashbackServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cashback.CashbackService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CashbackServiceServer).Create(ctx, req.(*Cashback))
	}
	return interceptor(ctx, in, info, handler)
}

func _CashbackService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cashback)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CashbackServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cashback.CashbackService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CashbackServiceServer).Update(ctx, req.(*Cashback))
	}
	return interceptor(ctx, in, info, handler)
}

func _CashbackService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CashbackIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CashbackServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cashback.CashbackService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CashbackServiceServer).Delete(ctx, req.(*CashbackIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CashbackService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CashbackIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CashbackServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cashback.CashbackService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CashbackServiceServer).GetById(ctx, req.(*CashbackIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CashbackService_GetSoldListByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CashbackUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CashbackServiceServer).GetSoldListByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cashback.CashbackService/GetSoldListByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CashbackServiceServer).GetSoldListByUser(ctx, req.(*CashbackUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CashbackService_GetPurchasedListByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CashbackUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CashbackServiceServer).GetPurchasedListByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cashback.CashbackService/GetPurchasedListByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CashbackServiceServer).GetPurchasedListByUser(ctx, req.(*CashbackUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CashbackService_GetListByIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CashbackIdListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CashbackServiceServer).GetListByIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cashback.CashbackService/GetListByIdList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CashbackServiceServer).GetListByIdList(ctx, req.(*CashbackIdListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CashbackService_ServiceDesc is the grpc.ServiceDesc for CashbackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CashbackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cashback.CashbackService",
	HandlerType: (*CashbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CashbackService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CashbackService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CashbackService_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _CashbackService_GetById_Handler,
		},
		{
			MethodName: "GetSoldListByUser",
			Handler:    _CashbackService_GetSoldListByUser_Handler,
		},
		{
			MethodName: "GetPurchasedListByUser",
			Handler:    _CashbackService_GetPurchasedListByUser_Handler,
		},
		{
			MethodName: "GetListByIdList",
			Handler:    _CashbackService_GetListByIdList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cashback.proto",
}
