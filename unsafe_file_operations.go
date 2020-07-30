package files

import "io/ioutil"

func readFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

func writeFile(file string, data []byte) error {
	return ioutil.WriteFile(file, data, 0777)
}
