package panelcrud

import (
	"net/http/httptest"
	"testing"
)

func TestTargetsIsListRequestByTarget(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Target", "#providers-body")

	targets := Targets{
		ListBody: "providers-body",
	}

	if !targets.IsListRequest(req) {
		t.Fatalf("expected list request")
	}
}

func TestTargetsIsListRequestByTriggerID(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Trigger", "provider-search")

	targets := Targets{
		ListBody:       "providers-body",
		ListTriggerIDs: []string{"provider-search"},
	}

	if !targets.IsListRequest(req) {
		t.Fatalf("expected list request from trigger ID")
	}
}

func TestTargetsIsListRequestByTriggerName(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Trigger-Name", "q")

	targets := Targets{
		ListBody:         "providers-body",
		ListTriggerNames: []string{"q"},
	}

	if !targets.IsListRequest(req) {
		t.Fatalf("expected list request from trigger name")
	}
}

func TestTargetsIsEditorRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers/new", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Target", "#provider-editor")

	targets := Targets{
		Editor: "provider-editor",
	}

	if !targets.IsEditorRequest(req) {
		t.Fatalf("expected editor request")
	}
}

func TestEditorWithListBuildsPage(t *testing.T) {
	targets := Targets{
		ListBody: "providers-body",
		Editor:   "provider-editor",
	}

	page, err := targets.EditorWithList(Update{
		Editor: []byte(`<section id="provider-editor">Editor</section>`),
		List:   []byte(`<tr><td>A</td></tr>`),
		OOB: [][]byte{
			[]byte(`<output id="notice" hx-swap-oob="true">Saved</output>`),
		},
	})
	if err != nil {
		t.Fatalf("EditorWithList returned error: %v", err)
	}

	if got := string(page.Main); got != `<section id="provider-editor">Editor</section>` {
		t.Fatalf("expected editor main fragment, got %q", got)
	}

	if len(page.OOB) != 2 {
		t.Fatalf("expected 2 OOB fragments, got %d", len(page.OOB))
	}

	wantList := `<table hidden aria-hidden="true"><tbody id="providers-body" hx-swap-oob="innerHTML"><tr><td>A</td></tr></tbody></table>`
	if got := string(page.OOB[0]); got != wantList {
		t.Fatalf("expected %q, got %q", wantList, got)
	}

	wantNotice := `<output id="notice" hx-swap-oob="true">Saved</output>`
	if got := string(page.OOB[1]); got != wantNotice {
		t.Fatalf("expected %q, got %q", wantNotice, got)
	}
}

func TestEditorWithListRejectsEmptyListBody(t *testing.T) {
	targets := Targets{}

	_, err := targets.EditorWithList(Update{
		Editor: []byte(`<section>Editor</section>`),
		List:   []byte(`<tr><td>A</td></tr>`),
	})
	if err == nil {
		t.Fatalf("expected error for empty list body target")
	}
}

func TestWriteEditorWithListWritesHTML(t *testing.T) {
	rec := httptest.NewRecorder()

	targets := Targets{
		ListBody: "providers-body",
		Editor:   "provider-editor",
	}

	err := targets.WriteEditorWithList(rec, Update{
		Editor: []byte(`<section id="provider-editor">Editor</section>`),
		List:   []byte(`<tr><td>A</td></tr>`),
		OOB: [][]byte{
			[]byte(`<output id="notice" hx-swap-oob="true">Saved</output>`),
		},
	})
	if err != nil {
		t.Fatalf("WriteEditorWithList returned error: %v", err)
	}

	if got := rec.Header().Get("Content-Type"); got != "text/html; charset=utf-8" {
		t.Fatalf("expected HTML content type, got %q", got)
	}

	want := `<section id="provider-editor">Editor</section><table hidden aria-hidden="true"><tbody id="providers-body" hx-swap-oob="innerHTML"><tr><td>A</td></tr></tbody></table><output id="notice" hx-swap-oob="true">Saved</output>`
	if got := rec.Body.String(); got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}
