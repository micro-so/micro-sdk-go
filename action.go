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

// Actions are tasks and to-dos that can be assigned to contacts, organizations, or
// deals, with a status, due date, and priority.
//
// ActionService contains methods and other services that help with interacting
// with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionService] method instead.
type ActionService struct {
	Options []option.RequestOption
}

// NewActionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewActionService(opts ...option.RequestOption) (r *ActionService) {
	r = &ActionService{}
	r.Options = opts
	return
}

// Create Action
func (r *ActionService) New(ctx context.Context, params ActionNewParams, opts ...option.RequestOption) (res *ActionNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/action", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update Action
func (r *ActionService) Update(ctx context.Context, actionID string, params ActionUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return err
	}
	requestconfig.UseDefaultParam(&params.TeamID, precfg.TeamID)
	if params.TeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return err
	}
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/action/%s", params.TeamID, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return err
}

// List Actions
func (r *ActionService) List(ctx context.Context, params ActionListParams, opts ...option.RequestOption) (res *ActionListResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/query/%s/action", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Delete Action
func (r *ActionService) Delete(ctx context.Context, actionID string, body ActionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return err
	}
	requestconfig.UseDefaultParam(&body.TeamID, precfg.TeamID)
	if body.TeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return err
	}
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/action/%s", body.TeamID, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ActionNewResponse struct {
	ID   string                `json:"id" format:"uuid"`
	JSON actionNewResponseJSON `json:"-"`
}

// actionNewResponseJSON contains the JSON metadata for the struct
// [ActionNewResponse]
type actionNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r actionNewResponseJSON) RawJSON() string {
	return r.raw
}

type ActionListResponse struct {
	Data       []interface{}          `json:"data"`
	NextCursor string                 `json:"next_cursor" api:"nullable"`
	Total      int64                  `json:"total"`
	JSON       actionListResponseJSON `json:"-"`
}

// actionListResponseJSON contains the JSON metadata for the struct
// [ActionListResponse]
type actionListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r actionListResponseJSON) RawJSON() string {
	return r.raw
}

type ActionNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r ActionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type ActionUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r ActionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type ActionListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                  `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[ActionListParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[ActionListParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                `json:"boxes"`
	Deleted param.Field[bool]                    `json:"deleted"`
	Sources param.Field[[]string]                `json:"sources" format:"uuid"`
}

func (r ActionListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ActionListParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[ActionListParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                          `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values must be option slugs
	Filter param.Field[[]map[string]map[string]ActionListParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                    `json:"limit"`
	Page   param.Field[int64]                                                    `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]ActionListParamsQuerySort] `json:"sort"`
}

func (r ActionListParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type ActionListParamsQueryCombinator string

const (
	ActionListParamsQueryCombinatorAnd ActionListParamsQueryCombinator = "AND"
	ActionListParamsQueryCombinatorOr  ActionListParamsQueryCombinator = "OR"
)

func (r ActionListParamsQueryCombinator) IsKnown() bool {
	switch r {
	case ActionListParamsQueryCombinatorAnd, ActionListParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [ActionListParamsQueryFilterArray].
type ActionListParamsQueryFilterUnion interface {
	ImplementsActionListParamsQueryFilterUnion()
}

type ActionListParamsQueryFilterArray []string

func (r ActionListParamsQueryFilterArray) ImplementsActionListParamsQueryFilterUnion() {}

type ActionListParamsQuerySort string

const (
	ActionListParamsQuerySortAsc  ActionListParamsQuerySort = "asc"
	ActionListParamsQuerySortDesc ActionListParamsQuerySort = "desc"
)

func (r ActionListParamsQuerySort) IsKnown() bool {
	switch r {
	case ActionListParamsQuerySortAsc, ActionListParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [ActionListParamsIDArray].
type ActionListParamsIDUnion interface {
	ImplementsActionListParamsIDUnion()
}

type ActionListParamsIDArray []string

func (r ActionListParamsIDArray) ImplementsActionListParamsIDUnion() {}

type ActionDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
