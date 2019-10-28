package protocol

import (
	"bytes"
	"encoding/binary"
)

type ReadObj struct {
	BufferCap uint // buffer的容量
	DataLen   uint // 数据长度
	Buffer    []byte
	Rpos      uint
	isReadOk  bool
}

func NewReadObj() *ReadObj {
	obj := &ReadObj{
		DataLen:  0,
		Rpos:     0,
		isReadOk: true,
	}

	obj.Buffer = make([]byte, 1024)
	obj.BufferCap = uint(cap(obj.Buffer))

	return obj
}

func (self *ReadObj) PushData(data []byte) {
	len := uint(len(data))
	if len > self.BufferCap {
		self.Buffer = make([]byte, 0) // 重置
		self.Buffer = append(self.Buffer, data...)
		self.BufferCap = uint(cap(self.Buffer))
	} else {
		copy(self.Buffer, data)
	}

	self.Rpos = 0
	self.DataLen = uint(len)
}

func (self *ReadObj) SetDataLen(dataLen uint16) {
	self.DataLen = uint(dataLen)
}

func (self *ReadObj) DecodeHeader(h *Header) bool {

	for i := 0; i < 2; i++ {
		self.ReadInt8(&h.Flag[i])
	}
	self.ReadUInt16(&h.Cmd)
	self.ReadInt8(&h.Ver)
	self.ReadUInt8(&h.Link)
	self.ReadUInt16(&h.BodyLen)

	for i := 0; i < 8; i++ {
		self.ReadInt8(&h.Reverse[i])
	}

	return self.isReadOk
}

func (self *ReadObj) skip(count uint) {
	self.Rpos += count
}

func (self *ReadObj) ReadInt8(dst *int8) {
	if self.Rpos > self.DataLen {
		self.isReadOk = false
		return
	}
	buf := bytes.NewBuffer(self.Buffer[self.Rpos:])
	binary.Read(buf, binary.BigEndian, dst)
	self.Rpos += 1
}

func (self *ReadObj) ReadUInt8(dst *uint8) {
	if self.Rpos > self.DataLen {
		self.isReadOk = false
		return
	}
	buf := bytes.NewBuffer(self.Buffer[self.Rpos:])
	binary.Read(buf, binary.BigEndian, dst)
	self.Rpos += 1
}

func (self *ReadObj) ReadInt16(dst *int16) {
	if self.Rpos > self.DataLen {
		self.isReadOk = false
		return
	}
	buf := bytes.NewBuffer(self.Buffer[self.Rpos:])
	binary.Read(buf, binary.BigEndian, dst)
	self.Rpos += 2
}

func (self *ReadObj) ReadUInt16(dst *uint16) {
	if self.Rpos > self.DataLen {
		self.isReadOk = false
		return
	}
	buf := bytes.NewBuffer(self.Buffer[self.Rpos:])
	binary.Read(buf, binary.BigEndian, dst)
	self.Rpos += 2
}

func (self *ReadObj) ReadInt32(dst *int32) {
	if self.Rpos > self.DataLen {
		self.isReadOk = false
		return
	}
	buf := bytes.NewBuffer(self.Buffer[self.Rpos:])
	binary.Read(buf, binary.BigEndian, dst)
	self.Rpos += 4
}

func (self *ReadObj) ReadUInt32(dst *uint32) {
	if self.Rpos > self.DataLen {
		self.isReadOk = false
		return
	}
	buf := bytes.NewBuffer(self.Buffer[self.Rpos:])
	binary.Read(buf, binary.BigEndian, dst)
	self.Rpos += 4
}

func (self *ReadObj) ReadInt64(dst *int64) {
	if self.Rpos > self.DataLen {
		self.isReadOk = false
		return
	}
	buf := bytes.NewBuffer(self.Buffer[self.Rpos:])
	binary.Read(buf, binary.BigEndian, dst)
	self.Rpos += 8
}

func (self *ReadObj) ReadUInt64(dst *uint64) {
	if self.Rpos > self.DataLen {
		self.isReadOk = false
		return
	}
	buf := bytes.NewBuffer(self.Buffer[self.Rpos:])
	binary.Read(buf, binary.BigEndian, dst)
	self.Rpos += 8
}

func (self *ReadObj) ReadString(dst *string) {
	var strLen uint16
	self.ReadUInt16(&strLen)
	if self.isReadOk {
		if self.Rpos+uint(strLen) > self.DataLen {
			self.isReadOk = false
			return
		}
		// buf := bytes.NewBuffer(self.Buffer[self.Rpos : self.Rpos+uint(strLen)-1])
		// binary.Read(buf, binary.BigEndian, dst)
		*dst = string(self.Buffer[self.Rpos : self.Rpos+uint(strLen)-1])
		self.Rpos += uint(strLen)
	} else {
		self.isReadOk = false
	}
}

func (self *ReadObj) IsReadOk() bool {
	return self.isReadOk
}
