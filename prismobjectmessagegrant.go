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

// PrismObjectMessageGrantService contains methods and other services that help
// with interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectMessageGrantService] method instead.
type PrismObjectMessageGrantService struct {
	Options []option.RequestOption
}

// NewPrismObjectMessageGrantService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectMessageGrantService(opts ...option.RequestOption) (r *PrismObjectMessageGrantService) {
	r = &PrismObjectMessageGrantService{}
	r.Options = opts
	return
}

// Update grant
func (r *PrismObjectMessageGrantService) Update(ctx context.Context, messageID string, params PrismObjectMessageGrantUpdateParams, opts ...option.RequestOption) (res *PrismObjectMessageGrantUpdateResponse, err error) {
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
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/message/%s/grant", params.PathTeamID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Get grant
func (r *PrismObjectMessageGrantService) Get(ctx context.Context, messageID string, query PrismObjectMessageGrantGetParams, opts ...option.RequestOption) (res *PrismObjectMessageGrantGetResponse, err error) {
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
	if messageID == "" {
		err = errors.New("missing required messageId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/message/%s/grant", query.TeamID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The grants on a record. For `message`, also carries the entity ids of everyone
// on the message, resolved from its address headers when the grant was written.
// The id arrays are read-only and are null when participant resolution was
// unavailable (for example the mailbox had no Gmail token at the time).
type PrismObjectMessageGrantUpdateResponse struct {
	ContactIDs      []string                                                `json:"contact_ids" api:"nullable" format:"uuid"`
	GroupID         map[string]PrismObjectMessageGrantUpdateResponseGroupID `json:"group_id"`
	IdentityIDs     []string                                                `json:"identity_ids" api:"nullable" format:"uuid"`
	OrganizationIDs []string                                                `json:"organization_ids" api:"nullable" format:"uuid"`
	// How much of the record the grant exposes. `metadata` shares only the record's
	// headers and participants; `full` shares its contents. Currently recorded on the
	// access row and returned on read — it is not yet enforced by the read path.
	// Applies to `message` grants; ignored for other object types.
	ShareLevel PrismObjectMessageGrantUpdateResponseShareLevel        `json:"share_level"`
	TeamID     map[string]PrismObjectMessageGrantUpdateResponseTeamID `json:"team_id"`
	UserID     map[string]PrismObjectMessageGrantUpdateResponseUserID `json:"user_id"`
	JSON       prismObjectMessageGrantUpdateResponseJSON              `json:"-"`
}

// prismObjectMessageGrantUpdateResponseJSON contains the JSON metadata for the
// struct [PrismObjectMessageGrantUpdateResponse]
type prismObjectMessageGrantUpdateResponseJSON struct {
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

func (r *PrismObjectMessageGrantUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectMessageGrantUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectMessageGrantUpdateResponseGroupID string

const (
	PrismObjectMessageGrantUpdateResponseGroupIDA PrismObjectMessageGrantUpdateResponseGroupID = "a"
	PrismObjectMessageGrantUpdateResponseGroupIDR PrismObjectMessageGrantUpdateResponseGroupID = "r"
	PrismObjectMessageGrantUpdateResponseGroupIDW PrismObjectMessageGrantUpdateResponseGroupID = "w"
)

func (r PrismObjectMessageGrantUpdateResponseGroupID) IsKnown() bool {
	switch r {
	case PrismObjectMessageGrantUpdateResponseGroupIDA, PrismObjectMessageGrantUpdateResponseGroupIDR, PrismObjectMessageGrantUpdateResponseGroupIDW:
		return true
	}
	return false
}

// How much of the record the grant exposes. `metadata` shares only the record's
// headers and participants; `full` shares its contents. Currently recorded on the
// access row and returned on read — it is not yet enforced by the read path.
// Applies to `message` grants; ignored for other object types.
type PrismObjectMessageGrantUpdateResponseShareLevel string

const (
	PrismObjectMessageGrantUpdateResponseShareLevelMetadata PrismObjectMessageGrantUpdateResponseShareLevel = "metadata"
	PrismObjectMessageGrantUpdateResponseShareLevelFull     PrismObjectMessageGrantUpdateResponseShareLevel = "full"
)

func (r PrismObjectMessageGrantUpdateResponseShareLevel) IsKnown() bool {
	switch r {
	case PrismObjectMessageGrantUpdateResponseShareLevelMetadata, PrismObjectMessageGrantUpdateResponseShareLevelFull:
		return true
	}
	return false
}

type PrismObjectMessageGrantUpdateResponseTeamID string

const (
	PrismObjectMessageGrantUpdateResponseTeamIDA PrismObjectMessageGrantUpdateResponseTeamID = "a"
	PrismObjectMessageGrantUpdateResponseTeamIDR PrismObjectMessageGrantUpdateResponseTeamID = "r"
	PrismObjectMessageGrantUpdateResponseTeamIDW PrismObjectMessageGrantUpdateResponseTeamID = "w"
)

func (r PrismObjectMessageGrantUpdateResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectMessageGrantUpdateResponseTeamIDA, PrismObjectMessageGrantUpdateResponseTeamIDR, PrismObjectMessageGrantUpdateResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectMessageGrantUpdateResponseUserID string

const (
	PrismObjectMessageGrantUpdateResponseUserIDA PrismObjectMessageGrantUpdateResponseUserID = "a"
	PrismObjectMessageGrantUpdateResponseUserIDR PrismObjectMessageGrantUpdateResponseUserID = "r"
	PrismObjectMessageGrantUpdateResponseUserIDW PrismObjectMessageGrantUpdateResponseUserID = "w"
)

func (r PrismObjectMessageGrantUpdateResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectMessageGrantUpdateResponseUserIDA, PrismObjectMessageGrantUpdateResponseUserIDR, PrismObjectMessageGrantUpdateResponseUserIDW:
		return true
	}
	return false
}

// The grants on a record. For `message`, also carries the entity ids of everyone
// on the message, resolved from its address headers when the grant was written.
// The id arrays are read-only and are null when participant resolution was
// unavailable (for example the mailbox had no Gmail token at the time).
type PrismObjectMessageGrantGetResponse struct {
	ContactIDs      []string                                             `json:"contact_ids" api:"nullable" format:"uuid"`
	GroupID         map[string]PrismObjectMessageGrantGetResponseGroupID `json:"group_id"`
	IdentityIDs     []string                                             `json:"identity_ids" api:"nullable" format:"uuid"`
	OrganizationIDs []string                                             `json:"organization_ids" api:"nullable" format:"uuid"`
	// How much of the record the grant exposes. `metadata` shares only the record's
	// headers and participants; `full` shares its contents. Currently recorded on the
	// access row and returned on read — it is not yet enforced by the read path.
	// Applies to `message` grants; ignored for other object types.
	ShareLevel PrismObjectMessageGrantGetResponseShareLevel        `json:"share_level"`
	TeamID     map[string]PrismObjectMessageGrantGetResponseTeamID `json:"team_id"`
	UserID     map[string]PrismObjectMessageGrantGetResponseUserID `json:"user_id"`
	JSON       prismObjectMessageGrantGetResponseJSON              `json:"-"`
}

// prismObjectMessageGrantGetResponseJSON contains the JSON metadata for the struct
// [PrismObjectMessageGrantGetResponse]
type prismObjectMessageGrantGetResponseJSON struct {
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

func (r *PrismObjectMessageGrantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectMessageGrantGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectMessageGrantGetResponseGroupID string

const (
	PrismObjectMessageGrantGetResponseGroupIDA PrismObjectMessageGrantGetResponseGroupID = "a"
	PrismObjectMessageGrantGetResponseGroupIDR PrismObjectMessageGrantGetResponseGroupID = "r"
	PrismObjectMessageGrantGetResponseGroupIDW PrismObjectMessageGrantGetResponseGroupID = "w"
)

func (r PrismObjectMessageGrantGetResponseGroupID) IsKnown() bool {
	switch r {
	case PrismObjectMessageGrantGetResponseGroupIDA, PrismObjectMessageGrantGetResponseGroupIDR, PrismObjectMessageGrantGetResponseGroupIDW:
		return true
	}
	return false
}

// How much of the record the grant exposes. `metadata` shares only the record's
// headers and participants; `full` shares its contents. Currently recorded on the
// access row and returned on read — it is not yet enforced by the read path.
// Applies to `message` grants; ignored for other object types.
type PrismObjectMessageGrantGetResponseShareLevel string

const (
	PrismObjectMessageGrantGetResponseShareLevelMetadata PrismObjectMessageGrantGetResponseShareLevel = "metadata"
	PrismObjectMessageGrantGetResponseShareLevelFull     PrismObjectMessageGrantGetResponseShareLevel = "full"
)

func (r PrismObjectMessageGrantGetResponseShareLevel) IsKnown() bool {
	switch r {
	case PrismObjectMessageGrantGetResponseShareLevelMetadata, PrismObjectMessageGrantGetResponseShareLevelFull:
		return true
	}
	return false
}

type PrismObjectMessageGrantGetResponseTeamID string

const (
	PrismObjectMessageGrantGetResponseTeamIDA PrismObjectMessageGrantGetResponseTeamID = "a"
	PrismObjectMessageGrantGetResponseTeamIDR PrismObjectMessageGrantGetResponseTeamID = "r"
	PrismObjectMessageGrantGetResponseTeamIDW PrismObjectMessageGrantGetResponseTeamID = "w"
)

func (r PrismObjectMessageGrantGetResponseTeamID) IsKnown() bool {
	switch r {
	case PrismObjectMessageGrantGetResponseTeamIDA, PrismObjectMessageGrantGetResponseTeamIDR, PrismObjectMessageGrantGetResponseTeamIDW:
		return true
	}
	return false
}

type PrismObjectMessageGrantGetResponseUserID string

const (
	PrismObjectMessageGrantGetResponseUserIDA PrismObjectMessageGrantGetResponseUserID = "a"
	PrismObjectMessageGrantGetResponseUserIDR PrismObjectMessageGrantGetResponseUserID = "r"
	PrismObjectMessageGrantGetResponseUserIDW PrismObjectMessageGrantGetResponseUserID = "w"
)

func (r PrismObjectMessageGrantGetResponseUserID) IsKnown() bool {
	switch r {
	case PrismObjectMessageGrantGetResponseUserIDA, PrismObjectMessageGrantGetResponseUserIDR, PrismObjectMessageGrantGetResponseUserIDW:
		return true
	}
	return false
}

type PrismObjectMessageGrantUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	PathTeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// How much of the record the grant exposes. `metadata` shares only the record's
	// headers and participants; `full` shares its contents. Currently recorded on the
	// access row and returned on read — it is not yet enforced by the read path.
	// Applies to `message` grants; ignored for other object types.
	ShareLevel     param.Field[PrismObjectMessageGrantUpdateParamsShareLevel]               `json:"share_level"`
	TeamGroupID    param.Field[[]map[string]PrismObjectMessageGrantUpdateParamsTeamGroupID] `json:"team_group_id"`
	BodyTeamID     param.Field[map[string]PrismObjectMessageGrantUpdateParamsTeamID]        `json:"team_id"`
	UserID         param.Field[[]map[string]PrismObjectMessageGrantUpdateParamsUserID]      `json:"user_id"`
	IdempotencyKey param.Field[string]                                                      `header:"Idempotency-Key"`
}

func (r PrismObjectMessageGrantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How much of the record the grant exposes. `metadata` shares only the record's
// headers and participants; `full` shares its contents. Currently recorded on the
// access row and returned on read — it is not yet enforced by the read path.
// Applies to `message` grants; ignored for other object types.
type PrismObjectMessageGrantUpdateParamsShareLevel string

const (
	PrismObjectMessageGrantUpdateParamsShareLevelMetadata PrismObjectMessageGrantUpdateParamsShareLevel = "metadata"
	PrismObjectMessageGrantUpdateParamsShareLevelFull     PrismObjectMessageGrantUpdateParamsShareLevel = "full"
)

func (r PrismObjectMessageGrantUpdateParamsShareLevel) IsKnown() bool {
	switch r {
	case PrismObjectMessageGrantUpdateParamsShareLevelMetadata, PrismObjectMessageGrantUpdateParamsShareLevelFull:
		return true
	}
	return false
}

type PrismObjectMessageGrantUpdateParamsTeamGroupID string

const (
	PrismObjectMessageGrantUpdateParamsTeamGroupIDA PrismObjectMessageGrantUpdateParamsTeamGroupID = "a"
	PrismObjectMessageGrantUpdateParamsTeamGroupIDR PrismObjectMessageGrantUpdateParamsTeamGroupID = "r"
	PrismObjectMessageGrantUpdateParamsTeamGroupIDW PrismObjectMessageGrantUpdateParamsTeamGroupID = "w"
)

func (r PrismObjectMessageGrantUpdateParamsTeamGroupID) IsKnown() bool {
	switch r {
	case PrismObjectMessageGrantUpdateParamsTeamGroupIDA, PrismObjectMessageGrantUpdateParamsTeamGroupIDR, PrismObjectMessageGrantUpdateParamsTeamGroupIDW:
		return true
	}
	return false
}

type PrismObjectMessageGrantUpdateParamsTeamID string

const (
	PrismObjectMessageGrantUpdateParamsTeamIDA PrismObjectMessageGrantUpdateParamsTeamID = "a"
	PrismObjectMessageGrantUpdateParamsTeamIDR PrismObjectMessageGrantUpdateParamsTeamID = "r"
	PrismObjectMessageGrantUpdateParamsTeamIDW PrismObjectMessageGrantUpdateParamsTeamID = "w"
)

func (r PrismObjectMessageGrantUpdateParamsTeamID) IsKnown() bool {
	switch r {
	case PrismObjectMessageGrantUpdateParamsTeamIDA, PrismObjectMessageGrantUpdateParamsTeamIDR, PrismObjectMessageGrantUpdateParamsTeamIDW:
		return true
	}
	return false
}

type PrismObjectMessageGrantUpdateParamsUserID string

const (
	PrismObjectMessageGrantUpdateParamsUserIDA PrismObjectMessageGrantUpdateParamsUserID = "a"
	PrismObjectMessageGrantUpdateParamsUserIDR PrismObjectMessageGrantUpdateParamsUserID = "r"
	PrismObjectMessageGrantUpdateParamsUserIDW PrismObjectMessageGrantUpdateParamsUserID = "w"
)

func (r PrismObjectMessageGrantUpdateParamsUserID) IsKnown() bool {
	switch r {
	case PrismObjectMessageGrantUpdateParamsUserIDA, PrismObjectMessageGrantUpdateParamsUserIDR, PrismObjectMessageGrantUpdateParamsUserIDW:
		return true
	}
	return false
}

type PrismObjectMessageGrantGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
