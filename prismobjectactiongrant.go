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

// PrismObjectActionGrantService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectActionGrantService] method instead.
type PrismObjectActionGrantService struct {
	Options []option.RequestOption
}

// NewPrismObjectActionGrantService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectActionGrantService(opts ...option.RequestOption) (r *PrismObjectActionGrantService) {
	r = &PrismObjectActionGrantService{}
	r.Options = opts
	return
}

// Update grant
func (r *PrismObjectActionGrantService) Update(ctx context.Context, actionID string, params PrismObjectActionGrantUpdateParams, opts ...option.RequestOption) (res *PrismObjectActionGrantUpdateResponse, err error) {
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
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/grant/%s/action/%s", params.PathTeamID, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Get grant
func (r *PrismObjectActionGrantService) Get(ctx context.Context, actionID string, query PrismObjectActionGrantGetParams, opts ...option.RequestOption) (res *PrismObjectActionGrantGetResponse, err error) {
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
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/grant/%s/action/%s", query.TeamID, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PrismObjectActionGrantUpdateResponse struct {
	TeamGroupID []map[string]PrismObjectActionGrantUpdateResponseTeamGroupID `json:"team_group_id"`
	TeamID      map[string]PrismObjectActionGrantUpdateResponseTeamID        `json:"team_id"`
	UserID      []map[string]PrismObjectActionGrantUpdateResponseUserID      `json:"user_id"`
	JSON        prismObjectActionGrantUpdateResponseJSON                     `json:"-"`
}

// prismObjectActionGrantUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectActionGrantUpdateResponse]
type prismObjectActionGrantUpdateResponseJSON struct {
	TeamGroupID apijson.Field
	TeamID      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionGrantUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionGrantUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionGrantUpdateResponseTeamGroupID string

const (
	PrismObjectActionGrantUpdateResponseTeamGroupIDA PrismObjectActionGrantUpdateResponseTeamGroupID = "a"
	PrismObjectActionGrantUpdateResponseTeamGroupIDR PrismObjectActionGrantUpdateResponseTeamGroupID = "r"
	PrismObjectActionGrantUpdateResponseTeamGroupIDW PrismObjectActionGrantUpdateResponseTeamGroupID = "w"
)

func (r PrismObjectActionGrantUpdateResponseTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectActionGrantUpdateResponseTeamGroupIDA, PrismObjectActionGrantUpdateResponseTeamGroupIDR, PrismObjectActionGrantUpdateResponseTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectActionGrantUpdateResponseTeamID string

const (
	PrismObjectActionGrantUpdateResponseTeamIDA PrismObjectActionGrantUpdateResponseTeamID = "a"
	PrismObjectActionGrantUpdateResponseTeamIDR PrismObjectActionGrantUpdateResponseTeamID = "r"
	PrismObjectActionGrantUpdateResponseTeamIDW PrismObjectActionGrantUpdateResponseTeamID = "w"
)

func (r PrismObjectActionGrantUpdateResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectActionGrantUpdateResponseTeamIDA, PrismObjectActionGrantUpdateResponseTeamIDR, PrismObjectActionGrantUpdateResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectActionGrantUpdateResponseUserID string

const (
	PrismObjectActionGrantUpdateResponseUserIDA PrismObjectActionGrantUpdateResponseUserID = "a"
	PrismObjectActionGrantUpdateResponseUserIDR PrismObjectActionGrantUpdateResponseUserID = "r"
	PrismObjectActionGrantUpdateResponseUserIDW PrismObjectActionGrantUpdateResponseUserID = "w"
)

func (r PrismObjectActionGrantUpdateResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectActionGrantUpdateResponseUserIDA, PrismObjectActionGrantUpdateResponseUserIDR, PrismObjectActionGrantUpdateResponseUserIDW:
		return true
	}
	return false
}

type PrismObjectActionGrantGetResponse struct {
	TeamGroupID []map[string]PrismObjectActionGrantGetResponseTeamGroupID `json:"team_group_id"`
	TeamID      map[string]PrismObjectActionGrantGetResponseTeamID        `json:"team_id"`
	UserID      []map[string]PrismObjectActionGrantGetResponseUserID      `json:"user_id"`
	JSON        prismObjectActionGrantGetResponseJSON                     `json:"-"`
}

// prismObjectActionGrantGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectActionGrantGetResponse]
type prismObjectActionGrantGetResponseJSON struct {
	TeamGroupID apijson.Field
	TeamID      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectActionGrantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectActionGrantGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectActionGrantGetResponseTeamGroupID string

const (
	PrismObjectActionGrantGetResponseTeamGroupIDA PrismObjectActionGrantGetResponseTeamGroupID = "a"
	PrismObjectActionGrantGetResponseTeamGroupIDR PrismObjectActionGrantGetResponseTeamGroupID = "r"
	PrismObjectActionGrantGetResponseTeamGroupIDW PrismObjectActionGrantGetResponseTeamGroupID = "w"
)

func (r PrismObjectActionGrantGetResponseTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectActionGrantGetResponseTeamGroupIDA, PrismObjectActionGrantGetResponseTeamGroupIDR, PrismObjectActionGrantGetResponseTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectActionGrantGetResponseTeamID string

const (
	PrismObjectActionGrantGetResponseTeamIDA PrismObjectActionGrantGetResponseTeamID = "a"
	PrismObjectActionGrantGetResponseTeamIDR PrismObjectActionGrantGetResponseTeamID = "r"
	PrismObjectActionGrantGetResponseTeamIDW PrismObjectActionGrantGetResponseTeamID = "w"
)

func (r PrismObjectActionGrantGetResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectActionGrantGetResponseTeamIDA, PrismObjectActionGrantGetResponseTeamIDR, PrismObjectActionGrantGetResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectActionGrantGetResponseUserID string

const (
	PrismObjectActionGrantGetResponseUserIDA PrismObjectActionGrantGetResponseUserID = "a"
	PrismObjectActionGrantGetResponseUserIDR PrismObjectActionGrantGetResponseUserID = "r"
	PrismObjectActionGrantGetResponseUserIDW PrismObjectActionGrantGetResponseUserID = "w"
)

func (r PrismObjectActionGrantGetResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectActionGrantGetResponseUserIDA, PrismObjectActionGrantGetResponseUserIDR, PrismObjectActionGrantGetResponseUserIDW:
		return true
	}
	return false
}

type PrismObjectActionGrantUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	PathTeamID  param.Field[string]                                                     `path:"teamId" api:"required" format:"uuid"`
	TeamGroupID param.Field[[]map[string]PrismObjectActionGrantUpdateParamsTeamGroupID] `json:"team_group_id"`
	BodyTeamID  param.Field[map[string]PrismObjectActionGrantUpdateParamsTeamID]        `json:"team_id"`
	UserID      param.Field[[]map[string]PrismObjectActionGrantUpdateParamsUserID]      `json:"user_id"`
}

func (r PrismObjectActionGrantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectActionGrantUpdateParamsTeamGroupID string

const (
	PrismObjectActionGrantUpdateParamsTeamGroupIDA PrismObjectActionGrantUpdateParamsTeamGroupID = "a"
	PrismObjectActionGrantUpdateParamsTeamGroupIDR PrismObjectActionGrantUpdateParamsTeamGroupID = "r"
	PrismObjectActionGrantUpdateParamsTeamGroupIDW PrismObjectActionGrantUpdateParamsTeamGroupID = "w"
)

func (r PrismObjectActionGrantUpdateParamsTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectActionGrantUpdateParamsTeamGroupIDA, PrismObjectActionGrantUpdateParamsTeamGroupIDR, PrismObjectActionGrantUpdateParamsTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectActionGrantUpdateParamsTeamID string

const (
	PrismObjectActionGrantUpdateParamsTeamIDA PrismObjectActionGrantUpdateParamsTeamID = "a"
	PrismObjectActionGrantUpdateParamsTeamIDR PrismObjectActionGrantUpdateParamsTeamID = "r"
	PrismObjectActionGrantUpdateParamsTeamIDW PrismObjectActionGrantUpdateParamsTeamID = "w"
)

func (r PrismObjectActionGrantUpdateParamsTeamID) IsKnown() bool {
	switch r {
	case PrismObjectActionGrantUpdateParamsTeamIDA, PrismObjectActionGrantUpdateParamsTeamIDR, PrismObjectActionGrantUpdateParamsTeamIDW:
		return true
	}
	return false
}

type PrismObjectActionGrantUpdateParamsUserID string

const (
	PrismObjectActionGrantUpdateParamsUserIDA PrismObjectActionGrantUpdateParamsUserID = "a"
	PrismObjectActionGrantUpdateParamsUserIDR PrismObjectActionGrantUpdateParamsUserID = "r"
	PrismObjectActionGrantUpdateParamsUserIDW PrismObjectActionGrantUpdateParamsUserID = "w"
)

func (r PrismObjectActionGrantUpdateParamsUserID) IsKnown() bool {
	switch r {
	case PrismObjectActionGrantUpdateParamsUserIDA, PrismObjectActionGrantUpdateParamsUserIDR, PrismObjectActionGrantUpdateParamsUserIDW:
		return true
	}
	return false
}

type PrismObjectActionGrantGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
