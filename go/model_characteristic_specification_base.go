/*
 * Service Catalog Management
 *
 * ## TMF API Reference: TMF633 - Service Catalog Management   Version 4.0   The Service Catalog API is one of Catalog Management API Family. Service Catalog API goal is to provide a catalog of services.  Service Catalog API performs the following operations on the resource : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity (for administration purposes) - Manage notification of events. .  Copyright © TM Forum 2020. All Rights Reserved. 
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf633

// CharacteristicSpecificationBase - This class defines a characteristic specification.
type CharacteristicSpecificationBase struct {

	// Unique ID for the characteristic
	Id string `json:"id,omitempty"`

	// If true, the Boolean indicates that the target Characteristic is configurable
	Configurable bool `json:"configurable,omitempty"`

	// A narrative that explains the CharacteristicSpecification.
	Description string `json:"description,omitempty"`

	// An indicator that specifies that the values for the characteristic can be extended by adding new values when instantiating a characteristic for a resource.
	Extensible bool `json:"extensible,omitempty"`

	// An indicator that specifies if a value is unique for the specification. Possible values are; \"unique while value is in effect\" and \"unique whether value is in effect or not\"
	IsUnique bool `json:"isUnique,omitempty"`

	// The maximum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where five is the value for the maxCardinality.
	MaxCardinality int32 `json:"maxCardinality,omitempty"`

	// The minimum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where zero is the value for the minCardinality.
	MinCardinality int32 `json:"minCardinality,omitempty"`

	// A word, term, or phrase by which this characteristic specification is known and distinguished from other characteristic specifications.
	Name string `json:"name,omitempty"`

	// A rule or principle represented in regular expression used to derive the value of a characteristic value.
	Regex string `json:"regex,omitempty"`

	// A kind of value that the characteristic can take on, such as numeric, text and so forth
	ValueType string `json:"valueType,omitempty"`

	ValidFor TimePeriod `json:"validFor,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type,omitempty"`

	// This (optional) field provides a link to the schema describing the value type.
	ValueSchemaLocation string `json:"@valueSchemaLocation,omitempty"`
}

// AssertCharacteristicSpecificationBaseRequired checks if the required fields are not zero-ed
func AssertCharacteristicSpecificationBaseRequired(obj CharacteristicSpecificationBase) error {
	if err := AssertTimePeriodRequired(obj.ValidFor); err != nil {
		return err
	}
	return nil
}

// AssertRecurseCharacteristicSpecificationBaseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of CharacteristicSpecificationBase (e.g. [][]CharacteristicSpecificationBase), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCharacteristicSpecificationBaseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aCharacteristicSpecificationBase, ok := obj.(CharacteristicSpecificationBase)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCharacteristicSpecificationBaseRequired(aCharacteristicSpecificationBase)
	})
}
