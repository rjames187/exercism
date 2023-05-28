package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	apps := map[string]string{
		"❗": "recommendation",
		"🔍": "search",
		"☀": "weather",
	}
	for _, run := range log {
		char := string(run)
		app, exists := apps[char]
		if exists {
			return app
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	res := ""
	for _, run := range log {
		if run == oldRune {
			res += string(newRune)
			continue
		}
		res += string(run)
	}
	return res
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	count := utf8.RuneCountInString(log)
	return count <= limit
}
