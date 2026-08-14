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

// PrismObjectEngagementGrantService contains methods and other services that help
// with interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectEngagementGrantService] method instead.
type PrismObjectEngagementGrantService struct {
	Options []option.RequestOption
}

// NewPrismObjectEngagementGrantService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPrismObjectEngagementGrantService(opts ...option.RequestOption) (r *PrismObjectEngagementGrantService) {
	r = &PrismObjectEngagementGrantService{}
	r.Options = opts
	return
}

// Update grant
func (r *PrismObjectEngagementGrantService) Update(ctx context.Context, engagementID string, params PrismObjectEngagementGrantUpdateParams, opts ...option.RequestOption) (res *PrismObjectEngagementGrantUpdateResponse, err error) {
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
	if engagementID == "" {
		err = errors.New("missing required engagementId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/engagement/%s/grant", params.PathTeamID, engagementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Get grant
func (r *PrismObjectEngagementGrantService) Get(ctx context.Context, engagementID string, query PrismObjectEngagementGrantGetParams, opts ...option.RequestOption) (res *PrismObjectEngagementGrantGetResponse, err error) {
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
	if engagementID == "" {
		err = errors.New("missing required engagementId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/engagement/%s/grant", query.TeamID, engagementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The grants on a record. For `message`, also carries the entity ids of everyone
// on the message, resolved from its address headers when the grant was written.
// The id arrays are read-only and are null when participant resolution was
// unavailable (for example the mailbox had no Gmail token at the time).
type PrismObjectEngagementGrantUpdateResponse struct {
	ContactIDs      []string                                                   `json:"contact_ids" api:"nullable" format:"uuid"`
	GroupID         map[string]PrismObjectEngagementGrantUpdateResponseGroupID `json:"group_id"`
	IdentityIDs     []string                                                   `json:"identity_ids" api:"nullable" format:"uuid"`
	OrganizationIDs []string                                                   `json:"organization_ids" api:"nullable" format:"uuid"`
	// How much of the record the grant exposes. `metadata` shares only the record's
	// headers and participants; `full` shares its contents. Currently recorded on the
	// access row and returned on read — it is not yet enforced by the read path.
	// Applies to `message` grants; ignored for other object types.
	ShareLevel PrismObjectEngagementGrantUpdateResponseShareLevel        `json:"share_level"`
	TeamID     map[string]PrismObjectEngagementGrantUpdateResponseTeamID `json:"team_id"`
	UserID     map[string]PrismObjectEngagementGrantUpdateResponseUserID `json:"user_id"`
	JSON       prismObjectEngagementGrantUpdateResponseJSON              `json:"-"`
}

// prismObjectEngagementGrantUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectEngagementGrantUpdateResponse]
type prismObjectEngagementGrantUpdateResponseJSON struct {
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

func (r *PrismObjectEngagementGrantUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementGrantUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementGrantUpdateResponseGroupID string

const (
	PrismObjectEngagementGrantUpdateResponseGroupIDA PrismObjectEngagementGrantUpdateResponseGroupID = "a"
	PrismObjectEngagementGrantUpdateResponseGroupIDR PrismObjectEngagementGrantUpdateResponseGroupID = "r"
	PrismObjectEngagementGrantUpdateResponseGroupIDW PrismObjectEngagementGrantUpdateResponseGroupID = "w"
)

func (r PrismObjectEngagementGrantUpdateResponseGroupID) IsKnown() bool {
	switch r {
	case PrismObjectEngagementGrantUpdateResponseGroupIDA, PrismObjectEngagementGrantUpdateResponseGroupIDR, PrismObjectEngagementGrantUpdateResponseGroupIDW:
		return true
	}
	return false
}

// How much of the record the grant exposes. `metadata` shares only the record's
// headers and participants; `full` shares its contents. Currently recorded on the
// access row and returned on read — it is not yet enforced by the read path.
// Applies to `message` grants; ignored for other object types.
type PrismObjectEngagementGrantUpdateResponseShareLevel string

const (
	PrismObjectEngagementGrantUpdateResponseShareLevelMetadata PrismObjectEngagementGrantUpdateResponseShareLevel = "metadata"
	PrismObjectEngagementGrantUpdateResponseShareLevelFull     PrismObjectEngagementGrantUpdateResponseShareLevel = "full"
)

func (r PrismObjectEngagementGrantUpdateResponseShareLevel) IsKnown() bool {
	switch r {
	case PrismObjectEngagementGrantUpdateResponseShareLevelMetadata, PrismObjectEngagementGrantUpdateResponseShareLevelFull:
		return true
	}
	return false
}

type PrismObjectEngagementGrantUpdateResponseTeamID string

const (
	PrismObjectEngagementGrantUpdateResponseTeamIDA PrismObjectEngagementGrantUpdateResponseTeamID = "a"
	PrismObjectEngagementGrantUpdateResponseTeamIDR PrismObjectEngagementGrantUpdateResponseTeamID = "r"
	PrismObjectEngagementGrantUpdateResponseTeamIDW PrismObjectEngagementGrantUpdateResponseTeamID = "w"
)

func (r PrismObjectEngagementGrantUpdateResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectEngagementGrantUpdateResponseTeamIDA, PrismObjectEngagementGrantUpdateResponseTeamIDR, PrismObjectEngagementGrantUpdateResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectEngagementGrantUpdateResponseUserID string

const (
	PrismObjectEngagementGrantUpdateResponseUserIDA PrismObjectEngagementGrantUpdateResponseUserID = "a"
	PrismObjectEngagementGrantUpdateResponseUserIDR PrismObjectEngagementGrantUpdateResponseUserID = "r"
	PrismObjectEngagementGrantUpdateResponseUserIDW PrismObjectEngagementGrantUpdateResponseUserID = "w"
)

func (r PrismObjectEngagementGrantUpdateResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectEngagementGrantUpdateResponseUserIDA, PrismObjectEngagementGrantUpdateResponseUserIDR, PrismObjectEngagementGrantUpdateResponseUserIDW:
		return true
	}
	return false
}

// The grants on a record. For `message`, also carries the entity ids of everyone
// on the message, resolved from its address headers when the grant was written.
// The id arrays are read-only and are null when participant resolution was
// unavailable (for example the mailbox had no Gmail token at the time).
type PrismObjectEngagementGrantGetResponse struct {
	ContactIDs      []string                                                `json:"contact_ids" api:"nullable" format:"uuid"`
	GroupID         map[string]PrismObjectEngagementGrantGetResponseGroupID `json:"group_id"`
	IdentityIDs     []string                                                `json:"identity_ids" api:"nullable" format:"uuid"`
	OrganizationIDs []string                                                `json:"organization_ids" api:"nullable" format:"uuid"`
	// How much of the record the grant exposes. `metadata` shares only the record's
	// headers and participants; `full` shares its contents. Currently recorded on the
	// access row and returned on read — it is not yet enforced by the read path.
	// Applies to `message` grants; ignored for other object types.
	ShareLevel PrismObjectEngagementGrantGetResponseShareLevel        `json:"share_level"`
	TeamID     map[string]PrismObjectEngagementGrantGetResponseTeamID `json:"team_id"`
	UserID     map[string]PrismObjectEngagementGrantGetResponseUserID `json:"user_id"`
	JSON       prismObjectEngagementGrantGetResponseJSON              `json:"-"`
}

// prismObjectEngagementGrantGetResponseJSON contains the JSON metadata for the
// struct [PrismObjectEngagementGrantGetResponse]
type prismObjectEngagementGrantGetResponseJSON struct {
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

func (r *PrismObjectEngagementGrantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEngagementGrantGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEngagementGrantGetResponseGroupID string

const (
	PrismObjectEngagementGrantGetResponseGroupIDA PrismObjectEngagementGrantGetResponseGroupID = "a"
	PrismObjectEngagementGrantGetResponseGroupIDR PrismObjectEngagementGrantGetResponseGroupID = "r"
	PrismObjectEngagementGrantGetResponseGroupIDW PrismObjectEngagementGrantGetResponseGroupID = "w"
)

func (r PrismObjectEngagementGrantGetResponseGroupID) IsKnown() bool {
	switch r {
	case PrismObjectEngagementGrantGetResponseGroupIDA, PrismObjectEngagementGrantGetResponseGroupIDR, PrismObjectEngagementGrantGetResponseGroupIDW:
		return true
	}
	return false
}

// How much of the record the grant exposes. `metadata` shares only the record's
// headers and participants; `full` shares its contents. Currently recorded on the
// access row and returned on read — it is not yet enforced by the read path.
// Applies to `message` grants; ignored for other object types.
type PrismObjectEngagementGrantGetResponseShareLevel string

const (
	PrismObjectEngagementGrantGetResponseShareLevelMetadata PrismObjectEngagementGrantGetResponseShareLevel = "metadata"
	PrismObjectEngagementGrantGetResponseShareLevelFull     PrismObjectEngagementGrantGetResponseShareLevel = "full"
)

func (r PrismObjectEngagementGrantGetResponseShareLevel) IsKnown() bool {
	switch r {
	case PrismObjectEngagementGrantGetResponseShareLevelMetadata, PrismObjectEngagementGrantGetResponseShareLevelFull:
		return true
	}
	return false
}

type PrismObjectEngagementGrantGetResponseTeamID string

const (
	PrismObjectEngagementGrantGetResponseTeamIDA PrismObjectEngagementGrantGetResponseTeamID = "a"
	PrismObjectEngagementGrantGetResponseTeamIDR PrismObjectEngagementGrantGetResponseTeamID = "r"
	PrismObjectEngagementGrantGetResponseTeamIDW PrismObjectEngagementGrantGetResponseTeamID = "w"
)

func (r PrismObjectEngagementGrantGetResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectEngagementGrantGetResponseTeamIDA, PrismObjectEngagementGrantGetResponseTeamIDR, PrismObjectEngagementGrantGetResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectEngagementGrantGetResponseUserID string

const (
	PrismObjectEngagementGrantGetResponseUserIDA PrismObjectEngagementGrantGetResponseUserID = "a"
	PrismObjectEngagementGrantGetResponseUserIDR PrismObjectEngagementGrantGetResponseUserID = "r"
	PrismObjectEngagementGrantGetResponseUserIDW PrismObjectEngagementGrantGetResponseUserID = "w"
)

func (r PrismObjectEngagementGrantGetResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectEngagementGrantGetResponseUserIDA, PrismObjectEngagementGrantGetResponseUserIDR, PrismObjectEngagementGrantGetResponseUserIDW:
		return true
	}
	return false
}

type PrismObjectEngagementGrantUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	PathTeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// How much of the record the grant exposes. `metadata` shares only the record's
	// headers and participants; `full` shares its contents. Currently recorded on the
	// access row and returned on read — it is not yet enforced by the read path.
	// Applies to `message` grants; ignored for other object types.
	ShareLevel     param.Field[PrismObjectEngagementGrantUpdateParamsShareLevel]               `json:"share_level"`
	TeamGroupID    param.Field[[]map[string]PrismObjectEngagementGrantUpdateParamsTeamGroupID] `json:"team_group_id"`
	BodyTeamID     param.Field[map[string]PrismObjectEngagementGrantUpdateParamsTeamID]        `json:"team_id"`
	UserID         param.Field[[]map[string]PrismObjectEngagementGrantUpdateParamsUserID]      `json:"user_id"`
	IdempotencyKey param.Field[string]                                                         `header:"Idempotency-Key"`
}

func (r PrismObjectEngagementGrantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How much of the record the grant exposes. `metadata` shares only the record's
// headers and participants; `full` shares its contents. Currently recorded on the
// access row and returned on read — it is not yet enforced by the read path.
// Applies to `message` grants; ignored for other object types.
type PrismObjectEngagementGrantUpdateParamsShareLevel string

const (
	PrismObjectEngagementGrantUpdateParamsShareLevelMetadata PrismObjectEngagementGrantUpdateParamsShareLevel = "metadata"
	PrismObjectEngagementGrantUpdateParamsShareLevelFull     PrismObjectEngagementGrantUpdateParamsShareLevel = "full"
)

func (r PrismObjectEngagementGrantUpdateParamsShareLevel) IsKnown() bool {
	switch r {
	case PrismObjectEngagementGrantUpdateParamsShareLevelMetadata, PrismObjectEngagementGrantUpdateParamsShareLevelFull:
		return true
	}
	return false
}

type PrismObjectEngagementGrantUpdateParamsTeamGroupID string

const (
	PrismObjectEngagementGrantUpdateParamsTeamGroupIDA PrismObjectEngagementGrantUpdateParamsTeamGroupID = "a"
	PrismObjectEngagementGrantUpdateParamsTeamGroupIDR PrismObjectEngagementGrantUpdateParamsTeamGroupID = "r"
	PrismObjectEngagementGrantUpdateParamsTeamGroupIDW PrismObjectEngagementGrantUpdateParamsTeamGroupID = "w"
)

func (r PrismObjectEngagementGrantUpdateParamsTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectEngagementGrantUpdateParamsTeamGroupIDA, PrismObjectEngagementGrantUpdateParamsTeamGroupIDR, PrismObjectEngagementGrantUpdateParamsTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectEngagementGrantUpdateParamsTeamID string

const (
	PrismObjectEngagementGrantUpdateParamsTeamIDA PrismObjectEngagementGrantUpdateParamsTeamID = "a"
	PrismObjectEngagementGrantUpdateParamsTeamIDR PrismObjectEngagementGrantUpdateParamsTeamID = "r"
	PrismObjectEngagementGrantUpdateParamsTeamIDW PrismObjectEngagementGrantUpdateParamsTeamID = "w"
)

func (r PrismObjectEngagementGrantUpdateParamsTeamID) IsKnown() bool {
	switch r {
	case PrismObjectEngagementGrantUpdateParamsTeamIDA, PrismObjectEngagementGrantUpdateParamsTeamIDR, PrismObjectEngagementGrantUpdateParamsTeamIDW:
		return true
	}
	return false
}

type PrismObjectEngagementGrantUpdateParamsUserID string

const (
	PrismObjectEngagementGrantUpdateParamsUserIDA PrismObjectEngagementGrantUpdateParamsUserID = "a"
	PrismObjectEngagementGrantUpdateParamsUserIDR PrismObjectEngagementGrantUpdateParamsUserID = "r"
	PrismObjectEngagementGrantUpdateParamsUserIDW PrismObjectEngagementGrantUpdateParamsUserID = "w"
)

func (r PrismObjectEngagementGrantUpdateParamsUserID) IsKnown() bool {
	switch r {
	case PrismObjectEngagementGrantUpdateParamsUserIDA, PrismObjectEngagementGrantUpdateParamsUserIDR, PrismObjectEngagementGrantUpdateParamsUserIDW:
		return true
	}
	return false
}

type PrismObjectEngagementGrantGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
