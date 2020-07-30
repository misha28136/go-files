package files

//ReadFile thread-safe reading file
func (o *Operations) ReadFile(file string) ([]byte, error) {
	var b []byte
	var err error

	o.ReadLock(file, func() {
		b, err = readFile(file)
	})
	return b, err
}

//WriteFile thread-safe write file
func (o *Operations) WriteFile(file string, data []byte) error {
	var err error
	o.WriteLock(file, func() {
		err = writeFile(file, data)
	})
	return err
}
