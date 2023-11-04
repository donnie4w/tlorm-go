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
	"errors"
	"reflect"
	"strings"

	. "github.com/donnie4w/tlcli-go/tlcli"
)

type Orm[T any] interface {
	Create() (err error)
	Insert(a any) (seq int64, err error)
	Update(a any) (err error)
	UpdateNonzero(a any) (err error)
	Delete(id int64) (err error)
	Drop() (err error)
	AlterTable() (err error)
	SelectId() (id int64, err error)
	SelectIdByIdx(columnName string, columnValue any) (id int64, err error)
	SelectById(id int64) (a *T, err error)
	SelectsByIdLimit(startId, limit int64) (as []*T, err error)
	SelectByIdx(columnName string, columnValue any) (a *T, err error)
	SelectAllByIdx(columnName string, columnValue any) (as []*T, err error)
	SelectByIdxLimit(startId, limit int64, columnName string, columnValue ...any) (as []*T, err error)
	DeleteBatch(ids ...int64) (err error)
	SelectByIdxDescLimit(columnName string, columnValue any, startId int64, limit int64) (as []*T, err error)
	SelectByIdxAscLimit(columnName string, columnValue any, startId int64, limit int64) (as []*T, err error)
	SelectIdByIdxSeq(columnName string, columnValue any, seq int64) (id int64, err error)
}

func NewConn(tls bool, addr string, auth string) (conn *Client, err error) {
	conn, err = NewConnect(tls, addr, auth)
	return
}

func Table[T any](conn *Client) Orm[T] {
	return source[T]{conn}
}

type source[T any] struct {
	conn *Client
}

func (this source[T]) Create() (err error) {
	var a T
	table_name := getObjectName(a)
	if columns, indexs, er := this.getFieldIndex(); er == nil {
		err = this.conn.CreateTable(table_name, columns, indexs)
	} else {
		err = er
	}
	return
}

func (this source[T]) getFieldIndex() (columns map[string]COLUMNTYPE, indexs []string, err error) {
	var a T
	hasId := false
	v := reflect.ValueOf(a)
	v.FieldByNameFunc(func(s string) bool {
		if len(s) == 2 && strings.ToLower(s) == "id" {
			hasId = true
			return true
		}
		return false
	})
	if !hasId {
		err = errors.New("Id not found")
		return
	}
	t := reflect.TypeOf(a)
	columns = make(map[string]COLUMNTYPE, 0)
	indexs = make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		if idxName := t.Field(i).Name; strings.ToLower(idxName) != "id" {
			field := v.FieldByName(idxName)
			columns[idxName] = fieldToColumnType(field)
			// columns = append(columns, idxName)
			if checkIndexField(idxName, t.Field(i).Tag) {
				indexs = append(indexs, idxName)
			}
		}
	}
	return
}

func (this source[T]) Insert(a any) (seq int64, err error) {
	if isPointer(a) {
		table_name := getObjectName(a)
		v := reflect.ValueOf(a).Elem()
		t := reflect.TypeOf(a).Elem()
		dm := make(map[string][]byte, 0)
		for i := 0; i < t.NumField(); i++ {
			if fName := t.Field(i).Name; strings.ToLower(fName) != "id" {
				f := v.FieldByName(fName)
				if idx_value, err := getBytesValueFromkind(f); err == nil {
					dm[fName] = idx_value
				}
			}
		}
		seq, err = this.conn.Insert(table_name, dm)
	} else {
		err = errors.New("insert object must be pointer")
	}
	return
}

func (this source[T]) Update(a any) (err error) {
	return this._update(a, false)
}

func (this source[T]) UpdateNonzero(a any) (err error) {
	return this._update(a, true)
}

func (this source[T]) _update(a any, nonzero bool) (err error) {
	if isPointer(a) {
		table_name := getObjectName(a)
		v := reflect.ValueOf(a).Elem()

		hasId := false
		id_v := v.FieldByNameFunc(func(s string) bool {
			if len(s) == 2 && strings.ToLower(s) == "id" {
				hasId = true
				return true
			}
			return false
		})
		if !hasId {
			err = errors.New("id not found")
			return
		}
		id := id_v.Int()
		t := reflect.TypeOf(a).Elem()
		dm := make(map[string][]byte, 0)
		for i := 0; i < t.NumField(); i++ {
			if fName := t.Field(i).Name; strings.ToLower(fName) != "id" {
				f := v.FieldByName(fName)
				if nonzero && f.IsZero() {
					continue
				}
				if idx_value, err := getBytesValueFromkind(f); err == nil && idx_value != nil {
					dm[fName] = idx_value
				}
			}
		}
		err = this.conn.Update(table_name, id, dm)
	} else {
		err = errors.New("update object must be pointer")
	}
	return
}

func (this source[T]) SelectId() (id int64, err error) {
	var a T
	table_name := getObjectName(a)
	id, err = this.conn.SelectId(table_name)
	return
}

func (this source[T]) SelectIdByIdx(columnName string, columnValue any) (id int64, err error) {
	var a T
	table_name := getObjectName(a)
	v := reflect.ValueOf(a)
	field := v.FieldByName(columnName)
	var bs []byte
	if bs, err = anyTobyte(field, columnValue); err == nil {
		id, err = this.conn.SelectIdByIdx(table_name, columnName, bs)
	}
	return
}

func (this source[T]) SelectById(id int64) (a *T, err error) {
	table_name := getObjectName(a)
	var db *DataBean
	if db, err = this.conn.SelectById(table_name, id); err == nil {
		a, err = tBeanToStruct[T](id, db.GetTBean())
	}
	return
}

func (this source[T]) SelectsByIdLimit(startId, limit int64) (as []*T, err error) {
	var t T
	table_name := getObjectName(t)
	var dblist []*DataBean
	if dblist, err = this.conn.SelectsByIdLimit(table_name, startId, limit); err == nil {
		as = make([]*T, 0)
		for _, db := range dblist {
			var a *T
			if a, err = tBeanToStruct[T](db.GetID(), db.GetTBean()); err == nil {
				as = append(as, a)
			}
		}
	}
	return
}

func (this source[T]) SelectByIdx(columnName string, columnValue any) (a *T, err error) {
	var _a T
	table_name := getObjectName(_a)
	v := reflect.ValueOf(_a)
	field := v.FieldByName(columnName)
	var bs []byte
	if bs, err = anyTobyte(field, columnValue); err == nil {
		var db *DataBean
		if db, err = this.conn.SelectByIdx(table_name, columnName, bs); err == nil {
			a, err = tBeanToStruct[T](db.GetID(), db.GetTBean())
		}
	}
	return
}

func (this source[T]) SelectAllByIdx(columnName string, columnValue any) (as []*T, err error) {
	var a T
	table_name := getObjectName(a)
	v := reflect.ValueOf(a)
	field := v.FieldByName(columnName)
	var bs []byte
	if bs, err = anyTobyte(field, columnValue); err == nil {
		var dblist []*DataBean
		if dblist, err = this.conn.SelectAllByIdx(table_name, columnName, bs); err == nil {
			as = make([]*T, 0)
			for _, db := range dblist {
				var a *T
				if a, err = tBeanToStruct[T](db.GetID(), db.GetTBean()); err == nil {
					as = append(as, a)
				}
			}
		}
	}
	return
}

func (this source[T]) SelectByIdxLimit(startId, limit int64, columnName string, columnValue ...any) (as []*T, err error) {
	var a T
	table_name := getObjectName(a)
	v := reflect.ValueOf(a)
	field := v.FieldByName(columnName)
	bss := make([][]byte, 0)
	for _, cv := range columnValue {
		var bs []byte
		if bs, err = anyTobyte(field, cv); err == nil {
			bss = append(bss, bs)
		}
	}
	if len(bss) > 0 {
		var dblist []*DataBean
		if dblist, err = this.conn.SelectByIdxLimit(table_name, columnName, bss, startId, limit); err == nil {
			as = make([]*T, 0)
			for _, db := range dblist {
				var a *T
				if a, err = tBeanToStruct[T](db.GetID(), db.GetTBean()); err == nil {
					as = append(as, a)
				}
			}
		}
	}
	return
}

func (this source[T]) Delete(id int64) (err error) {
	var a T
	table_name := getObjectName(a)
	return this.conn.Delete(table_name, id)
}

func (this source[T]) Drop() (err error) {
	var a T
	table_name := getObjectName(a)
	return this.conn.Drop(table_name)
}

func (this source[T]) AlterTable() (err error) {
	var a T
	table_name := getObjectName(a)
	if columns, indexs, er := this.getFieldIndex(); er == nil {
		err = this.conn.AlterTable(table_name, columns, indexs)
	} else {
		err = er
	}
	return
}

func (this source[T]) DeleteBatch(ids ...int64) (err error) {
	var a T
	table_name := getObjectName(a)
	return this.conn.DeleteBatch(table_name, ids...)
}

func (this source[T]) SelectByIdxDescLimit(columnName string, columnValue any, startId int64, limit int64) (as []*T, err error) {
	var a T
	table_name := getObjectName(a)
	v := reflect.ValueOf(a)
	field := v.FieldByName(columnName)
	var bs []byte
	if bs, err = anyTobyte(field, columnValue); err == nil {
		var dblist []*DataBean
		if dblist, err = this.conn.SelectByIdxDescLimit(table_name, columnName, bs, startId, limit); err == nil {
			as = make([]*T, 0)
			for _, db := range dblist {
				var a *T
				if a, err = tBeanToStruct[T](db.GetID(), db.GetTBean()); err == nil {
					as = append(as, a)
				}
			}
		}
	}
	return
}

func (this source[T]) SelectByIdxAscLimit(columnName string, columnValue any, startId int64, limit int64) (as []*T, err error) {
	var a T
	table_name := getObjectName(a)
	v := reflect.ValueOf(a)
	field := v.FieldByName(columnName)
	var bs []byte
	if bs, err = anyTobyte(field, columnValue); err == nil {
		var dblist []*DataBean
		if dblist, err = this.conn.SelectByIdxAscLimit(table_name, columnName, bs, startId, limit); err == nil {
			as = make([]*T, 0)
			for _, db := range dblist {
				var a *T
				if a, err = tBeanToStruct[T](db.GetID(), db.GetTBean()); err == nil {
					as = append(as, a)
				}
			}
		}
	}
	return
}

func (this source[T]) SelectIdByIdxSeq(columnName string, columnValue any, seq int64) (id int64, err error) {
	var a T
	table_name := getObjectName(a)
	v := reflect.ValueOf(a)
	field := v.FieldByName(columnName)
	var bs []byte
	if bs, err = anyTobyte(field, columnValue); err == nil {
		id, err = this.conn.SelectIdByIdxSeq(table_name, columnName, bs, seq)
	}
	return
}
