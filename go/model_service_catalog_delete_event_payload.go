/*
 * Service Catalog Management
 *
 * ## TMF API Reference: TMF633 - Service Catalog Management   Version 4.0   The Service Catalog API is one of Catalog Management API Family. Service Catalog API goal is to provide a catalog of services.  Service Catalog API performs the following operations on the resource : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity (for administration purposes) - Manage notification of events. .  Copyright © TM Forum 2020. All Rights Reserved. 
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf633

// ServiceCatalogDeleteEventPayload - The event data structure
type ServiceCatalogDeleteEventPayload struct {

	ServiceCatalog ServiceCatalog `json:"serviceCatalog,omitempty"`
}

// AssertServiceCatalogDeleteEventPayloadRequired checks if the required fields are not zero-ed
func AssertServiceCatalogDeleteEventPayloadRequired(obj ServiceCatalogDeleteEventPayload) error {
	if err := AssertServiceCatalogRequired(obj.ServiceCatalog); err != nil {
		return err
	}
	return nil
}

// AssertRecurseServiceCatalogDeleteEventPayloadRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ServiceCatalogDeleteEventPayload (e.g. [][]ServiceCatalogDeleteEventPayload), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServiceCatalogDeleteEventPayloadRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aServiceCatalogDeleteEventPayload, ok := obj.(ServiceCatalogDeleteEventPayload)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServiceCatalogDeleteEventPayloadRequired(aServiceCatalogDeleteEventPayload)
	})
}
