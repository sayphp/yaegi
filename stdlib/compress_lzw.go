package stdlib

// Generated by 'goexports compress/lzw'. Do not edit!

import (
	"compress/lzw"
	"reflect"
)

func init() {
	Value["compress/lzw"] = map[string]reflect.Value{
		"LSB":       reflect.ValueOf(lzw.LSB),
		"MSB":       reflect.ValueOf(lzw.MSB),
		"NewReader": reflect.ValueOf(lzw.NewReader),
		"NewWriter": reflect.ValueOf(lzw.NewWriter),
	}
	Type["compress/lzw"] = map[string]reflect.Type{
		"Order": reflect.TypeOf((*lzw.Order)(nil)).Elem(),
	}
}