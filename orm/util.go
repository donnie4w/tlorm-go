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
	"errors"
	"reflect"
	"strings"
)

func isPointer(a any) bool {
	return reflect.TypeOf(a).Kind() == reflect.Pointer
}

func isStruct(a any) bool {
	return reflect.TypeOf(a).Kind() == reflect.Struct
}

func getObjectName(a any) (tname string) {
	t := reflect.TypeOf(a)
	if t.Kind() != reflect.Pointer {
		tname = strings.ToLower(t.Name())
	} else {
		tname = strings.ToLower(t.Elem().Name())
	}
	if tname == "" {
		panic("error: table name is empty")
	}
	return
}

func checkIndexField(field_name string, tag reflect.StructTag) (b bool) {
	return strings.HasSuffix(field_name, "_") || string(tag) == "idx" || tag.Get("idx") == "1"
}

func getBytesValueFromkind(f reflect.Value) (_v []byte, e error) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	isSet := false
	if f.CanInt() {
		_v, isSet = IntToBytes[int64](f.Int()), true
	} else if f.CanFloat() {
		_v, isSet = IntToBytes[float64](f.Float()), true
	} else if f.CanUint() {
		_v, isSet = IntToBytes[uint64](f.Uint()), true
	} else if f.Kind() == reflect.String {
		_v, isSet = []byte(f.String()), true
	} else if f.Kind() == reflect.Bool {
		i := byte(0)
		if f.Bool() {
			i = 1
		}
		_v, isSet = []byte{i}, true
	} else if f.Kind() == reflect.Pointer {
		switch f.Interface().(type) {
		case *int:
			_v, isSet = IntToBytes[int64](int64(*(*int)(f.UnsafePointer()))), true
		case *int8:
			_v, isSet = IntToBytes[int8](*(*int8)(f.UnsafePointer())), true
		case *int16:
			_v, isSet = IntToBytes[int16](*(*int16)(f.UnsafePointer())), true
		case *int32:
			_v, isSet = IntToBytes[int32](*(*int32)(f.UnsafePointer())), true
		case *int64:
			_v, isSet = IntToBytes[int64](*(*int64)(f.UnsafePointer())), true
		case *uint:
			_v, isSet = IntToBytes[uint64](uint64(*(*uint)(f.UnsafePointer()))), true
		case *uint16:
			_v, isSet = IntToBytes[uint16](*(*uint16)(f.UnsafePointer())), true
		case *uint32:
			_v, isSet = IntToBytes[uint32](*(*uint32)(f.UnsafePointer())), true
		case *uint64:
			_v, isSet = IntToBytes[uint64](*(*uint64)(f.UnsafePointer())), true
		case *float32:
			_v, isSet = IntToBytes[float32](*(*float32)(f.UnsafePointer())), true
		case *float64:
			_v, isSet = IntToBytes[float64](*(*float64)(f.UnsafePointer())), true
		case *string:
			_v, isSet = []byte(*(*string)(f.UnsafePointer())), true
		}
	} else if f.Kind() == reflect.Slice {
		switch f.Interface().(type) {
		case []uint8:
			_v, isSet = f.Bytes(), true
		}
	}
	if !isSet {
		e = errors.New("type error")
	}
	return
}

func tBeanToStruct[T any](id int64, dm map[string][]byte) (a *T, e error) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	a = new(T)
	if isPointer(a) {
		v := reflect.ValueOf(a).Elem()
		t := reflect.TypeOf(a).Elem()
		for i := 0; i < t.NumField(); i++ {
			if idxName := t.Field(i).Name; strings.ToLower(idxName) != "id" {
				if _v, ok := dm[idxName]; ok {
					f := v.FieldByName(idxName)
					setBytesValueFromkind(f, _v)
				}
			}
		}
		id_v := v.FieldByNameFunc(func(s string) bool {
			return strings.ToLower(s) == "id"
		})
		if id_v.Kind() == reflect.Pointer {
			id_v.Set(reflect.ValueOf(&id))
		} else {
			id_v.SetInt(id)
		}
	}
	return
}

func setBytesValueFromkind(f reflect.Value, bs []byte) (_v []byte, e error) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	if f.CanInt() {
		f.SetInt(BytesToInt[int64](bs))
	} else if f.CanFloat() {
		f.SetFloat(BytesToInt[float64](bs))
	} else if f.CanUint() {
		f.SetUint(BytesToInt[uint64](bs))
	} else if f.Kind() == reflect.String {
		f.SetString(string(bs))
	} else if f.Kind() == reflect.Bool {
		if bs[0] == 1 {
			f.SetBool(true)
		}
	} else if f.Kind() == reflect.Pointer {
		switch f.Interface().(type) {
		case *int:
			i := int(BytesToInt[int64](bs))
			f.Set(reflect.ValueOf(&i))
		case *int8:
			i := BytesToInt[int8](bs)
			f.Set(reflect.ValueOf(&i))
		case *int16:
			i := BytesToInt[int16](bs)
			f.Set(reflect.ValueOf(&i))
		case *int32:
			i := BytesToInt[int32](bs)
			f.Set(reflect.ValueOf(&i))
		case *int64:
			i := BytesToInt[int64](bs)
			f.Set(reflect.ValueOf(&i))
		case *uint:
			i := uint(BytesToInt[uint64](bs))
			f.Set(reflect.ValueOf(&i))
		case *uint16:
			i := BytesToInt[uint16](bs)
			f.Set(reflect.ValueOf(&i))
		case *uint32:
			i := BytesToInt[uint32](bs)
			f.Set(reflect.ValueOf(&i))
		case *uint64:
			i := BytesToInt[uint64](bs)
			f.Set(reflect.ValueOf(&i))
		case *float32:
			i := BytesToInt[float32](bs)
			f.Set(reflect.ValueOf(&i))
		case *float64:
			i := BytesToInt[float64](bs)
			f.Set(reflect.ValueOf(&i))
		case *string:
			s := string(bs)
			f.Set(reflect.ValueOf(&s))
		}
	} else if f.Kind() == reflect.Slice {
		switch f.Interface().(type) {
		case []uint8:
			f.SetBytes(bs)
		}
	}
	return
}
