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

// PrismGrantService contains methods and other services that help with interacting
// with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismGrantService] method instead.
type PrismGrantService struct {
	Options []option.RequestOption
}

// NewPrismGrantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrismGrantService(opts ...option.RequestOption) (r *PrismGrantService) {
	r = &PrismGrantService{}
	r.Options = opts
	return
}

// Get grant
func (r *PrismGrantService) GetGrant(ctx context.Context, objectType ObjectType, objectID string, query PrismGrantGetGrantParams, opts ...option.RequestOption) (res *PrismGrantGetGrantResponse, err error) {
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
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/grant/%s/%v/%s", query.TeamID, objectType, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update grant
func (r *PrismGrantService) UpdateGrant(ctx context.Context, objectType ObjectType, objectID string, params PrismGrantUpdateGrantParams, opts ...option.RequestOption) (res *PrismGrantUpdateGrantResponse, err error) {
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
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/grant/%s/%v/%s", params.PathTeamID, objectType, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type PrismGrantGetGrantResponse struct {
	TeamGroupID []map[string]PrismGrantGetGrantResponseTeamGroupID `json:"team_group_id"`
	TeamID      map[string]PrismGrantGetGrantResponseTeamID        `json:"team_id"`
	UserID      []map[string]PrismGrantGetGrantResponseUserID      `json:"user_id"`
	JSON        prismGrantGetGrantResponseJSON                     `json:"-"`
}

// prismGrantGetGrantResponseJSON contains the JSON metadata for the struct
// [PrismGrantGetGrantResponse]
type prismGrantGetGrantResponseJSON struct {
	TeamGroupID apijson.Field
	TeamID      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismGrantGetGrantResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismGrantGetGrantResponseJSON) RawJSON() string {
	return r.raw
}

type PrismGrantGetGrantResponseTeamGroupID string

const (
	PrismGrantGetGrantResponseTeamGroupIDA PrismGrantGetGrantResponseTeamGroupID = "a"
	PrismGrantGetGrantResponseTeamGroupIDR PrismGrantGetGrantResponseTeamGroupID = "r"
	PrismGrantGetGrantResponseTeamGroupIDW PrismGrantGetGrantResponseTeamGroupID = "w"
)

func (r PrismGrantGetGrantResponseTeamGroupID) IsKnown() bool {
	switch r {
	case PrismGrantGetGrantResponseTeamGroupIDA, PrismGrantGetGrantResponseTeamGroupIDR, PrismGrantGetGrantResponseTeamGroupIDW:
		return true
	}
	return false
}

type PrismGrantGetGrantResponseTeamID string

const (
	PrismGrantGetGrantResponseTeamIDA PrismGrantGetGrantResponseTeamID = "a"
	PrismGrantGetGrantResponseTeamIDR PrismGrantGetGrantResponseTeamID = "r"
	PrismGrantGetGrantResponseTeamIDW PrismGrantGetGrantResponseTeamID = "w"
)

func (r PrismGrantGetGrantResponseTeamID) IsKnown() bool {
	switch r {
	case PrismGrantGetGrantResponseTeamIDA, PrismGrantGetGrantResponseTeamIDR, PrismGrantGetGrantResponseTeamIDW:
		return true
	}
	return false
}

type PrismGrantGetGrantResponseUserID string

const (
	PrismGrantGetGrantResponseUserIDA PrismGrantGetGrantResponseUserID = "a"
	PrismGrantGetGrantResponseUserIDR PrismGrantGetGrantResponseUserID = "r"
	PrismGrantGetGrantResponseUserIDW PrismGrantGetGrantResponseUserID = "w"
)

func (r PrismGrantGetGrantResponseUserID) IsKnown() bool {
	switch r {
	case PrismGrantGetGrantResponseUserIDA, PrismGrantGetGrantResponseUserIDR, PrismGrantGetGrantResponseUserIDW:
		return true
	}
	return false
}

type PrismGrantUpdateGrantResponse struct {
	TeamGroupID []map[string]PrismGrantUpdateGrantResponseTeamGroupID `json:"team_group_id"`
	TeamID      map[string]PrismGrantUpdateGrantResponseTeamID        `json:"team_id"`
	UserID      []map[string]PrismGrantUpdateGrantResponseUserID      `json:"user_id"`
	JSON        prismGrantUpdateGrantResponseJSON                     `json:"-"`
}

// prismGrantUpdateGrantResponseJSON contains the JSON metadata for the struct
// [PrismGrantUpdateGrantResponse]
type prismGrantUpdateGrantResponseJSON struct {
	TeamGroupID apijson.Field
	TeamID      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismGrantUpdateGrantResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismGrantUpdateGrantResponseJSON) RawJSON() string {
	return r.raw
}

type PrismGrantUpdateGrantResponseTeamGroupID string

const (
	PrismGrantUpdateGrantResponseTeamGroupIDA PrismGrantUpdateGrantResponseTeamGroupID = "a"
	PrismGrantUpdateGrantResponseTeamGroupIDR PrismGrantUpdateGrantResponseTeamGroupID = "r"
	PrismGrantUpdateGrantResponseTeamGroupIDW PrismGrantUpdateGrantResponseTeamGroupID = "w"
)

func (r PrismGrantUpdateGrantResponseTeamGroupID) IsKnown() bool {
	switch r {
	case PrismGrantUpdateGrantResponseTeamGroupIDA, PrismGrantUpdateGrantResponseTeamGroupIDR, PrismGrantUpdateGrantResponseTeamGroupIDW:
		return true
	}
	return false
}

type PrismGrantUpdateGrantResponseTeamID string

const (
	PrismGrantUpdateGrantResponseTeamIDA PrismGrantUpdateGrantResponseTeamID = "a"
	PrismGrantUpdateGrantResponseTeamIDR PrismGrantUpdateGrantResponseTeamID = "r"
	PrismGrantUpdateGrantResponseTeamIDW PrismGrantUpdateGrantResponseTeamID = "w"
)

func (r PrismGrantUpdateGrantResponseTeamID) IsKnown() bool {
	switch r {
	case PrismGrantUpdateGrantResponseTeamIDA, PrismGrantUpdateGrantResponseTeamIDR, PrismGrantUpdateGrantResponseTeamIDW:
		return true
	}
	return false
}

type PrismGrantUpdateGrantResponseUserID string

const (
	PrismGrantUpdateGrantResponseUserIDA PrismGrantUpdateGrantResponseUserID = "a"
	PrismGrantUpdateGrantResponseUserIDR PrismGrantUpdateGrantResponseUserID = "r"
	PrismGrantUpdateGrantResponseUserIDW PrismGrantUpdateGrantResponseUserID = "w"
)

func (r PrismGrantUpdateGrantResponseUserID) IsKnown() bool {
	switch r {
	case PrismGrantUpdateGrantResponseUserIDA, PrismGrantUpdateGrantResponseUserIDR, PrismGrantUpdateGrantResponseUserIDW:
		return true
	}
	return false
}

type PrismGrantGetGrantParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismGrantUpdateGrantParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	PathTeamID  param.Field[string]                                              `path:"teamId" api:"required" format:"uuid"`
	TeamGroupID param.Field[[]map[string]PrismGrantUpdateGrantParamsTeamGroupID] `json:"team_group_id"`
	BodyTeamID  param.Field[map[string]PrismGrantUpdateGrantParamsTeamID]        `json:"team_id"`
	UserID      param.Field[[]map[string]PrismGrantUpdateGrantParamsUserID]      `json:"user_id"`
}

func (r PrismGrantUpdateGrantParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismGrantUpdateGrantParamsTeamGroupID string

const (
	PrismGrantUpdateGrantParamsTeamGroupIDA PrismGrantUpdateGrantParamsTeamGroupID = "a"
	PrismGrantUpdateGrantParamsTeamGroupIDR PrismGrantUpdateGrantParamsTeamGroupID = "r"
	PrismGrantUpdateGrantParamsTeamGroupIDW PrismGrantUpdateGrantParamsTeamGroupID = "w"
)

func (r PrismGrantUpdateGrantParamsTeamGroupID) IsKnown() bool {
	switch r {
	case PrismGrantUpdateGrantParamsTeamGroupIDA, PrismGrantUpdateGrantParamsTeamGroupIDR, PrismGrantUpdateGrantParamsTeamGroupIDW:
		return true
	}
	return false
}

type PrismGrantUpdateGrantParamsTeamID string

const (
	PrismGrantUpdateGrantParamsTeamIDA PrismGrantUpdateGrantParamsTeamID = "a"
	PrismGrantUpdateGrantParamsTeamIDR PrismGrantUpdateGrantParamsTeamID = "r"
	PrismGrantUpdateGrantParamsTeamIDW PrismGrantUpdateGrantParamsTeamID = "w"
)

func (r PrismGrantUpdateGrantParamsTeamID) IsKnown() bool {
	switch r {
	case PrismGrantUpdateGrantParamsTeamIDA, PrismGrantUpdateGrantParamsTeamIDR, PrismGrantUpdateGrantParamsTeamIDW:
		return true
	}
	return false
}

type PrismGrantUpdateGrantParamsUserID string

const (
	PrismGrantUpdateGrantParamsUserIDA PrismGrantUpdateGrantParamsUserID = "a"
	PrismGrantUpdateGrantParamsUserIDR PrismGrantUpdateGrantParamsUserID = "r"
	PrismGrantUpdateGrantParamsUserIDW PrismGrantUpdateGrantParamsUserID = "w"
)

func (r PrismGrantUpdateGrantParamsUserID) IsKnown() bool {
	switch r {
	case PrismGrantUpdateGrantParamsUserIDA, PrismGrantUpdateGrantParamsUserIDR, PrismGrantUpdateGrantParamsUserIDW:
		return true
	}
	return false
}
