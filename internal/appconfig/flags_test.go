package appconfig

import (
	"testing"

	"github.com/spf13/pflag"
	"gotest.tools/v3/assert"
)

func TestDefaultsFromEnv(t *testing.T) {
	type target struct {
		OneMore string
		Next    int32
		Many    []string
	}

	setup := func(tg *target) *pflag.FlagSet {
		flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
		flags.String("word", "default-value", "")
		flags.Int("count", 12, "")
		flags.String("two-parts", "", "")
		flags.StringVar(&tg.OneMore, "one-more-extra-long", "", "")
		flags.Int32Var(&tg.Next, "next", 0, "")
		flags.StringSliceVar(&tg.Many, "many", nil, "")
		flags.StringSlice("many-others", nil, "")
		return flags
	}

	t.Run("values from flags", func(t *testing.T) {
		t.Setenv("MYAPP_WORD", "from-env")
		t.Setenv("MYAPP_COUNT", "3")
		t.Setenv("MYAPP_TWO_PARTS", "from-env-2")
		t.Setenv("MYAPP_ONE_MORE_EXTRA_LONG", "from-env-3")
		t.Setenv("MYAPP_NEXT", "4")
		t.Setenv("MYAPP_MANY", "a,b,c")
		t.Setenv("MYAPP_MANY_OTHERS", "d,e,f")

		tg := target{}
		flags := setup(&tg)
		args := []string{
			"--word", "the-value",
			"--count", "22",
			"--two-parts", "one-two",
			"--one-more-extra-long", "value-one",
			"--next", "23",
			"--many", "one,two,three",
			"--many-others", "four",
			"--many-others", "five",
		}
		err := flags.Parse(args)
		assert.NilError(t, err)

		err = DefaultsFromEnv("MYAPP", flags)
		assert.NilError(t, err)

		expected := target{
			OneMore: "value-one",
			Next:    23,
			Many:    []string{"one", "two", "three"},
		}
		assert.DeepEqual(t, tg, expected)

		v, err := flags.GetString("word")
		assert.NilError(t, err)
		assert.Equal(t, v, "the-value")

		i, err := flags.GetInt("count")
		assert.NilError(t, err)
		assert.Equal(t, i, 22)

		v, err = flags.GetString("two-parts")
		assert.NilError(t, err)
		assert.Equal(t, v, "one-two")

		s, err := flags.GetStringSlice("many-others")
		assert.NilError(t, err)
		assert.DeepEqual(t, s, []string{"four", "five"})
	})

	t.Run("defaults from flags", func(t *testing.T) {
		tg := target{}
		flags := setup(&tg)
		err := flags.Parse(nil)
		assert.NilError(t, err)

		err = DefaultsFromEnv("MYAPP", flags)
		assert.NilError(t, err)

		expected := target{}
		assert.DeepEqual(t, tg, expected)

		v, err := flags.GetString("word")
		assert.NilError(t, err)
		assert.Equal(t, v, "default-value")

		i, err := flags.GetInt("count")
		assert.NilError(t, err)
		assert.Equal(t, i, 12)

		v, err = flags.GetString("two-parts")
		assert.NilError(t, err)
		assert.Equal(t, v, "")

		s, err := flags.GetStringSlice("many-others")
		assert.NilError(t, err)
		assert.DeepEqual(t, s, []string{})
	})

	t.Run("values from env", func(t *testing.T) {
		t.Setenv("MYAPP_WORD", "from-env")
		t.Setenv("MYAPP_COUNT", "3")
		t.Setenv("MYAPP_TWO_PARTS", "from-env-2")
		t.Setenv("MYAPP_ONE_MORE_EXTRA_LONG", "from-env-3")
		t.Setenv("MYAPP_NEXT", "4")
		t.Setenv("MYAPP_MANY", "a,b,c")
		t.Setenv("MYAPP_MANY_OTHERS", "d,e,f")

		tg := target{}
		flags := setup(&tg)
		err := flags.Parse(nil)
		assert.NilError(t, err)

		err = DefaultsFromEnv("MYAPP", flags)
		assert.NilError(t, err)

		expected := target{
			OneMore: "from-env-3",
			Next:    4,
			Many:    []string{"a", "b", "c"},
		}
		assert.DeepEqual(t, tg, expected)

		v, err := flags.GetString("word")
		assert.NilError(t, err)
		assert.Equal(t, v, "from-env")

		i, err := flags.GetInt("count")
		assert.NilError(t, err)
		assert.Equal(t, i, 3)

		v, err = flags.GetString("two-parts")
		assert.NilError(t, err)
		assert.Equal(t, v, "from-env-2")

		s, err := flags.GetStringSlice("many-others")
		assert.NilError(t, err)
		assert.DeepEqual(t, s, []string{"d", "e", "f"})
	})

	t.Run("errors setting value from env var", func(t *testing.T) {
		t.Setenv("MYAPP_WORD", "from-env")
		t.Setenv("MYAPP_COUNT", "not-a-number")
		t.Setenv("MYAPP_NEXT", "true")
		t.Setenv("MYAPP_MANY_OTHERS", "d")

		tg := target{}
		flags := setup(&tg)
		err := flags.Parse(nil)
		assert.NilError(t, err)

		err = DefaultsFromEnv("MYAPP", flags)
		assert.ErrorContains(t, err, `failed to set count from environment variable: strconv.ParseInt: parsing "not-a-number": invalid syntax`)
		assert.ErrorContains(t, err, `failed to set next from environment variable: strconv.ParseInt: parsing "true": invalid syntax`)
	})

}
