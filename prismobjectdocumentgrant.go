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

// PrismObjectDocumentGrantService contains methods and other services that help
// with interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectDocumentGrantService] method instead.
type PrismObjectDocumentGrantService struct {
	Options []option.RequestOption
}

// NewPrismObjectDocumentGrantService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPrismObjectDocumentGrantService(opts ...option.RequestOption) (r *PrismObjectDocumentGrantService) {
	r = &PrismObjectDocumentGrantService{}
	r.Options = opts
	return
}

// Update grant
func (r *PrismObjectDocumentGrantService) Update(ctx context.Context, documentID string, params PrismObjectDocumentGrantUpdateParams, opts ...option.RequestOption) (res *PrismObjectDocumentGrantUpdateResponse, err error) {
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
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/grant/%s/document/%s", params.PathTeamID, documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Get grant
func (r *PrismObjectDocumentGrantService) Get(ctx context.Context, documentID string, query PrismObjectDocumentGrantGetParams, opts ...option.RequestOption) (res *PrismObjectDocumentGrantGetResponse, err error) {
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
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/grant/%s/document/%s", query.TeamID, documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type PrismObjectDocumentGrantUpdateResponse struct {
	TeamGroupID []map[string]PrismObjectDocumentGrantUpdateResponseTeamGroupID `json:"team_group_id"`
	TeamID      map[string]PrismObjectDocumentGrantUpdateResponseTeamID        `json:"team_id"`
	UserID      []map[string]PrismObjectDocumentGrantUpdateResponseUserID      `json:"user_id"`
	JSON        prismObjectDocumentGrantUpdateResponseJSON                     `json:"-"`
}

// prismObjectDocumentGrantUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectDocumentGrantUpdateResponse]
type prismObjectDocumentGrantUpdateResponseJSON struct {
	TeamGroupID apijson.Field
	TeamID      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentGrantUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentGrantUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentGrantUpdateResponseTeamGroupID string

const (
	PrismObjectDocumentGrantUpdateResponseTeamGroupIDA PrismObjectDocumentGrantUpdateResponseTeamGroupID = "a"
	PrismObjectDocumentGrantUpdateResponseTeamGroupIDR PrismObjectDocumentGrantUpdateResponseTeamGroupID = "r"
	PrismObjectDocumentGrantUpdateResponseTeamGroupIDW PrismObjectDocumentGrantUpdateResponseTeamGroupID = "w"
)

func (r PrismObjectDocumentGrantUpdateResponseTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectDocumentGrantUpdateResponseTeamGroupIDA, PrismObjectDocumentGrantUpdateResponseTeamGroupIDR, PrismObjectDocumentGrantUpdateResponseTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectDocumentGrantUpdateResponseTeamID string

const (
	PrismObjectDocumentGrantUpdateResponseTeamIDA PrismObjectDocumentGrantUpdateResponseTeamID = "a"
	PrismObjectDocumentGrantUpdateResponseTeamIDR PrismObjectDocumentGrantUpdateResponseTeamID = "r"
	PrismObjectDocumentGrantUpdateResponseTeamIDW PrismObjectDocumentGrantUpdateResponseTeamID = "w"
)

func (r PrismObjectDocumentGrantUpdateResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectDocumentGrantUpdateResponseTeamIDA, PrismObjectDocumentGrantUpdateResponseTeamIDR, PrismObjectDocumentGrantUpdateResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectDocumentGrantUpdateResponseUserID string

const (
	PrismObjectDocumentGrantUpdateResponseUserIDA PrismObjectDocumentGrantUpdateResponseUserID = "a"
	PrismObjectDocumentGrantUpdateResponseUserIDR PrismObjectDocumentGrantUpdateResponseUserID = "r"
	PrismObjectDocumentGrantUpdateResponseUserIDW PrismObjectDocumentGrantUpdateResponseUserID = "w"
)

func (r PrismObjectDocumentGrantUpdateResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectDocumentGrantUpdateResponseUserIDA, PrismObjectDocumentGrantUpdateResponseUserIDR, PrismObjectDocumentGrantUpdateResponseUserIDW:
		return true
	}
	return false
}

type PrismObjectDocumentGrantGetResponse struct {
	TeamGroupID []map[string]PrismObjectDocumentGrantGetResponseTeamGroupID `json:"team_group_id"`
	TeamID      map[string]PrismObjectDocumentGrantGetResponseTeamID        `json:"team_id"`
	UserID      []map[string]PrismObjectDocumentGrantGetResponseUserID      `json:"user_id"`
	JSON        prismObjectDocumentGrantGetResponseJSON                     `json:"-"`
}

// prismObjectDocumentGrantGetResponseJSON contains the JSON metadata for the
// struct [PrismObjectDocumentGrantGetResponse]
type prismObjectDocumentGrantGetResponseJSON struct {
	TeamGroupID apijson.Field
	TeamID      apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectDocumentGrantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectDocumentGrantGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectDocumentGrantGetResponseTeamGroupID string

const (
	PrismObjectDocumentGrantGetResponseTeamGroupIDA PrismObjectDocumentGrantGetResponseTeamGroupID = "a"
	PrismObjectDocumentGrantGetResponseTeamGroupIDR PrismObjectDocumentGrantGetResponseTeamGroupID = "r"
	PrismObjectDocumentGrantGetResponseTeamGroupIDW PrismObjectDocumentGrantGetResponseTeamGroupID = "w"
)

func (r PrismObjectDocumentGrantGetResponseTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectDocumentGrantGetResponseTeamGroupIDA, PrismObjectDocumentGrantGetResponseTeamGroupIDR, PrismObjectDocumentGrantGetResponseTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectDocumentGrantGetResponseTeamID string

const (
	PrismObjectDocumentGrantGetResponseTeamIDA PrismObjectDocumentGrantGetResponseTeamID = "a"
	PrismObjectDocumentGrantGetResponseTeamIDR PrismObjectDocumentGrantGetResponseTeamID = "r"
	PrismObjectDocumentGrantGetResponseTeamIDW PrismObjectDocumentGrantGetResponseTeamID = "w"
)

func (r PrismObjectDocumentGrantGetResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectDocumentGrantGetResponseTeamIDA, PrismObjectDocumentGrantGetResponseTeamIDR, PrismObjectDocumentGrantGetResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectDocumentGrantGetResponseUserID string

const (
	PrismObjectDocumentGrantGetResponseUserIDA PrismObjectDocumentGrantGetResponseUserID = "a"
	PrismObjectDocumentGrantGetResponseUserIDR PrismObjectDocumentGrantGetResponseUserID = "r"
	PrismObjectDocumentGrantGetResponseUserIDW PrismObjectDocumentGrantGetResponseUserID = "w"
)

func (r PrismObjectDocumentGrantGetResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectDocumentGrantGetResponseUserIDA, PrismObjectDocumentGrantGetResponseUserIDR, PrismObjectDocumentGrantGetResponseUserIDW:
		return true
	}
	return false
}

type PrismObjectDocumentGrantUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	PathTeamID  param.Field[string]                                                       `path:"teamId" api:"required" format:"uuid"`
	TeamGroupID param.Field[[]map[string]PrismObjectDocumentGrantUpdateParamsTeamGroupID] `json:"team_group_id"`
	BodyTeamID  param.Field[map[string]PrismObjectDocumentGrantUpdateParamsTeamID]        `json:"team_id"`
	UserID      param.Field[[]map[string]PrismObjectDocumentGrantUpdateParamsUserID]      `json:"user_id"`
}

func (r PrismObjectDocumentGrantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectDocumentGrantUpdateParamsTeamGroupID string

const (
	PrismObjectDocumentGrantUpdateParamsTeamGroupIDA PrismObjectDocumentGrantUpdateParamsTeamGroupID = "a"
	PrismObjectDocumentGrantUpdateParamsTeamGroupIDR PrismObjectDocumentGrantUpdateParamsTeamGroupID = "r"
	PrismObjectDocumentGrantUpdateParamsTeamGroupIDW PrismObjectDocumentGrantUpdateParamsTeamGroupID = "w"
)

func (r PrismObjectDocumentGrantUpdateParamsTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectDocumentGrantUpdateParamsTeamGroupIDA, PrismObjectDocumentGrantUpdateParamsTeamGroupIDR, PrismObjectDocumentGrantUpdateParamsTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectDocumentGrantUpdateParamsTeamID string

const (
	PrismObjectDocumentGrantUpdateParamsTeamIDA PrismObjectDocumentGrantUpdateParamsTeamID = "a"
	PrismObjectDocumentGrantUpdateParamsTeamIDR PrismObjectDocumentGrantUpdateParamsTeamID = "r"
	PrismObjectDocumentGrantUpdateParamsTeamIDW PrismObjectDocumentGrantUpdateParamsTeamID = "w"
)

func (r PrismObjectDocumentGrantUpdateParamsTeamID) IsKnown() bool {
	switch r {
	case PrismObjectDocumentGrantUpdateParamsTeamIDA, PrismObjectDocumentGrantUpdateParamsTeamIDR, PrismObjectDocumentGrantUpdateParamsTeamIDW:
		return true
	}
	return false
}

type PrismObjectDocumentGrantUpdateParamsUserID string

const (
	PrismObjectDocumentGrantUpdateParamsUserIDA PrismObjectDocumentGrantUpdateParamsUserID = "a"
	PrismObjectDocumentGrantUpdateParamsUserIDR PrismObjectDocumentGrantUpdateParamsUserID = "r"
	PrismObjectDocumentGrantUpdateParamsUserIDW PrismObjectDocumentGrantUpdateParamsUserID = "w"
)

func (r PrismObjectDocumentGrantUpdateParamsUserID) IsKnown() bool {
	switch r {
	case PrismObjectDocumentGrantUpdateParamsUserIDA, PrismObjectDocumentGrantUpdateParamsUserIDR, PrismObjectDocumentGrantUpdateParamsUserIDW:
		return true
	}
	return false
}

type PrismObjectDocumentGrantGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
