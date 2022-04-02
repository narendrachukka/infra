package appconfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

type Options struct {
	FieldTagName string
	Filename     string
	EnvPrefix    string
	Flags        func(fn func(*pflag.Flag))
	Overrides    []string
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
	if opts.Flags != nil {
		if err := loadFromFlags(target, opts); err != nil {
			return err
		}
	}
	if len(opts.Overrides) > 0 {
		if err := loadFromOverrides(target, opts); err != nil {
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

	if err := decode(target, raw, opts); err != nil {
		return fmt.Errorf("failed to decode from %s: %w", opts.Filename, err)
	}
	return nil
}

func decode(target interface{}, raw map[string]interface{}, opts Options) error {
	cfg := mapstructure.DecoderConfig{
		Squash:  true,
		Result:  target,
		TagName: opts.FieldTagName,
		MatchName: func(key string, fieldName string) bool {
			return strings.EqualFold(key, strcase.ToSnake(fieldName))
		},
	}
	decoder, err := mapstructure.NewDecoder(&cfg)
	if err != nil {
		return err
	}
	return decoder.Decode(raw)
}
