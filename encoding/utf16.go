// UTF-16 to UTF-8 decoding
// Ref: https://go.dev/play/p/CDbNzxBB63Z
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var bin = bytes.NewReader([]byte{
	0x7b, 0x00, 0x22, 0x00, 0x6d, 0x00, 0x65, 0x00, 0x73, 0x00, 0x73, 0x00,
	0x61, 0x00, 0x67, 0x00, 0x65, 0x00, 0x22, 0x00, 0x3a, 0x00, 0x20, 0x00,
	0x22, 0x00, 0x48, 0x00, 0x65, 0x00, 0x6c, 0x00, 0x6c, 0x00, 0x6f, 0x00,
	0x2c, 0x00, 0x20, 0x00, 0x47, 0x00, 0x6f, 0x00, 0x6c, 0x00, 0x61, 0x00,
	0x6e, 0x00, 0x67, 0x00, 0x22, 0x00, 0x7d, 0x00, 0x0a, 0x00,
})

func main() {
	r := transform.NewReader(bin, unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder().Transformer)
	s := struct {
		Message string `json:"message"`
	}{}
	err := json.NewDecoder(r).Decode(&s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.Message)
}
