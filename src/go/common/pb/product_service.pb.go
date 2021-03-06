// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product_service.proto

package pb

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

type GetProductRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Barcode              string   `protobuf:"bytes,2,opt,name=barcode,proto3" json:"barcode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductRequest) Reset()         { *m = GetProductRequest{} }
func (m *GetProductRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductRequest) ProtoMessage()    {}
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64a1a24e6b7d7ed5, []int{0}
}

func (m *GetProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductRequest.Unmarshal(m, b)
}
func (m *GetProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductRequest.Marshal(b, m, deterministic)
}
func (m *GetProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductRequest.Merge(m, src)
}
func (m *GetProductRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductRequest.Size(m)
}
func (m *GetProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductRequest proto.InternalMessageInfo

func (m *GetProductRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetProductRequest) GetBarcode() string {
	if m != nil {
		return m.Barcode
	}
	return ""
}

type GetProductResponse struct {
	Product              *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductResponse) Reset()         { *m = GetProductResponse{} }
func (m *GetProductResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductResponse) ProtoMessage()    {}
func (*GetProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64a1a24e6b7d7ed5, []int{1}
}

func (m *GetProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductResponse.Unmarshal(m, b)
}
func (m *GetProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductResponse.Marshal(b, m, deterministic)
}
func (m *GetProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductResponse.Merge(m, src)
}
func (m *GetProductResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductResponse.Size(m)
}
func (m *GetProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductResponse proto.InternalMessageInfo

func (m *GetProductResponse) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type GetProductsRequest struct {
	DisplayName          string      `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Pagination           *Pagination `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetProductsRequest) Reset()         { *m = GetProductsRequest{} }
func (m *GetProductsRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductsRequest) ProtoMessage()    {}
func (*GetProductsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64a1a24e6b7d7ed5, []int{2}
}

func (m *GetProductsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductsRequest.Unmarshal(m, b)
}
func (m *GetProductsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductsRequest.Marshal(b, m, deterministic)
}
func (m *GetProductsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductsRequest.Merge(m, src)
}
func (m *GetProductsRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductsRequest.Size(m)
}
func (m *GetProductsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductsRequest proto.InternalMessageInfo

func (m *GetProductsRequest) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *GetProductsRequest) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type GetProductsResponse struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetProductsResponse) Reset()         { *m = GetProductsResponse{} }
func (m *GetProductsResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductsResponse) ProtoMessage()    {}
func (*GetProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64a1a24e6b7d7ed5, []int{3}
}

func (m *GetProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductsResponse.Unmarshal(m, b)
}
func (m *GetProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductsResponse.Marshal(b, m, deterministic)
}
func (m *GetProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductsResponse.Merge(m, src)
}
func (m *GetProductsResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductsResponse.Size(m)
}
func (m *GetProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductsResponse proto.InternalMessageInfo

func (m *GetProductsResponse) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*GetProductRequest)(nil), "pb.GetProductRequest")
	proto.RegisterType((*GetProductResponse)(nil), "pb.GetProductResponse")
	proto.RegisterType((*GetProductsRequest)(nil), "pb.GetProductsRequest")
	proto.RegisterType((*GetProductsResponse)(nil), "pb.GetProductsResponse")
}

func init() {
	proto.RegisterFile("product_service.proto", fileDescriptor_64a1a24e6b7d7ed5)
}

var fileDescriptor_64a1a24e6b7d7ed5 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x49, 0x2a, 0x51, 0x98, 0x40, 0x04, 0x83, 0x0a, 0x56, 0x56, 0xc5, 0x12, 0x82, 0x55,
	0x16, 0x65, 0x89, 0x8a, 0xd8, 0xb1, 0x43, 0x28, 0x1c, 0xa0, 0x72, 0x62, 0xab, 0xb2, 0x44, 0x63,
	0x13, 0xbb, 0x48, 0x1c, 0x83, 0x1b, 0x23, 0xfc, 0xd3, 0xa6, 0xc9, 0x72, 0xbe, 0x79, 0x33, 0xef,
	0x79, 0x0c, 0x33, 0xdd, 0x29, 0xbe, 0x6d, 0xec, 0xca, 0x88, 0xee, 0x5b, 0x36, 0xa2, 0xd4, 0x9d,
	0xb2, 0x0a, 0x53, 0x5d, 0x17, 0x17, 0x9a, 0xad, 0x65, 0xcb, 0xac, 0x54, 0xad, 0xa7, 0xc5, 0x79,
	0x10, 0xfb, 0x92, 0x2e, 0xe1, 0xf2, 0x55, 0xd8, 0x77, 0xcf, 0x2a, 0xf1, 0xb5, 0x15, 0xc6, 0x62,
	0x0e, 0xa9, 0xe4, 0x24, 0x99, 0x27, 0x0f, 0xa7, 0x55, 0x2a, 0x39, 0x12, 0x98, 0xd6, 0xac, 0x6b,
	0x14, 0x17, 0x24, 0x75, 0x30, 0x96, 0xf4, 0x09, 0xb0, 0x3f, 0x6e, 0xb4, 0x6a, 0x8d, 0xc0, 0x3b,
	0x98, 0x06, 0x17, 0xb7, 0x24, 0x5b, 0x64, 0xa5, 0xae, 0xcb, 0xa8, 0x8a, 0x3d, 0xba, 0xee, 0x0f,
	0x9b, 0x68, 0x7e, 0x0b, 0x67, 0x5c, 0x1a, 0xfd, 0xc9, 0x7e, 0x56, 0x2d, 0xdb, 0x44, 0xc7, 0x2c,
	0xb0, 0x37, 0xb6, 0x11, 0x58, 0x02, 0xec, 0xdf, 0x45, 0x26, 0xce, 0x22, 0x77, 0x16, 0x3b, 0x5a,
	0xf5, 0x14, 0xf4, 0x19, 0xae, 0x0e, 0x8c, 0x42, 0xcc, 0x7b, 0x38, 0x09, 0x51, 0x0c, 0x49, 0xe6,
	0x93, 0x61, 0xce, 0x5d, 0x73, 0xf1, 0x9b, 0x40, 0x1e, 0xe8, 0x87, 0x3f, 0x31, 0x2e, 0x01, 0xf6,
	0x2b, 0x71, 0xf6, 0x3f, 0x37, 0xba, 0x63, 0x71, 0x3d, 0xc4, 0xde, 0x98, 0x1e, 0xe1, 0x0b, 0x64,
	0xbd, 0x44, 0x38, 0x10, 0xc6, 0x5b, 0x14, 0x37, 0x23, 0x1e, 0x37, 0xd4, 0xc7, 0xee, 0xff, 0x1e,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xeb, 0x9b, 0xc2, 0xfd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, "/pb.ProductService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, "/pb.ProductService/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) GetProduct(ctx context.Context, req *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (*UnimplementedProductServiceServer) GetProducts(ctx context.Context, req *GetProductsRequest) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductService/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProducts(ctx, req.(*GetProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _ProductService_GetProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_service.proto",
}
