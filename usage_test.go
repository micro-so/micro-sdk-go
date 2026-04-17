// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/micro-go"
	"github.com/stainless-sdks/micro-go/internal/testutil"
	"github.com/stainless-sdks/micro-go/option"
)

func TestUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := micro.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithTeamID("My Team ID"),
	)
	err := client.Prism.RestoreObject(
		context.TODO(),
		micro.ObjectTypeDeal,
		"REPLACE_ME",
		micro.PrismRestoreObjectParams{},
	)
	if err != nil {
		t.Error(err)
		return
	}
}
