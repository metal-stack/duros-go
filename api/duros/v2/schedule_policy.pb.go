// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: lightbits/api/duros/v2/schedule_policy.proto

package v2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DayOfWeek int32

const (
	DayOfWeek_DayOfWeekUnspecified DayOfWeek = 0
	DayOfWeek_Sunday               DayOfWeek = 1
	DayOfWeek_Monday               DayOfWeek = 2
	DayOfWeek_Tuesday              DayOfWeek = 3
	DayOfWeek_Wednesday            DayOfWeek = 4
	DayOfWeek_Thursday             DayOfWeek = 5
	DayOfWeek_Friday               DayOfWeek = 6
	DayOfWeek_Saturday             DayOfWeek = 7
)

// Enum value maps for DayOfWeek.
var (
	DayOfWeek_name = map[int32]string{
		0: "DayOfWeekUnspecified",
		1: "Sunday",
		2: "Monday",
		3: "Tuesday",
		4: "Wednesday",
		5: "Thursday",
		6: "Friday",
		7: "Saturday",
	}
	DayOfWeek_value = map[string]int32{
		"DayOfWeekUnspecified": 0,
		"Sunday":               1,
		"Monday":               2,
		"Tuesday":              3,
		"Wednesday":            4,
		"Thursday":             5,
		"Friday":               6,
		"Saturday":             7,
	}
)

func (x DayOfWeek) Enum() *DayOfWeek {
	p := new(DayOfWeek)
	*p = x
	return p
}

func (x DayOfWeek) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DayOfWeek) Descriptor() protoreflect.EnumDescriptor {
	return file_lightbits_api_duros_v2_schedule_policy_proto_enumTypes[0].Descriptor()
}

func (DayOfWeek) Type() protoreflect.EnumType {
	return &file_lightbits_api_duros_v2_schedule_policy_proto_enumTypes[0]
}

func (x DayOfWeek) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DayOfWeek.Descriptor instead.
func (DayOfWeek) EnumDescriptor() ([]byte, []int) {
	return file_lightbits_api_duros_v2_schedule_policy_proto_rawDescGZIP(), []int{0}
}

type HourlySchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	HoursInCycle uint32                 `protobuf:"varint,2,opt,name=hoursInCycle,proto3" json:"hoursInCycle,omitempty"`
}

func (x *HourlySchedule) Reset() {
	*x = HourlySchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HourlySchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HourlySchedule) ProtoMessage() {}

func (x *HourlySchedule) ProtoReflect() protoreflect.Message {
	mi := &file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HourlySchedule.ProtoReflect.Descriptor instead.
func (*HourlySchedule) Descriptor() ([]byte, []int) {
	return file_lightbits_api_duros_v2_schedule_policy_proto_rawDescGZIP(), []int{0}
}

func (x *HourlySchedule) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *HourlySchedule) GetHoursInCycle() uint32 {
	if x != nil {
		return x.HoursInCycle
	}
	return 0
}

type DailySchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	DaysInCycle uint32                 `protobuf:"varint,2,opt,name=daysInCycle,proto3" json:"daysInCycle,omitempty"`
}

func (x *DailySchedule) Reset() {
	*x = DailySchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailySchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailySchedule) ProtoMessage() {}

func (x *DailySchedule) ProtoReflect() protoreflect.Message {
	mi := &file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailySchedule.ProtoReflect.Descriptor instead.
func (*DailySchedule) Descriptor() ([]byte, []int) {
	return file_lightbits_api_duros_v2_schedule_policy_proto_rawDescGZIP(), []int{1}
}

func (x *DailySchedule) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *DailySchedule) GetDaysInCycle() uint32 {
	if x != nil {
		return x.DaysInCycle
	}
	return 0
}

type DayOfWeekEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	Day       DayOfWeek              `protobuf:"varint,2,opt,name=day,proto3,enum=lightbits.api.duros.v2.DayOfWeek" json:"day,omitempty"`
}

func (x *DayOfWeekEntry) Reset() {
	*x = DayOfWeekEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DayOfWeekEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DayOfWeekEntry) ProtoMessage() {}

func (x *DayOfWeekEntry) ProtoReflect() protoreflect.Message {
	mi := &file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DayOfWeekEntry.ProtoReflect.Descriptor instead.
func (*DayOfWeekEntry) Descriptor() ([]byte, []int) {
	return file_lightbits_api_duros_v2_schedule_policy_proto_rawDescGZIP(), []int{2}
}

func (x *DayOfWeekEntry) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *DayOfWeekEntry) GetDay() DayOfWeek {
	if x != nil {
		return x.Day
	}
	return DayOfWeek_DayOfWeekUnspecified
}

type WeeklySchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DaysOfWeek []*DayOfWeekEntry `protobuf:"bytes,1,rep,name=daysOfWeek,proto3" json:"daysOfWeek,omitempty"`
}

func (x *WeeklySchedule) Reset() {
	*x = WeeklySchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeeklySchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeeklySchedule) ProtoMessage() {}

func (x *WeeklySchedule) ProtoReflect() protoreflect.Message {
	mi := &file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeeklySchedule.ProtoReflect.Descriptor instead.
func (*WeeklySchedule) Descriptor() ([]byte, []int) {
	return file_lightbits_api_duros_v2_schedule_policy_proto_rawDescGZIP(), []int{3}
}

func (x *WeeklySchedule) GetDaysOfWeek() []*DayOfWeekEntry {
	if x != nil {
		return x.DaysOfWeek
	}
	return nil
}

type SnapshotSchedulePolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to SchedulePolicies:
	//	*SnapshotSchedulePolicy_HourlySchedule
	//	*SnapshotSchedulePolicy_DailySchedule
	//	*SnapshotSchedulePolicy_WeeklySchedule
	SchedulePolicies isSnapshotSchedulePolicy_SchedulePolicies `protobuf_oneof:"schedulePolicies"`
}

func (x *SnapshotSchedulePolicy) Reset() {
	*x = SnapshotSchedulePolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotSchedulePolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotSchedulePolicy) ProtoMessage() {}

func (x *SnapshotSchedulePolicy) ProtoReflect() protoreflect.Message {
	mi := &file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotSchedulePolicy.ProtoReflect.Descriptor instead.
func (*SnapshotSchedulePolicy) Descriptor() ([]byte, []int) {
	return file_lightbits_api_duros_v2_schedule_policy_proto_rawDescGZIP(), []int{4}
}

func (m *SnapshotSchedulePolicy) GetSchedulePolicies() isSnapshotSchedulePolicy_SchedulePolicies {
	if m != nil {
		return m.SchedulePolicies
	}
	return nil
}

func (x *SnapshotSchedulePolicy) GetHourlySchedule() *HourlySchedule {
	if x, ok := x.GetSchedulePolicies().(*SnapshotSchedulePolicy_HourlySchedule); ok {
		return x.HourlySchedule
	}
	return nil
}

func (x *SnapshotSchedulePolicy) GetDailySchedule() *DailySchedule {
	if x, ok := x.GetSchedulePolicies().(*SnapshotSchedulePolicy_DailySchedule); ok {
		return x.DailySchedule
	}
	return nil
}

func (x *SnapshotSchedulePolicy) GetWeeklySchedule() *WeeklySchedule {
	if x, ok := x.GetSchedulePolicies().(*SnapshotSchedulePolicy_WeeklySchedule); ok {
		return x.WeeklySchedule
	}
	return nil
}

type isSnapshotSchedulePolicy_SchedulePolicies interface {
	isSnapshotSchedulePolicy_SchedulePolicies()
}

type SnapshotSchedulePolicy_HourlySchedule struct {
	HourlySchedule *HourlySchedule `protobuf:"bytes,1,opt,name=hourlySchedule,proto3,oneof"`
}

type SnapshotSchedulePolicy_DailySchedule struct {
	DailySchedule *DailySchedule `protobuf:"bytes,2,opt,name=dailySchedule,proto3,oneof"`
}

type SnapshotSchedulePolicy_WeeklySchedule struct {
	WeeklySchedule *WeeklySchedule `protobuf:"bytes,3,opt,name=weeklySchedule,proto3,oneof"`
}

func (*SnapshotSchedulePolicy_HourlySchedule) isSnapshotSchedulePolicy_SchedulePolicies() {}

func (*SnapshotSchedulePolicy_DailySchedule) isSnapshotSchedulePolicy_SchedulePolicies() {}

func (*SnapshotSchedulePolicy_WeeklySchedule) isSnapshotSchedulePolicy_SchedulePolicies() {}

type SchedulePolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to SchedulePolicies:
	//	*SchedulePolicy_SnapshotSchedulePolicy
	SchedulePolicies isSchedulePolicy_SchedulePolicies `protobuf_oneof:"schedulePolicies"`
	RetentionTime    *durationpb.Duration              `protobuf:"bytes,2,opt,name=retentionTime,proto3" json:"retentionTime,omitempty"`
}

func (x *SchedulePolicy) Reset() {
	*x = SchedulePolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulePolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulePolicy) ProtoMessage() {}

func (x *SchedulePolicy) ProtoReflect() protoreflect.Message {
	mi := &file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulePolicy.ProtoReflect.Descriptor instead.
func (*SchedulePolicy) Descriptor() ([]byte, []int) {
	return file_lightbits_api_duros_v2_schedule_policy_proto_rawDescGZIP(), []int{5}
}

func (m *SchedulePolicy) GetSchedulePolicies() isSchedulePolicy_SchedulePolicies {
	if m != nil {
		return m.SchedulePolicies
	}
	return nil
}

func (x *SchedulePolicy) GetSnapshotSchedulePolicy() *SnapshotSchedulePolicy {
	if x, ok := x.GetSchedulePolicies().(*SchedulePolicy_SnapshotSchedulePolicy); ok {
		return x.SnapshotSchedulePolicy
	}
	return nil
}

func (x *SchedulePolicy) GetRetentionTime() *durationpb.Duration {
	if x != nil {
		return x.RetentionTime
	}
	return nil
}

type isSchedulePolicy_SchedulePolicies interface {
	isSchedulePolicy_SchedulePolicies()
}

type SchedulePolicy_SnapshotSchedulePolicy struct {
	SnapshotSchedulePolicy *SnapshotSchedulePolicy `protobuf:"bytes,1,opt,name=snapshotSchedulePolicy,proto3,oneof"`
}

func (*SchedulePolicy_SnapshotSchedulePolicy) isSchedulePolicy_SchedulePolicies() {}

var File_lightbits_api_duros_v2_schedule_policy_proto protoreflect.FileDescriptor

var file_lightbits_api_duros_v2_schedule_policy_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x75, 0x72, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x69, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x75,
	0x72, 0x6f, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x0e, 0x48, 0x6f, 0x75, 0x72, 0x6c,
	0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x49, 0x6e, 0x43, 0x79,
	0x63, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x68, 0x6f, 0x75, 0x72, 0x73,
	0x49, 0x6e, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x22, 0x6b, 0x0a, 0x0d, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x49, 0x6e, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x49, 0x6e, 0x43,
	0x79, 0x63, 0x6c, 0x65, 0x22, 0x7f, 0x0a, 0x0e, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65,
	0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x33, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x69, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x75,
	0x72, 0x6f, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b,
	0x52, 0x03, 0x64, 0x61, 0x79, 0x22, 0x58, 0x0a, 0x0e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x64, 0x61, 0x79, 0x73, 0x4f,
	0x66, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x62, 0x69, 0x74, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x75, 0x72, 0x6f,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x22,
	0x9f, 0x02, 0x0a, 0x16, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x50, 0x0a, 0x0e, 0x68, 0x6f,
	0x75, 0x72, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x69, 0x74, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x75, 0x72, 0x6f, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x48, 0x6f, 0x75, 0x72,
	0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x68, 0x6f,
	0x75, 0x72, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x4d, 0x0a, 0x0d,
	0x64, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x69, 0x74, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x75, 0x72, 0x6f, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x50, 0x0a, 0x0e, 0x77,
	0x65, 0x65, 0x6b, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x69, 0x74, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x75, 0x72, 0x6f, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x65, 0x65,
	0x6b, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x77,
	0x65, 0x65, 0x6b, 0x6c, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x42, 0x12, 0x0a,
	0x10, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x22, 0xcf, 0x01, 0x0a, 0x0e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x68, 0x0a, 0x16, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x69, 0x74, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x75, 0x72, 0x6f, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x48, 0x00, 0x52, 0x16, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3f,
	0x0a, 0x0d, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42,
	0x12, 0x0a, 0x10, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x2a, 0x81, 0x01, 0x0a, 0x09, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65,
	0x6b, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x55, 0x6e,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53,
	0x75, 0x6e, 0x64, 0x61, 0x79, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x6f, 0x6e, 0x64, 0x61,
	0x79, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x75, 0x65, 0x73, 0x64, 0x61, 0x79, 0x10, 0x03,
	0x12, 0x0d, 0x0a, 0x09, 0x57, 0x65, 0x64, 0x6e, 0x65, 0x73, 0x64, 0x61, 0x79, 0x10, 0x04, 0x12,
	0x0c, 0x0a, 0x08, 0x54, 0x68, 0x75, 0x72, 0x73, 0x64, 0x61, 0x79, 0x10, 0x05, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x72, 0x69, 0x64, 0x61, 0x79, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x61, 0x74,
	0x75, 0x72, 0x64, 0x61, 0x79, 0x10, 0x07, 0x42, 0x0d, 0x5a, 0x0b, 0x64, 0x75, 0x72, 0x6f, 0x73,
	0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lightbits_api_duros_v2_schedule_policy_proto_rawDescOnce sync.Once
	file_lightbits_api_duros_v2_schedule_policy_proto_rawDescData = file_lightbits_api_duros_v2_schedule_policy_proto_rawDesc
)

func file_lightbits_api_duros_v2_schedule_policy_proto_rawDescGZIP() []byte {
	file_lightbits_api_duros_v2_schedule_policy_proto_rawDescOnce.Do(func() {
		file_lightbits_api_duros_v2_schedule_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_lightbits_api_duros_v2_schedule_policy_proto_rawDescData)
	})
	return file_lightbits_api_duros_v2_schedule_policy_proto_rawDescData
}

var file_lightbits_api_duros_v2_schedule_policy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_lightbits_api_duros_v2_schedule_policy_proto_goTypes = []interface{}{
	(DayOfWeek)(0),                 // 0: lightbits.api.duros.v2.DayOfWeek
	(*HourlySchedule)(nil),         // 1: lightbits.api.duros.v2.HourlySchedule
	(*DailySchedule)(nil),          // 2: lightbits.api.duros.v2.DailySchedule
	(*DayOfWeekEntry)(nil),         // 3: lightbits.api.duros.v2.DayOfWeekEntry
	(*WeeklySchedule)(nil),         // 4: lightbits.api.duros.v2.WeeklySchedule
	(*SnapshotSchedulePolicy)(nil), // 5: lightbits.api.duros.v2.SnapshotSchedulePolicy
	(*SchedulePolicy)(nil),         // 6: lightbits.api.duros.v2.SchedulePolicy
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),    // 8: google.protobuf.Duration
}
var file_lightbits_api_duros_v2_schedule_policy_proto_depIdxs = []int32{
	7,  // 0: lightbits.api.duros.v2.HourlySchedule.startTime:type_name -> google.protobuf.Timestamp
	7,  // 1: lightbits.api.duros.v2.DailySchedule.startTime:type_name -> google.protobuf.Timestamp
	7,  // 2: lightbits.api.duros.v2.DayOfWeekEntry.startTime:type_name -> google.protobuf.Timestamp
	0,  // 3: lightbits.api.duros.v2.DayOfWeekEntry.day:type_name -> lightbits.api.duros.v2.DayOfWeek
	3,  // 4: lightbits.api.duros.v2.WeeklySchedule.daysOfWeek:type_name -> lightbits.api.duros.v2.DayOfWeekEntry
	1,  // 5: lightbits.api.duros.v2.SnapshotSchedulePolicy.hourlySchedule:type_name -> lightbits.api.duros.v2.HourlySchedule
	2,  // 6: lightbits.api.duros.v2.SnapshotSchedulePolicy.dailySchedule:type_name -> lightbits.api.duros.v2.DailySchedule
	4,  // 7: lightbits.api.duros.v2.SnapshotSchedulePolicy.weeklySchedule:type_name -> lightbits.api.duros.v2.WeeklySchedule
	5,  // 8: lightbits.api.duros.v2.SchedulePolicy.snapshotSchedulePolicy:type_name -> lightbits.api.duros.v2.SnapshotSchedulePolicy
	8,  // 9: lightbits.api.duros.v2.SchedulePolicy.retentionTime:type_name -> google.protobuf.Duration
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_lightbits_api_duros_v2_schedule_policy_proto_init() }
func file_lightbits_api_duros_v2_schedule_policy_proto_init() {
	if File_lightbits_api_duros_v2_schedule_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HourlySchedule); i {
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
		file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailySchedule); i {
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
		file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DayOfWeekEntry); i {
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
		file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeeklySchedule); i {
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
		file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotSchedulePolicy); i {
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
		file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulePolicy); i {
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
	file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*SnapshotSchedulePolicy_HourlySchedule)(nil),
		(*SnapshotSchedulePolicy_DailySchedule)(nil),
		(*SnapshotSchedulePolicy_WeeklySchedule)(nil),
	}
	file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*SchedulePolicy_SnapshotSchedulePolicy)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lightbits_api_duros_v2_schedule_policy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lightbits_api_duros_v2_schedule_policy_proto_goTypes,
		DependencyIndexes: file_lightbits_api_duros_v2_schedule_policy_proto_depIdxs,
		EnumInfos:         file_lightbits_api_duros_v2_schedule_policy_proto_enumTypes,
		MessageInfos:      file_lightbits_api_duros_v2_schedule_policy_proto_msgTypes,
	}.Build()
	File_lightbits_api_duros_v2_schedule_policy_proto = out.File
	file_lightbits_api_duros_v2_schedule_policy_proto_rawDesc = nil
	file_lightbits_api_duros_v2_schedule_policy_proto_goTypes = nil
	file_lightbits_api_duros_v2_schedule_policy_proto_depIdxs = nil
}