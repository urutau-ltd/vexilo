package fragment

import (
	"errors"
	"net/http"
)

const htmlContentType = "text/html; charset=utf-8"

// Page carries the HTML bytes to write in a single response.
//
// Main is written first. OOB fragments are written afterward in slice order.
// Vexilo does not transform, escape, or sanitize these bytes.
type Page struct {
	Main []byte
	OOB  [][]byte
}

// WriteHTML writes page to w with the exact HTML content type.
//
// WriteHTML always sets `Content-Type: text/html; charset=utf-8`, writes
// Page.Main first, and then writes each entry in Page.OOB in order.
func WriteHTML(w http.ResponseWriter, page Page) error {
	if w == nil {
		return errors.New("fragment: nil response writer")
	}

	w.Header().Set("Content-Type", htmlContentType)

	if _, err := w.Write(page.Main); err != nil {
		return err
	}

	for _, fragment := range page.OOB {
		if _, err := w.Write(fragment); err != nil {
			return err
		}
	}

	return nil
}

// OOB builds an HTMX out-of-band fragment.
//
// The returned bytes contain the caller's inner HTML literally. For most tags,
// OOB emits:
//
//	<tag id="..." hx-swap-oob="...">inner</tag>
//
// When tag is `tbody`, OOB wraps the fragment in a hidden table so the output
// remains valid HTML for HTMX to process.
func OOB(tag, id, swap string, inner []byte) []byte {
	if tag == "tbody" {
		return tbodyOOB(id, swap, inner)
	}
	return genericOOB(tag, id, swap, inner)
}

func tbodyOOB(id, swap string, inner []byte) []byte {
	out := make([]byte, 0, len(inner)+len(id)+len(swap)+77)
	out = append(out, `<table hidden aria-hidden="true"><tbody id="`...)
	out = append(out, id...)
	out = append(out, `" hx-swap-oob="`...)
	out = append(out, swap...)
	out = append(out, `">`...)
	out = append(out, inner...)
	out = append(out, `</tbody></table>`...)
	return out
}

func genericOOB(tag, id, swap string, inner []byte) []byte {
	out := make([]byte, 0, len(inner)+len(tag)*2+len(id)+len(swap)+31)
	out = append(out, '<')
	out = append(out, tag...)
	out = append(out, ` id="`...)
	out = append(out, id...)
	out = append(out, `" hx-swap-oob="`...)
	out = append(out, swap...)
	out = append(out, `">`...)
	out = append(out, inner...)
	out = append(out, `</`...)
	out = append(out, tag...)
	out = append(out, '>')
	return out
}
