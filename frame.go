package packet

import "io"
import "fmt"
import "errors"
import "encoding/binary"
import "github.com/golang/protobuf/proto"

type Frame struct {
	Type uint8
	Data []uint8
}

func NewFrameWithPacket(m proto.Message) (*Frame, error) {
	var t uint8 = 0
	switch m.(type) {
	case *PacketHandshake:
		{
			t = 1
		}
	case *PacketReady:
		{
			t = 2
		}
	case *PacketCallerSend:
		{
			t = 3
		}
	case *PacketCallerRecv:
		{
			t = 4
		}
	case *PacketCalleeRecv:
		{
			t = 5
		}
	case *PacketCalleeSend:
		{
			t = 6
		}
	case *ClusterPacketNodeInfo:
		{
			t = 101
		}
	case *ClusterPacketNodeInfoBatch:
		{
			t = 102
		}
	case *ClusterPacketCalleeInfo:
		{
			t = 103
		}
	case *ClusterPacketCalleeInfoBatch:
		{
			t = 104
		}
	case *ClusterPacketCalleeRemove:
		{
			t = 105
		}
	case *ClusterPacketInvocation:
		{
			t = 106
		}
	case *ClusterPacketResult:
		{
			t = 107
		}
	default:
		{
			return nil, errors.New("Unsupported Packet")
		}
	}

	data, err := proto.Marshal(m)
	if err != nil {
		return nil, err
	}

	return &Frame{Type: t, Data: data}, nil
}

func ReadFrame(r io.Reader) (*Frame, error) {
	var err error
	var head = make([]uint8, 1)
	var dataLen = make([]uint8, 4)

	// seek head (1 byte)
	_, err = r.Read(head)
	if err != nil {
		return nil, err
	}
	if head[0] != 0xAA {
		return nil, UnsynchronizedError{}
	}

	// read packet type (1 byte)
	_, err = r.Read(head)
	if err != nil {
		return nil, err
	}

	// read dataLen (4 byte, uint32)
	_, err = r.Read(dataLen)
	if err != nil {
		return nil, err
	}
	len := binary.BigEndian.Uint32(dataLen)

	// read data
	var data = make([]byte, len)
	_, err = r.Read(data)
	if err != nil {
		return nil, err
	}

	// build frame
	return &Frame{Type: head[0], Data: data}, err
}

func (f *Frame) Write(w io.Writer) error {
	var err error

	_, err = w.Write([]uint8{0xAA, f.Type})
	if err != nil {
		return err
	}

	dataLen := make([]byte, 4)
	binary.BigEndian.PutUint32(dataLen, uint32(len(f.Data)))
	w.Write(dataLen)
	if err != nil {
		return err
	}

	w.Write(f.Data)
	if err != nil {
		return err
	}

	return err
}

func (f *Frame) Parse() (proto.Message, error) {
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
			return nil, errors.New(fmt.Sprintf("Unsupported Packet %d", f.Type))
		}
	}
}
