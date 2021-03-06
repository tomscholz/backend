// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store_product_service.proto

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

type GetStoreProductRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStoreProductRequest) Reset()         { *m = GetStoreProductRequest{} }
func (m *GetStoreProductRequest) String() string { return proto.CompactTextString(m) }
func (*GetStoreProductRequest) ProtoMessage()    {}
func (*GetStoreProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_314586fa579a3cca, []int{0}
}

func (m *GetStoreProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoreProductRequest.Unmarshal(m, b)
}
func (m *GetStoreProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoreProductRequest.Marshal(b, m, deterministic)
}
func (m *GetStoreProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoreProductRequest.Merge(m, src)
}
func (m *GetStoreProductRequest) XXX_Size() int {
	return xxx_messageInfo_GetStoreProductRequest.Size(m)
}
func (m *GetStoreProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoreProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoreProductRequest proto.InternalMessageInfo

func (m *GetStoreProductRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetStoreProductResponse struct {
	StoreProduct         *StoreProduct `protobuf:"bytes,1,opt,name=store_product,json=storeProduct,proto3" json:"store_product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetStoreProductResponse) Reset()         { *m = GetStoreProductResponse{} }
func (m *GetStoreProductResponse) String() string { return proto.CompactTextString(m) }
func (*GetStoreProductResponse) ProtoMessage()    {}
func (*GetStoreProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_314586fa579a3cca, []int{1}
}

func (m *GetStoreProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoreProductResponse.Unmarshal(m, b)
}
func (m *GetStoreProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoreProductResponse.Marshal(b, m, deterministic)
}
func (m *GetStoreProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoreProductResponse.Merge(m, src)
}
func (m *GetStoreProductResponse) XXX_Size() int {
	return xxx_messageInfo_GetStoreProductResponse.Size(m)
}
func (m *GetStoreProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoreProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoreProductResponse proto.InternalMessageInfo

func (m *GetStoreProductResponse) GetStoreProduct() *StoreProduct {
	if m != nil {
		return m.StoreProduct
	}
	return nil
}

type GetStoreProductsRequest struct {
	StoreId              string                `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	ProductId            string                `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity             StoreProduct_Quantity `protobuf:"varint,3,opt,name=quantity,proto3,enum=pb.StoreProduct_Quantity" json:"quantity,omitempty"`
	Pagination           *Pagination           `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetStoreProductsRequest) Reset()         { *m = GetStoreProductsRequest{} }
func (m *GetStoreProductsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStoreProductsRequest) ProtoMessage()    {}
func (*GetStoreProductsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_314586fa579a3cca, []int{2}
}

func (m *GetStoreProductsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoreProductsRequest.Unmarshal(m, b)
}
func (m *GetStoreProductsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoreProductsRequest.Marshal(b, m, deterministic)
}
func (m *GetStoreProductsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoreProductsRequest.Merge(m, src)
}
func (m *GetStoreProductsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStoreProductsRequest.Size(m)
}
func (m *GetStoreProductsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoreProductsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoreProductsRequest proto.InternalMessageInfo

func (m *GetStoreProductsRequest) GetStoreId() string {
	if m != nil {
		return m.StoreId
	}
	return ""
}

func (m *GetStoreProductsRequest) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *GetStoreProductsRequest) GetQuantity() StoreProduct_Quantity {
	if m != nil {
		return m.Quantity
	}
	return StoreProduct_NONE
}

func (m *GetStoreProductsRequest) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type GetStoreProductsResponse struct {
	StoreProducts        []*StoreProduct `protobuf:"bytes,1,rep,name=store_products,json=storeProducts,proto3" json:"store_products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetStoreProductsResponse) Reset()         { *m = GetStoreProductsResponse{} }
func (m *GetStoreProductsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStoreProductsResponse) ProtoMessage()    {}
func (*GetStoreProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_314586fa579a3cca, []int{3}
}

func (m *GetStoreProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoreProductsResponse.Unmarshal(m, b)
}
func (m *GetStoreProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoreProductsResponse.Marshal(b, m, deterministic)
}
func (m *GetStoreProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoreProductsResponse.Merge(m, src)
}
func (m *GetStoreProductsResponse) XXX_Size() int {
	return xxx_messageInfo_GetStoreProductsResponse.Size(m)
}
func (m *GetStoreProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoreProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoreProductsResponse proto.InternalMessageInfo

func (m *GetStoreProductsResponse) GetStoreProducts() []*StoreProduct {
	if m != nil {
		return m.StoreProducts
	}
	return nil
}

func init() {
	proto.RegisterType((*GetStoreProductRequest)(nil), "pb.GetStoreProductRequest")
	proto.RegisterType((*GetStoreProductResponse)(nil), "pb.GetStoreProductResponse")
	proto.RegisterType((*GetStoreProductsRequest)(nil), "pb.GetStoreProductsRequest")
	proto.RegisterType((*GetStoreProductsResponse)(nil), "pb.GetStoreProductsResponse")
}

func init() {
	proto.RegisterFile("store_product_service.proto", fileDescriptor_314586fa579a3cca)
}

var fileDescriptor_314586fa579a3cca = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xdd, 0x4e, 0x83, 0x30,
	0x14, 0xb6, 0xcc, 0xe8, 0x76, 0x74, 0xb8, 0x74, 0x89, 0x76, 0x4c, 0x93, 0x85, 0x2b, 0xae, 0xb8,
	0xc0, 0x2c, 0x3e, 0x82, 0x59, 0x62, 0x22, 0xc2, 0x03, 0x2c, 0xb0, 0x36, 0xa6, 0x37, 0xb4, 0xa3,
	0xc5, 0xc4, 0xf7, 0xf2, 0xca, 0xa7, 0x33, 0x50, 0x40, 0xfe, 0x2e, 0x7b, 0xbe, 0x9f, 0x7e, 0xfd,
	0x4e, 0x61, 0xab, 0xb4, 0xc8, 0xd9, 0x51, 0xe6, 0x82, 0x16, 0x27, 0x7d, 0x54, 0x2c, 0xff, 0xe2,
	0x27, 0xe6, 0xcb, 0x5c, 0x68, 0x81, 0x2d, 0x99, 0x3a, 0x2b, 0x99, 0x7c, 0xf2, 0x2c, 0xd1, 0x5c,
	0x64, 0x66, 0xea, 0xac, 0x7b, 0x12, 0x33, 0x74, 0x3d, 0xb8, 0x7f, 0x65, 0x3a, 0x2e, 0x91, 0xd0,
	0x00, 0x11, 0x3b, 0x17, 0x4c, 0x69, 0x6c, 0x83, 0xc5, 0x29, 0x41, 0x3b, 0xe4, 0x2d, 0x22, 0x8b,
	0x53, 0x37, 0x84, 0x87, 0x11, 0x53, 0x49, 0x91, 0x29, 0x86, 0xf7, 0xb0, 0xec, 0x79, 0x57, 0xaa,
	0x9b, 0x60, 0xe5, 0xcb, 0xd4, 0xef, 0x09, 0x6e, 0x55, 0xe7, 0xe4, 0xfe, 0xa2, 0x91, 0xa5, 0x6a,
	0x6e, 0xdf, 0xc0, 0xdc, 0x58, 0xb6, 0x19, 0xae, 0xab, 0xf3, 0x81, 0xe2, 0x27, 0x80, 0xe6, 0xd9,
	0x9c, 0x12, 0xab, 0x02, 0x17, 0xf5, 0xe4, 0x40, 0xf1, 0x1e, 0xe6, 0xe7, 0x22, 0xc9, 0x34, 0xd7,
	0xdf, 0x64, 0xb6, 0x43, 0x9e, 0x1d, 0x6c, 0x86, 0x39, 0xfc, 0x8f, 0x9a, 0x10, 0xb5, 0x54, 0xec,
	0x03, 0xfc, 0x37, 0x46, 0x2e, 0xab, 0x07, 0xd8, 0xa5, 0x30, 0x6c, 0xa7, 0x51, 0x87, 0xe1, 0xc6,
	0x40, 0xc6, 0xd9, 0xeb, 0x3e, 0x5e, 0xc0, 0xee, 0xf5, 0xa1, 0x08, 0xda, 0xcd, 0x26, 0x0b, 0x59,
	0x76, 0x0b, 0x51, 0xc1, 0x0f, 0x82, 0x75, 0x17, 0x8f, 0xcd, 0x5a, 0xf1, 0x1b, 0xdc, 0x0d, 0x2e,
	0xc3, 0x4e, 0xe9, 0x35, 0xbd, 0x3a, 0x67, 0x3b, 0x89, 0x99, 0x70, 0xee, 0x05, 0x7e, 0x87, 0xd5,
	0x30, 0x3a, 0x9e, 0x92, 0x34, 0xcb, 0x70, 0x1e, 0xa7, 0xc1, 0xc6, 0x30, 0xbd, 0xaa, 0xfe, 0xd2,
	0xf3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x53, 0x75, 0xea, 0x95, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StoreProductServiceClient is the client API for StoreProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreProductServiceClient interface {
	GetStoreProduct(ctx context.Context, in *GetStoreProductRequest, opts ...grpc.CallOption) (*GetStoreProductResponse, error)
	GetStoreProducts(ctx context.Context, in *GetStoreProductsRequest, opts ...grpc.CallOption) (*GetStoreProductsResponse, error)
}

type storeProductServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreProductServiceClient(cc grpc.ClientConnInterface) StoreProductServiceClient {
	return &storeProductServiceClient{cc}
}

func (c *storeProductServiceClient) GetStoreProduct(ctx context.Context, in *GetStoreProductRequest, opts ...grpc.CallOption) (*GetStoreProductResponse, error) {
	out := new(GetStoreProductResponse)
	err := c.cc.Invoke(ctx, "/pb.StoreProductService/GetStoreProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeProductServiceClient) GetStoreProducts(ctx context.Context, in *GetStoreProductsRequest, opts ...grpc.CallOption) (*GetStoreProductsResponse, error) {
	out := new(GetStoreProductsResponse)
	err := c.cc.Invoke(ctx, "/pb.StoreProductService/GetStoreProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreProductServiceServer is the server API for StoreProductService service.
type StoreProductServiceServer interface {
	GetStoreProduct(context.Context, *GetStoreProductRequest) (*GetStoreProductResponse, error)
	GetStoreProducts(context.Context, *GetStoreProductsRequest) (*GetStoreProductsResponse, error)
}

// UnimplementedStoreProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStoreProductServiceServer struct {
}

func (*UnimplementedStoreProductServiceServer) GetStoreProduct(ctx context.Context, req *GetStoreProductRequest) (*GetStoreProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStoreProduct not implemented")
}
func (*UnimplementedStoreProductServiceServer) GetStoreProducts(ctx context.Context, req *GetStoreProductsRequest) (*GetStoreProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStoreProducts not implemented")
}

func RegisterStoreProductServiceServer(s *grpc.Server, srv StoreProductServiceServer) {
	s.RegisterService(&_StoreProductService_serviceDesc, srv)
}

func _StoreProductService_GetStoreProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoreProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreProductServiceServer).GetStoreProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StoreProductService/GetStoreProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreProductServiceServer).GetStoreProduct(ctx, req.(*GetStoreProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreProductService_GetStoreProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoreProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreProductServiceServer).GetStoreProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StoreProductService/GetStoreProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreProductServiceServer).GetStoreProducts(ctx, req.(*GetStoreProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StoreProductService",
	HandlerType: (*StoreProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStoreProduct",
			Handler:    _StoreProductService_GetStoreProduct_Handler,
		},
		{
			MethodName: "GetStoreProducts",
			Handler:    _StoreProductService_GetStoreProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store_product_service.proto",
}
