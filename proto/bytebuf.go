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

func (self *ByteBuf) GetCapacity() int {
	return len(self.buf)
}

func (self *ByteBuf) Reserve(cap int) {
	GetCapacity := self.GetCapacity()
	if GetCapacity >= cap {
		return
	}
	for GetCapacity < cap {
		GetCapacity = (GetCapacity + 1) * 2
	}
	var newBuf = make([]byte, GetCapacity)
	copy(newBuf, self.buf)
}

func (self *ByteBuf) Clear() {
	self.readIndex = 0
	self.writeIndex = 0
}

func (self *ByteBuf) GetReadableBuf() []byte {
	return self.buf[self.readIndex:self.writeIndex]
}

func (self *ByteBuf) GetWritableBuf(n int) []byte {
	return self.buf[self.writeIndex:self.writeIndex + n]
}

func (self *ByteBuf) GetWritedBuf() []byte {
	return self.buf[:self.writeIndex]
}

func (self *ByteBuf) GetReadableLen() int {
	return self.writeIndex - self.readIndex
}

func (self *ByteBuf) GetWritableLen() int {
	return self.GetCapacity() - self.writeIndex
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

func (self *ByteBuf) ReadBool() bool {
	x := self.buf[self.readIndex]
	self.readIndex += 1
	return x != 0
}

func (self *ByteBuf) WriteBool(v bool) {
	if v {
		self.buf[self.writeIndex] = 1
	} else {
		self.buf[self.writeIndex] = 0
	}
	self.writeIndex += 1
}

func (self *ByteBuf) ReadUint8() uint8 {
	x := self.buf[self.readIndex]
	self.readIndex += 1
	return x
}

func (self *ByteBuf) WriteUint8(v uint8) {
	self.buf[self.writeIndex] = v
	self.writeIndex += 1
}

func (self *ByteBuf) ReadUint16() uint16 {
	x := self.order.Uint16(self.buf[self.readIndex:self.readIndex + 2])
	self.readIndex += 2
	return x
}

func (self *ByteBuf) WriteUint16(v uint16) {
	self.order.PutUint16(self.buf[self.writeIndex:self.writeIndex + 2], v)
	self.writeIndex += 2
}

func (self *ByteBuf) ReadUint32() uint32 {
	x := self.order.Uint32(self.buf[self.readIndex:self.readIndex + 4])
	self.readIndex += 4
	return x
}

func (self *ByteBuf) WriteUint32(v uint32) {
	self.order.PutUint32(self.buf[self.writeIndex:self.writeIndex + 4], v)
	self.writeIndex += 4
}

func (self *ByteBuf) ReadUint64() uint64 {
	x := self.order.Uint64(self.buf[self.readIndex:self.readIndex + 8])
	self.readIndex += 8
	return x
}

func (self *ByteBuf) WriteUint64(v uint64) {
	self.order.PutUint64(self.buf[self.writeIndex:self.writeIndex + 8], v)
	self.writeIndex += 8
}

func (self *ByteBuf) ReadBytes(n int) []byte {
	x := self.buf[self.readIndex:self.readIndex + n]
	self.readIndex += n
	return x
}

func (self *ByteBuf) WriteBytes(v []byte) {
	copy(self.buf[self.writeIndex:], v)
	self.writeIndex += len(v)
}

func (self *ByteBuf) ReadByte() byte {
	return byte(self.ReadUint8())
}

func (self *ByteBuf) WriteByte(v byte) {
	self.WriteUint8(uint8(v))
}

func (self *ByteBuf) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *ByteBuf) WriteInt8(v int8) {
	self.WriteUint8(uint8(v))
}

func (self *ByteBuf) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *ByteBuf) WriteInt16(v int16) {
	self.WriteUint16(uint16(v))
}

func (self *ByteBuf) ReadInt32() int32 {
	return int32(self.ReadUint32())
}

func (self *ByteBuf) WriteInt32(v int32) {
	self.WriteUint32(uint32(v))
}

func (self *ByteBuf) ReadInt64() int64 {
	return int64(self.ReadUint64())
}

func (self *ByteBuf) WriteInt64(v int64) {
	self.WriteUint64(uint64(v))
}

func (self *ByteBuf) ReadFloat32() float32 {
    return math.Float32frombits(self.ReadUint32())
}

func (self *ByteBuf) WriteFloat32(v float32) {
	self.WriteUint32(math.Float32bits(v))
}

func (self *ByteBuf) ReadFloat64() float64 {
    return math.Float64frombits(self.ReadUint64())
}

func (self *ByteBuf) WriteFloat64(v float64) {
	self.WriteUint64(math.Float64bits(v))
}

func (self *ByteBuf) ReadRune() rune {
    return rune(self.ReadUint32())
}

func (self *ByteBuf) WriteRune(v rune) {
	self.WriteUint32(uint32(v))
}

func (self *ByteBuf) ReadString(n int) string {
	return string(self.ReadBytes(n))
}

func (self *ByteBuf) WriteString(v string) {
	self.WriteBytes([]byte(v))
}

