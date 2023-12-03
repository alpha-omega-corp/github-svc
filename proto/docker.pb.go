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

type CreatePackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dockerfile []byte `protobuf:"bytes,1,opt,name=dockerfile,proto3" json:"dockerfile,omitempty"`
	Workdir    string `protobuf:"bytes,2,opt,name=workdir,proto3" json:"workdir,omitempty"`
	Tag        string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *CreatePackageRequest) Reset() {
	*x = CreatePackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_docker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePackageRequest) ProtoMessage() {}

func (x *CreatePackageRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreatePackageRequest.ProtoReflect.Descriptor instead.
func (*CreatePackageRequest) Descriptor() ([]byte, []int) {
	return file_proto_docker_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePackageRequest) GetDockerfile() []byte {
	if x != nil {
		return x.Dockerfile
	}
	return nil
}

func (x *CreatePackageRequest) GetWorkdir() string {
	if x != nil {
		return x.Workdir
	}
	return ""
}

func (x *CreatePackageRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type CreatePackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error     string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Container string `protobuf:"bytes,3,opt,name=container,proto3" json:"container,omitempty"`
}

func (x *CreatePackageResponse) Reset() {
	*x = CreatePackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_docker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePackageResponse) ProtoMessage() {}

func (x *CreatePackageResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreatePackageResponse.ProtoReflect.Descriptor instead.
func (*CreatePackageResponse) Descriptor() ([]byte, []int) {
	return file_proto_docker_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePackageResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreatePackageResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreatePackageResponse) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

type GetPackagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPackagesRequest) Reset() {
	*x = GetPackagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_docker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackagesRequest) ProtoMessage() {}

func (x *GetPackagesRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPackagesRequest.ProtoReflect.Descriptor instead.
func (*GetPackagesRequest) Descriptor() ([]byte, []int) {
	return file_proto_docker_proto_rawDescGZIP(), []int{2}
}

type GetPackagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packages []*Package `protobuf:"bytes,1,rep,name=packages,proto3" json:"packages,omitempty"`
}

func (x *GetPackagesResponse) Reset() {
	*x = GetPackagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_docker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackagesResponse) ProtoMessage() {}

func (x *GetPackagesResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPackagesResponse.ProtoReflect.Descriptor instead.
func (*GetPackagesResponse) Descriptor() ([]byte, []int) {
	return file_proto_docker_proto_rawDescGZIP(), []int{3}
}

func (x *GetPackagesResponse) GetPackages() []*Package {
	if x != nil {
		return x.Packages
	}
	return nil
}

type GetContainersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetContainersRequest) Reset() {
	*x = GetContainersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_docker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContainersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContainersRequest) ProtoMessage() {}

func (x *GetContainersRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetContainersRequest.ProtoReflect.Descriptor instead.
func (*GetContainersRequest) Descriptor() ([]byte, []int) {
	return file_proto_docker_proto_rawDescGZIP(), []int{4}
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
		mi := &file_proto_docker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContainersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContainersResponse) ProtoMessage() {}

func (x *GetContainersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_docker_proto_msgTypes[5]
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
	return file_proto_docker_proto_rawDescGZIP(), []int{5}
}

func (x *GetContainersResponse) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

type GetContainerLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerId string `protobuf:"bytes,1,opt,name=containerId,proto3" json:"containerId,omitempty"`
}

func (x *GetContainerLogsRequest) Reset() {
	*x = GetContainerLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_docker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContainerLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContainerLogsRequest) ProtoMessage() {}

func (x *GetContainerLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_docker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContainerLogsRequest.ProtoReflect.Descriptor instead.
func (*GetContainerLogsRequest) Descriptor() ([]byte, []int) {
	return file_proto_docker_proto_rawDescGZIP(), []int{6}
}

func (x *GetContainerLogsRequest) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

type GetContainerLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs string `protobuf:"bytes,1,opt,name=logs,proto3" json:"logs,omitempty"`
}

func (x *GetContainerLogsResponse) Reset() {
	*x = GetContainerLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_docker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContainerLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContainerLogsResponse) ProtoMessage() {}

func (x *GetContainerLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_docker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContainerLogsResponse.ProtoReflect.Descriptor instead.
func (*GetContainerLogsResponse) Descriptor() ([]byte, []int) {
	return file_proto_docker_proto_rawDescGZIP(), []int{7}
}

func (x *GetContainerLogsResponse) GetLogs() string {
	if x != nil {
		return x.Logs
	}
	return ""
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
		mi := &file_proto_docker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_proto_docker_proto_msgTypes[8]
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
	return file_proto_docker_proto_rawDescGZIP(), []int{8}
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

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Tag        string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Dockerfile []byte `protobuf:"bytes,4,opt,name=dockerfile,proto3" json:"dockerfile,omitempty"`
	Makefile   []byte `protobuf:"bytes,5,opt,name=makefile,proto3" json:"makefile,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_docker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_proto_docker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_proto_docker_proto_rawDescGZIP(), []int{9}
}

func (x *Package) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Package) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetDockerfile() []byte {
	if x != nil {
		return x.Dockerfile
	}
	return nil
}

func (x *Package) GetMakefile() []byte {
	if x != nil {
		return x.Makefile
	}
	return nil
}

var File_proto_docker_proto protoreflect.FileDescriptor

var file_proto_docker_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x22, 0x62, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x64, 0x69, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x64, 0x69, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67,
	0x22, 0x63, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x22, 0x3b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x2e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73,
	0x22, 0xa9, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x7b, 0x0a, 0x07,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x6b, 0x65, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x6d, 0x61, 0x6b, 0x65, 0x66, 0x69, 0x6c, 0x65, 0x32, 0xd2, 0x02, 0x0a, 0x0d, 0x44, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1f, 0x2e, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e,
	0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2d, 0x6f, 0x6d, 0x65, 0x67, 0x61, 0x2d, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x64, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x2d, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
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

var file_proto_docker_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_docker_proto_goTypes = []interface{}{
	(*CreatePackageRequest)(nil),     // 0: docker.CreatePackageRequest
	(*CreatePackageResponse)(nil),    // 1: docker.CreatePackageResponse
	(*GetPackagesRequest)(nil),       // 2: docker.GetPackagesRequest
	(*GetPackagesResponse)(nil),      // 3: docker.GetPackagesResponse
	(*GetContainersRequest)(nil),     // 4: docker.GetContainersRequest
	(*GetContainersResponse)(nil),    // 5: docker.GetContainersResponse
	(*GetContainerLogsRequest)(nil),  // 6: docker.GetContainerLogsRequest
	(*GetContainerLogsResponse)(nil), // 7: docker.GetContainerLogsResponse
	(*Container)(nil),                // 8: docker.Container
	(*Package)(nil),                  // 9: docker.Package
}
var file_proto_docker_proto_depIdxs = []int32{
	9, // 0: docker.GetPackagesResponse.packages:type_name -> docker.Package
	8, // 1: docker.GetContainersResponse.containers:type_name -> docker.Container
	0, // 2: docker.DockerService.CreatePackage:input_type -> docker.CreatePackageRequest
	2, // 3: docker.DockerService.GetPackages:input_type -> docker.GetPackagesRequest
	4, // 4: docker.DockerService.GetContainers:input_type -> docker.GetContainersRequest
	6, // 5: docker.DockerService.GetContainerLogs:input_type -> docker.GetContainerLogsRequest
	1, // 6: docker.DockerService.CreatePackage:output_type -> docker.CreatePackageResponse
	3, // 7: docker.DockerService.GetPackages:output_type -> docker.GetPackagesResponse
	5, // 8: docker.DockerService.GetContainers:output_type -> docker.GetContainersResponse
	7, // 9: docker.DockerService.GetContainerLogs:output_type -> docker.GetContainerLogsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_docker_proto_init() }
func file_proto_docker_proto_init() {
	if File_proto_docker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_docker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePackageRequest); i {
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
			switch v := v.(*CreatePackageResponse); i {
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
			switch v := v.(*GetPackagesRequest); i {
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
			switch v := v.(*GetPackagesResponse); i {
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
		file_proto_docker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_docker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContainerLogsRequest); i {
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
		file_proto_docker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContainerLogsResponse); i {
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
		file_proto_docker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_docker_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
			NumMessages:   10,
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
