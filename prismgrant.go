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
func (r *PrismGrantService) GetGrant(ctx context.Context, objectType ObjectType, objectID string, query PrismGrantGetGrantParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return err
	}
	requestconfig.UseDefaultParam(&query.TeamID, precfg.TeamID)
	if query.TeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/grant/%s/%v/%s", query.TeamID, objectType, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Update grant
func (r *PrismGrantService) UpdateGrant(ctx context.Context, objectType ObjectType, objectID string, params PrismGrantUpdateGrantParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return err
	}
	requestconfig.UseDefaultParam(&params.PathTeamID, precfg.TeamID)
	if params.PathTeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/grant/%s/%v/%s", params.PathTeamID, objectType, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return err
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
