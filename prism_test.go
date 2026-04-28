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

func TestPrismNewObjectWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.NewObject(
		context.TODO(),
		micro.ObjectTypeDeal,
		micro.PrismNewObjectParams{
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

func TestPrismDeleteObject(t *testing.T) {
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
	err := client.Prism.DeleteObject(
		context.TODO(),
		micro.ObjectTypeDeal,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismDeleteObjectParams{},
	)
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrismDuplicateObject(t *testing.T) {
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
	_, err := client.Prism.DuplicateObject(
		context.TODO(),
		micro.ObjectTypeDeal,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismDuplicateObjectParams{},
	)
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrismImportObjectsWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.ImportObjects(
		context.TODO(),
		micro.PrismImportObjectsParamsObjectTypeIdentity,
		micro.PrismImportObjectsParams{
			Objects: micro.F([]micro.PrismObjectPropertiesParam{{
				ID:  micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				CRM: micro.F[any](map[string]interface{}{}),
				Default: micro.F(map[string]interface{}{
					"foo": "bar",
				}),
				Extended: micro.F[any](map[string]interface{}{}),
			}}),
			Options: micro.F(micro.PrismImportObjectsParamsOptions{
				CaseInsensitive: micro.F(true),
				CRMID:           micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				DedupeBy:        micro.F("dedupe_by"),
			}),
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

func TestPrismPatchObjectWithOptionalParams(t *testing.T) {
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
	_, err := client.Prism.PatchObject(
		context.TODO(),
		micro.ObjectTypeDeal,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismPatchObjectParams{
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

func TestPrismRestoreObject(t *testing.T) {
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
	_, err := client.Prism.RestoreObject(
		context.TODO(),
		micro.ObjectTypeDeal,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.PrismRestoreObjectParams{},
	)
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
