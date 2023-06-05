package io

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

type alphaReader struct {
	reader io.Reader
}

func newAlphaReader(r io.Reader) *alphaReader {
	return &alphaReader{reader: r}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (n int, err error) {
	n, err = a.reader.Read(p)
	if err != nil {
		return 0, err
	}

	buf := make([]byte, n)
	var j int
	for i := 0; i < n; i++ {
		b := alpha(p[i])

		if b != 0 {
			buf[j] = b
			j++
		}
	}

	copy(p, buf)

	return j, nil
}

func TestHW(t *testing.T) {
	reader := strings.NewReader("Clear is better than clever")

	p := make([]byte, 6)

	for {
		// 将reader中的字符串发到字节数组p中
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n, string(p[:n]))
	}
}

func TestNewAndMake(t *testing.T) {
	b := make([]bool, 6)
	fmt.Printf("%v\n", b)

	n := new(bool)
	fmt.Printf("%v\n", n)
	fmt.Println(*n)
}

func TestCustomizeReader(t *testing.T) {
	r := strings.NewReader("Hello! It's 9am, where is the sun?")
	a := newAlphaReader(r)

	for true {
		b := make([]byte, 6)
		n, err := a.Read(b)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
		fmt.Println(string(b[:n]))
	}
}

func TestWriter(t *testing.T) {
	proverbs := []string{"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic"}

	var b bytes.Buffer
	for _, p := range proverbs {
		n, err := b.Write([]byte(p))
		if err != nil {
			os.Exit(1)
		}

		fmt.Println(n)
	}
	fmt.Println(b.String())
}
