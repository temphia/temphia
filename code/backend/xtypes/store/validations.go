package store

import (
	"regexp"
	"strings"
)

var (
	SlugExp        = regexp.MustCompile("^[a-z0-9-]*$")
	ReplaceSlugExp = regexp.MustCompile("[^a-z0-9]+")
	ReservedSlugs  = []string{"page", "limit", "cursor"} // fixme => copy all reserved words from from autoapi and other
)

func Slugify(s string) string {
	return strings.Trim(ReplaceSlugExp.ReplaceAllString(strings.ToLower(s), "-"), "-")
}
