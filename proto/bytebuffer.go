package tcp


import (
	"encoding/binary"
	"math"
)

type ByteBuffer struct {
	buf []byte

	order binary.ByteOrder
	readPos int
	writePos int
	capacity int
}

func NewByteBuffer(cap int) *ByteBuffer {
	return &ByteBuffer {
		buf: make([]byte, cap),
		order: binary.BigEndian,
		readPos: 0,
		writePos: 0,
		capacity: cap,
	}
}

func (self *ByteBuffer) GetBuf() []byte {
	return self.buf
}

func (self *ByteBuffer) GetCapacity() int {
	return self.capacity
}

func (self *ByteBuffer) Reserve(cap int) {
	if self.capacity >= cap {
		return
	}
	var newBuf = make([]byte, cap)
	copy(newBuf, self.buf)
	self.capacity = cap
}

func (self *ByteBuffer) Reset() {
	self.readPos = 0
	self.writePos = 0
}

func (self *ByteBuffer) StepRead(step int) {
	self.readPos += step
}

func (self *ByteBuffer) StepWrite(step int) {
	self.writePos += step
}

func (self *ByteBuffer) ReadBuf(n int) []byte {
	return self.buf[self.readPos:self.readPos + n]
}

func (self *ByteBuffer) WriteBuf(n int) []byte {
	return self.buf[self.writePos:self.writePos + n]
}

func (self *ByteBuffer) GetBool() bool {
	x := self.buf[self.readPos]
	self.readPos += 1
	return x != 0
}

func (self *ByteBuffer) PutBool(v bool) {
	if v == true {
		self.buf[self.writePos] = 1
	} else {
		self.buf[self.writePos] = 0
	}
	self.writePos += 1
}

func (self *ByteBuffer) GetUint8() uint8 {
	x := self.buf[self.readPos]
	self.readPos += 1
	return x
}

func (self *ByteBuffer) PutUint8(v uint8) {
	self.buf[self.writePos] = v
	self.writePos += 1
}

func (self *ByteBuffer) GetUint16() uint16 {
	x := self.order.Uint16(self.buf[self.readPos:self.readPos + 2])
	self.readPos += 2
	return x
}

func (self *ByteBuffer) PutUint16(v uint16) {
	self.order.PutUint16(self.buf[self.writePos:self.writePos + 2], v)
	self.writePos += 2
}

func (self *ByteBuffer) GetUint32() uint32 {
	x := self.order.Uint32(self.buf[self.readPos:self.readPos + 4])
	self.readPos += 4
	return x
}

func (self *ByteBuffer) PutUint32(v uint32) {
	self.order.PutUint32(self.buf[self.writePos:self.writePos + 4], v)
	self.writePos += 4
}

func (self *ByteBuffer) GetUint64() uint64 {
	x := self.order.Uint64(self.buf[self.readPos:self.readPos + 8])
	self.readPos += 8
	return x
}

func (self *ByteBuffer) PutUint64(v uint64) {
	self.order.PutUint64(self.buf[self.writePos:self.writePos + 8], v)
	self.writePos += 8
}

func (self *ByteBuffer) GetBytes(n int) []byte {
	x := self.buf[self.readPos:self.readPos + n]
	self.readPos += n
	return x
}

func (self *ByteBuffer) PutBytes(v []byte) {
	copy(self.buf[self.writePos:], v)
	self.writePos += len(v)
}

func (self *ByteBuffer) GetByte() byte {
	return byte(self.GetUint8())
}

func (self *ByteBuffer) PutByte(v byte) {
	self.PutUint8(uint8(v))
}

func (self *ByteBuffer) GetInt8() int8 {
	return int8(self.GetUint8())
}

func (self *ByteBuffer) PutInt8(v int8) {
	self.PutUint8(uint8(v))
}

func (self *ByteBuffer) GetInt16() int16 {
	return int16(self.GetUint16())
}

func (self *ByteBuffer) PutInt16(v int16) {
	self.PutUint16(uint16(v))
}

func (self *ByteBuffer) GetInt32() int32 {
	return int32(self.GetUint32())
}

func (self *ByteBuffer) PutInt32(v int32) {
	self.PutUint32(uint32(v))
}

func (self *ByteBuffer) GetInt64() int64 {
	return int64(self.GetUint64())
}

func (self *ByteBuffer) PutInt64(v int64) {
	self.PutUint64(uint64(v))
}

func (self *ByteBuffer) GetFloat32() float32 {
    return math.Float32frombits(self.GetUint32())
}

func (self *ByteBuffer) PutFloat32(v float32) {
	self.PutUint32(math.Float32bits(v))
}

func (self *ByteBuffer) GetFloat64() float64 {
    return math.Float64frombits(self.GetUint64())
}

func (self *ByteBuffer) PutFloat64(v float64) {
	self.PutUint64(math.Float64bits(v))
}

func (self *ByteBuffer) GetRune() rune {
    return rune(self.GetUint32())
}

func (self *ByteBuffer) PutRune(v rune) {
	self.PutUint32(uint32(v))
}

func (self *ByteBuffer) GetString(n int) string {
	return string(self.GetBytes(n))
}

func (self *ByteBuffer) PutString(v string) {
	self.PutBytes([]byte(v))
}

