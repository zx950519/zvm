package classfile
/*
	常量池的定义，实质上是一个表
 */
type ConstantPool []ConstantInfo
/*
	读取常量池
 */
func readConstantPool(reader *ClassReader)  {
	cpCount := int(reader.readUnit16())
	cp := make([]ConstantPool, cpCount)
	for i := 1; i<cpCount; i++ {

	}
}

func (self ConstantPool) getConstantInfo(index uint16)  {

}

func (self ConstantPool) getNameAndType(index uint16) (string, string)  {

}

func (self ConstantPool) getClassName(index uint16) string {

}

func (self ConstantPool) getUtf8(index uint16) string {

}
