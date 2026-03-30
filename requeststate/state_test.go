package requeststate

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestParseQueryString(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers?q=abc&limit=25", nil)

	state := Parse(req, 10, 100)

	if state.Query != "abc" {
		t.Fatalf("expected query abc, got %q", state.Query)
	}
	if state.Limit != 25 {
		t.Fatalf("expected limit 25, got %d", state.Limit)
	}
}

func TestParsePostFormBody(t *testing.T) {
	req := httptest.NewRequest("POST", "/providers", strings.NewReader("q=abc&limit=20"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	state := Parse(req, 10, 100)

	if state.Query != "abc" {
		t.Fatalf("expected query abc, got %q", state.Query)
	}
	if state.Limit != 20 {
		t.Fatalf("expected limit 20, got %d", state.Limit)
	}
}

func TestParseTrimsQuery(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers?q=++abc++", nil)

	state := Parse(req, 10, 100)

	if state.Query != "abc" {
		t.Fatalf("expected trimmed query abc, got %q", state.Query)
	}
}

func TestParseFallsBackOnInvalidLimit(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers?limit=wat", nil)

	state := Parse(req, 10, 100)

	if state.Limit != 10 {
		t.Fatalf("expected default limit 10, got %d", state.Limit)
	}
}

func TestParseClampsLimit(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers?q=abc&limit=999", nil)

	state := Parse(req, 10, 100)

	if state.Query != "abc" {
		t.Fatalf("expected query abc, got %q", state.Query)
	}
	if state.Limit != 100 {
		t.Fatalf("expected clamped limit 100, got %d", state.Limit)
	}
}

func TestParseFallsBackOnNonPositiveLimit(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers?limit=0", nil)

	state := Parse(req, 10, 100)

	if state.Limit != 10 {
		t.Fatalf("expected default limit 10, got %d", state.Limit)
	}
}

func TestParseNilRequest(t *testing.T) {
	state := Parse(nil, 10, 100)

	if state.Query != "" {
		t.Fatalf("expected empty query, got %q", state.Query)
	}
	if state.Limit != 10 {
		t.Fatalf("expected default limit 10, got %d", state.Limit)
	}
}

func TestParsePrefersFormBodyOverQuery(t *testing.T) {
	req := httptest.NewRequest("POST", "/providers?q=url&limit=15", strings.NewReader("q=body&limit=30"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Method = http.MethodPost

	state := Parse(req, 10, 100)

	if state.Query != "body" {
		t.Fatalf("expected body query body, got %q", state.Query)
	}
	if state.Limit != 30 {
		t.Fatalf("expected body limit 30, got %d", state.Limit)
	}
}
