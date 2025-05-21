package dog

import (
	"strings"
)

// DOG_YEARS is a constant representing 7 years.
const DOG_YEARS = 7

// pushing this
func WhenGrownUp(s string) string {
	return "When the puppy grows up, it says " + strings.ToUpper(s)
}

// Years takes human years as an input and converts to dog years.
func Years(human int) int {
	return human * DOG_YEARS
}
