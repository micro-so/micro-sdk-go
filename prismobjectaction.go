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

// PrismObjectActionService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectActionService] method instead.
type PrismObjectActionService struct {
	Options []option.RequestOption
	Grant   *PrismObjectActionGrantService
}

// NewPrismObjectActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectActionService(opts ...option.RequestOption) (r *PrismObjectActionService) {
	r = &PrismObjectActionService{}
	r.Options = opts
	r.Grant = NewPrismObjectActionGrantService(opts...)
	return
}

// Create object
func (r *PrismObjectActionService) New(ctx context.Context, params PrismObjectActionNewParams, opts ...option.RequestOption) (res *PrismObjectActionNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/action", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch object
func (r *PrismObjectActionService) Update(ctx context.Context, actionID string, params PrismObjectActionUpdateParams, opts ...option.RequestOption) (res *PrismObjectActionUpdateResponse, err error) {
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
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/action/%s", params.TeamID, actionID)
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
func (r *PrismObjectActionService) List(ctx context.Context, params PrismObjectActionListParams, opts ...option.RequestOption) (res *PrismObjectActionListResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/action", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete object
func (r *PrismObjectActionService) Delete(ctx context.Context, actionID string, params PrismObjectActionDeleteParams, opts ...option.RequestOption) (err error) {
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
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/action/%s", params.TeamID, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: small batches complete synchronously and return 200 with
// the final `ImportJob`; large batches start an async job, return 202 with
// `status: processing` and a `Location` header, and can be polled via
// `GET /v2/prism/{teamId}/imports/{jobId}`.
func (r *PrismObjectActionService) BulkNew(ctx context.Context, params PrismObjectActionBulkNewParams, opts ...option.RequestOption) (res *PrismObjectActionBulkNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/action/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Soft-delete up to 100 records in a single call. Same partial-success contract as
// batch/update.
func (r *PrismObjectActionService) BulkDelete(ctx context.Context, params PrismObjectActionBulkDeleteParams, opts ...option.RequestOption) (res *PrismObjectActionBulkDeleteResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/action/batch/delete", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch up to 100 records in a single call. Each item is attempted independently —
// failures don't abort the batch. Inspect `results[].status` per item.
func (r *PrismObjectActionService) BulkUpdate(ctx context.Context, params PrismObjectActionBulkUpdateParams, opts ...option.RequestOption) (res *PrismObjectActionBulkUpdateResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/action/batch/update", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the total number of records of this object type that the caller can see.
// Avoids the page-overshoot anti-pattern — clients no longer need to keep paging
// until `has_more` flips false to discover the total. Currently does not apply
// query filters; for a filtered total, pass `include_total: true` in a POST
// `/query` body.
func (r *PrismObjectActionService) Count(ctx context.Context, params PrismObjectActionCountParams, opts ...option.RequestOption) (res *PrismObjectActionCountResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/action/count", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Duplicate object
func (r *PrismObjectActionService) Duplicate(ctx context.Context, actionID string, params PrismObjectActionDuplicateParams, opts ...option.RequestOption) (res *PrismObjectActionDuplicateResponse, err error) {
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
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/action/%s/duplicate", params.TeamID, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns the single record whose property `{slug}` equals `{value}`. 404 if
// nothing matches; 409 if more than one record matches.
func (r *PrismObjectActionService) Find(ctx context.Context, slug string, value string, params PrismObjectActionFindParams, opts ...option.RequestOption) (res *PrismObjectActionFindResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/action/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectActionService) Get(ctx context.Context, actionID string, params PrismObjectActionGetParams, opts ...option.RequestOption) (res *PrismObjectActionGetResponse, err error) {
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
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/action/%s", params.TeamID, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Query
func (r *PrismObjectActionService) Query(ctx context.Context, params PrismObjectActionQueryParams, opts ...option.RequestOption) (res *PrismObjectActionQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/action/query", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Restore object
func (r *PrismObjectActionService) Restore(ctx context.Context, actionID string, params PrismObjectActionRestoreParams, opts ...option.RequestOption) (res *PrismObjectActionRestoreResponse, err error) {
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
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/action/%s/restore", params.TeamID, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Idempotent create-or-update keyed on `{slug}={value}`. If exactly one record
// matches, it is patched and 200 is returned. If none match, a new record is
// created (with the lookup property set if absent) and 201 is returned. If
// multiple records match, 409 is returned and you should patch by id instead.
func (r *PrismObjectActionService) Upsert(ctx context.Context, slug string, value string, params PrismObjectActionUpsertParams, opts ...option.RequestOption) (res *PrismObjectActionUpsertResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/action/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectActionNewResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}           `json:"default"`
	List    interface{}                      `json:"list"`
	JSON    prismObjectActionNewResponseJSON `json:"-"`
}

// prismObjectActionNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionNewResponse]
type prismObjectActionNewResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionNewResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectActionUpdateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}              `json:"default"`
	List    interface{}                         `json:"list"`
	JSON    prismObjectActionUpdateResponseJSON `json:"-"`
}

// prismObjectActionUpdateResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionUpdateResponse]
type prismObjectActionUpdateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionListResponse struct {
	Data []PrismObjectActionListResponseData `json:"data" api:"required"`
	// Accurate end-of-data signal — false on the last page, never forces clients to
	// overshoot.
	HasMore    bool   `json:"has_more" api:"required"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// Populated only when `?include_total=true` was passed.
	Total int64                             `json:"total" api:"nullable"`
	JSON  prismObjectActionListResponseJSON `json:"-"`
}

// prismObjectActionListResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionListResponse]
type prismObjectActionListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionListResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectActionListResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                `json:"properties"`
	Source     []string                              `json:"source" api:"nullable"`
	JSON       prismObjectActionListResponseDataJSON `json:"-"`
}

// prismObjectActionListResponseDataJSON contains the JSON metadata for the struct
// [PrismObjectActionListResponseData]
type prismObjectActionListResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectActionListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Status snapshot of an import job. Same shape used by the POST /import response
// and by GET /imports/{jobId}.
type PrismObjectActionBulkNewResponse struct {
	// Null for sync imports (results inlined). Set for async imports.
	JobID  string                                 `json:"job_id" api:"required,nullable"`
	Status PrismObjectActionBulkNewResponseStatus `json:"status" api:"required"`
	// Total number of rows in the import.
	Total     int64     `json:"total" api:"required"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Set when status=failed; describes the job-level failure (not per-row).
	Error     PrismObjectActionBulkNewResponseError `json:"error"`
	ExpiresAt time.Time                             `json:"expires_at" format:"date-time"`
	Failed    int64                                 `json:"failed"`
	// Rows that have been attempted (succeeded + failed).
	Processed int64 `json:"processed"`
	// Per-row outcomes. Always present for sync imports; populated for async imports
	// once the job reaches `complete`.
	Results   []PrismObjectActionBulkNewResponseResult `json:"results"`
	Succeeded int64                                    `json:"succeeded"`
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      prismObjectActionBulkNewResponseJSON     `json:"-"`
}

// prismObjectActionBulkNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionBulkNewResponse]
type prismObjectActionBulkNewResponseJSON struct {
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

func (r *PrismObjectActionBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionBulkNewResponseStatus string

const (
	PrismObjectActionBulkNewResponseStatusComplete   PrismObjectActionBulkNewResponseStatus = "complete"
	PrismObjectActionBulkNewResponseStatusProcessing PrismObjectActionBulkNewResponseStatus = "processing"
	PrismObjectActionBulkNewResponseStatusFailed     PrismObjectActionBulkNewResponseStatus = "failed"
)

func (r PrismObjectActionBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectActionBulkNewResponseStatusComplete, PrismObjectActionBulkNewResponseStatusProcessing, PrismObjectActionBulkNewResponseStatusFailed:
		return true
	}
	return false
}

// Set when status=failed; describes the job-level failure (not per-row).
type PrismObjectActionBulkNewResponseError struct {
	Code    string                                    `json:"code"`
	Message string                                    `json:"message"`
	JSON    prismObjectActionBulkNewResponseErrorJSON `json:"-"`
}

// prismObjectActionBulkNewResponseErrorJSON contains the JSON metadata for the
// struct [PrismObjectActionBulkNewResponseError]
type prismObjectActionBulkNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionBulkNewResponseResult struct {
	ID      string                                       `json:"id" api:"nullable" format:"uuid"`
	Created bool                                         `json:"created"`
	Error   PrismObjectActionBulkNewResponseResultsError `json:"error"`
	// True if the row matched an existing record via the dedupe key.
	Existing bool `json:"existing"`
	// Zero-based position of this row in the request.
	InputIndex int64 `json:"input_index"`
	// True if a matching record was updated.
	Updated bool                                       `json:"updated"`
	JSON    prismObjectActionBulkNewResponseResultJSON `json:"-"`
}

// prismObjectActionBulkNewResponseResultJSON contains the JSON metadata for the
// struct [PrismObjectActionBulkNewResponseResult]
type prismObjectActionBulkNewResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	InputIndex  apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionBulkNewResponseResultsError struct {
	Code    string                                           `json:"code"`
	Message string                                           `json:"message"`
	JSON    prismObjectActionBulkNewResponseResultsErrorJSON `json:"-"`
}

// prismObjectActionBulkNewResponseResultsErrorJSON contains the JSON metadata for
// the struct [PrismObjectActionBulkNewResponseResultsError]
type prismObjectActionBulkNewResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkNewResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkNewResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectActionBulkDeleteResponse struct {
	Results []PrismObjectActionBulkDeleteResponseResult `json:"results" api:"required"`
	Summary PrismObjectActionBulkDeleteResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectActionBulkDeleteResponseJSON     `json:"-"`
}

// prismObjectActionBulkDeleteResponseJSON contains the JSON metadata for the
// struct [PrismObjectActionBulkDeleteResponse]
type prismObjectActionBulkDeleteResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionBulkDeleteResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                           `json:"id" api:"required,nullable"`
	Status PrismObjectActionBulkDeleteResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectActionBulkDeleteResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectActionBulkDeleteResponseResultsRecord `json:"record"`
	JSON   prismObjectActionBulkDeleteResponseResultJSON    `json:"-"`
}

// prismObjectActionBulkDeleteResponseResultJSON contains the JSON metadata for the
// struct [PrismObjectActionBulkDeleteResponseResult]
type prismObjectActionBulkDeleteResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionBulkDeleteResponseResultsStatus string

const (
	PrismObjectActionBulkDeleteResponseResultsStatusOk    PrismObjectActionBulkDeleteResponseResultsStatus = "ok"
	PrismObjectActionBulkDeleteResponseResultsStatusError PrismObjectActionBulkDeleteResponseResultsStatus = "error"
)

func (r PrismObjectActionBulkDeleteResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectActionBulkDeleteResponseResultsStatusOk, PrismObjectActionBulkDeleteResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectActionBulkDeleteResponseResultsError struct {
	Code    string                                              `json:"code"`
	Message string                                              `json:"message"`
	JSON    prismObjectActionBulkDeleteResponseResultsErrorJSON `json:"-"`
}

// prismObjectActionBulkDeleteResponseResultsErrorJSON contains the JSON metadata
// for the struct [PrismObjectActionBulkDeleteResponseResultsError]
type prismObjectActionBulkDeleteResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkDeleteResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkDeleteResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectActionBulkDeleteResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                               `json:"default"`
	List    interface{}                                          `json:"list"`
	JSON    prismObjectActionBulkDeleteResponseResultsRecordJSON `json:"-"`
}

// prismObjectActionBulkDeleteResponseResultsRecordJSON contains the JSON metadata
// for the struct [PrismObjectActionBulkDeleteResponseResultsRecord]
type prismObjectActionBulkDeleteResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkDeleteResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkDeleteResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionBulkDeleteResponseSummary struct {
	Failed    int64                                          `json:"failed" api:"required"`
	Succeeded int64                                          `json:"succeeded" api:"required"`
	Total     int64                                          `json:"total" api:"required"`
	JSON      prismObjectActionBulkDeleteResponseSummaryJSON `json:"-"`
}

// prismObjectActionBulkDeleteResponseSummaryJSON contains the JSON metadata for
// the struct [PrismObjectActionBulkDeleteResponseSummary]
type prismObjectActionBulkDeleteResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkDeleteResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkDeleteResponseSummaryJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectActionBulkUpdateResponse struct {
	Results []PrismObjectActionBulkUpdateResponseResult `json:"results" api:"required"`
	Summary PrismObjectActionBulkUpdateResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectActionBulkUpdateResponseJSON     `json:"-"`
}

// prismObjectActionBulkUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectActionBulkUpdateResponse]
type prismObjectActionBulkUpdateResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionBulkUpdateResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                           `json:"id" api:"required,nullable"`
	Status PrismObjectActionBulkUpdateResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectActionBulkUpdateResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectActionBulkUpdateResponseResultsRecord `json:"record"`
	JSON   prismObjectActionBulkUpdateResponseResultJSON    `json:"-"`
}

// prismObjectActionBulkUpdateResponseResultJSON contains the JSON metadata for the
// struct [PrismObjectActionBulkUpdateResponseResult]
type prismObjectActionBulkUpdateResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionBulkUpdateResponseResultsStatus string

const (
	PrismObjectActionBulkUpdateResponseResultsStatusOk    PrismObjectActionBulkUpdateResponseResultsStatus = "ok"
	PrismObjectActionBulkUpdateResponseResultsStatusError PrismObjectActionBulkUpdateResponseResultsStatus = "error"
)

func (r PrismObjectActionBulkUpdateResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectActionBulkUpdateResponseResultsStatusOk, PrismObjectActionBulkUpdateResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectActionBulkUpdateResponseResultsError struct {
	Code    string                                              `json:"code"`
	Message string                                              `json:"message"`
	JSON    prismObjectActionBulkUpdateResponseResultsErrorJSON `json:"-"`
}

// prismObjectActionBulkUpdateResponseResultsErrorJSON contains the JSON metadata
// for the struct [PrismObjectActionBulkUpdateResponseResultsError]
type prismObjectActionBulkUpdateResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkUpdateResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkUpdateResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectActionBulkUpdateResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                               `json:"default"`
	List    interface{}                                          `json:"list"`
	JSON    prismObjectActionBulkUpdateResponseResultsRecordJSON `json:"-"`
}

// prismObjectActionBulkUpdateResponseResultsRecordJSON contains the JSON metadata
// for the struct [PrismObjectActionBulkUpdateResponseResultsRecord]
type prismObjectActionBulkUpdateResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkUpdateResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkUpdateResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionBulkUpdateResponseSummary struct {
	Failed    int64                                          `json:"failed" api:"required"`
	Succeeded int64                                          `json:"succeeded" api:"required"`
	Total     int64                                          `json:"total" api:"required"`
	JSON      prismObjectActionBulkUpdateResponseSummaryJSON `json:"-"`
}

// prismObjectActionBulkUpdateResponseSummaryJSON contains the JSON metadata for
// the struct [PrismObjectActionBulkUpdateResponseSummary]
type prismObjectActionBulkUpdateResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionBulkUpdateResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionBulkUpdateResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionCountResponse struct {
	// Number of records matching the access scope.
	Total int64                              `json:"total" api:"required"`
	JSON  prismObjectActionCountResponseJSON `json:"-"`
}

// prismObjectActionCountResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionCountResponse]
type prismObjectActionCountResponseJSON struct {
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionCountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionCountResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectActionDuplicateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                 `json:"default"`
	List    interface{}                            `json:"list"`
	JSON    prismObjectActionDuplicateResponseJSON `json:"-"`
}

// prismObjectActionDuplicateResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionDuplicateResponse]
type prismObjectActionDuplicateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionDuplicateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionDuplicateResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectActionFindResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}            `json:"default"`
	List    interface{}                       `json:"list"`
	JSON    prismObjectActionFindResponseJSON `json:"-"`
}

// prismObjectActionFindResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionFindResponse]
type prismObjectActionFindResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionFindResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionFindResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectActionGetResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}           `json:"default"`
	List    interface{}                      `json:"list"`
	JSON    prismObjectActionGetResponseJSON `json:"-"`
}

// prismObjectActionGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionGetResponse]
type prismObjectActionGetResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionQueryResponse struct {
	Data []PrismObjectActionQueryResponseData `json:"data" api:"required"`
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
	Total int64                              `json:"total" api:"nullable"`
	JSON  prismObjectActionQueryResponseJSON `json:"-"`
}

// prismObjectActionQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionQueryResponse]
type prismObjectActionQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectActionQueryResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                 `json:"properties"`
	Source     []string                               `json:"source" api:"nullable"`
	JSON       prismObjectActionQueryResponseDataJSON `json:"-"`
}

// prismObjectActionQueryResponseDataJSON contains the JSON metadata for the struct
// [PrismObjectActionQueryResponseData]
type prismObjectActionQueryResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectActionQueryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionQueryResponseDataJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectActionRestoreResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}               `json:"default"`
	List    interface{}                          `json:"list"`
	JSON    prismObjectActionRestoreResponseJSON `json:"-"`
}

// prismObjectActionRestoreResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionRestoreResponse]
type prismObjectActionRestoreResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionRestoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionRestoreResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectActionUpsertResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}              `json:"default"`
	List    interface{}                         `json:"list"`
	JSON    prismObjectActionUpsertResponseJSON `json:"-"`
}

// prismObjectActionUpsertResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionUpsertResponse]
type prismObjectActionUpsertResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionUpsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionUpsertResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
}

func (r PrismObjectActionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectActionUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
	IfMatch               param.Field[string]        `header:"If-Match"`
}

func (r PrismObjectActionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectActionListParams struct {
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

// URLQuery serializes [PrismObjectActionListParams]'s query parameters as
// `url.Values`.
func (r PrismObjectActionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectActionDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IfMatch param.Field[string] `header:"If-Match"`
}

type PrismObjectActionBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects        param.Field[[]PrismObjectPropertiesParam]          `json:"objects" api:"required"`
	Options        param.Field[PrismObjectActionBulkNewParamsOptions] `json:"options"`
	IdempotencyKey param.Field[string]                                `header:"Idempotency-Key"`
}

func (r PrismObjectActionBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectActionBulkNewParamsOptions struct {
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
	DedupeBy param.Field[PrismObjectActionBulkNewParamsOptionsDedupeByUnion] `json:"dedupe_by"`
	// App/CRM ID for context (optional)
	ListID param.Field[string] `json:"list_id" format:"uuid"`
	// Require app_stage for every row in the selected list. app_stage is a reserved
	// list-scoped alias for native status.
	RequireListStage param.Field[bool] `json:"require_list_stage"`
	// Patch a deduplicated record with the supplied properties instead of skipping it.
	UpdateExisting param.Field[bool] `json:"update_existing"`
}

func (r PrismObjectActionBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Property slug to deduplicate on. A single-element array is also accepted;
// compound (multi-slug) dedupe is not supported yet and is rejected with guidance.
//
// Satisfied by [shared.UnionString],
// [PrismObjectActionBulkNewParamsOptionsDedupeByArray].
type PrismObjectActionBulkNewParamsOptionsDedupeByUnion interface {
	ImplementsPrismObjectActionBulkNewParamsOptionsDedupeByUnion()
}

type PrismObjectActionBulkNewParamsOptionsDedupeByArray []string

func (r PrismObjectActionBulkNewParamsOptionsDedupeByArray) ImplementsPrismObjectActionBulkNewParamsOptionsDedupeByUnion() {
}

type PrismObjectActionBulkDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]   `path:"teamId" api:"required" format:"uuid"`
	IDs            param.Field[[]string] `json:"ids" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string]   `header:"Idempotency-Key"`
}

func (r PrismObjectActionBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectActionBulkUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]                                  `path:"teamId" api:"required" format:"uuid"`
	Items          param.Field[[]PrismObjectActionBulkUpdateParamsItem] `json:"items" api:"required"`
	IdempotencyKey param.Field[string]                                  `header:"Idempotency-Key"`
}

func (r PrismObjectActionBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object with `id` plus the same property body shape as PATCH
// (`default`/`list`/`extended`).
type PrismObjectActionBulkUpdateParamsItem struct {
	ID          param.Field[string]    `json:"id" api:"required" format:"uuid"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r PrismObjectActionBulkUpdateParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectActionCountParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the count to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectActionCountParams]'s query parameters as
// `url.Values`.
func (r PrismObjectActionCountParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectActionDuplicateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectActionFindParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the lookup to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectActionFindParams]'s query parameters as
// `url.Values`.
func (r PrismObjectActionFindParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectActionGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Comma-separated property slugs to return. Use dot notation for relationships.
	// `id` is always returned at the top level. Defaults to all properties.
	Select param.Field[string] `query:"select"`
}

// URLQuery serializes [PrismObjectActionGetParams]'s query parameters as
// `url.Values`.
func (r PrismObjectActionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectActionQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string]                              `path:"teamId" api:"required" format:"uuid"`
	Query  param.Field[PrismObjectActionQueryParamsQuery]   `json:"query" api:"required"`
	ID     param.Field[PrismObjectActionQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes  param.Field[[]string]                            `json:"boxes"`
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

func (r PrismObjectActionQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectActionQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name). `id` is always returned at the top level of each
	// row and does not need to be selected.
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectActionQueryParamsQueryCombinator] `json:"combinator"`
	// Opaque cursor from a previous response's `next_cursor`. Pass it back unchanged
	// to fetch the next page. When set, `page` and `limit` are derived from the cursor
	// and any explicit values are ignored.
	Cursor param.Field[string] `json:"cursor"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectActionQueryParamsQueryFilterUnion] `json:"filter"`
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
	Sort param.Field[[]map[string]PrismObjectActionQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectActionQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectActionQueryParamsQueryCombinator string

const (
	PrismObjectActionQueryParamsQueryCombinatorAnd PrismObjectActionQueryParamsQueryCombinator = "AND"
	PrismObjectActionQueryParamsQueryCombinatorOr  PrismObjectActionQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectActionQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectActionQueryParamsQueryCombinatorAnd, PrismObjectActionQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

type PrismObjectActionQueryParamsQueryFilter struct {
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

func (r PrismObjectActionQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilter) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

// Satisfied by [PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEq],
// [PrismObjectActionQueryParamsQueryFilterPrismQueryFilterNe],
// [PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLt],
// [PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGt],
// [PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLte],
// [PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGte],
// [PrismObjectActionQueryParamsQueryFilterContains],
// [PrismObjectActionQueryParamsQueryFilterBeginsWith],
// [PrismObjectActionQueryParamsQueryFilterEndsWith],
// [PrismObjectActionQueryParamsQueryFilterNotContains],
// [PrismObjectActionQueryParamsQueryFilterExists],
// [PrismObjectActionQueryParamsQueryFilterNotExists],
// [PrismObjectActionQueryParamsQueryFilterIsNull],
// [PrismObjectActionQueryParamsQueryFilterIsNotNull],
// [PrismObjectActionQueryParamsQueryFilterBetween],
// [PrismObjectActionQueryParamsQueryFilterIn],
// [PrismObjectActionQueryParamsQueryFilterNotIn],
// [PrismObjectActionQueryParamsQueryFilter].
type PrismObjectActionQueryParamsQueryFilterUnion interface {
	implementsPrismObjectActionQueryParamsQueryFilterUnion()
}

type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEq struct {
	Equals param.Field[PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion] `json:"=" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEq) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEq) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion interface {
	ImplementsPrismObjectActionQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion()
}

type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterNe struct {
	NotEquals param.Field[PrismObjectActionQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion] `json:"!=" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterNe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterNe) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion interface {
	ImplementsPrismObjectActionQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion()
}

type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLt struct {
	Less param.Field[string] `json:"<" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLt) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGt struct {
	Greater param.Field[string] `json:">" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGt) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLte struct {
	LessOrEquals param.Field[string] `json:"<=" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterLte) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGte struct {
	GreaterOrEquals param.Field[string] `json:">=" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterPrismQueryFilterGte) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterContains struct {
	Contains param.Field[PrismObjectActionQueryParamsQueryFilterContainsContainsUnion] `json:"contains" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterContains) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectActionQueryParamsQueryFilterContainsContainsArray].
type PrismObjectActionQueryParamsQueryFilterContainsContainsUnion interface {
	ImplementsPrismObjectActionQueryParamsQueryFilterContainsContainsUnion()
}

type PrismObjectActionQueryParamsQueryFilterContainsContainsArray []string

func (r PrismObjectActionQueryParamsQueryFilterContainsContainsArray) ImplementsPrismObjectActionQueryParamsQueryFilterContainsContainsUnion() {
}

type PrismObjectActionQueryParamsQueryFilterBeginsWith struct {
	BeginsWith param.Field[string] `json:"begins_with" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterBeginsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterBeginsWith) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterEndsWith struct {
	EndsWith param.Field[string] `json:"ends_with" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterEndsWith) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterEndsWith) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterNotContains struct {
	NotContains param.Field[string] `json:"not_contains" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterNotContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterNotContains) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterExists struct {
	Exists param.Field[bool] `json:"exists" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterExists) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterNotExists struct {
	NotExists param.Field[bool] `json:"not_exists" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterNotExists) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterNotExists) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterIsNull struct {
	IsNull param.Field[PrismObjectActionQueryParamsQueryFilterIsNullIsNullUnion] `json:"is_null" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterIsNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterIsNull) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectActionQueryParamsQueryFilterIsNullIsNullArray].
type PrismObjectActionQueryParamsQueryFilterIsNullIsNullUnion interface {
	ImplementsPrismObjectActionQueryParamsQueryFilterIsNullIsNullUnion()
}

type PrismObjectActionQueryParamsQueryFilterIsNullIsNullArray []string

func (r PrismObjectActionQueryParamsQueryFilterIsNullIsNullArray) ImplementsPrismObjectActionQueryParamsQueryFilterIsNullIsNullUnion() {
}

type PrismObjectActionQueryParamsQueryFilterIsNotNull struct {
	IsNotNull param.Field[PrismObjectActionQueryParamsQueryFilterIsNotNullIsNotNullUnion] `json:"is_not_null" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterIsNotNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterIsNotNull) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectActionQueryParamsQueryFilterIsNotNullIsNotNullArray].
type PrismObjectActionQueryParamsQueryFilterIsNotNullIsNotNullUnion interface {
	ImplementsPrismObjectActionQueryParamsQueryFilterIsNotNullIsNotNullUnion()
}

type PrismObjectActionQueryParamsQueryFilterIsNotNullIsNotNullArray []string

func (r PrismObjectActionQueryParamsQueryFilterIsNotNullIsNotNullArray) ImplementsPrismObjectActionQueryParamsQueryFilterIsNotNullIsNotNullUnion() {
}

type PrismObjectActionQueryParamsQueryFilterBetween struct {
	Between param.Field[PrismObjectActionQueryParamsQueryFilterBetweenBetweenUnion] `json:"between" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterBetween) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterBetween) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectActionQueryParamsQueryFilterBetweenBetweenArray].
type PrismObjectActionQueryParamsQueryFilterBetweenBetweenUnion interface {
	ImplementsPrismObjectActionQueryParamsQueryFilterBetweenBetweenUnion()
}

type PrismObjectActionQueryParamsQueryFilterBetweenBetweenArray []string

func (r PrismObjectActionQueryParamsQueryFilterBetweenBetweenArray) ImplementsPrismObjectActionQueryParamsQueryFilterBetweenBetweenUnion() {
}

type PrismObjectActionQueryParamsQueryFilterIn struct {
	In param.Field[[]string] `json:"in" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterIn) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQueryFilterNotIn struct {
	NotIn param.Field[[]string] `json:"not_in" api:"required"`
}

func (r PrismObjectActionQueryParamsQueryFilterNotIn) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectActionQueryParamsQueryFilterNotIn) implementsPrismObjectActionQueryParamsQueryFilterUnion() {
}

type PrismObjectActionQueryParamsQuerySort string

const (
	PrismObjectActionQueryParamsQuerySortAsc  PrismObjectActionQueryParamsQuerySort = "asc"
	PrismObjectActionQueryParamsQuerySortDesc PrismObjectActionQueryParamsQuerySort = "desc"
)

func (r PrismObjectActionQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectActionQueryParamsQuerySortAsc, PrismObjectActionQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectActionQueryParamsIDArray].
type PrismObjectActionQueryParamsIDUnion interface {
	ImplementsPrismObjectActionQueryParamsIDUnion()
}

type PrismObjectActionQueryParamsIDArray []string

func (r PrismObjectActionQueryParamsIDArray) ImplementsPrismObjectActionQueryParamsIDUnion() {}

type PrismObjectActionRestoreParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectActionUpsertParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	// Scope the upsert to a specific list/app. Required to match or write list-scoped
	// properties, including `app_stage`.
	ListID         param.Field[string] `query:"list_id" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

func (r PrismObjectActionUpsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

// URLQuery serializes [PrismObjectActionUpsertParams]'s query parameters as
// `url.Values`.
func (r PrismObjectActionUpsertParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
