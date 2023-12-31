package dynamic

import (
	"errors"
	"fmt"
	"reflect"
)

type StructFieldIterator func(tag string, ft reflect.StructField, fv reflect.Value) error

var ErrCanNotIterateNilPointer = errors.New("can not iterate struct on a nil pointer")

func IterateFields(obj interface{}, cb func(ft reflect.StructField, fv reflect.Value) error) error {
	if obj == nil {
		return errors.New("can not iterate field, given object is nil")
	}

	sv := reflect.ValueOf(obj)
	st := reflect.TypeOf(obj)

	if st.Kind() != reflect.Ptr {
		return fmt.Errorf("f should be a pointer of a struct, %s given", st)
	}

	// for pointer, check if it's nil
	if sv.IsNil() {
		return ErrCanNotIterateNilPointer
	}

	// solve the reference
	st = st.Elem()
	sv = sv.Elem()

	if st.Kind() != reflect.Struct {
		return fmt.Errorf("f should be a struct, %s given", st)
	}

	for i := 0; i < sv.NumField(); i++ {
		fv := sv.Field(i)
		ft := st.Field(i)

		// skip unexported fields
		if !st.Field(i).IsExported() {
			continue
		}

		if err := cb(ft, fv); err != nil {
			return err
		}
	}

	return nil
}

func isStructPtr(tpe reflect.Type) bool {
	return tpe.Kind() == reflect.Ptr && tpe.Elem().Kind() == reflect.Struct
}

func IterateFieldsByTag(obj interface{}, tagName string, children bool, cb StructFieldIterator) error {
	sv := reflect.ValueOf(obj)
	st := reflect.TypeOf(obj)

	if st.Kind() != reflect.Ptr {
		return fmt.Errorf("obj should be a pointer of a struct, %s given", st)
	}

	// for pointer, check if it's nil
	if sv.IsNil() {
		return ErrCanNotIterateNilPointer
	}

	// solve the reference
	st = st.Elem()
	sv = sv.Elem()

	if st.Kind() != reflect.Struct {
		return fmt.Errorf("f should be a struct, %s given", st)
	}

	for i := 0; i < sv.NumField(); i++ {
		fv := sv.Field(i)
		ft := st.Field(i)

		// skip unexported fields
		if !st.Field(i).IsExported() {
			continue
		}

		if children && isStructPtr(ft.Type) && !fv.IsNil() {
			// recursive iterate the struct field
			if err := IterateFieldsByTag(fv.Interface(), tagName, false, cb); err != nil {
				return fmt.Errorf("unable to iterate struct fields over the type %v: %v", ft, err)
			}
		}

		tag, ok := ft.Tag.Lookup(tagName)
		if !ok {
			continue
		}

		// call the iterator
		if err := cb(tag, ft, fv); err != nil {
			return err
		}
	}

	return nil
}
