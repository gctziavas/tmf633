/*
 * Service Catalog Management
 *
 * ## TMF API Reference: TMF633 - Service Catalog Management   Version 4.0   The Service Catalog API is one of Catalog Management API Family. Service Catalog API goal is to provide a catalog of services.  Service Catalog API performs the following operations on the resource : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity (for administration purposes) - Manage notification of events. .  Copyright © TM Forum 2020. All Rights Reserved. 
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf633

// ServiceCandidateDeleteEventPayload - The event data structure
type ServiceCandidateDeleteEventPayload struct {

	ServiceCandidate ServiceCandidate `json:"serviceCandidate,omitempty"`
}

// AssertServiceCandidateDeleteEventPayloadRequired checks if the required fields are not zero-ed
func AssertServiceCandidateDeleteEventPayloadRequired(obj ServiceCandidateDeleteEventPayload) error {
	if err := AssertServiceCandidateRequired(obj.ServiceCandidate); err != nil {
		return err
	}
	return nil
}

// AssertRecurseServiceCandidateDeleteEventPayloadRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ServiceCandidateDeleteEventPayload (e.g. [][]ServiceCandidateDeleteEventPayload), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServiceCandidateDeleteEventPayloadRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aServiceCandidateDeleteEventPayload, ok := obj.(ServiceCandidateDeleteEventPayload)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServiceCandidateDeleteEventPayloadRequired(aServiceCandidateDeleteEventPayload)
	})
}
