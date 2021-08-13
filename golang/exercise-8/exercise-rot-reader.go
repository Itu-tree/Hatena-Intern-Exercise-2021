package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot13.r.Read(p)
	for i, v := range p {

		switch {
		case v >= 'A' && v <= 'Z':
			p[i] = byte('A') + (v-byte('A')+13)%26
		case v >= 'a' && v <= 'z':
			p[i] = byte('a') + (v-byte('a')+13)%26
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
