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

func TestPrismPropertyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Properties.New(
		context.TODO(),
		micro.PrismPropertyNewParamsObjectTypeComment,
		micro.PrismPropertyNewParams{
			PropertyDefinitionCreate: micro.PropertyDefinitionCreateParam{
				Name:   micro.F("name"),
				Type:   micro.F(micro.PropertyDefinitionCreateTypeNum),
				Icon:   micro.F("icon"),
				ListID: micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Options: micro.F([]micro.PropertyDefinitionCreateOptionParam{{
					Value:       micro.F("value"),
					ColorScheme: micro.F("color_scheme"),
					Description: micro.F("description"),
					Icon:        micro.F("icon"),
					OptionGroup: micro.F("option_group"),
					Slug:        micro.F("slug"),
					SortIndex:   micro.F(int64(0)),
				}}),
				Required: micro.F(true),
				RoleID:   micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Slug:     micro.F("slug"),
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

func TestPrismPropertyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Properties.Update(
		context.TODO(),
		micro.PrismPropertyUpdateParamsObjectTypeComment,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismPropertyUpdateParams{
			PropertyDefinitionPatch: micro.PropertyDefinitionPatchParam{
				Type:     micro.F(micro.PropertyDefinitionPatchTypeNum),
				Enabled:  micro.F(true),
				Icon:     micro.F("icon"),
				ListID:   micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Name:     micro.F("name"),
				Required: micro.F(true),
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

func TestPrismPropertyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Properties.List(
		context.TODO(),
		micro.PrismPropertyListParamsObjectTypeComment,
		micro.PrismPropertyListParams{
			Autofill:       micro.F(true),
			IncludeOptions: micro.F[micro.PrismPropertyListParamsIncludeOptionsUnion](micro.PrismPropertyListParamsIncludeOptionsString(micro.PrismPropertyListParamsIncludeOptionsStringTrue)),
			ListID:         micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Term:           micro.F("term"),
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

func TestPrismPropertyDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Prism.Properties.Delete(
		context.TODO(),
		micro.PrismPropertyDeleteParamsObjectTypeComment,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismPropertyDeleteParams{
			Type:   micro.F(micro.PrismPropertyDeleteParamsTypeNum),
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

func TestPrismPropertyListAllWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.Properties.ListAll(context.TODO(), micro.PrismPropertyListAllParams{
		Autofill:       micro.F(true),
		IncludeOptions: micro.F[micro.PrismPropertyListAllParamsIncludeOptionsUnion](micro.PrismPropertyListAllParamsIncludeOptionsString(micro.PrismPropertyListAllParamsIncludeOptionsStringTrue)),
		ListID:         micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Term:           micro.F("term"),
	})
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
