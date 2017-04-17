package packet

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func frameTypeFromPacket(m Packet) uint8 {
	switch m.(type) {

	case *PacketHandshake:
		{
			return 1
		}

	case *PacketReady:
		{
			return 2
		}

	case *PacketCallerSend:
		{
			return 3
		}

	case *PacketCallerRecv:
		{
			return 4
		}

	case *PacketCalleeRecv:
		{
			return 5
		}

	case *PacketCalleeSend:
		{
			return 6
		}

	case *ClusterPacketNodeInfo:
		{
			return 101
		}

	case *ClusterPacketNodeInfoBatch:
		{
			return 102
		}

	case *ClusterPacketCalleeInfo:
		{
			return 103
		}

	case *ClusterPacketCalleeInfoBatch:
		{
			return 104
		}

	case *ClusterPacketCalleeRemove:
		{
			return 105
		}

	case *ClusterPacketInvocation:
		{
			return 106
		}

	case *ClusterPacketResult:
		{
			return 107
		}

	default:
		{
			return 0
		}
	}
}

func packetFromFrame(f *Frame) (Packet, error) {
	var err error
	switch f.Type {

	case 1:
		{
			var p PacketHandshake
			err = proto.Unmarshal(f.Data, &p)
			return &p, err
		}

	case 2:
		{
			var p PacketReady
			err = proto.Unmarshal(f.Data, &p)
			return &p, err
		}

	case 3:
		{
			var p PacketCallerSend
			err = proto.Unmarshal(f.Data, &p)
			return &p, err
		}

	case 4:
		{
			var p PacketCallerRecv
			err = proto.Unmarshal(f.Data, &p)
			return &p, err
		}

	case 5:
		{
			var p PacketCalleeRecv
			err = proto.Unmarshal(f.Data, &p)
			return &p, err
		}

	case 6:
		{
			var p PacketCalleeSend
			err = proto.Unmarshal(f.Data, &p)
			return &p, err
		}

	case 101:
		{
			var p ClusterPacketNodeInfo
			err = proto.Unmarshal(f.Data, &p)
			return &p, err
		}

	case 102:
		{
			var p ClusterPacketNodeInfoBatch
			err = proto.Unmarshal(f.Data, &p)
			return &p, err
		}

	case 103:
		{
			var p ClusterPacketCalleeInfo
			err = proto.Unmarshal(f.Data, &p)
			return &p, err
		}

	case 104:
		{
			var p ClusterPacketCalleeInfoBatch
			err = proto.Unmarshal(f.Data, &p)
			return &p, err
		}

	case 105:
		{
			var p ClusterPacketCalleeRemove
			err = proto.Unmarshal(f.Data, &p)
			return &p, err
		}

	case 106:
		{
			var p ClusterPacketInvocation
			err = proto.Unmarshal(f.Data, &p)
			return &p, err
		}

	case 107:
		{
			var p ClusterPacketResult
			err = proto.Unmarshal(f.Data, &p)
			return &p, err
		}

	default:
		{
			return nil, fmt.Errorf("Packet Type %d Not Supported", f.Type)
		}
	}
}
