package fmap

import "testing"

type Person struct {
	Name   string
	Age    int    `fmap:"age"`
	Gender string `fmap:"gender"`
}

func TestConvert(t *testing.T) {
	formValue := map[string][]string{
		"person[name]":   []string{"Iwark"},
		"person[age]":    []string{"24"},
		"person[gender]": []string{"man", "woman"},
		"person[hobby]":  []string{"playing the piano"},
		"not_person":     []string{"hoge", "fuga"},
	}

	result := &Person{}
	err := ConvertToStruct(formValue, result)
	if err != nil {
		t.Error("Convert Error:", err)
	}
	if result.Name != "" || result.Age != 24 || result.Gender != "man" {
		t.Error("Map to struct failed, got:", result)
	}
}
