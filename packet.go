package packet

import "github.com/golang/protobuf/proto"

// Packet virtual interface for all Packet types
type Packet interface {
	proto.Message
}
