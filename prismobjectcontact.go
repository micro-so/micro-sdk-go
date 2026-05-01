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

// PrismObjectContactService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectContactService] method instead.
type PrismObjectContactService struct {
	Options []option.RequestOption
}

// NewPrismObjectContactService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectContactService(opts ...option.RequestOption) (r *PrismObjectContactService) {
	r = &PrismObjectContactService{}
	r.Options = opts
	return
}

// Query v2
func (r *PrismObjectContactService) Query(ctx context.Context, params PrismObjectContactQueryParams, opts ...option.RequestOption) (res *PrismObjectContactQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/query/%s/contact", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type PrismObjectContactQueryResponse struct {
	Data  []interface{}                       `json:"data"`
	Total int64                               `json:"total"`
	JSON  prismObjectContactQueryResponseJSON `json:"-"`
}

// prismObjectContactQueryResponseJSON contains the JSON metadata for the struct
// [PrismObjectContactQueryResponse]
type prismObjectContactQueryResponseJSON struct {
	Data        apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectContactQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectContactQueryResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectContactQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                               `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[PrismObjectContactQueryParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[PrismObjectContactQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                             `json:"boxes"`
	Deleted param.Field[bool]                                 `json:"deleted"`
	Sources param.Field[[]string]                             `json:"sources" format:"uuid"`
}

func (r PrismObjectContactQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectContactQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectContactQueryParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                                       `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values must be option slugs
	Filter param.Field[[]map[string]map[string]PrismObjectContactQueryParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                                 `json:"limit"`
	Page   param.Field[int64]                                                                 `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]PrismObjectContactQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectContactQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectContactQueryParamsQueryCombinator string

const (
	PrismObjectContactQueryParamsQueryCombinatorAnd PrismObjectContactQueryParamsQueryCombinator = "AND"
	PrismObjectContactQueryParamsQueryCombinatorOr  PrismObjectContactQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectContactQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectContactQueryParamsQueryCombinatorAnd, PrismObjectContactQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectContactQueryParamsQueryFilterArray].
type PrismObjectContactQueryParamsQueryFilterUnion interface {
	ImplementsPrismObjectContactQueryParamsQueryFilterUnion()
}

type PrismObjectContactQueryParamsQueryFilterArray []string

func (r PrismObjectContactQueryParamsQueryFilterArray) ImplementsPrismObjectContactQueryParamsQueryFilterUnion() {
}

type PrismObjectContactQueryParamsQuerySort string

const (
	PrismObjectContactQueryParamsQuerySortAsc  PrismObjectContactQueryParamsQuerySort = "asc"
	PrismObjectContactQueryParamsQuerySortDesc PrismObjectContactQueryParamsQuerySort = "desc"
)

func (r PrismObjectContactQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectContactQueryParamsQuerySortAsc, PrismObjectContactQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectContactQueryParamsIDArray].
type PrismObjectContactQueryParamsIDUnion interface {
	ImplementsPrismObjectContactQueryParamsIDUnion()
}

type PrismObjectContactQueryParamsIDArray []string

func (r PrismObjectContactQueryParamsIDArray) ImplementsPrismObjectContactQueryParamsIDUnion() {}
