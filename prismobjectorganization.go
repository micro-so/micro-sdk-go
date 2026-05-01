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

// PrismObjectOrganizationService contains methods and other services that help
// with interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectOrganizationService] method instead.
type PrismObjectOrganizationService struct {
	Options []option.RequestOption
}

// NewPrismObjectOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectOrganizationService(opts ...option.RequestOption) (r *PrismObjectOrganizationService) {
	r = &PrismObjectOrganizationService{}
	r.Options = opts
	return
}

// Query v2
func (r *PrismObjectOrganizationService) Query(ctx context.Context, params PrismObjectOrganizationQueryParams, opts ...option.RequestOption) (res *PrismObjectOrganizationQueryResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/query/%s/organization", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type PrismObjectOrganizationQueryResponse struct {
	Data  []interface{}                            `json:"data"`
	Total int64                                    `json:"total"`
	JSON  prismObjectOrganizationQueryResponseJSON `json:"-"`
}

// prismObjectOrganizationQueryResponseJSON contains the JSON metadata for the
// struct [PrismObjectOrganizationQueryResponse]
type prismObjectOrganizationQueryResponseJSON struct {
	Data        apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismObjectOrganizationQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismObjectOrganizationQueryResponseJSON) RawJSON() string {
	return r.raw
}

type PrismObjectOrganizationQueryParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                                    `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[PrismObjectOrganizationQueryParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[PrismObjectOrganizationQueryParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                                  `json:"boxes"`
	Deleted param.Field[bool]                                      `json:"deleted"`
	Sources param.Field[[]string]                                  `json:"sources" format:"uuid"`
}

func (r PrismObjectOrganizationQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismObjectOrganizationQueryParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[PrismObjectOrganizationQueryParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                                            `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values must be option slugs
	Filter param.Field[[]map[string]map[string]PrismObjectOrganizationQueryParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                                      `json:"limit"`
	Page   param.Field[int64]                                                                      `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]PrismObjectOrganizationQueryParamsQuerySort] `json:"sort"`
}

func (r PrismObjectOrganizationQueryParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type PrismObjectOrganizationQueryParamsQueryCombinator string

const (
	PrismObjectOrganizationQueryParamsQueryCombinatorAnd PrismObjectOrganizationQueryParamsQueryCombinator = "AND"
	PrismObjectOrganizationQueryParamsQueryCombinatorOr  PrismObjectOrganizationQueryParamsQueryCombinator = "OR"
)

func (r PrismObjectOrganizationQueryParamsQueryCombinator) IsKnown() bool {
	switch r {
	case PrismObjectOrganizationQueryParamsQueryCombinatorAnd, PrismObjectOrganizationQueryParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [PrismObjectOrganizationQueryParamsQueryFilterArray].
type PrismObjectOrganizationQueryParamsQueryFilterUnion interface {
	ImplementsPrismObjectOrganizationQueryParamsQueryFilterUnion()
}

type PrismObjectOrganizationQueryParamsQueryFilterArray []string

func (r PrismObjectOrganizationQueryParamsQueryFilterArray) ImplementsPrismObjectOrganizationQueryParamsQueryFilterUnion() {
}

type PrismObjectOrganizationQueryParamsQuerySort string

const (
	PrismObjectOrganizationQueryParamsQuerySortAsc  PrismObjectOrganizationQueryParamsQuerySort = "asc"
	PrismObjectOrganizationQueryParamsQuerySortDesc PrismObjectOrganizationQueryParamsQuerySort = "desc"
)

func (r PrismObjectOrganizationQueryParamsQuerySort) IsKnown() bool {
	switch r {
	case PrismObjectOrganizationQueryParamsQuerySortAsc, PrismObjectOrganizationQueryParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [PrismObjectOrganizationQueryParamsIDArray].
type PrismObjectOrganizationQueryParamsIDUnion interface {
	ImplementsPrismObjectOrganizationQueryParamsIDUnion()
}

type PrismObjectOrganizationQueryParamsIDArray []string

func (r PrismObjectOrganizationQueryParamsIDArray) ImplementsPrismObjectOrganizationQueryParamsIDUnion() {
}
