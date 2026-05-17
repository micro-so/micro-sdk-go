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

// PrismObjectActionService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectActionService] method instead.
type PrismObjectActionService struct {
	Options []option.RequestOption
	Grant   *PrismObjectActionGrantService
}

// NewPrismObjectActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectActionService(opts ...option.RequestOption) (r *PrismObjectActionService) {
	r = &PrismObjectActionService{}
	r.Options = opts
	r.Grant = NewPrismObjectActionGrantService(opts...)
	return
}

// Create object
func (r *PrismObjectActionService) New(ctx context.Context, params PrismObjectActionNewParams, opts ...option.RequestOption) (res *PrismObjectActionNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/action", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch object
func (r *PrismObjectActionService) Update(ctx context.Context, actionID string, params PrismObjectActionUpdateParams, opts ...option.RequestOption) (res *PrismObjectActionUpdateResponse, err error) {
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
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/action/%s", params.TeamID, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Delete object
func (r *PrismObjectActionService) Delete(ctx context.Context, actionID string, body PrismObjectActionDeleteParams, opts ...option.RequestOption) (err error) {
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
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/action/%s", body.TeamID, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: <100 records sync (immediate response), >=100 records
// async (S3/Lambda with WebSocket progress)
func (r *PrismObjectActionService) BulkNew(ctx context.Context, params PrismObjectActionBulkNewParams, opts ...option.RequestOption) (res *PrismObjectActionBulkNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/action/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Duplicate object
func (r *PrismObjectActionService) Duplicate(ctx context.Context, actionID string, body PrismObjectActionDuplicateParams, opts ...option.RequestOption) (res *PrismObjectActionDuplicateResponse, err error) {
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
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/action/%s/duplicate", body.TeamID, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectActionService) Get(ctx context.Context, actionID string, query PrismObjectActionGetParams, opts ...option.RequestOption) (res *PrismObjectActionGetResponse, err error) {
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
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/action/%s", query.TeamID, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Query
func (r *PrismObjectActionService) Query(ctx context.Context, params PrismObjectActionQueryParams, opts ...option.RequestOption) (res *PrismObjectActionQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/query/%s/action", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Restore object
func (r *PrismObjectActionService) Restore(ctx context.Context, actionID string, body PrismObjectActionRestoreParams, opts ...option.RequestOption) (res *PrismObjectActionRestoreResponse, err error) {
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
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/action/%s/restore", body.TeamID, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectActionNewResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}           `json:"default"`
	List    interface{}                      `json:"list"`
	JSON    prismObjectActionNewResponseJSON `json:"-"`
}

// prismObjectActionNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionNewResponse]
type prismObjectActionNewResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionNewResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectActionUpdateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}              `json:"default"`
	List    interface{}                         `json:"list"`
	JSON    prismObjectActionUpdateResponseJSON `json:"-"`
}

// prismObjectActionUpdateResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionUpdateResponse]
type prismObjectActionUpdateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionBulkNewResponse struct {
	Results []PrismObjectActionBulkNewResponseResult `json:"results"`
	Status  PrismObjectActionBulkNewResponseStatus   `json:"status"`
	Summary PrismObjectActionBulkNewResponseSummary  `json:"summary"`
	JSON    prismObjectActionBulkNewResponseJSON     `json:"-"`
}

// prismObjectActionBulkNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionBulkNewResponse]
type prismObjectActionBulkNewResponseJSON struct {
	Results     apijson.Field
	Status      apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionBulkNewResponseResult struct {
	ID       string                                     `json:"id" api:"nullable" format:"uuid"`
	Created  bool                                       `json:"created"`
	Error    string                                     `json:"error"`
	Existing bool                                       `json:"existing"`
	JSON     prismObjectActionBulkNewResponseResultJSON `json:"-"`
}

// prismObjectActionBulkNewResponseResultJSON contains the JSON metadata for the
// struct [PrismObjectActionBulkNewResponseResult]
type prismObjectActionBulkNewResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionBulkNewResponseStatus string

const (
	PrismObjectActionBulkNewResponseStatusComplete PrismObjectActionBulkNewResponseStatus = "complete"
)

func (r PrismObjectActionBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectActionBulkNewResponseStatusComplete:
		return true
	}
	return false
}

type PrismObjectActionBulkNewResponseSummary struct {
	Created  int64                                       `json:"created"`
	Errors   int64                                       `json:"errors"`
	Existing int64                                       `json:"existing"`
	Total    int64                                       `json:"total"`
	JSON     prismObjectActionBulkNewResponseSummaryJSON `json:"-"`
}

// prismObjectActionBulkNewResponseSummaryJSON contains the JSON metadata for the
// struct [PrismObjectActionBulkNewResponseSummary]
type prismObjectActionBulkNewResponseSummaryJSON struct {
	Created     apijson.Field
	Errors      apijson.Field
	Existing    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkNewResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkNewResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionDuplicateResponse struct {
	ID   string                                 `json:"id" format:"uuid"`
	JSON prismObjectActionDuplicateResponseJSON `json:"-"`
}

// prismObjectActionDuplicateResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionDuplicateResponse]
type prismObjectActionDuplicateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionDuplicateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionDuplicateResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectActionGetResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}           `json:"default"`
	List    interface{}                      `json:"list"`
	JSON    prismObjectActionGetResponseJSON `json:"-"`
}

// prismObjectActionGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionGetResponse]
type prismObjectActionGetResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionQueryResponse struct {
	Data []PrismObjectActionQueryResponseData `json:"data" api:"required"`
	// True when the page returned the maximum number of rows; another page may exist.
	HasMore bool                               `json:"has_more"`
	JSON    prismObjectActionQueryResponseJSON `json:"-"`
}

// prismObjectActionQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionQueryResponse]
type prismObjectActionQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectActionQueryResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                 `json:"properties"`
	Source     string                                 `json:"source" api:"nullable" format:"uuid"`
	JSON       prismObjectActionQueryResponseDataJSON `json:"-"`
}

// prismObjectActionQueryResponseDataJSON contains the JSON metadata for the struct
// [PrismObjectActionQueryResponseData]
type prismObjectActionQueryResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectActionQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectActionRestoreResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}               `json:"default"`
	List    interface{}                          `json:"list"`
	JSON    prismObjectActionRestoreResponseJSON `json:"-"`
}

// prismObjectActionRestoreResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionRestoreResponse]
type prismObjectActionRestoreResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionRestoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionRestoreResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismObjectActionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectActionUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismObjectActionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectActionDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectActionBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects param.Field[[]PrismObjectPropertiesParam]          `json:"objects" api:"required"`
	Options param.Field[PrismObjectActionBulkNewParamsOptions] `json:"options"`
}

func (r PrismObjectActionBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectActionBulkNewParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
	// App/CRM ID for context (optional)
	ListID param.Field[string] `json:"list_id" format:"uuid"`
}

func (r PrismObjectActionBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectActionDuplicateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectActionGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectActionQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                              `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[PrismObjectActionQueryParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[PrismObjectActionQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                            `json:"boxes"`
	Deleted param.Field[bool]                                `json:"deleted"`
	Sources param.Field[[]string]                            `json:"sources" format:"uuid"`
}

func (r PrismObjectActionQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectActionQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name). `id` is always returned at the top level of each
	// row and does not need to be selected.
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectActionQueryParamsQueryCombinator] `json:"combinator"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectActionQueryParamsQueryFilterUnion] `json:"filter"`
	// Maximum number of rows to return. Capped server-side at 50; requests above the
	// cap are rejected.
	Limit  param.Field[int64]  `json:"limit"`
	ListID param.Field[string] `json:"list_id" format:"uuid"`
	Page   param.Field[int64]  `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]PrismObjectActionQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectActionQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectActionQueryParamsQueryCombinator string

const (
	PrismObjectActionQueryParamsQueryCombinatorAnd PrismObjectActionQueryParamsQueryCombinator = "AND"
	PrismObjectActionQueryParamsQueryCombinatorOr  PrismObjectActionQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectActionQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectActionQueryParamsQueryCombinatorAnd, PrismObjectActionQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

type PrismObjectActionQueryParamsQueryFilter struct {
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

func (r PrismObjectActionQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilter) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

// Satisfied by [PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEq],
// [PrismObjectActionQueryParamsQueryFilterPrismQueryFilterNe],
// [PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLt],
// [PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGt],
// [PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLte],
// [PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGte],
// [PrismObjectActionQueryParamsQueryFilterContains],
// [PrismObjectActionQueryParamsQueryFilterBeginsWith],
// [PrismObjectActionQueryParamsQueryFilterEndsWith],
// [PrismObjectActionQueryParamsQueryFilterNotContains],
// [PrismObjectActionQueryParamsQueryFilterExists],
// [PrismObjectActionQueryParamsQueryFilterNotExists],
// [PrismObjectActionQueryParamsQueryFilterIn],
// [PrismObjectActionQueryParamsQueryFilterNotIn],
// [PrismObjectActionQueryParamsQueryFilter].
type PrismObjectActionQueryParamsQueryFilterUnion interface {
	implementsPrismObjectActionQueryParamsQueryFilterUnion()
}

type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEq struct {
	Equals param.Field[PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEqUnion] `json:"=" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEq) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEq) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEqUnion interface {
	ImplementsPrismObjectActionQueryParamsQueryFilterPrismQueryFilterEqUnion()
}

type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterNe struct {
	NotEquals param.Field[PrismObjectActionQueryParamsQueryFilterPrismQueryFilterNeUnion] `json:"!=" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterNe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterNe) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterNeUnion interface {
	ImplementsPrismObjectActionQueryParamsQueryFilterPrismQueryFilterNeUnion()
}

type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLt struct {
	Less param.Field[string] `json:"<" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLt) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGt struct {
	Greater param.Field[string] `json:">" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGt) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLte struct {
	LessOrEquals param.Field[string] `json:"<=" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLte) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGte struct {
	GreaterOrEquals param.Field[string] `json:">=" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGte) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterContains struct {
	Contains param.Field[PrismObjectActionQueryParamsQueryFilterContainsContainsUnion] `json:"contains" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterContains) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectActionQueryParamsQueryFilterContainsContainsArray].
type PrismObjectActionQueryParamsQueryFilterContainsContainsUnion interface {
	ImplementsPrismObjectActionQueryParamsQueryFilterContainsContainsUnion()
}

type PrismObjectActionQueryParamsQueryFilterContainsContainsArray []string

func (r PrismObjectActionQueryParamsQueryFilterContainsContainsArray) ImplementsPrismObjectActionQueryParamsQueryFilterContainsContainsUnion() {
}

type PrismObjectActionQueryParamsQueryFilterBeginsWith struct {
	BeginsWith param.Field[string] `json:"begins_with" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterBeginsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterBeginsWith) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterEndsWith struct {
	EndsWith param.Field[string] `json:"ends_with" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterEndsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterEndsWith) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterNotContains struct {
	NotContains param.Field[string] `json:"not_contains" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterNotContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterNotContains) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterExists struct {
	Exists param.Field[bool] `json:"exists" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterExists) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterNotExists struct {
	NotExists param.Field[bool] `json:"not_exists" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterNotExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterNotExists) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterIn struct {
	In param.Field[[]string] `json:"in" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterIn) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterNotIn struct {
	NotIn param.Field[[]string] `json:"not_in" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterNotIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterNotIn) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQuerySort string

const (
	PrismObjectActionQueryParamsQuerySortAsc  PrismObjectActionQueryParamsQuerySort = "asc"
	PrismObjectActionQueryParamsQuerySortDesc PrismObjectActionQueryParamsQuerySort = "desc"
)

func (r PrismObjectActionQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectActionQueryParamsQuerySortAsc, PrismObjectActionQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectActionQueryParamsIDArray].
type PrismObjectActionQueryParamsIDUnion interface {
	ImplementsPrismObjectActionQueryParamsIDUnion()
}

type PrismObjectActionQueryParamsIDArray []string

func (r PrismObjectActionQueryParamsIDArray) ImplementsPrismObjectActionQueryParamsIDUnion() {}

type PrismObjectActionRestoreParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
