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

func TestTriggeredAutomationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.TriggeredAutomations.New(
		context.TODO(),
		micro.TriggeredAutomationNewParamsAutomationObjectTypeMessage,
		micro.TriggeredAutomationNewParams{
			TriggeredAutomation: micro.TriggeredAutomationParam{
				Kind: micro.F(micro.TriggeredAutomationKindUpdate),
				Name: micro.F("name"),
				ID:   micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Actions: micro.F([]micro.TriggeredAutomationActionParam{{
					Type:                       micro.F(micro.TriggeredAutomationActionsTypeAgent),
					AgentID:                    micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					CronExpression:             micro.F("cron_expression"),
					DelaySeconds:               micro.F(int64(0)),
					RecipientEmailPropDefID:    micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					RecipientProviderPropDefID: micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					RecipientViewID:            micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					RecipientViewObjectType:    micro.F("recipient_view_object_type"),
					SendAsUserID:               micro.F("send_as_user_id"),
					Subject:                    micro.F("subject"),
					TemplateID:                 micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Timezone:                   micro.F("timezone"),
					WebhookID:                  micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				}}),
				Changeset: micro.F(micro.TriggeredAutomationChangesetParam{
					Combinator: micro.F(micro.TriggeredAutomationChangesetCombinatorAnd),
					Filter: micro.F([]map[string]interface{}{{
						"foo": "bar",
					}}),
				}),
				CreatedAt: micro.F("created_at"),
				Enabled:   micro.F(true),
				ListID:    micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				OnCreate:  micro.F(true),
				OnDelete:  micro.F(true),
				State: micro.F(micro.TriggeredAutomationStateParam{
					Combinator: micro.F(micro.TriggeredAutomationStateCombinatorAnd),
					Filter: micro.F([]map[string]interface{}{{
						"foo": "bar",
					}}),
				}),
				TeamID:    micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				UpdatedAt: micro.F("updated_at"),
				UserID:    micro.F("user_id"),
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

func TestTriggeredAutomationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.TriggeredAutomations.Update(
		context.TODO(),
		micro.TriggeredAutomationUpdateParamsAutomationObjectTypeMessage,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.TriggeredAutomationUpdateParams{
			TriggeredAutomation: micro.TriggeredAutomationParam{
				Kind: micro.F(micro.TriggeredAutomationKindUpdate),
				Name: micro.F("name"),
				ID:   micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Actions: micro.F([]micro.TriggeredAutomationActionParam{{
					Type:                       micro.F(micro.TriggeredAutomationActionsTypeAgent),
					AgentID:                    micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					CronExpression:             micro.F("cron_expression"),
					DelaySeconds:               micro.F(int64(0)),
					RecipientEmailPropDefID:    micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					RecipientProviderPropDefID: micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					RecipientViewID:            micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					RecipientViewObjectType:    micro.F("recipient_view_object_type"),
					SendAsUserID:               micro.F("send_as_user_id"),
					Subject:                    micro.F("subject"),
					TemplateID:                 micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Timezone:                   micro.F("timezone"),
					WebhookID:                  micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				}}),
				Changeset: micro.F(micro.TriggeredAutomationChangesetParam{
					Combinator: micro.F(micro.TriggeredAutomationChangesetCombinatorAnd),
					Filter: micro.F([]map[string]interface{}{{
						"foo": "bar",
					}}),
				}),
				CreatedAt: micro.F("created_at"),
				Enabled:   micro.F(true),
				ListID:    micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				OnCreate:  micro.F(true),
				OnDelete:  micro.F(true),
				State: micro.F(micro.TriggeredAutomationStateParam{
					Combinator: micro.F(micro.TriggeredAutomationStateCombinatorAnd),
					Filter: micro.F([]map[string]interface{}{{
						"foo": "bar",
					}}),
				}),
				TeamID:    micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				UpdatedAt: micro.F("updated_at"),
				UserID:    micro.F("user_id"),
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

func TestTriggeredAutomationListWithOptionalParams(t *testing.T) {
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
	_, err := client.TriggeredAutomations.List(
		context.TODO(),
		micro.TriggeredAutomationListParamsAutomationObjectTypeMessage,
		micro.TriggeredAutomationListParams{
			Cursor: micro.F("cursor"),
			Kind:   micro.F(micro.TriggeredAutomationListParamsKindUpdate),
			Limit:  micro.F(int64(0)),
			ListID: micro.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Page:   micro.F(int64(1)),
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

func TestTriggeredAutomationDelete(t *testing.T) {
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
	err := client.TriggeredAutomations.Delete(
		context.TODO(),
		micro.TriggeredAutomationDeleteParamsAutomationObjectTypeMessage,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.TriggeredAutomationDeleteParams{},
	)
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTriggeredAutomationGet(t *testing.T) {
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
	_, err := client.TriggeredAutomations.Get(
		context.TODO(),
		micro.TriggeredAutomationGetParamsAutomationObjectTypeMessage,
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		micro.TriggeredAutomationGetParams{},
	)
	if err != nil {
		var apierr *micro.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
