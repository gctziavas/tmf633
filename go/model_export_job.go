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

// ExportJob - Represents a task used to export resources to a file
type ExportJob struct {

	// Identifier of the export job
	Id string `json:"id,omitempty"`

	// Reference of the export job
	Href string `json:"href,omitempty"`

	// Data at which the job was completed
	CompletionDate time.Time `json:"completionDate,omitempty"`

	// The format of the exported data
	ContentType string `json:"contentType,omitempty"`

	// Date at which the job was created
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Reason for failure
	ErrorLog string `json:"errorLog,omitempty"`

	// URL of the root resource acting as the source for streaming content to the file specified by the export job
	Path string `json:"path,omitempty"`

	// Used to scope the exported data
	Query string `json:"query,omitempty"`

	// URL of the file containing the data to be exported
	Url string `json:"url,omitempty"`

	Status JobStateType `json:"status,omitempty"`
}

// AssertExportJobRequired checks if the required fields are not zero-ed
func AssertExportJobRequired(obj ExportJob) error {
	return nil
}

// AssertRecurseExportJobRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ExportJob (e.g. [][]ExportJob), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseExportJobRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aExportJob, ok := obj.(ExportJob)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertExportJobRequired(aExportJob)
	})
}
