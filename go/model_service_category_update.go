/*
 * Service Catalog Management
 *
 * ## TMF API Reference: TMF633 - Service Catalog Management   Version 4.0   The Service Catalog API is one of Catalog Management API Family. Service Catalog API goal is to provide a catalog of services.  Service Catalog API performs the following operations on the resource : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity (for administration purposes) - Manage notification of events. .  Copyright © TM Forum 2020. All Rights Reserved. 
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf633

// ServiceCategoryUpdate - The (service) category resource is used to group service candidates in logical containers. Categories can contain other categories. Skipped properties: id,href,lastUpdate
type ServiceCategoryUpdate struct {

	// Description of the category
	Description string `json:"description,omitempty"`

	// If true, this Boolean indicates that the category is a root of categories
	IsRoot bool `json:"isRoot,omitempty"`

	// Used to indicate the current lifecycle status
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`

	// Name of the category
	Name string `json:"name,omitempty"`

	// Unique identifier of the parent category
	ParentId string `json:"parentId,omitempty"`

	// ServiceCategory version
	Version string `json:"version,omitempty"`

	// List of child categories in the tree for in this category
	Category []ServiceCategoryRef `json:"category,omitempty"`

	// List of service candidates associated with this category
	ServiceCandidate []ServiceCandidateRef `json:"serviceCandidate,omitempty"`

	ValidFor TimePeriod `json:"validFor,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type,omitempty"`
}

// AssertServiceCategoryUpdateRequired checks if the required fields are not zero-ed
func AssertServiceCategoryUpdateRequired(obj ServiceCategoryUpdate) error {
	for _, el := range obj.Category {
		if err := AssertServiceCategoryRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ServiceCandidate {
		if err := AssertServiceCandidateRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertTimePeriodRequired(obj.ValidFor); err != nil {
		return err
	}
	return nil
}

// AssertRecurseServiceCategoryUpdateRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ServiceCategoryUpdate (e.g. [][]ServiceCategoryUpdate), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServiceCategoryUpdateRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aServiceCategoryUpdate, ok := obj.(ServiceCategoryUpdate)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServiceCategoryUpdateRequired(aServiceCategoryUpdate)
	})
}
