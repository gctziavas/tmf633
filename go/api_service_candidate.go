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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ServiceCandidateApiController binds http requests to an api service and writes the service results to the http response
type ServiceCandidateApiController struct {
	service ServiceCandidateApiServicer
	errorHandler ErrorHandler
}

// ServiceCandidateApiOption for how the controller is set up.
type ServiceCandidateApiOption func(*ServiceCandidateApiController)

// WithServiceCandidateApiErrorHandler inject ErrorHandler into controller
func WithServiceCandidateApiErrorHandler(h ErrorHandler) ServiceCandidateApiOption {
	return func(c *ServiceCandidateApiController) {
		c.errorHandler = h
	}
}

// NewServiceCandidateApiController creates a default api controller
func NewServiceCandidateApiController(s ServiceCandidateApiServicer, opts ...ServiceCandidateApiOption) Router {
	controller := &ServiceCandidateApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ServiceCandidateApiController
func (c *ServiceCandidateApiController) Routes() Routes {
	return Routes{ 
		{
			"CreateServiceCandidate",
			strings.ToUpper("Post"),
			"/tmf-api/serviceCatalogManagement/v4/serviceCandidate",
			c.CreateServiceCandidate,
		},
		{
			"DeleteServiceCandidate",
			strings.ToUpper("Delete"),
			"/tmf-api/serviceCatalogManagement/v4/serviceCandidate/{id}",
			c.DeleteServiceCandidate,
		},
		{
			"ListServiceCandidate",
			strings.ToUpper("Get"),
			"/tmf-api/serviceCatalogManagement/v4/serviceCandidate",
			c.ListServiceCandidate,
		},
		{
			"PatchServiceCandidate",
			strings.ToUpper("Patch"),
			"/tmf-api/serviceCatalogManagement/v4/serviceCandidate/{id}",
			c.PatchServiceCandidate,
		},
		{
			"RetrieveServiceCandidate",
			strings.ToUpper("Get"),
			"/tmf-api/serviceCatalogManagement/v4/serviceCandidate/{id}",
			c.RetrieveServiceCandidate,
		},
	}
}

// CreateServiceCandidate - Creates a ServiceCandidate
func (c *ServiceCandidateApiController) CreateServiceCandidate(w http.ResponseWriter, r *http.Request) {
	serviceCandidateParam := ServiceCandidateCreate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&serviceCandidateParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertServiceCandidateCreateRequired(serviceCandidateParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateServiceCandidate(r.Context(), serviceCandidateParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteServiceCandidate - Deletes a ServiceCandidate
func (c *ServiceCandidateApiController) DeleteServiceCandidate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	
	result, err := c.service.DeleteServiceCandidate(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ListServiceCandidate - List or find ServiceCandidate objects
func (c *ServiceCandidateApiController) ListServiceCandidate(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fieldsParam := query.Get("fields")
	offsetParam, err := parseInt32Parameter(query.Get("offset"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	limitParam, err := parseInt32Parameter(query.Get("limit"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.ListServiceCandidate(r.Context(), fieldsParam, offsetParam, limitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PatchServiceCandidate - Updates partially a ServiceCandidate
func (c *ServiceCandidateApiController) PatchServiceCandidate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	
	serviceCandidateParam := ServiceCandidateUpdate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&serviceCandidateParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertServiceCandidateUpdateRequired(serviceCandidateParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PatchServiceCandidate(r.Context(), idParam, serviceCandidateParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// RetrieveServiceCandidate - Retrieves a ServiceCandidate by ID
func (c *ServiceCandidateApiController) RetrieveServiceCandidate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	idParam := params["id"]
	
	fieldsParam := query.Get("fields")
	result, err := c.service.RetrieveServiceCandidate(r.Context(), idParam, fieldsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
