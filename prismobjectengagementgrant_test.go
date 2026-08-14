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

func TestPrismObjectEngagementGrantUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Objects.Engagements.Grant.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismObjectEngagementGrantUpdateParams{
			ShareLevel: micro.F(micro.PrismObjectEngagementGrantUpdateParamsShareLevelMetadata),
			TeamGroupID: micro.F([]map[string]micro.PrismObjectEngagementGrantUpdateParamsTeamGroupID{{
				"foo": micro.PrismObjectEngagementGrantUpdateParamsTeamGroupIDA,
			}}),
			BodyTeamID: micro.F(map[string]micro.PrismObjectEngagementGrantUpdateParamsTeamID{
				"foo": micro.PrismObjectEngagementGrantUpdateParamsTeamIDA,
			}),
			UserID: micro.F([]map[string]micro.PrismObjectEngagementGrantUpdateParamsUserID{{
				"foo": micro.PrismObjectEngagementGrantUpdateParamsUserIDA,
			}}),
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

func TestPrismObjectEngagementGrantGet(t *testing.T) {
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
	_, err := client.Prism.Objects.Engagements.Grant.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismObjectEngagementGrantGetParams{},
	)
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
