// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: like.proto

package pb

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

// 点赞或取消点赞
type InsertUpvoteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Action bool   `protobuf:"varint,3,opt,name=Action,proto3" json:"Action,omitempty"`
}

func (x *InsertUpvoteReq) Reset() {
	*x = InsertUpvoteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertUpvoteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertUpvoteReq) ProtoMessage() {}

func (x *InsertUpvoteReq) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertUpvoteReq.ProtoReflect.Descriptor instead.
func (*InsertUpvoteReq) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{0}
}

func (x *InsertUpvoteReq) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *InsertUpvoteReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *InsertUpvoteReq) GetAction() bool {
	if x != nil {
		return x.Action
	}
	return false
}

type InsertUpvoteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *InsertUpvoteResp) Reset() {
	*x = InsertUpvoteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertUpvoteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertUpvoteResp) ProtoMessage() {}

func (x *InsertUpvoteResp) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertUpvoteResp.ProtoReflect.Descriptor instead.
func (*InsertUpvoteResp) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{1}
}

func (x *InsertUpvoteResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 查询用户是否已经点赞动态
type GetStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetStatusReq) Reset() {
	*x = GetStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusReq) ProtoMessage() {}

func (x *GetStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusReq.ProtoReflect.Descriptor instead.
func (*GetStatusReq) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{2}
}

func (x *GetStatusReq) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *GetStatusReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *GetStatusResp) Reset() {
	*x = GetStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusResp) ProtoMessage() {}

func (x *GetStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusResp.ProtoReflect.Descriptor instead.
func (*GetStatusResp) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{3}
}

func (x *GetStatusResp) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

// 获取点赞数量
type GetNumReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
}

func (x *GetNumReq) Reset() {
	*x = GetNumReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNumReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNumReq) ProtoMessage() {}

func (x *GetNumReq) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNumReq.ProtoReflect.Descriptor instead.
func (*GetNumReq) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{4}
}

func (x *GetNumReq) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type GetNumResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (x *GetNumResp) Reset() {
	*x = GetNumResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNumResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNumResp) ProtoMessage() {}

func (x *GetNumResp) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNumResp.ProtoReflect.Descriptor instead.
func (*GetNumResp) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{5}
}

func (x *GetNumResp) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_like_proto protoreflect.FileDescriptor

var file_like_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x59, 0x0a, 0x0f, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x70, 0x76, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x0a,
	0x10, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x3e, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x50,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x23, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x32, 0xa8, 0x01, 0x0a, 0x04, 0x4c, 0x69,
	0x6b, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x0f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x3d, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x12,
	0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x70, 0x76,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_like_proto_rawDescOnce sync.Once
	file_like_proto_rawDescData = file_like_proto_rawDesc
)

func file_like_proto_rawDescGZIP() []byte {
	file_like_proto_rawDescOnce.Do(func() {
		file_like_proto_rawDescData = protoimpl.X.CompressGZIP(file_like_proto_rawDescData)
	})
	return file_like_proto_rawDescData
}

var file_like_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_like_proto_goTypes = []interface{}{
	(*InsertUpvoteReq)(nil),  // 0: user.InsertUpvoteReq
	(*InsertUpvoteResp)(nil), // 1: user.InsertUpvoteResp
	(*GetStatusReq)(nil),     // 2: user.GetStatusReq
	(*GetStatusResp)(nil),    // 3: user.GetStatusResp
	(*GetNumReq)(nil),        // 4: user.GetNumReq
	(*GetNumResp)(nil),       // 5: user.GetNumResp
}
var file_like_proto_depIdxs = []int32{
	4, // 0: user.Like.GetNum:input_type -> user.GetNumReq
	0, // 1: user.Like.InsertUpvote:input_type -> user.InsertUpvoteReq
	2, // 2: user.Like.GetStatus:input_type -> user.GetStatusReq
	5, // 3: user.Like.GetNum:output_type -> user.GetNumResp
	1, // 4: user.Like.InsertUpvote:output_type -> user.InsertUpvoteResp
	3, // 5: user.Like.GetStatus:output_type -> user.GetStatusResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_like_proto_init() }
func file_like_proto_init() {
	if File_like_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_like_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertUpvoteReq); i {
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
		file_like_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertUpvoteResp); i {
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
		file_like_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusReq); i {
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
		file_like_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusResp); i {
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
		file_like_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNumReq); i {
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
		file_like_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNumResp); i {
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
			RawDescriptor: file_like_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_like_proto_goTypes,
		DependencyIndexes: file_like_proto_depIdxs,
		MessageInfos:      file_like_proto_msgTypes,
	}.Build()
	File_like_proto = out.File
	file_like_proto_rawDesc = nil
	file_like_proto_goTypes = nil
	file_like_proto_depIdxs = nil
}
