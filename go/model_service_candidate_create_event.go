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

// ServiceCandidateCreateEvent - The notification data structure
type ServiceCandidateCreateEvent struct {

	Event ServiceCandidateCreateEventPayload `json:"event,omitempty"`

	// The identifier of the notification.
	EventId string `json:"eventId,omitempty"`

	// Time of the event occurrence.
	EventTime time.Time `json:"eventTime,omitempty"`

	// The type of the notification.
	EventType string `json:"eventType,omitempty"`

	// The correlation id for this event.
	CorrelationId string `json:"correlationId,omitempty"`

	// The domain of the event.
	Domain string `json:"domain,omitempty"`

	// The title of the event.
	Title string `json:"title,omitempty"`

	// An explnatory of the event.
	Description string `json:"description,omitempty"`

	// A priority.
	Priority string `json:"priority,omitempty"`

	// The time the event occured.
	TimeOcurred time.Time `json:"timeOcurred,omitempty"`
}

// AssertServiceCandidateCreateEventRequired checks if the required fields are not zero-ed
func AssertServiceCandidateCreateEventRequired(obj ServiceCandidateCreateEvent) error {
	if err := AssertServiceCandidateCreateEventPayloadRequired(obj.Event); err != nil {
		return err
	}
	return nil
}

// AssertRecurseServiceCandidateCreateEventRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ServiceCandidateCreateEvent (e.g. [][]ServiceCandidateCreateEvent), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServiceCandidateCreateEventRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aServiceCandidateCreateEvent, ok := obj.(ServiceCandidateCreateEvent)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServiceCandidateCreateEventRequired(aServiceCandidateCreateEvent)
	})
}