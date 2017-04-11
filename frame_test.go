package packet

import "testing"
import "github.com/stvp/assert"
import "github.com/golang/protobuf/proto"

func TestNewFrame(t *testing.T) {
	p := PacketHandshake{ Provides: []string{ "calculator" } }
	f, err := NewFrameWithPacket(&p)
	assert.Equal(t, f.Type, uint8(1), "Frame Type should be 1")
	assert.Equal(t, err, nil, "Error should be nil")
}

func TestFrameParsing(t *testing.T) {
	p := PacketHandshake{ Provides: []string{ "calculator" } }
	f, err := NewFrameWithPacket(&p)
	assert.Equal(t, f.Type, uint8(1), "Frame Type should be 1")
	assert.Equal(t, err, nil, "Error should be nil")

	var m proto.Message
	m, err = f.Parse()
	assert.Equal(t, err, nil, "Error should be nil")

	p2, ok := m.(*PacketHandshake)
	assert.Equal(t, ok, true, "Downcast should be OK")
	assert.NotEqual(t, p2, nil, "Downcast should be OK")

	assert.Equal(t, len(p2.Provides), 1, "Provides should be OK")
	assert.Equal(t, p2.Provides[0], "calculator", "Provides should be OK")
}

