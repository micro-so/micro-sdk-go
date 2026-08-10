// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/micro-so/micro-sdk-go/internal/apijson"
	"github.com/micro-so/micro-sdk-go/internal/param"
	"github.com/micro-so/micro-sdk-go/internal/requestconfig"
	"github.com/micro-so/micro-sdk-go/option"
)

// PrismListService contains methods and other services that help with interacting
// with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismListService] method instead.
type PrismListService struct {
	Options []option.RequestOption
}

// NewPrismListService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrismListService(opts ...option.RequestOption) (r *PrismListService) {
	r = &PrismListService{}
	r.Options = opts
	return
}

// Creates a list from a template. Seeds properties, pipeline stages (when
// applicable), and default views — identical to the session-auth
// `/default_app/create` path. API-key callers are fully supported; `type` is
// derived from `template_id` and must not be supplied.
func (r *PrismListService) New(ctx context.Context, params PrismListNewParams, opts ...option.RequestOption) (res *List, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
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
	path := fmt.Sprintf("v2/prism/%s/lists", params.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns non-core lists the caller can access in the workspace. Core system apps
// (Messages, All Inbox) are excluded.
func (r *PrismListService) List(ctx context.Context, query PrismListListParams, opts ...option.RequestOption) (res *PrismListListResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/lists", query.TeamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get a list by id
func (r *PrismListService) Get(ctx context.Context, listID string, query PrismListGetParams, opts ...option.RequestOption) (res *List, err error) {
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
	if listID == "" {
		err = errors.New("missing required listId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/lists/%s", query.TeamID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type List struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// Prism object type this list holds.
	ObjectType  ListObjectType `json:"object_type" api:"required"`
	TeamID      string         `json:"team_id" api:"required" format:"uuid"`
	CreatedAt   time.Time      `json:"created_at" api:"nullable" format:"date-time"`
	Description string         `json:"description" api:"nullable"`
	// Emoji or icon key for the list.
	Icon string `json:"icon" api:"nullable"`
	// Internal template type (e.g. dealFlow, hiring). Derived from template_id on
	// create.
	Type  string     `json:"type" api:"nullable"`
	Views []ListView `json:"views"`
	JSON  listJSON   `json:"-"`
}

// listJSON contains the JSON metadata for the struct [List]
type listJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	ObjectType  apijson.Field
	TeamID      apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Icon        apijson.Field
	Type        apijson.Field
	Views       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *List) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listJSON) RawJSON() string {
	return r.raw
}

// Prism object type this list holds.
type ListObjectType string

const (
	ListObjectTypeOrganization ListObjectType = "organization"
	ListObjectTypeIdentity     ListObjectType = "identity"
	ListObjectTypeAction       ListObjectType = "action"
	ListObjectTypeDocument     ListObjectType = "document"
	ListObjectTypeDeal         ListObjectType = "deal"
)

func (r ListObjectType) IsKnown() bool {
	switch r {
	case ListObjectTypeOrganization, ListObjectTypeIdentity, ListObjectTypeAction, ListObjectTypeDocument, ListObjectTypeDeal:
		return true
	}
	return false
}

type ListView struct {
	ID   string       `json:"id" api:"required" format:"uuid"`
	Name string       `json:"name" api:"nullable"`
	JSON listViewJSON `json:"-"`
}

// listViewJSON contains the JSON metadata for the struct [ListView]
type listViewJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListView) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listViewJSON) RawJSON() string {
	return r.raw
}

type ListCreateParam struct {
	// Template to seed the list from. `type` is derived server-side from this
	// template.
	TemplateID param.Field[ListCreateTemplateID] `json:"template_id" api:"required"`
	// Emoji or icon override.
	Icon param.Field[string] `json:"icon"`
	Name param.Field[string] `json:"name"`
	// Required only when template_id is `custom`.
	ObjectType param.Field[ListCreateObjectType] `json:"object_type"`
}

func (r ListCreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Template to seed the list from. `type` is derived server-side from this
// template.
type ListCreateTemplateID string

const (
	ListCreateTemplateIDSalesDeals      ListCreateTemplateID = "sales_deals"
	ListCreateTemplateIDRecruiting      ListCreateTemplateID = "recruiting"
	ListCreateTemplateIDPartnerships    ListCreateTemplateID = "partnerships"
	ListCreateTemplateIDFundraising     ListCreateTemplateID = "fundraising"
	ListCreateTemplateIDKnowledgeBase   ListCreateTemplateID = "knowledge_base"
	ListCreateTemplateIDIssueTracker    ListCreateTemplateID = "issue_tracker"
	ListCreateTemplateIDContentCalendar ListCreateTemplateID = "content_calendar"
	ListCreateTemplateIDJobApplications ListCreateTemplateID = "job_applications"
	ListCreateTemplateIDProjectTracker  ListCreateTemplateID = "project_tracker"
	ListCreateTemplateIDFeedback        ListCreateTemplateID = "feedback"
	ListCreateTemplateIDPortcoTracker   ListCreateTemplateID = "portco_tracker"
	ListCreateTemplateIDDealFlow        ListCreateTemplateID = "deal_flow"
	ListCreateTemplateIDLpFundraising   ListCreateTemplateID = "lp_fundraising"
	ListCreateTemplateIDCustom          ListCreateTemplateID = "custom"
)

func (r ListCreateTemplateID) IsKnown() bool {
	switch r {
	case ListCreateTemplateIDSalesDeals, ListCreateTemplateIDRecruiting, ListCreateTemplateIDPartnerships, ListCreateTemplateIDFundraising, ListCreateTemplateIDKnowledgeBase, ListCreateTemplateIDIssueTracker, ListCreateTemplateIDContentCalendar, ListCreateTemplateIDJobApplications, ListCreateTemplateIDProjectTracker, ListCreateTemplateIDFeedback, ListCreateTemplateIDPortcoTracker, ListCreateTemplateIDDealFlow, ListCreateTemplateIDLpFundraising, ListCreateTemplateIDCustom:
		return true
	}
	return false
}

// Required only when template_id is `custom`.
type ListCreateObjectType string

const (
	ListCreateObjectTypeOrganization ListCreateObjectType = "organization"
	ListCreateObjectTypeIdentity     ListCreateObjectType = "identity"
	ListCreateObjectTypeAction       ListCreateObjectType = "action"
	ListCreateObjectTypeDocument     ListCreateObjectType = "document"
	ListCreateObjectTypeDeal         ListCreateObjectType = "deal"
)

func (r ListCreateObjectType) IsKnown() bool {
	switch r {
	case ListCreateObjectTypeOrganization, ListCreateObjectTypeIdentity, ListCreateObjectTypeAction, ListCreateObjectTypeDocument, ListCreateObjectTypeDeal:
		return true
	}
	return false
}

type PrismListListResponse struct {
	Data []List                    `json:"data" api:"required"`
	JSON prismListListResponseJSON `json:"-"`
}

// prismListListResponseJSON contains the JSON metadata for the struct
// [PrismListListResponse]
type prismListListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismListListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismListListResponseJSON) RawJSON() string {
	return r.raw
}

type PrismListNewParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID         param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	ListCreate     ListCreateParam     `json:"list_create" api:"required"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

func (r PrismListNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ListCreate)
}

type PrismListListParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismListGetParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
