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

// PrismObjectIdentityService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectIdentityService] method instead.
type PrismObjectIdentityService struct {
	Options []option.RequestOption
}

// NewPrismObjectIdentityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectIdentityService(opts ...option.RequestOption) (r *PrismObjectIdentityService) {
	r = &PrismObjectIdentityService{}
	r.Options = opts
	return
}

// Create object
func (r *PrismObjectIdentityService) New(ctx context.Context, params PrismObjectIdentityNewParams, opts ...option.RequestOption) (res *PrismObjectIdentityNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/identity", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch object
func (r *PrismObjectIdentityService) Update(ctx context.Context, identityID string, params PrismObjectIdentityUpdateParams, opts ...option.RequestOption) (res *PrismObjectIdentityUpdateResponse, err error) {
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
	if identityID == "" {
		err = errors.New("missing required identityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/identity/%s", params.TeamID, identityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Delete object
func (r *PrismObjectIdentityService) Delete(ctx context.Context, identityID string, body PrismObjectIdentityDeleteParams, opts ...option.RequestOption) (err error) {
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
	if identityID == "" {
		err = errors.New("missing required identityId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/identity/%s", body.TeamID, identityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: <100 records sync (immediate response), >=100 records
// async (S3/Lambda with WebSocket progress)
func (r *PrismObjectIdentityService) BulkNew(ctx context.Context, params PrismObjectIdentityBulkNewParams, opts ...option.RequestOption) (res *PrismObjectIdentityBulkNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/identity/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Duplicate object
func (r *PrismObjectIdentityService) Duplicate(ctx context.Context, identityID string, body PrismObjectIdentityDuplicateParams, opts ...option.RequestOption) (res *PrismObjectIdentityDuplicateResponse, err error) {
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
	if identityID == "" {
		err = errors.New("missing required identityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/identity/%s/duplicate", body.TeamID, identityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectIdentityService) Get(ctx context.Context, identityID string, query PrismObjectIdentityGetParams, opts ...option.RequestOption) (res *PrismObjectIdentityGetResponse, err error) {
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
	if identityID == "" {
		err = errors.New("missing required identityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/identity/%s", query.TeamID, identityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Query
func (r *PrismObjectIdentityService) Query(ctx context.Context, params PrismObjectIdentityQueryParams, opts ...option.RequestOption) (res *PrismObjectIdentityQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/query/%s/identity", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Restore object
func (r *PrismObjectIdentityService) Restore(ctx context.Context, identityID string, body PrismObjectIdentityRestoreParams, opts ...option.RequestOption) (res *PrismObjectIdentityRestoreResponse, err error) {
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
	if identityID == "" {
		err = errors.New("missing required identityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/identity/%s/restore", body.TeamID, identityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityNewResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}             `json:"default"`
	List    interface{}                        `json:"list"`
	JSON    prismObjectIdentityNewResponseJSON `json:"-"`
}

// prismObjectIdentityNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityNewResponse]
type prismObjectIdentityNewResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityNewResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityUpdateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                `json:"default"`
	List    interface{}                           `json:"list"`
	JSON    prismObjectIdentityUpdateResponseJSON `json:"-"`
}

// prismObjectIdentityUpdateResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityUpdateResponse]
type prismObjectIdentityUpdateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityBulkNewResponse struct {
	Results []PrismObjectIdentityBulkNewResponseResult `json:"results"`
	Status  PrismObjectIdentityBulkNewResponseStatus   `json:"status"`
	Summary PrismObjectIdentityBulkNewResponseSummary  `json:"summary"`
	JSON    prismObjectIdentityBulkNewResponseJSON     `json:"-"`
}

// prismObjectIdentityBulkNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityBulkNewResponse]
type prismObjectIdentityBulkNewResponseJSON struct {
	Results     apijson.Field
	Status      apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityBulkNewResponseResult struct {
	ID       string                                       `json:"id" api:"nullable" format:"uuid"`
	Created  bool                                         `json:"created"`
	Error    string                                       `json:"error"`
	Existing bool                                         `json:"existing"`
	JSON     prismObjectIdentityBulkNewResponseResultJSON `json:"-"`
}

// prismObjectIdentityBulkNewResponseResultJSON contains the JSON metadata for the
// struct [PrismObjectIdentityBulkNewResponseResult]
type prismObjectIdentityBulkNewResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityBulkNewResponseStatus string

const (
	PrismObjectIdentityBulkNewResponseStatusComplete PrismObjectIdentityBulkNewResponseStatus = "complete"
)

func (r PrismObjectIdentityBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectIdentityBulkNewResponseStatusComplete:
		return true
	}
	return false
}

type PrismObjectIdentityBulkNewResponseSummary struct {
	Created  int64                                         `json:"created"`
	Errors   int64                                         `json:"errors"`
	Existing int64                                         `json:"existing"`
	Total    int64                                         `json:"total"`
	JSON     prismObjectIdentityBulkNewResponseSummaryJSON `json:"-"`
}

// prismObjectIdentityBulkNewResponseSummaryJSON contains the JSON metadata for the
// struct [PrismObjectIdentityBulkNewResponseSummary]
type prismObjectIdentityBulkNewResponseSummaryJSON struct {
	Created     apijson.Field
	Errors      apijson.Field
	Existing    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkNewResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkNewResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityDuplicateResponse struct {
	ID   string                                   `json:"id" format:"uuid"`
	JSON prismObjectIdentityDuplicateResponseJSON `json:"-"`
}

// prismObjectIdentityDuplicateResponseJSON contains the JSON metadata for the
// struct [PrismObjectIdentityDuplicateResponse]
type prismObjectIdentityDuplicateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityDuplicateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityDuplicateResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityGetResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}             `json:"default"`
	List    interface{}                        `json:"list"`
	JSON    prismObjectIdentityGetResponseJSON `json:"-"`
}

// prismObjectIdentityGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityGetResponse]
type prismObjectIdentityGetResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityQueryResponse struct {
	Data []PrismObjectIdentityQueryResponseData `json:"data" api:"required"`
	// True when the page returned the maximum number of rows; another page may exist.
	HasMore bool                                 `json:"has_more"`
	JSON    prismObjectIdentityQueryResponseJSON `json:"-"`
}

// prismObjectIdentityQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityQueryResponse]
type prismObjectIdentityQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityQueryResponseData struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                   `json:"default"`
	List    interface{}                              `json:"list"`
	JSON    prismObjectIdentityQueryResponseDataJSON `json:"-"`
}

// prismObjectIdentityQueryResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectIdentityQueryResponseData]
type prismObjectIdentityQueryResponseDataJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityRestoreResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                 `json:"default"`
	List    interface{}                            `json:"list"`
	JSON    prismObjectIdentityRestoreResponseJSON `json:"-"`
}

// prismObjectIdentityRestoreResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityRestoreResponse]
type prismObjectIdentityRestoreResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityRestoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityRestoreResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismObjectIdentityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectIdentityUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismObjectIdentityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectIdentityDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectIdentityBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects param.Field[[]PrismObjectPropertiesParam]            `json:"objects" api:"required"`
	Options param.Field[PrismObjectIdentityBulkNewParamsOptions] `json:"options"`
}

func (r PrismObjectIdentityBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectIdentityBulkNewParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
	// App/CRM ID for context (optional)
	ListID param.Field[string] `json:"list_id" format:"uuid"`
}

func (r PrismObjectIdentityBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectIdentityDuplicateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectIdentityGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectIdentityQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                                `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[PrismObjectIdentityQueryParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[PrismObjectIdentityQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                              `json:"boxes"`
	Deleted param.Field[bool]                                  `json:"deleted"`
	Sources param.Field[[]string]                              `json:"sources" format:"uuid"`
}

func (r PrismObjectIdentityQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectIdentityQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectIdentityQueryParamsQueryCombinator] `json:"combinator"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectIdentityQueryParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                       `json:"limit"`
	ListID param.Field[string]                                                      `json:"list_id" format:"uuid"`
	Page   param.Field[int64]                                                       `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]PrismObjectIdentityQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectIdentityQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectIdentityQueryParamsQueryCombinator string

const (
	PrismObjectIdentityQueryParamsQueryCombinatorAnd PrismObjectIdentityQueryParamsQueryCombinator = "AND"
	PrismObjectIdentityQueryParamsQueryCombinatorOr  PrismObjectIdentityQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectIdentityQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectIdentityQueryParamsQueryCombinatorAnd, PrismObjectIdentityQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

type PrismObjectIdentityQueryParamsQueryFilter struct {
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

func (r PrismObjectIdentityQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilter) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

// Satisfied by [PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEq],
// [PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNe],
// [PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLt],
// [PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGt],
// [PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLte],
// [PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGte],
// [PrismObjectIdentityQueryParamsQueryFilterLikeRegex],
// [PrismObjectIdentityQueryParamsQueryFilterBeginsWith],
// [PrismObjectIdentityQueryParamsQueryFilterEndsWith],
// [PrismObjectIdentityQueryParamsQueryFilterNotContains],
// [PrismObjectIdentityQueryParamsQueryFilterExists],
// [PrismObjectIdentityQueryParamsQueryFilterNotExists],
// [PrismObjectIdentityQueryParamsQueryFilterIn],
// [PrismObjectIdentityQueryParamsQueryFilterNotIn],
// [PrismObjectIdentityQueryParamsQueryFilter].
type PrismObjectIdentityQueryParamsQueryFilterUnion interface {
	implementsPrismObjectIdentityQueryParamsQueryFilterUnion()
}

type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEq struct {
	Equals param.Field[PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEqUnion] `json:"=" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEq) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEq) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEqUnion interface {
	ImplementsPrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEqUnion()
}

type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNe struct {
	NotEquals param.Field[PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNeUnion] `json:"!=" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNe) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNeUnion interface {
	ImplementsPrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNeUnion()
}

type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLt struct {
	Less param.Field[string] `json:"<" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLt) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGt struct {
	Greater param.Field[string] `json:">" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGt) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLte struct {
	LessOrEquals param.Field[string] `json:"<=" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLte) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGte struct {
	GreaterOrEquals param.Field[string] `json:">=" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGte) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterLikeRegex struct {
	LikeRegex param.Field[string] `json:"like_regex" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterLikeRegex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterLikeRegex) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterBeginsWith struct {
	BeginsWith param.Field[string] `json:"begins_with" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterBeginsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterBeginsWith) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterEndsWith struct {
	EndsWith param.Field[string] `json:"ends_with" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterEndsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterEndsWith) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterNotContains struct {
	NotContains param.Field[string] `json:"not_contains" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterNotContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterNotContains) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterExists struct {
	Exists param.Field[bool] `json:"exists" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterExists) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterNotExists struct {
	NotExists param.Field[bool] `json:"not_exists" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterNotExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterNotExists) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterIn struct {
	In param.Field[[]string] `json:"in" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterIn) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterNotIn struct {
	NotIn param.Field[[]string] `json:"not_in" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterNotIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterNotIn) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQuerySort string

const (
	PrismObjectIdentityQueryParamsQuerySortAsc  PrismObjectIdentityQueryParamsQuerySort = "asc"
	PrismObjectIdentityQueryParamsQuerySortDesc PrismObjectIdentityQueryParamsQuerySort = "desc"
)

func (r PrismObjectIdentityQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectIdentityQueryParamsQuerySortAsc, PrismObjectIdentityQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectIdentityQueryParamsIDArray].
type PrismObjectIdentityQueryParamsIDUnion interface {
	ImplementsPrismObjectIdentityQueryParamsIDUnion()
}

type PrismObjectIdentityQueryParamsIDArray []string

func (r PrismObjectIdentityQueryParamsIDArray) ImplementsPrismObjectIdentityQueryParamsIDUnion() {}

type PrismObjectIdentityRestoreParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
