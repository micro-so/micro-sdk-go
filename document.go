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

// Documents are rich-text notes attached to contacts, organizations, or deals,
// used for meeting notes, research, or context.
//
// DocumentService contains methods and other services that help with interacting
// with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentService] method instead.
type DocumentService struct {
	Options []option.RequestOption
}

// NewDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentService(opts ...option.RequestOption) (r *DocumentService) {
	r = &DocumentService{}
	r.Options = opts
	return
}

// Create Document
func (r *DocumentService) New(ctx context.Context, params DocumentNewParams, opts ...option.RequestOption) (res *DocumentNewResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/document", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update Document
func (r *DocumentService) Update(ctx context.Context, documentID string, params DocumentUpdateParams, opts ...option.RequestOption) (err error) {
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
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/document/%s", params.TeamID, documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return err
}

// List Documents
func (r *DocumentService) List(ctx context.Context, params DocumentListParams, opts ...option.RequestOption) (res *DocumentListResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/query/%s/document", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Delete Document
func (r *DocumentService) Delete(ctx context.Context, documentID string, body DocumentDeleteParams, opts ...option.RequestOption) (err error) {
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
	if documentID == "" {
		err = errors.New("missing required documentId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/document/%s", body.TeamID, documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type DocumentNewResponse struct {
	ID   string                  `json:"id" format:"uuid"`
	JSON documentNewResponseJSON `json:"-"`
}

// documentNewResponseJSON contains the JSON metadata for the struct
// [DocumentNewResponse]
type documentNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DocumentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r documentNewResponseJSON) RawJSON() string {
	return r.raw
}

type DocumentListResponse struct {
	Data       []interface{}            `json:"data"`
	NextCursor string                   `json:"next_cursor" api:"nullable"`
	Total      int64                    `json:"total"`
	JSON       documentListResponseJSON `json:"-"`
}

// documentListResponseJSON contains the JSON metadata for the struct
// [DocumentListResponse]
type documentListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DocumentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r documentListResponseJSON) RawJSON() string {
	return r.raw
}

type DocumentNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r DocumentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type DocumentUpdateParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r DocumentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type DocumentListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID  param.Field[string]                    `path:"teamId" api:"required" format:"uuid"`
	Query   param.Field[DocumentListParamsQuery]   `json:"query" api:"required"`
	ID      param.Field[DocumentListParamsIDUnion] `json:"id" format:"uuid"`
	Boxes   param.Field[[]string]                  `json:"boxes"`
	Deleted param.Field[bool]                      `json:"deleted"`
	Sources param.Field[[]string]                  `json:"sources" format:"uuid"`
}

func (r DocumentListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DocumentListParamsQuery struct {
	// Property slugs to select. Use dot notation for relationships (e.g.
	// attendee.contact.first_name)
	Select param.Field[[]string] `json:"select" api:"required"`
	// Logical operator for combining filters
	Combinator param.Field[DocumentListParamsQueryCombinator] `json:"combinator"`
	CRMID      param.Field[string]                            `json:"crm_id" format:"uuid"`
	// Filters as [{ slug: { operator: value } }]. For select/multiselect properties,
	// values must be option slugs
	Filter param.Field[[]map[string]map[string]DocumentListParamsQueryFilterUnion] `json:"filter"`
	Limit  param.Field[int64]                                                      `json:"limit"`
	Page   param.Field[int64]                                                      `json:"page"`
	// Sort order as [{ slug: direction }]. Array order determines sort priority
	Sort param.Field[[]map[string]DocumentListParamsQuerySort] `json:"sort"`
}

func (r DocumentListParamsQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logical operator for combining filters
type DocumentListParamsQueryCombinator string

const (
	DocumentListParamsQueryCombinatorAnd DocumentListParamsQueryCombinator = "AND"
	DocumentListParamsQueryCombinatorOr  DocumentListParamsQueryCombinator = "OR"
)

func (r DocumentListParamsQueryCombinator) IsKnown() bool {
	switch r {
	case DocumentListParamsQueryCombinatorAnd, DocumentListParamsQueryCombinatorOr:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionBool],
// [DocumentListParamsQueryFilterArray].
type DocumentListParamsQueryFilterUnion interface {
	ImplementsDocumentListParamsQueryFilterUnion()
}

type DocumentListParamsQueryFilterArray []string

func (r DocumentListParamsQueryFilterArray) ImplementsDocumentListParamsQueryFilterUnion() {}

type DocumentListParamsQuerySort string

const (
	DocumentListParamsQuerySortAsc  DocumentListParamsQuerySort = "asc"
	DocumentListParamsQuerySortDesc DocumentListParamsQuerySort = "desc"
)

func (r DocumentListParamsQuerySort) IsKnown() bool {
	switch r {
	case DocumentListParamsQuerySortAsc, DocumentListParamsQuerySortDesc:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [DocumentListParamsIDArray].
type DocumentListParamsIDUnion interface {
	ImplementsDocumentListParamsIDUnion()
}

type DocumentListParamsIDArray []string

func (r DocumentListParamsIDArray) ImplementsDocumentListParamsIDUnion() {}

type DocumentDeleteParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
