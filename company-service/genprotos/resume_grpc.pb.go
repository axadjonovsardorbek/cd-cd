// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: resume.proto

package genprotos

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

const (
	ResumeService_Create_FullMethodName  = "/staff.ResumeService/Create"
	ResumeService_GetById_FullMethodName = "/staff.ResumeService/GetById"
	ResumeService_GetAll_FullMethodName  = "/staff.ResumeService/GetAll"
	ResumeService_Update_FullMethodName  = "/staff.ResumeService/Update"
	ResumeService_Delete_FullMethodName  = "/staff.ResumeService/Delete"
)

// ResumeServiceClient is the client API for ResumeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResumeServiceClient interface {
	Create(ctx context.Context, in *ResumeCreateReq, opts ...grpc.CallOption) (*ResumeRes, error)
	GetById(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*ResumeGetByIdRes, error)
	GetAll(ctx context.Context, in *ResumeGetAllReq, opts ...grpc.CallOption) (*ResumeGetAllRes, error)
	Update(ctx context.Context, in *ResumeUpdateReq, opts ...grpc.CallOption) (*ResumeRes, error)
	Delete(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*Void, error)
}

type resumeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResumeServiceClient(cc grpc.ClientConnInterface) ResumeServiceClient {
	return &resumeServiceClient{cc}
}

func (c *resumeServiceClient) Create(ctx context.Context, in *ResumeCreateReq, opts ...grpc.CallOption) (*ResumeRes, error) {
	out := new(ResumeRes)
	err := c.cc.Invoke(ctx, ResumeService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeServiceClient) GetById(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*ResumeGetByIdRes, error) {
	out := new(ResumeGetByIdRes)
	err := c.cc.Invoke(ctx, ResumeService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeServiceClient) GetAll(ctx context.Context, in *ResumeGetAllReq, opts ...grpc.CallOption) (*ResumeGetAllRes, error) {
	out := new(ResumeGetAllRes)
	err := c.cc.Invoke(ctx, ResumeService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeServiceClient) Update(ctx context.Context, in *ResumeUpdateReq, opts ...grpc.CallOption) (*ResumeRes, error) {
	out := new(ResumeRes)
	err := c.cc.Invoke(ctx, ResumeService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeServiceClient) Delete(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, ResumeService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResumeServiceServer is the server API for ResumeService service.
// All implementations must embed UnimplementedResumeServiceServer
// for forward compatibility
type ResumeServiceServer interface {
	Create(context.Context, *ResumeCreateReq) (*ResumeRes, error)
	GetById(context.Context, *Byid) (*ResumeGetByIdRes, error)
	GetAll(context.Context, *ResumeGetAllReq) (*ResumeGetAllRes, error)
	Update(context.Context, *ResumeUpdateReq) (*ResumeRes, error)
	Delete(context.Context, *Byid) (*Void, error)
	mustEmbedUnimplementedResumeServiceServer()
}

// UnimplementedResumeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResumeServiceServer struct {
}

func (UnimplementedResumeServiceServer) Create(context.Context, *ResumeCreateReq) (*ResumeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedResumeServiceServer) GetById(context.Context, *Byid) (*ResumeGetByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedResumeServiceServer) GetAll(context.Context, *ResumeGetAllReq) (*ResumeGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedResumeServiceServer) Update(context.Context, *ResumeUpdateReq) (*ResumeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedResumeServiceServer) Delete(context.Context, *Byid) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedResumeServiceServer) mustEmbedUnimplementedResumeServiceServer() {}

// UnsafeResumeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResumeServiceServer will
// result in compilation errors.
type UnsafeResumeServiceServer interface {
	mustEmbedUnimplementedResumeServiceServer()
}

func RegisterResumeServiceServer(s grpc.ServiceRegistrar, srv ResumeServiceServer) {
	s.RegisterService(&ResumeService_ServiceDesc, srv)
}

func _ResumeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).Create(ctx, req.(*ResumeCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Byid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).GetById(ctx, req.(*Byid))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeGetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).GetAll(ctx, req.(*ResumeGetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).Update(ctx, req.(*ResumeUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Byid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).Delete(ctx, req.(*Byid))
	}
	return interceptor(ctx, in, info, handler)
}

// ResumeService_ServiceDesc is the grpc.ServiceDesc for ResumeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResumeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "staff.ResumeService",
	HandlerType: (*ResumeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ResumeService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ResumeService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ResumeService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResumeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResumeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resume.proto",
}
