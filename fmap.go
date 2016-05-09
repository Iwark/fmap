package fmap

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/huandu/xstrings"
)

// const (
// 	// SnakeCase uses snake_case to convert
// 	SnakeCase = iota
// )

// // Case represents the case of form names
// var Case = SnakeCase

// Converter converts the form values to the struct
type Converter struct {
	withStructName bool
}

// New creates a Converter
func New() *Converter {
	return &Converter{}
}

// WithStructName set whether the form key has a struct name or not
// when set to true, the input name should be like: user[name]
// when set to false, the input name should be like: name
// defaults to true
func (c *Converter) WithStructName() *Converter {
	c.withStructName = true
	return c
}

// ConvertToStruct converts req.Form to the struct which has the json tags
func (c *Converter) ConvertToStruct(m url.Values, s interface{}) error {
	tagInfo := make(map[string]string)
	rt := reflect.TypeOf(s).Elem()
	sName := xstrings.ToSnakeCase(rt.Name())
	for i := 0; i < rt.NumField(); i++ {
		var key string
		if val := rt.Field(i).Tag.Get("fmap"); val != "" {
			if val == "-" {
				continue
			}
			key = val
		} else {
			key = xstrings.ToSnakeCase(rt.Field(i).Name)
		}
		if c.withStructName {
			key = fmt.Sprintf("%s[%s]", sName, key)
		}
		tagInfo[key] = rt.Field(i).Name
	}
	for k := range m {
		var err error
		// If the struct has the key, set value
		if name, ok := tagInfo[k]; ok {
			err = setField(s, name, m.Get(k))
		} else {
			err = setField(s, k, m.Get(k))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func setField(obj interface{}, name string, value interface{}) error {
	rv := reflect.Indirect(reflect.ValueOf(obj)).FieldByName(name)
	if !rv.IsValid() {
		// No such field => SKIP
		return nil
	}
	if !rv.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	val := reflect.ValueOf(value)
	switch rv.Type() {
	case reflect.TypeOf(int(0)):
		intVal, err := strconv.Atoi(val.Interface().(string))
		if err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(intVal))
	case reflect.TypeOf(int32(0)):
		intVal, err := strconv.Atoi(val.Interface().(string))
		if err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(int32(intVal)))
	case reflect.TypeOf(int64(0)):
		intVal, err := strconv.Atoi(val.Interface().(string))
		if err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(int64(intVal)))
	case reflect.TypeOf(uint(0)):
		intVal, err := strconv.Atoi(val.Interface().(string))
		if err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(uint(intVal)))
	case reflect.TypeOf(uint32(0)):
		intVal, err := strconv.Atoi(val.Interface().(string))
		if err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(uint32(intVal)))
	case reflect.TypeOf(uint64(0)):
		intVal, err := strconv.Atoi(val.Interface().(string))
		if err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(uint64(intVal)))
	case reflect.TypeOf(""):
		rv.Set(val)
	case reflect.TypeOf(time.Time{}):
		t, err := parseTime(val.String())
		if err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(t))
	case reflect.TypeOf(&time.Time{}):
		t, err := parseTime(val.String())
		if err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(&t))
	case reflect.TypeOf(true):
		if val.String() == "true" {
			rv.SetBool(true)
		} else {
			rv.SetBool(false)
		}
	default:
		return fmt.Errorf("Provided value type didn't match obj field type val_type: %v, rv_type: %v", val.Type(), rv.Type())
	}

	return nil
}
