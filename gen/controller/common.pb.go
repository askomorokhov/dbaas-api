// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: controller/common.proto

package controllerv1beta1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// RunningOperation respresents a long-running operation.
type RunningOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Progress from 0.0 to 1.0; can decrease compared to the previous value.
	Progress float32 `protobuf:"fixed32,1,opt,name=progress,proto3" json:"progress,omitempty"`
	// Text describing the current operation progress step.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RunningOperation) Reset() {
	*x = RunningOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunningOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunningOperation) ProtoMessage() {}

func (x *RunningOperation) ProtoReflect() protoreflect.Message {
	mi := &file_controller_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunningOperation.ProtoReflect.Descriptor instead.
func (*RunningOperation) Descriptor() ([]byte, []int) {
	return file_controller_common_proto_rawDescGZIP(), []int{0}
}

func (x *RunningOperation) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *RunningOperation) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// KubeAuth represents Kubernetes / kubectl authentication and authorization information.
type KubeAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubeconfig file content.
	Kubeconfig string `protobuf:"bytes,1,opt,name=kubeconfig,proto3" json:"kubeconfig,omitempty"`
}

func (x *KubeAuth) Reset() {
	*x = KubeAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeAuth) ProtoMessage() {}

func (x *KubeAuth) ProtoReflect() protoreflect.Message {
	mi := &file_controller_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeAuth.ProtoReflect.Descriptor instead.
func (*KubeAuth) Descriptor() ([]byte, []int) {
	return file_controller_common_proto_rawDescGZIP(), []int{1}
}

func (x *KubeAuth) GetKubeconfig() string {
	if x != nil {
		return x.Kubeconfig
	}
	return ""
}

// ComputeResources represents container computer resources requests or limits.
type ComputeResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CPUs in milliCPUs; 1000m = 1 vCPU.
	CpuM int32 `protobuf:"varint,1,opt,name=cpu_m,json=cpuM,proto3" json:"cpu_m,omitempty"`
	// Memory in bytes.
	MemoryBytes int64 `protobuf:"varint,2,opt,name=memory_bytes,json=memoryBytes,proto3" json:"memory_bytes,omitempty"`
}

func (x *ComputeResources) Reset() {
	*x = ComputeResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeResources) ProtoMessage() {}

func (x *ComputeResources) ProtoReflect() protoreflect.Message {
	mi := &file_controller_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeResources.ProtoReflect.Descriptor instead.
func (*ComputeResources) Descriptor() ([]byte, []int) {
	return file_controller_common_proto_rawDescGZIP(), []int{2}
}

func (x *ComputeResources) GetCpuM() int32 {
	if x != nil {
		return x.CpuM
	}
	return 0
}

func (x *ComputeResources) GetMemoryBytes() int64 {
	if x != nil {
		return x.MemoryBytes
	}
	return 0
}

var File_controller_common_proto protoreflect.FileDescriptor

var file_controller_common_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x70, 0x65, 0x72, 0x63, 0x6f,
	0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61,
	0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x10,
	0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x32, 0x0a, 0x08, 0x4b, 0x75, 0x62, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x26, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x0a,
	0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x4a, 0x0a, 0x10, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x13,
	0x0a, 0x05, 0x63, 0x70, 0x75, 0x5f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x70, 0x75, 0x4d, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x1e, 0x5a, 0x1c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_common_proto_rawDescOnce sync.Once
	file_controller_common_proto_rawDescData = file_controller_common_proto_rawDesc
)

func file_controller_common_proto_rawDescGZIP() []byte {
	file_controller_common_proto_rawDescOnce.Do(func() {
		file_controller_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_common_proto_rawDescData)
	})
	return file_controller_common_proto_rawDescData
}

var file_controller_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_controller_common_proto_goTypes = []interface{}{
	(*RunningOperation)(nil), // 0: percona.platform.dbaas.controller.v1beta1.RunningOperation
	(*KubeAuth)(nil),         // 1: percona.platform.dbaas.controller.v1beta1.KubeAuth
	(*ComputeResources)(nil), // 2: percona.platform.dbaas.controller.v1beta1.ComputeResources
}
var file_controller_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_controller_common_proto_init() }
func file_controller_common_proto_init() {
	if File_controller_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunningOperation); i {
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
		file_controller_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeAuth); i {
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
		file_controller_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeResources); i {
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
			RawDescriptor: file_controller_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_common_proto_goTypes,
		DependencyIndexes: file_controller_common_proto_depIdxs,
		MessageInfos:      file_controller_common_proto_msgTypes,
	}.Build()
	File_controller_common_proto = out.File
	file_controller_common_proto_rawDesc = nil
	file_controller_common_proto_goTypes = nil
	file_controller_common_proto_depIdxs = nil
}
