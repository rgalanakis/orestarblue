package main

import "fmt"

func Ptr[T any](value T) *T {
	return &value
}

func ErrWrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func Substr(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length]
}
