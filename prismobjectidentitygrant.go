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

// PrismObjectIdentityGrantService contains methods and other services that help
// with interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectIdentityGrantService] method instead.
type PrismObjectIdentityGrantService struct {
	Options []option.RequestOption
}

// NewPrismObjectIdentityGrantService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPrismObjectIdentityGrantService(opts ...option.RequestOption) (r *PrismObjectIdentityGrantService) {
	r = &PrismObjectIdentityGrantService{}
	r.Options = opts
	return
}

// Update grant
func (r *PrismObjectIdentityGrantService) Update(ctx context.Context, identityID string, params PrismObjectIdentityGrantUpdateParams, opts ...option.RequestOption) (res *PrismObjectIdentityGrantUpdateResponse, err error) {
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
	if identityID == "" {
		err = errors.New("missing required identityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/grant/%s/identity/%s", params.PathTeamID, identityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Get grant
func (r *PrismObjectIdentityGrantService) Get(ctx context.Context, identityID string, query PrismObjectIdentityGrantGetParams, opts ...option.RequestOption) (res *PrismObjectIdentityGrantGetResponse, err error) {
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
	if identityID == "" {
		err = errors.New("missing required identityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/grant/%s/identity/%s", query.TeamID, identityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PrismObjectIdentityGrantUpdateResponse struct {
	TeamGroupID []map[string]PrismObjectIdentityGrantUpdateResponseTeamGroupID `json:"team_group_id"`
	TeamID      map[string]PrismObjectIdentityGrantUpdateResponseTeamID        `json:"team_id"`
	UserID      []map[string]PrismObjectIdentityGrantUpdateResponseUserID      `json:"user_id"`
	JSON        prismObjectIdentityGrantUpdateResponseJSON                     `json:"-"`
}

// prismObjectIdentityGrantUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectIdentityGrantUpdateResponse]
type prismObjectIdentityGrantUpdateResponseJSON struct {
	TeamGroupID apijson.Field
	TeamID      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityGrantUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityGrantUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityGrantUpdateResponseTeamGroupID string

const (
	PrismObjectIdentityGrantUpdateResponseTeamGroupIDA PrismObjectIdentityGrantUpdateResponseTeamGroupID = "a"
	PrismObjectIdentityGrantUpdateResponseTeamGroupIDR PrismObjectIdentityGrantUpdateResponseTeamGroupID = "r"
	PrismObjectIdentityGrantUpdateResponseTeamGroupIDW PrismObjectIdentityGrantUpdateResponseTeamGroupID = "w"
)

func (r PrismObjectIdentityGrantUpdateResponseTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectIdentityGrantUpdateResponseTeamGroupIDA, PrismObjectIdentityGrantUpdateResponseTeamGroupIDR, PrismObjectIdentityGrantUpdateResponseTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectIdentityGrantUpdateResponseTeamID string

const (
	PrismObjectIdentityGrantUpdateResponseTeamIDA PrismObjectIdentityGrantUpdateResponseTeamID = "a"
	PrismObjectIdentityGrantUpdateResponseTeamIDR PrismObjectIdentityGrantUpdateResponseTeamID = "r"
	PrismObjectIdentityGrantUpdateResponseTeamIDW PrismObjectIdentityGrantUpdateResponseTeamID = "w"
)

func (r PrismObjectIdentityGrantUpdateResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectIdentityGrantUpdateResponseTeamIDA, PrismObjectIdentityGrantUpdateResponseTeamIDR, PrismObjectIdentityGrantUpdateResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectIdentityGrantUpdateResponseUserID string

const (
	PrismObjectIdentityGrantUpdateResponseUserIDA PrismObjectIdentityGrantUpdateResponseUserID = "a"
	PrismObjectIdentityGrantUpdateResponseUserIDR PrismObjectIdentityGrantUpdateResponseUserID = "r"
	PrismObjectIdentityGrantUpdateResponseUserIDW PrismObjectIdentityGrantUpdateResponseUserID = "w"
)

func (r PrismObjectIdentityGrantUpdateResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectIdentityGrantUpdateResponseUserIDA, PrismObjectIdentityGrantUpdateResponseUserIDR, PrismObjectIdentityGrantUpdateResponseUserIDW:
		return true
	}
	return false
}

type PrismObjectIdentityGrantGetResponse struct {
	TeamGroupID []map[string]PrismObjectIdentityGrantGetResponseTeamGroupID `json:"team_group_id"`
	TeamID      map[string]PrismObjectIdentityGrantGetResponseTeamID        `json:"team_id"`
	UserID      []map[string]PrismObjectIdentityGrantGetResponseUserID      `json:"user_id"`
	JSON        prismObjectIdentityGrantGetResponseJSON                     `json:"-"`
}

// prismObjectIdentityGrantGetResponseJSON contains the JSON metadata for the
// struct [PrismObjectIdentityGrantGetResponse]
type prismObjectIdentityGrantGetResponseJSON struct {
	TeamGroupID apijson.Field
	TeamID      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectIdentityGrantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectIdentityGrantGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectIdentityGrantGetResponseTeamGroupID string

const (
	PrismObjectIdentityGrantGetResponseTeamGroupIDA PrismObjectIdentityGrantGetResponseTeamGroupID = "a"
	PrismObjectIdentityGrantGetResponseTeamGroupIDR PrismObjectIdentityGrantGetResponseTeamGroupID = "r"
	PrismObjectIdentityGrantGetResponseTeamGroupIDW PrismObjectIdentityGrantGetResponseTeamGroupID = "w"
)

func (r PrismObjectIdentityGrantGetResponseTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectIdentityGrantGetResponseTeamGroupIDA, PrismObjectIdentityGrantGetResponseTeamGroupIDR, PrismObjectIdentityGrantGetResponseTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectIdentityGrantGetResponseTeamID string

const (
	PrismObjectIdentityGrantGetResponseTeamIDA PrismObjectIdentityGrantGetResponseTeamID = "a"
	PrismObjectIdentityGrantGetResponseTeamIDR PrismObjectIdentityGrantGetResponseTeamID = "r"
	PrismObjectIdentityGrantGetResponseTeamIDW PrismObjectIdentityGrantGetResponseTeamID = "w"
)

func (r PrismObjectIdentityGrantGetResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectIdentityGrantGetResponseTeamIDA, PrismObjectIdentityGrantGetResponseTeamIDR, PrismObjectIdentityGrantGetResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectIdentityGrantGetResponseUserID string

const (
	PrismObjectIdentityGrantGetResponseUserIDA PrismObjectIdentityGrantGetResponseUserID = "a"
	PrismObjectIdentityGrantGetResponseUserIDR PrismObjectIdentityGrantGetResponseUserID = "r"
	PrismObjectIdentityGrantGetResponseUserIDW PrismObjectIdentityGrantGetResponseUserID = "w"
)

func (r PrismObjectIdentityGrantGetResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectIdentityGrantGetResponseUserIDA, PrismObjectIdentityGrantGetResponseUserIDR, PrismObjectIdentityGrantGetResponseUserIDW:
		return true
	}
	return false
}

type PrismObjectIdentityGrantUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	PathTeamID  param.Field[string]                                                       `path:"teamId" api:"required" format:"uuid"`
	TeamGroupID param.Field[[]map[string]PrismObjectIdentityGrantUpdateParamsTeamGroupID] `json:"team_group_id"`
	BodyTeamID  param.Field[map[string]PrismObjectIdentityGrantUpdateParamsTeamID]        `json:"team_id"`
	UserID      param.Field[[]map[string]PrismObjectIdentityGrantUpdateParamsUserID]      `json:"user_id"`
}

func (r PrismObjectIdentityGrantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectIdentityGrantUpdateParamsTeamGroupID string

const (
	PrismObjectIdentityGrantUpdateParamsTeamGroupIDA PrismObjectIdentityGrantUpdateParamsTeamGroupID = "a"
	PrismObjectIdentityGrantUpdateParamsTeamGroupIDR PrismObjectIdentityGrantUpdateParamsTeamGroupID = "r"
	PrismObjectIdentityGrantUpdateParamsTeamGroupIDW PrismObjectIdentityGrantUpdateParamsTeamGroupID = "w"
)

func (r PrismObjectIdentityGrantUpdateParamsTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectIdentityGrantUpdateParamsTeamGroupIDA, PrismObjectIdentityGrantUpdateParamsTeamGroupIDR, PrismObjectIdentityGrantUpdateParamsTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectIdentityGrantUpdateParamsTeamID string

const (
	PrismObjectIdentityGrantUpdateParamsTeamIDA PrismObjectIdentityGrantUpdateParamsTeamID = "a"
	PrismObjectIdentityGrantUpdateParamsTeamIDR PrismObjectIdentityGrantUpdateParamsTeamID = "r"
	PrismObjectIdentityGrantUpdateParamsTeamIDW PrismObjectIdentityGrantUpdateParamsTeamID = "w"
)

func (r PrismObjectIdentityGrantUpdateParamsTeamID) IsKnown() bool {
	switch r {
	case PrismObjectIdentityGrantUpdateParamsTeamIDA, PrismObjectIdentityGrantUpdateParamsTeamIDR, PrismObjectIdentityGrantUpdateParamsTeamIDW:
		return true
	}
	return false
}

type PrismObjectIdentityGrantUpdateParamsUserID string

const (
	PrismObjectIdentityGrantUpdateParamsUserIDA PrismObjectIdentityGrantUpdateParamsUserID = "a"
	PrismObjectIdentityGrantUpdateParamsUserIDR PrismObjectIdentityGrantUpdateParamsUserID = "r"
	PrismObjectIdentityGrantUpdateParamsUserIDW PrismObjectIdentityGrantUpdateParamsUserID = "w"
)

func (r PrismObjectIdentityGrantUpdateParamsUserID) IsKnown() bool {
	switch r {
	case PrismObjectIdentityGrantUpdateParamsUserIDA, PrismObjectIdentityGrantUpdateParamsUserIDR, PrismObjectIdentityGrantUpdateParamsUserIDW:
		return true
	}
	return false
}

type PrismObjectIdentityGrantGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
