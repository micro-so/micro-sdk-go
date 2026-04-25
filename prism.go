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

// The Prism query engine provides generic read/write access to any object type
// using a single unified API surface.
//
// PrismService contains methods and other services that help with interacting with
// the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismService] method instead.
type PrismService struct {
	Options []option.RequestOption
	// The Prism query engine provides generic read/write access to any object type
	// using a single unified API surface.
	Grant *PrismGrantService
	// The Prism query engine provides generic read/write access to any object type
	// using a single unified API surface.
	Query *PrismQueryService
	// The Prism query engine provides generic read/write access to any object type
	// using a single unified API surface.
	Metadata *PrismMetadataService
}

// NewPrismService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPrismService(opts ...option.RequestOption) (r *PrismService) {
	r = &PrismService{}
	r.Options = opts
	r.Grant = NewPrismGrantService(opts...)
	r.Query = NewPrismQueryService(opts...)
	r.Metadata = NewPrismMetadataService(opts...)
	return
}

// Create object
func (r *PrismService) NewObject(ctx context.Context, objectType ObjectType, params PrismNewObjectParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("v2/prism/%s/%v", params.TeamID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Delete object
func (r *PrismService) DeleteObject(ctx context.Context, objectType ObjectType, objectID string, body PrismDeleteObjectParams, opts ...option.RequestOption) (err error) {
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
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/%s", body.TeamID, objectType, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Duplicate object
func (r *PrismService) DuplicateObject(ctx context.Context, objectType ObjectType, objectID string, body PrismDuplicateObjectParams, opts ...option.RequestOption) (res *PrismDuplicateObjectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.TeamID, precfg.TeamID)
	if body.TeamID.Value == "" {
		err = errors.New("missing required teamId parameter")
		return nil, err
	}
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/%s/duplicate", body.TeamID, objectType, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Import multiple objects in batch. Properties are keyed by slug. Automatically
// routes based on size: <100 records sync (immediate response), >=100 records
// async (S3/Lambda with WebSocket progress)
func (r *PrismService) ImportObjects(ctx context.Context, objectType PrismImportObjectsParamsObjectType, params PrismImportObjectsParams, opts ...option.RequestOption) (res *PrismImportObjectsResponse, err error) {
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
	path := fmt.Sprintf("v2/prism/%s/%v/import", params.TeamID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Patch object
func (r *PrismService) PatchObject(ctx context.Context, objectType ObjectType, objectID string, params PrismPatchObjectParams, opts ...option.RequestOption) (err error) {
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
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/%s", params.TeamID, objectType, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return err
}

// Restore object
func (r *PrismService) RestoreObject(ctx context.Context, objectType ObjectType, objectID string, body PrismRestoreObjectParams, opts ...option.RequestOption) (err error) {
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
	if objectID == "" {
		err = errors.New("missing required objectId parameter")
		return err
	}
	path := fmt.Sprintf("v2/prism/%s/%v/%s/restore", body.TeamID, objectType, objectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

type ObjectType string

const (
	ObjectTypeDeal          ObjectType = "deal"
	ObjectTypeIdentity      ObjectType = "identity"
	ObjectTypeAIChatThread  ObjectType = "ai_chat_thread"
	ObjectTypeAIChatMessage ObjectType = "ai_chat_message"
	ObjectTypeDocument      ObjectType = "document"
	ObjectTypeAction        ObjectType = "action"
	ObjectTypeEvent         ObjectType = "event"
)

func (r ObjectType) IsKnown() bool {
	switch r {
	case ObjectTypeDeal, ObjectTypeIdentity, ObjectTypeAIChatThread, ObjectTypeAIChatMessage, ObjectTypeDocument, ObjectTypeAction, ObjectTypeEvent:
		return true
	}
	return false
}

type PrismObjectPropertiesParam struct {
	ID  param.Field[string]      `json:"id" format:"uuid"`
	CRM param.Field[interface{}] `json:"crm"`
	// Properties keyed by property slug. Values can be strings, numbers, booleans,
	// arrays, or null.
	Default  param.Field[map[string]interface{}] `json:"default"`
	Extended param.Field[interface{}]            `json:"extended"`
}

func (r PrismObjectPropertiesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismDuplicateObjectResponse struct {
	ID   string                           `json:"id" format:"uuid"`
	JSON prismDuplicateObjectResponseJSON `json:"-"`
}

// prismDuplicateObjectResponseJSON contains the JSON metadata for the struct
// [PrismDuplicateObjectResponse]
type prismDuplicateObjectResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismDuplicateObjectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismDuplicateObjectResponseJSON) RawJSON() string {
	return r.raw
}

type PrismImportObjectsResponse struct {
	Results []PrismImportObjectsResponseResult `json:"results"`
	Status  PrismImportObjectsResponseStatus   `json:"status"`
	Summary PrismImportObjectsResponseSummary  `json:"summary"`
	JSON    prismImportObjectsResponseJSON     `json:"-"`
}

// prismImportObjectsResponseJSON contains the JSON metadata for the struct
// [PrismImportObjectsResponse]
type prismImportObjectsResponseJSON struct {
	Results     apijson.Field
	Status      apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismImportObjectsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismImportObjectsResponseJSON) RawJSON() string {
	return r.raw
}

type PrismImportObjectsResponseResult struct {
	ID       string                               `json:"id" api:"nullable" format:"uuid"`
	Created  bool                                 `json:"created"`
	Error    string                               `json:"error"`
	Existing bool                                 `json:"existing"`
	JSON     prismImportObjectsResponseResultJSON `json:"-"`
}

// prismImportObjectsResponseResultJSON contains the JSON metadata for the struct
// [PrismImportObjectsResponseResult]
type prismImportObjectsResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Error       apijson.Field
	Existing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismImportObjectsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismImportObjectsResponseResultJSON) RawJSON() string {
	return r.raw
}

type PrismImportObjectsResponseStatus string

const (
	PrismImportObjectsResponseStatusComplete PrismImportObjectsResponseStatus = "complete"
)

func (r PrismImportObjectsResponseStatus) IsKnown() bool {
	switch r {
	case PrismImportObjectsResponseStatusComplete:
		return true
	}
	return false
}

type PrismImportObjectsResponseSummary struct {
	Created  int64                                 `json:"created"`
	Errors   int64                                 `json:"errors"`
	Existing int64                                 `json:"existing"`
	Total    int64                                 `json:"total"`
	JSON     prismImportObjectsResponseSummaryJSON `json:"-"`
}

// prismImportObjectsResponseSummaryJSON contains the JSON metadata for the struct
// [PrismImportObjectsResponseSummary]
type prismImportObjectsResponseSummaryJSON struct {
	Created     apijson.Field
	Errors      apijson.Field
	Existing    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrismImportObjectsResponseSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prismImportObjectsResponseSummaryJSON) RawJSON() string {
	return r.raw
}

type PrismNewObjectParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismNewObjectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismDeleteObjectParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismDuplicateObjectParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}

type PrismImportObjectsParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
	// Array of objects to import with property values keyed by slug
	Objects param.Field[[]PrismObjectPropertiesParam]    `json:"objects" api:"required"`
	Options param.Field[PrismImportObjectsParamsOptions] `json:"options"`
}

func (r PrismImportObjectsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismImportObjectsParamsObjectType string

const (
	PrismImportObjectsParamsObjectTypeIdentity     PrismImportObjectsParamsObjectType = "identity"
	PrismImportObjectsParamsObjectTypeOrganization PrismImportObjectsParamsObjectType = "organization"
	PrismImportObjectsParamsObjectTypeContact      PrismImportObjectsParamsObjectType = "contact"
	PrismImportObjectsParamsObjectTypeAction       PrismImportObjectsParamsObjectType = "action"
	PrismImportObjectsParamsObjectTypeDocument     PrismImportObjectsParamsObjectType = "document"
	PrismImportObjectsParamsObjectTypeDeal         PrismImportObjectsParamsObjectType = "deal"
)

func (r PrismImportObjectsParamsObjectType) IsKnown() bool {
	switch r {
	case PrismImportObjectsParamsObjectTypeIdentity, PrismImportObjectsParamsObjectTypeOrganization, PrismImportObjectsParamsObjectTypeContact, PrismImportObjectsParamsObjectTypeAction, PrismImportObjectsParamsObjectTypeDocument, PrismImportObjectsParamsObjectTypeDeal:
		return true
	}
	return false
}

type PrismImportObjectsParamsOptions struct {
	// Whether deduplication should be case insensitive
	CaseInsensitive param.Field[bool] `json:"caseInsensitive"`
	// App/CRM ID for context (optional)
	CRMID param.Field[string] `json:"crm_id" format:"uuid"`
	// Property slug to deduplicate on
	DedupeBy param.Field[string] `json:"dedupe_by"`
}

func (r PrismImportObjectsParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrismPatchObjectParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID                param.Field[string]        `path:"teamId" api:"required" format:"uuid"`
	PrismObjectProperties PrismObjectPropertiesParam `json:"prism_object_properties" api:"required"`
}

func (r PrismPatchObjectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrismObjectProperties)
}

type PrismRestoreObjectParams struct {
	// Use [option.WithTeamID] on the client to set a global default for this field.
	TeamID param.Field[string] `path:"teamId" api:"required" format:"uuid"`
}
