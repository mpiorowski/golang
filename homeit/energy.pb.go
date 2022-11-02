// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: energy.proto

package homeit

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

type TariffId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	TariffId string `protobuf:"bytes,2,opt,name=tariffId,proto3" json:"tariffId,omitempty"`
}

func (x *TariffId) Reset() {
	*x = TariffId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_energy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TariffId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TariffId) ProtoMessage() {}

func (x *TariffId) ProtoReflect() protoreflect.Message {
	mi := &file_energy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TariffId.ProtoReflect.Descriptor instead.
func (*TariffId) Descriptor() ([]byte, []int) {
	return file_energy_proto_rawDescGZIP(), []int{0}
}

func (x *TariffId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TariffId) GetTariffId() string {
	if x != nil {
		return x.TariffId
	}
	return ""
}

type MeterId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	MeterId string `protobuf:"bytes,2,opt,name=meterId,proto3" json:"meterId,omitempty"`
}

func (x *MeterId) Reset() {
	*x = MeterId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_energy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeterId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeterId) ProtoMessage() {}

func (x *MeterId) ProtoReflect() protoreflect.Message {
	mi := &file_energy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeterId.ProtoReflect.Descriptor instead.
func (*MeterId) Descriptor() ([]byte, []int) {
	return file_energy_proto_rawDescGZIP(), []int{1}
}

func (x *MeterId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MeterId) GetMeterId() string {
	if x != nil {
		return x.MeterId
	}
	return ""
}

type MeasurementId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	MeasurementId string `protobuf:"bytes,2,opt,name=measurementId,proto3" json:"measurementId,omitempty"`
}

func (x *MeasurementId) Reset() {
	*x = MeasurementId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_energy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeasurementId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasurementId) ProtoMessage() {}

func (x *MeasurementId) ProtoReflect() protoreflect.Message {
	mi := &file_energy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasurementId.ProtoReflect.Descriptor instead.
func (*MeasurementId) Descriptor() ([]byte, []int) {
	return file_energy_proto_rawDescGZIP(), []int{2}
}

func (x *MeasurementId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MeasurementId) GetMeasurementId() string {
	if x != nil {
		return x.MeasurementId
	}
	return ""
}

type Tariff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created        string          `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated        string          `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted        *string         `protobuf:"bytes,4,opt,name=deleted,proto3,oneof" json:"deleted,omitempty"`
	UserId         *string         `protobuf:"bytes,5,opt,name=userId,proto3,oneof" json:"userId,omitempty"`
	Global         bool            `protobuf:"varint,6,opt,name=global,proto3" json:"global,omitempty"`
	Name           string          `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Sectors        []*TariffSector `protobuf:"bytes,8,rep,name=sectors,proto3" json:"sectors,omitempty"`
	PricesPerKwh   []*TariffPrice  `protobuf:"bytes,9,rep,name=pricesPerKwh,proto3" json:"pricesPerKwh,omitempty"`
	PricesPerMonth []*TariffPrice  `protobuf:"bytes,10,rep,name=pricesPerMonth,proto3" json:"pricesPerMonth,omitempty"`
}

func (x *Tariff) Reset() {
	*x = Tariff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_energy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tariff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tariff) ProtoMessage() {}

func (x *Tariff) ProtoReflect() protoreflect.Message {
	mi := &file_energy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tariff.ProtoReflect.Descriptor instead.
func (*Tariff) Descriptor() ([]byte, []int) {
	return file_energy_proto_rawDescGZIP(), []int{3}
}

func (x *Tariff) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tariff) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *Tariff) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *Tariff) GetDeleted() string {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return ""
}

func (x *Tariff) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *Tariff) GetGlobal() bool {
	if x != nil {
		return x.Global
	}
	return false
}

func (x *Tariff) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tariff) GetSectors() []*TariffSector {
	if x != nil {
		return x.Sectors
	}
	return nil
}

func (x *Tariff) GetPricesPerKwh() []*TariffPrice {
	if x != nil {
		return x.PricesPerKwh
	}
	return nil
}

func (x *Tariff) GetPricesPerMonth() []*TariffPrice {
	if x != nil {
		return x.PricesPerMonth
	}
	return nil
}

type TariffSector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weekends      bool     `protobuf:"varint,2,opt,name=weekends,proto3" json:"weekends,omitempty"`
	StablePrice   float64  `protobuf:"fixed64,3,opt,name=stablePrice,proto3" json:"stablePrice,omitempty"`
	VariablePrice float64  `protobuf:"fixed64,4,opt,name=variablePrice,proto3" json:"variablePrice,omitempty"`
	TimeSector    []string `protobuf:"bytes,5,rep,name=timeSector,proto3" json:"timeSector,omitempty"`
}

func (x *TariffSector) Reset() {
	*x = TariffSector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_energy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TariffSector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TariffSector) ProtoMessage() {}

func (x *TariffSector) ProtoReflect() protoreflect.Message {
	mi := &file_energy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TariffSector.ProtoReflect.Descriptor instead.
func (*TariffSector) Descriptor() ([]byte, []int) {
	return file_energy_proto_rawDescGZIP(), []int{4}
}

func (x *TariffSector) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TariffSector) GetWeekends() bool {
	if x != nil {
		return x.Weekends
	}
	return false
}

func (x *TariffSector) GetStablePrice() float64 {
	if x != nil {
		return x.StablePrice
	}
	return 0
}

func (x *TariffSector) GetVariablePrice() float64 {
	if x != nil {
		return x.VariablePrice
	}
	return 0
}

func (x *TariffSector) GetTimeSector() []string {
	if x != nil {
		return x.TimeSector
	}
	return nil
}

type TariffPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *TariffPrice) Reset() {
	*x = TariffPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_energy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TariffPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TariffPrice) ProtoMessage() {}

func (x *TariffPrice) ProtoReflect() protoreflect.Message {
	mi := &file_energy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TariffPrice.ProtoReflect.Descriptor instead.
func (*TariffPrice) Descriptor() ([]byte, []int) {
	return file_energy_proto_rawDescGZIP(), []int{5}
}

func (x *TariffPrice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TariffPrice) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Meter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created string  `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated string  `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted *string `protobuf:"bytes,4,opt,name=deleted,proto3,oneof" json:"deleted,omitempty"`
	UserId  string  `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
	Name    string  `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Meter) Reset() {
	*x = Meter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_energy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meter) ProtoMessage() {}

func (x *Meter) ProtoReflect() protoreflect.Message {
	mi := &file_energy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meter.ProtoReflect.Descriptor instead.
func (*Meter) Descriptor() ([]byte, []int) {
	return file_energy_proto_rawDescGZIP(), []int{6}
}

func (x *Meter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Meter) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *Meter) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *Meter) GetDeleted() string {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return ""
}

func (x *Meter) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Meter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Measurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created string   `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated string   `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Deleted *string  `protobuf:"bytes,4,opt,name=deleted,proto3,oneof" json:"deleted,omitempty"`
	MeterId string   `protobuf:"bytes,5,opt,name=meterId,proto3" json:"meterId,omitempty"`
	Value   *float64 `protobuf:"fixed64,6,opt,name=value,proto3,oneof" json:"value,omitempty"`
	Date    string   `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Measurement) Reset() {
	*x = Measurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_energy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Measurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measurement) ProtoMessage() {}

func (x *Measurement) ProtoReflect() protoreflect.Message {
	mi := &file_energy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measurement.ProtoReflect.Descriptor instead.
func (*Measurement) Descriptor() ([]byte, []int) {
	return file_energy_proto_rawDescGZIP(), []int{7}
}

func (x *Measurement) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Measurement) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *Measurement) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *Measurement) GetDeleted() string {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return ""
}

func (x *Measurement) GetMeterId() string {
	if x != nil {
		return x.MeterId
	}
	return ""
}

func (x *Measurement) GetValue() float64 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

func (x *Measurement) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type EnergySectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	TariffId string `protobuf:"bytes,2,opt,name=tariffId,proto3" json:"tariffId,omitempty"`
	MeterId  string `protobuf:"bytes,3,opt,name=meterId,proto3" json:"meterId,omitempty"`
	DateFrom string `protobuf:"bytes,4,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"`
	DateTo   string `protobuf:"bytes,5,opt,name=dateTo,proto3" json:"dateTo,omitempty"`
}

func (x *EnergySectorRequest) Reset() {
	*x = EnergySectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_energy_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnergySectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnergySectorRequest) ProtoMessage() {}

func (x *EnergySectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_energy_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnergySectorRequest.ProtoReflect.Descriptor instead.
func (*EnergySectorRequest) Descriptor() ([]byte, []int) {
	return file_energy_proto_rawDescGZIP(), []int{8}
}

func (x *EnergySectorRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EnergySectorRequest) GetTariffId() string {
	if x != nil {
		return x.TariffId
	}
	return ""
}

func (x *EnergySectorRequest) GetMeterId() string {
	if x != nil {
		return x.MeterId
	}
	return ""
}

func (x *EnergySectorRequest) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *EnergySectorRequest) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

type EnergySectorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TariffId string  `protobuf:"bytes,1,opt,name=tariffId,proto3" json:"tariffId,omitempty"`
	MeterId  string  `protobuf:"bytes,2,opt,name=meterId,proto3" json:"meterId,omitempty"`
	Name     string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Value    float64 `protobuf:"fixed64,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EnergySectorResponse) Reset() {
	*x = EnergySectorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_energy_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnergySectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnergySectorResponse) ProtoMessage() {}

func (x *EnergySectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_energy_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnergySectorResponse.ProtoReflect.Descriptor instead.
func (*EnergySectorResponse) Descriptor() ([]byte, []int) {
	return file_energy_proto_rawDescGZIP(), []int{9}
}

func (x *EnergySectorResponse) GetTariffId() string {
	if x != nil {
		return x.TariffId
	}
	return ""
}

func (x *EnergySectorResponse) GetMeterId() string {
	if x != nil {
		return x.MeterId
	}
	return ""
}

func (x *EnergySectorResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnergySectorResponse) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_energy_proto protoreflect.FileDescriptor

var file_energy_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x22, 0x3e, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x0d, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d,
	0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0xf1, 0x02, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1d, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x6e, 0x65,
	0x72, 0x67, 0x79, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x50, 0x65, 0x72, 0x4b, 0x77, 0x68, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x50, 0x65, 0x72, 0x4b,
	0x77, 0x68, 0x12, 0x3b, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x50, 0x65, 0x72, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6e, 0x65,
	0x72, 0x67, 0x79, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52,
	0x0e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa6, 0x01, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77,
	0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x77,
	0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x73, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22,
	0x37, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x05, 0x4d, 0x65, 0x74,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0xcf, 0x01,
	0x0a, 0x0b, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1d, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x97, 0x01, 0x0a, 0x13, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x22, 0x76, 0x0a, 0x14, 0x45, 0x6e, 0x65,
	0x72, 0x67, 0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x4a, 0x0a, 0x14, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x68, 0x6f, 0x6d,
	0x65, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0b, 0x48, 0x6f, 0x6d, 0x65, 0x49,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x70, 0x69, 0x6f, 0x72, 0x6f, 0x77, 0x73, 0x6b, 0x69, 0x2f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x69, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_energy_proto_rawDescOnce sync.Once
	file_energy_proto_rawDescData = file_energy_proto_rawDesc
)

func file_energy_proto_rawDescGZIP() []byte {
	file_energy_proto_rawDescOnce.Do(func() {
		file_energy_proto_rawDescData = protoimpl.X.CompressGZIP(file_energy_proto_rawDescData)
	})
	return file_energy_proto_rawDescData
}

var file_energy_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_energy_proto_goTypes = []interface{}{
	(*TariffId)(nil),             // 0: energy.TariffId
	(*MeterId)(nil),              // 1: energy.MeterId
	(*MeasurementId)(nil),        // 2: energy.MeasurementId
	(*Tariff)(nil),               // 3: energy.Tariff
	(*TariffSector)(nil),         // 4: energy.TariffSector
	(*TariffPrice)(nil),          // 5: energy.TariffPrice
	(*Meter)(nil),                // 6: energy.Meter
	(*Measurement)(nil),          // 7: energy.Measurement
	(*EnergySectorRequest)(nil),  // 8: energy.EnergySectorRequest
	(*EnergySectorResponse)(nil), // 9: energy.EnergySectorResponse
}
var file_energy_proto_depIdxs = []int32{
	4, // 0: energy.Tariff.sectors:type_name -> energy.TariffSector
	5, // 1: energy.Tariff.pricesPerKwh:type_name -> energy.TariffPrice
	5, // 2: energy.Tariff.pricesPerMonth:type_name -> energy.TariffPrice
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_energy_proto_init() }
func file_energy_proto_init() {
	if File_energy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_energy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TariffId); i {
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
		file_energy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeterId); i {
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
		file_energy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeasurementId); i {
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
		file_energy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tariff); i {
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
		file_energy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TariffSector); i {
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
		file_energy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TariffPrice); i {
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
		file_energy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meter); i {
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
		file_energy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Measurement); i {
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
		file_energy_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnergySectorRequest); i {
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
		file_energy_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnergySectorResponse); i {
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
	file_energy_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_energy_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_energy_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_energy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_energy_proto_goTypes,
		DependencyIndexes: file_energy_proto_depIdxs,
		MessageInfos:      file_energy_proto_msgTypes,
	}.Build()
	File_energy_proto = out.File
	file_energy_proto_rawDesc = nil
	file_energy_proto_goTypes = nil
	file_energy_proto_depIdxs = nil
}
