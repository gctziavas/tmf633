/*
 * Service Catalog Management
 *
 * ## TMF API Reference: TMF633 - Service Catalog Management   Version 4.0   The Service Catalog API is one of Catalog Management API Family. Service Catalog API goal is to provide a catalog of services.  Service Catalog API performs the following operations on the resource : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity (for administration purposes) - Manage notification of events. .  Copyright © TM Forum 2020. All Rights Reserved. 
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf633

// ServiceSpecificationUpdate - ServiceSpecification is a class that offers characteristics to describe a type of service. Functionally, it acts as a template by which Services may be instantiated. By sharing the same specification, these services would therefore share the same set of characteristics. Skipped properties: id,href,lastUpdate
type ServiceSpecificationUpdate struct {

	// Description of the specification
	Description string `json:"description,omitempty"`

	// isBundle determines whether specification represents a single specification (false), or a bundle of specifications (true).
	IsBundle bool `json:"isBundle,omitempty"`

	// Used to indicate the current lifecycle status of this catalog item
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`

	// Name given to the specification
	Name string `json:"name,omitempty"`

	// specification version
	Version string `json:"version,omitempty"`

	// Attachments that may be of relevance to this specification, such as picture, document, media
	Attachment []AttachmentRefOrValue `json:"attachment,omitempty"`

	// This is a list of constraint references applied to this specification
	Constraint []ConstraintRef `json:"constraint,omitempty"`

	// Relationship to another specification
	EntitySpecRelationship []EntitySpecificationRelationship `json:"entitySpecRelationship,omitempty"`

	// A list of Features for this specification.
	FeatureSpecification []FeatureSpecification `json:"featureSpecification,omitempty"`

	// Parties who manage or otherwise have an interest in this specification
	RelatedParty []RelatedParty `json:"relatedParty,omitempty"`

	// A list of resource specification references (ResourceSpecificationRef [*]). The ResourceSpecification is required for a service specification with type ResourceFacingServiceSpecification (RFSS).
	ResourceSpecification []ResourceSpecificationRef `json:"resourceSpecification,omitempty"`

	// A list of service level specifications related to this service specification, and which will need to be satisifiable for corresponding service instances; e.g. Gold, Platinum
	ServiceLevelSpecification []ServiceLevelSpecificationRef `json:"serviceLevelSpecification,omitempty"`

	// A list of service specifications related to this specification, e.g. migration, substitution, dependency or exclusivity relationship
	ServiceSpecRelationship []ServiceSpecRelationship `json:"serviceSpecRelationship,omitempty"`

	// List of characteristics that the entity can take
	SpecCharacteristic []CharacteristicSpecification `json:"specCharacteristic,omitempty"`

	TargetEntitySchema TargetEntitySchema `json:"targetEntitySchema,omitempty"`

	ValidFor TimePeriod `json:"validFor,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type,omitempty"`
}

// AssertServiceSpecificationUpdateRequired checks if the required fields are not zero-ed
func AssertServiceSpecificationUpdateRequired(obj ServiceSpecificationUpdate) error {
	for _, el := range obj.Attachment {
		if err := AssertAttachmentRefOrValueRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Constraint {
		if err := AssertConstraintRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EntitySpecRelationship {
		if err := AssertEntitySpecificationRelationshipRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.FeatureSpecification {
		if err := AssertFeatureSpecificationRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RelatedParty {
		if err := AssertRelatedPartyRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ResourceSpecification {
		if err := AssertResourceSpecificationRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ServiceLevelSpecification {
		if err := AssertServiceLevelSpecificationRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ServiceSpecRelationship {
		if err := AssertServiceSpecRelationshipRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.SpecCharacteristic {
		if err := AssertCharacteristicSpecificationRequired(el); err != nil {
			return err
		}
	}
	if err := AssertTargetEntitySchemaRequired(obj.TargetEntitySchema); err != nil {
		return err
	}
	if err := AssertTimePeriodRequired(obj.ValidFor); err != nil {
		return err
	}
	return nil
}

// AssertRecurseServiceSpecificationUpdateRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ServiceSpecificationUpdate (e.g. [][]ServiceSpecificationUpdate), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServiceSpecificationUpdateRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aServiceSpecificationUpdate, ok := obj.(ServiceSpecificationUpdate)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServiceSpecificationUpdateRequired(aServiceSpecificationUpdate)
	})
}
