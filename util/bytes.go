package util

import (
	"reflect"
	"unsafe"
)

func ByteSliceToString(bs []byte) string {
	bh := *(*reflect.SliceHeader)(unsafe.Pointer(&bs))
	sh := reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}
