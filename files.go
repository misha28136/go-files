package files

import (
	"sync"

	"github.com/cornelk/hashmap"
)

//Files thread-safe struct package
type Files struct {
	operation *hashmap.HashMap
}

//New created struct Files
func New() *Files {
	return &Files{
		operation: &hashmap.HashMap{},
	}
}

//File contains a mutex for each file the package worked with
type File struct {
	sync.RWMutex
}

//ReadFile thread-safe reading file
func (f *Files) ReadFile(file string) ([]byte, error) {
	amount, ok := f.operation.Get(file)
	if ok {
		q := amount.(*File)
		q.RLock()
		b, err := readFile(file)
		q.RUnlock()
		return b, err
	}
	q := new(File)
	f.operation.Set(file, q)
	return f.ReadFile(file)

}

//WriteFile thread-safe write file
func (f *Files) WriteFile(file string, data []byte) error {
	amount, ok := f.operation.Get(file)
	if ok {
		q := amount.(*File)
		q.Lock()
		err := writeFile(file, data)
		q.Unlock()
		return err
	}
	q := new(File)
	f.operation.Set(file, q)
	return f.WriteFile(file, data)

}
