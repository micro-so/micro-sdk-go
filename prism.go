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
	Grant    *PrismGrantService
	Query    *PrismQueryService
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

type PrismObjectPropertiesParam struct {
	ID       param.Field[string]      `json:"id" format:"uuid"`
	CRM      param.Field[interface{}] `json:"crm"`
	Default  param.Field[interface{}] `json:"default"`
	Extended param.Field[interface{}] `json:"extended"`
}

func (r PrismObjectPropertiesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
