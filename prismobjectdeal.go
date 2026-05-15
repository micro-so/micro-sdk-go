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

// PrismObjectDealService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectDealService] method instead.
type PrismObjectDealService struct {
	Options []option.RequestOption
	Grant   *PrismObjectDealGrantService
}

// NewPrismObjectDealService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrismObjectDealService(opts ...option.RequestOption) (r *PrismObjectDealService) {
	r = &PrismObjectDealService{}
	r.Options = opts
	r.Grant = NewPrismObjectDealGrantService(opts...)
	return
}

// Create object
func (r *PrismObjectDealService) New(ctx context.Context, params PrismObjectDealNewParams, opts ...option.RequestOption) (res *PrismObjectDealNewResponse, err error) {
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
func (r *PrismObjectDealService) Update(ctx context.Context, dealID string, params PrismObjectDealUpdateParams, opts ...option.RequestOption) (res *PrismObjectDealUpdateResponse, err error) {
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

// Delete object
func (r *PrismObjectDealService) Delete(ctx context.Context, dealID string, body PrismObjectDealDeleteParams, opts ...option.RequestOption) (err error) {
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
func (r *PrismObjectDealService) BulkNew(ctx context.Context, params PrismObjectDealBulkNewParams, opts ...option.RequestOption) (res *PrismObjectDealBulkNewResponse, err error) {
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

// Duplicate object
func (r *PrismObjectDealService) Duplicate(ctx context.Context, dealID string, body PrismObjectDealDuplicateParams, opts ...option.RequestOption) (res *PrismObjectDealDuplicateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.TeamID, precfg.TeamID)
	if body.TeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return nil, err
	}
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/deal/%s/duplicate", body.TeamID, dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectDealService) Get(ctx context.Context, dealID string, query PrismObjectDealGetParams, opts ...option.RequestOption) (res *PrismObjectDealGetResponse, err error) {
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
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/deal/%s", query.TeamID, dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Query
func (r *PrismObjectDealService) Query(ctx context.Context, params PrismObjectDealQueryParams, opts ...option.RequestOption) (res *PrismObjectDealQueryResponse, err error) {
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

// Restore object
func (r *PrismObjectDealService) Restore(ctx context.Context, dealID string, body PrismObjectDealRestoreParams, opts ...option.RequestOption) (res *PrismObjectDealRestoreResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.TeamID, precfg.TeamID)
	if body.TeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return nil, err
	}
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/deal/%s/restore", body.TeamID, dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealNewResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}         `json:"default"`
	List    interface{}                    `json:"list"`
	JSON    prismObjectDealNewResponseJSON `json:"-"`
}

// prismObjectDealNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealNewResponse]
type prismObjectDealNewResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealNewResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealUpdateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}            `json:"default"`
	List    interface{}                       `json:"list"`
	JSON    prismObjectDealUpdateResponseJSON `json:"-"`
}

// prismObjectDealUpdateResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealUpdateResponse]
type prismObjectDealUpdateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealBulkNewResponse struct {
	Results []PrismObjectDealBulkNewResponseResult `json:"results"`
	Status  PrismObjectDealBulkNewResponseStatus   `json:"status"`
	Summary PrismObjectDealBulkNewResponseSummary  `json:"summary"`
	JSON    prismObjectDealBulkNewResponseJSON     `json:"-"`
}

// prismObjectDealBulkNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealBulkNewResponse]
type prismObjectDealBulkNewResponseJSON struct {
	Results     apijson.Field
	Status      apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealBulkNewResponseResult struct {
	ID       string                                   `json:"id" api:"nullable" format:"uuid"`
	Created  bool                                     `json:"created"`
	Error    string                                   `json:"error"`
	Existing bool                                     `json:"existing"`
	JSON     prismObjectDealBulkNewResponseResultJSON `json:"-"`
}

// prismObjectDealBulkNewResponseResultJSON contains the JSON metadata for the
// struct [PrismObjectDealBulkNewResponseResult]
type prismObjectDealBulkNewResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealBulkNewResponseStatus string

const (
	PrismObjectDealBulkNewResponseStatusComplete PrismObjectDealBulkNewResponseStatus = "complete"
)

func (r PrismObjectDealBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectDealBulkNewResponseStatusComplete:
		return true
	}
	return false
}

type PrismObjectDealBulkNewResponseSummary struct {
	Created  int64                                     `json:"created"`
	Errors   int64                                     `json:"errors"`
	Existing int64                                     `json:"existing"`
	Total    int64                                     `json:"total"`
	JSON     prismObjectDealBulkNewResponseSummaryJSON `json:"-"`
}

// prismObjectDealBulkNewResponseSummaryJSON contains the JSON metadata for the
// struct [PrismObjectDealBulkNewResponseSummary]
type prismObjectDealBulkNewResponseSummaryJSON struct {
	Created     apijson.Field
	Errors      apijson.Field
	Existing    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkNewResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkNewResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealDuplicateResponse struct {
	ID   string                               `json:"id" format:"uuid"`
	JSON prismObjectDealDuplicateResponseJSON `json:"-"`
}

// prismObjectDealDuplicateResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealDuplicateResponse]
type prismObjectDealDuplicateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealDuplicateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealDuplicateResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealGetResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}         `json:"default"`
	List    interface{}                    `json:"list"`
	JSON    prismObjectDealGetResponseJSON `json:"-"`
}

// prismObjectDealGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealGetResponse]
type prismObjectDealGetResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealQueryResponse struct {
	Data []PrismObjectDealQueryResponseData `json:"data" api:"required"`
	// True when the page returned the maximum number of rows; another page may exist.
	HasMore bool                             `json:"has_more"`
	JSON    prismObjectDealQueryResponseJSON `json:"-"`
}

// prismObjectDealQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealQueryResponse]
type prismObjectDealQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealQueryResponseData struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}               `json:"default"`
	List    interface{}                          `json:"list"`
	JSON    prismObjectDealQueryResponseDataJSON `json:"-"`
}

// prismObjectDealQueryResponseDataJSON contains the JSON metadata for the struct
// [PrismObjectDealQueryResponseData]
type prismObjectDealQueryResponseDataJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealRestoreResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}             `json:"default"`
	List    interface{}                        `json:"list"`
	JSON    prismObjectDealRestoreResponseJSON `json:"-"`
}

// prismObjectDealRestoreResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealRestoreResponse]
type prismObjectDealRestoreResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealRestoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealRestoreResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismObjectDealNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectDealUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismObjectDealUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectDealDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectDealBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects param.Field[[]PrismObjectPropertiesParam]        `json:"objects" api:"required"`
	Options param.Field[PrismObjectDealBulkNewParamsOptions] `json:"options"`
}

func (r PrismObjectDealBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDealBulkNewParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
	// App/CRM ID for context (optional)
	ListID param.Field[string] `json:"list_id" format:"uuid"`
}

func (r PrismObjectDealBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDealDuplicateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectDealGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectDealQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                            `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[PrismObjectDealQueryParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[PrismObjectDealQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                          `json:"boxes"`
	Deleted param.Field[bool]                              `json:"deleted"`
	Sources param.Field[[]string]                          `json:"sources" format:"uuid"`
}

func (r PrismObjectDealQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDealQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectDealQueryParamsQueryCombinator] `json:"combinator"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectDealQueryParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                   `json:"limit"`
	ListID param.Field[string]                                                  `json:"list_id" format:"uuid"`
	Page   param.Field[int64]                                                   `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]PrismObjectDealQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectDealQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectDealQueryParamsQueryCombinator string

const (
	PrismObjectDealQueryParamsQueryCombinatorAnd PrismObjectDealQueryParamsQueryCombinator = "AND"
	PrismObjectDealQueryParamsQueryCombinatorOr  PrismObjectDealQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectDealQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectDealQueryParamsQueryCombinatorAnd, PrismObjectDealQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

type PrismObjectDealQueryParamsQueryFilter struct {
	NotEquals       param.Field[interface{}] `json:"!="`
	Less            param.Field[string]      `json:"<"`
	LessOrEquals    param.Field[string]      `json:"<="`
	Equals          param.Field[interface{}] `json:"="`
	Greater         param.Field[string]      `json:">"`
	GreaterOrEquals param.Field[string]      `json:">="`
	BeginsWith      param.Field[string]      `json:"begins_with"`
	EndsWith        param.Field[string]      `json:"ends_with"`
	Exists          param.Field[bool]        `json:"exists"`
	In              param.Field[interface{}] `json:"in"`
	LikeRegex       param.Field[string]      `json:"like_regex"`
	NotContains     param.Field[string]      `json:"not_contains"`
	NotExists       param.Field[bool]        `json:"not_exists"`
	NotIn           param.Field[interface{}] `json:"not_in"`
}

func (r PrismObjectDealQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilter) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

// Satisfied by [PrismObjectDealQueryParamsQueryFilterPrismQueryFilterEq],
// [PrismObjectDealQueryParamsQueryFilterPrismQueryFilterNe],
// [PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLt],
// [PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGt],
// [PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLte],
// [PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGte],
// [PrismObjectDealQueryParamsQueryFilterLikeRegex],
// [PrismObjectDealQueryParamsQueryFilterBeginsWith],
// [PrismObjectDealQueryParamsQueryFilterEndsWith],
// [PrismObjectDealQueryParamsQueryFilterNotContains],
// [PrismObjectDealQueryParamsQueryFilterExists],
// [PrismObjectDealQueryParamsQueryFilterNotExists],
// [PrismObjectDealQueryParamsQueryFilterIn],
// [PrismObjectDealQueryParamsQueryFilterNotIn],
// [PrismObjectDealQueryParamsQueryFilter].
type PrismObjectDealQueryParamsQueryFilterUnion interface {
	implementsPrismObjectDealQueryParamsQueryFilterUnion()
}

type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterEq struct {
	Equals param.Field[PrismObjectDealQueryParamsQueryFilterPrismQueryFilterEqUnion] `json:"=" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterEq) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterEq) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterEqUnion interface {
	ImplementsPrismObjectDealQueryParamsQueryFilterPrismQueryFilterEqUnion()
}

type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterNe struct {
	NotEquals param.Field[PrismObjectDealQueryParamsQueryFilterPrismQueryFilterNeUnion] `json:"!=" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterNe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterNe) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterNeUnion interface {
	ImplementsPrismObjectDealQueryParamsQueryFilterPrismQueryFilterNeUnion()
}

type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLt struct {
	Less param.Field[string] `json:"<" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLt) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGt struct {
	Greater param.Field[string] `json:">" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGt) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLte struct {
	LessOrEquals param.Field[string] `json:"<=" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLte) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGte struct {
	GreaterOrEquals param.Field[string] `json:">=" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGte) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterLikeRegex struct {
	LikeRegex param.Field[string] `json:"like_regex" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterLikeRegex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterLikeRegex) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterBeginsWith struct {
	BeginsWith param.Field[string] `json:"begins_with" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterBeginsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterBeginsWith) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterEndsWith struct {
	EndsWith param.Field[string] `json:"ends_with" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterEndsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterEndsWith) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterNotContains struct {
	NotContains param.Field[string] `json:"not_contains" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterNotContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterNotContains) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterExists struct {
	Exists param.Field[bool] `json:"exists" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterExists) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterNotExists struct {
	NotExists param.Field[bool] `json:"not_exists" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterNotExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterNotExists) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterIn struct {
	In param.Field[[]string] `json:"in" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterIn) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterNotIn struct {
	NotIn param.Field[[]string] `json:"not_in" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterNotIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterNotIn) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQuerySort string

const (
	PrismObjectDealQueryParamsQuerySortAsc  PrismObjectDealQueryParamsQuerySort = "asc"
	PrismObjectDealQueryParamsQuerySortDesc PrismObjectDealQueryParamsQuerySort = "desc"
)

func (r PrismObjectDealQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectDealQueryParamsQuerySortAsc, PrismObjectDealQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectDealQueryParamsIDArray].
type PrismObjectDealQueryParamsIDUnion interface {
	ImplementsPrismObjectDealQueryParamsIDUnion()
}

type PrismObjectDealQueryParamsIDArray []string

func (r PrismObjectDealQueryParamsIDArray) ImplementsPrismObjectDealQueryParamsIDUnion() {}

type PrismObjectDealRestoreParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
