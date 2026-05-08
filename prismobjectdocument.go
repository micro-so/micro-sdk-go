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

// PrismObjectDocumentService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectDocumentService] method instead.
type PrismObjectDocumentService struct {
	Options []option.RequestOption
	Grant   *PrismObjectDocumentGrantService
}

// NewPrismObjectDocumentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectDocumentService(opts ...option.RequestOption) (r *PrismObjectDocumentService) {
	r = &PrismObjectDocumentService{}
	r.Options = opts
	r.Grant = NewPrismObjectDocumentGrantService(opts...)
	return
}

// Create object
func (r *PrismObjectDocumentService) New(ctx context.Context, params PrismObjectDocumentNewParams, opts ...option.RequestOption) (res *PrismObjectDocumentNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/document", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch object
func (r *PrismObjectDocumentService) Update(ctx context.Context, documentID string, params PrismObjectDocumentUpdateParams, opts ...option.RequestOption) (res *PrismObjectDocumentUpdateResponse, err error) {
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
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/document/%s", params.TeamID, documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Delete object
func (r *PrismObjectDocumentService) Delete(ctx context.Context, documentID string, body PrismObjectDocumentDeleteParams, opts ...option.RequestOption) (err error) {
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
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/document/%s", body.TeamID, documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: <100 records sync (immediate response), >=100 records
// async (S3/Lambda with WebSocket progress)
func (r *PrismObjectDocumentService) BulkNew(ctx context.Context, params PrismObjectDocumentBulkNewParams, opts ...option.RequestOption) (res *PrismObjectDocumentBulkNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/document/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Duplicate object
func (r *PrismObjectDocumentService) Duplicate(ctx context.Context, documentID string, body PrismObjectDocumentDuplicateParams, opts ...option.RequestOption) (res *PrismObjectDocumentDuplicateResponse, err error) {
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
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/document/%s/duplicate", body.TeamID, documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectDocumentService) Get(ctx context.Context, documentID string, query PrismObjectDocumentGetParams, opts ...option.RequestOption) (res *PrismObjectDocumentGetResponse, err error) {
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
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/document/%s", query.TeamID, documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Query v2
func (r *PrismObjectDocumentService) Query(ctx context.Context, params PrismObjectDocumentQueryParams, opts ...option.RequestOption) (res *PrismObjectDocumentQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/query/%s/document", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Restore object
func (r *PrismObjectDocumentService) Restore(ctx context.Context, documentID string, body PrismObjectDocumentRestoreParams, opts ...option.RequestOption) (res *PrismObjectDocumentRestoreResponse, err error) {
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
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/document/%s/restore", body.TeamID, documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDocumentNewResponse struct {
	ID  string      `json:"id" api:"required" format:"uuid"`
	CRM interface{} `json:"crm"`
	// Properties keyed by property slug.
	Default  map[string]interface{}             `json:"default"`
	Extended interface{}                        `json:"extended"`
	JSON     prismObjectDocumentNewResponseJSON `json:"-"`
}

// prismObjectDocumentNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentNewResponse]
type prismObjectDocumentNewResponseJSON struct {
	ID          apijson.Field
	CRM         apijson.Field
	Default     apijson.Field
	Extended    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentNewResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDocumentUpdateResponse struct {
	ID  string      `json:"id" api:"required" format:"uuid"`
	CRM interface{} `json:"crm"`
	// Properties keyed by property slug.
	Default  map[string]interface{}                `json:"default"`
	Extended interface{}                           `json:"extended"`
	JSON     prismObjectDocumentUpdateResponseJSON `json:"-"`
}

// prismObjectDocumentUpdateResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentUpdateResponse]
type prismObjectDocumentUpdateResponseJSON struct {
	ID          apijson.Field
	CRM         apijson.Field
	Default     apijson.Field
	Extended    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentBulkNewResponse struct {
	Results []PrismObjectDocumentBulkNewResponseResult `json:"results"`
	Status  PrismObjectDocumentBulkNewResponseStatus   `json:"status"`
	Summary PrismObjectDocumentBulkNewResponseSummary  `json:"summary"`
	JSON    prismObjectDocumentBulkNewResponseJSON     `json:"-"`
}

// prismObjectDocumentBulkNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentBulkNewResponse]
type prismObjectDocumentBulkNewResponseJSON struct {
	Results     apijson.Field
	Status      apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentBulkNewResponseResult struct {
	ID       string                                       `json:"id" api:"nullable" format:"uuid"`
	Created  bool                                         `json:"created"`
	Error    string                                       `json:"error"`
	Existing bool                                         `json:"existing"`
	JSON     prismObjectDocumentBulkNewResponseResultJSON `json:"-"`
}

// prismObjectDocumentBulkNewResponseResultJSON contains the JSON metadata for the
// struct [PrismObjectDocumentBulkNewResponseResult]
type prismObjectDocumentBulkNewResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentBulkNewResponseStatus string

const (
	PrismObjectDocumentBulkNewResponseStatusComplete PrismObjectDocumentBulkNewResponseStatus = "complete"
)

func (r PrismObjectDocumentBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectDocumentBulkNewResponseStatusComplete:
		return true
	}
	return false
}

type PrismObjectDocumentBulkNewResponseSummary struct {
	Created  int64                                         `json:"created"`
	Errors   int64                                         `json:"errors"`
	Existing int64                                         `json:"existing"`
	Total    int64                                         `json:"total"`
	JSON     prismObjectDocumentBulkNewResponseSummaryJSON `json:"-"`
}

// prismObjectDocumentBulkNewResponseSummaryJSON contains the JSON metadata for the
// struct [PrismObjectDocumentBulkNewResponseSummary]
type prismObjectDocumentBulkNewResponseSummaryJSON struct {
	Created     apijson.Field
	Errors      apijson.Field
	Existing    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkNewResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkNewResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentDuplicateResponse struct {
	ID   string                                   `json:"id" format:"uuid"`
	JSON prismObjectDocumentDuplicateResponseJSON `json:"-"`
}

// prismObjectDocumentDuplicateResponseJSON contains the JSON metadata for the
// struct [PrismObjectDocumentDuplicateResponse]
type prismObjectDocumentDuplicateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentDuplicateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentDuplicateResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDocumentGetResponse struct {
	ID  string      `json:"id" api:"required" format:"uuid"`
	CRM interface{} `json:"crm"`
	// Properties keyed by property slug.
	Default  map[string]interface{}             `json:"default"`
	Extended interface{}                        `json:"extended"`
	JSON     prismObjectDocumentGetResponseJSON `json:"-"`
}

// prismObjectDocumentGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentGetResponse]
type prismObjectDocumentGetResponseJSON struct {
	ID          apijson.Field
	CRM         apijson.Field
	Default     apijson.Field
	Extended    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentQueryResponse struct {
	Data []PrismObjectDocumentQueryResponseData `json:"data" api:"required"`
	// True when the page returned the maximum number of rows; another page may exist.
	HasMore bool                                 `json:"has_more"`
	JSON    prismObjectDocumentQueryResponseJSON `json:"-"`
}

// prismObjectDocumentQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentQueryResponse]
type prismObjectDocumentQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDocumentQueryResponseData struct {
	ID  string      `json:"id" api:"required" format:"uuid"`
	CRM interface{} `json:"crm"`
	// Properties keyed by property slug.
	Default  map[string]interface{}                   `json:"default"`
	Extended interface{}                              `json:"extended"`
	JSON     prismObjectDocumentQueryResponseDataJSON `json:"-"`
}

// prismObjectDocumentQueryResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectDocumentQueryResponseData]
type prismObjectDocumentQueryResponseDataJSON struct {
	ID          apijson.Field
	CRM         apijson.Field
	Default     apijson.Field
	Extended    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDocumentRestoreResponse struct {
	ID  string      `json:"id" api:"required" format:"uuid"`
	CRM interface{} `json:"crm"`
	// Properties keyed by property slug.
	Default  map[string]interface{}                 `json:"default"`
	Extended interface{}                            `json:"extended"`
	JSON     prismObjectDocumentRestoreResponseJSON `json:"-"`
}

// prismObjectDocumentRestoreResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentRestoreResponse]
type prismObjectDocumentRestoreResponseJSON struct {
	ID          apijson.Field
	CRM         apijson.Field
	Default     apijson.Field
	Extended    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentRestoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentRestoreResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismObjectDocumentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectDocumentUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismObjectDocumentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectDocumentDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectDocumentBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects param.Field[[]PrismObjectPropertiesParam]            `json:"objects" api:"required"`
	Options param.Field[PrismObjectDocumentBulkNewParamsOptions] `json:"options"`
}

func (r PrismObjectDocumentBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDocumentBulkNewParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// App/CRM ID for context (optional)
	CRMID param.Field[string] `json:"crm_id" format:"uuid"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
}

func (r PrismObjectDocumentBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDocumentDuplicateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectDocumentGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectDocumentQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                                `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[PrismObjectDocumentQueryParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[PrismObjectDocumentQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                              `json:"boxes"`
	Deleted param.Field[bool]                                  `json:"deleted"`
	Sources param.Field[[]string]                              `json:"sources" format:"uuid"`
}

func (r PrismObjectDocumentQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDocumentQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectDocumentQueryParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                                        `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]map[string]PrismObjectDocumentQueryParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                                  `json:"limit"`
	Page   param.Field[int64]                                                                  `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]PrismObjectDocumentQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectDocumentQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectDocumentQueryParamsQueryCombinator string

const (
	PrismObjectDocumentQueryParamsQueryCombinatorAnd PrismObjectDocumentQueryParamsQueryCombinator = "AND"
	PrismObjectDocumentQueryParamsQueryCombinatorOr  PrismObjectDocumentQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectDocumentQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectDocumentQueryParamsQueryCombinatorAnd, PrismObjectDocumentQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectDocumentQueryParamsQueryFilterArray].
type PrismObjectDocumentQueryParamsQueryFilterUnion interface {
	ImplementsPrismObjectDocumentQueryParamsQueryFilterUnion()
}

type PrismObjectDocumentQueryParamsQueryFilterArray []string

func (r PrismObjectDocumentQueryParamsQueryFilterArray) ImplementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

type PrismObjectDocumentQueryParamsQuerySort string

const (
	PrismObjectDocumentQueryParamsQuerySortAsc  PrismObjectDocumentQueryParamsQuerySort = "asc"
	PrismObjectDocumentQueryParamsQuerySortDesc PrismObjectDocumentQueryParamsQuerySort = "desc"
)

func (r PrismObjectDocumentQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectDocumentQueryParamsQuerySortAsc, PrismObjectDocumentQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectDocumentQueryParamsIDArray].
type PrismObjectDocumentQueryParamsIDUnion interface {
	ImplementsPrismObjectDocumentQueryParamsIDUnion()
}

type PrismObjectDocumentQueryParamsIDArray []string

func (r PrismObjectDocumentQueryParamsIDArray) ImplementsPrismObjectDocumentQueryParamsIDUnion() {}

type PrismObjectDocumentRestoreParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
