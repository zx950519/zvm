package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)	// 用于存放路径分隔符的常量

/*
	表示类路径的接口
	该接口有四个实现
 */
type Entry interface {
	readClass(className string)([]byte, Entry, error)	// 寻找和加载class文件
	String() string										// toString()方法
}

func newEntry(path string) Entry{
	if strings.Contains(path, pathListSeparator){
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}