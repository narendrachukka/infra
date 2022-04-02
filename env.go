package appconfig

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"github.com/mitchellh/reflectwalk"
)

func loadFromEnv(target interface{}, opts Options) error {
	opts.EnvPrefix = strings.ToUpper(opts.EnvPrefix)
	walker := &envWalker{
		vars:     toMap(opts.EnvPrefix, os.Environ()),
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
		MatchName:        w.matchName,
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
		if field.Anonymous { // embedded struct
			w.location = append(w.location, "")
			return nil
		}
		w.location = append(w.location, strings.ToUpper(strcase.ToSnake(field.Name)))
	}
	return nil
}

func (w *envWalker) matchName(key string, fieldName string) bool {
	var sb strings.Builder
	for _, part := range w.location {
		if part == "" {
			continue
		}
		sb.WriteString(part)
		sb.WriteString("_")
	}
	sb.WriteString(strings.ToUpper(strcase.ToSnake(fieldName)))
	return key == sb.String()
}

func isPtrToStruct(value reflect.Value) bool {
	return value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Struct
}

// toMap converts an environment variable slice to a map of key/value pairs.
// The environment slice is filtered to only include keys that match the prefix
// because mapstructure iterates over this map, and any keys that do not match
// the prefix will never be used.
func toMap(prefix string, env []string) map[string]string {
	result := map[string]string{}
	for _, raw := range env {
		key, value := getParts(raw)
		if strings.HasPrefix(key, prefix) {
			result[key] = value
		}
	}
	return result
}

func getParts(raw string) (string, string) {
	if raw == "" {
		return "", ""
	}
	// Environment variables on windows can begin with =
	// http://blogs.msdn.com/b/oldnewthing/archive/2010/05/06/10008132.aspx
	parts := strings.SplitN(raw[1:], "=", 2)
	key := raw[:1] + parts[0]
	if len(parts) == 1 {
		return key, ""
	}
	return key, parts[1]
}
