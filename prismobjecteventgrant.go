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
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/event/%s/grant", params.PathTeamID, eventID)
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
	path := fmt.Sprintf("v2/prism/%s/event/%s/grant", query.TeamID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The grants on a record. For `message`, also carries the entity ids of everyone
// on the message, resolved from its address headers when the grant was written.
// The id arrays are read-only and are null when participant resolution was
// unavailable (for example the mailbox had no Gmail token at the time).
type PrismObjectEventGrantUpdateResponse struct {
	ContactIDs      []string                                              `json:"contact_ids" api:"nullable" format:"uuid"`
	GroupID         map[string]PrismObjectEventGrantUpdateResponseGroupID `json:"group_id"`
	IdentityIDs     []string                                              `json:"identity_ids" api:"nullable" format:"uuid"`
	OrganizationIDs []string                                              `json:"organization_ids" api:"nullable" format:"uuid"`
	// How much of the record the grant exposes. `metadata` shares only the record's
	// headers and participants; `full` shares its contents. Currently recorded on the
	// access row and returned on read — it is not yet enforced by the read path.
	// Applies to `message` grants; ignored for other object types.
	ShareLevel PrismObjectEventGrantUpdateResponseShareLevel        `json:"share_level"`
	TeamID     map[string]PrismObjectEventGrantUpdateResponseTeamID `json:"team_id"`
	UserID     map[string]PrismObjectEventGrantUpdateResponseUserID `json:"user_id"`
	JSON       prismObjectEventGrantUpdateResponseJSON              `json:"-"`
}

// prismObjectEventGrantUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectEventGrantUpdateResponse]
type prismObjectEventGrantUpdateResponseJSON struct {
	ContactIDs      apijson.Field
	GroupID         apijson.Field
	IdentityIDs     apijson.Field
	OrganizationIDs apijson.Field
	ShareLevel      apijson.Field
	TeamID          apijson.Field
	UserID          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PrismObjectEventGrantUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEventGrantUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEventGrantUpdateResponseGroupID string

const (
	PrismObjectEventGrantUpdateResponseGroupIDA PrismObjectEventGrantUpdateResponseGroupID = "a"
	PrismObjectEventGrantUpdateResponseGroupIDR PrismObjectEventGrantUpdateResponseGroupID = "r"
	PrismObjectEventGrantUpdateResponseGroupIDW PrismObjectEventGrantUpdateResponseGroupID = "w"
)

func (r PrismObjectEventGrantUpdateResponseGroupID) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantUpdateResponseGroupIDA, PrismObjectEventGrantUpdateResponseGroupIDR, PrismObjectEventGrantUpdateResponseGroupIDW:
		return true
	}
	return false
}

// How much of the record the grant exposes. `metadata` shares only the record's
// headers and participants; `full` shares its contents. Currently recorded on the
// access row and returned on read — it is not yet enforced by the read path.
// Applies to `message` grants; ignored for other object types.
type PrismObjectEventGrantUpdateResponseShareLevel string

const (
	PrismObjectEventGrantUpdateResponseShareLevelMetadata PrismObjectEventGrantUpdateResponseShareLevel = "metadata"
	PrismObjectEventGrantUpdateResponseShareLevelFull     PrismObjectEventGrantUpdateResponseShareLevel = "full"
)

func (r PrismObjectEventGrantUpdateResponseShareLevel) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantUpdateResponseShareLevelMetadata, PrismObjectEventGrantUpdateResponseShareLevelFull:
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

// The grants on a record. For `message`, also carries the entity ids of everyone
// on the message, resolved from its address headers when the grant was written.
// The id arrays are read-only and are null when participant resolution was
// unavailable (for example the mailbox had no Gmail token at the time).
type PrismObjectEventGrantGetResponse struct {
	ContactIDs      []string                                           `json:"contact_ids" api:"nullable" format:"uuid"`
	GroupID         map[string]PrismObjectEventGrantGetResponseGroupID `json:"group_id"`
	IdentityIDs     []string                                           `json:"identity_ids" api:"nullable" format:"uuid"`
	OrganizationIDs []string                                           `json:"organization_ids" api:"nullable" format:"uuid"`
	// How much of the record the grant exposes. `metadata` shares only the record's
	// headers and participants; `full` shares its contents. Currently recorded on the
	// access row and returned on read — it is not yet enforced by the read path.
	// Applies to `message` grants; ignored for other object types.
	ShareLevel PrismObjectEventGrantGetResponseShareLevel        `json:"share_level"`
	TeamID     map[string]PrismObjectEventGrantGetResponseTeamID `json:"team_id"`
	UserID     map[string]PrismObjectEventGrantGetResponseUserID `json:"user_id"`
	JSON       prismObjectEventGrantGetResponseJSON              `json:"-"`
}

// prismObjectEventGrantGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectEventGrantGetResponse]
type prismObjectEventGrantGetResponseJSON struct {
	ContactIDs      apijson.Field
	GroupID         apijson.Field
	IdentityIDs     apijson.Field
	OrganizationIDs apijson.Field
	ShareLevel      apijson.Field
	TeamID          apijson.Field
	UserID          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PrismObjectEventGrantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEventGrantGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEventGrantGetResponseGroupID string

const (
	PrismObjectEventGrantGetResponseGroupIDA PrismObjectEventGrantGetResponseGroupID = "a"
	PrismObjectEventGrantGetResponseGroupIDR PrismObjectEventGrantGetResponseGroupID = "r"
	PrismObjectEventGrantGetResponseGroupIDW PrismObjectEventGrantGetResponseGroupID = "w"
)

func (r PrismObjectEventGrantGetResponseGroupID) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantGetResponseGroupIDA, PrismObjectEventGrantGetResponseGroupIDR, PrismObjectEventGrantGetResponseGroupIDW:
		return true
	}
	return false
}

// How much of the record the grant exposes. `metadata` shares only the record's
// headers and participants; `full` shares its contents. Currently recorded on the
// access row and returned on read — it is not yet enforced by the read path.
// Applies to `message` grants; ignored for other object types.
type PrismObjectEventGrantGetResponseShareLevel string

const (
	PrismObjectEventGrantGetResponseShareLevelMetadata PrismObjectEventGrantGetResponseShareLevel = "metadata"
	PrismObjectEventGrantGetResponseShareLevelFull     PrismObjectEventGrantGetResponseShareLevel = "full"
)

func (r PrismObjectEventGrantGetResponseShareLevel) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantGetResponseShareLevelMetadata, PrismObjectEventGrantGetResponseShareLevelFull:
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
	PathTeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// How much of the record the grant exposes. `metadata` shares only the record's
	// headers and participants; `full` shares its contents. Currently recorded on the
	// access row and returned on read — it is not yet enforced by the read path.
	// Applies to `message` grants; ignored for other object types.
	ShareLevel     param.Field[PrismObjectEventGrantUpdateParamsShareLevel]               `json:"share_level"`
	TeamGroupID    param.Field[[]map[string]PrismObjectEventGrantUpdateParamsTeamGroupID] `json:"team_group_id"`
	BodyTeamID     param.Field[map[string]PrismObjectEventGrantUpdateParamsTeamID]        `json:"team_id"`
	UserID         param.Field[[]map[string]PrismObjectEventGrantUpdateParamsUserID]      `json:"user_id"`
	IdempotencyKey param.Field[string]                                                    `header:"Idempotency-Key"`
}

func (r PrismObjectEventGrantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How much of the record the grant exposes. `metadata` shares only the record's
// headers and participants; `full` shares its contents. Currently recorded on the
// access row and returned on read — it is not yet enforced by the read path.
// Applies to `message` grants; ignored for other object types.
type PrismObjectEventGrantUpdateParamsShareLevel string

const (
	PrismObjectEventGrantUpdateParamsShareLevelMetadata PrismObjectEventGrantUpdateParamsShareLevel = "metadata"
	PrismObjectEventGrantUpdateParamsShareLevelFull     PrismObjectEventGrantUpdateParamsShareLevel = "full"
)

func (r PrismObjectEventGrantUpdateParamsShareLevel) IsKnown() bool {
	switch r {
	case PrismObjectEventGrantUpdateParamsShareLevelMetadata, PrismObjectEventGrantUpdateParamsShareLevelFull:
		return true
	}
	return false
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
