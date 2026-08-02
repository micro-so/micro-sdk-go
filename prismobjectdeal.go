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

// PrismObjectDealService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectDealService] method instead.
type PrismObjectDealService struct {
	Options []option.RequestOption
	Grant   *PrismObjectDealGrantService
}

// NewPrismObjectDealService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrismObjectDealService(opts ...option.RequestOption) (r *PrismObjectDealService) {
	r = &PrismObjectDealService{}
	r.Options = opts
	r.Grant = NewPrismObjectDealGrantService(opts...)
	return
}

// Create object
func (r *PrismObjectDealService) New(ctx context.Context, params PrismObjectDealNewParams, opts ...option.RequestOption) (res *PrismObjectDealNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/deal", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch object
func (r *PrismObjectDealService) Update(ctx context.Context, dealID string, params PrismObjectDealUpdateParams, opts ...option.RequestOption) (res *PrismObjectDealUpdateResponse, err error) {
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
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/deal/%s", params.TeamID, dealID)
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
func (r *PrismObjectDealService) List(ctx context.Context, params PrismObjectDealListParams, opts ...option.RequestOption) (res *PrismObjectDealListResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/deal", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete object
func (r *PrismObjectDealService) Delete(ctx context.Context, dealID string, params PrismObjectDealDeleteParams, opts ...option.RequestOption) (err error) {
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
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/deal/%s", params.TeamID, dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: small batches complete synchronously and return 200 with
// the final `ImportJob`; large batches start an async job, return 202 with
// `status: processing` and a `Location` header, and can be polled via
// `GET /v2/prism/{teamId}/imports/{jobId}`.
func (r *PrismObjectDealService) BulkNew(ctx context.Context, params PrismObjectDealBulkNewParams, opts ...option.RequestOption) (res *PrismObjectDealBulkNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/deal/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Soft-delete up to 100 records in a single call. Same partial-success contract as
// batch/update.
func (r *PrismObjectDealService) BulkDelete(ctx context.Context, params PrismObjectDealBulkDeleteParams, opts ...option.RequestOption) (res *PrismObjectDealBulkDeleteResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/deal/batch/delete", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch up to 100 records in a single call. Each item is attempted independently —
// failures don't abort the batch. Inspect `results[].status` per item.
func (r *PrismObjectDealService) BulkUpdate(ctx context.Context, params PrismObjectDealBulkUpdateParams, opts ...option.RequestOption) (res *PrismObjectDealBulkUpdateResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/deal/batch/update", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the total number of records of this object type that the caller can see.
// Avoids the page-overshoot anti-pattern — clients no longer need to keep paging
// until `has_more` flips false to discover the total. Currently does not apply
// query filters; for a filtered total, pass `include_total: true` in a POST
// `/query` body.
func (r *PrismObjectDealService) Count(ctx context.Context, params PrismObjectDealCountParams, opts ...option.RequestOption) (res *PrismObjectDealCountResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/deal/count", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Duplicate object
func (r *PrismObjectDealService) Duplicate(ctx context.Context, dealID string, params PrismObjectDealDuplicateParams, opts ...option.RequestOption) (res *PrismObjectDealDuplicateResponse, err error) {
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
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/deal/%s/duplicate", params.TeamID, dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns the single record whose property `{slug}` equals `{value}`. 404 if
// nothing matches; 409 if more than one record matches.
func (r *PrismObjectDealService) Find(ctx context.Context, slug string, value string, params PrismObjectDealFindParams, opts ...option.RequestOption) (res *PrismObjectDealFindResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/deal/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectDealService) Get(ctx context.Context, dealID string, params PrismObjectDealGetParams, opts ...option.RequestOption) (res *PrismObjectDealGetResponse, err error) {
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
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/deal/%s", params.TeamID, dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Query
func (r *PrismObjectDealService) Query(ctx context.Context, params PrismObjectDealQueryParams, opts ...option.RequestOption) (res *PrismObjectDealQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/deal/query", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Restore object
func (r *PrismObjectDealService) Restore(ctx context.Context, dealID string, params PrismObjectDealRestoreParams, opts ...option.RequestOption) (res *PrismObjectDealRestoreResponse, err error) {
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
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/deal/%s/restore", params.TeamID, dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Idempotent create-or-update keyed on `{slug}={value}`. If exactly one record
// matches, it is patched and 200 is returned. If none match, a new record is
// created (with the lookup property set if absent) and 201 is returned. If
// multiple records match, 409 is returned and you should patch by id instead.
func (r *PrismObjectDealService) Upsert(ctx context.Context, slug string, value string, params PrismObjectDealUpsertParams, opts ...option.RequestOption) (res *PrismObjectDealUpsertResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/deal/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealNewResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}         `json:"default"`
	List    interface{}                    `json:"list"`
	JSON    prismObjectDealNewResponseJSON `json:"-"`
}

// prismObjectDealNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealNewResponse]
type prismObjectDealNewResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealNewResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealUpdateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}            `json:"default"`
	List    interface{}                       `json:"list"`
	JSON    prismObjectDealUpdateResponseJSON `json:"-"`
}

// prismObjectDealUpdateResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealUpdateResponse]
type prismObjectDealUpdateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealListResponse struct {
	Data []PrismObjectDealListResponseData `json:"data" api:"required"`
	// Accurate end-of-data signal — false on the last page, never forces clients to
	// overshoot.
	HasMore    bool   `json:"has_more" api:"required"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// Populated only when `?include_total=true` was passed.
	Total int64                           `json:"total" api:"nullable"`
	JSON  prismObjectDealListResponseJSON `json:"-"`
}

// prismObjectDealListResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealListResponse]
type prismObjectDealListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealListResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectDealListResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}              `json:"properties"`
	Source     []string                            `json:"source" api:"nullable"`
	JSON       prismObjectDealListResponseDataJSON `json:"-"`
}

// prismObjectDealListResponseDataJSON contains the JSON metadata for the struct
// [PrismObjectDealListResponseData]
type prismObjectDealListResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectDealListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Status snapshot of an import job. Same shape used by the POST /import response
// and by GET /imports/{jobId}.
type PrismObjectDealBulkNewResponse struct {
	// Null for sync imports (results inlined). Set for async imports.
	JobID  string                               `json:"job_id" api:"required,nullable"`
	Status PrismObjectDealBulkNewResponseStatus `json:"status" api:"required"`
	// Total number of rows in the import.
	Total     int64     `json:"total" api:"required"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Set when status=failed; describes the job-level failure (not per-row).
	Error     PrismObjectDealBulkNewResponseError `json:"error"`
	ExpiresAt time.Time                           `json:"expires_at" format:"date-time"`
	Failed    int64                               `json:"failed"`
	// Rows that have been attempted (succeeded + failed).
	Processed int64 `json:"processed"`
	// Per-row outcomes. Always present for sync imports; populated for async imports
	// once the job reaches `complete`.
	Results   []PrismObjectDealBulkNewResponseResult `json:"results"`
	Succeeded int64                                  `json:"succeeded"`
	UpdatedAt time.Time                              `json:"updated_at" format:"date-time"`
	JSON      prismObjectDealBulkNewResponseJSON     `json:"-"`
}

// prismObjectDealBulkNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealBulkNewResponse]
type prismObjectDealBulkNewResponseJSON struct {
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

func (r *PrismObjectDealBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealBulkNewResponseStatus string

const (
	PrismObjectDealBulkNewResponseStatusComplete   PrismObjectDealBulkNewResponseStatus = "complete"
	PrismObjectDealBulkNewResponseStatusProcessing PrismObjectDealBulkNewResponseStatus = "processing"
	PrismObjectDealBulkNewResponseStatusFailed     PrismObjectDealBulkNewResponseStatus = "failed"
)

func (r PrismObjectDealBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectDealBulkNewResponseStatusComplete, PrismObjectDealBulkNewResponseStatusProcessing, PrismObjectDealBulkNewResponseStatusFailed:
		return true
	}
	return false
}

// Set when status=failed; describes the job-level failure (not per-row).
type PrismObjectDealBulkNewResponseError struct {
	Code    string                                  `json:"code"`
	Message string                                  `json:"message"`
	JSON    prismObjectDealBulkNewResponseErrorJSON `json:"-"`
}

// prismObjectDealBulkNewResponseErrorJSON contains the JSON metadata for the
// struct [PrismObjectDealBulkNewResponseError]
type prismObjectDealBulkNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealBulkNewResponseResult struct {
	ID      string                                     `json:"id" api:"nullable" format:"uuid"`
	Created bool                                       `json:"created"`
	Error   PrismObjectDealBulkNewResponseResultsError `json:"error"`
	// True if the row matched an existing record via the dedupe key.
	Existing bool `json:"existing"`
	// Zero-based position of this row in the request.
	InputIndex int64 `json:"input_index"`
	// True if a matching record was updated.
	Updated bool                                     `json:"updated"`
	JSON    prismObjectDealBulkNewResponseResultJSON `json:"-"`
}

// prismObjectDealBulkNewResponseResultJSON contains the JSON metadata for the
// struct [PrismObjectDealBulkNewResponseResult]
type prismObjectDealBulkNewResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	InputIndex  apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealBulkNewResponseResultsError struct {
	Code    string                                         `json:"code"`
	Message string                                         `json:"message"`
	JSON    prismObjectDealBulkNewResponseResultsErrorJSON `json:"-"`
}

// prismObjectDealBulkNewResponseResultsErrorJSON contains the JSON metadata for
// the struct [PrismObjectDealBulkNewResponseResultsError]
type prismObjectDealBulkNewResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkNewResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkNewResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectDealBulkDeleteResponse struct {
	Results []PrismObjectDealBulkDeleteResponseResult `json:"results" api:"required"`
	Summary PrismObjectDealBulkDeleteResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectDealBulkDeleteResponseJSON     `json:"-"`
}

// prismObjectDealBulkDeleteResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealBulkDeleteResponse]
type prismObjectDealBulkDeleteResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealBulkDeleteResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                         `json:"id" api:"required,nullable"`
	Status PrismObjectDealBulkDeleteResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectDealBulkDeleteResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectDealBulkDeleteResponseResultsRecord `json:"record"`
	JSON   prismObjectDealBulkDeleteResponseResultJSON    `json:"-"`
}

// prismObjectDealBulkDeleteResponseResultJSON contains the JSON metadata for the
// struct [PrismObjectDealBulkDeleteResponseResult]
type prismObjectDealBulkDeleteResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealBulkDeleteResponseResultsStatus string

const (
	PrismObjectDealBulkDeleteResponseResultsStatusOk    PrismObjectDealBulkDeleteResponseResultsStatus = "ok"
	PrismObjectDealBulkDeleteResponseResultsStatusError PrismObjectDealBulkDeleteResponseResultsStatus = "error"
)

func (r PrismObjectDealBulkDeleteResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectDealBulkDeleteResponseResultsStatusOk, PrismObjectDealBulkDeleteResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectDealBulkDeleteResponseResultsError struct {
	Code    string                                            `json:"code"`
	Message string                                            `json:"message"`
	JSON    prismObjectDealBulkDeleteResponseResultsErrorJSON `json:"-"`
}

// prismObjectDealBulkDeleteResponseResultsErrorJSON contains the JSON metadata for
// the struct [PrismObjectDealBulkDeleteResponseResultsError]
type prismObjectDealBulkDeleteResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkDeleteResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkDeleteResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealBulkDeleteResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                             `json:"default"`
	List    interface{}                                        `json:"list"`
	JSON    prismObjectDealBulkDeleteResponseResultsRecordJSON `json:"-"`
}

// prismObjectDealBulkDeleteResponseResultsRecordJSON contains the JSON metadata
// for the struct [PrismObjectDealBulkDeleteResponseResultsRecord]
type prismObjectDealBulkDeleteResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkDeleteResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkDeleteResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealBulkDeleteResponseSummary struct {
	Failed    int64                                        `json:"failed" api:"required"`
	Succeeded int64                                        `json:"succeeded" api:"required"`
	Total     int64                                        `json:"total" api:"required"`
	JSON      prismObjectDealBulkDeleteResponseSummaryJSON `json:"-"`
}

// prismObjectDealBulkDeleteResponseSummaryJSON contains the JSON metadata for the
// struct [PrismObjectDealBulkDeleteResponseSummary]
type prismObjectDealBulkDeleteResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkDeleteResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkDeleteResponseSummaryJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectDealBulkUpdateResponse struct {
	Results []PrismObjectDealBulkUpdateResponseResult `json:"results" api:"required"`
	Summary PrismObjectDealBulkUpdateResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectDealBulkUpdateResponseJSON     `json:"-"`
}

// prismObjectDealBulkUpdateResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealBulkUpdateResponse]
type prismObjectDealBulkUpdateResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealBulkUpdateResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                         `json:"id" api:"required,nullable"`
	Status PrismObjectDealBulkUpdateResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectDealBulkUpdateResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectDealBulkUpdateResponseResultsRecord `json:"record"`
	JSON   prismObjectDealBulkUpdateResponseResultJSON    `json:"-"`
}

// prismObjectDealBulkUpdateResponseResultJSON contains the JSON metadata for the
// struct [PrismObjectDealBulkUpdateResponseResult]
type prismObjectDealBulkUpdateResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealBulkUpdateResponseResultsStatus string

const (
	PrismObjectDealBulkUpdateResponseResultsStatusOk    PrismObjectDealBulkUpdateResponseResultsStatus = "ok"
	PrismObjectDealBulkUpdateResponseResultsStatusError PrismObjectDealBulkUpdateResponseResultsStatus = "error"
)

func (r PrismObjectDealBulkUpdateResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectDealBulkUpdateResponseResultsStatusOk, PrismObjectDealBulkUpdateResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectDealBulkUpdateResponseResultsError struct {
	Code    string                                            `json:"code"`
	Message string                                            `json:"message"`
	JSON    prismObjectDealBulkUpdateResponseResultsErrorJSON `json:"-"`
}

// prismObjectDealBulkUpdateResponseResultsErrorJSON contains the JSON metadata for
// the struct [PrismObjectDealBulkUpdateResponseResultsError]
type prismObjectDealBulkUpdateResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkUpdateResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkUpdateResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealBulkUpdateResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                             `json:"default"`
	List    interface{}                                        `json:"list"`
	JSON    prismObjectDealBulkUpdateResponseResultsRecordJSON `json:"-"`
}

// prismObjectDealBulkUpdateResponseResultsRecordJSON contains the JSON metadata
// for the struct [PrismObjectDealBulkUpdateResponseResultsRecord]
type prismObjectDealBulkUpdateResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkUpdateResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkUpdateResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealBulkUpdateResponseSummary struct {
	Failed    int64                                        `json:"failed" api:"required"`
	Succeeded int64                                        `json:"succeeded" api:"required"`
	Total     int64                                        `json:"total" api:"required"`
	JSON      prismObjectDealBulkUpdateResponseSummaryJSON `json:"-"`
}

// prismObjectDealBulkUpdateResponseSummaryJSON contains the JSON metadata for the
// struct [PrismObjectDealBulkUpdateResponseSummary]
type prismObjectDealBulkUpdateResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealBulkUpdateResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealBulkUpdateResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealCountResponse struct {
	// Number of records matching the access scope.
	Total int64                            `json:"total" api:"required"`
	JSON  prismObjectDealCountResponseJSON `json:"-"`
}

// prismObjectDealCountResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealCountResponse]
type prismObjectDealCountResponseJSON struct {
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealCountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealCountResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealDuplicateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}               `json:"default"`
	List    interface{}                          `json:"list"`
	JSON    prismObjectDealDuplicateResponseJSON `json:"-"`
}

// prismObjectDealDuplicateResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealDuplicateResponse]
type prismObjectDealDuplicateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealDuplicateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealDuplicateResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealFindResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}          `json:"default"`
	List    interface{}                     `json:"list"`
	JSON    prismObjectDealFindResponseJSON `json:"-"`
}

// prismObjectDealFindResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealFindResponse]
type prismObjectDealFindResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealFindResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealFindResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealGetResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}         `json:"default"`
	List    interface{}                    `json:"list"`
	JSON    prismObjectDealGetResponseJSON `json:"-"`
}

// prismObjectDealGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealGetResponse]
type prismObjectDealGetResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealQueryResponse struct {
	Data []PrismObjectDealQueryResponseData `json:"data" api:"required"`
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
	Total int64                            `json:"total" api:"nullable"`
	JSON  prismObjectDealQueryResponseJSON `json:"-"`
}

// prismObjectDealQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealQueryResponse]
type prismObjectDealQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectDealQueryResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}               `json:"properties"`
	Source     []string                             `json:"source" api:"nullable"`
	JSON       prismObjectDealQueryResponseDataJSON `json:"-"`
}

// prismObjectDealQueryResponseDataJSON contains the JSON metadata for the struct
// [PrismObjectDealQueryResponseData]
type prismObjectDealQueryResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectDealQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealRestoreResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}             `json:"default"`
	List    interface{}                        `json:"list"`
	JSON    prismObjectDealRestoreResponseJSON `json:"-"`
}

// prismObjectDealRestoreResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealRestoreResponse]
type prismObjectDealRestoreResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealRestoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealRestoreResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDealUpsertResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}            `json:"default"`
	List    interface{}                       `json:"list"`
	JSON    prismObjectDealUpsertResponseJSON `json:"-"`
}

// prismObjectDealUpsertResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealUpsertResponse]
type prismObjectDealUpsertResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealUpsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealUpsertResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
}

func (r PrismObjectDealNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectDealUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
	IfMatch               param.Field[string]        `header:"If-Match"`
}

func (r PrismObjectDealUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectDealListParams struct {
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

// URLQuery serializes [PrismObjectDealListParams]'s query parameters as
// `url.Values`.
func (r PrismObjectDealListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectDealDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IfMatch param.Field[string] `header:"If-Match"`
}

type PrismObjectDealBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects        param.Field[[]PrismObjectPropertiesParam]        `json:"objects" api:"required"`
	Options        param.Field[PrismObjectDealBulkNewParamsOptions] `json:"options"`
	IdempotencyKey param.Field[string]                              `header:"Idempotency-Key"`
}

func (r PrismObjectDealBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDealBulkNewParamsOptions struct {
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
	DedupeBy param.Field[PrismObjectDealBulkNewParamsOptionsDedupeByUnion] `json:"dedupe_by"`
	// App/CRM ID for context (optional)
	ListID param.Field[string] `json:"list_id" format:"uuid"`
	// Require app_stage for every row in the selected list. app_stage is a reserved
	// list-scoped alias for native status.
	RequireListStage param.Field[bool] `json:"require_list_stage"`
	// Patch a deduplicated record with the supplied properties instead of skipping it.
	UpdateExisting param.Field[bool] `json:"update_existing"`
}

func (r PrismObjectDealBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Property slug to deduplicate on. A single-element array is also accepted;
// compound (multi-slug) dedupe is not supported yet and is rejected with guidance.
//
// Satisfied by [shared.UnionString],
// [PrismObjectDealBulkNewParamsOptionsDedupeByArray].
type PrismObjectDealBulkNewParamsOptionsDedupeByUnion interface {
	ImplementsPrismObjectDealBulkNewParamsOptionsDedupeByUnion()
}

type PrismObjectDealBulkNewParamsOptionsDedupeByArray []string

func (r PrismObjectDealBulkNewParamsOptionsDedupeByArray) ImplementsPrismObjectDealBulkNewParamsOptionsDedupeByUnion() {
}

type PrismObjectDealBulkDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]   `path:"teamId" api:"required" format:"uuid"`
	IDs            param.Field[[]string] `json:"ids" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string]   `header:"Idempotency-Key"`
}

func (r PrismObjectDealBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDealBulkUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]                                `path:"teamId" api:"required" format:"uuid"`
	Items          param.Field[[]PrismObjectDealBulkUpdateParamsItem] `json:"items" api:"required"`
	IdempotencyKey param.Field[string]                                `header:"Idempotency-Key"`
}

func (r PrismObjectDealBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object with `id` plus the same property body shape as PATCH
// (`default`/`list`/`extended`).
type PrismObjectDealBulkUpdateParamsItem struct {
	ID          param.Field[string]    `json:"id" api:"required" format:"uuid"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r PrismObjectDealBulkUpdateParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDealCountParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the count to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectDealCountParams]'s query parameters as
// `url.Values`.
func (r PrismObjectDealCountParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectDealDuplicateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectDealFindParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the lookup to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectDealFindParams]'s query parameters as
// `url.Values`.
func (r PrismObjectDealFindParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectDealGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Comma-separated property slugs to return. Use dot notation for relationships.
	// `id` is always returned at the top level. Defaults to all properties.
	Select param.Field[string] `query:"select"`
}

// URLQuery serializes [PrismObjectDealGetParams]'s query parameters as
// `url.Values`.
func (r PrismObjectDealGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectDealQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string]                            `path:"teamId" api:"required" format:"uuid"`
	Query  param.Field[PrismObjectDealQueryParamsQuery]   `json:"query" api:"required"`
	ID     param.Field[PrismObjectDealQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes  param.Field[[]string]                          `json:"boxes"`
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

func (r PrismObjectDealQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDealQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name). `id` is always returned at the top level of each
	// row and does not need to be selected.
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectDealQueryParamsQueryCombinator] `json:"combinator"`
	// Opaque cursor from a previous response's `next_cursor`. Pass it back unchanged
	// to fetch the next page. When set, `page` and `limit` are derived from the cursor
	// and any explicit values are ignored.
	Cursor param.Field[string] `json:"cursor"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectDealQueryParamsQueryFilterUnion] `json:"filter"`
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
	Sort param.Field[[]map[string]PrismObjectDealQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectDealQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectDealQueryParamsQueryCombinator string

const (
	PrismObjectDealQueryParamsQueryCombinatorAnd PrismObjectDealQueryParamsQueryCombinator = "AND"
	PrismObjectDealQueryParamsQueryCombinatorOr  PrismObjectDealQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectDealQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectDealQueryParamsQueryCombinatorAnd, PrismObjectDealQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

type PrismObjectDealQueryParamsQueryFilter struct {
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

func (r PrismObjectDealQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilter) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

// Satisfied by [PrismObjectDealQueryParamsQueryFilterPrismQueryFilterEq],
// [PrismObjectDealQueryParamsQueryFilterPrismQueryFilterNe],
// [PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLt],
// [PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGt],
// [PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLte],
// [PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGte],
// [PrismObjectDealQueryParamsQueryFilterContains],
// [PrismObjectDealQueryParamsQueryFilterBeginsWith],
// [PrismObjectDealQueryParamsQueryFilterEndsWith],
// [PrismObjectDealQueryParamsQueryFilterNotContains],
// [PrismObjectDealQueryParamsQueryFilterExists],
// [PrismObjectDealQueryParamsQueryFilterNotExists],
// [PrismObjectDealQueryParamsQueryFilterIsNull],
// [PrismObjectDealQueryParamsQueryFilterIsNotNull],
// [PrismObjectDealQueryParamsQueryFilterBetween],
// [PrismObjectDealQueryParamsQueryFilterIn],
// [PrismObjectDealQueryParamsQueryFilterNotIn],
// [PrismObjectDealQueryParamsQueryFilter].
type PrismObjectDealQueryParamsQueryFilterUnion interface {
	implementsPrismObjectDealQueryParamsQueryFilterUnion()
}

type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterEq struct {
	Equals param.Field[PrismObjectDealQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion] `json:"=" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterEq) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterEq) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion interface {
	ImplementsPrismObjectDealQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion()
}

type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterNe struct {
	NotEquals param.Field[PrismObjectDealQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion] `json:"!=" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterNe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterNe) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion interface {
	ImplementsPrismObjectDealQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion()
}

type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLt struct {
	Less param.Field[string] `json:"<" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLt) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGt struct {
	Greater param.Field[string] `json:">" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGt) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLte struct {
	LessOrEquals param.Field[string] `json:"<=" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterLte) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGte struct {
	GreaterOrEquals param.Field[string] `json:">=" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterPrismQueryFilterGte) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterContains struct {
	Contains param.Field[PrismObjectDealQueryParamsQueryFilterContainsContainsUnion] `json:"contains" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterContains) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectDealQueryParamsQueryFilterContainsContainsArray].
type PrismObjectDealQueryParamsQueryFilterContainsContainsUnion interface {
	ImplementsPrismObjectDealQueryParamsQueryFilterContainsContainsUnion()
}

type PrismObjectDealQueryParamsQueryFilterContainsContainsArray []string

func (r PrismObjectDealQueryParamsQueryFilterContainsContainsArray) ImplementsPrismObjectDealQueryParamsQueryFilterContainsContainsUnion() {
}

type PrismObjectDealQueryParamsQueryFilterBeginsWith struct {
	BeginsWith param.Field[string] `json:"begins_with" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterBeginsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterBeginsWith) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterEndsWith struct {
	EndsWith param.Field[string] `json:"ends_with" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterEndsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterEndsWith) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterNotContains struct {
	NotContains param.Field[string] `json:"not_contains" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterNotContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterNotContains) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterExists struct {
	Exists param.Field[bool] `json:"exists" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterExists) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterNotExists struct {
	NotExists param.Field[bool] `json:"not_exists" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterNotExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterNotExists) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterIsNull struct {
	IsNull param.Field[PrismObjectDealQueryParamsQueryFilterIsNullIsNullUnion] `json:"is_null" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterIsNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterIsNull) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectDealQueryParamsQueryFilterIsNullIsNullArray].
type PrismObjectDealQueryParamsQueryFilterIsNullIsNullUnion interface {
	ImplementsPrismObjectDealQueryParamsQueryFilterIsNullIsNullUnion()
}

type PrismObjectDealQueryParamsQueryFilterIsNullIsNullArray []string

func (r PrismObjectDealQueryParamsQueryFilterIsNullIsNullArray) ImplementsPrismObjectDealQueryParamsQueryFilterIsNullIsNullUnion() {
}

type PrismObjectDealQueryParamsQueryFilterIsNotNull struct {
	IsNotNull param.Field[PrismObjectDealQueryParamsQueryFilterIsNotNullIsNotNullUnion] `json:"is_not_null" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterIsNotNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterIsNotNull) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectDealQueryParamsQueryFilterIsNotNullIsNotNullArray].
type PrismObjectDealQueryParamsQueryFilterIsNotNullIsNotNullUnion interface {
	ImplementsPrismObjectDealQueryParamsQueryFilterIsNotNullIsNotNullUnion()
}

type PrismObjectDealQueryParamsQueryFilterIsNotNullIsNotNullArray []string

func (r PrismObjectDealQueryParamsQueryFilterIsNotNullIsNotNullArray) ImplementsPrismObjectDealQueryParamsQueryFilterIsNotNullIsNotNullUnion() {
}

type PrismObjectDealQueryParamsQueryFilterBetween struct {
	Between param.Field[PrismObjectDealQueryParamsQueryFilterBetweenBetweenUnion] `json:"between" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterBetween) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterBetween) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectDealQueryParamsQueryFilterBetweenBetweenArray].
type PrismObjectDealQueryParamsQueryFilterBetweenBetweenUnion interface {
	ImplementsPrismObjectDealQueryParamsQueryFilterBetweenBetweenUnion()
}

type PrismObjectDealQueryParamsQueryFilterBetweenBetweenArray []string

func (r PrismObjectDealQueryParamsQueryFilterBetweenBetweenArray) ImplementsPrismObjectDealQueryParamsQueryFilterBetweenBetweenUnion() {
}

type PrismObjectDealQueryParamsQueryFilterIn struct {
	In param.Field[[]string] `json:"in" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterIn) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQueryFilterNotIn struct {
	NotIn param.Field[[]string] `json:"not_in" api:"required"`
}

func (r PrismObjectDealQueryParamsQueryFilterNotIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDealQueryParamsQueryFilterNotIn) implementsPrismObjectDealQueryParamsQueryFilterUnion() {
}

type PrismObjectDealQueryParamsQuerySort string

const (
	PrismObjectDealQueryParamsQuerySortAsc  PrismObjectDealQueryParamsQuerySort = "asc"
	PrismObjectDealQueryParamsQuerySortDesc PrismObjectDealQueryParamsQuerySort = "desc"
)

func (r PrismObjectDealQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectDealQueryParamsQuerySortAsc, PrismObjectDealQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectDealQueryParamsIDArray].
type PrismObjectDealQueryParamsIDUnion interface {
	ImplementsPrismObjectDealQueryParamsIDUnion()
}

type PrismObjectDealQueryParamsIDArray []string

func (r PrismObjectDealQueryParamsIDArray) ImplementsPrismObjectDealQueryParamsIDUnion() {}

type PrismObjectDealRestoreParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectDealUpsertParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	// Scope the upsert to a specific list/app. Required to match or write list-scoped
	// properties, including `app_stage`.
	ListID         param.Field[string] `query:"list_id" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

func (r PrismObjectDealUpsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

// URLQuery serializes [PrismObjectDealUpsertParams]'s query parameters as
// `url.Values`.
func (r PrismObjectDealUpsertParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
