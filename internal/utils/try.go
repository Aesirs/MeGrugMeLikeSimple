package helpers

// Try calls fn with err if err is non-nil, otherwise returns v.
func Try[T any](v T, err error, fn func(error)) T {
	if err != nil {
		fn(err)
	}
	return v
}

// TryE calls fn with err if err is non-nil.
func TryE(err error, fn func(error)) {
	if err != nil {
		fn(err)
	}
}
