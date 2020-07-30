package files

import "github.com/v-grabko1999/operations"

type Files struct {
	op *operations.Operations
}

func New() *Files {
	return &Files{
		op: operations.New(),
	}
}

//ReadFile thread-safe reading file
func (f *Files) ReadFile(file string) ([]byte, error) {
	var b []byte
	var err error

	f.op.ReadLock(file, func() {
		b, err = readFile(file)
	})
	return b, err
}

//WriteFile thread-safe write file
func (f *Files) WriteFile(file string, data []byte) error {
	var err error
	f.op.WriteLock(file, func() {
		err = writeFile(file, data)
	})
	return err
}
