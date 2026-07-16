// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/micro-so/micro-sdk-go/internal/apijson"
	"github.com/micro-so/micro-sdk-go/internal/apiquery"
	"github.com/micro-so/micro-sdk-go/internal/param"
	"github.com/micro-so/micro-sdk-go/internal/requestconfig"
	"github.com/micro-so/micro-sdk-go/option"
)

// TriggeredAutomationService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTriggeredAutomationService] method instead.
type TriggeredAutomationService struct {
	Options []option.RequestOption
}

// NewTriggeredAutomationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTriggeredAutomationService(opts ...option.RequestOption) (r *TriggeredAutomationService) {
	r = &TriggeredAutomationService{}
	r.Options = opts
	return
}

// Create a triggered automation (state + changeset filter trees)
func (r *TriggeredAutomationService) New(ctx context.Context, automationObjectType TriggeredAutomationNewParamsAutomationObjectType, params TriggeredAutomationNewParams, opts ...option.RequestOption) (res *TriggeredAutomation, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.TeamID, precfg.TeamID)
	if params.TeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/triggered_automations", params.TeamID, automationObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Replace a triggered automation (idempotent full write of the whole tree)
func (r *TriggeredAutomationService) Update(ctx context.Context, automationObjectType TriggeredAutomationUpdateParamsAutomationObjectType, automationID string, params TriggeredAutomationUpdateParams, opts ...option.RequestOption) (res *TriggeredAutomation, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.TeamID, precfg.TeamID)
	if params.TeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return nil, err
	}
	if automationID == "" {
		err = errors.New("missing required automationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/triggered_automations/%s", params.TeamID, automationObjectType, automationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List triggered automations for an owner
func (r *TriggeredAutomationService) List(ctx context.Context, automationObjectType TriggeredAutomationListParamsAutomationObjectType, params TriggeredAutomationListParams, opts ...option.RequestOption) (res *TriggeredAutomationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.TeamID, precfg.TeamID)
	if params.TeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/triggered_automations", params.TeamID, automationObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a triggered automation and its filter trees
func (r *TriggeredAutomationService) Delete(ctx context.Context, automationObjectType TriggeredAutomationDeleteParamsAutomationObjectType, automationID string, body TriggeredAutomationDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return err
	}
	requestconfig.UseDefaultParam(&body.TeamID, precfg.TeamID)
	if body.TeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return err
	}
	if automationID == "" {
		err = errors.New("missing required automationId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/triggered_automations/%s", body.TeamID, automationObjectType, automationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Read a triggered automation
func (r *TriggeredAutomationService) Get(ctx context.Context, automationObjectType TriggeredAutomationGetParamsAutomationObjectType, automationID string, query TriggeredAutomationGetParams, opts ...option.RequestOption) (res *TriggeredAutomation, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.TeamID, precfg.TeamID)
	if query.TeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return nil, err
	}
	if automationID == "" {
		err = errors.New("missing required automationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/triggered_automations/%s", query.TeamID, automationObjectType, automationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A triggered automation. `kind` selects the shape: `update` fires on object
// updates and requires a `changeset` (from/to transition) filter plus an optional
// `state` precondition; `lifecycle` fires on create and/or delete
// (`on_create`/`on_delete`) and requires a `state` filter (no changeset). `state`
// permits dot-paths (nested reference filters); `changeset` is direct properties
// only. Object type is taken from the path.
type TriggeredAutomation struct {
	Kind TriggeredAutomationKind `json:"kind" api:"required"`
	Name string                  `json:"name" api:"required"`
	ID   string                  `json:"id" format:"uuid"`
	// Actions to run when the automation fires; each item has a `type` plus
	// type-specific fields.
	Actions []TriggeredAutomationAction `json:"actions"`
	// A changeset filter group (update automations only): a combinator plus an array
	// of transition clauses matching what is changing. Dot-paths (nested reference
	// filters) are NOT permitted — direct properties only.
	Changeset TriggeredAutomationChangeset `json:"changeset"`
	CreatedAt string                       `json:"created_at"`
	Enabled   bool                         `json:"enabled"`
	ListID    string                       `json:"list_id" api:"nullable" format:"uuid"`
	// Lifecycle automations only.
	OnCreate bool `json:"on_create"`
	// Lifecycle automations only.
	OnDelete bool `json:"on_delete"`
	// A filter group: a combinator plus an array of slug-based clauses. Dot-paths
	// (e.g. `organization.location`) express nested reference filters.
	State     TriggeredAutomationState `json:"state"`
	TeamID    string                   `json:"team_id" api:"nullable" format:"uuid"`
	UpdatedAt string                   `json:"updated_at" api:"nullable"`
	UserID    string                   `json:"user_id" api:"nullable"`
	JSON      triggeredAutomationJSON  `json:"-"`
}

// triggeredAutomationJSON contains the JSON metadata for the struct
// [TriggeredAutomation]
type triggeredAutomationJSON struct {
	Kind        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Actions     apijson.Field
	Changeset   apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ListID      apijson.Field
	OnCreate    apijson.Field
	OnDelete    apijson.Field
	State       apijson.Field
	TeamID      apijson.Field
	UpdatedAt   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TriggeredAutomation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r triggeredAutomationJSON) RawJSON() string {
	return r.raw
}

type TriggeredAutomationKind string

const (
	TriggeredAutomationKindUpdate    TriggeredAutomationKind = "update"
	TriggeredAutomationKindLifecycle TriggeredAutomationKind = "lifecycle"
)

func (r TriggeredAutomationKind) IsKnown() bool {
	switch r {
	case TriggeredAutomationKindUpdate, TriggeredAutomationKindLifecycle:
		return true
	}
	return false
}

// An action the automation runs when it fires. `type` selects the kind; the
// remaining fields are type-specific (`agent` → `agent_id`, `webhook` →
// `webhook_id`). Generic: new action types add fields here.
type TriggeredAutomationAction struct {
	Type TriggeredAutomationActionsType `json:"type" api:"required"`
	// Required when `type` is `agent`. The agent to run.
	AgentID string `json:"agent_id" api:"nullable" format:"uuid"`
	// wait: cron schedule for the resume time. Exactly one of delay_seconds or
	// cron_expression.
	CronExpression string `json:"cron_expression" api:"nullable"`
	// wait: relative delay in seconds. Exactly one of delay_seconds or
	// cron_expression.
	DelaySeconds int64 `json:"delay_seconds" api:"nullable"`
	// wait: IANA timezone for evaluating cron_expression (optional).
	Timezone string `json:"timezone" api:"nullable"`
	// Required when `type` is `webhook`. The id of the webhook the event is dispatched
	// to (async) when the automation fires.
	WebhookID   string                        `json:"webhook_id" api:"nullable" format:"uuid"`
	ExtraFields map[string]interface{}        `json:"-" api:"extrafields"`
	JSON        triggeredAutomationActionJSON `json:"-"`
}

// triggeredAutomationActionJSON contains the JSON metadata for the struct
// [TriggeredAutomationAction]
type triggeredAutomationActionJSON struct {
	Type           apijson.Field
	AgentID        apijson.Field
	CronExpression apijson.Field
	DelaySeconds   apijson.Field
	Timezone       apijson.Field
	WebhookID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TriggeredAutomationAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r triggeredAutomationActionJSON) RawJSON() string {
	return r.raw
}

type TriggeredAutomationActionsType string

const (
	TriggeredAutomationActionsTypeAgent   TriggeredAutomationActionsType = "agent"
	TriggeredAutomationActionsTypeWebhook TriggeredAutomationActionsType = "webhook"
	TriggeredAutomationActionsTypeWait    TriggeredAutomationActionsType = "wait"
)

func (r TriggeredAutomationActionsType) IsKnown() bool {
	switch r {
	case TriggeredAutomationActionsTypeAgent, TriggeredAutomationActionsTypeWebhook, TriggeredAutomationActionsTypeWait:
		return true
	}
	return false
}

// A changeset filter group (update automations only): a combinator plus an array
// of transition clauses matching what is changing. Dot-paths (nested reference
// filters) are NOT permitted — direct properties only.
type TriggeredAutomationChangeset struct {
	Combinator TriggeredAutomationChangesetCombinator `json:"combinator"`
	// Each entry is a transition clause { slug: { from?: { comparator: value }, to?: {
	// comparator: value } } }. `from` matches the prior value, `to` the new value; an
	// empty body { slug: {} } matches any change to that property.
	Filter []map[string]interface{}         `json:"filter"`
	JSON   triggeredAutomationChangesetJSON `json:"-"`
}

// triggeredAutomationChangesetJSON contains the JSON metadata for the struct
// [TriggeredAutomationChangeset]
type triggeredAutomationChangesetJSON struct {
	Combinator  apijson.Field
	Filter      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TriggeredAutomationChangeset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r triggeredAutomationChangesetJSON) RawJSON() string {
	return r.raw
}

type TriggeredAutomationChangesetCombinator string

const (
	TriggeredAutomationChangesetCombinatorAnd TriggeredAutomationChangesetCombinator = "AND"
	TriggeredAutomationChangesetCombinatorOr  TriggeredAutomationChangesetCombinator = "OR"
)

func (r TriggeredAutomationChangesetCombinator) IsKnown() bool {
	switch r {
	case TriggeredAutomationChangesetCombinatorAnd, TriggeredAutomationChangesetCombinatorOr:
		return true
	}
	return false
}

// A filter group: a combinator plus an array of slug-based clauses. Dot-paths
// (e.g. `organization.location`) express nested reference filters.
type TriggeredAutomationState struct {
	Combinator TriggeredAutomationStateCombinator `json:"combinator"`
	// Each entry is { slug: { comparator: value } }
	Filter []map[string]interface{}     `json:"filter"`
	JSON   triggeredAutomationStateJSON `json:"-"`
}

// triggeredAutomationStateJSON contains the JSON metadata for the struct
// [TriggeredAutomationState]
type triggeredAutomationStateJSON struct {
	Combinator  apijson.Field
	Filter      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TriggeredAutomationState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r triggeredAutomationStateJSON) RawJSON() string {
	return r.raw
}

type TriggeredAutomationStateCombinator string

const (
	TriggeredAutomationStateCombinatorAnd TriggeredAutomationStateCombinator = "AND"
	TriggeredAutomationStateCombinatorOr  TriggeredAutomationStateCombinator = "OR"
)

func (r TriggeredAutomationStateCombinator) IsKnown() bool {
	switch r {
	case TriggeredAutomationStateCombinatorAnd, TriggeredAutomationStateCombinatorOr:
		return true
	}
	return false
}

// A triggered automation. `kind` selects the shape: `update` fires on object
// updates and requires a `changeset` (from/to transition) filter plus an optional
// `state` precondition; `lifecycle` fires on create and/or delete
// (`on_create`/`on_delete`) and requires a `state` filter (no changeset). `state`
// permits dot-paths (nested reference filters); `changeset` is direct properties
// only. Object type is taken from the path.
type TriggeredAutomationParam struct {
	Kind param.Field[TriggeredAutomationKind] `json:"kind" api:"required"`
	Name param.Field[string]                  `json:"name" api:"required"`
	ID   param.Field[string]                  `json:"id" format:"uuid"`
	// Actions to run when the automation fires; each item has a `type` plus
	// type-specific fields.
	Actions param.Field[[]TriggeredAutomationActionParam] `json:"actions"`
	// A changeset filter group (update automations only): a combinator plus an array
	// of transition clauses matching what is changing. Dot-paths (nested reference
	// filters) are NOT permitted — direct properties only.
	Changeset param.Field[TriggeredAutomationChangesetParam] `json:"changeset"`
	CreatedAt param.Field[string]                            `json:"created_at"`
	Enabled   param.Field[bool]                              `json:"enabled"`
	ListID    param.Field[string]                            `json:"list_id" format:"uuid"`
	// Lifecycle automations only.
	OnCreate param.Field[bool] `json:"on_create"`
	// Lifecycle automations only.
	OnDelete param.Field[bool] `json:"on_delete"`
	// A filter group: a combinator plus an array of slug-based clauses. Dot-paths
	// (e.g. `organization.location`) express nested reference filters.
	State     param.Field[TriggeredAutomationStateParam] `json:"state"`
	TeamID    param.Field[string]                        `json:"team_id" format:"uuid"`
	UpdatedAt param.Field[string]                        `json:"updated_at"`
	UserID    param.Field[string]                        `json:"user_id"`
}

func (r TriggeredAutomationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An action the automation runs when it fires. `type` selects the kind; the
// remaining fields are type-specific (`agent` → `agent_id`, `webhook` →
// `webhook_id`). Generic: new action types add fields here.
type TriggeredAutomationActionParam struct {
	Type param.Field[TriggeredAutomationActionsType] `json:"type" api:"required"`
	// Required when `type` is `agent`. The agent to run.
	AgentID param.Field[string] `json:"agent_id" format:"uuid"`
	// wait: cron schedule for the resume time. Exactly one of delay_seconds or
	// cron_expression.
	CronExpression param.Field[string] `json:"cron_expression"`
	// wait: relative delay in seconds. Exactly one of delay_seconds or
	// cron_expression.
	DelaySeconds param.Field[int64] `json:"delay_seconds"`
	// wait: IANA timezone for evaluating cron_expression (optional).
	Timezone param.Field[string] `json:"timezone"`
	// Required when `type` is `webhook`. The id of the webhook the event is dispatched
	// to (async) when the automation fires.
	WebhookID   param.Field[string]    `json:"webhook_id" format:"uuid"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r TriggeredAutomationActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A changeset filter group (update automations only): a combinator plus an array
// of transition clauses matching what is changing. Dot-paths (nested reference
// filters) are NOT permitted — direct properties only.
type TriggeredAutomationChangesetParam struct {
	Combinator param.Field[TriggeredAutomationChangesetCombinator] `json:"combinator"`
	// Each entry is a transition clause { slug: { from?: { comparator: value }, to?: {
	// comparator: value } } }. `from` matches the prior value, `to` the new value; an
	// empty body { slug: {} } matches any change to that property.
	Filter param.Field[[]map[string]interface{}] `json:"filter"`
}

func (r TriggeredAutomationChangesetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A filter group: a combinator plus an array of slug-based clauses. Dot-paths
// (e.g. `organization.location`) express nested reference filters.
type TriggeredAutomationStateParam struct {
	Combinator param.Field[TriggeredAutomationStateCombinator] `json:"combinator"`
	// Each entry is { slug: { comparator: value } }
	Filter param.Field[[]map[string]interface{}] `json:"filter"`
}

func (r TriggeredAutomationStateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TriggeredAutomationListResponse struct {
	Data []TriggeredAutomation `json:"data" api:"required"`
	// True if more automations exist beyond this page.
	HasMore bool `json:"has_more" api:"required"`
	// Opaque cursor for the next page; null when has_more is false.
	NextCursor string                              `json:"next_cursor" api:"nullable"`
	JSON       triggeredAutomationListResponseJSON `json:"-"`
}

// triggeredAutomationListResponseJSON contains the JSON metadata for the struct
// [TriggeredAutomationListResponse]
type triggeredAutomationListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TriggeredAutomationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r triggeredAutomationListResponseJSON) RawJSON() string {
	return r.raw
}

type TriggeredAutomationNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// A triggered automation. `kind` selects the shape: `update` fires on object
	// updates and requires a `changeset` (from/to transition) filter plus an optional
	// `state` precondition; `lifecycle` fires on create and/or delete
	// (`on_create`/`on_delete`) and requires a `state` filter (no changeset). `state`
	// permits dot-paths (nested reference filters); `changeset` is direct properties
	// only. Object type is taken from the path.
	TriggeredAutomation TriggeredAutomationParam `json:"triggered_automation" api:"required"`
	IdempotencyKey      param.Field[string]      `header:"Idempotency-Key"`
}

func (r TriggeredAutomationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TriggeredAutomation)
}

// Object types that support triggered automations. Must match the
// triggered-automation whitelist in @micro/database migrate-sql
// (TRIGGERED_AUTOMATION_OBJECTS).
type TriggeredAutomationNewParamsAutomationObjectType string

const (
	TriggeredAutomationNewParamsAutomationObjectTypeMessage         TriggeredAutomationNewParamsAutomationObjectType = "message"
	TriggeredAutomationNewParamsAutomationObjectTypeAction          TriggeredAutomationNewParamsAutomationObjectType = "action"
	TriggeredAutomationNewParamsAutomationObjectTypeEvent           TriggeredAutomationNewParamsAutomationObjectType = "event"
	TriggeredAutomationNewParamsAutomationObjectTypeDocument        TriggeredAutomationNewParamsAutomationObjectType = "document"
	TriggeredAutomationNewParamsAutomationObjectTypeIdentity        TriggeredAutomationNewParamsAutomationObjectType = "identity"
	TriggeredAutomationNewParamsAutomationObjectTypeLinkedinMessage TriggeredAutomationNewParamsAutomationObjectType = "linkedin_message"
	TriggeredAutomationNewParamsAutomationObjectTypeDeal            TriggeredAutomationNewParamsAutomationObjectType = "deal"
	TriggeredAutomationNewParamsAutomationObjectTypeOrganization    TriggeredAutomationNewParamsAutomationObjectType = "organization"
	TriggeredAutomationNewParamsAutomationObjectTypeContact         TriggeredAutomationNewParamsAutomationObjectType = "contact"
)

func (r TriggeredAutomationNewParamsAutomationObjectType) IsKnown() bool {
	switch r {
	case TriggeredAutomationNewParamsAutomationObjectTypeMessage, TriggeredAutomationNewParamsAutomationObjectTypeAction, TriggeredAutomationNewParamsAutomationObjectTypeEvent, TriggeredAutomationNewParamsAutomationObjectTypeDocument, TriggeredAutomationNewParamsAutomationObjectTypeIdentity, TriggeredAutomationNewParamsAutomationObjectTypeLinkedinMessage, TriggeredAutomationNewParamsAutomationObjectTypeDeal, TriggeredAutomationNewParamsAutomationObjectTypeOrganization, TriggeredAutomationNewParamsAutomationObjectTypeContact:
		return true
	}
	return false
}

type TriggeredAutomationUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// A triggered automation. `kind` selects the shape: `update` fires on object
	// updates and requires a `changeset` (from/to transition) filter plus an optional
	// `state` precondition; `lifecycle` fires on create and/or delete
	// (`on_create`/`on_delete`) and requires a `state` filter (no changeset). `state`
	// permits dot-paths (nested reference filters); `changeset` is direct properties
	// only. Object type is taken from the path.
	TriggeredAutomation TriggeredAutomationParam `json:"triggered_automation" api:"required"`
}

func (r TriggeredAutomationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TriggeredAutomation)
}

// Object types that support triggered automations. Must match the
// triggered-automation whitelist in @micro/database migrate-sql
// (TRIGGERED_AUTOMATION_OBJECTS).
type TriggeredAutomationUpdateParamsAutomationObjectType string

const (
	TriggeredAutomationUpdateParamsAutomationObjectTypeMessage         TriggeredAutomationUpdateParamsAutomationObjectType = "message"
	TriggeredAutomationUpdateParamsAutomationObjectTypeAction          TriggeredAutomationUpdateParamsAutomationObjectType = "action"
	TriggeredAutomationUpdateParamsAutomationObjectTypeEvent           TriggeredAutomationUpdateParamsAutomationObjectType = "event"
	TriggeredAutomationUpdateParamsAutomationObjectTypeDocument        TriggeredAutomationUpdateParamsAutomationObjectType = "document"
	TriggeredAutomationUpdateParamsAutomationObjectTypeIdentity        TriggeredAutomationUpdateParamsAutomationObjectType = "identity"
	TriggeredAutomationUpdateParamsAutomationObjectTypeLinkedinMessage TriggeredAutomationUpdateParamsAutomationObjectType = "linkedin_message"
	TriggeredAutomationUpdateParamsAutomationObjectTypeDeal            TriggeredAutomationUpdateParamsAutomationObjectType = "deal"
	TriggeredAutomationUpdateParamsAutomationObjectTypeOrganization    TriggeredAutomationUpdateParamsAutomationObjectType = "organization"
	TriggeredAutomationUpdateParamsAutomationObjectTypeContact         TriggeredAutomationUpdateParamsAutomationObjectType = "contact"
)

func (r TriggeredAutomationUpdateParamsAutomationObjectType) IsKnown() bool {
	switch r {
	case TriggeredAutomationUpdateParamsAutomationObjectTypeMessage, TriggeredAutomationUpdateParamsAutomationObjectTypeAction, TriggeredAutomationUpdateParamsAutomationObjectTypeEvent, TriggeredAutomationUpdateParamsAutomationObjectTypeDocument, TriggeredAutomationUpdateParamsAutomationObjectTypeIdentity, TriggeredAutomationUpdateParamsAutomationObjectTypeLinkedinMessage, TriggeredAutomationUpdateParamsAutomationObjectTypeDeal, TriggeredAutomationUpdateParamsAutomationObjectTypeOrganization, TriggeredAutomationUpdateParamsAutomationObjectTypeContact:
		return true
	}
	return false
}

type TriggeredAutomationListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Opaque pagination cursor (from a prior response's next_cursor); supersedes
	// page/limit when present.
	Cursor param.Field[string] `query:"cursor"`
	// Optional filter to a single automation kind. When omitted, both kinds are
	// returned.
	Kind param.Field[TriggeredAutomationListParamsKind] `query:"kind"`
	// Maximum items per page (<= 50; defaults to 50).
	Limit param.Field[int64] `query:"limit"`
	// List (CRM) id to scope the listing to. When omitted, automations owned by the
	// path team are returned.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
	// 1-based page number. Prefer cursor.
	Page param.Field[int64] `query:"page"`
}

// URLQuery serializes [TriggeredAutomationListParams]'s query parameters as
// `url.Values`.
func (r TriggeredAutomationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Object types that support triggered automations. Must match the
// triggered-automation whitelist in @micro/database migrate-sql
// (TRIGGERED_AUTOMATION_OBJECTS).
type TriggeredAutomationListParamsAutomationObjectType string

const (
	TriggeredAutomationListParamsAutomationObjectTypeMessage         TriggeredAutomationListParamsAutomationObjectType = "message"
	TriggeredAutomationListParamsAutomationObjectTypeAction          TriggeredAutomationListParamsAutomationObjectType = "action"
	TriggeredAutomationListParamsAutomationObjectTypeEvent           TriggeredAutomationListParamsAutomationObjectType = "event"
	TriggeredAutomationListParamsAutomationObjectTypeDocument        TriggeredAutomationListParamsAutomationObjectType = "document"
	TriggeredAutomationListParamsAutomationObjectTypeIdentity        TriggeredAutomationListParamsAutomationObjectType = "identity"
	TriggeredAutomationListParamsAutomationObjectTypeLinkedinMessage TriggeredAutomationListParamsAutomationObjectType = "linkedin_message"
	TriggeredAutomationListParamsAutomationObjectTypeDeal            TriggeredAutomationListParamsAutomationObjectType = "deal"
	TriggeredAutomationListParamsAutomationObjectTypeOrganization    TriggeredAutomationListParamsAutomationObjectType = "organization"
	TriggeredAutomationListParamsAutomationObjectTypeContact         TriggeredAutomationListParamsAutomationObjectType = "contact"
)

func (r TriggeredAutomationListParamsAutomationObjectType) IsKnown() bool {
	switch r {
	case TriggeredAutomationListParamsAutomationObjectTypeMessage, TriggeredAutomationListParamsAutomationObjectTypeAction, TriggeredAutomationListParamsAutomationObjectTypeEvent, TriggeredAutomationListParamsAutomationObjectTypeDocument, TriggeredAutomationListParamsAutomationObjectTypeIdentity, TriggeredAutomationListParamsAutomationObjectTypeLinkedinMessage, TriggeredAutomationListParamsAutomationObjectTypeDeal, TriggeredAutomationListParamsAutomationObjectTypeOrganization, TriggeredAutomationListParamsAutomationObjectTypeContact:
		return true
	}
	return false
}

// Optional filter to a single automation kind. When omitted, both kinds are
// returned.
type TriggeredAutomationListParamsKind string

const (
	TriggeredAutomationListParamsKindUpdate    TriggeredAutomationListParamsKind = "update"
	TriggeredAutomationListParamsKindLifecycle TriggeredAutomationListParamsKind = "lifecycle"
)

func (r TriggeredAutomationListParamsKind) IsKnown() bool {
	switch r {
	case TriggeredAutomationListParamsKindUpdate, TriggeredAutomationListParamsKindLifecycle:
		return true
	}
	return false
}

type TriggeredAutomationDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

// Object types that support triggered automations. Must match the
// triggered-automation whitelist in @micro/database migrate-sql
// (TRIGGERED_AUTOMATION_OBJECTS).
type TriggeredAutomationDeleteParamsAutomationObjectType string

const (
	TriggeredAutomationDeleteParamsAutomationObjectTypeMessage         TriggeredAutomationDeleteParamsAutomationObjectType = "message"
	TriggeredAutomationDeleteParamsAutomationObjectTypeAction          TriggeredAutomationDeleteParamsAutomationObjectType = "action"
	TriggeredAutomationDeleteParamsAutomationObjectTypeEvent           TriggeredAutomationDeleteParamsAutomationObjectType = "event"
	TriggeredAutomationDeleteParamsAutomationObjectTypeDocument        TriggeredAutomationDeleteParamsAutomationObjectType = "document"
	TriggeredAutomationDeleteParamsAutomationObjectTypeIdentity        TriggeredAutomationDeleteParamsAutomationObjectType = "identity"
	TriggeredAutomationDeleteParamsAutomationObjectTypeLinkedinMessage TriggeredAutomationDeleteParamsAutomationObjectType = "linkedin_message"
	TriggeredAutomationDeleteParamsAutomationObjectTypeDeal            TriggeredAutomationDeleteParamsAutomationObjectType = "deal"
	TriggeredAutomationDeleteParamsAutomationObjectTypeOrganization    TriggeredAutomationDeleteParamsAutomationObjectType = "organization"
	TriggeredAutomationDeleteParamsAutomationObjectTypeContact         TriggeredAutomationDeleteParamsAutomationObjectType = "contact"
)

func (r TriggeredAutomationDeleteParamsAutomationObjectType) IsKnown() bool {
	switch r {
	case TriggeredAutomationDeleteParamsAutomationObjectTypeMessage, TriggeredAutomationDeleteParamsAutomationObjectTypeAction, TriggeredAutomationDeleteParamsAutomationObjectTypeEvent, TriggeredAutomationDeleteParamsAutomationObjectTypeDocument, TriggeredAutomationDeleteParamsAutomationObjectTypeIdentity, TriggeredAutomationDeleteParamsAutomationObjectTypeLinkedinMessage, TriggeredAutomationDeleteParamsAutomationObjectTypeDeal, TriggeredAutomationDeleteParamsAutomationObjectTypeOrganization, TriggeredAutomationDeleteParamsAutomationObjectTypeContact:
		return true
	}
	return false
}

type TriggeredAutomationGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

// Object types that support triggered automations. Must match the
// triggered-automation whitelist in @micro/database migrate-sql
// (TRIGGERED_AUTOMATION_OBJECTS).
type TriggeredAutomationGetParamsAutomationObjectType string

const (
	TriggeredAutomationGetParamsAutomationObjectTypeMessage         TriggeredAutomationGetParamsAutomationObjectType = "message"
	TriggeredAutomationGetParamsAutomationObjectTypeAction          TriggeredAutomationGetParamsAutomationObjectType = "action"
	TriggeredAutomationGetParamsAutomationObjectTypeEvent           TriggeredAutomationGetParamsAutomationObjectType = "event"
	TriggeredAutomationGetParamsAutomationObjectTypeDocument        TriggeredAutomationGetParamsAutomationObjectType = "document"
	TriggeredAutomationGetParamsAutomationObjectTypeIdentity        TriggeredAutomationGetParamsAutomationObjectType = "identity"
	TriggeredAutomationGetParamsAutomationObjectTypeLinkedinMessage TriggeredAutomationGetParamsAutomationObjectType = "linkedin_message"
	TriggeredAutomationGetParamsAutomationObjectTypeDeal            TriggeredAutomationGetParamsAutomationObjectType = "deal"
	TriggeredAutomationGetParamsAutomationObjectTypeOrganization    TriggeredAutomationGetParamsAutomationObjectType = "organization"
	TriggeredAutomationGetParamsAutomationObjectTypeContact         TriggeredAutomationGetParamsAutomationObjectType = "contact"
)

func (r TriggeredAutomationGetParamsAutomationObjectType) IsKnown() bool {
	switch r {
	case TriggeredAutomationGetParamsAutomationObjectTypeMessage, TriggeredAutomationGetParamsAutomationObjectTypeAction, TriggeredAutomationGetParamsAutomationObjectTypeEvent, TriggeredAutomationGetParamsAutomationObjectTypeDocument, TriggeredAutomationGetParamsAutomationObjectTypeIdentity, TriggeredAutomationGetParamsAutomationObjectTypeLinkedinMessage, TriggeredAutomationGetParamsAutomationObjectTypeDeal, TriggeredAutomationGetParamsAutomationObjectTypeOrganization, TriggeredAutomationGetParamsAutomationObjectTypeContact:
		return true
	}
	return false
}
