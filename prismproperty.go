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

// PrismPropertyService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismPropertyService] method instead.
type PrismPropertyService struct {
	Options []option.RequestOption
	Options *PrismPropertyOptionService
}

// NewPrismPropertyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrismPropertyService(opts ...option.RequestOption) (r *PrismPropertyService) {
	r = &PrismPropertyService{}
	r.Options = opts
	r.Options = NewPrismPropertyOptionService(opts...)
	return
}

// Define a new property on this object type, scoped to the calling team. Search
// the existing properties first (GET this path with `term`) and reuse a match
// instead of defining a second property for the same fact. Pass `list_id` in the
// body to scope the definition to one list/app; without it the property is
// workspace-global and appears on every list. A name already used in that scope,
// an explicitly requested slug already taken, or a slug that a shared property
// already owns all return 409 naming the definition to use instead. The property's
// display format is resolved from `type` automatically — pass `role_id` only to
// override it. For `select_str` and `multiselect_str` types you may pre-seed the
// choices via `options`.
func (r *PrismPropertyService) New(ctx context.Context, objectType PrismPropertyNewParamsObjectType, params PrismPropertyNewParams, opts ...option.RequestOption) (res *PropertyDefinition, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/%v/properties", params.TeamID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patches the editable fields (`name`, `icon`, `enabled`) of a property
// definition. `type` and scoping fields are immutable; `type` must be supplied in
// the body so the server knows which per-type table to write.
func (r *PrismPropertyService) Update(ctx context.Context, objectType PrismPropertyUpdateParamsObjectType, propertyID string, params PrismPropertyUpdateParams, opts ...option.RequestOption) (res *PropertyDefinition, err error) {
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
	if propertyID == "" {
		err = errors.New("missing required propertyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/properties/%s", params.TeamID, objectType, propertyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Get metadata properties by object type
func (r *PrismPropertyService) List(ctx context.Context, objectType PrismPropertyListParamsObjectType, params PrismPropertyListParams, opts ...option.RequestOption) (res *PrismPropertyListResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/%v/properties", params.TeamID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Removes the property definition and any of its options. Fails with 409
// `property_in_use` if records still reference the property.
func (r *PrismPropertyService) Delete(ctx context.Context, objectType PrismPropertyDeleteParamsObjectType, propertyID string, params PrismPropertyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return err
	}
	requestconfig.UseDefaultParam(&params.TeamID, precfg.TeamID)
	if params.TeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return err
	}
	if propertyID == "" {
		err = errors.New("missing required propertyId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/properties/%s", params.TeamID, objectType, propertyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Lists property definitions across every object type the engine knows about,
// including pipeline-owned types that are not queryable or CRUD-capable
// (`message`, `thread`, `linkedin_thread`, and others). Only the `ObjectType` enum
// (12 types) can be queried, created, updated, or listed. Contacts point at
// `message` via `last_email`; that relationship cannot be followed with `/query`.
func (r *PrismPropertyService) ListAll(ctx context.Context, params PrismPropertyListAllParams, opts ...option.RequestOption) (res *PrismPropertyListAllResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/properties", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Definition for a single property on an object type. Definitions with team_id and
// crm_id null are shared defaults; values may be scoped to a team and/or list
// (crm).
type PropertyDefinition struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Slug string `json:"slug" api:"required"`
	// Storage type for a property definition. Determines which per-type table holds
	// the values, and which display formats the property can take.
	Type PropertyDefinitionType `json:"type" api:"required"`
	// Reserved alias naming this definition, or null. `app_stage` marks the list
	// pipeline stage definition. Resolve stages by this field rather than by name,
	// slug, or team_id: a superseded native `status` definition can coexist with the
	// pipeline one and is otherwise identical on the wire.
	Alias PropertyDefinitionAlias `json:"alias" api:"nullable"`
	// Identifier of the list this definition is scoped to, when applicable.
	CRMID string `json:"crm_id" api:"nullable" format:"uuid"`
	// Canonical identifier of the list this definition is scoped to.
	ListID string `json:"list_id" api:"nullable" format:"uuid"`
	Locked bool   `json:"locked"`
	Name   string `json:"name" api:"nullable"`
	Native bool   `json:"native"`
	// Present only for select_str and multiselect_str types.
	Options []PropertyOption `json:"options"`
	// When true, records of this object type must carry a non-empty value for this
	// property on create, and a patch may not clear it.
	Required bool `json:"required"`
	// The property's display format. Always populated on definitions created through
	// this API; a null here means the definition predates that and will render as an
	// unknown format until it is patched.
	RoleID string                 `json:"role_id" api:"nullable" format:"uuid"`
	TeamID string                 `json:"team_id" api:"nullable" format:"uuid"`
	JSON   propertyDefinitionJSON `json:"-"`
}

// propertyDefinitionJSON contains the JSON metadata for the struct
// [PropertyDefinition]
type propertyDefinitionJSON struct {
	ID          apijson.Field
	Slug        apijson.Field
	Type        apijson.Field
	Alias       apijson.Field
	CRMID       apijson.Field
	ListID      apijson.Field
	Locked      apijson.Field
	Name        apijson.Field
	Native      apijson.Field
	Options     apijson.Field
	Required    apijson.Field
	RoleID      apijson.Field
	TeamID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PropertyDefinition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r propertyDefinitionJSON) RawJSON() string {
	return r.raw
}

// Storage type for a property definition. Determines which per-type table holds
// the values, and which display formats the property can take.
type PropertyDefinitionType string

const (
	PropertyDefinitionTypeNum                   PropertyDefinitionType = "num"
	PropertyDefinitionTypeStr                   PropertyDefinitionType = "str"
	PropertyDefinitionTypeBool                  PropertyDefinitionType = "bool"
	PropertyDefinitionTypeDate                  PropertyDefinitionType = "date"
	PropertyDefinitionTypeText                  PropertyDefinitionType = "text"
	PropertyDefinitionTypeByte                  PropertyDefinitionType = "byte"
	PropertyDefinitionTypeSelectStr             PropertyDefinitionType = "select_str"
	PropertyDefinitionTypeMultiStr              PropertyDefinitionType = "multi_str"
	PropertyDefinitionTypeMultiselectStr        PropertyDefinitionType = "multiselect_str"
	PropertyDefinitionTypeJsonb                 PropertyDefinitionType = "jsonb"
	PropertyDefinitionTypeRefIdentity           PropertyDefinitionType = "ref_identity"
	PropertyDefinitionTypeRefUser               PropertyDefinitionType = "ref_user"
	PropertyDefinitionTypeRefOrganization       PropertyDefinitionType = "ref_organization"
	PropertyDefinitionTypeRefContact            PropertyDefinitionType = "ref_contact"
	PropertyDefinitionTypeRefThread             PropertyDefinitionType = "ref_thread"
	PropertyDefinitionTypeRefMessage            PropertyDefinitionType = "ref_message"
	PropertyDefinitionTypeRefEvent              PropertyDefinitionType = "ref_event"
	PropertyDefinitionTypeRefAccount            PropertyDefinitionType = "ref_account"
	PropertyDefinitionTypeRefAIChatThread       PropertyDefinitionType = "ref_ai_chat_thread"
	PropertyDefinitionTypeRefAIChatMessage      PropertyDefinitionType = "ref_ai_chat_message"
	PropertyDefinitionTypeMultirefAIChatMessage PropertyDefinitionType = "multiref_ai_chat_message"
	PropertyDefinitionTypeMultirefAgentSite     PropertyDefinitionType = "multiref_agent_site"
	PropertyDefinitionTypeMultirefAction        PropertyDefinitionType = "multiref_action"
	PropertyDefinitionTypeMultirefComment       PropertyDefinitionType = "multiref_comment"
	PropertyDefinitionTypeMultirefContact       PropertyDefinitionType = "multiref_contact"
	PropertyDefinitionTypeMultirefLabel         PropertyDefinitionType = "multiref_label"
	PropertyDefinitionTypeMultirefThread        PropertyDefinitionType = "multiref_thread"
	PropertyDefinitionTypeMultirefMessages      PropertyDefinitionType = "multiref_messages"
	PropertyDefinitionTypeMultirefDocument      PropertyDefinitionType = "multiref_document"
	PropertyDefinitionTypeMultirefIdentity      PropertyDefinitionType = "multiref_identity"
	PropertyDefinitionTypeMultirefOrganization  PropertyDefinitionType = "multiref_organization"
	PropertyDefinitionTypeMultirefEngagement    PropertyDefinitionType = "multiref_engagement"
	PropertyDefinitionTypeMultirefAttendee      PropertyDefinitionType = "multiref_attendee"
	PropertyDefinitionTypeMultirefMeetingEntry  PropertyDefinitionType = "multiref_meeting_entry"
	PropertyDefinitionTypeMultirefReadReceipt   PropertyDefinitionType = "multiref_read_receipt"
	PropertyDefinitionTypeMultirefAccount       PropertyDefinitionType = "multiref_account"
	PropertyDefinitionTypeMultirefSource        PropertyDefinitionType = "multiref_source"
)

func (r PropertyDefinitionType) IsKnown() bool {
	switch r {
	case PropertyDefinitionTypeNum, PropertyDefinitionTypeStr, PropertyDefinitionTypeBool, PropertyDefinitionTypeDate, PropertyDefinitionTypeText, PropertyDefinitionTypeByte, PropertyDefinitionTypeSelectStr, PropertyDefinitionTypeMultiStr, PropertyDefinitionTypeMultiselectStr, PropertyDefinitionTypeJsonb, PropertyDefinitionTypeRefIdentity, PropertyDefinitionTypeRefUser, PropertyDefinitionTypeRefOrganization, PropertyDefinitionTypeRefContact, PropertyDefinitionTypeRefThread, PropertyDefinitionTypeRefMessage, PropertyDefinitionTypeRefEvent, PropertyDefinitionTypeRefAccount, PropertyDefinitionTypeRefAIChatThread, PropertyDefinitionTypeRefAIChatMessage, PropertyDefinitionTypeMultirefAIChatMessage, PropertyDefinitionTypeMultirefAgentSite, PropertyDefinitionTypeMultirefAction, PropertyDefinitionTypeMultirefComment, PropertyDefinitionTypeMultirefContact, PropertyDefinitionTypeMultirefLabel, PropertyDefinitionTypeMultirefThread, PropertyDefinitionTypeMultirefMessages, PropertyDefinitionTypeMultirefDocument, PropertyDefinitionTypeMultirefIdentity, PropertyDefinitionTypeMultirefOrganization, PropertyDefinitionTypeMultirefEngagement, PropertyDefinitionTypeMultirefAttendee, PropertyDefinitionTypeMultirefMeetingEntry, PropertyDefinitionTypeMultirefReadReceipt, PropertyDefinitionTypeMultirefAccount, PropertyDefinitionTypeMultirefSource:
		return true
	}
	return false
}

// Reserved alias naming this definition, or null. `app_stage` marks the list
// pipeline stage definition. Resolve stages by this field rather than by name,
// slug, or team_id: a superseded native `status` definition can coexist with the
// pipeline one and is otherwise identical on the wire.
type PropertyDefinitionAlias string

const (
	PropertyDefinitionAliasAppStage PropertyDefinitionAlias = "app_stage"
)

func (r PropertyDefinitionAlias) IsKnown() bool {
	switch r {
	case PropertyDefinitionAliasAppStage:
		return true
	}
	return false
}

// New property definition. Check for an existing property first (GET the same path
// with `term`) and reuse it rather than defining a near-duplicate — writes address
// properties by slug, so two definitions sharing a slug leave no addressable
// winner. For `select_str`/`multiselect_str` types you may pre-seed choices via
// `options`.
type PropertyDefinitionCreateParam struct {
	// Human-readable name, unique within the scope the definition is created in. A
	// name already taken in that scope returns 409; the message names the existing
	// definition's id, slug and type so you can write to it instead.
	Name param.Field[string] `json:"name" api:"required"`
	// Storage type for a property definition. Determines which per-type table holds
	// the values, and which display formats the property can take.
	Type param.Field[PropertyDefinitionCreateType] `json:"type" api:"required"`
	Icon param.Field[string]                       `json:"icon"`
	// Scopes the definition to one list/app. Omit it only for a property that
	// genuinely belongs to the whole workspace: a definition created without `list_id`
	// is workspace-global and surfaces on every list of this object type.
	ListID param.Field[string] `json:"list_id" format:"uuid"`
	// Only honored when `type` is `select_str` or `multiselect_str`.
	Options param.Field[[]PropertyDefinitionCreateOptionParam] `json:"options"`
	// When true, records must carry a non-empty value for this property on create.
	// Defaults to false.
	Required param.Field[bool] `json:"required"`
	// Optional display format for the property, drawn from the workspace's property
	// roles. Omit it and the canonical role for `type` is applied (plain text, plain
	// number, checkbox). Supply it only to pick a narrower format such as email, URL
	// or currency; the role's data type must match `type`.
	RoleID param.Field[string] `json:"role_id" format:"uuid"`
	// URL-safe identifier. When omitted it defaults to a slugified `name` and is
	// disambiguated with a numeric suffix on conflict. When supplied explicitly it is
	// treated as part of your write contract and is never silently renamed — a
	// collision returns 409 instead.
	Slug param.Field[string] `json:"slug"`
}

func (r PropertyDefinitionCreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Storage type for a property definition. Determines which per-type table holds
// the values, and which display formats the property can take.
type PropertyDefinitionCreateType string

const (
	PropertyDefinitionCreateTypeNum                   PropertyDefinitionCreateType = "num"
	PropertyDefinitionCreateTypeStr                   PropertyDefinitionCreateType = "str"
	PropertyDefinitionCreateTypeBool                  PropertyDefinitionCreateType = "bool"
	PropertyDefinitionCreateTypeDate                  PropertyDefinitionCreateType = "date"
	PropertyDefinitionCreateTypeText                  PropertyDefinitionCreateType = "text"
	PropertyDefinitionCreateTypeByte                  PropertyDefinitionCreateType = "byte"
	PropertyDefinitionCreateTypeSelectStr             PropertyDefinitionCreateType = "select_str"
	PropertyDefinitionCreateTypeMultiStr              PropertyDefinitionCreateType = "multi_str"
	PropertyDefinitionCreateTypeMultiselectStr        PropertyDefinitionCreateType = "multiselect_str"
	PropertyDefinitionCreateTypeJsonb                 PropertyDefinitionCreateType = "jsonb"
	PropertyDefinitionCreateTypeRefIdentity           PropertyDefinitionCreateType = "ref_identity"
	PropertyDefinitionCreateTypeRefUser               PropertyDefinitionCreateType = "ref_user"
	PropertyDefinitionCreateTypeRefOrganization       PropertyDefinitionCreateType = "ref_organization"
	PropertyDefinitionCreateTypeRefContact            PropertyDefinitionCreateType = "ref_contact"
	PropertyDefinitionCreateTypeRefThread             PropertyDefinitionCreateType = "ref_thread"
	PropertyDefinitionCreateTypeRefMessage            PropertyDefinitionCreateType = "ref_message"
	PropertyDefinitionCreateTypeRefEvent              PropertyDefinitionCreateType = "ref_event"
	PropertyDefinitionCreateTypeRefAccount            PropertyDefinitionCreateType = "ref_account"
	PropertyDefinitionCreateTypeRefAIChatThread       PropertyDefinitionCreateType = "ref_ai_chat_thread"
	PropertyDefinitionCreateTypeRefAIChatMessage      PropertyDefinitionCreateType = "ref_ai_chat_message"
	PropertyDefinitionCreateTypeMultirefAIChatMessage PropertyDefinitionCreateType = "multiref_ai_chat_message"
	PropertyDefinitionCreateTypeMultirefAgentSite     PropertyDefinitionCreateType = "multiref_agent_site"
	PropertyDefinitionCreateTypeMultirefAction        PropertyDefinitionCreateType = "multiref_action"
	PropertyDefinitionCreateTypeMultirefComment       PropertyDefinitionCreateType = "multiref_comment"
	PropertyDefinitionCreateTypeMultirefContact       PropertyDefinitionCreateType = "multiref_contact"
	PropertyDefinitionCreateTypeMultirefLabel         PropertyDefinitionCreateType = "multiref_label"
	PropertyDefinitionCreateTypeMultirefThread        PropertyDefinitionCreateType = "multiref_thread"
	PropertyDefinitionCreateTypeMultirefMessages      PropertyDefinitionCreateType = "multiref_messages"
	PropertyDefinitionCreateTypeMultirefDocument      PropertyDefinitionCreateType = "multiref_document"
	PropertyDefinitionCreateTypeMultirefIdentity      PropertyDefinitionCreateType = "multiref_identity"
	PropertyDefinitionCreateTypeMultirefOrganization  PropertyDefinitionCreateType = "multiref_organization"
	PropertyDefinitionCreateTypeMultirefEngagement    PropertyDefinitionCreateType = "multiref_engagement"
	PropertyDefinitionCreateTypeMultirefAttendee      PropertyDefinitionCreateType = "multiref_attendee"
	PropertyDefinitionCreateTypeMultirefMeetingEntry  PropertyDefinitionCreateType = "multiref_meeting_entry"
	PropertyDefinitionCreateTypeMultirefReadReceipt   PropertyDefinitionCreateType = "multiref_read_receipt"
	PropertyDefinitionCreateTypeMultirefAccount       PropertyDefinitionCreateType = "multiref_account"
	PropertyDefinitionCreateTypeMultirefSource        PropertyDefinitionCreateType = "multiref_source"
)

func (r PropertyDefinitionCreateType) IsKnown() bool {
	switch r {
	case PropertyDefinitionCreateTypeNum, PropertyDefinitionCreateTypeStr, PropertyDefinitionCreateTypeBool, PropertyDefinitionCreateTypeDate, PropertyDefinitionCreateTypeText, PropertyDefinitionCreateTypeByte, PropertyDefinitionCreateTypeSelectStr, PropertyDefinitionCreateTypeMultiStr, PropertyDefinitionCreateTypeMultiselectStr, PropertyDefinitionCreateTypeJsonb, PropertyDefinitionCreateTypeRefIdentity, PropertyDefinitionCreateTypeRefUser, PropertyDefinitionCreateTypeRefOrganization, PropertyDefinitionCreateTypeRefContact, PropertyDefinitionCreateTypeRefThread, PropertyDefinitionCreateTypeRefMessage, PropertyDefinitionCreateTypeRefEvent, PropertyDefinitionCreateTypeRefAccount, PropertyDefinitionCreateTypeRefAIChatThread, PropertyDefinitionCreateTypeRefAIChatMessage, PropertyDefinitionCreateTypeMultirefAIChatMessage, PropertyDefinitionCreateTypeMultirefAgentSite, PropertyDefinitionCreateTypeMultirefAction, PropertyDefinitionCreateTypeMultirefComment, PropertyDefinitionCreateTypeMultirefContact, PropertyDefinitionCreateTypeMultirefLabel, PropertyDefinitionCreateTypeMultirefThread, PropertyDefinitionCreateTypeMultirefMessages, PropertyDefinitionCreateTypeMultirefDocument, PropertyDefinitionCreateTypeMultirefIdentity, PropertyDefinitionCreateTypeMultirefOrganization, PropertyDefinitionCreateTypeMultirefEngagement, PropertyDefinitionCreateTypeMultirefAttendee, PropertyDefinitionCreateTypeMultirefMeetingEntry, PropertyDefinitionCreateTypeMultirefReadReceipt, PropertyDefinitionCreateTypeMultirefAccount, PropertyDefinitionCreateTypeMultirefSource:
		return true
	}
	return false
}

type PropertyDefinitionCreateOptionParam struct {
	Value       param.Field[string] `json:"value" api:"required"`
	ColorScheme param.Field[string] `json:"color_scheme"`
	Description param.Field[string] `json:"description"`
	Icon        param.Field[string] `json:"icon"`
	OptionGroup param.Field[string] `json:"option_group"`
	Slug        param.Field[string] `json:"slug"`
	SortIndex   param.Field[int64]  `json:"sort_index"`
}

func (r PropertyDefinitionCreateOptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Partial update of a property definition. Only `name`, `icon`, `enabled`, and
// `required` are editable. `type` identifies the per-type table to write.
type PropertyDefinitionPatchParam struct {
	// Storage type for a property definition. Determines which per-type table holds
	// the values, and which display formats the property can take.
	Type     param.Field[PropertyDefinitionPatchType] `json:"type" api:"required"`
	Enabled  param.Field[bool]                        `json:"enabled"`
	Icon     param.Field[string]                      `json:"icon"`
	ListID   param.Field[string]                      `json:"list_id" format:"uuid"`
	Name     param.Field[string]                      `json:"name"`
	Required param.Field[bool]                        `json:"required"`
}

func (r PropertyDefinitionPatchParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Storage type for a property definition. Determines which per-type table holds
// the values, and which display formats the property can take.
type PropertyDefinitionPatchType string

const (
	PropertyDefinitionPatchTypeNum                   PropertyDefinitionPatchType = "num"
	PropertyDefinitionPatchTypeStr                   PropertyDefinitionPatchType = "str"
	PropertyDefinitionPatchTypeBool                  PropertyDefinitionPatchType = "bool"
	PropertyDefinitionPatchTypeDate                  PropertyDefinitionPatchType = "date"
	PropertyDefinitionPatchTypeText                  PropertyDefinitionPatchType = "text"
	PropertyDefinitionPatchTypeByte                  PropertyDefinitionPatchType = "byte"
	PropertyDefinitionPatchTypeSelectStr             PropertyDefinitionPatchType = "select_str"
	PropertyDefinitionPatchTypeMultiStr              PropertyDefinitionPatchType = "multi_str"
	PropertyDefinitionPatchTypeMultiselectStr        PropertyDefinitionPatchType = "multiselect_str"
	PropertyDefinitionPatchTypeJsonb                 PropertyDefinitionPatchType = "jsonb"
	PropertyDefinitionPatchTypeRefIdentity           PropertyDefinitionPatchType = "ref_identity"
	PropertyDefinitionPatchTypeRefUser               PropertyDefinitionPatchType = "ref_user"
	PropertyDefinitionPatchTypeRefOrganization       PropertyDefinitionPatchType = "ref_organization"
	PropertyDefinitionPatchTypeRefContact            PropertyDefinitionPatchType = "ref_contact"
	PropertyDefinitionPatchTypeRefThread             PropertyDefinitionPatchType = "ref_thread"
	PropertyDefinitionPatchTypeRefMessage            PropertyDefinitionPatchType = "ref_message"
	PropertyDefinitionPatchTypeRefEvent              PropertyDefinitionPatchType = "ref_event"
	PropertyDefinitionPatchTypeRefAccount            PropertyDefinitionPatchType = "ref_account"
	PropertyDefinitionPatchTypeRefAIChatThread       PropertyDefinitionPatchType = "ref_ai_chat_thread"
	PropertyDefinitionPatchTypeRefAIChatMessage      PropertyDefinitionPatchType = "ref_ai_chat_message"
	PropertyDefinitionPatchTypeMultirefAIChatMessage PropertyDefinitionPatchType = "multiref_ai_chat_message"
	PropertyDefinitionPatchTypeMultirefAgentSite     PropertyDefinitionPatchType = "multiref_agent_site"
	PropertyDefinitionPatchTypeMultirefAction        PropertyDefinitionPatchType = "multiref_action"
	PropertyDefinitionPatchTypeMultirefComment       PropertyDefinitionPatchType = "multiref_comment"
	PropertyDefinitionPatchTypeMultirefContact       PropertyDefinitionPatchType = "multiref_contact"
	PropertyDefinitionPatchTypeMultirefLabel         PropertyDefinitionPatchType = "multiref_label"
	PropertyDefinitionPatchTypeMultirefThread        PropertyDefinitionPatchType = "multiref_thread"
	PropertyDefinitionPatchTypeMultirefMessages      PropertyDefinitionPatchType = "multiref_messages"
	PropertyDefinitionPatchTypeMultirefDocument      PropertyDefinitionPatchType = "multiref_document"
	PropertyDefinitionPatchTypeMultirefIdentity      PropertyDefinitionPatchType = "multiref_identity"
	PropertyDefinitionPatchTypeMultirefOrganization  PropertyDefinitionPatchType = "multiref_organization"
	PropertyDefinitionPatchTypeMultirefEngagement    PropertyDefinitionPatchType = "multiref_engagement"
	PropertyDefinitionPatchTypeMultirefAttendee      PropertyDefinitionPatchType = "multiref_attendee"
	PropertyDefinitionPatchTypeMultirefMeetingEntry  PropertyDefinitionPatchType = "multiref_meeting_entry"
	PropertyDefinitionPatchTypeMultirefReadReceipt   PropertyDefinitionPatchType = "multiref_read_receipt"
	PropertyDefinitionPatchTypeMultirefAccount       PropertyDefinitionPatchType = "multiref_account"
	PropertyDefinitionPatchTypeMultirefSource        PropertyDefinitionPatchType = "multiref_source"
)

func (r PropertyDefinitionPatchType) IsKnown() bool {
	switch r {
	case PropertyDefinitionPatchTypeNum, PropertyDefinitionPatchTypeStr, PropertyDefinitionPatchTypeBool, PropertyDefinitionPatchTypeDate, PropertyDefinitionPatchTypeText, PropertyDefinitionPatchTypeByte, PropertyDefinitionPatchTypeSelectStr, PropertyDefinitionPatchTypeMultiStr, PropertyDefinitionPatchTypeMultiselectStr, PropertyDefinitionPatchTypeJsonb, PropertyDefinitionPatchTypeRefIdentity, PropertyDefinitionPatchTypeRefUser, PropertyDefinitionPatchTypeRefOrganization, PropertyDefinitionPatchTypeRefContact, PropertyDefinitionPatchTypeRefThread, PropertyDefinitionPatchTypeRefMessage, PropertyDefinitionPatchTypeRefEvent, PropertyDefinitionPatchTypeRefAccount, PropertyDefinitionPatchTypeRefAIChatThread, PropertyDefinitionPatchTypeRefAIChatMessage, PropertyDefinitionPatchTypeMultirefAIChatMessage, PropertyDefinitionPatchTypeMultirefAgentSite, PropertyDefinitionPatchTypeMultirefAction, PropertyDefinitionPatchTypeMultirefComment, PropertyDefinitionPatchTypeMultirefContact, PropertyDefinitionPatchTypeMultirefLabel, PropertyDefinitionPatchTypeMultirefThread, PropertyDefinitionPatchTypeMultirefMessages, PropertyDefinitionPatchTypeMultirefDocument, PropertyDefinitionPatchTypeMultirefIdentity, PropertyDefinitionPatchTypeMultirefOrganization, PropertyDefinitionPatchTypeMultirefEngagement, PropertyDefinitionPatchTypeMultirefAttendee, PropertyDefinitionPatchTypeMultirefMeetingEntry, PropertyDefinitionPatchTypeMultirefReadReceipt, PropertyDefinitionPatchTypeMultirefAccount, PropertyDefinitionPatchTypeMultirefSource:
		return true
	}
	return false
}

type PrismPropertyListResponse map[string]interface{}

type PrismPropertyListAllResponse map[string]interface{}

type PrismPropertyNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// New property definition. Check for an existing property first (GET the same path
	// with `term`) and reuse it rather than defining a near-duplicate — writes address
	// properties by slug, so two definitions sharing a slug leave no addressable
	// winner. For `select_str`/`multiselect_str` types you may pre-seed choices via
	// `options`.
	PropertyDefinitionCreate PropertyDefinitionCreateParam `json:"property_definition_create" api:"required"`
	IdempotencyKey           param.Field[string]           `header:"Idempotency-Key"`
}

func (r PrismPropertyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PropertyDefinitionCreate)
}

// Object types that support CRUD, query, list, and per-type property metadata.
// `GET /v2/prism/{teamId}/properties` (list-all) also returns definitions for
// pipeline-owned types that are not in this set — including `message`, `thread`,
// and `linkedin_thread`. Those types are not queryable. Contacts expose
// `last_email` as a `ref_message`; you cannot query `message` to follow it.
type PrismPropertyNewParamsObjectType string

const (
	PrismPropertyNewParamsObjectTypeComment       PrismPropertyNewParamsObjectType = "comment"
	PrismPropertyNewParamsObjectTypeDeal          PrismPropertyNewParamsObjectType = "deal"
	PrismPropertyNewParamsObjectTypeEngagement    PrismPropertyNewParamsObjectType = "engagement"
	PrismPropertyNewParamsObjectTypeIdentity      PrismPropertyNewParamsObjectType = "identity"
	PrismPropertyNewParamsObjectTypeAIChatThread  PrismPropertyNewParamsObjectType = "ai_chat_thread"
	PrismPropertyNewParamsObjectTypeAIChatMessage PrismPropertyNewParamsObjectType = "ai_chat_message"
	PrismPropertyNewParamsObjectTypeAgentSite     PrismPropertyNewParamsObjectType = "agent_site"
	PrismPropertyNewParamsObjectTypeDocument      PrismPropertyNewParamsObjectType = "document"
	PrismPropertyNewParamsObjectTypeAction        PrismPropertyNewParamsObjectType = "action"
	PrismPropertyNewParamsObjectTypeEvent         PrismPropertyNewParamsObjectType = "event"
	PrismPropertyNewParamsObjectTypeOrganization  PrismPropertyNewParamsObjectType = "organization"
	PrismPropertyNewParamsObjectTypeContact       PrismPropertyNewParamsObjectType = "contact"
)

func (r PrismPropertyNewParamsObjectType) IsKnown() bool {
	switch r {
	case PrismPropertyNewParamsObjectTypeComment, PrismPropertyNewParamsObjectTypeDeal, PrismPropertyNewParamsObjectTypeEngagement, PrismPropertyNewParamsObjectTypeIdentity, PrismPropertyNewParamsObjectTypeAIChatThread, PrismPropertyNewParamsObjectTypeAIChatMessage, PrismPropertyNewParamsObjectTypeAgentSite, PrismPropertyNewParamsObjectTypeDocument, PrismPropertyNewParamsObjectTypeAction, PrismPropertyNewParamsObjectTypeEvent, PrismPropertyNewParamsObjectTypeOrganization, PrismPropertyNewParamsObjectTypeContact:
		return true
	}
	return false
}

type PrismPropertyUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Partial update of a property definition. Only `name`, `icon`, `enabled`, and
	// `required` are editable. `type` identifies the per-type table to write.
	PropertyDefinitionPatch PropertyDefinitionPatchParam `json:"property_definition_patch" api:"required"`
	IdempotencyKey          param.Field[string]          `header:"Idempotency-Key"`
}

func (r PrismPropertyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PropertyDefinitionPatch)
}

// Object types that support CRUD, query, list, and per-type property metadata.
// `GET /v2/prism/{teamId}/properties` (list-all) also returns definitions for
// pipeline-owned types that are not in this set — including `message`, `thread`,
// and `linkedin_thread`. Those types are not queryable. Contacts expose
// `last_email` as a `ref_message`; you cannot query `message` to follow it.
type PrismPropertyUpdateParamsObjectType string

const (
	PrismPropertyUpdateParamsObjectTypeComment       PrismPropertyUpdateParamsObjectType = "comment"
	PrismPropertyUpdateParamsObjectTypeDeal          PrismPropertyUpdateParamsObjectType = "deal"
	PrismPropertyUpdateParamsObjectTypeEngagement    PrismPropertyUpdateParamsObjectType = "engagement"
	PrismPropertyUpdateParamsObjectTypeIdentity      PrismPropertyUpdateParamsObjectType = "identity"
	PrismPropertyUpdateParamsObjectTypeAIChatThread  PrismPropertyUpdateParamsObjectType = "ai_chat_thread"
	PrismPropertyUpdateParamsObjectTypeAIChatMessage PrismPropertyUpdateParamsObjectType = "ai_chat_message"
	PrismPropertyUpdateParamsObjectTypeAgentSite     PrismPropertyUpdateParamsObjectType = "agent_site"
	PrismPropertyUpdateParamsObjectTypeDocument      PrismPropertyUpdateParamsObjectType = "document"
	PrismPropertyUpdateParamsObjectTypeAction        PrismPropertyUpdateParamsObjectType = "action"
	PrismPropertyUpdateParamsObjectTypeEvent         PrismPropertyUpdateParamsObjectType = "event"
	PrismPropertyUpdateParamsObjectTypeOrganization  PrismPropertyUpdateParamsObjectType = "organization"
	PrismPropertyUpdateParamsObjectTypeContact       PrismPropertyUpdateParamsObjectType = "contact"
)

func (r PrismPropertyUpdateParamsObjectType) IsKnown() bool {
	switch r {
	case PrismPropertyUpdateParamsObjectTypeComment, PrismPropertyUpdateParamsObjectTypeDeal, PrismPropertyUpdateParamsObjectTypeEngagement, PrismPropertyUpdateParamsObjectTypeIdentity, PrismPropertyUpdateParamsObjectTypeAIChatThread, PrismPropertyUpdateParamsObjectTypeAIChatMessage, PrismPropertyUpdateParamsObjectTypeAgentSite, PrismPropertyUpdateParamsObjectTypeDocument, PrismPropertyUpdateParamsObjectTypeAction, PrismPropertyUpdateParamsObjectTypeEvent, PrismPropertyUpdateParamsObjectTypeOrganization, PrismPropertyUpdateParamsObjectTypeContact:
		return true
	}
	return false
}

type PrismPropertyListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID   param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	Autofill param.Field[bool]   `query:"autofill"`
	// When false, return property definitions without hydrating select/multiselect
	// option rows. Defaults to true server-side (parseIncludeOptions). Accepts boolean
	// or query-string forms (true/false/0/1). Uses anyOf (not oneOf) so qs/AJV
	// boolean-vs-string ambiguity does not 400 when Speakeasy SDKs send
	// include_options=true.
	IncludeOptions param.Field[PrismPropertyListParamsIncludeOptionsUnion] `query:"include_options"`
	// Scope properties to a specific list/app. Scoping is strict: the response carries
	// only that list's definitions, not the workspace-global ones that also apply to
	// its records. Call once with `list_id` and once without to see everything a write
	// could resolve against.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
	// Case-insensitive substring match on the property name. Use this to find an
	// existing property before creating a new one.
	Term param.Field[string] `query:"term"`
}

// URLQuery serializes [PrismPropertyListParams]'s query parameters as
// `url.Values`.
func (r PrismPropertyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Object types that support CRUD, query, list, and per-type property metadata.
// `GET /v2/prism/{teamId}/properties` (list-all) also returns definitions for
// pipeline-owned types that are not in this set — including `message`, `thread`,
// and `linkedin_thread`. Those types are not queryable. Contacts expose
// `last_email` as a `ref_message`; you cannot query `message` to follow it.
type PrismPropertyListParamsObjectType string

const (
	PrismPropertyListParamsObjectTypeComment       PrismPropertyListParamsObjectType = "comment"
	PrismPropertyListParamsObjectTypeDeal          PrismPropertyListParamsObjectType = "deal"
	PrismPropertyListParamsObjectTypeEngagement    PrismPropertyListParamsObjectType = "engagement"
	PrismPropertyListParamsObjectTypeIdentity      PrismPropertyListParamsObjectType = "identity"
	PrismPropertyListParamsObjectTypeAIChatThread  PrismPropertyListParamsObjectType = "ai_chat_thread"
	PrismPropertyListParamsObjectTypeAIChatMessage PrismPropertyListParamsObjectType = "ai_chat_message"
	PrismPropertyListParamsObjectTypeAgentSite     PrismPropertyListParamsObjectType = "agent_site"
	PrismPropertyListParamsObjectTypeDocument      PrismPropertyListParamsObjectType = "document"
	PrismPropertyListParamsObjectTypeAction        PrismPropertyListParamsObjectType = "action"
	PrismPropertyListParamsObjectTypeEvent         PrismPropertyListParamsObjectType = "event"
	PrismPropertyListParamsObjectTypeOrganization  PrismPropertyListParamsObjectType = "organization"
	PrismPropertyListParamsObjectTypeContact       PrismPropertyListParamsObjectType = "contact"
)

func (r PrismPropertyListParamsObjectType) IsKnown() bool {
	switch r {
	case PrismPropertyListParamsObjectTypeComment, PrismPropertyListParamsObjectTypeDeal, PrismPropertyListParamsObjectTypeEngagement, PrismPropertyListParamsObjectTypeIdentity, PrismPropertyListParamsObjectTypeAIChatThread, PrismPropertyListParamsObjectTypeAIChatMessage, PrismPropertyListParamsObjectTypeAgentSite, PrismPropertyListParamsObjectTypeDocument, PrismPropertyListParamsObjectTypeAction, PrismPropertyListParamsObjectTypeEvent, PrismPropertyListParamsObjectTypeOrganization, PrismPropertyListParamsObjectTypeContact:
		return true
	}
	return false
}

// When false, return property definitions without hydrating select/multiselect
// option rows. Defaults to true server-side (parseIncludeOptions). Accepts boolean
// or query-string forms (true/false/0/1). Uses anyOf (not oneOf) so qs/AJV
// boolean-vs-string ambiguity does not 400 when Speakeasy SDKs send
// include_options=true.
//
// Satisfied by [shared.UnionBool], [PrismPropertyListParamsIncludeOptionsString].
type PrismPropertyListParamsIncludeOptionsUnion interface {
	ImplementsPrismPropertyListParamsIncludeOptionsUnion()
}

type PrismPropertyListParamsIncludeOptionsString string

const (
	PrismPropertyListParamsIncludeOptionsStringTrue  PrismPropertyListParamsIncludeOptionsString = "true"
	PrismPropertyListParamsIncludeOptionsStringFalse PrismPropertyListParamsIncludeOptionsString = "false"
	PrismPropertyListParamsIncludeOptionsString0     PrismPropertyListParamsIncludeOptionsString = "0"
	PrismPropertyListParamsIncludeOptionsString1     PrismPropertyListParamsIncludeOptionsString = "1"
)

func (r PrismPropertyListParamsIncludeOptionsString) IsKnown() bool {
	switch r {
	case PrismPropertyListParamsIncludeOptionsStringTrue, PrismPropertyListParamsIncludeOptionsStringFalse, PrismPropertyListParamsIncludeOptionsString0, PrismPropertyListParamsIncludeOptionsString1:
		return true
	}
	return false
}

func (r PrismPropertyListParamsIncludeOptionsString) ImplementsPrismPropertyListParamsIncludeOptionsUnion() {
}

type PrismPropertyDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Storage type of this property definition.
	Type   param.Field[PrismPropertyDeleteParamsType] `query:"type" api:"required"`
	ListID param.Field[string]                        `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismPropertyDeleteParams]'s query parameters as
// `url.Values`.
func (r PrismPropertyDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Object types that support CRUD, query, list, and per-type property metadata.
// `GET /v2/prism/{teamId}/properties` (list-all) also returns definitions for
// pipeline-owned types that are not in this set — including `message`, `thread`,
// and `linkedin_thread`. Those types are not queryable. Contacts expose
// `last_email` as a `ref_message`; you cannot query `message` to follow it.
type PrismPropertyDeleteParamsObjectType string

const (
	PrismPropertyDeleteParamsObjectTypeComment       PrismPropertyDeleteParamsObjectType = "comment"
	PrismPropertyDeleteParamsObjectTypeDeal          PrismPropertyDeleteParamsObjectType = "deal"
	PrismPropertyDeleteParamsObjectTypeEngagement    PrismPropertyDeleteParamsObjectType = "engagement"
	PrismPropertyDeleteParamsObjectTypeIdentity      PrismPropertyDeleteParamsObjectType = "identity"
	PrismPropertyDeleteParamsObjectTypeAIChatThread  PrismPropertyDeleteParamsObjectType = "ai_chat_thread"
	PrismPropertyDeleteParamsObjectTypeAIChatMessage PrismPropertyDeleteParamsObjectType = "ai_chat_message"
	PrismPropertyDeleteParamsObjectTypeAgentSite     PrismPropertyDeleteParamsObjectType = "agent_site"
	PrismPropertyDeleteParamsObjectTypeDocument      PrismPropertyDeleteParamsObjectType = "document"
	PrismPropertyDeleteParamsObjectTypeAction        PrismPropertyDeleteParamsObjectType = "action"
	PrismPropertyDeleteParamsObjectTypeEvent         PrismPropertyDeleteParamsObjectType = "event"
	PrismPropertyDeleteParamsObjectTypeOrganization  PrismPropertyDeleteParamsObjectType = "organization"
	PrismPropertyDeleteParamsObjectTypeContact       PrismPropertyDeleteParamsObjectType = "contact"
)

func (r PrismPropertyDeleteParamsObjectType) IsKnown() bool {
	switch r {
	case PrismPropertyDeleteParamsObjectTypeComment, PrismPropertyDeleteParamsObjectTypeDeal, PrismPropertyDeleteParamsObjectTypeEngagement, PrismPropertyDeleteParamsObjectTypeIdentity, PrismPropertyDeleteParamsObjectTypeAIChatThread, PrismPropertyDeleteParamsObjectTypeAIChatMessage, PrismPropertyDeleteParamsObjectTypeAgentSite, PrismPropertyDeleteParamsObjectTypeDocument, PrismPropertyDeleteParamsObjectTypeAction, PrismPropertyDeleteParamsObjectTypeEvent, PrismPropertyDeleteParamsObjectTypeOrganization, PrismPropertyDeleteParamsObjectTypeContact:
		return true
	}
	return false
}

// Storage type of this property definition.
type PrismPropertyDeleteParamsType string

const (
	PrismPropertyDeleteParamsTypeNum                   PrismPropertyDeleteParamsType = "num"
	PrismPropertyDeleteParamsTypeStr                   PrismPropertyDeleteParamsType = "str"
	PrismPropertyDeleteParamsTypeBool                  PrismPropertyDeleteParamsType = "bool"
	PrismPropertyDeleteParamsTypeDate                  PrismPropertyDeleteParamsType = "date"
	PrismPropertyDeleteParamsTypeText                  PrismPropertyDeleteParamsType = "text"
	PrismPropertyDeleteParamsTypeByte                  PrismPropertyDeleteParamsType = "byte"
	PrismPropertyDeleteParamsTypeSelectStr             PrismPropertyDeleteParamsType = "select_str"
	PrismPropertyDeleteParamsTypeMultiStr              PrismPropertyDeleteParamsType = "multi_str"
	PrismPropertyDeleteParamsTypeMultiselectStr        PrismPropertyDeleteParamsType = "multiselect_str"
	PrismPropertyDeleteParamsTypeJsonb                 PrismPropertyDeleteParamsType = "jsonb"
	PrismPropertyDeleteParamsTypeRefIdentity           PrismPropertyDeleteParamsType = "ref_identity"
	PrismPropertyDeleteParamsTypeRefUser               PrismPropertyDeleteParamsType = "ref_user"
	PrismPropertyDeleteParamsTypeRefOrganization       PrismPropertyDeleteParamsType = "ref_organization"
	PrismPropertyDeleteParamsTypeRefContact            PrismPropertyDeleteParamsType = "ref_contact"
	PrismPropertyDeleteParamsTypeRefThread             PrismPropertyDeleteParamsType = "ref_thread"
	PrismPropertyDeleteParamsTypeRefMessage            PrismPropertyDeleteParamsType = "ref_message"
	PrismPropertyDeleteParamsTypeRefEvent              PrismPropertyDeleteParamsType = "ref_event"
	PrismPropertyDeleteParamsTypeRefAccount            PrismPropertyDeleteParamsType = "ref_account"
	PrismPropertyDeleteParamsTypeRefAIChatThread       PrismPropertyDeleteParamsType = "ref_ai_chat_thread"
	PrismPropertyDeleteParamsTypeRefAIChatMessage      PrismPropertyDeleteParamsType = "ref_ai_chat_message"
	PrismPropertyDeleteParamsTypeMultirefAIChatMessage PrismPropertyDeleteParamsType = "multiref_ai_chat_message"
	PrismPropertyDeleteParamsTypeMultirefAgentSite     PrismPropertyDeleteParamsType = "multiref_agent_site"
	PrismPropertyDeleteParamsTypeMultirefAction        PrismPropertyDeleteParamsType = "multiref_action"
	PrismPropertyDeleteParamsTypeMultirefComment       PrismPropertyDeleteParamsType = "multiref_comment"
	PrismPropertyDeleteParamsTypeMultirefContact       PrismPropertyDeleteParamsType = "multiref_contact"
	PrismPropertyDeleteParamsTypeMultirefLabel         PrismPropertyDeleteParamsType = "multiref_label"
	PrismPropertyDeleteParamsTypeMultirefThread        PrismPropertyDeleteParamsType = "multiref_thread"
	PrismPropertyDeleteParamsTypeMultirefMessages      PrismPropertyDeleteParamsType = "multiref_messages"
	PrismPropertyDeleteParamsTypeMultirefDocument      PrismPropertyDeleteParamsType = "multiref_document"
	PrismPropertyDeleteParamsTypeMultirefIdentity      PrismPropertyDeleteParamsType = "multiref_identity"
	PrismPropertyDeleteParamsTypeMultirefOrganization  PrismPropertyDeleteParamsType = "multiref_organization"
	PrismPropertyDeleteParamsTypeMultirefEngagement    PrismPropertyDeleteParamsType = "multiref_engagement"
	PrismPropertyDeleteParamsTypeMultirefAttendee      PrismPropertyDeleteParamsType = "multiref_attendee"
	PrismPropertyDeleteParamsTypeMultirefMeetingEntry  PrismPropertyDeleteParamsType = "multiref_meeting_entry"
	PrismPropertyDeleteParamsTypeMultirefReadReceipt   PrismPropertyDeleteParamsType = "multiref_read_receipt"
	PrismPropertyDeleteParamsTypeMultirefAccount       PrismPropertyDeleteParamsType = "multiref_account"
	PrismPropertyDeleteParamsTypeMultirefSource        PrismPropertyDeleteParamsType = "multiref_source"
)

func (r PrismPropertyDeleteParamsType) IsKnown() bool {
	switch r {
	case PrismPropertyDeleteParamsTypeNum, PrismPropertyDeleteParamsTypeStr, PrismPropertyDeleteParamsTypeBool, PrismPropertyDeleteParamsTypeDate, PrismPropertyDeleteParamsTypeText, PrismPropertyDeleteParamsTypeByte, PrismPropertyDeleteParamsTypeSelectStr, PrismPropertyDeleteParamsTypeMultiStr, PrismPropertyDeleteParamsTypeMultiselectStr, PrismPropertyDeleteParamsTypeJsonb, PrismPropertyDeleteParamsTypeRefIdentity, PrismPropertyDeleteParamsTypeRefUser, PrismPropertyDeleteParamsTypeRefOrganization, PrismPropertyDeleteParamsTypeRefContact, PrismPropertyDeleteParamsTypeRefThread, PrismPropertyDeleteParamsTypeRefMessage, PrismPropertyDeleteParamsTypeRefEvent, PrismPropertyDeleteParamsTypeRefAccount, PrismPropertyDeleteParamsTypeRefAIChatThread, PrismPropertyDeleteParamsTypeRefAIChatMessage, PrismPropertyDeleteParamsTypeMultirefAIChatMessage, PrismPropertyDeleteParamsTypeMultirefAgentSite, PrismPropertyDeleteParamsTypeMultirefAction, PrismPropertyDeleteParamsTypeMultirefComment, PrismPropertyDeleteParamsTypeMultirefContact, PrismPropertyDeleteParamsTypeMultirefLabel, PrismPropertyDeleteParamsTypeMultirefThread, PrismPropertyDeleteParamsTypeMultirefMessages, PrismPropertyDeleteParamsTypeMultirefDocument, PrismPropertyDeleteParamsTypeMultirefIdentity, PrismPropertyDeleteParamsTypeMultirefOrganization, PrismPropertyDeleteParamsTypeMultirefEngagement, PrismPropertyDeleteParamsTypeMultirefAttendee, PrismPropertyDeleteParamsTypeMultirefMeetingEntry, PrismPropertyDeleteParamsTypeMultirefReadReceipt, PrismPropertyDeleteParamsTypeMultirefAccount, PrismPropertyDeleteParamsTypeMultirefSource:
		return true
	}
	return false
}

type PrismPropertyListAllParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID   param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	Autofill param.Field[bool]   `query:"autofill"`
	// When false, return property definitions without hydrating select/multiselect
	// option rows. Defaults to true server-side (parseIncludeOptions). Accepts boolean
	// or query-string forms (true/false/0/1). Uses anyOf (not oneOf) so qs/AJV
	// boolean-vs-string ambiguity does not 400 when Speakeasy SDKs send
	// include_options=true.
	IncludeOptions param.Field[PrismPropertyListAllParamsIncludeOptionsUnion] `query:"include_options"`
	// Scope properties to a specific list/app. Scoping is strict: the response carries
	// only that list's definitions, not the workspace-global ones that also apply to
	// its records. Call once with `list_id` and once without to see everything a write
	// could resolve against.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
	// Case-insensitive substring match on the property name. Use this to find an
	// existing property before creating a new one.
	Term param.Field[string] `query:"term"`
}

// URLQuery serializes [PrismPropertyListAllParams]'s query parameters as
// `url.Values`.
func (r PrismPropertyListAllParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// When false, return property definitions without hydrating select/multiselect
// option rows. Defaults to true server-side (parseIncludeOptions). Accepts boolean
// or query-string forms (true/false/0/1). Uses anyOf (not oneOf) so qs/AJV
// boolean-vs-string ambiguity does not 400 when Speakeasy SDKs send
// include_options=true.
//
// Satisfied by [shared.UnionBool],
// [PrismPropertyListAllParamsIncludeOptionsString].
type PrismPropertyListAllParamsIncludeOptionsUnion interface {
	ImplementsPrismPropertyListAllParamsIncludeOptionsUnion()
}

type PrismPropertyListAllParamsIncludeOptionsString string

const (
	PrismPropertyListAllParamsIncludeOptionsStringTrue  PrismPropertyListAllParamsIncludeOptionsString = "true"
	PrismPropertyListAllParamsIncludeOptionsStringFalse PrismPropertyListAllParamsIncludeOptionsString = "false"
	PrismPropertyListAllParamsIncludeOptionsString0     PrismPropertyListAllParamsIncludeOptionsString = "0"
	PrismPropertyListAllParamsIncludeOptionsString1     PrismPropertyListAllParamsIncludeOptionsString = "1"
)

func (r PrismPropertyListAllParamsIncludeOptionsString) IsKnown() bool {
	switch r {
	case PrismPropertyListAllParamsIncludeOptionsStringTrue, PrismPropertyListAllParamsIncludeOptionsStringFalse, PrismPropertyListAllParamsIncludeOptionsString0, PrismPropertyListAllParamsIncludeOptionsString1:
		return true
	}
	return false
}

func (r PrismPropertyListAllParamsIncludeOptionsString) ImplementsPrismPropertyListAllParamsIncludeOptionsUnion() {
}
