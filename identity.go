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

// Identities link multiple contacts together as the same real-world person,
// deduplicating people who appear in different contexts.
//
// IdentityService contains methods and other services that help with interacting
// with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityService] method instead.
type IdentityService struct {
	Options []option.RequestOption
}

// NewIdentityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIdentityService(opts ...option.RequestOption) (r *IdentityService) {
	r = &IdentityService{}
	r.Options = opts
	return
}

// Create Identity
func (r *IdentityService) New(ctx context.Context, params IdentityNewParams, opts ...option.RequestOption) (res *IdentityNewResponse, err error) {
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

// Update Identity
func (r *IdentityService) Update(ctx context.Context, identityID string, params IdentityUpdateParams, opts ...option.RequestOption) (err error) {
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
	if identityID == "" {
		err = errors.New("missing required identityId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/identity/%s", params.TeamID, identityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return err
}

// List Identities
func (r *IdentityService) List(ctx context.Context, params IdentityListParams, opts ...option.RequestOption) (res *IdentityListResponse, err error) {
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

// Delete Identity
func (r *IdentityService) Delete(ctx context.Context, identityID string, body IdentityDeleteParams, opts ...option.RequestOption) (err error) {
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

// Import Identities
func (r *IdentityService) Import(ctx context.Context, params IdentityImportParams, opts ...option.RequestOption) (res *IdentityImportResponse, err error) {
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

type IdentityNewResponse struct {
	ID   string                  `json:"id" format:"uuid"`
	JSON identityNewResponseJSON `json:"-"`
}

// identityNewResponseJSON contains the JSON metadata for the struct
// [IdentityNewResponse]
type identityNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityNewResponseJSON) RawJSON() string {
	return r.raw
}

type IdentityListResponse struct {
	Data       []interface{}            `json:"data"`
	NextCursor string                   `json:"next_cursor" api:"nullable"`
	Total      int64                    `json:"total"`
	JSON       identityListResponseJSON `json:"-"`
}

// identityListResponseJSON contains the JSON metadata for the struct
// [IdentityListResponse]
type identityListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityListResponseJSON) RawJSON() string {
	return r.raw
}

type IdentityImportResponse struct {
	Results []IdentityImportResponseResult `json:"results"`
	Status  IdentityImportResponseStatus   `json:"status"`
	Summary IdentityImportResponseSummary  `json:"summary"`
	JSON    identityImportResponseJSON     `json:"-"`
}

// identityImportResponseJSON contains the JSON metadata for the struct
// [IdentityImportResponse]
type identityImportResponseJSON struct {
	Results     apijson.Field
	Status      apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityImportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityImportResponseJSON) RawJSON() string {
	return r.raw
}

type IdentityImportResponseResult struct {
	ID       string                           `json:"id" api:"nullable" format:"uuid"`
	Created  bool                             `json:"created"`
	Error    string                           `json:"error"`
	Existing bool                             `json:"existing"`
	JSON     identityImportResponseResultJSON `json:"-"`
}

// identityImportResponseResultJSON contains the JSON metadata for the struct
// [IdentityImportResponseResult]
type identityImportResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityImportResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityImportResponseResultJSON) RawJSON() string {
	return r.raw
}

type IdentityImportResponseStatus string

const (
	IdentityImportResponseStatusComplete IdentityImportResponseStatus = "complete"
)

func (r IdentityImportResponseStatus) IsKnown() bool {
	switch r {
	case IdentityImportResponseStatusComplete:
		return true
	}
	return false
}

type IdentityImportResponseSummary struct {
	Created  int64                             `json:"created"`
	Errors   int64                             `json:"errors"`
	Existing int64                             `json:"existing"`
	Total    int64                             `json:"total"`
	JSON     identityImportResponseSummaryJSON `json:"-"`
}

// identityImportResponseSummaryJSON contains the JSON metadata for the struct
// [IdentityImportResponseSummary]
type identityImportResponseSummaryJSON struct {
	Created     apijson.Field
	Errors      apijson.Field
	Existing    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityImportResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityImportResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type IdentityNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r IdentityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type IdentityUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r IdentityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type IdentityListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                    `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[IdentityListParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[IdentityListParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                  `json:"boxes"`
	Deleted param.Field[bool]                      `json:"deleted"`
	Sources param.Field[[]string]                  `json:"sources" format:"uuid"`
}

func (r IdentityListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityListParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[IdentityListParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                            `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values must be option slugs
	Filter param.Field[[]map[string]map[string]IdentityListParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                      `json:"limit"`
	Page   param.Field[int64]                                                      `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]IdentityListParamsQuerySort] `json:"sort"`
}

func (r IdentityListParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type IdentityListParamsQueryCombinator string

const (
	IdentityListParamsQueryCombinatorAnd IdentityListParamsQueryCombinator = "AND"
	IdentityListParamsQueryCombinatorOr  IdentityListParamsQueryCombinator = "OR"
)

func (r IdentityListParamsQueryCombinator) IsKnown() bool {
	switch r {
	case IdentityListParamsQueryCombinatorAnd, IdentityListParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [IdentityListParamsQueryFilterArray].
type IdentityListParamsQueryFilterUnion interface {
	ImplementsIdentityListParamsQueryFilterUnion()
}

type IdentityListParamsQueryFilterArray []string

func (r IdentityListParamsQueryFilterArray) ImplementsIdentityListParamsQueryFilterUnion() {}

type IdentityListParamsQuerySort string

const (
	IdentityListParamsQuerySortAsc  IdentityListParamsQuerySort = "asc"
	IdentityListParamsQuerySortDesc IdentityListParamsQuerySort = "desc"
)

func (r IdentityListParamsQuerySort) IsKnown() bool {
	switch r {
	case IdentityListParamsQuerySortAsc, IdentityListParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [IdentityListParamsIDArray].
type IdentityListParamsIDUnion interface {
	ImplementsIdentityListParamsIDUnion()
}

type IdentityListParamsIDArray []string

func (r IdentityListParamsIDArray) ImplementsIdentityListParamsIDUnion() {}

type IdentityDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type IdentityImportParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects param.Field[[]PrismObjectPropertiesParam] `json:"objects" api:"required"`
	Options param.Field[IdentityImportParamsOptions]  `json:"options"`
}

func (r IdentityImportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityImportParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// App/CRM ID for context (optional)
	CRMID param.Field[string] `json:"crm_id" format:"uuid"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
}

func (r IdentityImportParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
