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

// EventService contains methods and other services that help with interacting with
// the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	return
}

// List Events
func (r *EventService) List(ctx context.Context, params EventListParams, opts ...option.RequestOption) (res *EventListResponse, err error) {
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

type EventListResponse struct {
	Data       []interface{}         `json:"data"`
	NextCursor string                `json:"next_cursor" api:"nullable"`
	Total      int64                 `json:"total"`
	JSON       eventListResponseJSON `json:"-"`
}

// eventListResponseJSON contains the JSON metadata for the struct
// [EventListResponse]
type eventListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseJSON) RawJSON() string {
	return r.raw
}

type EventListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                 `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[EventListParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[EventListParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]               `json:"boxes"`
	Deleted param.Field[bool]                   `json:"deleted"`
	Sources param.Field[[]string]               `json:"sources" format:"uuid"`
}

func (r EventListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventListParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[EventListParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                         `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values must be option slugs
	Filter param.Field[[]map[string]map[string]EventListParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                   `json:"limit"`
	Page   param.Field[int64]                                                   `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]EventListParamsQuerySort] `json:"sort"`
}

func (r EventListParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type EventListParamsQueryCombinator string

const (
	EventListParamsQueryCombinatorAnd EventListParamsQueryCombinator = "AND"
	EventListParamsQueryCombinatorOr  EventListParamsQueryCombinator = "OR"
)

func (r EventListParamsQueryCombinator) IsKnown() bool {
	switch r {
	case EventListParamsQueryCombinatorAnd, EventListParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [EventListParamsQueryFilterArray].
type EventListParamsQueryFilterUnion interface {
	ImplementsEventListParamsQueryFilterUnion()
}

type EventListParamsQueryFilterArray []string

func (r EventListParamsQueryFilterArray) ImplementsEventListParamsQueryFilterUnion() {}

type EventListParamsQuerySort string

const (
	EventListParamsQuerySortAsc  EventListParamsQuerySort = "asc"
	EventListParamsQuerySortDesc EventListParamsQuerySort = "desc"
)

func (r EventListParamsQuerySort) IsKnown() bool {
	switch r {
	case EventListParamsQuerySortAsc, EventListParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [EventListParamsIDArray].
type EventListParamsIDUnion interface {
	ImplementsEventListParamsIDUnion()
}

type EventListParamsIDArray []string

func (r EventListParamsIDArray) ImplementsEventListParamsIDUnion() {}
