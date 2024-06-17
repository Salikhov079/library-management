// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: borrowers.proto

package genprotos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BorrowerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookId     string `protobuf:"bytes,3,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	BorrowDate string `protobuf:"bytes,4,opt,name=borrow_date,json=borrowDate,proto3" json:"borrow_date,omitempty"`
	ReturnDate string `protobuf:"bytes,5,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
}

func (x *BorrowerReq) Reset() {
	*x = BorrowerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrowers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowerReq) ProtoMessage() {}

func (x *BorrowerReq) ProtoReflect() protoreflect.Message {
	mi := &file_borrowers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowerReq.ProtoReflect.Descriptor instead.
func (*BorrowerReq) Descriptor() ([]byte, []int) {
	return file_borrowers_proto_rawDescGZIP(), []int{0}
}

func (x *BorrowerReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BorrowerReq) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *BorrowerReq) GetBorrowDate() string {
	if x != nil {
		return x.BorrowDate
	}
	return ""
}

func (x *BorrowerReq) GetReturnDate() string {
	if x != nil {
		return x.ReturnDate
	}
	return ""
}

type BorrowerRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User       *UserRes `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Book       *BookRes `protobuf:"bytes,3,opt,name=book,proto3" json:"book,omitempty"`
	BorrowDate string   `protobuf:"bytes,4,opt,name=borrow_date,json=borrowDate,proto3" json:"borrow_date,omitempty"`
	ReturnDate string   `protobuf:"bytes,5,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
}

func (x *BorrowerRes) Reset() {
	*x = BorrowerRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrowers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowerRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowerRes) ProtoMessage() {}

func (x *BorrowerRes) ProtoReflect() protoreflect.Message {
	mi := &file_borrowers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowerRes.ProtoReflect.Descriptor instead.
func (*BorrowerRes) Descriptor() ([]byte, []int) {
	return file_borrowers_proto_rawDescGZIP(), []int{1}
}

func (x *BorrowerRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BorrowerRes) GetUser() *UserRes {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *BorrowerRes) GetBook() *BookRes {
	if x != nil {
		return x.Book
	}
	return nil
}

func (x *BorrowerRes) GetBorrowDate() string {
	if x != nil {
		return x.BorrowDate
	}
	return ""
}

func (x *BorrowerRes) GetReturnDate() string {
	if x != nil {
		return x.ReturnDate
	}
	return ""
}

type AllBorrowers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Borrowers []*BorrowerRes `protobuf:"bytes,1,rep,name=borrowers,proto3" json:"borrowers,omitempty"`
}

func (x *AllBorrowers) Reset() {
	*x = AllBorrowers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrowers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllBorrowers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllBorrowers) ProtoMessage() {}

func (x *AllBorrowers) ProtoReflect() protoreflect.Message {
	mi := &file_borrowers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllBorrowers.ProtoReflect.Descriptor instead.
func (*AllBorrowers) Descriptor() ([]byte, []int) {
	return file_borrowers_proto_rawDescGZIP(), []int{2}
}

func (x *AllBorrowers) GetBorrowers() []*BorrowerRes {
	if x != nil {
		return x.Borrowers
	}
	return nil
}

type FilterBorrower struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BorrowDate string `protobuf:"bytes,2,opt,name=borrow_date,json=borrowDate,proto3" json:"borrow_date,omitempty"`
	ReturnDate string `protobuf:"bytes,3,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	DeletedAt  string `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *FilterBorrower) Reset() {
	*x = FilterBorrower{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrowers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterBorrower) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterBorrower) ProtoMessage() {}

func (x *FilterBorrower) ProtoReflect() protoreflect.Message {
	mi := &file_borrowers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterBorrower.ProtoReflect.Descriptor instead.
func (*FilterBorrower) Descriptor() ([]byte, []int) {
	return file_borrowers_proto_rawDescGZIP(), []int{3}
}

func (x *FilterBorrower) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FilterBorrower) GetBorrowDate() string {
	if x != nil {
		return x.BorrowDate
	}
	return ""
}

func (x *FilterBorrower) GetReturnDate() string {
	if x != nil {
		return x.ReturnDate
	}
	return ""
}

func (x *FilterBorrower) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

var File_borrowers_proto protoreflect.FileDescriptor

var file_borrowers_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x04,
	0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x52, 0x04, 0x62, 0x6f, 0x6f,
	0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x22, 0x41, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x09, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x52, 0x09, 0x62, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x32, 0xe5, 0x02, 0x0a, 0x0f, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72,
	0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x73,
	0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x64, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x6c, 0x6c,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_borrowers_proto_rawDescOnce sync.Once
	file_borrowers_proto_rawDescData = file_borrowers_proto_rawDesc
)

func file_borrowers_proto_rawDescGZIP() []byte {
	file_borrowers_proto_rawDescOnce.Do(func() {
		file_borrowers_proto_rawDescData = protoimpl.X.CompressGZIP(file_borrowers_proto_rawDescData)
	})
	return file_borrowers_proto_rawDescData
}

var file_borrowers_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_borrowers_proto_goTypes = []interface{}{
	(*BorrowerReq)(nil),    // 0: protos.BorrowerReq
	(*BorrowerRes)(nil),    // 1: protos.BorrowerRes
	(*AllBorrowers)(nil),   // 2: protos.AllBorrowers
	(*FilterBorrower)(nil), // 3: protos.FilterBorrower
	(*UserRes)(nil),        // 4: protos.UserRes
	(*BookRes)(nil),        // 5: protos.BookRes
	(*ById)(nil),           // 6: protos.ById
	(*Void)(nil),           // 7: protos.Void
}
var file_borrowers_proto_depIdxs = []int32{
	4, // 0: protos.BorrowerRes.user:type_name -> protos.UserRes
	5, // 1: protos.BorrowerRes.book:type_name -> protos.BookRes
	1, // 2: protos.AllBorrowers.borrowers:type_name -> protos.BorrowerRes
	0, // 3: protos.BorrowerService.CreateBorrower:input_type -> protos.BorrowerReq
	1, // 4: protos.BorrowerService.UpdateBorrower:input_type -> protos.BorrowerRes
	6, // 5: protos.BorrowerService.DeleteBorrower:input_type -> protos.ById
	6, // 6: protos.BorrowerService.GetByIdBorrower:input_type -> protos.ById
	3, // 7: protos.BorrowerService.GetAllBorrowers:input_type -> protos.FilterBorrower
	7, // 8: protos.BorrowerService.GetAllIdBorrowers:input_type -> protos.Void
	7, // 9: protos.BorrowerService.CreateBorrower:output_type -> protos.Void
	7, // 10: protos.BorrowerService.UpdateBorrower:output_type -> protos.Void
	7, // 11: protos.BorrowerService.DeleteBorrower:output_type -> protos.Void
	1, // 12: protos.BorrowerService.GetByIdBorrower:output_type -> protos.BorrowerRes
	2, // 13: protos.BorrowerService.GetAllBorrowers:output_type -> protos.AllBorrowers
	2, // 14: protos.BorrowerService.GetAllIdBorrowers:output_type -> protos.AllBorrowers
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_borrowers_proto_init() }
func file_borrowers_proto_init() {
	if File_borrowers_proto != nil {
		return
	}
	file_users_proto_init()
	file_books_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_borrowers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowerReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_borrowers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowerRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_borrowers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllBorrowers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_borrowers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterBorrower); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_borrowers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_borrowers_proto_goTypes,
		DependencyIndexes: file_borrowers_proto_depIdxs,
		MessageInfos:      file_borrowers_proto_msgTypes,
	}.Build()
	File_borrowers_proto = out.File
	file_borrowers_proto_rawDesc = nil
	file_borrowers_proto_goTypes = nil
	file_borrowers_proto_depIdxs = nil
}
