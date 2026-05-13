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

// Patch object
func (r *PrismObjectOrganizationService) Update(ctx context.Context, organizationID string, params PrismObjectOrganizationUpdateParams, opts ...option.RequestOption) (res *PrismObjectOrganizationUpdateResponse, err error) {
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
	if organizationID == "" {
		err = errors.New("missing required organizationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/organization/%s", params.TeamID, organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Delete object
func (r *PrismObjectOrganizationService) Delete(ctx context.Context, organizationID string, body PrismObjectOrganizationDeleteParams, opts ...option.RequestOption) (err error) {
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

// Duplicate object
func (r *PrismObjectOrganizationService) Duplicate(ctx context.Context, organizationID string, body PrismObjectOrganizationDuplicateParams, opts ...option.RequestOption) (res *PrismObjectOrganizationDuplicateResponse, err error) {
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
	if organizationID == "" {
		err = errors.New("missing required organizationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/organization/%s/duplicate", body.TeamID, organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectOrganizationService) Get(ctx context.Context, organizationID string, query PrismObjectOrganizationGetParams, opts ...option.RequestOption) (res *PrismObjectOrganizationGetResponse, err error) {
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
	if organizationID == "" {
		err = errors.New("missing required organizationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/organization/%s", query.TeamID, organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Query
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

// Restore object
func (r *PrismObjectOrganizationService) Restore(ctx context.Context, organizationID string, body PrismObjectOrganizationRestoreParams, opts ...option.RequestOption) (res *PrismObjectOrganizationRestoreResponse, err error) {
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
	if organizationID == "" {
		err = errors.New("missing required organizationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/organization/%s/restore", body.TeamID, organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectOrganizationNewResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                 `json:"default"`
	List    interface{}                            `json:"list"`
	JSON    prismObjectOrganizationNewResponseJSON `json:"-"`
}

// prismObjectOrganizationNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectOrganizationNewResponse]
type prismObjectOrganizationNewResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationNewResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectOrganizationUpdateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                    `json:"default"`
	List    interface{}                               `json:"list"`
	JSON    prismObjectOrganizationUpdateResponseJSON `json:"-"`
}

// prismObjectOrganizationUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationUpdateResponse]
type prismObjectOrganizationUpdateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationUpdateResponseJSON) RawJSON() string {
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

type PrismObjectOrganizationDuplicateResponse struct {
	ID   string                                       `json:"id" format:"uuid"`
	JSON prismObjectOrganizationDuplicateResponseJSON `json:"-"`
}

// prismObjectOrganizationDuplicateResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationDuplicateResponse]
type prismObjectOrganizationDuplicateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationDuplicateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationDuplicateResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectOrganizationGetResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                 `json:"default"`
	List    interface{}                            `json:"list"`
	JSON    prismObjectOrganizationGetResponseJSON `json:"-"`
}

// prismObjectOrganizationGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectOrganizationGetResponse]
type prismObjectOrganizationGetResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationGetResponseJSON) RawJSON() string {
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
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                       `json:"default"`
	List    interface{}                                  `json:"list"`
	JSON    prismObjectOrganizationQueryResponseDataJSON `json:"-"`
}

// prismObjectOrganizationQueryResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationQueryResponseData]
type prismObjectOrganizationQueryResponseDataJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectOrganizationRestoreResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                     `json:"default"`
	List    interface{}                                `json:"list"`
	JSON    prismObjectOrganizationRestoreResponseJSON `json:"-"`
}

// prismObjectOrganizationRestoreResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationRestoreResponse]
type prismObjectOrganizationRestoreResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationRestoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationRestoreResponseJSON) RawJSON() string {
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

type PrismObjectOrganizationUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismObjectOrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectOrganizationDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
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
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
	// App/CRM ID for context (optional)
	ListID param.Field[string] `json:"list_id" format:"uuid"`
}

func (r PrismObjectOrganizationBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectOrganizationDuplicateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectOrganizationGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
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
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectOrganizationQueryParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                           `json:"limit"`
	ListID param.Field[string]                                                          `json:"list_id" format:"uuid"`
	Page   param.Field[int64]                                                           `json:"page"`
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

type PrismObjectOrganizationQueryParamsQueryFilter struct {
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

func (r PrismObjectOrganizationQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilter) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

// Satisfied by [PrismObjectOrganizationQueryParamsQueryFilter],
// [PrismObjectOrganizationQueryParamsQueryFilter],
// [PrismObjectOrganizationQueryParamsQueryFilter],
// [PrismObjectOrganizationQueryParamsQueryFilter],
// [PrismObjectOrganizationQueryParamsQueryFilter],
// [PrismObjectOrganizationQueryParamsQueryFilter],
// [PrismObjectOrganizationQueryParamsQueryFilterLikeRegex],
// [PrismObjectOrganizationQueryParamsQueryFilterBeginsWith],
// [PrismObjectOrganizationQueryParamsQueryFilterEndsWith],
// [PrismObjectOrganizationQueryParamsQueryFilterNotContains],
// [PrismObjectOrganizationQueryParamsQueryFilterExists],
// [PrismObjectOrganizationQueryParamsQueryFilterNotExists],
// [PrismObjectOrganizationQueryParamsQueryFilterIn],
// [PrismObjectOrganizationQueryParamsQueryFilterNotIn],
// [PrismObjectOrganizationQueryParamsQueryFilter].
type PrismObjectOrganizationQueryParamsQueryFilterUnion interface {
	implementsPrismObjectOrganizationQueryParamsQueryFilterUnion()
}

type PrismObjectOrganizationQueryParamsQueryFilter struct {
	Equals param.Field[PrismObjectOrganizationQueryParamsQueryFilterUnion] `json:"=" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilter) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterLikeRegex struct {
	LikeRegex param.Field[string] `json:"like_regex" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterLikeRegex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterLikeRegex) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterBeginsWith struct {
	BeginsWith param.Field[string] `json:"begins_with" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterBeginsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterBeginsWith) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterEndsWith struct {
	EndsWith param.Field[string] `json:"ends_with" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterEndsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterEndsWith) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterNotContains struct {
	NotContains param.Field[string] `json:"not_contains" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterNotContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterNotContains) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterExists struct {
	Exists param.Field[bool] `json:"exists" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterExists) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterNotExists struct {
	NotExists param.Field[bool] `json:"not_exists" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterNotExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterNotExists) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterIn struct {
	In param.Field[[]string] `json:"in" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterIn) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterNotIn struct {
	NotIn param.Field[[]string] `json:"not_in" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterNotIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterNotIn) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
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

type PrismObjectOrganizationRestoreParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
