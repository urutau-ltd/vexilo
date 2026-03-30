// Package requeststate parses the shared request parameters used by list views.
//
// Vexilo keeps this package intentionally small: it only reads the common `q`
// and `limit` fields that HTML-first list pages tend to share across GET query
// strings and `application/x-www-form-urlencoded` submissions.
//
// Parse normalizes `q` with strings.TrimSpace, falls back to the caller's
// default limit for invalid or non-positive values, and clamps to maxLimit when
// maxLimit is positive.
package requeststate
