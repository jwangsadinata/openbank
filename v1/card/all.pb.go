// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/card/all.proto

package card

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	types "github.com/openbank/openbank/v1/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Card holds all details about a card.
type Card struct {
	// ID is a unique identifier of a card.
	ID string `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	// Account is the ID of the account associated with the card.
	Account string `protobuf:"bytes,2,opt,name=Account,json=account,proto3" json:"Account,omitempty"`
	// OwnerName is the name of the card owner.
	OwnerName string `protobuf:"bytes,3,opt,name=OwnerName,json=owner_name,proto3" json:"OwnerName,omitempty"`
	// ContactNumber is the contact number of the card owner.
	ContactNumber string `protobuf:"bytes,4,opt,name=ContactNumber,json=contact_number,proto3" json:"ContactNumber,omitempty"`
	// FirstName is the first name of card owner.
	FirstName string `protobuf:"bytes,5,opt,name=FirstName,json=first_name,proto3" json:"FirstName,omitempty"`
	// LastName is the last name of the card owner.
	LastName string `protobuf:"bytes,6,opt,name=LastName,json=last_name,proto3" json:"LastName,omitempty"`
	// Expiry is an expiry date of the card.
	Expiry *timestamp.Timestamp `protobuf:"bytes,7,opt,name=Expiry,json=expiry,proto3" json:"Expiry,omitempty"`
	// Status is the status of the card.
	Status types.CardStatus `protobuf:"varint,8,opt,name=Status,json=status,proto3,enum=types.CardStatus" json:"Status,omitempty"`
	// AccessStatus is the access status of the card.
	AccessStatus         types.CardAccessStatus `protobuf:"varint,9,opt,name=AccessStatus,json=access_status,proto3,enum=types.CardAccessStatus" json:"AccessStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d7396a66589e9e, []int{0}
}

func (m *Card) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Card.Unmarshal(m, b)
}
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Card.Marshal(b, m, deterministic)
}
func (m *Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Card.Merge(m, src)
}
func (m *Card) XXX_Size() int {
	return xxx_messageInfo_Card.Size(m)
}
func (m *Card) XXX_DiscardUnknown() {
	xxx_messageInfo_Card.DiscardUnknown(m)
}

var xxx_messageInfo_Card proto.InternalMessageInfo

func (m *Card) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Card) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Card) GetOwnerName() string {
	if m != nil {
		return m.OwnerName
	}
	return ""
}

func (m *Card) GetContactNumber() string {
	if m != nil {
		return m.ContactNumber
	}
	return ""
}

func (m *Card) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Card) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Card) GetExpiry() *timestamp.Timestamp {
	if m != nil {
		return m.Expiry
	}
	return nil
}

func (m *Card) GetStatus() types.CardStatus {
	if m != nil {
		return m.Status
	}
	return types.CardStatus_UnknownCardStatus
}

func (m *Card) GetAccessStatus() types.CardAccessStatus {
	if m != nil {
		return m.AccessStatus
	}
	return types.CardAccessStatus_UnknownCardAccessStatus
}

// GetCardRequest is the request envelope to retrieve card information.
type GetCardRequest struct {
	// CardToken is the identifier to get the card information.
	CardToken            string   `protobuf:"bytes,1,opt,name=CardToken,json=card_token,proto3" json:"CardToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCardRequest) Reset()         { *m = GetCardRequest{} }
func (m *GetCardRequest) String() string { return proto.CompactTextString(m) }
func (*GetCardRequest) ProtoMessage()    {}
func (*GetCardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d7396a66589e9e, []int{1}
}

func (m *GetCardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCardRequest.Unmarshal(m, b)
}
func (m *GetCardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCardRequest.Marshal(b, m, deterministic)
}
func (m *GetCardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCardRequest.Merge(m, src)
}
func (m *GetCardRequest) XXX_Size() int {
	return xxx_messageInfo_GetCardRequest.Size(m)
}
func (m *GetCardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCardRequest proto.InternalMessageInfo

func (m *GetCardRequest) GetCardToken() string {
	if m != nil {
		return m.CardToken
	}
	return ""
}

// UpdateCardStatusRequest is the request envelope to update card status.
type UpdateCardStatusRequest struct {
	// CardToken is the identifier to get the card information.
	CardToken string `protobuf:"bytes,1,opt,name=CardToken,json=card_token,proto3" json:"CardToken,omitempty"`
	// Status is the new card status.
	Status               types.CardStatus `protobuf:"varint,2,opt,name=Status,json=status,proto3,enum=types.CardStatus" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpdateCardStatusRequest) Reset()         { *m = UpdateCardStatusRequest{} }
func (m *UpdateCardStatusRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCardStatusRequest) ProtoMessage()    {}
func (*UpdateCardStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d7396a66589e9e, []int{2}
}

func (m *UpdateCardStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCardStatusRequest.Unmarshal(m, b)
}
func (m *UpdateCardStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCardStatusRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCardStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCardStatusRequest.Merge(m, src)
}
func (m *UpdateCardStatusRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCardStatusRequest.Size(m)
}
func (m *UpdateCardStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCardStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCardStatusRequest proto.InternalMessageInfo

func (m *UpdateCardStatusRequest) GetCardToken() string {
	if m != nil {
		return m.CardToken
	}
	return ""
}

func (m *UpdateCardStatusRequest) GetStatus() types.CardStatus {
	if m != nil {
		return m.Status
	}
	return types.CardStatus_UnknownCardStatus
}

// UpdateCardAccessStatusRequest is the request envelope to update card access status.
type UpdateCardAccessStatusRequest struct {
	// CardToken is the identifier to get the card information.
	CardToken string `protobuf:"bytes,1,opt,name=CardToken,json=card_token,proto3" json:"CardToken,omitempty"`
	// AccessStatus is the new card access status.
	AccessStatus         types.CardAccessStatus `protobuf:"varint,2,opt,name=AccessStatus,json=access_status,proto3,enum=types.CardAccessStatus" json:"AccessStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UpdateCardAccessStatusRequest) Reset()         { *m = UpdateCardAccessStatusRequest{} }
func (m *UpdateCardAccessStatusRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCardAccessStatusRequest) ProtoMessage()    {}
func (*UpdateCardAccessStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d7396a66589e9e, []int{3}
}

func (m *UpdateCardAccessStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCardAccessStatusRequest.Unmarshal(m, b)
}
func (m *UpdateCardAccessStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCardAccessStatusRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCardAccessStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCardAccessStatusRequest.Merge(m, src)
}
func (m *UpdateCardAccessStatusRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCardAccessStatusRequest.Size(m)
}
func (m *UpdateCardAccessStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCardAccessStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCardAccessStatusRequest proto.InternalMessageInfo

func (m *UpdateCardAccessStatusRequest) GetCardToken() string {
	if m != nil {
		return m.CardToken
	}
	return ""
}

func (m *UpdateCardAccessStatusRequest) GetAccessStatus() types.CardAccessStatus {
	if m != nil {
		return m.AccessStatus
	}
	return types.CardAccessStatus_UnknownCardAccessStatus
}

// Result is result of the update operation.
type Result struct {
	// Success is a boolean indicating the success of the operation.
	Success bool `protobuf:"varint,1,opt,name=Success,json=success,proto3" json:"Success,omitempty"`
	// ErrorCode is the code of the error.
	ErrorCode string `protobuf:"bytes,2,opt,name=ErrorCode,json=error_code,proto3" json:"ErrorCode,omitempty"`
	// ErrorMessage is the message of the error.
	ErrorMessage         string   `protobuf:"bytes,3,opt,name=ErrorMessage,json=error_message,proto3" json:"ErrorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d7396a66589e9e, []int{4}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *Result) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*Card)(nil), "card.Card")
	proto.RegisterType((*GetCardRequest)(nil), "card.GetCardRequest")
	proto.RegisterType((*UpdateCardStatusRequest)(nil), "card.UpdateCardStatusRequest")
	proto.RegisterType((*UpdateCardAccessStatusRequest)(nil), "card.UpdateCardAccessStatusRequest")
	proto.RegisterType((*Result)(nil), "card.Result")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/card/all.proto", fileDescriptor_41d7396a66589e9e)
}

var fileDescriptor_41d7396a66589e9e = []byte{
	// 1251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xde, 0x5d, 0x27, 0x8e, 0x33, 0xfd, 0x41, 0x3a, 0xad, 0x52, 0xb3, 0x6a, 0x61, 0x94, 0x22,
	0x35, 0x54, 0xed, 0xda, 0x71, 0xdb, 0x03, 0x11, 0x14, 0xb9, 0x26, 0xfd, 0x11, 0x95, 0x36, 0x72,
	0x0b, 0x42, 0xbd, 0x44, 0xe3, 0xdd, 0x67, 0x7b, 0xe9, 0x7a, 0x66, 0x99, 0x99, 0x4d, 0x1a, 0x10,
	0x12, 0x02, 0xa9, 0xca, 0xb1, 0x0a, 0x57, 0x10, 0x7f, 0x05, 0x12, 0x47, 0x24, 0xc4, 0x89, 0x03,
	0x48, 0x5c, 0x38, 0x71, 0x82, 0x03, 0x12, 0x37, 0x0e, 0x70, 0x44, 0x33, 0xb3, 0xfe, 0x15, 0x3b,
	0xaa, 0x52, 0x38, 0x25, 0x3b, 0xef, 0x7b, 0xdf, 0x7c, 0xf3, 0xbd, 0xf7, 0xc6, 0x83, 0x82, 0x4e,
	0xac, 0xba, 0x59, 0x2b, 0x08, 0x79, 0xaf, 0xc2, 0x53, 0x60, 0x2d, 0xca, 0x1e, 0x0d, 0xff, 0xd9,
	0x5a, 0xa9, 0x84, 0x54, 0x44, 0x15, 0x9a, 0x24, 0x41, 0x2a, 0xb8, 0xe2, 0x78, 0x46, 0x7f, 0xfb,
	0x2f, 0x77, 0x38, 0xef, 0x24, 0x50, 0x31, 0x6b, 0xad, 0xac, 0x5d, 0x51, 0x71, 0x0f, 0xa4, 0xa2,
	0xbd, 0xd4, 0xc2, 0xfc, 0x33, 0x39, 0x80, 0xa6, 0x71, 0x85, 0x32, 0xc6, 0x15, 0x55, 0x31, 0x67,
	0x32, 0x8f, 0x5e, 0x34, 0x7f, 0xc2, 0x4b, 0x1d, 0x60, 0x97, 0xe4, 0x36, 0xed, 0x74, 0x40, 0x54,
	0x78, 0x6a, 0x10, 0x53, 0xd0, 0x95, 0x67, 0x49, 0x54, 0x3b, 0x29, 0xc8, 0xa1, 0xc6, 0xa5, 0xef,
	0x0a, 0x68, 0xa6, 0x41, 0x45, 0x84, 0x7d, 0xe4, 0xdd, 0x7e, 0xab, 0xec, 0x12, 0x77, 0x79, 0xfe,
	0x3a, 0x2a, 0x39, 0x65, 0x67, 0xd9, 0xa9, 0x3a, 0x1b, 0x4e, 0xd3, 0x8b, 0x23, 0xfc, 0x0a, 0x9a,
	0xab, 0x87, 0x21, 0xcf, 0x98, 0x2a, 0x7b, 0x13, 0x80, 0x39, 0x6a, 0x43, 0xf8, 0x55, 0x34, 0x7f,
	0x6f, 0x9b, 0x81, 0xb8, 0x4b, 0x7b, 0x50, 0x2e, 0x4c, 0xe0, 0x10, 0xd7, 0xc1, 0x4d, 0x46, 0x7b,
	0x80, 0x57, 0xd0, 0xb1, 0x06, 0x67, 0x8a, 0x86, 0xea, 0x6e, 0xd6, 0x6b, 0x81, 0x28, 0xcf, 0x4c,
	0xc0, 0x8f, 0x87, 0x16, 0xb0, 0xc9, 0x0c, 0x42, 0xb3, 0xdf, 0x88, 0x85, 0x54, 0x86, 0x7d, 0x76,
	0x92, 0xbd, 0xad, 0x83, 0x96, 0xfd, 0x3c, 0x2a, 0xdd, 0xa1, 0x39, 0xb2, 0x38, 0x81, 0x9c, 0x4f,
	0x68, 0x1f, 0x78, 0x0d, 0x15, 0xd7, 0x1e, 0xa7, 0xb1, 0xd8, 0x29, 0xcf, 0x11, 0x77, 0xf9, 0x48,
	0xcd, 0x0f, 0x6c, 0x29, 0x82, 0x7e, 0xad, 0x82, 0x07, 0xfd, 0x5a, 0x8d, 0x51, 0x14, 0xc1, 0x64,
	0xe1, 0xab, 0xa8, 0x78, 0x5f, 0x51, 0x95, 0xc9, 0x72, 0x89, 0xb8, 0xcb, 0xc7, 0x6b, 0x27, 0x02,
	0x63, 0x6f, 0xa0, 0x0d, 0xb5, 0x81, 0xf1, 0x34, 0x69, 0xd6, 0xf0, 0x2d, 0x74, 0xb4, 0x1e, 0x86,
	0x20, 0x65, 0x9e, 0x3c, 0x6f, 0x92, 0x4f, 0x8f, 0x24, 0x8f, 0x86, 0xc7, 0x28, 0x8e, 0x51, 0x13,
	0xd9, 0xb4, 0x4c, 0xab, 0xc5, 0x92, 0xb3, 0xe0, 0x94, 0x9d, 0xa5, 0x06, 0x3a, 0x7e, 0x13, 0x94,
	0xce, 0x6c, 0xc2, 0x07, 0x19, 0x48, 0x53, 0x0c, 0xfd, 0xf9, 0x80, 0x3f, 0x02, 0x36, 0xa5, 0xaa,
	0x48, 0xb7, 0xe6, 0xa6, 0xd2, 0xd1, 0x01, 0xc9, 0x67, 0x2e, 0x3a, 0xfd, 0x4e, 0x1a, 0x51, 0x05,
	0x43, 0xfd, 0x87, 0xa7, 0x1b, 0x31, 0xc5, 0x3b, 0x84, 0x29, 0x03, 0x15, 0x5f, 0xb8, 0xe8, 0xec,
	0x50, 0xc5, 0xa8, 0x11, 0xcf, 0xa1, 0x65, 0xbf, 0xd3, 0xde, 0x7f, 0x76, 0x7a, 0xcf, 0x45, 0xc5,
	0x26, 0xc8, 0x2c, 0x51, 0x7a, 0x2a, 0xee, 0x67, 0x06, 0x64, 0x54, 0x94, 0xc6, 0xa7, 0x42, 0xda,
	0x90, 0x56, 0xbb, 0x26, 0x04, 0x17, 0x0d, 0x1e, 0xc1, 0x94, 0xe9, 0x41, 0xa0, 0x83, 0x9b, 0x21,
	0x8f, 0x00, 0x57, 0xd0, 0x51, 0x03, 0x7d, 0x1b, 0xa4, 0xa4, 0x9d, 0x69, 0x33, 0x74, 0xcc, 0xa2,
	0x7b, 0x16, 0xd0, 0x17, 0x55, 0xfb, 0x7a, 0x0e, 0x1d, 0x31, 0xf6, 0x82, 0xd8, 0x8a, 0x43, 0xc0,
	0x3f, 0x7a, 0x68, 0x2e, 0xef, 0x07, 0x7c, 0x2a, 0xd0, 0x7e, 0x04, 0xe3, 0xed, 0xe1, 0x23, 0xbb,
	0xaa, 0x97, 0x96, 0xbe, 0xf4, 0xf6, 0xea, 0x7f, 0xbb, 0xf9, 0x35, 0xf0, 0x62, 0x13, 0x94, 0x88,
	0x61, 0x0b, 0x88, 0x06, 0x90, 0x98, 0xb5, 0xb9, 0xe8, 0x99, 0x2b, 0xc6, 0xbf, 0xd5, 0x0f, 0x49,
	0x42, 0x93, 0x84, 0x44, 0x54, 0x51, 0xd2, 0x16, 0xbc, 0x47, 0xa8, 0xc1, 0x5e, 0x24, 0x12, 0x12,
	0x08, 0x15, 0x44, 0xa4, 0xb5, 0x43, 0x54, 0xd7, 0x32, 0xd8, 0x42, 0x90, 0x1d, 0x9e, 0x11, 0x99,
	0xa5, 0x69, 0x12, 0x43, 0x14, 0xac, 0x37, 0x50, 0xa1, 0x56, 0xad, 0xe2, 0xd7, 0xd1, 0x4b, 0xb9,
	0x1e, 0x02, 0x8f, 0x21, 0xcc, 0x74, 0x6a, 0x6e, 0x5b, 0x3b, 0x4b, 0x92, 0x9d, 0x00, 0xfb, 0xa8,
	0xec, 0x2f, 0x9e, 0xab, 0x44, 0xd0, 0x8e, 0x59, 0x6c, 0xef, 0x3c, 0x4d, 0xaa, 0x95, 0xae, 0xaf,
	0xa0, 0xc2, 0x95, 0xea, 0x15, 0x7c, 0x01, 0x2d, 0x37, 0x41, 0x65, 0x82, 0x41, 0x44, 0xb6, 0xbb,
	0xc0, 0xcc, 0xce, 0x02, 0x24, 0xcf, 0x44, 0x08, 0x24, 0x96, 0x84, 0x71, 0x45, 0xda, 0x3c, 0x63,
	0x51, 0xd0, 0xc2, 0x68, 0x01, 0x15, 0xef, 0xd5, 0x33, 0xd5, 0xad, 0xe1, 0x22, 0x9a, 0x11, 0x40,
	0xa3, 0x4f, 0x7f, 0xfe, 0xed, 0x73, 0x6f, 0x11, 0x9f, 0x1a, 0x5c, 0xdf, 0x1f, 0x0d, 0x7a, 0xeb,
	0xe3, 0x5d, 0xcf, 0x79, 0xea, 0x19, 0xeb, 0xf1, 0x13, 0x0f, 0x2d, 0xec, 0x9f, 0x0d, 0x7c, 0xd6,
	0x9a, 0x78, 0xc0, 0xcc, 0xf8, 0x47, 0x6d, 0xd8, 0x76, 0xcb, 0xd2, 0xb7, 0xee, 0x5e, 0xfd, 0xab,
	0xbe, 0xcb, 0xd8, 0xe6, 0x58, 0x8f, 0x6d, 0x87, 0xf9, 0x27, 0xf3, 0xb5, 0xc6, 0x70, 0x2d, 0x58,
	0x3f, 0x6f, 0x9d, 0x22, 0xcf, 0x72, 0xea, 0xff, 0x74, 0xe3, 0xc4, 0xd2, 0x0b, 0x03, 0x37, 0xac,
	0x8e, 0x11, 0x23, 0xbe, 0xf7, 0xd0, 0xe2, 0xf4, 0xf1, 0xc4, 0xe7, 0xf6, 0xdb, 0x31, 0x65, 0x78,
	0xf7, 0x99, 0xf2, 0x97, 0xbb, 0x57, 0xff, 0x61, 0xd0, 0x7a, 0xa3, 0xa6, 0xd8, 0xe9, 0xeb, 0x7b,
	0xe3, 0x1f, 0x18, 0x7a, 0x9e, 0x93, 0xaf, 0xaf, 0x59, 0x57, 0xaf, 0x3d, 0xb3, 0xff, 0xce, 0x20,
	0xdf, 0x2f, 0x4f, 0xf6, 0x9f, 0x55, 0x7f, 0xa0, 0x81, 0xe5, 0xa5, 0xc5, 0xe1, 0x6b, 0x60, 0xf4,
	0x16, 0x19, 0xfa, 0xe8, 0x17, 0x76, 0x3d, 0xe7, 0xfa, 0x6e, 0x71, 0xaf, 0xfe, 0xfb, 0x2c, 0x2a,
	0xd4, 0x82, 0x2a, 0xbe, 0x83, 0x4a, 0xa6, 0xf2, 0xf5, 0x8d, 0xdb, 0xf8, 0xb5, 0x0d, 0xc1, 0xb7,
	0xe2, 0x08, 0x24, 0x09, 0x05, 0xe8, 0x73, 0x53, 0x16, 0x11, 0x4d, 0x4f, 0x78, 0x0a, 0xc2, 0xfe,
	0xe0, 0x13, 0xce, 0x06, 0x43, 0x36, 0x38, 0x67, 0x50, 0x9b, 0x5d, 0x09, 0xaa, 0x41, 0xf5, 0x82,
	0xeb, 0xd5, 0x16, 0xa8, 0x1e, 0xb7, 0xd0, 0xa0, 0x2b, 0xef, 0x4b, 0xce, 0x56, 0x27, 0x56, 0x9a,
	0x37, 0x50, 0xe1, 0x6a, 0xb5, 0x8a, 0xdf, 0x44, 0x6f, 0x8c, 0x7b, 0x47, 0x19, 0xc9, 0x18, 0x3c,
	0x4e, 0xed, 0x38, 0x9b, 0x4b, 0x87, 0xf0, 0x30, 0xcc, 0x04, 0x44, 0xfd, 0x7d, 0x25, 0x88, 0x2d,
	0x10, 0x44, 0xc6, 0x11, 0x04, 0xcd, 0x4d, 0x5d, 0x83, 0x2a, 0x7e, 0x0f, 0xbd, 0x3b, 0xad, 0x06,
	0xd6, 0xe2, 0x16, 0x8f, 0x76, 0x74, 0x1d, 0x7a, 0x34, 0xb1, 0x77, 0x8a, 0xa6, 0xe6, 0x82, 0x44,
	0x1c, 0x6c, 0x71, 0x7a, 0x54, 0x85, 0x5d, 0x93, 0x32, 0xd8, 0x39, 0xcf, 0x0d, 0x9a, 0x77, 0xf4,
	0x06, 0x2b, 0x78, 0x0d, 0x35, 0x0e, 0xde, 0x60, 0x40, 0x64, 0x5e, 0x0e, 0x31, 0x93, 0x26, 0x9a,
	0x49, 0x10, 0xe7, 0x8d, 0x91, 0x11, 0x30, 0x15, 0xd3, 0x44, 0x06, 0xcd, 0x0d, 0xcd, 0x76, 0x19,
	0xdf, 0x46, 0x37, 0x27, 0xd9, 0x34, 0x7e, 0x48, 0xd5, 0xa5, 0x5b, 0x40, 0x52, 0x10, 0xbd, 0x58,
	0xca, 0x58, 0x9f, 0x9c, 0xf7, 0xdb, 0x6f, 0xb4, 0xbd, 0x82, 0xe6, 0xe1, 0x9b, 0xf0, 0xe1, 0x9f,
	0x2e, 0xfa, 0xc3, 0x1d, 0xf4, 0xcf, 0xaf, 0x6e, 0xa9, 0x80, 0x9f, 0xb8, 0xf5, 0x9c, 0x9c, 0x13,
	0x25, 0x28, 0x93, 0x34, 0xb4, 0xb5, 0xee, 0x53, 0x48, 0xcd, 0x21, 0x40, 0x2a, 0x11, 0x1b, 0x7f,
	0xb4, 0x9c, 0x4c, 0x75, 0xf5, 0xc1, 0x42, 0xaa, 0x17, 0xb4, 0x7a, 0x19, 0x90, 0x07, 0x5d, 0x18,
	0x0d, 0x68, 0xe5, 0xa9, 0xe0, 0x86, 0xba, 0xcd, 0x93, 0x84, 0x6f, 0x5b, 0xfd, 0x7a, 0x6b, 0x2e,
	0xe2, 0x0f, 0x2d, 0x42, 0xff, 0x42, 0x91, 0x76, 0xc2, 0xb7, 0x83, 0xe5, 0x99, 0x5a, 0x49, 0x77,
	0xb0, 0xa6, 0x58, 0x9d, 0x37, 0xcf, 0x46, 0x7d, 0x21, 0x5e, 0x5f, 0x45, 0xbe, 0x6d, 0x73, 0x8c,
	0x6f, 0x0a, 0xca, 0x94, 0xb4, 0x4d, 0x69, 0x1d, 0x41, 0x67, 0xd0, 0xec, 0xb6, 0x88, 0x15, 0xe0,
	0x93, 0x79, 0xd0, 0x7c, 0xe5, 0xd1, 0x5b, 0xee, 0x86, 0xf3, 0xd0, 0x3c, 0x88, 0x3f, 0x71, 0x9d,
	0x5d, 0xd7, 0x79, 0xea, 0x3a, 0xdf, 0xb8, 0xce, 0x2f, 0xae, 0xf3, 0x8f, 0xeb, 0xfc, 0xe4, 0x39,
	0xad, 0xa2, 0x79, 0x72, 0x5d, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x67, 0x7c, 0xbe, 0x0c, 0x64,
	0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CardServiceClient is the client API for CardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CardServiceClient interface {
	// GetCard retrieves the detail of a card information, selected by its token.
	GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*Card, error)
	// UpdateCardStatus update the card status.
	UpdateCardStatus(ctx context.Context, in *UpdateCardStatusRequest, opts ...grpc.CallOption) (*Result, error)
	// UpdateCardAccessStatus update the card access status.
	UpdateCardAccessStatus(ctx context.Context, in *UpdateCardAccessStatusRequest, opts ...grpc.CallOption) (*Result, error)
}

type cardServiceClient struct {
	cc *grpc.ClientConn
}

func NewCardServiceClient(cc *grpc.ClientConn) CardServiceClient {
	return &cardServiceClient{cc}
}

func (c *cardServiceClient) GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*Card, error) {
	out := new(Card)
	err := c.cc.Invoke(ctx, "/card.CardService/GetCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) UpdateCardStatus(ctx context.Context, in *UpdateCardStatusRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/card.CardService/UpdateCardStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) UpdateCardAccessStatus(ctx context.Context, in *UpdateCardAccessStatusRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/card.CardService/UpdateCardAccessStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardServiceServer is the server API for CardService service.
type CardServiceServer interface {
	// GetCard retrieves the detail of a card information, selected by its token.
	GetCard(context.Context, *GetCardRequest) (*Card, error)
	// UpdateCardStatus update the card status.
	UpdateCardStatus(context.Context, *UpdateCardStatusRequest) (*Result, error)
	// UpdateCardAccessStatus update the card access status.
	UpdateCardAccessStatus(context.Context, *UpdateCardAccessStatusRequest) (*Result, error)
}

// UnimplementedCardServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCardServiceServer struct {
}

func (*UnimplementedCardServiceServer) GetCard(ctx context.Context, req *GetCardRequest) (*Card, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCard not implemented")
}
func (*UnimplementedCardServiceServer) UpdateCardStatus(ctx context.Context, req *UpdateCardStatusRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCardStatus not implemented")
}
func (*UnimplementedCardServiceServer) UpdateCardAccessStatus(ctx context.Context, req *UpdateCardAccessStatusRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCardAccessStatus not implemented")
}

func RegisterCardServiceServer(s *grpc.Server, srv CardServiceServer) {
	s.RegisterService(&_CardService_serviceDesc, srv)
}

func _CardService_GetCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).GetCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/card.CardService/GetCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).GetCard(ctx, req.(*GetCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_UpdateCardStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCardStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).UpdateCardStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/card.CardService/UpdateCardStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).UpdateCardStatus(ctx, req.(*UpdateCardStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_UpdateCardAccessStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCardAccessStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).UpdateCardAccessStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/card.CardService/UpdateCardAccessStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).UpdateCardAccessStatus(ctx, req.(*UpdateCardAccessStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CardService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "card.CardService",
	HandlerType: (*CardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCard",
			Handler:    _CardService_GetCard_Handler,
		},
		{
			MethodName: "UpdateCardStatus",
			Handler:    _CardService_UpdateCardStatus_Handler,
		},
		{
			MethodName: "UpdateCardAccessStatus",
			Handler:    _CardService_UpdateCardAccessStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/card/all.proto",
}
