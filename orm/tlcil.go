/*
*
* Copyright 2023 tldb Author. All Rights Reserved.
* email: donnie4w@gmail.com
* github.com/donnie4w/tldb
* github.com/donnie4w/tlorm
*
 */

package orm

import (
	"github.com/donnie4w/tlcli-go/tlcli"
)

var defaultConn *tlcli.Client

func RegisterDefaultResource(tls bool, addr string, auth string) (err error) {
	defaultConn, err = tlcli.NewConnect(tls, addr, auth)
	return
}

func Create[T any]() (err error) {
	return Table[T](defaultConn).Create()
}

func Insert(a any) (seq int64, err error) {
	return Table[byte](defaultConn).Insert(a)
}

func Update(a any) (err error) {
	return Table[byte](defaultConn).Update(a)
}

func SelectById[T any](id int64) (a *T, err error) {
	return Table[T](defaultConn).SelectById(id)
}

func SelectsById[T any](startId, limit int64) (as []*T, err error) {
	return Table[T](defaultConn).SelectsById(startId, limit)
}

func SelectByIdx[T any](columnName string, columnValue []byte) (a *T, err error) {
	return Table[T](defaultConn).SelectByIdx(columnName, columnValue)
}

func SelectAllByIdx[T any](columnName string, columnValue []byte) (as []*T, err error) {
	return Table[T](defaultConn).SelectAllByIdx(columnName, columnValue)
}

func SelectByIdxLimit[T any](columnName string, columnValue [][]byte, startId, limit int64) (as []*T, err error) {
	return Table[T](defaultConn).SelectByIdxLimit(columnName, columnValue, startId, limit)
}

func Delete[T any](id int64) (err error) {
	return Table[T](defaultConn).Delete(id)
}

func Truncate[T any]() (err error) {
	return Table[T](defaultConn).Truncate()
}

func AlterTable[T any]() (err error) {
	return Table[T](defaultConn).AlterTable()
}
