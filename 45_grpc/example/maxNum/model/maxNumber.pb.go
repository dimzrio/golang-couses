// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/maxNumber.proto

package model

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type MaxNumReq struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaxNumReq) Reset()         { *m = MaxNumReq{} }
func (m *MaxNumReq) String() string { return proto.CompactTextString(m) }
func (*MaxNumReq) ProtoMessage()    {}
func (*MaxNumReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3e575eda194cd86, []int{0}
}

func (m *MaxNumReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaxNumReq.Unmarshal(m, b)
}
func (m *MaxNumReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaxNumReq.Marshal(b, m, deterministic)
}
func (m *MaxNumReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxNumReq.Merge(m, src)
}
func (m *MaxNumReq) XXX_Size() int {
	return xxx_messageInfo_MaxNumReq.Size(m)
}
func (m *MaxNumReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxNumReq.DiscardUnknown(m)
}

var xxx_messageInfo_MaxNumReq proto.InternalMessageInfo

func (m *MaxNumReq) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type MaxNumResp struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaxNumResp) Reset()         { *m = MaxNumResp{} }
func (m *MaxNumResp) String() string { return proto.CompactTextString(m) }
func (*MaxNumResp) ProtoMessage()    {}
func (*MaxNumResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3e575eda194cd86, []int{1}
}

func (m *MaxNumResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaxNumResp.Unmarshal(m, b)
}
func (m *MaxNumResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaxNumResp.Marshal(b, m, deterministic)
}
func (m *MaxNumResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaxNumResp.Merge(m, src)
}
func (m *MaxNumResp) XXX_Size() int {
	return xxx_messageInfo_MaxNumResp.Size(m)
}
func (m *MaxNumResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MaxNumResp.DiscardUnknown(m)
}

var xxx_messageInfo_MaxNumResp proto.InternalMessageInfo

func (m *MaxNumResp) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*MaxNumReq)(nil), "model.maxNumReq")
	proto.RegisterType((*MaxNumResp)(nil), "model.maxNumResp")
}

func init() { proto.RegisterFile("model/maxNumber.proto", fileDescriptor_e3e575eda194cd86) }

var fileDescriptor_e3e575eda194cd86 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0xcf, 0x4d, 0xac, 0xf0, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x05, 0x0b, 0x2b, 0x29, 0x73, 0x71, 0x42, 0x64, 0x82, 0x52, 0x0b, 0x85, 0xc4,
	0xb8, 0xd8, 0xf2, 0xc0, 0x6a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c, 0x25, 0x15,
	0x2e, 0x2e, 0x98, 0xa2, 0xe2, 0x02, 0x90, 0xaa, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0x98, 0x2a,
	0x08, 0xcf, 0xc8, 0x01, 0x66, 0x54, 0x52, 0x6a, 0x91, 0x90, 0x31, 0x17, 0x5b, 0x7a, 0x6a, 0x89,
	0x6f, 0x62, 0x85, 0x90, 0x80, 0x1e, 0xd8, 0x26, 0x3d, 0xb8, 0x35, 0x52, 0x82, 0x68, 0x22, 0xc5,
	0x05, 0x4a, 0x0c, 0x1a, 0x8c, 0x06, 0x8c, 0x4e, 0xec, 0x51, 0x10, 0x57, 0x25, 0xb1, 0x81, 0xdd,
	0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xe1, 0x92, 0xc6, 0xbc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MaxNumberClient is the client API for MaxNumber service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MaxNumberClient interface {
	GetMax(ctx context.Context, opts ...grpc.CallOption) (MaxNumber_GetMaxClient, error)
}

type maxNumberClient struct {
	cc *grpc.ClientConn
}

func NewMaxNumberClient(cc *grpc.ClientConn) MaxNumberClient {
	return &maxNumberClient{cc}
}

func (c *maxNumberClient) GetMax(ctx context.Context, opts ...grpc.CallOption) (MaxNumber_GetMaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MaxNumber_serviceDesc.Streams[0], "/model.maxNumber/getMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &maxNumberGetMaxClient{stream}
	return x, nil
}

type MaxNumber_GetMaxClient interface {
	Send(*MaxNumReq) error
	Recv() (*MaxNumResp, error)
	grpc.ClientStream
}

type maxNumberGetMaxClient struct {
	grpc.ClientStream
}

func (x *maxNumberGetMaxClient) Send(m *MaxNumReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *maxNumberGetMaxClient) Recv() (*MaxNumResp, error) {
	m := new(MaxNumResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MaxNumberServer is the server API for MaxNumber service.
type MaxNumberServer interface {
	GetMax(MaxNumber_GetMaxServer) error
}

func RegisterMaxNumberServer(s *grpc.Server, srv MaxNumberServer) {
	s.RegisterService(&_MaxNumber_serviceDesc, srv)
}

func _MaxNumber_GetMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MaxNumberServer).GetMax(&maxNumberGetMaxServer{stream})
}

type MaxNumber_GetMaxServer interface {
	Send(*MaxNumResp) error
	Recv() (*MaxNumReq, error)
	grpc.ServerStream
}

type maxNumberGetMaxServer struct {
	grpc.ServerStream
}

func (x *maxNumberGetMaxServer) Send(m *MaxNumResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *maxNumberGetMaxServer) Recv() (*MaxNumReq, error) {
	m := new(MaxNumReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MaxNumber_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.maxNumber",
	HandlerType: (*MaxNumberServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getMax",
			Handler:       _MaxNumber_GetMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "model/maxNumber.proto",
}
