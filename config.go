package appconfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
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
	// TODO: filename is optional
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

// TODO: support nested structs in target
func loadFromEnv(target interface{}, opts Options) error {
	opts.EnvPrefix = strings.ToUpper(opts.EnvPrefix)
	// TODO: probably copy this instead of import
	raw := env.ToMap(os.Environ())

	cfg := mapstructure.DecoderConfig{
		// TODO: Squash:           false,
		Result:           target,
		TagName:          opts.FieldTagName,
		WeaklyTypedInput: true,
		MatchName: func(key string, fieldName string) bool {
			name := opts.EnvPrefix + "_" + strings.ToUpper(strcase.ToSnake(fieldName))
			return key == name
		},
	}
	decoder, err := mapstructure.NewDecoder(&cfg)
	if err != nil {
		return fmt.Errorf("failed to create decoder: %w", err)
	}
	if err := decoder.Decode(raw); err != nil {
		return fmt.Errorf("failed to decode from environment variables: %w", err)
	}
	return nil
}
