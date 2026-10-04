package htmx

import (
	"net/http"
	"slices"
	"strings"
)

// IsBoosted reports whether HTMX marked the request as boosted navigation.
//
// It returns true only when the request contains `HX-Boosted: true`.
func IsBoosted(r *http.Request) bool {
	return headerEquals(r, "HX-Boosted", "true")
}

// IsListRequest reports whether r is an HTMX request for a list fragment.
//
// It returns true only when:
//   - `HX-Request: true` is present
//   - the request is not boosted navigation
//   - `HX-Target` matches target, or `HX-Trigger` matches one of triggerIDs, or
//     `HX-Trigger-Name` matches one of triggerNames
//
// Target matching accepts both `providers-body` and `#providers-body` forms.
func IsListRequest(r *http.Request, target string, triggerIDs []string, triggerNames []string) bool {
	if !headerEquals(r, "HX-Request", "true") || IsBoosted(r) {
		return false
	}

	if sameTarget(r.Header.Get("HX-Target"), target) {
		return true
	}

	trigger := r.Header.Get("HX-Trigger")
	if contains(triggerIDs, trigger) {
		return true
	}

	triggerName := r.Header.Get("HX-Trigger-Name")
	if contains(triggerNames, triggerName) {
		return true
	}

	return false
}

// IsEditorRequest reports whether r is an HTMX request for the editor fragment.
//
// It returns true only when:
//   - `HX-Request: true` is present
//   - the request is not boosted navigation
//   - `HX-Target` matches target
//
// Target matching accepts both `provider-editor` and `#provider-editor` forms.
func IsEditorRequest(r *http.Request, target string) bool {
	if target == "" || !headerEquals(r, "HX-Request", "true") || IsBoosted(r) {
		return false
	}

	return sameTarget(r.Header.Get("HX-Target"), target)
}

func headerEquals(r *http.Request, key, want string) bool {
	if r == nil {
		return false
	}
	return r.Header.Get(key) == want
}

func contains(values []string, want string) bool {
	if want == "" {
		return false
	}

	return slices.Contains(values, want)
}

func sameTarget(got, want string) bool {
	if got == "" || want == "" {
		return false
	}

	return normalizeTarget(got) == normalizeTarget(want)
}

func normalizeTarget(target string) string {
	return strings.TrimPrefix(target, "#")
}
