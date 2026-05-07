// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"github.com/stainless-sdks/micro-go/internal/apijson"
	"github.com/stainless-sdks/micro-go/internal/param"
	"github.com/stainless-sdks/micro-go/option"
)

// PrismService contains methods and other services that help with interacting with
// the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismService] method instead.
type PrismService struct {
	Options  []option.RequestOption
	Metadata *PrismMetadataService
	Objects  *PrismObjectService
}

// NewPrismService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPrismService(opts ...option.RequestOption) (r *PrismService) {
	r = &PrismService{}
	r.Options = opts
	r.Metadata = NewPrismMetadataService(opts...)
	r.Objects = NewPrismObjectService(opts...)
	return
}

type PrismObjectPropertiesParam struct {
	ID  param.Field[string]      `json:"id" format:"uuid"`
	CRM param.Field[interface{}] `json:"crm"`
	// Properties keyed by property slug. Values can be strings, numbers, booleans,
	// arrays, or null. For select/multiselect properties, values may be option slugs
	// or option UUIDs on write; option slugs are returned on read.
	Default  param.Field[map[string]interface{}] `json:"default"`
	Extended param.Field[interface{}]            `json:"extended"`
}

func (r PrismObjectPropertiesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
