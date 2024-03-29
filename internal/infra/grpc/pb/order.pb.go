// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/infra/grpc/protofiles/order.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Blank struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blank) Reset()         { *m = Blank{} }
func (m *Blank) String() string { return proto.CompactTextString(m) }
func (*Blank) ProtoMessage()    {}
func (*Blank) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{0}
}

func (m *Blank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blank.Unmarshal(m, b)
}
func (m *Blank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blank.Marshal(b, m, deterministic)
}
func (m *Blank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blank.Merge(m, src)
}
func (m *Blank) XXX_Size() int {
	return xxx_messageInfo_Blank.Size(m)
}
func (m *Blank) XXX_DiscardUnknown() {
	xxx_messageInfo_Blank.DiscardUnknown(m)
}

var xxx_messageInfo_Blank proto.InternalMessageInfo

type CreateOrderRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax                  float32  `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{1}
}

func (m *CreateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequest.Unmarshal(m, b)
}
func (m *CreateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequest.Merge(m, src)
}
func (m *CreateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequest.Size(m)
}
func (m *CreateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequest proto.InternalMessageInfo

func (m *CreateOrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateOrderRequest) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CreateOrderRequest) GetTax() float32 {
	if m != nil {
		return m.Tax
	}
	return 0
}

type CreateOrderResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax                  float32  `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice           float32  `protobuf:"fixed32,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderResponse) Reset()         { *m = CreateOrderResponse{} }
func (m *CreateOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOrderResponse) ProtoMessage()    {}
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{2}
}

func (m *CreateOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderResponse.Unmarshal(m, b)
}
func (m *CreateOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderResponse.Marshal(b, m, deterministic)
}
func (m *CreateOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderResponse.Merge(m, src)
}
func (m *CreateOrderResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOrderResponse.Size(m)
}
func (m *CreateOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderResponse proto.InternalMessageInfo

func (m *CreateOrderResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateOrderResponse) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CreateOrderResponse) GetTax() float32 {
	if m != nil {
		return m.Tax
	}
	return 0
}

func (m *CreateOrderResponse) GetFinalPrice() float32 {
	if m != nil {
		return m.FinalPrice
	}
	return 0
}

type ListOrderResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax                  float32  `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice           float32  `protobuf:"fixed32,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOrderResponse) Reset()         { *m = ListOrderResponse{} }
func (m *ListOrderResponse) String() string { return proto.CompactTextString(m) }
func (*ListOrderResponse) ProtoMessage()    {}
func (*ListOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{3}
}

func (m *ListOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOrderResponse.Unmarshal(m, b)
}
func (m *ListOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOrderResponse.Marshal(b, m, deterministic)
}
func (m *ListOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOrderResponse.Merge(m, src)
}
func (m *ListOrderResponse) XXX_Size() int {
	return xxx_messageInfo_ListOrderResponse.Size(m)
}
func (m *ListOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListOrderResponse proto.InternalMessageInfo

func (m *ListOrderResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ListOrderResponse) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ListOrderResponse) GetTax() float32 {
	if m != nil {
		return m.Tax
	}
	return 0
}

func (m *ListOrderResponse) GetFinalPrice() float32 {
	if m != nil {
		return m.FinalPrice
	}
	return 0
}

type ListOrdersResponse struct {
	Orders               []*ListOrderResponse `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListOrdersResponse) Reset()         { *m = ListOrdersResponse{} }
func (m *ListOrdersResponse) String() string { return proto.CompactTextString(m) }
func (*ListOrdersResponse) ProtoMessage()    {}
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{4}
}

func (m *ListOrdersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOrdersResponse.Unmarshal(m, b)
}
func (m *ListOrdersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOrdersResponse.Marshal(b, m, deterministic)
}
func (m *ListOrdersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOrdersResponse.Merge(m, src)
}
func (m *ListOrdersResponse) XXX_Size() int {
	return xxx_messageInfo_ListOrdersResponse.Size(m)
}
func (m *ListOrdersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOrdersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListOrdersResponse proto.InternalMessageInfo

func (m *ListOrdersResponse) GetOrders() []*ListOrderResponse {
	if m != nil {
		return m.Orders
	}
	return nil
}

func init() {
	proto.RegisterType((*Blank)(nil), "pb.blank")
	proto.RegisterType((*CreateOrderRequest)(nil), "pb.CreateOrderRequest")
	proto.RegisterType((*CreateOrderResponse)(nil), "pb.CreateOrderResponse")
	proto.RegisterType((*ListOrderResponse)(nil), "pb.ListOrderResponse")
	proto.RegisterType((*ListOrdersResponse)(nil), "pb.ListOrdersResponse")
}

func init() {
	proto.RegisterFile("internal/infra/grpc/protofiles/order.proto", fileDescriptor_a81083cb550ac9a5)
}

var fileDescriptor_a81083cb550ac9a5 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0x41, 0x4b, 0xf4, 0x30,
	0x10, 0xfd, 0xda, 0x7e, 0xbb, 0xd2, 0xa9, 0x88, 0x8e, 0x5a, 0xcb, 0x5e, 0x2c, 0x3d, 0x15, 0xc1,
	0x56, 0xd6, 0x3f, 0x20, 0xee, 0x75, 0x41, 0xa9, 0x37, 0x2f, 0x92, 0x6e, 0xa7, 0x12, 0x2c, 0x69,
	0x4c, 0xa2, 0x78, 0xf6, 0x97, 0x4b, 0x52, 0xa9, 0x2e, 0xf5, 0xe4, 0xc1, 0x5b, 0xf2, 0xe6, 0xcd,
	0x7b, 0x8f, 0x99, 0x81, 0x33, 0x2e, 0x0c, 0x29, 0xc1, 0xba, 0x92, 0x8b, 0x56, 0xb1, 0xf2, 0x51,
	0xc9, 0x4d, 0x29, 0x55, 0x6f, 0xfa, 0x96, 0x77, 0xa4, 0xcb, 0x5e, 0x35, 0xa4, 0x0a, 0x07, 0xa0,
	0x2f, 0xeb, 0x6c, 0x07, 0x66, 0x75, 0xc7, 0xc4, 0x53, 0xb6, 0x06, 0x5c, 0x29, 0x62, 0x86, 0x6e,
	0x2c, 0xa3, 0xa2, 0xe7, 0x17, 0xd2, 0x06, 0xf7, 0xc0, 0xe7, 0x4d, 0xe2, 0xa5, 0x5e, 0x1e, 0x56,
	0x3e, 0x6f, 0xf0, 0x08, 0x66, 0x52, 0xf1, 0x0d, 0x25, 0x7e, 0xea, 0xe5, 0x7e, 0x35, 0x7c, 0x70,
	0x1f, 0x02, 0xc3, 0xde, 0x92, 0xc0, 0x61, 0xf6, 0x99, 0x09, 0x38, 0xdc, 0x52, 0xd3, 0xb2, 0x17,
	0x9a, 0x7e, 0x2b, 0x87, 0xa7, 0x10, 0xb5, 0x5c, 0xb0, 0xee, 0x61, 0x60, 0xff, 0x77, 0x15, 0x70,
	0xd0, 0xad, 0x45, 0xb2, 0x0e, 0x0e, 0xd6, 0x5c, 0x9b, 0x3f, 0x72, 0x5b, 0x01, 0x8e, 0x6e, 0x7a,
	0xb4, 0x3b, 0x87, 0xb9, 0x9b, 0xae, 0x4e, 0xbc, 0x34, 0xc8, 0xa3, 0xe5, 0x71, 0x21, 0xeb, 0x62,
	0x92, 0xaa, 0xfa, 0x24, 0x2d, 0xdf, 0x3d, 0xd8, 0x75, 0x95, 0x3b, 0x52, 0xaf, 0x36, 0xc8, 0x15,
	0x44, 0xdf, 0x66, 0x86, 0xb1, 0x6d, 0x9f, 0xae, 0x64, 0x71, 0x32, 0xc1, 0x07, 0xe1, 0xec, 0x1f,
	0x5e, 0x40, 0x38, 0xfa, 0x61, 0x68, 0x79, 0x6e, 0xb7, 0x8b, 0x78, 0x2b, 0x89, 0xfe, 0xea, 0xb8,
	0x4e, 0xee, 0xe3, 0x1f, 0x0f, 0xa6, 0xae, 0xe7, 0xee, 0x46, 0x2e, 0x3f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x1c, 0xb4, 0x2b, 0xe8, 0x51, 0x02, 0x00, 0x00,
}
