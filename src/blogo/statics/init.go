package statics

import (
	"encoding/base64"
	"fmt"
)

var Bytes = map[string][]byte{}

func init() {
	for k, v := range Files {
		d, err := base64.StdEncoding.DecodeString(v)
		if nil != err {
			fmt.Println("Error decoding file '%s': %s", k, err)
		} else {
			Bytes[k] = d
		}
	}
}
