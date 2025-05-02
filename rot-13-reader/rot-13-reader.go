package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(u uint8) uint8 {
	ru := u

	if u >= 65 && u <= 90 {
		ru = ru + uint8(13)
		if ru > 90 {
			ru = 64 + (ru - 90)
		}
	} else if u >= 97 && u <= 122 {
		ru = ru + uint8(13)
		if ru > 122 {
			ru = 96 + (ru - 122)
		}
	}

	return ru
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)

	for i, v := range b {
		b[i] = rot13(v)
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
