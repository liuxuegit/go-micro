// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go-micro/network/router/proto/router.proto

package go_micro_router

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// AdvertType defines the type of advert
type AdvertType int32

const (
	AdvertType_AdvertAnnounce AdvertType = 0
	AdvertType_AdvertUpdate   AdvertType = 1
)

var AdvertType_name = map[int32]string{
	0: "AdvertAnnounce",
	1: "AdvertUpdate",
}

var AdvertType_value = map[string]int32{
	"AdvertAnnounce": 0,
	"AdvertUpdate":   1,
}

func (x AdvertType) String() string {
	return proto.EnumName(AdvertType_name, int32(x))
}

func (AdvertType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc08514fc6dadd29, []int{0}
}

// EventType defines the type of event
type EventType int32

const (
	EventType_Create EventType = 0
	EventType_Delete EventType = 1
	EventType_Update EventType = 2
)

var EventType_name = map[int32]string{
	0: "Create",
	1: "Delete",
	2: "Update",
}

var EventType_value = map[string]int32{
	"Create": 0,
	"Delete": 1,
	"Update": 2,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc08514fc6dadd29, []int{1}
}

// LookupRequest is made to Lookup
type LookupRequest struct {
	Query                *Query   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupRequest) Reset()         { *m = LookupRequest{} }
func (m *LookupRequest) String() string { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()    {}
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc08514fc6dadd29, []int{0}
}

func (m *LookupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupRequest.Unmarshal(m, b)
}
func (m *LookupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupRequest.Marshal(b, m, deterministic)
}
func (m *LookupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupRequest.Merge(m, src)
}
func (m *LookupRequest) XXX_Size() int {
	return xxx_messageInfo_LookupRequest.Size(m)
}
func (m *LookupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookupRequest proto.InternalMessageInfo

func (m *LookupRequest) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

// LookupResponse is returned by Lookup
type LookupResponse struct {
	Routes               []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupResponse) Reset()         { *m = LookupResponse{} }
func (m *LookupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()    {}
func (*LookupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc08514fc6dadd29, []int{1}
}

func (m *LookupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupResponse.Unmarshal(m, b)
}
func (m *LookupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupResponse.Marshal(b, m, deterministic)
}
func (m *LookupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupResponse.Merge(m, src)
}
func (m *LookupResponse) XXX_Size() int {
	return xxx_messageInfo_LookupResponse.Size(m)
}
func (m *LookupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LookupResponse proto.InternalMessageInfo

func (m *LookupResponse) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

// WatchRequest is made to Watch Router
type WatchRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc08514fc6dadd29, []int{2}
}

func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchRequest.Unmarshal(m, b)
}
func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
}
func (m *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(m, src)
}
func (m *WatchRequest) XXX_Size() int {
	return xxx_messageInfo_WatchRequest.Size(m)
}
func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

// AdvertiseRequest request a stream of Adverts
type AdvertiseRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdvertiseRequest) Reset()         { *m = AdvertiseRequest{} }
func (m *AdvertiseRequest) String() string { return proto.CompactTextString(m) }
func (*AdvertiseRequest) ProtoMessage()    {}
func (*AdvertiseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc08514fc6dadd29, []int{3}
}

func (m *AdvertiseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdvertiseRequest.Unmarshal(m, b)
}
func (m *AdvertiseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdvertiseRequest.Marshal(b, m, deterministic)
}
func (m *AdvertiseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdvertiseRequest.Merge(m, src)
}
func (m *AdvertiseRequest) XXX_Size() int {
	return xxx_messageInfo_AdvertiseRequest.Size(m)
}
func (m *AdvertiseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AdvertiseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AdvertiseRequest proto.InternalMessageInfo

// Advert is router advertsement streamed by Watch
type Advert struct {
	// id of the advertising router
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// type of advertisement
	Type AdvertType `protobuf:"varint,2,opt,name=type,proto3,enum=go.micro.router.AdvertType" json:"type,omitempty"`
	// unix timestamp of the advertisement
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// TTL of the Advert
	Ttl int64 `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// events is a list of advertised events
	Events               []*Event `protobuf:"bytes,5,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Advert) Reset()         { *m = Advert{} }
func (m *Advert) String() string { return proto.CompactTextString(m) }
func (*Advert) ProtoMessage()    {}
func (*Advert) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc08514fc6dadd29, []int{4}
}

func (m *Advert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Advert.Unmarshal(m, b)
}
func (m *Advert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Advert.Marshal(b, m, deterministic)
}
func (m *Advert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Advert.Merge(m, src)
}
func (m *Advert) XXX_Size() int {
	return xxx_messageInfo_Advert.Size(m)
}
func (m *Advert) XXX_DiscardUnknown() {
	xxx_messageInfo_Advert.DiscardUnknown(m)
}

var xxx_messageInfo_Advert proto.InternalMessageInfo

func (m *Advert) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Advert) GetType() AdvertType {
	if m != nil {
		return m.Type
	}
	return AdvertType_AdvertAnnounce
}

func (m *Advert) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Advert) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *Advert) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

// ProcessResponse is returned by Process
type ProcessResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessResponse) Reset()         { *m = ProcessResponse{} }
func (m *ProcessResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessResponse) ProtoMessage()    {}
func (*ProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc08514fc6dadd29, []int{5}
}

func (m *ProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessResponse.Unmarshal(m, b)
}
func (m *ProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessResponse.Marshal(b, m, deterministic)
}
func (m *ProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessResponse.Merge(m, src)
}
func (m *ProcessResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessResponse.Size(m)
}
func (m *ProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessResponse proto.InternalMessageInfo

// Event is routing table event
type Event struct {
	// type of event
	Type EventType `protobuf:"varint,1,opt,name=type,proto3,enum=go.micro.router.EventType" json:"type,omitempty"`
	// unix timestamp of event
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// service route
	Route                *Route   `protobuf:"bytes,3,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc08514fc6dadd29, []int{6}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_Create
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetRoute() *Route {
	if m != nil {
		return m.Route
	}
	return nil
}

// ListRequest is made to List routes
type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc08514fc6dadd29, []int{7}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

// ListResponse is returned by List
type ListResponse struct {
	Routes               []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc08514fc6dadd29, []int{8}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

// Query is passed in a LookupRequest
type Query struct {
	// service to lookup
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc08514fc6dadd29, []int{9}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

// Route is a service route
type Route struct {
	// service for the route
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// the address that advertise this route
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// gateway as the next hop
	Gateway string `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	// the network for this destination
	Network string `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	// the network link
	Link string `protobuf:"bytes,5,opt,name=link,proto3" json:"link,omitempty"`
	// the metric / score of this route
	Metric               int64    `protobuf:"varint,6,opt,name=metric,proto3" json:"metric,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc08514fc6dadd29, []int{10}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Route) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Route) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Route) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Route) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Route) GetMetric() int64 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func init() {
	proto.RegisterEnum("go.micro.router.AdvertType", AdvertType_name, AdvertType_value)
	proto.RegisterEnum("go.micro.router.EventType", EventType_name, EventType_value)
	proto.RegisterType((*LookupRequest)(nil), "go.micro.router.LookupRequest")
	proto.RegisterType((*LookupResponse)(nil), "go.micro.router.LookupResponse")
	proto.RegisterType((*WatchRequest)(nil), "go.micro.router.WatchRequest")
	proto.RegisterType((*AdvertiseRequest)(nil), "go.micro.router.AdvertiseRequest")
	proto.RegisterType((*Advert)(nil), "go.micro.router.Advert")
	proto.RegisterType((*ProcessResponse)(nil), "go.micro.router.ProcessResponse")
	proto.RegisterType((*Event)(nil), "go.micro.router.Event")
	proto.RegisterType((*ListRequest)(nil), "go.micro.router.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.router.ListResponse")
	proto.RegisterType((*Query)(nil), "go.micro.router.Query")
	proto.RegisterType((*Route)(nil), "go.micro.router.Route")
}

func init() {
	proto.RegisterFile("go-micro/network/router/proto/router.proto", fileDescriptor_fc08514fc6dadd29)
}

var fileDescriptor_fc08514fc6dadd29 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x9d, 0xd8, 0x95, 0xa7, 0x6d, 0x1a, 0xe6, 0x50, 0x2c, 0xd3, 0x42, 0xea, 0x53, 0x55,
	0x15, 0x07, 0x85, 0x33, 0x88, 0x02, 0xe5, 0xd2, 0x1e, 0xc0, 0x02, 0x71, 0x36, 0xf6, 0x28, 0x58,
	0x49, 0xbc, 0xee, 0xee, 0x26, 0x55, 0xce, 0x7c, 0x06, 0x5f, 0xc0, 0x07, 0x72, 0x47, 0x3b, 0xb6,
	0x13, 0x48, 0xea, 0x0b, 0xa7, 0xec, 0x9b, 0xf7, 0x66, 0x3d, 0x33, 0x3b, 0x2f, 0x70, 0x31, 0x11,
	0xcf, 0xe7, 0x79, 0x2a, 0xc5, 0xa8, 0x20, 0x7d, 0x2f, 0xe4, 0x74, 0x24, 0xc5, 0x42, 0x93, 0x1c,
	0x95, 0x52, 0x68, 0x51, 0x83, 0x88, 0x01, 0x1e, 0x4d, 0x44, 0xc4, 0xda, 0xa8, 0x0a, 0x87, 0xaf,
	0xe0, 0xf0, 0x56, 0x88, 0xe9, 0xa2, 0x8c, 0xe9, 0x6e, 0x41, 0x4a, 0xe3, 0x25, 0x38, 0x77, 0x0b,
	0x92, 0x2b, 0xdf, 0x1a, 0x5a, 0xe7, 0xfb, 0xe3, 0xe3, 0x68, 0x2b, 0x23, 0xfa, 0x64, 0xd8, 0xb8,
	0x12, 0x85, 0x6f, 0xa0, 0xdf, 0xa4, 0xab, 0x52, 0x14, 0x8a, 0x30, 0x02, 0x97, 0x85, 0xca, 0xb7,
	0x86, 0xdd, 0x07, 0x2f, 0x88, 0xcd, 0x4f, 0x5c, 0xab, 0xc2, 0x3e, 0x1c, 0x7c, 0x4d, 0x74, 0xfa,
	0xbd, 0xfe, 0x7e, 0x88, 0x30, 0xb8, 0xca, 0x96, 0x24, 0x75, 0xae, 0xa8, 0x89, 0xfd, 0xb2, 0xc0,
	0xad, 0x82, 0xd8, 0x07, 0x3b, 0xcf, 0xb8, 0x36, 0x2f, 0xb6, 0xf3, 0x0c, 0x47, 0xd0, 0xd3, 0xab,
	0x92, 0x7c, 0x7b, 0x68, 0x9d, 0xf7, 0xc7, 0x4f, 0x76, 0x3e, 0x56, 0xa5, 0x7d, 0x5e, 0x95, 0x14,
	0xb3, 0x10, 0x4f, 0xc0, 0xd3, 0xf9, 0x9c, 0x94, 0x4e, 0xe6, 0xa5, 0xdf, 0x1d, 0x5a, 0xe7, 0xdd,
	0x78, 0x13, 0xc0, 0x01, 0x74, 0xb5, 0x9e, 0xf9, 0x3d, 0x8e, 0x9b, 0xa3, 0xe9, 0x87, 0x96, 0x54,
	0x68, 0xe5, 0x3b, 0x2d, 0xfd, 0x5c, 0x1b, 0x3a, 0xae, 0x55, 0xe1, 0x23, 0x38, 0xfa, 0x28, 0x45,
	0x4a, 0x4a, 0x35, 0x23, 0x09, 0x7f, 0x58, 0xe0, 0xb0, 0x08, 0xa3, 0xba, 0x5a, 0x8b, 0xab, 0x0d,
	0x1e, 0xbe, 0xaa, 0xad, 0x58, 0x7b, 0xbb, 0xd8, 0x4b, 0x70, 0x38, 0x8f, 0xdb, 0x68, 0x9f, 0x74,
	0x25, 0x0a, 0x0f, 0x61, 0xff, 0x36, 0x57, 0xba, 0x99, 0xe9, 0x6b, 0x38, 0xa8, 0xe0, 0x7f, 0xbe,
	0xdb, 0x19, 0x38, 0xbc, 0x09, 0xe8, 0xc3, 0x9e, 0x22, 0xb9, 0xcc, 0x53, 0xaa, 0x9f, 0xa5, 0x81,
	0xe1, 0x4f, 0x0b, 0x1c, 0x4e, 0x6a, 0xd7, 0x18, 0x26, 0xc9, 0x32, 0x49, 0x4a, 0x71, 0x7f, 0x5e,
	0xdc, 0x40, 0xc3, 0x4c, 0x12, 0x4d, 0xf7, 0xc9, 0x8a, 0xfb, 0xf3, 0xe2, 0x06, 0x1a, 0xa6, 0xde,
	0x74, 0x7e, 0x28, 0x2f, 0x6e, 0x20, 0x22, 0xf4, 0x66, 0x79, 0x31, 0xf5, 0x1d, 0x0e, 0xf3, 0x19,
	0x8f, 0xc1, 0x9d, 0x93, 0x96, 0x79, 0xea, 0xbb, 0x3c, 0xc0, 0x1a, 0x5d, 0x8c, 0x01, 0x36, 0xcb,
	0x81, 0x08, 0xfd, 0x0a, 0x5d, 0x15, 0x85, 0x58, 0x14, 0x29, 0x0d, 0x3a, 0x38, 0x80, 0x83, 0x2a,
	0xf6, 0xa5, 0xcc, 0x12, 0x4d, 0x03, 0xeb, 0x62, 0x04, 0xde, 0xfa, 0x89, 0x10, 0xc0, 0x7d, 0x27,
	0xc9, 0x10, 0x1d, 0x73, 0x7e, 0x4f, 0x33, 0x32, 0x22, 0x73, 0xae, 0x13, 0xec, 0xf1, 0x6f, 0x1b,
	0x5c, 0x1e, 0x81, 0xc4, 0xb7, 0xe0, 0xf0, 0xa2, 0xe3, 0xe9, 0xce, 0x64, 0xff, 0x36, 0x40, 0xd0,
	0xb2, 0x60, 0x61, 0xe7, 0x85, 0x85, 0x37, 0xe0, 0x56, 0x76, 0xc3, 0xa7, 0x3b, 0xaa, 0x7f, 0x6c,
	0x1c, 0x3c, 0x6b, 0xe5, 0xeb, 0xa5, 0xec, 0xe0, 0x35, 0xf4, 0xcc, 0x06, 0xe0, 0xc9, 0xae, 0x74,
	0xb3, 0x27, 0xc1, 0x69, 0x0b, 0xbb, 0xbe, 0xe6, 0x06, 0xbc, 0xb5, 0x61, 0xf1, 0xac, 0xc5, 0x80,
	0x1b, 0x33, 0x07, 0x8f, 0x5b, 0x24, 0xdc, 0xe0, 0x07, 0xd8, 0xab, 0xdd, 0x83, 0x6d, 0xba, 0x60,
	0xb8, 0x43, 0x6c, 0x1b, 0xae, 0xf3, 0xcd, 0xe5, 0xbf, 0xbb, 0x97, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x99, 0x8e, 0xb9, 0x97, 0x1c, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Router_WatchClient, error)
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Advertise(ctx context.Context, in *AdvertiseRequest, opts ...grpc.CallOption) (Router_AdvertiseClient, error)
	Process(ctx context.Context, in *Advert, opts ...grpc.CallOption) (*ProcessResponse, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Router_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Router_serviceDesc.Streams[0], "/go.micro.router.Router/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &routerWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Router_WatchClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type routerWatchClient struct {
	grpc.ClientStream
}

func (x *routerWatchClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routerClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := c.cc.Invoke(ctx, "/go.micro.router.Router/Lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/go.micro.router.Router/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Advertise(ctx context.Context, in *AdvertiseRequest, opts ...grpc.CallOption) (Router_AdvertiseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Router_serviceDesc.Streams[1], "/go.micro.router.Router/Advertise", opts...)
	if err != nil {
		return nil, err
	}
	x := &routerAdvertiseClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Router_AdvertiseClient interface {
	Recv() (*Advert, error)
	grpc.ClientStream
}

type routerAdvertiseClient struct {
	grpc.ClientStream
}

func (x *routerAdvertiseClient) Recv() (*Advert, error) {
	m := new(Advert)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routerClient) Process(ctx context.Context, in *Advert, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, "/go.micro.router.Router/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	Watch(*WatchRequest, Router_WatchServer) error
	Lookup(context.Context, *LookupRequest) (*LookupResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Advertise(*AdvertiseRequest, Router_AdvertiseServer) error
	Process(context.Context, *Advert) (*ProcessResponse, error)
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouterServer).Watch(m, &routerWatchServer{stream})
}

type Router_WatchServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type routerWatchServer struct {
	grpc.ServerStream
}

func (x *routerWatchServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Router_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.router.Router/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.router.Router/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Advertise_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AdvertiseRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouterServer).Advertise(m, &routerAdvertiseServer{stream})
}

type Router_AdvertiseServer interface {
	Send(*Advert) error
	grpc.ServerStream
}

type routerAdvertiseServer struct {
	grpc.ServerStream
}

func (x *routerAdvertiseServer) Send(m *Advert) error {
	return x.ServerStream.SendMsg(m)
}

func _Router_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Advert)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.router.Router/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Process(ctx, req.(*Advert))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.router.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lookup",
			Handler:    _Router_Lookup_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Router_List_Handler,
		},
		{
			MethodName: "Process",
			Handler:    _Router_Process_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Router_Watch_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Advertise",
			Handler:       _Router_Advertise_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "go-micro/network/router/proto/router.proto",
}