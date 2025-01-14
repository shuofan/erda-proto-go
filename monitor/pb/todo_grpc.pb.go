// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: todo.proto

package pb

import (
	context "context"
	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// TODOServiceClient is the client API for TODOService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TODOServiceClient interface {
	Do(ctx context.Context, in *TODORequst, opts ...grpc.CallOption) (*TODOResponse, error)
}

type tODOServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewTODOServiceClient(cc grpc1.ClientConnInterface) TODOServiceClient {
	return &tODOServiceClient{cc}
}

func (c *tODOServiceClient) Do(ctx context.Context, in *TODORequst, opts ...grpc.CallOption) (*TODOResponse, error) {
	out := new(TODOResponse)
	err := c.cc.Invoke(ctx, "/erda.monitor.TODOService/Do", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TODOServiceServer is the server API for TODOService service.
// All implementations should embed UnimplementedTODOServiceServer
// for forward compatibility
type TODOServiceServer interface {
	Do(context.Context, *TODORequst) (*TODOResponse, error)
}

// UnimplementedTODOServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTODOServiceServer struct {
}

func (*UnimplementedTODOServiceServer) Do(context.Context, *TODORequst) (*TODOResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Do not implemented")
}

func RegisterTODOServiceServer(s grpc1.ServiceRegistrar, srv TODOServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_TODOService_serviceDesc(srv, opts...), srv)
}

var _TODOService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.monitor.TODOService",
	HandlerType: (*TODOServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "todo.proto",
}

func _get_TODOService_serviceDesc(srv TODOServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_TODOService_Do_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Do(ctx, req.(*TODORequst))
	}
	var _TODOService_Do_info transport.ServiceInfo
	if h.Interceptor != nil {
		_TODOService_Do_info = transport.NewServiceInfo("erda.monitor.TODOService", "Do", srv)
		_TODOService_Do_Handler = h.Interceptor(_TODOService_Do_Handler)
	}

	var serviceDesc = _TODOService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(TODORequst)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(TODOServiceServer).Do(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _TODOService_Do_info)
				}
				if interceptor == nil {
					return _TODOService_Do_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.monitor.TODOService/Do",
				}
				return interceptor(ctx, in, info, _TODOService_Do_Handler)
			},
		},
	}
	return &serviceDesc
}
