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

// ViewService contains methods and other services that help with interacting with
// the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewViewService] method instead.
type ViewService struct {
	Options []option.RequestOption
	Records *ViewRecordService
}

// NewViewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewViewService(opts ...option.RequestOption) (r *ViewService) {
	r = &ViewService{}
	r.Options = opts
	r.Records = NewViewRecordService(opts...)
	return
}

// Create a view bundle (view + select/filter/sort)
func (r *ViewService) New(ctx context.Context, viewObjectType ViewNewParamsViewObjectType, params ViewNewParams, opts ...option.RequestOption) (res *ViewNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/view/%v", params.PathTeamID, viewObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a view bundle (select/filter/sort arrays are replaced wholesale when
// supplied)
func (r *ViewService) Update(ctx context.Context, viewObjectType ViewUpdateParamsViewObjectType, viewID string, params ViewUpdateParams, opts ...option.RequestOption) (res *ViewUpdateResponse, err error) {
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
	if viewID == "" {
		err = errors.New("missing required viewId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/view/%v/%s", params.PathTeamID, viewObjectType, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Delete a view bundle
func (r *ViewService) Delete(ctx context.Context, viewObjectType ViewDeleteParamsViewObjectType, viewID string, body ViewDeleteParams, opts ...option.RequestOption) (err error) {
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
	if viewID == "" {
		err = errors.New("missing required viewId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/view/%v/%s", body.TeamID, viewObjectType, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Read a view bundle
func (r *ViewService) Get(ctx context.Context, viewObjectType ViewGetParamsViewObjectType, viewID string, query ViewGetParams, opts ...option.RequestOption) (res *ViewGetResponse, err error) {
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
	if viewID == "" {
		err = errors.New("missing required viewId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/view/%v/%s", query.TeamID, viewObjectType, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A view (saved configuration for displaying records of a given object type) plus
// its select/filter/sort children. Properties in select/filter/sort are referenced
// by slug.
type ViewNewResponse struct {
	Name                 string                    `json:"name" api:"required"`
	ViewType             string                    `json:"view_type" api:"required"`
	ID                   string                    `json:"id" format:"uuid"`
	AggregationPropDefID string                    `json:"aggregation_prop_def_id" api:"nullable" format:"uuid"`
	AggregationType      string                    `json:"aggregation_type" api:"nullable"`
	ColumnLayout         map[string]interface{}    `json:"column_layout" api:"nullable"`
	Combinator           ViewNewResponseCombinator `json:"combinator"`
	CreatedAt            string                    `json:"created_at"`
	CRMID                string                    `json:"crm_id" api:"nullable" format:"uuid"`
	// Each entry is { slug: { comparator: value } }
	Filter []map[string]interface{} `json:"filter"`
	// Property slug to group by
	GroupBy              string        `json:"group_by" api:"nullable"`
	GroupHiddenOptionIDs []interface{} `json:"group_hidden_option_ids" api:"nullable"`
	GroupHideEmpty       bool          `json:"group_hide_empty" api:"nullable"`
	GroupSort            string        `json:"group_sort" api:"nullable"`
	Icon                 string        `json:"icon" api:"nullable"`
	// Property slugs (dot-paths permitted for refs)
	Select []string `json:"select"`
	// Each entry is { slug: 'asc' | 'desc' }
	Sort      []map[string]interface{} `json:"sort"`
	SortOrder int64                    `json:"sort_order" api:"nullable"`
	TeamID    string                   `json:"team_id" api:"nullable" format:"uuid"`
	UpdatedAt string                   `json:"updated_at" api:"nullable"`
	UserID    string                   `json:"user_id" api:"nullable"`
	JSON      viewNewResponseJSON      `json:"-"`
}

// viewNewResponseJSON contains the JSON metadata for the struct [ViewNewResponse]
type viewNewResponseJSON struct {
	Name                 apijson.Field
	ViewType             apijson.Field
	ID                   apijson.Field
	AggregationPropDefID apijson.Field
	AggregationType      apijson.Field
	ColumnLayout         apijson.Field
	Combinator           apijson.Field
	CreatedAt            apijson.Field
	CRMID                apijson.Field
	Filter               apijson.Field
	GroupBy              apijson.Field
	GroupHiddenOptionIDs apijson.Field
	GroupHideEmpty       apijson.Field
	GroupSort            apijson.Field
	Icon                 apijson.Field
	Select               apijson.Field
	Sort                 apijson.Field
	SortOrder            apijson.Field
	TeamID               apijson.Field
	UpdatedAt            apijson.Field
	UserID               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ViewNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewNewResponseJSON) RawJSON() string {
	return r.raw
}

type ViewNewResponseCombinator string

const (
	ViewNewResponseCombinatorAnd ViewNewResponseCombinator = "AND"
	ViewNewResponseCombinatorOr  ViewNewResponseCombinator = "OR"
)

func (r ViewNewResponseCombinator) IsKnown() bool {
	switch r {
	case ViewNewResponseCombinatorAnd, ViewNewResponseCombinatorOr:
		return true
	}
	return false
}

// A view (saved configuration for displaying records of a given object type) plus
// its select/filter/sort children. Properties in select/filter/sort are referenced
// by slug.
type ViewUpdateResponse struct {
	Name                 string                       `json:"name" api:"required"`
	ViewType             string                       `json:"view_type" api:"required"`
	ID                   string                       `json:"id" format:"uuid"`
	AggregationPropDefID string                       `json:"aggregation_prop_def_id" api:"nullable" format:"uuid"`
	AggregationType      string                       `json:"aggregation_type" api:"nullable"`
	ColumnLayout         map[string]interface{}       `json:"column_layout" api:"nullable"`
	Combinator           ViewUpdateResponseCombinator `json:"combinator"`
	CreatedAt            string                       `json:"created_at"`
	CRMID                string                       `json:"crm_id" api:"nullable" format:"uuid"`
	// Each entry is { slug: { comparator: value } }
	Filter []map[string]interface{} `json:"filter"`
	// Property slug to group by
	GroupBy              string        `json:"group_by" api:"nullable"`
	GroupHiddenOptionIDs []interface{} `json:"group_hidden_option_ids" api:"nullable"`
	GroupHideEmpty       bool          `json:"group_hide_empty" api:"nullable"`
	GroupSort            string        `json:"group_sort" api:"nullable"`
	Icon                 string        `json:"icon" api:"nullable"`
	// Property slugs (dot-paths permitted for refs)
	Select []string `json:"select"`
	// Each entry is { slug: 'asc' | 'desc' }
	Sort      []map[string]interface{} `json:"sort"`
	SortOrder int64                    `json:"sort_order" api:"nullable"`
	TeamID    string                   `json:"team_id" api:"nullable" format:"uuid"`
	UpdatedAt string                   `json:"updated_at" api:"nullable"`
	UserID    string                   `json:"user_id" api:"nullable"`
	JSON      viewUpdateResponseJSON   `json:"-"`
}

// viewUpdateResponseJSON contains the JSON metadata for the struct
// [ViewUpdateResponse]
type viewUpdateResponseJSON struct {
	Name                 apijson.Field
	ViewType             apijson.Field
	ID                   apijson.Field
	AggregationPropDefID apijson.Field
	AggregationType      apijson.Field
	ColumnLayout         apijson.Field
	Combinator           apijson.Field
	CreatedAt            apijson.Field
	CRMID                apijson.Field
	Filter               apijson.Field
	GroupBy              apijson.Field
	GroupHiddenOptionIDs apijson.Field
	GroupHideEmpty       apijson.Field
	GroupSort            apijson.Field
	Icon                 apijson.Field
	Select               apijson.Field
	Sort                 apijson.Field
	SortOrder            apijson.Field
	TeamID               apijson.Field
	UpdatedAt            apijson.Field
	UserID               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ViewUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ViewUpdateResponseCombinator string

const (
	ViewUpdateResponseCombinatorAnd ViewUpdateResponseCombinator = "AND"
	ViewUpdateResponseCombinatorOr  ViewUpdateResponseCombinator = "OR"
)

func (r ViewUpdateResponseCombinator) IsKnown() bool {
	switch r {
	case ViewUpdateResponseCombinatorAnd, ViewUpdateResponseCombinatorOr:
		return true
	}
	return false
}

// A view (saved configuration for displaying records of a given object type) plus
// its select/filter/sort children. Properties in select/filter/sort are referenced
// by slug.
type ViewGetResponse struct {
	Name                 string                    `json:"name" api:"required"`
	ViewType             string                    `json:"view_type" api:"required"`
	ID                   string                    `json:"id" format:"uuid"`
	AggregationPropDefID string                    `json:"aggregation_prop_def_id" api:"nullable" format:"uuid"`
	AggregationType      string                    `json:"aggregation_type" api:"nullable"`
	ColumnLayout         map[string]interface{}    `json:"column_layout" api:"nullable"`
	Combinator           ViewGetResponseCombinator `json:"combinator"`
	CreatedAt            string                    `json:"created_at"`
	CRMID                string                    `json:"crm_id" api:"nullable" format:"uuid"`
	// Each entry is { slug: { comparator: value } }
	Filter []map[string]interface{} `json:"filter"`
	// Property slug to group by
	GroupBy              string        `json:"group_by" api:"nullable"`
	GroupHiddenOptionIDs []interface{} `json:"group_hidden_option_ids" api:"nullable"`
	GroupHideEmpty       bool          `json:"group_hide_empty" api:"nullable"`
	GroupSort            string        `json:"group_sort" api:"nullable"`
	Icon                 string        `json:"icon" api:"nullable"`
	// Property slugs (dot-paths permitted for refs)
	Select []string `json:"select"`
	// Each entry is { slug: 'asc' | 'desc' }
	Sort      []map[string]interface{} `json:"sort"`
	SortOrder int64                    `json:"sort_order" api:"nullable"`
	TeamID    string                   `json:"team_id" api:"nullable" format:"uuid"`
	UpdatedAt string                   `json:"updated_at" api:"nullable"`
	UserID    string                   `json:"user_id" api:"nullable"`
	JSON      viewGetResponseJSON      `json:"-"`
}

// viewGetResponseJSON contains the JSON metadata for the struct [ViewGetResponse]
type viewGetResponseJSON struct {
	Name                 apijson.Field
	ViewType             apijson.Field
	ID                   apijson.Field
	AggregationPropDefID apijson.Field
	AggregationType      apijson.Field
	ColumnLayout         apijson.Field
	Combinator           apijson.Field
	CreatedAt            apijson.Field
	CRMID                apijson.Field
	Filter               apijson.Field
	GroupBy              apijson.Field
	GroupHiddenOptionIDs apijson.Field
	GroupHideEmpty       apijson.Field
	GroupSort            apijson.Field
	Icon                 apijson.Field
	Select               apijson.Field
	Sort                 apijson.Field
	SortOrder            apijson.Field
	TeamID               apijson.Field
	UpdatedAt            apijson.Field
	UserID               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ViewGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewGetResponseJSON) RawJSON() string {
	return r.raw
}

type ViewGetResponseCombinator string

const (
	ViewGetResponseCombinatorAnd ViewGetResponseCombinator = "AND"
	ViewGetResponseCombinatorOr  ViewGetResponseCombinator = "OR"
)

func (r ViewGetResponseCombinator) IsKnown() bool {
	switch r {
	case ViewGetResponseCombinatorAnd, ViewGetResponseCombinatorOr:
		return true
	}
	return false
}

type ViewNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	PathTeamID           param.Field[string]                  `path:"teamId" api:"required" format:"uuid"`
	Name                 param.Field[string]                  `json:"name" api:"required"`
	ViewType             param.Field[string]                  `json:"view_type" api:"required"`
	ID                   param.Field[string]                  `json:"id" format:"uuid"`
	AggregationPropDefID param.Field[string]                  `json:"aggregation_prop_def_id" format:"uuid"`
	AggregationType      param.Field[string]                  `json:"aggregation_type"`
	ColumnLayout         param.Field[map[string]interface{}]  `json:"column_layout"`
	Combinator           param.Field[ViewNewParamsCombinator] `json:"combinator"`
	CreatedAt            param.Field[string]                  `json:"created_at"`
	CRMID                param.Field[string]                  `json:"crm_id" format:"uuid"`
	// Each entry is { slug: { comparator: value } }
	Filter param.Field[[]map[string]interface{}] `json:"filter"`
	// Property slug to group by
	GroupBy              param.Field[string]        `json:"group_by"`
	GroupHiddenOptionIDs param.Field[[]interface{}] `json:"group_hidden_option_ids"`
	GroupHideEmpty       param.Field[bool]          `json:"group_hide_empty"`
	GroupSort            param.Field[string]        `json:"group_sort"`
	Icon                 param.Field[string]        `json:"icon"`
	// Property slugs (dot-paths permitted for refs)
	Select param.Field[[]string] `json:"select"`
	// Each entry is { slug: 'asc' | 'desc' }
	Sort       param.Field[[]map[string]interface{}] `json:"sort"`
	SortOrder  param.Field[int64]                    `json:"sort_order"`
	BodyTeamID param.Field[string]                   `json:"team_id" format:"uuid"`
	UpdatedAt  param.Field[string]                   `json:"updated_at"`
	UserID     param.Field[string]                   `json:"user_id"`
}

func (r ViewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ViewNewParamsViewObjectType string

const (
	ViewNewParamsViewObjectTypeAction       ViewNewParamsViewObjectType = "action"
	ViewNewParamsViewObjectTypeDeal         ViewNewParamsViewObjectType = "deal"
	ViewNewParamsViewObjectTypeDocument     ViewNewParamsViewObjectType = "document"
	ViewNewParamsViewObjectTypeEvent        ViewNewParamsViewObjectType = "event"
	ViewNewParamsViewObjectTypeIdentity     ViewNewParamsViewObjectType = "identity"
	ViewNewParamsViewObjectTypeOrganization ViewNewParamsViewObjectType = "organization"
)

func (r ViewNewParamsViewObjectType) IsKnown() bool {
	switch r {
	case ViewNewParamsViewObjectTypeAction, ViewNewParamsViewObjectTypeDeal, ViewNewParamsViewObjectTypeDocument, ViewNewParamsViewObjectTypeEvent, ViewNewParamsViewObjectTypeIdentity, ViewNewParamsViewObjectTypeOrganization:
		return true
	}
	return false
}

type ViewNewParamsCombinator string

const (
	ViewNewParamsCombinatorAnd ViewNewParamsCombinator = "AND"
	ViewNewParamsCombinatorOr  ViewNewParamsCombinator = "OR"
)

func (r ViewNewParamsCombinator) IsKnown() bool {
	switch r {
	case ViewNewParamsCombinatorAnd, ViewNewParamsCombinatorOr:
		return true
	}
	return false
}

type ViewUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	PathTeamID           param.Field[string]                     `path:"teamId" api:"required" format:"uuid"`
	AggregationPropDefID param.Field[string]                     `json:"aggregation_prop_def_id" format:"uuid"`
	AggregationType      param.Field[string]                     `json:"aggregation_type"`
	ColumnLayout         param.Field[map[string]interface{}]     `json:"column_layout"`
	Combinator           param.Field[ViewUpdateParamsCombinator] `json:"combinator"`
	CRMID                param.Field[string]                     `json:"crm_id" format:"uuid"`
	Filter               param.Field[[]map[string]interface{}]   `json:"filter"`
	GroupBy              param.Field[string]                     `json:"group_by"`
	GroupHiddenOptionIDs param.Field[[]interface{}]              `json:"group_hidden_option_ids"`
	GroupHideEmpty       param.Field[bool]                       `json:"group_hide_empty"`
	GroupSort            param.Field[string]                     `json:"group_sort"`
	Icon                 param.Field[string]                     `json:"icon"`
	Name                 param.Field[string]                     `json:"name"`
	Select               param.Field[[]string]                   `json:"select"`
	Sort                 param.Field[[]map[string]interface{}]   `json:"sort"`
	SortOrder            param.Field[int64]                      `json:"sort_order"`
	BodyTeamID           param.Field[string]                     `json:"team_id" format:"uuid"`
	UserID               param.Field[string]                     `json:"user_id"`
	ViewType             param.Field[string]                     `json:"view_type"`
}

func (r ViewUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ViewUpdateParamsViewObjectType string

const (
	ViewUpdateParamsViewObjectTypeAction       ViewUpdateParamsViewObjectType = "action"
	ViewUpdateParamsViewObjectTypeDeal         ViewUpdateParamsViewObjectType = "deal"
	ViewUpdateParamsViewObjectTypeDocument     ViewUpdateParamsViewObjectType = "document"
	ViewUpdateParamsViewObjectTypeEvent        ViewUpdateParamsViewObjectType = "event"
	ViewUpdateParamsViewObjectTypeIdentity     ViewUpdateParamsViewObjectType = "identity"
	ViewUpdateParamsViewObjectTypeOrganization ViewUpdateParamsViewObjectType = "organization"
)

func (r ViewUpdateParamsViewObjectType) IsKnown() bool {
	switch r {
	case ViewUpdateParamsViewObjectTypeAction, ViewUpdateParamsViewObjectTypeDeal, ViewUpdateParamsViewObjectTypeDocument, ViewUpdateParamsViewObjectTypeEvent, ViewUpdateParamsViewObjectTypeIdentity, ViewUpdateParamsViewObjectTypeOrganization:
		return true
	}
	return false
}

type ViewUpdateParamsCombinator string

const (
	ViewUpdateParamsCombinatorAnd ViewUpdateParamsCombinator = "AND"
	ViewUpdateParamsCombinatorOr  ViewUpdateParamsCombinator = "OR"
)

func (r ViewUpdateParamsCombinator) IsKnown() bool {
	switch r {
	case ViewUpdateParamsCombinatorAnd, ViewUpdateParamsCombinatorOr:
		return true
	}
	return false
}

type ViewDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type ViewDeleteParamsViewObjectType string

const (
	ViewDeleteParamsViewObjectTypeAction       ViewDeleteParamsViewObjectType = "action"
	ViewDeleteParamsViewObjectTypeDeal         ViewDeleteParamsViewObjectType = "deal"
	ViewDeleteParamsViewObjectTypeDocument     ViewDeleteParamsViewObjectType = "document"
	ViewDeleteParamsViewObjectTypeEvent        ViewDeleteParamsViewObjectType = "event"
	ViewDeleteParamsViewObjectTypeIdentity     ViewDeleteParamsViewObjectType = "identity"
	ViewDeleteParamsViewObjectTypeOrganization ViewDeleteParamsViewObjectType = "organization"
)

func (r ViewDeleteParamsViewObjectType) IsKnown() bool {
	switch r {
	case ViewDeleteParamsViewObjectTypeAction, ViewDeleteParamsViewObjectTypeDeal, ViewDeleteParamsViewObjectTypeDocument, ViewDeleteParamsViewObjectTypeEvent, ViewDeleteParamsViewObjectTypeIdentity, ViewDeleteParamsViewObjectTypeOrganization:
		return true
	}
	return false
}

type ViewGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type ViewGetParamsViewObjectType string

const (
	ViewGetParamsViewObjectTypeAction       ViewGetParamsViewObjectType = "action"
	ViewGetParamsViewObjectTypeDeal         ViewGetParamsViewObjectType = "deal"
	ViewGetParamsViewObjectTypeDocument     ViewGetParamsViewObjectType = "document"
	ViewGetParamsViewObjectTypeEvent        ViewGetParamsViewObjectType = "event"
	ViewGetParamsViewObjectTypeIdentity     ViewGetParamsViewObjectType = "identity"
	ViewGetParamsViewObjectTypeOrganization ViewGetParamsViewObjectType = "organization"
)

func (r ViewGetParamsViewObjectType) IsKnown() bool {
	switch r {
	case ViewGetParamsViewObjectTypeAction, ViewGetParamsViewObjectTypeDeal, ViewGetParamsViewObjectTypeDocument, ViewGetParamsViewObjectTypeEvent, ViewGetParamsViewObjectTypeIdentity, ViewGetParamsViewObjectTypeOrganization:
		return true
	}
	return false
}
