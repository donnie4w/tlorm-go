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
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, t)
	return bytesBuffer.Bytes()
}
