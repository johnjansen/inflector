package inflector

import (
	"regexp"
	"strings"
)

const (
	regexFlags = "(?i)"
	regexEnd   = "$"
)

var (
	uncountables   words
	plurals        rules
	singulars      rules
	pluralsCache   = cache{}
	singularsCache = cache{}
)

type cache map[string]string

func (c cache) get(key string, r rules) string {
	value, ok := c[key]
	if !ok {
		value = r.convert(key)
		c[key] = value
	}
	return value
}

type rules []rule

func (r rules) convert(text string) string {
	if len(text) > 0 && !uncountables.contains(text) {
		for _, rule := range r {
			if newtext, ok := rule.apply(text); ok {
				return newtext
			}
		}
	}
	return text
}

type words []string

func (w words) contains(value string) bool {
	for _, word := range w {
		if strings.HasSuffix(value, word) {
			return true
		}
	}
	return false
}

type rule struct {
	regex       *regexp.Regexp
	replacement string
	same        bool
}

func (r rule) apply(text string) (string, bool) {
	if r.regex.MatchString(text) {
		return r.regex.ReplaceAllString(text, r.replacement), true
	}
	return text, false
}

func plural(regex, replacement string) {
	r := rule{regex: regexp.MustCompile(regexFlags + regex + regexEnd), replacement: replacement}
	plurals = append(rules{r}, plurals...)
}

func singular(regex, replacement string) {
	r := rule{regex: regexp.MustCompile(regexFlags + regex + regexEnd), replacement: replacement}
	singulars = append(rules{r}, singulars...)
}

func irregular(sing, plu string) {
	plural(regexp.QuoteMeta(sing), plu)
	plural(regexp.QuoteMeta(plu), plu)
	singular(regexp.QuoteMeta(sing), sing)
	singular(regexp.QuoteMeta(plu), sing)
}

func uncountable(words ...string) {
	uncountables = append(uncountables, words...)
}

// Pluralize ...
func Pluralize(singular string) string {
	return pluralsCache.get(singular, plurals)
}

// Singularize ...
func Singularize(plural string) string {
	return singularsCache.get(plural, singulars)
}
