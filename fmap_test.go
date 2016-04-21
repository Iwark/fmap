package fmap

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Person struct {
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Age       int    `fmap:"great_age"`
	Gender    string `fmap:"gender"`
	Birthday  time.Time
}

func TestConvert(t *testing.T) {
	assert := assert.New(t)

	formValue := map[string][]string{
		"person[updated_at]": []string{"2016-04-13 15:24"},
		"person[deleted_at]": []string{"2016-04-13"},
		"person[name]":       []string{"Iwark"},
		"person[great_age]":  []string{"24"},
		"person[gender]":     []string{"man", "woman"},
		"person[hobby]":      []string{"playing the piano"},
		"not_person":         []string{"hoge", "fuga"},
	}

	result := &Person{}
	err := ConvertToStruct(formValue, result)
	assert.NoError(err)
	require.NotNil(t, result.DeletedAt)
	deleted := result.DeletedAt.Format("2006-01-02")
	assert.Equal("2016-04-13", deleted)
	updated := result.UpdatedAt.Format("2006-01-02 15:04")
	assert.Equal("2016-04-13 15:24", updated)
	assert.Equal("Iwark", result.Name)
	assert.Equal(24, result.Age)
	assert.Equal("man", result.Gender)
}
