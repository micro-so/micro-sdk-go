package requestconfig_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/micro-so/micro-sdk-go/internal/requestconfig"
	"github.com/micro-so/micro-sdk-go/option"
)

type retryTransport struct {
	attempts int
	failure  string
}

func (h *retryTransport) Do(req *http.Request) (*http.Response, error) {
	h.attempts++
	if h.attempts == 1 && h.failure == "transport" {
		return nil, errors.New("connection lost after send")
	}
	status := 200
	retry := "false"
	if h.attempts == 1 {
		status = 500
		retry = "true"
	}
	return &http.Response{StatusCode: status, Header: http.Header{"Content-Type": []string{"application/json"}, "Retry-After": []string{"0"}, "X-Should-Retry": []string{retry}}, Body: io.NopCloser(strings.NewReader("{}")), Request: req}, nil
}

func TestRetriesDoNotRepeatUncertainWrites(t *testing.T) {
	cases := []struct {
		name, method, path string
		explicit           bool
		attempts           int
	}{
		{"create", "POST", "/v2/prism/team/document", false, 1},
		{"patch", "PATCH", "/v2/prism/team/document/id", false, 1},
		{"delete", "DELETE", "/v2/prism/team/document/id", false, 1},
		{"webhook", "POST", "/v2/webhooks/team", false, 1},
		{"query-like-write", "POST", "/v2/prism/team/document/query/extra", false, 1},
		{"get", "GET", "/v2/prism/team/document", false, 2},
		{"query", "POST", "/v2/prism/team/document/query", false, 2},
		{"query-alias", "POST", "/v2/prism/query/team/document", false, 2},
		{"explicit-write-retry", "POST", "/v2/prism/team/document", true, 2},
	}
	for _, tc := range cases {
		for _, failure := range []string{"status", "transport"} {
			t.Run(tc.name+"/"+failure, func(t *testing.T) {
				transport := &retryTransport{failure: failure}
				opts := []option.RequestOption{option.WithBaseURL("https://micro.test"), option.WithHTTPClient(transport), option.WithHeader("Idempotency-Key", "operation-id")}
				if tc.explicit {
					opts = append(opts, option.WithMaxRetries(1))
				}
				var result map[string]interface{}
				err := requestconfig.ExecuteNewRequest(context.Background(), tc.method, tc.path, map[string]string{"value": "test"}, &result, opts...)
				if transport.attempts != tc.attempts {
					t.Fatalf("attempts=%d want %d", transport.attempts, tc.attempts)
				}
				if (err != nil) != (tc.attempts == 1) {
					t.Fatalf("unexpected error: %v", err)
				}
			})
		}
	}
}

func TestExplicitZeroDisablesReadRetries(t *testing.T) {
	transport := &retryTransport{failure: "status"}
	var result map[string]interface{}
	err := requestconfig.ExecuteNewRequest(context.Background(), "GET", "/v2/prism/team/document", nil, &result, option.WithBaseURL("https://micro.test"), option.WithHTTPClient(transport), option.WithMaxRetries(0))
	if err == nil || transport.attempts != 1 {
		t.Fatalf("err=%v attempts=%d", err, transport.attempts)
	}
}
