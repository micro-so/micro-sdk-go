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

func TestContactNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Contacts.New(context.TODO(), micro.ContactNewParams{
		PrismObjectProperties: micro.PrismObjectPropertiesParam{
			ID:       micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			CRM:      micro.F[any](map[string]interface{}{}),
			Default:  micro.F[any](map[string]interface{}{}),
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

func TestContactUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Contacts.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.ContactUpdateParams{
			PrismObjectProperties: micro.PrismObjectPropertiesParam{
				ID:       micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				CRM:      micro.F[any](map[string]interface{}{}),
				Default:  micro.F[any](map[string]interface{}{}),
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

func TestContactListWithOptionalParams(t *testing.T) {
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
	_, err := client.Contacts.List(context.TODO(), micro.ContactListParams{
		Query: micro.F(micro.ContactListParamsQuery{
			Select:     micro.F([]string{"string"}),
			Combinator: micro.F(micro.ContactListParamsQueryCombinatorAnd),
			CRMID:      micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Filter: micro.F([]map[string]map[string]micro.ContactListParamsQueryFilterUnion{{
				"foo": {
					"foo": shared.UnionString("string"),
				},
			}}),
			Limit: micro.F(int64(0)),
			Page:  micro.F(int64(0)),
			Sort: micro.F([]map[string]micro.ContactListParamsQuerySort{{
				"foo": micro.ContactListParamsQuerySortAsc,
			}}),
		}),
		ID:      micro.F[micro.ContactListParamsIDUnion](shared.UnionString("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")),
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

func TestContactDelete(t *testing.T) {
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
	err := client.Contacts.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.ContactDeleteParams{},
	)
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContactImportWithOptionalParams(t *testing.T) {
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
	_, err := client.Contacts.Import(context.TODO(), micro.ContactImportParams{
		Objects: micro.F([]micro.PrismObjectPropertiesParam{{
			ID:       micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			CRM:      micro.F[any](map[string]interface{}{}),
			Default:  micro.F[any](map[string]interface{}{}),
			Extended: micro.F[any](map[string]interface{}{}),
		}}),
		Options: micro.F(micro.ContactImportParamsOptions{
			CaseInsensitive: micro.F(true),
			CRMID:           micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DedupeBy:        micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DedupeType:      micro.F(micro.ContactImportParamsOptionsDedupeTypeStr),
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
