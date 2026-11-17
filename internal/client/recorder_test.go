package client

import (
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

type recordedRequest struct {
	Method        string
	Path          string
	RawQuery      string
	Authorization string
	ContentType   string
	Accept        string
	Body          string
}

type reply struct {
	status  int
	headers map[string]string
	body    string
}

var tokenReply = reply{status: http.StatusOK, body: `{"access_token":"tok-1","token_type":"bearer","expires_in":3600}`}

// recorder serves scripted replies in order and records every request.
type recorder struct {
	t        *testing.T
	server   *httptest.Server
	mu       sync.Mutex
	requests []recordedRequest
	replies  []reply
}

func newRecorder(t *testing.T, replies ...reply) *recorder {
	t.Helper()
	r := &recorder{t: t, replies: replies}
	r.server = httptest.NewServer(http.HandlerFunc(r.serve))
	t.Cleanup(r.server.Close)
	return r
}

func (r *recorder) serve(w http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	r.mu.Lock()
	r.requests = append(r.requests, recordedRequest{
		Method:        req.Method,
		Path:          req.URL.Path,
		RawQuery:      req.URL.RawQuery,
		Authorization: req.Header.Get("Authorization"),
		ContentType:   req.Header.Get("Content-Type"),
		Accept:        req.Header.Get("Accept"),
		Body:          string(body),
	})
	var rep reply
	if len(r.replies) > 0 {
		rep, r.replies = r.replies[0], r.replies[1:]
	} else {
		rep = reply{status: http.StatusOK, body: `{}`}
	}
	r.mu.Unlock()
	for k, v := range rep.headers {
		w.Header().Set(k, v)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rep.status)
	_, _ = io.WriteString(w, rep.body)
}

func (r *recorder) recorded() []recordedRequest {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]recordedRequest(nil), r.requests...)
}
