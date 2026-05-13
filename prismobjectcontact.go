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

// PrismObjectContactService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectContactService] method instead.
type PrismObjectContactService struct {
	Options []option.RequestOption
}

// NewPrismObjectContactService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectContactService(opts ...option.RequestOption) (r *PrismObjectContactService) {
	r = &PrismObjectContactService{}
	r.Options = opts
	return
}

// Create object
func (r *PrismObjectContactService) New(ctx context.Context, params PrismObjectContactNewParams, opts ...option.RequestOption) (res *PrismObjectContactNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/contact", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch object
func (r *PrismObjectContactService) Update(ctx context.Context, contactID string, params PrismObjectContactUpdateParams, opts ...option.RequestOption) (res *PrismObjectContactUpdateResponse, err error) {
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
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/contact/%s", params.TeamID, contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Delete object
func (r *PrismObjectContactService) Delete(ctx context.Context, contactID string, body PrismObjectContactDeleteParams, opts ...option.RequestOption) (err error) {
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
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/contact/%s", body.TeamID, contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: <100 records sync (immediate response), >=100 records
// async (S3/Lambda with WebSocket progress)
func (r *PrismObjectContactService) BulkNew(ctx context.Context, params PrismObjectContactBulkNewParams, opts ...option.RequestOption) (res *PrismObjectContactBulkNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/contact/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Duplicate object
func (r *PrismObjectContactService) Duplicate(ctx context.Context, contactID string, body PrismObjectContactDuplicateParams, opts ...option.RequestOption) (res *PrismObjectContactDuplicateResponse, err error) {
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
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/contact/%s/duplicate", body.TeamID, contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectContactService) Get(ctx context.Context, contactID string, query PrismObjectContactGetParams, opts ...option.RequestOption) (res *PrismObjectContactGetResponse, err error) {
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
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/contact/%s", query.TeamID, contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Query
func (r *PrismObjectContactService) Query(ctx context.Context, params PrismObjectContactQueryParams, opts ...option.RequestOption) (res *PrismObjectContactQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/query/%s/contact", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Restore object
func (r *PrismObjectContactService) Restore(ctx context.Context, contactID string, body PrismObjectContactRestoreParams, opts ...option.RequestOption) (res *PrismObjectContactRestoreResponse, err error) {
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
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/contact/%s/restore", body.TeamID, contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectContactNewResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}            `json:"default"`
	List    interface{}                       `json:"list"`
	JSON    prismObjectContactNewResponseJSON `json:"-"`
}

// prismObjectContactNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactNewResponse]
type prismObjectContactNewResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactNewResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectContactUpdateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}               `json:"default"`
	List    interface{}                          `json:"list"`
	JSON    prismObjectContactUpdateResponseJSON `json:"-"`
}

// prismObjectContactUpdateResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactUpdateResponse]
type prismObjectContactUpdateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkNewResponse struct {
	Results []PrismObjectContactBulkNewResponseResult `json:"results"`
	Status  PrismObjectContactBulkNewResponseStatus   `json:"status"`
	Summary PrismObjectContactBulkNewResponseSummary  `json:"summary"`
	JSON    prismObjectContactBulkNewResponseJSON     `json:"-"`
}

// prismObjectContactBulkNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactBulkNewResponse]
type prismObjectContactBulkNewResponseJSON struct {
	Results     apijson.Field
	Status      apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkNewResponseResult struct {
	ID       string                                      `json:"id" api:"nullable" format:"uuid"`
	Created  bool                                        `json:"created"`
	Error    string                                      `json:"error"`
	Existing bool                                        `json:"existing"`
	JSON     prismObjectContactBulkNewResponseResultJSON `json:"-"`
}

// prismObjectContactBulkNewResponseResultJSON contains the JSON metadata for the
// struct [PrismObjectContactBulkNewResponseResult]
type prismObjectContactBulkNewResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkNewResponseStatus string

const (
	PrismObjectContactBulkNewResponseStatusComplete PrismObjectContactBulkNewResponseStatus = "complete"
)

func (r PrismObjectContactBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectContactBulkNewResponseStatusComplete:
		return true
	}
	return false
}

type PrismObjectContactBulkNewResponseSummary struct {
	Created  int64                                        `json:"created"`
	Errors   int64                                        `json:"errors"`
	Existing int64                                        `json:"existing"`
	Total    int64                                        `json:"total"`
	JSON     prismObjectContactBulkNewResponseSummaryJSON `json:"-"`
}

// prismObjectContactBulkNewResponseSummaryJSON contains the JSON metadata for the
// struct [PrismObjectContactBulkNewResponseSummary]
type prismObjectContactBulkNewResponseSummaryJSON struct {
	Created     apijson.Field
	Errors      apijson.Field
	Existing    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkNewResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkNewResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactDuplicateResponse struct {
	ID   string                                  `json:"id" format:"uuid"`
	JSON prismObjectContactDuplicateResponseJSON `json:"-"`
}

// prismObjectContactDuplicateResponseJSON contains the JSON metadata for the
// struct [PrismObjectContactDuplicateResponse]
type prismObjectContactDuplicateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactDuplicateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactDuplicateResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectContactGetResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}            `json:"default"`
	List    interface{}                       `json:"list"`
	JSON    prismObjectContactGetResponseJSON `json:"-"`
}

// prismObjectContactGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactGetResponse]
type prismObjectContactGetResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactQueryResponse struct {
	Data []PrismObjectContactQueryResponseData `json:"data" api:"required"`
	// True when the page returned the maximum number of rows; another page may exist.
	HasMore bool                                `json:"has_more"`
	JSON    prismObjectContactQueryResponseJSON `json:"-"`
}

// prismObjectContactQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactQueryResponse]
type prismObjectContactQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectContactQueryResponseData struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                  `json:"default"`
	List    interface{}                             `json:"list"`
	JSON    prismObjectContactQueryResponseDataJSON `json:"-"`
}

// prismObjectContactQueryResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectContactQueryResponseData]
type prismObjectContactQueryResponseDataJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectContactRestoreResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                `json:"default"`
	List    interface{}                           `json:"list"`
	JSON    prismObjectContactRestoreResponseJSON `json:"-"`
}

// prismObjectContactRestoreResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactRestoreResponse]
type prismObjectContactRestoreResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactRestoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactRestoreResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismObjectContactNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectContactUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismObjectContactUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectContactDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectContactBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects param.Field[[]PrismObjectPropertiesParam]           `json:"objects" api:"required"`
	Options param.Field[PrismObjectContactBulkNewParamsOptions] `json:"options"`
}

func (r PrismObjectContactBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectContactBulkNewParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
	// App/CRM ID for context (optional)
	ListID param.Field[string] `json:"list_id" format:"uuid"`
}

func (r PrismObjectContactBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectContactDuplicateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectContactGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectContactQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                               `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[PrismObjectContactQueryParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[PrismObjectContactQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                             `json:"boxes"`
	Deleted param.Field[bool]                                 `json:"deleted"`
	Sources param.Field[[]string]                             `json:"sources" format:"uuid"`
}

func (r PrismObjectContactQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectContactQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectContactQueryParamsQueryCombinator] `json:"combinator"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectContactQueryParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                      `json:"limit"`
	ListID param.Field[string]                                                     `json:"list_id" format:"uuid"`
	Page   param.Field[int64]                                                      `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]PrismObjectContactQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectContactQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectContactQueryParamsQueryCombinator string

const (
	PrismObjectContactQueryParamsQueryCombinatorAnd PrismObjectContactQueryParamsQueryCombinator = "AND"
	PrismObjectContactQueryParamsQueryCombinatorOr  PrismObjectContactQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectContactQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectContactQueryParamsQueryCombinatorAnd, PrismObjectContactQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

type PrismObjectContactQueryParamsQueryFilter struct {
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

func (r PrismObjectContactQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilter) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

// Satisfied by [PrismObjectContactQueryParamsQueryFilter],
// [PrismObjectContactQueryParamsQueryFilter],
// [PrismObjectContactQueryParamsQueryFilter],
// [PrismObjectContactQueryParamsQueryFilter],
// [PrismObjectContactQueryParamsQueryFilter],
// [PrismObjectContactQueryParamsQueryFilter],
// [PrismObjectContactQueryParamsQueryFilterLikeRegex],
// [PrismObjectContactQueryParamsQueryFilterBeginsWith],
// [PrismObjectContactQueryParamsQueryFilterEndsWith],
// [PrismObjectContactQueryParamsQueryFilterNotContains],
// [PrismObjectContactQueryParamsQueryFilterExists],
// [PrismObjectContactQueryParamsQueryFilterNotExists],
// [PrismObjectContactQueryParamsQueryFilterIn],
// [PrismObjectContactQueryParamsQueryFilterNotIn],
// [PrismObjectContactQueryParamsQueryFilter].
type PrismObjectContactQueryParamsQueryFilterUnion interface {
	implementsPrismObjectContactQueryParamsQueryFilterUnion()
}

type PrismObjectContactQueryParamsQueryFilter struct {
	Equals param.Field[PrismObjectContactQueryParamsQueryFilterUnion] `json:"=" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilter) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQueryFilterLikeRegex struct {
	LikeRegex param.Field[string] `json:"like_regex" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterLikeRegex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterLikeRegex) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQueryFilterBeginsWith struct {
	BeginsWith param.Field[string] `json:"begins_with" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterBeginsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterBeginsWith) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQueryFilterEndsWith struct {
	EndsWith param.Field[string] `json:"ends_with" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterEndsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterEndsWith) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQueryFilterNotContains struct {
	NotContains param.Field[string] `json:"not_contains" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterNotContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterNotContains) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQueryFilterExists struct {
	Exists param.Field[bool] `json:"exists" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterExists) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQueryFilterNotExists struct {
	NotExists param.Field[bool] `json:"not_exists" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterNotExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterNotExists) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQueryFilterIn struct {
	In param.Field[[]string] `json:"in" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterIn) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQueryFilterNotIn struct {
	NotIn param.Field[[]string] `json:"not_in" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterNotIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterNotIn) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQuerySort string

const (
	PrismObjectContactQueryParamsQuerySortAsc  PrismObjectContactQueryParamsQuerySort = "asc"
	PrismObjectContactQueryParamsQuerySortDesc PrismObjectContactQueryParamsQuerySort = "desc"
)

func (r PrismObjectContactQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectContactQueryParamsQuerySortAsc, PrismObjectContactQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectContactQueryParamsIDArray].
type PrismObjectContactQueryParamsIDUnion interface {
	ImplementsPrismObjectContactQueryParamsIDUnion()
}

type PrismObjectContactQueryParamsIDArray []string

func (r PrismObjectContactQueryParamsIDArray) ImplementsPrismObjectContactQueryParamsIDUnion() {}

type PrismObjectContactRestoreParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
