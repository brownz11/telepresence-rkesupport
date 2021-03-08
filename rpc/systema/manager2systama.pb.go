// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: rpc/systema/manager2systama.proto

// The "systema" package describes the server implemented by the
// SystemA cloud service, which is called to by the in-cluster
// Telepsresence manager and the in-cluster Ambassador Telepresence
// agent sidecars.

package systema

import (
	common "github.com/telepresenceio/telepresence/rpc/v2/common"
	manager "github.com/telepresenceio/telepresence/rpc/v2/manager"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_systema_manager2systama_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_systema_manager2systama_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_rpc_systema_manager2systama_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type CreateDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterceptId   string                 `protobuf:"bytes,1,opt,name=intercept_id,json=interceptId,proto3" json:"intercept_id,omitempty"`
	DisplayBanner bool                   `protobuf:"varint,2,opt,name=display_banner,json=displayBanner,proto3" json:"display_banner,omitempty"`
	InterceptSpec *manager.InterceptSpec `protobuf:"bytes,3,opt,name=intercept_spec,json=interceptSpec,proto3" json:"intercept_spec,omitempty"`
	Host          string                 `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *CreateDomainRequest) Reset() {
	*x = CreateDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_systema_manager2systama_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDomainRequest) ProtoMessage() {}

func (x *CreateDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_systema_manager2systama_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDomainRequest.ProtoReflect.Descriptor instead.
func (*CreateDomainRequest) Descriptor() ([]byte, []int) {
	return file_rpc_systema_manager2systama_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDomainRequest) GetInterceptId() string {
	if x != nil {
		return x.InterceptId
	}
	return ""
}

func (x *CreateDomainRequest) GetDisplayBanner() bool {
	if x != nil {
		return x.DisplayBanner
	}
	return false
}

func (x *CreateDomainRequest) GetInterceptSpec() *manager.InterceptSpec {
	if x != nil {
		return x.InterceptSpec
	}
	return nil
}

func (x *CreateDomainRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type CreateDomainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *CreateDomainResponse) Reset() {
	*x = CreateDomainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_systema_manager2systama_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDomainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDomainResponse) ProtoMessage() {}

func (x *CreateDomainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_systema_manager2systama_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDomainResponse.ProtoReflect.Descriptor instead.
func (*CreateDomainResponse) Descriptor() ([]byte, []int) {
	return file_rpc_systema_manager2systama_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDomainResponse) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type RemoveDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *RemoveDomainRequest) Reset() {
	*x = RemoveDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_systema_manager2systama_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDomainRequest) ProtoMessage() {}

func (x *RemoveDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_systema_manager2systama_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDomainRequest.ProtoReflect.Descriptor instead.
func (*RemoveDomainRequest) Descriptor() ([]byte, []int) {
	return file_rpc_systema_manager2systama_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveDomainRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type PreferredAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageName string `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
}

func (x *PreferredAgentResponse) Reset() {
	*x = PreferredAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_systema_manager2systama_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreferredAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreferredAgentResponse) ProtoMessage() {}

func (x *PreferredAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_systema_manager2systama_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreferredAgentResponse.ProtoReflect.Descriptor instead.
func (*PreferredAgentResponse) Descriptor() ([]byte, []int) {
	return file_rpc_systema_manager2systama_proto_rawDescGZIP(), []int{4}
}

func (x *PreferredAgentResponse) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

var File_rpc_systema_manager2systama_proto protoreflect.FileDescriptor

var file_rpc_systema_manager2systama_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x32, 0x73, 0x79, 0x73, 0x74, 0x61, 0x6d, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x05, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xbf,
	0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63,
	0x65, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x12, 0x4a, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0d, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x22, 0x2e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x22, 0x2d, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22,
	0x37, 0x0a, 0x16, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xa9, 0x02, 0x0a, 0x0b, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x41, 0x43, 0x52, 0x55, 0x44, 0x12, 0x65, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x29, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x51, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x29, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x60, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x2c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x61, 0x0a, 0x0c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x12, 0x51, 0x0a, 0x11, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x61,
	0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x1b, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x28, 0x01, 0x30, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x77, 0x69, 0x72, 0x65, 0x2f, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x32, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x76, 0x32, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rpc_systema_manager2systama_proto_rawDescOnce sync.Once
	file_rpc_systema_manager2systama_proto_rawDescData = file_rpc_systema_manager2systama_proto_rawDesc
)

func file_rpc_systema_manager2systama_proto_rawDescGZIP() []byte {
	file_rpc_systema_manager2systama_proto_rawDescOnce.Do(func() {
		file_rpc_systema_manager2systama_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_systema_manager2systama_proto_rawDescData)
	})
	return file_rpc_systema_manager2systama_proto_rawDescData
}

var file_rpc_systema_manager2systama_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_systema_manager2systama_proto_goTypes = []interface{}{
	(*Chunk)(nil),                  // 0: telepresence.systema.Chunk
	(*CreateDomainRequest)(nil),    // 1: telepresence.systema.CreateDomainRequest
	(*CreateDomainResponse)(nil),   // 2: telepresence.systema.CreateDomainResponse
	(*RemoveDomainRequest)(nil),    // 3: telepresence.systema.RemoveDomainRequest
	(*PreferredAgentResponse)(nil), // 4: telepresence.systema.PreferredAgentResponse
	(*manager.InterceptSpec)(nil),  // 5: telepresence.manager.InterceptSpec
	(*common.VersionInfo)(nil),     // 6: telepresence.common.VersionInfo
	(*empty.Empty)(nil),            // 7: google.protobuf.Empty
}
var file_rpc_systema_manager2systama_proto_depIdxs = []int32{
	5, // 0: telepresence.systema.CreateDomainRequest.intercept_spec:type_name -> telepresence.manager.InterceptSpec
	1, // 1: telepresence.systema.SystemACRUD.CreateDomain:input_type -> telepresence.systema.CreateDomainRequest
	3, // 2: telepresence.systema.SystemACRUD.RemoveDomain:input_type -> telepresence.systema.RemoveDomainRequest
	6, // 3: telepresence.systema.SystemACRUD.PreferredAgent:input_type -> telepresence.common.VersionInfo
	0, // 4: telepresence.systema.SystemAProxy.ReverseConnection:input_type -> telepresence.systema.Chunk
	2, // 5: telepresence.systema.SystemACRUD.CreateDomain:output_type -> telepresence.systema.CreateDomainResponse
	7, // 6: telepresence.systema.SystemACRUD.RemoveDomain:output_type -> google.protobuf.Empty
	4, // 7: telepresence.systema.SystemACRUD.PreferredAgent:output_type -> telepresence.systema.PreferredAgentResponse
	0, // 8: telepresence.systema.SystemAProxy.ReverseConnection:output_type -> telepresence.systema.Chunk
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_systema_manager2systama_proto_init() }
func file_rpc_systema_manager2systama_proto_init() {
	if File_rpc_systema_manager2systama_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_systema_manager2systama_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_rpc_systema_manager2systama_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDomainRequest); i {
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
		file_rpc_systema_manager2systama_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDomainResponse); i {
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
		file_rpc_systema_manager2systama_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveDomainRequest); i {
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
		file_rpc_systema_manager2systama_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreferredAgentResponse); i {
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
			RawDescriptor: file_rpc_systema_manager2systama_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_rpc_systema_manager2systama_proto_goTypes,
		DependencyIndexes: file_rpc_systema_manager2systama_proto_depIdxs,
		MessageInfos:      file_rpc_systema_manager2systama_proto_msgTypes,
	}.Build()
	File_rpc_systema_manager2systama_proto = out.File
	file_rpc_systema_manager2systama_proto_rawDesc = nil
	file_rpc_systema_manager2systama_proto_goTypes = nil
	file_rpc_systema_manager2systama_proto_depIdxs = nil
}
