package EncodeAndDecodeStrings

import (
	"strconv"
	"strings"
)

const split = ','

type Codec struct{}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(len(strs)))
	sb.WriteByte(split)
	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteByte(split)
		sb.WriteString(s)
	}
	return sb.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) (data []string) {
	if len(strs) == 0 {
		return nil
	}
	var (
		position  = strings.IndexByte(strs, split)
		length, _ = strconv.Atoi(strs[:position])
	)
	position++
	data = make([]string, length)
	for i := 0; i < length; i++ {
		offset := position + strings.IndexByte(strs[position:], split)
		lenOfS, _ := strconv.Atoi(strs[position:offset])
		position = offset + 1 + lenOfS
		data[i] = strs[offset+1 : position]
	}
	return
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
