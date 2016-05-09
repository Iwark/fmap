package fmap

import "reflect"

func contains(list interface{}, element interface{}) (found bool) {
	listValue := reflect.ValueOf(list)
	for i := 0; i < listValue.Len(); i++ {
		if reflect.DeepEqual(listValue.Index(i).Interface(), element) {
			return true
		}
	}
	return false
}
