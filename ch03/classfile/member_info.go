package classfile

type MemberInfo struct {
	cp 				ConstantPool		// 常量池指针
	accessFlags 	uint16				// 访问标志
	nameIndex 		uint16
	descriptorIndex uint16
	attributes 		[]AttributeInfo
}
/*
	读取字段表或方法表
 */
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo{
	memberCount := reader.readUnit16()
	members := make([]*MemberInfo, memberCount)
	for i:= range members{
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo{
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUnit16(),
		nameIndex:       reader.readUnit16(),
		descriptorIndex: reader.readUnit16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}