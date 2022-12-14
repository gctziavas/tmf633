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

// TimePeriod - A period of time, either as a deadline (endDateTime only) a startDateTime only, or both
type TimePeriod struct {

	// End of the time period, using IETC-RFC-3339 format
	EndDateTime time.Time `json:"endDateTime,omitempty"`

	// Start of the time period, using IETC-RFC-3339 format
	StartDateTime time.Time `json:"startDateTime,omitempty"`
}

// AssertTimePeriodRequired checks if the required fields are not zero-ed
func AssertTimePeriodRequired(obj TimePeriod) error {
	return nil
}

// AssertRecurseTimePeriodRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TimePeriod (e.g. [][]TimePeriod), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTimePeriodRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTimePeriod, ok := obj.(TimePeriod)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTimePeriodRequired(aTimePeriod)
	})
}
