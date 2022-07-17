package protocus

import (
	"encoding/binary"
	"testing"
)

func TestBroadcast(t *testing.T) {
	output, err := Broadcast("testchan", []byte{})
	if err != nil {
		t.Error(err)
	}

	command := int8(binary.BigEndian.Uint64(output[0:8]))
	if command != int8(broadcast) {
		t.Error("command code is not broadcast")
	}
}
