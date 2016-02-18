package paramstostruct

import "testing"

type Person struct {
	Name   string `json:"-"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func TestConvert(t *testing.T) {
	params := map[string][]string{
		"person[name]":   []string{"Iwark"},
		"person[age]":    []string{"24"},
		"person[gender]": []string{"man"},
		"person[hobby]":  []string{"playing the piano"},
	}

	result := &Person{}
	err := Convert(params, result)
	if err != nil {
		t.Error("Convert Error:", err)
	}
	if result.Name != "" || result.Age != 24 || result.Gender != "man" {
		t.Error("Map to struct failed, got:", result)
	}
}
