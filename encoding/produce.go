// Produce a compressed string using specified encoding/compression schemes, and encode it with std base64.
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	original := "test"
	var buffer bytes.Buffer
	w := gzip.NewWriter(&buffer)
	w.Write([]byte(original))
	w.Close()

	r, _ := gzip.NewReader(bytes.NewReader(buffer.Bytes()))
	output, _ := io.ReadAll(r)
	fmt.Println(string(output))

	result := base64.StdEncoding.EncodeToString(buffer.Bytes())
	fmt.Println(result)
}
