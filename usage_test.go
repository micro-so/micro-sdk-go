// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro_test

import (
	"context"
	"os"
	"testing"

	"github.com/micro-so/micro-sdk-go"
	"github.com/micro-so/micro-sdk-go/internal/testutil"
	"github.com/micro-so/micro-sdk-go/option"
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
	response, err := client.Prism.Objects.Deals.Query(context.TODO(), micro.PrismObjectDealQueryParams{
		Query: micro.F(micro.PrismObjectDealQueryParamsQuery{
			Select: micro.F([]string{"id", "name"}),
		}),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", response.Data)
}
