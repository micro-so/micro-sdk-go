// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/micro-so/micro-sdk-go/internal/apijson"
	"github.com/micro-so/micro-sdk-go/internal/apiquery"
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
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
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
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	if params.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
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

// Convenience list endpoint. Equivalent to
// `POST /v2/prism/{teamId}/{objectType}/query` with an empty body, plus
// query-string sugar for the common cases. Any unrecognized query parameter is
// interpreted as an equality filter on a property of that name; pass arrays for
// `in`. Values are received as strings, so non-string property filters via this
// endpoint may not work — use the `query` endpoint for typed comparisons or
// anything beyond simple equality.
func (r *PrismObjectOrganizationService) List(ctx context.Context, params PrismObjectOrganizationListParams, opts ...option.RequestOption) (res *PrismObjectOrganizationListResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete object
func (r *PrismObjectOrganizationService) Delete(ctx context.Context, organizationID string, params PrismObjectOrganizationDeleteParams, opts ...option.RequestOption) (err error) {
	if params.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: small batches complete synchronously and return 200 with
// the final `ImportJob`; large batches start an async job, return 202 with
// `status: processing` and a `Location` header, and can be polled via
// `GET /v2/prism/{teamId}/imports/{jobId}`.
func (r *PrismObjectOrganizationService) BulkNew(ctx context.Context, params PrismObjectOrganizationBulkNewParams, opts ...option.RequestOption) (res *PrismObjectOrganizationBulkNewResponse, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
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

// Soft-delete up to 100 records in a single call. Same partial-success contract as
// batch/update.
func (r *PrismObjectOrganizationService) BulkDelete(ctx context.Context, params PrismObjectOrganizationBulkDeleteParams, opts ...option.RequestOption) (res *PrismObjectOrganizationBulkDeleteResponse, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
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
	path := fmt.Sprintf("v2/prism/%s/organization/batch/delete", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch up to 100 records in a single call. Each item is attempted independently —
// failures don't abort the batch. Inspect `results[].status` per item.
func (r *PrismObjectOrganizationService) BulkUpdate(ctx context.Context, params PrismObjectOrganizationBulkUpdateParams, opts ...option.RequestOption) (res *PrismObjectOrganizationBulkUpdateResponse, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
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
	path := fmt.Sprintf("v2/prism/%s/organization/batch/update", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the total number of records of this object type that the caller can see.
// Avoids the page-overshoot anti-pattern — clients no longer need to keep paging
// until `has_more` flips false to discover the total. Currently does not apply
// query filters; for a filtered total, pass `include_total: true` in a POST
// `/query` body.
func (r *PrismObjectOrganizationService) Count(ctx context.Context, params PrismObjectOrganizationCountParams, opts ...option.RequestOption) (res *PrismObjectOrganizationCountResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/organization/count", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Duplicate object
func (r *PrismObjectOrganizationService) Duplicate(ctx context.Context, organizationID string, params PrismObjectOrganizationDuplicateParams, opts ...option.RequestOption) (res *PrismObjectOrganizationDuplicateResponse, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
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
	path := fmt.Sprintf("v2/prism/%s/organization/%s/duplicate", params.TeamID, organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns the single record whose property `{slug}` equals `{value}`. 404 if
// nothing matches; 409 if more than one record matches.
func (r *PrismObjectOrganizationService) Find(ctx context.Context, slug string, value string, params PrismObjectOrganizationFindParams, opts ...option.RequestOption) (res *PrismObjectOrganizationFindResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/organization/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectOrganizationService) Get(ctx context.Context, organizationID string, params PrismObjectOrganizationGetParams, opts ...option.RequestOption) (res *PrismObjectOrganizationGetResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Query
func (r *PrismObjectOrganizationService) Query(ctx context.Context, params PrismObjectOrganizationQueryParams, opts ...option.RequestOption) (res *PrismObjectOrganizationQueryResponse, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
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
	path := fmt.Sprintf("v2/prism/%s/organization/query", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Restore object
func (r *PrismObjectOrganizationService) Restore(ctx context.Context, organizationID string, params PrismObjectOrganizationRestoreParams, opts ...option.RequestOption) (res *PrismObjectOrganizationRestoreResponse, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
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
	path := fmt.Sprintf("v2/prism/%s/organization/%s/restore", params.TeamID, organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Idempotent create-or-update keyed on `{slug}={value}`. If exactly one record
// matches, it is patched and 200 is returned. If none match, a new record is
// created (with the lookup property set if absent) and 201 is returned. If
// multiple records match, 409 is returned and you should patch by id instead.
func (r *PrismObjectOrganizationService) Upsert(ctx context.Context, slug string, value string, params PrismObjectOrganizationUpsertParams, opts ...option.RequestOption) (res *PrismObjectOrganizationUpsertResponse, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
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
	path := fmt.Sprintf("v2/prism/%s/organization/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
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

type PrismObjectOrganizationListResponse struct {
	Data []PrismObjectOrganizationListResponseData `json:"data" api:"required"`
	// Accurate end-of-data signal — false on the last page, never forces clients to
	// overshoot.
	HasMore    bool   `json:"has_more" api:"required"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// Populated only when `?include_total=true` was passed.
	Total int64                                   `json:"total" api:"nullable"`
	JSON  prismObjectOrganizationListResponseJSON `json:"-"`
}

// prismObjectOrganizationListResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationListResponse]
type prismObjectOrganizationListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationListResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectOrganizationListResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                      `json:"properties"`
	Source     []string                                    `json:"source" api:"nullable"`
	JSON       prismObjectOrganizationListResponseDataJSON `json:"-"`
}

// prismObjectOrganizationListResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationListResponseData]
type prismObjectOrganizationListResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectOrganizationListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Status snapshot of an import job. Same shape used by the POST /import response
// and by GET /imports/{jobId}.
type PrismObjectOrganizationBulkNewResponse struct {
	// Null for sync imports (results inlined). Set for async imports.
	JobID  string                                       `json:"job_id" api:"required,nullable"`
	Status PrismObjectOrganizationBulkNewResponseStatus `json:"status" api:"required"`
	// Total number of rows in the import.
	Total     int64     `json:"total" api:"required"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Set when status=failed; describes the job-level failure (not per-row).
	Error     PrismObjectOrganizationBulkNewResponseError `json:"error"`
	ExpiresAt time.Time                                   `json:"expires_at" format:"date-time"`
	Failed    int64                                       `json:"failed"`
	// Rows that have been attempted (succeeded + failed).
	Processed int64 `json:"processed"`
	// Per-row outcomes. Always present for sync imports; populated for async imports
	// once the job reaches `complete`.
	Results   []PrismObjectOrganizationBulkNewResponseResult `json:"results"`
	Succeeded int64                                          `json:"succeeded"`
	UpdatedAt time.Time                                      `json:"updated_at" format:"date-time"`
	JSON      prismObjectOrganizationBulkNewResponseJSON     `json:"-"`
}

// prismObjectOrganizationBulkNewResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationBulkNewResponse]
type prismObjectOrganizationBulkNewResponseJSON struct {
	JobID       apijson.Field
	Status      apijson.Field
	Total       apijson.Field
	CreatedAt   apijson.Field
	Error       apijson.Field
	ExpiresAt   apijson.Field
	Failed      apijson.Field
	Processed   apijson.Field
	Results     apijson.Field
	Succeeded   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationBulkNewResponseStatus string

const (
	PrismObjectOrganizationBulkNewResponseStatusComplete   PrismObjectOrganizationBulkNewResponseStatus = "complete"
	PrismObjectOrganizationBulkNewResponseStatusProcessing PrismObjectOrganizationBulkNewResponseStatus = "processing"
	PrismObjectOrganizationBulkNewResponseStatusFailed     PrismObjectOrganizationBulkNewResponseStatus = "failed"
)

func (r PrismObjectOrganizationBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectOrganizationBulkNewResponseStatusComplete, PrismObjectOrganizationBulkNewResponseStatusProcessing, PrismObjectOrganizationBulkNewResponseStatusFailed:
		return true
	}
	return false
}

// Set when status=failed; describes the job-level failure (not per-row).
type PrismObjectOrganizationBulkNewResponseError struct {
	Code    string                                          `json:"code"`
	Message string                                          `json:"message"`
	JSON    prismObjectOrganizationBulkNewResponseErrorJSON `json:"-"`
}

// prismObjectOrganizationBulkNewResponseErrorJSON contains the JSON metadata for
// the struct [PrismObjectOrganizationBulkNewResponseError]
type prismObjectOrganizationBulkNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationBulkNewResponseResult struct {
	ID      string                                             `json:"id" api:"nullable" format:"uuid"`
	Created bool                                               `json:"created"`
	Error   PrismObjectOrganizationBulkNewResponseResultsError `json:"error"`
	// True if the row matched an existing record via the dedupe key.
	Existing bool `json:"existing"`
	// Zero-based position of this row in the request.
	InputIndex int64 `json:"input_index"`
	// True if a matching record was updated.
	Updated bool                                             `json:"updated"`
	JSON    prismObjectOrganizationBulkNewResponseResultJSON `json:"-"`
}

// prismObjectOrganizationBulkNewResponseResultJSON contains the JSON metadata for
// the struct [PrismObjectOrganizationBulkNewResponseResult]
type prismObjectOrganizationBulkNewResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	InputIndex  apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationBulkNewResponseResultsError struct {
	Code    string                                                 `json:"code"`
	Message string                                                 `json:"message"`
	JSON    prismObjectOrganizationBulkNewResponseResultsErrorJSON `json:"-"`
}

// prismObjectOrganizationBulkNewResponseResultsErrorJSON contains the JSON
// metadata for the struct [PrismObjectOrganizationBulkNewResponseResultsError]
type prismObjectOrganizationBulkNewResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkNewResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkNewResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectOrganizationBulkDeleteResponse struct {
	Results []PrismObjectOrganizationBulkDeleteResponseResult `json:"results" api:"required"`
	Summary PrismObjectOrganizationBulkDeleteResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectOrganizationBulkDeleteResponseJSON     `json:"-"`
}

// prismObjectOrganizationBulkDeleteResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationBulkDeleteResponse]
type prismObjectOrganizationBulkDeleteResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationBulkDeleteResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                                 `json:"id" api:"required,nullable"`
	Status PrismObjectOrganizationBulkDeleteResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectOrganizationBulkDeleteResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectOrganizationBulkDeleteResponseResultsRecord `json:"record"`
	JSON   prismObjectOrganizationBulkDeleteResponseResultJSON    `json:"-"`
}

// prismObjectOrganizationBulkDeleteResponseResultJSON contains the JSON metadata
// for the struct [PrismObjectOrganizationBulkDeleteResponseResult]
type prismObjectOrganizationBulkDeleteResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationBulkDeleteResponseResultsStatus string

const (
	PrismObjectOrganizationBulkDeleteResponseResultsStatusOk    PrismObjectOrganizationBulkDeleteResponseResultsStatus = "ok"
	PrismObjectOrganizationBulkDeleteResponseResultsStatusError PrismObjectOrganizationBulkDeleteResponseResultsStatus = "error"
)

func (r PrismObjectOrganizationBulkDeleteResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectOrganizationBulkDeleteResponseResultsStatusOk, PrismObjectOrganizationBulkDeleteResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectOrganizationBulkDeleteResponseResultsError struct {
	Code    string                                                    `json:"code"`
	Message string                                                    `json:"message"`
	JSON    prismObjectOrganizationBulkDeleteResponseResultsErrorJSON `json:"-"`
}

// prismObjectOrganizationBulkDeleteResponseResultsErrorJSON contains the JSON
// metadata for the struct [PrismObjectOrganizationBulkDeleteResponseResultsError]
type prismObjectOrganizationBulkDeleteResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkDeleteResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkDeleteResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectOrganizationBulkDeleteResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                                     `json:"default"`
	List    interface{}                                                `json:"list"`
	JSON    prismObjectOrganizationBulkDeleteResponseResultsRecordJSON `json:"-"`
}

// prismObjectOrganizationBulkDeleteResponseResultsRecordJSON contains the JSON
// metadata for the struct [PrismObjectOrganizationBulkDeleteResponseResultsRecord]
type prismObjectOrganizationBulkDeleteResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkDeleteResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkDeleteResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationBulkDeleteResponseSummary struct {
	Failed    int64                                                `json:"failed" api:"required"`
	Succeeded int64                                                `json:"succeeded" api:"required"`
	Total     int64                                                `json:"total" api:"required"`
	JSON      prismObjectOrganizationBulkDeleteResponseSummaryJSON `json:"-"`
}

// prismObjectOrganizationBulkDeleteResponseSummaryJSON contains the JSON metadata
// for the struct [PrismObjectOrganizationBulkDeleteResponseSummary]
type prismObjectOrganizationBulkDeleteResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkDeleteResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkDeleteResponseSummaryJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectOrganizationBulkUpdateResponse struct {
	Results []PrismObjectOrganizationBulkUpdateResponseResult `json:"results" api:"required"`
	Summary PrismObjectOrganizationBulkUpdateResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectOrganizationBulkUpdateResponseJSON     `json:"-"`
}

// prismObjectOrganizationBulkUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationBulkUpdateResponse]
type prismObjectOrganizationBulkUpdateResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationBulkUpdateResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                                 `json:"id" api:"required,nullable"`
	Status PrismObjectOrganizationBulkUpdateResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectOrganizationBulkUpdateResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectOrganizationBulkUpdateResponseResultsRecord `json:"record"`
	JSON   prismObjectOrganizationBulkUpdateResponseResultJSON    `json:"-"`
}

// prismObjectOrganizationBulkUpdateResponseResultJSON contains the JSON metadata
// for the struct [PrismObjectOrganizationBulkUpdateResponseResult]
type prismObjectOrganizationBulkUpdateResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationBulkUpdateResponseResultsStatus string

const (
	PrismObjectOrganizationBulkUpdateResponseResultsStatusOk    PrismObjectOrganizationBulkUpdateResponseResultsStatus = "ok"
	PrismObjectOrganizationBulkUpdateResponseResultsStatusError PrismObjectOrganizationBulkUpdateResponseResultsStatus = "error"
)

func (r PrismObjectOrganizationBulkUpdateResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectOrganizationBulkUpdateResponseResultsStatusOk, PrismObjectOrganizationBulkUpdateResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectOrganizationBulkUpdateResponseResultsError struct {
	Code    string                                                    `json:"code"`
	Message string                                                    `json:"message"`
	JSON    prismObjectOrganizationBulkUpdateResponseResultsErrorJSON `json:"-"`
}

// prismObjectOrganizationBulkUpdateResponseResultsErrorJSON contains the JSON
// metadata for the struct [PrismObjectOrganizationBulkUpdateResponseResultsError]
type prismObjectOrganizationBulkUpdateResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkUpdateResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkUpdateResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectOrganizationBulkUpdateResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                                     `json:"default"`
	List    interface{}                                                `json:"list"`
	JSON    prismObjectOrganizationBulkUpdateResponseResultsRecordJSON `json:"-"`
}

// prismObjectOrganizationBulkUpdateResponseResultsRecordJSON contains the JSON
// metadata for the struct [PrismObjectOrganizationBulkUpdateResponseResultsRecord]
type prismObjectOrganizationBulkUpdateResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkUpdateResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkUpdateResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationBulkUpdateResponseSummary struct {
	Failed    int64                                                `json:"failed" api:"required"`
	Succeeded int64                                                `json:"succeeded" api:"required"`
	Total     int64                                                `json:"total" api:"required"`
	JSON      prismObjectOrganizationBulkUpdateResponseSummaryJSON `json:"-"`
}

// prismObjectOrganizationBulkUpdateResponseSummaryJSON contains the JSON metadata
// for the struct [PrismObjectOrganizationBulkUpdateResponseSummary]
type prismObjectOrganizationBulkUpdateResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationBulkUpdateResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationBulkUpdateResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationCountResponse struct {
	// Number of records matching the access scope.
	Total int64                                    `json:"total" api:"required"`
	JSON  prismObjectOrganizationCountResponseJSON `json:"-"`
}

// prismObjectOrganizationCountResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationCountResponse]
type prismObjectOrganizationCountResponseJSON struct {
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationCountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationCountResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectOrganizationDuplicateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                       `json:"default"`
	List    interface{}                                  `json:"list"`
	JSON    prismObjectOrganizationDuplicateResponseJSON `json:"-"`
}

// prismObjectOrganizationDuplicateResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationDuplicateResponse]
type prismObjectOrganizationDuplicateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
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
type PrismObjectOrganizationFindResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                  `json:"default"`
	List    interface{}                             `json:"list"`
	JSON    prismObjectOrganizationFindResponseJSON `json:"-"`
}

// prismObjectOrganizationFindResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationFindResponse]
type prismObjectOrganizationFindResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationFindResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationFindResponseJSON) RawJSON() string {
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
	Total int64                                    `json:"total" api:"nullable"`
	JSON  prismObjectOrganizationQueryResponseJSON `json:"-"`
}

// prismObjectOrganizationQueryResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationQueryResponse]
type prismObjectOrganizationQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectOrganizationQueryResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                       `json:"properties"`
	Source     []string                                     `json:"source" api:"nullable"`
	JSON       prismObjectOrganizationQueryResponseDataJSON `json:"-"`
}

// prismObjectOrganizationQueryResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationQueryResponseData]
type prismObjectOrganizationQueryResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
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

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectOrganizationUpsertResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                    `json:"default"`
	List    interface{}                               `json:"list"`
	JSON    prismObjectOrganizationUpsertResponseJSON `json:"-"`
}

// prismObjectOrganizationUpsertResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationUpsertResponse]
type prismObjectOrganizationUpsertResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationUpsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationUpsertResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
}

func (r PrismObjectOrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectOrganizationUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
	IfMatch               param.Field[string]        `header:"If-Match"`
}

func (r PrismObjectOrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectOrganizationListParams struct {
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

// URLQuery serializes [PrismObjectOrganizationListParams]'s query parameters as
// `url.Values`.
func (r PrismObjectOrganizationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectOrganizationDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IfMatch param.Field[string] `header:"If-Match"`
}

type PrismObjectOrganizationBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects        param.Field[[]PrismObjectPropertiesParam]                `json:"objects" api:"required"`
	Options        param.Field[PrismObjectOrganizationBulkNewParamsOptions] `json:"options"`
	IdempotencyKey param.Field[string]                                      `header:"Idempotency-Key"`
}

func (r PrismObjectOrganizationBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectOrganizationBulkNewParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// When true, unknown values for select/multiselect properties are created as new
	// options instead of failing the import
	CreateMissingOptions param.Field[bool] `json:"create_missing_options"`
	// Deprecated alias for list_id.
	//
	// Deprecated: deprecated
	CRMID param.Field[string] `json:"crm_id" format:"uuid"`
	// Property slug to deduplicate on. A single-element array is also accepted;
	// compound (multi-slug) dedupe is not supported yet and is rejected with guidance.
	DedupeBy param.Field[PrismObjectOrganizationBulkNewParamsOptionsDedupeByUnion] `json:"dedupe_by"`
	// App/CRM ID for context (optional)
	ListID param.Field[string] `json:"list_id" format:"uuid"`
	// Require app_stage for every row in the selected list. app_stage is a reserved
	// list-scoped alias for native status.
	RequireListStage param.Field[bool] `json:"require_list_stage"`
	// Patch a deduplicated record with the supplied properties instead of skipping it.
	UpdateExisting param.Field[bool] `json:"update_existing"`
}

func (r PrismObjectOrganizationBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Property slug to deduplicate on. A single-element array is also accepted;
// compound (multi-slug) dedupe is not supported yet and is rejected with guidance.
//
// Satisfied by [shared.UnionString],
// [PrismObjectOrganizationBulkNewParamsOptionsDedupeByArray].
type PrismObjectOrganizationBulkNewParamsOptionsDedupeByUnion interface {
	ImplementsPrismObjectOrganizationBulkNewParamsOptionsDedupeByUnion()
}

type PrismObjectOrganizationBulkNewParamsOptionsDedupeByArray []string

func (r PrismObjectOrganizationBulkNewParamsOptionsDedupeByArray) ImplementsPrismObjectOrganizationBulkNewParamsOptionsDedupeByUnion() {
}

type PrismObjectOrganizationBulkDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]   `path:"teamId" api:"required" format:"uuid"`
	IDs            param.Field[[]string] `json:"ids" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string]   `header:"Idempotency-Key"`
}

func (r PrismObjectOrganizationBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectOrganizationBulkUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]                                        `path:"teamId" api:"required" format:"uuid"`
	Items          param.Field[[]PrismObjectOrganizationBulkUpdateParamsItem] `json:"items" api:"required"`
	IdempotencyKey param.Field[string]                                        `header:"Idempotency-Key"`
}

func (r PrismObjectOrganizationBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object with `id` plus the same property body shape as PATCH
// (`default`/`list`/`extended`).
type PrismObjectOrganizationBulkUpdateParamsItem struct {
	ID          param.Field[string]    `json:"id" api:"required" format:"uuid"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r PrismObjectOrganizationBulkUpdateParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectOrganizationCountParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the count to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectOrganizationCountParams]'s query parameters as
// `url.Values`.
func (r PrismObjectOrganizationCountParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectOrganizationDuplicateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectOrganizationFindParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the lookup to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectOrganizationFindParams]'s query parameters as
// `url.Values`.
func (r PrismObjectOrganizationFindParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectOrganizationGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Comma-separated property slugs to return. Use dot notation for relationships.
	// `id` is always returned at the top level. Defaults to all properties.
	Select param.Field[string] `query:"select"`
}

// URLQuery serializes [PrismObjectOrganizationGetParams]'s query parameters as
// `url.Values`.
func (r PrismObjectOrganizationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectOrganizationQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string]                                    `path:"teamId" api:"required" format:"uuid"`
	Query  param.Field[PrismObjectOrganizationQueryParamsQuery]   `json:"query" api:"required"`
	ID     param.Field[PrismObjectOrganizationQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes  param.Field[[]string]                                  `json:"boxes"`
	// Alternative location for the opaque cursor (a sibling of `query`). Use whichever
	// feels more natural; if both are present, `query.cursor` wins.
	Cursor  param.Field[string] `json:"cursor"`
	Deleted param.Field[bool]   `json:"deleted"`
	// When true, the response includes a `total` field with the unpaginated row count.
	// Costs an additional pass over the result set — for unfiltered totals prefer
	// `GET /v2/prism/{teamId}/{objectType}/count` instead.
	IncludeTotal   param.Field[bool]     `json:"include_total"`
	Sources        param.Field[[]string] `json:"sources" format:"uuid"`
	IdempotencyKey param.Field[string]   `header:"Idempotency-Key"`
}

func (r PrismObjectOrganizationQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectOrganizationQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name). `id` is always returned at the top level of each
	// row and does not need to be selected.
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectOrganizationQueryParamsQueryCombinator] `json:"combinator"`
	// Opaque cursor from a previous response's `next_cursor`. Pass it back unchanged
	// to fetch the next page. When set, `page` and `limit` are derived from the cursor
	// and any explicit values are ignored.
	Cursor param.Field[string] `json:"cursor"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectOrganizationQueryParamsQueryFilterUnion] `json:"filter"`
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

func (r PrismObjectOrganizationQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilter) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

// Satisfied by [PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterEq],
// [PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterNe],
// [PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterLt],
// [PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterGt],
// [PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterLte],
// [PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterGte],
// [PrismObjectOrganizationQueryParamsQueryFilterContains],
// [PrismObjectOrganizationQueryParamsQueryFilterBeginsWith],
// [PrismObjectOrganizationQueryParamsQueryFilterEndsWith],
// [PrismObjectOrganizationQueryParamsQueryFilterNotContains],
// [PrismObjectOrganizationQueryParamsQueryFilterExists],
// [PrismObjectOrganizationQueryParamsQueryFilterNotExists],
// [PrismObjectOrganizationQueryParamsQueryFilterIsNull],
// [PrismObjectOrganizationQueryParamsQueryFilterIsNotNull],
// [PrismObjectOrganizationQueryParamsQueryFilterBetween],
// [PrismObjectOrganizationQueryParamsQueryFilterIn],
// [PrismObjectOrganizationQueryParamsQueryFilterNotIn],
// [PrismObjectOrganizationQueryParamsQueryFilter].
type PrismObjectOrganizationQueryParamsQueryFilterUnion interface {
	implementsPrismObjectOrganizationQueryParamsQueryFilterUnion()
}

type PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterEq struct {
	Equals param.Field[PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion] `json:"=" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterEq) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterEq) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion interface {
	ImplementsPrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion()
}

type PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterNe struct {
	NotEquals param.Field[PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion] `json:"!=" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterNe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterNe) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion interface {
	ImplementsPrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion()
}

type PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterLt struct {
	Less param.Field[string] `json:"<" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterLt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterLt) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterGt struct {
	Greater param.Field[string] `json:">" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterGt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterGt) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterLte struct {
	LessOrEquals param.Field[string] `json:"<=" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterLte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterLte) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterGte struct {
	GreaterOrEquals param.Field[string] `json:">=" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterGte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterPrismQueryFilterGte) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterContains struct {
	Contains param.Field[PrismObjectOrganizationQueryParamsQueryFilterContainsContainsUnion] `json:"contains" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterContains) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectOrganizationQueryParamsQueryFilterContainsContainsArray].
type PrismObjectOrganizationQueryParamsQueryFilterContainsContainsUnion interface {
	ImplementsPrismObjectOrganizationQueryParamsQueryFilterContainsContainsUnion()
}

type PrismObjectOrganizationQueryParamsQueryFilterContainsContainsArray []string

func (r PrismObjectOrganizationQueryParamsQueryFilterContainsContainsArray) ImplementsPrismObjectOrganizationQueryParamsQueryFilterContainsContainsUnion() {
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

type PrismObjectOrganizationQueryParamsQueryFilterIsNull struct {
	IsNull param.Field[PrismObjectOrganizationQueryParamsQueryFilterIsNullIsNullUnion] `json:"is_null" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterIsNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterIsNull) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectOrganizationQueryParamsQueryFilterIsNullIsNullArray].
type PrismObjectOrganizationQueryParamsQueryFilterIsNullIsNullUnion interface {
	ImplementsPrismObjectOrganizationQueryParamsQueryFilterIsNullIsNullUnion()
}

type PrismObjectOrganizationQueryParamsQueryFilterIsNullIsNullArray []string

func (r PrismObjectOrganizationQueryParamsQueryFilterIsNullIsNullArray) ImplementsPrismObjectOrganizationQueryParamsQueryFilterIsNullIsNullUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterIsNotNull struct {
	IsNotNull param.Field[PrismObjectOrganizationQueryParamsQueryFilterIsNotNullIsNotNullUnion] `json:"is_not_null" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterIsNotNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterIsNotNull) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectOrganizationQueryParamsQueryFilterIsNotNullIsNotNullArray].
type PrismObjectOrganizationQueryParamsQueryFilterIsNotNullIsNotNullUnion interface {
	ImplementsPrismObjectOrganizationQueryParamsQueryFilterIsNotNullIsNotNullUnion()
}

type PrismObjectOrganizationQueryParamsQueryFilterIsNotNullIsNotNullArray []string

func (r PrismObjectOrganizationQueryParamsQueryFilterIsNotNullIsNotNullArray) ImplementsPrismObjectOrganizationQueryParamsQueryFilterIsNotNullIsNotNullUnion() {
}

type PrismObjectOrganizationQueryParamsQueryFilterBetween struct {
	Between param.Field[PrismObjectOrganizationQueryParamsQueryFilterBetweenBetweenUnion] `json:"between" api:"required"`
}

func (r PrismObjectOrganizationQueryParamsQueryFilterBetween) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectOrganizationQueryParamsQueryFilterBetween) implementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectOrganizationQueryParamsQueryFilterBetweenBetweenArray].
type PrismObjectOrganizationQueryParamsQueryFilterBetweenBetweenUnion interface {
	ImplementsPrismObjectOrganizationQueryParamsQueryFilterBetweenBetweenUnion()
}

type PrismObjectOrganizationQueryParamsQueryFilterBetweenBetweenArray []string

func (r PrismObjectOrganizationQueryParamsQueryFilterBetweenBetweenArray) ImplementsPrismObjectOrganizationQueryParamsQueryFilterBetweenBetweenUnion() {
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
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectOrganizationUpsertParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	// Scope the upsert to a specific list/app. Required to match or write list-scoped
	// properties, including `app_stage`.
	ListID         param.Field[string] `query:"list_id" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

func (r PrismObjectOrganizationUpsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

// URLQuery serializes [PrismObjectOrganizationUpsertParams]'s query parameters as
// `url.Values`.
func (r PrismObjectOrganizationUpsertParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
