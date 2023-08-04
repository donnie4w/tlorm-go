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

func SelectId[T any]() (id int64, err error) {
	return Table[T](defaultConn).SelectId()
}

func SelectIdByIdx[T any](columnName string, columnValue any) (id int64, err error) {
	return Table[T](defaultConn).SelectIdByIdx(columnName, columnValue)
}

func SelectById[T any](id int64) (a *T, err error) {
	return Table[T](defaultConn).SelectById(id)
}

func SelectsByIdLimit[T any](startId, limit int64) (as []*T, err error) {
	return Table[T](defaultConn).SelectsByIdLimit(startId, limit)
}

func SelectByIdx[T any](columnName string, columnValue any) (a *T, err error) {
	return Table[T](defaultConn).SelectByIdx(columnName, columnValue)
}

func SelectAllByIdx[T any](columnName string, columnValue any) (as []*T, err error) {
	return Table[T](defaultConn).SelectAllByIdx(columnName, columnValue)
}

func SelectByIdxLimit[T any](startId, limit int64, columnName string, columnValue ...any) (as []*T, err error) {
	return Table[T](defaultConn).SelectByIdxLimit(startId, limit, columnName, columnValue...)
}

func Delete[T any](id int64) (err error) {
	return Table[T](defaultConn).Delete(id)
}

func Drop[T any]() (err error) {
	return Table[T](defaultConn).Drop()
}

func AlterTable[T any]() (err error) {
	return Table[T](defaultConn).AlterTable()
}
