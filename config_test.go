package appconfig

import (
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/fs"
)

type Example struct {
	One         string
	StringField string
	BoolField   bool
	Int32Field  int32
	Singleword  int
	// TODO: other types
	// TODO: nested struct
	// TODO: squashed struct
}

func TestLoad(t *testing.T) {
	content := `
string_field: from-file
bool_field: true
int_32_field: 2
singleword: 3
`
	f := fs.NewFile(t, t.Name(), fs.WithContent(content))

	target := Example{
		One:         "left-as-default",
		StringField: "default-1",
		Int32Field:  1,
	}
	opts := Options{
		Filename: f.Path(),
	}
	err := Load(&target, opts)
	assert.NilError(t, err)
	expected := Example{
		One:         "left-as-default",
		StringField: "from-file",
		BoolField:   true,
		Int32Field:  2,
		Singleword:  3,
	}
	assert.DeepEqual(t, target, expected)
}
