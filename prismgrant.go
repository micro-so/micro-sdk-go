// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"github.com/stainless-sdks/micro-go/option"
)

// PrismGrantService contains methods and other services that help with interacting
// with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismGrantService] method instead.
type PrismGrantService struct {
	Options []option.RequestOption
}

// NewPrismGrantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrismGrantService(opts ...option.RequestOption) (r *PrismGrantService) {
	r = &PrismGrantService{}
	r.Options = opts
	return
}
