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
	path := fmt.Sprintf("v2/prism/%s/document", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch object
func (r *PrismObjectDocumentService) Update(ctx context.Context, documentID string, params PrismObjectDocumentUpdateParams, opts ...option.RequestOption) (res *PrismObjectDocumentUpdateResponse, err error) {
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
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/document/%s", params.TeamID, documentID)
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
func (r *PrismObjectDocumentService) List(ctx context.Context, params PrismObjectDocumentListParams, opts ...option.RequestOption) (res *PrismObjectDocumentListResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete object
func (r *PrismObjectDocumentService) Delete(ctx context.Context, documentID string, params PrismObjectDocumentDeleteParams, opts ...option.RequestOption) (err error) {
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
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/document/%s", params.TeamID, documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: small batches complete synchronously and return 200 with
// the final `ImportJob`; large batches start an async job, return 202 with
// `status: processing` and a `Location` header, and can be polled via
// `GET /v2/prism/{teamId}/imports/{jobId}`.
func (r *PrismObjectDocumentService) BulkNew(ctx context.Context, params PrismObjectDocumentBulkNewParams, opts ...option.RequestOption) (res *PrismObjectDocumentBulkNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/document/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Soft-delete up to 100 records in a single call. Same partial-success contract as
// batch/update.
func (r *PrismObjectDocumentService) BulkDelete(ctx context.Context, params PrismObjectDocumentBulkDeleteParams, opts ...option.RequestOption) (res *PrismObjectDocumentBulkDeleteResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/document/batch/delete", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch up to 100 records in a single call. Each item is attempted independently —
// failures don't abort the batch. Inspect `results[].status` per item.
func (r *PrismObjectDocumentService) BulkUpdate(ctx context.Context, params PrismObjectDocumentBulkUpdateParams, opts ...option.RequestOption) (res *PrismObjectDocumentBulkUpdateResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/document/batch/update", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the total number of records of this object type that the caller can see.
// Avoids the page-overshoot anti-pattern — clients no longer need to keep paging
// until `has_more` flips false to discover the total. Currently does not apply
// query filters; for a filtered total, pass `include_total: true` in a POST
// `/query` body.
func (r *PrismObjectDocumentService) Count(ctx context.Context, params PrismObjectDocumentCountParams, opts ...option.RequestOption) (res *PrismObjectDocumentCountResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/document/count", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Duplicate object
func (r *PrismObjectDocumentService) Duplicate(ctx context.Context, documentID string, params PrismObjectDocumentDuplicateParams, opts ...option.RequestOption) (res *PrismObjectDocumentDuplicateResponse, err error) {
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
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/document/%s/duplicate", params.TeamID, documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns the single record whose property `{slug}` equals `{value}`. 404 if
// nothing matches; 409 if more than one record matches.
func (r *PrismObjectDocumentService) Find(ctx context.Context, slug string, value string, params PrismObjectDocumentFindParams, opts ...option.RequestOption) (res *PrismObjectDocumentFindResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/document/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectDocumentService) Get(ctx context.Context, documentID string, params PrismObjectDocumentGetParams, opts ...option.RequestOption) (res *PrismObjectDocumentGetResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Query
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
	path := fmt.Sprintf("v2/prism/%s/document/query", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Restore object
func (r *PrismObjectDocumentService) Restore(ctx context.Context, documentID string, params PrismObjectDocumentRestoreParams, opts ...option.RequestOption) (res *PrismObjectDocumentRestoreResponse, err error) {
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
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/document/%s/restore", params.TeamID, documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Idempotent create-or-update keyed on `{slug}={value}`. If exactly one record
// matches, it is patched and 200 is returned. If none match, a new record is
// created (with the lookup property set if absent) and 201 is returned. If
// multiple records match, 409 is returned and you should patch by id instead.
func (r *PrismObjectDocumentService) Upsert(ctx context.Context, slug string, value string, params PrismObjectDocumentUpsertParams, opts ...option.RequestOption) (res *PrismObjectDocumentUpsertResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/document/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDocumentNewResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}             `json:"default"`
	List    interface{}                        `json:"list"`
	JSON    prismObjectDocumentNewResponseJSON `json:"-"`
}

// prismObjectDocumentNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentNewResponse]
type prismObjectDocumentNewResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
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
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                `json:"default"`
	List    interface{}                           `json:"list"`
	JSON    prismObjectDocumentUpdateResponseJSON `json:"-"`
}

// prismObjectDocumentUpdateResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentUpdateResponse]
type prismObjectDocumentUpdateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentListResponse struct {
	Data []PrismObjectDocumentListResponseData `json:"data" api:"required"`
	// Accurate end-of-data signal — false on the last page, never forces clients to
	// overshoot.
	HasMore    bool   `json:"has_more" api:"required"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// Populated only when `?include_total=true` was passed.
	Total int64                               `json:"total" api:"nullable"`
	JSON  prismObjectDocumentListResponseJSON `json:"-"`
}

// prismObjectDocumentListResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentListResponse]
type prismObjectDocumentListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentListResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectDocumentListResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                  `json:"properties"`
	Source     []string                                `json:"source" api:"nullable"`
	JSON       prismObjectDocumentListResponseDataJSON `json:"-"`
}

// prismObjectDocumentListResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectDocumentListResponseData]
type prismObjectDocumentListResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectDocumentListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Status snapshot of an import job. Same shape used by the POST /import response
// and by GET /imports/{jobId}.
type PrismObjectDocumentBulkNewResponse struct {
	// Null for sync imports (results inlined). Set for async imports.
	JobID  string                                   `json:"job_id" api:"required,nullable"`
	Status PrismObjectDocumentBulkNewResponseStatus `json:"status" api:"required"`
	// Total number of rows in the import.
	Total     int64     `json:"total" api:"required"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Set when status=failed; describes the job-level failure (not per-row).
	Error     PrismObjectDocumentBulkNewResponseError `json:"error"`
	ExpiresAt time.Time                               `json:"expires_at" format:"date-time"`
	Failed    int64                                   `json:"failed"`
	// Rows that have been attempted (succeeded + failed).
	Processed int64 `json:"processed"`
	// Per-row outcomes. Always present for sync imports; populated for async imports
	// once the job reaches `complete`.
	Results   []PrismObjectDocumentBulkNewResponseResult `json:"results"`
	Succeeded int64                                      `json:"succeeded"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      prismObjectDocumentBulkNewResponseJSON     `json:"-"`
}

// prismObjectDocumentBulkNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentBulkNewResponse]
type prismObjectDocumentBulkNewResponseJSON struct {
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

func (r *PrismObjectDocumentBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentBulkNewResponseStatus string

const (
	PrismObjectDocumentBulkNewResponseStatusComplete   PrismObjectDocumentBulkNewResponseStatus = "complete"
	PrismObjectDocumentBulkNewResponseStatusProcessing PrismObjectDocumentBulkNewResponseStatus = "processing"
	PrismObjectDocumentBulkNewResponseStatusFailed     PrismObjectDocumentBulkNewResponseStatus = "failed"
)

func (r PrismObjectDocumentBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectDocumentBulkNewResponseStatusComplete, PrismObjectDocumentBulkNewResponseStatusProcessing, PrismObjectDocumentBulkNewResponseStatusFailed:
		return true
	}
	return false
}

// Set when status=failed; describes the job-level failure (not per-row).
type PrismObjectDocumentBulkNewResponseError struct {
	Code    string                                      `json:"code"`
	Message string                                      `json:"message"`
	JSON    prismObjectDocumentBulkNewResponseErrorJSON `json:"-"`
}

// prismObjectDocumentBulkNewResponseErrorJSON contains the JSON metadata for the
// struct [PrismObjectDocumentBulkNewResponseError]
type prismObjectDocumentBulkNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentBulkNewResponseResult struct {
	ID      string                                         `json:"id" api:"nullable" format:"uuid"`
	Created bool                                           `json:"created"`
	Error   PrismObjectDocumentBulkNewResponseResultsError `json:"error"`
	// True if the row matched an existing record via the dedupe key.
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

type PrismObjectDocumentBulkNewResponseResultsError struct {
	Code    string                                             `json:"code"`
	Message string                                             `json:"message"`
	JSON    prismObjectDocumentBulkNewResponseResultsErrorJSON `json:"-"`
}

// prismObjectDocumentBulkNewResponseResultsErrorJSON contains the JSON metadata
// for the struct [PrismObjectDocumentBulkNewResponseResultsError]
type prismObjectDocumentBulkNewResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkNewResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkNewResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectDocumentBulkDeleteResponse struct {
	Results []PrismObjectDocumentBulkDeleteResponseResult `json:"results" api:"required"`
	Summary PrismObjectDocumentBulkDeleteResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectDocumentBulkDeleteResponseJSON     `json:"-"`
}

// prismObjectDocumentBulkDeleteResponseJSON contains the JSON metadata for the
// struct [PrismObjectDocumentBulkDeleteResponse]
type prismObjectDocumentBulkDeleteResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentBulkDeleteResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                             `json:"id" api:"required,nullable"`
	Status PrismObjectDocumentBulkDeleteResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectDocumentBulkDeleteResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectDocumentBulkDeleteResponseResultsRecord `json:"record"`
	JSON   prismObjectDocumentBulkDeleteResponseResultJSON    `json:"-"`
}

// prismObjectDocumentBulkDeleteResponseResultJSON contains the JSON metadata for
// the struct [PrismObjectDocumentBulkDeleteResponseResult]
type prismObjectDocumentBulkDeleteResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentBulkDeleteResponseResultsStatus string

const (
	PrismObjectDocumentBulkDeleteResponseResultsStatusOk    PrismObjectDocumentBulkDeleteResponseResultsStatus = "ok"
	PrismObjectDocumentBulkDeleteResponseResultsStatusError PrismObjectDocumentBulkDeleteResponseResultsStatus = "error"
)

func (r PrismObjectDocumentBulkDeleteResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectDocumentBulkDeleteResponseResultsStatusOk, PrismObjectDocumentBulkDeleteResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectDocumentBulkDeleteResponseResultsError struct {
	Code    string                                                `json:"code"`
	Message string                                                `json:"message"`
	JSON    prismObjectDocumentBulkDeleteResponseResultsErrorJSON `json:"-"`
}

// prismObjectDocumentBulkDeleteResponseResultsErrorJSON contains the JSON metadata
// for the struct [PrismObjectDocumentBulkDeleteResponseResultsError]
type prismObjectDocumentBulkDeleteResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkDeleteResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkDeleteResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDocumentBulkDeleteResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                                 `json:"default"`
	List    interface{}                                            `json:"list"`
	JSON    prismObjectDocumentBulkDeleteResponseResultsRecordJSON `json:"-"`
}

// prismObjectDocumentBulkDeleteResponseResultsRecordJSON contains the JSON
// metadata for the struct [PrismObjectDocumentBulkDeleteResponseResultsRecord]
type prismObjectDocumentBulkDeleteResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkDeleteResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkDeleteResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentBulkDeleteResponseSummary struct {
	Failed    int64                                            `json:"failed" api:"required"`
	Succeeded int64                                            `json:"succeeded" api:"required"`
	Total     int64                                            `json:"total" api:"required"`
	JSON      prismObjectDocumentBulkDeleteResponseSummaryJSON `json:"-"`
}

// prismObjectDocumentBulkDeleteResponseSummaryJSON contains the JSON metadata for
// the struct [PrismObjectDocumentBulkDeleteResponseSummary]
type prismObjectDocumentBulkDeleteResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkDeleteResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkDeleteResponseSummaryJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectDocumentBulkUpdateResponse struct {
	Results []PrismObjectDocumentBulkUpdateResponseResult `json:"results" api:"required"`
	Summary PrismObjectDocumentBulkUpdateResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectDocumentBulkUpdateResponseJSON     `json:"-"`
}

// prismObjectDocumentBulkUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectDocumentBulkUpdateResponse]
type prismObjectDocumentBulkUpdateResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentBulkUpdateResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                             `json:"id" api:"required,nullable"`
	Status PrismObjectDocumentBulkUpdateResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectDocumentBulkUpdateResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectDocumentBulkUpdateResponseResultsRecord `json:"record"`
	JSON   prismObjectDocumentBulkUpdateResponseResultJSON    `json:"-"`
}

// prismObjectDocumentBulkUpdateResponseResultJSON contains the JSON metadata for
// the struct [PrismObjectDocumentBulkUpdateResponseResult]
type prismObjectDocumentBulkUpdateResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentBulkUpdateResponseResultsStatus string

const (
	PrismObjectDocumentBulkUpdateResponseResultsStatusOk    PrismObjectDocumentBulkUpdateResponseResultsStatus = "ok"
	PrismObjectDocumentBulkUpdateResponseResultsStatusError PrismObjectDocumentBulkUpdateResponseResultsStatus = "error"
)

func (r PrismObjectDocumentBulkUpdateResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectDocumentBulkUpdateResponseResultsStatusOk, PrismObjectDocumentBulkUpdateResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectDocumentBulkUpdateResponseResultsError struct {
	Code    string                                                `json:"code"`
	Message string                                                `json:"message"`
	JSON    prismObjectDocumentBulkUpdateResponseResultsErrorJSON `json:"-"`
}

// prismObjectDocumentBulkUpdateResponseResultsErrorJSON contains the JSON metadata
// for the struct [PrismObjectDocumentBulkUpdateResponseResultsError]
type prismObjectDocumentBulkUpdateResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkUpdateResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkUpdateResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDocumentBulkUpdateResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                                 `json:"default"`
	List    interface{}                                            `json:"list"`
	JSON    prismObjectDocumentBulkUpdateResponseResultsRecordJSON `json:"-"`
}

// prismObjectDocumentBulkUpdateResponseResultsRecordJSON contains the JSON
// metadata for the struct [PrismObjectDocumentBulkUpdateResponseResultsRecord]
type prismObjectDocumentBulkUpdateResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkUpdateResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkUpdateResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentBulkUpdateResponseSummary struct {
	Failed    int64                                            `json:"failed" api:"required"`
	Succeeded int64                                            `json:"succeeded" api:"required"`
	Total     int64                                            `json:"total" api:"required"`
	JSON      prismObjectDocumentBulkUpdateResponseSummaryJSON `json:"-"`
}

// prismObjectDocumentBulkUpdateResponseSummaryJSON contains the JSON metadata for
// the struct [PrismObjectDocumentBulkUpdateResponseSummary]
type prismObjectDocumentBulkUpdateResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentBulkUpdateResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentBulkUpdateResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentCountResponse struct {
	// Number of records matching the access scope.
	Total int64                                `json:"total" api:"required"`
	JSON  prismObjectDocumentCountResponseJSON `json:"-"`
}

// prismObjectDocumentCountResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentCountResponse]
type prismObjectDocumentCountResponseJSON struct {
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentCountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentCountResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDocumentDuplicateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                   `json:"default"`
	List    interface{}                              `json:"list"`
	JSON    prismObjectDocumentDuplicateResponseJSON `json:"-"`
}

// prismObjectDocumentDuplicateResponseJSON contains the JSON metadata for the
// struct [PrismObjectDocumentDuplicateResponse]
type prismObjectDocumentDuplicateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
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
type PrismObjectDocumentFindResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}              `json:"default"`
	List    interface{}                         `json:"list"`
	JSON    prismObjectDocumentFindResponseJSON `json:"-"`
}

// prismObjectDocumentFindResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentFindResponse]
type prismObjectDocumentFindResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentFindResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentFindResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDocumentGetResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}             `json:"default"`
	List    interface{}                        `json:"list"`
	JSON    prismObjectDocumentGetResponseJSON `json:"-"`
}

// prismObjectDocumentGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentGetResponse]
type prismObjectDocumentGetResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
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
	JSON  prismObjectDocumentQueryResponseJSON `json:"-"`
}

// prismObjectDocumentQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentQueryResponse]
type prismObjectDocumentQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectDocumentQueryResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                   `json:"properties"`
	Source     []string                                 `json:"source" api:"nullable"`
	JSON       prismObjectDocumentQueryResponseDataJSON `json:"-"`
}

// prismObjectDocumentQueryResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectDocumentQueryResponseData]
type prismObjectDocumentQueryResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectDocumentQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDocumentRestoreResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                 `json:"default"`
	List    interface{}                            `json:"list"`
	JSON    prismObjectDocumentRestoreResponseJSON `json:"-"`
}

// prismObjectDocumentRestoreResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentRestoreResponse]
type prismObjectDocumentRestoreResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentRestoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentRestoreResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectDocumentUpsertResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                `json:"default"`
	List    interface{}                           `json:"list"`
	JSON    prismObjectDocumentUpsertResponseJSON `json:"-"`
}

// prismObjectDocumentUpsertResponseJSON contains the JSON metadata for the struct
// [PrismObjectDocumentUpsertResponse]
type prismObjectDocumentUpsertResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentUpsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentUpsertResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
}

func (r PrismObjectDocumentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectDocumentUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
	IfMatch               param.Field[string]        `header:"If-Match"`
}

func (r PrismObjectDocumentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectDocumentListParams struct {
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

// URLQuery serializes [PrismObjectDocumentListParams]'s query parameters as
// `url.Values`.
func (r PrismObjectDocumentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectDocumentDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IfMatch param.Field[string] `header:"If-Match"`
}

type PrismObjectDocumentBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects        param.Field[[]PrismObjectPropertiesParam]            `json:"objects" api:"required"`
	Options        param.Field[PrismObjectDocumentBulkNewParamsOptions] `json:"options"`
	IdempotencyKey param.Field[string]                                  `header:"Idempotency-Key"`
}

func (r PrismObjectDocumentBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDocumentBulkNewParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// When true, unknown values for select/multiselect properties are created as new
	// options instead of failing the import
	CreateMissingOptions param.Field[bool] `json:"create_missing_options"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
	// App/CRM ID for context (optional)
	ListID param.Field[string] `json:"list_id" format:"uuid"`
}

func (r PrismObjectDocumentBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDocumentBulkDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]   `path:"teamId" api:"required" format:"uuid"`
	IDs            param.Field[[]string] `json:"ids" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string]   `header:"Idempotency-Key"`
}

func (r PrismObjectDocumentBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDocumentBulkUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]                                    `path:"teamId" api:"required" format:"uuid"`
	Items          param.Field[[]PrismObjectDocumentBulkUpdateParamsItem] `json:"items" api:"required"`
	IdempotencyKey param.Field[string]                                    `header:"Idempotency-Key"`
}

func (r PrismObjectDocumentBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object with `id` plus the same property body shape as PATCH
// (`default`/`list`/`extended`).
type PrismObjectDocumentBulkUpdateParamsItem struct {
	ID          param.Field[string]    `json:"id" api:"required" format:"uuid"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r PrismObjectDocumentBulkUpdateParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDocumentCountParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the count to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectDocumentCountParams]'s query parameters as
// `url.Values`.
func (r PrismObjectDocumentCountParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectDocumentDuplicateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectDocumentFindParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the lookup to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectDocumentFindParams]'s query parameters as
// `url.Values`.
func (r PrismObjectDocumentFindParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectDocumentGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Comma-separated property slugs to return. Use dot notation for relationships.
	// `id` is always returned at the top level. Defaults to all properties.
	Select param.Field[string] `query:"select"`
}

// URLQuery serializes [PrismObjectDocumentGetParams]'s query parameters as
// `url.Values`.
func (r PrismObjectDocumentGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectDocumentQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string]                                `path:"teamId" api:"required" format:"uuid"`
	Query  param.Field[PrismObjectDocumentQueryParamsQuery]   `json:"query" api:"required"`
	ID     param.Field[PrismObjectDocumentQueryParamsIDUnion] `json:"id" format:"uuid"`
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

func (r PrismObjectDocumentQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDocumentQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name). `id` is always returned at the top level of each
	// row and does not need to be selected.
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectDocumentQueryParamsQueryCombinator] `json:"combinator"`
	// Opaque cursor from a previous response's `next_cursor`. Pass it back unchanged
	// to fetch the next page. When set, `page` and `limit` are derived from the cursor
	// and any explicit values are ignored.
	Cursor param.Field[string] `json:"cursor"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectDocumentQueryParamsQueryFilterUnion] `json:"filter"`
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

type PrismObjectDocumentQueryParamsQueryFilter struct {
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

func (r PrismObjectDocumentQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilter) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

// Satisfied by [PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterEq],
// [PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterNe],
// [PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterLt],
// [PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterGt],
// [PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterLte],
// [PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterGte],
// [PrismObjectDocumentQueryParamsQueryFilterContains],
// [PrismObjectDocumentQueryParamsQueryFilterBeginsWith],
// [PrismObjectDocumentQueryParamsQueryFilterEndsWith],
// [PrismObjectDocumentQueryParamsQueryFilterNotContains],
// [PrismObjectDocumentQueryParamsQueryFilterExists],
// [PrismObjectDocumentQueryParamsQueryFilterNotExists],
// [PrismObjectDocumentQueryParamsQueryFilterIsNull],
// [PrismObjectDocumentQueryParamsQueryFilterIsNotNull],
// [PrismObjectDocumentQueryParamsQueryFilterBetween],
// [PrismObjectDocumentQueryParamsQueryFilterIn],
// [PrismObjectDocumentQueryParamsQueryFilterNotIn],
// [PrismObjectDocumentQueryParamsQueryFilter].
type PrismObjectDocumentQueryParamsQueryFilterUnion interface {
	implementsPrismObjectDocumentQueryParamsQueryFilterUnion()
}

type PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterEq struct {
	Equals param.Field[PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterEqUnion] `json:"=" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterEq) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterEq) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterEqUnion interface {
	ImplementsPrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterEqUnion()
}

type PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterNe struct {
	NotEquals param.Field[PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterNeUnion] `json:"!=" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterNe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterNe) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterNeUnion interface {
	ImplementsPrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterNeUnion()
}

type PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterLt struct {
	Less param.Field[string] `json:"<" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterLt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterLt) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterGt struct {
	Greater param.Field[string] `json:">" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterGt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterGt) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterLte struct {
	LessOrEquals param.Field[string] `json:"<=" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterLte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterLte) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterGte struct {
	GreaterOrEquals param.Field[string] `json:">=" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterGte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterPrismQueryFilterGte) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterContains struct {
	Contains param.Field[PrismObjectDocumentQueryParamsQueryFilterContainsContainsUnion] `json:"contains" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterContains) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectDocumentQueryParamsQueryFilterContainsContainsArray].
type PrismObjectDocumentQueryParamsQueryFilterContainsContainsUnion interface {
	ImplementsPrismObjectDocumentQueryParamsQueryFilterContainsContainsUnion()
}

type PrismObjectDocumentQueryParamsQueryFilterContainsContainsArray []string

func (r PrismObjectDocumentQueryParamsQueryFilterContainsContainsArray) ImplementsPrismObjectDocumentQueryParamsQueryFilterContainsContainsUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterBeginsWith struct {
	BeginsWith param.Field[string] `json:"begins_with" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterBeginsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterBeginsWith) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterEndsWith struct {
	EndsWith param.Field[string] `json:"ends_with" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterEndsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterEndsWith) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterNotContains struct {
	NotContains param.Field[string] `json:"not_contains" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterNotContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterNotContains) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterExists struct {
	Exists param.Field[bool] `json:"exists" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterExists) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterNotExists struct {
	NotExists param.Field[bool] `json:"not_exists" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterNotExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterNotExists) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterIsNull struct {
	IsNull param.Field[PrismObjectDocumentQueryParamsQueryFilterIsNullIsNullUnion] `json:"is_null" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterIsNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterIsNull) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectDocumentQueryParamsQueryFilterIsNullIsNullArray].
type PrismObjectDocumentQueryParamsQueryFilterIsNullIsNullUnion interface {
	ImplementsPrismObjectDocumentQueryParamsQueryFilterIsNullIsNullUnion()
}

type PrismObjectDocumentQueryParamsQueryFilterIsNullIsNullArray []string

func (r PrismObjectDocumentQueryParamsQueryFilterIsNullIsNullArray) ImplementsPrismObjectDocumentQueryParamsQueryFilterIsNullIsNullUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterIsNotNull struct {
	IsNotNull param.Field[PrismObjectDocumentQueryParamsQueryFilterIsNotNullIsNotNullUnion] `json:"is_not_null" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterIsNotNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterIsNotNull) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectDocumentQueryParamsQueryFilterIsNotNullIsNotNullArray].
type PrismObjectDocumentQueryParamsQueryFilterIsNotNullIsNotNullUnion interface {
	ImplementsPrismObjectDocumentQueryParamsQueryFilterIsNotNullIsNotNullUnion()
}

type PrismObjectDocumentQueryParamsQueryFilterIsNotNullIsNotNullArray []string

func (r PrismObjectDocumentQueryParamsQueryFilterIsNotNullIsNotNullArray) ImplementsPrismObjectDocumentQueryParamsQueryFilterIsNotNullIsNotNullUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterBetween struct {
	Between param.Field[PrismObjectDocumentQueryParamsQueryFilterBetweenBetweenUnion] `json:"between" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterBetween) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterBetween) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectDocumentQueryParamsQueryFilterBetweenBetweenArray].
type PrismObjectDocumentQueryParamsQueryFilterBetweenBetweenUnion interface {
	ImplementsPrismObjectDocumentQueryParamsQueryFilterBetweenBetweenUnion()
}

type PrismObjectDocumentQueryParamsQueryFilterBetweenBetweenArray []string

func (r PrismObjectDocumentQueryParamsQueryFilterBetweenBetweenArray) ImplementsPrismObjectDocumentQueryParamsQueryFilterBetweenBetweenUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterIn struct {
	In param.Field[[]string] `json:"in" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterIn) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
}

type PrismObjectDocumentQueryParamsQueryFilterNotIn struct {
	NotIn param.Field[[]string] `json:"not_in" api:"required"`
}

func (r PrismObjectDocumentQueryParamsQueryFilterNotIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectDocumentQueryParamsQueryFilterNotIn) implementsPrismObjectDocumentQueryParamsQueryFilterUnion() {
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
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectDocumentUpsertParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
}

func (r PrismObjectDocumentUpsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}
