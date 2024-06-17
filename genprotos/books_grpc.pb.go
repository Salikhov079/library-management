// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: books.proto

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
	BookService_CreateBook_FullMethodName  = "/protos.BookService/CreateBook"
	BookService_UpdateBook_FullMethodName  = "/protos.BookService/UpdateBook"
	BookService_DeleteBook_FullMethodName  = "/protos.BookService/DeleteBook"
	BookService_GetByIdBook_FullMethodName = "/protos.BookService/GetByIdBook"
	BookService_GetAllBooks_FullMethodName = "/protos.BookService/GetAllBooks"
)

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	CreateBook(ctx context.Context, in *BookReq, opts ...grpc.CallOption) (*Void, error)
	UpdateBook(ctx context.Context, in *BookRes, opts ...grpc.CallOption) (*Void, error)
	DeleteBook(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	GetByIdBook(ctx context.Context, in *ById, opts ...grpc.CallOption) (*BookRes, error)
	GetAllBooks(ctx context.Context, in *FilterBook, opts ...grpc.CallOption) (*AllBooks, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) CreateBook(ctx context.Context, in *BookReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, BookService_CreateBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdateBook(ctx context.Context, in *BookRes, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, BookService_UpdateBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) DeleteBook(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, BookService_DeleteBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetByIdBook(ctx context.Context, in *ById, opts ...grpc.CallOption) (*BookRes, error) {
	out := new(BookRes)
	err := c.cc.Invoke(ctx, BookService_GetByIdBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetAllBooks(ctx context.Context, in *FilterBook, opts ...grpc.CallOption) (*AllBooks, error) {
	out := new(AllBooks)
	err := c.cc.Invoke(ctx, BookService_GetAllBooks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility
type BookServiceServer interface {
	CreateBook(context.Context, *BookReq) (*Void, error)
	UpdateBook(context.Context, *BookRes) (*Void, error)
	DeleteBook(context.Context, *ById) (*Void, error)
	GetByIdBook(context.Context, *ById) (*BookRes, error)
	GetAllBooks(context.Context, *FilterBook) (*AllBooks, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (UnimplementedBookServiceServer) CreateBook(context.Context, *BookReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookServiceServer) UpdateBook(context.Context, *BookRes) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookServiceServer) DeleteBook(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookServiceServer) GetByIdBook(context.Context, *ById) (*BookRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdBook not implemented")
}
func (UnimplementedBookServiceServer) GetAllBooks(context.Context, *FilterBook) (*AllBooks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBooks not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	s.RegisterService(&BookService_ServiceDesc, srv)
}

func _BookService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_CreateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CreateBook(ctx, req.(*BookReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_UpdateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdateBook(ctx, req.(*BookRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_DeleteBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).DeleteBook(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetByIdBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetByIdBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_GetByIdBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetByIdBook(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetAllBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterBook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetAllBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookService_GetAllBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetAllBooks(ctx, req.(*FilterBook))
	}
	return interceptor(ctx, in, info, handler)
}

// BookService_ServiceDesc is the grpc.ServiceDesc for BookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _BookService_CreateBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookService_DeleteBook_Handler,
		},
		{
			MethodName: "GetByIdBook",
			Handler:    _BookService_GetByIdBook_Handler,
		},
		{
			MethodName: "GetAllBooks",
			Handler:    _BookService_GetAllBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books.proto",
}
