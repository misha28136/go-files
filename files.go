package files

import (
	"sync"

	"github.com/cornelk/hashmap"
)

type Files struct {
	cache *cache
}

type cache struct {
	is      bool
	storage *hashmap.HashMap
}

func New() *Files {
	return &Files{
		cache: &cache{
			is:      false,
			storage: &hashmap.HashMap{},
		},
	}
}

type File struct {
	sync.RWMutex
	data []byte
}

func (f *Files) ReadFile(file string) ([]byte, error) {
	amount, ok := f.cache.storage.Get(file)
	if ok {
		q := amount.(*File)
		q.RLock()
		b, err := readFile(file)
		q.RUnlock()
		return b, err
	} else {
		q := new(File)
		f.cache.storage.Set(file, q)
		return f.ReadFile(file)
	}
}
