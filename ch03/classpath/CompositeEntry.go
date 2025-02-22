package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry{
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(className string)([]byte, Entry, error) {
	// 依次调用每个子路径的readClass方法
	for _, entry := range self  {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func(self CompositeEntry) String() string {
	strs := make([]string, len(self))
	// 依次调用每个子路径的String方法
	for i, entry := range self{
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
