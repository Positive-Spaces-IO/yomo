package frame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandshakeFrameEncode(t *testing.T) {
	expectedName := "1234"
	var expectedType byte = 0xD3
	m := NewHandshakeFrame(expectedName, "", expectedType, []Tag{0x01, 0x02}, "token", "a")
	assert.Equal(t, []byte{
		0x80 | byte(TagOfHandshakeFrame), 0x1f,
		byte(TagOfHandshakeName), 0x04, 0x31, 0x32, 0x33, 0x34,
		byte(TagOfHandshakeID), 0x0,
		byte(TagOfHandshakeType), 0x01, 0xD3,
		byte(TagOfHandshakeObserveDataTags), 0x8, 0x1, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0,
		byte(TagOfHandshakeAuthName), 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
		byte(TagOfHandshakeAuthPayload), 0x01, 0x61,
	},
		m.Encode(),
	)

	Handshake, err := DecodeToHandshakeFrame(m.Encode())
	assert.NoError(t, err)
	assert.EqualValues(t, expectedName, Handshake.Name)
	assert.EqualValues(t, expectedType, Handshake.ClientType)
	assert.EqualValues(t, "token", Handshake.AuthName())
	assert.EqualValues(t, "a", Handshake.AuthPayload())
}
