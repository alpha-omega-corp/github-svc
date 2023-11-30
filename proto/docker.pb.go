// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: proto/docker.proto

package proto

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

type CreateContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *CreateContainerRequest) Reset() {
	*x = CreateContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_docker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContainerRequest) ProtoMessage() {}

func (x *CreateContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_docker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContainerRequest.ProtoReflect.Descriptor instead.
func (*CreateContainerRequest) Descriptor() ([]byte, []int) {
	return file_proto_docker_proto_rawDescGZIP(), []int{0}
}

func (x *CreateContainerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateContainerRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type CreateContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error     string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Container string `protobuf:"bytes,3,opt,name=container,proto3" json:"container,omitempty"`
}

func (x *CreateContainerResponse) Reset() {
	*x = CreateContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_docker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContainerResponse) ProtoMessage() {}

func (x *CreateContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_docker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContainerResponse.ProtoReflect.Descriptor instead.
func (*CreateContainerResponse) Descriptor() ([]byte, []int) {
	return file_proto_docker_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContainerResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateContainerResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateContainerResponse) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

type GetContainersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetContainersRequest) Reset() {
	*x = GetContainersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_docker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContainersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContainersRequest) ProtoMessage() {}

func (x *GetContainersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_docker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContainersRequest.ProtoReflect.Descriptor instead.
func (*GetContainersRequest) Descriptor() ([]byte, []int) {
	return file_proto_docker_proto_rawDescGZIP(), []int{2}
}

type GetContainersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Containers []*Container `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
}

func (x *GetContainersResponse) Reset() {
	*x = GetContainersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_docker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContainersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContainersResponse) ProtoMessage() {}

func (x *GetContainersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_docker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContainersResponse.ProtoReflect.Descriptor instead.
func (*GetContainersResponse) Descriptor() ([]byte, []int) {
	return file_proto_docker_proto_rawDescGZIP(), []int{3}
}

func (x *GetContainersResponse) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Image   string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Status  string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Command string   `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"`
	Created int64    `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	State   string   `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	Names   []string `protobuf:"bytes,7,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_docker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_proto_docker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_proto_docker_proto_rawDescGZIP(), []int{4}
}

func (x *Container) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Container) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Container) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Container) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Container) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Container) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Container) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

var File_proto_docker_proto protoreflect.FileDescriptor

var file_proto_docker_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x22, 0x65, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x4a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0xb5, 0x01, 0x0a, 0x0d, 0x44, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x12, 0x1c, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_docker_proto_rawDescOnce sync.Once
	file_proto_docker_proto_rawDescData = file_proto_docker_proto_rawDesc
)

func file_proto_docker_proto_rawDescGZIP() []byte {
	file_proto_docker_proto_rawDescOnce.Do(func() {
		file_proto_docker_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_docker_proto_rawDescData)
	})
	return file_proto_docker_proto_rawDescData
}

var file_proto_docker_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_docker_proto_goTypes = []interface{}{
	(*CreateContainerRequest)(nil),  // 0: docker.CreateContainerRequest
	(*CreateContainerResponse)(nil), // 1: docker.CreateContainerResponse
	(*GetContainersRequest)(nil),    // 2: docker.GetContainersRequest
	(*GetContainersResponse)(nil),   // 3: docker.GetContainersResponse
	(*Container)(nil),               // 4: docker.Container
}
var file_proto_docker_proto_depIdxs = []int32{
	4, // 0: docker.GetContainersResponse.containers:type_name -> docker.Container
	0, // 1: docker.DockerService.CreateContainer:input_type -> docker.CreateContainerRequest
	2, // 2: docker.DockerService.GetContainers:input_type -> docker.GetContainersRequest
	1, // 3: docker.DockerService.CreateContainer:output_type -> docker.CreateContainerResponse
	3, // 4: docker.DockerService.GetContainers:output_type -> docker.GetContainersResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_docker_proto_init() }
func file_proto_docker_proto_init() {
	if File_proto_docker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_docker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContainerRequest); i {
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
		file_proto_docker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContainerResponse); i {
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
		file_proto_docker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContainersRequest); i {
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
		file_proto_docker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContainersResponse); i {
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
		file_proto_docker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container); i {
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
			RawDescriptor: file_proto_docker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_docker_proto_goTypes,
		DependencyIndexes: file_proto_docker_proto_depIdxs,
		MessageInfos:      file_proto_docker_proto_msgTypes,
	}.Build()
	File_proto_docker_proto = out.File
	file_proto_docker_proto_rawDesc = nil
	file_proto_docker_proto_goTypes = nil
	file_proto_docker_proto_depIdxs = nil
}
