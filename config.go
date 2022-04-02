package appconfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

type Options struct {
	Filename string
}

func Load(target interface{}, opts Options) error {
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
		// TODO: Metadata:         nil,
		Result:    target,
		TagName:   "config",
		MatchName: matchConfigFileName,
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

func matchConfigFileName(key string, fieldName string) bool {
	return strings.EqualFold(key, strcase.ToSnake(fieldName))
}
