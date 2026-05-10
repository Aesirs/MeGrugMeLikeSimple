package helpers

import "log"

// Log returns v. If err is non-nil, it logs the error first.
// An optional msg replaces "error" in the log line.
func Log[T any](v T, err error, msg ...string) T {
	if err != nil {
		label := "error"
		if len(msg) > 0 {
			label = msg[0]
		}
		log.Printf("%s: %v", label, err)
	}
	return v
}

// LogE logs a plain error if non-nil.
// An optional msg replaces "error" in the log line.
func LogE(err error, msg ...string) {
	if err != nil {
		label := "error"
		if len(msg) > 0 {
			label = msg[0]
		}
		log.Printf("%s: %v", label, err)
	}
}
