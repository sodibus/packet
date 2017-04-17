package packet

import "io"
import "fmt"

import "encoding/binary"
import "github.com/golang/protobuf/proto"

// Frame ...
// Frame represents a single frame data on TCP, containing a type identifier and
// a byte array of ambiguous data.
type Frame struct {
	Type uint8
	Data []uint8
}

// NewFrameWithPacket ...
// Create a new Frame from a Packet
func NewFrameWithPacket(m Packet) (*Frame, error) {
	t := frameTypeFromPacket(m)
	if t == 0 {
		return nil, fmt.Errorf("Packet Type %v Not Supported", m)
	}

	data, err := proto.Marshal(m)
	if err != nil {
		return nil, err
	}

	return &Frame{Type: t, Data: data}, nil
}

// ReadFrame ...
// Read a Frame from io.Reader
// if first byte is not 0xAA, a UnsynchronizedError will be returned
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

// ReadAndParse ...
// A combination of ReadFrame and Frame#Parse
func ReadAndParse(r io.Reader) (Packet, error) {
	f, err := ReadFrame(r)
	if err != nil {
		return nil, err
	}
	return f.Parse()
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

// Parse ...
// Parse the frame and return the underlying PacketXXXX structure
func (f *Frame) Parse() (Packet, error) {
	return packetFromFrame(f)
}
