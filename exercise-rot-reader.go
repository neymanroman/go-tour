package main

import (
	"io"
	"os"
	"strings"
)

func rot13(b byte) byte {
	if (b >= 'a' && b <= 'm') || (b >= 'A' && b <= 'M') {
		b += 13
	} else if (b >= 'n' && b <= 'z') || (b >= 'N' && b <= 'Z') {
		b -= 13
	}

	return b
}

type rot13Reader struct {
	r io.Reader
}

func (r13Reader rot13Reader) Read(b []byte) (int, error) {
	n, err := r13Reader.r.Read(b)

	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
