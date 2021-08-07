// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package friends

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FriendsClient is the client API for Friends service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendsClient interface {
	Add(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*Empty, error)
	Remove(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*Empty, error)
	GetFriends(ctx context.Context, in *User, opts ...grpc.CallOption) (*FriendsList, error)
}

type friendsClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendsClient(cc grpc.ClientConnInterface) FriendsClient {
	return &friendsClient{cc}
}

func (c *friendsClient) Add(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cloudstate.samples.chat.friends.Friends/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsClient) Remove(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cloudstate.samples.chat.friends.Friends/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsClient) GetFriends(ctx context.Context, in *User, opts ...grpc.CallOption) (*FriendsList, error) {
	out := new(FriendsList)
	err := c.cc.Invoke(ctx, "/cloudstate.samples.chat.friends.Friends/GetFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendsServer is the server API for Friends service.
// All implementations must embed UnimplementedFriendsServer
// for forward compatibility
type FriendsServer interface {
	Add(context.Context, *FriendRequest) (*Empty, error)
	Remove(context.Context, *FriendRequest) (*Empty, error)
	GetFriends(context.Context, *User) (*FriendsList, error)
	mustEmbedUnimplementedFriendsServer()
}

// UnimplementedFriendsServer must be embedded to have forward compatible implementations.
type UnimplementedFriendsServer struct {
}

func (UnimplementedFriendsServer) Add(context.Context, *FriendRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedFriendsServer) Remove(context.Context, *FriendRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedFriendsServer) GetFriends(context.Context, *User) (*FriendsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriends not implemented")
}
func (UnimplementedFriendsServer) mustEmbedUnimplementedFriendsServer() {}

// UnsafeFriendsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendsServer will
// result in compilation errors.
type UnsafeFriendsServer interface {
	mustEmbedUnimplementedFriendsServer()
}

func RegisterFriendsServer(s grpc.ServiceRegistrar, srv FriendsServer) {
	s.RegisterService(&_Friends_serviceDesc, srv)
}

func _Friends_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudstate.samples.chat.friends.Friends/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).Add(ctx, req.(*FriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friends_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudstate.samples.chat.friends.Friends/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).Remove(ctx, req.(*FriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friends_GetFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).GetFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudstate.samples.chat.friends.Friends/GetFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).GetFriends(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _Friends_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloudstate.samples.chat.friends.Friends",
	HandlerType: (*FriendsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Friends_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Friends_Remove_Handler,
		},
		{
			MethodName: "GetFriends",
			Handler:    _Friends_GetFriends_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "friends.proto",
}