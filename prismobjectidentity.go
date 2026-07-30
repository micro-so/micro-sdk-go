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

// PrismObjectIdentityService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectIdentityService] method instead.
type PrismObjectIdentityService struct {
	Options []option.RequestOption
}

// NewPrismObjectIdentityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectIdentityService(opts ...option.RequestOption) (r *PrismObjectIdentityService) {
	r = &PrismObjectIdentityService{}
	r.Options = opts
	return
}

// Create object
func (r *PrismObjectIdentityService) New(ctx context.Context, params PrismObjectIdentityNewParams, opts ...option.RequestOption) (res *PrismObjectIdentityNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/identity", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch object
func (r *PrismObjectIdentityService) Update(ctx context.Context, identityID string, params PrismObjectIdentityUpdateParams, opts ...option.RequestOption) (res *PrismObjectIdentityUpdateResponse, err error) {
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
	if identityID == "" {
		err = errors.New("missing required identityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/identity/%s", params.TeamID, identityID)
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
func (r *PrismObjectIdentityService) List(ctx context.Context, params PrismObjectIdentityListParams, opts ...option.RequestOption) (res *PrismObjectIdentityListResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete object
func (r *PrismObjectIdentityService) Delete(ctx context.Context, identityID string, params PrismObjectIdentityDeleteParams, opts ...option.RequestOption) (err error) {
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
	if identityID == "" {
		err = errors.New("missing required identityId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/identity/%s", params.TeamID, identityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: small batches complete synchronously and return 200 with
// the final `ImportJob`; large batches start an async job, return 202 with
// `status: processing` and a `Location` header, and can be polled via
// `GET /v2/prism/{teamId}/imports/{jobId}`.
func (r *PrismObjectIdentityService) BulkNew(ctx context.Context, params PrismObjectIdentityBulkNewParams, opts ...option.RequestOption) (res *PrismObjectIdentityBulkNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/identity/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Soft-delete up to 100 records in a single call. Same partial-success contract as
// batch/update.
func (r *PrismObjectIdentityService) BulkDelete(ctx context.Context, params PrismObjectIdentityBulkDeleteParams, opts ...option.RequestOption) (res *PrismObjectIdentityBulkDeleteResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/identity/batch/delete", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch up to 100 records in a single call. Each item is attempted independently —
// failures don't abort the batch. Inspect `results[].status` per item.
func (r *PrismObjectIdentityService) BulkUpdate(ctx context.Context, params PrismObjectIdentityBulkUpdateParams, opts ...option.RequestOption) (res *PrismObjectIdentityBulkUpdateResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/identity/batch/update", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the total number of records of this object type that the caller can see.
// Avoids the page-overshoot anti-pattern — clients no longer need to keep paging
// until `has_more` flips false to discover the total. Currently does not apply
// query filters; for a filtered total, pass `include_total: true` in a POST
// `/query` body.
func (r *PrismObjectIdentityService) Count(ctx context.Context, params PrismObjectIdentityCountParams, opts ...option.RequestOption) (res *PrismObjectIdentityCountResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/identity/count", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Duplicate object
func (r *PrismObjectIdentityService) Duplicate(ctx context.Context, identityID string, params PrismObjectIdentityDuplicateParams, opts ...option.RequestOption) (res *PrismObjectIdentityDuplicateResponse, err error) {
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
	if identityID == "" {
		err = errors.New("missing required identityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/identity/%s/duplicate", params.TeamID, identityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns the single record whose property `{slug}` equals `{value}`. 404 if
// nothing matches; 409 if more than one record matches.
func (r *PrismObjectIdentityService) Find(ctx context.Context, slug string, value string, params PrismObjectIdentityFindParams, opts ...option.RequestOption) (res *PrismObjectIdentityFindResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/identity/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectIdentityService) Get(ctx context.Context, identityID string, params PrismObjectIdentityGetParams, opts ...option.RequestOption) (res *PrismObjectIdentityGetResponse, err error) {
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
	if identityID == "" {
		err = errors.New("missing required identityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/identity/%s", params.TeamID, identityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Query
func (r *PrismObjectIdentityService) Query(ctx context.Context, params PrismObjectIdentityQueryParams, opts ...option.RequestOption) (res *PrismObjectIdentityQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/identity/query", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Restore object
func (r *PrismObjectIdentityService) Restore(ctx context.Context, identityID string, params PrismObjectIdentityRestoreParams, opts ...option.RequestOption) (res *PrismObjectIdentityRestoreResponse, err error) {
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
	if identityID == "" {
		err = errors.New("missing required identityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/identity/%s/restore", params.TeamID, identityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Idempotent create-or-update keyed on `{slug}={value}`. If exactly one record
// matches, it is patched and 200 is returned. If none match, a new record is
// created (with the lookup property set if absent) and 201 is returned. If
// multiple records match, 409 is returned and you should patch by id instead.
func (r *PrismObjectIdentityService) Upsert(ctx context.Context, slug string, value string, params PrismObjectIdentityUpsertParams, opts ...option.RequestOption) (res *PrismObjectIdentityUpsertResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/identity/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityNewResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}             `json:"default"`
	List    interface{}                        `json:"list"`
	JSON    prismObjectIdentityNewResponseJSON `json:"-"`
}

// prismObjectIdentityNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityNewResponse]
type prismObjectIdentityNewResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityNewResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityUpdateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                `json:"default"`
	List    interface{}                           `json:"list"`
	JSON    prismObjectIdentityUpdateResponseJSON `json:"-"`
}

// prismObjectIdentityUpdateResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityUpdateResponse]
type prismObjectIdentityUpdateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityListResponse struct {
	Data []PrismObjectIdentityListResponseData `json:"data" api:"required"`
	// Accurate end-of-data signal — false on the last page, never forces clients to
	// overshoot.
	HasMore    bool   `json:"has_more" api:"required"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// Populated only when `?include_total=true` was passed.
	Total int64                               `json:"total" api:"nullable"`
	JSON  prismObjectIdentityListResponseJSON `json:"-"`
}

// prismObjectIdentityListResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityListResponse]
type prismObjectIdentityListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityListResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectIdentityListResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                  `json:"properties"`
	Source     []string                                `json:"source" api:"nullable"`
	JSON       prismObjectIdentityListResponseDataJSON `json:"-"`
}

// prismObjectIdentityListResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectIdentityListResponseData]
type prismObjectIdentityListResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectIdentityListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Status snapshot of an import job. Same shape used by the POST /import response
// and by GET /imports/{jobId}.
type PrismObjectIdentityBulkNewResponse struct {
	// Null for sync imports (results inlined). Set for async imports.
	JobID  string                                   `json:"job_id" api:"required,nullable"`
	Status PrismObjectIdentityBulkNewResponseStatus `json:"status" api:"required"`
	// Total number of rows in the import.
	Total     int64     `json:"total" api:"required"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Set when status=failed; describes the job-level failure (not per-row).
	Error     PrismObjectIdentityBulkNewResponseError `json:"error"`
	ExpiresAt time.Time                               `json:"expires_at" format:"date-time"`
	Failed    int64                                   `json:"failed"`
	// Rows that have been attempted (succeeded + failed).
	Processed int64 `json:"processed"`
	// Per-row outcomes. Always present for sync imports; populated for async imports
	// once the job reaches `complete`.
	Results   []PrismObjectIdentityBulkNewResponseResult `json:"results"`
	Succeeded int64                                      `json:"succeeded"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      prismObjectIdentityBulkNewResponseJSON     `json:"-"`
}

// prismObjectIdentityBulkNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityBulkNewResponse]
type prismObjectIdentityBulkNewResponseJSON struct {
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

func (r *PrismObjectIdentityBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityBulkNewResponseStatus string

const (
	PrismObjectIdentityBulkNewResponseStatusComplete   PrismObjectIdentityBulkNewResponseStatus = "complete"
	PrismObjectIdentityBulkNewResponseStatusProcessing PrismObjectIdentityBulkNewResponseStatus = "processing"
	PrismObjectIdentityBulkNewResponseStatusFailed     PrismObjectIdentityBulkNewResponseStatus = "failed"
)

func (r PrismObjectIdentityBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectIdentityBulkNewResponseStatusComplete, PrismObjectIdentityBulkNewResponseStatusProcessing, PrismObjectIdentityBulkNewResponseStatusFailed:
		return true
	}
	return false
}

// Set when status=failed; describes the job-level failure (not per-row).
type PrismObjectIdentityBulkNewResponseError struct {
	Code    string                                      `json:"code"`
	Message string                                      `json:"message"`
	JSON    prismObjectIdentityBulkNewResponseErrorJSON `json:"-"`
}

// prismObjectIdentityBulkNewResponseErrorJSON contains the JSON metadata for the
// struct [PrismObjectIdentityBulkNewResponseError]
type prismObjectIdentityBulkNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityBulkNewResponseResult struct {
	ID      string                                         `json:"id" api:"nullable" format:"uuid"`
	Created bool                                           `json:"created"`
	Error   PrismObjectIdentityBulkNewResponseResultsError `json:"error"`
	// True if the row matched an existing record via the dedupe key.
	Existing bool `json:"existing"`
	// Zero-based position of this row in the request.
	InputIndex int64 `json:"input_index"`
	// True if a matching record was updated.
	Updated bool                                         `json:"updated"`
	JSON    prismObjectIdentityBulkNewResponseResultJSON `json:"-"`
}

// prismObjectIdentityBulkNewResponseResultJSON contains the JSON metadata for the
// struct [PrismObjectIdentityBulkNewResponseResult]
type prismObjectIdentityBulkNewResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	InputIndex  apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityBulkNewResponseResultsError struct {
	Code    string                                             `json:"code"`
	Message string                                             `json:"message"`
	JSON    prismObjectIdentityBulkNewResponseResultsErrorJSON `json:"-"`
}

// prismObjectIdentityBulkNewResponseResultsErrorJSON contains the JSON metadata
// for the struct [PrismObjectIdentityBulkNewResponseResultsError]
type prismObjectIdentityBulkNewResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkNewResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkNewResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectIdentityBulkDeleteResponse struct {
	Results []PrismObjectIdentityBulkDeleteResponseResult `json:"results" api:"required"`
	Summary PrismObjectIdentityBulkDeleteResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectIdentityBulkDeleteResponseJSON     `json:"-"`
}

// prismObjectIdentityBulkDeleteResponseJSON contains the JSON metadata for the
// struct [PrismObjectIdentityBulkDeleteResponse]
type prismObjectIdentityBulkDeleteResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityBulkDeleteResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                             `json:"id" api:"required,nullable"`
	Status PrismObjectIdentityBulkDeleteResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectIdentityBulkDeleteResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectIdentityBulkDeleteResponseResultsRecord `json:"record"`
	JSON   prismObjectIdentityBulkDeleteResponseResultJSON    `json:"-"`
}

// prismObjectIdentityBulkDeleteResponseResultJSON contains the JSON metadata for
// the struct [PrismObjectIdentityBulkDeleteResponseResult]
type prismObjectIdentityBulkDeleteResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityBulkDeleteResponseResultsStatus string

const (
	PrismObjectIdentityBulkDeleteResponseResultsStatusOk    PrismObjectIdentityBulkDeleteResponseResultsStatus = "ok"
	PrismObjectIdentityBulkDeleteResponseResultsStatusError PrismObjectIdentityBulkDeleteResponseResultsStatus = "error"
)

func (r PrismObjectIdentityBulkDeleteResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectIdentityBulkDeleteResponseResultsStatusOk, PrismObjectIdentityBulkDeleteResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectIdentityBulkDeleteResponseResultsError struct {
	Code    string                                                `json:"code"`
	Message string                                                `json:"message"`
	JSON    prismObjectIdentityBulkDeleteResponseResultsErrorJSON `json:"-"`
}

// prismObjectIdentityBulkDeleteResponseResultsErrorJSON contains the JSON metadata
// for the struct [PrismObjectIdentityBulkDeleteResponseResultsError]
type prismObjectIdentityBulkDeleteResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkDeleteResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkDeleteResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityBulkDeleteResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                                 `json:"default"`
	List    interface{}                                            `json:"list"`
	JSON    prismObjectIdentityBulkDeleteResponseResultsRecordJSON `json:"-"`
}

// prismObjectIdentityBulkDeleteResponseResultsRecordJSON contains the JSON
// metadata for the struct [PrismObjectIdentityBulkDeleteResponseResultsRecord]
type prismObjectIdentityBulkDeleteResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkDeleteResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkDeleteResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityBulkDeleteResponseSummary struct {
	Failed    int64                                            `json:"failed" api:"required"`
	Succeeded int64                                            `json:"succeeded" api:"required"`
	Total     int64                                            `json:"total" api:"required"`
	JSON      prismObjectIdentityBulkDeleteResponseSummaryJSON `json:"-"`
}

// prismObjectIdentityBulkDeleteResponseSummaryJSON contains the JSON metadata for
// the struct [PrismObjectIdentityBulkDeleteResponseSummary]
type prismObjectIdentityBulkDeleteResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkDeleteResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkDeleteResponseSummaryJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectIdentityBulkUpdateResponse struct {
	Results []PrismObjectIdentityBulkUpdateResponseResult `json:"results" api:"required"`
	Summary PrismObjectIdentityBulkUpdateResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectIdentityBulkUpdateResponseJSON     `json:"-"`
}

// prismObjectIdentityBulkUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectIdentityBulkUpdateResponse]
type prismObjectIdentityBulkUpdateResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityBulkUpdateResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                             `json:"id" api:"required,nullable"`
	Status PrismObjectIdentityBulkUpdateResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectIdentityBulkUpdateResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectIdentityBulkUpdateResponseResultsRecord `json:"record"`
	JSON   prismObjectIdentityBulkUpdateResponseResultJSON    `json:"-"`
}

// prismObjectIdentityBulkUpdateResponseResultJSON contains the JSON metadata for
// the struct [PrismObjectIdentityBulkUpdateResponseResult]
type prismObjectIdentityBulkUpdateResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityBulkUpdateResponseResultsStatus string

const (
	PrismObjectIdentityBulkUpdateResponseResultsStatusOk    PrismObjectIdentityBulkUpdateResponseResultsStatus = "ok"
	PrismObjectIdentityBulkUpdateResponseResultsStatusError PrismObjectIdentityBulkUpdateResponseResultsStatus = "error"
)

func (r PrismObjectIdentityBulkUpdateResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectIdentityBulkUpdateResponseResultsStatusOk, PrismObjectIdentityBulkUpdateResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectIdentityBulkUpdateResponseResultsError struct {
	Code    string                                                `json:"code"`
	Message string                                                `json:"message"`
	JSON    prismObjectIdentityBulkUpdateResponseResultsErrorJSON `json:"-"`
}

// prismObjectIdentityBulkUpdateResponseResultsErrorJSON contains the JSON metadata
// for the struct [PrismObjectIdentityBulkUpdateResponseResultsError]
type prismObjectIdentityBulkUpdateResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkUpdateResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkUpdateResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityBulkUpdateResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                                 `json:"default"`
	List    interface{}                                            `json:"list"`
	JSON    prismObjectIdentityBulkUpdateResponseResultsRecordJSON `json:"-"`
}

// prismObjectIdentityBulkUpdateResponseResultsRecordJSON contains the JSON
// metadata for the struct [PrismObjectIdentityBulkUpdateResponseResultsRecord]
type prismObjectIdentityBulkUpdateResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkUpdateResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkUpdateResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityBulkUpdateResponseSummary struct {
	Failed    int64                                            `json:"failed" api:"required"`
	Succeeded int64                                            `json:"succeeded" api:"required"`
	Total     int64                                            `json:"total" api:"required"`
	JSON      prismObjectIdentityBulkUpdateResponseSummaryJSON `json:"-"`
}

// prismObjectIdentityBulkUpdateResponseSummaryJSON contains the JSON metadata for
// the struct [PrismObjectIdentityBulkUpdateResponseSummary]
type prismObjectIdentityBulkUpdateResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityBulkUpdateResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityBulkUpdateResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityCountResponse struct {
	// Number of records matching the access scope.
	Total int64                                `json:"total" api:"required"`
	JSON  prismObjectIdentityCountResponseJSON `json:"-"`
}

// prismObjectIdentityCountResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityCountResponse]
type prismObjectIdentityCountResponseJSON struct {
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityCountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityCountResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityDuplicateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                   `json:"default"`
	List    interface{}                              `json:"list"`
	JSON    prismObjectIdentityDuplicateResponseJSON `json:"-"`
}

// prismObjectIdentityDuplicateResponseJSON contains the JSON metadata for the
// struct [PrismObjectIdentityDuplicateResponse]
type prismObjectIdentityDuplicateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityDuplicateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityDuplicateResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityFindResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}              `json:"default"`
	List    interface{}                         `json:"list"`
	JSON    prismObjectIdentityFindResponseJSON `json:"-"`
}

// prismObjectIdentityFindResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityFindResponse]
type prismObjectIdentityFindResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityFindResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityFindResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityGetResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}             `json:"default"`
	List    interface{}                        `json:"list"`
	JSON    prismObjectIdentityGetResponseJSON `json:"-"`
}

// prismObjectIdentityGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityGetResponse]
type prismObjectIdentityGetResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityQueryResponse struct {
	Data []PrismObjectIdentityQueryResponseData `json:"data" api:"required"`
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
	Total int64                                `json:"total" api:"nullable"`
	JSON  prismObjectIdentityQueryResponseJSON `json:"-"`
}

// prismObjectIdentityQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityQueryResponse]
type prismObjectIdentityQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectIdentityQueryResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                   `json:"properties"`
	Source     []string                                 `json:"source" api:"nullable"`
	JSON       prismObjectIdentityQueryResponseDataJSON `json:"-"`
}

// prismObjectIdentityQueryResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectIdentityQueryResponseData]
type prismObjectIdentityQueryResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectIdentityQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityRestoreResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                 `json:"default"`
	List    interface{}                            `json:"list"`
	JSON    prismObjectIdentityRestoreResponseJSON `json:"-"`
}

// prismObjectIdentityRestoreResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityRestoreResponse]
type prismObjectIdentityRestoreResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityRestoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityRestoreResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectIdentityUpsertResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                `json:"default"`
	List    interface{}                           `json:"list"`
	JSON    prismObjectIdentityUpsertResponseJSON `json:"-"`
}

// prismObjectIdentityUpsertResponseJSON contains the JSON metadata for the struct
// [PrismObjectIdentityUpsertResponse]
type prismObjectIdentityUpsertResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityUpsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityUpsertResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
}

func (r PrismObjectIdentityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectIdentityUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
	IfMatch               param.Field[string]        `header:"If-Match"`
}

func (r PrismObjectIdentityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectIdentityListParams struct {
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

// URLQuery serializes [PrismObjectIdentityListParams]'s query parameters as
// `url.Values`.
func (r PrismObjectIdentityListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectIdentityDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IfMatch param.Field[string] `header:"If-Match"`
}

type PrismObjectIdentityBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects        param.Field[[]PrismObjectPropertiesParam]            `json:"objects" api:"required"`
	Options        param.Field[PrismObjectIdentityBulkNewParamsOptions] `json:"options"`
	IdempotencyKey param.Field[string]                                  `header:"Idempotency-Key"`
}

func (r PrismObjectIdentityBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectIdentityBulkNewParamsOptions struct {
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
	DedupeBy param.Field[PrismObjectIdentityBulkNewParamsOptionsDedupeByUnion] `json:"dedupe_by"`
	// App/CRM ID for context (optional)
	ListID param.Field[string] `json:"list_id" format:"uuid"`
	// Require app_stage for every row in the selected list. app_stage is a reserved
	// list-scoped alias for native status.
	RequireListStage param.Field[bool] `json:"require_list_stage"`
	// Patch a deduplicated record with the supplied properties instead of skipping it.
	UpdateExisting param.Field[bool] `json:"update_existing"`
}

func (r PrismObjectIdentityBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Property slug to deduplicate on. A single-element array is also accepted;
// compound (multi-slug) dedupe is not supported yet and is rejected with guidance.
//
// Satisfied by [shared.UnionString],
// [PrismObjectIdentityBulkNewParamsOptionsDedupeByArray].
type PrismObjectIdentityBulkNewParamsOptionsDedupeByUnion interface {
	ImplementsPrismObjectIdentityBulkNewParamsOptionsDedupeByUnion()
}

type PrismObjectIdentityBulkNewParamsOptionsDedupeByArray []string

func (r PrismObjectIdentityBulkNewParamsOptionsDedupeByArray) ImplementsPrismObjectIdentityBulkNewParamsOptionsDedupeByUnion() {
}

type PrismObjectIdentityBulkDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]   `path:"teamId" api:"required" format:"uuid"`
	IDs            param.Field[[]string] `json:"ids" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string]   `header:"Idempotency-Key"`
}

func (r PrismObjectIdentityBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectIdentityBulkUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]                                    `path:"teamId" api:"required" format:"uuid"`
	Items          param.Field[[]PrismObjectIdentityBulkUpdateParamsItem] `json:"items" api:"required"`
	IdempotencyKey param.Field[string]                                    `header:"Idempotency-Key"`
}

func (r PrismObjectIdentityBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object with `id` plus the same property body shape as PATCH
// (`default`/`list`/`extended`).
type PrismObjectIdentityBulkUpdateParamsItem struct {
	ID          param.Field[string]    `json:"id" api:"required" format:"uuid"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r PrismObjectIdentityBulkUpdateParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectIdentityCountParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the count to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectIdentityCountParams]'s query parameters as
// `url.Values`.
func (r PrismObjectIdentityCountParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectIdentityDuplicateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectIdentityFindParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the lookup to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectIdentityFindParams]'s query parameters as
// `url.Values`.
func (r PrismObjectIdentityFindParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectIdentityGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Comma-separated property slugs to return. Use dot notation for relationships.
	// `id` is always returned at the top level. Defaults to all properties.
	Select param.Field[string] `query:"select"`
}

// URLQuery serializes [PrismObjectIdentityGetParams]'s query parameters as
// `url.Values`.
func (r PrismObjectIdentityGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectIdentityQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string]                                `path:"teamId" api:"required" format:"uuid"`
	Query  param.Field[PrismObjectIdentityQueryParamsQuery]   `json:"query" api:"required"`
	ID     param.Field[PrismObjectIdentityQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes  param.Field[[]string]                              `json:"boxes"`
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

func (r PrismObjectIdentityQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectIdentityQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name). `id` is always returned at the top level of each
	// row and does not need to be selected.
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectIdentityQueryParamsQueryCombinator] `json:"combinator"`
	// Opaque cursor from a previous response's `next_cursor`. Pass it back unchanged
	// to fetch the next page. When set, `page` and `limit` are derived from the cursor
	// and any explicit values are ignored.
	Cursor param.Field[string] `json:"cursor"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectIdentityQueryParamsQueryFilterUnion] `json:"filter"`
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
	Sort param.Field[[]map[string]PrismObjectIdentityQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectIdentityQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectIdentityQueryParamsQueryCombinator string

const (
	PrismObjectIdentityQueryParamsQueryCombinatorAnd PrismObjectIdentityQueryParamsQueryCombinator = "AND"
	PrismObjectIdentityQueryParamsQueryCombinatorOr  PrismObjectIdentityQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectIdentityQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectIdentityQueryParamsQueryCombinatorAnd, PrismObjectIdentityQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

type PrismObjectIdentityQueryParamsQueryFilter struct {
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

func (r PrismObjectIdentityQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilter) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

// Satisfied by [PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEq],
// [PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNe],
// [PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLt],
// [PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGt],
// [PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLte],
// [PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGte],
// [PrismObjectIdentityQueryParamsQueryFilterContains],
// [PrismObjectIdentityQueryParamsQueryFilterBeginsWith],
// [PrismObjectIdentityQueryParamsQueryFilterEndsWith],
// [PrismObjectIdentityQueryParamsQueryFilterNotContains],
// [PrismObjectIdentityQueryParamsQueryFilterExists],
// [PrismObjectIdentityQueryParamsQueryFilterNotExists],
// [PrismObjectIdentityQueryParamsQueryFilterIsNull],
// [PrismObjectIdentityQueryParamsQueryFilterIsNotNull],
// [PrismObjectIdentityQueryParamsQueryFilterBetween],
// [PrismObjectIdentityQueryParamsQueryFilterIn],
// [PrismObjectIdentityQueryParamsQueryFilterNotIn],
// [PrismObjectIdentityQueryParamsQueryFilter].
type PrismObjectIdentityQueryParamsQueryFilterUnion interface {
	implementsPrismObjectIdentityQueryParamsQueryFilterUnion()
}

type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEq struct {
	Equals param.Field[PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion] `json:"=" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEq) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEq) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion interface {
	ImplementsPrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion()
}

type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNe struct {
	NotEquals param.Field[PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion] `json:"!=" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNe) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion interface {
	ImplementsPrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion()
}

type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLt struct {
	Less param.Field[string] `json:"<" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLt) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGt struct {
	Greater param.Field[string] `json:">" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGt) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLte struct {
	LessOrEquals param.Field[string] `json:"<=" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterLte) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGte struct {
	GreaterOrEquals param.Field[string] `json:">=" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterPrismQueryFilterGte) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterContains struct {
	Contains param.Field[PrismObjectIdentityQueryParamsQueryFilterContainsContainsUnion] `json:"contains" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterContains) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectIdentityQueryParamsQueryFilterContainsContainsArray].
type PrismObjectIdentityQueryParamsQueryFilterContainsContainsUnion interface {
	ImplementsPrismObjectIdentityQueryParamsQueryFilterContainsContainsUnion()
}

type PrismObjectIdentityQueryParamsQueryFilterContainsContainsArray []string

func (r PrismObjectIdentityQueryParamsQueryFilterContainsContainsArray) ImplementsPrismObjectIdentityQueryParamsQueryFilterContainsContainsUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterBeginsWith struct {
	BeginsWith param.Field[string] `json:"begins_with" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterBeginsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterBeginsWith) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterEndsWith struct {
	EndsWith param.Field[string] `json:"ends_with" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterEndsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterEndsWith) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterNotContains struct {
	NotContains param.Field[string] `json:"not_contains" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterNotContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterNotContains) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterExists struct {
	Exists param.Field[bool] `json:"exists" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterExists) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterNotExists struct {
	NotExists param.Field[bool] `json:"not_exists" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterNotExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterNotExists) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterIsNull struct {
	IsNull param.Field[PrismObjectIdentityQueryParamsQueryFilterIsNullIsNullUnion] `json:"is_null" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterIsNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterIsNull) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectIdentityQueryParamsQueryFilterIsNullIsNullArray].
type PrismObjectIdentityQueryParamsQueryFilterIsNullIsNullUnion interface {
	ImplementsPrismObjectIdentityQueryParamsQueryFilterIsNullIsNullUnion()
}

type PrismObjectIdentityQueryParamsQueryFilterIsNullIsNullArray []string

func (r PrismObjectIdentityQueryParamsQueryFilterIsNullIsNullArray) ImplementsPrismObjectIdentityQueryParamsQueryFilterIsNullIsNullUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterIsNotNull struct {
	IsNotNull param.Field[PrismObjectIdentityQueryParamsQueryFilterIsNotNullIsNotNullUnion] `json:"is_not_null" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterIsNotNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterIsNotNull) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectIdentityQueryParamsQueryFilterIsNotNullIsNotNullArray].
type PrismObjectIdentityQueryParamsQueryFilterIsNotNullIsNotNullUnion interface {
	ImplementsPrismObjectIdentityQueryParamsQueryFilterIsNotNullIsNotNullUnion()
}

type PrismObjectIdentityQueryParamsQueryFilterIsNotNullIsNotNullArray []string

func (r PrismObjectIdentityQueryParamsQueryFilterIsNotNullIsNotNullArray) ImplementsPrismObjectIdentityQueryParamsQueryFilterIsNotNullIsNotNullUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterBetween struct {
	Between param.Field[PrismObjectIdentityQueryParamsQueryFilterBetweenBetweenUnion] `json:"between" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterBetween) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterBetween) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectIdentityQueryParamsQueryFilterBetweenBetweenArray].
type PrismObjectIdentityQueryParamsQueryFilterBetweenBetweenUnion interface {
	ImplementsPrismObjectIdentityQueryParamsQueryFilterBetweenBetweenUnion()
}

type PrismObjectIdentityQueryParamsQueryFilterBetweenBetweenArray []string

func (r PrismObjectIdentityQueryParamsQueryFilterBetweenBetweenArray) ImplementsPrismObjectIdentityQueryParamsQueryFilterBetweenBetweenUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterIn struct {
	In param.Field[[]string] `json:"in" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterIn) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQueryFilterNotIn struct {
	NotIn param.Field[[]string] `json:"not_in" api:"required"`
}

func (r PrismObjectIdentityQueryParamsQueryFilterNotIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectIdentityQueryParamsQueryFilterNotIn) implementsPrismObjectIdentityQueryParamsQueryFilterUnion() {
}

type PrismObjectIdentityQueryParamsQuerySort string

const (
	PrismObjectIdentityQueryParamsQuerySortAsc  PrismObjectIdentityQueryParamsQuerySort = "asc"
	PrismObjectIdentityQueryParamsQuerySortDesc PrismObjectIdentityQueryParamsQuerySort = "desc"
)

func (r PrismObjectIdentityQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectIdentityQueryParamsQuerySortAsc, PrismObjectIdentityQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectIdentityQueryParamsIDArray].
type PrismObjectIdentityQueryParamsIDUnion interface {
	ImplementsPrismObjectIdentityQueryParamsIDUnion()
}

type PrismObjectIdentityQueryParamsIDArray []string

func (r PrismObjectIdentityQueryParamsIDArray) ImplementsPrismObjectIdentityQueryParamsIDUnion() {}

type PrismObjectIdentityRestoreParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectIdentityUpsertParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	// Scope the upsert to a specific list/app. Required to match or write list-scoped
	// properties, including `app_stage`.
	ListID         param.Field[string] `query:"list_id" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

func (r PrismObjectIdentityUpsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

// URLQuery serializes [PrismObjectIdentityUpsertParams]'s query parameters as
// `url.Values`.
func (r PrismObjectIdentityUpsertParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
