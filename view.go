// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/micro-so/micro-sdk-go/internal/apijson"
	"github.com/micro-so/micro-sdk-go/internal/apiquery"
	"github.com/micro-so/micro-sdk-go/internal/param"
	"github.com/micro-so/micro-sdk-go/internal/requestconfig"
	"github.com/micro-so/micro-sdk-go/option"
	"github.com/tidwall/gjson"
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
	path := fmt.Sprintf("v2/prism/%s/%v/views", params.PathTeamID, viewObjectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a view bundle (select/filter/sort arrays are replaced wholesale when
// supplied)
func (r *ViewService) Update(ctx context.Context, viewObjectType ViewUpdateParamsViewObjectType, viewID string, params ViewUpdateParams, opts ...option.RequestOption) (res *ViewUpdateResponse, err error) {
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
	if viewID == "" {
		err = errors.New("missing required viewId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/views/%s", params.PathTeamID, viewObjectType, viewID)
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
	path := fmt.Sprintf("v2/prism/%s/%v/views/%s", body.TeamID, viewObjectType, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns the view bundle. Pass `?include=records` to also fetch a page of records
// selected by the view in the same call; the response is then wrapped as
// `{view, records}`.
func (r *ViewService) Get(ctx context.Context, viewObjectType ViewGetParamsViewObjectType, viewID string, params ViewGetParams, opts ...option.RequestOption) (res *ViewGetResponse, err error) {
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
	if viewID == "" {
		err = errors.New("missing required viewId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/views/%s", params.TeamID, viewObjectType, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
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
	// Each entry is { slug: { comparator: value } }
	Filter []map[string]interface{} `json:"filter"`
	// Property slug to group by
	GroupBy              string        `json:"group_by" api:"nullable"`
	GroupHiddenOptionIDs []interface{} `json:"group_hidden_option_ids" api:"nullable"`
	GroupHideEmpty       bool          `json:"group_hide_empty" api:"nullable"`
	GroupSort            string        `json:"group_sort" api:"nullable"`
	Icon                 string        `json:"icon" api:"nullable"`
	ListID               string        `json:"list_id" api:"nullable" format:"uuid"`
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
	Filter               apijson.Field
	GroupBy              apijson.Field
	GroupHiddenOptionIDs apijson.Field
	GroupHideEmpty       apijson.Field
	GroupSort            apijson.Field
	Icon                 apijson.Field
	ListID               apijson.Field
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
	// Each entry is { slug: { comparator: value } }
	Filter []map[string]interface{} `json:"filter"`
	// Property slug to group by
	GroupBy              string        `json:"group_by" api:"nullable"`
	GroupHiddenOptionIDs []interface{} `json:"group_hidden_option_ids" api:"nullable"`
	GroupHideEmpty       bool          `json:"group_hide_empty" api:"nullable"`
	GroupSort            string        `json:"group_sort" api:"nullable"`
	Icon                 string        `json:"icon" api:"nullable"`
	ListID               string        `json:"list_id" api:"nullable" format:"uuid"`
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
	Filter               apijson.Field
	GroupBy              apijson.Field
	GroupHiddenOptionIDs apijson.Field
	GroupHideEmpty       apijson.Field
	GroupSort            apijson.Field
	Icon                 apijson.Field
	ListID               apijson.Field
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
	ID                   string `json:"id" format:"uuid"`
	AggregationPropDefID string `json:"aggregation_prop_def_id" api:"nullable" format:"uuid"`
	AggregationType      string `json:"aggregation_type" api:"nullable"`
	// This field can have the runtime type of [map[string]interface{}].
	ColumnLayout interface{}               `json:"column_layout"`
	Combinator   ViewGetResponseCombinator `json:"combinator"`
	CreatedAt    string                    `json:"created_at"`
	// This field can have the runtime type of [[]map[string]interface{}].
	Filter interface{} `json:"filter"`
	// Property slug to group by
	GroupBy string `json:"group_by" api:"nullable"`
	// This field can have the runtime type of [[]interface{}].
	GroupHiddenOptionIDs interface{} `json:"group_hidden_option_ids"`
	GroupHideEmpty       bool        `json:"group_hide_empty" api:"nullable"`
	GroupSort            string      `json:"group_sort" api:"nullable"`
	Icon                 string      `json:"icon" api:"nullable"`
	ListID               string      `json:"list_id" api:"nullable" format:"uuid"`
	Name                 string      `json:"name"`
	// This field can have the runtime type of
	// [ViewGetResponseViewBundleWithRecordsRecords].
	Records interface{} `json:"records"`
	// This field can have the runtime type of [[]string].
	Select interface{} `json:"select"`
	// This field can have the runtime type of [[]map[string]interface{}].
	Sort      interface{} `json:"sort"`
	SortOrder int64       `json:"sort_order" api:"nullable"`
	TeamID    string      `json:"team_id" api:"nullable" format:"uuid"`
	UpdatedAt string      `json:"updated_at" api:"nullable"`
	UserID    string      `json:"user_id" api:"nullable"`
	// This field can have the runtime type of
	// [ViewGetResponseViewBundleWithRecordsView].
	View     interface{}         `json:"view"`
	ViewType string              `json:"view_type"`
	JSON     viewGetResponseJSON `json:"-"`
	union    ViewGetResponseUnion
}

// viewGetResponseJSON contains the JSON metadata for the struct [ViewGetResponse]
type viewGetResponseJSON struct {
	ID                   apijson.Field
	AggregationPropDefID apijson.Field
	AggregationType      apijson.Field
	ColumnLayout         apijson.Field
	Combinator           apijson.Field
	CreatedAt            apijson.Field
	Filter               apijson.Field
	GroupBy              apijson.Field
	GroupHiddenOptionIDs apijson.Field
	GroupHideEmpty       apijson.Field
	GroupSort            apijson.Field
	Icon                 apijson.Field
	ListID               apijson.Field
	Name                 apijson.Field
	Records              apijson.Field
	Select               apijson.Field
	Sort                 apijson.Field
	SortOrder            apijson.Field
	TeamID               apijson.Field
	UpdatedAt            apijson.Field
	UserID               apijson.Field
	View                 apijson.Field
	ViewType             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r viewGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ViewGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ViewGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ViewGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ViewGetResponseViewBundle],
// [ViewGetResponseViewBundleWithRecords].
func (r ViewGetResponse) AsUnion() ViewGetResponseUnion {
	return r.union
}

// A view (saved configuration for displaying records of a given object type) plus
// its select/filter/sort children. Properties in select/filter/sort are referenced
// by slug.
//
// Union satisfied by [ViewGetResponseViewBundle] or
// [ViewGetResponseViewBundleWithRecords].
type ViewGetResponseUnion interface {
	implementsViewGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ViewGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ViewGetResponseViewBundle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ViewGetResponseViewBundleWithRecords{}),
		},
	)
}

// A view (saved configuration for displaying records of a given object type) plus
// its select/filter/sort children. Properties in select/filter/sort are referenced
// by slug.
type ViewGetResponseViewBundle struct {
	Name                 string                              `json:"name" api:"required"`
	ViewType             string                              `json:"view_type" api:"required"`
	ID                   string                              `json:"id" format:"uuid"`
	AggregationPropDefID string                              `json:"aggregation_prop_def_id" api:"nullable" format:"uuid"`
	AggregationType      string                              `json:"aggregation_type" api:"nullable"`
	ColumnLayout         map[string]interface{}              `json:"column_layout" api:"nullable"`
	Combinator           ViewGetResponseViewBundleCombinator `json:"combinator"`
	CreatedAt            string                              `json:"created_at"`
	// Each entry is { slug: { comparator: value } }
	Filter []map[string]interface{} `json:"filter"`
	// Property slug to group by
	GroupBy              string        `json:"group_by" api:"nullable"`
	GroupHiddenOptionIDs []interface{} `json:"group_hidden_option_ids" api:"nullable"`
	GroupHideEmpty       bool          `json:"group_hide_empty" api:"nullable"`
	GroupSort            string        `json:"group_sort" api:"nullable"`
	Icon                 string        `json:"icon" api:"nullable"`
	ListID               string        `json:"list_id" api:"nullable" format:"uuid"`
	// Property slugs (dot-paths permitted for refs)
	Select []string `json:"select"`
	// Each entry is { slug: 'asc' | 'desc' }
	Sort      []map[string]interface{}      `json:"sort"`
	SortOrder int64                         `json:"sort_order" api:"nullable"`
	TeamID    string                        `json:"team_id" api:"nullable" format:"uuid"`
	UpdatedAt string                        `json:"updated_at" api:"nullable"`
	UserID    string                        `json:"user_id" api:"nullable"`
	JSON      viewGetResponseViewBundleJSON `json:"-"`
}

// viewGetResponseViewBundleJSON contains the JSON metadata for the struct
// [ViewGetResponseViewBundle]
type viewGetResponseViewBundleJSON struct {
	Name                 apijson.Field
	ViewType             apijson.Field
	ID                   apijson.Field
	AggregationPropDefID apijson.Field
	AggregationType      apijson.Field
	ColumnLayout         apijson.Field
	Combinator           apijson.Field
	CreatedAt            apijson.Field
	Filter               apijson.Field
	GroupBy              apijson.Field
	GroupHiddenOptionIDs apijson.Field
	GroupHideEmpty       apijson.Field
	GroupSort            apijson.Field
	Icon                 apijson.Field
	ListID               apijson.Field
	Select               apijson.Field
	Sort                 apijson.Field
	SortOrder            apijson.Field
	TeamID               apijson.Field
	UpdatedAt            apijson.Field
	UserID               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ViewGetResponseViewBundle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewGetResponseViewBundleJSON) RawJSON() string {
	return r.raw
}

func (r ViewGetResponseViewBundle) implementsViewGetResponse() {}

type ViewGetResponseViewBundleCombinator string

const (
	ViewGetResponseViewBundleCombinatorAnd ViewGetResponseViewBundleCombinator = "AND"
	ViewGetResponseViewBundleCombinatorOr  ViewGetResponseViewBundleCombinator = "OR"
)

func (r ViewGetResponseViewBundleCombinator) IsKnown() bool {
	switch r {
	case ViewGetResponseViewBundleCombinatorAnd, ViewGetResponseViewBundleCombinatorOr:
		return true
	}
	return false
}

// Returned by `GET /views/{viewId}?include=records`. Same `records` shape as the
// standalone list-view-records endpoint.
type ViewGetResponseViewBundleWithRecords struct {
	Records ViewGetResponseViewBundleWithRecordsRecords `json:"records" api:"required"`
	// A view (saved configuration for displaying records of a given object type) plus
	// its select/filter/sort children. Properties in select/filter/sort are referenced
	// by slug.
	View ViewGetResponseViewBundleWithRecordsView `json:"view" api:"required"`
	JSON viewGetResponseViewBundleWithRecordsJSON `json:"-"`
}

// viewGetResponseViewBundleWithRecordsJSON contains the JSON metadata for the
// struct [ViewGetResponseViewBundleWithRecords]
type viewGetResponseViewBundleWithRecordsJSON struct {
	Records     apijson.Field
	View        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ViewGetResponseViewBundleWithRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewGetResponseViewBundleWithRecordsJSON) RawJSON() string {
	return r.raw
}

func (r ViewGetResponseViewBundleWithRecords) implementsViewGetResponse() {}

type ViewGetResponseViewBundleWithRecordsRecords struct {
	Data       []map[string]interface{}                        `json:"data" api:"required"`
	HasMore    bool                                            `json:"has_more" api:"required"`
	NextCursor string                                          `json:"next_cursor" api:"nullable"`
	JSON       viewGetResponseViewBundleWithRecordsRecordsJSON `json:"-"`
}

// viewGetResponseViewBundleWithRecordsRecordsJSON contains the JSON metadata for
// the struct [ViewGetResponseViewBundleWithRecordsRecords]
type viewGetResponseViewBundleWithRecordsRecordsJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ViewGetResponseViewBundleWithRecordsRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewGetResponseViewBundleWithRecordsRecordsJSON) RawJSON() string {
	return r.raw
}

// A view (saved configuration for displaying records of a given object type) plus
// its select/filter/sort children. Properties in select/filter/sort are referenced
// by slug.
type ViewGetResponseViewBundleWithRecordsView struct {
	Name                 string                                             `json:"name" api:"required"`
	ViewType             string                                             `json:"view_type" api:"required"`
	ID                   string                                             `json:"id" format:"uuid"`
	AggregationPropDefID string                                             `json:"aggregation_prop_def_id" api:"nullable" format:"uuid"`
	AggregationType      string                                             `json:"aggregation_type" api:"nullable"`
	ColumnLayout         map[string]interface{}                             `json:"column_layout" api:"nullable"`
	Combinator           ViewGetResponseViewBundleWithRecordsViewCombinator `json:"combinator"`
	CreatedAt            string                                             `json:"created_at"`
	// Each entry is { slug: { comparator: value } }
	Filter []map[string]interface{} `json:"filter"`
	// Property slug to group by
	GroupBy              string        `json:"group_by" api:"nullable"`
	GroupHiddenOptionIDs []interface{} `json:"group_hidden_option_ids" api:"nullable"`
	GroupHideEmpty       bool          `json:"group_hide_empty" api:"nullable"`
	GroupSort            string        `json:"group_sort" api:"nullable"`
	Icon                 string        `json:"icon" api:"nullable"`
	ListID               string        `json:"list_id" api:"nullable" format:"uuid"`
	// Property slugs (dot-paths permitted for refs)
	Select []string `json:"select"`
	// Each entry is { slug: 'asc' | 'desc' }
	Sort      []map[string]interface{}                     `json:"sort"`
	SortOrder int64                                        `json:"sort_order" api:"nullable"`
	TeamID    string                                       `json:"team_id" api:"nullable" format:"uuid"`
	UpdatedAt string                                       `json:"updated_at" api:"nullable"`
	UserID    string                                       `json:"user_id" api:"nullable"`
	JSON      viewGetResponseViewBundleWithRecordsViewJSON `json:"-"`
}

// viewGetResponseViewBundleWithRecordsViewJSON contains the JSON metadata for the
// struct [ViewGetResponseViewBundleWithRecordsView]
type viewGetResponseViewBundleWithRecordsViewJSON struct {
	Name                 apijson.Field
	ViewType             apijson.Field
	ID                   apijson.Field
	AggregationPropDefID apijson.Field
	AggregationType      apijson.Field
	ColumnLayout         apijson.Field
	Combinator           apijson.Field
	CreatedAt            apijson.Field
	Filter               apijson.Field
	GroupBy              apijson.Field
	GroupHiddenOptionIDs apijson.Field
	GroupHideEmpty       apijson.Field
	GroupSort            apijson.Field
	Icon                 apijson.Field
	ListID               apijson.Field
	Select               apijson.Field
	Sort                 apijson.Field
	SortOrder            apijson.Field
	TeamID               apijson.Field
	UpdatedAt            apijson.Field
	UserID               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ViewGetResponseViewBundleWithRecordsView) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r viewGetResponseViewBundleWithRecordsViewJSON) RawJSON() string {
	return r.raw
}

type ViewGetResponseViewBundleWithRecordsViewCombinator string

const (
	ViewGetResponseViewBundleWithRecordsViewCombinatorAnd ViewGetResponseViewBundleWithRecordsViewCombinator = "AND"
	ViewGetResponseViewBundleWithRecordsViewCombinatorOr  ViewGetResponseViewBundleWithRecordsViewCombinator = "OR"
)

func (r ViewGetResponseViewBundleWithRecordsViewCombinator) IsKnown() bool {
	switch r {
	case ViewGetResponseViewBundleWithRecordsViewCombinatorAnd, ViewGetResponseViewBundleWithRecordsViewCombinatorOr:
		return true
	}
	return false
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
	// Each entry is { slug: { comparator: value } }
	Filter param.Field[[]map[string]interface{}] `json:"filter"`
	// Property slug to group by
	GroupBy              param.Field[string]        `json:"group_by"`
	GroupHiddenOptionIDs param.Field[[]interface{}] `json:"group_hidden_option_ids"`
	GroupHideEmpty       param.Field[bool]          `json:"group_hide_empty"`
	GroupSort            param.Field[string]        `json:"group_sort"`
	Icon                 param.Field[string]        `json:"icon"`
	ListID               param.Field[string]        `json:"list_id" format:"uuid"`
	// Property slugs (dot-paths permitted for refs)
	Select param.Field[[]string] `json:"select"`
	// Each entry is { slug: 'asc' | 'desc' }
	Sort           param.Field[[]map[string]interface{}] `json:"sort"`
	SortOrder      param.Field[int64]                    `json:"sort_order"`
	BodyTeamID     param.Field[string]                   `json:"team_id" format:"uuid"`
	UpdatedAt      param.Field[string]                   `json:"updated_at"`
	UserID         param.Field[string]                   `json:"user_id"`
	IdempotencyKey param.Field[string]                   `header:"Idempotency-Key"`
}

func (r ViewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ViewNewParamsViewObjectType string

const (
	ViewNewParamsViewObjectTypeComment      ViewNewParamsViewObjectType = "comment"
	ViewNewParamsViewObjectTypeAction       ViewNewParamsViewObjectType = "action"
	ViewNewParamsViewObjectTypeDeal         ViewNewParamsViewObjectType = "deal"
	ViewNewParamsViewObjectTypeEngagement   ViewNewParamsViewObjectType = "engagement"
	ViewNewParamsViewObjectTypeDocument     ViewNewParamsViewObjectType = "document"
	ViewNewParamsViewObjectTypeEvent        ViewNewParamsViewObjectType = "event"
	ViewNewParamsViewObjectTypeIdentity     ViewNewParamsViewObjectType = "identity"
	ViewNewParamsViewObjectTypeOrganization ViewNewParamsViewObjectType = "organization"
)

func (r ViewNewParamsViewObjectType) IsKnown() bool {
	switch r {
	case ViewNewParamsViewObjectTypeComment, ViewNewParamsViewObjectTypeAction, ViewNewParamsViewObjectTypeDeal, ViewNewParamsViewObjectTypeEngagement, ViewNewParamsViewObjectTypeDocument, ViewNewParamsViewObjectTypeEvent, ViewNewParamsViewObjectTypeIdentity, ViewNewParamsViewObjectTypeOrganization:
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
	Filter               param.Field[[]map[string]interface{}]   `json:"filter"`
	GroupBy              param.Field[string]                     `json:"group_by"`
	GroupHiddenOptionIDs param.Field[[]interface{}]              `json:"group_hidden_option_ids"`
	GroupHideEmpty       param.Field[bool]                       `json:"group_hide_empty"`
	GroupSort            param.Field[string]                     `json:"group_sort"`
	Icon                 param.Field[string]                     `json:"icon"`
	ListID               param.Field[string]                     `json:"list_id" format:"uuid"`
	Name                 param.Field[string]                     `json:"name"`
	Select               param.Field[[]string]                   `json:"select"`
	Sort                 param.Field[[]map[string]interface{}]   `json:"sort"`
	SortOrder            param.Field[int64]                      `json:"sort_order"`
	BodyTeamID           param.Field[string]                     `json:"team_id" format:"uuid"`
	UserID               param.Field[string]                     `json:"user_id"`
	ViewType             param.Field[string]                     `json:"view_type"`
	IdempotencyKey       param.Field[string]                     `header:"Idempotency-Key"`
}

func (r ViewUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ViewUpdateParamsViewObjectType string

const (
	ViewUpdateParamsViewObjectTypeComment      ViewUpdateParamsViewObjectType = "comment"
	ViewUpdateParamsViewObjectTypeAction       ViewUpdateParamsViewObjectType = "action"
	ViewUpdateParamsViewObjectTypeDeal         ViewUpdateParamsViewObjectType = "deal"
	ViewUpdateParamsViewObjectTypeEngagement   ViewUpdateParamsViewObjectType = "engagement"
	ViewUpdateParamsViewObjectTypeDocument     ViewUpdateParamsViewObjectType = "document"
	ViewUpdateParamsViewObjectTypeEvent        ViewUpdateParamsViewObjectType = "event"
	ViewUpdateParamsViewObjectTypeIdentity     ViewUpdateParamsViewObjectType = "identity"
	ViewUpdateParamsViewObjectTypeOrganization ViewUpdateParamsViewObjectType = "organization"
)

func (r ViewUpdateParamsViewObjectType) IsKnown() bool {
	switch r {
	case ViewUpdateParamsViewObjectTypeComment, ViewUpdateParamsViewObjectTypeAction, ViewUpdateParamsViewObjectTypeDeal, ViewUpdateParamsViewObjectTypeEngagement, ViewUpdateParamsViewObjectTypeDocument, ViewUpdateParamsViewObjectTypeEvent, ViewUpdateParamsViewObjectTypeIdentity, ViewUpdateParamsViewObjectTypeOrganization:
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
	ViewDeleteParamsViewObjectTypeComment      ViewDeleteParamsViewObjectType = "comment"
	ViewDeleteParamsViewObjectTypeAction       ViewDeleteParamsViewObjectType = "action"
	ViewDeleteParamsViewObjectTypeDeal         ViewDeleteParamsViewObjectType = "deal"
	ViewDeleteParamsViewObjectTypeEngagement   ViewDeleteParamsViewObjectType = "engagement"
	ViewDeleteParamsViewObjectTypeDocument     ViewDeleteParamsViewObjectType = "document"
	ViewDeleteParamsViewObjectTypeEvent        ViewDeleteParamsViewObjectType = "event"
	ViewDeleteParamsViewObjectTypeIdentity     ViewDeleteParamsViewObjectType = "identity"
	ViewDeleteParamsViewObjectTypeOrganization ViewDeleteParamsViewObjectType = "organization"
)

func (r ViewDeleteParamsViewObjectType) IsKnown() bool {
	switch r {
	case ViewDeleteParamsViewObjectTypeComment, ViewDeleteParamsViewObjectTypeAction, ViewDeleteParamsViewObjectTypeDeal, ViewDeleteParamsViewObjectTypeEngagement, ViewDeleteParamsViewObjectTypeDocument, ViewDeleteParamsViewObjectTypeEvent, ViewDeleteParamsViewObjectTypeIdentity, ViewDeleteParamsViewObjectTypeOrganization:
		return true
	}
	return false
}

type ViewGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Forwarded to the records sub-resource when `include=records`.
	Cursor param.Field[string] `query:"cursor"`
	// Comma-separated list of optional sub-resources to inline. Currently the only
	// recognized value is `records` — when present, the response is `{view, records}`
	// rather than the bare view bundle.
	Include param.Field[string] `query:"include"`
	// Forwarded to the records sub-resource when `include=records`.
	Limit param.Field[int64] `query:"limit"`
	// Forwarded to the records sub-resource when `include=records`.
	Page param.Field[int64] `query:"page"`
}

// URLQuery serializes [ViewGetParams]'s query parameters as `url.Values`.
func (r ViewGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ViewGetParamsViewObjectType string

const (
	ViewGetParamsViewObjectTypeComment      ViewGetParamsViewObjectType = "comment"
	ViewGetParamsViewObjectTypeAction       ViewGetParamsViewObjectType = "action"
	ViewGetParamsViewObjectTypeDeal         ViewGetParamsViewObjectType = "deal"
	ViewGetParamsViewObjectTypeEngagement   ViewGetParamsViewObjectType = "engagement"
	ViewGetParamsViewObjectTypeDocument     ViewGetParamsViewObjectType = "document"
	ViewGetParamsViewObjectTypeEvent        ViewGetParamsViewObjectType = "event"
	ViewGetParamsViewObjectTypeIdentity     ViewGetParamsViewObjectType = "identity"
	ViewGetParamsViewObjectTypeOrganization ViewGetParamsViewObjectType = "organization"
)

func (r ViewGetParamsViewObjectType) IsKnown() bool {
	switch r {
	case ViewGetParamsViewObjectTypeComment, ViewGetParamsViewObjectTypeAction, ViewGetParamsViewObjectTypeDeal, ViewGetParamsViewObjectTypeEngagement, ViewGetParamsViewObjectTypeDocument, ViewGetParamsViewObjectTypeEvent, ViewGetParamsViewObjectTypeIdentity, ViewGetParamsViewObjectTypeOrganization:
		return true
	}
	return false
}
