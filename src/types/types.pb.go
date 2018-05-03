// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/types.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	types/types.proto

It has these top-level messages:
	BluePrintCreateAccountTx
	BluePrintStateTx
	BluePrintAppState
	StateQueryParams
	StateQueryResult
*/
package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type BluePrintCreateAccountTx struct {
	Version int32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Owner   string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BluePrintCreateAccountTx) Reset()                    { *m = BluePrintCreateAccountTx{} }
func (m *BluePrintCreateAccountTx) String() string            { return proto.CompactTextString(m) }
func (*BluePrintCreateAccountTx) ProtoMessage()               {}
func (*BluePrintCreateAccountTx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *BluePrintCreateAccountTx) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BluePrintCreateAccountTx) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *BluePrintCreateAccountTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BluePrintStateTx struct {
	Version int32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Owner   string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BluePrintStateTx) Reset()                    { *m = BluePrintStateTx{} }
func (m *BluePrintStateTx) String() string            { return proto.CompactTextString(m) }
func (*BluePrintStateTx) ProtoMessage()               {}
func (*BluePrintStateTx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *BluePrintStateTx) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BluePrintStateTx) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *BluePrintStateTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BluePrintAppState struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Blob    []byte `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (m *BluePrintAppState) Reset()                    { *m = BluePrintAppState{} }
func (m *BluePrintAppState) String() string            { return proto.CompactTextString(m) }
func (*BluePrintAppState) ProtoMessage()               {}
func (*BluePrintAppState) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

func (m *BluePrintAppState) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *BluePrintAppState) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

type StateQueryParams struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *StateQueryParams) Reset()                    { *m = StateQueryParams{} }
func (m *StateQueryParams) String() string            { return proto.CompactTextString(m) }
func (*StateQueryParams) ProtoMessage()               {}
func (*StateQueryParams) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{3} }

func (m *StateQueryParams) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type StateQueryResult struct {
	State []byte `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (m *StateQueryResult) Reset()                    { *m = StateQueryResult{} }
func (m *StateQueryResult) String() string            { return proto.CompactTextString(m) }
func (*StateQueryResult) ProtoMessage()               {}
func (*StateQueryResult) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{4} }

func (m *StateQueryResult) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func init() {
	proto.RegisterType((*BluePrintCreateAccountTx)(nil), "BluePrintCreateAccountTx")
	proto.RegisterType((*BluePrintStateTx)(nil), "BluePrintStateTx")
	proto.RegisterType((*BluePrintAppState)(nil), "BluePrintAppState")
	proto.RegisterType((*StateQueryParams)(nil), "StateQueryParams")
	proto.RegisterType((*StateQueryResult)(nil), "StateQueryResult")
}

func init() { proto.RegisterFile("types/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd0, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0x19, 0x28, 0x88, 0x53, 0x86, 0xd6, 0x62, 0xf0, 0x18, 0x65, 0xf2, 0x04, 0x03,
	0x4f, 0x10, 0x78, 0x81, 0x62, 0x10, 0x23, 0x92, 0xd3, 0xdc, 0x50, 0x29, 0xd8, 0xd6, 0xf9, 0x0c,
	0xf4, 0xed, 0x91, 0xaf, 0x24, 0xc0, 0xde, 0xc5, 0xba, 0xdf, 0x3a, 0x7f, 0x3e, 0x1d, 0x6c, 0xf8,
	0x90, 0x30, 0xdf, 0xc9, 0x79, 0x9b, 0x28, 0x72, 0xec, 0xde, 0xc0, 0x3c, 0x4c, 0x05, 0xb7, 0xb4,
	0x0f, 0xfc, 0x48, 0xe8, 0x19, 0xfb, 0xdd, 0x2e, 0x96, 0xc0, 0x2f, 0x5f, 0xda, 0xc0, 0xd5, 0x07,
	0x52, 0xde, 0xc7, 0x60, 0x54, 0xab, 0xec, 0xca, 0xcd, 0x51, 0xdf, 0xc0, 0x2a, 0x7e, 0x06, 0x24,
	0x73, 0xd6, 0x2a, 0x7b, 0xed, 0x8e, 0x41, 0x6b, 0xb8, 0x18, 0x3d, 0x7b, 0x73, 0xde, 0x2a, 0xdb,
	0x38, 0xa9, 0xbb, 0x57, 0x58, 0x2f, 0xfe, 0x33, 0x7b, 0xc6, 0x13, 0xb9, 0x3d, 0x6c, 0x16, 0xb7,
	0x4f, 0x49, 0xe8, 0x0a, 0xfb, 0x71, 0x24, 0xcc, 0x59, 0xe0, 0xc6, 0xcd, 0xb1, 0x12, 0xc3, 0x14,
	0x07, 0x71, 0x1b, 0x27, 0x75, 0x67, 0x61, 0x2d, 0xcf, 0x9e, 0x0a, 0xd2, 0x61, 0xeb, 0xc9, 0xbf,
	0xe7, 0xdf, 0x01, 0xd4, 0x9f, 0x01, 0xfe, 0x77, 0x3a, 0xcc, 0x65, 0xe2, 0xda, 0x99, 0xeb, 0xdd,
	0xcf, 0x4f, 0xc7, 0x30, 0x5c, 0xca, 0x56, 0xef, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x38,
	0x82, 0xa3, 0x6a, 0x01, 0x00, 0x00,
}