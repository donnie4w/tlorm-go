/*
*
* Copyright 2023 tldb Author. All Rights Reserved.
* email: donnie4w@gmail.com
* github.com/donnie4w/tldb
* github.com/donnie4w/tlorm-go
*
 */

package orm

import (
	"bytes"
	"encoding/binary"
)

func BytesToInt[T int64 | int32 | int16 | int8 | uint64 | uint32 | uint16 | uint8 | float64 | float32](bs []byte) (_r T) {
	bytesBuffer := bytes.NewBuffer(bs)
	binary.Read(bytesBuffer, binary.BigEndian, &_r)
	return
}

func IntToBytes[T int64 | int32 | int16 | int8 | uint64 | uint32 | uint16 | uint8 | float64 | float32](t T) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, t)
	return buf.Bytes()
}
