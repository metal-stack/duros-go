// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: lightbits/api/duros/v1/statisticsapi.proto

package v1

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

///////////////////////////////////////
//   API structures (external facing)
///////////////////////////////////////
type ClusterStatisticsApi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Installed Physical Storage
	//
	// All installed SSDs capacities over all servers in cluster, given in bytes.
	InstalledPhysicalStorage uint64 `protobuf:"varint,1,opt,name=installedPhysicalStorage,proto3" json:"installedPhysicalStorage,omitempty"`
	// Total Attached Physical Storage
	//
	// Sum of all managed and healthy SSDs capacities, given in bytes.
	ManagedPhysicalStorage uint64 `protobuf:"varint,2,opt,name=managedPhysicalStorage,proto3" json:"managedPhysicalStorage,omitempty"`
	// Effective Physical Storage
	//
	// Effective Physical storage excluding overhead of OVP and Parity, given in bytes.
	EffectivePhysicalStorage uint64 `protobuf:"varint,3,opt,name=effectivePhysicalStorage,proto3" json:"effectivePhysicalStorage,omitempty"`
	// Logical Storage
	//
	// Sum of capacities of all allocated volumes, given in bytes.
	LogicalStorage uint64 `protobuf:"varint,4,opt,name=logicalStorage,proto3" json:"logicalStorage,omitempty"`
	// Logical Used Storage
	//
	// Logical storage space used by all volumes (n of LBAs x 4096), given in bytes.
	LogicalUsedStorage uint64 `protobuf:"varint,5,opt,name=logicalUsedStorage,proto3" json:"logicalUsedStorage,omitempty"`
	// Physical Used Storage Excluding Parity
	//
	// Physical storage space occupied by all volumes (data only), given in bytes.
	PhysicalUsedStorage uint64 `protobuf:"varint,6,opt,name=physicalUsedStorage,proto3" json:"physicalUsedStorage,omitempty"`
	// Physical Used Storage
	//
	// Physical storage space occupied by all data including Parity overhead when EC enabled (physical n*disks/(n*disks -1)), given in bytes.
	PhysicalUsedStorageIncludingParity uint64 `protobuf:"varint,7,opt,name=physicalUsedStorageIncludingParity,proto3" json:"physicalUsedStorageIncludingParity,omitempty"`
	// Free Physical Storage
	//
	// Free storage before entering to read-only mode , given in bytes.
	FreePhysicalStorage uint64 `protobuf:"varint,8,opt,name=freePhysicalStorage,proto3" json:"freePhysicalStorage,omitempty"`
	// Estimated Free Logical Storage
	//
	// Estimated free storage before entering to read-only mode assuming current compression ratio, given in bytes.
	EstimatedFreeLogicalStorage uint64 `protobuf:"varint,9,opt,name=estimatedFreeLogicalStorage,proto3" json:"estimatedFreeLogicalStorage,omitempty"`
	// Estimated Total Available Logical Storage
	//
	// Estimate of total available logical storage based on current compression ratio (effective * compression)
	EstimatedLogicalStorage uint64 `protobuf:"varint,10,opt,name=estimatedLogicalStorage,proto3" json:"estimatedLogicalStorage,omitempty"`
	// commpression ratio
	//
	// compression ratio logicalUsedStorage/physicalUsedStorage
	CompressionRatio float64 `protobuf:"fixed64,11,opt,name=compressionRatio,proto3" json:"compressionRatio,omitempty"`
}

func (x *ClusterStatisticsApi) Reset() {
	*x = ClusterStatisticsApi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lightbits_api_duros_v1_statisticsapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterStatisticsApi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterStatisticsApi) ProtoMessage() {}

func (x *ClusterStatisticsApi) ProtoReflect() protoreflect.Message {
	mi := &file_lightbits_api_duros_v1_statisticsapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterStatisticsApi.ProtoReflect.Descriptor instead.
func (*ClusterStatisticsApi) Descriptor() ([]byte, []int) {
	return file_lightbits_api_duros_v1_statisticsapi_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterStatisticsApi) GetInstalledPhysicalStorage() uint64 {
	if x != nil {
		return x.InstalledPhysicalStorage
	}
	return 0
}

func (x *ClusterStatisticsApi) GetManagedPhysicalStorage() uint64 {
	if x != nil {
		return x.ManagedPhysicalStorage
	}
	return 0
}

func (x *ClusterStatisticsApi) GetEffectivePhysicalStorage() uint64 {
	if x != nil {
		return x.EffectivePhysicalStorage
	}
	return 0
}

func (x *ClusterStatisticsApi) GetLogicalStorage() uint64 {
	if x != nil {
		return x.LogicalStorage
	}
	return 0
}

func (x *ClusterStatisticsApi) GetLogicalUsedStorage() uint64 {
	if x != nil {
		return x.LogicalUsedStorage
	}
	return 0
}

func (x *ClusterStatisticsApi) GetPhysicalUsedStorage() uint64 {
	if x != nil {
		return x.PhysicalUsedStorage
	}
	return 0
}

func (x *ClusterStatisticsApi) GetPhysicalUsedStorageIncludingParity() uint64 {
	if x != nil {
		return x.PhysicalUsedStorageIncludingParity
	}
	return 0
}

func (x *ClusterStatisticsApi) GetFreePhysicalStorage() uint64 {
	if x != nil {
		return x.FreePhysicalStorage
	}
	return 0
}

func (x *ClusterStatisticsApi) GetEstimatedFreeLogicalStorage() uint64 {
	if x != nil {
		return x.EstimatedFreeLogicalStorage
	}
	return 0
}

func (x *ClusterStatisticsApi) GetEstimatedLogicalStorage() uint64 {
	if x != nil {
		return x.EstimatedLogicalStorage
	}
	return 0
}

func (x *ClusterStatisticsApi) GetCompressionRatio() float64 {
	if x != nil {
		return x.CompressionRatio
	}
	return 0
}

type NodeStatisticsApi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total Attached Physical Total Storage
	//
	// Sum of all managed and healthy SSDs capacities, given in bytes.
	ManagedPhysicalStorage uint64 `protobuf:"varint,1,opt,name=managedPhysicalStorage,proto3" json:"managedPhysicalStorage,omitempty"`
	// Effective Physical Storage
	//
	// Effective Physical storage excluding overhead of OVP and Parity, given in bytes.
	EffectivePhysicalStorage uint64 `protobuf:"varint,2,opt,name=effectivePhysicalStorage,proto3" json:"effectivePhysicalStorage,omitempty"`
	// Logical Storage
	//
	// Sum of capacities of all allocated volumes, given in bytes.
	LogicalStorage uint64 `protobuf:"varint,3,opt,name=logicalStorage,proto3" json:"logicalStorage,omitempty"`
	// Logical Used Storage
	//
	// Logical storage space used by all volumes (n of LBAs x 4096), given in bytes.
	LogicalUsedStorage uint64 `protobuf:"varint,4,opt,name=logicalUsedStorage,proto3" json:"logicalUsedStorage,omitempty"`
	// Physical Used Storage Excluding Parity
	//
	// Physical storage space occupied by all volumes (data only), given in bytes.
	PhysicalUsedStorage uint64 `protobuf:"varint,5,opt,name=physicalUsedStorage,proto3" json:"physicalUsedStorage,omitempty"`
	// Physical Used Storage
	//
	// Physical storage space occupied by all data including Parity overhead when EC enabled (physical n*disks/(n*disks -1)), given in bytes.
	PhysicalUsedStorageIncludingParity uint64 `protobuf:"varint,6,opt,name=physicalUsedStorageIncludingParity,proto3" json:"physicalUsedStorageIncludingParity,omitempty"`
	// Free Physical Storage
	//
	// Free storage before entering to read-only mode, given in bytes.
	FreePhysicalStorage uint64 `protobuf:"varint,7,opt,name=freePhysicalStorage,proto3" json:"freePhysicalStorage,omitempty"`
	// Estimated Free Logical Storage
	//
	// Estimated free storage before entering to read-only mode assuming current compression ratio, given in bytes.
	EstimatedFreeLogicalStorage uint64 `protobuf:"varint,8,opt,name=estimatedFreeLogicalStorage,proto3" json:"estimatedFreeLogicalStorage,omitempty"`
	// Estimated Total Available Logical Storage
	//
	// Estimate of total available logical storage based on current compression ratio (effective * compression)
	EstimatedLogicalStorage uint64 `protobuf:"varint,9,opt,name=estimatedLogicalStorage,proto3" json:"estimatedLogicalStorage,omitempty"`
	// commpression ratio
	//
	// compression ratio logicalUsedStorage/physicalUsedStorage
	CompressionRatio float64 `protobuf:"fixed64,10,opt,name=compressionRatio,proto3" json:"compressionRatio,omitempty"`
}

func (x *NodeStatisticsApi) Reset() {
	*x = NodeStatisticsApi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lightbits_api_duros_v1_statisticsapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStatisticsApi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatisticsApi) ProtoMessage() {}

func (x *NodeStatisticsApi) ProtoReflect() protoreflect.Message {
	mi := &file_lightbits_api_duros_v1_statisticsapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatisticsApi.ProtoReflect.Descriptor instead.
func (*NodeStatisticsApi) Descriptor() ([]byte, []int) {
	return file_lightbits_api_duros_v1_statisticsapi_proto_rawDescGZIP(), []int{1}
}

func (x *NodeStatisticsApi) GetManagedPhysicalStorage() uint64 {
	if x != nil {
		return x.ManagedPhysicalStorage
	}
	return 0
}

func (x *NodeStatisticsApi) GetEffectivePhysicalStorage() uint64 {
	if x != nil {
		return x.EffectivePhysicalStorage
	}
	return 0
}

func (x *NodeStatisticsApi) GetLogicalStorage() uint64 {
	if x != nil {
		return x.LogicalStorage
	}
	return 0
}

func (x *NodeStatisticsApi) GetLogicalUsedStorage() uint64 {
	if x != nil {
		return x.LogicalUsedStorage
	}
	return 0
}

func (x *NodeStatisticsApi) GetPhysicalUsedStorage() uint64 {
	if x != nil {
		return x.PhysicalUsedStorage
	}
	return 0
}

func (x *NodeStatisticsApi) GetPhysicalUsedStorageIncludingParity() uint64 {
	if x != nil {
		return x.PhysicalUsedStorageIncludingParity
	}
	return 0
}

func (x *NodeStatisticsApi) GetFreePhysicalStorage() uint64 {
	if x != nil {
		return x.FreePhysicalStorage
	}
	return 0
}

func (x *NodeStatisticsApi) GetEstimatedFreeLogicalStorage() uint64 {
	if x != nil {
		return x.EstimatedFreeLogicalStorage
	}
	return 0
}

func (x *NodeStatisticsApi) GetEstimatedLogicalStorage() uint64 {
	if x != nil {
		return x.EstimatedLogicalStorage
	}
	return 0
}

func (x *NodeStatisticsApi) GetCompressionRatio() float64 {
	if x != nil {
		return x.CompressionRatio
	}
	return 0
}

type VolumeStatisticsApi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Logical Used Storage
	//
	// Logical storage space used by volume, given in bytes.
	LogicalUsedStorage uint64 `protobuf:"varint,1,opt,name=logicalUsedStorage,proto3" json:"logicalUsedStorage,omitempty"`
	// Physical Used Storage
	//
	// Physical storage space used by volume excluding parity, given in bytes.
	PhysicalUsedStorage uint64 `protobuf:"varint,2,opt,name=physicalUsedStorage,proto3" json:"physicalUsedStorage,omitempty"`
	// commpression ratio
	//
	// compression ratio logicalUsedStorage/physicalUsedStorage
	CompressionRatio float64 `protobuf:"fixed64,3,opt,name=compressionRatio,proto3" json:"compressionRatio,omitempty"`
}

func (x *VolumeStatisticsApi) Reset() {
	*x = VolumeStatisticsApi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lightbits_api_duros_v1_statisticsapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeStatisticsApi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeStatisticsApi) ProtoMessage() {}

func (x *VolumeStatisticsApi) ProtoReflect() protoreflect.Message {
	mi := &file_lightbits_api_duros_v1_statisticsapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeStatisticsApi.ProtoReflect.Descriptor instead.
func (*VolumeStatisticsApi) Descriptor() ([]byte, []int) {
	return file_lightbits_api_duros_v1_statisticsapi_proto_rawDescGZIP(), []int{2}
}

func (x *VolumeStatisticsApi) GetLogicalUsedStorage() uint64 {
	if x != nil {
		return x.LogicalUsedStorage
	}
	return 0
}

func (x *VolumeStatisticsApi) GetPhysicalUsedStorage() uint64 {
	if x != nil {
		return x.PhysicalUsedStorage
	}
	return 0
}

func (x *VolumeStatisticsApi) GetCompressionRatio() float64 {
	if x != nil {
		return x.CompressionRatio
	}
	return 0
}

var File_lightbits_api_duros_v1_statisticsapi_proto protoreflect.FileDescriptor

var file_lightbits_api_duros_v1_statisticsapi_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x75, 0x72, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x62, 0x69, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x75, 0x72, 0x6f,
	0x73, 0x2e, 0x76, 0x31, 0x22, 0xfa, 0x04, 0x0a, 0x14, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x41, 0x70, 0x69, 0x12, 0x3a, 0x0a,
	0x18, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63,
	0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x18, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63,
	0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x12, 0x3a, 0x0a, 0x18, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x68,
	0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x18, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x68,
	0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c,
	0x55, 0x73, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x12, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x64, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61,
	0x6c, 0x55, 0x73, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x13, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x64,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x22, 0x70, 0x68, 0x79, 0x73, 0x69,
	0x63, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x22, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x55, 0x73, 0x65,
	0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e,
	0x67, 0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x72, 0x65, 0x65, 0x50,
	0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x66, 0x72, 0x65, 0x65, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63,
	0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x1b, 0x65, 0x73, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x46, 0x72, 0x65, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x61,
	0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1b,
	0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x46, 0x72, 0x65, 0x65, 0x4c, 0x6f, 0x67,
	0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x65,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x65, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69,
	0x6f, 0x22, 0xbb, 0x04, 0x0a, 0x11, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x41, 0x70, 0x69, 0x12, 0x36, 0x0a, 0x16, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x16, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x3a, 0x0a, 0x18, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x68, 0x79, 0x73,
	0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x18, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x68, 0x79, 0x73,
	0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6c,
	0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x55, 0x73,
	0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x12, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x55,
	0x73, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x13, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x64, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x22, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61,
	0x6c, 0x55, 0x73, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x22, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x64, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x50,
	0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x72, 0x65, 0x65, 0x50, 0x68, 0x79,
	0x73, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x13, 0x66, 0x72, 0x65, 0x65, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x1b, 0x65, 0x73, 0x74, 0x69, 0x6d,
	0x61, 0x74, 0x65, 0x64, 0x46, 0x72, 0x65, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1b, 0x65, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x46, 0x72, 0x65, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x63,
	0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x65, 0x73, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x65, 0x73, 0x74, 0x69,
	0x6d, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x22,
	0xa3, 0x01, 0x0a, 0x13, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x41, 0x70, 0x69, 0x12, 0x2e, 0x0a, 0x12, 0x6c, 0x6f, 0x67, 0x69, 0x63,
	0x61, 0x6c, 0x55, 0x73, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x12, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x64,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x68, 0x79, 0x73, 0x69,
	0x63, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x55, 0x73,
	0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x61, 0x74, 0x69, 0x6f, 0x42, 0x0d, 0x5a, 0x0b, 0x64, 0x75, 0x72, 0x6f, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lightbits_api_duros_v1_statisticsapi_proto_rawDescOnce sync.Once
	file_lightbits_api_duros_v1_statisticsapi_proto_rawDescData = file_lightbits_api_duros_v1_statisticsapi_proto_rawDesc
)

func file_lightbits_api_duros_v1_statisticsapi_proto_rawDescGZIP() []byte {
	file_lightbits_api_duros_v1_statisticsapi_proto_rawDescOnce.Do(func() {
		file_lightbits_api_duros_v1_statisticsapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_lightbits_api_duros_v1_statisticsapi_proto_rawDescData)
	})
	return file_lightbits_api_duros_v1_statisticsapi_proto_rawDescData
}

var file_lightbits_api_duros_v1_statisticsapi_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_lightbits_api_duros_v1_statisticsapi_proto_goTypes = []interface{}{
	(*ClusterStatisticsApi)(nil), // 0: lightbits.api.duros.v1.ClusterStatisticsApi
	(*NodeStatisticsApi)(nil),    // 1: lightbits.api.duros.v1.NodeStatisticsApi
	(*VolumeStatisticsApi)(nil),  // 2: lightbits.api.duros.v1.VolumeStatisticsApi
}
var file_lightbits_api_duros_v1_statisticsapi_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lightbits_api_duros_v1_statisticsapi_proto_init() }
func file_lightbits_api_duros_v1_statisticsapi_proto_init() {
	if File_lightbits_api_duros_v1_statisticsapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lightbits_api_duros_v1_statisticsapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterStatisticsApi); i {
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
		file_lightbits_api_duros_v1_statisticsapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStatisticsApi); i {
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
		file_lightbits_api_duros_v1_statisticsapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeStatisticsApi); i {
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
			RawDescriptor: file_lightbits_api_duros_v1_statisticsapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lightbits_api_duros_v1_statisticsapi_proto_goTypes,
		DependencyIndexes: file_lightbits_api_duros_v1_statisticsapi_proto_depIdxs,
		MessageInfos:      file_lightbits_api_duros_v1_statisticsapi_proto_msgTypes,
	}.Build()
	File_lightbits_api_duros_v1_statisticsapi_proto = out.File
	file_lightbits_api_duros_v1_statisticsapi_proto_rawDesc = nil
	file_lightbits_api_duros_v1_statisticsapi_proto_goTypes = nil
	file_lightbits_api_duros_v1_statisticsapi_proto_depIdxs = nil
}
