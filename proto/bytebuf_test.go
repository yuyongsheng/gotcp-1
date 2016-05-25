package proto

import (
	"testing"
	"strconv"
)

func TestByteBuf(t *testing.T) {
	var proto = NewByteBuf(20)

	proto.WriteUint32(9000)
	proto.WriteInt32(-8000)
	proto.WriteBool(true)
	proto.WriteFloat32(3.141592)

	if proto.ReadUint32() != 9000 {
		t.Error("FAILED")
	}

	if proto.ReadInt32() != -8000 {
		t.Error("FAILED")
	}

	if proto.ReadBool() != true {
		t.Error("FAILED")
	}

	if f := proto.ReadFloat32(); f != 3.141592 {
		s := strconv.FormatFloat(float64(f), 'E', -1, 32)
		t.Error("FAILED:" + s)
	}

	if proto.writeIndex != 13 {
		t.Error("FAILED")
	}

	if proto.readIndex != 13 {
		t.Error("FAILED")
	}

	if proto.GetCapacity() != 20 {
		t.Error("FAILED")
	}
}