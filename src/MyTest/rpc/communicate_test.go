package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func BenchmarkStr1(b *testing.B) {
	//serial := make([]byte,4,23)
	var pkg = new(bytes.Buffer)
	binary.Write(pkg, binary.BigEndian, uint32(len(l1))+4)
	fmt.Println(pkg.Bytes())
	//binary.Write(pkg, binary.BigEndian, serial)
	//fmt.Println(pkg.Bytes())
	//binary.Write(pkg, binary.BigEndian, l1)
	//fmt.Println(pkg.Bytes())
}

func BenchmarkStr2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sli := []byte(l4)
		_ = string(sli)
	}
}
