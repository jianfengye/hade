package formatter

import (
	"bytes"
	"fmt"
)

func TextFormatter(msg string, fields []interface{}) ([]byte, error) {
	bf := bytes.NewBuffer([]byte(msg))
	bf.Write([]byte{':'})
	bf.WriteString(fmt.Sprint(fields))
	return bf.Bytes(), nil
}
