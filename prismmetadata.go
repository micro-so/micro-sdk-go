// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"github.com/stainless-sdks/micro-go/option"
)

// PrismMetadataService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismMetadataService] method instead.
type PrismMetadataService struct {
	Options []option.RequestOption
}

// NewPrismMetadataService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrismMetadataService(opts ...option.RequestOption) (r *PrismMetadataService) {
	r = &PrismMetadataService{}
	r.Options = opts
	return
}
