package panelcrud

import (
	"fmt"
	"net/http/httptest"
)

func ExampleTargets_WriteEditorWithList() {
	targets := Targets{
		ListBody: "providers-body",
		Editor:   "provider-editor",
	}

	rec := httptest.NewRecorder()

	_ = targets.WriteEditorWithList(rec, Update{
		Editor: []byte(`<section id="provider-editor">Editor</section>`),
		List:   []byte(`<tr><td>A</td></tr>`),
		OOB: [][]byte{
			[]byte(`<output id="notice" hx-swap-oob="true">Saved</output>`),
		},
	})

	fmt.Println(rec.Body.String())

	// Output:
	// <section id="provider-editor">Editor</section><table hidden aria-hidden="true"><tbody id="providers-body" hx-swap-oob="innerHTML"><tr><td>A</td></tr></tbody></table><output id="notice" hx-swap-oob="true">Saved</output>
}
