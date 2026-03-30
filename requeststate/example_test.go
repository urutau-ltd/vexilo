package requeststate

import (
	"fmt"
	"net/http/httptest"
	"strings"
)

func ExampleParse() {
	req := httptest.NewRequest("POST", "/providers", strings.NewReader("q=++Acme++&limit=250"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	state := Parse(req, 25, 100)

	fmt.Printf("query=%q limit=%d\n", state.Query, state.Limit)

	// Output:
	// query="Acme" limit=100
}
