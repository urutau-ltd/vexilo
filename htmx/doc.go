// Package htmx detects a small set of HTMX request intents.
//
// The package is intentionally mechanical. It does not route requests or build
// UI; it only reads HTMX headers to answer narrow questions such as:
//
//   - whether a request is boosted navigation
//   - whether a request targets a list fragment
//   - whether a request targets an editor fragment
//
// Target matching accepts both `id` and `#id` forms so applications can pass
// whichever representation they already use.
package htmx
