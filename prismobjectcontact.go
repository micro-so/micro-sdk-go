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

// PrismObjectContactService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectContactService] method instead.
type PrismObjectContactService struct {
	Options []option.RequestOption
}

// NewPrismObjectContactService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectContactService(opts ...option.RequestOption) (r *PrismObjectContactService) {
	r = &PrismObjectContactService{}
	r.Options = opts
	return
}

// Create object
func (r *PrismObjectContactService) New(ctx context.Context, params PrismObjectContactNewParams, opts ...option.RequestOption) (res *PrismObjectContactNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/contact", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: <100 records sync (immediate response), >=100 records
// async (S3/Lambda with WebSocket progress)
func (r *PrismObjectContactService) BulkNew(ctx context.Context, params PrismObjectContactBulkNewParams, opts ...option.RequestOption) (res *PrismObjectContactBulkNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/contact/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Query v2
func (r *PrismObjectContactService) Query(ctx context.Context, params PrismObjectContactQueryParams, opts ...option.RequestOption) (res *PrismObjectContactQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/query/%s/contact", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectContactNewResponse struct {
	ID  string      `json:"id" api:"required" format:"uuid"`
	CRM interface{} `json:"crm"`
	// Properties keyed by property slug.
	Default  map[string]interface{}            `json:"default"`
	Extended interface{}                       `json:"extended"`
	JSON     prismObjectContactNewResponseJSON `json:"-"`
}

// prismObjectContactNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactNewResponse]
type prismObjectContactNewResponseJSON struct {
	ID          apijson.Field
	CRM         apijson.Field
	Default     apijson.Field
	Extended    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkNewResponse struct {
	Results []PrismObjectContactBulkNewResponseResult `json:"results"`
	Status  PrismObjectContactBulkNewResponseStatus   `json:"status"`
	Summary PrismObjectContactBulkNewResponseSummary  `json:"summary"`
	JSON    prismObjectContactBulkNewResponseJSON     `json:"-"`
}

// prismObjectContactBulkNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactBulkNewResponse]
type prismObjectContactBulkNewResponseJSON struct {
	Results     apijson.Field
	Status      apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkNewResponseResult struct {
	ID       string                                      `json:"id" api:"nullable" format:"uuid"`
	Created  bool                                        `json:"created"`
	Error    string                                      `json:"error"`
	Existing bool                                        `json:"existing"`
	JSON     prismObjectContactBulkNewResponseResultJSON `json:"-"`
}

// prismObjectContactBulkNewResponseResultJSON contains the JSON metadata for the
// struct [PrismObjectContactBulkNewResponseResult]
type prismObjectContactBulkNewResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkNewResponseStatus string

const (
	PrismObjectContactBulkNewResponseStatusComplete PrismObjectContactBulkNewResponseStatus = "complete"
)

func (r PrismObjectContactBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectContactBulkNewResponseStatusComplete:
		return true
	}
	return false
}

type PrismObjectContactBulkNewResponseSummary struct {
	Created  int64                                        `json:"created"`
	Errors   int64                                        `json:"errors"`
	Existing int64                                        `json:"existing"`
	Total    int64                                        `json:"total"`
	JSON     prismObjectContactBulkNewResponseSummaryJSON `json:"-"`
}

// prismObjectContactBulkNewResponseSummaryJSON contains the JSON metadata for the
// struct [PrismObjectContactBulkNewResponseSummary]
type prismObjectContactBulkNewResponseSummaryJSON struct {
	Created     apijson.Field
	Errors      apijson.Field
	Existing    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkNewResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkNewResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactQueryResponse struct {
	Data []PrismObjectContactQueryResponseData `json:"data" api:"required"`
	// True when the page returned the maximum number of rows; another page may exist.
	HasMore bool                                `json:"has_more"`
	JSON    prismObjectContactQueryResponseJSON `json:"-"`
}

// prismObjectContactQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactQueryResponse]
type prismObjectContactQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectContactQueryResponseData struct {
	ID  string      `json:"id" api:"required" format:"uuid"`
	CRM interface{} `json:"crm"`
	// Properties keyed by property slug.
	Default  map[string]interface{}                  `json:"default"`
	Extended interface{}                             `json:"extended"`
	JSON     prismObjectContactQueryResponseDataJSON `json:"-"`
}

// prismObjectContactQueryResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectContactQueryResponseData]
type prismObjectContactQueryResponseDataJSON struct {
	ID          apijson.Field
	CRM         apijson.Field
	Default     apijson.Field
	Extended    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismObjectContactNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectContactBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects param.Field[[]PrismObjectPropertiesParam]           `json:"objects" api:"required"`
	Options param.Field[PrismObjectContactBulkNewParamsOptions] `json:"options"`
}

func (r PrismObjectContactBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectContactBulkNewParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// App/CRM ID for context (optional)
	CRMID param.Field[string] `json:"crm_id" format:"uuid"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
}

func (r PrismObjectContactBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectContactQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                               `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[PrismObjectContactQueryParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[PrismObjectContactQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                             `json:"boxes"`
	Deleted param.Field[bool]                                 `json:"deleted"`
	Sources param.Field[[]string]                             `json:"sources" format:"uuid"`
}

func (r PrismObjectContactQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectContactQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectContactQueryParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                                       `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]map[string]PrismObjectContactQueryParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                                 `json:"limit"`
	Page   param.Field[int64]                                                                 `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]PrismObjectContactQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectContactQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectContactQueryParamsQueryCombinator string

const (
	PrismObjectContactQueryParamsQueryCombinatorAnd PrismObjectContactQueryParamsQueryCombinator = "AND"
	PrismObjectContactQueryParamsQueryCombinatorOr  PrismObjectContactQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectContactQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectContactQueryParamsQueryCombinatorAnd, PrismObjectContactQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectContactQueryParamsQueryFilterArray].
type PrismObjectContactQueryParamsQueryFilterUnion interface {
	ImplementsPrismObjectContactQueryParamsQueryFilterUnion()
}

type PrismObjectContactQueryParamsQueryFilterArray []string

func (r PrismObjectContactQueryParamsQueryFilterArray) ImplementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQuerySort string

const (
	PrismObjectContactQueryParamsQuerySortAsc  PrismObjectContactQueryParamsQuerySort = "asc"
	PrismObjectContactQueryParamsQuerySortDesc PrismObjectContactQueryParamsQuerySort = "desc"
)

func (r PrismObjectContactQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectContactQueryParamsQuerySortAsc, PrismObjectContactQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectContactQueryParamsIDArray].
type PrismObjectContactQueryParamsIDUnion interface {
	ImplementsPrismObjectContactQueryParamsIDUnion()
}

type PrismObjectContactQueryParamsIDArray []string

func (r PrismObjectContactQueryParamsIDArray) ImplementsPrismObjectContactQueryParamsIDUnion() {}
