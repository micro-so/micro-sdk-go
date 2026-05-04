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
func (r *PrismObjectActionService) New(ctx context.Context, params PrismObjectActionNewParams, opts ...option.RequestOption) (res *PrismObjectProperties, err error) {
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
func (r *PrismObjectActionService) Update(ctx context.Context, actionID string, params PrismObjectActionUpdateParams, opts ...option.RequestOption) (res *PrismObjectProperties, err error) {
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
func (r *PrismObjectActionService) Get(ctx context.Context, actionID string, query PrismObjectActionGetParams, opts ...option.RequestOption) (res *PrismObjectProperties, err error) {
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

// Query v2
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
func (r *PrismObjectActionService) Restore(ctx context.Context, actionID string, body PrismObjectActionRestoreParams, opts ...option.RequestOption) (res *PrismObjectProperties, err error) {
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

type PrismObjectActionQueryResponse struct {
	Data  []interface{}                      `json:"data"`
	Total int64                              `json:"total"`
	JSON  prismObjectActionQueryResponseJSON `json:"-"`
}

// prismObjectActionQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionQueryResponse]
type prismObjectActionQueryResponseJSON struct {
	Data        apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionQueryResponseJSON) RawJSON() string {
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
	// App/CRM ID for context (optional)
	CRMID param.Field[string] `json:"crm_id" format:"uuid"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
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
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectActionQueryParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                                      `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values must be option slugs
	Filter param.Field[[]map[string]map[string]PrismObjectActionQueryParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                                `json:"limit"`
	Page   param.Field[int64]                                                                `json:"page"`
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

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectActionQueryParamsQueryFilterArray].
type PrismObjectActionQueryParamsQueryFilterUnion interface {
	ImplementsPrismObjectActionQueryParamsQueryFilterUnion()
}

type PrismObjectActionQueryParamsQueryFilterArray []string

func (r PrismObjectActionQueryParamsQueryFilterArray) ImplementsPrismObjectActionQueryParamsQueryFilterUnion() {
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
