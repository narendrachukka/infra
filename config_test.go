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

	StringFromEnv string
	BoolFromEnv   bool
	UintFromEnv   uint

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

string_from_env: from-file-2
bool_from_env: false
uint_from_env: 5
`
	f := fs.NewFile(t, t.Name(), fs.WithContent(content))

	t.Setenv("APPNAME_STRING_FROM_ENV", "from-env-1")
	t.Setenv("APPNAME_BOOL_FROM_ENV", "true")
	t.Setenv("APPNAME_UINT_FROM_ENV", "412")

	target := Example{
		One:         "left-as-default",
		StringField: "default-1",
		Int32Field:  1,
	}
	opts := Options{
		Filename:  f.Path(),
		EnvPrefix: "APPNAME",
	}
	err := Load(&target, opts)
	assert.NilError(t, err)
	expected := Example{
		One:           "left-as-default",
		StringField:   "from-file",
		BoolField:     true,
		Int32Field:    2,
		Singleword:    3,
		StringFromEnv: "from-env-1",
		BoolFromEnv:   true,
		UintFromEnv:   412,
	}
	assert.DeepEqual(t, target, expected)
}
