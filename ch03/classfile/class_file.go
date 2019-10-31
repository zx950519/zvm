package classfile

import "fmt"

/*
	Class文件结构的定义
 */
type ClassFile struct {
	//magic uint32
	minorVersion 	uint16
	majorVersion 	uint16
	constantPool 	ConstantPool
	accessFlags 	uint16
	thisClass 		uint16
	superClass 		uint16
	interfaces 		[]uint16
	fields 			[]*MemberInfo
	methods 		[]*MemberInfo
	attributes		[]*AttributeInfo
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) ThisClass() uint16 {
	return self.thisClass
}

func (self *ClassFile) SuperClass() uint16 {
	return self.superClass
}

func (self *ClassFile) Interfaces() []uint16 {
	return self.interfaces
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) Attributes() []*AttributeInfo {
	return self.attributes
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
/*
	Parse()函数将[]byte解析成ClassFile结构体
 */
func Parse(classData []byte)(cf *ClassFile, err error) {
	defer func() {
		if r:= recover(); r!=nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}
/*
	依次调用其他方法解析class文件
 */
func (self *ClassFile) read(reader *ClassReader) {
	// 检查魔数 0xCAFEBABE
	self.readAndCheckMagic(reader)
	// 检查主版本号和副版本号
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	// 检查类访问标志，一个16位的 bitmask ，表明文件是类还是接口以及访问级别
	self.accessFlags = reader.readUnit16()
	// 类名
	self.thisClass = reader.readUnit16()
	// 超类名
	self.superClass = reader.readUnit16()
	// 接口索引表
	self.interfaces = reader.readUnit16s()
	// 字段表
	self.fields = readMembers(reader, self.constantPool)
	// 方法表
	self.methods = readerMembers(reader, self.constantPool)
	self.attributes = readerAttributes(reader, self.constantPool)
} 
/*
	从常量池中查找超类名
 */
func(self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""	// Object 没有超类
}
/*
	从常量池查找接口名
 */
func (self *ClassFile) InterfaceNames() []string {
	interfacesNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfacesNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfacesNames
}
/*
	检查魔数
 */
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUnit32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError:magic!")	// 使用panic模拟抛出异常
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUnit16()
	self.majorVersion = reader.readUnit16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0{
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}