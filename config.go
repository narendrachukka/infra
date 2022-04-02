package appconfig

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"github.com/mitchellh/reflectwalk"
	"gopkg.in/yaml.v2"
	"gotest.tools/v3/env"
)

type Options struct {
	FieldTagName string
	Filename     string
	EnvPrefix    string
}

func Load(target interface{}, opts Options) error {
	if opts.FieldTagName == "" {
		opts.FieldTagName = "config"
	}
	if opts.Filename != "" {
		if err := loadFromFile(target, opts); err != nil {
			return err
		}
	}
	if opts.EnvPrefix != "" {
		if err := loadFromEnv(target, opts); err != nil {
			return err
		}
	}
	return nil
}

func loadFromFile(target interface{}, opts Options) error {
	fh, err := os.Open(opts.Filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	var raw map[string]interface{}
	if err := yaml.NewDecoder(fh).Decode(&raw); err != nil {
		return fmt.Errorf("failed to decode yaml from %s: %w", opts.Filename, err)
	}

	cfg := mapstructure.DecoderConfig{
		// TODO: Squash:           false,
		Result:  target,
		TagName: opts.FieldTagName,
		MatchName: func(key string, fieldName string) bool {
			return strings.EqualFold(key, strcase.ToSnake(fieldName))
		},
	}
	decoder, err := mapstructure.NewDecoder(&cfg)
	if err != nil {
		return fmt.Errorf("failed to create decoder: %w", err)
	}
	if err := decoder.Decode(raw); err != nil {
		return fmt.Errorf("failed to decode from %s: %w", opts.Filename, err)
	}

	return nil
}

func loadFromEnv(target interface{}, opts Options) error {
	walker := &envWalker{
		// TODO: probably copy this instead of import
		// TODO: filter vars to only those that start with EnvPrefix
		vars:     env.ToMap(os.Environ()),
		location: []string{opts.EnvPrefix},
	}

	if err := reflectwalk.Walk(target, walker); err != nil {
		return fmt.Errorf("failed to decode from environment variables: %w", err)
	}
	return nil
}

type envWalker struct {
	opts     Options
	location []string
	vars     map[string]string
}

func (w *envWalker) Enter(reflectwalk.Location) error {
	return nil
}

func (w *envWalker) Exit(loc reflectwalk.Location) error {
	if loc == reflectwalk.Struct && len(w.location) > 0 {
		w.location = w.location[:len(w.location)-1]
	}
	return nil
}

func (w *envWalker) Struct(value reflect.Value) error {
	cfg := mapstructure.DecoderConfig{
		Result:           value.Addr().Interface(),
		TagName:          w.opts.FieldTagName,
		WeaklyTypedInput: true,
		MatchName: func(key string, fieldName string) bool {
			path := strings.Join(w.location, "_")
			name := strings.ToUpper(path + "_" + strcase.ToSnake(fieldName))
			return key == name
		},
	}
	decoder, err := mapstructure.NewDecoder(&cfg)
	if err != nil {
		return fmt.Errorf("failed to create decoder for struct: %w", err)
	}
	if err := decoder.Decode(w.vars); err != nil {
		return fmt.Errorf("failed to decode into struct: %w", err)
	}
	return nil
}

func (w *envWalker) StructField(field reflect.StructField, value reflect.Value) error {
	if value.Kind() == reflect.Struct || isPtrToStruct(value) {
		w.location = append(w.location, strcase.ToSnake(field.Name))
	}
	return nil
}

func isPtrToStruct(value reflect.Value) bool {
	return value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Struct
}
