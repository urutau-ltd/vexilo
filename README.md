# Vexilo

Vexilo is a Go library for HTML-first applications built with [`aile`](https://codeberg.org/urutau-ltd/aile) and [HTMX](https://htmx.org).

It handles a narrow slice of server-side UI mechanics:

- Shared list state parsing (`q`, `limit`)
- HTMX request intent detection
- HTML response writing with `text/html; charset=utf-8`
- Out-of-band fragment composition for HTMX

## Design rules

- Small API surface
- Explicit behaviour
- Exact HTML contracts where needed
- No application-specific CSS classes
- No presentation helpers or generated app UI

## Packages

- `requeststate`: parse `q` and `limit`
- `htmx`: detect list vs editor HTMX requests
- `fragment`: write HTML and compose OOB fragments

## Intended usage

Use `aile` for routing and middleware, `vexilo` for HTMX/HTML mechanics, and keep templates, CSS, domain logic, persistence, auth, and layout in your application.

## HTML ownership

- `fragment.Page.Main` and `fragment.Page.OOB` are caller-provided HTML
- `fragment.OOB(...)` inserts inner HTML literally into the emitted wrapper
- Vexilo does not escape or sanitize markup
- The application is responsible for the HTML it passes to Vexilo

```go
state := requeststate.Parse(r, 10, 100)
if htmx.IsListRequest(r, "providers-body", []string{"provider-search"}, nil) {
	_ = fragment.WriteHTML(w, fragment.Page{
		Main: renderList(state),
		OOB: [][]byte{
			fragment.OOB("tbody", "providers-body", "innerHTML", renderRows(state)),
		},
	})
}
```
