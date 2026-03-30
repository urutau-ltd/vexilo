// Package panelcrud provides a small helper layer for HTML-first CRUD screens
// that use a split list/editor view.
//
// The package builds on Vexilo's lower-level packages instead of replacing
// them. Applications still own full-page rendering, templates, notices, and
// all markup. panelcrud only helps with two repetitive tasks:
//
//   - detect whether an HTMX request is aimed at the list or the editor
//   - compose the common response shape of editor main fragment plus list tbody
//     out-of-band refresh
//
// The package is intentionally narrow. It does not create handlers, forms,
// routes, repositories, or presentation components.
package panelcrud
