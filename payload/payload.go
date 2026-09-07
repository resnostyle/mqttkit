// Package payload provides small helpers for MQTT JSON payloads.
package payload

import (
	"regexp"
	"strings"
	"time"
)

var nonSlug = regexp.MustCompile(`[^a-z0-9_]+`)

// NilIfEmpty returns nil for empty strings, otherwise the string (for JSON omit-as-null).
func NilIfEmpty(s string) any {
	if s == "" {
		return nil
	}
	return s
}

// UTCNowISO returns the current UTC time in RFC3339.
func UTCNowISO() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// Slugify lowercases value, replaces non-alphanumeric runs with underscores, and trims edges.
// Empty results become fallback (or "device" if fallback is empty).
func Slugify(value, fallback string) string {
	slug := nonSlug.ReplaceAllString(strings.ToLower(strings.TrimSpace(value)), "_")
	slug = strings.Trim(slug, "_")
	if slug == "" {
		if fallback == "" {
			return "device"
		}
		return fallback
	}
	return slug
}
