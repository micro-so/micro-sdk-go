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
)

func TestPrismGrantGetGrant(t *testing.T) {
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
	_, err := client.Prism.Grant.GetGrant(
		context.TODO(),
		micro.ObjectTypeDeal,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismGrantGetGrantParams{},
	)
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrismGrantUpdateGrantWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Grant.UpdateGrant(
		context.TODO(),
		micro.ObjectTypeDeal,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismGrantUpdateGrantParams{
			TeamGroupID: micro.F([]map[string]micro.PrismGrantUpdateGrantParamsTeamGroupID{{
				"foo": micro.PrismGrantUpdateGrantParamsTeamGroupIDA,
			}}),
			BodyTeamID: micro.F(map[string]micro.PrismGrantUpdateGrantParamsTeamID{
				"foo": micro.PrismGrantUpdateGrantParamsTeamIDA,
			}),
			UserID: micro.F([]map[string]micro.PrismGrantUpdateGrantParamsUserID{{
				"foo": micro.PrismGrantUpdateGrantParamsUserIDA,
			}}),
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
