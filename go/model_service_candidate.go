/*
 * Service Catalog Management
 *
 * ## TMF API Reference: TMF633 - Service Catalog Management   Version 4.0   The Service Catalog API is one of Catalog Management API Family. Service Catalog API goal is to provide a catalog of services.  Service Catalog API performs the following operations on the resource : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity (for administration purposes) - Manage notification of events. .  Copyright © TM Forum 2020. All Rights Reserved. 
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf633

import (
	"time"
)

// ServiceCandidate - ServiceCandidate is an entity that makes a service specification available to a catalog. A ServiceCandidate and its associated service specification may be published - made visible - in any number of service catalogs, or in none. One service specification can be composed of other service specifications.
type ServiceCandidate struct {

	// unique identifier
	Id string `json:"id,omitempty"`

	// Hyperlink reference
	Href string `json:"href,omitempty"`

	// Description of this REST resource
	Description string `json:"description,omitempty"`

	// Date and time of the last update of this REST resource
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	// Used to indicate the current lifecycle status of the service candidate.
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`

	// Name given to this REST resource
	Name string `json:"name,omitempty"`

	// the version of service candidate
	Version string `json:"version,omitempty"`

	// List of categories for this candidate
	Category []ServiceCategoryRef `json:"category,omitempty"`

	ServiceSpecification ServiceSpecificationRef `json:"serviceSpecification,omitempty"`

	ValidFor TimePeriod `json:"validFor,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type,omitempty"`
}

// AssertServiceCandidateRequired checks if the required fields are not zero-ed
func AssertServiceCandidateRequired(obj ServiceCandidate) error {
	for _, el := range obj.Category {
		if err := AssertServiceCategoryRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertServiceSpecificationRefRequired(obj.ServiceSpecification); err != nil {
		return err
	}
	if err := AssertTimePeriodRequired(obj.ValidFor); err != nil {
		return err
	}
	return nil
}

// AssertRecurseServiceCandidateRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ServiceCandidate (e.g. [][]ServiceCandidate), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServiceCandidateRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aServiceCandidate, ok := obj.(ServiceCandidate)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServiceCandidateRequired(aServiceCandidate)
	})
}
