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
	return string(tag) == "idx" || tag.Get("idx") == "1"
}

func getBytesValueFromkind(f reflect.Value) (_v []byte, e error) {
	defer recover()
	isSet := false
	switch f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		_v, isSet = IntToBytes[int64](f.Int()), true
	case reflect.Float32, reflect.Float64:
		_v, isSet = IntToBytes[float64](f.Float()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		_v, isSet = IntToBytes[uint64](f.Uint()), true
	case reflect.String:
		_v, isSet = []byte(f.String()), true
	case reflect.Bool:
		i := byte(0)
		if f.Bool() {
			i = 1
		}
		_v, isSet = []byte{i}, true
	case reflect.Pointer:
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
	case reflect.Slice:
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

func anyTobyte(f reflect.Value, v any) (_v []byte, e error) {
	defer recover()
	isSet := false
	switch f.Kind() {
	case reflect.Bool:
		if v.(bool) {
			_v, isSet = []byte{1}, true
		} else {
			_v, isSet = []byte{0}, true
		}
	case reflect.Int:
		_v, isSet = IntToBytes[int64](int64(v.(int))), true
	case reflect.Int8:
		_v, isSet = IntToBytes[int8](int8(v.(int8))), true
	case reflect.Int16:
		_v, isSet = IntToBytes[int16](int16(v.(int16))), true
	case reflect.Int32:
		_v, isSet = IntToBytes[int32](int32(v.(int32))), true
	case reflect.Int64:
		_v, isSet = IntToBytes[int64](int64(v.(int64))), true
	case reflect.Uint:
		_v, isSet = IntToBytes[uint64](uint64(v.(uint))), true
	case reflect.Uint8:
		_v, isSet = IntToBytes[uint8](uint8(v.(uint8))), true
	case reflect.Uint16:
		_v, isSet = IntToBytes[uint16](uint16(v.(uint16))), true
	case reflect.Uint32:
		_v, isSet = IntToBytes[uint32](uint32(v.(uint32))), true
	case reflect.Uint64:
		_v, isSet = IntToBytes[uint64](uint64(v.(uint64))), true
	case reflect.Float32:
		_v, isSet = IntToBytes[float32](float32(v.(float32))), true
	case reflect.Float64:
		_v, isSet = IntToBytes[float64](float64(v.(float64))), true
	case reflect.String:
		_v, isSet = []byte(v.(string)), true
	case reflect.Pointer:
		switch f.Interface().(type) {
		case *int:
			_v, isSet = IntToBytes[int64](int64(*(v.(*int)))), true
		case *int8:
			_v, isSet = IntToBytes[int8](*(v.(*int8))), true
		case *int16:
			_v, isSet = IntToBytes[int16](*(v.(*int16))), true
		case *int32:
			_v, isSet = IntToBytes[int32](*(v.(*int32))), true
		case *int64:
			_v, isSet = IntToBytes[int64](*(v.(*int64))), true
		case *uint:
			_v, isSet = IntToBytes[uint64](uint64(*(v.(*uint)))), true
		case *uint16:
			_v, isSet = IntToBytes[uint16](*(v.(*uint16))), true
		case *uint32:
			_v, isSet = IntToBytes[uint32](*(v.(*uint32))), true
		case *uint64:
			_v, isSet = IntToBytes[uint64](*(v.(*uint64))), true
		case *float32:
			_v, isSet = IntToBytes[float32](*(v.(*float32))), true
		case *float64:
			_v, isSet = IntToBytes[float64](*(v.(*float64))), true
		case *string:
			_v, isSet = []byte(*(v.(*string))), true
		}
	case reflect.Slice:
		switch f.Interface().(type) {
		case []uint8:
			_v, isSet = v.([]byte), true
		}
	}

	if !isSet {
		e = errors.New("type error")
	}
	return
}

func tBeanToStruct[T any](id int64, dm map[string][]byte) (a *T, e error) {
	defer recover()
	if dm != nil {
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
	}
	return
}

func setBytesValueFromkind(f reflect.Value, bs []byte) (e error) {
	defer recover()
	switch f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		f.SetInt(BytesToInt[int64](bs))
	case reflect.Float32, reflect.Float64:
		f.SetFloat(BytesToInt[float64](bs))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		f.SetUint(BytesToInt[uint64](bs))
	case reflect.String:
		f.SetString(string(bs))
	case reflect.Bool:
		if bs[0] == 1 {
			f.SetBool(true)
		}
	case reflect.Pointer:
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
	case reflect.Slice:
		switch f.Interface().(type) {
		case []uint8:
			f.SetBytes(bs)
		}
	}
	return
}
