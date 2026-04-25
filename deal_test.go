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

func TestDealNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Deals.New(context.TODO(), micro.DealNewParams{
		PrismObjectProperties: micro.PrismObjectPropertiesParam{
			ID:  micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			CRM: micro.F[any](map[string]interface{}{}),
			Default: micro.F(map[string]interface{}{
				"foo": "bar",
			}),
			Extended: micro.F[any](map[string]interface{}{}),
		},
	})
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDealUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Deals.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.DealUpdateParams{
			PrismObjectProperties: micro.PrismObjectPropertiesParam{
				ID:  micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				CRM: micro.F[any](map[string]interface{}{}),
				Default: micro.F(map[string]interface{}{
					"foo": "bar",
				}),
				Extended: micro.F[any](map[string]interface{}{}),
			},
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

func TestDealListWithOptionalParams(t *testing.T) {
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
	_, err := client.Deals.List(context.TODO(), micro.DealListParams{
		Query: micro.F(micro.DealListParamsQuery{
			Select:     micro.F([]string{"string"}),
			Combinator: micro.F(micro.DealListParamsQueryCombinatorAnd),
			CRMID:      micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Filter: micro.F([]map[string]map[string]micro.DealListParamsQueryFilterUnion{{
				"foo": {
					"foo": shared.UnionString("string"),
				},
			}}),
			Limit: micro.F(int64(0)),
			Page:  micro.F(int64(0)),
			Sort: micro.F([]map[string]micro.DealListParamsQuerySort{{
				"foo": micro.DealListParamsQuerySortAsc,
			}}),
		}),
		ID:      micro.F[micro.DealListParamsIDUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
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

func TestDealDelete(t *testing.T) {
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
	err := client.Deals.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.DealDeleteParams{},
	)
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDealImportWithOptionalParams(t *testing.T) {
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
	_, err := client.Deals.Import(context.TODO(), micro.DealImportParams{
		Objects: micro.F([]micro.PrismObjectPropertiesParam{{
			ID:  micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			CRM: micro.F[any](map[string]interface{}{}),
			Default: micro.F(map[string]interface{}{
				"foo": "bar",
			}),
			Extended: micro.F[any](map[string]interface{}{}),
		}}),
		Options: micro.F(micro.DealImportParamsOptions{
			CaseInsensitive: micro.F(true),
			CRMID:           micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DedupeBy:        micro.F("dedupe_by"),
		}),
	})
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
