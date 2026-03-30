package panelcrud

import (
	"errors"
	"net/http"

	"codeberg.org/urutau-ltd/vexilo/fragment"
	"codeberg.org/urutau-ltd/vexilo/htmx"
)

var errEmptyListBody = errors.New("panelcrud: empty list body target")

// Targets describes the HTMX targets and triggers for a split list/editor view.
type Targets struct {
	// ListBody is the id of the tbody refreshed out-of-band after mutations.
	ListBody string

	// Editor is the HTMX target used for editor requests.
	Editor string

	// ListTriggerIDs matches HX-Trigger values that should be treated as list
	// requests.
	ListTriggerIDs []string

	// ListTriggerNames matches HX-Trigger-Name values that should be treated as
	// list requests.
	ListTriggerNames []string
}

// Update carries caller-owned HTML for a CRUD mutation response.
type Update struct {
	// Editor is written as the main fragment.
	Editor []byte

	// List is inserted as the inner HTML of the tbody identified by Targets.ListBody.
	List []byte

	// OOB appends extra caller-owned out-of-band fragments after the list refresh.
	OOB [][]byte
}

// IsListRequest reports whether r is an HTMX request for the list side of the panel.
func (t Targets) IsListRequest(r *http.Request) bool {
	return htmx.IsListRequest(r, t.ListBody, t.ListTriggerIDs, t.ListTriggerNames)
}

// IsEditorRequest reports whether r is an HTMX request for the editor side of the panel.
func (t Targets) IsEditorRequest(r *http.Request) bool {
	return htmx.IsEditorRequest(r, t.Editor)
}

// EditorWithList builds the common HTMX response shape for a CRUD mutation:
// editor HTML as the main fragment plus a tbody out-of-band refresh for the list.
func (t Targets) EditorWithList(update Update) (fragment.Page, error) {
	if t.ListBody == "" {
		return fragment.Page{}, errEmptyListBody
	}

	oob := make([][]byte, 0, 1+len(update.OOB))
	oob = append(oob, fragment.OOB("tbody", t.ListBody, "innerHTML", update.List))
	oob = append(oob, update.OOB...)

	return fragment.Page{
		Main: update.Editor,
		OOB:  oob,
	}, nil
}

// WriteEditorWithList writes the HTMX response for EditorWithList using fragment.WriteHTML.
func (t Targets) WriteEditorWithList(w http.ResponseWriter, update Update) error {
	page, err := t.EditorWithList(update)
	if err != nil {
		return err
	}

	return fragment.WriteHTML(w, page)
}
