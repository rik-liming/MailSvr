package protocol

import (
	"bytes"
	"encoding/binary"
)

const PACK_HEADER_LEN = 16

type Header struct {
	Flag    [2]int8
	Cmd     uint16
	Ver     int8
	Link    uint8
	BodyLen uint16
	Reverse [8]int8
}

type WriteObj struct {
	Header         // 继承
	BufferCap uint // buffer的容量
	Buffer    []byte
	Wpos      uint
	isWriteOk bool
}

func NewWriteObj() *WriteObj {
	obj := &WriteObj{
		Wpos:      0,
		isWriteOk: true,
	}

	obj.Flag = [2]int8{'Y', 'H'}
	obj.Buffer = make([]byte, 1024)
	obj.BufferCap = uint(cap(obj.Buffer))

	return obj
}

func (self *WriteObj) Init(cmd uint16) {

	self.Wpos = 0
	self.isWriteOk = true

	self.Cmd = cmd
	self.write_header()
}

func (self *WriteObj) skip(count uint) {
	if self.Wpos+count > self.BufferCap {
		tmp := make([]byte, self.Wpos+count+1024)
		copy(tmp, self.Buffer)
		self.Buffer = tmp
		self.BufferCap = uint(cap(self.Buffer))
	}
	self.Wpos += count
}

func (self *WriteObj) write_header() {
	self.WriteInt8(self.Flag[0])
	self.WriteInt8(self.Flag[1])
	self.WriteUInt16(self.Cmd)
	self.WriteInt8(self.Ver)
	self.WriteUInt8(self.Link)
	self.WriteUInt16(self.BodyLen)
	self.skip(uint(len(self.Reverse)))
}

func (self *WriteObj) get_header_size() uint {
	return 16
}

func (self *WriteObj) set_body_len() {
	self.BodyLen = uint16(self.Wpos - self.get_header_size())
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, self.BodyLen)
	copy(self.Buffer[6:], buf.Bytes())
}

func (self *WriteObj) WriteByte(src []byte) error {
	if self.Wpos+uint(len(src)) > self.BufferCap {
		self.Buffer = append(self.Buffer, src...)
		self.BufferCap = uint(cap(self.Buffer))
	} else {
		copy(self.Buffer[self.Wpos:], src)
	}
	self.Wpos += uint(len(src))
	return nil
}

func (self *WriteObj) WriteInt8(src int8) {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, src)
	self.WriteByte(buf.Bytes())
}

func (self *WriteObj) WriteInt16(src int16) {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, src)
	self.WriteByte(buf.Bytes())
}

func (self *WriteObj) WriteInt32(src int32) {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, src)
	self.WriteByte(buf.Bytes())
}

func (self *WriteObj) WriteInt64(src int64) {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, src)
	self.WriteByte(buf.Bytes())
}

func (self *WriteObj) WriteUInt8(src uint8) {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, src)
	self.WriteByte(buf.Bytes())
}

func (self *WriteObj) WriteUInt16(src uint16) {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, src)
	self.WriteByte(buf.Bytes())
}

func (self *WriteObj) WriteUInt32(src uint32) {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, src)
	self.WriteByte(buf.Bytes())
}

func (self *WriteObj) WriteUInt64(src uint64) {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, src)
	self.WriteByte(buf.Bytes())
}

func (self *WriteObj) WriteString(src string) {
	buf := []byte(src)
	buf = append(buf, 0)
	len := uint16(len(buf))
	self.WriteUInt16(len)
	self.WriteByte(buf)
}

func (self *WriteObj) WriteData(src interface{}) {
	switch src := src.(type) {
	case int8:
		self.WriteInt8(src)
	case uint8:
		self.WriteUInt8(src)
	case int16:
		self.WriteInt16(src)
	case uint16:
		self.WriteUInt16(src)
	case int32:
		self.WriteInt32(src)
	case uint32:
		self.WriteUInt32(src)
	case int64:
		self.WriteInt64(src)
	case uint64:
		self.WriteUInt64(src)
	case string:
		self.WriteString(src)
	default:
		// not support type
		self.isWriteOk = false
	}
}

func (self *WriteObj) GetBuf() []byte {
	self.set_body_len()
	return self.Buffer[:self.Wpos]
}

func (self *WriteObj) GetDataLen() uint {
	return self.Wpos
}

func (self *WriteObj) IsWriteOk() bool {
	return self.isWriteOk
}
