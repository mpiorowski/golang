// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: grpc.proto

package homeit

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

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServiceClient interface {
	CreateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*Token, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	CreateUser(ctx context.Context, opts ...grpc.CallOption) (UsersService_CreateUserClient, error)
	GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UsersService_GetUsersClient, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) CreateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/homeit.UsersService/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/homeit.UsersService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) CreateUser(ctx context.Context, opts ...grpc.CallOption) (UsersService_CreateUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &UsersService_ServiceDesc.Streams[0], "/homeit.UsersService/CreateUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &usersServiceCreateUserClient{stream}
	return x, nil
}

type UsersService_CreateUserClient interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ClientStream
}

type usersServiceCreateUserClient struct {
	grpc.ClientStream
}

func (x *usersServiceCreateUserClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *usersServiceCreateUserClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *usersServiceClient) GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UsersService_GetUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &UsersService_ServiceDesc.Streams[1], "/homeit.UsersService/GetUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &usersServiceGetUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UsersService_GetUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type usersServiceGetUsersClient struct {
	grpc.ClientStream
}

func (x *usersServiceGetUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations must embed UnimplementedUsersServiceServer
// for forward compatibility
type UsersServiceServer interface {
	CreateToken(context.Context, *TokenRequest) (*Token, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	CreateUser(UsersService_CreateUserServer) error
	GetUsers(*Empty, UsersService_GetUsersServer) error
	mustEmbedUnimplementedUsersServiceServer()
}

// UnimplementedUsersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (UnimplementedUsersServiceServer) CreateToken(context.Context, *TokenRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedUsersServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUsersServiceServer) CreateUser(UsersService_CreateUserServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUsersServiceServer) GetUsers(*Empty, UsersService_GetUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUsersServiceServer) mustEmbedUnimplementedUsersServiceServer() {}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homeit.UsersService/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).CreateToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homeit.UsersService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_CreateUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UsersServiceServer).CreateUser(&usersServiceCreateUserServer{stream})
}

type UsersService_CreateUserServer interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ServerStream
}

type usersServiceCreateUserServer struct {
	grpc.ServerStream
}

func (x *usersServiceCreateUserServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func (x *usersServiceCreateUserServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UsersService_GetUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UsersServiceServer).GetUsers(m, &usersServiceGetUsersServer{stream})
}

type UsersService_GetUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type usersServiceGetUsersServer struct {
	grpc.ServerStream
}

func (x *usersServiceGetUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "homeit.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _UsersService_CreateToken_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UsersService_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateUser",
			Handler:       _UsersService_CreateUser_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetUsers",
			Handler:       _UsersService_GetUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc.proto",
}

// BillServiceClient is the client API for BillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillServiceClient interface {
	GetBills(ctx context.Context, in *UserId, opts ...grpc.CallOption) (BillService_GetBillsClient, error)
	CreateBill(ctx context.Context, in *Bill, opts ...grpc.CallOption) (*Bill, error)
	DeleteBill(ctx context.Context, in *BillId, opts ...grpc.CallOption) (*Bill, error)
}

type billServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBillServiceClient(cc grpc.ClientConnInterface) BillServiceClient {
	return &billServiceClient{cc}
}

func (c *billServiceClient) GetBills(ctx context.Context, in *UserId, opts ...grpc.CallOption) (BillService_GetBillsClient, error) {
	stream, err := c.cc.NewStream(ctx, &BillService_ServiceDesc.Streams[0], "/homeit.BillService/GetBills", opts...)
	if err != nil {
		return nil, err
	}
	x := &billServiceGetBillsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BillService_GetBillsClient interface {
	Recv() (*Bill, error)
	grpc.ClientStream
}

type billServiceGetBillsClient struct {
	grpc.ClientStream
}

func (x *billServiceGetBillsClient) Recv() (*Bill, error) {
	m := new(Bill)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *billServiceClient) CreateBill(ctx context.Context, in *Bill, opts ...grpc.CallOption) (*Bill, error) {
	out := new(Bill)
	err := c.cc.Invoke(ctx, "/homeit.BillService/CreateBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billServiceClient) DeleteBill(ctx context.Context, in *BillId, opts ...grpc.CallOption) (*Bill, error) {
	out := new(Bill)
	err := c.cc.Invoke(ctx, "/homeit.BillService/DeleteBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillServiceServer is the server API for BillService service.
// All implementations must embed UnimplementedBillServiceServer
// for forward compatibility
type BillServiceServer interface {
	GetBills(*UserId, BillService_GetBillsServer) error
	CreateBill(context.Context, *Bill) (*Bill, error)
	DeleteBill(context.Context, *BillId) (*Bill, error)
	mustEmbedUnimplementedBillServiceServer()
}

// UnimplementedBillServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBillServiceServer struct {
}

func (UnimplementedBillServiceServer) GetBills(*UserId, BillService_GetBillsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBills not implemented")
}
func (UnimplementedBillServiceServer) CreateBill(context.Context, *Bill) (*Bill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBill not implemented")
}
func (UnimplementedBillServiceServer) DeleteBill(context.Context, *BillId) (*Bill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBill not implemented")
}
func (UnimplementedBillServiceServer) mustEmbedUnimplementedBillServiceServer() {}

// UnsafeBillServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillServiceServer will
// result in compilation errors.
type UnsafeBillServiceServer interface {
	mustEmbedUnimplementedBillServiceServer()
}

func RegisterBillServiceServer(s grpc.ServiceRegistrar, srv BillServiceServer) {
	s.RegisterService(&BillService_ServiceDesc, srv)
}

func _BillService_GetBills_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BillServiceServer).GetBills(m, &billServiceGetBillsServer{stream})
}

type BillService_GetBillsServer interface {
	Send(*Bill) error
	grpc.ServerStream
}

type billServiceGetBillsServer struct {
	grpc.ServerStream
}

func (x *billServiceGetBillsServer) Send(m *Bill) error {
	return x.ServerStream.SendMsg(m)
}

func _BillService_CreateBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bill)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillServiceServer).CreateBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homeit.BillService/CreateBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillServiceServer).CreateBill(ctx, req.(*Bill))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillService_DeleteBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BillId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillServiceServer).DeleteBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homeit.BillService/DeleteBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillServiceServer).DeleteBill(ctx, req.(*BillId))
	}
	return interceptor(ctx, in, info, handler)
}

// BillService_ServiceDesc is the grpc.ServiceDesc for BillService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BillService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "homeit.BillService",
	HandlerType: (*BillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBill",
			Handler:    _BillService_CreateBill_Handler,
		},
		{
			MethodName: "DeleteBill",
			Handler:    _BillService_DeleteBill_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBills",
			Handler:       _BillService_GetBills_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc.proto",
}

// FoodServiceClient is the client API for FoodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FoodServiceClient interface {
	GetSupplies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Supplies, error)
	CreateSupply(ctx context.Context, in *Supply, opts ...grpc.CallOption) (*Supply, error)
	DeleteSupply(ctx context.Context, in *SupplyId, opts ...grpc.CallOption) (*Supply, error)
}

type foodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFoodServiceClient(cc grpc.ClientConnInterface) FoodServiceClient {
	return &foodServiceClient{cc}
}

func (c *foodServiceClient) GetSupplies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Supplies, error) {
	out := new(Supplies)
	err := c.cc.Invoke(ctx, "/homeit.FoodService/GetSupplies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodServiceClient) CreateSupply(ctx context.Context, in *Supply, opts ...grpc.CallOption) (*Supply, error) {
	out := new(Supply)
	err := c.cc.Invoke(ctx, "/homeit.FoodService/CreateSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodServiceClient) DeleteSupply(ctx context.Context, in *SupplyId, opts ...grpc.CallOption) (*Supply, error) {
	out := new(Supply)
	err := c.cc.Invoke(ctx, "/homeit.FoodService/DeleteSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoodServiceServer is the server API for FoodService service.
// All implementations must embed UnimplementedFoodServiceServer
// for forward compatibility
type FoodServiceServer interface {
	GetSupplies(context.Context, *Empty) (*Supplies, error)
	CreateSupply(context.Context, *Supply) (*Supply, error)
	DeleteSupply(context.Context, *SupplyId) (*Supply, error)
	mustEmbedUnimplementedFoodServiceServer()
}

// UnimplementedFoodServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFoodServiceServer struct {
}

func (UnimplementedFoodServiceServer) GetSupplies(context.Context, *Empty) (*Supplies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupplies not implemented")
}
func (UnimplementedFoodServiceServer) CreateSupply(context.Context, *Supply) (*Supply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSupply not implemented")
}
func (UnimplementedFoodServiceServer) DeleteSupply(context.Context, *SupplyId) (*Supply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSupply not implemented")
}
func (UnimplementedFoodServiceServer) mustEmbedUnimplementedFoodServiceServer() {}

// UnsafeFoodServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FoodServiceServer will
// result in compilation errors.
type UnsafeFoodServiceServer interface {
	mustEmbedUnimplementedFoodServiceServer()
}

func RegisterFoodServiceServer(s grpc.ServiceRegistrar, srv FoodServiceServer) {
	s.RegisterService(&FoodService_ServiceDesc, srv)
}

func _FoodService_GetSupplies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServiceServer).GetSupplies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homeit.FoodService/GetSupplies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServiceServer).GetSupplies(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoodService_CreateSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Supply)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServiceServer).CreateSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homeit.FoodService/CreateSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServiceServer).CreateSupply(ctx, req.(*Supply))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoodService_DeleteSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServiceServer).DeleteSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homeit.FoodService/DeleteSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServiceServer).DeleteSupply(ctx, req.(*SupplyId))
	}
	return interceptor(ctx, in, info, handler)
}

// FoodService_ServiceDesc is the grpc.ServiceDesc for FoodService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FoodService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "homeit.FoodService",
	HandlerType: (*FoodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSupplies",
			Handler:    _FoodService_GetSupplies_Handler,
		},
		{
			MethodName: "CreateSupply",
			Handler:    _FoodService_CreateSupply_Handler,
		},
		{
			MethodName: "DeleteSupply",
			Handler:    _FoodService_DeleteSupply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
