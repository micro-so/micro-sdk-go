// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"context"
	"net/http"
	"slices"

	"github.com/micro-so/micro-sdk-go/internal/apijson"
	"github.com/micro-so/micro-sdk-go/internal/requestconfig"
	"github.com/micro-so/micro-sdk-go/option"
)

// RealtimeService contains methods and other services that help with interacting
// with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRealtimeService] method instead.
type RealtimeService struct {
	Options []option.RequestOption
}

// NewRealtimeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRealtimeService(opts ...option.RequestOption) (r *RealtimeService) {
	r = &RealtimeService{}
	r.Options = opts
	return
}

// Exchange your API key (or session) for a short-lived ticket that authenticates a
// connection to the realtime object-change stream. Open a WebSocket to the push
// endpoint with the returned ticket as the `token` query parameter. The ticket is
// single-purpose and expires quickly; call this again to obtain a fresh one before
// reconnecting.
func (r *RealtimeService) NewTicket(ctx context.Context, opts ...option.RequestOption) (res *RealtimeNewTicketResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/realtime/ticket"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type RealtimeNewTicketResponse struct {
	// Seconds until the ticket expires. Refresh (call the endpoint again) before
	// reconnecting.
	ExpiresIn int64 `json:"expires_in" api:"required"`
	// Short-lived token authenticating a realtime WebSocket connection. Pass as the
	// `token` query parameter when connecting.
	Ticket string `json:"ticket" api:"required"`
	// WebSocket URL for this environment (wss://stream.developers[.staging].micro.so).
	// Connect here with the ticket as the `token` query parameter.
	WsURL string                        `json:"ws_url" api:"required,nullable"`
	JSON  realtimeNewTicketResponseJSON `json:"-"`
}

// realtimeNewTicketResponseJSON contains the JSON metadata for the struct
// [RealtimeNewTicketResponse]
type realtimeNewTicketResponseJSON struct {
	ExpiresIn   apijson.Field
	Ticket      apijson.Field
	WsURL       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RealtimeNewTicketResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r realtimeNewTicketResponseJSON) RawJSON() string {
	return r.raw
}
