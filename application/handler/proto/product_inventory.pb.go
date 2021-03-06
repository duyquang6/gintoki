// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: product_inventory.proto

package proto

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type GetMultiProductInventoryRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	ProductIDs           []int64  `protobuf:"varint,2,rep,packed,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMultiProductInventoryRequest) Reset()         { *m = GetMultiProductInventoryRequest{} }
func (m *GetMultiProductInventoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetMultiProductInventoryRequest) ProtoMessage()    {}
func (*GetMultiProductInventoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8bd4aae1db7d82e, []int{0}
}
func (m *GetMultiProductInventoryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetMultiProductInventoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetMultiProductInventoryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetMultiProductInventoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMultiProductInventoryRequest.Merge(m, src)
}
func (m *GetMultiProductInventoryRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetMultiProductInventoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMultiProductInventoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMultiProductInventoryRequest proto.InternalMessageInfo

func (m *GetMultiProductInventoryRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetMultiProductInventoryRequest) GetProductIDs() []int64 {
	if m != nil {
		return m.ProductIDs
	}
	return nil
}

type GetMultiProductInventoryResponse struct {
	Meta                 *Meta                           `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Data                 []*GetMultiProductInventoryData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *GetMultiProductInventoryResponse) Reset()         { *m = GetMultiProductInventoryResponse{} }
func (m *GetMultiProductInventoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetMultiProductInventoryResponse) ProtoMessage()    {}
func (*GetMultiProductInventoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8bd4aae1db7d82e, []int{1}
}
func (m *GetMultiProductInventoryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetMultiProductInventoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetMultiProductInventoryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetMultiProductInventoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMultiProductInventoryResponse.Merge(m, src)
}
func (m *GetMultiProductInventoryResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetMultiProductInventoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMultiProductInventoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMultiProductInventoryResponse proto.InternalMessageInfo

func (m *GetMultiProductInventoryResponse) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetMultiProductInventoryResponse) GetData() []*GetMultiProductInventoryData {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetMultiProductInventoryData struct {
	ProductID            int64                 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Data                 []*WarehouseInventory `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetMultiProductInventoryData) Reset()         { *m = GetMultiProductInventoryData{} }
func (m *GetMultiProductInventoryData) String() string { return proto.CompactTextString(m) }
func (*GetMultiProductInventoryData) ProtoMessage()    {}
func (*GetMultiProductInventoryData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8bd4aae1db7d82e, []int{2}
}
func (m *GetMultiProductInventoryData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetMultiProductInventoryData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetMultiProductInventoryData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetMultiProductInventoryData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMultiProductInventoryData.Merge(m, src)
}
func (m *GetMultiProductInventoryData) XXX_Size() int {
	return m.Size()
}
func (m *GetMultiProductInventoryData) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMultiProductInventoryData.DiscardUnknown(m)
}

var xxx_messageInfo_GetMultiProductInventoryData proto.InternalMessageInfo

func (m *GetMultiProductInventoryData) GetProductID() int64 {
	if m != nil {
		return m.ProductID
	}
	return 0
}

func (m *GetMultiProductInventoryData) GetData() []*WarehouseInventory {
	if m != nil {
		return m.Data
	}
	return nil
}

type WarehouseInventory struct {
	WarehouseCode        string   `protobuf:"bytes,1,opt,name=warehouse_code,json=warehouseCode,proto3" json:"warehouse_code,omitempty"`
	QtySalable           int32    `protobuf:"varint,2,opt,name=qty_salable,json=qtySalable,proto3" json:"qty_salable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WarehouseInventory) Reset()         { *m = WarehouseInventory{} }
func (m *WarehouseInventory) String() string { return proto.CompactTextString(m) }
func (*WarehouseInventory) ProtoMessage()    {}
func (*WarehouseInventory) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8bd4aae1db7d82e, []int{3}
}
func (m *WarehouseInventory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WarehouseInventory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WarehouseInventory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WarehouseInventory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WarehouseInventory.Merge(m, src)
}
func (m *WarehouseInventory) XXX_Size() int {
	return m.Size()
}
func (m *WarehouseInventory) XXX_DiscardUnknown() {
	xxx_messageInfo_WarehouseInventory.DiscardUnknown(m)
}

var xxx_messageInfo_WarehouseInventory proto.InternalMessageInfo

func (m *WarehouseInventory) GetWarehouseCode() string {
	if m != nil {
		return m.WarehouseCode
	}
	return ""
}

func (m *WarehouseInventory) GetQtySalable() int32 {
	if m != nil {
		return m.QtySalable
	}
	return 0
}

func init() {
	proto.RegisterType((*GetMultiProductInventoryRequest)(nil), "proto.GetMultiProductInventoryRequest")
	proto.RegisterType((*GetMultiProductInventoryResponse)(nil), "proto.GetMultiProductInventoryResponse")
	proto.RegisterType((*GetMultiProductInventoryData)(nil), "proto.GetMultiProductInventoryData")
	proto.RegisterType((*WarehouseInventory)(nil), "proto.WarehouseInventory")
}

func init() { proto.RegisterFile("product_inventory.proto", fileDescriptor_d8bd4aae1db7d82e) }

var fileDescriptor_d8bd4aae1db7d82e = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0x76, 0x9a, 0x7b, 0x85, 0x9e, 0xd0, 0x52, 0x06, 0xa4, 0x31, 0x94, 0x24, 0x44, 0xd4, 0x22,
	0xb6, 0xc1, 0xba, 0x10, 0x5c, 0xd6, 0x82, 0x74, 0x51, 0x90, 0x74, 0xe1, 0x46, 0x28, 0x93, 0x64,
	0x4c, 0x03, 0x69, 0x26, 0xcd, 0x4c, 0x2a, 0x45, 0x57, 0xee, 0x5d, 0xb9, 0xf1, 0x15, 0x7c, 0x13,
	0x97, 0x82, 0xfb, 0x22, 0xd1, 0x07, 0x91, 0x4c, 0xd2, 0x14, 0x95, 0xde, 0xae, 0x72, 0xe6, 0xfc,
	0x7d, 0x3f, 0x27, 0xd0, 0x4f, 0x33, 0x16, 0xe4, 0xbe, 0x58, 0x45, 0xc9, 0x8e, 0x26, 0x82, 0x65,
	0xfb, 0x71, 0x9a, 0x31, 0xc1, 0xf0, 0xb5, 0xfc, 0xe8, 0xa3, 0x30, 0x12, 0xeb, 0xdc, 0x1b, 0xfb,
	0x6c, 0xe3, 0x84, 0x2c, 0x64, 0x8e, 0x4c, 0x7b, 0xf9, 0x5b, 0xf9, 0x92, 0x0f, 0x19, 0x55, 0x53,
	0x3a, 0x6c, 0xa8, 0x20, 0x75, 0x3c, 0x08, 0x19, 0x0b, 0x63, 0xea, 0x90, 0x34, 0x72, 0x48, 0x92,
	0x30, 0x41, 0x44, 0xc4, 0x12, 0x5e, 0x55, 0xed, 0x00, 0xcc, 0x97, 0x54, 0x2c, 0xf2, 0x58, 0x44,
	0xaf, 0x2a, 0x0a, 0xf3, 0x23, 0x03, 0x97, 0x6e, 0x73, 0xca, 0x05, 0xee, 0x81, 0x42, 0xd2, 0x48,
	0x43, 0x16, 0x1a, 0xb6, 0xdd, 0x32, 0xc4, 0x0e, 0xa8, 0x0d, 0xdf, 0x80, 0x6b, 0x2d, 0x4b, 0x19,
	0x2a, 0xd3, 0x6e, 0x71, 0x30, 0xe1, 0xb8, 0x63, 0xc6, 0x5d, 0xa8, 0x5b, 0xe6, 0x01, 0xb7, 0x3f,
	0x80, 0x75, 0x1e, 0x85, 0xa7, 0x2c, 0xe1, 0x14, 0x9b, 0x70, 0x55, 0xb2, 0x96, 0x38, 0xea, 0x44,
	0xad, 0xf8, 0x8d, 0x17, 0x54, 0x10, 0x57, 0x16, 0xf0, 0x33, 0xb8, 0x0a, 0x88, 0x20, 0x12, 0x4e,
	0x9d, 0xdc, 0xab, 0x1b, 0xce, 0xed, 0x9d, 0x91, 0x72, 0xb0, 0x1c, 0xb0, 0xdf, 0xc3, 0xe0, 0xa6,
	0x2e, 0xfc, 0x18, 0xe0, 0x24, 0x47, 0xe2, 0x2b, 0xd3, 0x4e, 0x71, 0x30, 0xdb, 0x8d, 0x1a, 0xb7,
	0xdd, 0x88, 0xc1, 0xa3, 0xbf, 0x68, 0xdc, 0xad, 0x69, 0xbc, 0x26, 0x19, 0x5d, 0xb3, 0x9c, 0xd3,
	0x93, 0xb0, 0x0a, 0xfc, 0x0d, 0xe0, 0xff, 0x6b, 0xf8, 0x3e, 0x74, 0xdf, 0x1d, 0xb3, 0x2b, 0x9f,
	0x05, 0xb4, 0xb6, 0xb7, 0xd3, 0x64, 0x5f, 0xb0, 0xa0, 0xf4, 0x44, 0xdd, 0x8a, 0xfd, 0x8a, 0x93,
	0x98, 0x78, 0x31, 0xd5, 0x5a, 0x16, 0x1a, 0x5e, 0xbb, 0xb0, 0x15, 0xfb, 0x65, 0x95, 0x99, 0x7c,
	0x45, 0xd0, 0xff, 0x57, 0xd3, 0x92, 0x66, 0xbb, 0xc8, 0xa7, 0xf8, 0x13, 0x02, 0xed, 0x9c, 0x6e,
	0xfc, 0xe0, 0x82, 0x7d, 0xf5, 0xf1, 0xf5, 0x87, 0x17, 0xfb, 0xaa, 0xf3, 0xd9, 0xd6, 0xc7, 0x1f,
	0xbf, 0x3f, 0xb7, 0x74, 0xfb, 0x8e, 0xb3, 0x7b, 0xe2, 0xd4, 0x6e, 0x8d, 0x9a, 0xbf, 0xf9, 0x39,
	0x7a, 0x34, 0xed, 0x7d, 0x2b, 0x0c, 0xf4, 0xbd, 0x30, 0xd0, 0xcf, 0xc2, 0x40, 0x5f, 0x7e, 0x19,
	0xb7, 0xbc, 0xdb, 0x72, 0xf7, 0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xdc, 0xaa, 0xc9,
	0xfe, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductInventoryServiceClient is the client API for ProductInventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductInventoryServiceClient interface {
	GetMultiProductInventory(ctx context.Context, in *GetMultiProductInventoryRequest, opts ...grpc.CallOption) (*GetMultiProductInventoryResponse, error)
}

type productInventoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductInventoryServiceClient(cc *grpc.ClientConn) ProductInventoryServiceClient {
	return &productInventoryServiceClient{cc}
}

func (c *productInventoryServiceClient) GetMultiProductInventory(ctx context.Context, in *GetMultiProductInventoryRequest, opts ...grpc.CallOption) (*GetMultiProductInventoryResponse, error) {
	out := new(GetMultiProductInventoryResponse)
	err := c.cc.Invoke(ctx, "/proto.ProductInventoryService/GetMultiProductInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductInventoryServiceServer is the server API for ProductInventoryService service.
type ProductInventoryServiceServer interface {
	GetMultiProductInventory(context.Context, *GetMultiProductInventoryRequest) (*GetMultiProductInventoryResponse, error)
}

// UnimplementedProductInventoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductInventoryServiceServer struct {
}

func (*UnimplementedProductInventoryServiceServer) GetMultiProductInventory(ctx context.Context, req *GetMultiProductInventoryRequest) (*GetMultiProductInventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultiProductInventory not implemented")
}

func RegisterProductInventoryServiceServer(s *grpc.Server, srv ProductInventoryServiceServer) {
	s.RegisterService(&_ProductInventoryService_serviceDesc, srv)
}

func _ProductInventoryService_GetMultiProductInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMultiProductInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInventoryServiceServer).GetMultiProductInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductInventoryService/GetMultiProductInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInventoryServiceServer).GetMultiProductInventory(ctx, req.(*GetMultiProductInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductInventoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProductInventoryService",
	HandlerType: (*ProductInventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMultiProductInventory",
			Handler:    _ProductInventoryService_GetMultiProductInventory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_inventory.proto",
}

func (m *GetMultiProductInventoryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMultiProductInventoryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetMultiProductInventoryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ProductIDs) > 0 {
		dAtA2 := make([]byte, len(m.ProductIDs)*10)
		var j1 int
		for _, num1 := range m.ProductIDs {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintProductInventory(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Api) > 0 {
		i -= len(m.Api)
		copy(dAtA[i:], m.Api)
		i = encodeVarintProductInventory(dAtA, i, uint64(len(m.Api)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetMultiProductInventoryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMultiProductInventoryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetMultiProductInventoryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProductInventory(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Meta != nil {
		{
			size, err := m.Meta.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProductInventory(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetMultiProductInventoryData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMultiProductInventoryData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetMultiProductInventoryData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProductInventory(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.ProductID != 0 {
		i = encodeVarintProductInventory(dAtA, i, uint64(m.ProductID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *WarehouseInventory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WarehouseInventory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WarehouseInventory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.QtySalable != 0 {
		i = encodeVarintProductInventory(dAtA, i, uint64(m.QtySalable))
		i--
		dAtA[i] = 0x10
	}
	if len(m.WarehouseCode) > 0 {
		i -= len(m.WarehouseCode)
		copy(dAtA[i:], m.WarehouseCode)
		i = encodeVarintProductInventory(dAtA, i, uint64(len(m.WarehouseCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProductInventory(dAtA []byte, offset int, v uint64) int {
	offset -= sovProductInventory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetMultiProductInventoryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Api)
	if l > 0 {
		n += 1 + l + sovProductInventory(uint64(l))
	}
	if len(m.ProductIDs) > 0 {
		l = 0
		for _, e := range m.ProductIDs {
			l += sovProductInventory(uint64(e))
		}
		n += 1 + sovProductInventory(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetMultiProductInventoryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovProductInventory(uint64(l))
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovProductInventory(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetMultiProductInventoryData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProductID != 0 {
		n += 1 + sovProductInventory(uint64(m.ProductID))
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovProductInventory(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WarehouseInventory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.WarehouseCode)
	if l > 0 {
		n += 1 + l + sovProductInventory(uint64(l))
	}
	if m.QtySalable != 0 {
		n += 1 + sovProductInventory(uint64(m.QtySalable))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProductInventory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProductInventory(x uint64) (n int) {
	return sovProductInventory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetMultiProductInventoryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProductInventory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetMultiProductInventoryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMultiProductInventoryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Api", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductInventory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProductInventory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProductInventory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Api = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProductInventory
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ProductIDs = append(m.ProductIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProductInventory
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthProductInventory
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthProductInventory
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ProductIDs) == 0 {
					m.ProductIDs = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProductInventory
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ProductIDs = append(m.ProductIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductIDs", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProductInventory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProductInventory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProductInventory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetMultiProductInventoryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProductInventory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetMultiProductInventoryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMultiProductInventoryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductInventory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProductInventory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProductInventory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = &Meta{}
			}
			if err := m.Meta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductInventory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProductInventory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProductInventory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &GetMultiProductInventoryData{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProductInventory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProductInventory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProductInventory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetMultiProductInventoryData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProductInventory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetMultiProductInventoryData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMultiProductInventoryData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductID", wireType)
			}
			m.ProductID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductInventory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProductID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductInventory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProductInventory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProductInventory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &WarehouseInventory{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProductInventory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProductInventory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProductInventory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WarehouseInventory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProductInventory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WarehouseInventory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WarehouseInventory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WarehouseCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductInventory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProductInventory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProductInventory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WarehouseCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QtySalable", wireType)
			}
			m.QtySalable = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProductInventory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QtySalable |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProductInventory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProductInventory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProductInventory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProductInventory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProductInventory
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProductInventory
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProductInventory
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthProductInventory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProductInventory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProductInventory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProductInventory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProductInventory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProductInventory = fmt.Errorf("proto: unexpected end of group")
)
