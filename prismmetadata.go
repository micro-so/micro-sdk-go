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

// PrismMetadataService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismMetadataService] method instead.
type PrismMetadataService struct {
	Options []option.RequestOption
}

// NewPrismMetadataService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrismMetadataService(opts ...option.RequestOption) (r *PrismMetadataService) {
	r = &PrismMetadataService{}
	r.Options = opts
	return
}

// Get metadata properties by object type
func (r *PrismMetadataService) List(ctx context.Context, objectType PrismMetadataListParamsObjectType, params PrismMetadataListParams, opts ...option.RequestOption) (res *PrismMetadataListResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/metadata/properties/%s/%v", params.TeamID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type PrismMetadataListResponse map[string]interface{}

type PrismMetadataListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID   param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	Autofill param.Field[bool]   `query:"autofill"`
	CRMID    param.Field[string] `query:"crmId" format:"uuid"`
	Term     param.Field[string] `query:"term"`
}

// URLQuery serializes [PrismMetadataListParams]'s query parameters as
// `url.Values`.
func (r PrismMetadataListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismMetadataListParamsObjectType string

const (
	PrismMetadataListParamsObjectTypeDeal          PrismMetadataListParamsObjectType = "deal"
	PrismMetadataListParamsObjectTypeIdentity      PrismMetadataListParamsObjectType = "identity"
	PrismMetadataListParamsObjectTypeAIChatThread  PrismMetadataListParamsObjectType = "ai_chat_thread"
	PrismMetadataListParamsObjectTypeAIChatMessage PrismMetadataListParamsObjectType = "ai_chat_message"
	PrismMetadataListParamsObjectTypeDocument      PrismMetadataListParamsObjectType = "document"
	PrismMetadataListParamsObjectTypeAction        PrismMetadataListParamsObjectType = "action"
	PrismMetadataListParamsObjectTypeEvent         PrismMetadataListParamsObjectType = "event"
	PrismMetadataListParamsObjectTypeOrganization  PrismMetadataListParamsObjectType = "organization"
	PrismMetadataListParamsObjectTypeContact       PrismMetadataListParamsObjectType = "contact"
)

func (r PrismMetadataListParamsObjectType) IsKnown() bool {
	switch r {
	case PrismMetadataListParamsObjectTypeDeal, PrismMetadataListParamsObjectTypeIdentity, PrismMetadataListParamsObjectTypeAIChatThread, PrismMetadataListParamsObjectTypeAIChatMessage, PrismMetadataListParamsObjectTypeDocument, PrismMetadataListParamsObjectTypeAction, PrismMetadataListParamsObjectTypeEvent, PrismMetadataListParamsObjectTypeOrganization, PrismMetadataListParamsObjectTypeContact:
		return true
	}
	return false
}
