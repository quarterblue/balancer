// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type KVRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVRequest) Reset()         { *m = KVRequest{} }
func (m *KVRequest) String() string { return proto.CompactTextString(m) }
func (*KVRequest) ProtoMessage()    {}
func (*KVRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *KVRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVRequest.Unmarshal(m, b)
}
func (m *KVRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVRequest.Marshal(b, m, deterministic)
}
func (m *KVRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVRequest.Merge(m, src)
}
func (m *KVRequest) XXX_Size() int {
	return xxx_messageInfo_KVRequest.Size(m)
}
func (m *KVRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KVRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KVRequest proto.InternalMessageInfo

func (m *KVRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KVRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KVResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ping                 bool     `protobuf:"varint,3,opt,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVResponse) Reset()         { *m = KVResponse{} }
func (m *KVResponse) String() string { return proto.CompactTextString(m) }
func (*KVResponse) ProtoMessage()    {}
func (*KVResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *KVResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVResponse.Unmarshal(m, b)
}
func (m *KVResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVResponse.Marshal(b, m, deterministic)
}
func (m *KVResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVResponse.Merge(m, src)
}
func (m *KVResponse) XXX_Size() int {
	return xxx_messageInfo_KVResponse.Size(m)
}
func (m *KVResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KVResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KVResponse proto.InternalMessageInfo

func (m *KVResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *KVResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *KVResponse) GetPing() bool {
	if m != nil {
		return m.Ping
	}
	return false
}

type NodeRequest struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRequest) Reset()         { *m = NodeRequest{} }
func (m *NodeRequest) String() string { return proto.CompactTextString(m) }
func (*NodeRequest) ProtoMessage()    {}
func (*NodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *NodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRequest.Unmarshal(m, b)
}
func (m *NodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRequest.Marshal(b, m, deterministic)
}
func (m *NodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRequest.Merge(m, src)
}
func (m *NodeRequest) XXX_Size() int {
	return xxx_messageInfo_NodeRequest.Size(m)
}
func (m *NodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRequest proto.InternalMessageInfo

func (m *NodeRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type NodeResponse struct {
	Notified             bool     `protobuf:"varint,1,opt,name=notified,proto3" json:"notified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeResponse) Reset()         { *m = NodeResponse{} }
func (m *NodeResponse) String() string { return proto.CompactTextString(m) }
func (*NodeResponse) ProtoMessage()    {}
func (*NodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *NodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeResponse.Unmarshal(m, b)
}
func (m *NodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeResponse.Marshal(b, m, deterministic)
}
func (m *NodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeResponse.Merge(m, src)
}
func (m *NodeResponse) XXX_Size() int {
	return xxx_messageInfo_NodeResponse.Size(m)
}
func (m *NodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeResponse proto.InternalMessageInfo

func (m *NodeResponse) GetNotified() bool {
	if m != nil {
		return m.Notified
	}
	return false
}

type Node struct {
	Ipaddr               string   `protobuf:"bytes,1,opt,name=ipaddr,proto3" json:"ipaddr,omitempty"`
	Port                 string   `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetIpaddr() string {
	if m != nil {
		return m.Ipaddr
	}
	return ""
}

func (m *Node) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type SuccessorResponse struct {
	Successorlist        []*Node  `protobuf:"bytes,1,rep,name=successorlist,proto3" json:"successorlist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuccessorResponse) Reset()         { *m = SuccessorResponse{} }
func (m *SuccessorResponse) String() string { return proto.CompactTextString(m) }
func (*SuccessorResponse) ProtoMessage()    {}
func (*SuccessorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *SuccessorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessorResponse.Unmarshal(m, b)
}
func (m *SuccessorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessorResponse.Marshal(b, m, deterministic)
}
func (m *SuccessorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessorResponse.Merge(m, src)
}
func (m *SuccessorResponse) XXX_Size() int {
	return xxx_messageInfo_SuccessorResponse.Size(m)
}
func (m *SuccessorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessorResponse proto.InternalMessageInfo

func (m *SuccessorResponse) GetSuccessorlist() []*Node {
	if m != nil {
		return m.Successorlist
	}
	return nil
}

func init() {
	proto.RegisterType((*KVRequest)(nil), "proto.KVRequest")
	proto.RegisterType((*KVResponse)(nil), "proto.KVResponse")
	proto.RegisterType((*NodeRequest)(nil), "proto.NodeRequest")
	proto.RegisterType((*NodeResponse)(nil), "proto.NodeResponse")
	proto.RegisterType((*Node)(nil), "proto.Node")
	proto.RegisterType((*SuccessorResponse)(nil), "proto.SuccessorResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0x89, 0xed, 0xca, 0x76, 0x6a, 0x61, 0x0b, 0x22, 0x65, 0x57, 0xb3, 0x57, 0x73, 0xe2,
	0xd0, 0xee, 0xda, 0x0b, 0x41, 0xa6, 0xa0, 0x8c, 0xd1, 0x81, 0xf7, 0x73, 0x39, 0x8e, 0xe0, 0x6c,
	0x6a, 0x93, 0x0e, 0xf6, 0x76, 0xbe, 0x86, 0x6f, 0x23, 0x49, 0xd3, 0x5a, 0x51, 0xd8, 0xae, 0xfa,
	0xe5, 0xf4, 0x77, 0xfe, 0x7d, 0x07, 0x02, 0x89, 0xf9, 0x96, 0xaf, 0x70, 0x9c, 0xe5, 0x42, 0x09,
	0xda, 0x32, 0x9f, 0x68, 0x02, 0x9d, 0xc7, 0xe7, 0x04, 0x3f, 0x0a, 0x94, 0x8a, 0x76, 0xc1, 0x79,
	0xc3, 0x5d, 0x48, 0x06, 0x64, 0xd8, 0x49, 0xb4, 0xa4, 0x27, 0xd0, 0xda, 0x2e, 0x37, 0x05, 0x86,
	0x47, 0x26, 0x56, 0x3e, 0xa2, 0x07, 0x00, 0x9d, 0x24, 0x33, 0x91, 0x4a, 0xfc, 0x61, 0x48, 0x83,
	0xd1, 0xb5, 0xde, 0xe5, 0xda, 0xe6, 0x69, 0x49, 0x29, 0xb8, 0x19, 0x4f, 0xd7, 0xa1, 0x33, 0x20,
	0xc3, 0x76, 0x62, 0x74, 0x74, 0x06, 0xfe, 0x4c, 0x30, 0xac, 0x06, 0xa0, 0xe0, 0xa6, 0x82, 0x55,
	0x95, 0x8c, 0x8e, 0x46, 0x70, 0x5c, 0x22, 0xb6, 0x5d, 0x1f, 0xda, 0xa9, 0x50, 0xfc, 0x95, 0x23,
	0x33, 0x5c, 0x3b, 0xa9, 0xdf, 0x51, 0x0c, 0xae, 0x66, 0xe9, 0x29, 0x78, 0x3c, 0x5b, 0x32, 0x96,
	0xdb, 0x4a, 0xf6, 0x65, 0x46, 0x10, 0xb9, 0xb2, 0x53, 0x19, 0x1d, 0x4d, 0xa1, 0xb7, 0x28, 0x56,
	0x2b, 0x94, 0x52, 0xe4, 0x75, 0x93, 0x6b, 0x08, 0x64, 0x15, 0xdc, 0x70, 0xa9, 0x42, 0x32, 0x70,
	0x86, 0x7e, 0xec, 0x97, 0xe6, 0x8d, 0xcd, 0x40, 0xbf, 0x89, 0xf8, 0x93, 0x00, 0xdc, 0x32, 0xb6,
	0x28, 0x5d, 0xa6, 0x17, 0xe0, 0xce, 0x79, 0xba, 0xa6, 0x5d, 0x9b, 0x52, 0xbb, 0xdc, 0xef, 0x35,
	0x22, 0xb6, 0xdd, 0x08, 0x9c, 0x7b, 0x54, 0x07, 0xb3, 0xf3, 0xe2, 0x40, 0xf6, 0x12, 0xbc, 0x3b,
	0xdc, 0xa0, 0xc2, 0x83, 0xf0, 0xf8, 0x8b, 0x94, 0xe7, 0xa8, 0x76, 0x18, 0x83, 0x3f, 0xcf, 0x91,
	0x61, 0xb9, 0x25, 0xa5, 0xcd, 0xed, 0x6d, 0x95, 0xa6, 0x23, 0xf4, 0x06, 0x82, 0xda, 0xca, 0x27,
	0xae, 0xef, 0xf9, 0x4f, 0x46, 0x68, 0x63, 0x7f, 0x4d, 0xbf, 0x82, 0x60, 0xca, 0x53, 0x56, 0xff,
	0xd8, 0xdf, 0xf0, 0x1c, 0xbc, 0x99, 0xbe, 0xfd, 0x6e, 0x2f, 0xfa, 0xe2, 0x19, 0x3d, 0xf9, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xb4, 0x45, 0x09, 0x40, 0x07, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddServiceClient interface {
	Ping(ctx context.Context, in *KVRequest, opts ...grpc.CallOption) (*KVResponse, error)
	Get(ctx context.Context, in *KVRequest, opts ...grpc.CallOption) (*KVResponse, error)
	Put(ctx context.Context, in *KVRequest, opts ...grpc.CallOption) (*KVResponse, error)
	Delete(ctx context.Context, in *KVRequest, opts ...grpc.CallOption) (*KVResponse, error)
}

type addServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddServiceClient(cc *grpc.ClientConn) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) Ping(ctx context.Context, in *KVRequest, opts ...grpc.CallOption) (*KVResponse, error) {
	out := new(KVResponse)
	err := c.cc.Invoke(ctx, "/proto.AddService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) Get(ctx context.Context, in *KVRequest, opts ...grpc.CallOption) (*KVResponse, error) {
	out := new(KVResponse)
	err := c.cc.Invoke(ctx, "/proto.AddService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) Put(ctx context.Context, in *KVRequest, opts ...grpc.CallOption) (*KVResponse, error) {
	out := new(KVResponse)
	err := c.cc.Invoke(ctx, "/proto.AddService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) Delete(ctx context.Context, in *KVRequest, opts ...grpc.CallOption) (*KVResponse, error) {
	out := new(KVResponse)
	err := c.cc.Invoke(ctx, "/proto.AddService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceServer is the server API for AddService service.
type AddServiceServer interface {
	Ping(context.Context, *KVRequest) (*KVResponse, error)
	Get(context.Context, *KVRequest) (*KVResponse, error)
	Put(context.Context, *KVRequest) (*KVResponse, error)
	Delete(context.Context, *KVRequest) (*KVResponse, error)
}

// UnimplementedAddServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddServiceServer struct {
}

func (*UnimplementedAddServiceServer) Ping(ctx context.Context, req *KVRequest) (*KVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedAddServiceServer) Get(ctx context.Context, req *KVRequest) (*KVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAddServiceServer) Put(ctx context.Context, req *KVRequest) (*KVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedAddServiceServer) Delete(ctx context.Context, req *KVRequest) (*KVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Ping(ctx, req.(*KVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Get(ctx, req.(*KVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Put(ctx, req.(*KVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Delete(ctx, req.(*KVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AddService_Ping_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AddService_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _AddService_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AddService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeServiceClient interface {
	Predecessor(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*Node, error)
	SuccessorList(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*SuccessorResponse, error)
	FindSuccessor(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*Node, error)
	Notify(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*Node, error)
}

type nodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewNodeServiceClient(cc *grpc.ClientConn) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) Predecessor(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/proto.NodeService/Predecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) SuccessorList(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*SuccessorResponse, error) {
	out := new(SuccessorResponse)
	err := c.cc.Invoke(ctx, "/proto.NodeService/SuccessorList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) FindSuccessor(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/proto.NodeService/FindSuccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Notify(ctx context.Context, in *NodeRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/proto.NodeService/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
type NodeServiceServer interface {
	Predecessor(context.Context, *NodeRequest) (*Node, error)
	SuccessorList(context.Context, *NodeRequest) (*SuccessorResponse, error)
	FindSuccessor(context.Context, *NodeRequest) (*Node, error)
	Notify(context.Context, *NodeRequest) (*Node, error)
}

// UnimplementedNodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (*UnimplementedNodeServiceServer) Predecessor(ctx context.Context, req *NodeRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predecessor not implemented")
}
func (*UnimplementedNodeServiceServer) SuccessorList(ctx context.Context, req *NodeRequest) (*SuccessorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuccessorList not implemented")
}
func (*UnimplementedNodeServiceServer) FindSuccessor(ctx context.Context, req *NodeRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessor not implemented")
}
func (*UnimplementedNodeServiceServer) Notify(ctx context.Context, req *NodeRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_Predecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Predecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/Predecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Predecessor(ctx, req.(*NodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_SuccessorList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).SuccessorList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/SuccessorList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).SuccessorList(ctx, req.(*NodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_FindSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).FindSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/FindSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).FindSuccessor(ctx, req.(*NodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Notify(ctx, req.(*NodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predecessor",
			Handler:    _NodeService_Predecessor_Handler,
		},
		{
			MethodName: "SuccessorList",
			Handler:    _NodeService_SuccessorList_Handler,
		},
		{
			MethodName: "FindSuccessor",
			Handler:    _NodeService_FindSuccessor_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _NodeService_Notify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
