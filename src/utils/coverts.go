package utils

import (
	"bytes"
)

func String2Byte(str string) []byte{
	return bytes.NewBufferString(str).Bytes()
}

func Byte2String(b []byte) string{
	return string(b)
}

