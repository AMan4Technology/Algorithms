package EncodeAndDecodeStrings

import (
	"fmt"
	"testing"
)

func TestCodec(t *testing.T) {
	a := []string{"hello", "world"}
	c := new(Codec)
	fmt.Println(c.Encode(a))
	fmt.Println(c.Decode(c.Encode(a)))
}
