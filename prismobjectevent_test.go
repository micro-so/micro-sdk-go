// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/micro-so/micro-sdk-go"
	"github.com/micro-so/micro-sdk-go/internal/testutil"
	"github.com/micro-so/micro-sdk-go/option"
	"github.com/micro-so/micro-sdk-go/shared"
)

func TestPrismObjectEventGet(t *testing.T) {
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
	_, err := client.Prism.Objects.Events.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismObjectEventGetParams{},
	)
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrismObjectEventQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Events.Query(context.TODO(), micro.PrismObjectEventQueryParams{
		Query: micro.F(micro.PrismObjectEventQueryParamsQuery{
			Select:     micro.F([]string{"string"}),
			Combinator: micro.F(micro.PrismObjectEventQueryParamsQueryCombinatorAnd),
			Filter: micro.F([]map[string]micro.PrismObjectEventQueryParamsQueryFilterUnion{{
				"foo": micro.PrismObjectEventQueryParamsQueryFilter{
					Equals: micro.F[micro.PrismObjectEventQueryParamsQueryFilterUnion](shared.UnionString("string")),
				},
			}}),
			Limit:  micro.F(int64(1)),
			ListID: micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Page:   micro.F(int64(0)),
			Sort: micro.F([]map[string]micro.PrismObjectEventQueryParamsQuerySort{{
				"foo": micro.PrismObjectEventQueryParamsQuerySortAsc,
			}}),
		}),
		ID:      micro.F[micro.PrismObjectEventQueryParamsIDUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
		Boxes:   micro.F([]string{"string"}),
		Deleted: micro.F(true),
		Sources: micro.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
	})
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
