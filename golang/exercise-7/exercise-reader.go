package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (n int, e error) {
	for i := range b {
		b[i] = 'A'
		n += 1
	}
	return
}

func main() {
	reader.Validate(MyReader{})
}
