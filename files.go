package files

import (
	"sync"

	"github.com/cornelk/hashmap"
)

//Files thread-safe struct package
type Files struct {
	cache *cache
}

type cache struct {
	is      bool
	storage *hashmap.HashMap
}

//New created struct Files
func New() *Files {
	return &Files{
		cache: &cache{
			is:      false,
			storage: &hashmap.HashMap{},
		},
	}
}

//File contains a mutex for each file the package worked with
type File struct {
	sync.RWMutex
	data []byte
}

//ReadFile thread-safe reading file
func (f *Files) ReadFile(file string) ([]byte, error) {
	amount, ok := f.cache.storage.Get(file)
	if ok {
		q := amount.(*File)
		q.RLock()
		b, err := readFile(file)
		q.RUnlock()
		return b, err
	}
	q := new(File)
	f.cache.storage.Set(file, q)
	return f.ReadFile(file)

}

//WriteFile thread-safe write file
func (f *Files) WriteFile(file string, data []byte) error {
	amount, ok := f.cache.storage.Get(file)
	if ok {
		q := amount.(*File)
		q.Lock()
		err := writeFile(file, data)
		q.Unlock()
		return err
	}
	q := new(File)
	f.cache.storage.Set(file, q)
	return f.WriteFile(file, data)

}
