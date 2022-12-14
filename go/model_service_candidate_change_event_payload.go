/*
 * Service Catalog Management
 *
 * ## TMF API Reference: TMF633 - Service Catalog Management   Version 4.0   The Service Catalog API is one of Catalog Management API Family. Service Catalog API goal is to provide a catalog of services.  Service Catalog API performs the following operations on the resource : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity (for administration purposes) - Manage notification of events. .  Copyright © TM Forum 2020. All Rights Reserved. 
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf633

// ServiceCandidateChangeEventPayload - The event data structure
type ServiceCandidateChangeEventPayload struct {

	ServiceCandidate ServiceCandidate `json:"serviceCandidate,omitempty"`
}

// AssertServiceCandidateChangeEventPayloadRequired checks if the required fields are not zero-ed
func AssertServiceCandidateChangeEventPayloadRequired(obj ServiceCandidateChangeEventPayload) error {
	if err := AssertServiceCandidateRequired(obj.ServiceCandidate); err != nil {
		return err
	}
	return nil
}

// AssertRecurseServiceCandidateChangeEventPayloadRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ServiceCandidateChangeEventPayload (e.g. [][]ServiceCandidateChangeEventPayload), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServiceCandidateChangeEventPayloadRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aServiceCandidateChangeEventPayload, ok := obj.(ServiceCandidateChangeEventPayload)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServiceCandidateChangeEventPayloadRequired(aServiceCandidateChangeEventPayload)
	})
}
