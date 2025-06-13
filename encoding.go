package sbi

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

type SbiIE interface{} //any SBI data models

// for now we use JSON
func Encode(ie SbiIE) (int64, io.Reader, error) {
	if buf, err := json.Marshal(ie); err != nil {
		return -1, nil, err
	} else {
		return int64(len(buf)), bytes.NewBuffer(buf), nil
	}
}

func Decode(length int64, r io.Reader, ie SbiIE) error {
	if buf, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		return json.Unmarshal(buf, ie)
	}
}
