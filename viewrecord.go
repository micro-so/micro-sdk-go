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

// ViewRecordService contains methods and other services that help with interacting
// with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewViewRecordService] method instead.
type ViewRecordService struct {
	Options []option.RequestOption
}

// NewViewRecordService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewViewRecordService(opts ...option.RequestOption) (r *ViewRecordService) {
	r = &ViewRecordService{}
	r.Options = opts
	return
}

// List records selected by a view (filters and sorts applied; pinned record_order
// overlaid first)
func (r *ViewRecordService) List(ctx context.Context, objectType ViewRecordListParamsObjectType, viewID string, params ViewRecordListParams, opts ...option.RequestOption) (res *ViewRecordListResponse, err error) {
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
	if viewID == "" {
		err = errors.New("missing required viewId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/views/%s/records", params.TeamID, objectType, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Pin a record to the view (append to record_order)
func (r *ViewRecordService) Pin(ctx context.Context, objectType ViewRecordPinParamsObjectType, viewID string, objectID string, params ViewRecordPinParams, opts ...option.RequestOption) (err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
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
	if viewID == "" {
		err = errors.New("missing required viewId parameter")
		return err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/views/%s/records/%s", params.TeamID, objectType, viewID, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Bulk reorder pinned records
func (r *ViewRecordService) Reorder(ctx context.Context, objectType ViewRecordReorderParamsObjectType, viewID string, params ViewRecordReorderParams, opts ...option.RequestOption) (err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
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
	if viewID == "" {
		err = errors.New("missing required viewId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/views/%s/records", params.TeamID, objectType, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return err
}

// Unpin a record from the view
func (r *ViewRecordService) Unpin(ctx context.Context, objectType ViewRecordUnpinParamsObjectType, viewID string, objectID string, body ViewRecordUnpinParams, opts ...option.RequestOption) (err error) {
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
	if viewID == "" {
		err = errors.New("missing required viewId parameter")
		return err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/views/%s/records/%s", body.TeamID, objectType, viewID, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ViewRecordListResponse struct {
	Data []map[string]interface{} `json:"data" api:"required"`
	// True if more records exist beyond this page.
	HasMore bool `json:"has_more" api:"required"`
	// Opaque cursor for the next page; null when `has_more` is false.
	NextCursor string                     `json:"next_cursor" api:"nullable"`
	JSON       viewRecordListResponseJSON `json:"-"`
}

// viewRecordListResponseJSON contains the JSON metadata for the struct
// [ViewRecordListResponse]
type viewRecordListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ViewRecordListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewRecordListResponseJSON) RawJSON() string {
	return r.raw
}

type ViewRecordListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Opaque cursor from a previous response's `next_cursor`. Pass it back unchanged
	// to fetch the next page. When set, `page` and `limit` are derived from the
	// cursor.
	Cursor param.Field[string] `query:"cursor"`
	Limit  param.Field[int64]  `query:"limit"`
	// Page number (1-based). Prefer `cursor`.
	Page param.Field[int64] `query:"page"`
}

// URLQuery serializes [ViewRecordListParams]'s query parameters as `url.Values`.
func (r ViewRecordListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ViewRecordListParamsObjectType string

const (
	ViewRecordListParamsObjectTypeComment      ViewRecordListParamsObjectType = "comment"
	ViewRecordListParamsObjectTypeAction       ViewRecordListParamsObjectType = "action"
	ViewRecordListParamsObjectTypeDeal         ViewRecordListParamsObjectType = "deal"
	ViewRecordListParamsObjectTypeEngagement   ViewRecordListParamsObjectType = "engagement"
	ViewRecordListParamsObjectTypeDocument     ViewRecordListParamsObjectType = "document"
	ViewRecordListParamsObjectTypeEvent        ViewRecordListParamsObjectType = "event"
	ViewRecordListParamsObjectTypeIdentity     ViewRecordListParamsObjectType = "identity"
	ViewRecordListParamsObjectTypeOrganization ViewRecordListParamsObjectType = "organization"
)

func (r ViewRecordListParamsObjectType) IsKnown() bool {
	switch r {
	case ViewRecordListParamsObjectTypeComment, ViewRecordListParamsObjectTypeAction, ViewRecordListParamsObjectTypeDeal, ViewRecordListParamsObjectTypeEngagement, ViewRecordListParamsObjectTypeDocument, ViewRecordListParamsObjectTypeEvent, ViewRecordListParamsObjectTypeIdentity, ViewRecordListParamsObjectTypeOrganization:
		return true
	}
	return false
}

type ViewRecordPinParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type ViewRecordPinParamsObjectType string

const (
	ViewRecordPinParamsObjectTypeComment      ViewRecordPinParamsObjectType = "comment"
	ViewRecordPinParamsObjectTypeAction       ViewRecordPinParamsObjectType = "action"
	ViewRecordPinParamsObjectTypeDeal         ViewRecordPinParamsObjectType = "deal"
	ViewRecordPinParamsObjectTypeEngagement   ViewRecordPinParamsObjectType = "engagement"
	ViewRecordPinParamsObjectTypeDocument     ViewRecordPinParamsObjectType = "document"
	ViewRecordPinParamsObjectTypeEvent        ViewRecordPinParamsObjectType = "event"
	ViewRecordPinParamsObjectTypeIdentity     ViewRecordPinParamsObjectType = "identity"
	ViewRecordPinParamsObjectTypeOrganization ViewRecordPinParamsObjectType = "organization"
)

func (r ViewRecordPinParamsObjectType) IsKnown() bool {
	switch r {
	case ViewRecordPinParamsObjectTypeComment, ViewRecordPinParamsObjectTypeAction, ViewRecordPinParamsObjectTypeDeal, ViewRecordPinParamsObjectTypeEngagement, ViewRecordPinParamsObjectTypeDocument, ViewRecordPinParamsObjectTypeEvent, ViewRecordPinParamsObjectTypeIdentity, ViewRecordPinParamsObjectTypeOrganization:
		return true
	}
	return false
}

type ViewRecordReorderParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]   `path:"teamId" api:"required" format:"uuid"`
	ObjectIDs      param.Field[[]string] `json:"object_ids" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string]   `header:"Idempotency-Key"`
}

func (r ViewRecordReorderParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ViewRecordReorderParamsObjectType string

const (
	ViewRecordReorderParamsObjectTypeComment      ViewRecordReorderParamsObjectType = "comment"
	ViewRecordReorderParamsObjectTypeAction       ViewRecordReorderParamsObjectType = "action"
	ViewRecordReorderParamsObjectTypeDeal         ViewRecordReorderParamsObjectType = "deal"
	ViewRecordReorderParamsObjectTypeEngagement   ViewRecordReorderParamsObjectType = "engagement"
	ViewRecordReorderParamsObjectTypeDocument     ViewRecordReorderParamsObjectType = "document"
	ViewRecordReorderParamsObjectTypeEvent        ViewRecordReorderParamsObjectType = "event"
	ViewRecordReorderParamsObjectTypeIdentity     ViewRecordReorderParamsObjectType = "identity"
	ViewRecordReorderParamsObjectTypeOrganization ViewRecordReorderParamsObjectType = "organization"
)

func (r ViewRecordReorderParamsObjectType) IsKnown() bool {
	switch r {
	case ViewRecordReorderParamsObjectTypeComment, ViewRecordReorderParamsObjectTypeAction, ViewRecordReorderParamsObjectTypeDeal, ViewRecordReorderParamsObjectTypeEngagement, ViewRecordReorderParamsObjectTypeDocument, ViewRecordReorderParamsObjectTypeEvent, ViewRecordReorderParamsObjectTypeIdentity, ViewRecordReorderParamsObjectTypeOrganization:
		return true
	}
	return false
}

type ViewRecordUnpinParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type ViewRecordUnpinParamsObjectType string

const (
	ViewRecordUnpinParamsObjectTypeComment      ViewRecordUnpinParamsObjectType = "comment"
	ViewRecordUnpinParamsObjectTypeAction       ViewRecordUnpinParamsObjectType = "action"
	ViewRecordUnpinParamsObjectTypeDeal         ViewRecordUnpinParamsObjectType = "deal"
	ViewRecordUnpinParamsObjectTypeEngagement   ViewRecordUnpinParamsObjectType = "engagement"
	ViewRecordUnpinParamsObjectTypeDocument     ViewRecordUnpinParamsObjectType = "document"
	ViewRecordUnpinParamsObjectTypeEvent        ViewRecordUnpinParamsObjectType = "event"
	ViewRecordUnpinParamsObjectTypeIdentity     ViewRecordUnpinParamsObjectType = "identity"
	ViewRecordUnpinParamsObjectTypeOrganization ViewRecordUnpinParamsObjectType = "organization"
)

func (r ViewRecordUnpinParamsObjectType) IsKnown() bool {
	switch r {
	case ViewRecordUnpinParamsObjectTypeComment, ViewRecordUnpinParamsObjectTypeAction, ViewRecordUnpinParamsObjectTypeDeal, ViewRecordUnpinParamsObjectTypeEngagement, ViewRecordUnpinParamsObjectTypeDocument, ViewRecordUnpinParamsObjectTypeEvent, ViewRecordUnpinParamsObjectTypeIdentity, ViewRecordUnpinParamsObjectTypeOrganization:
		return true
	}
	return false
}
