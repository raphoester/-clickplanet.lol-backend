// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: clicks/v1/clicks.proto

package clicksv1

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

type ClickRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TileId    string `protobuf:"bytes,1,opt,name=tile_id,json=tileId,proto3" json:"tile_id,omitempty"`
	CountryId string `protobuf:"bytes,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
}

func (x *ClickRequest) Reset() {
	*x = ClickRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clicks_v1_clicks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickRequest) ProtoMessage() {}

func (x *ClickRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clicks_v1_clicks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickRequest.ProtoReflect.Descriptor instead.
func (*ClickRequest) Descriptor() ([]byte, []int) {
	return file_clicks_v1_clicks_proto_rawDescGZIP(), []int{0}
}

func (x *ClickRequest) GetTileId() string {
	if x != nil {
		return x.TileId
	}
	return ""
}

func (x *ClickRequest) GetCountryId() string {
	if x != nil {
		return x.CountryId
	}
	return ""
}

type Map struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regions []*Region `protobuf:"bytes,1,rep,name=regions,proto3" json:"regions,omitempty"`
}

func (x *Map) Reset() {
	*x = Map{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clicks_v1_clicks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Map) ProtoMessage() {}

func (x *Map) ProtoReflect() protoreflect.Message {
	mi := &file_clicks_v1_clicks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Map.ProtoReflect.Descriptor instead.
func (*Map) Descriptor() ([]byte, []int) {
	return file_clicks_v1_clicks_proto_rawDescGZIP(), []int{1}
}

func (x *Map) GetRegions() []*Region {
	if x != nil {
		return x.Regions
	}
	return nil
}

type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epicenter *GeodesicCoordinates `protobuf:"bytes,1,opt,name=epicenter,proto3" json:"epicenter,omitempty"`
	Tiles     []*Tile              `protobuf:"bytes,2,rep,name=tiles,proto3" json:"tiles,omitempty"`
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clicks_v1_clicks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_clicks_v1_clicks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_clicks_v1_clicks_proto_rawDescGZIP(), []int{2}
}

func (x *Region) GetEpicenter() *GeodesicCoordinates {
	if x != nil {
		return x.Epicenter
	}
	return nil
}

func (x *Region) GetTiles() []*Tile {
	if x != nil {
		return x.Tiles
	}
	return nil
}

type Tile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SouthWest *GeodesicCoordinates `protobuf:"bytes,1,opt,name=south_west,json=southWest,proto3" json:"south_west,omitempty"`
	NorthEast *GeodesicCoordinates `protobuf:"bytes,2,opt,name=north_east,json=northEast,proto3" json:"north_east,omitempty"`
	Id        string               `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Tile) Reset() {
	*x = Tile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clicks_v1_clicks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tile) ProtoMessage() {}

func (x *Tile) ProtoReflect() protoreflect.Message {
	mi := &file_clicks_v1_clicks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tile.ProtoReflect.Descriptor instead.
func (*Tile) Descriptor() ([]byte, []int) {
	return file_clicks_v1_clicks_proto_rawDescGZIP(), []int{3}
}

func (x *Tile) GetSouthWest() *GeodesicCoordinates {
	if x != nil {
		return x.SouthWest
	}
	return nil
}

func (x *Tile) GetNorthEast() *GeodesicCoordinates {
	if x != nil {
		return x.NorthEast
	}
	return nil
}

func (x *Tile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GeodesicCoordinates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon float64 `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *GeodesicCoordinates) Reset() {
	*x = GeodesicCoordinates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clicks_v1_clicks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeodesicCoordinates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeodesicCoordinates) ProtoMessage() {}

func (x *GeodesicCoordinates) ProtoReflect() protoreflect.Message {
	mi := &file_clicks_v1_clicks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeodesicCoordinates.ProtoReflect.Descriptor instead.
func (*GeodesicCoordinates) Descriptor() ([]byte, []int) {
	return file_clicks_v1_clicks_proto_rawDescGZIP(), []int{4}
}

func (x *GeodesicCoordinates) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *GeodesicCoordinates) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

type Ownerships struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bindings map[string]string `protobuf:"bytes,1,rep,name=bindings,proto3" json:"bindings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // TODO: add country leaderboard
}

func (x *Ownerships) Reset() {
	*x = Ownerships{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clicks_v1_clicks_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ownerships) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ownerships) ProtoMessage() {}

func (x *Ownerships) ProtoReflect() protoreflect.Message {
	mi := &file_clicks_v1_clicks_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ownerships.ProtoReflect.Descriptor instead.
func (*Ownerships) Descriptor() ([]byte, []int) {
	return file_clicks_v1_clicks_proto_rawDescGZIP(), []int{5}
}

func (x *Ownerships) GetBindings() map[string]string {
	if x != nil {
		return x.Bindings
	}
	return nil
}

type TileUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TileId    string `protobuf:"bytes,1,opt,name=tile_id,json=tileId,proto3" json:"tile_id,omitempty"`
	CountryId string `protobuf:"bytes,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
}

func (x *TileUpdate) Reset() {
	*x = TileUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clicks_v1_clicks_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TileUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TileUpdate) ProtoMessage() {}

func (x *TileUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_clicks_v1_clicks_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TileUpdate.ProtoReflect.Descriptor instead.
func (*TileUpdate) Descriptor() ([]byte, []int) {
	return file_clicks_v1_clicks_proto_rawDescGZIP(), []int{6}
}

func (x *TileUpdate) GetTileId() string {
	if x != nil {
		return x.TileId
	}
	return ""
}

func (x *TileUpdate) GetCountryId() string {
	if x != nil {
		return x.CountryId
	}
	return ""
}

var File_clicks_v1_clicks_proto protoreflect.FileDescriptor

var file_clicks_v1_clicks_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x63,
	0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73,
	0x2e, 0x76, 0x31, 0x22, 0x46, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x03, 0x4d,
	0x61, 0x70, 0x12, 0x2b, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x6d, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x09, 0x65, 0x70, 0x69,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63,
	0x6c, 0x69, 0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x64, 0x65, 0x73, 0x69,
	0x63, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x09, 0x65, 0x70,
	0x69, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x74, 0x69, 0x6c, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x94,
	0x01, 0x0a, 0x04, 0x54, 0x69, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x74, 0x68,
	0x5f, 0x77, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6c,
	0x69, 0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x64, 0x65, 0x73, 0x69, 0x63,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x09, 0x73, 0x6f, 0x75,
	0x74, 0x68, 0x57, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0a, 0x6e, 0x6f, 0x72, 0x74, 0x68, 0x5f,
	0x65, 0x61, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6c, 0x69,
	0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x64, 0x65, 0x73, 0x69, 0x63, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x52, 0x09, 0x6e, 0x6f, 0x72, 0x74,
	0x68, 0x45, 0x61, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x13, 0x47, 0x65, 0x6f, 0x64, 0x65, 0x73, 0x69,
	0x63, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e,
	0x22, 0x8a, 0x01, 0x0a, 0x0a, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12,
	0x3f, 0x0a, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x1a, 0x3b, 0x0a, 0x0d, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x44, 0x0a,
	0x0a, 0x54, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x49, 0x64, 0x42, 0xb3, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x63,
	0x6b, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x63,
	0x6b, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x2e, 0x6c, 0x6f, 0x6c, 0x2d, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6c,
	0x69, 0x63, 0x6b, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x43,
	0x6c, 0x69, 0x63, 0x6b, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x43, 0x6c, 0x69, 0x63, 0x6b,
	0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x43,
	0x6c, 0x69, 0x63, 0x6b, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_clicks_v1_clicks_proto_rawDescOnce sync.Once
	file_clicks_v1_clicks_proto_rawDescData = file_clicks_v1_clicks_proto_rawDesc
)

func file_clicks_v1_clicks_proto_rawDescGZIP() []byte {
	file_clicks_v1_clicks_proto_rawDescOnce.Do(func() {
		file_clicks_v1_clicks_proto_rawDescData = protoimpl.X.CompressGZIP(file_clicks_v1_clicks_proto_rawDescData)
	})
	return file_clicks_v1_clicks_proto_rawDescData
}

var file_clicks_v1_clicks_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_clicks_v1_clicks_proto_goTypes = []interface{}{
	(*ClickRequest)(nil),        // 0: clicks.v1.ClickRequest
	(*Map)(nil),                 // 1: clicks.v1.Map
	(*Region)(nil),              // 2: clicks.v1.Region
	(*Tile)(nil),                // 3: clicks.v1.Tile
	(*GeodesicCoordinates)(nil), // 4: clicks.v1.GeodesicCoordinates
	(*Ownerships)(nil),          // 5: clicks.v1.Ownerships
	(*TileUpdate)(nil),          // 6: clicks.v1.TileUpdate
	nil,                         // 7: clicks.v1.Ownerships.BindingsEntry
}
var file_clicks_v1_clicks_proto_depIdxs = []int32{
	2, // 0: clicks.v1.Map.regions:type_name -> clicks.v1.Region
	4, // 1: clicks.v1.Region.epicenter:type_name -> clicks.v1.GeodesicCoordinates
	3, // 2: clicks.v1.Region.tiles:type_name -> clicks.v1.Tile
	4, // 3: clicks.v1.Tile.south_west:type_name -> clicks.v1.GeodesicCoordinates
	4, // 4: clicks.v1.Tile.north_east:type_name -> clicks.v1.GeodesicCoordinates
	7, // 5: clicks.v1.Ownerships.bindings:type_name -> clicks.v1.Ownerships.BindingsEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_clicks_v1_clicks_proto_init() }
func file_clicks_v1_clicks_proto_init() {
	if File_clicks_v1_clicks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clicks_v1_clicks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClickRequest); i {
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
		file_clicks_v1_clicks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Map); i {
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
		file_clicks_v1_clicks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
		file_clicks_v1_clicks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tile); i {
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
		file_clicks_v1_clicks_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeodesicCoordinates); i {
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
		file_clicks_v1_clicks_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ownerships); i {
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
		file_clicks_v1_clicks_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TileUpdate); i {
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
			RawDescriptor: file_clicks_v1_clicks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_clicks_v1_clicks_proto_goTypes,
		DependencyIndexes: file_clicks_v1_clicks_proto_depIdxs,
		MessageInfos:      file_clicks_v1_clicks_proto_msgTypes,
	}.Build()
	File_clicks_v1_clicks_proto = out.File
	file_clicks_v1_clicks_proto_rawDesc = nil
	file_clicks_v1_clicks_proto_goTypes = nil
	file_clicks_v1_clicks_proto_depIdxs = nil
}
