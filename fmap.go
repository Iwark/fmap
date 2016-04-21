package fmap

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/huandu/xstrings"
)

var timeFormats = []string{
	"1/2/2006",
	"1/2/2006 15:4:5",
	"2006-1-2 15:4:5",
	"2006-1-2 15:4",
	"2006-1-2",
	"1-2",
	"15:4:5",
	"15:4",
	"15",
	"15:4:5 Jan 2, 2006 MST",
}

// const (
// 	// SnakeCase uses snake_case to convert
// 	SnakeCase = iota
// )

// // Case represents the case of form names
// var Case = SnakeCase

// ConvertToStruct converts req.Form to the struct which has the json tags
func ConvertToStruct(m map[string][]string, s interface{}) error {

	tagInfo := make(map[string]string)
	rt := reflect.TypeOf(s).Elem()
	sName := xstrings.ToSnakeCase(rt.Name())
	for i := 0; i < rt.NumField(); i++ {
		var key string
		if val := rt.Field(i).Tag.Get("fmap"); val != "" {
			if val == "-" {
				continue
			}
			key = fmt.Sprintf("%s[%s]", sName, val)
		} else {
			key = fmt.Sprintf("%s[%s]", sName, xstrings.ToSnakeCase(rt.Field(i).Name))
		}
		tagInfo[key] = rt.Field(i).Name
	}
	for k, v := range m {
		var err error
		// If the struct has the key, set value
		if name, ok := tagInfo[k]; ok {
			err = setField(s, name, v[0])
		} else {
			err = setField(s, k, v[0])
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
	default:
		return fmt.Errorf("Provided value type didn't match obj field type val_type: %v, rv_type: %v", val.Type(), rv.Type())
	}

	return nil
}

func parseTime(str string) (t time.Time, err error) {
	for _, format := range timeFormats {
		t, err = time.Parse(format, str)
		if err == nil {
			location := t.Location()
			if location.String() == "UTC" {
				location = time.Now().Location()
			}
			pt := []int{t.Second(), t.Minute(), t.Hour(), t.Day(), int(t.Month()), t.Year()}
			t = time.Date(pt[5], time.Month(pt[4]), pt[3], pt[2], pt[1], pt[0], 0, location)
			return
		}
	}
	err = errors.New("Can't parse string as time: " + str)
	return
}
