// Code generated by protoc-gen-go.
// source: cluster_packet.proto
// DO NOT EDIT!

package packet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ClusterPacketNodeInfo struct {
	// uinque uint64 id
	NodeId uint64 `protobuf:"varint,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	// unique ip:port address of node cluster port
	NodeAddr string `protobuf:"bytes,2,opt,name=node_addr,json=nodeAddr" json:"node_addr,omitempty"`
	// whether this message is forwarded from other node
	IsForwarded bool `protobuf:"varint,3,opt,name=is_forwarded,json=isForwarded" json:"is_forwarded,omitempty"`
}

func (m *ClusterPacketNodeInfo) Reset()                    { *m = ClusterPacketNodeInfo{} }
func (m *ClusterPacketNodeInfo) String() string            { return proto.CompactTextString(m) }
func (*ClusterPacketNodeInfo) ProtoMessage()               {}
func (*ClusterPacketNodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ClusterPacketNodeInfo) GetNodeId() uint64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *ClusterPacketNodeInfo) GetNodeAddr() string {
	if m != nil {
		return m.NodeAddr
	}
	return ""
}

func (m *ClusterPacketNodeInfo) GetIsForwarded() bool {
	if m != nil {
		return m.IsForwarded
	}
	return false
}

type ClusterPacketNodeInfoBatch struct {
	Nodes []*ClusterPacketNodeInfo `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *ClusterPacketNodeInfoBatch) Reset()                    { *m = ClusterPacketNodeInfoBatch{} }
func (m *ClusterPacketNodeInfoBatch) String() string            { return proto.CompactTextString(m) }
func (*ClusterPacketNodeInfoBatch) ProtoMessage()               {}
func (*ClusterPacketNodeInfoBatch) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ClusterPacketNodeInfoBatch) GetNodes() []*ClusterPacketNodeInfo {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type ClusterPacketCalleeInfo struct {
	Id       *CalleeId `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Provides []string  `protobuf:"bytes,2,rep,name=provides" json:"provides,omitempty"`
}

func (m *ClusterPacketCalleeInfo) Reset()                    { *m = ClusterPacketCalleeInfo{} }
func (m *ClusterPacketCalleeInfo) String() string            { return proto.CompactTextString(m) }
func (*ClusterPacketCalleeInfo) ProtoMessage()               {}
func (*ClusterPacketCalleeInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ClusterPacketCalleeInfo) GetId() *CalleeId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ClusterPacketCalleeInfo) GetProvides() []string {
	if m != nil {
		return m.Provides
	}
	return nil
}

type ClusterPacketCalleeInfoBatch struct {
	Callees []*ClusterPacketCalleeInfo `protobuf:"bytes,1,rep,name=callees" json:"callees,omitempty"`
}

func (m *ClusterPacketCalleeInfoBatch) Reset()                    { *m = ClusterPacketCalleeInfoBatch{} }
func (m *ClusterPacketCalleeInfoBatch) String() string            { return proto.CompactTextString(m) }
func (*ClusterPacketCalleeInfoBatch) ProtoMessage()               {}
func (*ClusterPacketCalleeInfoBatch) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ClusterPacketCalleeInfoBatch) GetCallees() []*ClusterPacketCalleeInfo {
	if m != nil {
		return m.Callees
	}
	return nil
}

type ClusterPacketCalleeRemove struct {
	Id *CalleeId `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ClusterPacketCalleeRemove) Reset()                    { *m = ClusterPacketCalleeRemove{} }
func (m *ClusterPacketCalleeRemove) String() string            { return proto.CompactTextString(m) }
func (*ClusterPacketCalleeRemove) ProtoMessage()               {}
func (*ClusterPacketCalleeRemove) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ClusterPacketCalleeRemove) GetId() *CalleeId {
	if m != nil {
		return m.Id
	}
	return nil
}

type ClusterPacketInvocation struct {
	Id         *InvocationId `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Invocation *Invocation   `protobuf:"bytes,2,opt,name=invocation" json:"invocation,omitempty"`
}

func (m *ClusterPacketInvocation) Reset()                    { *m = ClusterPacketInvocation{} }
func (m *ClusterPacketInvocation) String() string            { return proto.CompactTextString(m) }
func (*ClusterPacketInvocation) ProtoMessage()               {}
func (*ClusterPacketInvocation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ClusterPacketInvocation) GetId() *InvocationId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ClusterPacketInvocation) GetInvocation() *Invocation {
	if m != nil {
		return m.Invocation
	}
	return nil
}

type ClusterPacketResult struct {
	Id     *InvocationId `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Code   ErrorCode     `protobuf:"varint,2,opt,name=code,enum=sodibus.ErrorCode" json:"code,omitempty"`
	Result string        `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
}

func (m *ClusterPacketResult) Reset()                    { *m = ClusterPacketResult{} }
func (m *ClusterPacketResult) String() string            { return proto.CompactTextString(m) }
func (*ClusterPacketResult) ProtoMessage()               {}
func (*ClusterPacketResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ClusterPacketResult) GetId() *InvocationId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ClusterPacketResult) GetCode() ErrorCode {
	if m != nil {
		return m.Code
	}
	return ErrorCode_OK
}

func (m *ClusterPacketResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*ClusterPacketNodeInfo)(nil), "sodibus.ClusterPacketNodeInfo")
	proto.RegisterType((*ClusterPacketNodeInfoBatch)(nil), "sodibus.ClusterPacketNodeInfoBatch")
	proto.RegisterType((*ClusterPacketCalleeInfo)(nil), "sodibus.ClusterPacketCalleeInfo")
	proto.RegisterType((*ClusterPacketCalleeInfoBatch)(nil), "sodibus.ClusterPacketCalleeInfoBatch")
	proto.RegisterType((*ClusterPacketCalleeRemove)(nil), "sodibus.ClusterPacketCalleeRemove")
	proto.RegisterType((*ClusterPacketInvocation)(nil), "sodibus.ClusterPacketInvocation")
	proto.RegisterType((*ClusterPacketResult)(nil), "sodibus.ClusterPacketResult")
}

func init() { proto.RegisterFile("cluster_packet.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x25, 0xe9, 0x9a, 0xb6, 0x5f, 0x16, 0xc1, 0x59, 0xd7, 0x8d, 0x5d, 0x91, 0x6c, 0x60, 0x25,
	0xa7, 0x1c, 0xb2, 0x9e, 0x3c, 0x08, 0xb6, 0x28, 0xf4, 0x22, 0x32, 0x27, 0xe9, 0xa5, 0x4c, 0xf3,
	0x4d, 0xed, 0x60, 0x9a, 0x09, 0x33, 0x49, 0x45, 0xf0, 0xc7, 0x4b, 0x66, 0x92, 0xb4, 0x95, 0x08,
	0x3d, 0x85, 0x79, 0xdf, 0x7b, 0xdf, 0x7b, 0x8f, 0x7c, 0xf0, 0x32, 0xcb, 0x6b, 0x5d, 0x71, 0xb5,
	0x2e, 0x59, 0xf6, 0x93, 0x57, 0x49, 0xa9, 0x64, 0x25, 0xc9, 0x58, 0x4b, 0x14, 0x9b, 0x5a, 0xcf,
	0xae, 0xf5, 0x8e, 0x29, 0x8e, 0x16, 0x8e, 0x4a, 0xb8, 0x5d, 0x58, 0xfa, 0x37, 0xc3, 0xfe, 0x2a,
	0x91, 0x2f, 0x8b, 0xad, 0x24, 0x77, 0x30, 0x2e, 0x24, 0xf2, 0xb5, 0xc0, 0xc0, 0x09, 0x9d, 0xf8,
	0x8a, 0x7a, 0xcd, 0x73, 0x89, 0xe4, 0x1e, 0xa6, 0x66, 0xc0, 0x10, 0x55, 0xe0, 0x86, 0x4e, 0x3c,
	0xa5, 0x93, 0x06, 0xf8, 0x84, 0xa8, 0xc8, 0x03, 0x5c, 0x0b, 0xbd, 0xde, 0x4a, 0xf5, 0x8b, 0x29,
	0xe4, 0x18, 0x8c, 0x42, 0x27, 0x9e, 0x50, 0x5f, 0xe8, 0x2f, 0x1d, 0x14, 0x51, 0x98, 0x0d, 0x3a,
	0xce, 0x59, 0x95, 0xed, 0xc8, 0x7b, 0x78, 0xd6, 0x2c, 0xd3, 0x81, 0x13, 0x8e, 0x62, 0x3f, 0x7d,
	0x9b, 0xb4, 0xb1, 0x93, 0x41, 0x0d, 0xb5, 0xe4, 0xe8, 0x3b, 0xdc, 0x9d, 0xcd, 0x17, 0x2c, 0xcf,
	0xb9, 0xed, 0xf1, 0x00, 0x6e, 0x5b, 0xc1, 0x4f, 0x5f, 0x1c, 0xb7, 0x59, 0x02, 0x52, 0x57, 0x20,
	0x99, 0xc1, 0xa4, 0x54, 0xf2, 0x20, 0x1a, 0x5b, 0x37, 0x1c, 0x35, 0x85, 0xba, 0x77, 0xb4, 0x82,
	0x37, 0xff, 0xd9, 0x6c, 0xf3, 0x7e, 0x80, 0x71, 0x66, 0xa0, 0x2e, 0x71, 0x38, 0x9c, 0xf8, 0xa8,
	0xa3, 0x9d, 0x20, 0xfa, 0x08, 0xaf, 0x07, 0x38, 0x94, 0xef, 0xe5, 0x81, 0x5f, 0x90, 0x3b, 0xaa,
	0xff, 0x69, 0xbd, 0x2c, 0x0e, 0x32, 0x63, 0x95, 0x90, 0x05, 0x79, 0x3c, 0x51, 0xdf, 0xf6, 0xea,
	0x23, 0xa1, 0x6d, 0xfe, 0x04, 0x20, 0x7a, 0xcc, 0xfc, 0x4c, 0x3f, 0xbd, 0x19, 0xa0, 0xd3, 0x13,
	0x5a, 0xf4, 0x07, 0x6e, 0xce, 0x6c, 0x29, 0xd7, 0x75, 0x5e, 0x5d, 0x6a, 0xf9, 0x0e, 0xae, 0x32,
	0x89, 0xdc, 0x98, 0x3d, 0x4f, 0x49, 0x4f, 0xfc, 0xac, 0x94, 0x54, 0x0b, 0x89, 0x9c, 0x9a, 0x39,
	0x79, 0x05, 0x9e, 0x32, 0x8b, 0xcd, 0x0d, 0x4d, 0x69, 0xfb, 0x9a, 0x3f, 0xc2, 0xbd, 0xfe, 0xad,
	0x2b, 0xbe, 0xd7, 0xc9, 0x9e, 0xfd, 0x10, 0xbd, 0xde, 0x1e, 0xfb, 0xca, 0xb3, 0xdf, 0x8d, 0x67,
	0xce, 0xfb, 0xe9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xe1, 0xc4, 0x2f, 0x0d, 0x03, 0x00,
	0x00,
}
