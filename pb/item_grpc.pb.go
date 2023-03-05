// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: item.proto

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

// ItemServiceClient is the client API for ItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemServiceClient interface {
	// general
	// create
	Create(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error)
	// update
	Update(ctx context.Context, in *Item, opts ...grpc.CallOption) (*ItemBoolResponse, error)
	// delete
	Delete(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemBoolResponse, error)
	// get by id
	GetById(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*Item, error)
	// get by user
	GetListByUser(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error)
	// get by type
	GetListByCategory(ctx context.Context, in *ItemCategoryRequest, opts ...grpc.CallOption) (*ItemList, error)
	// get by latest id=user_id
	GetLatestList(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error)
	// get by trend id=user_id
	GetTrendList(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error)
	// get by recommend id=user_id
	GetRecommendedListByUser(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error)
	// get by sold id=user_id
	GetSoldListByUser(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error)
	GetPurchasedListByUser(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error)
	// get by liked id=user_id
	GetLikedListByUser(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error)
	// get list by id list
	GetListByIdList(ctx context.Context, in *ItemIdListRequest, opts ...grpc.CallOption) (*ItemList, error)
	// admin
	// get all
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ItemList, error)
}

type itemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewItemServiceClient(cc grpc.ClientConnInterface) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) Create(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/item.ItemService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) Update(ctx context.Context, in *Item, opts ...grpc.CallOption) (*ItemBoolResponse, error) {
	out := new(ItemBoolResponse)
	err := c.cc.Invoke(ctx, "/item.ItemService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) Delete(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemBoolResponse, error) {
	out := new(ItemBoolResponse)
	err := c.cc.Invoke(ctx, "/item.ItemService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetById(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/item.ItemService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetListByUser(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error) {
	out := new(ItemList)
	err := c.cc.Invoke(ctx, "/item.ItemService/GetListByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetListByCategory(ctx context.Context, in *ItemCategoryRequest, opts ...grpc.CallOption) (*ItemList, error) {
	out := new(ItemList)
	err := c.cc.Invoke(ctx, "/item.ItemService/GetListByCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetLatestList(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error) {
	out := new(ItemList)
	err := c.cc.Invoke(ctx, "/item.ItemService/GetLatestList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetTrendList(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error) {
	out := new(ItemList)
	err := c.cc.Invoke(ctx, "/item.ItemService/GetTrendList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetRecommendedListByUser(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error) {
	out := new(ItemList)
	err := c.cc.Invoke(ctx, "/item.ItemService/GetRecommendedListByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetSoldListByUser(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error) {
	out := new(ItemList)
	err := c.cc.Invoke(ctx, "/item.ItemService/GetSoldListByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetPurchasedListByUser(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error) {
	out := new(ItemList)
	err := c.cc.Invoke(ctx, "/item.ItemService/GetPurchasedListByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetLikedListByUser(ctx context.Context, in *ItemIdRequest, opts ...grpc.CallOption) (*ItemList, error) {
	out := new(ItemList)
	err := c.cc.Invoke(ctx, "/item.ItemService/GetLikedListByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetListByIdList(ctx context.Context, in *ItemIdListRequest, opts ...grpc.CallOption) (*ItemList, error) {
	out := new(ItemList)
	err := c.cc.Invoke(ctx, "/item.ItemService/GetListByIdList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ItemList, error) {
	out := new(ItemList)
	err := c.cc.Invoke(ctx, "/item.ItemService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemServiceServer is the server API for ItemService service.
// All implementations should embed UnimplementedItemServiceServer
// for forward compatibility
type ItemServiceServer interface {
	// general
	// create
	Create(context.Context, *Item) (*Item, error)
	// update
	Update(context.Context, *Item) (*ItemBoolResponse, error)
	// delete
	Delete(context.Context, *ItemIdRequest) (*ItemBoolResponse, error)
	// get by id
	GetById(context.Context, *ItemIdRequest) (*Item, error)
	// get by user
	GetListByUser(context.Context, *ItemIdRequest) (*ItemList, error)
	// get by type
	GetListByCategory(context.Context, *ItemCategoryRequest) (*ItemList, error)
	// get by latest id=user_id
	GetLatestList(context.Context, *ItemIdRequest) (*ItemList, error)
	// get by trend id=user_id
	GetTrendList(context.Context, *ItemIdRequest) (*ItemList, error)
	// get by recommend id=user_id
	GetRecommendedListByUser(context.Context, *ItemIdRequest) (*ItemList, error)
	// get by sold id=user_id
	GetSoldListByUser(context.Context, *ItemIdRequest) (*ItemList, error)
	GetPurchasedListByUser(context.Context, *ItemIdRequest) (*ItemList, error)
	// get by liked id=user_id
	GetLikedListByUser(context.Context, *ItemIdRequest) (*ItemList, error)
	// get list by id list
	GetListByIdList(context.Context, *ItemIdListRequest) (*ItemList, error)
	// admin
	// get all
	GetAll(context.Context, *emptypb.Empty) (*ItemList, error)
}

// UnimplementedItemServiceServer should be embedded to have forward compatible implementations.
type UnimplementedItemServiceServer struct {
}

func (UnimplementedItemServiceServer) Create(context.Context, *Item) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedItemServiceServer) Update(context.Context, *Item) (*ItemBoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedItemServiceServer) Delete(context.Context, *ItemIdRequest) (*ItemBoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedItemServiceServer) GetById(context.Context, *ItemIdRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedItemServiceServer) GetListByUser(context.Context, *ItemIdRequest) (*ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByUser not implemented")
}
func (UnimplementedItemServiceServer) GetListByCategory(context.Context, *ItemCategoryRequest) (*ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByCategory not implemented")
}
func (UnimplementedItemServiceServer) GetLatestList(context.Context, *ItemIdRequest) (*ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestList not implemented")
}
func (UnimplementedItemServiceServer) GetTrendList(context.Context, *ItemIdRequest) (*ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrendList not implemented")
}
func (UnimplementedItemServiceServer) GetRecommendedListByUser(context.Context, *ItemIdRequest) (*ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendedListByUser not implemented")
}
func (UnimplementedItemServiceServer) GetSoldListByUser(context.Context, *ItemIdRequest) (*ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSoldListByUser not implemented")
}
func (UnimplementedItemServiceServer) GetPurchasedListByUser(context.Context, *ItemIdRequest) (*ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPurchasedListByUser not implemented")
}
func (UnimplementedItemServiceServer) GetLikedListByUser(context.Context, *ItemIdRequest) (*ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikedListByUser not implemented")
}
func (UnimplementedItemServiceServer) GetListByIdList(context.Context, *ItemIdListRequest) (*ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListByIdList not implemented")
}
func (UnimplementedItemServiceServer) GetAll(context.Context, *emptypb.Empty) (*ItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

// UnsafeItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemServiceServer will
// result in compilation errors.
type UnsafeItemServiceServer interface {
	mustEmbedUnimplementedItemServiceServer()
}

func RegisterItemServiceServer(s grpc.ServiceRegistrar, srv ItemServiceServer) {
	s.RegisterService(&ItemService_ServiceDesc, srv)
}

func _ItemService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).Create(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).Update(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).Delete(ctx, req.(*ItemIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetById(ctx, req.(*ItemIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetListByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetListByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/GetListByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetListByUser(ctx, req.(*ItemIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetListByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetListByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/GetListByCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetListByCategory(ctx, req.(*ItemCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetLatestList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetLatestList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/GetLatestList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetLatestList(ctx, req.(*ItemIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetTrendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetTrendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/GetTrendList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetTrendList(ctx, req.(*ItemIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetRecommendedListByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetRecommendedListByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/GetRecommendedListByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetRecommendedListByUser(ctx, req.(*ItemIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetSoldListByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetSoldListByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/GetSoldListByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetSoldListByUser(ctx, req.(*ItemIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetPurchasedListByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetPurchasedListByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/GetPurchasedListByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetPurchasedListByUser(ctx, req.(*ItemIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetLikedListByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetLikedListByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/GetLikedListByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetLikedListByUser(ctx, req.(*ItemIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetListByIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemIdListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetListByIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/GetListByIdList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetListByIdList(ctx, req.(*ItemIdListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/item.ItemService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ItemService_ServiceDesc is the grpc.ServiceDesc for ItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "item.ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ItemService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ItemService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ItemService_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ItemService_GetById_Handler,
		},
		{
			MethodName: "GetListByUser",
			Handler:    _ItemService_GetListByUser_Handler,
		},
		{
			MethodName: "GetListByCategory",
			Handler:    _ItemService_GetListByCategory_Handler,
		},
		{
			MethodName: "GetLatestList",
			Handler:    _ItemService_GetLatestList_Handler,
		},
		{
			MethodName: "GetTrendList",
			Handler:    _ItemService_GetTrendList_Handler,
		},
		{
			MethodName: "GetRecommendedListByUser",
			Handler:    _ItemService_GetRecommendedListByUser_Handler,
		},
		{
			MethodName: "GetSoldListByUser",
			Handler:    _ItemService_GetSoldListByUser_Handler,
		},
		{
			MethodName: "GetPurchasedListByUser",
			Handler:    _ItemService_GetPurchasedListByUser_Handler,
		},
		{
			MethodName: "GetLikedListByUser",
			Handler:    _ItemService_GetLikedListByUser_Handler,
		},
		{
			MethodName: "GetListByIdList",
			Handler:    _ItemService_GetListByIdList_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ItemService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "item.proto",
}
