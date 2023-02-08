package argument

import (
	"fmt"
	"regexp"
	"time"
)

// Error is the custom error type for argument package.
type Error string

// Error implements the standard error interface.
func (err Error) Error() string {
	return string(err)
}

const (
	// ErrInvalid is the "base" error returned from argument validations.
	ErrInvalid = Error("invalid argument")
)

type requiredError struct {
	argument string
}

func (err requiredError) Error() string {
	return fmt.Sprintf("the argument %s is required", err.argument)
}

func (err requiredError) Unwrap() error {
	return ErrInvalid
}

type notEqualError struct {
	argument string
	value    any
}

func (err notEqualError) Error() string {
	return fmt.Sprintf("the argument %q must be equal to %v", err.argument, err.value)
}

func (err notEqualError) Unwrap() error {
	return ErrInvalid
}

type equalError struct {
	argument string
	value    any
}

func (err equalError) Error() string {
	return fmt.Sprintf("the argument %q cannot be equal to %v", err.argument, err.value)
}

func (err equalError) Unwrap() error {
	return ErrInvalid
}

type maxLengthError struct {
	argument  string
	maxLength int
}

func (err maxLengthError) Error() string {
	return fmt.Sprintf("the argument %s must be long at most %d characters", err.argument, err.maxLength)
}

func (err maxLengthError) Unwrap() error {
	return ErrInvalid
}

type lengthError struct {
	argument  string
	minLength int
	maxLength int
}

func (err lengthError) Error() string {
	return fmt.Sprintf("the argument %s must be long between %d and %d characters", err.argument, err.minLength, err.maxLength)
}

func (err lengthError) Unwrap() error {
	return ErrInvalid
}

type patternError struct {
	argument string
	pattern  *regexp.Regexp
}

func (err patternError) Error() string {
	return fmt.Sprintf("the argument %s must match the pattern %q", err.argument, err.pattern.String())
}

func (err patternError) Unwrap() error {
	return ErrInvalid
}

type minError struct {
	argument string
	min      any
}

func (err minError) Error() string {
	return fmt.Sprintf("the argument %s must be at least %v", err.argument, err.min)
}

func (err minError) Unwrap() error {
	return ErrInvalid
}

type rangeError struct {
	argument string
	min      any
	max      any
}

func (err rangeError) Error() string {
	return fmt.Sprintf("the argument %s must be between %v and %v", err.argument, err.min, err.max)
}

func (err rangeError) Unwrap() error {
	return ErrInvalid
}

type beforeError struct {
	argument  string
	reference time.Time
}

func (err beforeError) Error() string {
	return fmt.Sprintf("the argument %s must occurs before %s", err.argument, err.reference.Format(time.RFC3339))
}

func (err beforeError) Unwrap() error {
	return ErrInvalid
}

type afterError struct {
	argument  string
	reference time.Time
}

func (err afterError) Error() string {
	return fmt.Sprintf("the argument %s must occurs after %s", err.argument, err.reference.Format(time.RFC3339))
}

func (err afterError) Unwrap() error {
	return ErrInvalid
}
