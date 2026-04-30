// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/micro-go/internal/apijson"
	"github.com/stainless-sdks/micro-go/internal/param"
	"github.com/stainless-sdks/micro-go/internal/requestconfig"
	"github.com/stainless-sdks/micro-go/option"
)

// DealService contains methods and other services that help with interacting with
// the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDealService] method instead.
type DealService struct {
	Options []option.RequestOption
}

// NewDealService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDealService(opts ...option.RequestOption) (r *DealService) {
	r = &DealService{}
	r.Options = opts
	return
}

// Create object
func (r *DealService) New(ctx context.Context, params DealNewParams, opts ...option.RequestOption) (res *PrismObjectProperties, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/deal", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch object
func (r *DealService) Update(ctx context.Context, dealID string, params DealUpdateParams, opts ...option.RequestOption) (res *PrismObjectProperties, err error) {
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
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/deal/%s", params.TeamID, dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Query v2
func (r *DealService) List(ctx context.Context, params DealListParams, opts ...option.RequestOption) (res *DealListResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/query/%s/deal", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Delete object
func (r *DealService) Delete(ctx context.Context, dealID string, body DealDeleteParams, opts ...option.RequestOption) (err error) {
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
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/deal/%s", body.TeamID, dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: <100 records sync (immediate response), >=100 records
// async (S3/Lambda with WebSocket progress)
func (r *DealService) Import(ctx context.Context, params DealImportParams, opts ...option.RequestOption) (res *DealImportResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/deal/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type DealListResponse struct {
	Data  []interface{}        `json:"data"`
	Total int64                `json:"total"`
	JSON  dealListResponseJSON `json:"-"`
}

// dealListResponseJSON contains the JSON metadata for the struct
// [DealListResponse]
type dealListResponseJSON struct {
	Data        apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DealListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dealListResponseJSON) RawJSON() string {
	return r.raw
}

type DealImportResponse struct {
	Results []DealImportResponseResult `json:"results"`
	Status  DealImportResponseStatus   `json:"status"`
	Summary DealImportResponseSummary  `json:"summary"`
	JSON    dealImportResponseJSON     `json:"-"`
}

// dealImportResponseJSON contains the JSON metadata for the struct
// [DealImportResponse]
type dealImportResponseJSON struct {
	Results     apijson.Field
	Status      apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DealImportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dealImportResponseJSON) RawJSON() string {
	return r.raw
}

type DealImportResponseResult struct {
	ID       string                       `json:"id" api:"nullable" format:"uuid"`
	Created  bool                         `json:"created"`
	Error    string                       `json:"error"`
	Existing bool                         `json:"existing"`
	JSON     dealImportResponseResultJSON `json:"-"`
}

// dealImportResponseResultJSON contains the JSON metadata for the struct
// [DealImportResponseResult]
type dealImportResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DealImportResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dealImportResponseResultJSON) RawJSON() string {
	return r.raw
}

type DealImportResponseStatus string

const (
	DealImportResponseStatusComplete DealImportResponseStatus = "complete"
)

func (r DealImportResponseStatus) IsKnown() bool {
	switch r {
	case DealImportResponseStatusComplete:
		return true
	}
	return false
}

type DealImportResponseSummary struct {
	Created  int64                         `json:"created"`
	Errors   int64                         `json:"errors"`
	Existing int64                         `json:"existing"`
	Total    int64                         `json:"total"`
	JSON     dealImportResponseSummaryJSON `json:"-"`
}

// dealImportResponseSummaryJSON contains the JSON metadata for the struct
// [DealImportResponseSummary]
type dealImportResponseSummaryJSON struct {
	Created     apijson.Field
	Errors      apijson.Field
	Existing    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DealImportResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dealImportResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type DealNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r DealNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type DealUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r DealUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type DealListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[DealListParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[DealListParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]              `json:"boxes"`
	Deleted param.Field[bool]                  `json:"deleted"`
	Sources param.Field[[]string]              `json:"sources" format:"uuid"`
}

func (r DealListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DealListParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[DealListParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                        `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values must be option slugs
	Filter param.Field[[]map[string]map[string]DealListParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                  `json:"limit"`
	Page   param.Field[int64]                                                  `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]DealListParamsQuerySort] `json:"sort"`
}

func (r DealListParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type DealListParamsQueryCombinator string

const (
	DealListParamsQueryCombinatorAnd DealListParamsQueryCombinator = "AND"
	DealListParamsQueryCombinatorOr  DealListParamsQueryCombinator = "OR"
)

func (r DealListParamsQueryCombinator) IsKnown() bool {
	switch r {
	case DealListParamsQueryCombinatorAnd, DealListParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [DealListParamsQueryFilterArray].
type DealListParamsQueryFilterUnion interface {
	ImplementsDealListParamsQueryFilterUnion()
}

type DealListParamsQueryFilterArray []string

func (r DealListParamsQueryFilterArray) ImplementsDealListParamsQueryFilterUnion() {}

type DealListParamsQuerySort string

const (
	DealListParamsQuerySortAsc  DealListParamsQuerySort = "asc"
	DealListParamsQuerySortDesc DealListParamsQuerySort = "desc"
)

func (r DealListParamsQuerySort) IsKnown() bool {
	switch r {
	case DealListParamsQuerySortAsc, DealListParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [DealListParamsIDArray].
type DealListParamsIDUnion interface {
	ImplementsDealListParamsIDUnion()
}

type DealListParamsIDArray []string

func (r DealListParamsIDArray) ImplementsDealListParamsIDUnion() {}

type DealDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type DealImportParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects param.Field[[]PrismObjectPropertiesParam] `json:"objects" api:"required"`
	Options param.Field[DealImportParamsOptions]      `json:"options"`
}

func (r DealImportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DealImportParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// App/CRM ID for context (optional)
	CRMID param.Field[string] `json:"crm_id" format:"uuid"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
}

func (r DealImportParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
