/*
 * Service Catalog Management
 *
 * ## TMF API Reference: TMF633 - Service Catalog Management   Version 4.0   The Service Catalog API is one of Catalog Management API Family. Service Catalog API goal is to provide a catalog of services.  Service Catalog API performs the following operations on the resource : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity (for administration purposes) - Manage notification of events. .  Copyright © TM Forum 2020. All Rights Reserved. 
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf633

// AttachmentRefOrValue - An attachment by value or by reference. An attachment complements the description of an element, for example through a document, a video, a picture.
type AttachmentRefOrValue struct {

	// Unique identifier for this particular attachment
	Id string `json:"id,omitempty"`

	// URI for this Attachment
	Href string `json:"href,omitempty"`

	// Attachment type such as video, picture
	AttachmentType string `json:"attachmentType,omitempty"`

	// The actual contents of the attachment object, if embedded, encoded as base64
	Content string `json:"content,omitempty"`

	// A narrative text describing the content of the attachment
	Description string `json:"description,omitempty"`

	// Attachment mime type such as extension file for video, picture and document
	MimeType string `json:"mimeType,omitempty"`

	// The name of the attachment
	Name string `json:"name,omitempty"`

	// Uniform Resource Locator, is a web page address (a subset of URI)
	Url string `json:"url,omitempty"`

	Size Quantity `json:"size,omitempty"`

	ValidFor TimePeriod `json:"validFor,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	Type string `json:"@type,omitempty"`

	// The actual type of the target instance when needed for disambiguation.
	ReferredType string `json:"@referredType,omitempty"`
}

// AssertAttachmentRefOrValueRequired checks if the required fields are not zero-ed
func AssertAttachmentRefOrValueRequired(obj AttachmentRefOrValue) error {
	if err := AssertQuantityRequired(obj.Size); err != nil {
		return err
	}
	if err := AssertTimePeriodRequired(obj.ValidFor); err != nil {
		return err
	}
	return nil
}

// AssertRecurseAttachmentRefOrValueRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AttachmentRefOrValue (e.g. [][]AttachmentRefOrValue), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAttachmentRefOrValueRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAttachmentRefOrValue, ok := obj.(AttachmentRefOrValue)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAttachmentRefOrValueRequired(aAttachmentRefOrValue)
	})
}