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

// PrismPropertyOptionService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismPropertyOptionService] method instead.
type PrismPropertyOptionService struct {
	Options []option.RequestOption
}

// NewPrismPropertyOptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismPropertyOptionService(opts ...option.RequestOption) (r *PrismPropertyOptionService) {
	r = &PrismPropertyOptionService{}
	r.Options = opts
	return
}

// Adds a single option to a `select_str` or `multiselect_str` property definition.
// Body must include `type` so the server knows which per-type option table to
// write.
func (r *PrismPropertyOptionService) New(ctx context.Context, objectType PrismPropertyOptionNewParamsObjectType, propertyID string, params PrismPropertyOptionNewParams, opts ...option.RequestOption) (res *PropertyOption, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/%v/properties/%s/options", params.TeamID, objectType, propertyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a property option
func (r *PrismPropertyOptionService) Update(ctx context.Context, objectType PrismPropertyOptionUpdateParamsObjectType, propertyID string, optionID string, params PrismPropertyOptionUpdateParams, opts ...option.RequestOption) (res *PropertyOption, err error) {
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
	if optionID == "" {
		err = errors.New("missing required optionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/properties/%s/options/%s", params.TeamID, objectType, propertyID, optionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Delete a property option
func (r *PrismPropertyOptionService) Delete(ctx context.Context, objectType PrismPropertyOptionDeleteParamsObjectType, propertyID string, optionID string, params PrismPropertyOptionDeleteParams, opts ...option.RequestOption) (err error) {
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
	if optionID == "" {
		err = errors.New("missing required optionId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/properties/%s/options/%s", params.TeamID, objectType, propertyID, optionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// An enabled option for a select_str or multiselect_str property definition.
type PropertyOption struct {
	ID          string `json:"id" api:"required" format:"uuid"`
	Slug        string `json:"slug" api:"required"`
	ColorScheme string `json:"color_scheme" api:"nullable"`
	// Deprecated: deprecated
	CRMID       string `json:"crm_id" api:"nullable" format:"uuid"`
	Description string `json:"description" api:"nullable"`
	Icon        string `json:"icon" api:"nullable"`
	ListID      string `json:"list_id" api:"nullable" format:"uuid"`
	OptionGroup string `json:"option_group" api:"nullable"`
	SortIndex   int64  `json:"sort_index" api:"nullable"`
	// Display value for the option.
	Value string             `json:"value" api:"nullable"`
	JSON  propertyOptionJSON `json:"-"`
}

// propertyOptionJSON contains the JSON metadata for the struct [PropertyOption]
type propertyOptionJSON struct {
	ID          apijson.Field
	Slug        apijson.Field
	ColorScheme apijson.Field
	CRMID       apijson.Field
	Description apijson.Field
	Icon        apijson.Field
	ListID      apijson.Field
	OptionGroup apijson.Field
	SortIndex   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PropertyOption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r propertyOptionJSON) RawJSON() string {
	return r.raw
}

// New option for a `select_str` or `multiselect_str` property. `type` identifies
// the per-type option table to write.
type PropertyOptionCreateParam struct {
	// Storage type for a property definition. Determines which per-type table holds
	// the values, and which display formats the property can take.
	Type param.Field[PropertyOptionCreateType] `json:"type" api:"required"`
	// Display value for the option.
	Value       param.Field[string] `json:"value" api:"required"`
	ColorScheme param.Field[string] `json:"color_scheme"`
	Description param.Field[string] `json:"description"`
	Icon        param.Field[string] `json:"icon"`
	// Scope the option to a specific list/app.
	ListID      param.Field[string] `json:"list_id" format:"uuid"`
	OptionGroup param.Field[string] `json:"option_group"`
	// URL-safe identifier. Defaults to a slugified `value`.
	Slug      param.Field[string] `json:"slug"`
	SortIndex param.Field[int64]  `json:"sort_index"`
}

func (r PropertyOptionCreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Storage type for a property definition. Determines which per-type table holds
// the values, and which display formats the property can take.
type PropertyOptionCreateType string

const (
	PropertyOptionCreateTypeNum                   PropertyOptionCreateType = "num"
	PropertyOptionCreateTypeStr                   PropertyOptionCreateType = "str"
	PropertyOptionCreateTypeBool                  PropertyOptionCreateType = "bool"
	PropertyOptionCreateTypeDate                  PropertyOptionCreateType = "date"
	PropertyOptionCreateTypeText                  PropertyOptionCreateType = "text"
	PropertyOptionCreateTypeByte                  PropertyOptionCreateType = "byte"
	PropertyOptionCreateTypeSelectStr             PropertyOptionCreateType = "select_str"
	PropertyOptionCreateTypeMultiStr              PropertyOptionCreateType = "multi_str"
	PropertyOptionCreateTypeMultiselectStr        PropertyOptionCreateType = "multiselect_str"
	PropertyOptionCreateTypeJsonb                 PropertyOptionCreateType = "jsonb"
	PropertyOptionCreateTypeRefIdentity           PropertyOptionCreateType = "ref_identity"
	PropertyOptionCreateTypeRefUser               PropertyOptionCreateType = "ref_user"
	PropertyOptionCreateTypeRefOrganization       PropertyOptionCreateType = "ref_organization"
	PropertyOptionCreateTypeRefContact            PropertyOptionCreateType = "ref_contact"
	PropertyOptionCreateTypeRefThread             PropertyOptionCreateType = "ref_thread"
	PropertyOptionCreateTypeRefMessage            PropertyOptionCreateType = "ref_message"
	PropertyOptionCreateTypeRefEvent              PropertyOptionCreateType = "ref_event"
	PropertyOptionCreateTypeRefAccount            PropertyOptionCreateType = "ref_account"
	PropertyOptionCreateTypeRefAIChatThread       PropertyOptionCreateType = "ref_ai_chat_thread"
	PropertyOptionCreateTypeRefAIChatMessage      PropertyOptionCreateType = "ref_ai_chat_message"
	PropertyOptionCreateTypeMultirefAIChatMessage PropertyOptionCreateType = "multiref_ai_chat_message"
	PropertyOptionCreateTypeMultirefAgentArtifact PropertyOptionCreateType = "multiref_agent_artifact"
	PropertyOptionCreateTypeMultirefAction        PropertyOptionCreateType = "multiref_action"
	PropertyOptionCreateTypeMultirefComment       PropertyOptionCreateType = "multiref_comment"
	PropertyOptionCreateTypeMultirefContact       PropertyOptionCreateType = "multiref_contact"
	PropertyOptionCreateTypeMultirefLabel         PropertyOptionCreateType = "multiref_label"
	PropertyOptionCreateTypeMultirefThread        PropertyOptionCreateType = "multiref_thread"
	PropertyOptionCreateTypeMultirefMessages      PropertyOptionCreateType = "multiref_messages"
	PropertyOptionCreateTypeMultirefDocument      PropertyOptionCreateType = "multiref_document"
	PropertyOptionCreateTypeMultirefIdentity      PropertyOptionCreateType = "multiref_identity"
	PropertyOptionCreateTypeMultirefOrganization  PropertyOptionCreateType = "multiref_organization"
	PropertyOptionCreateTypeMultirefEngagement    PropertyOptionCreateType = "multiref_engagement"
	PropertyOptionCreateTypeMultirefAttendee      PropertyOptionCreateType = "multiref_attendee"
	PropertyOptionCreateTypeMultirefMeetingEntry  PropertyOptionCreateType = "multiref_meeting_entry"
	PropertyOptionCreateTypeMultirefReadReceipt   PropertyOptionCreateType = "multiref_read_receipt"
	PropertyOptionCreateTypeMultirefAccount       PropertyOptionCreateType = "multiref_account"
	PropertyOptionCreateTypeMultirefSource        PropertyOptionCreateType = "multiref_source"
)

func (r PropertyOptionCreateType) IsKnown() bool {
	switch r {
	case PropertyOptionCreateTypeNum, PropertyOptionCreateTypeStr, PropertyOptionCreateTypeBool, PropertyOptionCreateTypeDate, PropertyOptionCreateTypeText, PropertyOptionCreateTypeByte, PropertyOptionCreateTypeSelectStr, PropertyOptionCreateTypeMultiStr, PropertyOptionCreateTypeMultiselectStr, PropertyOptionCreateTypeJsonb, PropertyOptionCreateTypeRefIdentity, PropertyOptionCreateTypeRefUser, PropertyOptionCreateTypeRefOrganization, PropertyOptionCreateTypeRefContact, PropertyOptionCreateTypeRefThread, PropertyOptionCreateTypeRefMessage, PropertyOptionCreateTypeRefEvent, PropertyOptionCreateTypeRefAccount, PropertyOptionCreateTypeRefAIChatThread, PropertyOptionCreateTypeRefAIChatMessage, PropertyOptionCreateTypeMultirefAIChatMessage, PropertyOptionCreateTypeMultirefAgentArtifact, PropertyOptionCreateTypeMultirefAction, PropertyOptionCreateTypeMultirefComment, PropertyOptionCreateTypeMultirefContact, PropertyOptionCreateTypeMultirefLabel, PropertyOptionCreateTypeMultirefThread, PropertyOptionCreateTypeMultirefMessages, PropertyOptionCreateTypeMultirefDocument, PropertyOptionCreateTypeMultirefIdentity, PropertyOptionCreateTypeMultirefOrganization, PropertyOptionCreateTypeMultirefEngagement, PropertyOptionCreateTypeMultirefAttendee, PropertyOptionCreateTypeMultirefMeetingEntry, PropertyOptionCreateTypeMultirefReadReceipt, PropertyOptionCreateTypeMultirefAccount, PropertyOptionCreateTypeMultirefSource:
		return true
	}
	return false
}

// Partial update of a property option. `type` identifies the per-type option table
// to write.
type PropertyOptionPatchParam struct {
	// Storage type for a property definition. Determines which per-type table holds
	// the values, and which display formats the property can take.
	Type        param.Field[PropertyOptionPatchType] `json:"type" api:"required"`
	ColorScheme param.Field[string]                  `json:"color_scheme"`
	Description param.Field[string]                  `json:"description"`
	Enabled     param.Field[bool]                    `json:"enabled"`
	Icon        param.Field[string]                  `json:"icon"`
	ListID      param.Field[string]                  `json:"list_id" format:"uuid"`
	OptionGroup param.Field[string]                  `json:"option_group"`
	Slug        param.Field[string]                  `json:"slug"`
	SortIndex   param.Field[int64]                   `json:"sort_index"`
	Value       param.Field[string]                  `json:"value"`
}

func (r PropertyOptionPatchParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Storage type for a property definition. Determines which per-type table holds
// the values, and which display formats the property can take.
type PropertyOptionPatchType string

const (
	PropertyOptionPatchTypeNum                   PropertyOptionPatchType = "num"
	PropertyOptionPatchTypeStr                   PropertyOptionPatchType = "str"
	PropertyOptionPatchTypeBool                  PropertyOptionPatchType = "bool"
	PropertyOptionPatchTypeDate                  PropertyOptionPatchType = "date"
	PropertyOptionPatchTypeText                  PropertyOptionPatchType = "text"
	PropertyOptionPatchTypeByte                  PropertyOptionPatchType = "byte"
	PropertyOptionPatchTypeSelectStr             PropertyOptionPatchType = "select_str"
	PropertyOptionPatchTypeMultiStr              PropertyOptionPatchType = "multi_str"
	PropertyOptionPatchTypeMultiselectStr        PropertyOptionPatchType = "multiselect_str"
	PropertyOptionPatchTypeJsonb                 PropertyOptionPatchType = "jsonb"
	PropertyOptionPatchTypeRefIdentity           PropertyOptionPatchType = "ref_identity"
	PropertyOptionPatchTypeRefUser               PropertyOptionPatchType = "ref_user"
	PropertyOptionPatchTypeRefOrganization       PropertyOptionPatchType = "ref_organization"
	PropertyOptionPatchTypeRefContact            PropertyOptionPatchType = "ref_contact"
	PropertyOptionPatchTypeRefThread             PropertyOptionPatchType = "ref_thread"
	PropertyOptionPatchTypeRefMessage            PropertyOptionPatchType = "ref_message"
	PropertyOptionPatchTypeRefEvent              PropertyOptionPatchType = "ref_event"
	PropertyOptionPatchTypeRefAccount            PropertyOptionPatchType = "ref_account"
	PropertyOptionPatchTypeRefAIChatThread       PropertyOptionPatchType = "ref_ai_chat_thread"
	PropertyOptionPatchTypeRefAIChatMessage      PropertyOptionPatchType = "ref_ai_chat_message"
	PropertyOptionPatchTypeMultirefAIChatMessage PropertyOptionPatchType = "multiref_ai_chat_message"
	PropertyOptionPatchTypeMultirefAgentArtifact PropertyOptionPatchType = "multiref_agent_artifact"
	PropertyOptionPatchTypeMultirefAction        PropertyOptionPatchType = "multiref_action"
	PropertyOptionPatchTypeMultirefComment       PropertyOptionPatchType = "multiref_comment"
	PropertyOptionPatchTypeMultirefContact       PropertyOptionPatchType = "multiref_contact"
	PropertyOptionPatchTypeMultirefLabel         PropertyOptionPatchType = "multiref_label"
	PropertyOptionPatchTypeMultirefThread        PropertyOptionPatchType = "multiref_thread"
	PropertyOptionPatchTypeMultirefMessages      PropertyOptionPatchType = "multiref_messages"
	PropertyOptionPatchTypeMultirefDocument      PropertyOptionPatchType = "multiref_document"
	PropertyOptionPatchTypeMultirefIdentity      PropertyOptionPatchType = "multiref_identity"
	PropertyOptionPatchTypeMultirefOrganization  PropertyOptionPatchType = "multiref_organization"
	PropertyOptionPatchTypeMultirefEngagement    PropertyOptionPatchType = "multiref_engagement"
	PropertyOptionPatchTypeMultirefAttendee      PropertyOptionPatchType = "multiref_attendee"
	PropertyOptionPatchTypeMultirefMeetingEntry  PropertyOptionPatchType = "multiref_meeting_entry"
	PropertyOptionPatchTypeMultirefReadReceipt   PropertyOptionPatchType = "multiref_read_receipt"
	PropertyOptionPatchTypeMultirefAccount       PropertyOptionPatchType = "multiref_account"
	PropertyOptionPatchTypeMultirefSource        PropertyOptionPatchType = "multiref_source"
)

func (r PropertyOptionPatchType) IsKnown() bool {
	switch r {
	case PropertyOptionPatchTypeNum, PropertyOptionPatchTypeStr, PropertyOptionPatchTypeBool, PropertyOptionPatchTypeDate, PropertyOptionPatchTypeText, PropertyOptionPatchTypeByte, PropertyOptionPatchTypeSelectStr, PropertyOptionPatchTypeMultiStr, PropertyOptionPatchTypeMultiselectStr, PropertyOptionPatchTypeJsonb, PropertyOptionPatchTypeRefIdentity, PropertyOptionPatchTypeRefUser, PropertyOptionPatchTypeRefOrganization, PropertyOptionPatchTypeRefContact, PropertyOptionPatchTypeRefThread, PropertyOptionPatchTypeRefMessage, PropertyOptionPatchTypeRefEvent, PropertyOptionPatchTypeRefAccount, PropertyOptionPatchTypeRefAIChatThread, PropertyOptionPatchTypeRefAIChatMessage, PropertyOptionPatchTypeMultirefAIChatMessage, PropertyOptionPatchTypeMultirefAgentArtifact, PropertyOptionPatchTypeMultirefAction, PropertyOptionPatchTypeMultirefComment, PropertyOptionPatchTypeMultirefContact, PropertyOptionPatchTypeMultirefLabel, PropertyOptionPatchTypeMultirefThread, PropertyOptionPatchTypeMultirefMessages, PropertyOptionPatchTypeMultirefDocument, PropertyOptionPatchTypeMultirefIdentity, PropertyOptionPatchTypeMultirefOrganization, PropertyOptionPatchTypeMultirefEngagement, PropertyOptionPatchTypeMultirefAttendee, PropertyOptionPatchTypeMultirefMeetingEntry, PropertyOptionPatchTypeMultirefReadReceipt, PropertyOptionPatchTypeMultirefAccount, PropertyOptionPatchTypeMultirefSource:
		return true
	}
	return false
}

type PrismPropertyOptionNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// New option for a `select_str` or `multiselect_str` property. `type` identifies
	// the per-type option table to write.
	PropertyOptionCreate PropertyOptionCreateParam `json:"property_option_create" api:"required"`
	IdempotencyKey       param.Field[string]       `header:"Idempotency-Key"`
}

func (r PrismPropertyOptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PropertyOptionCreate)
}

type PrismPropertyOptionNewParamsObjectType string

const (
	PrismPropertyOptionNewParamsObjectTypeComment       PrismPropertyOptionNewParamsObjectType = "comment"
	PrismPropertyOptionNewParamsObjectTypeDeal          PrismPropertyOptionNewParamsObjectType = "deal"
	PrismPropertyOptionNewParamsObjectTypeEngagement    PrismPropertyOptionNewParamsObjectType = "engagement"
	PrismPropertyOptionNewParamsObjectTypeIdentity      PrismPropertyOptionNewParamsObjectType = "identity"
	PrismPropertyOptionNewParamsObjectTypeAIChatThread  PrismPropertyOptionNewParamsObjectType = "ai_chat_thread"
	PrismPropertyOptionNewParamsObjectTypeAIChatMessage PrismPropertyOptionNewParamsObjectType = "ai_chat_message"
	PrismPropertyOptionNewParamsObjectTypeAgentArtifact PrismPropertyOptionNewParamsObjectType = "agent_artifact"
	PrismPropertyOptionNewParamsObjectTypeDocument      PrismPropertyOptionNewParamsObjectType = "document"
	PrismPropertyOptionNewParamsObjectTypeAction        PrismPropertyOptionNewParamsObjectType = "action"
	PrismPropertyOptionNewParamsObjectTypeEvent         PrismPropertyOptionNewParamsObjectType = "event"
	PrismPropertyOptionNewParamsObjectTypeOrganization  PrismPropertyOptionNewParamsObjectType = "organization"
	PrismPropertyOptionNewParamsObjectTypeContact       PrismPropertyOptionNewParamsObjectType = "contact"
)

func (r PrismPropertyOptionNewParamsObjectType) IsKnown() bool {
	switch r {
	case PrismPropertyOptionNewParamsObjectTypeComment, PrismPropertyOptionNewParamsObjectTypeDeal, PrismPropertyOptionNewParamsObjectTypeEngagement, PrismPropertyOptionNewParamsObjectTypeIdentity, PrismPropertyOptionNewParamsObjectTypeAIChatThread, PrismPropertyOptionNewParamsObjectTypeAIChatMessage, PrismPropertyOptionNewParamsObjectTypeAgentArtifact, PrismPropertyOptionNewParamsObjectTypeDocument, PrismPropertyOptionNewParamsObjectTypeAction, PrismPropertyOptionNewParamsObjectTypeEvent, PrismPropertyOptionNewParamsObjectTypeOrganization, PrismPropertyOptionNewParamsObjectTypeContact:
		return true
	}
	return false
}

type PrismPropertyOptionUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Partial update of a property option. `type` identifies the per-type option table
	// to write.
	PropertyOptionPatch PropertyOptionPatchParam `json:"property_option_patch" api:"required"`
	IdempotencyKey      param.Field[string]      `header:"Idempotency-Key"`
}

func (r PrismPropertyOptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PropertyOptionPatch)
}

type PrismPropertyOptionUpdateParamsObjectType string

const (
	PrismPropertyOptionUpdateParamsObjectTypeComment       PrismPropertyOptionUpdateParamsObjectType = "comment"
	PrismPropertyOptionUpdateParamsObjectTypeDeal          PrismPropertyOptionUpdateParamsObjectType = "deal"
	PrismPropertyOptionUpdateParamsObjectTypeEngagement    PrismPropertyOptionUpdateParamsObjectType = "engagement"
	PrismPropertyOptionUpdateParamsObjectTypeIdentity      PrismPropertyOptionUpdateParamsObjectType = "identity"
	PrismPropertyOptionUpdateParamsObjectTypeAIChatThread  PrismPropertyOptionUpdateParamsObjectType = "ai_chat_thread"
	PrismPropertyOptionUpdateParamsObjectTypeAIChatMessage PrismPropertyOptionUpdateParamsObjectType = "ai_chat_message"
	PrismPropertyOptionUpdateParamsObjectTypeAgentArtifact PrismPropertyOptionUpdateParamsObjectType = "agent_artifact"
	PrismPropertyOptionUpdateParamsObjectTypeDocument      PrismPropertyOptionUpdateParamsObjectType = "document"
	PrismPropertyOptionUpdateParamsObjectTypeAction        PrismPropertyOptionUpdateParamsObjectType = "action"
	PrismPropertyOptionUpdateParamsObjectTypeEvent         PrismPropertyOptionUpdateParamsObjectType = "event"
	PrismPropertyOptionUpdateParamsObjectTypeOrganization  PrismPropertyOptionUpdateParamsObjectType = "organization"
	PrismPropertyOptionUpdateParamsObjectTypeContact       PrismPropertyOptionUpdateParamsObjectType = "contact"
)

func (r PrismPropertyOptionUpdateParamsObjectType) IsKnown() bool {
	switch r {
	case PrismPropertyOptionUpdateParamsObjectTypeComment, PrismPropertyOptionUpdateParamsObjectTypeDeal, PrismPropertyOptionUpdateParamsObjectTypeEngagement, PrismPropertyOptionUpdateParamsObjectTypeIdentity, PrismPropertyOptionUpdateParamsObjectTypeAIChatThread, PrismPropertyOptionUpdateParamsObjectTypeAIChatMessage, PrismPropertyOptionUpdateParamsObjectTypeAgentArtifact, PrismPropertyOptionUpdateParamsObjectTypeDocument, PrismPropertyOptionUpdateParamsObjectTypeAction, PrismPropertyOptionUpdateParamsObjectTypeEvent, PrismPropertyOptionUpdateParamsObjectTypeOrganization, PrismPropertyOptionUpdateParamsObjectTypeContact:
		return true
	}
	return false
}

type PrismPropertyOptionDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Storage type for a property definition. Determines which per-type table holds
	// the values, and which display formats the property can take.
	Type   param.Field[PrismPropertyOptionDeleteParamsType] `query:"type" api:"required"`
	ListID param.Field[string]                              `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismPropertyOptionDeleteParams]'s query parameters as
// `url.Values`.
func (r PrismPropertyOptionDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismPropertyOptionDeleteParamsObjectType string

const (
	PrismPropertyOptionDeleteParamsObjectTypeComment       PrismPropertyOptionDeleteParamsObjectType = "comment"
	PrismPropertyOptionDeleteParamsObjectTypeDeal          PrismPropertyOptionDeleteParamsObjectType = "deal"
	PrismPropertyOptionDeleteParamsObjectTypeEngagement    PrismPropertyOptionDeleteParamsObjectType = "engagement"
	PrismPropertyOptionDeleteParamsObjectTypeIdentity      PrismPropertyOptionDeleteParamsObjectType = "identity"
	PrismPropertyOptionDeleteParamsObjectTypeAIChatThread  PrismPropertyOptionDeleteParamsObjectType = "ai_chat_thread"
	PrismPropertyOptionDeleteParamsObjectTypeAIChatMessage PrismPropertyOptionDeleteParamsObjectType = "ai_chat_message"
	PrismPropertyOptionDeleteParamsObjectTypeAgentArtifact PrismPropertyOptionDeleteParamsObjectType = "agent_artifact"
	PrismPropertyOptionDeleteParamsObjectTypeDocument      PrismPropertyOptionDeleteParamsObjectType = "document"
	PrismPropertyOptionDeleteParamsObjectTypeAction        PrismPropertyOptionDeleteParamsObjectType = "action"
	PrismPropertyOptionDeleteParamsObjectTypeEvent         PrismPropertyOptionDeleteParamsObjectType = "event"
	PrismPropertyOptionDeleteParamsObjectTypeOrganization  PrismPropertyOptionDeleteParamsObjectType = "organization"
	PrismPropertyOptionDeleteParamsObjectTypeContact       PrismPropertyOptionDeleteParamsObjectType = "contact"
)

func (r PrismPropertyOptionDeleteParamsObjectType) IsKnown() bool {
	switch r {
	case PrismPropertyOptionDeleteParamsObjectTypeComment, PrismPropertyOptionDeleteParamsObjectTypeDeal, PrismPropertyOptionDeleteParamsObjectTypeEngagement, PrismPropertyOptionDeleteParamsObjectTypeIdentity, PrismPropertyOptionDeleteParamsObjectTypeAIChatThread, PrismPropertyOptionDeleteParamsObjectTypeAIChatMessage, PrismPropertyOptionDeleteParamsObjectTypeAgentArtifact, PrismPropertyOptionDeleteParamsObjectTypeDocument, PrismPropertyOptionDeleteParamsObjectTypeAction, PrismPropertyOptionDeleteParamsObjectTypeEvent, PrismPropertyOptionDeleteParamsObjectTypeOrganization, PrismPropertyOptionDeleteParamsObjectTypeContact:
		return true
	}
	return false
}

// Storage type for a property definition. Determines which per-type table holds
// the values, and which display formats the property can take.
type PrismPropertyOptionDeleteParamsType string

const (
	PrismPropertyOptionDeleteParamsTypeNum                   PrismPropertyOptionDeleteParamsType = "num"
	PrismPropertyOptionDeleteParamsTypeStr                   PrismPropertyOptionDeleteParamsType = "str"
	PrismPropertyOptionDeleteParamsTypeBool                  PrismPropertyOptionDeleteParamsType = "bool"
	PrismPropertyOptionDeleteParamsTypeDate                  PrismPropertyOptionDeleteParamsType = "date"
	PrismPropertyOptionDeleteParamsTypeText                  PrismPropertyOptionDeleteParamsType = "text"
	PrismPropertyOptionDeleteParamsTypeByte                  PrismPropertyOptionDeleteParamsType = "byte"
	PrismPropertyOptionDeleteParamsTypeSelectStr             PrismPropertyOptionDeleteParamsType = "select_str"
	PrismPropertyOptionDeleteParamsTypeMultiStr              PrismPropertyOptionDeleteParamsType = "multi_str"
	PrismPropertyOptionDeleteParamsTypeMultiselectStr        PrismPropertyOptionDeleteParamsType = "multiselect_str"
	PrismPropertyOptionDeleteParamsTypeJsonb                 PrismPropertyOptionDeleteParamsType = "jsonb"
	PrismPropertyOptionDeleteParamsTypeRefIdentity           PrismPropertyOptionDeleteParamsType = "ref_identity"
	PrismPropertyOptionDeleteParamsTypeRefUser               PrismPropertyOptionDeleteParamsType = "ref_user"
	PrismPropertyOptionDeleteParamsTypeRefOrganization       PrismPropertyOptionDeleteParamsType = "ref_organization"
	PrismPropertyOptionDeleteParamsTypeRefContact            PrismPropertyOptionDeleteParamsType = "ref_contact"
	PrismPropertyOptionDeleteParamsTypeRefThread             PrismPropertyOptionDeleteParamsType = "ref_thread"
	PrismPropertyOptionDeleteParamsTypeRefMessage            PrismPropertyOptionDeleteParamsType = "ref_message"
	PrismPropertyOptionDeleteParamsTypeRefEvent              PrismPropertyOptionDeleteParamsType = "ref_event"
	PrismPropertyOptionDeleteParamsTypeRefAccount            PrismPropertyOptionDeleteParamsType = "ref_account"
	PrismPropertyOptionDeleteParamsTypeRefAIChatThread       PrismPropertyOptionDeleteParamsType = "ref_ai_chat_thread"
	PrismPropertyOptionDeleteParamsTypeRefAIChatMessage      PrismPropertyOptionDeleteParamsType = "ref_ai_chat_message"
	PrismPropertyOptionDeleteParamsTypeMultirefAIChatMessage PrismPropertyOptionDeleteParamsType = "multiref_ai_chat_message"
	PrismPropertyOptionDeleteParamsTypeMultirefAgentArtifact PrismPropertyOptionDeleteParamsType = "multiref_agent_artifact"
	PrismPropertyOptionDeleteParamsTypeMultirefAction        PrismPropertyOptionDeleteParamsType = "multiref_action"
	PrismPropertyOptionDeleteParamsTypeMultirefComment       PrismPropertyOptionDeleteParamsType = "multiref_comment"
	PrismPropertyOptionDeleteParamsTypeMultirefContact       PrismPropertyOptionDeleteParamsType = "multiref_contact"
	PrismPropertyOptionDeleteParamsTypeMultirefLabel         PrismPropertyOptionDeleteParamsType = "multiref_label"
	PrismPropertyOptionDeleteParamsTypeMultirefThread        PrismPropertyOptionDeleteParamsType = "multiref_thread"
	PrismPropertyOptionDeleteParamsTypeMultirefMessages      PrismPropertyOptionDeleteParamsType = "multiref_messages"
	PrismPropertyOptionDeleteParamsTypeMultirefDocument      PrismPropertyOptionDeleteParamsType = "multiref_document"
	PrismPropertyOptionDeleteParamsTypeMultirefIdentity      PrismPropertyOptionDeleteParamsType = "multiref_identity"
	PrismPropertyOptionDeleteParamsTypeMultirefOrganization  PrismPropertyOptionDeleteParamsType = "multiref_organization"
	PrismPropertyOptionDeleteParamsTypeMultirefEngagement    PrismPropertyOptionDeleteParamsType = "multiref_engagement"
	PrismPropertyOptionDeleteParamsTypeMultirefAttendee      PrismPropertyOptionDeleteParamsType = "multiref_attendee"
	PrismPropertyOptionDeleteParamsTypeMultirefMeetingEntry  PrismPropertyOptionDeleteParamsType = "multiref_meeting_entry"
	PrismPropertyOptionDeleteParamsTypeMultirefReadReceipt   PrismPropertyOptionDeleteParamsType = "multiref_read_receipt"
	PrismPropertyOptionDeleteParamsTypeMultirefAccount       PrismPropertyOptionDeleteParamsType = "multiref_account"
	PrismPropertyOptionDeleteParamsTypeMultirefSource        PrismPropertyOptionDeleteParamsType = "multiref_source"
)

func (r PrismPropertyOptionDeleteParamsType) IsKnown() bool {
	switch r {
	case PrismPropertyOptionDeleteParamsTypeNum, PrismPropertyOptionDeleteParamsTypeStr, PrismPropertyOptionDeleteParamsTypeBool, PrismPropertyOptionDeleteParamsTypeDate, PrismPropertyOptionDeleteParamsTypeText, PrismPropertyOptionDeleteParamsTypeByte, PrismPropertyOptionDeleteParamsTypeSelectStr, PrismPropertyOptionDeleteParamsTypeMultiStr, PrismPropertyOptionDeleteParamsTypeMultiselectStr, PrismPropertyOptionDeleteParamsTypeJsonb, PrismPropertyOptionDeleteParamsTypeRefIdentity, PrismPropertyOptionDeleteParamsTypeRefUser, PrismPropertyOptionDeleteParamsTypeRefOrganization, PrismPropertyOptionDeleteParamsTypeRefContact, PrismPropertyOptionDeleteParamsTypeRefThread, PrismPropertyOptionDeleteParamsTypeRefMessage, PrismPropertyOptionDeleteParamsTypeRefEvent, PrismPropertyOptionDeleteParamsTypeRefAccount, PrismPropertyOptionDeleteParamsTypeRefAIChatThread, PrismPropertyOptionDeleteParamsTypeRefAIChatMessage, PrismPropertyOptionDeleteParamsTypeMultirefAIChatMessage, PrismPropertyOptionDeleteParamsTypeMultirefAgentArtifact, PrismPropertyOptionDeleteParamsTypeMultirefAction, PrismPropertyOptionDeleteParamsTypeMultirefComment, PrismPropertyOptionDeleteParamsTypeMultirefContact, PrismPropertyOptionDeleteParamsTypeMultirefLabel, PrismPropertyOptionDeleteParamsTypeMultirefThread, PrismPropertyOptionDeleteParamsTypeMultirefMessages, PrismPropertyOptionDeleteParamsTypeMultirefDocument, PrismPropertyOptionDeleteParamsTypeMultirefIdentity, PrismPropertyOptionDeleteParamsTypeMultirefOrganization, PrismPropertyOptionDeleteParamsTypeMultirefEngagement, PrismPropertyOptionDeleteParamsTypeMultirefAttendee, PrismPropertyOptionDeleteParamsTypeMultirefMeetingEntry, PrismPropertyOptionDeleteParamsTypeMultirefReadReceipt, PrismPropertyOptionDeleteParamsTypeMultirefAccount, PrismPropertyOptionDeleteParamsTypeMultirefSource:
		return true
	}
	return false
}
