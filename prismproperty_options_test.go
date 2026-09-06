package micro_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/micro-so/micro-sdk-go"
	"github.com/micro-so/micro-sdk-go/option"
)

// Both the parent service and its Options resource must retain client defaults
// and allow per-request overrides after generation.
func TestPropertyOptionsInheritRequestConfiguration(t *testing.T) {
	var paths []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		if got := r.Header.Get("X-API-Key"); got != "test-key" {
			t.Errorf("API key = %q", got)
		}
		if got := r.Header.Get("X-Request-Override"); got != "request" {
			t.Errorf("override = %q", got)
		}
		if got := r.URL.Query().Get("type"); got != "select_str" {
			t.Errorf("property type = %q", got)
		}
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()
	client := micro.NewClient(option.WithBaseURL(server.URL), option.WithAPIKey("test-key"), option.WithTeamID("test-team"), option.WithHeader("X-Request-Override", "client"))
	override := option.WithHeader("X-Request-Override", "request")
	if err := client.Prism.Properties.Delete(context.Background(), "document", "property-id", micro.PrismPropertyDeleteParams{Type: micro.F(micro.PrismPropertyDeleteParamsTypeSelectStr)}, override); err != nil {
		t.Fatal(err)
	}
	if err := client.Prism.Properties.Options.Delete(context.Background(), "document", "property-id", "option-id", micro.PrismPropertyOptionDeleteParams{Type: micro.F(micro.PrismPropertyOptionDeleteParamsTypeSelectStr)}, override); err != nil {
		t.Fatal(err)
	}
	want := []string{"/v2/prism/test-team/document/properties/property-id", "/v2/prism/test-team/document/properties/property-id/options/option-id"}
	if len(paths) != len(want) {
		t.Fatalf("requests = %v", paths)
	}
	for i := range want {
		if paths[i] != want[i] {
			t.Errorf("path = %q; want %q", paths[i], want[i])
		}
	}
}
