package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string	// 用于存放目录的绝对路径
}

func newDirEntry(path string) *DirEntry{
	absDir, err := filepath.Abs(path)	// 将参数转化为绝对路径
	if err != nil{
		panic(err)		// 转换失败则调用panic()函数终止程序执行
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string)([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)	// 先将目录和class文件名拼成一个完整的路径
	data, err := ioutil.ReadFile(fileName)				// 读取class文件的具体内容
	return data, self, err
}

func(self *DirEntry) String() string {
	return self.absDir
}