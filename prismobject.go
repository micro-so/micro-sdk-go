// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package micro

import (
	"github.com/micro-so/micro-sdk-go/option"
)

// PrismObjectService contains methods and other services that help with
// interacting with the micro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrismObjectService] method instead.
type PrismObjectService struct {
	Options       []option.RequestOption
	Contacts      *PrismObjectContactService
	Organizations *PrismObjectOrganizationService
	Identities    *PrismObjectIdentityService
	Deals         *PrismObjectDealService
	Actions       *PrismObjectActionService
	Documents     *PrismObjectDocumentService
	Events        *PrismObjectEventService
}

// NewPrismObjectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrismObjectService(opts ...option.RequestOption) (r *PrismObjectService) {
	r = &PrismObjectService{}
	r.Options = opts
	r.Contacts = NewPrismObjectContactService(opts...)
	r.Organizations = NewPrismObjectOrganizationService(opts...)
	r.Identities = NewPrismObjectIdentityService(opts...)
	r.Deals = NewPrismObjectDealService(opts...)
	r.Actions = NewPrismObjectActionService(opts...)
	r.Documents = NewPrismObjectDocumentService(opts...)
	r.Events = NewPrismObjectEventService(opts...)
	return
}
