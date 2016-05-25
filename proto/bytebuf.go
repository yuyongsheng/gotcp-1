package proto


import (
	"encoding/binary"
	"math"
)

type ByteBuf struct {
	buf []byte

	order binary.ByteOrder
	readIndex int
	writeIndex int
	makedReadIndex int
}

func NewByteBuf(cap int) *ByteBuf {
	return &ByteBuf {
		buf: make([]byte, cap),
		order: binary.LittleEndian,
		readIndex: 0,
		writeIndex: 0,
		makedReadIndex: 0,
	}
}

func (self *ByteBuf) GetBuf() []byte {
	return self.buf
}

func (self *ByteBuf) Capacity() int {
	return len(self.buf)
}

func (self *ByteBuf) Reserve(cap int) {
	capacity := self.Capacity()
	if capacity >= cap {
		return
	}
	for capacity < cap {
		capacity = (capacity + 1) * 2
	}
	var newBuf = make([]byte, capacity)
	copy(newBuf, self.buf)
}

func (self *ByteBuf) Clear() {
	self.readIndex = 0
	self.writeIndex = 0
}

func (self *ByteBuf) GetWritableBuf(n int) []byte {
	return self.buf[self.writeIndex:self.writeIndex + n]
}

func (self *ByteBuf) GetReadableBytes() int {
	return self.writeIndex - self.readIndex
}

func (self *ByteBuf) GetWritableBytes() int {
	return self.Capacity() - self.writeIndex
}

func (self *ByteBuf) MarkReadIndex() {
	self.makedReadIndex = self.readIndex
}

func (self *ByteBuf) ResetReadIndex() {
	self.readIndex = self.makedReadIndex
}

func (self *ByteBuf) MoveReadIndex(step int) {
	self.readIndex += step
}

func (self *ByteBuf) MoveWriteIndex(step int) {
	self.writeIndex += step
}

// decoding & encoding

func (self *ByteBuf) GetBool() bool {
	x := self.buf[self.readIndex]
	self.readIndex += 1
	return x != 0
}

func (self *ByteBuf) PutBool(v bool) {
	if v {
		self.buf[self.writeIndex] = 1
	} else {
		self.buf[self.writeIndex] = 0
	}
	self.writeIndex += 1
}

func (self *ByteBuf) GetUint8() uint8 {
	x := self.buf[self.readIndex]
	self.readIndex += 1
	return x
}

func (self *ByteBuf) PutUint8(v uint8) {
	self.buf[self.writeIndex] = v
	self.writeIndex += 1
}

func (self *ByteBuf) GetUint16() uint16 {
	x := self.order.Uint16(self.buf[self.readIndex:self.readIndex + 2])
	self.readIndex += 2
	return x
}

func (self *ByteBuf) PutUint16(v uint16) {
	self.order.PutUint16(self.buf[self.writeIndex:self.writeIndex + 2], v)
	self.writeIndex += 2
}

func (self *ByteBuf) GetUint32() uint32 {
	x := self.order.Uint32(self.buf[self.readIndex:self.readIndex + 4])
	self.readIndex += 4
	return x
}

func (self *ByteBuf) PutUint32(v uint32) {
	self.order.PutUint32(self.buf[self.writeIndex:self.writeIndex + 4], v)
	self.writeIndex += 4
}

func (self *ByteBuf) GetUint64() uint64 {
	x := self.order.Uint64(self.buf[self.readIndex:self.readIndex + 8])
	self.readIndex += 8
	return x
}

func (self *ByteBuf) PutUint64(v uint64) {
	self.order.PutUint64(self.buf[self.writeIndex:self.writeIndex + 8], v)
	self.writeIndex += 8
}

func (self *ByteBuf) GetBytes(n int) []byte {
	x := self.buf[self.readIndex:self.readIndex + n]
	self.readIndex += n
	return x
}

func (self *ByteBuf) PutBytes(v []byte) {
	copy(self.buf[self.writeIndex:], v)
	self.writeIndex += len(v)
}

func (self *ByteBuf) GetByte() byte {
	return byte(self.GetUint8())
}

func (self *ByteBuf) PutByte(v byte) {
	self.PutUint8(uint8(v))
}

func (self *ByteBuf) GetInt8() int8 {
	return int8(self.GetUint8())
}

func (self *ByteBuf) PutInt8(v int8) {
	self.PutUint8(uint8(v))
}

func (self *ByteBuf) GetInt16() int16 {
	return int16(self.GetUint16())
}

func (self *ByteBuf) PutInt16(v int16) {
	self.PutUint16(uint16(v))
}

func (self *ByteBuf) GetInt32() int32 {
	return int32(self.GetUint32())
}

func (self *ByteBuf) PutInt32(v int32) {
	self.PutUint32(uint32(v))
}

func (self *ByteBuf) GetInt64() int64 {
	return int64(self.GetUint64())
}

func (self *ByteBuf) PutInt64(v int64) {
	self.PutUint64(uint64(v))
}

func (self *ByteBuf) GetFloat32() float32 {
    return math.Float32frombits(self.GetUint32())
}

func (self *ByteBuf) PutFloat32(v float32) {
	self.PutUint32(math.Float32bits(v))
}

func (self *ByteBuf) GetFloat64() float64 {
    return math.Float64frombits(self.GetUint64())
}

func (self *ByteBuf) PutFloat64(v float64) {
	self.PutUint64(math.Float64bits(v))
}

func (self *ByteBuf) GetRune() rune {
    return rune(self.GetUint32())
}

func (self *ByteBuf) PutRune(v rune) {
	self.PutUint32(uint32(v))
}

func (self *ByteBuf) GetString(n int) string {
	return string(self.GetBytes(n))
}

func (self *ByteBuf) PutString(v string) {
	self.PutBytes([]byte(v))
}

