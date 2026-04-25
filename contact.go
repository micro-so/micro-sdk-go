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

// ContactService contains methods and other services that help with interacting
// with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContactService] method instead.
type ContactService struct {
	Options []option.RequestOption
}

// NewContactService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewContactService(opts ...option.RequestOption) (r *ContactService) {
	r = &ContactService{}
	r.Options = opts
	return
}

// Create Contact
func (r *ContactService) New(ctx context.Context, params ContactNewParams, opts ...option.RequestOption) (res *ContactNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/contact", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update Contact
func (r *ContactService) Update(ctx context.Context, contactID string, params ContactUpdateParams, opts ...option.RequestOption) (err error) {
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
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/contact/%s", params.TeamID, contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return err
}

// List Contacts
func (r *ContactService) List(ctx context.Context, params ContactListParams, opts ...option.RequestOption) (res *ContactListResponse, err error) {
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

// Delete Contact
func (r *ContactService) Delete(ctx context.Context, contactID string, body ContactDeleteParams, opts ...option.RequestOption) (err error) {
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
	if contactID == "" {
		err = errors.New("missing required contactId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/contact/%s", body.TeamID, contactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Import Contacts
func (r *ContactService) Import(ctx context.Context, params ContactImportParams, opts ...option.RequestOption) (res *ContactImportResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/contact/import", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type ContactNewResponse struct {
	ID   string                 `json:"id" format:"uuid"`
	JSON contactNewResponseJSON `json:"-"`
}

// contactNewResponseJSON contains the JSON metadata for the struct
// [ContactNewResponse]
type contactNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContactNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contactNewResponseJSON) RawJSON() string {
	return r.raw
}

type ContactListResponse struct {
	Data       []interface{}           `json:"data"`
	NextCursor string                  `json:"next_cursor" api:"nullable"`
	Total      int64                   `json:"total"`
	JSON       contactListResponseJSON `json:"-"`
}

// contactListResponseJSON contains the JSON metadata for the struct
// [ContactListResponse]
type contactListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContactListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contactListResponseJSON) RawJSON() string {
	return r.raw
}

type ContactImportResponse struct {
	Results []ContactImportResponseResult `json:"results"`
	Status  ContactImportResponseStatus   `json:"status"`
	Summary ContactImportResponseSummary  `json:"summary"`
	JSON    contactImportResponseJSON     `json:"-"`
}

// contactImportResponseJSON contains the JSON metadata for the struct
// [ContactImportResponse]
type contactImportResponseJSON struct {
	Results     apijson.Field
	Status      apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContactImportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contactImportResponseJSON) RawJSON() string {
	return r.raw
}

type ContactImportResponseResult struct {
	ID       string                          `json:"id" api:"nullable" format:"uuid"`
	Created  bool                            `json:"created"`
	Error    string                          `json:"error"`
	Existing bool                            `json:"existing"`
	JSON     contactImportResponseResultJSON `json:"-"`
}

// contactImportResponseResultJSON contains the JSON metadata for the struct
// [ContactImportResponseResult]
type contactImportResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContactImportResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contactImportResponseResultJSON) RawJSON() string {
	return r.raw
}

type ContactImportResponseStatus string

const (
	ContactImportResponseStatusComplete ContactImportResponseStatus = "complete"
)

func (r ContactImportResponseStatus) IsKnown() bool {
	switch r {
	case ContactImportResponseStatusComplete:
		return true
	}
	return false
}

type ContactImportResponseSummary struct {
	Created  int64                            `json:"created"`
	Errors   int64                            `json:"errors"`
	Existing int64                            `json:"existing"`
	Total    int64                            `json:"total"`
	JSON     contactImportResponseSummaryJSON `json:"-"`
}

// contactImportResponseSummaryJSON contains the JSON metadata for the struct
// [ContactImportResponseSummary]
type contactImportResponseSummaryJSON struct {
	Created     apijson.Field
	Errors      apijson.Field
	Existing    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContactImportResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contactImportResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type ContactNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r ContactNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type ContactUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r ContactUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type ContactListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                   `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[ContactListParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[ContactListParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                 `json:"boxes"`
	Deleted param.Field[bool]                     `json:"deleted"`
	Sources param.Field[[]string]                 `json:"sources" format:"uuid"`
}

func (r ContactListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContactListParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[ContactListParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                           `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values must be option slugs
	Filter param.Field[[]map[string]map[string]ContactListParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                     `json:"limit"`
	Page   param.Field[int64]                                                     `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]ContactListParamsQuerySort] `json:"sort"`
}

func (r ContactListParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type ContactListParamsQueryCombinator string

const (
	ContactListParamsQueryCombinatorAnd ContactListParamsQueryCombinator = "AND"
	ContactListParamsQueryCombinatorOr  ContactListParamsQueryCombinator = "OR"
)

func (r ContactListParamsQueryCombinator) IsKnown() bool {
	switch r {
	case ContactListParamsQueryCombinatorAnd, ContactListParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [ContactListParamsQueryFilterArray].
type ContactListParamsQueryFilterUnion interface {
	ImplementsContactListParamsQueryFilterUnion()
}

type ContactListParamsQueryFilterArray []string

func (r ContactListParamsQueryFilterArray) ImplementsContactListParamsQueryFilterUnion() {}

type ContactListParamsQuerySort string

const (
	ContactListParamsQuerySortAsc  ContactListParamsQuerySort = "asc"
	ContactListParamsQuerySortDesc ContactListParamsQuerySort = "desc"
)

func (r ContactListParamsQuerySort) IsKnown() bool {
	switch r {
	case ContactListParamsQuerySortAsc, ContactListParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [ContactListParamsIDArray].
type ContactListParamsIDUnion interface {
	ImplementsContactListParamsIDUnion()
}

type ContactListParamsIDArray []string

func (r ContactListParamsIDArray) ImplementsContactListParamsIDUnion() {}

type ContactDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type ContactImportParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects param.Field[[]PrismObjectPropertiesParam] `json:"objects" api:"required"`
	Options param.Field[ContactImportParamsOptions]   `json:"options"`
}

func (r ContactImportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContactImportParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// App/CRM ID for context (optional)
	CRMID param.Field[string] `json:"crm_id" format:"uuid"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
}

func (r ContactImportParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
