package tcp

import (
	"testing"
	"strconv"
)

func TestByteBuffer(t *testing.T) {
	var proto = NewByteBuffer(20)

	proto.PutUint32(9000)
	proto.PutInt32(-8000)
	proto.PutBool(true)
	proto.PutFloat32(3.141592)

	if proto.GetUint32() != 9000 {
		t.Error("FAILED")
	}

	if proto.GetInt32() != -8000 {
		t.Error("FAILED")
	}

	if proto.GetBool() != true {
		t.Error("FAILED")
	}

	if f := proto.GetFloat32(); f != 3.141592 {
		s := strconv.FormatFloat(float64(f), 'E', -1, 32)
		t.Error("FAILED:" + s)
	}

	if proto.writePos != 13 {
		t.Error("FAILED")
	}

	if proto.readPos != 13 {
		t.Error("FAILED")
	}

	if proto.capacity != 20 {
		t.Error("FAILED")
	}
}