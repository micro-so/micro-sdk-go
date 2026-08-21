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

// PrismObjectEngagementService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectEngagementService] method instead.
type PrismObjectEngagementService struct {
	Options []option.RequestOption
	Grant   *PrismObjectEngagementGrantService
}

// NewPrismObjectEngagementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectEngagementService(opts ...option.RequestOption) (r *PrismObjectEngagementService) {
	r = &PrismObjectEngagementService{}
	r.Options = opts
	r.Grant = NewPrismObjectEngagementGrantService(opts...)
	return
}

// Creates a record. For `document`, writing `content` (or HTML) stores the
// property and reads back, but the in-app editor is CRDT-backed and will render a
// blank page until that document has been opened and saved in the app. Treat
// API-created docs as data records, not as collaboratively edited pages, unless
// you only need the stored property values.
func (r *PrismObjectEngagementService) New(ctx context.Context, params PrismObjectEngagementNewParams, opts ...option.RequestOption) (res *PrismObjectEngagementNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/engagement", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch object
func (r *PrismObjectEngagementService) Update(ctx context.Context, engagementID string, params PrismObjectEngagementUpdateParams, opts ...option.RequestOption) (res *PrismObjectEngagementUpdateResponse, err error) {
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
	if engagementID == "" {
		err = errors.New("missing required engagementId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/engagement/%s", params.TeamID, engagementID)
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
func (r *PrismObjectEngagementService) List(ctx context.Context, params PrismObjectEngagementListParams, opts ...option.RequestOption) (res *PrismObjectEngagementListResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/engagement", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete object
func (r *PrismObjectEngagementService) Delete(ctx context.Context, engagementID string, params PrismObjectEngagementDeleteParams, opts ...option.RequestOption) (err error) {
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
	if engagementID == "" {
		err = errors.New("missing required engagementId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/engagement/%s", params.TeamID, engagementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: small batches complete synchronously and return 200 with
// the final `ImportJob`; large batches start an async job, return 202 with
// `status: processing` and a `Location` header, and can be polled via
// `GET /v2/prism/{teamId}/imports/{jobId}`.
func (r *PrismObjectEngagementService) BulkNew(ctx context.Context, params PrismObjectEngagementBulkNewParams, opts ...option.RequestOption) (res *PrismObjectEngagementBulkNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/engagement/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Soft-delete up to 100 records in a single call. Same partial-success contract as
// batch/update.
func (r *PrismObjectEngagementService) BulkDelete(ctx context.Context, params PrismObjectEngagementBulkDeleteParams, opts ...option.RequestOption) (res *PrismObjectEngagementBulkDeleteResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/engagement/batch/delete", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch up to 100 records in a single call. Each item is attempted independently —
// failures don't abort the batch. Inspect `results[].status` per item.
func (r *PrismObjectEngagementService) BulkUpdate(ctx context.Context, params PrismObjectEngagementBulkUpdateParams, opts ...option.RequestOption) (res *PrismObjectEngagementBulkUpdateResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/engagement/batch/update", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the total number of records of this object type that the caller can see.
// Avoids the page-overshoot anti-pattern — clients no longer need to keep paging
// until `has_more` flips false to discover the total. Currently does not apply
// query filters; for a filtered total, pass `include_total: true` in a POST
// `/query` body. Unfiltered counts on high-cardinality types (especially
// `engagement`) scan the full access-scoped set and can take tens of seconds or
// time out; prefer a filtered `include_total` query or accept that this endpoint
// is expensive there.
func (r *PrismObjectEngagementService) Count(ctx context.Context, params PrismObjectEngagementCountParams, opts ...option.RequestOption) (res *PrismObjectEngagementCountResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/engagement/count", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Duplicate object
func (r *PrismObjectEngagementService) Duplicate(ctx context.Context, engagementID string, params PrismObjectEngagementDuplicateParams, opts ...option.RequestOption) (res *PrismObjectEngagementDuplicateResponse, err error) {
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
	if engagementID == "" {
		err = errors.New("missing required engagementId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/engagement/%s/duplicate", params.TeamID, engagementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns the single record whose property `{slug}` equals `{value}`. 404 if
// nothing matches; 409 if more than one record matches.
func (r *PrismObjectEngagementService) Find(ctx context.Context, slug string, value string, params PrismObjectEngagementFindParams, opts ...option.RequestOption) (res *PrismObjectEngagementFindResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/engagement/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectEngagementService) Get(ctx context.Context, engagementID string, params PrismObjectEngagementGetParams, opts ...option.RequestOption) (res *PrismObjectEngagementGetResponse, err error) {
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
	if engagementID == "" {
		err = errors.New("missing required engagementId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/engagement/%s", params.TeamID, engagementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Query
func (r *PrismObjectEngagementService) Query(ctx context.Context, params PrismObjectEngagementQueryParams, opts ...option.RequestOption) (res *PrismObjectEngagementQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/engagement/query", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Restore object
func (r *PrismObjectEngagementService) Restore(ctx context.Context, engagementID string, params PrismObjectEngagementRestoreParams, opts ...option.RequestOption) (res *PrismObjectEngagementRestoreResponse, err error) {
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
	if engagementID == "" {
		err = errors.New("missing required engagementId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/engagement/%s/restore", params.TeamID, engagementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Idempotent create-or-update keyed on `{slug}={value}`. If exactly one record
// matches, it is patched and 200 is returned. If none match, a new record is
// created (with the lookup property set if absent) and 201 is returned. If
// multiple records match, 409 is returned and you should patch by id instead.
func (r *PrismObjectEngagementService) Upsert(ctx context.Context, slug string, value string, params PrismObjectEngagementUpsertParams, opts ...option.RequestOption) (res *PrismObjectEngagementUpsertResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/engagement/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectEngagementNewResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}               `json:"default"`
	List    interface{}                          `json:"list"`
	JSON    prismObjectEngagementNewResponseJSON `json:"-"`
}

// prismObjectEngagementNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectEngagementNewResponse]
type prismObjectEngagementNewResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementNewResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectEngagementUpdateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                  `json:"default"`
	List    interface{}                             `json:"list"`
	JSON    prismObjectEngagementUpdateResponseJSON `json:"-"`
}

// prismObjectEngagementUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectEngagementUpdateResponse]
type prismObjectEngagementUpdateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementListResponse struct {
	Data []PrismObjectEngagementListResponseData `json:"data" api:"required"`
	// Accurate end-of-data signal — false on the last page, never forces clients to
	// overshoot.
	HasMore    bool   `json:"has_more" api:"required"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// Populated only when `?include_total=true` was passed.
	Total int64                                 `json:"total" api:"nullable"`
	JSON  prismObjectEngagementListResponseJSON `json:"-"`
}

// prismObjectEngagementListResponseJSON contains the JSON metadata for the struct
// [PrismObjectEngagementListResponse]
type prismObjectEngagementListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementListResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectEngagementListResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                    `json:"properties"`
	Source     []string                                  `json:"source" api:"nullable"`
	JSON       prismObjectEngagementListResponseDataJSON `json:"-"`
}

// prismObjectEngagementListResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectEngagementListResponseData]
type prismObjectEngagementListResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectEngagementListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Status snapshot of an import job. Same shape used by the POST /import response
// and by GET /imports/{jobId}.
type PrismObjectEngagementBulkNewResponse struct {
	// Null for sync imports (results inlined). Set for async imports.
	JobID  string                                     `json:"job_id" api:"required,nullable"`
	Status PrismObjectEngagementBulkNewResponseStatus `json:"status" api:"required"`
	// Total number of rows in the import.
	Total     int64     `json:"total" api:"required"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Set when status=failed; describes the job-level failure (not per-row).
	Error     PrismObjectEngagementBulkNewResponseError `json:"error"`
	ExpiresAt time.Time                                 `json:"expires_at" format:"date-time"`
	Failed    int64                                     `json:"failed"`
	// Rows that have been attempted (succeeded + failed).
	Processed int64 `json:"processed"`
	// Per-row outcomes. Always present for sync imports; populated for async imports
	// once the job reaches `complete`.
	Results   []PrismObjectEngagementBulkNewResponseResult `json:"results"`
	Succeeded int64                                        `json:"succeeded"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      prismObjectEngagementBulkNewResponseJSON     `json:"-"`
}

// prismObjectEngagementBulkNewResponseJSON contains the JSON metadata for the
// struct [PrismObjectEngagementBulkNewResponse]
type prismObjectEngagementBulkNewResponseJSON struct {
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

func (r *PrismObjectEngagementBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementBulkNewResponseStatus string

const (
	PrismObjectEngagementBulkNewResponseStatusComplete   PrismObjectEngagementBulkNewResponseStatus = "complete"
	PrismObjectEngagementBulkNewResponseStatusProcessing PrismObjectEngagementBulkNewResponseStatus = "processing"
	PrismObjectEngagementBulkNewResponseStatusFailed     PrismObjectEngagementBulkNewResponseStatus = "failed"
)

func (r PrismObjectEngagementBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectEngagementBulkNewResponseStatusComplete, PrismObjectEngagementBulkNewResponseStatusProcessing, PrismObjectEngagementBulkNewResponseStatusFailed:
		return true
	}
	return false
}

// Set when status=failed; describes the job-level failure (not per-row).
type PrismObjectEngagementBulkNewResponseError struct {
	Code    string                                        `json:"code"`
	Message string                                        `json:"message"`
	JSON    prismObjectEngagementBulkNewResponseErrorJSON `json:"-"`
}

// prismObjectEngagementBulkNewResponseErrorJSON contains the JSON metadata for the
// struct [PrismObjectEngagementBulkNewResponseError]
type prismObjectEngagementBulkNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementBulkNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementBulkNewResponseResult struct {
	ID      string                                           `json:"id" api:"nullable" format:"uuid"`
	Created bool                                             `json:"created"`
	Error   PrismObjectEngagementBulkNewResponseResultsError `json:"error"`
	// True if the row matched an existing record via the dedupe key.
	Existing bool `json:"existing"`
	// Zero-based position of this row in the request.
	InputIndex int64 `json:"input_index"`
	// True if a matching record was updated.
	Updated bool                                           `json:"updated"`
	JSON    prismObjectEngagementBulkNewResponseResultJSON `json:"-"`
}

// prismObjectEngagementBulkNewResponseResultJSON contains the JSON metadata for
// the struct [PrismObjectEngagementBulkNewResponseResult]
type prismObjectEngagementBulkNewResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	InputIndex  apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementBulkNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementBulkNewResponseResultsError struct {
	Code    string                                               `json:"code"`
	Message string                                               `json:"message"`
	JSON    prismObjectEngagementBulkNewResponseResultsErrorJSON `json:"-"`
}

// prismObjectEngagementBulkNewResponseResultsErrorJSON contains the JSON metadata
// for the struct [PrismObjectEngagementBulkNewResponseResultsError]
type prismObjectEngagementBulkNewResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementBulkNewResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkNewResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectEngagementBulkDeleteResponse struct {
	Results []PrismObjectEngagementBulkDeleteResponseResult `json:"results" api:"required"`
	Summary PrismObjectEngagementBulkDeleteResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectEngagementBulkDeleteResponseJSON     `json:"-"`
}

// prismObjectEngagementBulkDeleteResponseJSON contains the JSON metadata for the
// struct [PrismObjectEngagementBulkDeleteResponse]
type prismObjectEngagementBulkDeleteResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementBulkDeleteResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                               `json:"id" api:"required,nullable"`
	Status PrismObjectEngagementBulkDeleteResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectEngagementBulkDeleteResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectEngagementBulkDeleteResponseResultsRecord `json:"record"`
	JSON   prismObjectEngagementBulkDeleteResponseResultJSON    `json:"-"`
}

// prismObjectEngagementBulkDeleteResponseResultJSON contains the JSON metadata for
// the struct [PrismObjectEngagementBulkDeleteResponseResult]
type prismObjectEngagementBulkDeleteResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementBulkDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementBulkDeleteResponseResultsStatus string

const (
	PrismObjectEngagementBulkDeleteResponseResultsStatusOk    PrismObjectEngagementBulkDeleteResponseResultsStatus = "ok"
	PrismObjectEngagementBulkDeleteResponseResultsStatusError PrismObjectEngagementBulkDeleteResponseResultsStatus = "error"
)

func (r PrismObjectEngagementBulkDeleteResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectEngagementBulkDeleteResponseResultsStatusOk, PrismObjectEngagementBulkDeleteResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectEngagementBulkDeleteResponseResultsError struct {
	Code    string                                                  `json:"code"`
	Message string                                                  `json:"message"`
	JSON    prismObjectEngagementBulkDeleteResponseResultsErrorJSON `json:"-"`
}

// prismObjectEngagementBulkDeleteResponseResultsErrorJSON contains the JSON
// metadata for the struct [PrismObjectEngagementBulkDeleteResponseResultsError]
type prismObjectEngagementBulkDeleteResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementBulkDeleteResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkDeleteResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectEngagementBulkDeleteResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                                   `json:"default"`
	List    interface{}                                              `json:"list"`
	JSON    prismObjectEngagementBulkDeleteResponseResultsRecordJSON `json:"-"`
}

// prismObjectEngagementBulkDeleteResponseResultsRecordJSON contains the JSON
// metadata for the struct [PrismObjectEngagementBulkDeleteResponseResultsRecord]
type prismObjectEngagementBulkDeleteResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementBulkDeleteResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkDeleteResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementBulkDeleteResponseSummary struct {
	Failed    int64                                              `json:"failed" api:"required"`
	Succeeded int64                                              `json:"succeeded" api:"required"`
	Total     int64                                              `json:"total" api:"required"`
	JSON      prismObjectEngagementBulkDeleteResponseSummaryJSON `json:"-"`
}

// prismObjectEngagementBulkDeleteResponseSummaryJSON contains the JSON metadata
// for the struct [PrismObjectEngagementBulkDeleteResponseSummary]
type prismObjectEngagementBulkDeleteResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementBulkDeleteResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkDeleteResponseSummaryJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectEngagementBulkUpdateResponse struct {
	Results []PrismObjectEngagementBulkUpdateResponseResult `json:"results" api:"required"`
	Summary PrismObjectEngagementBulkUpdateResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectEngagementBulkUpdateResponseJSON     `json:"-"`
}

// prismObjectEngagementBulkUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectEngagementBulkUpdateResponse]
type prismObjectEngagementBulkUpdateResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementBulkUpdateResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                               `json:"id" api:"required,nullable"`
	Status PrismObjectEngagementBulkUpdateResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectEngagementBulkUpdateResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectEngagementBulkUpdateResponseResultsRecord `json:"record"`
	JSON   prismObjectEngagementBulkUpdateResponseResultJSON    `json:"-"`
}

// prismObjectEngagementBulkUpdateResponseResultJSON contains the JSON metadata for
// the struct [PrismObjectEngagementBulkUpdateResponseResult]
type prismObjectEngagementBulkUpdateResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementBulkUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementBulkUpdateResponseResultsStatus string

const (
	PrismObjectEngagementBulkUpdateResponseResultsStatusOk    PrismObjectEngagementBulkUpdateResponseResultsStatus = "ok"
	PrismObjectEngagementBulkUpdateResponseResultsStatusError PrismObjectEngagementBulkUpdateResponseResultsStatus = "error"
)

func (r PrismObjectEngagementBulkUpdateResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectEngagementBulkUpdateResponseResultsStatusOk, PrismObjectEngagementBulkUpdateResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectEngagementBulkUpdateResponseResultsError struct {
	Code    string                                                  `json:"code"`
	Message string                                                  `json:"message"`
	JSON    prismObjectEngagementBulkUpdateResponseResultsErrorJSON `json:"-"`
}

// prismObjectEngagementBulkUpdateResponseResultsErrorJSON contains the JSON
// metadata for the struct [PrismObjectEngagementBulkUpdateResponseResultsError]
type prismObjectEngagementBulkUpdateResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementBulkUpdateResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkUpdateResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectEngagementBulkUpdateResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                                   `json:"default"`
	List    interface{}                                              `json:"list"`
	JSON    prismObjectEngagementBulkUpdateResponseResultsRecordJSON `json:"-"`
}

// prismObjectEngagementBulkUpdateResponseResultsRecordJSON contains the JSON
// metadata for the struct [PrismObjectEngagementBulkUpdateResponseResultsRecord]
type prismObjectEngagementBulkUpdateResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementBulkUpdateResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkUpdateResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementBulkUpdateResponseSummary struct {
	Failed    int64                                              `json:"failed" api:"required"`
	Succeeded int64                                              `json:"succeeded" api:"required"`
	Total     int64                                              `json:"total" api:"required"`
	JSON      prismObjectEngagementBulkUpdateResponseSummaryJSON `json:"-"`
}

// prismObjectEngagementBulkUpdateResponseSummaryJSON contains the JSON metadata
// for the struct [PrismObjectEngagementBulkUpdateResponseSummary]
type prismObjectEngagementBulkUpdateResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementBulkUpdateResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementBulkUpdateResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementCountResponse struct {
	// Number of records matching the access scope.
	Total int64                                  `json:"total" api:"required"`
	JSON  prismObjectEngagementCountResponseJSON `json:"-"`
}

// prismObjectEngagementCountResponseJSON contains the JSON metadata for the struct
// [PrismObjectEngagementCountResponse]
type prismObjectEngagementCountResponseJSON struct {
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementCountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementCountResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectEngagementDuplicateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                     `json:"default"`
	List    interface{}                                `json:"list"`
	JSON    prismObjectEngagementDuplicateResponseJSON `json:"-"`
}

// prismObjectEngagementDuplicateResponseJSON contains the JSON metadata for the
// struct [PrismObjectEngagementDuplicateResponse]
type prismObjectEngagementDuplicateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementDuplicateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementDuplicateResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectEngagementFindResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                `json:"default"`
	List    interface{}                           `json:"list"`
	JSON    prismObjectEngagementFindResponseJSON `json:"-"`
}

// prismObjectEngagementFindResponseJSON contains the JSON metadata for the struct
// [PrismObjectEngagementFindResponse]
type prismObjectEngagementFindResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementFindResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementFindResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectEngagementGetResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}               `json:"default"`
	List    interface{}                          `json:"list"`
	JSON    prismObjectEngagementGetResponseJSON `json:"-"`
}

// prismObjectEngagementGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectEngagementGetResponse]
type prismObjectEngagementGetResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementQueryResponse struct {
	Data []PrismObjectEngagementQueryResponseData `json:"data" api:"required"`
	// Accurate end-of-data signal. False when this page contains the last record; true
	// only when at least one more record exists. (Implementation note: the server
	// fetches one extra row internally to determine this — clients never need to
	// overshoot to discover the end.)
	HasMore bool `json:"has_more" api:"required"`
	// Opaque cursor pointing at the next page. Pass it back unchanged. Do not parse
	// it. The current encoding is offset-based (page + limit), so it has the same
	// concurrent-write drift the deprecated `page` parameter has; treat it as a black
	// box so a future keyset cursor is a drop-in. Null when `has_more` is false.
	NextCursor string `json:"next_cursor" api:"nullable"`
	// Only populated when the request set `include_total: true`. Total number of
	// records matching the query, ignoring pagination. Opt-in because it costs an
	// additional pass over the result set.
	Total int64                                  `json:"total" api:"nullable"`
	JSON  prismObjectEngagementQueryResponseJSON `json:"-"`
}

// prismObjectEngagementQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectEngagementQueryResponse]
type prismObjectEngagementQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectEngagementQueryResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                     `json:"properties"`
	Source     []string                                   `json:"source" api:"nullable"`
	JSON       prismObjectEngagementQueryResponseDataJSON `json:"-"`
}

// prismObjectEngagementQueryResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectEngagementQueryResponseData]
type prismObjectEngagementQueryResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectEngagementQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectEngagementRestoreResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                   `json:"default"`
	List    interface{}                              `json:"list"`
	JSON    prismObjectEngagementRestoreResponseJSON `json:"-"`
}

// prismObjectEngagementRestoreResponseJSON contains the JSON metadata for the
// struct [PrismObjectEngagementRestoreResponse]
type prismObjectEngagementRestoreResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementRestoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementRestoreResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectEngagementUpsertResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                  `json:"default"`
	List    interface{}                             `json:"list"`
	JSON    prismObjectEngagementUpsertResponseJSON `json:"-"`
}

// prismObjectEngagementUpsertResponseJSON contains the JSON metadata for the
// struct [PrismObjectEngagementUpsertResponse]
type prismObjectEngagementUpsertResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEngagementUpsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementUpsertResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
}

func (r PrismObjectEngagementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectEngagementUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
	IfMatch               param.Field[string]        `header:"If-Match"`
}

func (r PrismObjectEngagementUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectEngagementListParams struct {
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

// URLQuery serializes [PrismObjectEngagementListParams]'s query parameters as
// `url.Values`.
func (r PrismObjectEngagementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectEngagementDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IfMatch param.Field[string] `header:"If-Match"`
}

type PrismObjectEngagementBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects        param.Field[[]PrismObjectPropertiesParam]              `json:"objects" api:"required"`
	Options        param.Field[PrismObjectEngagementBulkNewParamsOptions] `json:"options"`
	IdempotencyKey param.Field[string]                                    `header:"Idempotency-Key"`
}

func (r PrismObjectEngagementBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectEngagementBulkNewParamsOptions struct {
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
	DedupeBy param.Field[PrismObjectEngagementBulkNewParamsOptionsDedupeByUnion] `json:"dedupe_by"`
	// App/CRM ID for context (optional)
	ListID param.Field[string] `json:"list_id" format:"uuid"`
	// Require app_stage for every row in the selected list. app_stage is a reserved
	// list-scoped alias for native status.
	RequireListStage param.Field[bool] `json:"require_list_stage"`
	// Patch a deduplicated record with the supplied properties instead of skipping it.
	UpdateExisting param.Field[bool] `json:"update_existing"`
}

func (r PrismObjectEngagementBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Property slug to deduplicate on. A single-element array is also accepted;
// compound (multi-slug) dedupe is not supported yet and is rejected with guidance.
//
// Satisfied by [shared.UnionString],
// [PrismObjectEngagementBulkNewParamsOptionsDedupeByArray].
type PrismObjectEngagementBulkNewParamsOptionsDedupeByUnion interface {
	ImplementsPrismObjectEngagementBulkNewParamsOptionsDedupeByUnion()
}

type PrismObjectEngagementBulkNewParamsOptionsDedupeByArray []string

func (r PrismObjectEngagementBulkNewParamsOptionsDedupeByArray) ImplementsPrismObjectEngagementBulkNewParamsOptionsDedupeByUnion() {
}

type PrismObjectEngagementBulkDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]   `path:"teamId" api:"required" format:"uuid"`
	IDs            param.Field[[]string] `json:"ids" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string]   `header:"Idempotency-Key"`
}

func (r PrismObjectEngagementBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectEngagementBulkUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]                                      `path:"teamId" api:"required" format:"uuid"`
	Items          param.Field[[]PrismObjectEngagementBulkUpdateParamsItem] `json:"items" api:"required"`
	IdempotencyKey param.Field[string]                                      `header:"Idempotency-Key"`
}

func (r PrismObjectEngagementBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object with `id` plus the same property body shape as PATCH
// (`default`/`list`/`extended`).
type PrismObjectEngagementBulkUpdateParamsItem struct {
	ID          param.Field[string]    `json:"id" api:"required" format:"uuid"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r PrismObjectEngagementBulkUpdateParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectEngagementCountParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the count to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectEngagementCountParams]'s query parameters as
// `url.Values`.
func (r PrismObjectEngagementCountParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectEngagementDuplicateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectEngagementFindParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the lookup to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectEngagementFindParams]'s query parameters as
// `url.Values`.
func (r PrismObjectEngagementFindParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectEngagementGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Comma-separated property slugs to return. Use dot notation for relationships.
	// `id` is always returned at the top level. Defaults to all properties.
	Select param.Field[string] `query:"select"`
}

// URLQuery serializes [PrismObjectEngagementGetParams]'s query parameters as
// `url.Values`.
func (r PrismObjectEngagementGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectEngagementQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string]                                  `path:"teamId" api:"required" format:"uuid"`
	Query  param.Field[PrismObjectEngagementQueryParamsQuery]   `json:"query" api:"required"`
	ID     param.Field[PrismObjectEngagementQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes  param.Field[[]string]                                `json:"boxes"`
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

func (r PrismObjectEngagementQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectEngagementQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name). `id` is always returned at the top level of each
	// row and does not need to be selected.
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectEngagementQueryParamsQueryCombinator] `json:"combinator"`
	// Opaque cursor from a previous response's `next_cursor`. Pass it back unchanged
	// to fetch the next page. When set, `page` and `limit` are derived from the cursor
	// and any explicit values are ignored.
	Cursor param.Field[string] `json:"cursor"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectEngagementQueryParamsQueryFilterUnion] `json:"filter"`
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
	Sort param.Field[[]map[string]PrismObjectEngagementQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectEngagementQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectEngagementQueryParamsQueryCombinator string

const (
	PrismObjectEngagementQueryParamsQueryCombinatorAnd PrismObjectEngagementQueryParamsQueryCombinator = "AND"
	PrismObjectEngagementQueryParamsQueryCombinatorOr  PrismObjectEngagementQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectEngagementQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectEngagementQueryParamsQueryCombinatorAnd, PrismObjectEngagementQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

type PrismObjectEngagementQueryParamsQueryFilter struct {
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

func (r PrismObjectEngagementQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilter) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

// Satisfied by [PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterEq],
// [PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterNe],
// [PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterLt],
// [PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterGt],
// [PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterLte],
// [PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterGte],
// [PrismObjectEngagementQueryParamsQueryFilterContains],
// [PrismObjectEngagementQueryParamsQueryFilterBeginsWith],
// [PrismObjectEngagementQueryParamsQueryFilterEndsWith],
// [PrismObjectEngagementQueryParamsQueryFilterNotContains],
// [PrismObjectEngagementQueryParamsQueryFilterExists],
// [PrismObjectEngagementQueryParamsQueryFilterNotExists],
// [PrismObjectEngagementQueryParamsQueryFilterIsNull],
// [PrismObjectEngagementQueryParamsQueryFilterIsNotNull],
// [PrismObjectEngagementQueryParamsQueryFilterBetween],
// [PrismObjectEngagementQueryParamsQueryFilterIn],
// [PrismObjectEngagementQueryParamsQueryFilterNotIn],
// [PrismObjectEngagementQueryParamsQueryFilter].
type PrismObjectEngagementQueryParamsQueryFilterUnion interface {
	implementsPrismObjectEngagementQueryParamsQueryFilterUnion()
}

type PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterEq struct {
	Equals param.Field[PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion] `json:"=" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterEq) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterEq) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion interface {
	ImplementsPrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion()
}

type PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterNe struct {
	NotEquals param.Field[PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion] `json:"!=" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterNe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterNe) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion interface {
	ImplementsPrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion()
}

type PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterLt struct {
	Less param.Field[string] `json:"<" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterLt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterLt) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterGt struct {
	Greater param.Field[string] `json:">" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterGt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterGt) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterLte struct {
	LessOrEquals param.Field[string] `json:"<=" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterLte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterLte) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterGte struct {
	GreaterOrEquals param.Field[string] `json:">=" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterGte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterPrismQueryFilterGte) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterContains struct {
	Contains param.Field[PrismObjectEngagementQueryParamsQueryFilterContainsContainsUnion] `json:"contains" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterContains) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectEngagementQueryParamsQueryFilterContainsContainsArray].
type PrismObjectEngagementQueryParamsQueryFilterContainsContainsUnion interface {
	ImplementsPrismObjectEngagementQueryParamsQueryFilterContainsContainsUnion()
}

type PrismObjectEngagementQueryParamsQueryFilterContainsContainsArray []string

func (r PrismObjectEngagementQueryParamsQueryFilterContainsContainsArray) ImplementsPrismObjectEngagementQueryParamsQueryFilterContainsContainsUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterBeginsWith struct {
	BeginsWith param.Field[string] `json:"begins_with" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterBeginsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterBeginsWith) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterEndsWith struct {
	EndsWith param.Field[string] `json:"ends_with" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterEndsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterEndsWith) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterNotContains struct {
	NotContains param.Field[string] `json:"not_contains" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterNotContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterNotContains) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterExists struct {
	Exists param.Field[bool] `json:"exists" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterExists) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterNotExists struct {
	NotExists param.Field[bool] `json:"not_exists" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterNotExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterNotExists) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterIsNull struct {
	IsNull param.Field[PrismObjectEngagementQueryParamsQueryFilterIsNullIsNullUnion] `json:"is_null" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterIsNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterIsNull) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectEngagementQueryParamsQueryFilterIsNullIsNullArray].
type PrismObjectEngagementQueryParamsQueryFilterIsNullIsNullUnion interface {
	ImplementsPrismObjectEngagementQueryParamsQueryFilterIsNullIsNullUnion()
}

type PrismObjectEngagementQueryParamsQueryFilterIsNullIsNullArray []string

func (r PrismObjectEngagementQueryParamsQueryFilterIsNullIsNullArray) ImplementsPrismObjectEngagementQueryParamsQueryFilterIsNullIsNullUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterIsNotNull struct {
	IsNotNull param.Field[PrismObjectEngagementQueryParamsQueryFilterIsNotNullIsNotNullUnion] `json:"is_not_null" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterIsNotNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterIsNotNull) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectEngagementQueryParamsQueryFilterIsNotNullIsNotNullArray].
type PrismObjectEngagementQueryParamsQueryFilterIsNotNullIsNotNullUnion interface {
	ImplementsPrismObjectEngagementQueryParamsQueryFilterIsNotNullIsNotNullUnion()
}

type PrismObjectEngagementQueryParamsQueryFilterIsNotNullIsNotNullArray []string

func (r PrismObjectEngagementQueryParamsQueryFilterIsNotNullIsNotNullArray) ImplementsPrismObjectEngagementQueryParamsQueryFilterIsNotNullIsNotNullUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterBetween struct {
	Between param.Field[PrismObjectEngagementQueryParamsQueryFilterBetweenBetweenUnion] `json:"between" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterBetween) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterBetween) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectEngagementQueryParamsQueryFilterBetweenBetweenArray].
type PrismObjectEngagementQueryParamsQueryFilterBetweenBetweenUnion interface {
	ImplementsPrismObjectEngagementQueryParamsQueryFilterBetweenBetweenUnion()
}

type PrismObjectEngagementQueryParamsQueryFilterBetweenBetweenArray []string

func (r PrismObjectEngagementQueryParamsQueryFilterBetweenBetweenArray) ImplementsPrismObjectEngagementQueryParamsQueryFilterBetweenBetweenUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterIn struct {
	In param.Field[[]string] `json:"in" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterIn) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

type PrismObjectEngagementQueryParamsQueryFilterNotIn struct {
	NotIn param.Field[[]string] `json:"not_in" api:"required"`
}

func (r PrismObjectEngagementQueryParamsQueryFilterNotIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectEngagementQueryParamsQueryFilterNotIn) implementsPrismObjectEngagementQueryParamsQueryFilterUnion() {
}

type PrismObjectEngagementQueryParamsQuerySort string

const (
	PrismObjectEngagementQueryParamsQuerySortAsc  PrismObjectEngagementQueryParamsQuerySort = "asc"
	PrismObjectEngagementQueryParamsQuerySortDesc PrismObjectEngagementQueryParamsQuerySort = "desc"
)

func (r PrismObjectEngagementQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectEngagementQueryParamsQuerySortAsc, PrismObjectEngagementQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectEngagementQueryParamsIDArray].
type PrismObjectEngagementQueryParamsIDUnion interface {
	ImplementsPrismObjectEngagementQueryParamsIDUnion()
}

type PrismObjectEngagementQueryParamsIDArray []string

func (r PrismObjectEngagementQueryParamsIDArray) ImplementsPrismObjectEngagementQueryParamsIDUnion() {
}

type PrismObjectEngagementRestoreParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectEngagementUpsertParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	// Scope the upsert to a specific list/app. Required to match or write list-scoped
	// properties, including `app_stage`.
	ListID         param.Field[string] `query:"list_id" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

func (r PrismObjectEngagementUpsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

// URLQuery serializes [PrismObjectEngagementUpsertParams]'s query parameters as
// `url.Values`.
func (r PrismObjectEngagementUpsertParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
