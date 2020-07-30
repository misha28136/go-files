package files

import (
	"testing"
)

var f *Files = New()

func TestRead(t *testing.T) {
	b, err := f.ReadFile("./tests/test_file.txt")
	if err != nil {
		t.Error(err)
	}
	if string(b) == "test_file data" {
		t.Log("OK, read data: ", string(b))
	} else {
		t.Error("filr read data damage: Expected: `test_file data` Received: `" + string(b) + "`")
	}
}

func TestThreadSafeFile(t *testing.T) {
	r := func() {
		i := 0
		for {
			if i > 200 {
				t.Log("Read worker: ok")
				break
			}
			b, err := f.ReadFile("./tests/test_write_file.txt")
			if err != nil {
				t.Error("Worker read", err)
			}
			if string(b) == "test_write_file" {
				t.Log("OK, read data: ", string(b))
			} else {
				t.Error("filr read data damage: Expected: `test_write_file` Received: `" + string(b) + "`")
			}
			i++
			exit
		}
	}

	w := func() {
		i := 0
		for {
			if i > 500 {
				t.Log("Writer worker: ok")
				break
			}
			err := f.WriteFile("/tests/test_write_file.txt", []byte("test_write_file"))
			if err != nil {
				t.Error("Worker write", err)
			}
			i++
		        exit
		}
	}

	go r()
	go w()
	w()

}
