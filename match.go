package argument

import "regexp"

// Match check that given string value matches the provided pattern.
func Match(argument, value string, pattern *regexp.Regexp) error {
	if !pattern.MatchString(value) {
		return patternError{argument: argument, pattern: pattern}
	}
	return nil
}
