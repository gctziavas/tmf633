/*
 * Service Catalog Management
 *
 * ## TMF API Reference: TMF633 - Service Catalog Management   Version 4.0   The Service Catalog API is one of Catalog Management API Family. Service Catalog API goal is to provide a catalog of services.  Service Catalog API performs the following operations on the resource : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity (for administration purposes) - Manage notification of events. .  Copyright © TM Forum 2020. All Rights Reserved. 
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf633

// ServiceLevelSpecificationRef - A Service Level Specification represents a pre-defined or negotiated set of Service Level  Objectives. In addition, certain consequences are associated with not meeting the Service Level  Objectives. Service Level Agreements are expressed in terms of Service Level Specifications.
type ServiceLevelSpecificationRef struct {

	// unique identifier
	Id string `json:"id"`

	// Hyperlink reference
	Href string `json:"href,omitempty"`

	// Name of the related entity.
	Name string `json:"name,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type,omitempty"`

	// The actual type of the target instance when needed for disambiguation.
	ReferredType string `json:"@referredType,omitempty"`
}

// AssertServiceLevelSpecificationRefRequired checks if the required fields are not zero-ed
func AssertServiceLevelSpecificationRefRequired(obj ServiceLevelSpecificationRef) error {
	elements := map[string]interface{}{
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseServiceLevelSpecificationRefRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ServiceLevelSpecificationRef (e.g. [][]ServiceLevelSpecificationRef), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServiceLevelSpecificationRefRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aServiceLevelSpecificationRef, ok := obj.(ServiceLevelSpecificationRef)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServiceLevelSpecificationRefRequired(aServiceLevelSpecificationRef)
	})
}
