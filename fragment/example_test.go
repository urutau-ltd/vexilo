package fragment

import (
	"fmt"
	"net/http/httptest"
)

func ExampleWriteHTML() {
	rec := httptest.NewRecorder()

	_ = WriteHTML(rec, Page{
		Main: []byte("<section>List</section>"),
		OOB: [][]byte{
			OOB("tbody", "providers-body", "innerHTML", []byte(`<tr><td>A</td></tr>`)),
		},
	})

	fmt.Println(rec.Body.String())

	// Output:
	// <section>List</section><table hidden aria-hidden="true"><tbody id="providers-body" hx-swap-oob="innerHTML"><tr><td>A</td></tr></tbody></table>
}

func ExampleOOB() {
	fmt.Println(string(OOB("section", "provider-editor", "outerHTML", []byte(`<p>Hello</p>`))))

	// Output:
	// <section id="provider-editor" hx-swap-oob="outerHTML"><p>Hello</p></section>
}
