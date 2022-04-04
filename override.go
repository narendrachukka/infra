package appconfig

import (
	"fmt"
	"strings"
)

func loadFromOverrides(target interface{}, opts Options) error {
	raw := restructure(opts.Overrides)

	if err := decode(target, raw, opts); err != nil {
		return fmt.Errorf("failed to decode from overrides: %w", err)
	}
	return nil
}

// TODO: more tests for this with more nesting, and overlapping overrides
func restructure(flat []string) map[string]interface{} {
	raw := map[string]interface{}{}
	for _, item := range flat {
		key, value := parseOverride(item)
		if len(key) == 0 {
			continue
		}
		current := raw
		last := len(key) - 1
		for _, part := range key[:last] {
			if next, exists := current[part]; !exists {
				next := map[string]interface{}{}
				current[part] = next
				current = next
			} else {
				current = next.(map[string]interface{})
			}
		}
		current[key[last]] = value
	}
	return raw
}

func parseOverride(item string) ([]string, string) {
	parts := strings.SplitN(item, "=", 2)
	if len(parts) < 2 {
		return nil, ""
	}
	return strings.Split(parts[0], "."), parts[1]

}
