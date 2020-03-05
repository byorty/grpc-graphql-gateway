// Code generated by protoc-gen-go. DO NOT EDIT.
// source: author/author.proto

package author

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/ysugimoto/grpc-graphql-gateway/graphql"
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

type AuthorType int32

const (
	AuthorType_NORMAL  AuthorType = 0
	AuthorType_SPECIAL AuthorType = 1
)

var AuthorType_name = map[int32]string{
	0: "NORMAL",
	1: "SPECIAL",
}

var AuthorType_value = map[string]int32{
	"NORMAL":  0,
	"SPECIAL": 1,
}

func (x AuthorType) String() string {
	return proto.EnumName(AuthorType_name, int32(x))
}

func (AuthorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{0}
}

type Author struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{0}
}

func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (m *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(m, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListAuthorsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAuthorsRequest) Reset()         { *m = ListAuthorsRequest{} }
func (m *ListAuthorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAuthorsRequest) ProtoMessage()    {}
func (*ListAuthorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{1}
}

func (m *ListAuthorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAuthorsRequest.Unmarshal(m, b)
}
func (m *ListAuthorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAuthorsRequest.Marshal(b, m, deterministic)
}
func (m *ListAuthorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuthorsRequest.Merge(m, src)
}
func (m *ListAuthorsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAuthorsRequest.Size(m)
}
func (m *ListAuthorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAuthorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAuthorsRequest proto.InternalMessageInfo

type ListAuthorsResponse struct {
	Authors              []*Author `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListAuthorsResponse) Reset()         { *m = ListAuthorsResponse{} }
func (m *ListAuthorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAuthorsResponse) ProtoMessage()    {}
func (*ListAuthorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{2}
}

func (m *ListAuthorsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAuthorsResponse.Unmarshal(m, b)
}
func (m *ListAuthorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAuthorsResponse.Marshal(b, m, deterministic)
}
func (m *ListAuthorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAuthorsResponse.Merge(m, src)
}
func (m *ListAuthorsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAuthorsResponse.Size(m)
}
func (m *ListAuthorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAuthorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAuthorsResponse proto.InternalMessageInfo

func (m *ListAuthorsResponse) GetAuthors() []*Author {
	if m != nil {
		return m.Authors
	}
	return nil
}

type GetAuthorRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAuthorRequest) Reset()         { *m = GetAuthorRequest{} }
func (m *GetAuthorRequest) String() string { return proto.CompactTextString(m) }
func (*GetAuthorRequest) ProtoMessage()    {}
func (*GetAuthorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41efdaee09960382, []int{3}
}

func (m *GetAuthorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAuthorRequest.Unmarshal(m, b)
}
func (m *GetAuthorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAuthorRequest.Marshal(b, m, deterministic)
}
func (m *GetAuthorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthorRequest.Merge(m, src)
}
func (m *GetAuthorRequest) XXX_Size() int {
	return xxx_messageInfo_GetAuthorRequest.Size(m)
}
func (m *GetAuthorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthorRequest proto.InternalMessageInfo

func (m *GetAuthorRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("some.app.v1.author.AuthorType", AuthorType_name, AuthorType_value)
	proto.RegisterType((*Author)(nil), "some.app.v1.author.Author")
	proto.RegisterType((*ListAuthorsRequest)(nil), "some.app.v1.author.ListAuthorsRequest")
	proto.RegisterType((*ListAuthorsResponse)(nil), "some.app.v1.author.ListAuthorsResponse")
	proto.RegisterType((*GetAuthorRequest)(nil), "some.app.v1.author.GetAuthorRequest")
}

func init() { proto.RegisterFile("author/author.proto", fileDescriptor_41efdaee09960382) }

var fileDescriptor_41efdaee09960382 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xad, 0xca, 0x86, 0x77, 0x28, 0x25, 0x53, 0x29, 0xc5, 0x87, 0x51, 0x50, 0x87, 0xb0,
	0x64, 0x4e, 0x1f, 0x86, 0x6f, 0x75, 0x0c, 0x11, 0xeb, 0x1f, 0x3a, 0x9f, 0x7c, 0x4b, 0xcb, 0xa5,
	0x2d, 0xb6, 0x4b, 0xd6, 0xa4, 0xd3, 0x7e, 0x05, 0xbf, 0x8f, 0xdf, 0x4f, 0x68, 0xaa, 0x22, 0x3e,
	0x9d, 0x90, 0xfb, 0xbb, 0xf7, 0x1c, 0x0e, 0xf4, 0x79, 0xa5, 0x53, 0x51, 0x32, 0x23, 0x54, 0x96,
	0x42, 0x0b, 0x42, 0x94, 0x28, 0x90, 0x72, 0x29, 0xe9, 0xfa, 0x9c, 0x9a, 0x89, 0x7b, 0x90, 0x94,
	0x5c, 0xa6, 0xab, 0x9c, 0xb5, 0x6a, 0x50, 0xef, 0x08, 0x3a, 0x7e, 0x03, 0x10, 0x02, 0xdb, 0x4b,
	0x5e, 0xa0, 0xb3, 0x39, 0xb0, 0x86, 0x3b, 0x61, 0xf3, 0xf6, 0xf6, 0x81, 0x04, 0x99, 0xd2, 0x86,
	0x50, 0x21, 0xae, 0x2a, 0x54, 0xda, 0xbb, 0x83, 0xfe, 0x9f, 0x5f, 0x25, 0xc5, 0x52, 0x21, 0xb9,
	0x84, 0xae, 0xf1, 0x52, 0x8e, 0x35, 0xd8, 0x1a, 0xf6, 0x26, 0x2e, 0xfd, 0x9f, 0x83, 0x9a, 0xad,
	0xf0, 0x1b, 0xf5, 0x4e, 0xc0, 0xbe, 0xc1, 0xf6, 0x56, 0x6b, 0xf0, 0x13, 0xc5, 0xfa, 0x8d, 0x72,
	0x76, 0x0c, 0x60, 0xa0, 0xe7, 0x5a, 0x22, 0x01, 0xe8, 0x3c, 0x3c, 0x86, 0xf7, 0x7e, 0x60, 0x6f,
	0x90, 0x1e, 0x74, 0x17, 0x4f, 0xf3, 0xd9, 0xad, 0x1f, 0xd8, 0xd6, 0xe4, 0x14, 0x76, 0x0d, 0xb6,
	0xc0, 0x72, 0x9d, 0xc5, 0xe8, 0x1e, 0x7e, 0x7c, 0x3a, 0x04, 0xf6, 0x72, 0x11, 0xf3, 0x3c, 0x15,
	0x4a, 0x5f, 0x4d, 0xc7, 0xd3, 0xb1, 0x6d, 0x5d, 0xcf, 0x5f, 0x66, 0x49, 0xa6, 0xd3, 0x2a, 0xa2,
	0xb1, 0x28, 0x58, 0xad, 0xaa, 0x24, 0x2b, 0x84, 0x16, 0x2c, 0x29, 0x65, 0x3c, 0x6a, 0x3b, 0x1a,
	0x25, 0x5c, 0xe3, 0x1b, 0xaf, 0x19, 0xbe, 0xf3, 0x42, 0xe6, 0xc8, 0x22, 0x1e, 0xbf, 0x56, 0x92,
	0x71, 0x29, 0xdb, 0xc2, 0xa3, 0x4e, 0x53, 0xe3, 0xc5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a,
	0x03, 0x58, 0xa4, 0x88, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorServiceClient is the client API for AuthorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorServiceClient interface {
}

type authorServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthorServiceClient(cc *grpc.ClientConn) AuthorServiceClient {
	return &authorServiceClient{cc}
}

// AuthorServiceServer is the server API for AuthorService service.
type AuthorServiceServer interface {
}

// UnimplementedAuthorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorServiceServer struct {
}

func RegisterAuthorServiceServer(s *grpc.Server, srv AuthorServiceServer) {
	s.RegisterService(&_AuthorService_serviceDesc, srv)
}

var _AuthorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "some.app.v1.author.AuthorService",
	HandlerType: (*AuthorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "author/author.proto",
}
