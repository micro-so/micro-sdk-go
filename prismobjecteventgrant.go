// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/micro-go/internal/apijson"
	"github.com/stainless-sdks/micro-go/internal/param"
	"github.com/stainless-sdks/micro-go/internal/requestconfig"
	"github.com/stainless-sdks/micro-go/option"
)

// PrismObjectEventGrantService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectEventGrantService] method instead.
type PrismObjectEventGrantService struct {
	Options []option.RequestOption
}

// NewPrismObjectEventGrantService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectEventGrantService(opts ...option.RequestOption) (r *PrismObjectEventGrantService) {
	r = &PrismObjectEventGrantService{}
	r.Options = opts
	return
}

// Update grant
func (r *PrismObjectEventGrantService) Update(ctx context.Context, eventID string, params PrismObjectEventGrantUpdateParams, opts ...option.RequestOption) (res *PrismObjectEventGrantUpdateResponse, err error) {
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
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/grant/%s/event/%s", params.PathTeamID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Get grant
func (r *PrismObjectEventGrantService) Get(ctx context.Context, eventID string, query PrismObjectEventGrantGetParams, opts ...option.RequestOption) (res *PrismObjectEventGrantGetResponse, err error) {
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
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/grant/%s/event/%s", query.TeamID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PrismObjectEventGrantUpdateResponse struct {
	TeamGroupID []map[string]PrismObjectEventGrantUpdateResponseTeamGroupID `json:"team_group_id"`
	TeamID      map[string]PrismObjectEventGrantUpdateResponseTeamID        `json:"team_id"`
	UserID      []map[string]PrismObjectEventGrantUpdateResponseUserID      `json:"user_id"`
	JSON        prismObjectEventGrantUpdateResponseJSON                     `json:"-"`
}

// prismObjectEventGrantUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectEventGrantUpdateResponse]
type prismObjectEventGrantUpdateResponseJSON struct {
	TeamGroupID apijson.Field
	TeamID      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEventGrantUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEventGrantUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEventGrantUpdateResponseTeamGroupID string

const (
	PrismObjectEventGrantUpdateResponseTeamGroupIDA PrismObjectEventGrantUpdateResponseTeamGroupID = "a"
	PrismObjectEventGrantUpdateResponseTeamGroupIDR PrismObjectEventGrantUpdateResponseTeamGroupID = "r"
	PrismObjectEventGrantUpdateResponseTeamGroupIDW PrismObjectEventGrantUpdateResponseTeamGroupID = "w"
)

func (r PrismObjectEventGrantUpdateResponseTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantUpdateResponseTeamGroupIDA, PrismObjectEventGrantUpdateResponseTeamGroupIDR, PrismObjectEventGrantUpdateResponseTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectEventGrantUpdateResponseTeamID string

const (
	PrismObjectEventGrantUpdateResponseTeamIDA PrismObjectEventGrantUpdateResponseTeamID = "a"
	PrismObjectEventGrantUpdateResponseTeamIDR PrismObjectEventGrantUpdateResponseTeamID = "r"
	PrismObjectEventGrantUpdateResponseTeamIDW PrismObjectEventGrantUpdateResponseTeamID = "w"
)

func (r PrismObjectEventGrantUpdateResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantUpdateResponseTeamIDA, PrismObjectEventGrantUpdateResponseTeamIDR, PrismObjectEventGrantUpdateResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectEventGrantUpdateResponseUserID string

const (
	PrismObjectEventGrantUpdateResponseUserIDA PrismObjectEventGrantUpdateResponseUserID = "a"
	PrismObjectEventGrantUpdateResponseUserIDR PrismObjectEventGrantUpdateResponseUserID = "r"
	PrismObjectEventGrantUpdateResponseUserIDW PrismObjectEventGrantUpdateResponseUserID = "w"
)

func (r PrismObjectEventGrantUpdateResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantUpdateResponseUserIDA, PrismObjectEventGrantUpdateResponseUserIDR, PrismObjectEventGrantUpdateResponseUserIDW:
		return true
	}
	return false
}

type PrismObjectEventGrantGetResponse struct {
	TeamGroupID []map[string]PrismObjectEventGrantGetResponseTeamGroupID `json:"team_group_id"`
	TeamID      map[string]PrismObjectEventGrantGetResponseTeamID        `json:"team_id"`
	UserID      []map[string]PrismObjectEventGrantGetResponseUserID      `json:"user_id"`
	JSON        prismObjectEventGrantGetResponseJSON                     `json:"-"`
}

// prismObjectEventGrantGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectEventGrantGetResponse]
type prismObjectEventGrantGetResponseJSON struct {
	TeamGroupID apijson.Field
	TeamID      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEventGrantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEventGrantGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEventGrantGetResponseTeamGroupID string

const (
	PrismObjectEventGrantGetResponseTeamGroupIDA PrismObjectEventGrantGetResponseTeamGroupID = "a"
	PrismObjectEventGrantGetResponseTeamGroupIDR PrismObjectEventGrantGetResponseTeamGroupID = "r"
	PrismObjectEventGrantGetResponseTeamGroupIDW PrismObjectEventGrantGetResponseTeamGroupID = "w"
)

func (r PrismObjectEventGrantGetResponseTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantGetResponseTeamGroupIDA, PrismObjectEventGrantGetResponseTeamGroupIDR, PrismObjectEventGrantGetResponseTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectEventGrantGetResponseTeamID string

const (
	PrismObjectEventGrantGetResponseTeamIDA PrismObjectEventGrantGetResponseTeamID = "a"
	PrismObjectEventGrantGetResponseTeamIDR PrismObjectEventGrantGetResponseTeamID = "r"
	PrismObjectEventGrantGetResponseTeamIDW PrismObjectEventGrantGetResponseTeamID = "w"
)

func (r PrismObjectEventGrantGetResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantGetResponseTeamIDA, PrismObjectEventGrantGetResponseTeamIDR, PrismObjectEventGrantGetResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectEventGrantGetResponseUserID string

const (
	PrismObjectEventGrantGetResponseUserIDA PrismObjectEventGrantGetResponseUserID = "a"
	PrismObjectEventGrantGetResponseUserIDR PrismObjectEventGrantGetResponseUserID = "r"
	PrismObjectEventGrantGetResponseUserIDW PrismObjectEventGrantGetResponseUserID = "w"
)

func (r PrismObjectEventGrantGetResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantGetResponseUserIDA, PrismObjectEventGrantGetResponseUserIDR, PrismObjectEventGrantGetResponseUserIDW:
		return true
	}
	return false
}

type PrismObjectEventGrantUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	PathTeamID  param.Field[string]                                                    `path:"teamId" api:"required" format:"uuid"`
	TeamGroupID param.Field[[]map[string]PrismObjectEventGrantUpdateParamsTeamGroupID] `json:"team_group_id"`
	BodyTeamID  param.Field[map[string]PrismObjectEventGrantUpdateParamsTeamID]        `json:"team_id"`
	UserID      param.Field[[]map[string]PrismObjectEventGrantUpdateParamsUserID]      `json:"user_id"`
}

func (r PrismObjectEventGrantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectEventGrantUpdateParamsTeamGroupID string

const (
	PrismObjectEventGrantUpdateParamsTeamGroupIDA PrismObjectEventGrantUpdateParamsTeamGroupID = "a"
	PrismObjectEventGrantUpdateParamsTeamGroupIDR PrismObjectEventGrantUpdateParamsTeamGroupID = "r"
	PrismObjectEventGrantUpdateParamsTeamGroupIDW PrismObjectEventGrantUpdateParamsTeamGroupID = "w"
)

func (r PrismObjectEventGrantUpdateParamsTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantUpdateParamsTeamGroupIDA, PrismObjectEventGrantUpdateParamsTeamGroupIDR, PrismObjectEventGrantUpdateParamsTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectEventGrantUpdateParamsTeamID string

const (
	PrismObjectEventGrantUpdateParamsTeamIDA PrismObjectEventGrantUpdateParamsTeamID = "a"
	PrismObjectEventGrantUpdateParamsTeamIDR PrismObjectEventGrantUpdateParamsTeamID = "r"
	PrismObjectEventGrantUpdateParamsTeamIDW PrismObjectEventGrantUpdateParamsTeamID = "w"
)

func (r PrismObjectEventGrantUpdateParamsTeamID) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantUpdateParamsTeamIDA, PrismObjectEventGrantUpdateParamsTeamIDR, PrismObjectEventGrantUpdateParamsTeamIDW:
		return true
	}
	return false
}

type PrismObjectEventGrantUpdateParamsUserID string

const (
	PrismObjectEventGrantUpdateParamsUserIDA PrismObjectEventGrantUpdateParamsUserID = "a"
	PrismObjectEventGrantUpdateParamsUserIDR PrismObjectEventGrantUpdateParamsUserID = "r"
	PrismObjectEventGrantUpdateParamsUserIDW PrismObjectEventGrantUpdateParamsUserID = "w"
)

func (r PrismObjectEventGrantUpdateParamsUserID) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantUpdateParamsUserIDA, PrismObjectEventGrantUpdateParamsUserIDR, PrismObjectEventGrantUpdateParamsUserIDW:
		return true
	}
	return false
}

type PrismObjectEventGrantGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
