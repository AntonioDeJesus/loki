// FIXME Copied out of the loki project because of some import issues which caused this project to import all of cortex

package reader

import (
	"context"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Direction int32

const (
	FORWARD  Direction = 0
	BACKWARD Direction = 1
)

var Direction_name = map[int32]string{
	0: "FORWARD",
	1: "BACKWARD",
}

var Direction_value = map[string]int32{
	"FORWARD":  0,
	"BACKWARD": 1,
}

func (Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7a8976f235a02f79, []int{0}
}

type PushRequest struct {
	Streams []*Stream `protobuf:"bytes,1,rep,name=streams,proto3" json:"streams"`
}

func (m *PushRequest) Reset()      { *m = PushRequest{} }
func (*PushRequest) ProtoMessage() {}
func (*PushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8976f235a02f79, []int{0}
}
func (m *PushRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRequest.Merge(m, src)
}
func (m *PushRequest) XXX_Size() int {
	return m.Size()
}
func (m *PushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushRequest proto.InternalMessageInfo

func (m *PushRequest) GetStreams() []*Stream {
	if m != nil {
		return m.Streams
	}
	return nil
}

type PushResponse struct {
}

func (m *PushResponse) Reset()      { *m = PushResponse{} }
func (*PushResponse) ProtoMessage() {}
func (*PushResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8976f235a02f79, []int{1}
}
func (m *PushResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushResponse.Merge(m, src)
}
func (m *PushResponse) XXX_Size() int {
	return m.Size()
}
func (m *PushResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushResponse proto.InternalMessageInfo

type QueryRequest struct {
	Query     string    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Limit     uint32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Start     time.Time `protobuf:"bytes,3,opt,name=start,proto3,stdtime" json:"start"`
	End       time.Time `protobuf:"bytes,4,opt,name=end,proto3,stdtime" json:"end"`
	Direction Direction `protobuf:"varint,5,opt,name=direction,proto3,enum=logproto.Direction" json:"direction,omitempty"`
	Regex     string    `protobuf:"bytes,6,opt,name=regex,proto3" json:"regex,omitempty"`
}

func (m *QueryRequest) Reset()      { *m = QueryRequest{} }
func (*QueryRequest) ProtoMessage() {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8976f235a02f79, []int{2}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *QueryRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryRequest) GetStart() time.Time {
	if m != nil {
		return m.Start
	}
	return time.Time{}
}

func (m *QueryRequest) GetEnd() time.Time {
	if m != nil {
		return m.End
	}
	return time.Time{}
}

func (m *QueryRequest) GetDirection() Direction {
	if m != nil {
		return m.Direction
	}
	return FORWARD
}

func (m *QueryRequest) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

type QueryResponse struct {
	Streams []*Stream `protobuf:"bytes,1,rep,name=streams,proto3" json:"streams,omitempty"`
}

func (m *QueryResponse) Reset()      { *m = QueryResponse{} }
func (*QueryResponse) ProtoMessage() {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8976f235a02f79, []int{3}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetStreams() []*Stream {
	if m != nil {
		return m.Streams
	}
	return nil
}

type LabelRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Values bool   `protobuf:"varint,2,opt,name=values,proto3" json:"values,omitempty"`
}

func (m *LabelRequest) Reset()      { *m = LabelRequest{} }
func (*LabelRequest) ProtoMessage() {}
func (*LabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8976f235a02f79, []int{4}
}
func (m *LabelRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LabelRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelRequest.Merge(m, src)
}
func (m *LabelRequest) XXX_Size() int {
	return m.Size()
}
func (m *LabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LabelRequest proto.InternalMessageInfo

func (m *LabelRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LabelRequest) GetValues() bool {
	if m != nil {
		return m.Values
	}
	return false
}

type LabelResponse struct {
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (m *LabelResponse) Reset()      { *m = LabelResponse{} }
func (*LabelResponse) ProtoMessage() {}
func (*LabelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8976f235a02f79, []int{5}
}
func (m *LabelResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LabelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LabelResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LabelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelResponse.Merge(m, src)
}
func (m *LabelResponse) XXX_Size() int {
	return m.Size()
}
func (m *LabelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LabelResponse proto.InternalMessageInfo

func (m *LabelResponse) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Stream struct {
	Labels  string  `protobuf:"bytes,1,opt,name=labels,proto3" json:"labels"`
	Entries []Entry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries"`
}

func (m *Stream) Reset()      { *m = Stream{} }
func (*Stream) ProtoMessage() {}
func (*Stream) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8976f235a02f79, []int{6}
}
func (m *Stream) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stream.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stream.Merge(m, src)
}
func (m *Stream) XXX_Size() int {
	return m.Size()
}
func (m *Stream) XXX_DiscardUnknown() {
	xxx_messageInfo_Stream.DiscardUnknown(m)
}

var xxx_messageInfo_Stream proto.InternalMessageInfo

func (m *Stream) GetLabels() string {
	if m != nil {
		return m.Labels
	}
	return ""
}

func (m *Stream) GetEntries() []Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Entry struct {
	Timestamp time.Time `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"ts"`
	Line      string    `protobuf:"bytes,2,opt,name=line,proto3" json:"line"`
}

func (m *Entry) Reset()      { *m = Entry{} }
func (*Entry) ProtoMessage() {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8976f235a02f79, []int{7}
}
func (m *Entry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return m.Size()
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *Entry) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func init() {
	proto.RegisterEnum("logproto.Direction", Direction_name, Direction_value)
	proto.RegisterType((*PushRequest)(nil), "logproto.PushRequest")
	proto.RegisterType((*PushResponse)(nil), "logproto.PushResponse")
	proto.RegisterType((*QueryRequest)(nil), "logproto.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "logproto.QueryResponse")
	proto.RegisterType((*LabelRequest)(nil), "logproto.LabelRequest")
	proto.RegisterType((*LabelResponse)(nil), "logproto.LabelResponse")
	proto.RegisterType((*Stream)(nil), "logproto.Stream")
	proto.RegisterType((*Entry)(nil), "logproto.Entry")
}

func init() { proto.RegisterFile("logproto.proto", fileDescriptor_7a8976f235a02f79) }

var fileDescriptor_7a8976f235a02f79 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xf5, 0xb6, 0x89, 0x13, 0x4f, 0xd2, 0xb4, 0xda, 0xdf, 0x8f, 0x62, 0x45, 0x68, 0x1d, 0xf9,
	0x00, 0x51, 0x25, 0x5c, 0x08, 0x88, 0x4a, 0x85, 0x4b, 0x4d, 0xa9, 0x90, 0x40, 0x02, 0x16, 0x24,
	0xce, 0x4e, 0xbb, 0xb8, 0x96, 0xfc, 0xa7, 0xb5, 0xd7, 0x88, 0xde, 0x90, 0xf8, 0x02, 0xfd, 0x18,
	0x7c, 0x94, 0x1e, 0x73, 0xec, 0x29, 0x10, 0xe7, 0x82, 0x72, 0xea, 0x8d, 0x2b, 0xda, 0xb5, 0x1d,
	0x1b, 0x90, 0x40, 0x5c, 0x9c, 0x79, 0xbb, 0xef, 0xcd, 0xec, 0x9b, 0x99, 0x40, 0xcf, 0x8f, 0xdc,
	0x93, 0x38, 0xe2, 0x91, 0x25, 0xbf, 0xb8, 0x5d, 0xe2, 0xbe, 0xe1, 0x46, 0x91, 0xeb, 0xb3, 0x6d,
	0x89, 0xc6, 0xe9, 0xbb, 0x6d, 0xee, 0x05, 0x2c, 0xe1, 0x4e, 0x70, 0x92, 0x53, 0xfb, 0xb7, 0x5d,
	0x8f, 0x1f, 0xa7, 0x63, 0xeb, 0x30, 0x0a, 0xb6, 0xdd, 0xc8, 0x8d, 0x2a, 0xa6, 0x40, 0x12, 0xc8,
	0x28, 0xa7, 0x9b, 0x07, 0xd0, 0x79, 0x99, 0x26, 0xc7, 0x94, 0x9d, 0xa6, 0x2c, 0xe1, 0x78, 0x07,
	0x5a, 0x09, 0x8f, 0x99, 0x13, 0x24, 0x3a, 0x1a, 0xac, 0x0e, 0x3b, 0xa3, 0x0d, 0x6b, 0xf9, 0x94,
	0xd7, 0xf2, 0xc2, 0xee, 0x2c, 0xa6, 0x46, 0x49, 0xa2, 0x65, 0x60, 0xf6, 0xa0, 0x9b, 0xe7, 0x49,
	0x4e, 0xa2, 0x30, 0x61, 0xe6, 0x77, 0x04, 0xdd, 0x57, 0x29, 0x8b, 0xcf, 0xca, 0xcc, 0xff, 0x43,
	0xf3, 0x54, 0x60, 0x1d, 0x0d, 0xd0, 0x50, 0xa3, 0x39, 0x10, 0xa7, 0xbe, 0x17, 0x78, 0x5c, 0x5f,
	0x19, 0xa0, 0xe1, 0x1a, 0xcd, 0x01, 0xde, 0x85, 0x66, 0xc2, 0x9d, 0x98, 0xeb, 0xab, 0x03, 0x34,
	0xec, 0x8c, 0xfa, 0x56, 0x6e, 0xda, 0x2a, 0xad, 0x58, 0x6f, 0x4a, 0xd3, 0x76, 0xfb, 0x62, 0x6a,
	0x28, 0xe7, 0x5f, 0x0c, 0x44, 0x73, 0x09, 0x7e, 0x00, 0xab, 0x2c, 0x3c, 0xd2, 0x1b, 0xff, 0xa0,
	0x14, 0x02, 0x7c, 0x17, 0xb4, 0x23, 0x2f, 0x66, 0x87, 0xdc, 0x8b, 0x42, 0xbd, 0x39, 0x40, 0xc3,
	0xde, 0xe8, 0xbf, 0xca, 0xfb, 0x7e, 0x79, 0x45, 0x2b, 0x96, 0x78, 0x7c, 0xcc, 0x5c, 0xf6, 0x41,
	0x57, 0x73, 0x4b, 0x12, 0x98, 0x0f, 0x61, 0xad, 0x30, 0x9e, 0xb7, 0x02, 0x6f, 0xfd, 0xb5, 0xa7,
	0x55, 0x1b, 0x77, 0xa1, 0xfb, 0xdc, 0x19, 0x33, 0xbf, 0xec, 0x1a, 0x86, 0x46, 0xe8, 0x04, 0xac,
	0x68, 0x9a, 0x8c, 0xf1, 0x26, 0xa8, 0xef, 0x1d, 0x3f, 0x65, 0x89, 0x6c, 0x5a, 0x9b, 0x16, 0xc8,
	0xbc, 0x05, 0x6b, 0x85, 0xb6, 0x28, 0x5c, 0x11, 0x45, 0x5d, 0x6d, 0x49, 0x3c, 0x06, 0x35, 0xaf,
	0x8b, 0x4d, 0x50, 0x7d, 0x21, 0x49, 0xf2, 0x02, 0x36, 0x2c, 0xa6, 0x46, 0x71, 0x42, 0x8b, 0x5f,
	0xbc, 0x0b, 0x2d, 0x16, 0xf2, 0xd8, 0x93, 0xf5, 0xc4, 0xf3, 0xd7, 0xab, 0xe7, 0x3f, 0x09, 0x79,
	0x7c, 0x66, 0xaf, 0x8b, 0x4e, 0x8a, 0xad, 0x28, 0x78, 0xb4, 0x0c, 0xcc, 0x08, 0x9a, 0x92, 0x82,
	0x9f, 0x82, 0xb6, 0x5c, 0x54, 0x59, 0xeb, 0xcf, 0xb3, 0xe9, 0x15, 0x19, 0x57, 0x78, 0x22, 0x27,
	0x54, 0x89, 0xf1, 0x0d, 0x68, 0xf8, 0x5e, 0xc8, 0xa4, 0x77, 0xcd, 0x6e, 0x2f, 0xa6, 0x86, 0xc4,
	0x54, 0x7e, 0xb7, 0x6e, 0x82, 0xb6, 0x1c, 0x15, 0xee, 0x40, 0xeb, 0xe0, 0x05, 0x7d, 0xbb, 0x47,
	0xf7, 0x37, 0x14, 0xdc, 0x85, 0xb6, 0xbd, 0xf7, 0xf8, 0x99, 0x44, 0x68, 0xb4, 0x07, 0xaa, 0x58,
	0x57, 0x16, 0xe3, 0x1d, 0x68, 0x88, 0x08, 0x5f, 0xab, 0x5c, 0xd5, 0xfe, 0x10, 0xfd, 0xcd, 0x5f,
	0x8f, 0x8b, 0xfd, 0x56, 0x46, 0x9f, 0x10, 0xb4, 0xc4, 0xa0, 0x3d, 0x16, 0xe3, 0x47, 0xd0, 0x94,
	0x33, 0xc7, 0x35, 0x7a, 0x7d, 0xfb, 0xfb, 0xd7, 0x7f, 0x3b, 0x2f, 0xf3, 0xdc, 0x41, 0x62, 0xdd,
	0xe5, 0xe0, 0xea, 0xea, 0xfa, 0x16, 0xd4, 0xd5, 0x3f, 0x4d, 0xd8, 0x54, 0xec, 0xfb, 0x93, 0x19,
	0x51, 0x2e, 0x67, 0x44, 0xb9, 0x9a, 0x11, 0xf4, 0x31, 0x23, 0xe8, 0x73, 0x46, 0xd0, 0x45, 0x46,
	0xd0, 0x24, 0x23, 0xe8, 0x6b, 0x46, 0xd0, 0xb7, 0x8c, 0x28, 0x57, 0x19, 0x41, 0xe7, 0x73, 0xa2,
	0x4c, 0xe6, 0x44, 0xb9, 0x9c, 0x13, 0x65, 0xac, 0xca, 0x64, 0xf7, 0x7e, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x47, 0x69, 0x1e, 0x88, 0x68, 0x04, 0x00, 0x00,
}

func (x Direction) String() string {
	s, ok := Direction_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *PushRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PushRequest)
	if !ok {
		that2, ok := that.(PushRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Streams) != len(that1.Streams) {
		return false
	}
	for i := range this.Streams {
		if !this.Streams[i].Equal(that1.Streams[i]) {
			return false
		}
	}
	return true
}
func (this *PushResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PushResponse)
	if !ok {
		that2, ok := that.(PushResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *QueryRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryRequest)
	if !ok {
		that2, ok := that.(QueryRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Query != that1.Query {
		return false
	}
	if this.Limit != that1.Limit {
		return false
	}
	if !this.Start.Equal(that1.Start) {
		return false
	}
	if !this.End.Equal(that1.End) {
		return false
	}
	if this.Direction != that1.Direction {
		return false
	}
	if this.Regex != that1.Regex {
		return false
	}
	return true
}
func (this *QueryResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryResponse)
	if !ok {
		that2, ok := that.(QueryResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Streams) != len(that1.Streams) {
		return false
	}
	for i := range this.Streams {
		if !this.Streams[i].Equal(that1.Streams[i]) {
			return false
		}
	}
	return true
}
func (this *LabelRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LabelRequest)
	if !ok {
		that2, ok := that.(LabelRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Values != that1.Values {
		return false
	}
	return true
}
func (this *LabelResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LabelResponse)
	if !ok {
		that2, ok := that.(LabelResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Values) != len(that1.Values) {
		return false
	}
	for i := range this.Values {
		if this.Values[i] != that1.Values[i] {
			return false
		}
	}
	return true
}
func (this *Stream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Stream)
	if !ok {
		that2, ok := that.(Stream)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Labels != that1.Labels {
		return false
	}
	if len(this.Entries) != len(that1.Entries) {
		return false
	}
	for i := range this.Entries {
		if !this.Entries[i].Equal(&that1.Entries[i]) {
			return false
		}
	}
	return true
}
func (this *Entry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Entry)
	if !ok {
		that2, ok := that.(Entry)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Timestamp.Equal(that1.Timestamp) {
		return false
	}
	if this.Line != that1.Line {
		return false
	}
	return true
}
func (this *PushRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&logproto.PushRequest{")
	if this.Streams != nil {
		s = append(s, "Streams: "+fmt.Sprintf("%#v", this.Streams)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PushResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&logproto.PushResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueryRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&logproto.QueryRequest{")
	s = append(s, "Query: "+fmt.Sprintf("%#v", this.Query)+",\n")
	s = append(s, "Limit: "+fmt.Sprintf("%#v", this.Limit)+",\n")
	s = append(s, "Start: "+fmt.Sprintf("%#v", this.Start)+",\n")
	s = append(s, "End: "+fmt.Sprintf("%#v", this.End)+",\n")
	s = append(s, "Direction: "+fmt.Sprintf("%#v", this.Direction)+",\n")
	s = append(s, "Regex: "+fmt.Sprintf("%#v", this.Regex)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueryResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&logproto.QueryResponse{")
	if this.Streams != nil {
		s = append(s, "Streams: "+fmt.Sprintf("%#v", this.Streams)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *LabelRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&logproto.LabelRequest{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Values: "+fmt.Sprintf("%#v", this.Values)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *LabelResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&logproto.LabelResponse{")
	s = append(s, "Values: "+fmt.Sprintf("%#v", this.Values)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Stream) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&logproto.Stream{")
	s = append(s, "Labels: "+fmt.Sprintf("%#v", this.Labels)+",\n")
	if this.Entries != nil {
		vs := make([]*Entry, len(this.Entries))
		for i := range vs {
			vs[i] = &this.Entries[i]
		}
		s = append(s, "Entries: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Entry) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&logproto.Entry{")
	s = append(s, "Timestamp: "+fmt.Sprintf("%#v", this.Timestamp)+",\n")
	s = append(s, "Line: "+fmt.Sprintf("%#v", this.Line)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringLogproto(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PusherClient is the client API for Pusher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PusherClient interface {
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error)
}

type pusherClient struct {
	cc *grpc.ClientConn
}

func NewPusherClient(cc *grpc.ClientConn) PusherClient {
	return &pusherClient{cc}
}

func (c *pusherClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	out := new(PushResponse)
	err := c.cc.Invoke(ctx, "/logproto.Pusher/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PusherServer is the server API for Pusher service.
type PusherServer interface {
	Push(context.Context, *PushRequest) (*PushResponse, error)
}

func RegisterPusherServer(s *grpc.Server, srv PusherServer) {
	s.RegisterService(&_Pusher_serviceDesc, srv)
}

func _Pusher_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PusherServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logproto.Pusher/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PusherServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pusher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logproto.Pusher",
	HandlerType: (*PusherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Pusher_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logproto.proto",
}

// QuerierClient is the client API for Querier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuerierClient interface {
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Querier_QueryClient, error)
	Label(ctx context.Context, in *LabelRequest, opts ...grpc.CallOption) (*LabelResponse, error)
}

type querierClient struct {
	cc *grpc.ClientConn
}

func NewQuerierClient(cc *grpc.ClientConn) QuerierClient {
	return &querierClient{cc}
}

func (c *querierClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Querier_QueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Querier_serviceDesc.Streams[0], "/logproto.Querier/Query", opts...)
	if err != nil {
		return nil, err
	}
	x := &querierQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Querier_QueryClient interface {
	Recv() (*QueryResponse, error)
	grpc.ClientStream
}

type querierQueryClient struct {
	grpc.ClientStream
}

func (x *querierQueryClient) Recv() (*QueryResponse, error) {
	m := new(QueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *querierClient) Label(ctx context.Context, in *LabelRequest, opts ...grpc.CallOption) (*LabelResponse, error) {
	out := new(LabelResponse)
	err := c.cc.Invoke(ctx, "/logproto.Querier/Label", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuerierServer is the server API for Querier service.
type QuerierServer interface {
	Query(*QueryRequest, Querier_QueryServer) error
	Label(context.Context, *LabelRequest) (*LabelResponse, error)
}

func RegisterQuerierServer(s *grpc.Server, srv QuerierServer) {
	s.RegisterService(&_Querier_serviceDesc, srv)
}

func _Querier_Query_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QuerierServer).Query(m, &querierQueryServer{stream})
}

type Querier_QueryServer interface {
	Send(*QueryResponse) error
	grpc.ServerStream
}

type querierQueryServer struct {
	grpc.ServerStream
}

func (x *querierQueryServer) Send(m *QueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Querier_Label_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServer).Label(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logproto.Querier/Label",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServer).Label(ctx, req.(*LabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Querier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logproto.Querier",
	HandlerType: (*QuerierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Label",
			Handler:    _Querier_Label_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Query",
			Handler:       _Querier_Query_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "logproto.proto",
}

func (m *PushRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Streams) > 0 {
		for _, msg := range m.Streams {
			dAtA[i] = 0xa
			i++
			i = encodeVarintLogproto(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *PushResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *QueryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Query) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLogproto(dAtA, i, uint64(len(m.Query)))
		i += copy(dAtA[i:], m.Query)
	}
	if m.Limit != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLogproto(dAtA, i, uint64(m.Limit))
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintLogproto(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.Start)))
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Start, dAtA[i:])
	if err1 != nil {
		return 0, err1
	}
	i += n1
	dAtA[i] = 0x22
	i++
	i = encodeVarintLogproto(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.End)))
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.End, dAtA[i:])
	if err2 != nil {
		return 0, err2
	}
	i += n2
	if m.Direction != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintLogproto(dAtA, i, uint64(m.Direction))
	}
	if len(m.Regex) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintLogproto(dAtA, i, uint64(len(m.Regex)))
		i += copy(dAtA[i:], m.Regex)
	}
	return i, nil
}

func (m *QueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Streams) > 0 {
		for _, msg := range m.Streams {
			dAtA[i] = 0xa
			i++
			i = encodeVarintLogproto(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *LabelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LabelRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLogproto(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Values {
		dAtA[i] = 0x10
		i++
		if m.Values {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *LabelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LabelResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *Stream) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stream) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Labels) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLogproto(dAtA, i, uint64(len(m.Labels)))
		i += copy(dAtA[i:], m.Labels)
	}
	if len(m.Entries) > 0 {
		for _, msg := range m.Entries {
			dAtA[i] = 0x12
			i++
			i = encodeVarintLogproto(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintLogproto(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)))
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i:])
	if err3 != nil {
		return 0, err3
	}
	i += n3
	if len(m.Line) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLogproto(dAtA, i, uint64(len(m.Line)))
		i += copy(dAtA[i:], m.Line)
	}
	return i, nil
}

func encodeVarintLogproto(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PushRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Streams) > 0 {
		for _, e := range m.Streams {
			l = e.Size()
			n += 1 + l + sovLogproto(uint64(l))
		}
	}
	return n
}

func (m *PushResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovLogproto(uint64(l))
	}
	if m.Limit != 0 {
		n += 1 + sovLogproto(uint64(m.Limit))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Start)
	n += 1 + l + sovLogproto(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.End)
	n += 1 + l + sovLogproto(uint64(l))
	if m.Direction != 0 {
		n += 1 + sovLogproto(uint64(m.Direction))
	}
	l = len(m.Regex)
	if l > 0 {
		n += 1 + l + sovLogproto(uint64(l))
	}
	return n
}

func (m *QueryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Streams) > 0 {
		for _, e := range m.Streams {
			l = e.Size()
			n += 1 + l + sovLogproto(uint64(l))
		}
	}
	return n
}

func (m *LabelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovLogproto(uint64(l))
	}
	if m.Values {
		n += 2
	}
	return n
}

func (m *LabelResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovLogproto(uint64(l))
		}
	}
	return n
}

func (m *Stream) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Labels)
	if l > 0 {
		n += 1 + l + sovLogproto(uint64(l))
	}
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovLogproto(uint64(l))
		}
	}
	return n
}

func (m *Entry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovLogproto(uint64(l))
	l = len(m.Line)
	if l > 0 {
		n += 1 + l + sovLogproto(uint64(l))
	}
	return n
}

func sovLogproto(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLogproto(x uint64) (n int) {
	return sovLogproto(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PushRequest) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForStreams := "[]*Stream{"
	for _, f := range this.Streams {
		repeatedStringForStreams += strings.Replace(f.String(), "Stream", "Stream", 1) + ","
	}
	repeatedStringForStreams += "}"
	s := strings.Join([]string{`&PushRequest{`,
		`Streams:` + repeatedStringForStreams + `,`,
		`}`,
	}, "")
	return s
}
func (this *PushResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PushResponse{`,
		`}`,
	}, "")
	return s
}
func (this *QueryRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueryRequest{`,
		`Query:` + fmt.Sprintf("%v", this.Query) + `,`,
		`Limit:` + fmt.Sprintf("%v", this.Limit) + `,`,
		`Start:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Start), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`End:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.End), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`Direction:` + fmt.Sprintf("%v", this.Direction) + `,`,
		`Regex:` + fmt.Sprintf("%v", this.Regex) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QueryResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForStreams := "[]*Stream{"
	for _, f := range this.Streams {
		repeatedStringForStreams += strings.Replace(f.String(), "Stream", "Stream", 1) + ","
	}
	repeatedStringForStreams += "}"
	s := strings.Join([]string{`&QueryResponse{`,
		`Streams:` + repeatedStringForStreams + `,`,
		`}`,
	}, "")
	return s
}
func (this *LabelRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LabelRequest{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Values:` + fmt.Sprintf("%v", this.Values) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LabelResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LabelResponse{`,
		`Values:` + fmt.Sprintf("%v", this.Values) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Stream) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForEntries := "[]Entry{"
	for _, f := range this.Entries {
		repeatedStringForEntries += strings.Replace(strings.Replace(f.String(), "Entry", "Entry", 1), `&`, ``, 1) + ","
	}
	repeatedStringForEntries += "}"
	s := strings.Join([]string{`&Stream{`,
		`Labels:` + fmt.Sprintf("%v", this.Labels) + `,`,
		`Entries:` + repeatedStringForEntries + `,`,
		`}`,
	}, "")
	return s
}
func (this *Entry) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Entry{`,
		`Timestamp:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Timestamp), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`Line:` + fmt.Sprintf("%v", this.Line) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringLogproto(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PushRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogproto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PushRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Streams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Streams = append(m.Streams, &Stream{})
			if err := m.Streams[len(m.Streams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PushResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogproto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PushResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLogproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogproto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Start, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.End, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direction", wireType)
			}
			m.Direction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Direction |= Direction(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Regex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Regex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogproto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Streams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Streams = append(m.Streams, &Stream{})
			if err := m.Streams[len(m.Streams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LabelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogproto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LabelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LabelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Values = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipLogproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LabelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogproto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LabelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LabelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Stream) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogproto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Stream: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stream: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, Entry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogproto
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLogproto
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Line", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogproto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogproto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Line = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogproto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogproto
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLogproto(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogproto
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLogproto
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLogproto
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthLogproto
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLogproto
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLogproto(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthLogproto
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLogproto = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogproto   = fmt.Errorf("proto: integer overflow")
)
