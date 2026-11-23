package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteIDRequiresJSONResponse(t *testing.T) {
	cases := map[string]struct {
		contentType, body string
		wantErr           string
	}{
		"json":            {"application/json", `{}`, ""},
		"sign-in page":    {"text/html", "<html>sign in</html>", `expected a JSON response, got content-type "text/html"`},
		"no content type": {"", "", "expected a JSON response"},
		"invalid json":    {"application/json", "not json", "error parsing response json"},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, http.MethodDelete, r.Method)
				require.Equal(t, "/v2/tags/t1", r.URL.Path)
				w.Header()["Content-Type"] = []string{tc.contentType}
				_, _ = w.Write([]byte(tc.body))
			}))
			defer srv.Close()
			c, err := New(Config{BaseURL: srv.URL, BearerToken: "tok"})
			require.NoError(t, err)

			err = c.DeleteID(context.Background(), "v2/tags", "t1")
			if tc.wantErr == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, tc.wantErr)
		})
	}
}
