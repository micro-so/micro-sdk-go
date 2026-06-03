// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"github.com/micro-so/micro-sdk-go/option"
)

// PrismObjectEngagementService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectEngagementService] method instead.
type PrismObjectEngagementService struct {
	Options []option.RequestOption
	Grant   *PrismObjectEngagementGrantService
}

// NewPrismObjectEngagementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPrismObjectEngagementService(opts ...option.RequestOption) (r *PrismObjectEngagementService) {
	r = &PrismObjectEngagementService{}
	r.Options = opts
	r.Grant = NewPrismObjectEngagementGrantService(opts...)
	return
}
