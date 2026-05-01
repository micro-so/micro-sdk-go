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

// PrismObjectEventService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectEventService] method instead.
type PrismObjectEventService struct {
	Options []option.RequestOption
	Grant   *PrismObjectEventGrantService
}

// NewPrismObjectEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectEventService(opts ...option.RequestOption) (r *PrismObjectEventService) {
	r = &PrismObjectEventService{}
	r.Options = opts
	r.Grant = NewPrismObjectEventGrantService(opts...)
	return
}

// Get object
func (r *PrismObjectEventService) Get(ctx context.Context, eventID string, query PrismObjectEventGetParams, opts ...option.RequestOption) (res *PrismObjectProperties, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/event/%s", query.TeamID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Query v2
func (r *PrismObjectEventService) Query(ctx context.Context, params PrismObjectEventQueryParams, opts ...option.RequestOption) (res *PrismObjectEventQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/query/%s/event", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type PrismObjectEventQueryResponse struct {
	Data  []interface{}                     `json:"data"`
	Total int64                             `json:"total"`
	JSON  prismObjectEventQueryResponseJSON `json:"-"`
}

// prismObjectEventQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectEventQueryResponse]
type prismObjectEventQueryResponseJSON struct {
	Data        apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectEventQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectEventQueryResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectEventGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismObjectEventQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                             `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[PrismObjectEventQueryParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[PrismObjectEventQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                           `json:"boxes"`
	Deleted param.Field[bool]                               `json:"deleted"`
	Sources param.Field[[]string]                           `json:"sources" format:"uuid"`
}

func (r PrismObjectEventQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectEventQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectEventQueryParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                                     `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values must be option slugs
	Filter param.Field[[]map[string]map[string]PrismObjectEventQueryParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                               `json:"limit"`
	Page   param.Field[int64]                                                               `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]PrismObjectEventQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectEventQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectEventQueryParamsQueryCombinator string

const (
	PrismObjectEventQueryParamsQueryCombinatorAnd PrismObjectEventQueryParamsQueryCombinator = "AND"
	PrismObjectEventQueryParamsQueryCombinatorOr  PrismObjectEventQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectEventQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectEventQueryParamsQueryCombinatorAnd, PrismObjectEventQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectEventQueryParamsQueryFilterArray].
type PrismObjectEventQueryParamsQueryFilterUnion interface {
	ImplementsPrismObjectEventQueryParamsQueryFilterUnion()
}

type PrismObjectEventQueryParamsQueryFilterArray []string

func (r PrismObjectEventQueryParamsQueryFilterArray) ImplementsPrismObjectEventQueryParamsQueryFilterUnion() {
}

type PrismObjectEventQueryParamsQuerySort string

const (
	PrismObjectEventQueryParamsQuerySortAsc  PrismObjectEventQueryParamsQuerySort = "asc"
	PrismObjectEventQueryParamsQuerySortDesc PrismObjectEventQueryParamsQuerySort = "desc"
)

func (r PrismObjectEventQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectEventQueryParamsQuerySortAsc, PrismObjectEventQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectEventQueryParamsIDArray].
type PrismObjectEventQueryParamsIDUnion interface {
	ImplementsPrismObjectEventQueryParamsIDUnion()
}

type PrismObjectEventQueryParamsIDArray []string

func (r PrismObjectEventQueryParamsIDArray) ImplementsPrismObjectEventQueryParamsIDUnion() {}
