package requeststate

import (
	"net/http"
	"strconv"
	"strings"
)

// ListState carries the shared request parameters for a list view.
//
// Query comes from the `q` form field after strings.TrimSpace. Limit comes from
// the `limit` form field after applying the caller's default and max limits.
type ListState struct {
	Query string
	Limit int
}

// Parse reads `q` and `limit` from r.Form after calling r.ParseForm.
//
// Parse supports both query-string parameters and
// `application/x-www-form-urlencoded` bodies. If r is nil, Parse returns a
// zero query and defaultLimit. Invalid or non-positive limits fall back to
// defaultLimit. If maxLimit is positive, larger limits are clamped to it.
func Parse(r *http.Request, defaultLimit, maxLimit int) ListState {
	if r == nil {
		return ListState{Limit: defaultLimit}
	}

	_ = r.ParseForm()

	state := ListState{
		Query: strings.TrimSpace(r.Form.Get("q")),
		Limit: defaultLimit,
	}

	rawLimit := r.Form.Get("limit")
	if rawLimit == "" {
		return state
	}

	parsedLimit, err := strconv.Atoi(rawLimit)
	if err != nil || parsedLimit < 1 {
		return state
	}

	if maxLimit > 0 && parsedLimit > maxLimit {
		state.Limit = maxLimit
		return state
	}

	state.Limit = parsedLimit
	return state
}
