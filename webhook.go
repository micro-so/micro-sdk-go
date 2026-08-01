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

// WebhookService contains methods and other services that help with interacting
// with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options    []option.RequestOption
	Deliveries *WebhookDeliveryService
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	r.Deliveries = NewWebhookDeliveryService(opts...)
	return
}

// Registers a webhook and enqueues an asynchronous verification handshake (run by
// the dispatcher). The response includes the signing `secret`, shown only this
// once; `verified` is false until the handshake passes.
func (r *WebhookService) New(ctx context.Context, params WebhookNewParams, opts ...option.RequestOption) (res *WebhookWithSecret, err error) {
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
	path := fmt.Sprintf("v2/webhooks/%s", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Updates mutable fields. Changing `url` resets verification and re-runs the
// handshake.
func (r *WebhookService) Update(ctx context.Context, webhookID string, params WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookUpdateResponse, err error) {
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
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/webhooks/%s/%s", params.TeamID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the team's webhooks. Signing secrets are never included.
func (r *WebhookService) List(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
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
	path := fmt.Sprintf("v2/webhooks/%s", query.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a webhook
func (r *WebhookService) Delete(ctx context.Context, webhookID string, body WebhookDeleteParams, opts ...option.RequestOption) (err error) {
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
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return err
	}
	path := fmt.Sprintf("v2/webhooks/%s/%s", body.TeamID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get a webhook
func (r *WebhookService) Get(ctx context.Context, webhookID string, query WebhookGetParams, opts ...option.RequestOption) (res *Webhook, err error) {
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
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/webhooks/%s/%s", query.TeamID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Account-wide delivery feed across all of the team's webhooks, newest first.
func (r *WebhookService) ListDeliveries(ctx context.Context, params WebhookListDeliveriesParams, opts ...option.RequestOption) (res *WebhookListDeliveriesResponse, err error) {
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
	path := fmt.Sprintf("v2/webhooks/%s/deliveries", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Fire-and-forget test delivery through the async dispatcher. The webhook must be
// enabled and verified.
func (r *WebhookService) Ping(ctx context.Context, webhookID string, params WebhookPingParams, opts ...option.RequestOption) (res *WebhookPingResponse, err error) {
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
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/webhooks/%s/%s/ping", params.TeamID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Re-runs the GET challenge/echo handshake against the webhook's url and updates
// its verified state.
func (r *WebhookService) Verify(ctx context.Context, webhookID string, body WebhookVerifyParams, opts ...option.RequestOption) (res *WebhookVerifyResponse, err error) {
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
	if webhookID == "" {
		err = errors.New("missing required webhookId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/webhooks/%s/%s/verify", body.TeamID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// A registered webhook endpoint.
type Webhook struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Disabled webhooks are skipped at delivery time.
	Enabled bool   `json:"enabled" api:"required"`
	Name    string `json:"name" api:"required"`
	TeamID  string `json:"team_id" api:"required" format:"uuid"`
	// Endpoint events are delivered to.
	URL string `json:"url" api:"required" format:"uri"`
	// True once the endpoint has completed the verification handshake.
	Verified    bool      `json:"verified" api:"required"`
	Description string    `json:"description" api:"nullable"`
	UpdatedAt   time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// Stable token replayed to the endpoint (as the `micro_hook_token` query param)
	// during the verification handshake. The endpoint may check it to confirm the
	// request originated from Micro.
	VerificationToken string      `json:"verification_token"`
	VerifiedAt        time.Time   `json:"verified_at" api:"nullable" format:"date-time"`
	JSON              webhookJSON `json:"-"`
}

// webhookJSON contains the JSON metadata for the struct [Webhook]
type webhookJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	Enabled           apijson.Field
	Name              apijson.Field
	TeamID            apijson.Field
	URL               apijson.Field
	Verified          apijson.Field
	Description       apijson.Field
	UpdatedAt         apijson.Field
	VerificationToken apijson.Field
	VerifiedAt        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Webhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookJSON) RawJSON() string {
	return r.raw
}

// On create, the dispatcher asynchronously runs a verification handshake: it sends
// a GET to `url` with `micro_hook_mode=subscribe`, a one-time
// `micro_hook_challenge`, and the webhook's `micro_hook_token`. The endpoint must
// respond 200 and echo the challenge value verbatim in the body; on success the
// webhook's `verified` flag flips to true. A failed handshake does not fail
// creation — re-run it later via the verify endpoint.
type WebhookCreateParam struct {
	Name param.Field[string] `json:"name" api:"required"`
	// HTTP(S) endpoint. Rejected if it resolves to a private/internal address.
	URL         param.Field[string] `json:"url" api:"required" format:"uri"`
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
}

func (r WebhookCreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A webhook delivery — one logical event delivery to an endpoint, grouping its
// attempts. Status and status_code reflect the latest attempt.
type WebhookDelivery struct {
	CreatedAt  time.Time             `json:"created_at" api:"required" format:"date-time"`
	DeliveryID string                `json:"delivery_id" api:"required" format:"uuid"`
	Status     WebhookDeliveryStatus `json:"status" api:"required"`
	Type       WebhookDeliveryType   `json:"type" api:"required"`
	WebhookID  string                `json:"webhook_id" api:"required" format:"uuid"`
	// Number of attempts made so far (including async retries).
	Attempts int64 `json:"attempts" api:"nullable"`
	// Event name (e.g. `webhook.test`); `verification` for handshake runs.
	Event string `json:"event" api:"nullable"`
	// HTTP status of the latest attempt; null on a transport error.
	StatusCode int64               `json:"status_code" api:"nullable"`
	TeamID     string              `json:"team_id" api:"nullable" format:"uuid"`
	UpdatedAt  time.Time           `json:"updated_at" api:"nullable" format:"date-time"`
	URL        string              `json:"url"`
	JSON       webhookDeliveryJSON `json:"-"`
}

// webhookDeliveryJSON contains the JSON metadata for the struct [WebhookDelivery]
type webhookDeliveryJSON struct {
	CreatedAt   apijson.Field
	DeliveryID  apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	WebhookID   apijson.Field
	Attempts    apijson.Field
	Event       apijson.Field
	StatusCode  apijson.Field
	TeamID      apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookDelivery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookDeliveryJSON) RawJSON() string {
	return r.raw
}

type WebhookDeliveryStatus string

const (
	WebhookDeliveryStatusSuccess WebhookDeliveryStatus = "success"
	WebhookDeliveryStatusFailed  WebhookDeliveryStatus = "failed"
)

func (r WebhookDeliveryStatus) IsKnown() bool {
	switch r {
	case WebhookDeliveryStatusSuccess, WebhookDeliveryStatusFailed:
		return true
	}
	return false
}

type WebhookDeliveryType string

const (
	WebhookDeliveryTypeDelivery     WebhookDeliveryType = "delivery"
	WebhookDeliveryTypeVerification WebhookDeliveryType = "verification"
)

func (r WebhookDeliveryType) IsKnown() bool {
	switch r {
	case WebhookDeliveryTypeDelivery, WebhookDeliveryTypeVerification:
		return true
	}
	return false
}

// A delivery plus its full attempt timeline.
type WebhookDeliveryDetail struct {
	AttemptHistory []WebhookDeliveryDetailAttemptHistory `json:"attempt_history"`
	JSON           webhookDeliveryDetailJSON             `json:"-"`
	WebhookDelivery
}

// webhookDeliveryDetailJSON contains the JSON metadata for the struct
// [WebhookDeliveryDetail]
type webhookDeliveryDetailJSON struct {
	AttemptHistory apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WebhookDeliveryDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookDeliveryDetailJSON) RawJSON() string {
	return r.raw
}

// A single HTTP attempt within a delivery (including async retries).
type WebhookDeliveryDetailAttemptHistory struct {
	// 1-based attempt number.
	Attempt   int64                                     `json:"attempt" api:"required"`
	CreatedAt time.Time                                 `json:"created_at" api:"required" format:"date-time"`
	Status    WebhookDeliveryDetailAttemptHistoryStatus `json:"status" api:"required"`
	// Failure reason, when status is failed.
	Error string `json:"error" api:"nullable"`
	// Body sent to the endpoint (delivery only); may be truncated.
	RequestBody string `json:"request_body" api:"nullable"`
	// Body returned by the endpoint; may be truncated.
	ResponseBody string                                  `json:"response_body" api:"nullable"`
	StatusCode   int64                                   `json:"status_code" api:"nullable"`
	JSON         webhookDeliveryDetailAttemptHistoryJSON `json:"-"`
}

// webhookDeliveryDetailAttemptHistoryJSON contains the JSON metadata for the
// struct [WebhookDeliveryDetailAttemptHistory]
type webhookDeliveryDetailAttemptHistoryJSON struct {
	Attempt      apijson.Field
	CreatedAt    apijson.Field
	Status       apijson.Field
	Error        apijson.Field
	RequestBody  apijson.Field
	ResponseBody apijson.Field
	StatusCode   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WebhookDeliveryDetailAttemptHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookDeliveryDetailAttemptHistoryJSON) RawJSON() string {
	return r.raw
}

type WebhookDeliveryDetailAttemptHistoryStatus string

const (
	WebhookDeliveryDetailAttemptHistoryStatusSuccess WebhookDeliveryDetailAttemptHistoryStatus = "success"
	WebhookDeliveryDetailAttemptHistoryStatusFailed  WebhookDeliveryDetailAttemptHistoryStatus = "failed"
)

func (r WebhookDeliveryDetailAttemptHistoryStatus) IsKnown() bool {
	switch r {
	case WebhookDeliveryDetailAttemptHistoryStatusSuccess, WebhookDeliveryDetailAttemptHistoryStatusFailed:
		return true
	}
	return false
}

// Partial update. Changing `url` resets verification and re-runs the handshake.
type WebhookUpdateParam struct {
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
	Name        param.Field[string] `json:"name"`
	URL         param.Field[string] `json:"url" format:"uri"`
}

func (r WebhookUpdateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Returned ONLY on creation. Includes the signing secret (shown once) and the
// pending verification status.
type WebhookWithSecret struct {
	// HMAC signing secret (prefix `whsec_`). Store it now — it is never returned
	// again. The dispatcher signs each delivered payload with it so your endpoint can
	// verify authenticity.
	Secret string `json:"secret" api:"required"`
	// Status of the verification handshake enqueued by this request. The handshake
	// runs asynchronously in the dispatcher; poll the webhook (its `verified` flag
	// flips to true on success) to observe the outcome.
	Verification WebhookWithSecretVerification `json:"verification"`
	JSON         webhookWithSecretJSON         `json:"-"`
	Webhook
}

// webhookWithSecretJSON contains the JSON metadata for the struct
// [WebhookWithSecret]
type webhookWithSecretJSON struct {
	Secret       apijson.Field
	Verification apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WebhookWithSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookWithSecretJSON) RawJSON() string {
	return r.raw
}

// Status of the verification handshake enqueued by this request. The handshake
// runs asynchronously in the dispatcher; poll the webhook (its `verified` flag
// flips to true on success) to observe the outcome.
type WebhookWithSecretVerification struct {
	// Always `pending` at the moment of the response — the dispatcher has been asked
	// to run the handshake but has not reported back yet.
	Status WebhookWithSecretVerificationStatus `json:"status" api:"required"`
	JSON   webhookWithSecretVerificationJSON   `json:"-"`
}

// webhookWithSecretVerificationJSON contains the JSON metadata for the struct
// [WebhookWithSecretVerification]
type webhookWithSecretVerificationJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookWithSecretVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookWithSecretVerificationJSON) RawJSON() string {
	return r.raw
}

// Always `pending` at the moment of the response — the dispatcher has been asked
// to run the handshake but has not reported back yet.
type WebhookWithSecretVerificationStatus string

const (
	WebhookWithSecretVerificationStatusPending WebhookWithSecretVerificationStatus = "pending"
)

func (r WebhookWithSecretVerificationStatus) IsKnown() bool {
	switch r {
	case WebhookWithSecretVerificationStatusPending:
		return true
	}
	return false
}

// A webhook plus the status of a verification handshake enqueued by this request.
type WebhookUpdateResponse struct {
	// Status of the verification handshake enqueued by this request. The handshake
	// runs asynchronously in the dispatcher; poll the webhook (its `verified` flag
	// flips to true on success) to observe the outcome.
	Verification WebhookUpdateResponseVerification `json:"verification"`
	JSON         webhookUpdateResponseJSON         `json:"-"`
	Webhook
}

// webhookUpdateResponseJSON contains the JSON metadata for the struct
// [WebhookUpdateResponse]
type webhookUpdateResponseJSON struct {
	Verification apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WebhookUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the verification handshake enqueued by this request. The handshake
// runs asynchronously in the dispatcher; poll the webhook (its `verified` flag
// flips to true on success) to observe the outcome.
type WebhookUpdateResponseVerification struct {
	// Always `pending` at the moment of the response — the dispatcher has been asked
	// to run the handshake but has not reported back yet.
	Status WebhookUpdateResponseVerificationStatus `json:"status" api:"required"`
	JSON   webhookUpdateResponseVerificationJSON   `json:"-"`
}

// webhookUpdateResponseVerificationJSON contains the JSON metadata for the struct
// [WebhookUpdateResponseVerification]
type webhookUpdateResponseVerificationJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookUpdateResponseVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookUpdateResponseVerificationJSON) RawJSON() string {
	return r.raw
}

// Always `pending` at the moment of the response — the dispatcher has been asked
// to run the handshake but has not reported back yet.
type WebhookUpdateResponseVerificationStatus string

const (
	WebhookUpdateResponseVerificationStatusPending WebhookUpdateResponseVerificationStatus = "pending"
)

func (r WebhookUpdateResponseVerificationStatus) IsKnown() bool {
	switch r {
	case WebhookUpdateResponseVerificationStatusPending:
		return true
	}
	return false
}

type WebhookListResponse struct {
	Data []Webhook               `json:"data" api:"required"`
	JSON webhookListResponseJSON `json:"-"`
}

// webhookListResponseJSON contains the JSON metadata for the struct
// [WebhookListResponse]
type webhookListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookListResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookListDeliveriesResponse struct {
	Data []WebhookDelivery `json:"data" api:"required"`
	// Pass as `cursor` to fetch the next page; null when there are no more.
	NextCursor string                            `json:"next_cursor" api:"nullable"`
	JSON       webhookListDeliveriesResponseJSON `json:"-"`
}

// webhookListDeliveriesResponseJSON contains the JSON metadata for the struct
// [WebhookListDeliveriesResponse]
type webhookListDeliveriesResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookListDeliveriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookListDeliveriesResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookPingResponse struct {
	Dispatched bool                    `json:"dispatched" api:"required"`
	Event      string                  `json:"event" api:"required"`
	WebhookID  string                  `json:"webhook_id" api:"required" format:"uuid"`
	JSON       webhookPingResponseJSON `json:"-"`
}

// webhookPingResponseJSON contains the JSON metadata for the struct
// [WebhookPingResponse]
type webhookPingResponseJSON struct {
	Dispatched  apijson.Field
	Event       apijson.Field
	WebhookID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookPingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookPingResponseJSON) RawJSON() string {
	return r.raw
}

// A webhook plus the status of a verification handshake enqueued by this request.
type WebhookVerifyResponse struct {
	// Status of the verification handshake enqueued by this request. The handshake
	// runs asynchronously in the dispatcher; poll the webhook (its `verified` flag
	// flips to true on success) to observe the outcome.
	Verification WebhookVerifyResponseVerification `json:"verification"`
	JSON         webhookVerifyResponseJSON         `json:"-"`
	Webhook
}

// webhookVerifyResponseJSON contains the JSON metadata for the struct
// [WebhookVerifyResponse]
type webhookVerifyResponseJSON struct {
	Verification apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WebhookVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookVerifyResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the verification handshake enqueued by this request. The handshake
// runs asynchronously in the dispatcher; poll the webhook (its `verified` flag
// flips to true on success) to observe the outcome.
type WebhookVerifyResponseVerification struct {
	// Always `pending` at the moment of the response — the dispatcher has been asked
	// to run the handshake but has not reported back yet.
	Status WebhookVerifyResponseVerificationStatus `json:"status" api:"required"`
	JSON   webhookVerifyResponseVerificationJSON   `json:"-"`
}

// webhookVerifyResponseVerificationJSON contains the JSON metadata for the struct
// [WebhookVerifyResponseVerification]
type webhookVerifyResponseVerificationJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookVerifyResponseVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookVerifyResponseVerificationJSON) RawJSON() string {
	return r.raw
}

// Always `pending` at the moment of the response — the dispatcher has been asked
// to run the handshake but has not reported back yet.
type WebhookVerifyResponseVerificationStatus string

const (
	WebhookVerifyResponseVerificationStatusPending WebhookVerifyResponseVerificationStatus = "pending"
)

func (r WebhookVerifyResponseVerificationStatus) IsKnown() bool {
	switch r {
	case WebhookVerifyResponseVerificationStatusPending:
		return true
	}
	return false
}

type WebhookNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// On create, the dispatcher asynchronously runs a verification handshake: it sends
	// a GET to `url` with `micro_hook_mode=subscribe`, a one-time
	// `micro_hook_challenge`, and the webhook's `micro_hook_token`. The endpoint must
	// respond 200 and echo the challenge value verbatim in the body; on success the
	// webhook's `verified` flag flips to true. A failed handshake does not fail
	// creation — re-run it later via the verify endpoint.
	WebhookCreate WebhookCreateParam `json:"webhook_create" api:"required"`
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.WebhookCreate)
}

type WebhookUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Partial update. Changing `url` resets verification and re-runs the handshake.
	WebhookUpdate WebhookUpdateParam `json:"webhook_update" api:"required"`
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.WebhookUpdate)
}

type WebhookListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type WebhookDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type WebhookGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type WebhookListDeliveriesParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Opaque cursor from a previous response's `next_cursor`.
	Cursor param.Field[string] `query:"cursor"`
	// Page size (1–100, default 25).
	Limit param.Field[int64] `query:"limit"`
	// Filter by outcome.
	Status param.Field[WebhookListDeliveriesParamsStatus] `query:"status"`
	// Filter by run type. Defaults to `delivery` (event deliveries). Pass `all` to
	// include verification handshakes.
	Type param.Field[WebhookListDeliveriesParamsType] `query:"type"`
}

// URLQuery serializes [WebhookListDeliveriesParams]'s query parameters as
// `url.Values`.
func (r WebhookListDeliveriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by outcome.
type WebhookListDeliveriesParamsStatus string

const (
	WebhookListDeliveriesParamsStatusSuccess WebhookListDeliveriesParamsStatus = "success"
	WebhookListDeliveriesParamsStatusFailed  WebhookListDeliveriesParamsStatus = "failed"
)

func (r WebhookListDeliveriesParamsStatus) IsKnown() bool {
	switch r {
	case WebhookListDeliveriesParamsStatusSuccess, WebhookListDeliveriesParamsStatusFailed:
		return true
	}
	return false
}

// Filter by run type. Defaults to `delivery` (event deliveries). Pass `all` to
// include verification handshakes.
type WebhookListDeliveriesParamsType string

const (
	WebhookListDeliveriesParamsTypeDelivery     WebhookListDeliveriesParamsType = "delivery"
	WebhookListDeliveriesParamsTypeVerification WebhookListDeliveriesParamsType = "verification"
	WebhookListDeliveriesParamsTypeAll          WebhookListDeliveriesParamsType = "all"
)

func (r WebhookListDeliveriesParamsType) IsKnown() bool {
	switch r {
	case WebhookListDeliveriesParamsTypeDelivery, WebhookListDeliveriesParamsTypeVerification, WebhookListDeliveriesParamsTypeAll:
		return true
	}
	return false
}

type WebhookPingParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Arbitrary JSON payload body.
	Data param.Field[map[string]interface{}] `json:"data"`
	// Event name to send.
	Event param.Field[string] `json:"event"`
}

func (r WebhookPingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookVerifyParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
