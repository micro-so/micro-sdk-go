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
	path := fmt.Sprintf("v2/prism/%s/contact", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch object
func (r *PrismObjectContactService) Update(ctx context.Context, contactID string, params PrismObjectContactUpdateParams, opts ...option.RequestOption) (res *PrismObjectContactUpdateResponse, err error) {
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
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/contact/%s", params.TeamID, contactID)
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
func (r *PrismObjectContactService) List(ctx context.Context, params PrismObjectContactListParams, opts ...option.RequestOption) (res *PrismObjectContactListResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete object
func (r *PrismObjectContactService) Delete(ctx context.Context, contactID string, params PrismObjectContactDeleteParams, opts ...option.RequestOption) (err error) {
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
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/contact/%s", params.TeamID, contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: small batches complete synchronously and return 200 with
// the final `ImportJob`; large batches start an async job, return 202 with
// `status: processing` and a `Location` header, and can be polled via
// `GET /v2/prism/{teamId}/imports/{jobId}`.
func (r *PrismObjectContactService) BulkNew(ctx context.Context, params PrismObjectContactBulkNewParams, opts ...option.RequestOption) (res *PrismObjectContactBulkNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/contact/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Soft-delete up to 100 records in a single call. Same partial-success contract as
// batch/update.
func (r *PrismObjectContactService) BulkDelete(ctx context.Context, params PrismObjectContactBulkDeleteParams, opts ...option.RequestOption) (res *PrismObjectContactBulkDeleteResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/contact/batch/delete", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch up to 100 records in a single call. Each item is attempted independently —
// failures don't abort the batch. Inspect `results[].status` per item.
func (r *PrismObjectContactService) BulkUpdate(ctx context.Context, params PrismObjectContactBulkUpdateParams, opts ...option.RequestOption) (res *PrismObjectContactBulkUpdateResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/contact/batch/update", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the total number of records of this object type that the caller can see.
// Avoids the page-overshoot anti-pattern — clients no longer need to keep paging
// until `has_more` flips false to discover the total. Currently does not apply
// query filters; for a filtered total, pass `include_total: true` in a POST
// `/query` body.
func (r *PrismObjectContactService) Count(ctx context.Context, params PrismObjectContactCountParams, opts ...option.RequestOption) (res *PrismObjectContactCountResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/contact/count", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Duplicate object
func (r *PrismObjectContactService) Duplicate(ctx context.Context, contactID string, params PrismObjectContactDuplicateParams, opts ...option.RequestOption) (res *PrismObjectContactDuplicateResponse, err error) {
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
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/contact/%s/duplicate", params.TeamID, contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns the single record whose property `{slug}` equals `{value}`. 404 if
// nothing matches; 409 if more than one record matches.
func (r *PrismObjectContactService) Find(ctx context.Context, slug string, value string, params PrismObjectContactFindParams, opts ...option.RequestOption) (res *PrismObjectContactFindResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/contact/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get object
func (r *PrismObjectContactService) Get(ctx context.Context, contactID string, params PrismObjectContactGetParams, opts ...option.RequestOption) (res *PrismObjectContactGetResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
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
	path := fmt.Sprintf("v2/prism/%s/contact/query", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Restore object
func (r *PrismObjectContactService) Restore(ctx context.Context, contactID string, params PrismObjectContactRestoreParams, opts ...option.RequestOption) (res *PrismObjectContactRestoreResponse, err error) {
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
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/contact/%s/restore", params.TeamID, contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Idempotent create-or-update keyed on `{slug}={value}`. If exactly one record
// matches, it is patched and 200 is returned. If none match, a new record is
// created (with the lookup property set if absent) and 201 is returned. If
// multiple records match, 409 is returned and you should patch by id instead.
func (r *PrismObjectContactService) Upsert(ctx context.Context, slug string, value string, params PrismObjectContactUpsertParams, opts ...option.RequestOption) (res *PrismObjectContactUpsertResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/contact/by/%s/%s", params.TeamID, slug, value)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
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

type PrismObjectContactListResponse struct {
	Data []PrismObjectContactListResponseData `json:"data" api:"required"`
	// Accurate end-of-data signal — false on the last page, never forces clients to
	// overshoot.
	HasMore    bool   `json:"has_more" api:"required"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// Populated only when `?include_total=true` was passed.
	Total int64                              `json:"total" api:"nullable"`
	JSON  prismObjectContactListResponseJSON `json:"-"`
}

// prismObjectContactListResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactListResponse]
type prismObjectContactListResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactListResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectContactListResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                 `json:"properties"`
	Source     []string                               `json:"source" api:"nullable"`
	JSON       prismObjectContactListResponseDataJSON `json:"-"`
}

// prismObjectContactListResponseDataJSON contains the JSON metadata for the struct
// [PrismObjectContactListResponseData]
type prismObjectContactListResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrismObjectContactListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Status snapshot of an import job. Same shape used by the POST /import response
// and by GET /imports/{jobId}.
type PrismObjectContactBulkNewResponse struct {
	// Null for sync imports (results inlined). Set for async imports.
	JobID  string                                  `json:"job_id" api:"required,nullable"`
	Status PrismObjectContactBulkNewResponseStatus `json:"status" api:"required"`
	// Total number of rows in the import.
	Total     int64     `json:"total" api:"required"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Set when status=failed; describes the job-level failure (not per-row).
	Error     PrismObjectContactBulkNewResponseError `json:"error"`
	ExpiresAt time.Time                              `json:"expires_at" format:"date-time"`
	Failed    int64                                  `json:"failed"`
	// Rows that have been attempted (succeeded + failed).
	Processed int64 `json:"processed"`
	// Per-row outcomes. Always present for sync imports; populated for async imports
	// once the job reaches `complete`.
	Results   []PrismObjectContactBulkNewResponseResult `json:"results"`
	Succeeded int64                                     `json:"succeeded"`
	UpdatedAt time.Time                                 `json:"updated_at" format:"date-time"`
	JSON      prismObjectContactBulkNewResponseJSON     `json:"-"`
}

// prismObjectContactBulkNewResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactBulkNewResponse]
type prismObjectContactBulkNewResponseJSON struct {
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

func (r *PrismObjectContactBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkNewResponseStatus string

const (
	PrismObjectContactBulkNewResponseStatusComplete   PrismObjectContactBulkNewResponseStatus = "complete"
	PrismObjectContactBulkNewResponseStatusProcessing PrismObjectContactBulkNewResponseStatus = "processing"
	PrismObjectContactBulkNewResponseStatusFailed     PrismObjectContactBulkNewResponseStatus = "failed"
)

func (r PrismObjectContactBulkNewResponseStatus) IsKnown() bool {
	switch r {
	case PrismObjectContactBulkNewResponseStatusComplete, PrismObjectContactBulkNewResponseStatusProcessing, PrismObjectContactBulkNewResponseStatusFailed:
		return true
	}
	return false
}

// Set when status=failed; describes the job-level failure (not per-row).
type PrismObjectContactBulkNewResponseError struct {
	Code    string                                     `json:"code"`
	Message string                                     `json:"message"`
	JSON    prismObjectContactBulkNewResponseErrorJSON `json:"-"`
}

// prismObjectContactBulkNewResponseErrorJSON contains the JSON metadata for the
// struct [PrismObjectContactBulkNewResponseError]
type prismObjectContactBulkNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkNewResponseResult struct {
	ID      string                                        `json:"id" api:"nullable" format:"uuid"`
	Created bool                                          `json:"created"`
	Error   PrismObjectContactBulkNewResponseResultsError `json:"error"`
	// True if the row matched an existing record via the dedupe key.
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

type PrismObjectContactBulkNewResponseResultsError struct {
	Code    string                                            `json:"code"`
	Message string                                            `json:"message"`
	JSON    prismObjectContactBulkNewResponseResultsErrorJSON `json:"-"`
}

// prismObjectContactBulkNewResponseResultsErrorJSON contains the JSON metadata for
// the struct [PrismObjectContactBulkNewResponseResultsError]
type prismObjectContactBulkNewResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkNewResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkNewResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectContactBulkDeleteResponse struct {
	Results []PrismObjectContactBulkDeleteResponseResult `json:"results" api:"required"`
	Summary PrismObjectContactBulkDeleteResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectContactBulkDeleteResponseJSON     `json:"-"`
}

// prismObjectContactBulkDeleteResponseJSON contains the JSON metadata for the
// struct [PrismObjectContactBulkDeleteResponse]
type prismObjectContactBulkDeleteResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkDeleteResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                            `json:"id" api:"required,nullable"`
	Status PrismObjectContactBulkDeleteResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectContactBulkDeleteResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectContactBulkDeleteResponseResultsRecord `json:"record"`
	JSON   prismObjectContactBulkDeleteResponseResultJSON    `json:"-"`
}

// prismObjectContactBulkDeleteResponseResultJSON contains the JSON metadata for
// the struct [PrismObjectContactBulkDeleteResponseResult]
type prismObjectContactBulkDeleteResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkDeleteResponseResultsStatus string

const (
	PrismObjectContactBulkDeleteResponseResultsStatusOk    PrismObjectContactBulkDeleteResponseResultsStatus = "ok"
	PrismObjectContactBulkDeleteResponseResultsStatusError PrismObjectContactBulkDeleteResponseResultsStatus = "error"
)

func (r PrismObjectContactBulkDeleteResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectContactBulkDeleteResponseResultsStatusOk, PrismObjectContactBulkDeleteResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectContactBulkDeleteResponseResultsError struct {
	Code    string                                               `json:"code"`
	Message string                                               `json:"message"`
	JSON    prismObjectContactBulkDeleteResponseResultsErrorJSON `json:"-"`
}

// prismObjectContactBulkDeleteResponseResultsErrorJSON contains the JSON metadata
// for the struct [PrismObjectContactBulkDeleteResponseResultsError]
type prismObjectContactBulkDeleteResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkDeleteResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkDeleteResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectContactBulkDeleteResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                                `json:"default"`
	List    interface{}                                           `json:"list"`
	JSON    prismObjectContactBulkDeleteResponseResultsRecordJSON `json:"-"`
}

// prismObjectContactBulkDeleteResponseResultsRecordJSON contains the JSON metadata
// for the struct [PrismObjectContactBulkDeleteResponseResultsRecord]
type prismObjectContactBulkDeleteResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkDeleteResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkDeleteResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkDeleteResponseSummary struct {
	Failed    int64                                           `json:"failed" api:"required"`
	Succeeded int64                                           `json:"succeeded" api:"required"`
	Total     int64                                           `json:"total" api:"required"`
	JSON      prismObjectContactBulkDeleteResponseSummaryJSON `json:"-"`
}

// prismObjectContactBulkDeleteResponseSummaryJSON contains the JSON metadata for
// the struct [PrismObjectContactBulkDeleteResponseSummary]
type prismObjectContactBulkDeleteResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkDeleteResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkDeleteResponseSummaryJSON) RawJSON() string {
	return r.raw
}

// Partial-success bulk operation result. Inspect `results[].status` per item; the
// operation as a whole returns 200 even if some items failed.
type PrismObjectContactBulkUpdateResponse struct {
	Results []PrismObjectContactBulkUpdateResponseResult `json:"results" api:"required"`
	Summary PrismObjectContactBulkUpdateResponseSummary  `json:"summary" api:"required"`
	JSON    prismObjectContactBulkUpdateResponseJSON     `json:"-"`
}

// prismObjectContactBulkUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectContactBulkUpdateResponse]
type prismObjectContactBulkUpdateResponseJSON struct {
	Results     apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkUpdateResponseResult struct {
	// Item ID, or null if the input was unparseable.
	ID     string                                            `json:"id" api:"required,nullable"`
	Status PrismObjectContactBulkUpdateResponseResultsStatus `json:"status" api:"required"`
	Error  PrismObjectContactBulkUpdateResponseResultsError  `json:"error"`
	// Object returned by reads (get/create/patch/restore). id is always present.
	Record PrismObjectContactBulkUpdateResponseResultsRecord `json:"record"`
	JSON   prismObjectContactBulkUpdateResponseResultJSON    `json:"-"`
}

// prismObjectContactBulkUpdateResponseResultJSON contains the JSON metadata for
// the struct [PrismObjectContactBulkUpdateResponseResult]
type prismObjectContactBulkUpdateResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Error       apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkUpdateResponseResultsStatus string

const (
	PrismObjectContactBulkUpdateResponseResultsStatusOk    PrismObjectContactBulkUpdateResponseResultsStatus = "ok"
	PrismObjectContactBulkUpdateResponseResultsStatusError PrismObjectContactBulkUpdateResponseResultsStatus = "error"
)

func (r PrismObjectContactBulkUpdateResponseResultsStatus) IsKnown() bool {
	switch r {
	case PrismObjectContactBulkUpdateResponseResultsStatusOk, PrismObjectContactBulkUpdateResponseResultsStatusError:
		return true
	}
	return false
}

type PrismObjectContactBulkUpdateResponseResultsError struct {
	Code    string                                               `json:"code"`
	Message string                                               `json:"message"`
	JSON    prismObjectContactBulkUpdateResponseResultsErrorJSON `json:"-"`
}

// prismObjectContactBulkUpdateResponseResultsErrorJSON contains the JSON metadata
// for the struct [PrismObjectContactBulkUpdateResponseResultsError]
type prismObjectContactBulkUpdateResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkUpdateResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkUpdateResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectContactBulkUpdateResponseResultsRecord struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                                `json:"default"`
	List    interface{}                                           `json:"list"`
	JSON    prismObjectContactBulkUpdateResponseResultsRecordJSON `json:"-"`
}

// prismObjectContactBulkUpdateResponseResultsRecordJSON contains the JSON metadata
// for the struct [PrismObjectContactBulkUpdateResponseResultsRecord]
type prismObjectContactBulkUpdateResponseResultsRecordJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkUpdateResponseResultsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkUpdateResponseResultsRecordJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactBulkUpdateResponseSummary struct {
	Failed    int64                                           `json:"failed" api:"required"`
	Succeeded int64                                           `json:"succeeded" api:"required"`
	Total     int64                                           `json:"total" api:"required"`
	JSON      prismObjectContactBulkUpdateResponseSummaryJSON `json:"-"`
}

// prismObjectContactBulkUpdateResponseSummaryJSON contains the JSON metadata for
// the struct [PrismObjectContactBulkUpdateResponseSummary]
type prismObjectContactBulkUpdateResponseSummaryJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactBulkUpdateResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactBulkUpdateResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactCountResponse struct {
	// Number of records matching the access scope.
	Total int64                               `json:"total" api:"required"`
	JSON  prismObjectContactCountResponseJSON `json:"-"`
}

// prismObjectContactCountResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactCountResponse]
type prismObjectContactCountResponseJSON struct {
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactCountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactCountResponseJSON) RawJSON() string {
	return r.raw
}

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectContactDuplicateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}                  `json:"default"`
	List    interface{}                             `json:"list"`
	JSON    prismObjectContactDuplicateResponseJSON `json:"-"`
}

// prismObjectContactDuplicateResponseJSON contains the JSON metadata for the
// struct [PrismObjectContactDuplicateResponse]
type prismObjectContactDuplicateResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
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
type PrismObjectContactFindResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}             `json:"default"`
	List    interface{}                        `json:"list"`
	JSON    prismObjectContactFindResponseJSON `json:"-"`
}

// prismObjectContactFindResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactFindResponse]
type prismObjectContactFindResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactFindResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactFindResponseJSON) RawJSON() string {
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
	Total int64                               `json:"total" api:"nullable"`
	JSON  prismObjectContactQueryResponseJSON `json:"-"`
}

// prismObjectContactQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactQueryResponse]
type prismObjectContactQueryResponseJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Row returned by the query endpoint. `id` is always present at the top level.
// Selected property values are returned under `properties`, keyed by property
// slug. Reference-typed values are returned as nested `{ id, properties }`
// objects.
type PrismObjectContactQueryResponseData struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	IsUserObject bool   `json:"is_user_object"`
	// Selected property values keyed by property slug. For select/multiselect
	// properties, option slugs are returned. For reference properties, values are
	// nested `{ id, properties }` objects.
	Properties map[string]interface{}                  `json:"properties"`
	Source     []string                                `json:"source" api:"nullable"`
	JSON       prismObjectContactQueryResponseDataJSON `json:"-"`
}

// prismObjectContactQueryResponseDataJSON contains the JSON metadata for the
// struct [PrismObjectContactQueryResponseData]
type prismObjectContactQueryResponseDataJSON struct {
	ID           apijson.Field
	IsUserObject apijson.Field
	Properties   apijson.Field
	Source       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
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

// Object returned by reads (get/create/patch/restore). id is always present.
type PrismObjectContactUpsertResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Properties keyed by property slug.
	Default map[string]interface{}               `json:"default"`
	List    interface{}                          `json:"list"`
	JSON    prismObjectContactUpsertResponseJSON `json:"-"`
}

// prismObjectContactUpsertResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactUpsertResponse]
type prismObjectContactUpsertResponseJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactUpsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactUpsertResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
}

func (r PrismObjectContactNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectContactUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
	IfMatch               param.Field[string]        `header:"If-Match"`
}

func (r PrismObjectContactUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismObjectContactListParams struct {
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

// URLQuery serializes [PrismObjectContactListParams]'s query parameters as
// `url.Values`.
func (r PrismObjectContactListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectContactDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IfMatch param.Field[string] `header:"If-Match"`
}

type PrismObjectContactBulkNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects        param.Field[[]PrismObjectPropertiesParam]           `json:"objects" api:"required"`
	Options        param.Field[PrismObjectContactBulkNewParamsOptions] `json:"options"`
	IdempotencyKey param.Field[string]                                 `header:"Idempotency-Key"`
}

func (r PrismObjectContactBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectContactBulkNewParamsOptions struct {
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

func (r PrismObjectContactBulkNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectContactBulkDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]   `path:"teamId" api:"required" format:"uuid"`
	IDs            param.Field[[]string] `json:"ids" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string]   `header:"Idempotency-Key"`
}

func (r PrismObjectContactBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectContactBulkUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string]                                   `path:"teamId" api:"required" format:"uuid"`
	Items          param.Field[[]PrismObjectContactBulkUpdateParamsItem] `json:"items" api:"required"`
	IdempotencyKey param.Field[string]                                   `header:"Idempotency-Key"`
}

func (r PrismObjectContactBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Object with `id` plus the same property body shape as PATCH
// (`default`/`list`/`extended`).
type PrismObjectContactBulkUpdateParamsItem struct {
	ID          param.Field[string]    `json:"id" api:"required" format:"uuid"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r PrismObjectContactBulkUpdateParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectContactCountParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the count to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectContactCountParams]'s query parameters as
// `url.Values`.
func (r PrismObjectContactCountParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectContactDuplicateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectContactFindParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Scope the lookup to a specific list/app.
	ListID param.Field[string] `query:"list_id" format:"uuid"`
}

// URLQuery serializes [PrismObjectContactFindParams]'s query parameters as
// `url.Values`.
func (r PrismObjectContactFindParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectContactGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Comma-separated property slugs to return. Use dot notation for relationships.
	// `id` is always returned at the top level. Defaults to all properties.
	Select param.Field[string] `query:"select"`
}

// URLQuery serializes [PrismObjectContactGetParams]'s query parameters as
// `url.Values`.
func (r PrismObjectContactGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrismObjectContactQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string]                               `path:"teamId" api:"required" format:"uuid"`
	Query  param.Field[PrismObjectContactQueryParamsQuery]   `json:"query" api:"required"`
	ID     param.Field[PrismObjectContactQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes  param.Field[[]string]                             `json:"boxes"`
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

func (r PrismObjectContactQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectContactQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name). `id` is always returned at the top level of each
	// row and does not need to be selected.
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectContactQueryParamsQueryCombinator] `json:"combinator"`
	// Opaque cursor from a previous response's `next_cursor`. Pass it back unchanged
	// to fetch the next page. When set, `page` and `limit` are derived from the cursor
	// and any explicit values are ignored.
	Cursor param.Field[string] `json:"cursor"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values may be option slugs or option UUIDs.
	Filter param.Field[[]map[string]PrismObjectContactQueryParamsQueryFilterUnion] `json:"filter"`
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

func (r PrismObjectContactQueryParamsQueryFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilter) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

// Satisfied by [PrismObjectContactQueryParamsQueryFilterPrismQueryFilterEq],
// [PrismObjectContactQueryParamsQueryFilterPrismQueryFilterNe],
// [PrismObjectContactQueryParamsQueryFilterPrismQueryFilterLt],
// [PrismObjectContactQueryParamsQueryFilterPrismQueryFilterGt],
// [PrismObjectContactQueryParamsQueryFilterPrismQueryFilterLte],
// [PrismObjectContactQueryParamsQueryFilterPrismQueryFilterGte],
// [PrismObjectContactQueryParamsQueryFilterContains],
// [PrismObjectContactQueryParamsQueryFilterBeginsWith],
// [PrismObjectContactQueryParamsQueryFilterEndsWith],
// [PrismObjectContactQueryParamsQueryFilterNotContains],
// [PrismObjectContactQueryParamsQueryFilterExists],
// [PrismObjectContactQueryParamsQueryFilterNotExists],
// [PrismObjectContactQueryParamsQueryFilterIsNull],
// [PrismObjectContactQueryParamsQueryFilterIsNotNull],
// [PrismObjectContactQueryParamsQueryFilterBetween],
// [PrismObjectContactQueryParamsQueryFilterIn],
// [PrismObjectContactQueryParamsQueryFilterNotIn],
// [PrismObjectContactQueryParamsQueryFilter].
type PrismObjectContactQueryParamsQueryFilterUnion interface {
	implementsPrismObjectContactQueryParamsQueryFilterUnion()
}

type PrismObjectContactQueryParamsQueryFilterPrismQueryFilterEq struct {
	Equals param.Field[PrismObjectContactQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion] `json:"=" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterPrismQueryFilterEq) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterPrismQueryFilterEq) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectContactQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion interface {
	ImplementsPrismObjectContactQueryParamsQueryFilterPrismQueryFilterEqEqualsUnion()
}

type PrismObjectContactQueryParamsQueryFilterPrismQueryFilterNe struct {
	NotEquals param.Field[PrismObjectContactQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion] `json:"!=" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterPrismQueryFilterNe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterPrismQueryFilterNe) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type PrismObjectContactQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion interface {
	ImplementsPrismObjectContactQueryParamsQueryFilterPrismQueryFilterNeNotEqualsUnion()
}

type PrismObjectContactQueryParamsQueryFilterPrismQueryFilterLt struct {
	Less param.Field[string] `json:"<" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterPrismQueryFilterLt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterPrismQueryFilterLt) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQueryFilterPrismQueryFilterGt struct {
	Greater param.Field[string] `json:">" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterPrismQueryFilterGt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterPrismQueryFilterGt) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQueryFilterPrismQueryFilterLte struct {
	LessOrEquals param.Field[string] `json:"<=" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterPrismQueryFilterLte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterPrismQueryFilterLte) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQueryFilterPrismQueryFilterGte struct {
	GreaterOrEquals param.Field[string] `json:">=" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterPrismQueryFilterGte) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterPrismQueryFilterGte) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQueryFilterContains struct {
	Contains param.Field[PrismObjectContactQueryParamsQueryFilterContainsContainsUnion] `json:"contains" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterContains) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterContains) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectContactQueryParamsQueryFilterContainsContainsArray].
type PrismObjectContactQueryParamsQueryFilterContainsContainsUnion interface {
	ImplementsPrismObjectContactQueryParamsQueryFilterContainsContainsUnion()
}

type PrismObjectContactQueryParamsQueryFilterContainsContainsArray []string

func (r PrismObjectContactQueryParamsQueryFilterContainsContainsArray) ImplementsPrismObjectContactQueryParamsQueryFilterContainsContainsUnion() {
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

type PrismObjectContactQueryParamsQueryFilterIsNull struct {
	IsNull param.Field[PrismObjectContactQueryParamsQueryFilterIsNullIsNullUnion] `json:"is_null" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterIsNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterIsNull) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectContactQueryParamsQueryFilterIsNullIsNullArray].
type PrismObjectContactQueryParamsQueryFilterIsNullIsNullUnion interface {
	ImplementsPrismObjectContactQueryParamsQueryFilterIsNullIsNullUnion()
}

type PrismObjectContactQueryParamsQueryFilterIsNullIsNullArray []string

func (r PrismObjectContactQueryParamsQueryFilterIsNullIsNullArray) ImplementsPrismObjectContactQueryParamsQueryFilterIsNullIsNullUnion() {
}

type PrismObjectContactQueryParamsQueryFilterIsNotNull struct {
	IsNotNull param.Field[PrismObjectContactQueryParamsQueryFilterIsNotNullIsNotNullUnion] `json:"is_not_null" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterIsNotNull) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterIsNotNull) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectContactQueryParamsQueryFilterIsNotNullIsNotNullArray].
type PrismObjectContactQueryParamsQueryFilterIsNotNullIsNotNullUnion interface {
	ImplementsPrismObjectContactQueryParamsQueryFilterIsNotNullIsNotNullUnion()
}

type PrismObjectContactQueryParamsQueryFilterIsNotNullIsNotNullArray []string

func (r PrismObjectContactQueryParamsQueryFilterIsNotNullIsNotNullArray) ImplementsPrismObjectContactQueryParamsQueryFilterIsNotNullIsNotNullUnion() {
}

type PrismObjectContactQueryParamsQueryFilterBetween struct {
	Between param.Field[PrismObjectContactQueryParamsQueryFilterBetweenBetweenUnion] `json:"between" api:"required"`
}

func (r PrismObjectContactQueryParamsQueryFilterBetween) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrismObjectContactQueryParamsQueryFilterBetween) implementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectContactQueryParamsQueryFilterBetweenBetweenArray].
type PrismObjectContactQueryParamsQueryFilterBetweenBetweenUnion interface {
	ImplementsPrismObjectContactQueryParamsQueryFilterBetweenBetweenUnion()
}

type PrismObjectContactQueryParamsQueryFilterBetweenBetweenArray []string

func (r PrismObjectContactQueryParamsQueryFilterBetweenBetweenArray) ImplementsPrismObjectContactQueryParamsQueryFilterBetweenBetweenUnion() {
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
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type PrismObjectContactUpsertParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
	IdempotencyKey        param.Field[string]        `header:"Idempotency-Key"`
}

func (r PrismObjectContactUpsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}
