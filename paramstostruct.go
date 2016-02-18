package paramstostruct

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Convert converts req.Form to the structure which has the json tags
func Convert(m map[string][]string, s interface{}) error {

	tagInfo := make(map[string]string)
	rt := reflect.TypeOf(s).Elem()
	for i := 0; i < rt.NumField(); i++ {
		if json := rt.Field(i).Tag.Get("json"); json != "" && json != "-" {
			sName := strings.ToLower(rt.Name())
			key := fmt.Sprintf("%s[%s]", sName, json)
			tagInfo[key] = rt.Field(i).Name
		}
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
	rv := reflect.ValueOf(obj).Elem().FieldByName(name)

	if !rv.IsValid() {
		// No such field => SKIP
		return nil
	}

	if !rv.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	val := reflect.ValueOf(value)
	switch rv.Type().Name() {
	case "int":
		intVal, err := strconv.Atoi(val.Interface().(string))
		if err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(intVal))
	case "string":
		rv.Set(val)
	default:
		return errors.New(fmt.Sprintf("Provided value type didn't match obj field type val_type: %v, rv_type: %v", val.Type(), rv.Type()))
	}

	return nil
}
