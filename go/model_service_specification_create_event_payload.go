/*
 * Service Catalog Management
 *
 * ## TMF API Reference: TMF633 - Service Catalog Management   Version 4.0   The Service Catalog API is one of Catalog Management API Family. Service Catalog API goal is to provide a catalog of services.  Service Catalog API performs the following operations on the resource : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity (for administration purposes) - Manage notification of events. .  Copyright © TM Forum 2020. All Rights Reserved. 
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf633

// ServiceSpecificationCreateEventPayload - The event data structure
type ServiceSpecificationCreateEventPayload struct {

	ServiceSpecification ServiceSpecification `json:"serviceSpecification,omitempty"`
}

// AssertServiceSpecificationCreateEventPayloadRequired checks if the required fields are not zero-ed
func AssertServiceSpecificationCreateEventPayloadRequired(obj ServiceSpecificationCreateEventPayload) error {
	if err := AssertServiceSpecificationRequired(obj.ServiceSpecification); err != nil {
		return err
	}
	return nil
}

// AssertRecurseServiceSpecificationCreateEventPayloadRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ServiceSpecificationCreateEventPayload (e.g. [][]ServiceSpecificationCreateEventPayload), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServiceSpecificationCreateEventPayloadRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aServiceSpecificationCreateEventPayload, ok := obj.(ServiceSpecificationCreateEventPayload)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServiceSpecificationCreateEventPayloadRequired(aServiceSpecificationCreateEventPayload)
	})
}
