package patterns

import (
	"fmt"
	"regexp"
	"strings"
)

func Compile(pattern map[string]string) *regexp.Regexp {
	var builder strings.Builder
	builder.WriteString("^(?:")
	addOr := false

	for name, pattern := range pattern {
		if addOr {
			builder.WriteRune('|')
		} else {
			addOr = true
		}
		fmt.Fprintf(&builder, "(?P<%s>%s)", name, pattern)
	}
	builder.WriteRune(')')

	return regexp.MustCompile(builder.String())
}

type Raw map[string]struct {
	pattern string
	isRegex bool
}

func Construct(pattern Raw) *regexp.Regexp {
	var n = map[string]string{}

	for k, v := range pattern {
		if v.isRegex {
			n[k] = v.pattern
		} else {
			// Escape regex meta characters, as we want to match this string literally
			n[k] = regexp.QuoteMeta(v.pattern)
		}
	}

	return Compile(n)
}
