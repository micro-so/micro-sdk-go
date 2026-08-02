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

func TestPrismObjectActionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Actions.New(context.TODO(), micro.PrismObjectActionNewParams{
		PrismObjectProperties: micro.PrismObjectPropertiesParam{
			Default: micro.F(map[string]interface{}{
				"full_name":    "bar",
				"email":        "bar",
				"title":        "bar",
				"organization": "bar",
			}),
			List: micro.F[any](map[string]interface{}{}),
		},
		IdempotencyKey: micro.F("x"),
	})
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrismObjectActionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Actions.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismObjectActionUpdateParams{
			PrismObjectProperties: micro.PrismObjectPropertiesParam{
				Default: micro.F(map[string]interface{}{
					"title": "bar",
				}),
				List: micro.F[any](map[string]interface{}{}),
			},
			IdempotencyKey: micro.F("x"),
			IfMatch:        micro.F("If-Match"),
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

func TestPrismObjectActionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Actions.List(context.TODO(), micro.PrismObjectActionListParams{
		Cursor:       micro.F("cursor"),
		Deleted:      micro.F(true),
		IncludeTotal: micro.F(true),
		Limit:        micro.F(int64(1)),
		ListID:       micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Select:       micro.F("select"),
		Sort:         micro.F("sort"),
	})
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrismObjectActionDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Prism.Objects.Actions.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismObjectActionDeleteParams{
			IfMatch: micro.F("If-Match"),
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

func TestPrismObjectActionBulkNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Actions.BulkNew(context.TODO(), micro.PrismObjectActionBulkNewParams{
		Objects: micro.F([]micro.PrismObjectPropertiesParam{{
			Default: micro.F(map[string]interface{}{
				"foo": "bar",
			}),
			List: micro.F[any](map[string]interface{}{}),
		}}),
		Options: micro.F(micro.PrismObjectActionBulkNewParamsOptions{
			CaseInsensitive:      micro.F(true),
			CreateMissingOptions: micro.F(true),
			CRMID:                micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DedupeBy:             micro.F[micro.PrismObjectActionBulkNewParamsOptionsDedupeByUnion](shared.UnionString("string")),
			ListID:               micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			RequireListStage:     micro.F(true),
			UpdateExisting:       micro.F(true),
		}),
		IdempotencyKey: micro.F("x"),
	})
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrismObjectActionBulkDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Actions.BulkDelete(context.TODO(), micro.PrismObjectActionBulkDeleteParams{
		IDs:            micro.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		IdempotencyKey: micro.F("x"),
	})
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrismObjectActionBulkUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Actions.BulkUpdate(context.TODO(), micro.PrismObjectActionBulkUpdateParams{
		Items: micro.F([]micro.PrismObjectActionBulkUpdateParamsItem{{
			ID: micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}}),
		IdempotencyKey: micro.F("x"),
	})
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrismObjectActionCountWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Actions.Count(context.TODO(), micro.PrismObjectActionCountParams{
		ListID: micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrismObjectActionDuplicateWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Actions.Duplicate(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismObjectActionDuplicateParams{
			IdempotencyKey: micro.F("x"),
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

func TestPrismObjectActionFindWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Actions.Find(
		context.TODO(),
		"slug",
		"value",
		micro.PrismObjectActionFindParams{
			ListID: micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestPrismObjectActionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Actions.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismObjectActionGetParams{
			Select: micro.F("select"),
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

func TestPrismObjectActionQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Actions.Query(context.TODO(), micro.PrismObjectActionQueryParams{
		Query: micro.F(micro.PrismObjectActionQueryParamsQuery{
			Select:     micro.F([]string{"full_name", "email", "title", "organization"}),
			Combinator: micro.F(micro.PrismObjectActionQueryParamsQueryCombinatorAnd),
			Cursor:     micro.F("cursor"),
			Filter: micro.F([]map[string]micro.PrismObjectActionQueryParamsQueryFilterUnion{{
				"full_name": micro.PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEq{
					Equals: micro.F[micro.PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion](shared.UnionString("Sarah Chen")),
				},
			}}),
			Limit:  micro.F(int64(10)),
			ListID: micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Page:   micro.F(int64(0)),
			Sort: micro.F([]map[string]micro.PrismObjectActionQueryParamsQuerySort{{
				"foo": micro.PrismObjectActionQueryParamsQuerySortAsc,
			}}),
		}),
		ID:             micro.F[micro.PrismObjectActionQueryParamsIDUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
		Boxes:          micro.F([]string{"string"}),
		Cursor:         micro.F("cursor"),
		Deleted:        micro.F(true),
		IncludeTotal:   micro.F(true),
		Sources:        micro.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		IdempotencyKey: micro.F("x"),
	})
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrismObjectActionRestoreWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Actions.Restore(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismObjectActionRestoreParams{
			IdempotencyKey: micro.F("x"),
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

func TestPrismObjectActionUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Actions.Upsert(
		context.TODO(),
		"slug",
		"value",
		micro.PrismObjectActionUpsertParams{
			PrismObjectProperties: micro.PrismObjectPropertiesParam{
				Default: micro.F(map[string]interface{}{
					"foo": "bar",
				}),
				List: micro.F[any](map[string]interface{}{}),
			},
			ListID:         micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			IdempotencyKey: micro.F("x"),
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
