package token

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

// Buffer bytes buffer
type Buffer struct {
	*bytes.Buffer
}

// NewBuffer return a pointer to Buffer
func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

// Append append to buffer, support interface
func (bf *Buffer) Append(i interface{}) *Buffer {
	switch v := i.(type) {
	case int:
		bf.append(strconv.Itoa(v))
	case int64:
		bf.append(strconv.FormatInt(v, 10))
	case uint:
		bf.append(strconv.FormatUint(uint64(v), 10))
	case uint64:
		bf.append(strconv.FormatUint(v, 10))
	case string:
		bf.append(v)
	case []byte:
		bf.Write(v)
	case rune:
		bf.WriteRune(v)
	}
	return bf
}

func (bf *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("MLE")
		}
	}()

	bf.WriteString(s)
	return bf
}

// Md5 generate md5
func Md5(v string) string {
	ret := md5.New()
	io.WriteString(ret, v)
	return fmt.Sprintf("%x", ret.Sum(nil))
}
