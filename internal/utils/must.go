package helpers

// Must returns v if err is nil, otherwise panics.
// Use at program startup or when an error truly means "can't proceed."
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// MustE panics if err is non-nil.
// Use with functions that return only an error.
func MustE(err error) {
	if err != nil {
		panic(err)
	}
}

// Ignore discards the error from a (value, error) pair, returning just v.
// Use to explicitly signal "I know this returns an error, I'm ignoring it."
func Ignore[T any](v T, _ error) T {
	return v
}

// IgnoreE discards a plain error.
// Use with functions that return only an error.
func IgnoreE(err error) {
	// noop — explicitly discarded
}
