package classfile

import "encoding/binary"

/*
	ClassReader是byte[]的包装
 */
type ClassReader struct {
	data []byte
}

/*
	读取JVM中u1类型的数据
 */
func (self *ClassReader) readUnit8() uint8{
	val := self.data[0]
	self.data = self.data[1:]
	return val
}
/*
	读取JVM中u2类型的数据
*/
func (self *ClassReader) readUnit16() uint16{
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}
/*
	读取JVM中u4类型的数据
*/
func (self *ClassReader) readUnit32() uint32{
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}
/*
	读取u8类型的数据
*/
func (self *ClassReader) readUnit64() uint64{
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}
/*
	读取Uint16表
*/
func (self *ClassReader) readUnit16s() []uint16{
	n := self.readUnit16()
	s := make([]uint16, n)
	for i:= range s {
		s[i] = self.readUnit16()
	}
	return s
}
/*
	读取指定数量的字节
*/
func (self *ClassReader) readBytes(n uint32) []byte{
	bytes := self.data[:n]
	self.data = self.data[:n]
	return bytes
}