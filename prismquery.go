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

// PrismQueryService contains methods and other services that help with interacting
// with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismQueryService] method instead.
type PrismQueryService struct {
	Options []option.RequestOption
}

// NewPrismQueryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrismQueryService(opts ...option.RequestOption) (r *PrismQueryService) {
	r = &PrismQueryService{}
	r.Options = opts
	return
}

// Query v2
func (r *PrismQueryService) Execute(ctx context.Context, objectType PrismQueryExecuteParamsObjectType, params PrismQueryExecuteParams, opts ...option.RequestOption) (res *PrismQueryExecuteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.TeamID, precfg.TeamID)
	if params.TeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/query/%s/%v", params.TeamID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type PrismQueryExecuteResponse struct {
	Data  []interface{}                 `json:"data"`
	Total int64                         `json:"total"`
	JSON  prismQueryExecuteResponseJSON `json:"-"`
}

// prismQueryExecuteResponseJSON contains the JSON metadata for the struct
// [PrismQueryExecuteResponse]
type prismQueryExecuteResponseJSON struct {
	Data        apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismQueryExecuteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismQueryExecuteResponseJSON) RawJSON() string {
	return r.raw
}

type PrismQueryExecuteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                         `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[PrismQueryExecuteParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[PrismQueryExecuteParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                       `json:"boxes"`
	Deleted param.Field[bool]                           `json:"deleted"`
	Sources param.Field[[]string]                       `json:"sources" format:"uuid"`
}

func (r PrismQueryExecuteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismQueryExecuteParamsObjectType string

const (
	PrismQueryExecuteParamsObjectTypeDeal          PrismQueryExecuteParamsObjectType = "deal"
	PrismQueryExecuteParamsObjectTypeIdentity      PrismQueryExecuteParamsObjectType = "identity"
	PrismQueryExecuteParamsObjectTypeAIChatThread  PrismQueryExecuteParamsObjectType = "ai_chat_thread"
	PrismQueryExecuteParamsObjectTypeAIChatMessage PrismQueryExecuteParamsObjectType = "ai_chat_message"
	PrismQueryExecuteParamsObjectTypeDocument      PrismQueryExecuteParamsObjectType = "document"
	PrismQueryExecuteParamsObjectTypeOrganization  PrismQueryExecuteParamsObjectType = "organization"
	PrismQueryExecuteParamsObjectTypeContact       PrismQueryExecuteParamsObjectType = "contact"
	PrismQueryExecuteParamsObjectTypeAction        PrismQueryExecuteParamsObjectType = "action"
	PrismQueryExecuteParamsObjectTypeEvent         PrismQueryExecuteParamsObjectType = "event"
)

func (r PrismQueryExecuteParamsObjectType) IsKnown() bool {
	switch r {
	case PrismQueryExecuteParamsObjectTypeDeal, PrismQueryExecuteParamsObjectTypeIdentity, PrismQueryExecuteParamsObjectTypeAIChatThread, PrismQueryExecuteParamsObjectTypeAIChatMessage, PrismQueryExecuteParamsObjectTypeDocument, PrismQueryExecuteParamsObjectTypeOrganization, PrismQueryExecuteParamsObjectTypeContact, PrismQueryExecuteParamsObjectTypeAction, PrismQueryExecuteParamsObjectTypeEvent:
		return true
	}
	return false
}

type PrismQueryExecuteParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismQueryExecuteParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                                 `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values must be option slugs
	Filter param.Field[[]map[string]map[string]PrismQueryExecuteParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                           `json:"limit"`
	Page   param.Field[int64]                                                           `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]PrismQueryExecuteParamsQuerySort] `json:"sort"`
}

func (r PrismQueryExecuteParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismQueryExecuteParamsQueryCombinator string

const (
	PrismQueryExecuteParamsQueryCombinatorAnd PrismQueryExecuteParamsQueryCombinator = "AND"
	PrismQueryExecuteParamsQueryCombinatorOr  PrismQueryExecuteParamsQueryCombinator = "OR"
)

func (r PrismQueryExecuteParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismQueryExecuteParamsQueryCombinatorAnd, PrismQueryExecuteParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismQueryExecuteParamsQueryFilterArray].
type PrismQueryExecuteParamsQueryFilterUnion interface {
	ImplementsPrismQueryExecuteParamsQueryFilterUnion()
}

type PrismQueryExecuteParamsQueryFilterArray []string

func (r PrismQueryExecuteParamsQueryFilterArray) ImplementsPrismQueryExecuteParamsQueryFilterUnion() {
}

type PrismQueryExecuteParamsQuerySort string

const (
	PrismQueryExecuteParamsQuerySortAsc  PrismQueryExecuteParamsQuerySort = "asc"
	PrismQueryExecuteParamsQuerySortDesc PrismQueryExecuteParamsQuerySort = "desc"
)

func (r PrismQueryExecuteParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismQueryExecuteParamsQuerySortAsc, PrismQueryExecuteParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismQueryExecuteParamsIDArray].
type PrismQueryExecuteParamsIDUnion interface {
	ImplementsPrismQueryExecuteParamsIDUnion()
}

type PrismQueryExecuteParamsIDArray []string

func (r PrismQueryExecuteParamsIDArray) ImplementsPrismQueryExecuteParamsIDUnion() {}
