// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/micro-so/micro-sdk-go/internal/apijson"
	"github.com/micro-so/micro-sdk-go/internal/param"
	"github.com/micro-so/micro-sdk-go/internal/requestconfig"
	"github.com/micro-so/micro-sdk-go/option"
)

// PrismImportService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismImportService] method instead.
type PrismImportService struct {
	Options []option.RequestOption
}

// NewPrismImportService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrismImportService(opts ...option.RequestOption) (r *PrismImportService) {
	r = &PrismImportService{}
	r.Options = opts
	return
}

// Poll the status of an async import. Sync imports complete in the original
// response and don't appear here. Async jobs are retained for 7 days. Returns 404
// once the job has expired.
func (r *PrismImportService) Get(ctx context.Context, jobID string, query PrismImportGetParams, opts ...option.RequestOption) (res *PrismImportGetResponse, err error) {
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
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/imports/%s", query.TeamID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Status snapshot of an import job. Same shape used by the POST /import response
// and by GET /imports/{jobId}.
type PrismImportGetResponse struct {
	// Null for sync imports (results inlined). Set for async imports.
	JobID  string                       `json:"job_id" api:"required,nullable"`
	Status PrismImportGetResponseStatus `json:"status" api:"required"`
	// Total number of rows in the import.
	Total     int64     `json:"total" api:"required"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Set when status=failed; describes the job-level failure (not per-row).
	Error     PrismImportGetResponseError `json:"error"`
	ExpiresAt time.Time                   `json:"expires_at" format:"date-time"`
	Failed    int64                       `json:"failed"`
	// Rows that have been attempted (succeeded + failed).
	Processed int64 `json:"processed"`
	// Per-row outcomes. Always present for sync imports; populated for async imports
	// once the job reaches `complete`.
	Results   []PrismImportGetResponseResult `json:"results"`
	Succeeded int64                          `json:"succeeded"`
	UpdatedAt time.Time                      `json:"updated_at" format:"date-time"`
	JSON      prismImportGetResponseJSON     `json:"-"`
}

// prismImportGetResponseJSON contains the JSON metadata for the struct
// [PrismImportGetResponse]
type prismImportGetResponseJSON struct {
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

func (r *PrismImportGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismImportGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismImportGetResponseStatus string

const (
	PrismImportGetResponseStatusComplete   PrismImportGetResponseStatus = "complete"
	PrismImportGetResponseStatusProcessing PrismImportGetResponseStatus = "processing"
	PrismImportGetResponseStatusFailed     PrismImportGetResponseStatus = "failed"
)

func (r PrismImportGetResponseStatus) IsKnown() bool {
	switch r {
	case PrismImportGetResponseStatusComplete, PrismImportGetResponseStatusProcessing, PrismImportGetResponseStatusFailed:
		return true
	}
	return false
}

// Set when status=failed; describes the job-level failure (not per-row).
type PrismImportGetResponseError struct {
	Code    string                          `json:"code"`
	Message string                          `json:"message"`
	JSON    prismImportGetResponseErrorJSON `json:"-"`
}

// prismImportGetResponseErrorJSON contains the JSON metadata for the struct
// [PrismImportGetResponseError]
type prismImportGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismImportGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismImportGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type PrismImportGetResponseResult struct {
	ID      string                             `json:"id" api:"nullable" format:"uuid"`
	Created bool                               `json:"created"`
	Error   PrismImportGetResponseResultsError `json:"error"`
	// True if the row matched an existing record via the dedupe key.
	Existing bool                             `json:"existing"`
	JSON     prismImportGetResponseResultJSON `json:"-"`
}

// prismImportGetResponseResultJSON contains the JSON metadata for the struct
// [PrismImportGetResponseResult]
type prismImportGetResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismImportGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismImportGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismImportGetResponseResultsError struct {
	Code    string                                 `json:"code"`
	Message string                                 `json:"message"`
	JSON    prismImportGetResponseResultsErrorJSON `json:"-"`
}

// prismImportGetResponseResultsErrorJSON contains the JSON metadata for the struct
// [PrismImportGetResponseResultsError]
type prismImportGetResponseResultsErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismImportGetResponseResultsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismImportGetResponseResultsErrorJSON) RawJSON() string {
	return r.raw
}

type PrismImportGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
