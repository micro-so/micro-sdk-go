// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/micro-go/internal/apiquery"
	"github.com/stainless-sdks/micro-go/internal/param"
	"github.com/stainless-sdks/micro-go/internal/requestconfig"
	"github.com/stainless-sdks/micro-go/option"
)

// The Prism query engine provides generic read/write access to any object type
// using a single unified API surface.
//
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
func (r *PrismMetadataService) Properties(ctx context.Context, objectType ObjectType, params PrismMetadataPropertiesParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("v2/prism/metadata/properties/%s/%v", params.TeamID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return err
}

type PrismMetadataPropertiesParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID   param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	Autofill param.Field[bool]   `query:"autofill"`
	CRMID    param.Field[string] `query:"crmId" format:"uuid"`
	Term     param.Field[string] `query:"term"`
}

// URLQuery serializes [PrismMetadataPropertiesParams]'s query parameters as
// `url.Values`.
func (r PrismMetadataPropertiesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
