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

// OrganizationService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	return
}

// Create Organization
func (r *OrganizationService) New(ctx context.Context, params OrganizationNewParams, opts ...option.RequestOption) (res *OrganizationNewResponse, err error) {
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

// Update Organization
func (r *OrganizationService) Update(ctx context.Context, organizationID string, params OrganizationUpdateParams, opts ...option.RequestOption) (err error) {
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
	if organizationID == "" {
		err = errors.New("missing required organizationId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/organization/%s", params.TeamID, organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return err
}

// List Organizations
func (r *OrganizationService) List(ctx context.Context, params OrganizationListParams, opts ...option.RequestOption) (res *OrganizationListResponse, err error) {
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

// Delete Organization
func (r *OrganizationService) Delete(ctx context.Context, organizationID string, body OrganizationDeleteParams, opts ...option.RequestOption) (err error) {
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
	if organizationID == "" {
		err = errors.New("missing required organizationId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/organization/%s", body.TeamID, organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import Organizations
func (r *OrganizationService) Import(ctx context.Context, params OrganizationImportParams, opts ...option.RequestOption) (res *OrganizationImportResponse, err error) {
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

type OrganizationNewResponse struct {
	ID   string                      `json:"id" format:"uuid"`
	JSON organizationNewResponseJSON `json:"-"`
}

// organizationNewResponseJSON contains the JSON metadata for the struct
// [OrganizationNewResponse]
type organizationNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponse struct {
	Data       []interface{}                `json:"data"`
	NextCursor string                       `json:"next_cursor" api:"nullable"`
	Total      int64                        `json:"total"`
	JSON       organizationListResponseJSON `json:"-"`
}

// organizationListResponseJSON contains the JSON metadata for the struct
// [OrganizationListResponse]
type organizationListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationImportResponse struct {
	Results []OrganizationImportResponseResult `json:"results"`
	Status  OrganizationImportResponseStatus   `json:"status"`
	Summary OrganizationImportResponseSummary  `json:"summary"`
	JSON    organizationImportResponseJSON     `json:"-"`
}

// organizationImportResponseJSON contains the JSON metadata for the struct
// [OrganizationImportResponse]
type organizationImportResponseJSON struct {
	Results     apijson.Field
	Status      apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationImportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationImportResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationImportResponseResult struct {
	ID       string                               `json:"id" api:"nullable" format:"uuid"`
	Created  bool                                 `json:"created"`
	Error    string                               `json:"error"`
	Existing bool                                 `json:"existing"`
	JSON     organizationImportResponseResultJSON `json:"-"`
}

// organizationImportResponseResultJSON contains the JSON metadata for the struct
// [OrganizationImportResponseResult]
type organizationImportResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationImportResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationImportResponseResultJSON) RawJSON() string {
	return r.raw
}

type OrganizationImportResponseStatus string

const (
	OrganizationImportResponseStatusComplete OrganizationImportResponseStatus = "complete"
)

func (r OrganizationImportResponseStatus) IsKnown() bool {
	switch r {
	case OrganizationImportResponseStatusComplete:
		return true
	}
	return false
}

type OrganizationImportResponseSummary struct {
	Created  int64                                 `json:"created"`
	Errors   int64                                 `json:"errors"`
	Existing int64                                 `json:"existing"`
	Total    int64                                 `json:"total"`
	JSON     organizationImportResponseSummaryJSON `json:"-"`
}

// organizationImportResponseSummaryJSON contains the JSON metadata for the struct
// [OrganizationImportResponseSummary]
type organizationImportResponseSummaryJSON struct {
	Created     apijson.Field
	Errors      apijson.Field
	Existing    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationImportResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationImportResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type OrganizationUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type OrganizationListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                        `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[OrganizationListParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[OrganizationListParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                      `json:"boxes"`
	Deleted param.Field[bool]                          `json:"deleted"`
	Sources param.Field[[]string]                      `json:"sources" format:"uuid"`
}

func (r OrganizationListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationListParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[OrganizationListParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                                `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values must be option slugs
	Filter param.Field[[]map[string]map[string]OrganizationListParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                          `json:"limit"`
	Page   param.Field[int64]                                                          `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]OrganizationListParamsQuerySort] `json:"sort"`
}

func (r OrganizationListParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type OrganizationListParamsQueryCombinator string

const (
	OrganizationListParamsQueryCombinatorAnd OrganizationListParamsQueryCombinator = "AND"
	OrganizationListParamsQueryCombinatorOr  OrganizationListParamsQueryCombinator = "OR"
)

func (r OrganizationListParamsQueryCombinator) IsKnown() bool {
	switch r {
	case OrganizationListParamsQueryCombinatorAnd, OrganizationListParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [OrganizationListParamsQueryFilterArray].
type OrganizationListParamsQueryFilterUnion interface {
	ImplementsOrganizationListParamsQueryFilterUnion()
}

type OrganizationListParamsQueryFilterArray []string

func (r OrganizationListParamsQueryFilterArray) ImplementsOrganizationListParamsQueryFilterUnion() {}

type OrganizationListParamsQuerySort string

const (
	OrganizationListParamsQuerySortAsc  OrganizationListParamsQuerySort = "asc"
	OrganizationListParamsQuerySortDesc OrganizationListParamsQuerySort = "desc"
)

func (r OrganizationListParamsQuerySort) IsKnown() bool {
	switch r {
	case OrganizationListParamsQuerySortAsc, OrganizationListParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [OrganizationListParamsIDArray].
type OrganizationListParamsIDUnion interface {
	ImplementsOrganizationListParamsIDUnion()
}

type OrganizationListParamsIDArray []string

func (r OrganizationListParamsIDArray) ImplementsOrganizationListParamsIDUnion() {}

type OrganizationDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type OrganizationImportParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects param.Field[[]PrismObjectPropertiesParam]    `json:"objects" api:"required"`
	Options param.Field[OrganizationImportParamsOptions] `json:"options"`
}

func (r OrganizationImportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationImportParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// App/CRM ID for context (optional)
	CRMID param.Field[string] `json:"crm_id" format:"uuid"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
}

func (r OrganizationImportParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
