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

// Get object
func (r *PrismObjectEventService) Get(ctx context.Context, eventID string, query PrismObjectEventGetParams, opts ...option.RequestOption) (res *PrismObjectEventGetResponse, err error) {
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
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/event/%s", query.TeamID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
	path := fmt.Sprintf("v2/prism/query/%s/event", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
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
	// True when the page returned the maximum number of rows; another page may exist.
	HasMore bool                              `json:"has_more"`
	JSON    prismObjectEventQueryResponseJSON `json:"-"`
}

// prismObjectEventQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectEventQueryResponse]
type prismObjectEventQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
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
	Source     string                                `json:"source" api:"nullable" format:"uuid"`
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

type PrismObjectEventGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectEventQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                             `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[PrismObjectEventQueryParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[PrismObjectEventQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                           `json:"boxes"`
	Deleted param.Field[bool]                               `json:"deleted"`
	Sources param.Field[[]string]                           `json:"sources" format:"uuid"`
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
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectEventQueryParamsQueryFilterUnion] `json:"filter"`
	// Maximum number of rows to return. Capped server-side at 50; requests above the
	// cap are rejected.
	Limit  param.Field[int64]  `json:"limit"`
	ListID param.Field[string] `json:"list_id" format:"uuid"`
	Page   param.Field[int64]  `json:"page"`
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
	Contains        param.Field[interface{}] `json:"contains"`
	EndsWith        param.Field[string]      `json:"ends_with"`
	Exists          param.Field[bool]        `json:"exists"`
	In              param.Field[interface{}] `json:"in"`
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
