// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/micro-so/micro-sdk-go/internal/apijson"
	"github.com/micro-so/micro-sdk-go/internal/param"
	"github.com/micro-so/micro-sdk-go/internal/requestconfig"
	"github.com/micro-so/micro-sdk-go/option"
)

// PrismObjectDealGrantService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectDealGrantService] method instead.
type PrismObjectDealGrantService struct {
	Options []option.RequestOption
}

// NewPrismObjectDealGrantService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectDealGrantService(opts ...option.RequestOption) (r *PrismObjectDealGrantService) {
	r = &PrismObjectDealGrantService{}
	r.Options = opts
	return
}

// Update grant
func (r *PrismObjectDealGrantService) Update(ctx context.Context, dealID string, params PrismObjectDealGrantUpdateParams, opts ...option.RequestOption) (res *PrismObjectDealGrantUpdateResponse, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.PathTeamID, precfg.TeamID)
	if params.PathTeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return nil, err
	}
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/deal/%s/grant", params.PathTeamID, dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Get grant
func (r *PrismObjectDealGrantService) Get(ctx context.Context, dealID string, query PrismObjectDealGrantGetParams, opts ...option.RequestOption) (res *PrismObjectDealGrantGetResponse, err error) {
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
	if dealID == "" {
		err = errors.New("missing required dealId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/deal/%s/grant", query.TeamID, dealID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PrismObjectDealGrantUpdateResponse struct {
	TeamGroupID []map[string]PrismObjectDealGrantUpdateResponseTeamGroupID `json:"team_group_id"`
	TeamID      map[string]PrismObjectDealGrantUpdateResponseTeamID        `json:"team_id"`
	UserID      []map[string]PrismObjectDealGrantUpdateResponseUserID      `json:"user_id"`
	JSON        prismObjectDealGrantUpdateResponseJSON                     `json:"-"`
}

// prismObjectDealGrantUpdateResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealGrantUpdateResponse]
type prismObjectDealGrantUpdateResponseJSON struct {
	TeamGroupID apijson.Field
	TeamID      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealGrantUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealGrantUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealGrantUpdateResponseTeamGroupID string

const (
	PrismObjectDealGrantUpdateResponseTeamGroupIDA PrismObjectDealGrantUpdateResponseTeamGroupID = "a"
	PrismObjectDealGrantUpdateResponseTeamGroupIDR PrismObjectDealGrantUpdateResponseTeamGroupID = "r"
	PrismObjectDealGrantUpdateResponseTeamGroupIDW PrismObjectDealGrantUpdateResponseTeamGroupID = "w"
)

func (r PrismObjectDealGrantUpdateResponseTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectDealGrantUpdateResponseTeamGroupIDA, PrismObjectDealGrantUpdateResponseTeamGroupIDR, PrismObjectDealGrantUpdateResponseTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectDealGrantUpdateResponseTeamID string

const (
	PrismObjectDealGrantUpdateResponseTeamIDA PrismObjectDealGrantUpdateResponseTeamID = "a"
	PrismObjectDealGrantUpdateResponseTeamIDR PrismObjectDealGrantUpdateResponseTeamID = "r"
	PrismObjectDealGrantUpdateResponseTeamIDW PrismObjectDealGrantUpdateResponseTeamID = "w"
)

func (r PrismObjectDealGrantUpdateResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectDealGrantUpdateResponseTeamIDA, PrismObjectDealGrantUpdateResponseTeamIDR, PrismObjectDealGrantUpdateResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectDealGrantUpdateResponseUserID string

const (
	PrismObjectDealGrantUpdateResponseUserIDA PrismObjectDealGrantUpdateResponseUserID = "a"
	PrismObjectDealGrantUpdateResponseUserIDR PrismObjectDealGrantUpdateResponseUserID = "r"
	PrismObjectDealGrantUpdateResponseUserIDW PrismObjectDealGrantUpdateResponseUserID = "w"
)

func (r PrismObjectDealGrantUpdateResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectDealGrantUpdateResponseUserIDA, PrismObjectDealGrantUpdateResponseUserIDR, PrismObjectDealGrantUpdateResponseUserIDW:
		return true
	}
	return false
}

type PrismObjectDealGrantGetResponse struct {
	TeamGroupID []map[string]PrismObjectDealGrantGetResponseTeamGroupID `json:"team_group_id"`
	TeamID      map[string]PrismObjectDealGrantGetResponseTeamID        `json:"team_id"`
	UserID      []map[string]PrismObjectDealGrantGetResponseUserID      `json:"user_id"`
	JSON        prismObjectDealGrantGetResponseJSON                     `json:"-"`
}

// prismObjectDealGrantGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectDealGrantGetResponse]
type prismObjectDealGrantGetResponseJSON struct {
	TeamGroupID apijson.Field
	TeamID      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDealGrantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDealGrantGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDealGrantGetResponseTeamGroupID string

const (
	PrismObjectDealGrantGetResponseTeamGroupIDA PrismObjectDealGrantGetResponseTeamGroupID = "a"
	PrismObjectDealGrantGetResponseTeamGroupIDR PrismObjectDealGrantGetResponseTeamGroupID = "r"
	PrismObjectDealGrantGetResponseTeamGroupIDW PrismObjectDealGrantGetResponseTeamGroupID = "w"
)

func (r PrismObjectDealGrantGetResponseTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectDealGrantGetResponseTeamGroupIDA, PrismObjectDealGrantGetResponseTeamGroupIDR, PrismObjectDealGrantGetResponseTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectDealGrantGetResponseTeamID string

const (
	PrismObjectDealGrantGetResponseTeamIDA PrismObjectDealGrantGetResponseTeamID = "a"
	PrismObjectDealGrantGetResponseTeamIDR PrismObjectDealGrantGetResponseTeamID = "r"
	PrismObjectDealGrantGetResponseTeamIDW PrismObjectDealGrantGetResponseTeamID = "w"
)

func (r PrismObjectDealGrantGetResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectDealGrantGetResponseTeamIDA, PrismObjectDealGrantGetResponseTeamIDR, PrismObjectDealGrantGetResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectDealGrantGetResponseUserID string

const (
	PrismObjectDealGrantGetResponseUserIDA PrismObjectDealGrantGetResponseUserID = "a"
	PrismObjectDealGrantGetResponseUserIDR PrismObjectDealGrantGetResponseUserID = "r"
	PrismObjectDealGrantGetResponseUserIDW PrismObjectDealGrantGetResponseUserID = "w"
)

func (r PrismObjectDealGrantGetResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectDealGrantGetResponseUserIDA, PrismObjectDealGrantGetResponseUserIDR, PrismObjectDealGrantGetResponseUserIDW:
		return true
	}
	return false
}

type PrismObjectDealGrantUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	PathTeamID     param.Field[string]                                                   `path:"teamId" api:"required" format:"uuid"`
	TeamGroupID    param.Field[[]map[string]PrismObjectDealGrantUpdateParamsTeamGroupID] `json:"team_group_id"`
	BodyTeamID     param.Field[map[string]PrismObjectDealGrantUpdateParamsTeamID]        `json:"team_id"`
	UserID         param.Field[[]map[string]PrismObjectDealGrantUpdateParamsUserID]      `json:"user_id"`
	IdempotencyKey param.Field[string]                                                   `header:"Idempotency-Key"`
}

func (r PrismObjectDealGrantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDealGrantUpdateParamsTeamGroupID string

const (
	PrismObjectDealGrantUpdateParamsTeamGroupIDA PrismObjectDealGrantUpdateParamsTeamGroupID = "a"
	PrismObjectDealGrantUpdateParamsTeamGroupIDR PrismObjectDealGrantUpdateParamsTeamGroupID = "r"
	PrismObjectDealGrantUpdateParamsTeamGroupIDW PrismObjectDealGrantUpdateParamsTeamGroupID = "w"
)

func (r PrismObjectDealGrantUpdateParamsTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectDealGrantUpdateParamsTeamGroupIDA, PrismObjectDealGrantUpdateParamsTeamGroupIDR, PrismObjectDealGrantUpdateParamsTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectDealGrantUpdateParamsTeamID string

const (
	PrismObjectDealGrantUpdateParamsTeamIDA PrismObjectDealGrantUpdateParamsTeamID = "a"
	PrismObjectDealGrantUpdateParamsTeamIDR PrismObjectDealGrantUpdateParamsTeamID = "r"
	PrismObjectDealGrantUpdateParamsTeamIDW PrismObjectDealGrantUpdateParamsTeamID = "w"
)

func (r PrismObjectDealGrantUpdateParamsTeamID) IsKnown() bool {
	switch r {
	case PrismObjectDealGrantUpdateParamsTeamIDA, PrismObjectDealGrantUpdateParamsTeamIDR, PrismObjectDealGrantUpdateParamsTeamIDW:
		return true
	}
	return false
}

type PrismObjectDealGrantUpdateParamsUserID string

const (
	PrismObjectDealGrantUpdateParamsUserIDA PrismObjectDealGrantUpdateParamsUserID = "a"
	PrismObjectDealGrantUpdateParamsUserIDR PrismObjectDealGrantUpdateParamsUserID = "r"
	PrismObjectDealGrantUpdateParamsUserIDW PrismObjectDealGrantUpdateParamsUserID = "w"
)

func (r PrismObjectDealGrantUpdateParamsUserID) IsKnown() bool {
	switch r {
	case PrismObjectDealGrantUpdateParamsUserIDA, PrismObjectDealGrantUpdateParamsUserIDR, PrismObjectDealGrantUpdateParamsUserIDW:
		return true
	}
	return false
}

type PrismObjectDealGrantGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
