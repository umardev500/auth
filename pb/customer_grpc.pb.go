// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: pb/customer.proto

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

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServiceClient interface {
	Create(ctx context.Context, in *CustomerCreateRequest, opts ...grpc.CallOption) (*Empty, error)
	FindOne(ctx context.Context, in *CustomerFindOneRequest, opts ...grpc.CallOption) (*Customer, error)
	FindAll(ctx context.Context, in *CustomerFindAllRequest, opts ...grpc.CallOption) (*CustomerFindAllResponse, error)
	ChangeStatus(ctx context.Context, in *CustomerChangeStatusRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	UpdateDetail(ctx context.Context, in *CustomerUpdateDetailRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	Delete(ctx context.Context, in *CustomerDeleteRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	SetExp(ctx context.Context, in *CustomerSetExpRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	Login(ctx context.Context, in *CustomerLoginRequest, opts ...grpc.CallOption) (*CustomerLoginResponse, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) Create(ctx context.Context, in *CustomerCreateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/CustomerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) FindOne(ctx context.Context, in *CustomerFindOneRequest, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/CustomerService/FindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) FindAll(ctx context.Context, in *CustomerFindAllRequest, opts ...grpc.CallOption) (*CustomerFindAllResponse, error) {
	out := new(CustomerFindAllResponse)
	err := c.cc.Invoke(ctx, "/CustomerService/FindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) ChangeStatus(ctx context.Context, in *CustomerChangeStatusRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/CustomerService/ChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) UpdateDetail(ctx context.Context, in *CustomerUpdateDetailRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/CustomerService/UpdateDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) Delete(ctx context.Context, in *CustomerDeleteRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/CustomerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) SetExp(ctx context.Context, in *CustomerSetExpRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, "/CustomerService/SetExp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) Login(ctx context.Context, in *CustomerLoginRequest, opts ...grpc.CallOption) (*CustomerLoginResponse, error) {
	out := new(CustomerLoginResponse)
	err := c.cc.Invoke(ctx, "/CustomerService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations must embed UnimplementedCustomerServiceServer
// for forward compatibility
type CustomerServiceServer interface {
	Create(context.Context, *CustomerCreateRequest) (*Empty, error)
	FindOne(context.Context, *CustomerFindOneRequest) (*Customer, error)
	FindAll(context.Context, *CustomerFindAllRequest) (*CustomerFindAllResponse, error)
	ChangeStatus(context.Context, *CustomerChangeStatusRequest) (*OperationResponse, error)
	UpdateDetail(context.Context, *CustomerUpdateDetailRequest) (*OperationResponse, error)
	Delete(context.Context, *CustomerDeleteRequest) (*OperationResponse, error)
	SetExp(context.Context, *CustomerSetExpRequest) (*OperationResponse, error)
	Login(context.Context, *CustomerLoginRequest) (*CustomerLoginResponse, error)
	mustEmbedUnimplementedCustomerServiceServer()
}

// UnimplementedCustomerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (UnimplementedCustomerServiceServer) Create(context.Context, *CustomerCreateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCustomerServiceServer) FindOne(context.Context, *CustomerFindOneRequest) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (UnimplementedCustomerServiceServer) FindAll(context.Context, *CustomerFindAllRequest) (*CustomerFindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedCustomerServiceServer) ChangeStatus(context.Context, *CustomerChangeStatusRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (UnimplementedCustomerServiceServer) UpdateDetail(context.Context, *CustomerUpdateDetailRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDetail not implemented")
}
func (UnimplementedCustomerServiceServer) Delete(context.Context, *CustomerDeleteRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCustomerServiceServer) SetExp(context.Context, *CustomerSetExpRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetExp not implemented")
}
func (UnimplementedCustomerServiceServer) Login(context.Context, *CustomerLoginRequest) (*CustomerLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedCustomerServiceServer) mustEmbedUnimplementedCustomerServiceServer() {}

// UnsafeCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceServer will
// result in compilation errors.
type UnsafeCustomerServiceServer interface {
	mustEmbedUnimplementedCustomerServiceServer()
}

func RegisterCustomerServiceServer(s grpc.ServiceRegistrar, srv CustomerServiceServer) {
	s.RegisterService(&CustomerService_ServiceDesc, srv)
}

func _CustomerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Create(ctx, req.(*CustomerCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerFindOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerService/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).FindOne(ctx, req.(*CustomerFindOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerFindAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerService/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).FindAll(ctx, req.(*CustomerFindAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerService/ChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).ChangeStatus(ctx, req.(*CustomerChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_UpdateDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerUpdateDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).UpdateDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerService/UpdateDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).UpdateDetail(ctx, req.(*CustomerUpdateDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Delete(ctx, req.(*CustomerDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_SetExp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerSetExpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).SetExp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerService/SetExp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).SetExp(ctx, req.(*CustomerSetExpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CustomerService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).Login(ctx, req.(*CustomerLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CustomerService_Create_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _CustomerService_FindOne_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _CustomerService_FindAll_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _CustomerService_ChangeStatus_Handler,
		},
		{
			MethodName: "UpdateDetail",
			Handler:    _CustomerService_UpdateDetail_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CustomerService_Delete_Handler,
		},
		{
			MethodName: "SetExp",
			Handler:    _CustomerService_SetExp_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _CustomerService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/customer.proto",
}
