// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: energy.proto

package energy

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

// EnergyServiceClient is the client API for EnergyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnergyServiceClient interface {
	GetTariffs(ctx context.Context, in *UserId, opts ...grpc.CallOption) (EnergyService_GetTariffsClient, error)
	CreateTariff(ctx context.Context, in *Tariff, opts ...grpc.CallOption) (*Tariff, error)
	DeleteTariff(ctx context.Context, in *TariffId, opts ...grpc.CallOption) (*Tariff, error)
	GetMeters(ctx context.Context, in *UserId, opts ...grpc.CallOption) (EnergyService_GetMetersClient, error)
	CreateMeter(ctx context.Context, in *Meter, opts ...grpc.CallOption) (*Meter, error)
	DeleteMeter(ctx context.Context, in *MeterId, opts ...grpc.CallOption) (*Meter, error)
	GetMeasurements(ctx context.Context, in *MeterId, opts ...grpc.CallOption) (EnergyService_GetMeasurementsClient, error)
	CreateMeasurement(ctx context.Context, in *Measurement, opts ...grpc.CallOption) (*Measurement, error)
	DeleteMeasurement(ctx context.Context, in *MeasurementId, opts ...grpc.CallOption) (*Measurement, error)
}

type energyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEnergyServiceClient(cc grpc.ClientConnInterface) EnergyServiceClient {
	return &energyServiceClient{cc}
}

func (c *energyServiceClient) GetTariffs(ctx context.Context, in *UserId, opts ...grpc.CallOption) (EnergyService_GetTariffsClient, error) {
	stream, err := c.cc.NewStream(ctx, &EnergyService_ServiceDesc.Streams[0], "/energy.EnergyService/GetTariffs", opts...)
	if err != nil {
		return nil, err
	}
	x := &energyServiceGetTariffsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EnergyService_GetTariffsClient interface {
	Recv() (*Tariff, error)
	grpc.ClientStream
}

type energyServiceGetTariffsClient struct {
	grpc.ClientStream
}

func (x *energyServiceGetTariffsClient) Recv() (*Tariff, error) {
	m := new(Tariff)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *energyServiceClient) CreateTariff(ctx context.Context, in *Tariff, opts ...grpc.CallOption) (*Tariff, error) {
	out := new(Tariff)
	err := c.cc.Invoke(ctx, "/energy.EnergyService/CreateTariff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *energyServiceClient) DeleteTariff(ctx context.Context, in *TariffId, opts ...grpc.CallOption) (*Tariff, error) {
	out := new(Tariff)
	err := c.cc.Invoke(ctx, "/energy.EnergyService/DeleteTariff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *energyServiceClient) GetMeters(ctx context.Context, in *UserId, opts ...grpc.CallOption) (EnergyService_GetMetersClient, error) {
	stream, err := c.cc.NewStream(ctx, &EnergyService_ServiceDesc.Streams[1], "/energy.EnergyService/GetMeters", opts...)
	if err != nil {
		return nil, err
	}
	x := &energyServiceGetMetersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EnergyService_GetMetersClient interface {
	Recv() (*Meter, error)
	grpc.ClientStream
}

type energyServiceGetMetersClient struct {
	grpc.ClientStream
}

func (x *energyServiceGetMetersClient) Recv() (*Meter, error) {
	m := new(Meter)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *energyServiceClient) CreateMeter(ctx context.Context, in *Meter, opts ...grpc.CallOption) (*Meter, error) {
	out := new(Meter)
	err := c.cc.Invoke(ctx, "/energy.EnergyService/CreateMeter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *energyServiceClient) DeleteMeter(ctx context.Context, in *MeterId, opts ...grpc.CallOption) (*Meter, error) {
	out := new(Meter)
	err := c.cc.Invoke(ctx, "/energy.EnergyService/DeleteMeter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *energyServiceClient) GetMeasurements(ctx context.Context, in *MeterId, opts ...grpc.CallOption) (EnergyService_GetMeasurementsClient, error) {
	stream, err := c.cc.NewStream(ctx, &EnergyService_ServiceDesc.Streams[2], "/energy.EnergyService/GetMeasurements", opts...)
	if err != nil {
		return nil, err
	}
	x := &energyServiceGetMeasurementsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EnergyService_GetMeasurementsClient interface {
	Recv() (*Measurement, error)
	grpc.ClientStream
}

type energyServiceGetMeasurementsClient struct {
	grpc.ClientStream
}

func (x *energyServiceGetMeasurementsClient) Recv() (*Measurement, error) {
	m := new(Measurement)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *energyServiceClient) CreateMeasurement(ctx context.Context, in *Measurement, opts ...grpc.CallOption) (*Measurement, error) {
	out := new(Measurement)
	err := c.cc.Invoke(ctx, "/energy.EnergyService/CreateMeasurement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *energyServiceClient) DeleteMeasurement(ctx context.Context, in *MeasurementId, opts ...grpc.CallOption) (*Measurement, error) {
	out := new(Measurement)
	err := c.cc.Invoke(ctx, "/energy.EnergyService/DeleteMeasurement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnergyServiceServer is the server API for EnergyService service.
// All implementations must embed UnimplementedEnergyServiceServer
// for forward compatibility
type EnergyServiceServer interface {
	GetTariffs(*UserId, EnergyService_GetTariffsServer) error
	CreateTariff(context.Context, *Tariff) (*Tariff, error)
	DeleteTariff(context.Context, *TariffId) (*Tariff, error)
	GetMeters(*UserId, EnergyService_GetMetersServer) error
	CreateMeter(context.Context, *Meter) (*Meter, error)
	DeleteMeter(context.Context, *MeterId) (*Meter, error)
	GetMeasurements(*MeterId, EnergyService_GetMeasurementsServer) error
	CreateMeasurement(context.Context, *Measurement) (*Measurement, error)
	DeleteMeasurement(context.Context, *MeasurementId) (*Measurement, error)
	mustEmbedUnimplementedEnergyServiceServer()
}

// UnimplementedEnergyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEnergyServiceServer struct {
}

func (UnimplementedEnergyServiceServer) GetTariffs(*UserId, EnergyService_GetTariffsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTariffs not implemented")
}
func (UnimplementedEnergyServiceServer) CreateTariff(context.Context, *Tariff) (*Tariff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTariff not implemented")
}
func (UnimplementedEnergyServiceServer) DeleteTariff(context.Context, *TariffId) (*Tariff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTariff not implemented")
}
func (UnimplementedEnergyServiceServer) GetMeters(*UserId, EnergyService_GetMetersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMeters not implemented")
}
func (UnimplementedEnergyServiceServer) CreateMeter(context.Context, *Meter) (*Meter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMeter not implemented")
}
func (UnimplementedEnergyServiceServer) DeleteMeter(context.Context, *MeterId) (*Meter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMeter not implemented")
}
func (UnimplementedEnergyServiceServer) GetMeasurements(*MeterId, EnergyService_GetMeasurementsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMeasurements not implemented")
}
func (UnimplementedEnergyServiceServer) CreateMeasurement(context.Context, *Measurement) (*Measurement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMeasurement not implemented")
}
func (UnimplementedEnergyServiceServer) DeleteMeasurement(context.Context, *MeasurementId) (*Measurement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMeasurement not implemented")
}
func (UnimplementedEnergyServiceServer) mustEmbedUnimplementedEnergyServiceServer() {}

// UnsafeEnergyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnergyServiceServer will
// result in compilation errors.
type UnsafeEnergyServiceServer interface {
	mustEmbedUnimplementedEnergyServiceServer()
}

func RegisterEnergyServiceServer(s grpc.ServiceRegistrar, srv EnergyServiceServer) {
	s.RegisterService(&EnergyService_ServiceDesc, srv)
}

func _EnergyService_GetTariffs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EnergyServiceServer).GetTariffs(m, &energyServiceGetTariffsServer{stream})
}

type EnergyService_GetTariffsServer interface {
	Send(*Tariff) error
	grpc.ServerStream
}

type energyServiceGetTariffsServer struct {
	grpc.ServerStream
}

func (x *energyServiceGetTariffsServer) Send(m *Tariff) error {
	return x.ServerStream.SendMsg(m)
}

func _EnergyService_CreateTariff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tariff)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnergyServiceServer).CreateTariff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/energy.EnergyService/CreateTariff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnergyServiceServer).CreateTariff(ctx, req.(*Tariff))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnergyService_DeleteTariff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TariffId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnergyServiceServer).DeleteTariff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/energy.EnergyService/DeleteTariff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnergyServiceServer).DeleteTariff(ctx, req.(*TariffId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnergyService_GetMeters_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EnergyServiceServer).GetMeters(m, &energyServiceGetMetersServer{stream})
}

type EnergyService_GetMetersServer interface {
	Send(*Meter) error
	grpc.ServerStream
}

type energyServiceGetMetersServer struct {
	grpc.ServerStream
}

func (x *energyServiceGetMetersServer) Send(m *Meter) error {
	return x.ServerStream.SendMsg(m)
}

func _EnergyService_CreateMeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Meter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnergyServiceServer).CreateMeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/energy.EnergyService/CreateMeter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnergyServiceServer).CreateMeter(ctx, req.(*Meter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnergyService_DeleteMeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeterId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnergyServiceServer).DeleteMeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/energy.EnergyService/DeleteMeter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnergyServiceServer).DeleteMeter(ctx, req.(*MeterId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnergyService_GetMeasurements_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MeterId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EnergyServiceServer).GetMeasurements(m, &energyServiceGetMeasurementsServer{stream})
}

type EnergyService_GetMeasurementsServer interface {
	Send(*Measurement) error
	grpc.ServerStream
}

type energyServiceGetMeasurementsServer struct {
	grpc.ServerStream
}

func (x *energyServiceGetMeasurementsServer) Send(m *Measurement) error {
	return x.ServerStream.SendMsg(m)
}

func _EnergyService_CreateMeasurement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Measurement)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnergyServiceServer).CreateMeasurement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/energy.EnergyService/CreateMeasurement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnergyServiceServer).CreateMeasurement(ctx, req.(*Measurement))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnergyService_DeleteMeasurement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasurementId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnergyServiceServer).DeleteMeasurement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/energy.EnergyService/DeleteMeasurement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnergyServiceServer).DeleteMeasurement(ctx, req.(*MeasurementId))
	}
	return interceptor(ctx, in, info, handler)
}

// EnergyService_ServiceDesc is the grpc.ServiceDesc for EnergyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnergyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "energy.EnergyService",
	HandlerType: (*EnergyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTariff",
			Handler:    _EnergyService_CreateTariff_Handler,
		},
		{
			MethodName: "DeleteTariff",
			Handler:    _EnergyService_DeleteTariff_Handler,
		},
		{
			MethodName: "CreateMeter",
			Handler:    _EnergyService_CreateMeter_Handler,
		},
		{
			MethodName: "DeleteMeter",
			Handler:    _EnergyService_DeleteMeter_Handler,
		},
		{
			MethodName: "CreateMeasurement",
			Handler:    _EnergyService_CreateMeasurement_Handler,
		},
		{
			MethodName: "DeleteMeasurement",
			Handler:    _EnergyService_DeleteMeasurement_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTariffs",
			Handler:       _EnergyService_GetTariffs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMeters",
			Handler:       _EnergyService_GetMeters_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMeasurements",
			Handler:       _EnergyService_GetMeasurements_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "energy.proto",
}
