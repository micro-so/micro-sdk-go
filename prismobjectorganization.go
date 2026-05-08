// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/micro-so/micro-sdk-go/internal/apijson"
	"github.com/micro-so/micro-sdk-go/internal/param"
	"github.com/micro-so/micro-sdk-go/internal/requestconfig"
	"github.com/micro-so/micro-sdk-go/option"
)

// PrismObjectOrganizationService contains methods and other services that help
// with interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectOrganizationService] method instead.
type PrismObjectOrganizationService struct {
	Options []option.RequestOption
}

// NewPrismObjectOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectOrganizationService(opts ...option.RequestOption) (r *PrismObjectOrganizationService) {
	r = &PrismObjectOrganizationService{}
	r.Options = opts
	return
}

// Create object
func (r *PrismObjectOrganizationService) New(ctx context.Context, params PrismObjectOrganizationNewParams, opts ...option.RequestOption) (res *PrismObjectOrganizationNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/organization", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: <100 records sync (immediate response), >=100 records
// async (S3/Lambda with WebSocket progress)
func (r *PrismObjectOrganizationService) BulkNew(ctx context.Context, params PrismObjectOrganizationBulkNewParams, opts ...option.RequestOption) (res *PrismObjectOrganizationBulkNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/organization/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Query v2
func (r *PrismObjectOrganizationService) Query(ctx context.Context, params PrismObjectOrganizationQueryParams, opts ...option.RequestOption) (res *PrismObjectOrganizationQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/query/%s/organization", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectOrganizationNewResponse struct {
	ID  string      `json:"id" api:"required" format:"uuid"`
	CRM interface{} `json:"crm"`
	// Properties keyed by property slug.
	Default  map[string]interface{}                 `json:"default"`
	Extended interface{}                            `json:"extended"`
	JSON     prismObjectOrganizationNewResponseJSON `json:"-"`
}

// prismObjectOrganizationNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectOrganizationNewResponse]
type prismObjectOrganizationNewResponseJSON struct {
	ID          apijson.Field
	CRM         apijson.Field
	Default     apijson.Field
	Extended    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationBulkNewResponse struct {
	Results []PrismObjectOrganizationBulkNewResponseResult `json:"results"`
	Status  PrismObjectOrganizationBulkNewResponseStatus   `json:"status"`
	Summary PrismObjectOrganizationBulkNewResponseSummary  `json:"summary"`
	JSON    prismObjectOrganizationBulkNewResponseJSON     `json:"-"`
}

// prismObjectOrganizationBulkNewResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationBulkNewResponse]
type prismObjectOrganizationBulkNewResponseJSON struct {
	Results     apijson.Field
	Status      apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationBulkNewResponseResult struct {
	ID       string                                           `json:"id" api:"nullable" format:"uuid"`
	Created  bool                                             `json:"created"`
	Error    string                                           `json:"error"`
	Existing bool                                             `json:"existing"`
	JSON     prismObjectOrganizationBulkNewResponseResultJSON `json:"-"`
}

// prismObjectOrganizationBulkNewResponseResultJSON contains the JSON metadata for
// the struct [PrismObjectOrganizationBulkNewResponseResult]
type prismObjectOrganizationBulkNewResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationBulkNewResponseStatus string

const (
	PrismObjectOrganizationBulkNewResponseStatusComplete PrismObjectOrganizationBulkNewResponseStatus = "complete"
)

func (r PrismObjectOrganizationBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectOrganizationBulkNewResponseStatusComplete:
		return true
	}
	return false
}

type PrismObjectOrganizationBulkNewResponseSummary struct {
	Created  int64                                             `json:"created"`
	Errors   int64                                             `json:"errors"`
	Existing int64                                             `json:"existing"`
	Total    int64                                             `json:"total"`
	JSON     prismObjectOrganizationBulkNewResponseSummaryJSON `json:"-"`
}

// prismObjectOrganizationBulkNewResponseSummaryJSON contains the JSON metadata for
// the struct [PrismObjectOrganizationBulkNewResponseSummary]
type prismObjectOrganizationBulkNewResponseSummaryJSON struct {
	Created     apijson.Field
	Errors      apijson.Field
	Existing    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkNewResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkNewResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationQueryResponse struct {
	Data []PrismObjectOrganizationQueryResponseData `json:"data" api:"required"`
	// True when the page returned the maximum number of rows; another page may exist.
	HasMore bool                                     `json:"has_more"`
	JSON    prismObjectOrganizationQueryResponseJSON `json:"-"`
}

// prismObjectOrganizationQueryResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationQueryResponse]
type prismObjectOrganizationQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectOrganizationQueryResponseData struct {
	ID  string      `json:"id" api:"required" format:"uuid"`
	CRM interface{} `json:"crm"`
	// Properties keyed by property slug.
	Default  map[string]interface{}                       `json:"default"`
	Extended interface{}                                  `json:"extended"`
	JSON     prismObjectOrganizationQueryResponseDataJSON `json:"-"`
}

// prismObjectOrganizationQueryResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationQueryResponseData]
type prismObjectOrganizationQueryResponseDataJSON struct {
	ID          apijson.Field
	CRM         apijson.Field
	Default     apijson.Field
	Extended    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismObjectOrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectOrganizationBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects param.Field[[]PrismObjectPropertiesParam]                `json:"objects" api:"required"`
	Options param.Field[PrismObjectOrganizationBulkNewParamsOptions] `json:"options"`
}

func (r PrismObjectOrganizationBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectOrganizationBulkNewParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// App/CRM ID for context (optional)
	CRMID param.Field[string] `json:"crm_id" format:"uuid"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
}

func (r PrismObjectOrganizationBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectOrganizationQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                                    `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[PrismObjectOrganizationQueryParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[PrismObjectOrganizationQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                                  `json:"boxes"`
	Deleted param.Field[bool]                                      `json:"deleted"`
	Sources param.Field[[]string]                                  `json:"sources" format:"uuid"`
}

func (r PrismObjectOrganizationQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectOrganizationQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectOrganizationQueryParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                                            `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]map[string]PrismObjectOrganizationQueryParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                                      `json:"limit"`
	Page   param.Field[int64]                                                                      `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]PrismObjectOrganizationQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectOrganizationQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectOrganizationQueryParamsQueryCombinator string

const (
	PrismObjectOrganizationQueryParamsQueryCombinatorAnd PrismObjectOrganizationQueryParamsQueryCombinator = "AND"
	PrismObjectOrganizationQueryParamsQueryCombinatorOr  PrismObjectOrganizationQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectOrganizationQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectOrganizationQueryParamsQueryCombinatorAnd, PrismObjectOrganizationQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectOrganizationQueryParamsQueryFilterArray].
type PrismObjectOrganizationQueryParamsQueryFilterUnion interface {
	ImplementsPrismObjectOrganizationQueryParamsQueryFilterUnion()
}

type PrismObjectOrganizationQueryParamsQueryFilterArray []string

func (r PrismObjectOrganizationQueryParamsQueryFilterArray) ImplementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQuerySort string

const (
	PrismObjectOrganizationQueryParamsQuerySortAsc  PrismObjectOrganizationQueryParamsQuerySort = "asc"
	PrismObjectOrganizationQueryParamsQuerySortDesc PrismObjectOrganizationQueryParamsQuerySort = "desc"
)

func (r PrismObjectOrganizationQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectOrganizationQueryParamsQuerySortAsc, PrismObjectOrganizationQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectOrganizationQueryParamsIDArray].
type PrismObjectOrganizationQueryParamsIDUnion interface {
	ImplementsPrismObjectOrganizationQueryParamsIDUnion()
}

type PrismObjectOrganizationQueryParamsIDArray []string

func (r PrismObjectOrganizationQueryParamsIDArray) ImplementsPrismObjectOrganizationQueryParamsIDUnion() {
}
