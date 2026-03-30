package htmx

import (
	"fmt"
	"net/http/httptest"
)

func ExampleIsListRequest() {
	req := httptest.NewRequest("GET", "/providers", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Target", "#providers-body")

	ok := IsListRequest(req, "providers-body", []string{"provider-search"}, []string{"q"})
	fmt.Println(ok)

	// Output:
	// true
}

func ExampleIsEditorRequest() {
	req := httptest.NewRequest("GET", "/providers/new", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Target", "provider-editor")

	ok := IsEditorRequest(req, "#provider-editor")
	fmt.Println(ok)

	// Output:
	// true
}
