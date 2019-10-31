package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)
type ZipEntry struct {
	absPath string	// 用于存放目录的绝对路径
}

func newZipEntry(path string) *ZipEntry{
	absDir, err := filepath.Abs(path)	// 将参数转化为绝对路径
	if err != nil{
		panic(err)		// 转换失败则调用panic()函数终止程序执行
	}
	return &ZipEntry{absDir}
}

func (self *ZipEntry) readClass(className string)([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)	// 尝试打开zip文件
	if err != nil{
		return nil, nil, err
	}
	defer r.Close()		// 确保打开的文件被关闭
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()		// 确保打开的文件被关闭
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func(self *ZipEntry) String() string {
	return self.absPath
}