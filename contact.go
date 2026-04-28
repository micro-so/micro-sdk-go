// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
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
