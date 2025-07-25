// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: podcast/v2/podcast.proto

package podcast

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PodcastServiceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PodcastServiceRequest) Reset() {
	*x = PodcastServiceRequest{}
	mi := &file_podcast_v2_podcast_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PodcastServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodcastServiceRequest) ProtoMessage() {}

func (x *PodcastServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_podcast_v2_podcast_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodcastServiceRequest.ProtoReflect.Descriptor instead.
func (*PodcastServiceRequest) Descriptor() ([]byte, []int) {
	return file_podcast_v2_podcast_proto_rawDescGZIP(), []int{0}
}

func (x *PodcastServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PodcastServiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PodcastServiceResponse) Reset() {
	*x = PodcastServiceResponse{}
	mi := &file_podcast_v2_podcast_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PodcastServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodcastServiceResponse) ProtoMessage() {}

func (x *PodcastServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_podcast_v2_podcast_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodcastServiceResponse.ProtoReflect.Descriptor instead.
func (*PodcastServiceResponse) Descriptor() ([]byte, []int) {
	return file_podcast_v2_podcast_proto_rawDescGZIP(), []int{1}
}

func (x *PodcastServiceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_podcast_v2_podcast_proto protoreflect.FileDescriptor

const file_podcast_v2_podcast_proto_rawDesc = "" +
	"\n" +
	"\x18podcast/v2/podcast.proto\x12\n" +
	"podcast.v1\"+\n" +
	"\x15PodcastServiceRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"2\n" +
	"\x16PodcastServiceResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage2e\n" +
	"\x0ePodcastService\x12S\n" +
	"\bSayHello\x12!.podcast.v1.PodcastServiceRequest\x1a\".podcast.v1.PodcastServiceResponse\"\x00B8Z6github.com/reaksa-maii/one_digital_grpc_getway/podcastb\x06proto3"

var (
	file_podcast_v2_podcast_proto_rawDescOnce sync.Once
	file_podcast_v2_podcast_proto_rawDescData []byte
)

func file_podcast_v2_podcast_proto_rawDescGZIP() []byte {
	file_podcast_v2_podcast_proto_rawDescOnce.Do(func() {
		file_podcast_v2_podcast_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_podcast_v2_podcast_proto_rawDesc), len(file_podcast_v2_podcast_proto_rawDesc)))
	})
	return file_podcast_v2_podcast_proto_rawDescData
}

var file_podcast_v2_podcast_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_podcast_v2_podcast_proto_goTypes = []any{
	(*PodcastServiceRequest)(nil),  // 0: podcast.v1.PodcastServiceRequest
	(*PodcastServiceResponse)(nil), // 1: podcast.v1.PodcastServiceResponse
}
var file_podcast_v2_podcast_proto_depIdxs = []int32{
	0, // 0: podcast.v1.PodcastService.SayHello:input_type -> podcast.v1.PodcastServiceRequest
	1, // 1: podcast.v1.PodcastService.SayHello:output_type -> podcast.v1.PodcastServiceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_podcast_v2_podcast_proto_init() }
func file_podcast_v2_podcast_proto_init() {
	if File_podcast_v2_podcast_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_podcast_v2_podcast_proto_rawDesc), len(file_podcast_v2_podcast_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_podcast_v2_podcast_proto_goTypes,
		DependencyIndexes: file_podcast_v2_podcast_proto_depIdxs,
		MessageInfos:      file_podcast_v2_podcast_proto_msgTypes,
	}.Build()
	File_podcast_v2_podcast_proto = out.File
	file_podcast_v2_podcast_proto_goTypes = nil
	file_podcast_v2_podcast_proto_depIdxs = nil
}
