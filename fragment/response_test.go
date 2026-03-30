package fragment

import (
	"net/http/httptest"
	"testing"
)

func TestWriteHTMLSetsContentTypeAndWritesAllFragments(t *testing.T) {
	rec := httptest.NewRecorder()

	err := WriteHTML(rec, Page{
		Main: []byte("<section>Main</section>"),
		OOB: [][]byte{
			[]byte(`<div id="x" hx-swap-oob="true">X</div>`),
			[]byte(`<section id="y" hx-swap-oob="beforeend">Y</section>`),
		},
	})
	if err != nil {
		t.Fatalf("WriteHTML returned error: %v", err)
	}

	if got := rec.Header().Get("Content-Type"); got != "text/html; charset=utf-8" {
		t.Fatalf("expected HTML content type, got %q", got)
	}

	want := `<section>Main</section><div id="x" hx-swap-oob="true">X</div><section id="y" hx-swap-oob="beforeend">Y</section>`
	if rec.Body.String() != want {
		t.Fatalf("expected %q, got %q", want, rec.Body.String())
	}
}

func TestOOBWrapsTbodyExactly(t *testing.T) {
	got := string(OOB("tbody", "providers-body", "innerHTML", []byte(`<tr><td>A</td></tr>`)))
	want := `<table hidden aria-hidden="true"><tbody id="providers-body" hx-swap-oob="innerHTML"><tr><td>A</td></tr></tbody></table>`

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestOOBWrapsGenericTagExactly(t *testing.T) {
	got := string(OOB("section", "provider-editor", "outerHTML", []byte(`<p>Hello</p>`)))
	want := `<section id="provider-editor" hx-swap-oob="outerHTML"><p>Hello</p></section>`

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}
