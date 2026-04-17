// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/micro-go"
	"github.com/stainless-sdks/micro-go/internal/testutil"
	"github.com/stainless-sdks/micro-go/option"
	"github.com/stainless-sdks/micro-go/shared"
)

func TestPrismQueryExecuteWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Query.Execute(
		context.TODO(),
		micro.PrismQueryExecuteParamsObjectTypeDeal,
		micro.PrismQueryExecuteParams{
			Query: micro.F(micro.PrismQueryExecuteParamsQuery{
				Select:     micro.F([]string{"string"}),
				Combinator: micro.F(micro.PrismQueryExecuteParamsQueryCombinatorAnd),
				CRMID:      micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Filter: micro.F([]map[string]map[string]micro.PrismQueryExecuteParamsQueryFilterUnion{{
					"foo": {
						"foo": shared.UnionString("string"),
					},
				}}),
				Limit: micro.F(int64(0)),
				Page:  micro.F(int64(0)),
				Sort: micro.F([]map[string]micro.PrismQueryExecuteParamsQuerySort{{
					"foo": micro.PrismQueryExecuteParamsQuerySortAsc,
				}}),
			}),
			ID:      micro.F[micro.PrismQueryExecuteParamsIDUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
			Boxes:   micro.F([]string{"string"}),
			Deleted: micro.F(true),
			Sources: micro.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		},
	)
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
