package fmap

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Person struct {
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Name        string
	Age         int    `fmap:"great_age"`
	Mobilephone string `fmap:"mobilephone"`
	Birthday    time.Time
	Admin       bool
}

func TestConvert(t *testing.T) {
	assert := assert.New(t)

	formValue := map[string][]string{
		"person[updated_at]":  {"2016-04-13 15:24"},
		"person[deleted_at]":  {"2016-04-13"},
		"person[name]":        {"Iwark"},
		"person[great_age]":   {"24"},
		"person[mobilephone]": {"ios", "android"},
		"person[hobby]":       {"playing the piano"},
		"not_person":          {"hoge", "fuga"},
		"birthday":            {"2016-05-18"},
		"admin":               {"true"},
	}

	result := &Person{}
	err := New().WithStructName().ConvertToStruct(formValue, result)
	assert.NoError(err)
	require.NotNil(t, result.DeletedAt)
	deleted := result.DeletedAt.Format("2006-01-02")
	assert.Equal("2016-04-13", deleted)
	updated := result.UpdatedAt.Format("2006-01-02 15:04")
	assert.Equal("2016-04-13 15:24", updated)
	assert.Equal("Iwark", result.Name)
	assert.Equal(24, result.Age)
	assert.Equal("ios,android", result.Mobilephone)
	birthday := result.Birthday.Format("2006-01-02")
	assert.NotEqual("2016-05-18", birthday)
	assert.NotEqual(true, result.Admin)

	err = New().ConvertToStruct(formValue, result)
	assert.NoError(err)

	birthday = result.Birthday.Format("2006-01-02")
	assert.Equal("2016-05-18", birthday)
	assert.Equal(true, result.Admin)
}
