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

func TestPrismPropertyOptionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Properties.Options.New(
		context.TODO(),
		micro.PrismPropertyOptionNewParamsObjectTypeComment,
		"2fdcD1Dc-bbDb-2BBD-0Afa-1A3C33cFaADc",
		micro.PrismPropertyOptionNewParams{
			PropertyOptionCreate: micro.PropertyOptionCreateParam{
				Type:        micro.F(micro.PropertyOptionCreateTypeNum),
				Value:       micro.F("value"),
				ColorScheme: micro.F("color_scheme"),
				Description: micro.F("description"),
				Icon:        micro.F("icon"),
				ListID:      micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				OptionGroup: micro.F("option_group"),
				Slug:        micro.F("slug"),
				SortIndex:   micro.F(int64(0)),
			},
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

func TestPrismPropertyOptionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Properties.Options.Update(
		context.TODO(),
		micro.PrismPropertyOptionUpdateParamsObjectTypeComment,
		"2fdcD1Dc-bbDb-2BBD-0Afa-1A3C33cFaADc",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismPropertyOptionUpdateParams{
			PropertyOptionPatch: micro.PropertyOptionPatchParam{
				Type:        micro.F(micro.PropertyOptionPatchTypeNum),
				ColorScheme: micro.F("color_scheme"),
				Description: micro.F("description"),
				Enabled:     micro.F(true),
				Icon:        micro.F("icon"),
				ListID:      micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				OptionGroup: micro.F("option_group"),
				Slug:        micro.F("slug"),
				SortIndex:   micro.F(int64(0)),
				Value:       micro.F("value"),
			},
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

func TestPrismPropertyOptionDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Prism.Properties.Options.Delete(
		context.TODO(),
		micro.PrismPropertyOptionDeleteParamsObjectTypeComment,
		"2fdcD1Dc-bbDb-2BBD-0Afa-1A3C33cFaADc",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismPropertyOptionDeleteParams{
			Type:   micro.F(micro.PrismPropertyOptionDeleteParamsTypeNum),
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
