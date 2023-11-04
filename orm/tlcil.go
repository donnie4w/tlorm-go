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

// update data for non nil
func Update(a any) (err error) {
	return Table[byte](defaultConn).Update(a)
}

// update data for non zero
func UpdateNonzero(a any) (err error) {
	return Table[byte](defaultConn).UpdateNonzero(a)
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

func DeleteBatch[T any](ids ...int64) (err error) {
	return Table[T](defaultConn).DeleteBatch(ids...)
}
//Index fields that are updated frequently are not suitable for this method and may result in sorting errors
//频繁更新的索引字段不适合此方法，并且可能导致排序错误
func SelectByIdxDescLimit[T any](columnName string, columnValue any, startId int64, limit int64) (as []*T, err error) {
	return Table[T](defaultConn).SelectByIdxDescLimit(columnName, columnValue, startId, limit)
}
//Index fields that are updated frequently are not suitable for this method and may result in sorting errors
////频繁更新的索引字段不适合此方法，并且可能导致排序错误
func SelectByIdxAscLimit[T any](columnName string, columnValue any, startId int64, limit int64) (as []*T, err error) {
	return Table[T](defaultConn).SelectByIdxAscLimit(columnName, columnValue, startId, limit)
}
