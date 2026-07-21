// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

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
}

// NewPrismPropertyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrismPropertyService(opts ...option.RequestOption) (r *PrismPropertyService) {
	r = &PrismPropertyService{}
	r.Options = opts
	return
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

// Get metadata properties
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

type PrismPropertyListResponse map[string]interface{}

type PrismPropertyListAllResponse map[string]interface{}

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
	// Scope properties to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
	Term   param.Field[string] `query:"term"`
}

// URLQuery serializes [PrismPropertyListParams]'s query parameters as
// `url.Values`.
func (r PrismPropertyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismPropertyListParamsObjectType string

const (
	PrismPropertyListParamsObjectTypeComment       PrismPropertyListParamsObjectType = "comment"
	PrismPropertyListParamsObjectTypeDeal          PrismPropertyListParamsObjectType = "deal"
	PrismPropertyListParamsObjectTypeEngagement    PrismPropertyListParamsObjectType = "engagement"
	PrismPropertyListParamsObjectTypeIdentity      PrismPropertyListParamsObjectType = "identity"
	PrismPropertyListParamsObjectTypeAIChatThread  PrismPropertyListParamsObjectType = "ai_chat_thread"
	PrismPropertyListParamsObjectTypeAIChatMessage PrismPropertyListParamsObjectType = "ai_chat_message"
	PrismPropertyListParamsObjectTypeDocument      PrismPropertyListParamsObjectType = "document"
	PrismPropertyListParamsObjectTypeAction        PrismPropertyListParamsObjectType = "action"
	PrismPropertyListParamsObjectTypeEvent         PrismPropertyListParamsObjectType = "event"
	PrismPropertyListParamsObjectTypeOrganization  PrismPropertyListParamsObjectType = "organization"
	PrismPropertyListParamsObjectTypeContact       PrismPropertyListParamsObjectType = "contact"
)

func (r PrismPropertyListParamsObjectType) IsKnown() bool {
	switch r {
	case PrismPropertyListParamsObjectTypeComment, PrismPropertyListParamsObjectTypeDeal, PrismPropertyListParamsObjectTypeEngagement, PrismPropertyListParamsObjectTypeIdentity, PrismPropertyListParamsObjectTypeAIChatThread, PrismPropertyListParamsObjectTypeAIChatMessage, PrismPropertyListParamsObjectTypeDocument, PrismPropertyListParamsObjectTypeAction, PrismPropertyListParamsObjectTypeEvent, PrismPropertyListParamsObjectTypeOrganization, PrismPropertyListParamsObjectTypeContact:
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
	// Scope properties to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
	Term   param.Field[string] `query:"term"`
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
