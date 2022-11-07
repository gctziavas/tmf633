/*
 * Service Catalog Management
 *
 * ## TMF API Reference: TMF633 - Service Catalog Management   Version 4.0   The Service Catalog API is one of Catalog Management API Family. Service Catalog API goal is to provide a catalog of services.  Service Catalog API performs the following operations on the resource : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity (for administration purposes) - Manage notification of events. .  Copyright © TM Forum 2020. All Rights Reserved. 
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf633

// TargetEntitySchema - The reference object to the schema and type of target entity which is described by a specification
type TargetEntitySchema struct {

	// This field provides a link to the schema describing the target entity
	SchemaLocation string `json:"@schemaLocation"`

	// Class type of the target entity
	Type string `json:"@type"`
}

// AssertTargetEntitySchemaRequired checks if the required fields are not zero-ed
func AssertTargetEntitySchemaRequired(obj TargetEntitySchema) error {
	elements := map[string]interface{}{
		"@schemaLocation": obj.SchemaLocation,
		"@type": obj.Type,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseTargetEntitySchemaRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TargetEntitySchema (e.g. [][]TargetEntitySchema), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTargetEntitySchemaRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTargetEntitySchema, ok := obj.(TargetEntitySchema)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTargetEntitySchemaRequired(aTargetEntitySchema)
	})
}