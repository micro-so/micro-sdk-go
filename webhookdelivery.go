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

// WebhookDeliveryService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookDeliveryService] method instead.
type WebhookDeliveryService struct {
	Options []option.RequestOption
}

// NewWebhookDeliveryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookDeliveryService(opts ...option.RequestOption) (r *WebhookDeliveryService) {
	r = &WebhookDeliveryService{}
	r.Options = opts
	return
}

// An endpoint's deliveries, newest first, with optional status / type / time-range
// filters and cursor pagination.
func (r *WebhookDeliveryService) List(ctx context.Context, webhookID string, params WebhookDeliveryListParams, opts ...option.RequestOption) (res *WebhookDeliveryListResponse, err error) {
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
	path := fmt.Sprintf("v2/webhooks/%s/%s/deliveries", params.TeamID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// A single delivery plus its full attempt timeline (including async retries).
func (r *WebhookDeliveryService) Get(ctx context.Context, webhookID string, deliveryID string, query WebhookDeliveryGetParams, opts ...option.RequestOption) (res *WebhookDeliveryDetail, err error) {
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
	if deliveryID == "" {
		err = errors.New("missing required deliveryId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/webhooks/%s/%s/deliveries/%s", query.TeamID, webhookID, deliveryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type WebhookDeliveryListResponse struct {
	Data []WebhookDelivery `json:"data" api:"required"`
	// Pass as `cursor` to fetch the next page; null when there are no more.
	NextCursor string                          `json:"next_cursor" api:"nullable"`
	JSON       webhookDeliveryListResponseJSON `json:"-"`
}

// webhookDeliveryListResponseJSON contains the JSON metadata for the struct
// [WebhookDeliveryListResponse]
type webhookDeliveryListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookDeliveryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookDeliveryListResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookDeliveryListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Only deliveries at or after this ISO-8601 timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Only deliveries at or before this ISO-8601 timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Opaque cursor from a previous response's `next_cursor`.
	Cursor param.Field[string] `query:"cursor"`
	// Page size (1–100, default 25).
	Limit param.Field[int64] `query:"limit"`
	// Filter by outcome.
	Status param.Field[WebhookDeliveryListParamsStatus] `query:"status"`
	// Filter by run type. Defaults to `delivery` (event deliveries). Pass `all` to
	// include verification handshakes.
	Type param.Field[WebhookDeliveryListParamsType] `query:"type"`
}

// URLQuery serializes [WebhookDeliveryListParams]'s query parameters as
// `url.Values`.
func (r WebhookDeliveryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by outcome.
type WebhookDeliveryListParamsStatus string

const (
	WebhookDeliveryListParamsStatusSuccess WebhookDeliveryListParamsStatus = "success"
	WebhookDeliveryListParamsStatusFailed  WebhookDeliveryListParamsStatus = "failed"
)

func (r WebhookDeliveryListParamsStatus) IsKnown() bool {
	switch r {
	case WebhookDeliveryListParamsStatusSuccess, WebhookDeliveryListParamsStatusFailed:
		return true
	}
	return false
}

// Filter by run type. Defaults to `delivery` (event deliveries). Pass `all` to
// include verification handshakes.
type WebhookDeliveryListParamsType string

const (
	WebhookDeliveryListParamsTypeDelivery     WebhookDeliveryListParamsType = "delivery"
	WebhookDeliveryListParamsTypeVerification WebhookDeliveryListParamsType = "verification"
	WebhookDeliveryListParamsTypeAll          WebhookDeliveryListParamsType = "all"
)

func (r WebhookDeliveryListParamsType) IsKnown() bool {
	switch r {
	case WebhookDeliveryListParamsTypeDelivery, WebhookDeliveryListParamsTypeVerification, WebhookDeliveryListParamsTypeAll:
		return true
	}
	return false
}

type WebhookDeliveryGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
