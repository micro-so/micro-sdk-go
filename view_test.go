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
)

func TestViewNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.New(
		context.TODO(),
		micro.ViewNewParamsViewObjectTypeAction,
		micro.ViewNewParams{
			Name:                 micro.F("name"),
			ViewType:             micro.F("view_type"),
			ID:                   micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AggregationPropDefID: micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AggregationType:      micro.F("aggregation_type"),
			ColumnLayout: micro.F(map[string]interface{}{
				"foo": "bar",
			}),
			Combinator: micro.F(micro.ViewNewParamsCombinatorAnd),
			CreatedAt:  micro.F("created_at"),
			Filter: micro.F([]map[string]interface{}{{
				"foo": "bar",
			}}),
			GroupBy:              micro.F("group_by"),
			GroupHiddenOptionIDs: micro.F([]interface{}{map[string]interface{}{}}),
			GroupHideEmpty:       micro.F(true),
			GroupSort:            micro.F("group_sort"),
			Icon:                 micro.F("icon"),
			ListID:               micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Select:               micro.F([]string{"string"}),
			Sort: micro.F([]map[string]interface{}{{
				"foo": "bar",
			}}),
			SortOrder:  micro.F(int64(0)),
			BodyTeamID: micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			UpdatedAt:  micro.F("updated_at"),
			UserID:     micro.F("user_id"),
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

func TestViewUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.Update(
		context.TODO(),
		micro.ViewUpdateParamsViewObjectTypeAction,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.ViewUpdateParams{
			AggregationPropDefID: micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AggregationType:      micro.F("aggregation_type"),
			ColumnLayout: micro.F(map[string]interface{}{
				"foo": "bar",
			}),
			Combinator: micro.F(micro.ViewUpdateParamsCombinatorAnd),
			Filter: micro.F([]map[string]interface{}{{
				"foo": "bar",
			}}),
			GroupBy:              micro.F("group_by"),
			GroupHiddenOptionIDs: micro.F([]interface{}{map[string]interface{}{}}),
			GroupHideEmpty:       micro.F(true),
			GroupSort:            micro.F("group_sort"),
			Icon:                 micro.F("icon"),
			ListID:               micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:                 micro.F("name"),
			Select:               micro.F([]string{"string"}),
			Sort: micro.F([]map[string]interface{}{{
				"foo": "bar",
			}}),
			SortOrder:  micro.F(int64(0)),
			BodyTeamID: micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			UserID:     micro.F("user_id"),
			ViewType:   micro.F("view_type"),
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

func TestViewDelete(t *testing.T) {
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
	err := client.Views.Delete(
		context.TODO(),
		micro.ViewDeleteParamsViewObjectTypeAction,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.ViewDeleteParams{},
	)
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestViewGet(t *testing.T) {
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
	_, err := client.Views.Get(
		context.TODO(),
		micro.ViewGetParamsViewObjectTypeAction,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.ViewGetParams{},
	)
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
