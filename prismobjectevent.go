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

// PrismObjectEventService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectEventService] method instead.
type PrismObjectEventService struct {
	Options []option.RequestOption
	Grant   *PrismObjectEventGrantService
}

// NewPrismObjectEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectEventService(opts ...option.RequestOption) (r *PrismObjectEventService) {
	r = &PrismObjectEventService{}
	r.Options = opts
	r.Grant = NewPrismObjectEventGrantService(opts...)
	return
}

// Convenience list endpoint. Equivalent to
// `POST /v2/prism/{teamId}/{objectType}/query` with an empty body, plus
// query-string sugar for the common cases. Any unrecognized query parameter is
// interpreted as an equality filter on a property of that name; pass arrays for
// `in`. Values are received as strings, so non-string property filters via this
// endpoint may not work — use the `query` endpoint for typed comparisons or
// anything beyond simple equality.
func (r *PrismObjectEventService) List(ctx context.Context, params PrismObjectEventListParams, opts ...option.RequestOption) (res *PrismObjectEventListResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/event", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns the total number of records of this object type that the caller can see.
// Avoids the page-overshoot anti-pattern — clients no longer need to keep paging
// until `has_more` flips false to discover the total. Currently does not apply
// query filters; for a filtered total, pass `include_total: true` in a POST
// `/query` body.
func (r *PrismObjectEventService) Count(ctx context.Context, params PrismObjectEventCountParams, opts ...option.RequestOption) (res *PrismObjectEventCountResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/event/count", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns the single record whose property `{slug}` equals `{value}`. 404 if
// nothing matches; 409 if more than one record matches.
func (r *PrismObjectEventService) Find(ctx context.Context, slug string, value string, params PrismObjectEventFindParams, opts ...option.RequestOption) (res *PrismObjectEventFindResponse, err error) {
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
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	if value == "" {
		err = errors.New("missing required value parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/event/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectEventService) Get(ctx context.Context, eventID string, params PrismObjectEventGetParams, opts ...option.RequestOption) (res *PrismObjectEventGetResponse, err error) {
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
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/event/%s", params.TeamID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Query
func (r *PrismObjectEventService) Query(ctx context.Context, params PrismObjectEventQueryParams, opts ...option.RequestOption) (res *PrismObjectEventQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/event/query", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type PrismObjectEventListResponse struct {
	Data []PrismObjectEventListResponseData `json:"data" api:"required"`
	// Accurate end-of-data signal — false on the last page, never forces clients to
	// overshoot.
	HasMore    bool   `json:"has_more" api:"required"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// Populated only when `?include_total=true` was passed.
	Total int64                            `json:"total" api:"nullable"`
	JSON  prismObjectEventListResponseJSON `json:"-"`
}

// prismObjectEventListResponseJSON contains the JSON metadata for the struct
// [PrismObjectEventListResponse]
type prismObjectEventListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEventListResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectEventListResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}               `json:"properties"`
	Source     []string                             `json:"source" api:"nullable"`
	JSON       prismObjectEventListResponseDataJSON `json:"-"`
}

// prismObjectEventListResponseDataJSON contains the JSON metadata for the struct
// [PrismObjectEventListResponseData]
type prismObjectEventListResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectEventListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEventListResponseDataJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEventCountResponse struct {
	// Number of records matching the access scope.
	Total int64                             `json:"total" api:"required"`
	JSON  prismObjectEventCountResponseJSON `json:"-"`
}

// prismObjectEventCountResponseJSON contains the JSON metadata for the struct
// [PrismObjectEventCountResponse]
type prismObjectEventCountResponseJSON struct {
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEventCountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEventCountResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectEventFindResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}           `json:"default"`
	List    interface{}                      `json:"list"`
	JSON    prismObjectEventFindResponseJSON `json:"-"`
}

// prismObjectEventFindResponseJSON contains the JSON metadata for the struct
// [PrismObjectEventFindResponse]
type prismObjectEventFindResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEventFindResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEventFindResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectEventGetResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}          `json:"default"`
	List    interface{}                     `json:"list"`
	JSON    prismObjectEventGetResponseJSON `json:"-"`
}

// prismObjectEventGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectEventGetResponse]
type prismObjectEventGetResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEventGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEventGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEventQueryResponse struct {
	Data []PrismObjectEventQueryResponseData `json:"data" api:"required"`
	// Accurate end-of-data signal. False when this page contains the last record; true
	// only when at least one more record exists. (Implementation note: the server
	// fetches one extra row internally to determine this — clients never need to
	// overshoot to discover the end.)
	HasMore bool `json:"has_more" api:"required"`
	// Opaque cursor pointing at the next page. Pass it back unchanged in the request
	// body (`cursor`) of the next call. Null when `has_more` is false.
	NextCursor string `json:"next_cursor" api:"nullable"`
	// Only populated when the request set `include_total: true`. Total number of
	// records matching the query, ignoring pagination. Opt-in because it costs an
	// additional pass over the result set.
	Total int64                             `json:"total" api:"nullable"`
	JSON  prismObjectEventQueryResponseJSON `json:"-"`
}

// prismObjectEventQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectEventQueryResponse]
type prismObjectEventQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEventQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEventQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectEventQueryResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                `json:"properties"`
	Source     []string                              `json:"source" api:"nullable"`
	JSON       prismObjectEventQueryResponseDataJSON `json:"-"`
}

// prismObjectEventQueryResponseDataJSON contains the JSON metadata for the struct
// [PrismObjectEventQueryResponseData]
type prismObjectEventQueryResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectEventQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEventQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEventListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Opaque cursor from a previous response's `next_cursor`. Pass it back unchanged
	// to fetch the next page.
	Cursor param.Field[string] `query:"cursor"`
	// Include soft-deleted records. Pass the literal string `true`.
	Deleted param.Field[bool] `query:"deleted"`
	// When set to `true`, the response includes a `total` field with the unpaginated
	// row count. Costs an extra pass; prefer `GET .../count` for the unfiltered total.
	IncludeTotal param.Field[bool] `query:"include_total"`
	// Maximum number of rows to return. Capped server-side at 50.
	Limit param.Field[int64] `query:"limit"`
	// Scope properties to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
	// Comma-separated property slugs to return. Use dot notation for relationships.
	// `id` is always returned at the top level. Defaults to all properties.
	Select param.Field[string] `query:"select"`
	// Comma-separated list of slugs. Prefix with `-` for descending. Example:
	// `sort=-updated_at,name`.
	Sort param.Field[string] `query:"sort"`
}

// URLQuery serializes [PrismObjectEventListParams]'s query parameters as
// `url.Values`.
func (r PrismObjectEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectEventCountParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the count to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectEventCountParams]'s query parameters as
// `url.Values`.
func (r PrismObjectEventCountParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectEventFindParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the lookup to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectEventFindParams]'s query parameters as
// `url.Values`.
func (r PrismObjectEventFindParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectEventGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Comma-separated property slugs to return. Use dot notation for relationships.
	// `id` is always returned at the top level. Defaults to all properties.
	Select param.Field[string] `query:"select"`
}

// URLQuery serializes [PrismObjectEventGetParams]'s query parameters as
// `url.Values`.
func (r PrismObjectEventGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectEventQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string]                             `path:"teamId" api:"required" format:"uuid"`
	Query  param.Field[PrismObjectEventQueryParamsQuery]   `json:"query" api:"required"`
	ID     param.Field[PrismObjectEventQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes  param.Field[[]string]                           `json:"boxes"`
	// Alternative location for the opaque cursor (a sibling of `query`). Use whichever
	// feels more natural; if both are present, `query.cursor` wins.
	Cursor  param.Field[string] `json:"cursor"`
	Deleted param.Field[bool]   `json:"deleted"`
	// When true, the response includes a `total` field with the unpaginated row count.
	// Costs an additional pass over the result set — for unfiltered totals prefer
	// `GET /v2/prism/{teamId}/{objectType}/count` instead.
	IncludeTotal param.Field[bool]     `json:"include_total"`
	Sources      param.Field[[]string] `json:"sources" format:"uuid"`
}

func (r PrismObjectEventQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectEventQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name). `id` is always returned at the top level of each
	// row and does not need to be selected.
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectEventQueryParamsQueryCombinator] `json:"combinator"`
	// Opaque cursor from a previous response's `next_cursor`. Pass it back unchanged
	// to fetch the next page. When set, `page` and `limit` are derived from the cursor
	// and any explicit values are ignored.
	Cursor param.Field[string] `json:"cursor"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectEventQueryParamsQueryFilterUnion] `json:"filter"`
	// Maximum number of rows to return. Capped server-side at 50; requests above the
	// cap are rejected.
	Limit  param.Field[int64]  `json:"limit"`
	ListID param.Field[string] `json:"list_id" format:"uuid"`
	// Page number (1-based). Prefer `cursor`. Page-number pagination drifts under
	// concurrent writes; use it only for one-shot exports.
	//
	// Deprecated: deprecated
	Page param.Field[int64] `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]PrismObjectEventQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectEventQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectEventQueryParamsQueryCombinator string

const (
	PrismObjectEventQueryParamsQueryCombinatorAnd PrismObjectEventQueryParamsQueryCombinator = "AND"
	PrismObjectEventQueryParamsQueryCombinatorOr  PrismObjectEventQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectEventQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectEventQueryParamsQueryCombinatorAnd, PrismObjectEventQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

type PrismObjectEventQueryParamsQueryFilter struct {
	NotEquals       param.Field[interface{}] `json:"!="`
	Less            param.Field[string]      `json:"<"`
	LessOrEquals    param.Field[string]      `json:"<="`
	Equals          param.Field[interface{}] `json:"="`
	Greater         param.Field[string]      `json:">"`
	GreaterOrEquals param.Field[string]      `json:">="`
	BeginsWith      param.Field[string]      `json:"begins_with"`
	Between         param.Field[interface{}] `json:"between"`
	Contains        param.Field[interface{}] `json:"contains"`
	EndsWith        param.Field[string]      `json:"ends_with"`
	Exists          param.Field[bool]        `json:"exists"`
	In              param.Field[interface{}] `json:"in"`
	IsNotNull       param.Field[interface{}] `json:"is_not_null"`
	IsNull          param.Field[interface{}] `json:"is_null"`
	NotContains     param.Field[string]      `json:"not_contains"`
	NotExists       param.Field[bool]        `json:"not_exists"`
	NotIn           param.Field[interface{}] `json:"not_in"`
}

func (r PrismObjectEventQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilter) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

// Satisfied by [PrismObjectEventQueryParamsQueryFilterPrismQueryFilterEq],
// [PrismObjectEventQueryParamsQueryFilterPrismQueryFilterNe],
// [PrismObjectEventQueryParamsQueryFilterPrismQueryFilterLt],
// [PrismObjectEventQueryParamsQueryFilterPrismQueryFilterGt],
// [PrismObjectEventQueryParamsQueryFilterPrismQueryFilterLte],
// [PrismObjectEventQueryParamsQueryFilterPrismQueryFilterGte],
// [PrismObjectEventQueryParamsQueryFilterContains],
// [PrismObjectEventQueryParamsQueryFilterBeginsWith],
// [PrismObjectEventQueryParamsQueryFilterEndsWith],
// [PrismObjectEventQueryParamsQueryFilterNotContains],
// [PrismObjectEventQueryParamsQueryFilterExists],
// [PrismObjectEventQueryParamsQueryFilterNotExists],
// [PrismObjectEventQueryParamsQueryFilterIsNull],
// [PrismObjectEventQueryParamsQueryFilterIsNotNull],
// [PrismObjectEventQueryParamsQueryFilterBetween],
// [PrismObjectEventQueryParamsQueryFilterIn],
// [PrismObjectEventQueryParamsQueryFilterNotIn],
// [PrismObjectEventQueryParamsQueryFilter].
type PrismObjectEventQueryParamsQueryFilterUnion interface {
	implementsPrismObjectEventQueryParamsQueryFilterUnion()
}

type PrismObjectEventQueryParamsQueryFilterPrismQueryFilterEq struct {
	Equals param.Field[PrismObjectEventQueryParamsQueryFilterPrismQueryFilterEqUnion] `json:"=" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterPrismQueryFilterEq) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterPrismQueryFilterEq) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectEventQueryParamsQueryFilterPrismQueryFilterEqUnion interface {
	ImplementsPrismObjectEventQueryParamsQueryFilterPrismQueryFilterEqUnion()
}

type PrismObjectEventQueryParamsQueryFilterPrismQueryFilterNe struct {
	NotEquals param.Field[PrismObjectEventQueryParamsQueryFilterPrismQueryFilterNeUnion] `json:"!=" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterPrismQueryFilterNe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterPrismQueryFilterNe) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectEventQueryParamsQueryFilterPrismQueryFilterNeUnion interface {
	ImplementsPrismObjectEventQueryParamsQueryFilterPrismQueryFilterNeUnion()
}

type PrismObjectEventQueryParamsQueryFilterPrismQueryFilterLt struct {
	Less param.Field[string] `json:"<" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterPrismQueryFilterLt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterPrismQueryFilterLt) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

type PrismObjectEventQueryParamsQueryFilterPrismQueryFilterGt struct {
	Greater param.Field[string] `json:">" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterPrismQueryFilterGt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterPrismQueryFilterGt) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

type PrismObjectEventQueryParamsQueryFilterPrismQueryFilterLte struct {
	LessOrEquals param.Field[string] `json:"<=" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterPrismQueryFilterLte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterPrismQueryFilterLte) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

type PrismObjectEventQueryParamsQueryFilterPrismQueryFilterGte struct {
	GreaterOrEquals param.Field[string] `json:">=" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterPrismQueryFilterGte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterPrismQueryFilterGte) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

type PrismObjectEventQueryParamsQueryFilterContains struct {
	Contains param.Field[PrismObjectEventQueryParamsQueryFilterContainsContainsUnion] `json:"contains" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterContains) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectEventQueryParamsQueryFilterContainsContainsArray].
type PrismObjectEventQueryParamsQueryFilterContainsContainsUnion interface {
	ImplementsPrismObjectEventQueryParamsQueryFilterContainsContainsUnion()
}

type PrismObjectEventQueryParamsQueryFilterContainsContainsArray []string

func (r PrismObjectEventQueryParamsQueryFilterContainsContainsArray) ImplementsPrismObjectEventQueryParamsQueryFilterContainsContainsUnion() {
}

type PrismObjectEventQueryParamsQueryFilterBeginsWith struct {
	BeginsWith param.Field[string] `json:"begins_with" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterBeginsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterBeginsWith) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

type PrismObjectEventQueryParamsQueryFilterEndsWith struct {
	EndsWith param.Field[string] `json:"ends_with" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterEndsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterEndsWith) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

type PrismObjectEventQueryParamsQueryFilterNotContains struct {
	NotContains param.Field[string] `json:"not_contains" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterNotContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterNotContains) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

type PrismObjectEventQueryParamsQueryFilterExists struct {
	Exists param.Field[bool] `json:"exists" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterExists) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

type PrismObjectEventQueryParamsQueryFilterNotExists struct {
	NotExists param.Field[bool] `json:"not_exists" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterNotExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterNotExists) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

type PrismObjectEventQueryParamsQueryFilterIsNull struct {
	IsNull param.Field[PrismObjectEventQueryParamsQueryFilterIsNullIsNullUnion] `json:"is_null" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterIsNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterIsNull) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectEventQueryParamsQueryFilterIsNullIsNullArray].
type PrismObjectEventQueryParamsQueryFilterIsNullIsNullUnion interface {
	ImplementsPrismObjectEventQueryParamsQueryFilterIsNullIsNullUnion()
}

type PrismObjectEventQueryParamsQueryFilterIsNullIsNullArray []string

func (r PrismObjectEventQueryParamsQueryFilterIsNullIsNullArray) ImplementsPrismObjectEventQueryParamsQueryFilterIsNullIsNullUnion() {
}

type PrismObjectEventQueryParamsQueryFilterIsNotNull struct {
	IsNotNull param.Field[PrismObjectEventQueryParamsQueryFilterIsNotNullIsNotNullUnion] `json:"is_not_null" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterIsNotNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterIsNotNull) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectEventQueryParamsQueryFilterIsNotNullIsNotNullArray].
type PrismObjectEventQueryParamsQueryFilterIsNotNullIsNotNullUnion interface {
	ImplementsPrismObjectEventQueryParamsQueryFilterIsNotNullIsNotNullUnion()
}

type PrismObjectEventQueryParamsQueryFilterIsNotNullIsNotNullArray []string

func (r PrismObjectEventQueryParamsQueryFilterIsNotNullIsNotNullArray) ImplementsPrismObjectEventQueryParamsQueryFilterIsNotNullIsNotNullUnion() {
}

type PrismObjectEventQueryParamsQueryFilterBetween struct {
	Between param.Field[PrismObjectEventQueryParamsQueryFilterBetweenBetweenUnion] `json:"between" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterBetween) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterBetween) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectEventQueryParamsQueryFilterBetweenBetweenArray].
type PrismObjectEventQueryParamsQueryFilterBetweenBetweenUnion interface {
	ImplementsPrismObjectEventQueryParamsQueryFilterBetweenBetweenUnion()
}

type PrismObjectEventQueryParamsQueryFilterBetweenBetweenArray []string

func (r PrismObjectEventQueryParamsQueryFilterBetweenBetweenArray) ImplementsPrismObjectEventQueryParamsQueryFilterBetweenBetweenUnion() {
}

type PrismObjectEventQueryParamsQueryFilterIn struct {
	In param.Field[[]string] `json:"in" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterIn) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

type PrismObjectEventQueryParamsQueryFilterNotIn struct {
	NotIn param.Field[[]string] `json:"not_in" api:"required"`
}

func (r PrismObjectEventQueryParamsQueryFilterNotIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEventQueryParamsQueryFilterNotIn) implementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

type PrismObjectEventQueryParamsQuerySort string

const (
	PrismObjectEventQueryParamsQuerySortAsc  PrismObjectEventQueryParamsQuerySort = "asc"
	PrismObjectEventQueryParamsQuerySortDesc PrismObjectEventQueryParamsQuerySort = "desc"
)

func (r PrismObjectEventQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectEventQueryParamsQuerySortAsc, PrismObjectEventQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectEventQueryParamsIDArray].
type PrismObjectEventQueryParamsIDUnion interface {
	ImplementsPrismObjectEventQueryParamsIDUnion()
}

type PrismObjectEventQueryParamsIDArray []string

func (r PrismObjectEventQueryParamsIDArray) ImplementsPrismObjectEventQueryParamsIDUnion() {}
