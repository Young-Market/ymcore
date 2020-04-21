// Code generated by protoc-gen-go. DO NOT EDIT.
// source: waves/node/grpc/accounts_api.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	waves "github.com/wavesplatform/gowaves/pkg/grpc/generated/waves"
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

type AccountRequest struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRequest) Reset()         { *m = AccountRequest{} }
func (m *AccountRequest) String() string { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()    {}
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4fbafb2c22317fb, []int{0}
}

func (m *AccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRequest.Unmarshal(m, b)
}
func (m *AccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRequest.Marshal(b, m, deterministic)
}
func (m *AccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRequest.Merge(m, src)
}
func (m *AccountRequest) XXX_Size() int {
	return xxx_messageInfo_AccountRequest.Size(m)
}
func (m *AccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRequest proto.InternalMessageInfo

func (m *AccountRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type DataRequest struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataRequest) Reset()         { *m = DataRequest{} }
func (m *DataRequest) String() string { return proto.CompactTextString(m) }
func (*DataRequest) ProtoMessage()    {}
func (*DataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4fbafb2c22317fb, []int{1}
}

func (m *DataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRequest.Unmarshal(m, b)
}
func (m *DataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRequest.Marshal(b, m, deterministic)
}
func (m *DataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRequest.Merge(m, src)
}
func (m *DataRequest) XXX_Size() int {
	return xxx_messageInfo_DataRequest.Size(m)
}
func (m *DataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataRequest proto.InternalMessageInfo

func (m *DataRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *DataRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type BalancesRequest struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Assets               [][]byte `protobuf:"bytes,4,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalancesRequest) Reset()         { *m = BalancesRequest{} }
func (m *BalancesRequest) String() string { return proto.CompactTextString(m) }
func (*BalancesRequest) ProtoMessage()    {}
func (*BalancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4fbafb2c22317fb, []int{2}
}

func (m *BalancesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalancesRequest.Unmarshal(m, b)
}
func (m *BalancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalancesRequest.Marshal(b, m, deterministic)
}
func (m *BalancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalancesRequest.Merge(m, src)
}
func (m *BalancesRequest) XXX_Size() int {
	return xxx_messageInfo_BalancesRequest.Size(m)
}
func (m *BalancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalancesRequest proto.InternalMessageInfo

func (m *BalancesRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *BalancesRequest) GetAssets() [][]byte {
	if m != nil {
		return m.Assets
	}
	return nil
}

type BalanceResponse struct {
	// Types that are valid to be assigned to Balance:
	//	*BalanceResponse_Waves
	//	*BalanceResponse_Asset
	Balance              isBalanceResponse_Balance `protobuf_oneof:"balance"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BalanceResponse) Reset()         { *m = BalanceResponse{} }
func (m *BalanceResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceResponse) ProtoMessage()    {}
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4fbafb2c22317fb, []int{3}
}

func (m *BalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceResponse.Unmarshal(m, b)
}
func (m *BalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceResponse.Marshal(b, m, deterministic)
}
func (m *BalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceResponse.Merge(m, src)
}
func (m *BalanceResponse) XXX_Size() int {
	return xxx_messageInfo_BalanceResponse.Size(m)
}
func (m *BalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceResponse proto.InternalMessageInfo

type isBalanceResponse_Balance interface {
	isBalanceResponse_Balance()
}

type BalanceResponse_Waves struct {
	Waves *BalanceResponse_WavesBalances `protobuf:"bytes,1,opt,name=waves,proto3,oneof"`
}

type BalanceResponse_Asset struct {
	Asset *waves.Amount `protobuf:"bytes,2,opt,name=asset,proto3,oneof"`
}

func (*BalanceResponse_Waves) isBalanceResponse_Balance() {}

func (*BalanceResponse_Asset) isBalanceResponse_Balance() {}

func (m *BalanceResponse) GetBalance() isBalanceResponse_Balance {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *BalanceResponse) GetWaves() *BalanceResponse_WavesBalances {
	if x, ok := m.GetBalance().(*BalanceResponse_Waves); ok {
		return x.Waves
	}
	return nil
}

func (m *BalanceResponse) GetAsset() *waves.Amount {
	if x, ok := m.GetBalance().(*BalanceResponse_Asset); ok {
		return x.Asset
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BalanceResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BalanceResponse_Waves)(nil),
		(*BalanceResponse_Asset)(nil),
	}
}

type BalanceResponse_WavesBalances struct {
	Regular              int64    `protobuf:"varint,1,opt,name=regular,proto3" json:"regular,omitempty"`
	Generating           int64    `protobuf:"varint,2,opt,name=generating,proto3" json:"generating,omitempty"`
	Available            int64    `protobuf:"varint,3,opt,name=available,proto3" json:"available,omitempty"`
	Effective            int64    `protobuf:"varint,4,opt,name=effective,proto3" json:"effective,omitempty"`
	LeaseIn              int64    `protobuf:"varint,5,opt,name=lease_in,json=leaseIn,proto3" json:"lease_in,omitempty"`
	LeaseOut             int64    `protobuf:"varint,6,opt,name=lease_out,json=leaseOut,proto3" json:"lease_out,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceResponse_WavesBalances) Reset()         { *m = BalanceResponse_WavesBalances{} }
func (m *BalanceResponse_WavesBalances) String() string { return proto.CompactTextString(m) }
func (*BalanceResponse_WavesBalances) ProtoMessage()    {}
func (*BalanceResponse_WavesBalances) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4fbafb2c22317fb, []int{3, 0}
}

func (m *BalanceResponse_WavesBalances) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceResponse_WavesBalances.Unmarshal(m, b)
}
func (m *BalanceResponse_WavesBalances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceResponse_WavesBalances.Marshal(b, m, deterministic)
}
func (m *BalanceResponse_WavesBalances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceResponse_WavesBalances.Merge(m, src)
}
func (m *BalanceResponse_WavesBalances) XXX_Size() int {
	return xxx_messageInfo_BalanceResponse_WavesBalances.Size(m)
}
func (m *BalanceResponse_WavesBalances) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceResponse_WavesBalances.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceResponse_WavesBalances proto.InternalMessageInfo

func (m *BalanceResponse_WavesBalances) GetRegular() int64 {
	if m != nil {
		return m.Regular
	}
	return 0
}

func (m *BalanceResponse_WavesBalances) GetGenerating() int64 {
	if m != nil {
		return m.Generating
	}
	return 0
}

func (m *BalanceResponse_WavesBalances) GetAvailable() int64 {
	if m != nil {
		return m.Available
	}
	return 0
}

func (m *BalanceResponse_WavesBalances) GetEffective() int64 {
	if m != nil {
		return m.Effective
	}
	return 0
}

func (m *BalanceResponse_WavesBalances) GetLeaseIn() int64 {
	if m != nil {
		return m.LeaseIn
	}
	return 0
}

func (m *BalanceResponse_WavesBalances) GetLeaseOut() int64 {
	if m != nil {
		return m.LeaseOut
	}
	return 0
}

type DataEntryResponse struct {
	Address              []byte                               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Entry                *waves.DataTransactionData_DataEntry `protobuf:"bytes,2,opt,name=entry,proto3" json:"entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *DataEntryResponse) Reset()         { *m = DataEntryResponse{} }
func (m *DataEntryResponse) String() string { return proto.CompactTextString(m) }
func (*DataEntryResponse) ProtoMessage()    {}
func (*DataEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4fbafb2c22317fb, []int{4}
}

func (m *DataEntryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataEntryResponse.Unmarshal(m, b)
}
func (m *DataEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataEntryResponse.Marshal(b, m, deterministic)
}
func (m *DataEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataEntryResponse.Merge(m, src)
}
func (m *DataEntryResponse) XXX_Size() int {
	return xxx_messageInfo_DataEntryResponse.Size(m)
}
func (m *DataEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataEntryResponse proto.InternalMessageInfo

func (m *DataEntryResponse) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *DataEntryResponse) GetEntry() *waves.DataTransactionData_DataEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type ScriptData struct {
	ScriptBytes          []byte   `protobuf:"bytes,1,opt,name=script_bytes,json=scriptBytes,proto3" json:"script_bytes,omitempty"`
	ScriptText           string   `protobuf:"bytes,2,opt,name=script_text,json=scriptText,proto3" json:"script_text,omitempty"`
	Complexity           int64    `protobuf:"varint,3,opt,name=complexity,proto3" json:"complexity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScriptData) Reset()         { *m = ScriptData{} }
func (m *ScriptData) String() string { return proto.CompactTextString(m) }
func (*ScriptData) ProtoMessage()    {}
func (*ScriptData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4fbafb2c22317fb, []int{5}
}

func (m *ScriptData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScriptData.Unmarshal(m, b)
}
func (m *ScriptData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScriptData.Marshal(b, m, deterministic)
}
func (m *ScriptData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScriptData.Merge(m, src)
}
func (m *ScriptData) XXX_Size() int {
	return xxx_messageInfo_ScriptData.Size(m)
}
func (m *ScriptData) XXX_DiscardUnknown() {
	xxx_messageInfo_ScriptData.DiscardUnknown(m)
}

var xxx_messageInfo_ScriptData proto.InternalMessageInfo

func (m *ScriptData) GetScriptBytes() []byte {
	if m != nil {
		return m.ScriptBytes
	}
	return nil
}

func (m *ScriptData) GetScriptText() string {
	if m != nil {
		return m.ScriptText
	}
	return ""
}

func (m *ScriptData) GetComplexity() int64 {
	if m != nil {
		return m.Complexity
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountRequest)(nil), "waves.node.grpc.AccountRequest")
	proto.RegisterType((*DataRequest)(nil), "waves.node.grpc.DataRequest")
	proto.RegisterType((*BalancesRequest)(nil), "waves.node.grpc.BalancesRequest")
	proto.RegisterType((*BalanceResponse)(nil), "waves.node.grpc.BalanceResponse")
	proto.RegisterType((*BalanceResponse_WavesBalances)(nil), "waves.node.grpc.BalanceResponse.WavesBalances")
	proto.RegisterType((*DataEntryResponse)(nil), "waves.node.grpc.DataEntryResponse")
	proto.RegisterType((*ScriptData)(nil), "waves.node.grpc.ScriptData")
}

func init() { proto.RegisterFile("waves/node/grpc/accounts_api.proto", fileDescriptor_e4fbafb2c22317fb) }

var fileDescriptor_e4fbafb2c22317fb = []byte{
	// 661 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x3e, 0x69, 0x9a, 0xf6, 0x64, 0x92, 0x36, 0xe7, 0xec, 0x05, 0x18, 0xb7, 0x6a, 0x83, 0x55,
	0x50, 0xc5, 0x85, 0x8d, 0xca, 0x15, 0xdc, 0x25, 0x05, 0x42, 0x25, 0x04, 0x92, 0x5b, 0x81, 0xd4,
	0x9b, 0x68, 0xe3, 0x4c, 0xcc, 0xaa, 0x8e, 0xd7, 0xec, 0xae, 0xd3, 0xe6, 0x95, 0xfa, 0x0e, 0x3c,
	0x01, 0x2f, 0x85, 0x76, 0xd7, 0x4e, 0x9d, 0x84, 0x2a, 0x77, 0x9e, 0xf9, 0xbe, 0xf9, 0x76, 0xfe,
	0x3c, 0xe0, 0xdd, 0xd2, 0x19, 0xca, 0x20, 0xe5, 0x63, 0x0c, 0x62, 0x91, 0x45, 0x01, 0x8d, 0x22,
	0x9e, 0xa7, 0x4a, 0x0e, 0x69, 0xc6, 0xfc, 0x4c, 0x70, 0xc5, 0x49, 0xc7, 0x70, 0x7c, 0xcd, 0xf1,
	0x35, 0xc7, 0x7d, 0xb9, 0x1a, 0xa4, 0x04, 0x4d, 0x25, 0x8d, 0x14, 0xe3, 0x69, 0x25, 0xd0, 0x25,
	0x96, 0x47, 0xa7, 0x5a, 0xb1, 0xf0, 0x3d, 0xb5, 0xbe, 0x4a, 0x44, 0x01, 0x1c, 0xc5, 0x9c, 0xc7,
	0x09, 0x06, 0xc6, 0x1a, 0xe5, 0x93, 0xe0, 0x56, 0xd0, 0x2c, 0x43, 0x21, 0x2d, 0xee, 0xbd, 0x82,
	0xfd, 0x9e, 0xcd, 0x2d, 0xc4, 0x9f, 0x39, 0x4a, 0x45, 0x1c, 0xd8, 0xa5, 0xe3, 0xb1, 0x40, 0x29,
	0x9d, 0x5a, 0xb7, 0x76, 0xda, 0x0e, 0x4b, 0xd3, 0x7b, 0x0b, 0xad, 0xf7, 0x54, 0xd1, 0x8d, 0x44,
	0xf2, 0x1f, 0xd4, 0x6f, 0x70, 0xee, 0x6c, 0x75, 0x6b, 0xa7, 0xcd, 0x50, 0x7f, 0x7a, 0xe7, 0xd0,
	0xe9, 0xd3, 0x84, 0xa6, 0x11, 0xca, 0xcd, 0xe1, 0x4f, 0x60, 0x87, 0x4a, 0x89, 0x4a, 0x3a, 0xdb,
	0xdd, 0xfa, 0x69, 0x3b, 0x2c, 0x2c, 0xef, 0xf7, 0xd6, 0x42, 0x25, 0x44, 0x99, 0xf1, 0x54, 0x22,
	0xf9, 0x08, 0x0d, 0x53, 0xba, 0xd1, 0x68, 0x9d, 0xf9, 0xfe, 0x4a, 0x57, 0xfd, 0x95, 0x00, 0xff,
	0xbb, 0xc6, 0xcb, 0x5c, 0x3e, 0xfd, 0x13, 0xda, 0x70, 0xf2, 0x02, 0x1a, 0xe6, 0x15, 0x93, 0x74,
	0xeb, 0x6c, 0xaf, 0xd0, 0xe9, 0x99, 0x26, 0x6b, 0x9a, 0x41, 0xdd, 0x5f, 0x35, 0xd8, 0x5b, 0x52,
	0xd0, 0x65, 0x08, 0x8c, 0xf3, 0x84, 0x0a, 0x93, 0x42, 0x3d, 0x2c, 0x4d, 0x72, 0x04, 0x10, 0x63,
	0x8a, 0x82, 0x2a, 0x96, 0xc6, 0x46, 0xb7, 0x1e, 0x56, 0x3c, 0xe4, 0x10, 0x9a, 0x74, 0x46, 0x59,
	0x42, 0x47, 0x09, 0x3a, 0x75, 0x03, 0x3f, 0x38, 0x34, 0x8a, 0x93, 0x09, 0x46, 0x8a, 0xcd, 0xd0,
	0xd9, 0xb6, 0xe8, 0xc2, 0x41, 0x9e, 0xc1, 0xbf, 0x09, 0x52, 0x89, 0x43, 0x96, 0x3a, 0x0d, 0xfb,
	0xac, 0xb1, 0x2f, 0x52, 0x72, 0x00, 0x4d, 0x0b, 0xf1, 0x5c, 0x39, 0x3b, 0x06, 0xb3, 0xdc, 0xaf,
	0xb9, 0xea, 0x37, 0x61, 0x77, 0x64, 0x33, 0xf7, 0x18, 0xfc, 0xaf, 0xa7, 0xf9, 0x21, 0x55, 0x62,
	0xbe, 0x68, 0xe7, 0xe3, 0x43, 0x79, 0x07, 0x0d, 0xd4, 0xd4, 0xa2, 0x41, 0x27, 0x45, 0x83, 0xb4,
	0xc4, 0xd5, 0xc3, 0xd6, 0x69, 0xd3, 0x7f, 0x90, 0xb5, 0x21, 0x5e, 0x06, 0x70, 0x19, 0x09, 0x96,
	0x29, 0x8d, 0x90, 0xe7, 0xd0, 0x96, 0xc6, 0x1a, 0x8e, 0xe6, 0x0a, 0xcb, 0x87, 0x5a, 0xd6, 0xd7,
	0xd7, 0x2e, 0x72, 0x0c, 0x85, 0x39, 0x54, 0x78, 0xa7, 0x8a, 0x45, 0x02, 0xeb, 0xba, 0xc2, 0x3b,
	0xa5, 0x7b, 0x1b, 0xf1, 0x69, 0x96, 0xe0, 0x1d, 0x53, 0xf3, 0xa2, 0x79, 0x15, 0xcf, 0xd9, 0x7d,
	0x1d, 0x5a, 0xc5, 0x5e, 0xcb, 0x5e, 0xc6, 0xc8, 0x25, 0xb4, 0x06, 0xa8, 0x16, 0x43, 0xeb, 0x3e,
	0xb6, 0x26, 0xe5, 0x76, 0xba, 0xdd, 0x4d, 0x8b, 0xf4, 0xba, 0x46, 0x2e, 0xa0, 0x39, 0x40, 0x65,
	0x2b, 0x23, 0xc7, 0x6b, 0x01, 0xcb, 0xff, 0x95, 0x7b, 0xb0, 0x46, 0xa8, 0xf4, 0xe4, 0x1a, 0x3a,
	0x03, 0x54, 0x3d, 0x33, 0xdc, 0xcf, 0x7a, 0x58, 0x72, 0xb3, 0xe0, 0xc9, 0x1a, 0xa1, 0x32, 0x88,
	0x4a, 0x9a, 0x57, 0xb0, 0x3f, 0x40, 0x55, 0x0e, 0x85, 0xa1, 0x24, 0x87, 0x6b, 0x91, 0x95, 0xff,
	0xda, 0xf5, 0xfe, 0x8a, 0x2e, 0xed, 0x89, 0x29, 0xbe, 0x1d, 0xa2, 0xe4, 0xc9, 0x0c, 0x7b, 0x09,
	0xa3, 0x5a, 0xd3, 0x5e, 0x1a, 0xbf, 0xbc, 0x34, 0xfe, 0xa5, 0x12, 0x2c, 0x8d, 0xbf, 0xd1, 0x24,
	0x47, 0xf7, 0x60, 0x0d, 0x35, 0x83, 0x36, 0x60, 0x5f, 0x82, 0x1b, 0xf1, 0xa9, 0x7d, 0x35, 0x4b,
	0xa8, 0x9a, 0x70, 0x31, 0xf5, 0xf5, 0xbd, 0xd3, 0x8f, 0x5f, 0x9f, 0xc7, 0x4c, 0xfd, 0xc8, 0x47,
	0x7e, 0xc4, 0xa7, 0xc1, 0x12, 0x25, 0x88, 0xb9, 0xbd, 0x7a, 0xd9, 0x4d, 0x6c, 0x0f, 0x66, 0xf1,
	0x73, 0xe1, 0x38, 0x58, 0x39, 0xa5, 0xf7, 0x5b, 0x1d, 0xf3, 0xd7, 0xfa, 0x5f, 0x74, 0x4d, 0x03,
	0x91, 0x45, 0xa3, 0x1d, 0x93, 0xc9, 0x9b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xff, 0x8f, 0x19,
	0x77, 0xab, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountsApiClient is the client API for AccountsApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountsApiClient interface {
	GetBalances(ctx context.Context, in *BalancesRequest, opts ...grpc.CallOption) (AccountsApi_GetBalancesClient, error)
	GetScript(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*ScriptData, error)
	GetActiveLeases(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (AccountsApi_GetActiveLeasesClient, error)
	GetDataEntries(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (AccountsApi_GetDataEntriesClient, error)
	ResolveAlias(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*wrappers.BytesValue, error)
}

type accountsApiClient struct {
	cc *grpc.ClientConn
}

func NewAccountsApiClient(cc *grpc.ClientConn) AccountsApiClient {
	return &accountsApiClient{cc}
}

func (c *accountsApiClient) GetBalances(ctx context.Context, in *BalancesRequest, opts ...grpc.CallOption) (AccountsApi_GetBalancesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AccountsApi_serviceDesc.Streams[0], "/waves.node.grpc.AccountsApi/GetBalances", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountsApiGetBalancesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AccountsApi_GetBalancesClient interface {
	Recv() (*BalanceResponse, error)
	grpc.ClientStream
}

type accountsApiGetBalancesClient struct {
	grpc.ClientStream
}

func (x *accountsApiGetBalancesClient) Recv() (*BalanceResponse, error) {
	m := new(BalanceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *accountsApiClient) GetScript(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*ScriptData, error) {
	out := new(ScriptData)
	err := c.cc.Invoke(ctx, "/waves.node.grpc.AccountsApi/GetScript", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsApiClient) GetActiveLeases(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (AccountsApi_GetActiveLeasesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AccountsApi_serviceDesc.Streams[1], "/waves.node.grpc.AccountsApi/GetActiveLeases", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountsApiGetActiveLeasesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AccountsApi_GetActiveLeasesClient interface {
	Recv() (*TransactionResponse, error)
	grpc.ClientStream
}

type accountsApiGetActiveLeasesClient struct {
	grpc.ClientStream
}

func (x *accountsApiGetActiveLeasesClient) Recv() (*TransactionResponse, error) {
	m := new(TransactionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *accountsApiClient) GetDataEntries(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (AccountsApi_GetDataEntriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AccountsApi_serviceDesc.Streams[2], "/waves.node.grpc.AccountsApi/GetDataEntries", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountsApiGetDataEntriesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AccountsApi_GetDataEntriesClient interface {
	Recv() (*DataEntryResponse, error)
	grpc.ClientStream
}

type accountsApiGetDataEntriesClient struct {
	grpc.ClientStream
}

func (x *accountsApiGetDataEntriesClient) Recv() (*DataEntryResponse, error) {
	m := new(DataEntryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *accountsApiClient) ResolveAlias(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*wrappers.BytesValue, error) {
	out := new(wrappers.BytesValue)
	err := c.cc.Invoke(ctx, "/waves.node.grpc.AccountsApi/ResolveAlias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsApiServer is the server API for AccountsApi service.
type AccountsApiServer interface {
	GetBalances(*BalancesRequest, AccountsApi_GetBalancesServer) error
	GetScript(context.Context, *AccountRequest) (*ScriptData, error)
	GetActiveLeases(*AccountRequest, AccountsApi_GetActiveLeasesServer) error
	GetDataEntries(*DataRequest, AccountsApi_GetDataEntriesServer) error
	ResolveAlias(context.Context, *wrappers.StringValue) (*wrappers.BytesValue, error)
}

// UnimplementedAccountsApiServer can be embedded to have forward compatible implementations.
type UnimplementedAccountsApiServer struct {
}

func (*UnimplementedAccountsApiServer) GetBalances(req *BalancesRequest, srv AccountsApi_GetBalancesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBalances not implemented")
}
func (*UnimplementedAccountsApiServer) GetScript(ctx context.Context, req *AccountRequest) (*ScriptData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScript not implemented")
}
func (*UnimplementedAccountsApiServer) GetActiveLeases(req *AccountRequest, srv AccountsApi_GetActiveLeasesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetActiveLeases not implemented")
}
func (*UnimplementedAccountsApiServer) GetDataEntries(req *DataRequest, srv AccountsApi_GetDataEntriesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDataEntries not implemented")
}
func (*UnimplementedAccountsApiServer) ResolveAlias(ctx context.Context, req *wrappers.StringValue) (*wrappers.BytesValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveAlias not implemented")
}

func RegisterAccountsApiServer(s *grpc.Server, srv AccountsApiServer) {
	s.RegisterService(&_AccountsApi_serviceDesc, srv)
}

func _AccountsApi_GetBalances_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BalancesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountsApiServer).GetBalances(m, &accountsApiGetBalancesServer{stream})
}

type AccountsApi_GetBalancesServer interface {
	Send(*BalanceResponse) error
	grpc.ServerStream
}

type accountsApiGetBalancesServer struct {
	grpc.ServerStream
}

func (x *accountsApiGetBalancesServer) Send(m *BalanceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AccountsApi_GetScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsApiServer).GetScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/waves.node.grpc.AccountsApi/GetScript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsApiServer).GetScript(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsApi_GetActiveLeases_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AccountRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountsApiServer).GetActiveLeases(m, &accountsApiGetActiveLeasesServer{stream})
}

type AccountsApi_GetActiveLeasesServer interface {
	Send(*TransactionResponse) error
	grpc.ServerStream
}

type accountsApiGetActiveLeasesServer struct {
	grpc.ServerStream
}

func (x *accountsApiGetActiveLeasesServer) Send(m *TransactionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AccountsApi_GetDataEntries_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountsApiServer).GetDataEntries(m, &accountsApiGetDataEntriesServer{stream})
}

type AccountsApi_GetDataEntriesServer interface {
	Send(*DataEntryResponse) error
	grpc.ServerStream
}

type accountsApiGetDataEntriesServer struct {
	grpc.ServerStream
}

func (x *accountsApiGetDataEntriesServer) Send(m *DataEntryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AccountsApi_ResolveAlias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsApiServer).ResolveAlias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/waves.node.grpc.AccountsApi/ResolveAlias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsApiServer).ResolveAlias(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountsApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "waves.node.grpc.AccountsApi",
	HandlerType: (*AccountsApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScript",
			Handler:    _AccountsApi_GetScript_Handler,
		},
		{
			MethodName: "ResolveAlias",
			Handler:    _AccountsApi_ResolveAlias_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBalances",
			Handler:       _AccountsApi_GetBalances_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetActiveLeases",
			Handler:       _AccountsApi_GetActiveLeases_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetDataEntries",
			Handler:       _AccountsApi_GetDataEntries_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "waves/node/grpc/accounts_api.proto",
}