package htmx

import (
	"net/http/httptest"
	"testing"
)

func TestIsBoosted(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers", nil)
	req.Header.Set("HX-Boosted", "true")

	if !IsBoosted(req) {
		t.Fatalf("expected boosted request")
	}
}

func TestIsListRequestByTarget(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Target", "providers-body")

	if !IsListRequest(req, "providers-body", nil, nil) {
		t.Fatalf("expected list request")
	}
}

func TestIsListRequestByTargetWithHashHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Target", "#providers-body")

	if !IsListRequest(req, "providers-body", nil, nil) {
		t.Fatalf("expected list request with hashed HX-Target")
	}
}

func TestIsListRequestByTargetWithHashCaller(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Target", "providers-body")

	if !IsListRequest(req, "#providers-body", nil, nil) {
		t.Fatalf("expected list request with hashed caller target")
	}
}

func TestIsListRequestByTriggerID(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Trigger", "provider-search")

	if !IsListRequest(req, "providers-body", []string{"provider-search"}, nil) {
		t.Fatalf("expected list request from trigger ID")
	}
}

func TestIsListRequestByTriggerName(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Trigger-Name", "provider-search")

	if !IsListRequest(req, "providers-body", nil, []string{"provider-search"}) {
		t.Fatalf("expected list request from trigger name")
	}
}

func TestIsListRequestRejectsBoosted(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Boosted", "true")
	req.Header.Set("HX-Target", "providers-body")

	if IsListRequest(req, "providers-body", nil, nil) {
		t.Fatalf("expected boosted request to be ignored")
	}
}

func TestIsEditorRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers/new", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Target", "provider-editor")

	if !IsEditorRequest(req, "provider-editor") {
		t.Fatalf("expected editor request")
	}
}

func TestIsEditorRequestWithHashHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers/new", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Target", "#provider-editor")

	if !IsEditorRequest(req, "provider-editor") {
		t.Fatalf("expected editor request with hashed HX-Target")
	}
}

func TestIsEditorRequestWithHashCaller(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers/new", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Target", "provider-editor")

	if !IsEditorRequest(req, "#provider-editor") {
		t.Fatalf("expected editor request with hashed caller target")
	}
}

func TestIsEditorRequestRejectsBoosted(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers/new", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Boosted", "true")
	req.Header.Set("HX-Target", "provider-editor")

	if IsEditorRequest(req, "provider-editor") {
		t.Fatalf("expected boosted request to be ignored")
	}
}

func TestIsListRequestRejectsNonHTMX(t *testing.T) {
	req := httptest.NewRequest("GET", "/providers", nil)
	req.Header.Set("HX-Target", "providers-body")

	if IsListRequest(req, "providers-body", nil, nil) {
		t.Fatalf("expected plain request to be ignored")
	}
}
