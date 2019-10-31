package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]		// 去掉路径末尾的*
	compositeEntry := []Entry{}
	/*
		在walkFn中，根据后缀名选出JAR文件
	 */
	walkFn := func(path string, info os.FileInfo, err error) error{
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)		// 遍历baseDir创建ZipEntry
	return compositeEntry
}
