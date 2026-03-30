// Package fragment writes HTML responses and composes HTMX out-of-band
// fragments.
//
// The package is deliberately narrow and presentation-agnostic. Callers provide
// the HTML bytes for Page.Main and Page.OOB, and OOB inserts the provided inner
// HTML literally into the wrapper it emits. Vexilo does not escape or sanitize
// markup; the application owns the HTML it passes in.
//
// For `tbody` fragments, OOB emits a valid wrapper:
//
//	<table hidden aria-hidden="true"><tbody ...>...</tbody></table>
//
// Other tags are emitted directly with `id`, `hx-swap-oob`, and the caller's
// inner HTML.
package fragment
