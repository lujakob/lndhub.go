// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: lndhubrpc/lndhub.proto

package lndhubrpc

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

// InvoiceSubscriptionClient is the client API for InvoiceSubscription service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoiceSubscriptionClient interface {
	SubsribeInvoices(ctx context.Context, in *SubsribeInvoicesRequest, opts ...grpc.CallOption) (InvoiceSubscription_SubsribeInvoicesClient, error)
}

type invoiceSubscriptionClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoiceSubscriptionClient(cc grpc.ClientConnInterface) InvoiceSubscriptionClient {
	return &invoiceSubscriptionClient{cc}
}

func (c *invoiceSubscriptionClient) SubsribeInvoices(ctx context.Context, in *SubsribeInvoicesRequest, opts ...grpc.CallOption) (InvoiceSubscription_SubsribeInvoicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &InvoiceSubscription_ServiceDesc.Streams[0], "/lndhubrpc.InvoiceSubscription/SubsribeInvoices", opts...)
	if err != nil {
		return nil, err
	}
	x := &invoiceSubscriptionSubsribeInvoicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InvoiceSubscription_SubsribeInvoicesClient interface {
	Recv() (*Invoice, error)
	grpc.ClientStream
}

type invoiceSubscriptionSubsribeInvoicesClient struct {
	grpc.ClientStream
}

func (x *invoiceSubscriptionSubsribeInvoicesClient) Recv() (*Invoice, error) {
	m := new(Invoice)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InvoiceSubscriptionServer is the server API for InvoiceSubscription service.
// All implementations must embed UnimplementedInvoiceSubscriptionServer
// for forward compatibility
type InvoiceSubscriptionServer interface {
	SubsribeInvoices(*SubsribeInvoicesRequest, InvoiceSubscription_SubsribeInvoicesServer) error
	mustEmbedUnimplementedInvoiceSubscriptionServer()
}

// UnimplementedInvoiceSubscriptionServer must be embedded to have forward compatible implementations.
type UnimplementedInvoiceSubscriptionServer struct {
}

func (UnimplementedInvoiceSubscriptionServer) SubsribeInvoices(*SubsribeInvoicesRequest, InvoiceSubscription_SubsribeInvoicesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubsribeInvoices not implemented")
}
func (UnimplementedInvoiceSubscriptionServer) mustEmbedUnimplementedInvoiceSubscriptionServer() {}

// UnsafeInvoiceSubscriptionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoiceSubscriptionServer will
// result in compilation errors.
type UnsafeInvoiceSubscriptionServer interface {
	mustEmbedUnimplementedInvoiceSubscriptionServer()
}

func RegisterInvoiceSubscriptionServer(s grpc.ServiceRegistrar, srv InvoiceSubscriptionServer) {
	s.RegisterService(&InvoiceSubscription_ServiceDesc, srv)
}

func _InvoiceSubscription_SubsribeInvoices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubsribeInvoicesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InvoiceSubscriptionServer).SubsribeInvoices(m, &invoiceSubscriptionSubsribeInvoicesServer{stream})
}

type InvoiceSubscription_SubsribeInvoicesServer interface {
	Send(*Invoice) error
	grpc.ServerStream
}

type invoiceSubscriptionSubsribeInvoicesServer struct {
	grpc.ServerStream
}

func (x *invoiceSubscriptionSubsribeInvoicesServer) Send(m *Invoice) error {
	return x.ServerStream.SendMsg(m)
}

// InvoiceSubscription_ServiceDesc is the grpc.ServiceDesc for InvoiceSubscription service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InvoiceSubscription_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lndhubrpc.InvoiceSubscription",
	HandlerType: (*InvoiceSubscriptionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubsribeInvoices",
			Handler:       _InvoiceSubscription_SubsribeInvoices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lndhubrpc/lndhub.proto",
}
